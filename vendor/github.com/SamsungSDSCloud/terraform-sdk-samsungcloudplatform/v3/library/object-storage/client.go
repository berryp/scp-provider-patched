package objectstorage

import (
	scpsdk "github.com/SamsungSDSCloud/terraform-sdk-samsungcloudplatform/v3/client"
)

type APIClient struct {
	cfg    *scpsdk.Configuration
	common service
	ObjectStorageBucketV4ControllerApi *ObjectStorageBucketV4ControllerApiService
	ObjectStorageDrV4ControllerApi *ObjectStorageDrV4ControllerApiService
	ObjectStorageIpsV4ControllerApi *ObjectStorageIpsV4ControllerApiService
	ObjectStorageObjectV4ControllerApi *ObjectStorageObjectV4ControllerApiService
	ObjectStorageV4ControllerApi *ObjectStorageV4ControllerApiService
	Client *Client
}

type service struct {
	client *APIClient
}

type ObjectStorageBucketV4ControllerApiService service
type ObjectStorageDrV4ControllerApiService service
type ObjectStorageIpsV4ControllerApiService service
type ObjectStorageObjectV4ControllerApiService service
type ObjectStorageV4ControllerApiService service
type Client service

func NewAPIClient(cfg *scpsdk.Configuration) *APIClient {
	c := &APIClient{cfg: cfg}
	c.common.client = c
	c.ObjectStorageBucketV4ControllerApi = (*ObjectStorageBucketV4ControllerApiService)(&c.common)
	c.ObjectStorageDrV4ControllerApi = (*ObjectStorageDrV4ControllerApiService)(&c.common)
	c.ObjectStorageIpsV4ControllerApi = (*ObjectStorageIpsV4ControllerApiService)(&c.common)
	c.ObjectStorageObjectV4ControllerApi = (*ObjectStorageObjectV4ControllerApiService)(&c.common)
	c.ObjectStorageV4ControllerApi = (*ObjectStorageV4ControllerApiService)(&c.common)
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
