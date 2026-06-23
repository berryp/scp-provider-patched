package directconnect2

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

type DirectConnectConnectionCreateRequest struct {
	ApproverProjectId                  string                         `json:"approverProjectId"`
	ApproverVpcId                      string                         `json:"approverVpcId"`
	DirectConnectConnectionType        string                         `json:"directConnectConnectionType"`
	FirewallEnabled                    *bool                          `json:"firewallEnabled"`
	FirewallLoggable                   *bool                          `json:"firewallLoggable,omitempty"`
	RequesterDirectConnectId           string                         `json:"requesterDirectConnectId"`
	RequesterProjectId                 string                         `json:"requesterProjectId"`
	Tags                               []TagRequest `json:"tags,omitempty"`
	DirectConnectConnectionDescription string                         `json:"directConnectConnectionDescription,omitempty"`
}

type DirectConnectConnectionDescriptionUpdateRequest struct {
	DirectConnectConnectionDescription string `json:"directConnectConnectionDescription,omitempty"`
}

type DirectConnectConnectionDetailResponse struct {
	ProjectId                          string       `json:"projectId,omitempty"`
	ApproverProjectId                  string       `json:"approverProjectId,omitempty"`
	ApproverVpcId                      string       `json:"approverVpcId,omitempty"`
	BlockId                            string       `json:"blockId,omitempty"`
	CompletedDt                        time.Time `json:"completedDt,omitempty"`
	DirectConnectConnectionId          string       `json:"directConnectConnectionId,omitempty"`
	DirectConnectConnectionName        string       `json:"directConnectConnectionName,omitempty"`
	DirectConnectConnectionState       string       `json:"directConnectConnectionState,omitempty"`
	DirectConnectConnectionType        string       `json:"directConnectConnectionType,omitempty"`
	RequesterDirectConnectId           string       `json:"requesterDirectConnectId,omitempty"`
	RequesterProjectId                 string       `json:"requesterProjectId,omitempty"`
	ServiceZoneId                      string       `json:"serviceZoneId,omitempty"`
	DirectConnectConnectionDescription string       `json:"directConnectConnectionDescription,omitempty"`
	CreatedBy                          string       `json:"createdBy,omitempty"`
	CreatedDt                          time.Time `json:"createdDt,omitempty"`
	ModifiedBy                         string       `json:"modifiedBy,omitempty"`
	ModifiedDt                         time.Time `json:"modifiedDt,omitempty"`
}

type DirectConnectConnectionListResponse struct {
	ProjectId                          string       `json:"projectId,omitempty"`
	ApproverProjectId                  string       `json:"approverProjectId,omitempty"`
	ApproverVpcId                      string       `json:"approverVpcId,omitempty"`
	BlockId                            string       `json:"blockId,omitempty"`
	CompletedDt                        time.Time `json:"completedDt,omitempty"`
	DirectConnectConnectionId          string       `json:"directConnectConnectionId,omitempty"`
	DirectConnectConnectionName        string       `json:"directConnectConnectionName,omitempty"`
	DirectConnectConnectionState       string       `json:"directConnectConnectionState,omitempty"`
	DirectConnectConnectionType        string       `json:"directConnectConnectionType,omitempty"`
	RequesterDirectConnectId           string       `json:"requesterDirectConnectId,omitempty"`
	RequesterProjectId                 string       `json:"requesterProjectId,omitempty"`
	ServiceZoneId                      string       `json:"serviceZoneId,omitempty"`
	DirectConnectConnectionDescription string       `json:"directConnectConnectionDescription,omitempty"`
	CreatedBy                          string       `json:"createdBy,omitempty"`
	CreatedDt                          time.Time `json:"createdDt,omitempty"`
	ModifiedBy                         string       `json:"modifiedBy,omitempty"`
	ModifiedDt                         time.Time `json:"modifiedDt,omitempty"`
}

