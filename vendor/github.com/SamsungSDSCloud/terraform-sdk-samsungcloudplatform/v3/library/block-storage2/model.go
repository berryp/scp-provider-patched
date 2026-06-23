package blockstorage2

import (
	"github.com/antihax/optional"
	"time"
)

type APIKey struct {
	Key    string
	Prefix string
}

type AsyncResponse struct {
	ProjectId  string `json:"projectId,omitempty"`
	RequestId  string `json:"requestId,omitempty"`
	ResourceId string `json:"resourceId,omitempty"`
}

type BasicAuth struct {
	UserName string `json:"userName,omitempty"`
	Password string `json:"password,omitempty"`
}

type BlockStorageAttachRequest struct {
	VirtualServerId string `json:"virtualServerId"`
}

type BlockStorageControllerApiListBlockStorageVirtualServersOpts struct {
	Page optional.Int32
	Size optional.Int32
	Sort optional.Interface
}

type BlockStorageControllerApiListBlockStoragesOpts struct {
	BlockStorageName  optional.String
	VirtualServerId   optional.String
	VirtualServerName optional.String
	CreatedBy         optional.String
	Page              optional.Int32
	Size              optional.Int32
	Sort              optional.Interface
}

type BlockStorageCreateV3Request struct {
	BlockStorageName string                        `json:"blockStorageName"`
	BlockStorageSize int32                         `json:"blockStorageSize"`
	DiskType         string                        `json:"diskType"`
	EncryptEnabled   *bool                         `json:"encryptEnabled,omitempty"`
	SharedType       string                        `json:"sharedType,omitempty"`
	Tags             []TagRequest `json:"tags,omitempty"`
	VirtualServerId  string                        `json:"virtualServerId"`
}

type BlockStorageDetachRequest struct {
	VirtualServerId string `json:"virtualServerId"`
}

type BlockStorageMountResponse struct {
	MountState      string `json:"mountState,omitempty"`
	VirtualServerId string `json:"virtualServerId,omitempty"`
}

type BlockStorageResizeV3Request struct {
	BlockStorageSize int32 `json:"blockStorageSize"`
}

type BlockStorageResponse struct {
	ProjectId            string                                       `json:"projectId,omitempty"`
	AvailabilityZoneName string                                       `json:"availabilityZoneName,omitempty"`
	BlockId              string                                       `json:"blockId,omitempty"`
	BlockStorageId       string                                       `json:"blockStorageId,omitempty"`
	BlockStorageName     string                                       `json:"blockStorageName,omitempty"`
	BlockStorageSize     int32                                        `json:"blockStorageSize,omitempty"`
	BlockStorageState    string                                       `json:"blockStorageState,omitempty"`
	BlockStorageUuid     string                                       `json:"blockStorageUuid,omitempty"`
	DeviceNode           string                                       `json:"deviceNode,omitempty"`
	EncryptEnabled       *bool                                        `json:"encryptEnabled,omitempty"`
	IsBootDisk           *bool                                        `json:"isBootDisk,omitempty"`
	MountPath            string                                       `json:"mountPath,omitempty"`
	ProductId            string                                       `json:"productId,omitempty"`
	ServiceZoneId        string                                       `json:"serviceZoneId,omitempty"`
	SharedType           string                                       `json:"sharedType,omitempty"`
	VirtualServerId      string                                       `json:"virtualServerId,omitempty"`
	VirtualServers       []BlockStorageMountResponse `json:"virtualServers,omitempty"`
	CreatedBy            string                                       `json:"createdBy,omitempty"`
	CreatedDt            time.Time                                 `json:"createdDt,omitempty"`
	ModifiedBy           string                                       `json:"modifiedBy,omitempty"`
	ModifiedDt           time.Time                                 `json:"modifiedDt,omitempty"`
}

type BlockStorageVirtualServerResponse struct {
	ProjectId       string       `json:"projectId,omitempty"`
	BlockId         string       `json:"blockId,omitempty"`
	DeviceNode      string       `json:"deviceNode,omitempty"`
	MountState      string       `json:"mountState,omitempty"`
	ServiceZoneId   string       `json:"serviceZoneId,omitempty"`
	VirtualServerId string       `json:"virtualServerId,omitempty"`
	CreatedBy       string       `json:"createdBy,omitempty"`
	CreatedDt       time.Time `json:"createdDt,omitempty"`
	ModifiedBy      string       `json:"modifiedBy,omitempty"`
	ModifiedDt      time.Time `json:"modifiedDt,omitempty"`
}

type ListResponseBlockStorageResponse struct {
	Contents   []BlockStorageResponse `json:"contents,omitempty"`
	TotalCount int32                                   `json:"totalCount,omitempty"`
}

type ListResponseBlockStorageVirtualServerResponse struct {
	Contents   []BlockStorageVirtualServerResponse `json:"contents,omitempty"`
	TotalCount int32                                                `json:"totalCount,omitempty"`
}

type TagRequest struct {
	TagKey   string `json:"tagKey"`
	TagValue string `json:"tagValue,omitempty"`
}

// GenericSwaggerError is the per-service error type (the SDK's OAG template
// uses the older "swagger" naming).
type GenericSwaggerError struct {
	body  []byte
	error string
	model interface{}
}

func (e GenericSwaggerError) Body() []byte       { return e.body }
func (e GenericSwaggerError) Error() string      { return e.error }
func (e GenericSwaggerError) Model() interface{} { return e.model }
