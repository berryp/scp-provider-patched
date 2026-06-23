package gslb2

import (
	scpsdk "github.com/SamsungSDSCloud/terraform-sdk-samsungcloudplatform/v3/client"
)

type APIClient struct {
	cfg    *scpsdk.Configuration
	common service
	GslbOpenApiV2ControllerApi *GslbOpenApiV2ControllerApiService
	GslbOpenApiV3ControllerApi *GslbOpenApiV3ControllerApiService
}

type service struct {
	client *APIClient
}

type GslbOpenApiV2ControllerApiService service
type GslbOpenApiV3ControllerApiService service

func NewAPIClient(cfg *scpsdk.Configuration) *APIClient {
	c := &APIClient{cfg: cfg}
	c.common.client = c
	c.GslbOpenApiV2ControllerApi = (*GslbOpenApiV2ControllerApiService)(&c.common)
	c.GslbOpenApiV3ControllerApi = (*GslbOpenApiV3ControllerApiService)(&c.common)
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
