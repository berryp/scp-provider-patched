package kubernetesengine2

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

type CheckResponse struct {
	Result *bool `json:"result,omitempty"`
}

type ClusterCifsVolumeUpdateV2Request struct {
	CifsVolumeId        string `json:"cifsVolumeId,omitempty"`
	CifsVolumeIdEnabled *bool  `json:"cifsVolumeIdEnabled,omitempty"`
}

type ClusterCifsVolumeUpdateV3Request struct {
	CifsVolumeId string `json:"cifsVolumeId,omitempty"`
}

type ClusterCloudLoggingUpdateV2Request struct {
	CloudLoggingEnabled *bool `json:"cloudLoggingEnabled,omitempty"`
}

type ClusterCreateV2Request struct {
	CifsVolumeId         string                                                        `json:"cifsVolumeId,omitempty"`
	CloudLoggingEnabled  *bool                                                         `json:"cloudLoggingEnabled,omitempty"`
	K8sVersion           string                                                        `json:"k8sVersion"`
	KubernetesEngineName string                                                        `json:"kubernetesEngineName"`
	LbId                 string                                                        `json:"lbId,omitempty"`
	PrivateAclResources  []PrivateEndpointAccessControlResourceVo `json:"privateAclResources,omitempty"`
	PublicAclIpAddress   string                                                        `json:"publicAclIpAddress,omitempty"`
	SecurityGroupId      string                                                        `json:"securityGroupId"`
	SubnetId             string                                                        `json:"subnetId"`
	Tags                 []TagRequest                             `json:"tags,omitempty"`
	VolumeId             string                                                        `json:"volumeId"`
	VpcId                string                                                        `json:"vpcId"`
	ZoneId               string                                                        `json:"zoneId"`
}

type ClusterLoadBalancerUpdateV2Request struct {
	LbId                string `json:"lbId,omitempty"`
	LoadBalancerEnabled *bool  `json:"loadBalancerEnabled,omitempty"`
}

type ClusterLoadBalancerUpdateV3Request struct {
	LoadBalancerId string `json:"loadBalancerId,omitempty"`
}

type ClusterPrivateEndpointAccessControlUpdateV2Request struct {
	PrivateAclResources []PrivateEndpointAccessControlResourceVo `json:"privateAclResources,omitempty"`
}

type ClusterPrivateEndpointAccessControlUpdateV3Request struct {
	PrivateEndpointAccessControlResourceList []PrivateEndpointAccessControlResourceV3Vo `json:"privateEndpointAccessControlResourceList,omitempty"`
}

type ClusterPublicEndpointAccessV3UpdateRequest struct {
	PublicEndpointAccessControlIp string `json:"publicEndpointAccessControlIp,omitempty"`
}

type ClusterUpdateV2Request struct {
	K8sVersion         string `json:"k8sVersion,omitempty"`
	PublicAclIpAddress string `json:"publicAclIpAddress,omitempty"`
}

type ClusterUpgradeV3Request struct {
	UpgradeK8sVersion string `json:"upgradeK8sVersion"`
}

