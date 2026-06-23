package baremetalservervdc

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

type BaremetalServersTerminateRequest struct {
	BareMetalServerIds []string `json:"bareMetalServerIds,omitempty"`
}

type BasicAuth struct {
	UserName string `json:"userName,omitempty"`
	Password string `json:"password,omitempty"`
}

type ListResponseVxLanBmServerGridResponse struct {
	Contents   []VxLanBmServerGridResponse `json:"contents,omitempty"`
	TotalCount int32                                             `json:"totalCount,omitempty"`
}

type TagRequest struct {
	TagKey   string `json:"tagKey"`
	TagValue string `json:"tagValue,omitempty"`
}

type VxLanAdditionalBlockStorageCreateOpenApiRequest struct {
	BareMetalBlockStorageName string `json:"bareMetalBlockStorageName"`
	BareMetalBlockStorageSize int32  `json:"bareMetalBlockStorageSize"`
	BareMetalBlockStorageType string `json:"bareMetalBlockStorageType,omitempty"`
	EncryptionEnabled         *bool  `json:"encryptionEnabled"`
}

type VxLanAdditionalBlockStorageCreateRequest struct {
	BareMetalBlockStorageName   string `json:"bareMetalBlockStorageName"`
	BareMetalBlockStorageSize   int32  `json:"bareMetalBlockStorageSize"`
	BareMetalBlockStorageType   string `json:"bareMetalBlockStorageType,omitempty"`
	BareMetalBlockStorageTypeId string `json:"bareMetalBlockStorageTypeId"`
	EncryptionEnabled           *bool  `json:"encryptionEnabled"`
}

type VxLanBareMetalServerDetailResponse struct {
	ProjectId                 string       `json:"projectId,omitempty"`
	AllMountedStorage         *bool        `json:"allMountedStorage,omitempty"`
	BareMetalBlockStorageIds  []string     `json:"bareMetalBlockStorageIds,omitempty"`
	BareMetalServerId         string       `json:"bareMetalServerId,omitempty"`
	BareMetalServerName       string       `json:"bareMetalServerName,omitempty"`
	BareMetalServerState      string       `json:"bareMetalServerState,omitempty"`
	BlockId                   string       `json:"blockId,omitempty"`
	CheckCriticalError        *bool        `json:"checkCriticalError,omitempty"`
	Contract                  string       `json:"contract,omitempty"`
	ContractEndDate           string       `json:"contractEndDate,omitempty"`
	ContractId                string       `json:"contractId,omitempty"`
	ContractStartDate         string       `json:"contractStartDate,omitempty"`
	DeletionProtectionEnabled string       `json:"deletionProtectionEnabled,omitempty"`
	DnsEnabled                string       `json:"dnsEnabled,omitempty"`
	ErrorCheck                *bool        `json:"errorCheck,omitempty"`
	ImageId                   string       `json:"imageId,omitempty"`
	InitialScriptContent      string       `json:"initialScriptContent,omitempty"`
	IpAddress                 string       `json:"ipAddress,omitempty"`
	NatIpAddress              string       `json:"natIpAddress,omitempty"`
	NextContract              string       `json:"nextContract,omitempty"`
	NextContractEndDate       string       `json:"nextContractEndDate,omitempty"`
	OsType                    string       `json:"osType,omitempty"`
	OsVersion                 string       `json:"osVersion,omitempty"`
	ProductGroupId            string       `json:"productGroupId,omitempty"`
	ProductType               string       `json:"productType,omitempty"`
	PublicNatStatus           string       `json:"publicNatStatus,omitempty"`
	ServerTypeId              string       `json:"serverTypeId,omitempty"`
	ServiceLevelId            string       `json:"serviceLevelId,omitempty"`
	ServiceZoneId             string       `json:"serviceZoneId,omitempty"`
	SubnetId                  string       `json:"subnetId,omitempty"`
	Timezone                  string       `json:"timezone,omitempty"`
	UseHyperThreading         string       `json:"useHyperThreading,omitempty"`
	VdcId                     string       `json:"vdcId,omitempty"`
	CreatedBy                 string       `json:"createdBy,omitempty"`
	CreatedDt                 time.Time `json:"createdDt,omitempty"`
	ModifiedBy                string       `json:"modifiedBy,omitempty"`
	ModifiedDt                time.Time `json:"modifiedDt,omitempty"`
}

type VxLanBareMetalServerOperationRequest struct {
	BareMetalServerIds []string `json:"bareMetalServerIds,omitempty"`
}

type VxLanBareMetalServerSimpleTaskOpenApiControllerApiListVdcBareMetalServersOpts struct {
	BareMetalServerName   optional.String
	BareMetalServerStates optional.Interface
	BlockId               optional.String
	IpAddress             optional.String
	RegionName            optional.String
	ServiceZoneId         optional.String
	CreatedBy             optional.String
	Page                  optional.Int32
	Size                  optional.Int32
	Sort                  optional.Interface
}

type VxLanBmServerCreateOpenApiRequest struct {
	BlockId                   string                                                     `json:"blockId"`
	Contract                  string                                                     `json:"contract,omitempty"`
	DeletionProtectionEnabled *bool                                                      `json:"deletionProtectionEnabled"`
	ImageId                   string                                                     `json:"imageId"`
	InitScript                string                                                     `json:"initScript,omitempty"`
	OsUserId                  string                                                     `json:"osUserId"`
	OsUserPassword            string                                                     `json:"osUserPassword"`
	ServerDetails             []VxLanBmServerDetailsOpenApiRequest `json:"serverDetails"`
	ServiceZoneId             string                                                     `json:"serviceZoneId"`
	SubnetId                  string                                                     `json:"subnetId"`
	Tags                      []TagRequest                         `json:"tags,omitempty"`
	VdcId                     string                                                     `json:"vdcId"`
}

