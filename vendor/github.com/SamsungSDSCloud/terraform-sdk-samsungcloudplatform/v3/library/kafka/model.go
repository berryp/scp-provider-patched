package kafka

import (
	"github.com/antihax/optional"
	"time"
)

type APIKey struct {
	Key    string
	Prefix string
}

type AddNodeKafkaRequest struct {
	AvailabilityZoneConfig  *KafkaAvailabilityZoneConfig `json:"availabilityZoneConfig,omitempty"`
	Backup                  *DatabaseBackup              `json:"backup,omitempty"`
	BlockId                 string                                `json:"blockId"`
	BrokerNodeCount         int32                                 `json:"brokerNodeCount"`
	DbPort                  int32                                 `json:"dbPort,omitempty"`
	DeploymentEnvType       string                                `json:"deploymentEnvType"`
	ImageId                 string                                `json:"imageId"`
	KafkaAkhqNode           *KafkaAkhqNode               `json:"kafkaAkhqNode,omitempty"`
	KafkaAkhqPort           int32                                 `json:"kafkaAkhqPort,omitempty"`
	KafkaPort               int32                                 `json:"kafkaPort"`
	KafkaZookeeperNode      *KafkaZookeeperNode          `json:"kafkaZookeeperNode,omitempty"`
	Maintenance             *DatabaseMaintenance         `json:"maintenance,omitempty"`
	Network                 *DatabaseNetwork             `json:"network"`
	ProductGroupId          string                                `json:"productGroupId"`
	SecurityGroupIds        []string                              `json:"securityGroupIds"`
	ServerGroupName         string                                `json:"serverGroupName"`
	ServiceZoneId           string                                `json:"serviceZoneId"`
	Tags                    []TagRequest                 `json:"tags,omitempty"`
	Timezone                string                                `json:"timezone"`
	VirtualServer           *InstanceSpec                `json:"virtualServer,omitempty"`
	VirtualServerNamePrefix string                                `json:"virtualServerNamePrefix"`
	ZookeeperPort           int32                                 `json:"zookeeperPort"`
}

type AkhqInitialConfigCreateRequest struct {
	AkhqAccount  string `json:"akhqAccount"`
	AkhqPassword string `json:"akhqPassword"`
}

type AkhqInitialConfigDetailResponse struct {
	AkhqAccount string `json:"akhqAccount,omitempty"`
	AkhqPort    int32  `json:"akhqPort,omitempty"`
}

type AkhqNodeGroupCreateRequest struct {
	AkhqAvailabilityZoneName string `json:"akhqAvailabilityZoneName,omitempty"`
	AkhqNodeName             string `json:"akhqNodeName"`
	NatPublicIpId            string `json:"natPublicIpId,omitempty"`
}

