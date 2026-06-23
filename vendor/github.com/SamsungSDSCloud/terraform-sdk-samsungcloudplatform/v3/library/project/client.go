package project

import (
	scpsdk "github.com/SamsungSDSCloud/terraform-sdk-samsungcloudplatform/v3/client"
)

type APIClient struct {
	cfg    *scpsdk.Configuration
	common service
	AccountV3ControllerApi *AccountV3ControllerApiService
	CommonCodeControllerApi *CommonCodeControllerApiService
	PlatformControllerApi *PlatformControllerApiService
	ProjectControllerV2Api *ProjectControllerV2ApiService
	ProjectV3ControllerApi *ProjectV3ControllerApiService
	ZoneControllerV2Api *ZoneControllerV2ApiService
	ZoneV3ControllerApi *ZoneV3ControllerApiService
	Client *Client
}

type service struct {
	client *APIClient
}

type AccountV3ControllerApiService service
type CommonCodeControllerApiService service
type PlatformControllerApiService service
type ProjectControllerV2ApiService service
type ProjectV3ControllerApiService service
type ZoneControllerV2ApiService service
type ZoneV3ControllerApiService service
type Client service

func NewAPIClient(cfg *scpsdk.Configuration) *APIClient {
	c := &APIClient{cfg: cfg}
	c.common.client = c
	c.AccountV3ControllerApi = (*AccountV3ControllerApiService)(&c.common)
	c.CommonCodeControllerApi = (*CommonCodeControllerApiService)(&c.common)
	c.PlatformControllerApi = (*PlatformControllerApiService)(&c.common)
	c.ProjectControllerV2Api = (*ProjectControllerV2ApiService)(&c.common)
	c.ProjectV3ControllerApi = (*ProjectV3ControllerApiService)(&c.common)
	c.ZoneControllerV2Api = (*ZoneControllerV2ApiService)(&c.common)
	c.ZoneV3ControllerApi = (*ZoneV3ControllerApiService)(&c.common)
	c.Client = (*Client)(&c.common)
	return c
}

func (c *APIClient) ChangeBasePath(p string) { c.cfg.BasePath = p }
