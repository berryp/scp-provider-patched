package product

import (
	scpsdk "github.com/SamsungSDSCloud/terraform-sdk-samsungcloudplatform/v3/client"
)

type APIClient struct {
	cfg    *scpsdk.Configuration
	common service
	CommonCodeControllerApi *CommonCodeControllerApiService
	PlatformControllerApi *PlatformControllerApiService
	ProductAdminMenuControllerApi *ProductAdminMenuControllerApiService
	ProductV2ControllerApi *ProductV2ControllerApiService
	Client *Client
}

type service struct {
	client *APIClient
}

type CommonCodeControllerApiService service
type PlatformControllerApiService service
type ProductAdminMenuControllerApiService service
type ProductV2ControllerApiService service
type Client service

func NewAPIClient(cfg *scpsdk.Configuration) *APIClient {
	c := &APIClient{cfg: cfg}
	c.common.client = c
	c.CommonCodeControllerApi = (*CommonCodeControllerApiService)(&c.common)
	c.PlatformControllerApi = (*PlatformControllerApiService)(&c.common)
	c.ProductAdminMenuControllerApi = (*ProductAdminMenuControllerApiService)(&c.common)
	c.ProductV2ControllerApi = (*ProductV2ControllerApiService)(&c.common)
	c.Client = (*Client)(&c.common)
	return c
}

func (c *APIClient) ChangeBasePath(p string) { c.cfg.BasePath = p }
