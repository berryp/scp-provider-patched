package sqlserver

import (
	"github.com/antihax/optional"
	"time"
)

type APIKey struct {
	Key    string
	Prefix string
}

type ActiveDirectory struct {
	AdServerId          string   `json:"adServerId,omitempty"`
	AdServerPassword    string   `json:"adServerPassword,omitempty"`
	DnsServerIpList     []string `json:"dnsServerIpList,omitempty"`
	DomainName          string   `json:"domainName,omitempty"`
	DomainNetBiosName   string   `json:"domainNetBiosName,omitempty"`
	FailoverClusterName string   `json:"failoverClusterName,omitempty"`
}

type ApplyParametersRequest struct {
	DbParameters []DatabaseParameter `json:"dbParameters"`
}

type AsyncResponse struct {
	ProjectId  string `json:"projectId,omitempty"`
	RequestId  string `json:"requestId,omitempty"`
	ResourceId string `json:"resourceId,omitempty"`
}

type AttachSecurityGroupsRequest struct {
	SecurityGroupIds []string `json:"securityGroupIds"`
}

type AvailabilityZoneConfigResponse struct {
	AvailabilityZoneDeploymentType string `json:"availabilityZoneDeploymentType,omitempty"`
}

type BasicAuth struct {
	UserName string `json:"userName,omitempty"`
	Password string `json:"password,omitempty"`
}

type ContractResponseV2 struct {
	BillingResourceId   string `json:"billingResourceId,omitempty"`
	ContractEndDate     string `json:"contractEndDate,omitempty"`
	ContractStartDate   string `json:"contractStartDate,omitempty"`
	ContractType        string `json:"contractType,omitempty"`
	NextContractEndDate string `json:"nextContractEndDate,omitempty"`
	NextContractType    string `json:"nextContractType,omitempty"`
	ResourceId          string `json:"resourceId,omitempty"`
}

type CreateReplicaRequest struct {
	ActiveDirectory  *ActiveDirectory `json:"activeDirectory,omitempty"`
	Network          *DatabaseNetwork `json:"network,omitempty"`
	OtherRegion      *bool                         `json:"otherRegion,omitempty"`
	Replica          *ReplicaSpec     `json:"replica"`
	SecurityGroupIds []string                      `json:"securityGroupIds,omitempty"`
	ServiceZoneId    string                        `json:"serviceZoneId,omitempty"`
}

type CreateSqlServerRequest struct {
	ActiveDirectory         *ActiveDirectory      `json:"activeDirectory,omitempty"`
	AdditionalDb            []string                           `json:"additionalDb,omitempty"`
	Backup                  *DatabaseBackup       `json:"backup,omitempty"`
	BlockId                 string                             `json:"blockId"`
	DbDiskInfos             []SqlserverDbDiskInfo `json:"dbDiskInfos,omitempty"`
	DbName                  string                             `json:"dbName"`
	DbPort                  int32                              `json:"dbPort,omitempty"`
	DbServiceName           string                             `json:"dbServiceName"`
	DbSoftwareOptions       map[string]string                  `json:"dbSoftwareOptions,omitempty"`
	DbUserId                string                             `json:"dbUserId"`
	DbUserPassword          string                             `json:"dbUserPassword"`
	DeploymentEnvType       string                             `json:"deploymentEnvType"`
	HighAvailability        *HaNormal             `json:"highAvailability,omitempty"`
	ImageId                 string                             `json:"imageId"`
	LicenseKey              string                             `json:"licenseKey"`
	Maintenance             *DatabaseMaintenance  `json:"maintenance,omitempty"`
	Network                 *DatabaseNetwork      `json:"network"`
	ProductGroupId          string                             `json:"productGroupId"`
	Replica                 *ReplicaSpec          `json:"replica,omitempty"`
	SecurityGroupIds        []string                           `json:"securityGroupIds"`
	ServerGroupName         string                             `json:"serverGroupName"`
	ServiceZoneId           string                             `json:"serviceZoneId"`
	Tags                    []TagRequest          `json:"tags,omitempty"`
	Timezone                string                             `json:"timezone"`
	UseDbLoggingAudit       *bool                              `json:"useDbLoggingAudit"`
	VirtualServer           *InstanceSpec         `json:"virtualServer"`
	VirtualServerNamePrefix string                             `json:"virtualServerNamePrefix"`
}

type DatabaseBackup struct {
	BackupMethod                string `json:"backupMethod,omitempty"`
	BackupRetentionDay          int32  `json:"backupRetentionDay"`
	BackupScheduleFrequencyHour int32  `json:"backupScheduleFrequencyHour"`
	BackupStartHour             int32  `json:"backupStartHour"`
	DbBackupArchMin             int32  `json:"dbBackupArchMin,omitempty"`
	FullBackupDay               string `json:"fullBackupDay,omitempty"`
	ObjectStorageId             string `json:"objectStorageId"`
}

type DatabaseBackupApiListBackupsOpts struct {
	ServerGroupName optional.String
	Page            optional.Int32
	Size            optional.Int32
	Sort            optional.String
}

type DatabaseBackupsResponse struct {
	BackupEndDt       time.Time `json:"backupEndDt,omitempty"`
	BackupHistoryId   string       `json:"backupHistoryId,omitempty"`
	BackupSize        int32        `json:"backupSize,omitempty"`
	BackupStartDt     time.Time `json:"backupStartDt,omitempty"`
	BackupState       string       `json:"backupState,omitempty"`
	BackupStateDetail string       `json:"backupStateDetail,omitempty"`
	ServerGroupId     string       `json:"serverGroupId,omitempty"`
	SoftwareVersion   string       `json:"softwareVersion,omitempty"`
	CreatedBy         string       `json:"createdBy,omitempty"`
	CreatedDt         time.Time `json:"createdDt,omitempty"`
	ModifiedBy        string       `json:"modifiedBy,omitempty"`
	ModifiedDt        time.Time `json:"modifiedDt,omitempty"`
}

type DatabaseBlockStorage struct {
	BlockStorageSize int32  `json:"blockStorageSize"`
	BlockStorageType string `json:"blockStorageType"`
	DiskType         string `json:"diskType"`
	DriveLetter      string `json:"driveLetter,omitempty"`
}

type DatabaseEventControllerApiListEventsOpts struct {
	EventCategories  optional.Interface
	EventContent     optional.String
	EventEndDt       optional.Time
	EventLevels      optional.Interface
	EventRequestedBy optional.String
	EventStartDt     optional.Time
	Page             optional.Int32
	Size             optional.Int32
	Sort             optional.Interface
}

