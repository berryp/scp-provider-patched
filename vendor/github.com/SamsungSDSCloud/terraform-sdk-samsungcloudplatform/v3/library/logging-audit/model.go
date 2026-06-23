package loggingaudit

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

type CreateTrailRequest struct {
	IsLoggingTargetAllRegion   string                         `json:"isLoggingTargetAllRegion,omitempty"`
	IsLoggingTargetAllResource string                         `json:"isLoggingTargetAllResource,omitempty"`
	IsLoggingTargetAllUser     string                         `json:"isLoggingTargetAllUser,omitempty"`
	LoggingTargetProviders     []UserResponse `json:"loggingTargetProviders,omitempty"`
	LoggingTargetRegions       []string                       `json:"loggingTargetRegions"`
	LoggingTargetResourceIds   []string                       `json:"loggingTargetResourceIds,omitempty"`
	LoggingTargetUsers         []UserResponse `json:"loggingTargetUsers,omitempty"`
	ObsBucketId                string                         `json:"obsBucketId"`
	TagCreateRequests          []TagRequest   `json:"tagCreateRequests,omitempty"`
	TrailName                  string                         `json:"trailName"`
	TrailSaveType              string                         `json:"trailSaveType"`
	ValidationYn               string                         `json:"validationYn,omitempty"`
	TrailDescription           string                         `json:"trailDescription,omitempty"`
}

type LoggingControllerApiListLoggingMyProjectsStatisticsOpts struct {
	RequestEndDt   optional.Time
	RequestStartDt optional.Time
	Page           optional.Int32
	Size           optional.Int32
	Sort           optional.Interface
}

type LoggingControllerApiListLoggingStatisticsOpts struct {
	RequestEndDt   optional.Time
	RequestStartDt optional.Time
	Page           optional.Int32
	Size           optional.Int32
	Sort           optional.Interface
}

type LoggingSearchCriteria struct {
	LoggingObjectId          string       `json:"loggingObjectId,omitempty"`
	LoggingTargetProductName []string     `json:"loggingTargetProductName,omitempty"`
	LoggingTargetRegion      []string     `json:"loggingTargetRegion,omitempty"`
	LoggingTargetResource    []string     `json:"loggingTargetResource,omitempty"`
	ObjectName               string       `json:"objectName,omitempty"`
	Page                     int32        `json:"page,omitempty"`
	ProductOffering          string       `json:"productOffering,omitempty"`
	RequestClientType        string       `json:"requestClientType,omitempty"`
	RequestEndDt             time.Time `json:"requestEndDt,omitempty"`
	RequestStartDt           time.Time `json:"requestStartDt,omitempty"`
	RoleName                 string       `json:"roleName,omitempty"`
	Size                     int32        `json:"size,omitempty"`
	Sort                     []string     `json:"sort,omitempty"`
	State                    string       `json:"state,omitempty"`
	UserInfo                 string       `json:"userInfo,omitempty"`
	UserName                 string       `json:"userName,omitempty"`
}

type LoggingsResponse struct {
	ProjectId           string       `json:"projectId,omitempty"`
	AuditContent        string       `json:"auditContent,omitempty"`
	AuditDetailContent  string       `json:"auditDetailContent,omitempty"`
	ClusterId           string       `json:"clusterId,omitempty"`
	ClusterNamespaceId  string       `json:"clusterNamespaceId,omitempty"`
	EventTopicName      string       `json:"eventTopicName,omitempty"`
	Id                  string       `json:"id,omitempty"`
	LogErrorMessage     string       `json:"logErrorMessage,omitempty"`
	LoggingResourceType string       `json:"loggingResourceType,omitempty"`
	ObjectId            string       `json:"objectId,omitempty"`
	ObjectName          string       `json:"objectName,omitempty"`
	ProductName         string       `json:"productName,omitempty"`
	ProjectName         string       `json:"projectName,omitempty"`
	Region              string       `json:"region,omitempty"`
	RequestBy           string       `json:"requestBy,omitempty"`
	RequestClientType   string       `json:"requestClientType,omitempty"`
	RequestDt           time.Time `json:"requestDt,omitempty"`
	ResourceType        string       `json:"resourceType,omitempty"`
	RoleName            string       `json:"roleName,omitempty"`
	State               string       `json:"state,omitempty"`
	UserEmail           string       `json:"userEmail,omitempty"`
	UserInfo            string       `json:"userInfo,omitempty"`
	UserName            string       `json:"userName,omitempty"`
}