type ClusterV2Response struct {
	ProjectId                   string                                                         `json:"projectId,omitempty"`
	BlockId                     string                                                         `json:"blockId,omitempty"`
	CifsVolumeEnabled           *bool                                                          `json:"cifsVolumeEnabled,omitempty"`
	CifsVolumeId                string                                                         `json:"cifsVolumeId,omitempty"`
	CloudLoggingEnabled         *bool                                                          `json:"cloudLoggingEnabled,omitempty"`
	ClusterPrefix               string                                                         `json:"clusterPrefix,omitempty"`
	K8sVersion                  string                                                         `json:"k8sVersion,omitempty"`
	KubernetesEngineId          string                                                         `json:"kubernetesEngineId,omitempty"`
	KubernetesEngineName        string                                                         `json:"kubernetesEngineName,omitempty"`
	KubernetesEngineStatus      string                                                         `json:"kubernetesEngineStatus,omitempty"`
	LbId                        string                                                         `json:"lbId,omitempty"`
	MaxNodeCount                int32                                                          `json:"maxNodeCount,omitempty"`
	NodeCount                   int32                                                          `json:"nodeCount,omitempty"`
	PrivateAclResources         []PrivateEndpointAccessControlResourceVo1 `json:"privateAclResources,omitempty"`
	PrivateEndpointUrl          string                                                         `json:"privateEndpointUrl,omitempty"`
	PublicAclIpAddress          string                                                         `json:"publicAclIpAddress,omitempty"`
	PublicEndpointUrl           string                                                         `json:"publicEndpointUrl,omitempty"`
	Region                      string                                                         `json:"region,omitempty"`
	SecurityGroupId             string                                                         `json:"securityGroupId,omitempty"`
	SubnetId                    string                                                         `json:"subnetId,omitempty"`
	Upgradable                  *bool                                                          `json:"upgradable,omitempty"`
	VolumeId                    string                                                         `json:"volumeId,omitempty"`
	VpcId                       string                                                         `json:"vpcId,omitempty"`
	ZoneId                      string                                                         `json:"zoneId,omitempty"`
	KubernetesEngineDescription string                                                         `json:"kubernetesEngineDescription,omitempty"`
	CreatedBy                   string                                                         `json:"createdBy,omitempty"`
	CreatedDt                   time.Time                                                   `json:"createdDt,omitempty"`
	ModifiedBy                  string                                                         `json:"modifiedBy,omitempty"`
	ModifiedDt                  time.Time                                                   `json:"modifiedDt,omitempty"`
}

type ClusterV4Response struct {
	ProjectId                                string                                                         `json:"projectId,omitempty"`
	CifsVolumeId                             string                                                         `json:"cifsVolumeId,omitempty"`
	CloudLoggingEnabled                      *bool                                                          `json:"cloudLoggingEnabled,omitempty"`
	ClusterPrefix                            string                                                         `json:"clusterPrefix,omitempty"`
	K8sVersion                               string                                                         `json:"k8sVersion,omitempty"`
	KubernetesEngineId                       string                                                         `json:"kubernetesEngineId,omitempty"`
	KubernetesEngineName                     string                                                         `json:"kubernetesEngineName,omitempty"`
	KubernetesEngineStatus                   string                                                         `json:"kubernetesEngineStatus,omitempty"`
	LoadBalancerId                           string                                                         `json:"loadBalancerId,omitempty"`
	MaxNodeCount                             int32                                                          `json:"maxNodeCount,omitempty"`
	NodeCount                                int32                                                          `json:"nodeCount,omitempty"`
	PrivateEndpointAccessControlResourceList []PrivateEndpointAccessControlResourceVo1 `json:"privateEndpointAccessControlResourceList,omitempty"`
	PrivateEndpointUrl                       string                                                         `json:"privateEndpointUrl,omitempty"`
	PrivateKubeConfigDownloadYn              string                                                         `json:"privateKubeConfigDownloadYn,omitempty"`
	PublicEndpointAccessControlIp            string                                                         `json:"publicEndpointAccessControlIp,omitempty"`
	PublicEndpointUrl                        string                                                         `json:"publicEndpointUrl,omitempty"`
	PublicKubeConfigDownloadYn               string                                                         `json:"publicKubeConfigDownloadYn,omitempty"`
	Region                                   string                                                         `json:"region,omitempty"`
	SecurityGroupId                          string                                                         `json:"securityGroupId,omitempty"`
	SubnetId                                 string                                                         `json:"subnetId,omitempty"`
	Upgradable                               *bool                                                          `json:"upgradable,omitempty"`
	VolumeId                                 string                                                         `json:"volumeId,omitempty"`
	VpcId                                    string                                                         `json:"vpcId,omitempty"`
	ZoneId                                   string                                                         `json:"zoneId,omitempty"`
	K8sVersionDescription                    string                                                         `json:"k8sVersionDescription,omitempty"`
	CreatedBy                                string                                                         `json:"createdBy,omitempty"`
	CreatedDt                                time.Time                                                   `json:"createdDt,omitempty"`
	ModifiedBy                               string                                                         `json:"modifiedBy,omitempty"`
	ModifiedDt                               time.Time                                                   `json:"modifiedDt,omitempty"`
}