type DirectConnectConnectionOpenApiControllerApiListDirectConnectConnectionsOpts struct {
	ApproverVpcId                 optional.String
	DirectConnectConnectionName   optional.String
	DirectConnectConnectionStates optional.Interface
	RequesterDirectConnectId      optional.String
	CreatedBy                     optional.String
	Page                          optional.Int32
	Size                          optional.Int32
	Sort                          optional.Interface
}

type DirectConnectCreateV3Request struct {
	BandwidthGbps            int32                          `json:"bandwidthGbps"`
	DirectConnectName        string                         `json:"directConnectName"`
	ServiceZoneId            string                         `json:"serviceZoneId"`
	Tags                     []TagRequest `json:"tags,omitempty"`
	DirectConnectDescription string                         `json:"directConnectDescription"`
}

type DirectConnectDetailResponse struct {
	ProjectId                string                                                             `json:"projectId,omitempty"`
	BandwidthGbps            int32                                                              `json:"bandwidthGbps,omitempty"`
	BlockId                  string                                                             `json:"blockId,omitempty"`
	DirectConnectConnections *ListResponseDirectConnectConnectionListResponse `json:"directConnectConnections,omitempty"`
	DirectConnectId          string                                                             `json:"directConnectId,omitempty"`
	DirectConnectName        string                                                             `json:"directConnectName,omitempty"`
	DirectConnectState       string                                                             `json:"directConnectState,omitempty"`
	ProductGroupId           string                                                             `json:"productGroupId,omitempty"`
	ScaleProductId           string                                                             `json:"scaleProductId,omitempty"`
	ServiceZoneId            string                                                             `json:"serviceZoneId,omitempty"`
	UplinkEnabled            *bool                                                              `json:"uplinkEnabled,omitempty"`
	DirectConnectDescription string                                                             `json:"directConnectDescription,omitempty"`
	CreatedBy                string                                                             `json:"createdBy,omitempty"`
	CreatedDt                time.Time                                                       `json:"createdDt,omitempty"`
	ModifiedBy               string                                                             `json:"modifiedBy,omitempty"`
	ModifiedDt               time.Time                                                       `json:"modifiedDt,omitempty"`
}

type DirectConnectListItemResponse struct {
	ProjectId                string       `json:"projectId,omitempty"`
	BandwidthGbps            int32        `json:"bandwidthGbps,omitempty"`
	BlockId                  string       `json:"blockId,omitempty"`
	DirectConnectId          string       `json:"directConnectId,omitempty"`
	DirectConnectName        string       `json:"directConnectName,omitempty"`
	DirectConnectState       string       `json:"directConnectState,omitempty"`
	ProductGroupId           string       `json:"productGroupId,omitempty"`
	ServiceZoneId            string       `json:"serviceZoneId,omitempty"`
	UplinkEnabled            *bool        `json:"uplinkEnabled,omitempty"`
	DirectConnectDescription string       `json:"directConnectDescription,omitempty"`
	CreatedBy                string       `json:"createdBy,omitempty"`
	CreatedDt                time.Time `json:"createdDt,omitempty"`
	ModifiedBy               string       `json:"modifiedBy,omitempty"`
	ModifiedDt               time.Time `json:"modifiedDt,omitempty"`
}

type DirectConnectOpenApiControllerApiListDirectConnectsOpts struct {
	DirectConnectId   optional.String
	DirectConnectName optional.String
	CreatedBy         optional.String
	Page              optional.Int32
	Size              optional.Int32
	Sort              optional.Interface
}

type DirectConnectUpdateRequest struct {
	DirectConnectDescription string `json:"directConnectDescription,omitempty"`
}

type ListResponseDirectConnectConnectionListResponse struct {
	Contents   []DirectConnectConnectionListResponse `json:"contents,omitempty"`
	TotalCount int32                                                   `json:"totalCount,omitempty"`
}

type ListResponseDirectConnectListItemResponse struct {
	Contents   []DirectConnectListItemResponse `json:"contents,omitempty"`
	TotalCount int32                                             `json:"totalCount,omitempty"`
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
