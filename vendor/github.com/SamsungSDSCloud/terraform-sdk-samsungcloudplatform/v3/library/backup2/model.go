package backup2

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

type BackupAgentCandidateObjectsV1Response struct {
	ImageName           string `json:"imageName,omitempty"`
	IsAgentRegistered   *bool  `json:"isAgentRegistered,omitempty"`
	ObjectBackupAgentIp string `json:"objectBackupAgentIp,omitempty"`
	ObjectId            string `json:"objectId,omitempty"`
	ObjectName          string `json:"objectName,omitempty"`
}

type BackupAgentCreateV1Request struct {
	AzCode        string                  `json:"azCode,omitempty"`
	ObjectId      string                  `json:"objectId"`
	ObjectType    string                  `json:"objectType,omitempty"`
	ProductNames  []string                `json:"productNames"`
	ServiceZoneId string                  `json:"serviceZoneId"`
	Tags          []TagRequest `json:"tags,omitempty"`
}

type BackupAgentDetailResponse struct {
	ProjectId           string       `json:"projectId,omitempty"`
	AzCode              string       `json:"azCode,omitempty"`
	BackupAgentName     string       `json:"backupAgentName,omitempty"`
	BackupAgentState    string       `json:"backupAgentState,omitempty"`
	BackupMasterIp      string       `json:"backupMasterIp,omitempty"`
	BackupMasterName    string       `json:"backupMasterName,omitempty"`
	BackupPolicyCount   int64        `json:"backupPolicyCount,omitempty"`
	BlockId             string       `json:"blockId,omitempty"`
	ConnectionCheckDt   time.Time `json:"connectionCheckDt,omitempty"`
	ConnectionState     string       `json:"connectionState,omitempty"`
	ObjectBackupAgentIp string       `json:"objectBackupAgentIp,omitempty"`
	ObjectGatewayIp     string       `json:"objectGatewayIp,omitempty"`
	ObjectId            string       `json:"objectId,omitempty"`
	ObjectMacAddress    string       `json:"objectMacAddress,omitempty"`
	ObjectType          string       `json:"objectType,omitempty"`
	ServiceZoneId       string       `json:"serviceZoneId,omitempty"`
	CreatedBy           string       `json:"createdBy,omitempty"`
	CreatedDt           time.Time `json:"createdDt,omitempty"`
	ModifiedBy          string       `json:"modifiedBy,omitempty"`
	ModifiedDt          time.Time `json:"modifiedDt,omitempty"`
}

type BackupAgentInstallFilePathV1Response struct {
	AgentInstallFilePath string `json:"agentInstallFilePath,omitempty"`
}

type BackupAgentListResponse struct {
	ProjectId             string       `json:"projectId,omitempty"`
	BackupAgentId         string       `json:"backupAgentId,omitempty"`
	BackupAgentName       string       `json:"backupAgentName,omitempty"`
	BackupMasterId        string       `json:"backupMasterId,omitempty"`
	BlockId               string       `json:"blockId,omitempty"`
	ConnectionState       string       `json:"connectionState,omitempty"`
	ObjectBackupAgentIp   string       `json:"objectBackupAgentIp,omitempty"`
	ObjectBackupAgentIpId string       `json:"objectBackupAgentIpId,omitempty"`
	ObjectGatewayIp       string       `json:"objectGatewayIp,omitempty"`
	ObjectId              string       `json:"objectId,omitempty"`
	ObjectMacAddress      string       `json:"objectMacAddress,omitempty"`
	ObjectName            string       `json:"objectName,omitempty"`
	ObjectType            string       `json:"objectType,omitempty"`
	ProductGroupId        string       `json:"productGroupId,omitempty"`
	ServiceZoneId         string       `json:"serviceZoneId,omitempty"`
	State                 string       `json:"state,omitempty"`
	CreatedBy             string       `json:"createdBy,omitempty"`
	CreatedDt             time.Time `json:"createdDt,omitempty"`
	ModifiedBy            string       `json:"modifiedBy,omitempty"`
	ModifiedDt            time.Time `json:"modifiedDt,omitempty"`
}