type ClustersV2Response struct {
	ProjectId              string       `json:"projectId,omitempty"`
	CifsVolumeId           string       `json:"cifsVolumeId,omitempty"`
	K8sVersion             string       `json:"k8sVersion,omitempty"`
	KubernetesEngineId     string       `json:"kubernetesEngineId,omitempty"`
	KubernetesEngineName   string       `json:"kubernetesEngineName,omitempty"`
	KubernetesEngineStatus string       `json:"kubernetesEngineStatus,omitempty"`
	LoadBalancerId         string       `json:"loadBalancerId,omitempty"`
	NodeCount              int32        `json:"nodeCount,omitempty"`
	Region                 string       `json:"region,omitempty"`
	SecurityGroupId        string       `json:"securityGroupId,omitempty"`
	SubnetId               string       `json:"subnetId,omitempty"`
	VolumeId               string       `json:"volumeId,omitempty"`
	VpcId                  string       `json:"vpcId,omitempty"`
	CreatedBy              string       `json:"createdBy,omitempty"`
	CreatedDt              time.Time `json:"createdDt,omitempty"`
	ModifiedBy             string       `json:"modifiedBy,omitempty"`
	ModifiedDt             time.Time `json:"modifiedDt,omitempty"`
}

type K8sEngineV2ApiListKubernetesEnginesV2Opts struct {
	K8sVersion             optional.Interface
	KubernetesEngineName   optional.String
	KubernetesEngineStatus optional.Interface
	Region                 optional.Interface
	CreatedBy              optional.String
	Page                   optional.Int32
	Size                   optional.Int32
	Sort                   optional.String
}

type K8sTemplateV2ApiListKubernetesVersionV21Opts struct {
	Page optional.Int32
	Size optional.Int32
}

type K8sTemplateV2ApiListUpgradableKubernetesVersionV2Opts struct {
	Page optional.Int32
	Size optional.Int32
	Sort optional.String
}

type K8sVersionWithProjectIdResponse struct {
	ProjectId             string       `json:"projectId,omitempty"`
	EosDate               time.Time `json:"eosDate,omitempty"`
	K8sVersion            string       `json:"k8sVersion,omitempty"`
	K8sVersionDescription string       `json:"k8sVersionDescription,omitempty"`
}

type NodePoolAdvancedSettingsVo struct {
	AllowedUnsafeSysctls string `json:"allowedUnsafeSysctls"`
	ContainerLogMaxFiles int32  `json:"containerLogMaxFiles"`
	ContainerLogMaxSize  int32  `json:"containerLogMaxSize"`
	ImageGcHighThreshold int32  `json:"imageGcHighThreshold"`
	ImageGcLowThreshold  int32  `json:"imageGcLowThreshold"`
	MaxPods              int32  `json:"maxPods"`
	PodMaxPids           int32  `json:"podMaxPids"`
}

type NodePoolCreateV2Request struct {
	AutoRecovery         *bool  `json:"autoRecovery,omitempty"`
	AutoScale            *bool  `json:"autoScale,omitempty"`
	AvailabilityZoneName string `json:"availabilityZoneName,omitempty"`
	ContractId           string `json:"contractId"`
	DesiredNodeCount     int32  `json:"desiredNodeCount"`
	ImageId              string `json:"imageId"`
	MaxNodeCount         int32  `json:"maxNodeCount,omitempty"`
	MinNodeCount         int32  `json:"minNodeCount,omitempty"`
	NodePoolName         string `json:"nodePoolName"`
	ProductGroupId       string `json:"productGroupId"`
	ScaleId              string `json:"scaleId"`
	ServiceLevelId       string `json:"serviceLevelId"`
	StorageId            string `json:"storageId"`
	StorageSize          string `json:"storageSize"`
}

