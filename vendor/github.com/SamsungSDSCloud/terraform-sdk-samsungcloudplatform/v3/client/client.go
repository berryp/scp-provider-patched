package client

import (
	"bytes"
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"reflect"
	"strings"
	"time"
)

// ParameterToString formats a value the way the OAG-Go template does for
// query/form parameters: slices become comma-separated (with brackets
// trimmed), nil and empty slices become "" (so callers can skip Set), and
// scalars fall through fmt.Sprintf("%v", ...).
//
// Used by every generated *Opts encoder so the wire format matches what the
// SCP gateway expects.
func ParameterToString(obj interface{}) string {
	if obj == nil {
		return ""
	}
	val := reflect.ValueOf(obj)
	switch val.Kind() {
	case reflect.Slice, reflect.Array:
		if val.Len() == 0 {
			return ""
		}
		s := strings.ReplaceAll(fmt.Sprint(obj), " ", ",")
		return strings.Trim(s, "[]")
	case reflect.Ptr:
		if val.IsNil() {
			return ""
		}
		return ParameterToString(val.Elem().Interface())
	}
	return fmt.Sprintf("%v", obj)
}

// APIResponse wraps the raw http.Response and is what every per-service API
// method returns in addition to the decoded payload and error.
type APIResponse struct {
	*http.Response
	Message   string
	Operation string
}

// GenericOpenAPIError is the OpenAPI-Generator-standard error type. Per-service
// methods return *this* when an HTTP error or decode error occurs; callers can
// type-assert and inspect the body.
type GenericOpenAPIError struct {
	body  []byte
	error string
	model interface{}
}

func (e GenericOpenAPIError) Error() string      { return e.error }
func (e GenericOpenAPIError) Body() []byte       { return e.body }
func (e GenericOpenAPIError) Model() interface{} { return e.model }

// AddDefaultHeader installs a header that every subsequent request will carry.
// Mirrors the recovered (*Configuration).AddDefaultHeader behavior.
func (c *Configuration) AddDefaultHeader(key, value string) {
	if c.DefaultHeader == nil {
		c.DefaultHeader = make(map[string]string)
	}
	c.DefaultHeader[key] = value
}

// MakeHmacSignature is the recovered SCP signer. Validated to byte-equivalence
// against a captured real request (see harness/signer_probe).
//
//	HMAC-SHA256(secretKey, method + fullURL + timestampMillis + accessKey + projectId + "OpenApi")
//
// fullURL is the request's full URL: scheme + "://" + host[":"+port] + path[?query].
// timestampMillis is the Unix time in milliseconds as a decimal string.
// Output is base64-encoded with standard padding.
func MakeHmacSignature(method, fullURL, timestampMillis, accessKey, secretKey, projectId, clientType string) string {
	canonical := method + fullURL + timestampMillis + accessKey + projectId + clientType
	mac := hmac.New(sha256.New, []byte(secretKey))
	mac.Write([]byte(canonical))
	return base64.StdEncoding.EncodeToString(mac.Sum(nil))
}

// SetupRequestHeader installs SCP-style headers on req. Mirrors the recovered
// (*Configuration).SetupRequestHeader behavior: deletes "X-Cmp-Client" then
// repopulates ~10 X-Cmp-* headers, dispatching to MakeHmacSignature for
// "access-key" auth and bearer-token injection for "id-token".
//
// Header names captured from the wire (recon/traffic/real/0001_*.json):
//
//	X-Cmp-AccessKey, X-Cmp-ClientType, X-Cmp-CompanyId, X-Cmp-Language,
//	X-Cmp-LoginId, X-Cmp-ProjectId, X-Cmp-UserEmail, X-Cmp-UserId,
//	X-Cmp-Signature, X-Cmp-Timestamp
func (c *Configuration) SetupRequestHeader(method, _ignored string, req *http.Request) error {
	req.Header.Del("X-Cmp-Client")

	switch c.AuthMethod {
	case "access-key":
		if c.Credentials == nil {
			return errors.New("scpsdk: access-key auth requires non-nil Credentials")
		}
		ts := fmt.Sprintf("%d", time.Now().UnixMilli())
		fullURL := req.URL.String()
		sig := MakeHmacSignature(
			method, fullURL, ts,
			c.Credentials.AccessKey, c.Credentials.SecretKey,
			c.ProjectId, "OpenApi",
		)
		req.Header.Set("X-Cmp-AccessKey", c.Credentials.AccessKey)
		req.Header.Set("X-Cmp-ClientType", "OpenApi")
		req.Header.Set("X-Cmp-CompanyId", "")
		req.Header.Set("X-Cmp-Language", "en-US")
		req.Header.Set("X-Cmp-LoginId", c.LoginId)
		req.Header.Set("X-Cmp-ProjectId", c.ProjectId)
		req.Header.Set("X-Cmp-UserEmail", c.Email)
		req.Header.Set("X-Cmp-UserId", c.UserId)
		req.Header.Set("X-Cmp-Signature", sig)
		req.Header.Set("X-Cmp-Timestamp", ts)
	case "id-token", "":
		if c.Token != "" {
			req.Header.Set("Authorization", "Bearer "+c.Token)
		}
	default:
		return fmt.Errorf("scpsdk: unsupported AuthMethod %q", c.AuthMethod)
	}
	return nil
}