type DatabaseListResponseServiceParamItem struct {
	Contents   []ServiceParamItem `json:"contents,omitempty"`
	TotalCount int32                           `json:"totalCount,omitempty"`
	ModifiedDt time.Time                    `json:"modifiedDt,omitempty"`
}

type DatabaseMaintenance struct {
	MaintenanceStartDayOfWeek  string  `json:"maintenanceStartDayOfWeek"`
	MaintenanceStartHourMinute string  `json:"maintenanceStartHourMinute"`
	MaintenanceTermHour        float64 `json:"maintenanceTermHour"`
}

type DatabaseManagementApiSwitchOverOpts struct {
	InstanceId optional.String
}

type DatabaseNetwork struct {
	NetworkEnvType string                               `json:"networkEnvType"`
	ServerNetworks []DatabaseServerNetwork `json:"serverNetworks"`
	SubnetId       string                               `json:"subnetId"`
	UseNat         *bool                                `json:"useNat"`
	VpcId          string                               `json:"vpcId"`
}

type DatabaseParameter struct {
	DbParameterId     string `json:"dbParameterId"`
	NewParameterValue string `json:"newParameterValue"`
	OldParameterValue string `json:"oldParameterValue"`
}

type DatabaseParameterApiGetParamByParamGroupOpts struct {
	Page optional.Int32
	Size optional.Int32
	Sort optional.String
}

type DatabaseParameterGroupResponse struct {
	Contents             []DefaultParamItem `json:"contents,omitempty"`
	DbParameterGroupId   string                          `json:"dbParameterGroupId,omitempty"`
	DbParameterGroupName string                          `json:"dbParameterGroupName,omitempty"`
	TotalCount           int32                           `json:"totalCount,omitempty"`
}

type DatabaseParameterGroupsResponse struct {
	DbParameterGroupId   string `json:"dbParameterGroupId,omitempty"`
	DbParameterGroupName string `json:"dbParameterGroupName,omitempty"`
}

type DatabaseSecurityGroupResponse struct {
	SecurityGroupId          string `json:"securityGroupId,omitempty"`
	SecurityGroupName        string `json:"securityGroupName,omitempty"`
	SecurityGroupState       string `json:"securityGroupState,omitempty"`
	SecurityGroupDescription string `json:"securityGroupDescription,omitempty"`
}

type DatabaseServerNetwork struct {
	AvailabilityZoneName string `json:"availabilityZoneName,omitempty"`
	NatIpId              string `json:"natIpId,omitempty"`
	NodeType             string `json:"nodeType"`
	ServerName           string `json:"serverName"`
	ServiceIp            string `json:"serviceIp,omitempty"`
}

type DatabaseStroagesResponse struct {
	BlockStorageDiskType string       `json:"blockStorageDiskType,omitempty"`
	BlockStorageId       string       `json:"blockStorageId,omitempty"`
	BlockStorageName     string       `json:"blockStorageName,omitempty"`
	BlockStorageSize     int32        `json:"blockStorageSize,omitempty"`
	BlockStorageType     string       `json:"blockStorageType,omitempty"`
	DriveLetter          string       `json:"driveLetter,omitempty"`
	ServerGroupId        string       `json:"serverGroupId,omitempty"`
	VirtualServerId      string       `json:"virtualServerId,omitempty"`
	CreatedBy            string       `json:"createdBy,omitempty"`
	CreatedDt            time.Time `json:"createdDt,omitempty"`
	ModifiedBy           string       `json:"modifiedBy,omitempty"`
	ModifiedDt           time.Time `json:"modifiedDt,omitempty"`
}

type DatabaseUserApiGetUserInfoOpts struct {
	Distinct optional.String
}

type DatabaseUserResponse struct {
	DbLoginId   string       `json:"dbLoginId,omitempty"`
	DbLoginType string       `json:"dbLoginType,omitempty"`
	DbName      string       `json:"dbName,omitempty"`
	Name        string       `json:"name,omitempty"`
	Role        string       `json:"role,omitempty"`
	ServiceId   string       `json:"serviceId,omitempty"`
	State       string       `json:"state,omitempty"`
	Description string       `json:"description,omitempty"`
	CreatedBy   string       `json:"createdBy,omitempty"`
	CreatedDt   time.Time `json:"createdDt,omitempty"`
	ModifiedBy  string       `json:"modifiedBy,omitempty"`
	ModifiedDt  time.Time `json:"modifiedDt,omitempty"`
}

type DbBackupsDomain struct {
	BackupDir       string       `json:"backupDir,omitempty"`
	BackupEndTime   time.Time `json:"backupEndTime,omitempty"`
	BackupHistNo    string       `json:"backupHistNo,omitempty"`
	BackupServerIp  string       `json:"backupServerIp,omitempty"`
	BackupSize      int32        `json:"backupSize,omitempty"`
	BackupStartTime time.Time `json:"backupStartTime,omitempty"`
	ServiceId       string       `json:"serviceId,omitempty"`
	SettingHistNo   int32        `json:"settingHistNo,omitempty"`
	State           string       `json:"state,omitempty"`
	StateText       string       `json:"stateText,omitempty"`
	Version         string       `json:"version,omitempty"`
	CreatedBy       string       `json:"createdBy,omitempty"`
	CreatedDt       time.Time `json:"createdDt,omitempty"`
	ModifiedBy      string       `json:"modifiedBy,omitempty"`
	ModifiedDt      time.Time `json:"modifiedDt,omitempty"`
}

type DbBackupsRequest struct {
	BackupDir       string       `json:"backupDir,omitempty"`
	BackupEndTime   time.Time `json:"backupEndTime,omitempty"`
	BackupHistNo    string       `json:"backupHistNo,omitempty"`
	BackupServerIp  string       `json:"backupServerIp,omitempty"`
	BackupSize      int32        `json:"backupSize,omitempty"`
	BackupStartTime time.Time `json:"backupStartTime,omitempty"`
	ServiceId       string       `json:"serviceId,omitempty"`
	SettingHistNo   int32        `json:"settingHistNo,omitempty"`
	State           string       `json:"state,omitempty"`
	StateText       string       `json:"stateText,omitempty"`
	Version         string       `json:"version,omitempty"`
	CreatedBy       string       `json:"createdBy,omitempty"`
	CreatedDt       string       `json:"createdDt,omitempty"`
	ModifiedBy      string       `json:"modifiedBy,omitempty"`
	ModifiedDt      string       `json:"modifiedDt,omitempty"`
}