type NodePoolCreateV3Request struct {
	AutoRecovery         *bool  `json:"autoRecovery,omitempty"`
	AutoScale            *bool  `json:"autoScale,omitempty"`
	AvailabilityZoneName string `json:"availabilityZoneName,omitempty"`
	ContractName         string `json:"contractName"`
	DesiredNodeCount     int32  `json:"desiredNodeCount"`
	ImageId              string `json:"imageId"`
	MaxNodeCount         int32  `json:"maxNodeCount,omitempty"`
	MinNodeCount         int32  `json:"minNodeCount,omitempty"`
	NodePoolName         string `json:"nodePoolName"`
	ScaleName            string `json:"scaleName"`
	ServiceLevelName     string `json:"serviceLevelName"`
	StorageName          string `json:"storageName"`
	StorageSize          string `json:"storageSize"`
}

type NodePoolCreateV4Request struct {
	AdvancedSettings     *NodePoolAdvancedSettingsVo `json:"advancedSettings,omitempty"`
	AutoRecovery         *bool                                            `json:"autoRecovery,omitempty"`
	AutoScale            *bool                                            `json:"autoScale,omitempty"`
	AvailabilityZoneName string                                           `json:"availabilityZoneName,omitempty"`
	ContractName         string                                           `json:"contractName"`
	DesiredNodeCount     int32                                            `json:"desiredNodeCount"`
	EncryptEnabled       *bool                                            `json:"encryptEnabled,omitempty"`
	ImageId              string                                           `json:"imageId"`
	Labels               []NodePoolLabelVo           `json:"labels,omitempty"`
	MaxNodeCount         int32                                            `json:"maxNodeCount,omitempty"`
	MinNodeCount         int32                                            `json:"minNodeCount,omitempty"`
	NodePoolName         string                                           `json:"nodePoolName"`
	ScaleName            string                                           `json:"scaleName"`
	ServerType           string                                           `json:"serverType"`
	ServiceLevelName     string                                           `json:"serviceLevelName"`
	StorageName          string                                           `json:"storageName"`
	StorageSize          string                                           `json:"storageSize"`
	Taints               []NodePoolTaintVo           `json:"taints,omitempty"`
}

type NodePoolLabelVo struct {
	Key   string `json:"key"`
	Value string `json:"value,omitempty"`
}

type NodePoolLabelsUpdateV2Request struct {
	Labels []NodePoolLabelVo `json:"labels,omitempty"`
}

type NodePoolTaintVo struct {
	Effect string `json:"effect"`
	Key    string `json:"key"`
	Value  string `json:"value,omitempty"`
}

type NodePoolTaintsUpdateV2Request struct {
	Taints []NodePoolTaintVo `json:"taints,omitempty"`
}

type NodePoolUpdateV2Request struct {
	AutoRecovery     *bool  `json:"autoRecovery,omitempty"`
	AutoScale        *bool  `json:"autoScale,omitempty"`
	ContractId       string `json:"contractId"`
	DesiredNodeCount int32  `json:"desiredNodeCount"`
	ImageId          string `json:"imageId"`
	MaxNodeCount     int32  `json:"maxNodeCount,omitempty"`
	MinNodeCount     int32  `json:"minNodeCount,omitempty"`
	NodePoolName     string `json:"nodePoolName"`
	ProductGroupId   string `json:"productGroupId"`
	ScaleId          string `json:"scaleId"`
	ServiceLevelId   string `json:"serviceLevelId"`
	StorageId        string `json:"storageId"`
	StorageSize      string `json:"storageSize"`
}

type NodePoolUpdateV3Request struct {
	AutoRecovery     *bool  `json:"autoRecovery,omitempty"`
	AutoScale        *bool  `json:"autoScale,omitempty"`
	DesiredNodeCount int32  `json:"desiredNodeCount"`
	MaxNodeCount     int32  `json:"maxNodeCount,omitempty"`
	MinNodeCount     int32  `json:"minNodeCount,omitempty"`
	NodePoolName     string `json:"nodePoolName"`
}

