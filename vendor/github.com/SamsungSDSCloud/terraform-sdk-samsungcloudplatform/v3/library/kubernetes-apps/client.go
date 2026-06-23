package kubernetesapps

import (
	scpsdk "github.com/SamsungSDSCloud/terraform-sdk-samsungcloudplatform/v3/client"
)

type APIClient struct {
	cfg    *scpsdk.Configuration
	common service
	CommonCodeControllerApi *CommonCodeControllerApiService
	ImageApi *ImageApiService
	K8sAppsApi *K8sAppsApiService
	NamespaceApi *NamespaceApiService
	ReleaseApi *ReleaseApiService
	Client *Client
}

type service struct {
	client *APIClient
}

type CommonCodeControllerApiService service
type ImageApiService service
type K8sAppsApiService service
type NamespaceApiService service
type ReleaseApiService service
type Client service

func NewAPIClient(cfg *scpsdk.Configuration) *APIClient {
	c := &APIClient{cfg: cfg}
	c.common.client = c
	c.CommonCodeControllerApi = (*CommonCodeControllerApiService)(&c.common)
	c.ImageApi = (*ImageApiService)(&c.common)
	c.K8sAppsApi = (*K8sAppsApiService)(&c.common)
	c.NamespaceApi = (*NamespaceApiService)(&c.common)
	c.ReleaseApi = (*ReleaseApiService)(&c.common)
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
