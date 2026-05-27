package redis

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

type ApplyCommandsRequest struct {
	Commands []DatabaseCommand `json:"commands"`
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

type CheckDatabasePasswordRequest struct {
	DbNewPwd string `json:"dbNewPwd,omitempty"`
	DbPwd    string `json:"dbPwd"`
}

type CommandItem struct {
	AllowsValue  string `json:"allowsValue,omitempty"`
	AppliedValue string `json:"appliedValue,omitempty"`
	ApplyType    string `json:"applyType,omitempty"`
	CommandId    string `json:"commandId,omitempty"`
	CommandName  string `json:"commandName,omitempty"`
	DataType     string `json:"dataType,omitempty"`
	DefaultValue string `json:"defaultValue,omitempty"`
	Modifiable   *bool  `json:"modifiable,omitempty"`
	SoftwareType string `json:"softwareType,omitempty"`
	Description  string `json:"description,omitempty"`
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

type CreateRedisRequest struct {
	Backup                  *DatabaseBackup      `json:"backup,omitempty"`
	BlockId                 string                        `json:"blockId"`
	DbPort                  int32                         `json:"dbPort,omitempty"`
	DbUserPassword          string                        `json:"dbUserPassword"`
	DeploymentEnvType       string                        `json:"deploymentEnvType"`
	HighAvailability        *HaRedis             `json:"highAvailability,omitempty"`
	ImageId                 string                        `json:"imageId"`
	Maintenance             *DatabaseMaintenance `json:"maintenance,omitempty"`
	Network                 *DatabaseNetwork     `json:"network"`
	ProductGroupId          string                        `json:"productGroupId"`
	SecurityGroupIds        []string                      `json:"securityGroupIds"`
	ServerGroupName         string                        `json:"serverGroupName"`
	ServiceZoneId           string                        `json:"serviceZoneId"`
	Tags                    []TagRequest         `json:"tags,omitempty"`
	Timezone                string                        `json:"timezone"`
	VirtualServer           *InstanceSpec        `json:"virtualServer"`
	VirtualServerNamePrefix string                        `json:"virtualServerNamePrefix"`
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

type DatabaseCommand struct {
	CommandId  string `json:"commandId"`
	NewCommand string `json:"newCommand"`
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

type DatabaseListResponseCommandItem struct {
	Contents   []CommandItem `json:"contents,omitempty"`
	TotalCount int32                  `json:"totalCount,omitempty"`
	ModifiedDt time.Time           `json:"modifiedDt,omitempty"`
}

type DatabaseListResponseServiceParamItem struct {
	Contents   []ServiceParamItem `json:"contents,omitempty"`
	TotalCount int32                       `json:"totalCount,omitempty"`
	ModifiedDt time.Time                `json:"modifiedDt,omitempty"`
}

type DatabaseMaintenance struct {
	MaintenanceStartDayOfWeek  string  `json:"maintenanceStartDayOfWeek"`
	MaintenanceStartHourMinute string  `json:"maintenanceStartHourMinute"`
	MaintenanceTermHour        float64 `json:"maintenanceTermHour"`
}

type DatabaseNetwork struct {
	NetworkEnvType string                           `json:"networkEnvType"`
	ServerNetworks []DatabaseServerNetwork `json:"serverNetworks"`
	SubnetId       string                           `json:"subnetId"`
	UseNat         *bool                            `json:"useNat"`
	VpcId          string                           `json:"vpcId"`
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
	DbParameterGroupId   string                      `json:"dbParameterGroupId,omitempty"`
	DbParameterGroupName string                      `json:"dbParameterGroupName,omitempty"`
	TotalCount           int32                       `json:"totalCount,omitempty"`
}

type DatabaseParameterGroupsResponse struct {
	DbParameterGroupId   string `json:"dbParameterGroupId,omitempty"`
	DbParameterGroupName string `json:"dbParameterGroupName,omitempty"`
}

type DatabasePasswordResponse struct {
	ErrorCode string `json:"errorCode,omitempty"`
	ErrorText string `json:"errorText,omitempty"`
	IsRunning *bool  `json:"isRunning,omitempty"`
	Status    string `json:"status,omitempty"`
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
	DataBlockStorageSpec string                     `json:"dataBlockStorageSpec,omitempty"`
	DatabaseState        string                     `json:"databaseState,omitempty"`
	Region               string                     `json:"region,omitempty"`
	ServiceZoneId        string                     `json:"serviceZoneId,omitempty"`
	Software             *SoftwareResponse `json:"software,omitempty"`
	VirtualServerDbType  string                     `json:"virtualServerDbType,omitempty"`
	VirtualServerId      string                     `json:"virtualServerId,omitempty"`
	VirtualServerName    string                     `json:"virtualServerName,omitempty"`
	CreatedBy            string                     `json:"createdBy,omitempty"`
	CreatedDt            time.Time               `json:"createdDt,omitempty"`
	ModifiedBy           string                     `json:"modifiedBy,omitempty"`
	ModifiedDt           time.Time               `json:"modifiedDt,omitempty"`
}

type DbServerGroupsResponse struct {
	DbName               string                         `json:"dbName,omitempty"`
	Region               string                         `json:"region,omitempty"`
	ServerGroupId        string                         `json:"serverGroupId,omitempty"`
	ServerGroupName      string                         `json:"serverGroupName,omitempty"`
	ServerGroupState     string                         `json:"serverGroupState,omitempty"`
	SoftwareServiceState string                         `json:"softwareServiceState,omitempty"`
	VirtualServers       []DbInstancesResponse `json:"virtualServers,omitempty"`
	CreatedBy            string                         `json:"createdBy,omitempty"`
	CreatedDt            time.Time                   `json:"createdDt,omitempty"`
	ModifiedBy           string                         `json:"modifiedBy,omitempty"`
	ModifiedDt           time.Time                   `json:"modifiedDt,omitempty"`
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
	ProjectId               string                                   `json:"projectId,omitempty"`
	AkhqPort                string                                   `json:"akhqPort,omitempty"`
	AkhqUserId              string                                   `json:"akhqUserId,omitempty"`
	AvailabilityZoneConfig  *AvailabilityZoneConfigResponse `json:"availabilityZoneConfig,omitempty"`
	Backup                  *DatabaseBackup                 `json:"backup,omitempty"`
	BlockId                 string                                   `json:"blockId,omitempty"`
	BlockName               string                                   `json:"blockName,omitempty"`
	BrokerNatPort           string                                   `json:"brokerNatPort,omitempty"`
	ClusterVip              string                                   `json:"clusterVip,omitempty"`
	ContractInfo            *ContractResponseV2             `json:"contractInfo,omitempty"`
	CreatorEmail            string                                   `json:"creatorEmail,omitempty"`
	CreatorName             string                                   `json:"creatorName,omitempty"`
	DbConfigs               map[string]string                        `json:"dbConfigs,omitempty"`
	DbName                  string                                   `json:"dbName,omitempty"`
	DbPort                  string                                   `json:"dbPort,omitempty"`
	DbUserId                string                                   `json:"dbUserId,omitempty"`
	DeploymentEnvType       string                                   `json:"deploymentEnvType,omitempty"`
	ExternalVip             string                                   `json:"externalVip,omitempty"`
	ImageId                 string                                   `json:"imageId,omitempty"`
	IsVpcEnv                *bool                                    `json:"isVpcEnv,omitempty"`
	KibanaPort              string                                   `json:"kibanaPort,omitempty"`
	LoggingAuditPath        string                                   `json:"loggingAuditPath,omitempty"`
	Maintenance             *DatabaseMaintenance            `json:"maintenance,omitempty"`
	ManagementConsolePort   string                                   `json:"managementConsolePort,omitempty"`
	MasterServerGroup       *ServerGroupInfo                `json:"masterServerGroup,omitempty"`
	ObjectStorageBuckets    []ObjectStorageBucketResponse   `json:"objectStorageBuckets,omitempty"`
	ProductGroupId          string                                   `json:"productGroupId,omitempty"`
	Region                  string                                   `json:"region,omitempty"`
	ReplicaServerGroups     []ServerGroupInfo               `json:"replicaServerGroups,omitempty"`
	SecurityGroups          []DatabaseSecurityGroupResponse `json:"securityGroups,omitempty"`
	SentinelPort            string                                   `json:"sentinelPort,omitempty"`
	ServerGroupId           string                                   `json:"serverGroupId,omitempty"`
	ServerGroupName         string                                   `json:"serverGroupName,omitempty"`
	ServerGroupState        string                                   `json:"serverGroupState,omitempty"`
	ServiceZoneId           string                                   `json:"serviceZoneId,omitempty"`
	ServicedFor             string                                   `json:"servicedFor,omitempty"`
	Software                *SoftwareResponse               `json:"software,omitempty"`
	SqlServerServiceName    string                                   `json:"sqlServerServiceName,omitempty"`
	SqlserverDbDiskInfoList []SqlserverDbDiskInfo           `json:"sqlserverDbDiskInfoList,omitempty"`
	Srn                     string                                   `json:"srn,omitempty"`
	Timezone                string                                   `json:"timezone,omitempty"`
	UpgradableImages        []UpgradableImageResponseVo     `json:"upgradableImages,omitempty"`
	Vip                     string                                   `json:"vip,omitempty"`
	VirtualServers          []VirtualServerResponse         `json:"virtualServers,omitempty"`
	ZookeeperPort           string                                   `json:"zookeeperPort,omitempty"`
	ZookeeperUserId         string                                   `json:"zookeeperUserId,omitempty"`
	CreatedBy               string                                   `json:"createdBy,omitempty"`
	CreatedDt               time.Time                             `json:"createdDt,omitempty"`
	ModifiedBy              string                                   `json:"modifiedBy,omitempty"`
	ModifiedDt              time.Time                             `json:"modifiedDt,omitempty"`
}

type ElasticSearchMasterNode struct {
	AdditionalBlockStorages []DatabaseBlockStorage `json:"additionalBlockStorages,omitempty"`
	DataBlockStorageSize    int32                           `json:"dataBlockStorageSize"`
	DataDiskType            string                          `json:"dataDiskType"`
	EncryptEnabled          *bool                           `json:"encryptEnabled"`
	MasterNodeCount         int32                           `json:"masterNodeCount"`
	ScaleProductId          string                          `json:"scaleProductId"`
}

type ElasticsearchEntKibanaNode struct {
	KibanaNodeCount int32  `json:"kibanaNodeCount"`
	ScaleProductId  string `json:"scaleProductId"`
}

type HaRedis struct {
	ReplicaCount int32 `json:"replicaCount"`
	SentinelPort int32 `json:"sentinelPort,omitempty"`
}

type InstanceSpec struct {
	AdditionalBlockStorages []DatabaseBlockStorage `json:"additionalBlockStorages,omitempty"`
	DataBlockStorageSize    int32                           `json:"dataBlockStorageSize"`
	DataDiskType            string                          `json:"dataDiskType"`
	EncryptEnabled          *bool                           `json:"encryptEnabled"`
	ScaleProductId          string                          `json:"scaleProductId"`
}

type InstanceSpecWithAddtionalDisks struct {
	AdditionalBlockStorages []DatabaseBlockStorage `json:"additionalBlockStorages,omitempty"`
	DataBlockStorageSize    int32                           `json:"dataBlockStorageSize"`
	DataDiskType            string                          `json:"dataDiskType"`
	EncryptEnabled          *bool                           `json:"encryptEnabled"`
	ScaleProductId          string                          `json:"scaleProductId"`
}

type ListResponseDatabaseBackupsResponse struct {
	Contents   []DatabaseBackupsResponse `json:"contents,omitempty"`
	TotalCount int32                              `json:"totalCount,omitempty"`
}

type ListResponseDatabaseParameterGroupsResponse struct {
	Contents   []DatabaseParameterGroupsResponse `json:"contents,omitempty"`
	TotalCount int32                                      `json:"totalCount,omitempty"`
}

type ListResponseDatabaseStroagesResponse struct {
	Contents   []DatabaseStroagesResponse `json:"contents,omitempty"`
	TotalCount int32                               `json:"totalCount,omitempty"`
}

type ListResponseDbBackupsDomain struct {
	Contents   []DbBackupsDomain `json:"contents,omitempty"`
	TotalCount int32                      `json:"totalCount,omitempty"`
}

type ListResponseDbExtEventItemResponse struct {
	Contents   []DbExtEventItemResponse `json:"contents,omitempty"`
	TotalCount int32                             `json:"totalCount,omitempty"`
}

type ListResponseDbServerGroupsResponse struct {
	Contents   []DbServerGroupsResponse `json:"contents,omitempty"`
	TotalCount int32                             `json:"totalCount,omitempty"`
}

type ListResponseRedisListItemResponse struct {
	Contents   []RedisListItemResponse `json:"contents,omitempty"`
	TotalCount int32                            `json:"totalCount,omitempty"`
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

type RedisBackupConfigResponse struct {
	FullBackupConfig *RedisFullBackupConfigResponse `json:"fullBackupConfig,omitempty"`
}

type RedisBlockStorageGroupCreateRequest struct {
	BlockStorageSize int32  `json:"blockStorageSize"`
	BlockStorageType string `json:"blockStorageType,omitempty"`
}

type RedisBlockStorageGroupDetailResponse struct {
	BlockStorageGroupId  string `json:"blockStorageGroupId,omitempty"`
	BlockStorageName     string `json:"blockStorageName,omitempty"`
	BlockStorageRoleType string `json:"blockStorageRoleType,omitempty"`
	BlockStorageSize     int32  `json:"blockStorageSize,omitempty"`
	BlockStorageType     string `json:"blockStorageType,omitempty"`
}

type RedisClusterCreateRequest struct {
	ImageId            string                                    `json:"imageId"`
	NatEnabled         *bool                                     `json:"natEnabled,omitempty"`
	RedisInitialConfig *RedisInitialConfigCreateRequest `json:"redisInitialConfig"`
	RedisName          string                                    `json:"redisName"`
	RedisServerGroup   *RedisServerGroupCreateRequest   `json:"redisServerGroup"`
	SecurityGroupIds   []string                                  `json:"securityGroupIds"`
	ServiceZoneId      string                                    `json:"serviceZoneId"`
	ShardsCount        int32                                     `json:"shardsCount"`
	ShardsReplicaCount int32                                     `json:"shardsReplicaCount"`
	SubnetId           string                                    `json:"subnetId"`
	Tags               []TagRequest                     `json:"tags,omitempty"`
	Timezone           string                                    `json:"timezone,omitempty"`
}

type RedisClusterDetailResponse struct {
	ProjectId          string                                            `json:"projectId,omitempty"`
	BackupConfig       *RedisBackupConfigResponse               `json:"backupConfig,omitempty"`
	BlockId            string                                            `json:"blockId,omitempty"`
	Contract           *DbClusterContractDetailResponse         `json:"contract,omitempty"`
	DatabaseVersion    string                                            `json:"databaseVersion,omitempty"`
	ImageId            string                                            `json:"imageId,omitempty"`
	Maintenance        *DbClusterMaintenanceResponse            `json:"maintenance,omitempty"`
	RedisId            string                                            `json:"redisId,omitempty"`
	RedisInitialConfig *RedisClusterInitialConfigDetailResponse `json:"redisInitialConfig,omitempty"`
	RedisName          string                                            `json:"redisName,omitempty"`
	RedisServerGroup   *RedisServerGroupDetailResponse          `json:"redisServerGroup,omitempty"`
	RedisState         string                                            `json:"redisState,omitempty"`
	SecurityGroupIds   []string                                          `json:"securityGroupIds,omitempty"`
	ServiceZoneId      string                                            `json:"serviceZoneId,omitempty"`
	SubnetId           string                                            `json:"subnetId,omitempty"`
	Timezone           string                                            `json:"timezone,omitempty"`
	VpcId              string                                            `json:"vpcId,omitempty"`
	CreatedBy          string                                            `json:"createdBy,omitempty"`
	CreatedDt          time.Time                                      `json:"createdDt,omitempty"`
	ModifiedBy         string                                            `json:"modifiedBy,omitempty"`
	ModifiedDt         time.Time                                      `json:"modifiedDt,omitempty"`
}

type RedisClusterInitialConfigDetailResponse struct {
	DatabasePort int32 `json:"databasePort,omitempty"`
}

type RedisClusterSearchApiListRedisClusterOpts struct {
	RedisName optional.String
	Page      optional.Int32
	Size      optional.Int32
	Sort      optional.Interface
}

type RedisConfigurationControllerApiListRedis1Opts struct {
	Region            optional.String
	ServerGroupName   optional.String
	State             optional.String
	VirtualServerName optional.String
	CreatedBy         optional.String
	Page              optional.Int32
	Size              optional.Int32
	Sort              optional.Interface
}

type RedisCreateFullBackupConfigRequest struct {
	BackupRetentionPeriod string `json:"backupRetentionPeriod"`
	BackupStartHour       int32  `json:"backupStartHour"`
	ObjectStorageId       string `json:"objectStorageId"`
}

type RedisCreateRequest struct {
	ImageId             string                                     `json:"imageId"`
	NatEnabled          *bool                                      `json:"natEnabled,omitempty"`
	RedisInitialConfig  *RedisInitialConfigCreateRequest  `json:"redisInitialConfig"`
	RedisName           string                                     `json:"redisName"`
	RedisSentinelServer *RedisSentinelServerCreateRequest `json:"redisSentinelServer,omitempty"`
	RedisServerGroup    *RedisServerGroupCreateRequest    `json:"redisServerGroup"`
	SecurityGroupIds    []string                                   `json:"securityGroupIds"`
	ServiceZoneId       string                                     `json:"serviceZoneId"`
	SubnetId            string                                     `json:"subnetId"`
	Tags                []TagRequest                      `json:"tags,omitempty"`
	Timezone            string                                     `json:"timezone,omitempty"`
}

type RedisDetailResponse struct {
	ProjectId          string                                      `json:"projectId,omitempty"`
	BackupConfig       *RedisBackupConfigResponse         `json:"backupConfig,omitempty"`
	BlockId            string                                      `json:"blockId,omitempty"`
	Contract           *DbClusterContractDetailResponse   `json:"contract,omitempty"`
	DatabaseVersion    string                                      `json:"databaseVersion,omitempty"`
	ImageId            string                                      `json:"imageId,omitempty"`
	Maintenance        *DbClusterMaintenanceResponse      `json:"maintenance,omitempty"`
	RedisId            string                                      `json:"redisId,omitempty"`
	RedisInitialConfig *RedisInitialConfigDetailResponse  `json:"redisInitialConfig,omitempty"`
	RedisName          string                                      `json:"redisName,omitempty"`
	RedisServerGroup   *RedisServerGroupDetailResponse    `json:"redisServerGroup,omitempty"`
	RedisState         string                                      `json:"redisState,omitempty"`
	SecurityGroupIds   []string                                    `json:"securityGroupIds,omitempty"`
	SentinelServer     *RedisSentinelServerDetailResponse `json:"sentinelServer,omitempty"`
	ServiceZoneId      string                                      `json:"serviceZoneId,omitempty"`
	SubnetId           string                                      `json:"subnetId,omitempty"`
	Timezone           string                                      `json:"timezone,omitempty"`
	VpcId              string                                      `json:"vpcId,omitempty"`
	CreatedBy          string                                      `json:"createdBy,omitempty"`
	CreatedDt          time.Time                                `json:"createdDt,omitempty"`
	ModifiedBy         string                                      `json:"modifiedBy,omitempty"`
	ModifiedDt         time.Time                                `json:"modifiedDt,omitempty"`
}

type RedisFullBackupConfigResponse struct {
	BackupRetentionPeriod string `json:"backupRetentionPeriod,omitempty"`
	BackupStartHour       int32  `json:"backupStartHour,omitempty"`
	ObjectStorageBucketId string `json:"objectStorageBucketId,omitempty"`
}

type RedisInitialConfigCreateRequest struct {
	DatabasePort         int32  `json:"databasePort,omitempty"`
	DatabaseUserPassword string `json:"databaseUserPassword"`
}

type RedisInitialConfigDetailResponse struct {
	DatabasePort int32 `json:"databasePort,omitempty"`
	SentinelPort int32 `json:"sentinelPort,omitempty"`
}

type RedisListItemResponse struct {
	ProjectId     string       `json:"projectId,omitempty"`
	BlockId       string       `json:"blockId,omitempty"`
	RedisId       string       `json:"redisId,omitempty"`
	RedisName     string       `json:"redisName,omitempty"`
	RedisState    string       `json:"redisState,omitempty"`
	ServiceZoneId string       `json:"serviceZoneId,omitempty"`
	CreatedBy     string       `json:"createdBy,omitempty"`
	CreatedDt     time.Time `json:"createdDt,omitempty"`
	ModifiedBy    string       `json:"modifiedBy,omitempty"`
	ModifiedDt    time.Time `json:"modifiedDt,omitempty"`
}

type RedisModifyFullBackupConfigRequest struct {
	BackupRetentionPeriod string `json:"backupRetentionPeriod"`
	BackupStartHour       int32  `json:"backupStartHour"`
}

type RedisResizeBlockStoragesRequest struct {
	BlockStorageGroupId string `json:"blockStorageGroupId"`
	BlockStorageSize    int32  `json:"blockStorageSize"`
}

type RedisResizeVirtualServersRequest struct {
	ServerType string `json:"serverType"`
}

type RedisSearchApiListRedisOpts struct {
	RedisName optional.String
	Page      optional.Int32
	Size      optional.Int32
	Sort      optional.Interface
}

type RedisSentinelServerCreateRequest struct {
	AvailabilityZoneName  string `json:"availabilityZoneName,omitempty"`
	SentinelNatPublicIpId string `json:"sentinelNatPublicIpId,omitempty"`
	SentinelPort          int32  `json:"sentinelPort,omitempty"`
	SentinelServerName    string `json:"sentinelServerName"`
}

type RedisSentinelServerDetailResponse struct {
	AvailabilityZoneName string                                          `json:"availabilityZoneName,omitempty"`
	BlockStorages        []RedisBlockStorageGroupDetailResponse `json:"blockStorages,omitempty"`
	EncryptionEnabled    *bool                                           `json:"encryptionEnabled,omitempty"`
	NatPublicIpAddress   string                                          `json:"natPublicIpAddress,omitempty"`
	SentinelServerId     string                                          `json:"sentinelServerId,omitempty"`
	SentinelServerName   string                                          `json:"sentinelServerName,omitempty"`
	SentinelServerState  string                                          `json:"sentinelServerState,omitempty"`
	ServerType           string                                          `json:"serverType,omitempty"`
	SubnetIpAddress      string                                          `json:"subnetIpAddress,omitempty"`
	CreatedBy            string                                          `json:"createdBy,omitempty"`
	CreatedDt            time.Time                                    `json:"createdDt,omitempty"`
	ModifiedBy           string                                          `json:"modifiedBy,omitempty"`
	ModifiedDt           time.Time                                    `json:"modifiedDt,omitempty"`
}

type RedisServerCreateRequest struct {
	AvailabilityZoneName string `json:"availabilityZoneName,omitempty"`
	NatPublicIpId        string `json:"natPublicIpId,omitempty"`
	RedisServerName      string `json:"redisServerName"`
	ServerRoleType       string `json:"serverRoleType"`
}

type RedisServerDetailResponse struct {
	AvailabilityZoneName string       `json:"availabilityZoneName,omitempty"`
	NatPublicIpAddress   string       `json:"natPublicIpAddress,omitempty"`
	RedisServerId        string       `json:"redisServerId,omitempty"`
	RedisServerName      string       `json:"redisServerName,omitempty"`
	RedisServerState     string       `json:"redisServerState,omitempty"`
	ServerRoleType       string       `json:"serverRoleType,omitempty"`
	SubnetIpAddress      string       `json:"subnetIpAddress,omitempty"`
	CreatedBy            string       `json:"createdBy,omitempty"`
	CreatedDt            time.Time `json:"createdDt,omitempty"`
	ModifiedBy           string       `json:"modifiedBy,omitempty"`
	ModifiedDt           time.Time `json:"modifiedDt,omitempty"`
}

type RedisServerGroupCreateRequest struct {
	BlockStorages     []RedisBlockStorageGroupCreateRequest `json:"blockStorages"`
	EncryptionEnabled *bool                                          `json:"encryptionEnabled,omitempty"`
	RedisServers      []RedisServerCreateRequest            `json:"redisServers"`
	ServerType        string                                         `json:"serverType"`
}

type RedisServerGroupDetailResponse struct {
	BlockStorages       []RedisBlockStorageGroupDetailResponse `json:"blockStorages,omitempty"`
	EncryptionEnabled   *bool                                           `json:"encryptionEnabled,omitempty"`
	RedisServers        []RedisServerDetailResponse            `json:"redisServers,omitempty"`
	ServerGroupRoleType string                                          `json:"serverGroupRoleType,omitempty"`
	ServerType          string                                          `json:"serverType,omitempty"`
}

type RequestDomain struct {
	ProjectId       string                       `json:"projectId,omitempty"`
	AwxJobId        string                       `json:"awxJobId,omitempty"`
	CauseId         string                       `json:"causeId,omitempty"`
	ChildRequestIds string                       `json:"childRequestIds,omitempty"`
	ChildRequests   []RequestDomain     `json:"childRequests,omitempty"`
	Data            string                       `json:"data,omitempty"`
	Histories       []RequestHistDomain `json:"histories,omitempty"`
	Id              string                       `json:"id,omitempty"`
	ObjectId        string                       `json:"objectId,omitempty"`
	ObjectScope     string                       `json:"objectScope,omitempty"`
	Operation       string                       `json:"operation,omitempty"`
	ParentId        string                       `json:"parentId,omitempty"`
	PoolId          string                       `json:"poolId,omitempty"`
	RequestBy       string                       `json:"requestBy,omitempty"`
	RequestOrigin   string                       `json:"requestOrigin,omitempty"`
	RequestTime     string                       `json:"requestTime,omitempty"`
	ResponseTime    string                       `json:"responseTime,omitempty"`
	Retried         string                       `json:"retried,omitempty"`
	RootId          string                       `json:"rootId,omitempty"`
	State           string                       `json:"state,omitempty"`
	SystemId        string                       `json:"systemId,omitempty"`
	TopRequestId    string                       `json:"topRequestId,omitempty"`
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

type ResizeStorageRequest struct {
	BlockStorageId   string `json:"blockStorageId"`
	BlockStorageSize int32  `json:"blockStorageSize"`
	VirtualServerId  string `json:"virtualServerId"`
}

type RestoreDatabaseRequest struct {
	ActiveDirectory         *ActiveDirectory                `json:"activeDirectory,omitempty"`
	AvailabilityZoneName    string                                   `json:"availabilityZoneName,omitempty"`
	BackupHistoryId         string                                   `json:"backupHistoryId,omitempty"`
	BackupRecoveryDt        time.Time                             `json:"backupRecoveryDt,omitempty"`
	DeploymentEnvType       string                                   `json:"deploymentEnvType"`
	KibanaVirtualServer     *ElasticsearchEntKibanaNode     `json:"kibanaVirtualServer,omitempty"`
	LicenseKey              string                                   `json:"licenseKey,omitempty"`
	Maintenance             *DatabaseMaintenance            `json:"maintenance,omitempty"`
	MasterVirtualServer     *ElasticSearchMasterNode        `json:"masterVirtualServer,omitempty"`
	ServerGroupName         string                                   `json:"serverGroupName"`
	Tags                    []TagRequest                    `json:"tags,omitempty"`
	VirtualServer           *InstanceSpecWithAddtionalDisks `json:"virtualServer"`
	VirtualServerNamePrefix string                                   `json:"virtualServerNamePrefix"`
}

type SelectReplicaRequest struct {
	InstanceId string `json:"instanceId,omitempty"`
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

type SqlserverDbDiskInfo struct {
	DatabaseName string `json:"databaseName"`
	DriveLetter  string `json:"driveLetter"`
}

type TagRequest struct {
	TagKey   string `json:"tagKey"`
	TagValue string `json:"tagValue,omitempty"`
}

type UpdateBackupSettingRequest struct {
	Backup    *DatabaseBackup `json:"backup,omitempty"`
	UseBackup *bool                    `json:"useBackup"`
}

type UpdateMaintenanceRequest struct {
	Maintenance    *DatabaseMaintenance `json:"maintenance,omitempty"`
	UseMaintenance *bool                         `json:"useMaintenance"`
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
	AvailabilityZoneName      string                              `json:"availabilityZoneName,omitempty"`
	BlockStorages             []DatabaseStroagesResponse `json:"blockStorages,omitempty"`
	ContractDiscountName      string                              `json:"contractDiscountName,omitempty"`
	ContractDiscountProductId string                              `json:"contractDiscountProductId,omitempty"`
	Network                   *NetworkResponse           `json:"network,omitempty"`
	ScaleName                 string                              `json:"scaleName,omitempty"`
	ScaleProductId            string                              `json:"scaleProductId,omitempty"`
	ScaleType                 string                              `json:"scaleType,omitempty"`
	Software                  *SoftwareResponse          `json:"software,omitempty"`
	VirtualServerId           string                              `json:"virtualServerId,omitempty"`
	VirtualServerName         string                              `json:"virtualServerName,omitempty"`
	VirtualServerState        string                              `json:"virtualServerState,omitempty"`
	VirtualServerType         string                              `json:"virtualServerType,omitempty"`
	VpcId                     string                              `json:"vpcId,omitempty"`
	ScaleDescription          string                              `json:"scaleDescription,omitempty"`
	CreatedBy                 string                              `json:"createdBy,omitempty"`
	CreatedDt                 time.Time                        `json:"createdDt,omitempty"`
	ModifiedBy                string                              `json:"modifiedBy,omitempty"`
	ModifiedDt                time.Time                        `json:"modifiedDt,omitempty"`
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
