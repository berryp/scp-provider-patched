package iam

import (
	scpsdk "github.com/SamsungSDSCloud/terraform-sdk-samsungcloudplatform/v3/client"
)

type APIClient struct {
	cfg    *scpsdk.Configuration
	common service
	AccessKeyControllerApi *AccessKeyControllerApiService
	AuthorizationDashboardControllerApi *AuthorizationDashboardControllerApiService
	CommonCodeControllerApi *CommonCodeControllerApiService
	GroupControllerApi *GroupControllerApiService
	MemberControllerApi *MemberControllerApiService
	PdpControllerApi *PdpControllerApiService
	PlatformControllerApi *PlatformControllerApiService
	PolicyControllerApi *PolicyControllerApiService
	ProjectControllerApi *ProjectControllerApiService
	RoleControllerApi *RoleControllerApiService
	SamlControllerApi *SamlControllerApiService
	SecurityControllerApi *SecurityControllerApiService
	TemporaryAccessKeyControllerApi *TemporaryAccessKeyControllerApiService
	UserControllerApi *UserControllerApiService
	Client *Client
}

type service struct {
	client *APIClient
}

type AccessKeyControllerApiService service
type AuthorizationDashboardControllerApiService service
type CommonCodeControllerApiService service
type GroupControllerApiService service
type MemberControllerApiService service
type PdpControllerApiService service
type PlatformControllerApiService service
type PolicyControllerApiService service
type ProjectControllerApiService service
type RoleControllerApiService service
type SamlControllerApiService service
type SecurityControllerApiService service
type TemporaryAccessKeyControllerApiService service
type UserControllerApiService service
type Client service

func NewAPIClient(cfg *scpsdk.Configuration) *APIClient {
	c := &APIClient{cfg: cfg}
	c.common.client = c
	c.AccessKeyControllerApi = (*AccessKeyControllerApiService)(&c.common)
	c.AuthorizationDashboardControllerApi = (*AuthorizationDashboardControllerApiService)(&c.common)
	c.CommonCodeControllerApi = (*CommonCodeControllerApiService)(&c.common)
	c.GroupControllerApi = (*GroupControllerApiService)(&c.common)
	c.MemberControllerApi = (*MemberControllerApiService)(&c.common)
	c.PdpControllerApi = (*PdpControllerApiService)(&c.common)
	c.PlatformControllerApi = (*PlatformControllerApiService)(&c.common)
	c.PolicyControllerApi = (*PolicyControllerApiService)(&c.common)
	c.ProjectControllerApi = (*ProjectControllerApiService)(&c.common)
	c.RoleControllerApi = (*RoleControllerApiService)(&c.common)
	c.SamlControllerApi = (*SamlControllerApiService)(&c.common)
	c.SecurityControllerApi = (*SecurityControllerApiService)(&c.common)
	c.TemporaryAccessKeyControllerApi = (*TemporaryAccessKeyControllerApiService)(&c.common)
	c.UserControllerApi = (*UserControllerApiService)(&c.common)
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
