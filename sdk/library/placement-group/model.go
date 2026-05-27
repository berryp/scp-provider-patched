package placementgroup

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

type CreateRequest struct {
	AvailabilityZoneName      string
	PlacementGroupName        string
	ServiceZoneId             string
	Tags                      map[string]interface{}
	VirtualServerType         string
	PlacementGroupDescription string
}

type ListPlacementGroupsRequestParam struct {
	PlacementGroupName      string
	PlacementGroupStateList []string
	ServiceZoneId           string
	VirtualServerType       string
	CreatedBy               string
	Page                    int32
	Size                    int32
	Sort                    string
}

type PageResponseV2PlacementGroupsResponse struct {
	Contents   []PlacementGroupsResponse `json:"contents,omitempty"`
	Page       int32                                       `json:"page,omitempty"`
	Size       int32                                       `json:"size,omitempty"`
	Sort       []string                                    `json:"sort,omitempty"`
	TotalCount int64                                       `json:"totalCount,omitempty"`
}

type PlacementGroupCreateRequest struct {
	AvailabilityZoneName      string                         `json:"availabilityZoneName,omitempty"`
	PlacementGroupName        string                         `json:"placementGroupName"`
	ServiceZoneId             string                         `json:"serviceZoneId"`
	Tags                      []TagRequest `json:"tags,omitempty"`
	VirtualServerType         string                         `json:"virtualServerType"`
	PlacementGroupDescription string                         `json:"placementGroupDescription,omitempty"`
}

type PlacementGroupDescriptionUpdateRequest struct {
	PlacementGroupDescription string `json:"placementGroupDescription,omitempty"`
}

type PlacementGroupDetailResponse struct {
	ProjectId                 string       `json:"projectId,omitempty"`
	AvailabilityZoneName      string       `json:"availabilityZoneName,omitempty"`
	PlacementGroupId          string       `json:"placementGroupId,omitempty"`
	PlacementGroupName        string       `json:"placementGroupName,omitempty"`
	PlacementGroupState       string       `json:"placementGroupState,omitempty"`
	ServiceZoneId             string       `json:"serviceZoneId,omitempty"`
	VirtualServerIdList       []string     `json:"virtualServerIdList,omitempty"`
	VirtualServerType         string       `json:"virtualServerType,omitempty"`
	PlacementGroupDescription string       `json:"placementGroupDescription,omitempty"`
	CreatedBy                 string       `json:"createdBy,omitempty"`
	CreatedDt                 time.Time `json:"createdDt,omitempty"`
	ModifiedBy                string       `json:"modifiedBy,omitempty"`
	ModifiedDt                time.Time `json:"modifiedDt,omitempty"`
}

type PlacementGroupV1ApiListPlacementGroups1Opts struct {
	PlacementGroupName      optional.String
	PlacementGroupStateList optional.Interface
	ServiceZoneId           optional.String
	VirtualServerType       optional.String
	CreatedBy               optional.String
	Page                    optional.Int32
	Size                    optional.Int32
	Sort                    optional.Interface
}

type PlacementGroupsResponse struct {
	AvailabilityZoneName string       `json:"availabilityZoneName,omitempty"`
	PlacementGroupId     string       `json:"placementGroupId,omitempty"`
	PlacementGroupName   string       `json:"placementGroupName,omitempty"`
	PlacementGroupState  string       `json:"placementGroupState,omitempty"`
	VirtualServerIdList  []string     `json:"virtualServerIdList,omitempty"`
	CreatedBy            string       `json:"createdBy,omitempty"`
	CreatedDt            time.Time `json:"createdDt,omitempty"`
	ModifiedBy           string       `json:"modifiedBy,omitempty"`
	ModifiedDt           time.Time `json:"modifiedDt,omitempty"`
}

type TagRequest struct {
	TagKey   string `json:"tagKey"`
	TagValue string `json:"tagValue,omitempty"`
}

type UpdatePlacementGroupDescriptionRequest struct {
	PlacementGroupDescription string
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
