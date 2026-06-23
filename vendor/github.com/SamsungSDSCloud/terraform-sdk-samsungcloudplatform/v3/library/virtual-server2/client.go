package virtualserver2

import (
	scpsdk "github.com/SamsungSDSCloud/terraform-sdk-samsungcloudplatform/v3/client"
)

type APIClient struct {
	cfg    *scpsdk.Configuration
	common service
	VirtualServerCreateDeleteV2Api *VirtualServerCreateDeleteV2ApiService
	VirtualServerCreateV3Api *VirtualServerCreateV3ApiService
	VirtualServerCreateV4Api *VirtualServerCreateV4ApiService
	VirtualServerNicV2Api *VirtualServerNicV2ApiService
	VirtualServerOperateV2Api *VirtualServerOperateV2ApiService
	VirtualServerOperateV3Api *VirtualServerOperateV3ApiService
	VirtualServerRoleV2Api *VirtualServerRoleV2ApiService
	VirtualServerV2Api *VirtualServerV2ApiService
	VirtualServerV3Api *VirtualServerV3ApiService
}

type service struct {
	client *APIClient
}

type VirtualServerCreateDeleteV2ApiService service
type VirtualServerCreateV3ApiService service
type VirtualServerCreateV4ApiService service
type VirtualServerNicV2ApiService service
type VirtualServerOperateV2ApiService service
type VirtualServerOperateV3ApiService service
type VirtualServerRoleV2ApiService service
type VirtualServerV2ApiService service
type VirtualServerV3ApiService service

func NewAPIClient(cfg *scpsdk.Configuration) *APIClient {
	c := &APIClient{cfg: cfg}
	c.common.client = c
	c.VirtualServerCreateDeleteV2Api = (*VirtualServerCreateDeleteV2ApiService)(&c.common)
	c.VirtualServerCreateV3Api = (*VirtualServerCreateV3ApiService)(&c.common)
	c.VirtualServerCreateV4Api = (*VirtualServerCreateV4ApiService)(&c.common)
	c.VirtualServerNicV2Api = (*VirtualServerNicV2ApiService)(&c.common)
	c.VirtualServerOperateV2Api = (*VirtualServerOperateV2ApiService)(&c.common)
	c.VirtualServerOperateV3Api = (*VirtualServerOperateV3ApiService)(&c.common)
	c.VirtualServerRoleV2Api = (*VirtualServerRoleV2ApiService)(&c.common)
	c.VirtualServerV2Api = (*VirtualServerV2ApiService)(&c.common)
	c.VirtualServerV3Api = (*VirtualServerV3ApiService)(&c.common)
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
