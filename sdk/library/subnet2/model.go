package subnet2

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

type AttachSubnetPublicIpRequest struct {
	PublicIpAddressId string `json:"publicIpAddressId"`
}

type AttachSubnetSecurityGroupRequest struct {
	SecurityGroupId string `json:"securityGroupId"`
}

type BasicAuth struct {
	UserName string `json:"userName,omitempty"`
	Password string `json:"password,omitempty"`
}

type CheckResponse struct {
	Result *bool `json:"result,omitempty"`
}

type CreateSubnetRequest struct {
	SubnetCidrBlock   string                  `json:"subnetCidrBlock"`
	SubnetName        string                  `json:"subnetName"`
	SubnetType        string                  `json:"subnetType"`
	Tags              []TagRequest `json:"tags,omitempty"`
	VpcId             string                  `json:"vpcId"`
	SubnetDescription string                  `json:"subnetDescription"`
}

type ListResponseSubnetListItemResVo struct {
	Contents   []SubnetListItemResVo `json:"contents,omitempty"`
	TotalCount int32                            `json:"totalCount,omitempty"`
}

type ListResponseSubnetResourceIpListItemResVo struct {
	Contents   []SubnetResourceIpListItemResVo `json:"contents,omitempty"`
	TotalCount int32                                      `json:"totalCount,omitempty"`
}

type ListResponseSubnetVirtualIpAvailableListItemResVo struct {
	Contents   []SubnetVirtualIpAvailableListItemResVo `json:"contents,omitempty"`
	TotalCount int32                                              `json:"totalCount,omitempty"`
}

type ListResponseSubnetVirtualIpListItemResVo struct {
	Contents   []SubnetVirtualIpListItemResVo `json:"contents,omitempty"`
	TotalCount int32                                     `json:"totalCount,omitempty"`
}

type SecurityGroupIdsResponse struct {
	SecurityGroupId          string `json:"securityGroupId,omitempty"`
	SecurityGroupMemberState string `json:"securityGroupMemberState,omitempty"`
}

type SubnetDetailResVo struct {
	ProjectId         string       `json:"projectId,omitempty"`
	GatewayIpAddress  string       `json:"gatewayIpAddress,omitempty"`
	ServiceZoneId     string       `json:"serviceZoneId,omitempty"`
	SubnetCidrBlock   string       `json:"subnetCidrBlock,omitempty"`
	SubnetId          string       `json:"subnetId,omitempty"`
	SubnetName        string       `json:"subnetName,omitempty"`
	SubnetPurpose     string       `json:"subnetPurpose,omitempty"`
	SubnetState       string       `json:"subnetState,omitempty"`
	SubnetType        string       `json:"subnetType,omitempty"`
	VpcId             string       `json:"vpcId,omitempty"`
	SubnetDescription string       `json:"subnetDescription,omitempty"`
	CreatedBy         string       `json:"createdBy,omitempty"`
	CreatedDt         time.Time `json:"createdDt,omitempty"`
	ModifiedBy        string       `json:"modifiedBy,omitempty"`
	ModifiedDt        time.Time `json:"modifiedDt,omitempty"`
}

type SubnetListItemResVo struct {
	GatewayIpAddress string       `json:"gatewayIpAddress,omitempty"`
	SubnetCidrBlock  string       `json:"subnetCidrBlock,omitempty"`
	SubnetId         string       `json:"subnetId,omitempty"`
	SubnetName       string       `json:"subnetName,omitempty"`
	SubnetPurpose    string       `json:"subnetPurpose,omitempty"`
	SubnetState      string       `json:"subnetState,omitempty"`
	SubnetType       string       `json:"subnetType,omitempty"`
	VpcId            string       `json:"vpcId,omitempty"`
	CreatedDt        time.Time `json:"createdDt,omitempty"`
}

type SubnetOpenApiControllerApiListSubnetV2Opts struct {
	SubnetCidrBlock optional.String
	SubnetId        optional.String
	SubnetName      optional.String
	SubnetStates    optional.Interface
	SubnetTypes     optional.Interface
	VpcId           optional.String
	CreatedBy       optional.String
	Page            optional.Int32
	Size            optional.Int32
	Sort            optional.Interface
}

type SubnetResourceIpListItemResVo struct {
	IpAddress        string       `json:"ipAddress,omitempty"`
	IpId             string       `json:"ipId,omitempty"`
	IpState          string       `json:"ipState,omitempty"`
	LinkedObjectId   string       `json:"linkedObjectId,omitempty"`
	LinkedObjectName string       `json:"linkedObjectName,omitempty"`
	LinkedObjectType string       `json:"linkedObjectType,omitempty"`
	IpDescription    string       `json:"ipDescription,omitempty"`
	CreatedBy        string       `json:"createdBy,omitempty"`
	CreatedDt        time.Time `json:"createdDt,omitempty"`
	ModifiedBy       string       `json:"modifiedBy,omitempty"`
	ModifiedDt       time.Time `json:"modifiedDt,omitempty"`
}