type DbClusterAttachSecurityGroupRequest struct {
	SecurityGroupId string `json:"securityGroupId"`
}

type DbClusterContractDetailResponse struct {
	ContractEndDate     string `json:"contractEndDate,omitempty"`
	ContractPeriod      string `json:"contractPeriod,omitempty"`
	ContractStartDate   string `json:"contractStartDate,omitempty"`
	NextContractEndDate string `json:"nextContractEndDate,omitempty"`
	NextContractPeriod  string `json:"nextContractPeriod,omitempty"`
}

type DbClusterDetachSecurityGroupRequest struct {
	SecurityGroupId string `json:"securityGroupId"`
}

type DbClusterMaintenanceResponse struct {
	MaintenancePeriod         string `json:"maintenancePeriod,omitempty"`
	MaintenanceStartDayOfWeek string `json:"maintenanceStartDayOfWeek,omitempty"`
	MaintenanceStartTime      string `json:"maintenanceStartTime,omitempty"`
}

type DbClusterModifyContractRequest struct {
	ContractPeriod string `json:"contractPeriod"`
}

type DbClusterModifyNextContractRequest struct {
	NextContractPeriod string `json:"nextContractPeriod"`
}

type DbExtEventItemResponse struct {
	DbEventCategory  string       `json:"dbEventCategory,omitempty"`
	DbEventGuide     string       `json:"dbEventGuide,omitempty"`
	DbEventId        string       `json:"dbEventId,omitempty"`
	DbEventLevel     string       `json:"dbEventLevel,omitempty"`
	DbEventMessage   string       `json:"dbEventMessage,omitempty"`
	DbEventRequester string       `json:"dbEventRequester,omitempty"`
	CreatedDt        time.Time `json:"createdDt,omitempty"`
}

type DbInstancesResponse struct {
	DataBlockStorageSpec string                         `json:"dataBlockStorageSpec,omitempty"`
	DatabaseState        string                         `json:"databaseState,omitempty"`
	Region               string                         `json:"region,omitempty"`
	ServiceZoneId        string                         `json:"serviceZoneId,omitempty"`
	Software             *SoftwareResponse `json:"software,omitempty"`
	VirtualServerDbType  string                         `json:"virtualServerDbType,omitempty"`
	VirtualServerId      string                         `json:"virtualServerId,omitempty"`
	VirtualServerName    string                         `json:"virtualServerName,omitempty"`
	CreatedBy            string                         `json:"createdBy,omitempty"`
	CreatedDt            time.Time                   `json:"createdDt,omitempty"`
	ModifiedBy           string                         `json:"modifiedBy,omitempty"`
	ModifiedDt           time.Time                   `json:"modifiedDt,omitempty"`
}

type DbLogExportConfigConsoleResponse struct {
	DbLogExportConfigId   string `json:"dbLogExportConfigId,omitempty"`
	DbLogLabel            string `json:"dbLogLabel,omitempty"`
	DbLogPath             string `json:"dbLogPath,omitempty"`
	DbLogType             string `json:"dbLogType,omitempty"`
	ObjectStorageId       string `json:"objectStorageId,omitempty"`
	ObsBucketId           string `json:"obsBucketId,omitempty"`
	ObsBucketName         string `json:"obsBucketName,omitempty"`
	ScheduleDayOfMonth    int32  `json:"scheduleDayOfMonth,omitempty"`
	ScheduleDayOfWeek     string `json:"scheduleDayOfWeek,omitempty"`
	ScheduleFrequencyType string `json:"scheduleFrequencyType,omitempty"`
	ScheduleHour          int32  `json:"scheduleHour,omitempty"`
}

type DbLogExportConfigCreateRequest struct {
	DbLogType             string `json:"dbLogType"`
	ObjectStorageId       string `json:"objectStorageId"`
	ObsBucketId           string `json:"obsBucketId"`
	ScheduleDayOfMonth    int32  `json:"scheduleDayOfMonth,omitempty"`
	ScheduleDayOfWeek     string `json:"scheduleDayOfWeek,omitempty"`
	ScheduleFrequencyType string `json:"scheduleFrequencyType"`
	ScheduleHour          int32  `json:"scheduleHour"`
}

type DbLogExportConfigResponse struct {
	DbLogExportConfigId   string `json:"dbLogExportConfigId,omitempty"`
	DbLogLabel            string `json:"dbLogLabel,omitempty"`
	DbLogPath             string `json:"dbLogPath,omitempty"`
	DbLogType             string `json:"dbLogType,omitempty"`
	DeleteOnExport        *bool  `json:"deleteOnExport,omitempty"`
	ObjectStorageId       string `json:"objectStorageId,omitempty"`
	ObsBucketId           string `json:"obsBucketId,omitempty"`
	ScheduleDayOfMonth    int32  `json:"scheduleDayOfMonth,omitempty"`
	ScheduleDayOfWeek     string `json:"scheduleDayOfWeek,omitempty"`
	ScheduleFrequencyType string `json:"scheduleFrequencyType,omitempty"`
	ScheduleHour          int32  `json:"scheduleHour,omitempty"`
}

type DbLogExportConfigUpdateRequest struct {
	ScheduleDayOfMonth    int32  `json:"scheduleDayOfMonth,omitempty"`
	ScheduleDayOfWeek     string `json:"scheduleDayOfWeek,omitempty"`
	ScheduleFrequencyType string `json:"scheduleFrequencyType"`
	ScheduleHour          int32  `json:"scheduleHour"`
}

type DbReplicaInfoResponseVo struct {
	AccumulatedReplicaCount int32                      `json:"accumulatedReplicaCount,omitempty"`
	MasterServerGroupId     string                     `json:"masterServerGroupId,omitempty"`
	MasterServerGroupName   string                     `json:"masterServerGroupName,omitempty"`
	MasterServerGroupState  string                     `json:"masterServerGroupState,omitempty"`
	MasterUpgradeSequence   int64                      `json:"masterUpgradeSequence,omitempty"`
	Replicas                []ReplicaInfo `json:"replicas,omitempty"`
}