type BackupAgentSearchOpenApiApiGetBackupAgentInstallFilePathOpts struct {
	AzCode optional.String
}

type BackupAgentSearchOpenApiApiListBackupAgentsOpts struct {
	BackupAgentId       optional.String
	BackupAgentName     optional.String
	ObjectBackupAgentIp optional.String
	ObjectId            optional.String
	ObjectIds           optional.Interface
	ObjectName          optional.String
	CreatedBy           optional.String
	Page                optional.Int32
	Size                optional.Int32
	Sort                optional.Interface
}

type BackupAgentSearchOpenApiApiListCandidateObjectsForBackupAgentOpts struct {
	AzCode     optional.String
	ObjectName optional.String
	Page       optional.Int32
	Size       optional.Int32
	Sort       optional.Interface
}

type BackupCdpCreateRequest struct {
	BackupPolicyTypeCategory string                  `json:"backupPolicyTypeCategory"`
	CdpScheduleFrequency     string                  `json:"cdpScheduleFrequency"`
	ObjectId                 string                  `json:"objectId"`
	ObjectType               string                  `json:"objectType"`
	PolicyType               string                  `json:"policyType"`
	ProductNames             []string                `json:"productNames"`
	RetentionPeriod          string                  `json:"retentionPeriod"`
	ServiceZoneId            string                  `json:"serviceZoneId"`
	Tags                     []TagRequest `json:"tags,omitempty"`
}

type BackupCdpScheduleUpdateRequest struct {
	CdpScheduleFrequency string `json:"cdpScheduleFrequency"`
	RetentionPeriod      string `json:"retentionPeriod"`
}

type BackupCreateV3Request struct {
	AzCode                     string                          `json:"azCode,omitempty"`
	BackupDrZoneId             string                          `json:"backupDrZoneId,omitempty"`
	BackupName                 string                          `json:"backupName"`
	BackupPolicyTypeCategory   string                          `json:"backupPolicyTypeCategory"`
	BackupRepository           string                          `json:"backupRepository,omitempty"`
	DrAzCode                   string                          `json:"drAzCode,omitempty"`
	FileSystemBackupSelections []string                        `json:"fileSystemBackupSelections,omitempty"`
	IsBackupDrEnabled          string                          `json:"isBackupDrEnabled,omitempty"`
	ObjectId                   string                          `json:"objectId"`
	ObjectType                 string                          `json:"objectType"`
	PolicyType                 string                          `json:"policyType"`
	ProductNames               []string                        `json:"productNames"`
	RetentionPeriod            string                          `json:"retentionPeriod"`
	Schedules                  []BackupScheduleInfo `json:"schedules"`
	ServiceZoneId              string                          `json:"serviceZoneId"`
	Tags                       []TagRequest         `json:"tags,omitempty"`
}

type BackupCreateV4Request struct {
	AzCode                     string                          `json:"azCode,omitempty"`
	BackupDrZoneId             string                          `json:"backupDrZoneId,omitempty"`
	BackupName                 string                          `json:"backupName"`
	BackupPolicyTypeCategory   string                          `json:"backupPolicyTypeCategory"`
	BackupRepository           string                          `json:"backupRepository,omitempty"`
	DrAzCode                   string                          `json:"drAzCode,omitempty"`
	FileSystemBackupSelections []string                        `json:"fileSystemBackupSelections,omitempty"`
	IsBackupDrEnabled          string                          `json:"isBackupDrEnabled,omitempty"`
	ObjectId                   string                          `json:"objectId"`
	ObjectType                 string                          `json:"objectType"`
	PolicyType                 string                          `json:"policyType"`
	ProductNames               []string                        `json:"productNames"`
	RetentionPeriod            string                          `json:"retentionPeriod"`
	Schedules                  []BackupScheduleInfo `json:"schedules"`
	ServiceZoneId              string                          `json:"serviceZoneId"`
	Tags                       []TagRequest         `json:"tags,omitempty"`
}

