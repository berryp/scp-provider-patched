package transitgateway2

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

type ListResponseTransitGatewayConnectionListItemResponse struct {
	Contents   []TransitGatewayConnectionListItemResponse `json:"contents,omitempty"`
	TotalCount int32                                                         `json:"totalCount,omitempty"`
}

type ListResponseTransitGatewayListItemResponse struct {
	Contents   []TransitGatewayListItemResponse `json:"contents,omitempty"`
	TotalCount int32                                               `json:"totalCount,omitempty"`
}

type ListResponseTransitGatewayPeeringListItemResponse struct {
	Contents   []TransitGatewayPeeringListItemResponse `json:"contents,omitempty"`
	TotalCount int32                                                      `json:"totalCount,omitempty"`
}

type TagRequest struct {
	TagKey   string `json:"tagKey"`
	TagValue string `json:"tagValue,omitempty"`
}

type TransitGatewayConnectionApprovalResponse struct {
	Success                    *bool  `json:"success,omitempty"`
	TransitGatewayConnectionId string `json:"transitGatewayConnectionId,omitempty"`
}

type TransitGatewayConnectionCreateRequest struct {
	ApproverProjectId                   string                          `json:"approverProjectId"`
	ApproverVpcId                       string                          `json:"approverVpcId"`
	FirewallEnabled                     *bool                           `json:"firewallEnabled"`
	FirewallLoggable                    *bool                           `json:"firewallLoggable,omitempty"`
	RequesterProjectId                  string                          `json:"requesterProjectId"`
	RequesterTransitGatewayId           string                          `json:"requesterTransitGatewayId"`
	Tags                                []TagRequest `json:"tags,omitempty"`
	TransitGatewayConnectionType        string                          `json:"transitGatewayConnectionType"`
	TransitGatewayConnectionDescription string                          `json:"transitGatewayConnectionDescription,omitempty"`
}

type TransitGatewayConnectionDescriptionUpdateRequest struct {
	TransitGatewayConnectionDescription string `json:"transitGatewayConnectionDescription,omitempty"`
}

type TransitGatewayConnectionDetailResponse struct {
	ProjectId                           string       `json:"projectId,omitempty"`
	ApprovedBy                          string       `json:"approvedBy,omitempty"`
	ApprovedDt                          time.Time `json:"approvedDt,omitempty"`
	ApproverProjectId                   string       `json:"approverProjectId,omitempty"`
	ApproverVpcId                       string       `json:"approverVpcId,omitempty"`
	BlockId                             string       `json:"blockId,omitempty"`
	CompletedDt                         time.Time `json:"completedDt,omitempty"`
	RequestedBy                         string       `json:"requestedBy,omitempty"`
	RequestedDt                         time.Time `json:"requestedDt,omitempty"`
	RequesterProjectId                  string       `json:"requesterProjectId,omitempty"`
	RequesterTransitGatewayId           string       `json:"requesterTransitGatewayId,omitempty"`
	ServiceZoneId                       string       `json:"serviceZoneId,omitempty"`
	TransitGatewayConnectionId          string       `json:"transitGatewayConnectionId,omitempty"`
	TransitGatewayConnectionName        string       `json:"transitGatewayConnectionName,omitempty"`
	TransitGatewayConnectionState       string       `json:"transitGatewayConnectionState,omitempty"`
	TransitGatewayConnectionType        string       `json:"transitGatewayConnectionType,omitempty"`
	TransitGatewayConnectionDescription string       `json:"transitGatewayConnectionDescription,omitempty"`
	CreatedBy                           string       `json:"createdBy,omitempty"`
	CreatedDt                           time.Time `json:"createdDt,omitempty"`
	ModifiedBy                          string       `json:"modifiedBy,omitempty"`
	ModifiedDt                          time.Time `json:"modifiedDt,omitempty"`
}

type TransitGatewayConnectionListItemResponse struct {
	ProjectId                           string       `json:"projectId,omitempty"`
	ApproverProjectId                   string       `json:"approverProjectId,omitempty"`
	ApproverVpcId                       string       `json:"approverVpcId,omitempty"`
	BlockId                             string       `json:"blockId,omitempty"`
	CompletedDt                         time.Time `json:"completedDt,omitempty"`
	RequesterProjectId                  string       `json:"requesterProjectId,omitempty"`
	RequesterTransitGatewayId           string       `json:"requesterTransitGatewayId,omitempty"`
	ServiceZoneId                       string       `json:"serviceZoneId,omitempty"`
	TransitGatewayConnectionId          string       `json:"transitGatewayConnectionId,omitempty"`
	TransitGatewayConnectionName        string       `json:"transitGatewayConnectionName,omitempty"`
	TransitGatewayConnectionState       string       `json:"transitGatewayConnectionState,omitempty"`
	TransitGatewayConnectionType        string       `json:"transitGatewayConnectionType,omitempty"`
	TransitGatewayConnectionDescription string       `json:"transitGatewayConnectionDescription,omitempty"`
	CreatedBy                           string       `json:"createdBy,omitempty"`
	CreatedDt                           time.Time `json:"createdDt,omitempty"`
	ModifiedBy                          string       `json:"modifiedBy,omitempty"`
	ModifiedDt                          time.Time `json:"modifiedDt,omitempty"`
}

