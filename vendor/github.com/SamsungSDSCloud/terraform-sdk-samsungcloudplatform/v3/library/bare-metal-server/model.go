package baremetalserver

import (
	"github.com/antihax/optional"
	"time"
)

type APIKey struct {
	Key    string
	Prefix string
}

type AdditionalBlockStorageCreateOpenApiRequest struct {
	BareMetalBlockStorageName string `json:"bareMetalBlockStorageName"`
	BareMetalBlockStorageSize int32  `json:"bareMetalBlockStorageSize"`
	BareMetalBlockStorageType string `json:"bareMetalBlockStorageType"`
	EncryptionEnabled         *bool  `json:"encryptionEnabled,omitempty"`
}

type AdditionalBlockStorageCreateRequest struct {
	BareMetalBlockStorageName   string `json:"bareMetalBlockStorageName"`
	BareMetalBlockStorageSize   int32  `json:"bareMetalBlockStorageSize"`
	BareMetalBlockStorageType   string `json:"bareMetalBlockStorageType,omitempty"`
	BareMetalBlockStorageTypeId string `json:"bareMetalBlockStorageTypeId"`
	EncryptionEnabled           *bool  `json:"encryptionEnabled"`
}

type AsyncResponse struct {
	ProjectId  string `json:"projectId,omitempty"`
	RequestId  string `json:"requestId,omitempty"`
	ResourceId string `json:"resourceId,omitempty"`
}

type AttachLocalSubnetRequest struct {
	BareMetalLocalSubnetId string `json:"bareMetalLocalSubnetId"`
	IpAddress              string `json:"ipAddress,omitempty"`
	IpType                 string `json:"ipType"`
}

type AttachLocalSubnetResponse struct {
	Result string `json:"result,omitempty"`
}

type BareMetalServerAssignPublicNatRequest struct {
	NatEnabled        string `json:"natEnabled"`
	NatIpAddress      string `json:"natIpAddress,omitempty"`
	NatIpAddressType  string `json:"natIpAddressType"`
	PublicIpAddressId string `json:"publicIpAddressId,omitempty"`
}

type BareMetalServerContractPeriodUpdateResponse struct {
	Result string `json:"result,omitempty"`
}

type BareMetalServerCreateLocalSubnetRequest struct {
	BareMetalLocalSubnetIpRange string `json:"bareMetalLocalSubnetIpRange"`
	BareMetalLocalSubnetName    string `json:"bareMetalLocalSubnetName"`
	ServiceZoneId               string `json:"serviceZoneId"`
	TargetProduct               string `json:"targetProduct,omitempty"`
}

type BareMetalServerDeleteProtectionUpdateResponse struct {
	Result string `json:"result,omitempty"`
}

