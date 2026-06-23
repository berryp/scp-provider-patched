package baremetalservervdc

import (
	scpsdk "github.com/SamsungSDSCloud/terraform-sdk-samsungcloudplatform/v3/client"
)

type APIClient struct {
	cfg    *scpsdk.Configuration
	common service
	VxLanBareMetalServerCreateDeleteOpenApiControllerApi *VxLanBareMetalServerCreateDeleteOpenApiControllerApiService
	VxLanBareMetalServerCreateDeleteOpenApiV2ControllerApi *VxLanBareMetalServerCreateDeleteOpenApiV2ControllerApiService
	VxLanBareMetalServerCreateDeleteOpenApiV3ControllerApi *VxLanBareMetalServerCreateDeleteOpenApiV3ControllerApiService
	VxLanBareMetalServerLongRunningTaskOpenApiControllerApi *VxLanBareMetalServerLongRunningTaskOpenApiControllerApiService
	VxLanBareMetalServerSimpleTaskOpenApiControllerApi *VxLanBareMetalServerSimpleTaskOpenApiControllerApiService
}

type service struct {
	client *APIClient
}

type VxLanBareMetalServerCreateDeleteOpenApiControllerApiService service
type VxLanBareMetalServerCreateDeleteOpenApiV2ControllerApiService service
type VxLanBareMetalServerCreateDeleteOpenApiV3ControllerApiService service
type VxLanBareMetalServerLongRunningTaskOpenApiControllerApiService service
type VxLanBareMetalServerSimpleTaskOpenApiControllerApiService service

func NewAPIClient(cfg *scpsdk.Configuration) *APIClient {
	c := &APIClient{cfg: cfg}
	c.common.client = c
	c.VxLanBareMetalServerCreateDeleteOpenApiControllerApi = (*VxLanBareMetalServerCreateDeleteOpenApiControllerApiService)(&c.common)
	c.VxLanBareMetalServerCreateDeleteOpenApiV2ControllerApi = (*VxLanBareMetalServerCreateDeleteOpenApiV2ControllerApiService)(&c.common)
	c.VxLanBareMetalServerCreateDeleteOpenApiV3ControllerApi = (*VxLanBareMetalServerCreateDeleteOpenApiV3ControllerApiService)(&c.common)
	c.VxLanBareMetalServerLongRunningTaskOpenApiControllerApi = (*VxLanBareMetalServerLongRunningTaskOpenApiControllerApiService)(&c.common)
	c.VxLanBareMetalServerSimpleTaskOpenApiControllerApi = (*VxLanBareMetalServerSimpleTaskOpenApiControllerApiService)(&c.common)
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
