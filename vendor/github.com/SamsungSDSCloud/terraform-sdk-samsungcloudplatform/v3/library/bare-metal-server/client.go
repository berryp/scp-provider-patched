package baremetalserver

import (
	scpsdk "github.com/SamsungSDSCloud/terraform-sdk-samsungcloudplatform/v3/client"
)

type APIClient struct {
	cfg    *scpsdk.Configuration
	common service
	BareMetalServerCreateDeleteOpenApiControllerApi *BareMetalServerCreateDeleteOpenApiControllerApiService
	BareMetalServerCreateDeleteOpenApiV2ControllerApi *BareMetalServerCreateDeleteOpenApiV2ControllerApiService
	BareMetalServerCreateDeleteOpenApiV3ControllerApi *BareMetalServerCreateDeleteOpenApiV3ControllerApiService
	BareMetalServerLocalSubnetOpenApiControllerApi *BareMetalServerLocalSubnetOpenApiControllerApiService
	BareMetalServerLocalSubnetOpenApiV2ControllerApi *BareMetalServerLocalSubnetOpenApiV2ControllerApiService
	BareMetalServerLongRunningTaskOpenApiControllerApi *BareMetalServerLongRunningTaskOpenApiControllerApiService
	BareMetalServerLongRunningTaskOpenApiV2ControllerApi *BareMetalServerLongRunningTaskOpenApiV2ControllerApiService
	BareMetalServerSimpleTaskOpenApiControllerApi *BareMetalServerSimpleTaskOpenApiControllerApiService
	BareMetalServerSimpleTaskOpenApiV3ControllerApi *BareMetalServerSimpleTaskOpenApiV3ControllerApiService
	BareMetalServerStaticNatOpenApiControllerApi *BareMetalServerStaticNatOpenApiControllerApiService
	BareMetalServerStaticNatOpenApiV2ControllerApi *BareMetalServerStaticNatOpenApiV2ControllerApiService
	ImOpenApiControllerApi *ImOpenApiControllerApiService
}

type service struct {
	client *APIClient
}

type BareMetalServerCreateDeleteOpenApiControllerApiService service
type BareMetalServerCreateDeleteOpenApiV2ControllerApiService service
type BareMetalServerCreateDeleteOpenApiV3ControllerApiService service
type BareMetalServerLocalSubnetOpenApiControllerApiService service
type BareMetalServerLocalSubnetOpenApiV2ControllerApiService service
type BareMetalServerLongRunningTaskOpenApiControllerApiService service
type BareMetalServerLongRunningTaskOpenApiV2ControllerApiService service
type BareMetalServerSimpleTaskOpenApiControllerApiService service
type BareMetalServerSimpleTaskOpenApiV3ControllerApiService service
type BareMetalServerStaticNatOpenApiControllerApiService service
type BareMetalServerStaticNatOpenApiV2ControllerApiService service
type ImOpenApiControllerApiService service

func NewAPIClient(cfg *scpsdk.Configuration) *APIClient {
	c := &APIClient{cfg: cfg}
	c.common.client = c
	c.BareMetalServerCreateDeleteOpenApiControllerApi = (*BareMetalServerCreateDeleteOpenApiControllerApiService)(&c.common)
	c.BareMetalServerCreateDeleteOpenApiV2ControllerApi = (*BareMetalServerCreateDeleteOpenApiV2ControllerApiService)(&c.common)
	c.BareMetalServerCreateDeleteOpenApiV3ControllerApi = (*BareMetalServerCreateDeleteOpenApiV3ControllerApiService)(&c.common)
	c.BareMetalServerLocalSubnetOpenApiControllerApi = (*BareMetalServerLocalSubnetOpenApiControllerApiService)(&c.common)
	c.BareMetalServerLocalSubnetOpenApiV2ControllerApi = (*BareMetalServerLocalSubnetOpenApiV2ControllerApiService)(&c.common)
	c.BareMetalServerLongRunningTaskOpenApiControllerApi = (*BareMetalServerLongRunningTaskOpenApiControllerApiService)(&c.common)
	c.BareMetalServerLongRunningTaskOpenApiV2ControllerApi = (*BareMetalServerLongRunningTaskOpenApiV2ControllerApiService)(&c.common)
	c.BareMetalServerSimpleTaskOpenApiControllerApi = (*BareMetalServerSimpleTaskOpenApiControllerApiService)(&c.common)
	c.BareMetalServerSimpleTaskOpenApiV3ControllerApi = (*BareMetalServerSimpleTaskOpenApiV3ControllerApiService)(&c.common)
	c.BareMetalServerStaticNatOpenApiControllerApi = (*BareMetalServerStaticNatOpenApiControllerApiService)(&c.common)
	c.BareMetalServerStaticNatOpenApiV2ControllerApi = (*BareMetalServerStaticNatOpenApiV2ControllerApiService)(&c.common)
	c.ImOpenApiControllerApi = (*ImOpenApiControllerApiService)(&c.common)
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
