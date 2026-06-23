package tag

import (
	scpsdk "github.com/SamsungSDSCloud/terraform-sdk-samsungcloudplatform/v3/client"
)

type APIClient struct {
	cfg    *scpsdk.Configuration
	common service
	CommonCodeControllerApi *CommonCodeControllerApiService
	PlatformControllerApi *PlatformControllerApiService
	ResourceTagControllerApi *ResourceTagControllerApiService
	TagPolicyControllerApi *TagPolicyControllerApiService
	TagPolicyV3ControllerApi *TagPolicyV3ControllerApiService
	Client *Client
}

type service struct {
	client *APIClient
}

type CommonCodeControllerApiService service
type PlatformControllerApiService service
type ResourceTagControllerApiService service
type TagPolicyControllerApiService service
type TagPolicyV3ControllerApiService service
type Client service

func NewAPIClient(cfg *scpsdk.Configuration) *APIClient {
	c := &APIClient{cfg: cfg}
	c.common.client = c
	c.CommonCodeControllerApi = (*CommonCodeControllerApiService)(&c.common)
	c.PlatformControllerApi = (*PlatformControllerApiService)(&c.common)
	c.ResourceTagControllerApi = (*ResourceTagControllerApiService)(&c.common)
	c.TagPolicyControllerApi = (*TagPolicyControllerApiService)(&c.common)
	c.TagPolicyV3ControllerApi = (*TagPolicyV3ControllerApiService)(&c.common)
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
