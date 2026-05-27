package loadbalancer2

import (
	scpsdk "github.com/SamsungSDSCloud/terraform-sdk-samsungcloudplatform/v3/client"
)

type APIClient struct {
	cfg    *scpsdk.Configuration
	common service
	LbLogStorageOpenApiControllerApi *LbLogStorageOpenApiControllerApiService
	LbProfileOpenApiControllerApi *LbProfileOpenApiControllerApiService
	LbServerGroupOpenApiControllerApi *LbServerGroupOpenApiControllerApiService
	LbServiceOpenApiControllerApi *LbServiceOpenApiControllerApiService
	LoadBalancerOpenApiControllerApi *LoadBalancerOpenApiControllerApiService
	LoadBalancerOpenApiV3ControllerApi *LoadBalancerOpenApiV3ControllerApiService
}

type service struct {
	client *APIClient
}

type LbLogStorageOpenApiControllerApiService service
type LbProfileOpenApiControllerApiService service
type LbServerGroupOpenApiControllerApiService service
type LbServiceOpenApiControllerApiService service
type LoadBalancerOpenApiControllerApiService service
type LoadBalancerOpenApiV3ControllerApiService service

func NewAPIClient(cfg *scpsdk.Configuration) *APIClient {
	c := &APIClient{cfg: cfg}
	c.common.client = c
	c.LbLogStorageOpenApiControllerApi = (*LbLogStorageOpenApiControllerApiService)(&c.common)
	c.LbProfileOpenApiControllerApi = (*LbProfileOpenApiControllerApiService)(&c.common)
	c.LbServerGroupOpenApiControllerApi = (*LbServerGroupOpenApiControllerApiService)(&c.common)
	c.LbServiceOpenApiControllerApi = (*LbServiceOpenApiControllerApiService)(&c.common)
	c.LoadBalancerOpenApiControllerApi = (*LoadBalancerOpenApiControllerApiService)(&c.common)
	c.LoadBalancerOpenApiV3ControllerApi = (*LoadBalancerOpenApiV3ControllerApiService)(&c.common)
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
