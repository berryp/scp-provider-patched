package routing2

import (
	scpsdk "github.com/SamsungSDSCloud/terraform-sdk-samsungcloudplatform/v3/client"
)

type APIClient struct {
	cfg    *scpsdk.Configuration
	common service
	DirectConnectRoutingRuleOpenApiControllerApi *DirectConnectRoutingRuleOpenApiControllerApiService
	DirectConnectRoutingTableOpenApiControllerApi *DirectConnectRoutingTableOpenApiControllerApiService
	TransitGatewayRoutingRuleOpenApiControllerApi *TransitGatewayRoutingRuleOpenApiControllerApiService
	TransitGatewayRoutingTableOpenApiControllerApi *TransitGatewayRoutingTableOpenApiControllerApiService
	VpcRoutingRuleOpenApiControllerApi *VpcRoutingRuleOpenApiControllerApiService
	VpcRoutingTableOpenApiControllerApi *VpcRoutingTableOpenApiControllerApiService
}

type service struct {
	client *APIClient
}

type DirectConnectRoutingRuleOpenApiControllerApiService service
type DirectConnectRoutingTableOpenApiControllerApiService service
type TransitGatewayRoutingRuleOpenApiControllerApiService service
type TransitGatewayRoutingTableOpenApiControllerApiService service
type VpcRoutingRuleOpenApiControllerApiService service
type VpcRoutingTableOpenApiControllerApiService service

func NewAPIClient(cfg *scpsdk.Configuration) *APIClient {
	c := &APIClient{cfg: cfg}
	c.common.client = c
	c.DirectConnectRoutingRuleOpenApiControllerApi = (*DirectConnectRoutingRuleOpenApiControllerApiService)(&c.common)
	c.DirectConnectRoutingTableOpenApiControllerApi = (*DirectConnectRoutingTableOpenApiControllerApiService)(&c.common)
	c.TransitGatewayRoutingRuleOpenApiControllerApi = (*TransitGatewayRoutingRuleOpenApiControllerApiService)(&c.common)
	c.TransitGatewayRoutingTableOpenApiControllerApi = (*TransitGatewayRoutingTableOpenApiControllerApiService)(&c.common)
	c.VpcRoutingRuleOpenApiControllerApi = (*VpcRoutingRuleOpenApiControllerApiService)(&c.common)
	c.VpcRoutingTableOpenApiControllerApi = (*VpcRoutingTableOpenApiControllerApiService)(&c.common)
	return c
}

func (c *APIClient) ChangeBasePath(p string) { c.cfg.BasePath = p }
