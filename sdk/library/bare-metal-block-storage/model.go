package baremetalblockstorage

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

type BlockStorageListOutSo struct {
	BlockStorageId         string                                                  `json:"blockStorageId,omitempty"`
	BlockStorageName       string                                                  `json:"blockStorageName,omitempty"`
	Capacity               string                                                  `json:"capacity,omitempty"`
	LinkedBareMetalServers []LinkBareMetalServerListOutSo `json:"linkedBareMetalServers,omitempty"`
	SnapshotCapacity       string                                                  `json:"snapshotCapacity,omitempty"`
	SnapshotRate           string                                                  `json:"snapshotRate,omitempty"`
	State                  string                                                  `json:"state,omitempty"`
	Type_                  string                                                  `json:"type,omitempty"`
}

type BlockStorageListResponse struct {
	BlockStorageId         string                                                  `json:"blockStorageId,omitempty"`
	BlockStorageName       string                                                  `json:"blockStorageName,omitempty"`
	Capacity               string                                                  `json:"capacity,omitempty"`
	LinkedBareMetalServers []LinkBareMetalServerListOutSo `json:"linkedBareMetalServers,omitempty"`
	SnapshotCapacity       string                                                  `json:"snapshotCapacity,omitempty"`
	State                  string                                                  `json:"state,omitempty"`
	Type_                  string                                                  `json:"type,omitempty"`
}

type BlockStorageServerResponse struct {
	ServerId   string `json:"serverId,omitempty"`
	ServerType string `json:"serverType,omitempty"`
}

type BmBlockStoragDrCreateRequest struct {
	BackupRetentionCount      *int32                                     `json:"backupRetentionCount,omitempty"`
	BareMetalBlockStorageName string                                     `json:"bareMetalBlockStorageName"`
	RelationSchedule          *RelationSchedule `json:"relationSchedule"`
	RelationType              string                                     `json:"relationType"`
	ServiceZoneId             string                                     `json:"serviceZoneId"`
}

type BmBlockStoragDrRelationPolicyRequest struct {
	RelationPolicyType string `json:"relationPolicyType"`
}

type BmBlockStoragDrRelationScheduleRequest struct {
	BackupRetentionCount *int32                                     `json:"backupRetentionCount,omitempty"`
	RelationPolicyType   string                                     `json:"relationPolicyType"`
	RelationSchedule     *RelationSchedule `json:"relationSchedule"`
}

type BmBlockStorageAttachRequest struct {
	BareMetalServerIds []string `json:"bareMetalServerIds"`
}

type BmBlockStorageControllerApiListBareMetalBlockStoragesOpts struct {
	BareMetalBlockStorageName optional.String
	BareMetalServerName       optional.String
	CreatedBy                 optional.String
	Page                      optional.Int32
	Size                      optional.Int32
	Sort                      optional.Interface
}

type BmBlockStorageCreateRequest struct {
	BareMetalBlockStorageName string                                     `json:"bareMetalBlockStorageName"`
	BareMetalBlockStorageSize *int32                                     `json:"bareMetalBlockStorageSize"`
	BareMetalServerIds        []string                                   `json:"bareMetalServerIds"`
	EncryptionEnabled         *bool                                      `json:"encryptionEnabled"`
	IsSnapshotPolicy          *bool                                      `json:"isSnapshotPolicy,omitempty"`
	ProductId                 string                                     `json:"productId"`
	ServiceZoneId             string                                     `json:"serviceZoneId"`
	SnapshotCapacityRate      *int32                                     `json:"snapshotCapacityRate,omitempty"`
	SnapshotSchedule          *SnapshotSchedule `json:"snapshotSchedule,omitempty"`
	Tags                      []TagRequest      `json:"tags,omitempty"`
}