type DbServerGroupsResponse struct {
	DbName               string                             `json:"dbName,omitempty"`
	Region               string                             `json:"region,omitempty"`
	ServerGroupId        string                             `json:"serverGroupId,omitempty"`
	ServerGroupName      string                             `json:"serverGroupName,omitempty"`
	ServerGroupState     string                             `json:"serverGroupState,omitempty"`
	SoftwareServiceState string                             `json:"softwareServiceState,omitempty"`
	VirtualServers       []DbInstancesResponse `json:"virtualServers,omitempty"`
	CreatedBy            string                             `json:"createdBy,omitempty"`
	CreatedDt            time.Time                       `json:"createdDt,omitempty"`
	ModifiedBy           string                             `json:"modifiedBy,omitempty"`
	ModifiedDt           time.Time                       `json:"modifiedDt,omitempty"`
}

type DefaultParamItem struct {
	DefaultValue string `json:"defaultValue,omitempty"`
	Name         string `json:"name,omitempty"`
	SoftwareType string `json:"softwareType,omitempty"`
	Description  string `json:"description,omitempty"`
}

type DeleteBackupItem struct {
	BackupHistoryId string `json:"backupHistoryId"`
}

type DeleteBackupRequest struct {
	Backups []DeleteBackupItem `json:"backups,omitempty"`
}

type DetachSecurityGroupsRequest struct {
	SecurityGroupIds []string `json:"securityGroupIds"`
}

type DetailDatabaseResponse struct {
	ProjectId               string                                       `json:"projectId,omitempty"`
	AkhqPort                string                                       `json:"akhqPort,omitempty"`
	AkhqUserId              string                                       `json:"akhqUserId,omitempty"`
	AvailabilityZoneConfig  *AvailabilityZoneConfigResponse `json:"availabilityZoneConfig,omitempty"`
	Backup                  *DatabaseBackup                 `json:"backup,omitempty"`
	BlockId                 string                                       `json:"blockId,omitempty"`
	BlockName               string                                       `json:"blockName,omitempty"`
	BrokerNatPort           string                                       `json:"brokerNatPort,omitempty"`
	ClusterVip              string                                       `json:"clusterVip,omitempty"`
	ContractInfo            *ContractResponseV2             `json:"contractInfo,omitempty"`
	CreatorEmail            string                                       `json:"creatorEmail,omitempty"`
	CreatorName             string                                       `json:"creatorName,omitempty"`
	DbConfigs               map[string]string                            `json:"dbConfigs,omitempty"`
	DbName                  string                                       `json:"dbName,omitempty"`
	DbPort                  string                                       `json:"dbPort,omitempty"`
	DbUserId                string                                       `json:"dbUserId,omitempty"`
	DeploymentEnvType       string                                       `json:"deploymentEnvType,omitempty"`
	ExternalVip             string                                       `json:"externalVip,omitempty"`
	ImageId                 string                                       `json:"imageId,omitempty"`
	IsVpcEnv                *bool                                        `json:"isVpcEnv,omitempty"`
	KibanaPort              string                                       `json:"kibanaPort,omitempty"`
	LoggingAuditPath        string                                       `json:"loggingAuditPath,omitempty"`
	Maintenance             *DatabaseMaintenance            `json:"maintenance,omitempty"`
	ManagementConsolePort   string                                       `json:"managementConsolePort,omitempty"`
	MasterServerGroup       *ServerGroupInfo                `json:"masterServerGroup,omitempty"`
	ObjectStorageBuckets    []ObjectStorageBucketResponse   `json:"objectStorageBuckets,omitempty"`
	ProductGroupId          string                                       `json:"productGroupId,omitempty"`
	Region                  string                                       `json:"region,omitempty"`
	ReplicaServerGroups     []ServerGroupInfo               `json:"replicaServerGroups,omitempty"`
	SecurityGroups          []DatabaseSecurityGroupResponse `json:"securityGroups,omitempty"`
	SentinelPort            string                                       `json:"sentinelPort,omitempty"`
	ServerGroupId           string                                       `json:"serverGroupId,omitempty"`
	ServerGroupName         string                                       `json:"serverGroupName,omitempty"`
	ServerGroupState        string                                       `json:"serverGroupState,omitempty"`
	ServiceZoneId           string                                       `json:"serviceZoneId,omitempty"`
	ServicedFor             string                                       `json:"servicedFor,omitempty"`
	Software                *SoftwareResponse               `json:"software,omitempty"`
	SqlServerServiceName    string                                       `json:"sqlServerServiceName,omitempty"`
	SqlserverDbDiskInfoList []SqlserverDbDiskInfo           `json:"sqlserverDbDiskInfoList,omitempty"`
	Srn                     string                                       `json:"srn,omitempty"`
	Timezone                string                                       `json:"timezone,omitempty"`
	UpgradableImages        []UpgradableImageResponseVo     `json:"upgradableImages,omitempty"`
	Vip                     string                                       `json:"vip,omitempty"`
	VirtualServers          []VirtualServerResponse         `json:"virtualServers,omitempty"`
	ZookeeperPort           string                                       `json:"zookeeperPort,omitempty"`
	ZookeeperUserId         string                                       `json:"zookeeperUserId,omitempty"`
	CreatedBy               string                                       `json:"createdBy,omitempty"`
	CreatedDt               time.Time                                 `json:"createdDt,omitempty"`
	ModifiedBy              string                                       `json:"modifiedBy,omitempty"`
	ModifiedDt              time.Time                                 `json:"modifiedDt,omitempty"`
}

type ElasticSearchMasterNode struct {
	AdditionalBlockStorages []DatabaseBlockStorage `json:"additionalBlockStorages,omitempty"`
	DataBlockStorageSize    int32                               `json:"dataBlockStorageSize"`
	DataDiskType            string                              `json:"dataDiskType"`
	EncryptEnabled          *bool                               `json:"encryptEnabled"`
	MasterNodeCount         int32                               `json:"masterNodeCount"`
	ScaleProductId          string                              `json:"scaleProductId"`
}

type ElasticsearchEntKibanaNode struct {
	KibanaNodeCount int32  `json:"kibanaNodeCount"`
	ScaleProductId  string `json:"scaleProductId"`
}

type HaNormal struct {
	ReservedNatIpId string `json:"reservedNatIpId,omitempty"`
	UseVipNat       *bool  `json:"useVipNat,omitempty"`
	VirtualIp       string `json:"virtualIp,omitempty"`
}

type InstanceSpec struct {
	AdditionalBlockStorages []DatabaseBlockStorage `json:"additionalBlockStorages,omitempty"`
	DataBlockStorageSize    int32                               `json:"dataBlockStorageSize"`
	DataDiskType            string                              `json:"dataDiskType"`
	EncryptEnabled          *bool                               `json:"encryptEnabled"`
	ScaleProductId          string                              `json:"scaleProductId"`
}

