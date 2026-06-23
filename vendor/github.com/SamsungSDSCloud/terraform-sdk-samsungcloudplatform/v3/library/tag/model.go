package tag

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

type CommonCodeResponse struct {
	ClassificationCode string `json:"classificationCode,omitempty"`
	Code               string `json:"code,omitempty"`
	Name               string `json:"name,omitempty"`
	Description        string `json:"description,omitempty"`
}

type Filter struct {
	TagKey    string
	TagValues []string
}

type PageResponseV2TagResponse struct {
	Contents   []TagResponse `json:"contents,omitempty"`
	Page       int32                `json:"page,omitempty"`
	Size       int32                `json:"size,omitempty"`
	Sort       []string             `json:"sort,omitempty"`
	TotalCount int64                `json:"totalCount,omitempty"`
}

type PageResponseV2TagsResponse struct {
	Contents   []TagsResponse `json:"contents,omitempty"`
	Page       int32                 `json:"page,omitempty"`
	Size       int32                 `json:"size,omitempty"`
	Sort       []string              `json:"sort,omitempty"`
	TotalCount int64                 `json:"totalCount,omitempty"`
}

type PlatformControllerApiDetailApiDocsOpts struct {
	ApiScopes            optional.Interface
	BuildApiResponse     optional.Bool
	DisplayFieldRefDesc  optional.Bool
	ForceGenMsgCode      optional.Bool
	IgnoreInternalHeader optional.Bool
	MsgCodePrefix        optional.String
	State                optional.Interface
	UriPrefix            optional.Interface
	UseCliMsgCode        optional.Bool
	Version              optional.Interface
}

type ResourceSearchCriteria struct {
	Page                int32              `json:"page,omitempty"`
	ResourceIds         []string           `json:"resourceIds,omitempty"`
	ResourceTypeFilters []string           `json:"resourceTypeFilters,omitempty"`
	Size                int32              `json:"size,omitempty"`
	Sort                []string           `json:"sort,omitempty"`
	TagFilters          []TagFilter `json:"tagFilters,omitempty"`
}

type ResourceTagControllerApiListResourceTagsOpts struct {
	Page optional.Int32
	Size optional.Int32
	Sort optional.Interface
}

type ResourceTagCreateRequest struct {
	Tags []TagRequest `json:"tags"`
}

type ResourceTagUpdateRequest struct {
	Tags []TagRequest `json:"tags"`
}

type ResourceTagValidateRequest struct {
	CreateOrUpdate []TagRequest `json:"createOrUpdate,omitempty"`
	Delete         []string            `json:"delete,omitempty"`
}

type ResourcesTagValidateRequest struct {
	ResourceIds []string            `json:"resourceIds,omitempty"`
	Tags        []TagRequest `json:"tags,omitempty"`
}

type TagFilter struct {
	TagKey    string   `json:"tagKey"`
	TagValues []string `json:"tagValues,omitempty"`
}

type TagPolicyEditRequest struct {
	Tags []TagPolicyStmtRequest `json:"tags,omitempty"`
}

type TagPolicyEditRequest1 struct {
	Tags []TagPolicyStmtRequest `json:"tags"`
}

type TagPolicyErrorResponse struct {
	Location string `json:"location,omitempty"`
	Message  string `json:"message,omitempty"`
}

type TagPolicyRequirementDetail struct {
	Active          *bool    `json:"active,omitempty"`
	DefaultTagValue string   `json:"defaultTagValue,omitempty"`
	ResourceTypes   []string `json:"resourceTypes,omitempty"`
	TagValues       []string `json:"tagValues,omitempty"`
}

type TagPolicyResponse struct {
	ProjectId       string                         `json:"projectId,omitempty"`
	TagPolicyId     string                         `json:"tagPolicyId,omitempty"`
	Tags            []TagPolicyStmtResponse `json:"tags,omitempty"`
	CreatedBy       string                         `json:"createdBy"`
	CreatedByName   string                         `json:"createdByName"`
	CreatedByEmail  string                         `json:"createdByEmail"`
	CreatedDt       time.Time                   `json:"createdDt"`
	ModifiedBy      string                         `json:"modifiedBy"`
	ModifiedByName  string                         `json:"modifiedByName"`
	ModifiedByEmail string                         `json:"modifiedByEmail"`
	ModifiedDt      time.Time                   `json:"modifiedDt"`
}

type TagPolicyStmtDetailRequest struct {
	Active          *bool    `json:"active"`
	DefaultTagValue string   `json:"defaultTagValue"`
	ResourceTypes   []string `json:"resourceTypes"`
	TagValues       []string `json:"tagValues"`
}

type TagPolicyStmtDetailResponse struct {
	Active          *bool    `json:"active,omitempty"`
	DefaultTagValue string   `json:"defaultTagValue,omitempty"`
	ResourceTypes   []string `json:"resourceTypes,omitempty"`
	TagValues       []string `json:"tagValues,omitempty"`
}

type TagPolicyStmtRequest struct {
	Active     *bool                               `json:"active"`
	Conditions []TagPolicyStmtDetailRequest `json:"conditions"`
	TagKey     string                              `json:"tagKey"`
}

type TagPolicyStmtResponse struct {
	Active     *bool                                `json:"active"`
	Conditions []TagPolicyStmtDetailResponse `json:"conditions"`
	TagKey     string                               `json:"tagKey,omitempty"`
}

type TagRequest struct {
	TagKey   string `json:"tagKey"`
	TagValue string `json:"tagValue,omitempty"`
}

type TagResponse struct {
	TagKey   string `json:"tagKey"`
	TagValue string `json:"tagValue,omitempty"`
}

type TagsResponse struct {
	ProjectId    string               `json:"projectId"`
	ResourceId   string               `json:"resourceId"`
	ResourceType string               `json:"resourceType"`
	Tags         []TagResponse `json:"tags,omitempty"`
}

type VersionModel struct {
	BuildDate    string `json:"buildDate,omitempty"`
	BuildProfile string `json:"buildProfile,omitempty"`
	Title        string `json:"title,omitempty"`
	Version      string `json:"version,omitempty"`
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