type BmBlockStorageCreateV3Request struct {
	BareMetalBlockStorageName string                                     `json:"bareMetalBlockStorageName"`
	BareMetalBlockStorageSize *int32                                     `json:"bareMetalBlockStorageSize"`
	BareMetalBlockStorageType string                                     `json:"bareMetalBlockStorageType"`
	BareMetalServerIds        []string                                   `json:"bareMetalServerIds"`
	EncryptionEnabled         *bool                                      `json:"encryptionEnabled,omitempty"`
	ServiceZoneId             string                                     `json:"serviceZoneId"`
	SnapshotCapacityRate      *int32                                     `json:"snapshotCapacityRate,omitempty"`
	SnapshotPolicy            *bool                                      `json:"snapshotPolicy,omitempty"`
	SnapshotSchedule          *SnapshotSchedule `json:"snapshotSchedule,omitempty"`
	Tags                      []TagRequest      `json:"tags,omitempty"`
}

type BmBlockStorageCreateV4Request struct {
	BareMetalBlockStorageName string                                     `json:"bareMetalBlockStorageName"`
	BareMetalBlockStorageSize *int32                                     `json:"bareMetalBlockStorageSize"`
	BareMetalBlockStorageType string                                     `json:"bareMetalBlockStorageType"`
	BareMetalServerIds        []string                                   `json:"bareMetalServerIds"`
	EncryptionEnabled         *bool                                      `json:"encryptionEnabled,omitempty"`
	ServiceZoneId             string                                     `json:"serviceZoneId"`
	SnapshotCapacityRate      *int32                                     `json:"snapshotCapacityRate,omitempty"`
	SnapshotPolicy            *bool                                      `json:"snapshotPolicy,omitempty"`
	SnapshotSchedule          *SnapshotSchedule `json:"snapshotSchedule,omitempty"`
	Tags                      []TagRequest      `json:"tags,omitempty"`
}

type BmBlockStorageDetachRequest struct {
	BareMetalServerIds []string `json:"bareMetalServerIds"`
}

type BmBlockStorageDetailResponse struct {
	ProjectId                     string                                                         `json:"projectId,omitempty"`
	AzName                        string                                                         `json:"azName,omitempty"`
	BackupBareMetalBlockStorageId string                                                         `json:"backupBareMetalBlockStorageId,omitempty"`
	BareMetalBlockStorageId       string                                                         `json:"bareMetalBlockStorageId,omitempty"`
	BareMetalBlockStorageName     string                                                         `json:"bareMetalBlockStorageName,omitempty"`
	BareMetalBlockStoragePurpose  string                                                         `json:"bareMetalBlockStoragePurpose,omitempty"`
	BareMetalBlockStorageSize     *int32                                                         `json:"bareMetalBlockStorageSize,omitempty"`
	BareMetalBlockStorageState    string                                                         `json:"bareMetalBlockStorageState,omitempty"`
	BlockId                       string                                                         `json:"blockId,omitempty"`
	ConsistencyGroupId            string                                                         `json:"consistencyGroupId,omitempty"`
	ConsistencyGroupName          string                                                         `json:"consistencyGroupName,omitempty"`
	DrBareMetalBlockStorageId     string                                                         `json:"drBareMetalBlockStorageId,omitempty"`
	EncryptionEnabled             *bool                                                          `json:"encryptionEnabled,omitempty"`
	ErrorCheck                    *bool                                                          `json:"errorCheck,omitempty"`
	IscsiTargetIp                 []string                                                       `json:"iscsiTargetIp,omitempty"`
	OriginBareMetalBlockStorage   []OriginBareMetalBlockStorageResponse `json:"originBareMetalBlockStorage,omitempty"`
	ProductId                     string                                                         `json:"productId,omitempty"`
	Servers                       []BlockStorageServerResponse          `json:"servers,omitempty"`
	ServiceZoneId                 string                                                         `json:"serviceZoneId,omitempty"`
	CreatedBy                     string                                                         `json:"createdBy,omitempty"`
	CreatedDt                     time.Time                                                   `json:"createdDt,omitempty"`
	ModifiedBy                    string                                                         `json:"modifiedBy,omitempty"`
	ModifiedDt                    time.Time                                                   `json:"modifiedDt,omitempty"`
}

