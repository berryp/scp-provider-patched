package firewall2

import (
	scpsdk "github.com/SamsungSDSCloud/terraform-sdk-samsungcloudplatform/v3/client"
)

type APIClient struct {
	cfg    *scpsdk.Configuration
	common service
	FirewallLogStorageV2Api *FirewallLogStorageV2ApiService
	FirewallRuleV2Api *FirewallRuleV2ApiService
	FirewallV2Api *FirewallV2ApiService
}

type service struct {
	client *APIClient
}

type FirewallLogStorageV2ApiService service
type FirewallRuleV2ApiService service
type FirewallV2ApiService service

func NewAPIClient(cfg *scpsdk.Configuration) *APIClient {
	c := &APIClient{cfg: cfg}
	c.common.client = c
	c.FirewallLogStorageV2Api = (*FirewallLogStorageV2ApiService)(&c.common)
	c.FirewallRuleV2Api = (*FirewallRuleV2ApiService)(&c.common)
	c.FirewallV2Api = (*FirewallV2ApiService)(&c.common)
	return c
}

func (c *APIClient) ChangeBasePath(p string) { c.cfg.BasePath = p }
