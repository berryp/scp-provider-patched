package servergroup2

import (
	scpsdk "github.com/SamsungSDSCloud/terraform-sdk-samsungcloudplatform/v3/client"
)

type APIClient struct {
	cfg    *scpsdk.Configuration
	common service
	ServerGroupOperationControllerApi *ServerGroupOperationControllerApiService
	ServerGroupSearchControllerApi *ServerGroupSearchControllerApiService
}

type service struct {
	client *APIClient
}

type ServerGroupOperationControllerApiService service
type ServerGroupSearchControllerApiService service

func NewAPIClient(cfg *scpsdk.Configuration) *APIClient {
	c := &APIClient{cfg: cfg}
	c.common.client = c
	c.ServerGroupOperationControllerApi = (*ServerGroupOperationControllerApiService)(&c.common)
	c.ServerGroupSearchControllerApi = (*ServerGroupSearchControllerApiService)(&c.common)
	return c
}

func (c *APIClient) ChangeBasePath(p string) { c.cfg.BasePath = p }
