package vpc2

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

type CreateDnsUserZoneRequest struct {
	DnsUserZoneDomain   string `json:"dnsUserZoneDomain"`
	DnsUserZoneServerIp string `json:"dnsUserZoneServerIp"`
	DnsUserZoneSourceIp string `json:"dnsUserZoneSourceIp,omitempty"`
	SubnetId            string `json:"subnetId"`
}

type DetailVpcResponse struct {
	ProjectId      string       `json:"projectId,omitempty"`
	BlockId        string       `json:"blockId,omitempty"`
	DnsServiceIp   string       `json:"dnsServiceIp,omitempty"`
	ProductGroupId string       `json:"productGroupId,omitempty"`
	ServiceZoneId  string       `json:"serviceZoneId,omitempty"`
	T1RouterId     string       `json:"t1RouterId,omitempty"`
	VpcId          string       `json:"vpcId,omitempty"`
	VpcName        string       `json:"vpcName,omitempty"`
	VpcState       string       `json:"vpcState,omitempty"`
	VpcVersion     string       `json:"vpcVersion,omitempty"`
	VpcDescription string       `json:"vpcDescription,omitempty"`
	CreatedBy      string       `json:"createdBy,omitempty"`
	CreatedDt      time.Time `json:"createdDt,omitempty"`
	ModifiedBy     string       `json:"modifiedBy,omitempty"`
	ModifiedDt     time.Time `json:"modifiedDt,omitempty"`
}

type DnsUserZoneResponse struct {
	DnsUserZoneDomain   string `json:"dnsUserZoneDomain,omitempty"`
	DnsUserZoneId       string `json:"dnsUserZoneId,omitempty"`
	DnsUserZoneName     string `json:"dnsUserZoneName,omitempty"`
	DnsUserZoneServerIp string `json:"dnsUserZoneServerIp,omitempty"`
	DnsUserZoneSourceIp string `json:"dnsUserZoneSourceIp,omitempty"`
	DnsUserZoneState    string `json:"dnsUserZoneState,omitempty"`
}

type ListResponseDnsUserZoneResponse struct {
	Contents   []DnsUserZoneResponse `json:"contents,omitempty"`
	TotalCount int32                         `json:"totalCount,omitempty"`
}

type ListResponseVpcResponse struct {
	Contents   []VpcResponse `json:"contents,omitempty"`
	TotalCount int32                 `json:"totalCount,omitempty"`
}

type ModifyVpcDescriptionRequest struct {
	VpcDescription string `json:"vpcDescription"`
}

type TagRequest struct {
	TagKey   string `json:"tagKey"`
	TagValue string `json:"tagValue,omitempty"`
}

type VpcCreateV3Request struct {
	ServiceZoneId  string               `json:"serviceZoneId"`
	Tags           []TagRequest `json:"tags,omitempty"`
	VpcName        string               `json:"vpcName"`
	VpcDescription string               `json:"vpcDescription,omitempty"`
}

type VpcOpenApiControllerApiListVpcV2Opts struct {
	ServiceZoneId optional.String
	VpcId         optional.String
	VpcName       optional.String
	VpcStates     optional.Interface
	CreatedBy     optional.String
	Page          optional.Int32
	Size          optional.Int32
	Sort          optional.Interface
}

type VpcResponse struct {
	ProjectId     string       `json:"projectId,omitempty"`
	BlockId       string       `json:"blockId,omitempty"`
	ServiceZoneId string       `json:"serviceZoneId,omitempty"`
	VpcId         string       `json:"vpcId,omitempty"`
	VpcName       string       `json:"vpcName,omitempty"`
	VpcState      string       `json:"vpcState,omitempty"`
	CreatedBy     string       `json:"createdBy,omitempty"`
	CreatedDt     time.Time `json:"createdDt,omitempty"`
	ModifiedBy    string       `json:"modifiedBy,omitempty"`
	ModifiedDt    time.Time `json:"modifiedDt,omitempty"`
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