type InstanceSpecWithAddtionalDisks struct {
	AdditionalBlockStorages []DatabaseBlockStorage `json:"additionalBlockStorages,omitempty"`
	DataBlockStorageSize    int32                               `json:"dataBlockStorageSize"`
	DataDiskType            string                              `json:"dataDiskType"`
	EncryptEnabled          *bool                               `json:"encryptEnabled"`
	ScaleProductId          string                              `json:"scaleProductId"`
}

type ListResponseDatabaseBackupsResponse struct {
	Contents   []DatabaseBackupsResponse `json:"contents,omitempty"`
	TotalCount int32                                  `json:"totalCount,omitempty"`
}

type ListResponseDatabaseParameterGroupsResponse struct {
	Contents   []DatabaseParameterGroupsResponse `json:"contents,omitempty"`
	TotalCount int32                                          `json:"totalCount,omitempty"`
}

type ListResponseDatabaseStroagesResponse struct {
	Contents   []DatabaseStroagesResponse `json:"contents,omitempty"`
	TotalCount int32                                   `json:"totalCount,omitempty"`
}

type ListResponseDatabaseUserResponse struct {
	Contents   []DatabaseUserResponse `json:"contents,omitempty"`
	TotalCount int32                               `json:"totalCount,omitempty"`
}

type ListResponseDbBackupsDomain struct {
	Contents   []DbBackupsDomain `json:"contents,omitempty"`
	TotalCount int32                          `json:"totalCount,omitempty"`
}

type ListResponseDbExtEventItemResponse struct {
	Contents   []DbExtEventItemResponse `json:"contents,omitempty"`
	TotalCount int32                                 `json:"totalCount,omitempty"`
}

type ListResponseDbLogExportConfigConsoleResponse struct {
	Contents   []DbLogExportConfigConsoleResponse `json:"contents,omitempty"`
	TotalCount int32                                           `json:"totalCount,omitempty"`
}

type ListResponseDbLogExportConfigResponse struct {
	Contents   []DbLogExportConfigResponse `json:"contents,omitempty"`
	TotalCount int32                                    `json:"totalCount,omitempty"`
}

type ListResponseDbServerGroupsResponse struct {
	Contents   []DbServerGroupsResponse `json:"contents,omitempty"`
	TotalCount int32                                 `json:"totalCount,omitempty"`
}

type ListResponseLogConfigResponse struct {
	Contents   []LogConfigResponse `json:"contents,omitempty"`
	TotalCount int32                            `json:"totalCount,omitempty"`
}

type ListResponseSqlserverClusterListItemResponse struct {
	Contents   []SqlserverClusterListItemResponse `json:"contents,omitempty"`
	TotalCount int32                                           `json:"totalCount,omitempty"`
}

type LogConfigResponse struct {
	DbLogLabel string `json:"dbLogLabel,omitempty"`
	DbLogPath  string `json:"dbLogPath,omitempty"`
	DbLogType  string `json:"dbLogType,omitempty"`
}

type MsSqlConfigurationControllerApiListSqlserverOpts struct {
	DbName            optional.String
	Region            optional.String
	ServerGroupName   optional.String
	State             optional.String
	VirtualServerName optional.String
	CreatedBy         optional.String
	Page              optional.Int32
	Size              optional.Int32
	Sort              optional.Interface
}

type NetworkResponse struct {
	ExternalIp   string `json:"externalIp,omitempty"`
	ExternalIpId string `json:"externalIpId,omitempty"`
	Ip           string `json:"ip,omitempty"`
	IpId         string `json:"ipId,omitempty"`
	NetworkId    string `json:"networkId,omitempty"`
	NetworkName  string `json:"networkName,omitempty"`
	NicId        string `json:"nicId,omitempty"`
	UplinkType   string `json:"uplinkType,omitempty"`
	VpcId        string `json:"vpcId,omitempty"`
	VpcName      string `json:"vpcName,omitempty"`
}

type ObjectStorageBucketResponse struct {
	ObjectStorageBucketId   string       `json:"objectStorageBucketId,omitempty"`
	ObjectStorageBucketName string       `json:"objectStorageBucketName,omitempty"`
	ObjectStorageId         string       `json:"objectStorageId,omitempty"`
	ObjectStorageName       string       `json:"objectStorageName,omitempty"`
	ObsBucketUsedType       string       `json:"obsBucketUsedType,omitempty"`
	Region                  string       `json:"region,omitempty"`
	CreatedDt               time.Time `json:"createdDt,omitempty"`
	ModifiedDt              time.Time `json:"modifiedDt,omitempty"`
}

type ReplicaInfo struct {
	LastCheckDt        time.Time `json:"lastCheckDt,omitempty"`
	ReplicaState       string       `json:"replicaState,omitempty"`
	ReplicaStateDetail string       `json:"replicaStateDetail,omitempty"`
	ServerGroupId      string       `json:"serverGroupId,omitempty"`
	ServerGroupName    string       `json:"serverGroupName,omitempty"`
	VirtualServerId    string       `json:"virtualServerId,omitempty"`
	VirtualServerName  string       `json:"virtualServerName,omitempty"`
}

type ReplicaSpec struct {
	AdditionalBlockStorages []DatabaseBlockStorage `json:"additionalBlockStorages,omitempty"`
	Maintenance             *DatabaseMaintenance   `json:"maintenance,omitempty"`
	ScaleProductId          string                              `json:"scaleProductId"`
	Tags                    []TagRequest           `json:"tags,omitempty"`
	VirtualServerNames      []string                            `json:"virtualServerNames"`
}

type ReplicaStateResponse struct {
	Inprogress             *bool                      `json:"inprogress,omitempty"`
	MasterServerGroupId    string                     `json:"masterServerGroupId,omitempty"`
	MasterServerGroupName  string                     `json:"masterServerGroupName,omitempty"`
	MasterServerGroupState string                     `json:"masterServerGroupState,omitempty"`
	Replicas               []ReplicaInfo `json:"replicas,omitempty"`
}

