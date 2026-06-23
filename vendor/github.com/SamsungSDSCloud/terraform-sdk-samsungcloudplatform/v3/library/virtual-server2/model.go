package virtualserver2

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

type DetailVirtualServerV3Response struct {
	ProjectId                 string                                       `json:"projectId,omitempty"`
	AutoscalingEnabled        *bool                                        `json:"autoscalingEnabled,omitempty"`
	AvailabilityZoneName      string                                       `json:"availabilityZoneName,omitempty"`
	BlockId                   string                                       `json:"blockId,omitempty"`
	BlockStorageIds           []string                                     `json:"blockStorageIds,omitempty"`
	Contract                  string                                       `json:"contract,omitempty"`
	ContractEndDate           string                                       `json:"contractEndDate,omitempty"`
	ContractId                string                                       `json:"contractId,omitempty"`
	ContractStartDate         string                                       `json:"contractStartDate,omitempty"`
	DeletionProtectionEnabled *bool                                        `json:"deletionProtectionEnabled,omitempty"`
	DnsEnabled                *bool                                        `json:"dnsEnabled,omitempty"`
	EncryptEnabled            *bool                                        `json:"encryptEnabled,omitempty"`
	ImageId                   string                                       `json:"imageId,omitempty"`
	InitialScriptContent      string                                       `json:"initialScriptContent,omitempty"`
	Ip                        string                                       `json:"ip,omitempty"`
	IsDr                      *bool                                        `json:"isDr,omitempty"`
	KeyPairId                 string                                       `json:"keyPairId,omitempty"`
	NextContractEndDate       string                                       `json:"nextContractEndDate,omitempty"`
	NextContractId            string                                       `json:"nextContractId,omitempty"`
	NicIds                    []string                                     `json:"nicIds,omitempty"`
	OsType                    string                                       `json:"osType,omitempty"`
	PlacementGroupId          string                                       `json:"placementGroupId,omitempty"`
	ProductGroupId            string                                       `json:"productGroupId,omitempty"`
	RoleId                    string                                       `json:"roleId,omitempty"`
	SecurityGroupIds          []SecurityGroupIdsResponse `json:"securityGroupIds,omitempty"`
	ServerGroupId             string                                       `json:"serverGroupId,omitempty"`
	ServerType                string                                       `json:"serverType,omitempty"`
	ServerTypeId              string                                       `json:"serverTypeId,omitempty"`
	ServiceZoneId             string                                       `json:"serviceZoneId,omitempty"`
	ServicedFor               string                                       `json:"servicedFor,omitempty"`
	ServicedGroupFor          string                                       `json:"servicedGroupFor,omitempty"`
	VdcId                     string                                       `json:"vdcId,omitempty"`
	VirtualServerDrId         string                                       `json:"virtualServerDrId,omitempty"`
	VirtualServerId           string                                       `json:"virtualServerId,omitempty"`
	VirtualServerName         string                                       `json:"virtualServerName,omitempty"`
	VirtualServerState        string                                       `json:"virtualServerState,omitempty"`
	VpcId                     string                                       `json:"vpcId,omitempty"`
	CreatedBy                 string                                       `json:"createdBy,omitempty"`
	CreatedDt                 time.Time                                 `json:"createdDt,omitempty"`
	ModifiedBy                string                                       `json:"modifiedBy,omitempty"`
	ModifiedDt                time.Time                                 `json:"modifiedDt,omitempty"`
}

type ListResponseNicResponse struct {
	Contents   []NicResponse `json:"contents,omitempty"`
	TotalCount int32                           `json:"totalCount,omitempty"`
}

type ListResponseVirtualServersResponse struct {
	Contents   []VirtualServersResponse `json:"contents,omitempty"`
	TotalCount int32                                      `json:"totalCount,omitempty"`
}

type NicResponse struct {
	AttachableNat *bool  `json:"attachableNat,omitempty"`
	Ip            string `json:"ip,omitempty"`
	MacAddress    string `json:"macAddress,omitempty"`
	NatIp         string `json:"natIp,omitempty"`
	NatState      string `json:"natState,omitempty"`
	NicId         string `json:"nicId,omitempty"`
	Purpose       string `json:"purpose,omitempty"`
	SubnetId      string `json:"subnetId,omitempty"`
	SubnetType    string `json:"subnetType,omitempty"`
	VdcSubnetId   string `json:"vdcSubnetId,omitempty"`
}

type RdpAccountResponse struct {
	RdpUserId       string `json:"rdpUserId,omitempty"`
	RdpUserPassword string `json:"rdpUserPassword,omitempty"`
}

type SecurityGroupIdsResponse struct {
	SecurityGroupId          string `json:"securityGroupId,omitempty"`
	SecurityGroupMemberState string `json:"securityGroupMemberState,omitempty"`
}

type TagRequest struct {
	TagKey   string `json:"tagKey"`
	TagValue string `json:"tagValue,omitempty"`
}

type VirtualServerAttachNatIpRequest struct {
	PublicIpId string `json:"publicIpId,omitempty"`
}

type VirtualServerChangeAnsiblePasswordRequest struct {
	AnsiblePassword string `json:"ansiblePassword"`
}

type VirtualServerCreateBlockStorageV3Request struct {
	BlockStorageName string `json:"blockStorageName"`
	DiskSize         int32  `json:"diskSize"`
	DiskType         string `json:"diskType,omitempty"`
	EncryptEnabled   *bool  `json:"encryptEnabled,omitempty"`
}

type VirtualServerCreateInitialScriptRequest struct {
	EncodingType         string `json:"encodingType,omitempty"`
	InitialScriptContent string `json:"initialScriptContent,omitempty"`
	InitialScriptShell   string `json:"initialScriptShell"`
	InitialScriptType    string `json:"initialScriptType"`
}