type NodePoolUpdateV4Request struct {
	AutoRecovery     *bool `json:"autoRecovery,omitempty"`
	AutoScale        *bool `json:"autoScale,omitempty"`
	DesiredNodeCount int32 `json:"desiredNodeCount"`
	MaxNodeCount     int32 `json:"maxNodeCount,omitempty"`
	MinNodeCount     int32 `json:"minNodeCount,omitempty"`
}

type NodePoolUpgradeV3Request struct {
	UpgradeImageId string `json:"upgradeImageId,omitempty"`
}

type NodePoolV2ApiListNodePoolsV2Opts struct {
	NodePoolName optional.String
	CreatedBy    optional.String
	Page         optional.Int32
	Size         optional.Int32
	Sort         optional.String
}

type NodePoolV2Response struct {
	ProjectId        string                                           `json:"projectId,omitempty"`
	AdvancedSettings *NodePoolAdvancedSettingsVo `json:"advancedSettings,omitempty"`
	AutoRecovery     *bool                                            `json:"autoRecovery,omitempty"`
	AutoScale        *bool                                            `json:"autoScale,omitempty"`
	ContractId       string                                           `json:"contractId,omitempty"`
	CurrentNodeCount int32                                            `json:"currentNodeCount,omitempty"`
	DesiredNodeCount int32                                            `json:"desiredNodeCount,omitempty"`
	EncryptEnabled   *bool                                            `json:"encryptEnabled,omitempty"`
	ImageId          string                                           `json:"imageId,omitempty"`
	InProgress       *bool                                            `json:"inProgress,omitempty"`
	K8sVersion       string                                           `json:"k8sVersion,omitempty"`
	Labels           []NodePoolLabelVo           `json:"labels,omitempty"`
	MaxNodeCount     int32                                            `json:"maxNodeCount,omitempty"`
	MinNodeCount     int32                                            `json:"minNodeCount,omitempty"`
	NodePoolId       string                                           `json:"nodePoolId,omitempty"`
	NodePoolName     string                                           `json:"nodePoolName,omitempty"`
	NodePoolStatus   string                                           `json:"nodePoolStatus,omitempty"`
	ProductGroupId   string                                           `json:"productGroupId,omitempty"`
	ScaleId          string                                           `json:"scaleId,omitempty"`
	ServiceLevelId   string                                           `json:"serviceLevelId,omitempty"`
	StorageId        string                                           `json:"storageId,omitempty"`
	StorageSize      string                                           `json:"storageSize,omitempty"`
	Taints           []NodePoolTaintVo           `json:"taints,omitempty"`
	Upgradable       *bool                                            `json:"upgradable,omitempty"`
	CreatedBy        string                                           `json:"createdBy,omitempty"`
	CreatedDt        time.Time                                     `json:"createdDt,omitempty"`
	ModifiedBy       string                                           `json:"modifiedBy,omitempty"`
	ModifiedDt       time.Time                                     `json:"modifiedDt,omitempty"`
}