type RequestDomain struct {
	ProjectId       string                           `json:"projectId,omitempty"`
	AwxJobId        string                           `json:"awxJobId,omitempty"`
	CauseId         string                           `json:"causeId,omitempty"`
	ChildRequestIds string                           `json:"childRequestIds,omitempty"`
	ChildRequests   []RequestDomain     `json:"childRequests,omitempty"`
	Data            string                           `json:"data,omitempty"`
	Histories       []RequestHistDomain `json:"histories,omitempty"`
	Id              string                           `json:"id,omitempty"`
	ObjectId        string                           `json:"objectId,omitempty"`
	ObjectScope     string                           `json:"objectScope,omitempty"`
	Operation       string                           `json:"operation,omitempty"`
	ParentId        string                           `json:"parentId,omitempty"`
	PoolId          string                           `json:"poolId,omitempty"`
	RequestBy       string                           `json:"requestBy,omitempty"`
	RequestOrigin   string                           `json:"requestOrigin,omitempty"`
	RequestTime     string                           `json:"requestTime,omitempty"`
	ResponseTime    string                           `json:"responseTime,omitempty"`
	Retried         string                           `json:"retried,omitempty"`
	RootId          string                           `json:"rootId,omitempty"`
	State           string                           `json:"state,omitempty"`
	SystemId        string                           `json:"systemId,omitempty"`
	TopRequestId    string                           `json:"topRequestId,omitempty"`
}

type RequestHistDomain struct {
	AwxJobId     string       `json:"awxJobId,omitempty"`
	Data         string       `json:"data,omitempty"`
	ErrorCode    string       `json:"errorCode,omitempty"`
	ErrorDetail  string       `json:"errorDetail,omitempty"`
	ErrorText    string       `json:"errorText,omitempty"`
	Id           string       `json:"id,omitempty"`
	OrchestUrl   string       `json:"orchestUrl,omitempty"`
	ProgressRate int32        `json:"progressRate,omitempty"`
	RootId       string       `json:"rootId,omitempty"`
	State        string       `json:"state,omitempty"`
	StateText    string       `json:"stateText,omitempty"`
	Target       string       `json:"target,omitempty"`
	TraceInfo    string       `json:"traceInfo,omitempty"`
	CreatedDt    time.Time `json:"createdDt,omitempty"`
}

type ResizeScaleRequest struct {
	ScaleProductId  string `json:"scaleProductId"`
	VirtualServerId string `json:"virtualServerId"`
}

type ResizeSqlServerRequest struct {
	AddedDb          []string                           `json:"addedDb"`
	AddedDbDiskInfos []SqlserverDbDiskInfo `json:"addedDbDiskInfos"`
	DeletedDb        []string                           `json:"deletedDb"`
}

type ResizeStorageRequest struct {
	BlockStorageId   string `json:"blockStorageId"`
	BlockStorageSize int32  `json:"blockStorageSize"`
	VirtualServerId  string `json:"virtualServerId"`
}

type RestoreDatabaseRequest struct {
	ActiveDirectory         *ActiveDirectory                `json:"activeDirectory,omitempty"`
	AvailabilityZoneName    string                                       `json:"availabilityZoneName,omitempty"`
	BackupHistoryId         string                                       `json:"backupHistoryId,omitempty"`
	BackupRecoveryDt        time.Time                                 `json:"backupRecoveryDt,omitempty"`
	DeploymentEnvType       string                                       `json:"deploymentEnvType"`
	KibanaVirtualServer     *ElasticsearchEntKibanaNode     `json:"kibanaVirtualServer,omitempty"`
	LicenseKey              string                                       `json:"licenseKey,omitempty"`
	Maintenance             *DatabaseMaintenance            `json:"maintenance,omitempty"`
	MasterVirtualServer     *ElasticSearchMasterNode        `json:"masterVirtualServer,omitempty"`
	ServerGroupName         string                                       `json:"serverGroupName"`
	Tags                    []TagRequest                    `json:"tags,omitempty"`
	VirtualServer           *InstanceSpecWithAddtionalDisks `json:"virtualServer"`
	VirtualServerNamePrefix string                                       `json:"virtualServerNamePrefix"`
}

type ServerGroupInfo struct {
	ServerGroupId    string `json:"serverGroupId,omitempty"`
	ServerGroupName  string `json:"serverGroupName,omitempty"`
	ServerGroupState string `json:"serverGroupState,omitempty"`
}

type ServiceParamItem struct {
	AllowsValue     string `json:"allowsValue,omitempty"`
	AppliedValue    string `json:"appliedValue,omitempty"`
	ApplyType       string `json:"applyType,omitempty"`
	DataType        string `json:"dataType,omitempty"`
	DbParameterId   string `json:"dbParameterId,omitempty"`
	DbParameterName string `json:"dbParameterName,omitempty"`
	DefaultValue    string `json:"defaultValue,omitempty"`
	Modifiable      *bool  `json:"modifiable,omitempty"`
	SoftwareType    string `json:"softwareType,omitempty"`
	Description     string `json:"description,omitempty"`
}

type SoftwareResponse struct {
	SoftwareCategory     string            `json:"softwareCategory,omitempty"`
	SoftwareId           string            `json:"softwareId,omitempty"`
	SoftwareName         string            `json:"softwareName,omitempty"`
	SoftwareProperties   map[string]string `json:"softwareProperties,omitempty"`
	SoftwareServiceState string            `json:"softwareServiceState,omitempty"`
	SoftwareState        string            `json:"softwareState,omitempty"`
	SoftwareType         string            `json:"softwareType,omitempty"`
	SoftwareVersion      string            `json:"softwareVersion,omitempty"`
	CreatedBy            string            `json:"createdBy,omitempty"`
	CreatedDt            time.Time      `json:"createdDt,omitempty"`
	ModifiedBy           string            `json:"modifiedBy,omitempty"`
	ModifiedDt           time.Time      `json:"modifiedDt,omitempty"`
}

type SqlserverActiveDirectory struct {
	AdServerUserId       string   `json:"adServerUserId"`
	AdServerUserPassword string   `json:"adServerUserPassword"`
	DnsServerIps         []string `json:"dnsServerIps"`
	DomainName           string   `json:"domainName"`
	DomainNetBiosName    string   `json:"domainNetBiosName"`
	FailoverClusterName  string   `json:"failoverClusterName"`
}

type SqlserverAddStorageRequest struct {
	BlockStorageSize int32  `json:"blockStorageSize"`
	BlockStorageType string `json:"blockStorageType"`
	DiskType         string `json:"diskType"`
	DriveLetter      string `json:"driveLetter,omitempty"`
	VirtualServerId  string `json:"virtualServerId"`
}

type SqlserverBackupConfigResponse struct {
	FullBackupConfig *SqlserverFullBackupConfigResponse `json:"fullBackupConfig,omitempty"`
}

type SqlserverBlockStorageGroupCreateRequest struct {
	BlockStorageSize int32  `json:"blockStorageSize"`
	BlockStorageType string `json:"blockStorageType,omitempty"`
}

