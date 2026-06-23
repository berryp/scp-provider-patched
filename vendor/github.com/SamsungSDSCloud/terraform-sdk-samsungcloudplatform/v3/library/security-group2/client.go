package securitygroup2

import (
	scpsdk "github.com/SamsungSDSCloud/terraform-sdk-samsungcloudplatform/v3/client"
)

type APIClient struct {
	cfg    *scpsdk.Configuration
	common service
	SecurityGroupLogStorageOpenApiControllerV2Api *SecurityGroupLogStorageOpenApiControllerV2ApiService
	SecurityGroupOpenApiControllerV2Api *SecurityGroupOpenApiControllerV2ApiService
	SecurityGroupOpenApiControllerV3Api *SecurityGroupOpenApiControllerV3ApiService
}

type service struct {
	client *APIClient
}

type SecurityGroupLogStorageOpenApiControllerV2ApiService service
type SecurityGroupOpenApiControllerV2ApiService service
type SecurityGroupOpenApiControllerV3ApiService service

func NewAPIClient(cfg *scpsdk.Configuration) *APIClient {
	c := &APIClient{cfg: cfg}
	c.common.client = c
	c.SecurityGroupLogStorageOpenApiControllerV2Api = (*SecurityGroupLogStorageOpenApiControllerV2ApiService)(&c.common)
	c.SecurityGroupOpenApiControllerV2Api = (*SecurityGroupOpenApiControllerV2ApiService)(&c.common)
	c.SecurityGroupOpenApiControllerV3Api = (*SecurityGroupOpenApiControllerV3ApiService)(&c.common)
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
