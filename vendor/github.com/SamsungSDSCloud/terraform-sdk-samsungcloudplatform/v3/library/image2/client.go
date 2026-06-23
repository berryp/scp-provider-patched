package image2

import (
	scpsdk "github.com/SamsungSDSCloud/terraform-sdk-samsungcloudplatform/v3/client"
)

type APIClient struct {
	cfg    *scpsdk.Configuration
	common service
	CommonImageV2Api *CommonImageV2ApiService
	CustomImageCloneV2Api *CustomImageCloneV2ApiService
	CustomImageV2Api *CustomImageV2ApiService
	MigrationImageV2Api *MigrationImageV2ApiService
	StandardImageV2Api *StandardImageV2ApiService
}

type service struct {
	client *APIClient
}

type CommonImageV2ApiService service
type CustomImageCloneV2ApiService service
type CustomImageV2ApiService service
type MigrationImageV2ApiService service
type StandardImageV2ApiService service

func NewAPIClient(cfg *scpsdk.Configuration) *APIClient {
	c := &APIClient{cfg: cfg}
	c.common.client = c
	c.CommonImageV2Api = (*CommonImageV2ApiService)(&c.common)
	c.CustomImageCloneV2Api = (*CustomImageCloneV2ApiService)(&c.common)
	c.CustomImageV2Api = (*CustomImageV2ApiService)(&c.common)
	c.MigrationImageV2Api = (*MigrationImageV2ApiService)(&c.common)
	c.StandardImageV2Api = (*StandardImageV2ApiService)(&c.common)
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