type VirtualServerCreateLocalSubnetNicRequest struct {
	LocalSubnetIpAddress string `json:"localSubnetIpAddress,omitempty"`
	SubnetId             string `json:"subnetId"`
}

type VirtualServerCreateNicRequest struct {
	InternalIpAddress string `json:"internalIpAddress,omitempty"`
	NatEnabled        *bool  `json:"natEnabled,omitempty"`
	PublicIpId        string `json:"publicIpId,omitempty"`
	SubnetId          string `json:"subnetId"`
}

type VirtualServerCreateOsCredentialRequest struct {
	OsUserId       string `json:"osUserId,omitempty"`
	OsUserPassword string `json:"osUserPassword"`
}

type VirtualServerCreateV3Request struct {
	AvailabilityZoneName      string                                                       `json:"availabilityZoneName,omitempty"`
	BlockStorage              *VirtualServerCreateBlockStorageV3Request  `json:"blockStorage"`
	DeletionProtectionEnabled *bool                                                        `json:"deletionProtectionEnabled,omitempty"`
	DnsEnabled                *bool                                                        `json:"dnsEnabled,omitempty"`
	ExtraBlockStorages        []VirtualServerCreateBlockStorageV3Request `json:"extraBlockStorages,omitempty"`
	ImageId                   string                                                       `json:"imageId"`
	InitialScript             *VirtualServerCreateInitialScriptRequest   `json:"initialScript,omitempty"`
	LocalSubnet               *VirtualServerCreateLocalSubnetNicRequest  `json:"localSubnet,omitempty"`
	Nic                       *VirtualServerCreateNicRequest             `json:"nic"`
	OsAdmin                   *VirtualServerCreateOsCredentialRequest    `json:"osAdmin"`
	PlacementGroupId          string                                                       `json:"placementGroupId,omitempty"`
	RoleId                    string                                                       `json:"roleId,omitempty"`
	SecurityGroupIds          []string                                                     `json:"securityGroupIds"`
	ServerGroupId             string                                                       `json:"serverGroupId,omitempty"`
	ServerType                string                                                       `json:"serverType"`
	ServiceZoneId             string                                                       `json:"serviceZoneId"`
	Tags                      []TagRequest                               `json:"tags,omitempty"`
	VirtualServerName         string                                                       `json:"virtualServerName"`
}

type VirtualServerCreateV4Request struct {
	AvailabilityZoneName      string                                                       `json:"availabilityZoneName,omitempty"`
	BlockStorage              *VirtualServerCreateBlockStorageV3Request  `json:"blockStorage"`
	DeletionProtectionEnabled *bool                                                        `json:"deletionProtectionEnabled,omitempty"`
	ExtraBlockStorages        []VirtualServerCreateBlockStorageV3Request `json:"extraBlockStorages,omitempty"`
	ImageId                   string                                                       `json:"imageId"`
	InitialScript             *VirtualServerCreateInitialScriptRequest   `json:"initialScript,omitempty"`
	KeyPairId                 string                                                       `json:"keyPairId"`
	LocalSubnet               *VirtualServerCreateLocalSubnetNicRequest  `json:"localSubnet,omitempty"`
	Nic                       *VirtualServerCreateNicRequest             `json:"nic"`
	PlacementGroupId          string                                                       `json:"placementGroupId,omitempty"`
	RoleId                    string                                                       `json:"roleId,omitempty"`
	SecurityGroupIds          []string                                                     `json:"securityGroupIds"`
	ServerGroupId             string                                                       `json:"serverGroupId,omitempty"`
	ServerType                string                                                       `json:"serverType"`
	ServiceZoneId             string                                                       `json:"serviceZoneId"`
	Tags                      []TagRequest                               `json:"tags,omitempty"`
	VirtualServerName         string                                                       `json:"virtualServerName"`
}

type VirtualServerDecryptRdpAccountRequest struct {
	PrivateKey string `json:"privateKey"`
}

type VirtualServerDeletionProtectionEnabledUpdateRequest struct {
	DeletionProtectionEnabled *bool `json:"deletionProtectionEnabled"`
}

type VirtualServerResizeV3Request struct {
	ServerType string `json:"serverType"`
}

type VirtualServerSecurityGroupUpdateRequest struct {
	SecurityGroupId string `json:"securityGroupId"`
}

type VirtualServerSubnetIpUpdateRequest struct {
	InternalIpAddress string `json:"internalIpAddress,omitempty"`
	SubnetId          string `json:"subnetId"`
}

type VirtualServerV2ApiListVirtualServers2Opts struct {
	AutoScalingGroupId   optional.String
	AutoscalingEnabled   optional.Bool
	ServerGroupId        optional.String
	ServiceZoneId        optional.String
	ServicedForList      optional.Interface
	ServicedGroupForList optional.Interface
	VirtualServerName    optional.String
	Page                 optional.Int32
	Size                 optional.Int32
	Sort                 optional.Interface
}

type VirtualServersResponse struct {
	ImageId            string            `json:"imageId,omitempty"`
	Ip                 string            `json:"ip,omitempty"`
	Properties         map[string]string `json:"properties,omitempty"`
	ServerGroupId      string            `json:"serverGroupId,omitempty"`
	ServicedFor        string            `json:"servicedFor,omitempty"`
	ServicedGroupFor   string            `json:"servicedGroupFor,omitempty"`
	VirtualServerId    string            `json:"virtualServerId,omitempty"`
	VirtualServerName  string            `json:"virtualServerName,omitempty"`
	VirtualServerState string            `json:"virtualServerState,omitempty"`
	CreatedBy          string            `json:"createdBy,omitempty"`
	CreatedDt          time.Time      `json:"createdDt,omitempty"`
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
