package vpc2

import (
	scpsdk "github.com/SamsungSDSCloud/terraform-sdk-samsungcloudplatform/v3/client"
)

type APIClient struct {
	cfg    *scpsdk.Configuration
	common service
	VpcOpenApiControllerApi *VpcOpenApiControllerApiService
	VpcOpenApiV3ControllerApi *VpcOpenApiV3ControllerApiService
}

type service struct {
	client *APIClient
}

type VpcOpenApiControllerApiService service
type VpcOpenApiV3ControllerApiService service

func NewAPIClient(cfg *scpsdk.Configuration) *APIClient {
	c := &APIClient{cfg: cfg}
	c.common.client = c
	c.VpcOpenApiControllerApi = (*VpcOpenApiControllerApiService)(&c.common)
	c.VpcOpenApiV3ControllerApi = (*VpcOpenApiV3ControllerApiService)(&c.common)
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