type BmBlockStorageDetailV3Response struct {
	ProjectId                     string                                                         `json:"projectId,omitempty"`
	BackupBareMetalBlockStorageId string                                                         `json:"backupBareMetalBlockStorageId,omitempty"`
	BareMetalBlockStorageId       string                                                         `json:"bareMetalBlockStorageId,omitempty"`
	BareMetalBlockStorageName     string                                                         `json:"bareMetalBlockStorageName,omitempty"`
	BareMetalBlockStoragePurpose  string                                                         `json:"bareMetalBlockStoragePurpose,omitempty"`
	BareMetalBlockStorageSize     *int32                                                         `json:"bareMetalBlockStorageSize,omitempty"`
	BareMetalBlockStorageState    string                                                         `json:"bareMetalBlockStorageState,omitempty"`
	BareMetalBlockStorageType     string                                                         `json:"bareMetalBlockStorageType,omitempty"`
	BlockId                       string                                                         `json:"blockId,omitempty"`
	DrBareMetalBlockStorageId     string                                                         `json:"drBareMetalBlockStorageId,omitempty"`
	EncryptionEnabled             *bool                                                          `json:"encryptionEnabled,omitempty"`
	ErrorCheck                    *bool                                                          `json:"errorCheck,omitempty"`
	IscsiTargetIp                 []string                                                       `json:"iscsiTargetIp,omitempty"`
	OriginBareMetalBlockStorage   []OriginBareMetalBlockStorageResponse `json:"originBareMetalBlockStorage,omitempty"`
	Servers                       []BlockStorageServerResponse          `json:"servers,omitempty"`
	ServiceZoneId                 string                                                         `json:"serviceZoneId,omitempty"`
	CreatedBy                     string                                                         `json:"createdBy,omitempty"`
	CreatedDt                     time.Time                                                   `json:"createdDt,omitempty"`
	ModifiedBy                    string                                                         `json:"modifiedBy,omitempty"`
	ModifiedDt                    time.Time                                                   `json:"modifiedDt,omitempty"`
}

type BmBlockStorageDrRegionResponse struct {
	BlockId       string `json:"blockId,omitempty"`
	ServiceZoneId string `json:"serviceZoneId,omitempty"`
}

type BmBlockStorageDrRelationResponse struct {
	ProjectId                 string       `json:"projectId,omitempty"`
	BackupRetentionCount      *int32       `json:"backupRetentionCount,omitempty"`
	BareMetalBlockStorageId   string       `json:"bareMetalBlockStorageId,omitempty"`
	BareMetalBlockStorageName string       `json:"bareMetalBlockStorageName,omitempty"`
	BlockId                   string       `json:"blockId,omitempty"`
	EncryptionEnabled         string       `json:"encryptionEnabled,omitempty"`
	RelationLagTime           string       `json:"relationLagTime,omitempty"`
	RelationPolicyType        string       `json:"relationPolicyType,omitempty"`
	RelationSchedule          string       `json:"relationSchedule,omitempty"`
	RelationState             string       `json:"relationState,omitempty"`
	RelationType              string       `json:"relationType,omitempty"`
	ServiceZoneId             string       `json:"serviceZoneId,omitempty"`
	CreatedBy                 string       `json:"createdBy,omitempty"`
	CreatedDt                 time.Time `json:"createdDt,omitempty"`
	ModifiedBy                string       `json:"modifiedBy,omitempty"`
	ModifiedDt                time.Time `json:"modifiedDt,omitempty"`
}