type TransitGatewayConnectionOpenApiControllerApiListTransitGatewayConnectionsOpts struct {
	ApproverVpcId                  optional.String
	RequesterTransitGatewayId      optional.String
	TransitGatewayConnectionName   optional.String
	TransitGatewayConnectionStates optional.Interface
	CreatedBy                      optional.String
	Page                           optional.Int32
	Size                           optional.Int32
	Sort                           optional.Interface
}

type TransitGatewayCreateV3Request struct {
	BandwidthGbps             int32                           `json:"bandwidthGbps"`
	ServiceZoneId             string                          `json:"serviceZoneId"`
	Tags                      []TagRequest `json:"tags,omitempty"`
	TransitGatewayName        string                          `json:"transitGatewayName"`
	UplinkEnabled             *bool                           `json:"uplinkEnabled,omitempty"`
	TransitGatewayDescription string                          `json:"transitGatewayDescription,omitempty"`
}

type TransitGatewayDescriptionUpdateRequest struct {
	TransitGatewayDescription string `json:"transitGatewayDescription,omitempty"`
}

type TransitGatewayDetailResponse struct {
	ProjectId                   string       `json:"projectId,omitempty"`
	BandwidthGbps               int32        `json:"bandwidthGbps,omitempty"`
	BlockId                     string       `json:"blockId,omitempty"`
	ProductGroupId              string       `json:"productGroupId,omitempty"`
	ScaleProductId              string       `json:"scaleProductId,omitempty"`
	ServiceZoneId               string       `json:"serviceZoneId,omitempty"`
	TransitGatewayConnectionIds []string     `json:"transitGatewayConnectionIds,omitempty"`
	TransitGatewayDisplayName   string       `json:"transitGatewayDisplayName,omitempty"`
	TransitGatewayId            string       `json:"transitGatewayId,omitempty"`
	TransitGatewayName          string       `json:"transitGatewayName,omitempty"`
	TransitGatewayPeeringIds    []string     `json:"transitGatewayPeeringIds,omitempty"`
	TransitGatewayState         string       `json:"transitGatewayState,omitempty"`
	UplinkEnabled               *bool        `json:"uplinkEnabled,omitempty"`
	TransitGatewayDescription   string       `json:"transitGatewayDescription,omitempty"`
	CreatedBy                   string       `json:"createdBy,omitempty"`
	CreatedDt                   time.Time `json:"createdDt,omitempty"`
	ModifiedBy                  string       `json:"modifiedBy,omitempty"`
	ModifiedDt                  time.Time `json:"modifiedDt,omitempty"`
}

type TransitGatewayListItemResponse struct {
	ProjectId                 string       `json:"projectId,omitempty"`
	BandwidthGbps             int32        `json:"bandwidthGbps,omitempty"`
	BlockId                   string       `json:"blockId,omitempty"`
	ProductGroupId            string       `json:"productGroupId,omitempty"`
	ServiceZoneId             string       `json:"serviceZoneId,omitempty"`
	TransitGatewayId          string       `json:"transitGatewayId,omitempty"`
	TransitGatewayName        string       `json:"transitGatewayName,omitempty"`
	TransitGatewayState       string       `json:"transitGatewayState,omitempty"`
	UplinkEnabled             *bool        `json:"uplinkEnabled,omitempty"`
	VpcCount                  int64        `json:"vpcCount,omitempty"`
	TransitGatewayDescription string       `json:"transitGatewayDescription,omitempty"`
	CreatedBy                 string       `json:"createdBy,omitempty"`
	CreatedDt                 time.Time `json:"createdDt,omitempty"`
	ModifiedBy                string       `json:"modifiedBy,omitempty"`
	ModifiedDt                time.Time `json:"modifiedDt,omitempty"`
}

type TransitGatewayOpenApiControllerApiListTransitGatewaysOpts struct {
	ServiceZoneId        optional.Interface
	TransitGatewayId     optional.String
	TransitGatewayName   optional.String
	TransitGatewayStates optional.Interface
	CreatedBy            optional.String
	Page                 optional.Int32
	Size                 optional.Int32
	Sort                 optional.Interface
}

type TransitGatewayPeeringCreateRequest struct {
	ApproverProjectId                string                          `json:"approverProjectId"`
	ApproverTransitGatewayId         string                          `json:"approverTransitGatewayId"`
	RequesterProjectId               string                          `json:"requesterProjectId"`
	RequesterTransitGatewayId        string                          `json:"requesterTransitGatewayId"`
	Tags                             []TagRequest `json:"tags,omitempty"`
	TransitGatewayPeeringDescription string                          `json:"transitGatewayPeeringDescription,omitempty"`
}

