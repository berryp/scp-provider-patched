package internetgateway2

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

type InternetGatewayCreateRequest struct {
	FirewallEnabled            *bool                            `json:"firewallEnabled"`
	FirewallLoggable           *bool                            `json:"firewallLoggable,omitempty"`
	ServiceZoneId              string                           `json:"serviceZoneId"`
	Tags                       []TagRequest `json:"tags,omitempty"`
	VpcId                      string                           `json:"vpcId"`
	InternetGatewayDescription string                           `json:"internetGatewayDescription,omitempty"`
}

type InternetGatewayCreateV3Request struct {
	FirewallEnabled            *bool                            `json:"firewallEnabled"`
	FirewallLoggable           *bool                            `json:"firewallLoggable,omitempty"`
	InternetGatewayType        string                           `json:"internetGatewayType"`
	ServiceZoneId              string                           `json:"serviceZoneId"`
	Tags                       []TagRequest `json:"tags,omitempty"`
	VpcId                      string                           `json:"vpcId"`
	InternetGatewayDescription string                           `json:"internetGatewayDescription,omitempty"`
}

type InternetGatewayCreateV4Request struct {
	FirewallEnabled            *bool                            `json:"firewallEnabled"`
	FirewallLoggable           *bool                            `json:"firewallLoggable,omitempty"`
	InternetGatewayType        string                           `json:"internetGatewayType"`
	Tags                       []TagRequest `json:"tags,omitempty"`
	VpcId                      string                           `json:"vpcId"`
	InternetGatewayDescription string                           `json:"internetGatewayDescription,omitempty"`
}

type InternetGatewayDetailResponse struct {
	ProjectId                  string       `json:"projectId,omitempty"`
	BlockId                    string       `json:"blockId,omitempty"`
	InternetGatewayId          string       `json:"internetGatewayId,omitempty"`
	InternetGatewayName        string       `json:"internetGatewayName,omitempty"`
	InternetGatewayState       string       `json:"internetGatewayState,omitempty"`
	InternetGatewayType        string       `json:"internetGatewayType,omitempty"`
	ServiceZoneId              string       `json:"serviceZoneId,omitempty"`
	VpcId                      string       `json:"vpcId,omitempty"`
	InternetGatewayDescription string       `json:"internetGatewayDescription,omitempty"`
	CreatedBy                  string       `json:"createdBy,omitempty"`
	CreatedDt                  time.Time `json:"createdDt,omitempty"`
	ModifiedBy                 string       `json:"modifiedBy,omitempty"`
	ModifiedDt                 time.Time `json:"modifiedDt,omitempty"`
}

type InternetGatewayListItemResponse struct {
	ProjectId                  string       `json:"projectId,omitempty"`
	BlockId                    string       `json:"blockId,omitempty"`
	InternetGatewayId          string       `json:"internetGatewayId,omitempty"`
	InternetGatewayName        string       `json:"internetGatewayName,omitempty"`
	InternetGatewayState       string       `json:"internetGatewayState,omitempty"`
	InternetGatewayType        string       `json:"internetGatewayType,omitempty"`
	ServiceZoneId              string       `json:"serviceZoneId,omitempty"`
	VpcId                      string       `json:"vpcId,omitempty"`
	InternetGatewayDescription string       `json:"internetGatewayDescription,omitempty"`
	CreatedBy                  string       `json:"createdBy,omitempty"`
	CreatedDt                  time.Time `json:"createdDt,omitempty"`
	ModifiedBy                 string       `json:"modifiedBy,omitempty"`
	ModifiedDt                 time.Time `json:"modifiedDt,omitempty"`
}

type InternetGatewayModifyDescriptionRequest struct {
	InternetGatewayDescription string `json:"internetGatewayDescription,omitempty"`
}

type InternetGatewayV2ControllerV2ApiListInternetGatewaysOpts struct {
	InternetGatewayId   optional.String
	InternetGatewayName optional.String
	VpcId               optional.String
	CreatedBy           optional.String
	Page                optional.Int32
	Size                optional.Int32
	Sort                optional.Interface
}

type ListResponseInternetGatewayListItemResponse struct {
	Contents   []InternetGatewayListItemResponse `json:"contents,omitempty"`
	TotalCount int32                                                 `json:"totalCount,omitempty"`
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
