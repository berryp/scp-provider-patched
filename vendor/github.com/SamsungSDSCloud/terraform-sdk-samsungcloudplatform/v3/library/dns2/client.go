package dns2

import (
	scpsdk "github.com/SamsungSDSCloud/terraform-sdk-samsungcloudplatform/v3/client"
)

type APIClient struct {
	cfg    *scpsdk.Configuration
	common service
	DnsOpenApiV2ControllerApi *DnsOpenApiV2ControllerApiService
	DnsOpenApiV3ControllerApi *DnsOpenApiV3ControllerApiService
}

type service struct {
	client *APIClient
}

type DnsOpenApiV2ControllerApiService service
type DnsOpenApiV3ControllerApiService service

func NewAPIClient(cfg *scpsdk.Configuration) *APIClient {
	c := &APIClient{cfg: cfg}
	c.common.client = c
	c.DnsOpenApiV2ControllerApi = (*DnsOpenApiV2ControllerApiService)(&c.common)
	c.DnsOpenApiV3ControllerApi = (*DnsOpenApiV3ControllerApiService)(&c.common)
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