type BackupCreateV5Request struct {
	AzCode                     string                          `json:"azCode,omitempty"`
	BackupDrZoneId             string                          `json:"backupDrZoneId,omitempty"`
	BackupName                 string                          `json:"backupName"`
	BackupPolicyTypeCategory   string                          `json:"backupPolicyTypeCategory"`
	BackupRepository           string                          `json:"backupRepository,omitempty"`
	DrAzCode                   string                          `json:"drAzCode,omitempty"`
	FileSystemBackupSelections []string                        `json:"fileSystemBackupSelections,omitempty"`
	IncrementalRetentionPeriod string                          `json:"incrementalRetentionPeriod,omitempty"`
	IsBackupDrEnabled          string                          `json:"isBackupDrEnabled,omitempty"`
	ObjectId                   string                          `json:"objectId"`
	ObjectType                 string                          `json:"objectType"`
	PolicyType                 string                          `json:"policyType"`
	ProductNames               []string                        `json:"productNames"`
	RetentionPeriod            string                          `json:"retentionPeriod"`
	Schedules                  []BackupScheduleInfo `json:"schedules"`
	ServiceZoneId              string                          `json:"serviceZoneId"`
	Tags                       []TagRequest         `json:"tags,omitempty"`
}

type BackupFilesystemRestoreRequest struct {
	BackupAgentId  string   `json:"backupAgentId"`
	BackupImageId  string   `json:"backupImageId"`
	OriginPath     []string `json:"originPath"`
	OverWrite      *bool    `json:"overWrite"`
	ScheduleName   string   `json:"scheduleName"`
	ScheduleType   string   `json:"scheduleType"`
	StartTime      string   `json:"startTime"`
	TargetObjectId string   `json:"targetObjectId"`
	TargetPath     []string `json:"targetPath"`
}

type BackupImageDeleteRequest struct {
	BackupImages           []BackupImageInfo `json:"backupImages"`
	BackupImageDescription string                       `json:"backupImageDescription,omitempty"`
}

type BackupImageInfo struct {
	BackupImageId string `json:"backupImageId"`
	ScheduleName  string `json:"scheduleName"`
	StartTime     string `json:"startTime"`
}

type BackupImagesResponse struct {
	BackupImageId            string       `json:"backupImageId,omitempty"`
	EndTime                  time.Time `json:"endTime,omitempty"`
	ExpiredTime              time.Time `json:"expiredTime,omitempty"`
	IsRestoreAvailable       *bool        `json:"isRestoreAvailable,omitempty"`
	ScheduleName             string       `json:"scheduleName,omitempty"`
	ScheduleType             string       `json:"scheduleType,omitempty"`
	StartTime                time.Time `json:"startTime,omitempty"`
	TransferredDataKilobytes int64        `json:"transferredDataKilobytes,omitempty"`
	Usage                    string       `json:"usage,omitempty"`
}

type BackupJobHistoriesResponse struct {
	EndTime                  time.Time `json:"endTime,omitempty"`
	JobId                    int64        `json:"jobId,omitempty"`
	JobState                 string       `json:"jobState,omitempty"`
	ScheduleId               string       `json:"scheduleId,omitempty"`
	ScheduleName             string       `json:"scheduleName,omitempty"`
	ScheduleType             string       `json:"scheduleType,omitempty"`
	StartTime                time.Time `json:"startTime,omitempty"`
	TransferredDataKilobytes string       `json:"transferredDataKilobytes,omitempty"`
}

