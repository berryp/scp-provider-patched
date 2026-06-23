package kubernetesapps

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

type CheckResponse struct {
	Result *bool `json:"result,omitempty"`
}

type ClusterValidationResponse struct {
	ErrorDetail string `json:"errorDetail,omitempty"`
	ErrorType   string `json:"errorType,omitempty"`
	Result      *bool  `json:"result,omitempty"`
}

type CommonCodeResponse struct {
	ClassificationCode string `json:"classificationCode,omitempty"`
	Code               string `json:"code,omitempty"`
	Name               string `json:"name,omitempty"`
	Description        string `json:"description,omitempty"`
}

type ImageApiListImagesV1Opts struct {
	Category         optional.String
	ImageId          optional.String
	ImageName        optional.String
	IsCarepack       optional.String
	IsNew            optional.String
	IsRecommended    optional.String
	PricePolicy      optional.String
	ProductGroupName optional.String
	Provider         optional.String
	Page             optional.Int32
	Size             optional.Int32
	Sort             optional.String
}

type ImagesResponse struct {
	BaseImage        string            `json:"baseImage,omitempty"`
	Category         string            `json:"category,omitempty"`
	IconFile         string            `json:"iconFile,omitempty"`
	IconFileName     string            `json:"iconFileName,omitempty"`
	ImageAttr        map[string]string `json:"imageAttr,omitempty"`
	ImageId          string            `json:"imageId,omitempty"`
	ImageName        string            `json:"imageName,omitempty"`
	PoolId           string            `json:"poolId,omitempty"`
	ProductGroupAttr map[string]string `json:"productGroupAttr,omitempty"`
	ProductGroupId   string            `json:"productGroupId,omitempty"`
	ProductGroupName string            `json:"productGroupName,omitempty"`
	ZoneId           string            `json:"zoneId,omitempty"`
	Description      string            `json:"description,omitempty"`
	CreatedBy        string            `json:"createdBy,omitempty"`
	CreatedDt        time.Time      `json:"createdDt,omitempty"`
	ModifiedBy       string            `json:"modifiedBy,omitempty"`
	ModifiedDt       time.Time      `json:"modifiedDt,omitempty"`
}

type K8sAppsApiListK8sAppsV1Opts struct {
	ClusterName   optional.String
	NamespaceName optional.String
	ReleaseName   optional.String
	CreatedBy     optional.String
	Page          optional.Int32
	Size          optional.Int32
	Sort          optional.String
}

type K8sAppsResponse struct {
	Application    string                 `json:"application,omitempty"`
	ClusterId      string                 `json:"clusterId,omitempty"`
	ClusterName    string                 `json:"clusterName,omitempty"`
	NamespaceName  string                 `json:"namespaceName,omitempty"`
	ReleaseId      string                 `json:"releaseId,omitempty"`
	ReleaseName    string                 `json:"releaseName,omitempty"`
	ReleaseState   string                 `json:"releaseState,omitempty"`
	ResourceDetail map[string]interface{} `json:"resourceDetail,omitempty"`
	CreatedBy      string                 `json:"createdBy,omitempty"`
	CreatedDt      time.Time           `json:"createdDt,omitempty"`
	ModifiedBy     string                 `json:"modifiedBy,omitempty"`
	ModifiedDt     time.Time           `json:"modifiedDt,omitempty"`
}

type K8sAppsesResponse struct {
	ChartName     string       `json:"chartName,omitempty"`
	ClusterId     string       `json:"clusterId,omitempty"`
	ClusterName   string       `json:"clusterName,omitempty"`
	NamespaceName string       `json:"namespaceName,omitempty"`
	ReleaseId     string       `json:"releaseId,omitempty"`
	ReleaseName   string       `json:"releaseName,omitempty"`
	ReleaseState  string       `json:"releaseState,omitempty"`
	Version       string       `json:"version,omitempty"`
	CreatedBy     string       `json:"createdBy,omitempty"`
	CreatedDt     time.Time `json:"createdDt,omitempty"`
	ModifiedBy    string       `json:"modifiedBy,omitempty"`
	ModifiedDt    time.Time `json:"modifiedDt,omitempty"`
}

type ListStandardImageRequest struct {
	Category         string
	ImageId          string
	ImageName        string
	IsCarepack       string
	IsNew            string
	IsRecommended    string
	PricePolicy      string
	ProductGroupName string
	Page             int32
	Size             int32
	Sort             string
}

type PageResponseImagesResponse struct {
	Contents   []ImagesResponse `json:"contents,omitempty"`
	TotalCount int64                              `json:"totalCount,omitempty"`
}

type PageResponseK8sAppsesResponse struct {
	Contents   []K8sAppsesResponse `json:"contents,omitempty"`
	TotalCount int64                                 `json:"totalCount,omitempty"`
}

type PageResponseSyscodesResponse struct {
	Contents   []SyscodesResponse `json:"contents,omitempty"`
	TotalCount int64                                `json:"totalCount,omitempty"`
}

type ReleaseCreateRequest struct {
	ProjectId        string                         `json:"projectId"`
	AdditionalParams map[string]interface{}         `json:"additionalParams,omitempty"`
	ClusterId        string                         `json:"clusterId"`
	ImageId          string                         `json:"imageId"`
	NamespaceName    string                         `json:"namespaceName"`
	ProductGroupId   string                         `json:"productGroupId"`
	ReleaseName      string                         `json:"releaseName"`
	Tags             []TagRequest `json:"tags,omitempty"`
}

type SyscodesResponse struct {
	Id   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type TagRequest struct {
	TagKey   string `json:"tagKey"`
	TagValue string `json:"tagValue,omitempty"`
}

type ValidationResponse struct {
	Result *bool `json:"result,omitempty"`
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