type BmBlockStorageResponse struct {
	BareMetalBlockStorageId      string       `json:"bareMetalBlockStorageId,omitempty"`
	BareMetalBlockStorageName    string       `json:"bareMetalBlockStorageName,omitempty"`
	BareMetalBlockStoragePurpose string       `json:"bareMetalBlockStoragePurpose,omitempty"`
	BareMetalBlockStorageSize    *int32       `json:"bareMetalBlockStorageSize,omitempty"`
	BareMetalBlockStorageState   string       `json:"bareMetalBlockStorageState,omitempty"`
	BareMetalBlockStorageTypeId  string       `json:"bareMetalBlockStorageTypeId,omitempty"`
	BareMetalServerIds           []string     `json:"bareMetalServerIds,omitempty"`
	BlockId                      string       `json:"blockId,omitempty"`
	EncryptionEnabled            *bool        `json:"encryptionEnabled,omitempty"`
	Location                     string       `json:"location,omitempty"`
	ServiceZoneId                string       `json:"serviceZoneId,omitempty"`
	CreatedBy                    string       `json:"createdBy,omitempty"`
	CreatedDt                    time.Time `json:"createdDt,omitempty"`
	ModifiedBy                   string       `json:"modifiedBy,omitempty"`
	ModifiedDt                   time.Time `json:"modifiedDt,omitempty"`
}

type BmBlockStorageSnapshotAttributeDeleteRequest struct {
	IsSnapshotPolicy string `json:"isSnapshotPolicy,omitempty"`
}

type BmBlockStorageSnapshotAttributeRequest struct {
	IsSnapshotPolicy     string `json:"isSnapshotPolicy,omitempty"`
	SnapshotCapacityRate *int32 `json:"snapshotCapacityRate,omitempty"`
}

type BmBlockStorageSnapshotCreateResponse struct {
	ProjectId               string       `json:"projectId,omitempty"`
	BareMetalBlockStorageId string       `json:"bareMetalBlockStorageId,omitempty"`
	ServiceZoneId           string       `json:"serviceZoneId,omitempty"`
	SnapshotId              string       `json:"snapshotId,omitempty"`
	CreatedBy               string       `json:"createdBy,omitempty"`
	CreatedDt               time.Time `json:"createdDt,omitempty"`
	ModifiedBy              string       `json:"modifiedBy,omitempty"`
	ModifiedDt              time.Time `json:"modifiedDt,omitempty"`
}

type BmBlockStorageSnapshotDeleteResponse struct {
	ProjectId               string       `json:"projectId,omitempty"`
	BareMetalBlockStorageId string       `json:"bareMetalBlockStorageId,omitempty"`
	ServiceZoneId           string       `json:"serviceZoneId,omitempty"`
	SnapshotId              string       `json:"snapshotId,omitempty"`
	CreatedBy               string       `json:"createdBy,omitempty"`
	CreatedDt               time.Time `json:"createdDt,omitempty"`
	ModifiedBy              string       `json:"modifiedBy,omitempty"`
	ModifiedDt              time.Time `json:"modifiedDt,omitempty"`
}

type BmBlockStorageSnapshotItemResponse struct {
	SnapshotId     string       `json:"snapshotId,omitempty"`
	SnapshotName   string       `json:"snapshotName,omitempty"`
	SnapshotSizeMb float64      `json:"snapshotSizeMb,omitempty"`
	CreatedDt      time.Time `json:"createdDt,omitempty"`
}

type BmBlockStorageSnapshotRestoreResponse struct {
	ProjectId               string       `json:"projectId,omitempty"`
	BareMetalBlockStorageId string       `json:"bareMetalBlockStorageId,omitempty"`
	ServiceZoneId           string       `json:"serviceZoneId,omitempty"`
	SnapshotId              string       `json:"snapshotId,omitempty"`
	CreatedBy               string       `json:"createdBy,omitempty"`
	CreatedDt               time.Time `json:"createdDt,omitempty"`
	ModifiedBy              string       `json:"modifiedBy,omitempty"`
	ModifiedDt              time.Time `json:"modifiedDt,omitempty"`
}

type BmBlockStorageSnapshotScheduleRequest struct {
	SnapshotSchedule *SnapshotSchedule `json:"snapshotSchedule"`
}

