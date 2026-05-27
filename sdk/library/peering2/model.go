package peering2

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

type ListResponseVpcPeeringResponse struct {
	Contents   []VpcPeeringResponse `json:"contents,omitempty"`
	TotalCount int32                            `json:"totalCount,omitempty"`
}

type TagRequest struct {
	TagKey   string `json:"tagKey"`
	TagValue string `json:"tagValue,omitempty"`
}

type VpcPeeringApprovalResponse struct {
	Success      *bool  `json:"success"`
	VpcPeeringId string `json:"vpcPeeringId"`
}

type VpcPeeringApproveRequest struct {
	FirewallEnabled  *bool `json:"firewallEnabled"`
	FirewallLoggable *bool `json:"firewallLoggable,omitempty"`
}

type VpcPeeringCreateV3Request struct {
	ApproverProjectId     string                   `json:"approverProjectId"`
	ApproverVpcId         string                   `json:"approverVpcId"`
	FirewallEnabled       *bool                    `json:"firewallEnabled"`
	FirewallLoggable      *bool                    `json:"firewallLoggable,omitempty"`
	RequesterProjectId    string                   `json:"requesterProjectId"`
	RequesterVpcId        string                   `json:"requesterVpcId"`
	Tags                  []TagRequest `json:"tags,omitempty"`
	VpcPeeringDescription string                   `json:"vpcPeeringDescription,omitempty"`
}

type VpcPeeringDetailResponse struct {
	ProjectId                   string       `json:"projectId,omitempty"`
	ApprovedBy                  string       `json:"approvedBy,omitempty"`
	ApprovedDt                  time.Time `json:"approvedDt,omitempty"`
	ApproverProjectId           string       `json:"approverProjectId,omitempty"`
	ApproverVpcFirewallEnabled  *bool        `json:"approverVpcFirewallEnabled,omitempty"`
	ApproverVpcId               string       `json:"approverVpcId,omitempty"`
	BlockId                     string       `json:"blockId,omitempty"`
	CompletedDt                 time.Time `json:"completedDt,omitempty"`
	ProductGroupId              string       `json:"productGroupId,omitempty"`
	RequestedBy                 string       `json:"requestedBy,omitempty"`
	RequestedDt                 time.Time `json:"requestedDt,omitempty"`
	RequesterProjectId          string       `json:"requesterProjectId,omitempty"`
	RequesterVpcFirewallEnabled *bool        `json:"requesterVpcFirewallEnabled,omitempty"`
	RequesterVpcId              string       `json:"requesterVpcId,omitempty"`
	ServiceZoneId               string       `json:"serviceZoneId,omitempty"`
	VpcPeeringId                string       `json:"vpcPeeringId,omitempty"`
	VpcPeeringName              string       `json:"vpcPeeringName,omitempty"`
	VpcPeeringState             string       `json:"vpcPeeringState,omitempty"`
	VpcPeeringType              string       `json:"vpcPeeringType,omitempty"`
	VpcPeeringDescription       string       `json:"vpcPeeringDescription,omitempty"`
	CreatedBy                   string       `json:"createdBy,omitempty"`
	CreatedDt                   time.Time `json:"createdDt,omitempty"`
	ModifiedBy                  string       `json:"modifiedBy,omitempty"`
	ModifiedDt                  time.Time `json:"modifiedDt,omitempty"`
}

type VpcPeeringOpenApiControllerApiListVpcPeeringsOpts struct {
	ApproverVpcId    optional.String
	RequesterVpcId   optional.String
	VpcPeeringName   optional.String
	VpcPeeringStates optional.Interface
	CreatedBy        optional.String
	Page             optional.Int32
	Size             optional.Int32
	Sort             optional.Interface
}

type VpcPeeringResponse struct {
	ApproverProjectId     string       `json:"approverProjectId,omitempty"`
	ApproverVpcId         string       `json:"approverVpcId,omitempty"`
	Automated             *bool        `json:"automated,omitempty"`
	CompletedDt           time.Time `json:"completedDt,omitempty"`
	RequesterProjectId    string       `json:"requesterProjectId,omitempty"`
	RequesterVpcId        string       `json:"requesterVpcId,omitempty"`
	VpcPeeringId          string       `json:"vpcPeeringId,omitempty"`
	VpcPeeringName        string       `json:"vpcPeeringName,omitempty"`
	VpcPeeringState       string       `json:"vpcPeeringState,omitempty"`
	VpcPeeringType        string       `json:"vpcPeeringType,omitempty"`
	VpcPeeringDescription string       `json:"vpcPeeringDescription,omitempty"`
	CreatedBy             string       `json:"createdBy,omitempty"`
	CreatedDt             time.Time `json:"createdDt,omitempty"`
	ModifiedBy            string       `json:"modifiedBy,omitempty"`
	ModifiedDt            time.Time `json:"modifiedDt,omitempty"`
}

type VpcPeeringUpdateDescriptionRequest struct {
	VpcPeeringDescription string `json:"vpcPeeringDescription,omitempty"`
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
