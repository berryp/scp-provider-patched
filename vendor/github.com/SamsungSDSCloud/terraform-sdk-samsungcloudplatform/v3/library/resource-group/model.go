package resourcegroup

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

type ListResourceGroupRequest struct {
	CreatedById       string
	ModifiedByEmail   string
	ModifiedById      string
	ResourceGroupName string
}

type ListResourceGroupResourcesRequest struct {
	CreatedById  string
	ModifiedById string
	ResourceId   string
	ResourceName string
}

type ListResourceRequest struct {
	CreatedById         string
	DisplayServiceNames []interface{}
	FromCreatedAt       string
	IncludeDeleted      string
	Location            string
	ModifiedById        string
	MyCreate            string
	Partitions          []interface{}
	Regions             []interface{}
	ResourceId          string
	ResourceName        string
	ResourceTypes       []interface{}
	ServiceTypes        []interface{}
	ServiceZones        []interface{}
	Tags                []interface{}
	ToCreatedAt         string
}

type MyProjectResourceControllerApiDetailMyProjectsResourceOpts struct {
	IncludeDeleted optional.String
}

type MyProjectResourceControllerApiListMyProjectsResourcesOpts struct {
	CreatedById         optional.String
	DisplayServiceNames optional.Interface
	FromCreatedAt       optional.String
	IncludeDeleted      optional.String
	Location            optional.String
	ModifiedById        optional.String
	MyCreate            optional.String
	Partitions          optional.Interface
	ProjectIds          optional.Interface
	Regions             optional.Interface
	ResourceId          optional.String
	ResourceName        optional.String
	ResourceTypes       optional.Interface
	ServiceTypes        optional.Interface
	ServiceZones        optional.Interface
	Tags                optional.Interface
	ToCreatedAt         optional.String
	Page                optional.Int32
	Size                optional.Int32
	Sort                optional.Interface
}

type MyProjectResourceGroupControllerApiListMyProjectsResourceGroupResourcesOpts struct {
	CreatedById  optional.String
	ModifiedById optional.String
	ResourceId   optional.String
	ResourceName optional.String
	Page         optional.Int32
	Size         optional.Int32
	Sort         optional.Interface
}

type MyProjectResourceGroupControllerApiListMyProjectsResourceGroupsOpts struct {
	CreatedById       optional.String
	ModifiedByEmail   optional.String
	ModifiedById      optional.String
	ProjectIds        optional.Interface
	ResourceGroupName optional.String
	Page              optional.Int32
	Size              optional.Int32
	Sort              optional.Interface
}

type MyProjectsResourceGroupResponse struct {
	ProjectId                string                 `json:"projectId,omitempty"`
	AccountName              string                 `json:"accountName,omitempty"`
	CreatedById              string                 `json:"createdById,omitempty"`
	ModifiedById             string                 `json:"modifiedById,omitempty"`
	ProjectName              string                 `json:"projectName,omitempty"`
	ResourceGroupId          string                 `json:"resourceGroupId,omitempty"`
	ResourceGroupName        string                 `json:"resourceGroupName,omitempty"`
	TargetResourceTags       []Tag `json:"targetResourceTags,omitempty"`
	TargetResourceTypes      []string               `json:"targetResourceTypes,omitempty"`
	ResourceGroupDescription string                 `json:"resourceGroupDescription,omitempty"`
	CreatedByName            string                 `json:"createdByName,omitempty"`
	CreatedByEmail           string                 `json:"createdByEmail,omitempty"`
	CreatedDt                string                 `json:"createdDt,omitempty"`
	ModifiedByName           string                 `json:"modifiedByName,omitempty"`
	ModifiedByEmail          string                 `json:"modifiedByEmail,omitempty"`
	ModifiedDt               string                 `json:"modifiedDt,omitempty"`
}

type MyProjectsResourceGroupsResponse struct {
	ProjectId                string `json:"projectId,omitempty"`
	AccountName              string `json:"accountName,omitempty"`
	CreatedById              string `json:"createdById,omitempty"`
	ModifiedById             string `json:"modifiedById,omitempty"`
	ProjectName              string `json:"projectName,omitempty"`
	ResourceGroupId          string `json:"resourceGroupId,omitempty"`
	ResourceGroupName        string `json:"resourceGroupName,omitempty"`
	ResourceGroupDescription string `json:"resourceGroupDescription,omitempty"`
	CreatedByName            string `json:"createdByName,omitempty"`
	CreatedByEmail           string `json:"createdByEmail,omitempty"`
	CreatedDt                string `json:"createdDt,omitempty"`
	ModifiedByName           string `json:"modifiedByName,omitempty"`
	ModifiedByEmail          string `json:"modifiedByEmail,omitempty"`
	ModifiedDt               string `json:"modifiedDt,omitempty"`
}