type BmBlockStorageSnapshotScheduleResponse struct {
	ProjectId               string                                     `json:"projectId,omitempty"`
	BareMetalBlockStorageId string                                     `json:"bareMetalBlockStorageId,omitempty"`
	ServiceZoneId           string                                     `json:"serviceZoneId,omitempty"`
	SnapshotSchedule        *SnapshotSchedule `json:"snapshotSchedule,omitempty"`
	SnapshotScheduleId      string                                     `json:"snapshotScheduleId,omitempty"`
	CreatedBy               string                                     `json:"createdBy,omitempty"`
	CreatedDt               time.Time                               `json:"createdDt,omitempty"`
	ModifiedBy              string                                     `json:"modifiedBy,omitempty"`
	ModifiedDt              time.Time                               `json:"modifiedDt,omitempty"`
}

type BmBlockStorageSnapshotsAttributeResponse struct {
	ProjectId               string       `json:"projectId,omitempty"`
	BareMetalBlockStorageId string       `json:"bareMetalBlockStorageId,omitempty"`
	IsSnapshotPolicy        string       `json:"isSnapshotPolicy,omitempty"`
	ServiceZoneId           string       `json:"serviceZoneId,omitempty"`
	SnapshotCapacityRate    *int32       `json:"snapshotCapacityRate,omitempty"`
	CreatedBy               string       `json:"createdBy,omitempty"`
	CreatedDt               time.Time `json:"createdDt,omitempty"`
	ModifiedBy              string       `json:"modifiedBy,omitempty"`
	ModifiedDt              time.Time `json:"modifiedDt,omitempty"`
}

type BmBlockStorageSnapshotsResponse struct {
	ProjectId               string                                                        `json:"projectId,omitempty"`
	BareMetalBlockStorageId string                                                        `json:"bareMetalBlockStorageId,omitempty"`
	IsSnapshotPolicy        *bool                                                         `json:"isSnapshotPolicy,omitempty"`
	ServiceZoneId           string                                                        `json:"serviceZoneId,omitempty"`
	SnapshotCapacityMb      float64                                                       `json:"snapshotCapacityMb,omitempty"`
	SnapshotCapacityRate    *int32                                                        `json:"snapshotCapacityRate,omitempty"`
	SnapshotTotalUsage      float64                                                       `json:"snapshotTotalUsage,omitempty"`
	Snapshots               []BmBlockStorageSnapshotItemResponse `json:"snapshots,omitempty"`
	CreatedBy               string                                                        `json:"createdBy,omitempty"`
	CreatedDt               time.Time                                                  `json:"createdDt,omitempty"`
	ModifiedBy              string                                                        `json:"modifiedBy,omitempty"`
	ModifiedDt              time.Time                                                  `json:"modifiedDt,omitempty"`
}

type BmBlockStorageTerminatedSoBlockStorageResponse struct {
	ProjectId      string `json:"projectId,omitempty"`
	BlockStorageId string `json:"blockStorageId,omitempty"`
	RequestId      string `json:"requestId,omitempty"`
}

type BmBlockStorageV3ControllerApiListBareMetalBlockStoragesV3Opts struct {
	BareMetalBlockStorageName optional.String
	BareMetalServerName       optional.String
	CreatedBy                 optional.String
	Page                      optional.Int32
	Size                      optional.Int32
	Sort                      optional.Interface
}

type BmBlockStorageV3Response struct {
	BareMetalBlockStorageId      string       `json:"bareMetalBlockStorageId,omitempty"`
	BareMetalBlockStorageName    string       `json:"bareMetalBlockStorageName,omitempty"`
	BareMetalBlockStoragePurpose string       `json:"bareMetalBlockStoragePurpose,omitempty"`
	BareMetalBlockStorageSize    *int32       `json:"bareMetalBlockStorageSize,omitempty"`
	BareMetalBlockStorageState   string       `json:"bareMetalBlockStorageState,omitempty"`
	BareMetalBlockStorageType    string       `json:"bareMetalBlockStorageType,omitempty"`
	BareMetalServerIds           []string     `json:"bareMetalServerIds,omitempty"`
	BlockId                      string       `json:"blockId,omitempty"`
	EncryptionEnabled            *bool        `json:"encryptionEnabled,omitempty"`
	Location                     string       `json:"location,omitempty"`
	ServiceZoneId                string       `json:"serviceZoneId,omitempty"`
	CreatedBy                    string       `json:"createdBy,omitempty"`
	CreatedDt                    time.Time `json:"createdDt,omitempty"`
	ModifiedBy                   string       `json:"modifiedBy,omitempty"`
	ModifiedDt                   time.Time `json:"modifiedDt,omitempty"`
}

