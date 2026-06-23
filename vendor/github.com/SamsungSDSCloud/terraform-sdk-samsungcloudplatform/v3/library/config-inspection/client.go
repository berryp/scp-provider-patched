package configinspection

import (
	scpsdk "github.com/SamsungSDSCloud/terraform-sdk-samsungcloudplatform/v3/client"
)

type APIClient struct {
	cfg    *scpsdk.Configuration
	common service
	ConfigInspectionDashboardOpenApiControllerApi *ConfigInspectionDashboardOpenApiControllerApiService
	ConfigInspectionOpenApiControllerApi *ConfigInspectionOpenApiControllerApiService
	Client *Client
}

type service struct {
	client *APIClient
}

type ConfigInspectionDashboardOpenApiControllerApiService service
type ConfigInspectionOpenApiControllerApiService service
type Client service

func NewAPIClient(cfg *scpsdk.Configuration) *APIClient {
	c := &APIClient{cfg: cfg}
	c.common.client = c
	c.ConfigInspectionDashboardOpenApiControllerApi = (*ConfigInspectionDashboardOpenApiControllerApiService)(&c.common)
	c.ConfigInspectionOpenApiControllerApi = (*ConfigInspectionOpenApiControllerApiService)(&c.common)
	c.Client = (*Client)(&c.common)
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