type MyProjectsResourceResponse struct {
	ProjectId       string                 `json:"projectId,omitempty"`
	AccountName     string                 `json:"accountName,omitempty"`
	EventState      string                 `json:"eventState,omitempty"`
	Partition       string                 `json:"partition,omitempty"`
	ProjectName     string                 `json:"projectName,omitempty"`
	Region          string                 `json:"region,omitempty"`
	ResourceId      string                 `json:"resourceId,omitempty"`
	ResourceName    string                 `json:"resourceName,omitempty"`
	ResourceSrn     string                 `json:"resourceSrn,omitempty"`
	ResourceState   string                 `json:"resourceState,omitempty"`
	ResourceType    string                 `json:"resourceType,omitempty"`
	ServiceType     string                 `json:"serviceType,omitempty"`
	Tags            []Tag `json:"tags,omitempty"`
	Zone            string                 `json:"zone,omitempty"`
	CreatedBy       string                 `json:"createdBy,omitempty"`
	CreatedByName   string                 `json:"createdByName,omitempty"`
	CreatedByEmail  string                 `json:"createdByEmail,omitempty"`
	CreatedDt       time.Time           `json:"createdDt,omitempty"`
	ModifiedBy      string                 `json:"modifiedBy,omitempty"`
	ModifiedByName  string                 `json:"modifiedByName,omitempty"`
	ModifiedByEmail string                 `json:"modifiedByEmail,omitempty"`
	ModifiedDt      time.Time           `json:"modifiedDt,omitempty"`
}

