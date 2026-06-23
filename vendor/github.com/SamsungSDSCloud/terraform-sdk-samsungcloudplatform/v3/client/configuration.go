package client

import "net/http"

// Configuration mirrors the recovered shape of
// terraform-sdk-samsungcloudplatform/v3/client.Configuration from DWARF.
type Configuration struct {
	BasePath      string            `json:"basePath,omitempty"`
	Host          string            `json:"host,omitempty"`
	Scheme        string            `json:"scheme,omitempty"`
	DefaultHeader map[string]string `json:"defaultHeader,omitempty"`
	UserAgent     string            `json:"userAgent,omitempty"`
	HTTPClient    *http.Client
	ProjectId     string
	UserId        string
	Email         string
	LoginId       string
	AuthMethod    string
	Credentials   *Credentials
	Token         string
}

// Credentials carries the access-key auth pair used when AuthMethod == "access-key".
// Recovered from DWARF.
type Credentials struct {
	AccessKey string
	SecretKey string
}
