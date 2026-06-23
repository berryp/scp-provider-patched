package gslb2

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

type ChangeGslbAlgorithmRequest struct {
	GslbAlgorithm string `json:"gslbAlgorithm"`
}

type ChangeGslbHealthCheckRequest struct {
	GslbHealthCheckInterval     int32  `json:"gslbHealthCheckInterval,omitempty"`
	GslbHealthCheckTimeout      int32  `json:"gslbHealthCheckTimeout,omitempty"`
	GslbHealthCheckUserId       string `json:"gslbHealthCheckUserId,omitempty"`
	GslbHealthCheckUserPassword string `json:"gslbHealthCheckUserPassword,omitempty"`
	GslbResponseString          string `json:"gslbResponseString,omitempty"`
	GslbSendString              string `json:"gslbSendString,omitempty"`
	ProbeTimeout                int32  `json:"probeTimeout,omitempty"`
	Protocol                    string `json:"protocol"`
	ServicePort                 int32  `json:"servicePort,omitempty"`
}

type ChangeGslbResourceDisableRequest struct {
	GslbResourceDisable *bool `json:"gslbResourceDisable"`
}

type ChangeGslbResourceMappingRequest struct {
	GslbResources []GslbResourceMappingRequestVo `json:"gslbResources"`
}

type CreateGslbServiceV3Request struct {
	GslbAlgorithm   string                                `json:"gslbAlgorithm"`
	GslbEnvUsage    string                                `json:"gslbEnvUsage"`
	GslbHealthCheck *GslbHealthCheckReqVo1       `json:"gslbHealthCheck"`
	GslbName        string                                `json:"gslbName"`
	GslbResources   []GslbResourceMappingRequest `json:"gslbResources"`
	Tags            []TagRequest                 `json:"tags,omitempty"`
}

type GslbHealthCheckReqVo1 struct {
	GslbHealthCheckInterval     int32  `json:"gslbHealthCheckInterval,omitempty"`
	GslbHealthCheckTimeout      int32  `json:"gslbHealthCheckTimeout,omitempty"`
	GslbHealthCheckUserId       string `json:"gslbHealthCheckUserId,omitempty"`
	GslbHealthCheckUserPassword string `json:"gslbHealthCheckUserPassword,omitempty"`
	GslbResponseString          string `json:"gslbResponseString,omitempty"`
	GslbSendString              string `json:"gslbSendString,omitempty"`
	ProbeTimeout                int32  `json:"probeTimeout,omitempty"`
	Protocol                    string `json:"protocol"`
	ServicePort                 int32  `json:"servicePort,omitempty"`
}

type GslbHealthCheckResponse struct {
	GslbHealthCheckInterval int32  `json:"gslbHealthCheckInterval,omitempty"`
	GslbHealthCheckTimeout  int32  `json:"gslbHealthCheckTimeout,omitempty"`
	GslbHealthCheckUserId   string `json:"gslbHealthCheckUserId,omitempty"`
	GslbResponseString      string `json:"gslbResponseString,omitempty"`
	GslbSendString          string `json:"gslbSendString,omitempty"`
	Id                      string `json:"id,omitempty"`
	ProbeTimeout            int32  `json:"probeTimeout,omitempty"`
	Protocol                string `json:"protocol,omitempty"`
	ServicePort             int32  `json:"servicePort,omitempty"`
}

type GslbOpenApiV2ControllerApiListGslbResourcesOpts struct {
	Page optional.Int32
	Size optional.Int32
	Sort optional.Interface
}

type GslbOpenApiV2ControllerApiListGslbsOpts struct {
	GslbEnvUsage optional.String
	GslbName     optional.String
	CreatedBy    optional.String
	Page         optional.Int32
	Size         optional.Int32
	Sort         optional.Interface
}

type GslbResourceMappingRequest struct {
	GslbDestination         string `json:"gslbDestination"`
	GslbRegion              string `json:"gslbRegion"`
	GslbResourceWeight      int32  `json:"gslbResourceWeight,omitempty"`
	GslbResourceDescription string `json:"gslbResourceDescription,omitempty"`
}

type GslbResourceMappingRequestVo struct {
	GslbDestination         string `json:"gslbDestination"`
	GslbRegion              string `json:"gslbRegion"`
	GslbResourceDisable     *bool  `json:"gslbResourceDisable,omitempty"`
	GslbResourceWeight      int32  `json:"gslbResourceWeight,omitempty"`
	GslbResourceDescription string `json:"gslbResourceDescription,omitempty"`
}

type GslbResourceMappingResponse struct {
	GslbDestination         string `json:"gslbDestination,omitempty"`
	GslbRegion              string `json:"gslbRegion,omitempty"`
	GslbResourceDisable     *bool  `json:"gslbResourceDisable,omitempty"`
	GslbResourceId          string `json:"gslbResourceId,omitempty"`
	GslbResourceName        string `json:"gslbResourceName,omitempty"`
	GslbResourceWeight      int32  `json:"gslbResourceWeight,omitempty"`
	GslbResourceDescription string `json:"gslbResourceDescription,omitempty"`
}

type GslbServiceDetailResponse struct {
	ProjectId           string                            `json:"projectId,omitempty"`
	GslbAlgorithm       string                            `json:"gslbAlgorithm,omitempty"`
	GslbEnvUsage        string                            `json:"gslbEnvUsage,omitempty"`
	GslbHealthCheck     *GslbHealthCheckResponse `json:"gslbHealthCheck"`
	GslbId              string                            `json:"gslbId,omitempty"`
	GslbName            string                            `json:"gslbName,omitempty"`
	GslbState           string                            `json:"gslbState,omitempty"`
	LinkedResourceCount int32                             `json:"linkedResourceCount,omitempty"`
	CreatedBy           string                            `json:"createdBy,omitempty"`
	CreatedDt           time.Time                      `json:"createdDt,omitempty"`
	ModifiedBy          string                            `json:"modifiedBy,omitempty"`
	ModifiedDt          time.Time                      `json:"modifiedDt,omitempty"`
}

type GslbServiceListItemResponse struct {
	GslbAlgorithm       string       `json:"gslbAlgorithm,omitempty"`
	GslbEnvUsage        string       `json:"gslbEnvUsage,omitempty"`
	GslbId              string       `json:"gslbId,omitempty"`
	GslbName            string       `json:"gslbName,omitempty"`
	GslbState           string       `json:"gslbState,omitempty"`
	LinkedResourceCount int32        `json:"linkedResourceCount,omitempty"`
	CreatedDt           time.Time `json:"createdDt,omitempty"`
}

type ListResponseGslbResourceMappingResponse struct {
	Contents   []GslbResourceMappingResponse `json:"contents,omitempty"`
	TotalCount int32                                  `json:"totalCount,omitempty"`
}

type ListResponseGslbServiceListItemResponse struct {
	Contents   []GslbServiceListItemResponse `json:"contents,omitempty"`
	TotalCount int32                                  `json:"totalCount,omitempty"`
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