type MyProjectsResourcesResponse struct {
	ProjectId       string                 `json:"projectId,omitempty"`
	AccountName     string                 `json:"accountName,omitempty"`
	EventState      string                 `json:"eventState,omitempty"`
	Partition       string                 `json:"partition,omitempty"`
	ProjectName     string                 `json:"projectName,omitempty"`
	Region          string                 `json:"region,omitempty"`
	ResourceId      string                 `json:"resourceId,omitempty"`
	ResourceName    string                 `json:"resourceName,omitempty"`
	ResourceSrn     string                 `json:"resourceSrn,omitempty"`
	ResourceState   string                 `json:"resourceState,omitempty"`
	ResourceType    string                 `json:"resourceType,omitempty"`
	ServiceType     string                 `json:"serviceType,omitempty"`
	Tags            []Tag `json:"tags,omitempty"`
	Zone            string                 `json:"zone,omitempty"`
	CreatedBy       string                 `json:"createdBy,omitempty"`
	CreatedByName   string                 `json:"createdByName,omitempty"`
	CreatedByEmail  string                 `json:"createdByEmail,omitempty"`
	CreatedDt       time.Time           `json:"createdDt,omitempty"`
	ModifiedBy      string                 `json:"modifiedBy,omitempty"`
	ModifiedByName  string                 `json:"modifiedByName,omitempty"`
	ModifiedByEmail string                 `json:"modifiedByEmail,omitempty"`
	ModifiedDt      time.Time           `json:"modifiedDt,omitempty"`
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

type ResourceControllerApiDetailResourceOpts struct {
	IncludeDeleted optional.String
}

type ResourceControllerApiListResourcesOpts struct {
	CreatedById         optional.String
	DisplayServiceNames optional.Interface
	FromCreatedAt       optional.String
	IncludeDeleted      optional.String
	Location            optional.String
	ModifiedById        optional.String
	MyCreate            optional.String
	Partitions          optional.Interface
	Regions             optional.Interface
	ResourceId          optional.String
	ResourceName        optional.String
	ResourceTypes       optional.Interface
	ServiceTypes        optional.Interface
	ServiceZones        optional.Interface
	Tags                optional.Interface
	ToCreatedAt         optional.String
	Page                optional.Int32
	Size                optional.Int32
	Sort                optional.Interface
}

type ResourceControllerApiResourceSrn1Opts struct {
	IncludeDeleted optional.String
}

type ResourceControllerApiResourceSrnDeprecatedOpts struct {
	IncludeDeleted optional.String
}

type ResourceControllerApiResourceSrnOpts struct {
	IncludeDeleted optional.String
}

type ResourceGroupControllerApiListResourceGroupConditionResourcesOpts struct {
	CreatedById   optional.String
	ModifiedById  optional.String
	ResourceTypes optional.Interface
	Page          optional.Int32
	Size          optional.Int32
	Sort          optional.Interface
}

type ResourceGroupControllerApiListResourceGroupResourcesOpts struct {
	CreatedById  optional.String
	ModifiedById optional.String
	ResourceId   optional.String
	ResourceName optional.String
	Page         optional.Int32
	Size         optional.Int32
	Sort         optional.Interface
}

type ResourceGroupControllerApiListResourceGroupsOpts struct {
	CreatedById       optional.String
	ModifiedByEmail   optional.String
	ModifiedById      optional.String
	ResourceGroupName optional.String
	Page              optional.Int32
	Size              optional.Int32
	Sort              optional.Interface
}

type ResourceGroupCreateRequest struct {
	ResourceGroupName        string                 `json:"resourceGroupName"`
	TargetResourceTags       []Tag `json:"targetResourceTags"`
	TargetResourceTypes      []string               `json:"targetResourceTypes,omitempty"`
	ResourceGroupDescription string                 `json:"resourceGroupDescription,omitempty"`
}

type ResourceGroupRequest struct {
	ResourceGroupName        string
	TargetResourceTags       []interface{}
	TargetResourceTypes      []interface{}
	ResourceGroupDescription string
}

type ResourceGroupResponse struct {
	CreatedById              string                 `json:"createdById,omitempty"`
	ModifiedById             string                 `json:"modifiedById,omitempty"`
	ResourceGroupId          string                 `json:"resourceGroupId,omitempty"`
	ResourceGroupName        string                 `json:"resourceGroupName,omitempty"`
	TargetResourceTag        []Tag `json:"targetResourceTag,omitempty"`
	TargetResourceTypes      []string               `json:"targetResourceTypes,omitempty"`
	ResourceGroupDescription string                 `json:"resourceGroupDescription,omitempty"`
	CreatedByName            string                 `json:"createdByName,omitempty"`
	CreatedByEmail           string                 `json:"createdByEmail,omitempty"`
	CreatedDt                string                 `json:"createdDt,omitempty"`
	ModifiedByName           string                 `json:"modifiedByName,omitempty"`
	ModifiedByEmail          string                 `json:"modifiedByEmail,omitempty"`
	ModifiedDt               string                 `json:"modifiedDt,omitempty"`
}

type ResourceGroupUpdateRequest struct {
	ResourceGroupName        string                 `json:"resourceGroupName"`
	TargetResourceTags       []Tag `json:"targetResourceTags"`
	TargetResourceTypes      []string               `json:"targetResourceTypes,omitempty"`
	ResourceGroupDescription string                 `json:"resourceGroupDescription,omitempty"`
}

type ResourceGroupsResponse struct {
	CreatedById              string `json:"createdById,omitempty"`
	ModifiedById             string `json:"modifiedById,omitempty"`
	ResourceGroupId          string `json:"resourceGroupId,omitempty"`
	ResourceGroupName        string `json:"resourceGroupName,omitempty"`
	ResourceGroupDescription string `json:"resourceGroupDescription,omitempty"`
	CreatedByName            string `json:"createdByName,omitempty"`
	CreatedByEmail           string `json:"createdByEmail,omitempty"`
	CreatedDt                string `json:"createdDt,omitempty"`
	ModifiedByName           string `json:"modifiedByName,omitempty"`
	ModifiedByEmail          string `json:"modifiedByEmail,omitempty"`
	ModifiedDt               string `json:"modifiedDt,omitempty"`
}

type ResourceResponse struct {
	EventState      string                 `json:"eventState,omitempty"`
	Partition       string                 `json:"partition,omitempty"`
	Region          string                 `json:"region,omitempty"`
	ResourceId      string                 `json:"resourceId,omitempty"`
	ResourceName    string                 `json:"resourceName,omitempty"`
	ResourceSrn     string                 `json:"resourceSrn,omitempty"`
	ResourceState   string                 `json:"resourceState,omitempty"`
	ResourceType    string                 `json:"resourceType,omitempty"`
	ServiceType     string                 `json:"serviceType,omitempty"`
	Tags            []Tag `json:"tags,omitempty"`
	Zone            string                 `json:"zone,omitempty"`
	CreatedBy       string                 `json:"createdBy,omitempty"`
	CreatedByName   string                 `json:"createdByName,omitempty"`
	CreatedByEmail  string                 `json:"createdByEmail,omitempty"`
	CreatedDt       time.Time           `json:"createdDt,omitempty"`
	ModifiedBy      string                 `json:"modifiedBy,omitempty"`
	ModifiedByName  string                 `json:"modifiedByName,omitempty"`
	ModifiedByEmail string                 `json:"modifiedByEmail,omitempty"`
	ModifiedDt      time.Time           `json:"modifiedDt,omitempty"`
}

type ResourceTypeResponse struct {
	ResourceType string `json:"resourceType,omitempty"`
	ServiceType  string `json:"serviceType,omitempty"`
	TagPolicy    *bool  `json:"tagPolicy,omitempty"`
}

type ResourcesResponse struct {
	EventState      string                 `json:"eventState,omitempty"`
	Partition       string                 `json:"partition,omitempty"`
	Region          string                 `json:"region,omitempty"`
	ResourceId      string                 `json:"resourceId,omitempty"`
	ResourceName    string                 `json:"resourceName,omitempty"`
	ResourceSrn     string                 `json:"resourceSrn,omitempty"`
	ResourceState   string                 `json:"resourceState,omitempty"`
	ResourceType    string                 `json:"resourceType,omitempty"`
	ServiceType     string                 `json:"serviceType,omitempty"`
	Tags            []Tag `json:"tags,omitempty"`
	Zone            string                 `json:"zone,omitempty"`
	CreatedBy       string                 `json:"createdBy,omitempty"`
	CreatedByName   string                 `json:"createdByName,omitempty"`
	CreatedByEmail  string                 `json:"createdByEmail,omitempty"`
	CreatedDt       time.Time           `json:"createdDt,omitempty"`
	ModifiedBy      string                 `json:"modifiedBy,omitempty"`
	ModifiedByName  string                 `json:"modifiedByName,omitempty"`
	ModifiedByEmail string                 `json:"modifiedByEmail,omitempty"`
	ModifiedDt      time.Time           `json:"modifiedDt,omitempty"`
}

type RgPageResponseMyProjectsResourceGroupsResponse struct {
	Contents     []MyProjectsResourceGroupsResponse `json:"contents,omitempty"`
	CreatedById  string                                              `json:"createdById,omitempty"`
	ModifiedById string                                              `json:"modifiedById,omitempty"`
	Page         int32                                               `json:"page,omitempty"`
	Size         int32                                               `json:"size,omitempty"`
	Sort         []string                                            `json:"sort,omitempty"`
	TotalCount   int64                                               `json:"totalCount,omitempty"`
}

type RgPageResponseMyProjectsResourcesResponse struct {
	Contents     []MyProjectsResourcesResponse `json:"contents,omitempty"`
	CreatedById  string                                         `json:"createdById,omitempty"`
	ModifiedById string                                         `json:"modifiedById,omitempty"`
	Page         int32                                          `json:"page,omitempty"`
	Size         int32                                          `json:"size,omitempty"`
	Sort         []string                                       `json:"sort,omitempty"`
	TotalCount   int64                                          `json:"totalCount,omitempty"`
}

type RgPageResponseResourceGroupsResponse struct {
	Contents     []ResourceGroupsResponse `json:"contents,omitempty"`
	CreatedById  string                                    `json:"createdById,omitempty"`
	ModifiedById string                                    `json:"modifiedById,omitempty"`
	Page         int32                                     `json:"page,omitempty"`
	Size         int32                                     `json:"size,omitempty"`
	Sort         []string                                  `json:"sort,omitempty"`
	TotalCount   int64                                     `json:"totalCount,omitempty"`
}

type RgPageResponseResourcesResponse struct {
	Contents     []ResourcesResponse `json:"contents,omitempty"`
	CreatedById  string                               `json:"createdById,omitempty"`
	ModifiedById string                               `json:"modifiedById,omitempty"`
	Page         int32                                `json:"page,omitempty"`
	Size         int32                                `json:"size,omitempty"`
	Sort         []string                             `json:"sort,omitempty"`
	TotalCount   int64                                `json:"totalCount,omitempty"`
}

type Tag struct {
	TagKey   string `json:"tagKey,omitempty"`
	TagValue string `json:"tagValue,omitempty"`
}

type TypeControllerApiListResourceTypesOpts struct {
	ResourceType optional.String
	ServiceType  optional.String
}

type TypeControllerApiListServiceTypeResourcesOpts struct {
	ResourceType optional.String
}

type TypeControllerApiListServiceTypesOpts struct {
	ServiceType optional.String
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
