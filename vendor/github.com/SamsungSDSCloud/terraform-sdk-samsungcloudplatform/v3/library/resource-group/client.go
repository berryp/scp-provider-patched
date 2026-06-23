package resourcegroup

import (
	scpsdk "github.com/SamsungSDSCloud/terraform-sdk-samsungcloudplatform/v3/client"
)

type APIClient struct {
	cfg    *scpsdk.Configuration
	common service
	CommonCodeControllerApi *CommonCodeControllerApiService
	MyProjectResourceControllerApi *MyProjectResourceControllerApiService
	MyProjectResourceGroupControllerApi *MyProjectResourceGroupControllerApiService
	PlatformControllerApi *PlatformControllerApiService
	ResourceControllerApi *ResourceControllerApiService
	ResourceGroupControllerApi *ResourceGroupControllerApiService
	TypeControllerApi *TypeControllerApiService
	Client *Client
}

type service struct {
	client *APIClient
}

type CommonCodeControllerApiService service
type MyProjectResourceControllerApiService service
type MyProjectResourceGroupControllerApiService service
type PlatformControllerApiService service
type ResourceControllerApiService service
type ResourceGroupControllerApiService service
type TypeControllerApiService service
type Client service

func NewAPIClient(cfg *scpsdk.Configuration) *APIClient {
	c := &APIClient{cfg: cfg}
	c.common.client = c
	c.CommonCodeControllerApi = (*CommonCodeControllerApiService)(&c.common)
	c.MyProjectResourceControllerApi = (*MyProjectResourceControllerApiService)(&c.common)
	c.MyProjectResourceGroupControllerApi = (*MyProjectResourceGroupControllerApiService)(&c.common)
	c.PlatformControllerApi = (*PlatformControllerApiService)(&c.common)
	c.ResourceControllerApi = (*ResourceControllerApiService)(&c.common)
	c.ResourceGroupControllerApi = (*ResourceGroupControllerApiService)(&c.common)
	c.TypeControllerApi = (*TypeControllerApiService)(&c.common)
	c.Client = (*Client)(&c.common)
	return c
}

func (c *APIClient) ChangeBasePath(p string) { c.cfg.BasePath = p }
