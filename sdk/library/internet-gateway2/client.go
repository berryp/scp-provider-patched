package internetgateway2

import (
	scpsdk "github.com/SamsungSDSCloud/terraform-sdk-samsungcloudplatform/v3/client"
)

type APIClient struct {
	cfg    *scpsdk.Configuration
	common service
	InternetGatewayV2ControllerV2Api *InternetGatewayV2ControllerV2ApiService
	InternetGatewayV3ControllerV3Api *InternetGatewayV3ControllerV3ApiService
	InternetGatewayV4ControllerApi *InternetGatewayV4ControllerApiService
}

type service struct {
	client *APIClient
}

type InternetGatewayV2ControllerV2ApiService service
type InternetGatewayV3ControllerV3ApiService service
type InternetGatewayV4ControllerApiService service

func NewAPIClient(cfg *scpsdk.Configuration) *APIClient {
	c := &APIClient{cfg: cfg}
	c.common.client = c
	c.InternetGatewayV2ControllerV2Api = (*InternetGatewayV2ControllerV2ApiService)(&c.common)
	c.InternetGatewayV3ControllerV3Api = (*InternetGatewayV3ControllerV3ApiService)(&c.common)
	c.InternetGatewayV4ControllerApi = (*InternetGatewayV4ControllerApiService)(&c.common)
	return c
}

func (c *APIClient) ChangeBasePath(p string) { c.cfg.BasePath = p }

func (c *APIClient) ToTagRequestList(tags map[string]interface{}) []TagRequest {
	if len(tags) == 0 {
		return nil
	}
	out := make([]TagRequest, 0, len(tags))
	for k, v := range tags {
		t := TagRequest{TagKey: k}
		if s, ok := v.(string); ok {
			t.TagValue = s
		}
		out = append(out, t)
	}
	return out
}
