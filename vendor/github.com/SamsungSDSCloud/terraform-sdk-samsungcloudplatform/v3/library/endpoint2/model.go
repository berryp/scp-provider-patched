package endpoint2

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

type EndpointCreateV3Request struct {
	EndpointIpAddress   string                    `json:"endpointIpAddress"`
	EndpointName        string                    `json:"endpointName"`
	EndpointType        string                    `json:"endpointType"`
	ObjectId            string                    `json:"objectId,omitempty"`
	ServiceZoneId       string                    `json:"serviceZoneId"`
	Tags                []TagRequest `json:"tags,omitempty"`
	VpcId               string                    `json:"vpcId"`
	EndpointDescription string                    `json:"endpointDescription,omitempty"`
}

type EndpointDetailResponse struct {
	ProjectId           string       `json:"projectId,omitempty"`
	BlockId             string       `json:"blockId,omitempty"`
	EndpointId          string       `json:"endpointId,omitempty"`
	EndpointIpAddress   string       `json:"endpointIpAddress,omitempty"`
	EndpointName        string       `json:"endpointName,omitempty"`
	EndpointState       string       `json:"endpointState,omitempty"`
	EndpointType        string       `json:"endpointType,omitempty"`
	ObjectId            string       `json:"objectId,omitempty"`
	ProductGroupId      string       `json:"productGroupId,omitempty"`
	ServiceZoneId       string       `json:"serviceZoneId,omitempty"`
	VpcId               string       `json:"vpcId,omitempty"`
	EndpointDescription string       `json:"endpointDescription,omitempty"`
	CreatedBy           string       `json:"createdBy,omitempty"`
	CreatedDt           time.Time `json:"createdDt,omitempty"`
	ModifiedBy          string       `json:"modifiedBy,omitempty"`
	ModifiedDt          time.Time `json:"modifiedDt,omitempty"`
}

type EndpointModifyDescriptionRequest struct {
	EndpointDescription string `json:"endpointDescription"`
}

type EndpointOpenApiControllerApiListEndpointOpts struct {
	EndpointId        optional.String
	EndpointIpAddress optional.String
	EndpointName      optional.String
	EndpointStates    optional.Interface
	EndpointType      optional.String
	ObjectId          optional.String
	ServiceZoneId     optional.String
	VpcId             optional.String
	CreatedBy         optional.String
	Page              optional.Int32
	Size              optional.Int32
	Sort              optional.Interface
}

type EndpointResponse struct {
	ProjectId         string       `json:"projectId,omitempty"`
	BlockId           string       `json:"blockId,omitempty"`
	EndpointId        string       `json:"endpointId,omitempty"`
	EndpointIpAddress string       `json:"endpointIpAddress,omitempty"`
	EndpointName      string       `json:"endpointName,omitempty"`
	EndpointState     string       `json:"endpointState,omitempty"`
	EndpointType      string       `json:"endpointType,omitempty"`
	ObjectId          string       `json:"objectId,omitempty"`
	ServiceZoneId     string       `json:"serviceZoneId,omitempty"`
	VpcId             string       `json:"vpcId,omitempty"`
	CreatedBy         string       `json:"createdBy,omitempty"`
	CreatedDt         time.Time `json:"createdDt,omitempty"`
	ModifiedBy        string       `json:"modifiedBy,omitempty"`
	ModifiedDt        time.Time `json:"modifiedDt,omitempty"`
}

type ListResponseEndpointResponse struct {
	Contents   []EndpointResponse `json:"contents,omitempty"`
	TotalCount int32                           `json:"totalCount,omitempty"`
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