type NodePoolsV2Response struct {
	ProjectId            string       `json:"projectId,omitempty"`
	AutoRecovery         *bool        `json:"autoRecovery,omitempty"`
	AutoScale            *bool        `json:"autoScale,omitempty"`
	AvailabilityZoneName string       `json:"availabilityZoneName,omitempty"`
	ContractId           string       `json:"contractId,omitempty"`
	CurrentNodeCount     int32        `json:"currentNodeCount,omitempty"`
	DesiredNodeCount     int32        `json:"desiredNodeCount,omitempty"`
	ImageId              string       `json:"imageId,omitempty"`
	InProgress           *bool        `json:"inProgress,omitempty"`
	K8sVersion           string       `json:"k8sVersion,omitempty"`
	MaxNodeCount         int32        `json:"maxNodeCount,omitempty"`
	MinNodeCount         int32        `json:"minNodeCount,omitempty"`
	NodePoolId           string       `json:"nodePoolId,omitempty"`
	NodePoolName         string       `json:"nodePoolName,omitempty"`
	NodePoolState        string       `json:"nodePoolState,omitempty"`
	OsType               string       `json:"osType,omitempty"`
	ProductGroupId       string       `json:"productGroupId,omitempty"`
	Region               string       `json:"region,omitempty"`
	ScaleId              string       `json:"scaleId,omitempty"`
	ServiceLevelId       string       `json:"serviceLevelId,omitempty"`
	StorageId            string       `json:"storageId,omitempty"`
	StorageSize          string       `json:"storageSize,omitempty"`
	Upgradable           *bool        `json:"upgradable,omitempty"`
	CreatedBy            string       `json:"createdBy,omitempty"`
	CreatedDt            time.Time `json:"createdDt,omitempty"`
	ModifiedBy           string       `json:"modifiedBy,omitempty"`
	ModifiedDt           time.Time `json:"modifiedDt,omitempty"`
}

type NodeV2ApiListNodesV2Opts struct {
	NodeName  optional.String
	CreatedBy optional.String
	Page      optional.Int32
	Size      optional.Int32
	Sort      optional.String
}

type NodesV2Response struct {
	ProjectId  string       `json:"projectId,omitempty"`
	InstanceId string       `json:"instanceId,omitempty"`
	K8sVersion string       `json:"k8sVersion,omitempty"`
	NodeId     string       `json:"nodeId,omitempty"`
	NodeName   string       `json:"nodeName,omitempty"`
	NodePoolId string       `json:"nodePoolId,omitempty"`
	NodeStatus string       `json:"nodeStatus,omitempty"`
	CreatedBy  string       `json:"createdBy,omitempty"`
	CreatedDt  time.Time `json:"createdDt,omitempty"`
}

type PageResponseClustersV2Response struct {
	Contents   []ClustersV2Response `json:"contents,omitempty"`
	TotalCount int64                                     `json:"totalCount,omitempty"`
}

type PageResponseK8sVersionWithProjectIdResponse struct {
	Contents   []K8sVersionWithProjectIdResponse `json:"contents,omitempty"`
	TotalCount int64                                                  `json:"totalCount,omitempty"`
}

type PageResponseNodePoolsV2Response struct {
	Contents   []NodePoolsV2Response `json:"contents,omitempty"`
	TotalCount int64                                      `json:"totalCount,omitempty"`
}

type PageResponseNodesV2Response struct {
	Contents   []NodesV2Response `json:"contents,omitempty"`
	TotalCount int64                                  `json:"totalCount,omitempty"`
}

type PageResponseUpgradableK8sVersionsResponse struct {
	Contents   []UpgradableK8sVersionsResponse `json:"contents,omitempty"`
	TotalCount int64                                                `json:"totalCount,omitempty"`
}

type PrivateEndpointAccessControlResourceV3Vo struct {
	ResourceId   string `json:"resourceId"`
	ResourceName string `json:"resourceName"`
	ResourceType string `json:"resourceType"`
}

type PrivateEndpointAccessControlResourceVo struct {
	Id    string `json:"id"`
	Type_ string `json:"type"`
	Value string `json:"value"`
}

type PrivateEndpointAccessControlResourceVo1 struct {
	ResourceID   string `json:"resourceID,omitempty"`
	ResourceName string `json:"resourceName,omitempty"`
	ResourceType string `json:"resourceType,omitempty"`
}

type TagRequest struct {
	TagKey   string `json:"tagKey"`
	TagValue string `json:"tagValue,omitempty"`
}

type UpgradableK8sVersionsResponse struct {
	ProjectId  string `json:"projectId,omitempty"`
	K8sVersion string `json:"k8sVersion,omitempty"`
	Upgradable *bool  `json:"upgradable,omitempty"`
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