type MembersResponse struct {
	Email    string `json:"email,omitempty"`
	UserId   string `json:"userId,omitempty"`
	UserName string `json:"userName,omitempty"`
}

type PageResponseV2LoggingsResponse struct {
	Contents   []LoggingsResponse `json:"contents,omitempty"`
	Page       int32                              `json:"page,omitempty"`
	Size       int32                              `json:"size,omitempty"`
	Sort       []string                           `json:"sort,omitempty"`
	TotalCount int64                              `json:"totalCount,omitempty"`
}

type PageResponseV2MembersResponse struct {
	Contents   []MembersResponse `json:"contents,omitempty"`
	Page       int32                             `json:"page,omitempty"`
	Size       int32                             `json:"size,omitempty"`
	Sort       []string                          `json:"sort,omitempty"`
	TotalCount int64                             `json:"totalCount,omitempty"`
}

type PageResponseV2StatisticsLoggingResponse struct {
	Contents   []StatisticsLoggingResponse `json:"contents,omitempty"`
	Page       int32                                       `json:"page,omitempty"`
	Size       int32                                       `json:"size,omitempty"`
	Sort       []string                                    `json:"sort,omitempty"`
	TotalCount int64                                       `json:"totalCount,omitempty"`
}

type PageResponseV2TrailResponse struct {
	Contents   []TrailResponse `json:"contents,omitempty"`
	Page       int32                           `json:"page,omitempty"`
	Size       int32                           `json:"size,omitempty"`
	Sort       []string                        `json:"sort,omitempty"`
	TotalCount int64                           `json:"totalCount,omitempty"`
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

type StatisticsLoggingResponse struct {
	LogCount int32  `json:"logCount,omitempty"`
	LogDate  string `json:"logDate,omitempty"`
}

type TagRequest struct {
	TagKey   string `json:"tagKey"`
	TagValue string `json:"tagValue,omitempty"`
}

type TargetResourceResponse struct {
	LoggingItem      string `json:"loggingItem,omitempty"`
	ResourceTypeName string `json:"resourceTypeName,omitempty"`
}

type TrailControllerApiListUsersOpts struct {
	UserName optional.String
	Page     optional.Int32
	Size     optional.Int32
	Sort     optional.Interface
}

type TrailResponse struct {
	ProjectId                  string                                   `json:"projectId,omitempty"`
	IsLoggingTargetAllRegion   string                                   `json:"isLoggingTargetAllRegion,omitempty"`
	IsLoggingTargetAllResource string                                   `json:"isLoggingTargetAllResource,omitempty"`
	IsLoggingTargetAllUser     string                                   `json:"isLoggingTargetAllUser,omitempty"`
	IsTrailDeleted             string                                   `json:"isTrailDeleted,omitempty"`
	LastDigestFileCreateDt     time.Time                             `json:"lastDigestFileCreateDt,omitempty"`
	LoggingTargetProviders     []UserResponse           `json:"loggingTargetProviders,omitempty"`
	LoggingTargetRegions       []string                                 `json:"loggingTargetRegions,omitempty"`
	LoggingTargetResourceIds   []string                                 `json:"loggingTargetResourceIds,omitempty"`
	LoggingTargetUsers         []UserResponse           `json:"loggingTargetUsers,omitempty"`
	ObjectStorageFolderName    string                                   `json:"objectStorageFolderName,omitempty"`
	ObjectStorageName          string                                   `json:"objectStorageName,omitempty"`
	ObsBucketId                string                                   `json:"obsBucketId,omitempty"`
	ObsBucketName              string                                   `json:"obsBucketName,omitempty"`
	ProjectName                string                                   `json:"projectName,omitempty"`
	Region                     string                                   `json:"region,omitempty"`
	ServiceZoneId              string                                   `json:"serviceZoneId,omitempty"`
	TagCreateRequests          []TagRequest             `json:"tagCreateRequests,omitempty"`
	TargetLoggingResourceList  []TargetResourceResponse `json:"targetLoggingResourceList,omitempty"`
	TrailBatchEndDt            time.Time                             `json:"trailBatchEndDt,omitempty"`
	TrailBatchFirstStartDt     time.Time                             `json:"trailBatchFirstStartDt,omitempty"`
	TrailBatchLastSuccessDt    time.Time                             `json:"trailBatchLastSuccessDt,omitempty"`
	TrailBatchStartDt          time.Time                             `json:"trailBatchStartDt,omitempty"`
	TrailBatchState            string                                   `json:"trailBatchState,omitempty"`
	TrailId                    string                                   `json:"trailId,omitempty"`
	TrailName                  string                                   `json:"trailName,omitempty"`
	TrailSaveType              string                                   `json:"trailSaveType,omitempty"`
	TrailState                 string                                   `json:"trailState,omitempty"`
	TrailUpdateType            string                                   `json:"trailUpdateType,omitempty"`
	ValidationYn               string                                   `json:"validationYn,omitempty"`
	TrailDescription           string                                   `json:"trailDescription,omitempty"`
	CreatedBy                  string                                   `json:"createdBy,omitempty"`
	CreatedByName              string                                   `json:"createdByName,omitempty"`
	CreatedByEmail             string                                   `json:"createdByEmail,omitempty"`
	CreatedDt                  time.Time                             `json:"createdDt,omitempty"`
	ModifiedBy                 string                                   `json:"modifiedBy,omitempty"`
	ModifiedByName             string                                   `json:"modifiedByName,omitempty"`
	ModifiedByEmail            string                                   `json:"modifiedByEmail,omitempty"`
	ModifiedDt                 time.Time                             `json:"modifiedDt,omitempty"`
}

type TrailSearchCriteria struct {
	IsMy                     *bool    `json:"isMy,omitempty"`
	LoggingTargetRegions     []string `json:"loggingTargetRegions,omitempty"`
	LoggingTargetResourceIds []string `json:"loggingTargetResourceIds,omitempty"`
	Page                     int32    `json:"page,omitempty"`
	Size                     int32    `json:"size,omitempty"`
	Sort                     []string `json:"sort,omitempty"`
	StateList                []string `json:"stateList,omitempty"`
	TrailName                string   `json:"trailName,omitempty"`
}

type TrailStateRequest struct {
	ModifiedBy string `json:"modifiedBy,omitempty"`
}

type TrailStateResponse struct {
	TrailId    string `json:"trailId,omitempty"`
	TrailState string `json:"trailState,omitempty"`
}

type TrailUpdateResponse struct {
	IsLoggingTargetAllRegion   string                         `json:"isLoggingTargetAllRegion,omitempty"`
	IsLoggingTargetAllResource string                         `json:"isLoggingTargetAllResource,omitempty"`
	IsLoggingTargetAllUser     string                         `json:"isLoggingTargetAllUser,omitempty"`
	LoggingTargetRegions       []string                       `json:"loggingTargetRegions,omitempty"`
	LoggingTargetResourceIds   []string                       `json:"loggingTargetResourceIds,omitempty"`
	LoggingTargetUsers         []UserResponse `json:"loggingTargetUsers,omitempty"`
	TrailId                    string                         `json:"trailId,omitempty"`
	TrailSaveType              string                         `json:"trailSaveType,omitempty"`
	TrailUpdateType            string                         `json:"trailUpdateType,omitempty"`
	ValidationYn               string                         `json:"validationYn,omitempty"`
	TrailDescription           string                         `json:"trailDescription,omitempty"`
}

type TrailValidationRequest struct {
	Detail    *bool        `json:"detail,omitempty"`
	EndTime   time.Time `json:"endTime"`
	StartTime time.Time `json:"startTime"`
	TrailId   string       `json:"trailId"`
}

type TrailValidationResultResponse struct {
	TrailValidationId string `json:"trailValidationId,omitempty"`
}

type UpdateTrailRequest struct {
	IsLoggingTargetAllRegion   string                         `json:"isLoggingTargetAllRegion,omitempty"`
	IsLoggingTargetAllResource string                         `json:"isLoggingTargetAllResource,omitempty"`
	IsLoggingTargetAllUser     string                         `json:"isLoggingTargetAllUser,omitempty"`
	LoggingTargetProviders     []UserResponse `json:"loggingTargetProviders,omitempty"`
	LoggingTargetRegions       []string                       `json:"loggingTargetRegions,omitempty"`
	LoggingTargetResourceIds   []string                       `json:"loggingTargetResourceIds,omitempty"`
	LoggingTargetUsers         []UserResponse `json:"loggingTargetUsers,omitempty"`
	TrailSaveType              string                         `json:"trailSaveType,omitempty"`
	TrailUpdateType            string                         `json:"trailUpdateType"`
	ValidationYn               string                         `json:"validationYn,omitempty"`
	TrailDescription           string                         `json:"trailDescription,omitempty"`
}

type UserResponse struct {
	Email       string `json:"email,omitempty"`
	UserId      string `json:"userId,omitempty"`
	UserLoginId string `json:"userLoginId,omitempty"`
	UserName    string `json:"userName,omitempty"`
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