type BmCloneBlockStorageCreateRequest struct {
	CloneBlockStorageName string `json:"cloneBlockStorageName"`
}

type CloneConsistencyGroup struct {
	AzName                  string                                           `json:"azName,omitempty"`
	ConsistencyGroupId      string                                           `json:"consistencyGroupId,omitempty"`
	ConsistencyGroupName    string                                           `json:"consistencyGroupName,omitempty"`
	ConsistencyGroupPurpose string                                           `json:"consistencyGroupPurpose,omitempty"`
	MemberBlockStorageList  []BlockStorageListOutSo `json:"memberBlockStorageList,omitempty"`
	ZoneId                  string                                           `json:"zoneId,omitempty"`
}

type CloneConsistencyGroupCreateRequest struct {
	BlockStorageNamePrefix string `json:"blockStorageNamePrefix"`
	ConsistencyGroupName   string `json:"consistencyGroupName"`
}

type ConsistencyGroupControllerApiListConsistencyGroupOpts struct {
	AzName               optional.String
	ConsistencyGroupName optional.String
	Location             optional.String
	CreatedBy            optional.String
	Page                 optional.Int32
	Size                 optional.Int32
	Sort                 optional.Interface
}

type ConsistencyGroupCreateRequest struct {
	AzName                      string                                `json:"azName,omitempty"`
	BareMetalBlockStorageId     string                                `json:"bareMetalBlockStorageId,omitempty"`
	ConsistencyGroupName        string                                `json:"consistencyGroupName,omitempty"`
	MultiAvailabilityZone       *bool                                 `json:"multiAvailabilityZone,omitempty"`
	ServiceZoneId               string                                `json:"serviceZoneId,omitempty"`
	Tags                        []TagRequest `json:"tags,omitempty"`
	ConsistencyGroupDescription string                                `json:"consistencyGroupDescription,omitempty"`
}

type ConsistencyGroupDetailResponse struct {
	ProjectId                 string                                              `json:"projectId,omitempty"`
	AzName                    string                                              `json:"azName,omitempty"`
	BackupConsistencyGroupId  string                                              `json:"backupConsistencyGroupId,omitempty"`
	BlockId                   string                                              `json:"blockId,omitempty"`
	CloneConsistencyGroupList []CloneConsistencyGroup    `json:"cloneConsistencyGroupList,omitempty"`
	ConsistencyGroupId        string                                              `json:"consistencyGroupId,omitempty"`
	ConsistencyGroupName      string                                              `json:"consistencyGroupName,omitempty"`
	ConsistencyGroupPurpose   string                                              `json:"consistencyGroupPurpose,omitempty"`
	DrConsistencyGroupId      string                                              `json:"drConsistencyGroupId,omitempty"`
	HasClone                  *bool                                               `json:"hasClone,omitempty"`
	MemberBlockStorageList    []BlockStorageListResponse `json:"memberBlockStorageList,omitempty"`
	OrgConsistencyGroupId     string                                              `json:"orgConsistencyGroupId,omitempty"`
	OrgConsistencyGroupName   string                                              `json:"orgConsistencyGroupName,omitempty"`
	OrgConsistencyGroupZoneId string                                              `json:"orgConsistencyGroupZoneId,omitempty"`
	ServiceZoneId             string                                              `json:"serviceZoneId,omitempty"`
	State                     string                                              `json:"state,omitempty"`
	ZoneId                    string                                              `json:"zoneId,omitempty"`
	CreatedBy                 string                                              `json:"createdBy,omitempty"`
	CreatedDt                 time.Time                                        `json:"createdDt,omitempty"`
	ModifiedBy                string                                              `json:"modifiedBy,omitempty"`
	ModifiedDt                time.Time                                        `json:"modifiedDt,omitempty"`
}