type AkhqNodeGroupDetailResponse struct {
	AkhqNodeId           string `json:"akhqNodeId,omitempty"`
	AkhqNodeName         string `json:"akhqNodeName,omitempty"`
	AkhqNodeState        string `json:"akhqNodeState,omitempty"`
	AvailabilityZoneName string `json:"availabilityZoneName,omitempty"`
	BlockStorageGroupId  string `json:"blockStorageGroupId,omitempty"`
	BlockStorageName     string `json:"blockStorageName,omitempty"`
	BlockStorageRoleType string `json:"blockStorageRoleType,omitempty"`
	BlockStorageSize     int32  `json:"blockStorageSize,omitempty"`
	BlockStorageType     string `json:"blockStorageType,omitempty"`
	NodeRoleType         string `json:"nodeRoleType,omitempty"`
	ServerType           string `json:"serverType,omitempty"`
	SubnetIpAddress      string `json:"subnetIpAddress,omitempty"`
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

type BrokerInitialConfigCreateRequest struct {
	BrokerPort         int32  `json:"brokerPort,omitempty"`
	BrokerSaslAccount  string `json:"brokerSaslAccount"`
	BrokerSaslPassword string `json:"brokerSaslPassword"`
}

type BrokerInitialConfigDetailResponse struct {
	BrokerPort        int32  `json:"brokerPort,omitempty"`
	BrokerSaslAccount string `json:"brokerSaslAccount,omitempty"`
}

type BrokerNodeCreateRequest struct {
	BrokerNodeName string `json:"brokerNodeName"`
	NatPublicIpId  string `json:"natPublicIpId,omitempty"`
}

type BrokerNodeDetailResponse struct {
	AvailabilityZoneName string       `json:"availabilityZoneName,omitempty"`
	BrokerNodeId         string       `json:"brokerNodeId,omitempty"`
	BrokerNodeName       string       `json:"brokerNodeName,omitempty"`
	BrokerNodeState      string       `json:"brokerNodeState,omitempty"`
	SubnetIpAddress      string       `json:"subnetIpAddress,omitempty"`
	CreatedBy            string       `json:"createdBy,omitempty"`
	CreatedDt            time.Time `json:"createdDt,omitempty"`
	ModifiedBy           string       `json:"modifiedBy,omitempty"`
	ModifiedDt           time.Time `json:"modifiedDt,omitempty"`
}

type BrokerNodeGroupCreateRequest struct {
	BlockStorage *KafkaNodeGroupBlockStorageGroupCreateRequest `json:"blockStorage"`
	BrokerNodes  []BrokerNodeCreateRequest                     `json:"brokerNodes"`
	ServerType   string                                                 `json:"serverType"`
}

type BrokerNodeGroupDetailResponse struct {
	BlockStorages []KafkaNodeGroupBlockStorageGroupDetailResponse `json:"blockStorages,omitempty"`
	BrokerNodes   []BrokerNodeDetailResponse                      `json:"brokerNodes,omitempty"`
	NodeRoleType  string                                                   `json:"nodeRoleType,omitempty"`
	ServerType    string                                                   `json:"serverType,omitempty"`
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

type CreateKafkaRequest struct {
	AvailabilityZoneConfig  *KafkaAvailabilityZoneConfig    `json:"availabilityZoneConfig,omitempty"`
	Backup                  *DatabaseBackup                 `json:"backup,omitempty"`
	BlockId                 string                                   `json:"blockId"`
	BrokerNodeCount         int32                                    `json:"brokerNodeCount"`
	DbPort                  int32                                    `json:"dbPort,omitempty"`
	DeploymentEnvType       string                                   `json:"deploymentEnvType"`
	ImageId                 string                                   `json:"imageId"`
	KafkaAkhqNode           *KafkaAkhqNode                  `json:"kafkaAkhqNode,omitempty"`
	KafkaAkhqPort           int32                                    `json:"kafkaAkhqPort,omitempty"`
	KafkaAkhqUserId         string                                   `json:"kafkaAkhqUserId,omitempty"`
	KafkaAkhqUserPwd        string                                   `json:"kafkaAkhqUserPwd,omitempty"`
	KafkaPort               int32                                    `json:"kafkaPort"`
	KafkaUserId             string                                   `json:"kafkaUserId"`
	KafkaUserPwd            string                                   `json:"kafkaUserPwd"`
	KafkaZookeeperNode      *KafkaZookeeperNode             `json:"kafkaZookeeperNode,omitempty"`
	Maintenance             *DatabaseMaintenance            `json:"maintenance,omitempty"`
	Network                 *DatabaseNetwork                `json:"network"`
	ProductGroupId          string                                   `json:"productGroupId"`
	SecurityGroupIds        []string                                 `json:"securityGroupIds"`
	ServerGroupName         string                                   `json:"serverGroupName"`
	ServiceZoneId           string                                   `json:"serviceZoneId"`
	Tags                    []TagRequest                    `json:"tags,omitempty"`
	Timezone                string                                   `json:"timezone"`
	VirtualServer           *InstanceSpecWithAddtionalDisks `json:"virtualServer"`
	VirtualServerNamePrefix string                                   `json:"virtualServerNamePrefix"`
	ZookeeperPort           int32                                    `json:"zookeeperPort"`
	ZookeeperUserId         string                                   `json:"zookeeperUserId"`
	ZookeeperUserPwd        string                                   `json:"zookeeperUserPwd"`
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

type KafkaAkhqNode struct {
	DataBlockStorageSize int32  `json:"dataBlockStorageSize"`
	ScaleProductId       string `json:"scaleProductId"`
}

type KafkaAvailabilityZoneConfig struct {
	AvailabilityZoneDeploymentType string `json:"availabilityZoneDeploymentType"`
}

type KafkaClusterCreateAvailabilityZoneConfig struct {
	AvailabilityZoneDeploymentType string `json:"availabilityZoneDeploymentType"`
	AvailabilityZoneName           string `json:"availabilityZoneName,omitempty"`
}

type KafkaClusterCreateRequest struct {
	AkhqEnabled            *bool                                              `json:"akhqEnabled,omitempty"`
	AkhqNodeGroup          *AkhqNodeGroupCreateRequest               `json:"akhqNodeGroup,omitempty"`
	AvailabilityZoneConfig *KafkaClusterCreateAvailabilityZoneConfig `json:"availabilityZoneConfig,omitempty"`
	BrokerNodeGroup        *BrokerNodeGroupCreateRequest             `json:"brokerNodeGroup"`
	ImageId                string                                             `json:"imageId"`
	KafkaClusterName       string                                             `json:"kafkaClusterName"`
	KafkaInitialConfig     *KafkaInitialConfigCreateRequest          `json:"kafkaInitialConfig"`
	NatEnabled             *bool                                              `json:"natEnabled,omitempty"`
	SecurityGroupIds       []string                                           `json:"securityGroupIds"`
	ServiceZoneId          string                                             `json:"serviceZoneId"`
	SubnetId               string                                             `json:"subnetId"`
	Tags                   []TagRequest                              `json:"tags,omitempty"`
	Timezone               string                                             `json:"timezone,omitempty"`
	ZookeeperNodeGroup     *ZookeeperNodeGroupCreateRequest          `json:"zookeeperNodeGroup,omitempty"`
}

type KafkaClusterDetailResponse struct {
	ProjectId          string                                     `json:"projectId,omitempty"`
	AkhqNodeGroup      *AkhqNodeGroupDetailResponse      `json:"akhqNodeGroup,omitempty"`
	BlockId            string                                     `json:"blockId,omitempty"`
	BrokerNodeGroup    *BrokerNodeGroupDetailResponse    `json:"brokerNodeGroup,omitempty"`
	Contract           *DbClusterContractDetailResponse  `json:"contract,omitempty"`
	DatabaseVersion    string                                     `json:"databaseVersion,omitempty"`
	ImageId            string                                     `json:"imageId,omitempty"`
	KafkaClusterId     string                                     `json:"kafkaClusterId,omitempty"`
	KafkaClusterName   string                                     `json:"kafkaClusterName,omitempty"`
	KafkaClusterState  string                                     `json:"kafkaClusterState,omitempty"`
	KafkaInitialConfig *KafkaInitialConfigDetailResponse `json:"kafkaInitialConfig,omitempty"`
	Maintenance        *DbClusterMaintenanceResponse     `json:"maintenance,omitempty"`
	NatIpAddress       string                                     `json:"natIpAddress,omitempty"`
	SecurityGroupIds   []string                                   `json:"securityGroupIds,omitempty"`
	ServiceZoneId      string                                     `json:"serviceZoneId,omitempty"`
	SubnetId           string                                     `json:"subnetId,omitempty"`
	Timezone           string                                     `json:"timezone,omitempty"`
	VpcId              string                                     `json:"vpcId,omitempty"`
	ZookeeperNodeGroup *ZookeeperNodeGroupDetailResponse `json:"zookeeperNodeGroup,omitempty"`
	CreatedBy          string                                     `json:"createdBy,omitempty"`
	CreatedDt          time.Time                               `json:"createdDt,omitempty"`
	ModifiedBy         string                                     `json:"modifiedBy,omitempty"`
	ModifiedDt         time.Time                               `json:"modifiedDt,omitempty"`
}

type KafkaClusterListItemResponse struct {
	ProjectId         string       `json:"projectId,omitempty"`
	BlockId           string       `json:"blockId,omitempty"`
	KafkaClusterId    string       `json:"kafkaClusterId,omitempty"`
	KafkaClusterName  string       `json:"kafkaClusterName,omitempty"`
	KafkaClusterState string       `json:"kafkaClusterState,omitempty"`
	ServiceZoneId     string       `json:"serviceZoneId,omitempty"`
	CreatedBy         string       `json:"createdBy,omitempty"`
	CreatedDt         time.Time `json:"createdDt,omitempty"`
	ModifiedBy        string       `json:"modifiedBy,omitempty"`
	ModifiedDt        time.Time `json:"modifiedDt,omitempty"`
}

type KafkaClusterResizeBlockStoragesRequest struct {
	BlockStorageGroupId string `json:"blockStorageGroupId"`
	BlockStorageSize    int32  `json:"blockStorageSize"`
}

type KafkaClusterResizeVirtualServersRequest struct {
	NodeRoleType string `json:"nodeRoleType"`
	ServerType   string `json:"serverType"`
}

type KafkaConfigurationControllerApiListKafkaOpts struct {
	Region            optional.String
	ServerGroupName   optional.String
	State             optional.String
	VirtualServerName optional.String
	CreatedBy         optional.String
	Page              optional.Int32
	Size              optional.Int32
	Sort              optional.Interface
}

type KafkaInitialConfigCreateRequest struct {
	AkhqInitialConfig      *AkhqInitialConfigCreateRequest      `json:"akhqInitialConfig,omitempty"`
	BrokerInitialConfig    *BrokerInitialConfigCreateRequest    `json:"brokerInitialConfig"`
	ZookeeperInitialConfig *ZookeeperInitialConfigCreateRequest `json:"zookeeperInitialConfig"`
}

type KafkaInitialConfigDetailResponse struct {
	AkhqInitialConfig      *AkhqInitialConfigDetailResponse      `json:"akhqInitialConfig,omitempty"`
	BrokerInitialConfig    *BrokerInitialConfigDetailResponse    `json:"brokerInitialConfig,omitempty"`
	ZookeeperInitialConfig *ZookeeperInitialConfigDetailResponse `json:"zookeeperInitialConfig,omitempty"`
}

type KafkaNodeGroupBlockStorageGroupCreateRequest struct {
	BlockStorageSize int32  `json:"blockStorageSize"`
	BlockStorageType string `json:"blockStorageType,omitempty"`
}

type KafkaNodeGroupBlockStorageGroupDetailResponse struct {
	BlockStorageGroupId  string `json:"blockStorageGroupId,omitempty"`
	BlockStorageName     string `json:"blockStorageName,omitempty"`
	BlockStorageRoleType string `json:"blockStorageRoleType,omitempty"`
	BlockStorageSize     int32  `json:"blockStorageSize,omitempty"`
	BlockStorageType     string `json:"blockStorageType,omitempty"`
}

type KafkaSearchApiListKafkaClustersOpts struct {
	KafkaClusterName optional.String
	Page             optional.Int32
	Size             optional.Int32
	Sort             optional.Interface
}

type KafkaZookeeperNode struct {
	AdditionalBlockStorages []DatabaseBlockStorage `json:"additionalBlockStorages,omitempty"`
	DataBlockStorageSize    int32                           `json:"dataBlockStorageSize"`
	DataDiskType            string                          `json:"dataDiskType"`
	EncryptEnabled          *bool                           `json:"encryptEnabled"`
	ScaleProductId          string                          `json:"scaleProductId"`
	ZookeeperNodeCount      int32                           `json:"zookeeperNodeCount"`
}

type ListResponseDatabaseParameterGroupsResponse struct {
	Contents   []DatabaseParameterGroupsResponse `json:"contents,omitempty"`
	TotalCount int32                                      `json:"totalCount,omitempty"`
}

type ListResponseDbExtEventItemResponse struct {
	Contents   []DbExtEventItemResponse `json:"contents,omitempty"`
	TotalCount int32                             `json:"totalCount,omitempty"`
}

type ListResponseDbServerGroupsResponse struct {
	Contents   []DbServerGroupsResponse `json:"contents,omitempty"`
	TotalCount int32                             `json:"totalCount,omitempty"`
}

type ListResponseKafkaClusterListItemResponse struct {
	Contents   []KafkaClusterListItemResponse `json:"contents,omitempty"`
	TotalCount int32                                   `json:"totalCount,omitempty"`
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

type UpdateMaintenanceRequest struct {
	Maintenance    *DatabaseMaintenance `json:"maintenance,omitempty"`
	UseMaintenance *bool                         `json:"useMaintenance"`
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

type ZookeeperInitialConfigCreateRequest struct {
	ZookeeperPort         int32  `json:"zookeeperPort,omitempty"`
	ZookeeperSaslAccount  string `json:"zookeeperSaslAccount"`
	ZookeeperSaslPassword string `json:"zookeeperSaslPassword"`
}

type ZookeeperInitialConfigDetailResponse struct {
	ZookeeperPort        int32  `json:"zookeeperPort,omitempty"`
	ZookeeperSaslAccount string `json:"zookeeperSaslAccount,omitempty"`
}

type ZookeeperNodeCreateRequest struct {
	NatPublicIpId     string `json:"natPublicIpId,omitempty"`
	ZookeeperNodeName string `json:"zookeeperNodeName"`
}

type ZookeeperNodeDetailResponse struct {
	AvailabilityZoneName string       `json:"availabilityZoneName,omitempty"`
	SubnetIpAddress      string       `json:"subnetIpAddress,omitempty"`
	ZookeeperNodeId      string       `json:"zookeeperNodeId,omitempty"`
	ZookeeperNodeName    string       `json:"zookeeperNodeName,omitempty"`
	ZookeeperNodeState   string       `json:"zookeeperNodeState,omitempty"`
	CreatedBy            string       `json:"createdBy,omitempty"`
	CreatedDt            time.Time `json:"createdDt,omitempty"`
	ModifiedBy           string       `json:"modifiedBy,omitempty"`
	ModifiedDt           time.Time `json:"modifiedDt,omitempty"`
}

type ZookeeperNodeGroupCreateRequest struct {
	BlockStorage   *KafkaNodeGroupBlockStorageGroupCreateRequest `json:"blockStorage"`
	ServerType     string                                                 `json:"serverType"`
	ZookeeperNodes []ZookeeperNodeCreateRequest                  `json:"zookeeperNodes"`
}

type ZookeeperNodeGroupDetailResponse struct {
	BlockStorages  []KafkaNodeGroupBlockStorageGroupDetailResponse `json:"blockStorages,omitempty"`
	NodeRoleType   string                                                   `json:"nodeRoleType,omitempty"`
	ServerType     string                                                   `json:"serverType,omitempty"`
	ZookeeperNodes []ZookeeperNodeDetailResponse                   `json:"zookeeperNodes,omitempty"`
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
