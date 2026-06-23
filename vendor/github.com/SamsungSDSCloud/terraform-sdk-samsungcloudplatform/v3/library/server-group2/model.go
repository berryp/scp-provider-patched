package servergroup2

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

type DetailServerGroupResponse struct {
	ProjectId          string       `json:"projectId,omitempty"`
	AffinityPolicyType string       `json:"affinityPolicyType,omitempty"`
	AutoscalingEnabled string       `json:"autoscalingEnabled,omitempty"`
	BlockId            string       `json:"blockId,omitempty"`
	DeploymentEnvType  string       `json:"deploymentEnvType,omitempty"`
	ImageId            string       `json:"imageId,omitempty"`
	Prefix             string       `json:"prefix,omitempty"`
	ProductGroupId     string       `json:"productGroupId,omitempty"`
	ServerGroupId      string       `json:"serverGroupId,omitempty"`
	ServerGroupName    string       `json:"serverGroupName,omitempty"`
	ServerGroupState   string       `json:"serverGroupState,omitempty"`
	ServerGroupType    string       `json:"serverGroupType,omitempty"`
	ServiceZoneId      string       `json:"serviceZoneId,omitempty"`
	ServicedFor        string       `json:"servicedFor,omitempty"`
	ServicedGroupFor   string       `json:"servicedGroupFor,omitempty"`
	CreatedBy          string       `json:"createdBy,omitempty"`
	CreatedDt          time.Time `json:"createdDt,omitempty"`
	ModifiedBy         string       `json:"modifiedBy,omitempty"`
	ModifiedDt         time.Time `json:"modifiedDt,omitempty"`
}

type ListResponseServerGroupsResponse struct {
	Contents   []ServerGroupsResponse `json:"contents,omitempty"`
	TotalCount int32                                  `json:"totalCount,omitempty"`
}

type ServerGroupCreateRequest struct {
	AffinityPolicyType string `json:"affinityPolicyType,omitempty"`
	DeploymentEnvType  string `json:"deploymentEnvType,omitempty"`
	ServerGroupName    string `json:"serverGroupName"`
	ServerGroupType    string `json:"serverGroupType,omitempty"`
	ServiceZoneId      string `json:"serviceZoneId"`
	ServicedFor        string `json:"servicedFor"`
	ServicedGroupFor   string `json:"servicedGroupFor"`
}

type ServerGroupSearchControllerApiListServerGroupOpts struct {
	ServerGroupName optional.String
	ServicedForList optional.Interface
	Page            optional.Int32
	Size            optional.Int32
	Sort            optional.Interface
}

type ServerGroupUpdateRequest struct {
	AffinityPolicyType string `json:"affinityPolicyType,omitempty"`
	DeploymentEnvType  string `json:"deploymentEnvType,omitempty"`
	ServerGroupName    string `json:"serverGroupName,omitempty"`
	ServerGroupType    string `json:"serverGroupType,omitempty"`
}

type ServerGroupsResponse struct {
	ProjectId          string       `json:"projectId,omitempty"`
	AffinityPolicyType string       `json:"affinityPolicyType,omitempty"`
	BlockId            string       `json:"blockId,omitempty"`
	DeploymentEnvType  string       `json:"deploymentEnvType,omitempty"`
	ServerGroupId      string       `json:"serverGroupId,omitempty"`
	ServerGroupName    string       `json:"serverGroupName,omitempty"`
	ServerGroupState   string       `json:"serverGroupState,omitempty"`
	ServerGroupType    string       `json:"serverGroupType,omitempty"`
	ServiceZoneId      string       `json:"serviceZoneId,omitempty"`
	ServicedFor        string       `json:"servicedFor,omitempty"`
	ServicedGroupFor   string       `json:"servicedGroupFor,omitempty"`
	CreatedBy          string       `json:"createdBy,omitempty"`
	CreatedDt          time.Time `json:"createdDt,omitempty"`
	ModifiedBy         string       `json:"modifiedBy,omitempty"`
	ModifiedDt         time.Time `json:"modifiedDt,omitempty"`
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