type SqlserverBlockStorageGroupDetailResponse struct {
	BlockStorageGroupId  string `json:"blockStorageGroupId,omitempty"`
	BlockStorageName     string `json:"blockStorageName,omitempty"`
	BlockStorageRoleType string `json:"blockStorageRoleType,omitempty"`
	BlockStorageSize     int32  `json:"blockStorageSize,omitempty"`
	BlockStorageType     string `json:"blockStorageType,omitempty"`
}

type SqlserverClusterCreateRequest struct {
	AuditEnabled           *bool                                             `json:"auditEnabled,omitempty"`
	ImageId                string                                            `json:"imageId"`
	NatEnabled             *bool                                             `json:"natEnabled,omitempty"`
	NatPublicIpId          string                                            `json:"natPublicIpId,omitempty"`
	SecurityGroupIds       []string                                          `json:"securityGroupIds"`
	ServiceZoneId          string                                            `json:"serviceZoneId"`
	SqlserverClusterName   string                                            `json:"sqlserverClusterName"`
	SqlserverInitialConfig *SqlserverInitialConfigCreateRequest `json:"sqlserverInitialConfig"`
	SqlserverServerGroup   *SqlserverServerGroupCreateRequest   `json:"sqlserverServerGroup"`
	SubnetId               string                                            `json:"subnetId"`
	Tags                   []TagRequest                         `json:"tags,omitempty"`
	Timezone               string                                            `json:"timezone,omitempty"`
}

type SqlserverClusterDetailResponse struct {
	ProjectId                   string                                             `json:"projectId,omitempty"`
	AuditEnabled                *bool                                              `json:"auditEnabled,omitempty"`
	BackupConfig                *SqlserverBackupConfigResponse        `json:"backupConfig,omitempty"`
	BlockId                     string                                             `json:"blockId,omitempty"`
	Contract                    *DbClusterContractDetailResponse      `json:"contract,omitempty"`
	DatabaseVersion             string                                             `json:"databaseVersion,omitempty"`
	ImageId                     string                                             `json:"imageId,omitempty"`
	Maintenance                 *DbClusterMaintenanceResponse         `json:"maintenance,omitempty"`
	NatIpAddress                string                                             `json:"natIpAddress,omitempty"`
	QuorumServerGroup           *SqlserverServerGroupDetailResponse   `json:"quorumServerGroup,omitempty"`
	SecurityGroupIds            []string                                           `json:"securityGroupIds,omitempty"`
	ServiceZoneId               string                                             `json:"serviceZoneId,omitempty"`
	SqlserverClusterId          string                                             `json:"sqlserverClusterId,omitempty"`
	SqlserverClusterName        string                                             `json:"sqlserverClusterName,omitempty"`
	SqlserverClusterState       string                                             `json:"sqlserverClusterState,omitempty"`
	SqlserverInitialConfig      *SqlserverInitialConfigDetailResponse `json:"sqlserverInitialConfig,omitempty"`
	SqlserverMasterClusterId    string                                             `json:"sqlserverMasterClusterId,omitempty"`
	SqlserverSecondaryClusterId string                                             `json:"sqlserverSecondaryClusterId,omitempty"`
	SqlserverServerGroup        *SqlserverServerGroupDetailResponse   `json:"sqlserverServerGroup,omitempty"`
	SubnetId                    string                                             `json:"subnetId,omitempty"`
	Timezone                    string                                             `json:"timezone,omitempty"`
	VpcId                       string                                             `json:"vpcId,omitempty"`
	CreatedBy                   string                                             `json:"createdBy,omitempty"`
	CreatedDt                   time.Time                                       `json:"createdDt,omitempty"`
	ModifiedBy                  string                                             `json:"modifiedBy,omitempty"`
	ModifiedDt                  time.Time                                       `json:"modifiedDt,omitempty"`
}

type SqlserverClusterListItemResponse struct {
	ProjectId             string       `json:"projectId,omitempty"`
	BlockId               string       `json:"blockId,omitempty"`
	ServiceZoneId         string       `json:"serviceZoneId,omitempty"`
	SqlserverClusterId    string       `json:"sqlserverClusterId,omitempty"`
	SqlserverClusterName  string       `json:"sqlserverClusterName,omitempty"`
	SqlserverClusterState string       `json:"sqlserverClusterState,omitempty"`
	CreatedBy             string       `json:"createdBy,omitempty"`
	CreatedDt             time.Time `json:"createdDt,omitempty"`
	ModifiedBy            string       `json:"modifiedBy,omitempty"`
	ModifiedDt            time.Time `json:"modifiedDt,omitempty"`
}

type SqlserverClusterResizeBlockStoragesRequest struct {
	BlockStorageGroupId string `json:"blockStorageGroupId"`
	BlockStorageSize    int32  `json:"blockStorageSize"`
}

type SqlserverClusterResizeVirtualServersRequest struct {
	ServerType string `json:"serverType"`
}

type SqlserverCreateFullBackupConfigRequest struct {
	ArchiveBackupScheduleFrequency string `json:"archiveBackupScheduleFrequency"`
	BackupRetentionPeriod          string `json:"backupRetentionPeriod"`
	BackupStartHour                int32  `json:"backupStartHour"`
	FullBackupDayOfWeek            string `json:"fullBackupDayOfWeek"`
	ObjectStorageId                string `json:"objectStorageId"`
}

type SqlserverDbDiskInfo struct {
	DatabaseName string `json:"databaseName"`
	DriveLetter  string `json:"driveLetter"`
}

type SqlserverFullBackupConfigResponse struct {
	ArchiveBackupScheduleFrequency string `json:"archiveBackupScheduleFrequency,omitempty"`
	BackupRetentionPeriod          string `json:"backupRetentionPeriod,omitempty"`
	BackupStartHour                int32  `json:"backupStartHour,omitempty"`
	FullBackupDayOfWeek            string `json:"fullBackupDayOfWeek,omitempty"`
	ObjectStorageBucketId          string `json:"objectStorageBucketId,omitempty"`
}

type SqlserverInitialConfigCreateRequest struct {
	ActiveDirectory      *SqlserverActiveDirectory `json:"activeDirectory,omitempty"`
	DatabaseCollation    string                                 `json:"databaseCollation"`
	DatabaseNames        []string                               `json:"databaseNames"`
	DatabasePort         int32                                  `json:"databasePort,omitempty"`
	DatabaseServiceName  string                                 `json:"databaseServiceName"`
	DatabaseUserName     string                                 `json:"databaseUserName"`
	DatabaseUserPassword string                                 `json:"databaseUserPassword"`
	License              string                                 `json:"license"`
}

