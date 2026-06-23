package natgateway2

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

type ListResponseNatGatewayListItemResponse struct {
	Contents   []NatGatewayListItemResponse `json:"contents,omitempty"`
	TotalCount int32                                       `json:"totalCount,omitempty"`
}

type NatGatewayCreateRequest struct {
	PublicIpAddressId     string                      `json:"publicIpAddressId,omitempty"`
	SubnetId              string                      `json:"subnetId"`
	Tags                  []TagRequest `json:"tags,omitempty"`
	NatGatewayDescription string                      `json:"natGatewayDescription,omitempty"`
}

type NatGatewayDescriptionUpdateRequest struct {
	NatGatewayDescription string `json:"natGatewayDescription,omitempty"`
}

type NatGatewayDetailResponse struct {
	ProjectId             string       `json:"projectId,omitempty"`
	BlockId               string       `json:"blockId,omitempty"`
	NatGatewayId          string       `json:"natGatewayId,omitempty"`
	NatGatewayIpAddress   string       `json:"natGatewayIpAddress,omitempty"`
	NatGatewayName        string       `json:"natGatewayName,omitempty"`
	NatGatewayState       string       `json:"natGatewayState,omitempty"`
	ServiceZoneId         string       `json:"serviceZoneId,omitempty"`
	SubnetId              string       `json:"subnetId,omitempty"`
	VpcId                 string       `json:"vpcId,omitempty"`
	NatGatewayDescription string       `json:"natGatewayDescription,omitempty"`
	CreatedBy             string       `json:"createdBy,omitempty"`
	CreatedDt             time.Time `json:"createdDt,omitempty"`
	ModifiedBy            string       `json:"modifiedBy,omitempty"`
	ModifiedDt            time.Time `json:"modifiedDt,omitempty"`
}

type NatGatewayListItemResponse struct {
	ProjectId           string       `json:"projectId,omitempty"`
	BlockId             string       `json:"blockId,omitempty"`
	NatGatewayId        string       `json:"natGatewayId,omitempty"`
	NatGatewayIpAddress string       `json:"natGatewayIpAddress,omitempty"`
	NatGatewayName      string       `json:"natGatewayName,omitempty"`
	NatGatewayState     string       `json:"natGatewayState,omitempty"`
	ServiceZoneId       string       `json:"serviceZoneId,omitempty"`
	SubnetId            string       `json:"subnetId,omitempty"`
	VpcId               string       `json:"vpcId,omitempty"`
	CreatedBy           string       `json:"createdBy,omitempty"`
	CreatedDt           time.Time `json:"createdDt,omitempty"`
	ModifiedBy          string       `json:"modifiedBy,omitempty"`
	ModifiedDt          time.Time `json:"modifiedDt,omitempty"`
}

type NatGatewayV2ControllerV2ApiListNatGatewaysOpts struct {
	NatGatewayId    optional.String
	NatGatewayName  optional.String
	NatGatewayState optional.String
	SubnetId        optional.String
	VpcId           optional.String
	CreatedBy       optional.String
	Page            optional.Int32
	Size            optional.Int32
	Sort            optional.Interface
}

type TagRequest struct {
	TagKey   string `json:"tagKey"`
	TagValue string `json:"tagValue,omitempty"`
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