type ConsistencyGroupDrCreateRequest struct {
	AzName               string                                     `json:"azName,omitempty"`
	BackupRetentionCount *int32                                     `json:"backupRetentionCount,omitempty"`
	BlockStoragePrefix   string                                     `json:"blockStoragePrefix"`
	ConsistencyGroupName string                                     `json:"consistencyGroupName"`
	RelationSchedule     *RelationSchedule `json:"relationSchedule"`
	RelationType         string                                     `json:"relationType"`
	ServiceZoneId        string                                     `json:"serviceZoneId"`
}

type ConsistencyGroupDrCreateResponse struct {
	ConsistencyGroupId       string `json:"consistencyGroupId,omitempty"`
	ConsistencyGroupName     string `json:"consistencyGroupName,omitempty"`
	OriginConsistencyGroupId string `json:"originConsistencyGroupId,omitempty"`
	Purpose                  string `json:"purpose,omitempty"`
	State                    string `json:"state,omitempty"`
}

type ConsistencyGroupDrDetailResponse struct {
	ProjectId               string `json:"projectId,omitempty"`
	BackupRetentionCount    *int32 `json:"backupRetentionCount,omitempty"`
	ConsistencyGroupId      string `json:"consistencyGroupId,omitempty"`
	ConsistencyGroupName    string `json:"consistencyGroupName,omitempty"`
	ConsistencyGroupPurpose string `json:"consistencyGroupPurpose,omitempty"`
	RelationLagTime         string `json:"relationLagTime,omitempty"`
	RelationPolicyType      string `json:"relationPolicyType,omitempty"`
	RelationSchedule        string `json:"relationSchedule,omitempty"`
	RelationState           string `json:"relationState,omitempty"`
	ZoneId                  string `json:"zoneId,omitempty"`
}

type ConsistencyGroupDrRelationPolicyRequest struct {
	RelationPolicyType string `json:"relationPolicyType"`
}

type ConsistencyGroupDrRelationScheduleRequest struct {
	BackupRetentionCount *int32 `json:"backupRetentionCount,omitempty"`
	RelationSchedule     string `json:"relationSchedule,omitempty"`
}

type ConsistencyGroupListResponse struct {
	ConsistencyGroupId   string `json:"consistencyGroupId,omitempty"`
	ConsistencyGroupName string `json:"consistencyGroupName,omitempty"`
	Location             string `json:"location,omitempty"`
	NumOfBlockStorage    *int32 `json:"numOfBlockStorage,omitempty"`
	Purpose              string `json:"purpose,omitempty"`
	State                string `json:"state,omitempty"`
}

type ConsistencyGroupSnapshotCreateResponse struct {
	ConsistencyGroupId string `json:"consistencyGroupId,omitempty"`
	SnapshotId         string `json:"snapshotId,omitempty"`
	SnapshotName       string `json:"snapshotName,omitempty"`
}

type ConsistencyGroupSnapshotDeleteResponse struct {
	ConsistencyGroupId string `json:"consistencyGroupId,omitempty"`
	SnapshotId         string `json:"snapshotId,omitempty"`
}

type ConsistencyGroupSnapshotListResponse struct {
	ConsistencyGroupId string                                              `json:"consistencyGroupId,omitempty"`
	SnapshotList       []SoSnapshotDetailResponse `json:"snapshotList,omitempty"`
}

type ConsistencyGroupSnapshotScheduleDetailResponse struct {
	ConsistencyGroupId string                                     `json:"consistencyGroupId,omitempty"`
	SnapshotSchedule   *SnapshotSchedule `json:"snapshotSchedule,omitempty"`
	SnapshotScheduleId string                                     `json:"snapshotScheduleId,omitempty"`
}