type BareMetalServerDetailResponse struct {
	ProjectId                     string       `json:"projectId,omitempty"`
	AllMountedStorage             *bool        `json:"allMountedStorage,omitempty"`
	BareMetalBlockStorageIds      []string     `json:"bareMetalBlockStorageIds,omitempty"`
	BareMetalLocalSubnetId        string       `json:"bareMetalLocalSubnetId,omitempty"`
	BareMetalLocalSubnetIpAddress string       `json:"bareMetalLocalSubnetIpAddress,omitempty"`
	BareMetalServerId             string       `json:"bareMetalServerId,omitempty"`
	BareMetalServerName           string       `json:"bareMetalServerName,omitempty"`
	BareMetalServerState          string       `json:"bareMetalServerState,omitempty"`
	BlockId                       string       `json:"blockId,omitempty"`
	CheckCriticalError            *bool        `json:"checkCriticalError,omitempty"`
	Contract                      string       `json:"contract,omitempty"`
	ContractEndDate               string       `json:"contractEndDate,omitempty"`
	ContractId                    string       `json:"contractId,omitempty"`
	ContractStartDate             string       `json:"contractStartDate,omitempty"`
	DeletionProtectionEnabled     string       `json:"deletionProtectionEnabled,omitempty"`
	DnsEnabled                    string       `json:"dnsEnabled,omitempty"`
	ErrorCheck                    *bool        `json:"errorCheck,omitempty"`
	ImageId                       string       `json:"imageId,omitempty"`
	InitialScriptContent          string       `json:"initialScriptContent,omitempty"`
	IpAddress                     string       `json:"ipAddress,omitempty"`
	LocalSubnetStatus             string       `json:"localSubnetStatus,omitempty"`
	NatIpAddress                  string       `json:"natIpAddress,omitempty"`
	NextContract                  string       `json:"nextContract,omitempty"`
	NextContractEndDate           string       `json:"nextContractEndDate,omitempty"`
	OsType                        string       `json:"osType,omitempty"`
	OsVersion                     string       `json:"osVersion,omitempty"`
	PlacementGroupName            string       `json:"placementGroupName,omitempty"`
	ProductGroupId                string       `json:"productGroupId,omitempty"`
	ProductType                   string       `json:"productType,omitempty"`
	PublicNatStatus               string       `json:"publicNatStatus,omitempty"`
	ServerType                    string       `json:"serverType,omitempty"`
	ServerTypeId                  string       `json:"serverTypeId,omitempty"`
	ServiceLevelId                string       `json:"serviceLevelId,omitempty"`
	ServiceZoneId                 string       `json:"serviceZoneId,omitempty"`
	SubnetId                      string       `json:"subnetId,omitempty"`
	Timezone                      string       `json:"timezone,omitempty"`
	UseHyperThreading             string       `json:"useHyperThreading,omitempty"`
	VpcId                         string       `json:"vpcId,omitempty"`
	CreatedBy                     string       `json:"createdBy,omitempty"`
	CreatedDt                     time.Time `json:"createdDt,omitempty"`
	ModifiedBy                    string       `json:"modifiedBy,omitempty"`
	ModifiedDt                    time.Time `json:"modifiedDt,omitempty"`
}

type BareMetalServerLocalSubnetOpenApiControllerApiListBareMetalServerLocalSubnetsOpts struct {
	Page optional.Int32
	Size optional.Int32
	Sort optional.Interface
}

type BareMetalServerPublicNatResponse struct {
	Result string `json:"result,omitempty"`
}

type BareMetalServerResponse struct {
	ProjectId            string       `json:"projectId,omitempty"`
	BareMetalServerId    string       `json:"bareMetalServerId,omitempty"`
	BareMetalServerName  string       `json:"bareMetalServerName,omitempty"`
	BareMetalServerState string       `json:"bareMetalServerState,omitempty"`
	BlockId              string       `json:"blockId,omitempty"`
	ImageId              string       `json:"imageId,omitempty"`
	IpAddress            string       `json:"ipAddress,omitempty"`
	ServerTypeId         string       `json:"serverTypeId,omitempty"`
	ServiceZoneId        string       `json:"serviceZoneId,omitempty"`
	CreatedBy            string       `json:"createdBy,omitempty"`
	CreatedDt            time.Time `json:"createdDt,omitempty"`
	ModifiedBy           string       `json:"modifiedBy,omitempty"`
	ModifiedDt           time.Time `json:"modifiedDt,omitempty"`
}

type BareMetalServerSimpleTaskOpenApiControllerApiListBareMetalServersOpts struct {
	BareMetalServerName optional.String
	IpAddress           optional.String
	Page                optional.Int32
	Size                optional.Int32
	Sort                optional.Interface
}

type BareMetalServerStartStopRequest struct {
	BareMetalServerIds []string `json:"bareMetalServerIds,omitempty"`
}

type BaremetalServersTerminateRequest struct {
	BareMetalServerIds []string `json:"bareMetalServerIds,omitempty"`
}

type BasicAuth struct {
	UserName string `json:"userName,omitempty"`
	Password string `json:"password,omitempty"`
}

type BmServerContractPeriodUpdateOpenApiRequest struct {
	Contract string `json:"contract"`
}

