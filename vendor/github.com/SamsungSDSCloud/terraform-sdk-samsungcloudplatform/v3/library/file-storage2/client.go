package filestorage2

import (
	scpsdk "github.com/SamsungSDSCloud/terraform-sdk-samsungcloudplatform/v3/client"
)

type APIClient struct {
	cfg    *scpsdk.Configuration
	common service
	FileStorageOpenApiV3Api *FileStorageOpenApiV3ApiService
	FileStorageOpenApiV4Api *FileStorageOpenApiV4ApiService
}

type service struct {
	client *APIClient
}

type FileStorageOpenApiV3ApiService service
type FileStorageOpenApiV4ApiService service

func NewAPIClient(cfg *scpsdk.Configuration) *APIClient {
	c := &APIClient{cfg: cfg}
	c.common.client = c
	c.FileStorageOpenApiV3Api = (*FileStorageOpenApiV3ApiService)(&c.common)
	c.FileStorageOpenApiV4Api = (*FileStorageOpenApiV4ApiService)(&c.common)
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