type SqlserverInitialConfigDetailResponse struct {
	ActiveDirectory     *SqlserverActiveDirectory `json:"activeDirectory,omitempty"`
	DatabaseCollation   string                                 `json:"databaseCollation,omitempty"`
	DatabaseNames       []string                               `json:"databaseNames,omitempty"`
	DatabasePort        int32                                  `json:"databasePort,omitempty"`
	DatabaseServiceName string                                 `json:"databaseServiceName,omitempty"`
	DatabaseUserName    string                                 `json:"databaseUserName,omitempty"`
}

type SqlserverModifyFullBackupConfigRequest struct {
	ArchiveBackupScheduleFrequency string `json:"archiveBackupScheduleFrequency"`
	BackupRetentionPeriod          string `json:"backupRetentionPeriod"`
	BackupStartHour                int32  `json:"backupStartHour"`
	FullBackupDayOfWeek            string `json:"fullBackupDayOfWeek"`
}

type SqlserverSearchApiListSqlserverClustersOpts struct {
	SqlserverClusterName optional.String
	Page                 optional.Int32
	Size                 optional.Int32
	Sort                 optional.Interface
}

type SqlserverServerCreateRequest struct {
	AvailabilityZoneName string `json:"availabilityZoneName,omitempty"`
	ServerRoleType       string `json:"serverRoleType"`
	SqlserverServerName  string `json:"sqlserverServerName"`
}

type SqlserverServerDetailResponse struct {
	AvailabilityZoneName string       `json:"availabilityZoneName,omitempty"`
	ServerRoleType       string       `json:"serverRoleType,omitempty"`
	SqlserverServerId    string       `json:"sqlserverServerId,omitempty"`
	SqlserverServerName  string       `json:"sqlserverServerName,omitempty"`
	SqlserverServerState string       `json:"sqlserverServerState,omitempty"`
	SubnetIpAddress      string       `json:"subnetIpAddress,omitempty"`
	CreatedBy            string       `json:"createdBy,omitempty"`
	CreatedDt            time.Time `json:"createdDt,omitempty"`
	ModifiedBy           string       `json:"modifiedBy,omitempty"`
	ModifiedDt           time.Time `json:"modifiedDt,omitempty"`
}

type SqlserverServerGroupCreateRequest struct {
	BlockStorages     []SqlserverBlockStorageGroupCreateRequest `json:"blockStorages"`
	EncryptionEnabled *bool                                                  `json:"encryptionEnabled,omitempty"`
	ServerType        string                                                 `json:"serverType"`
	SqlserverServers  []SqlserverServerCreateRequest            `json:"sqlserverServers"`
}

type SqlserverServerGroupDetailResponse struct {
	BlockStorages       []SqlserverBlockStorageGroupDetailResponse `json:"blockStorages,omitempty"`
	EncryptionEnabled   *bool                                                   `json:"encryptionEnabled,omitempty"`
	ServerGroupRoleType string                                                  `json:"serverGroupRoleType,omitempty"`
	ServerType          string                                                  `json:"serverType,omitempty"`
	SqlserverServers    []SqlserverServerDetailResponse            `json:"sqlserverServers,omitempty"`
	VirtualIpAddress    string                                                  `json:"virtualIpAddress,omitempty"`
}

type TagRequest struct {
	TagKey   string `json:"tagKey"`
	TagValue string `json:"tagValue,omitempty"`
}

type UpdateBackupSettingRequest struct {
	Backup    *DatabaseBackup `json:"backup,omitempty"`
	UseBackup *bool                        `json:"useBackup"`
}

type UpdateLoggingAuditRequest struct {
	UseDbLoggingAudit *bool `json:"useDbLoggingAudit"`
}

type UpdateMaintenanceRequest struct {
	Maintenance    *DatabaseMaintenance `json:"maintenance,omitempty"`
	UseMaintenance *bool                             `json:"useMaintenance"`
}

type UpdateUserRequest struct {
	DbLoginId   string `json:"dbLoginId,omitempty"`
	DbLoginType string `json:"dbLoginType,omitempty"`
	DbName      string `json:"dbName,omitempty"`
	Name        string `json:"name,omitempty"`
	State       string `json:"state,omitempty"`
	Description string `json:"description,omitempty"`
}

type UpdateUsersRequest struct {
	UpdateUsersRequest []UpdateUserRequest `json:"updateUsersRequest,omitempty"`
}

type UpgradableImageResponseVo struct {
	ImageId   string `json:"imageId,omitempty"`
	ImageName string `json:"imageName,omitempty"`
}

type UpgradeDatabaseRequest struct {
	ImageId             string `json:"imageId"`
	PatchBackupIncluded *bool  `json:"patchBackupIncluded,omitempty"`
	Description         string `json:"description,omitempty"`
}

type VirtualServerResponse struct {
	AvailabilityZoneName      string                                  `json:"availabilityZoneName,omitempty"`
	BlockStorages             []DatabaseStroagesResponse `json:"blockStorages,omitempty"`
	ContractDiscountName      string                                  `json:"contractDiscountName,omitempty"`
	ContractDiscountProductId string                                  `json:"contractDiscountProductId,omitempty"`
	Network                   *NetworkResponse           `json:"network,omitempty"`
	ScaleName                 string                                  `json:"scaleName,omitempty"`
	ScaleProductId            string                                  `json:"scaleProductId,omitempty"`
	ScaleType                 string                                  `json:"scaleType,omitempty"`
	Software                  *SoftwareResponse          `json:"software,omitempty"`
	VirtualServerId           string                                  `json:"virtualServerId,omitempty"`
	VirtualServerName         string                                  `json:"virtualServerName,omitempty"`
	VirtualServerState        string                                  `json:"virtualServerState,omitempty"`
	VirtualServerType         string                                  `json:"virtualServerType,omitempty"`
	VpcId                     string                                  `json:"vpcId,omitempty"`
	ScaleDescription          string                                  `json:"scaleDescription,omitempty"`
	CreatedBy                 string                                  `json:"createdBy,omitempty"`
	CreatedDt                 time.Time                            `json:"createdDt,omitempty"`
	ModifiedBy                string                                  `json:"modifiedBy,omitempty"`
	ModifiedDt                time.Time                            `json:"modifiedDt,omitempty"`
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
