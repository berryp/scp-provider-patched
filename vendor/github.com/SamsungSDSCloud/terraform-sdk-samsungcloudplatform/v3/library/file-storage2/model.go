package filestorage2

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

type BackupRelationFileStorage struct {
	BackupFileStorageIds []string `json:"backupFileStorageIds,omitempty"`
	BackupRelationCount  *int32   `json:"backupRelationCount,omitempty"`
}

type BasicAuth struct {
	UserName string `json:"userName,omitempty"`
	Password string `json:"password,omitempty"`
}

type CheckResponse struct {
	Result *bool `json:"result,omitempty"`
}

type CreateFileStorageRelationV4Request struct {
	BackupRetentionCount *int32                            `json:"backupRetentionCount,omitempty"`
	CifsPassword         string                            `json:"cifsPassword,omitempty"`
	FileStorageName      string                            `json:"fileStorageName"`
	ProductNames         []string                          `json:"productNames"`
	RelationSchedule     *RelationSchedule `json:"relationSchedule"`
	RelationType         string                            `json:"relationType"`
	ServiceZoneId        string                            `json:"serviceZoneId"`
	Tags                 []TagRequest      `json:"tags,omitempty"`
	TieringLifeCycle     *int32                            `json:"tieringLifeCycle,omitempty"`
}

type CreateFileStorageTieringV4Request struct {
	ProductNames     []string `json:"productNames"`
	ServiceZoneId    string   `json:"serviceZoneId"`
	TieringLifeCycle *int32   `json:"tieringLifeCycle"`
}

type CreateFileStorageV4Request struct {
	CifsPassword           string                            `json:"cifsPassword,omitempty"`
	DiskType               string                            `json:"diskType"`
	FileStorageName        string                            `json:"fileStorageName"`
	FileStorageProtocol    string                            `json:"fileStorageProtocol"`
	MultiAvailabilityZone  *bool                             `json:"multiAvailabilityZone,omitempty"`
	ProductNames           []string                          `json:"productNames"`
	ServiceZoneId          string                            `json:"serviceZoneId"`
	SnapshotRetentionCount *int32                            `json:"snapshotRetentionCount,omitempty"`
	SnapshotSchedule       *SnapshotSchedule `json:"snapshotSchedule,omitempty"`
	Tags                   []TagRequest      `json:"tags,omitempty"`
}

type DrFileStorage struct {
	DrCount          *int32   `json:"drCount,omitempty"`
	DrFileStorageIds []string `json:"drFileStorageIds,omitempty"`
}

type FileStorageDetailResponseV3 struct {
	ProjectId                 string                                     `json:"projectId,omitempty"`
	BackupRelationFileStorage *BackupRelationFileStorage `json:"backupRelationFileStorage,omitempty"`
	BlockId                   string                                     `json:"blockId,omitempty"`
	CifsId                    string                                     `json:"cifsId,omitempty"`
	DiskType                  string                                     `json:"diskType,omitempty"`
	DrFileStorage             *DrFileStorage             `json:"drFileStorage,omitempty"`
	EncryptionEnabled         *bool                                      `json:"encryptionEnabled,omitempty"`
	FileStorageCategory       string                                     `json:"fileStorageCategory,omitempty"`
	FileStorageId             string                                     `json:"fileStorageId,omitempty"`
	FileStorageName           string                                     `json:"fileStorageName,omitempty"`
	FileStorageProtocol       string                                     `json:"fileStorageProtocol,omitempty"`
	FileStoragePurpose        string                                     `json:"fileStoragePurpose,omitempty"`
	FileStorageState          string                                     `json:"fileStorageState,omitempty"`
	FileUnitRecoveryEnabled   *bool                                      `json:"fileUnitRecoveryEnabled,omitempty"`
	IsSnapshotDeleted         *bool                                      `json:"isSnapshotDeleted,omitempty"`
	IsSnapshotRestore         *bool                                      `json:"isSnapshotRestore,omitempty"`
	LinkedObjects             []LinkedObject             `json:"linkedObjects,omitempty"`
	MountInfo                 string                                     `json:"mountInfo,omitempty"`
	MultiAvailabilityZone     *bool                                      `json:"multiAvailabilityZone,omitempty"`
	OriginFileStorage         *OriginFileStorage         `json:"originFileStorage,omitempty"`
	ProductGroupId            string                                     `json:"productGroupId,omitempty"`
	ServiceZoneId             string                                     `json:"serviceZoneId,omitempty"`
	TieringEnabled            *bool                                      `json:"tieringEnabled,omitempty"`
	VolumePoolId              string                                     `json:"volumePoolId,omitempty"`
	VpcEndpointInfo           string                                     `json:"vpcEndpointInfo,omitempty"`
	VpcEndpointVolumePoolId   string                                     `json:"vpcEndpointVolumePoolId,omitempty"`
	CreatedBy                 string                                     `json:"createdBy,omitempty"`
	CreatedDt                 time.Time                               `json:"createdDt,omitempty"`
	ModifiedBy                string                                     `json:"modifiedBy,omitempty"`
	ModifiedDt                time.Time                               `json:"modifiedDt,omitempty"`
}

