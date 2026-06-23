package kubernetesengine2

import (
	scpsdk "github.com/SamsungSDSCloud/terraform-sdk-samsungcloudplatform/v3/client"
)

type APIClient struct {
	cfg    *scpsdk.Configuration
	common service
	K8sEngineV2Api *K8sEngineV2ApiService
	K8sEngineV3Api *K8sEngineV3ApiService
	K8sEngineV4Api *K8sEngineV4ApiService
	K8sTemplateV2Api *K8sTemplateV2ApiService
	NodePoolV2Api *NodePoolV2ApiService
	NodePoolV3Api *NodePoolV3ApiService
	NodePoolV4Api *NodePoolV4ApiService
	NodeV2Api *NodeV2ApiService
}

type service struct {
	client *APIClient
}

type K8sEngineV2ApiService service
type K8sEngineV3ApiService service
type K8sEngineV4ApiService service
type K8sTemplateV2ApiService service
type NodePoolV2ApiService service
type NodePoolV3ApiService service
type NodePoolV4ApiService service
type NodeV2ApiService service

func NewAPIClient(cfg *scpsdk.Configuration) *APIClient {
	c := &APIClient{cfg: cfg}
	c.common.client = c
	c.K8sEngineV2Api = (*K8sEngineV2ApiService)(&c.common)
	c.K8sEngineV3Api = (*K8sEngineV3ApiService)(&c.common)
	c.K8sEngineV4Api = (*K8sEngineV4ApiService)(&c.common)
	c.K8sTemplateV2Api = (*K8sTemplateV2ApiService)(&c.common)
	c.NodePoolV2Api = (*NodePoolV2ApiService)(&c.common)
	c.NodePoolV3Api = (*NodePoolV3ApiService)(&c.common)
	c.NodePoolV4Api = (*NodePoolV4ApiService)(&c.common)
	c.NodeV2Api = (*NodeV2ApiService)(&c.common)
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