type BackupRestoreHistoriesResponse struct {
	BackupDrId               string       `json:"backupDrId,omitempty"`
	BackupId                 string       `json:"backupId,omitempty"`
	BackupTime               time.Time `json:"backupTime,omitempty"`
	IsBackupDrRestored       string       `json:"isBackupDrRestored,omitempty"`
	RestoreId                string       `json:"restoreId,omitempty"`
	RestoreState             string       `json:"restoreState,omitempty"`
	RestoreType              string       `json:"restoreType,omitempty"`
	RestoreVirtualServerId   string       `json:"restoreVirtualServerId,omitempty"`
	RestoreVirtualServerName string       `json:"restoreVirtualServerName,omitempty"`
	CreatedBy                string       `json:"createdBy,omitempty"`
	CreatedDt                time.Time `json:"createdDt,omitempty"`
}

type BackupRestoreRequest struct {
	BackupDrId          string   `json:"backupDrId,omitempty"`
	BackupImageId       string   `json:"backupImageId"`
	BackupRestoreName   string   `json:"backupRestoreName"`
	BackupTime          string   `json:"backupTime"`
	DnsEnabled          string   `json:"dnsEnabled,omitempty"`
	Ip                  string   `json:"ip,omitempty"`
	IpType              string   `json:"ipType,omitempty"`
	NatEnabled          string   `json:"natEnabled,omitempty"`
	ReservedIpAddressId string   `json:"reservedIpAddressId,omitempty"`
	ScheduleName        string   `json:"scheduleName"`
	ScheduleType        string   `json:"scheduleType"`
	SecurityGroupIds    []string `json:"securityGroupIds"`
	ServerGroupId       string   `json:"serverGroupId"`
	SubnetId            string   `json:"subnetId,omitempty"`
	VpcId               string   `json:"vpcId,omitempty"`
	RestoreDescription  string   `json:"restoreDescription,omitempty"`
}

type BackupScheduleInfo struct {
	ScheduleFrequency       string `json:"scheduleFrequency"`
	ScheduleFrequencyDetail string `json:"scheduleFrequencyDetail"`
	ScheduleId              string `json:"scheduleId,omitempty"`
	ScheduleName            string `json:"scheduleName,omitempty"`
	ScheduleType            string `json:"scheduleType"`
	StartTime               string `json:"startTime"`
}

type BackupScheduleUpdateRequest struct {
	Schedules           []BackupScheduleInfo `json:"schedules"`
	ScheduleDescription string                          `json:"scheduleDescription,omitempty"`
}

type BackupScheduleUpdateV3Request struct {
	IncrementalRetentionPeriod string                          `json:"incrementalRetentionPeriod,omitempty"`
	RetentionPeriod            string                          `json:"retentionPeriod"`
	Schedules                  []BackupScheduleInfo `json:"schedules"`
	ScheduleDescription        string                          `json:"scheduleDescription,omitempty"`
}

type BackupSchedulesResponse struct {
	BackupId                string       `json:"backupId,omitempty"`
	ScheduleFrequency       string       `json:"scheduleFrequency,omitempty"`
	ScheduleFrequencyDetail string       `json:"scheduleFrequencyDetail,omitempty"`
	ScheduleId              string       `json:"scheduleId,omitempty"`
	ScheduleName            string       `json:"scheduleName,omitempty"`
	ScheduleType            string       `json:"scheduleType,omitempty"`
	StartTime               string       `json:"startTime,omitempty"`
	CreatedBy               string       `json:"createdBy,omitempty"`
	CreatedDt               time.Time `json:"createdDt,omitempty"`
	ModifiedBy              string       `json:"modifiedBy,omitempty"`
	ModifiedDt              time.Time `json:"modifiedDt,omitempty"`
}

type BackupSearchOpenApiApiGetFilesystemRestoreCandidateObjectListOpts struct {
	AzCode     optional.String
	ObjectName optional.String
	Page       optional.Int32
	Size       optional.Int32
	Sort       optional.Interface
}

