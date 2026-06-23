package backup2

import (
	scpsdk "github.com/SamsungSDSCloud/terraform-sdk-samsungcloudplatform/v3/client"
)

type APIClient struct {
	cfg    *scpsdk.Configuration
	common service
	BackupAgentOperateOpenApiApi *BackupAgentOperateOpenApiApiService
	BackupAgentSearchOpenApiApi *BackupAgentSearchOpenApiApiService
	BackupDrOpenApiV2Api *BackupDrOpenApiV2ApiService
	BackupOperateOpenApiApi *BackupOperateOpenApiApiService
	BackupSearchOpenApiApi *BackupSearchOpenApiApiService
}

type service struct {
	client *APIClient
}

type BackupAgentOperateOpenApiApiService service
type BackupAgentSearchOpenApiApiService service
type BackupDrOpenApiV2ApiService service
type BackupOperateOpenApiApiService service
type BackupSearchOpenApiApiService service

func NewAPIClient(cfg *scpsdk.Configuration) *APIClient {
	c := &APIClient{cfg: cfg}
	c.common.client = c
	c.BackupAgentOperateOpenApiApi = (*BackupAgentOperateOpenApiApiService)(&c.common)
	c.BackupAgentSearchOpenApiApi = (*BackupAgentSearchOpenApiApiService)(&c.common)
	c.BackupDrOpenApiV2Api = (*BackupDrOpenApiV2ApiService)(&c.common)
	c.BackupOperateOpenApiApi = (*BackupOperateOpenApiApiService)(&c.common)
	c.BackupSearchOpenApiApi = (*BackupSearchOpenApiApiService)(&c.common)
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