type ConsistencyGroupSnapshotScheduleRequest struct {
	SnapshotSchedule *SnapshotSchedule `json:"snapshotSchedule,omitempty"`
}

type ConsistencyGroupVolumeAddRemoveRequest struct {
	BlockStorageIdList []string `json:"blockStorageIdList,omitempty"`
	Mode               string   `json:"mode,omitempty"`
}

type LinkBareMetalServerListOutSo struct {
	Id          string `json:"id,omitempty"`
	ImageName   string `json:"imageName,omitempty"`
	MountState  string `json:"mountState,omitempty"`
	Name        string `json:"name,omitempty"`
	OsType      string `json:"osType,omitempty"`
	ProductType string `json:"productType,omitempty"`
	State       string `json:"state,omitempty"`
}

type ListResponseBmBlockStorageDrRegionResponse struct {
	Contents   []BmBlockStorageDrRegionResponse `json:"contents,omitempty"`
	TotalCount *int32                                                    `json:"totalCount,omitempty"`
}

type ListResponseBmBlockStorageResponse struct {
	Contents   []BmBlockStorageResponse `json:"contents,omitempty"`
	TotalCount *int32                                            `json:"totalCount,omitempty"`
}

type ListResponseBmBlockStorageSnapshotScheduleResponse struct {
	Contents   []BmBlockStorageSnapshotScheduleResponse `json:"contents,omitempty"`
	TotalCount *int32                                                            `json:"totalCount,omitempty"`
}

type ListResponseBmBlockStorageSnapshotsResponse struct {
	Contents   []BmBlockStorageSnapshotsResponse `json:"contents,omitempty"`
	TotalCount *int32                                                     `json:"totalCount,omitempty"`
}

type ListResponseBmBlockStorageV3Response struct {
	Contents   []BmBlockStorageV3Response `json:"contents,omitempty"`
	TotalCount *int32                                              `json:"totalCount,omitempty"`
}

type ListResponseConsistencyGroupListResponse struct {
	Contents   []ConsistencyGroupListResponse `json:"contents,omitempty"`
	TotalCount *int32                                                  `json:"totalCount,omitempty"`
}

type OriginBareMetalBlockStorageResponse struct {
	AzName                    string `json:"azName,omitempty"`
	BareMetalBlockStorageId   string `json:"bareMetalBlockStorageId,omitempty"`
	BareMetalBlockStorageName string `json:"bareMetalBlockStorageName,omitempty"`
	BlockId                   string `json:"blockId,omitempty"`
	ServiceZoneId             string `json:"serviceZoneId,omitempty"`
}

type RelationSchedule struct {
	Frequency string `json:"frequency,omitempty"`
}

type SnapshotSchedule struct {
	DayOfWeek string `json:"dayOfWeek,omitempty"`
	Frequency string `json:"frequency,omitempty"`
	Hour      *int32 `json:"hour,omitempty"`
}

type SoSnapshotDetailResponse struct {
	ExpireDate   string       `json:"expireDate,omitempty"`
	Size         int64        `json:"size,omitempty"`
	SnapshotId   string       `json:"snapshotId,omitempty"`
	SnapshotName string       `json:"snapshotName,omitempty"`
	SnapshotTime time.Time `json:"snapshotTime,omitempty"`
}

type TagRequest struct {
	TagKey   string `json:"tagKey"`
	TagValue string `json:"tagValue,omitempty"`
}

type ValidateBlockStorageForConsistencyGroupMemberRequest struct {
	BaremetalBlockStorageIdList []string `json:"baremetalBlockStorageIdList,omitempty"`
}

type ValidateBlockStorageForConsistencyGroupMemberResponse struct {
	AllResult           []map[string]interface{} `json:"allResult,omitempty"`
	ComprehensiveResult *bool                    `json:"comprehensiveResult,omitempty"`
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