type BackupSearchOpenApiApiListBackupImagesOpts struct {
	EndDate      optional.String
	StartDate    optional.String
	TimeZoneName optional.String
	Page         optional.Int32
	Size         optional.Int32
	Sort         optional.Interface
}

type BackupSearchOpenApiApiListBackupsOpts struct {
	BackupName               optional.String
	BackupPolicyTypeCategory optional.String
	CreatedBy                optional.String
	Page                     optional.Int32
	Size                     optional.Int32
	Sort                     optional.Interface
}

type BackupSearchOpenApiApiListDrBackupImagesOpts struct {
	EndDate      optional.String
	StartDate    optional.String
	TimeZoneName optional.String
	Page         optional.Int32
	Size         optional.Int32
	Sort         optional.Interface
}

type BackupSearchOpenApiApiListDrRestoreHistoriesOpts struct {
	EndDate      optional.String
	StartDate    optional.String
	TimeZoneName optional.String
	Page         optional.Int32
	Size         optional.Int32
	Sort         optional.Interface
}

type BackupSearchOpenApiApiListJobHistoriesOpts struct {
	EndDate      optional.String
	StartDate    optional.String
	TimeZoneName optional.String
	Page         optional.Int32
	Size         optional.Int32
	Sort         optional.Interface
}

type BackupSearchOpenApiApiListRestoreHistoriesOpts struct {
	EndDate      optional.String
	StartDate    optional.String
	TimeZoneName optional.String
	Page         optional.Int32
	Size         optional.Int32
	Sort         optional.Interface
}

type BackupSearchOpenApiApiListSchedulesOpts struct {
	Page optional.Int32
	Size optional.Int32
	Sort optional.Interface
}

type BackupV3Response struct {
	ProjectId                string       `json:"projectId,omitempty"`
	BackupDrId               string       `json:"backupDrId,omitempty"`
	BackupId                 string       `json:"backupId,omitempty"`
	BackupName               string       `json:"backupName,omitempty"`
	BackupPolicyTypeCategory string       `json:"backupPolicyTypeCategory,omitempty"`
	BackupState              string       `json:"backupState,omitempty"`
	BlockId                  string       `json:"blockId,omitempty"`
	IsBackupDrEnabled        string       `json:"isBackupDrEnabled,omitempty"`
	IsBackupDrOrigin         string       `json:"isBackupDrOrigin,omitempty"`
	ObjectId                 string       `json:"objectId,omitempty"`
	ObjectType               string       `json:"objectType,omitempty"`
	PolicyType               string       `json:"policyType,omitempty"`
	Region                   string       `json:"region,omitempty"`
	RetentionPeriod          string       `json:"retentionPeriod,omitempty"`
	ServiceZoneId            string       `json:"serviceZoneId,omitempty"`
	ZoneName                 string       `json:"zoneName,omitempty"`
	CreatedBy                string       `json:"createdBy,omitempty"`
	CreatedDt                time.Time `json:"createdDt,omitempty"`
	ModifiedBy               string       `json:"modifiedBy,omitempty"`
	ModifiedDt               time.Time `json:"modifiedDt,omitempty"`
}

type BasicAuth struct {
	UserName string `json:"userName,omitempty"`
	Password string `json:"password,omitempty"`
}