type FileStorageFileUnitRecoveryRequest struct {
	FileUnitRecoveryEnabled *bool `json:"fileUnitRecoveryEnabled,omitempty"`
}

type FileStorageLinkMaxCountResponse struct {
	LinkMaxCount *int32 `json:"linkMaxCount,omitempty"`
}

type FileStorageOpenApiV3ApiListFileStoragesOpts struct {
	BlockId             optional.String
	FileStorageCategory optional.String
	FileStorageId       optional.String
	FileStorageName     optional.String
	FileStorageProtocol optional.String
	FileStoragePurposes optional.Interface
	FileStorageState    optional.String
	FileStorageStates   optional.Interface
	ServiceZoneId       optional.String
	CreatedBy           optional.String
	Page                optional.Int32
	Size                optional.Int32
	Sort                optional.Interface
}

type FileStorageRelationInfoResponse struct {
	ProjectId            string                            `json:"projectId,omitempty"`
	BackupRetentionCount *int32                            `json:"backupRetentionCount,omitempty"`
	BlockId              string                            `json:"blockId,omitempty"`
	EncryptionEnabled    *bool                             `json:"encryptionEnabled,omitempty"`
	FileStorageId        string                            `json:"fileStorageId,omitempty"`
	FileStorageName      string                            `json:"fileStorageName,omitempty"`
	RelationLagTime      string                            `json:"relationLagTime,omitempty"`
	RelationPolicyType   string                            `json:"relationPolicyType,omitempty"`
	RelationSchedule     *RelationSchedule `json:"relationSchedule,omitempty"`
	RelationState        string                            `json:"relationState,omitempty"`
	RelationType         string                            `json:"relationType,omitempty"`
	ServiceZoneId        string                            `json:"serviceZoneId,omitempty"`
	CreatedBy            string                            `json:"createdBy,omitempty"`
	CreatedDt            time.Time                      `json:"createdDt,omitempty"`
	ModifiedBy           string                            `json:"modifiedBy,omitempty"`
	ModifiedDt           time.Time                      `json:"modifiedDt,omitempty"`
}

type FileStorageRelationLocationResponse struct {
	BlockId       string `json:"blockId,omitempty"`
	ServiceZoneId string `json:"serviceZoneId,omitempty"`
}

type FileStorageRelationPolicyRequest struct {
	RelationPolicyType string `json:"relationPolicyType"`
}

type FileStorageRelationScheduleRequest struct {
	BackupRetentionCount *int32                            `json:"backupRetentionCount,omitempty"`
	RelationPolicyType   string                            `json:"relationPolicyType"`
	RelationSchedule     *RelationSchedule `json:"relationSchedule"`
}

type FileStorageSnapshotItemResponse struct {
	SnapshotId     string       `json:"snapshotId,omitempty"`
	SnapshotName   string       `json:"snapshotName,omitempty"`
	SnapshotSizeMb float64      `json:"snapshotSizeMb,omitempty"`
	CreatedDt      time.Time `json:"createdDt,omitempty"`
}

type FileStorageSnapshotScheduleResponse struct {
	ProjectId              string                            `json:"projectId,omitempty"`
	BlockId                string                            `json:"blockId,omitempty"`
	FileStorageId          string                            `json:"fileStorageId,omitempty"`
	IsSnapshotPolicy       *bool                             `json:"isSnapshotPolicy,omitempty"`
	ServiceZoneId          string                            `json:"serviceZoneId,omitempty"`
	SnapshotRetentionCount *int32                            `json:"snapshotRetentionCount,omitempty"`
	SnapshotSchedule       *SnapshotSchedule `json:"snapshotSchedule,omitempty"`
	CreatedBy              string                            `json:"createdBy,omitempty"`
	CreatedDt              time.Time                      `json:"createdDt,omitempty"`
	ModifiedBy             string                            `json:"modifiedBy,omitempty"`
	ModifiedDt             time.Time                      `json:"modifiedDt,omitempty"`
}

type FileStorageSnapshotsResponse struct {
	ProjectId          string                                            `json:"projectId,omitempty"`
	BlockId            string                                            `json:"blockId,omitempty"`
	FileStorageId      string                                            `json:"fileStorageId,omitempty"`
	ServiceZoneId      string                                            `json:"serviceZoneId,omitempty"`
	SnapshotTotalUsage float64                                           `json:"snapshotTotalUsage,omitempty"`
	Snapshots          []FileStorageSnapshotItemResponse `json:"snapshots,omitempty"`
	CreatedBy          string                                            `json:"createdBy,omitempty"`
	CreatedDt          time.Time                                      `json:"createdDt,omitempty"`
	ModifiedBy         string                                            `json:"modifiedBy,omitempty"`
	ModifiedDt         time.Time                                      `json:"modifiedDt,omitempty"`
}

