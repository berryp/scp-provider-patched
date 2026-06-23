package publicip2

import (
	"github.com/antihax/optional"
	"time"
)

type APIKey struct {
	Key    string
	Prefix string
}

type BasicAuth struct {
	UserName string `json:"userName,omitempty"`
	Password string `json:"password,omitempty"`
}

type ChangePublicIpRequest struct {
	PublicIpAddressDescription string `json:"publicIpAddressDescription"`
}

type CreatePublicIpV4Request struct {
	ServiceZoneId       string                    `json:"serviceZoneId"`
	Tags                []TagRequest `json:"tags,omitempty"`
	UplinkType          string                    `json:"uplinkType"`
	PublicIpDescription string                    `json:"publicIpDescription,omitempty"`
}

type DetailPublicIpResponse struct {
	ProjectId                  string       `json:"projectId,omitempty"`
	AttachedObjectName         string       `json:"attachedObjectName,omitempty"`
	IpAddress                  string       `json:"ipAddress,omitempty"`
	IpAddressId                string       `json:"ipAddressId,omitempty"`
	ProductGroupId             string       `json:"productGroupId,omitempty"`
	PublicIpAddressId          string       `json:"publicIpAddressId,omitempty"`
	PublicIpPurpose            string       `json:"publicIpPurpose,omitempty"`
	PublicIpState              string       `json:"publicIpState,omitempty"`
	ServiceZoneId              string       `json:"serviceZoneId,omitempty"`
	UplinkType                 string       `json:"uplinkType,omitempty"`
	PublicIpAddressDescription string       `json:"publicIpAddressDescription,omitempty"`
	CreatedBy                  string       `json:"createdBy,omitempty"`
	CreatedDt                  time.Time `json:"createdDt,omitempty"`
	ModifiedBy                 string       `json:"modifiedBy,omitempty"`
	ModifiedDt                 time.Time `json:"modifiedDt,omitempty"`
}

type ListResponseDetailPublicIpResponse struct {
	Contents   []DetailPublicIpResponse `json:"contents,omitempty"`
	TotalCount int32                                 `json:"totalCount,omitempty"`
}

type PublicIpOpenApiV3ControllerApiListPublicIpsV3Opts struct {
	IpAddress     optional.String
	PublicIpState optional.String
	ServiceZoneId optional.String
	UplinkType    optional.String
	VpcId         optional.String
	CreatedBy     optional.String
	Page          optional.Int32
	Size          optional.Int32
	Sort          optional.Interface
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