type SubnetVipOpenApiControllerApiListAvailableVipsV2Opts struct {
	SubnetIpAddress optional.String
	Page            optional.Int32
	Size            optional.Int32
	Sort            optional.Interface
}

type SubnetVipOpenApiControllerApiListSubnetResourcesV2Opts struct {
	IpAddress        optional.String
	LinkedObjectType optional.String
	Page             optional.Int32
	Size             optional.Int32
	Sort             optional.Interface
}

type SubnetVipOpenApiControllerApiListSubnetVipsV2Opts struct {
	SubnetIpAddress optional.String
	VipState        optional.String
	Page            optional.Int32
	Size            optional.Int32
	Sort            optional.Interface
}

type SubnetVirtualIpAvailableListItemResVo struct {
	IpId            string       `json:"ipId,omitempty"`
	SubnetIpAddress string       `json:"subnetIpAddress,omitempty"`
	VipState        string       `json:"vipState,omitempty"`
	VipDescription  string       `json:"vipDescription,omitempty"`
	CreatedBy       string       `json:"createdBy,omitempty"`
	CreatedDt       time.Time `json:"createdDt,omitempty"`
	ModifiedBy      string       `json:"modifiedBy,omitempty"`
	ModifiedDt      time.Time `json:"modifiedDt,omitempty"`
}

type SubnetVirtualIpDetailResVo struct {
	ProjectId        string                                `json:"projectId,omitempty"`
	BlockId          string                                `json:"blockId,omitempty"`
	NatIpAddress     string                                `json:"natIpAddress,omitempty"`
	NatIpId          string                                `json:"natIpId,omitempty"`
	SecurityGroupIds []SecurityGroupIdsResponse `json:"securityGroupIds,omitempty"`
	ServiceZoneId    string                                `json:"serviceZoneId,omitempty"`
	SubnetId         string                                `json:"subnetId,omitempty"`
	SubnetIpAddress  string                                `json:"subnetIpAddress,omitempty"`
	SubnetIpId       string                                `json:"subnetIpId,omitempty"`
	SubnetState      string                                `json:"subnetState,omitempty"`
	VipId            string                                `json:"vipId,omitempty"`
	VipState         string                                `json:"vipState,omitempty"`
	VipDescription   string                                `json:"vipDescription,omitempty"`
	CreatedBy        string                                `json:"createdBy,omitempty"`
	CreatedDt        time.Time                          `json:"createdDt,omitempty"`
	ModifiedBy       string                                `json:"modifiedBy,omitempty"`
	ModifiedDt       time.Time                          `json:"modifiedDt,omitempty"`
}

type SubnetVirtualIpListItemResVo struct {
	ProjectId        string                                `json:"projectId,omitempty"`
	BlockId          string                                `json:"blockId,omitempty"`
	NatIpAddress     string                                `json:"natIpAddress,omitempty"`
	NatIpId          string                                `json:"natIpId,omitempty"`
	SecurityGroupIds []SecurityGroupIdsResponse `json:"securityGroupIds,omitempty"`
	ServiceZoneId    string                                `json:"serviceZoneId,omitempty"`
	SubnetIpAddress  string                                `json:"subnetIpAddress,omitempty"`
	SubnetIpId       string                                `json:"subnetIpId,omitempty"`
	VipId            string                                `json:"vipId,omitempty"`
	VipState         string                                `json:"vipState,omitempty"`
	VipDescription   string                                `json:"vipDescription,omitempty"`
	CreatedBy        string                                `json:"createdBy,omitempty"`
	CreatedDt        time.Time                          `json:"createdDt,omitempty"`
	ModifiedBy       string                                `json:"modifiedBy,omitempty"`
	ModifiedDt       time.Time                          `json:"modifiedDt,omitempty"`
}

type SubnetVirtualIpReserveRequest struct {
	VipDescription string `json:"vipDescription"`
}

type TagRequest struct {
	TagKey   string `json:"tagKey"`
	TagValue string `json:"tagValue,omitempty"`
}

type UpdateSubnetDescriptionRequest struct {
	SubnetDescription string `json:"subnetDescription"`
}

type UpdateSubnetTypeRequest struct {
	SubnetType string `json:"subnetType"`
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
