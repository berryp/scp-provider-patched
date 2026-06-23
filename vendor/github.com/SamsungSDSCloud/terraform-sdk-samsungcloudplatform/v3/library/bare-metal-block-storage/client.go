package baremetalblockstorage

import (
	scpsdk "github.com/SamsungSDSCloud/terraform-sdk-samsungcloudplatform/v3/client"
)

type APIClient struct {
	cfg    *scpsdk.Configuration
	common service
	BareMetalBlockStorageOpenApiV1Api *BareMetalBlockStorageOpenApiV1ApiService
	BareMetalBlockStorageOpenApiV2Api *BareMetalBlockStorageOpenApiV2ApiService
	BareMetalBlockStorageSnapshotCloneApiV1Api *BareMetalBlockStorageSnapshotCloneApiV1ApiService
	BareMetalBlockStorageSnapshotScheduleOpenApiV1Api *BareMetalBlockStorageSnapshotScheduleOpenApiV1ApiService
	BmBlockStorageControllerApi *BmBlockStorageControllerApiService
	BmBlockStorageDrControllerApi *BmBlockStorageDrControllerApiService
	BmBlockStorageV2ControllerApi *BmBlockStorageV2ControllerApiService
	BmBlockStorageV3ControllerApi *BmBlockStorageV3ControllerApiService
	BmBlockStorageV4ControllerApi *BmBlockStorageV4ControllerApiService
	ConsistencyGroupControllerApi *ConsistencyGroupControllerApiService
	ConsistencyGroupDrControllerApi *ConsistencyGroupDrControllerApiService
	ConsistencyGroupSnapshotCloneOpenApiControllerApi *ConsistencyGroupSnapshotCloneOpenApiControllerApiService
	ConsistencyGroupSnapshotControllerApi *ConsistencyGroupSnapshotControllerApiService
	ConsistencyGroupSnapshotScheduleControllerApi *ConsistencyGroupSnapshotScheduleControllerApiService
}

type service struct {
	client *APIClient
}

type BareMetalBlockStorageOpenApiV1ApiService service
type BareMetalBlockStorageOpenApiV2ApiService service
type BareMetalBlockStorageSnapshotCloneApiV1ApiService service
type BareMetalBlockStorageSnapshotScheduleOpenApiV1ApiService service
type BmBlockStorageControllerApiService service
type BmBlockStorageDrControllerApiService service
type BmBlockStorageV2ControllerApiService service
type BmBlockStorageV3ControllerApiService service
type BmBlockStorageV4ControllerApiService service
type ConsistencyGroupControllerApiService service
type ConsistencyGroupDrControllerApiService service
type ConsistencyGroupSnapshotCloneOpenApiControllerApiService service
type ConsistencyGroupSnapshotControllerApiService service
type ConsistencyGroupSnapshotScheduleControllerApiService service

func NewAPIClient(cfg *scpsdk.Configuration) *APIClient {
	c := &APIClient{cfg: cfg}
	c.common.client = c
	c.BareMetalBlockStorageOpenApiV1Api = (*BareMetalBlockStorageOpenApiV1ApiService)(&c.common)
	c.BareMetalBlockStorageOpenApiV2Api = (*BareMetalBlockStorageOpenApiV2ApiService)(&c.common)
	c.BareMetalBlockStorageSnapshotCloneApiV1Api = (*BareMetalBlockStorageSnapshotCloneApiV1ApiService)(&c.common)
	c.BareMetalBlockStorageSnapshotScheduleOpenApiV1Api = (*BareMetalBlockStorageSnapshotScheduleOpenApiV1ApiService)(&c.common)
	c.BmBlockStorageControllerApi = (*BmBlockStorageControllerApiService)(&c.common)
	c.BmBlockStorageDrControllerApi = (*BmBlockStorageDrControllerApiService)(&c.common)
	c.BmBlockStorageV2ControllerApi = (*BmBlockStorageV2ControllerApiService)(&c.common)
	c.BmBlockStorageV3ControllerApi = (*BmBlockStorageV3ControllerApiService)(&c.common)
	c.BmBlockStorageV4ControllerApi = (*BmBlockStorageV4ControllerApiService)(&c.common)
	c.ConsistencyGroupControllerApi = (*ConsistencyGroupControllerApiService)(&c.common)
	c.ConsistencyGroupDrControllerApi = (*ConsistencyGroupDrControllerApiService)(&c.common)
	c.ConsistencyGroupSnapshotCloneOpenApiControllerApi = (*ConsistencyGroupSnapshotCloneOpenApiControllerApiService)(&c.common)
	c.ConsistencyGroupSnapshotControllerApi = (*ConsistencyGroupSnapshotControllerApiService)(&c.common)
	c.ConsistencyGroupSnapshotScheduleControllerApi = (*ConsistencyGroupSnapshotScheduleControllerApiService)(&c.common)
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