type BmServerContractPeriodUpdateRequest struct {
	ContractId string `json:"contractId"`
}

type BmServerCreateOpenApiRequest struct {
	BlockId                   string                                             `json:"blockId"`
	Contract                  string                                             `json:"contract,omitempty"`
	DeletionProtectionEnabled *bool                                              `json:"deletionProtectionEnabled,omitempty"`
	ImageId                   string                                             `json:"imageId"`
	InitScript                string                                             `json:"initScript,omitempty"`
	OsUserId                  string                                             `json:"osUserId"`
	OsUserPassword            string                                             `json:"osUserPassword"`
	ServerDetails             []BmServerDetailsOpenApiRequest `json:"serverDetails"`
	ServiceZoneId             string                                             `json:"serviceZoneId"`
	SubnetId                  string                                             `json:"subnetId"`
	Tags                      []TagRequest                    `json:"tags,omitempty"`
	VpcId                     string                                             `json:"vpcId"`
}

type BmServerCreateRequest struct {
	BlockId                   string                                      `json:"blockId"`
	ContractId                string                                      `json:"contractId"`
	DeletionProtectionEnabled *bool                                       `json:"deletionProtectionEnabled"`
	ImageId                   string                                      `json:"imageId"`
	InitScript                string                                      `json:"initScript,omitempty"`
	OsUserId                  string                                      `json:"osUserId"`
	OsUserPassword            string                                      `json:"osUserPassword"`
	ProductGroupId            string                                      `json:"productGroupId"`
	ServerDetails             []BmServerDetailsRequest `json:"serverDetails"`
	ServiceZoneId             string                                      `json:"serviceZoneId"`
	SubnetId                  string                                      `json:"subnetId"`
	Tags                      []TagRequest             `json:"tags,omitempty"`
	VpcId                     string                                      `json:"vpcId"`
}

type BmServerDeletionProtectionEnabledUpdateRequest struct {
	DeletionProtectionEnabled string `json:"deletionProtectionEnabled"`
}

type BmServerDeletionProtectionRequest struct {
	BareMetalServerId         string `json:"bareMetalServerId"`
	DeletionProtectionEnabled string `json:"deletionProtectionEnabled"`
}

type BmServerDetailsOpenApiRequest struct {
	BareMetalLocalSubnetEnabled   *bool                                                           `json:"bareMetalLocalSubnetEnabled,omitempty"`
	BareMetalLocalSubnetId        string                                                          `json:"bareMetalLocalSubnetId,omitempty"`
	BareMetalLocalSubnetIpAddress string                                                          `json:"bareMetalLocalSubnetIpAddress,omitempty"`
	BareMetalServerName           string                                                          `json:"bareMetalServerName"`
	DnsEnabled                    *bool                                                           `json:"dnsEnabled,omitempty"`
	IpAddress                     string                                                          `json:"ipAddress,omitempty"`
	LocalDiskPartitions           []LocalDiskPartitionCreateRequest            `json:"localDiskPartitions,omitempty"`
	LocalDiskType                 string                                                          `json:"localDiskType,omitempty"`
	NatEnabled                    *bool                                                           `json:"natEnabled,omitempty"`
	PlacementGroupName            string                                                          `json:"placementGroupName,omitempty"`
	PublicIpAddressId             string                                                          `json:"publicIpAddressId,omitempty"`
	ServerType                    string                                                          `json:"serverType"`
	StorageDetails                []AdditionalBlockStorageCreateOpenApiRequest `json:"storageDetails,omitempty"`
	UseHyperThreading             string                                                          `json:"useHyperThreading"`
}