type DetailBackupV4Response struct {
	ProjectId                  string       `json:"projectId,omitempty"`
	AzCode                     string       `json:"azCode,omitempty"`
	BackupAgentId              string       `json:"backupAgentId,omitempty"`
	BackupDrBlockId            string       `json:"backupDrBlockId,omitempty"`
	BackupDrId                 string       `json:"backupDrId,omitempty"`
	BackupDrRegion             string       `json:"backupDrRegion,omitempty"`
	BackupDrRetentionPeriod    string       `json:"backupDrRetentionPeriod,omitempty"`
	BackupDrState              string       `json:"backupDrState,omitempty"`
	BackupDrZone               string       `json:"backupDrZone,omitempty"`
	BackupDrZoneId             string       `json:"backupDrZoneId,omitempty"`
	BackupId                   string       `json:"backupId,omitempty"`
	BackupName                 string       `json:"backupName,omitempty"`
	BackupPolicyTypeCategory   string       `json:"backupPolicyTypeCategory,omitempty"`
	BackupRepository           string       `json:"backupRepository,omitempty"`
	BackupState                string       `json:"backupState,omitempty"`
	BlockId                    string       `json:"blockId,omitempty"`
	CdpScheduleFrequency       string       `json:"cdpScheduleFrequency,omitempty"`
	DrAzCode                   string       `json:"drAzCode,omitempty"`
	FileSystemBackupSelections []string     `json:"fileSystemBackupSelections,omitempty"`
	IncrementalRetentionPeriod string       `json:"incrementalRetentionPeriod,omitempty"`
	IsBackupDrDeleted          string       `json:"isBackupDrDeleted,omitempty"`
	IsBackupDrEnabled          string       `json:"isBackupDrEnabled,omitempty"`
	IsBackupDrOrigin           string       `json:"isBackupDrOrigin,omitempty"`
	ObjectId                   string       `json:"objectId,omitempty"`
	ObjectType                 string       `json:"objectType,omitempty"`
	PlanName                   string       `json:"planName,omitempty"`
	PolicyType                 string       `json:"policyType,omitempty"`
	ProductGroupId             string       `json:"productGroupId,omitempty"`
	RetentionPeriod            string       `json:"retentionPeriod,omitempty"`
	ScheduleIds                []string     `json:"scheduleIds,omitempty"`
	ServiceZoneId              string       `json:"serviceZoneId,omitempty"`
	CreatedBy                  string       `json:"createdBy,omitempty"`
	CreatedDt                  time.Time `json:"createdDt,omitempty"`
	ModifiedBy                 string       `json:"modifiedBy,omitempty"`
	ModifiedDt                 time.Time `json:"modifiedDt,omitempty"`
}

type ListResponseBackupAgentCandidateObjectsV1Response struct {
	Contents   []BackupAgentCandidateObjectsV1Response `json:"contents,omitempty"`
	TotalCount int32                                              `json:"totalCount,omitempty"`
}

type ListResponseBackupAgentListResponse struct {
	Contents   []BackupAgentListResponse `json:"contents,omitempty"`
	TotalCount int32                                `json:"totalCount,omitempty"`
}

type ListResponseBackupImagesResponse struct {
	Contents   []BackupImagesResponse `json:"contents,omitempty"`
	TotalCount int32                             `json:"totalCount,omitempty"`
}

type ListResponseBackupJobHistoriesResponse struct {
	Contents   []BackupJobHistoriesResponse `json:"contents,omitempty"`
	TotalCount int32                                   `json:"totalCount,omitempty"`
}

type ListResponseBackupRestoreHistoriesResponse struct {
	Contents   []BackupRestoreHistoriesResponse `json:"contents,omitempty"`
	TotalCount int32                                       `json:"totalCount,omitempty"`
}

type ListResponseBackupSchedulesResponse struct {
	Contents   []BackupSchedulesResponse `json:"contents,omitempty"`
	TotalCount int32                                `json:"totalCount,omitempty"`
}

type ListResponseBackupV3Response struct {
	Contents   []BackupV3Response `json:"contents,omitempty"`
	TotalCount int32                         `json:"totalCount,omitempty"`
}

type PolicyDupliCheckResponse struct {
	IsFileSystemBackupSelectionAvailable *bool `json:"isFileSystemBackupSelectionAvailable,omitempty"`
}

type TagRequest struct {
	TagKey   string `json:"tagKey"`
	TagValue string `json:"tagValue,omitempty"`
}

type UpdateAgentBackupFilesystemRequest struct {
	FileSystemBackupSelections []string `json:"fileSystemBackupSelections"`
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