// CallAPI is the central transport entry point shared by every per-service
// generated method. Shape mirrors the openapi-generator-cli "go" template v6.x.
// Each library/<service> package's per-method generated code calls into this.
func (c *Configuration) CallAPI(
	ctx context.Context,
	relativePath string,
	method string,
	postBody interface{},
	headerParams map[string]string,
	queryParams url.Values,
	formParams url.Values,
) (*http.Response, error) {
	if c.HTTPClient == nil {
		c.HTTPClient = http.DefaultClient
	}

	fullURL := strings.TrimRight(c.BasePath, "/") + "/" + strings.TrimLeft(relativePath, "/")
	if len(queryParams) > 0 {
		fullURL += "?" + queryParams.Encode()
	}

	var bodyReader io.Reader
	if postBody != nil {
		buf, err := json.Marshal(postBody)
		if err != nil {
			return nil, fmt.Errorf("scpsdk: marshal request body: %w", err)
		}
		bodyReader = bytes.NewReader(buf)
	}

	req, err := http.NewRequestWithContext(ctx, method, fullURL, bodyReader)
	if err != nil {
		return nil, err
	}

	for k, v := range c.DefaultHeader {
		req.Header.Set(k, v)
	}
	for k, v := range headerParams {
		req.Header.Set(k, v)
	}
	if req.Header.Get("Content-Type") == "" && postBody != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	if req.Header.Get("Accept") == "" {
		req.Header.Set("Accept", "application/json")
	}
	if c.UserAgent != "" {
		req.Header.Set("User-Agent", c.UserAgent)
	}

	if err := c.SetupRequestHeader(method, relativePath, req); err != nil {
		return nil, err
	}

	// Debug log every request when SCPSDK_TRACE is set (any non-empty value).
	if traceURLs {
		fmt.Fprintf(os.Stderr, "[scpsdk] %s %s\n", method, req.URL.String())
	}
	resp, err := c.HTTPClient.Do(req)
	if traceURLs && resp != nil {
		fmt.Fprintf(os.Stderr, "[scpsdk]   <- %d %s\n", resp.StatusCode, req.URL.String())
	}
	return resp, err
}

// traceURLs is set once at package init from $SCPSDK_TRACE.
var traceURLs = os.Getenv("SCPSDK_TRACE") != ""

// DecodeResponse helper used by generated per-service methods. Reads the
// response body, returns GenericOpenAPIError on non-2xx, decodes JSON into
// `out` on 2xx (skipping if out is nil or body is empty).
func DecodeResponse(resp *http.Response, out interface{}) error {
	if resp == nil {
		return errors.New("scpsdk: nil response")
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	if resp.StatusCode >= 400 {
		// Match the OAG-Go template's GenericSwaggerError format so upstream
		// provider helpers like common.IsDeleted (which checks prefix "404")
		// work as designed. Body is preserved on the struct for callers that
		// want it via .Body().
		return GenericOpenAPIError{body: body, error: resp.Status}
	}
	if out == nil || len(body) == 0 {
		return nil
	}
	return json.Unmarshal(body, out)
}

func truncate(b []byte, n int) string {
	if len(b) <= n {
		return string(b)
	}
	return string(b[:n]) + "...(truncated)"
}