type BmServerDetailsRequest struct {
	BareMetalLocalSubnetEnabled   *bool                                                    `json:"bareMetalLocalSubnetEnabled"`
	BareMetalLocalSubnetId        string                                                   `json:"bareMetalLocalSubnetId,omitempty"`
	BareMetalLocalSubnetIpAddress string                                                   `json:"bareMetalLocalSubnetIpAddress,omitempty"`
	BareMetalServerName           string                                                   `json:"bareMetalServerName"`
	DnsEnabled                    *bool                                                    `json:"dnsEnabled"`
	IpAddress                     string                                                   `json:"ipAddress,omitempty"`
	LocalDiskPartitions           []LocalDiskPartitionCreateRequest     `json:"localDiskPartitions,omitempty"`
	LocalDiskType                 string                                                   `json:"localDiskType,omitempty"`
	NatEnabled                    *bool                                                    `json:"natEnabled"`
	PlacementGroupName            string                                                   `json:"placementGroupName,omitempty"`
	PublicIpAddressId             string                                                   `json:"publicIpAddressId,omitempty"`
	ServerTypeId                  string                                                   `json:"serverTypeId"`
	StorageDetails                []AdditionalBlockStorageCreateRequest `json:"storageDetails"`
	UseHyperThreading             string                                                   `json:"useHyperThreading"`
}

type BmServerLocalSubnetResponse struct {
	BareMetalLocalSubnetId      string       `json:"bareMetalLocalSubnetId,omitempty"`
	BareMetalLocalSubnetIpRange string       `json:"bareMetalLocalSubnetIpRange,omitempty"`
	BareMetalLocalSubnetName    string       `json:"bareMetalLocalSubnetName,omitempty"`
	BlockId                     string       `json:"blockId,omitempty"`
	ServiceZoneId               string       `json:"serviceZoneId,omitempty"`
	CreatedBy                   string       `json:"createdBy,omitempty"`
	CreatedDt                   time.Time `json:"createdDt,omitempty"`
	ModifiedBy                  string       `json:"modifiedBy,omitempty"`
	ModifiedDt                  time.Time `json:"modifiedDt,omitempty"`
}

type BmServerNextContractPeriodUpdateOpenApiRequest struct {
	Contract string `json:"contract"`
}

type BmServerNextContractPeriodUpdateRequest struct {
	ContractId string `json:"contractId"`
}

type BmServersDeletionProtectionEnabledUpdateRequest struct {
	BmServerDeletionProtectionRequest []BmServerDeletionProtectionRequest `json:"bmServerDeletionProtectionRequest,omitempty"`
}

type DetachLocalSubnetOutSo struct {
	Result string `json:"result,omitempty"`
}

type ImErrorResponse struct {
	ErrorCode     string       `json:"errorCode,omitempty"`
	ErrorMessage  string       `json:"errorMessage,omitempty"`
	OperationId   string       `json:"operationId,omitempty"`
	OperationName string       `json:"operationName,omitempty"`
	CreatedDt     time.Time `json:"createdDt,omitempty"`
}

type ImOpenApiControllerApiListImErrorsOpts struct {
	Page optional.Int32
	Size optional.Int32
	Sort optional.Interface
}

type ListResponseBareMetalServerResponse struct {
	Contents   []BareMetalServerResponse `json:"contents,omitempty"`
	TotalCount int32                                        `json:"totalCount,omitempty"`
}

type ListResponseBmServerLocalSubnetResponse struct {
	Contents   []BmServerLocalSubnetResponse `json:"contents,omitempty"`
	TotalCount int32                                            `json:"totalCount,omitempty"`
}

type ListResponseImErrorResponse struct {
	Contents   []ImErrorResponse `json:"contents,omitempty"`
	TotalCount int32                                `json:"totalCount,omitempty"`
}

type LocalDiskPartitionCreateRequest struct {
	MountPoint    string `json:"mountPoint,omitempty"`
	PartitionName string `json:"partitionName,omitempty"`
	PartitionSize string `json:"partitionSize,omitempty"`
	PartitionType string `json:"partitionType,omitempty"`
}

type TagRequest struct {
	TagKey   string `json:"tagKey"`
	TagValue string `json:"tagValue,omitempty"`
}

type ValidationResponseDto struct {
	ErrorCode    string `json:"errorCode,omitempty"`
	ErrorMessage string `json:"errorMessage,omitempty"`
	Status       string `json:"status,omitempty"`
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