type FileStorageTieringItemResponse struct {
	TierType  string  `json:"tierType,omitempty"`
	TierUsage float64 `json:"tierUsage,omitempty"`
}

type FileStorageTieringResponse struct {
	TieringLifeCycle string                                           `json:"tieringLifeCycle,omitempty"`
	TieringPolicy    string                                           `json:"tieringPolicy,omitempty"`
	Tiers            []FileStorageTieringItemResponse `json:"tiers,omitempty"`
}

type FileStorageUsageResponse struct {
	ProjectId        string       `json:"projectId,omitempty"`
	BlockId          string       `json:"blockId,omitempty"`
	FileStorageUsage string       `json:"fileStorageUsage,omitempty"`
	ServiceZoneId    string       `json:"serviceZoneId,omitempty"`
	CreatedBy        string       `json:"createdBy,omitempty"`
	CreatedDt        time.Time `json:"createdDt,omitempty"`
	ModifiedBy       string       `json:"modifiedBy,omitempty"`
	ModifiedDt       time.Time `json:"modifiedDt,omitempty"`
}

type FileStoragesListResponse struct {
	ProjectId           string       `json:"projectId,omitempty"`
	BlockId             string       `json:"blockId,omitempty"`
	DiskType            string       `json:"diskType,omitempty"`
	EncryptionEnabled   *bool        `json:"encryptionEnabled,omitempty"`
	FileStorageCategory string       `json:"fileStorageCategory,omitempty"`
	FileStorageId       string       `json:"fileStorageId,omitempty"`
	FileStorageName     string       `json:"fileStorageName,omitempty"`
	FileStorageProtocol string       `json:"fileStorageProtocol,omitempty"`
	FileStoragePurpose  string       `json:"fileStoragePurpose,omitempty"`
	FileStorageState    string       `json:"fileStorageState,omitempty"`
	LinkedObjectCount   int64        `json:"linkedObjectCount,omitempty"`
	ProductGroupId      string       `json:"productGroupId,omitempty"`
	ServiceZoneId       string       `json:"serviceZoneId,omitempty"`
	TieringEnabled      *bool        `json:"tieringEnabled,omitempty"`
	CreatedBy           string       `json:"createdBy,omitempty"`
	CreatedDt           time.Time `json:"createdDt,omitempty"`
	ModifiedBy          string       `json:"modifiedBy,omitempty"`
	ModifiedDt          time.Time `json:"modifiedDt,omitempty"`
}

type FsSnapshotScheduleRequest struct {
	SnapshotRetentionCount *int32                            `json:"snapshotRetentionCount,omitempty"`
	SnapshotSchedule       *SnapshotSchedule `json:"snapshotSchedule,omitempty"`
}

type LinkFileStorageObjectV4Request struct {
	LinkObjects   []LinkObjectV4Request `json:"linkObjects,omitempty"`
	UnlinkObjects []LinkObjectV4Request `json:"unlinkObjects,omitempty"`
}

type LinkObjectV4Request struct {
	ObjectId string `json:"objectId"`
	Type_    string `json:"type"`
}

type LinkedObject struct {
	LinkedObjectId   string `json:"linkedObjectId,omitempty"`
	LinkedObjectType string `json:"linkedObjectType,omitempty"`
}

type ListResponseFileStorageRelationLocationResponse struct {
	Contents   []FileStorageRelationLocationResponse `json:"contents,omitempty"`
	TotalCount *int32                                                `json:"totalCount,omitempty"`
}

type ListResponseFileStorageSnapshotsResponse struct {
	Contents   []FileStorageSnapshotsResponse `json:"contents,omitempty"`
	TotalCount *int32                                         `json:"totalCount,omitempty"`
}

type ListResponseFileStoragesListResponse struct {
	Contents   []FileStoragesListResponse `json:"contents,omitempty"`
	TotalCount *int32                                     `json:"totalCount,omitempty"`
}

type OriginFileStorage struct {
	OrgFileStorageId            string `json:"orgFileStorageId,omitempty"`
	OrgFileStorageName          string `json:"orgFileStorageName,omitempty"`
	OrgFileStorageServiceZoneId string `json:"orgFileStorageServiceZoneId,omitempty"`
}

type RelationSchedule struct {
	Frequency string `json:"frequency,omitempty"`
}

type SnapshotSchedule struct {
	DayOfWeek string `json:"dayOfWeek,omitempty"`
	Frequency string `json:"frequency,omitempty"`
	Hour      *int32 `json:"hour,omitempty"`
}

type TagRequest struct {
	TagKey   string `json:"tagKey"`
	TagValue string `json:"tagValue,omitempty"`
}

type UpdateFileStorageTieringRequest struct {
	TieringLifeCycle *int32 `json:"tieringLifeCycle,omitempty"`
	TieringPolicy    string `json:"tieringPolicy"`
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