type TransitGatewayPeeringDescriptionUpdateRequest struct {
	TransitGatewayPeeringDescription string `json:"transitGatewayPeeringDescription,omitempty"`
}

type TransitGatewayPeeringDetailResponse struct {
	ProjectId                        string       `json:"projectId,omitempty"`
	ApprovedBy                       string       `json:"approvedBy,omitempty"`
	ApprovedDt                       time.Time `json:"approvedDt,omitempty"`
	ApproverProjectId                string       `json:"approverProjectId,omitempty"`
	ApproverProjectName              string       `json:"approverProjectName,omitempty"`
	ApproverTransitGatewayId         string       `json:"approverTransitGatewayId,omitempty"`
	ApproverTransitGatewayName       string       `json:"approverTransitGatewayName,omitempty"`
	ApproverZoneId                   string       `json:"approverZoneId,omitempty"`
	Automated                        *bool        `json:"automated,omitempty"`
	BlockId                          string       `json:"blockId,omitempty"`
	CompletedDt                      time.Time `json:"completedDt,omitempty"`
	RequestedBy                      string       `json:"requestedBy,omitempty"`
	RequestedDt                      time.Time `json:"requestedDt,omitempty"`
	RequesterProjectId               string       `json:"requesterProjectId,omitempty"`
	RequesterProjectName             string       `json:"requesterProjectName,omitempty"`
	RequesterTransitGatewayId        string       `json:"requesterTransitGatewayId,omitempty"`
	RequesterTransitGatewayName      string       `json:"requesterTransitGatewayName,omitempty"`
	RequesterZoneId                  string       `json:"requesterZoneId,omitempty"`
	ServiceZoneId                    string       `json:"serviceZoneId,omitempty"`
	TransitGatewayPeeringId          string       `json:"transitGatewayPeeringId,omitempty"`
	TransitGatewayPeeringName        string       `json:"transitGatewayPeeringName,omitempty"`
	TransitGatewayPeeringState       string       `json:"transitGatewayPeeringState,omitempty"`
	TransitGatewayPeeringDescription string       `json:"transitGatewayPeeringDescription,omitempty"`
	CreatedBy                        string       `json:"createdBy,omitempty"`
	CreatedDt                        time.Time `json:"createdDt,omitempty"`
	ModifiedBy                       string       `json:"modifiedBy,omitempty"`
	ModifiedDt                       time.Time `json:"modifiedDt,omitempty"`
}

type TransitGatewayPeeringListItemResponse struct {
	ProjectId                        string       `json:"projectId,omitempty"`
	ApproverProjectId                string       `json:"approverProjectId,omitempty"`
	ApproverTransitGatewayId         string       `json:"approverTransitGatewayId,omitempty"`
	ApproverZoneId                   string       `json:"approverZoneId,omitempty"`
	Automated                        *bool        `json:"automated,omitempty"`
	BlockId                          string       `json:"blockId,omitempty"`
	CompletedDt                      time.Time `json:"completedDt,omitempty"`
	RequesterProjectId               string       `json:"requesterProjectId,omitempty"`
	RequesterTransitGatewayId        string       `json:"requesterTransitGatewayId,omitempty"`
	RequesterZoneId                  string       `json:"requesterZoneId,omitempty"`
	ServiceZoneId                    string       `json:"serviceZoneId,omitempty"`
	TransitGatewayPeeringId          string       `json:"transitGatewayPeeringId,omitempty"`
	TransitGatewayPeeringName        string       `json:"transitGatewayPeeringName,omitempty"`
	TransitGatewayPeeringState       string       `json:"transitGatewayPeeringState,omitempty"`
	TransitGatewayPeeringDescription string       `json:"transitGatewayPeeringDescription,omitempty"`
	CreatedBy                        string       `json:"createdBy,omitempty"`
	CreatedDt                        time.Time `json:"createdDt,omitempty"`
	ModifiedBy                       string       `json:"modifiedBy,omitempty"`
	ModifiedDt                       time.Time `json:"modifiedDt,omitempty"`
}

type TransitGatewayPeeringOpenApiControllerApiListTransitGatewayPeeringOpts struct {
	ApproverTransitGatewayId  optional.String
	RequesterTransitGatewayId optional.String
	TransitGatewayPeeringName optional.String
	CreatedBy                 optional.String
	Page                      optional.Int32
	Size                      optional.Int32
	Sort                      optional.Interface
}

type TransitGatewayPeeringResponse struct {
	Success                 *bool  `json:"success,omitempty"`
	TransitGatewayPeeringId string `json:"transitGatewayPeeringId,omitempty"`
}

type TransitGatewayUplinkUpdateRequest struct {
	UplinkEnabled *bool `json:"uplinkEnabled,omitempty"`
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