type VxLanBmServerCreateOpenApiv3Request struct {
	BillingStartDate          string                                                       `json:"billingStartDate,omitempty"`
	BlockId                   string                                                       `json:"blockId"`
	Contract                  string                                                       `json:"contract,omitempty"`
	DeletionProtectionEnabled *bool                                                        `json:"deletionProtectionEnabled"`
	ImageId                   string                                                       `json:"imageId"`
	InitScript                string                                                       `json:"initScript,omitempty"`
	OsUserId                  string                                                       `json:"osUserId"`
	OsUserPassword            string                                                       `json:"osUserPassword"`
	ServerDetails             []VxLanBmServerDetailsOpenApiv3Request `json:"serverDetails"`
	ServiceZoneId             string                                                       `json:"serviceZoneId"`
	Tags                      []TagRequest                           `json:"tags,omitempty"`
	VdcId                     string                                                       `json:"vdcId"`
}

type VxLanBmServerCreateRequest struct {
	BlockId                   string                                              `json:"blockId"`
	ContractId                string                                              `json:"contractId"`
	DeletionProtectionEnabled *bool                                               `json:"deletionProtectionEnabled"`
	ImageId                   string                                              `json:"imageId"`
	InitScript                string                                              `json:"initScript,omitempty"`
	OsUserId                  string                                              `json:"osUserId"`
	OsUserPassword            string                                              `json:"osUserPassword"`
	ProductGroupId            string                                              `json:"productGroupId"`
	ServerDetails             []VxLanBmServerDetailsRequest `json:"serverDetails"`
	ServiceZoneId             string                                              `json:"serviceZoneId"`
	SubnetId                  string                                              `json:"subnetId"`
	Tags                      []TagRequest                  `json:"tags,omitempty"`
	VdcId                     string                                              `json:"vdcId"`
}

type VxLanBmServerDetailsOpenApiRequest struct {
	BareMetalServerName string                                                                  `json:"bareMetalServerName"`
	DnsEnabled          *bool                                                                   `json:"dnsEnabled,omitempty"`
	IpAddress           string                                                                  `json:"ipAddress,omitempty"`
	ServerType          string                                                                  `json:"serverType"`
	StorageDetails      []VxLanAdditionalBlockStorageCreateOpenApiRequest `json:"storageDetails"`
	UseHyperThreading   string                                                                  `json:"useHyperThreading"`
}

type VxLanBmServerDetailsOpenApiv3Request struct {
	BareMetalServerName string                                                                  `json:"bareMetalServerName"`
	DnsEnabled          *bool                                                                   `json:"dnsEnabled,omitempty"`
	IpAddress           string                                                                  `json:"ipAddress,omitempty"`
	ServerType          string                                                                  `json:"serverType"`
	StorageDetails      []VxLanAdditionalBlockStorageCreateOpenApiRequest `json:"storageDetails"`
	SubnetId            string                                                                  `json:"subnetId"`
	UseHyperThreading   string                                                                  `json:"useHyperThreading"`
}

type VxLanBmServerDetailsRequest struct {
	BareMetalServerName string                                                           `json:"bareMetalServerName"`
	DnsEnabled          *bool                                                            `json:"dnsEnabled"`
	IpAddress           string                                                           `json:"ipAddress,omitempty"`
	ServerTypeId        string                                                           `json:"serverTypeId"`
	StorageDetails      []VxLanAdditionalBlockStorageCreateRequest `json:"storageDetails"`
	UseHyperThreading   string                                                           `json:"useHyperThreading"`
}

type VxLanBmServerGridResponse struct {
	ProjectId            string       `json:"projectId,omitempty"`
	AllMountedStorage    *bool        `json:"allMountedStorage,omitempty"`
	BareMetalServerId    string       `json:"bareMetalServerId,omitempty"`
	BareMetalServerName  string       `json:"bareMetalServerName,omitempty"`
	BareMetalServerState string       `json:"bareMetalServerState,omitempty"`
	BlockId              string       `json:"blockId,omitempty"`
	CheckCriticalError   *bool        `json:"checkCriticalError,omitempty"`
	ErrorCheck           *bool        `json:"errorCheck,omitempty"`
	ImageId              string       `json:"imageId,omitempty"`
	ImageName            string       `json:"imageName,omitempty"`
	IpAddress            string       `json:"ipAddress,omitempty"`
	OsType               string       `json:"osType,omitempty"`
	ProductType          string       `json:"productType,omitempty"`
	RegionName           string       `json:"regionName,omitempty"`
	ServerType           string       `json:"serverType,omitempty"`
	ServiceZoneId        string       `json:"serviceZoneId,omitempty"`
	CreatedBy            string       `json:"createdBy,omitempty"`
	CreatedDt            time.Time `json:"createdDt,omitempty"`
	ModifiedBy           string       `json:"modifiedBy,omitempty"`
	ModifiedDt           time.Time `json:"modifiedDt,omitempty"`
}

type VxLanDeletionProtectionEnabledUpdateRequest struct {
	DeletionProtectionEnabled string `json:"deletionProtectionEnabled"`
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
