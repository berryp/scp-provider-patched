package objectstorage

import (
	"github.com/antihax/optional"
	"time"
)

type APIKey struct {
	Key    string
	Prefix string
}

type AccessControlRule struct {
	RuleType  string `json:"ruleType,omitempty"`
	RuleValue string `json:"ruleValue,omitempty"`
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

type CreateBucketRequest struct {
	ObjectStorageBucketAccessControlEnabled  bool
	AccessControlRules                       []AccessControlRule
	ObjectStorageBucketFileEncryptionEnabled bool
	ObjectStorageBucketName                  string
	ObjectStorageBucketUserPurpose           string
	ObjectStorageBucketVersionEnabled        bool
	ObjectStorageId                          string
	ServiceZoneId                            string
	ProductNames                             []string
	Tags                                     map[string]interface{}
}

type ListResponseObjectStorageBucketListV4Response struct {
	Contents   []ObjectStorageBucketListV4Response `json:"contents,omitempty"`
	TotalCount int32                                                `json:"totalCount,omitempty"`
}

type ListResponseObjectStorageListV4Response struct {
	Contents   []ObjectStorageListV4Response `json:"contents,omitempty"`
	TotalCount int32                                          `json:"totalCount,omitempty"`
}

type ObjectStorageBucketAccessControlV4Request struct {
	ObjectStorageBucketAccessControlEnabled *bool                                `json:"objectStorageBucketAccessControlEnabled"`
	ObjectStorageBucketAccessControlRules   []AccessControlRule `json:"objectStorageBucketAccessControlRules,omitempty"`
}

type ObjectStorageBucketAccessControlV4Response struct {
	ProjectId                               string                               `json:"projectId,omitempty"`
	AccessControlRules                      []AccessControlRule `json:"accessControlRules,omitempty"`
	BlockId                                 string                               `json:"blockId,omitempty"`
	ObjectStorageBucketAccessControlEnabled *bool                                `json:"objectStorageBucketAccessControlEnabled,omitempty"`
	ObjectStorageBucketId                   string                               `json:"objectStorageBucketId,omitempty"`
	ObjectStorageBucketName                 string                               `json:"objectStorageBucketName,omitempty"`
	ServiceZoneId                           string                               `json:"serviceZoneId,omitempty"`
	CreatedBy                               string                               `json:"createdBy,omitempty"`
	CreatedDt                               time.Time                         `json:"createdDt,omitempty"`
	ModifiedBy                              string                               `json:"modifiedBy,omitempty"`
	ModifiedDt                              time.Time                         `json:"modifiedDt,omitempty"`
}

type ObjectStorageBucketAccessInfoV4Response struct {
	ProjectId                             string `json:"projectId,omitempty"`
	ObjectStorageBucketAccessKey          string `json:"objectStorageBucketAccessKey,omitempty"`
	ObjectStorageBucketPrivateEndpointUrl string `json:"objectStorageBucketPrivateEndpointUrl,omitempty"`
	ObjectStorageBucketPublicEndpointUrl  string `json:"objectStorageBucketPublicEndpointUrl,omitempty"`
	ObjectStorageBucketSecretKey          string `json:"objectStorageBucketSecretKey,omitempty"`
	ObjectStorageHost                     string `json:"objectStorageHost,omitempty"`
	ObjectStoragePort                     string `json:"objectStoragePort,omitempty"`
	ServiceZoneId                         string `json:"serviceZoneId,omitempty"`
}

type ObjectStorageBucketCreateV4Request struct {
	AccessControlRules                       []AccessControlRule `json:"accessControlRules,omitempty"`
	ObjectStorageBucketAccessControlEnabled  *bool                                `json:"objectStorageBucketAccessControlEnabled"`
	ObjectStorageBucketFileEncryptionEnabled *bool                                `json:"objectStorageBucketFileEncryptionEnabled"`
	ObjectStorageBucketName                  string                               `json:"objectStorageBucketName"`
	ObjectStorageBucketUserPurpose           string                               `json:"objectStorageBucketUserPurpose,omitempty"`
	ObjectStorageBucketVersionEnabled        *bool                                `json:"objectStorageBucketVersionEnabled"`
	ObjectStorageId                          string                               `json:"objectStorageId"`
	ProductNames                             []string                             `json:"productNames"`
	ServiceZoneId                            string                               `json:"serviceZoneId"`
	Tags                                     []TagRequest        `json:"tags,omitempty"`
}

type ObjectStorageBucketCredentialV4Response struct {
	ObjectStorageBucketAccessKey string `json:"objectStorageBucketAccessKey,omitempty"`
	ObjectStorageBucketSecretKey string `json:"objectStorageBucketSecretKey,omitempty"`
}

type ObjectStorageBucketDetailV4Response struct {
	ProjectId                                  string                               `json:"projectId,omitempty"`
	IsMultiAvailabilityZone                    *bool                                `json:"isMultiAvailabilityZone,omitempty"`
	IsSyncInProgress                           *bool                                `json:"isSyncInProgress,omitempty"`
	ObjectStorageBucketAccessControlEnabled    *bool                                `json:"objectStorageBucketAccessControlEnabled,omitempty"`
	ObjectStorageBucketAccessControlRules      []AccessControlRule `json:"objectStorageBucketAccessControlRules,omitempty"`
	ObjectStorageBucketDrEnabled               *bool                                `json:"objectStorageBucketDrEnabled,omitempty"`
	ObjectStorageBucketDrType                  string                               `json:"objectStorageBucketDrType,omitempty"`
	ObjectStorageBucketFileEncryptionAlgorithm string                               `json:"objectStorageBucketFileEncryptionAlgorithm,omitempty"`
	ObjectStorageBucketFileEncryptionEnabled   *bool                                `json:"objectStorageBucketFileEncryptionEnabled,omitempty"`
	ObjectStorageBucketFileEncryptionType      string                               `json:"objectStorageBucketFileEncryptionType,omitempty"`
	ObjectStorageBucketId                      string                               `json:"objectStorageBucketId,omitempty"`
	ObjectStorageBucketName                    string                               `json:"objectStorageBucketName,omitempty"`
	ObjectStorageBucketObjectUploadEnabled     *bool                                `json:"objectStorageBucketObjectUploadEnabled,omitempty"`
	ObjectStorageBucketPrivateEndpointUrl      string                               `json:"objectStorageBucketPrivateEndpointUrl,omitempty"`
	ObjectStorageBucketPublicEndpointUrl       string                               `json:"objectStorageBucketPublicEndpointUrl,omitempty"`
	ObjectStorageBucketPurpose                 string                               `json:"objectStorageBucketPurpose,omitempty"`
	ObjectStorageBucketState                   string                               `json:"objectStorageBucketState,omitempty"`
	ObjectStorageBucketUsage                   int64                                `json:"objectStorageBucketUsage,omitempty"`
	ObjectStorageBucketUserPurpose             string                               `json:"objectStorageBucketUserPurpose,omitempty"`
	ObjectStorageBucketVersionEnabled          *bool                                `json:"objectStorageBucketVersionEnabled,omitempty"`
	ObjectStorageDeviceUserId                  string                               `json:"objectStorageDeviceUserId,omitempty"`
	ObjectStorageId                            string                               `json:"objectStorageId,omitempty"`
	ObjectStorageName                          string                               `json:"objectStorageName,omitempty"`
	ObjectStorageQuotaId                       string                               `json:"objectStorageQuotaId,omitempty"`
	ObjectStorageQuotaName                     string                               `json:"objectStorageQuotaName,omitempty"`
	ObjectStorageSystemBucketEnabled           *bool                                `json:"objectStorageSystemBucketEnabled,omitempty"`
	ObjectStorageTenantName                    string                               `json:"objectStorageTenantName,omitempty"`
	ServiceZoneId                              string                               `json:"serviceZoneId,omitempty"`
	SyncObjectStorageBucketId                  string                               `json:"syncObjectStorageBucketId,omitempty"`
	SyncObjectStorageBucketName                string                               `json:"syncObjectStorageBucketName,omitempty"`
	SyncObjectStorageBucketServiceZoneId       string                               `json:"syncObjectStorageBucketServiceZoneId,omitempty"`
	CreatedBy                                  string                               `json:"createdBy,omitempty"`
	CreatedDt                                  time.Time                         `json:"createdDt,omitempty"`
	ModifiedBy                                 string                               `json:"modifiedBy,omitempty"`
	ModifiedDt                                 time.Time                         `json:"modifiedDt,omitempty"`
}

type ObjectStorageBucketDrUpdateV4Request struct {
	ObjectStorageBucketDrEnabled *bool  `json:"objectStorageBucketDrEnabled"`
	SyncObjectStorageBucketId    string `json:"syncObjectStorageBucketId,omitempty"`
}

type ObjectStorageBucketEndpointV4Response struct {
	ObjectStorageBucketPrivateEndpointUrl string `json:"objectStorageBucketPrivateEndpointUrl,omitempty"`
	ObjectStorageBucketPublicEndpointUrl  string `json:"objectStorageBucketPublicEndpointUrl,omitempty"`
}

type ObjectStorageBucketFileEncryptionUpdateV4Request struct {
	ObjectStorageBucketFileEncryptionEnabled *bool `json:"objectStorageBucketFileEncryptionEnabled"`
}

type ObjectStorageBucketListV4Response struct {
	ProjectId                               string       `json:"projectId,omitempty"`
	ObjectStorageBucketAccessControlEnabled *bool        `json:"objectStorageBucketAccessControlEnabled,omitempty"`
	ObjectStorageBucketDrEnabled            *bool        `json:"objectStorageBucketDrEnabled,omitempty"`
	ObjectStorageBucketDrType               string       `json:"objectStorageBucketDrType,omitempty"`
	ObjectStorageBucketId                   string       `json:"objectStorageBucketId,omitempty"`
	ObjectStorageBucketName                 string       `json:"objectStorageBucketName,omitempty"`
	ObjectStorageBucketPurpose              string       `json:"objectStorageBucketPurpose,omitempty"`
	ObjectStorageBucketState                string       `json:"objectStorageBucketState,omitempty"`
	ObjectStorageBucketUserPurpose          string       `json:"objectStorageBucketUserPurpose,omitempty"`
	ObjectStorageBucketVersionEnabled       *bool        `json:"objectStorageBucketVersionEnabled,omitempty"`
	ObjectStorageId                         string       `json:"objectStorageId,omitempty"`
	ObjectStorageName                       string       `json:"objectStorageName,omitempty"`
	ObjectStorageQuotaId                    string       `json:"objectStorageQuotaId,omitempty"`
	ObjectStorageQuotaName                  string       `json:"objectStorageQuotaName,omitempty"`
	ObjectStorageSystemBucketEnabled        *bool        `json:"objectStorageSystemBucketEnabled,omitempty"`
	ObjectStorageTenantName                 string       `json:"objectStorageTenantName,omitempty"`
	ServiceZoneId                           string       `json:"serviceZoneId,omitempty"`
	CreatedBy                               string       `json:"createdBy,omitempty"`
	CreatedDt                               time.Time `json:"createdDt,omitempty"`
	ModifiedBy                              string       `json:"modifiedBy,omitempty"`
	ModifiedDt                              time.Time `json:"modifiedDt,omitempty"`
}

type ObjectStorageBucketObjectListV4ResponseObjectStorageObjectListResponse struct {
	Contents   []ObjectStorageObjectListResponse `json:"contents,omitempty"`
	Marker     string                                             `json:"marker,omitempty"`
	TotalCount int32                                              `json:"totalCount,omitempty"`
	Truncated  *bool                                              `json:"truncated,omitempty"`
}

type ObjectStorageBucketPreSignedUrlV4Response struct {
	ObjectStoragePresignedUrl string `json:"objectStoragePresignedUrl,omitempty"`
}

type ObjectStorageBucketV4ControllerApiListObjectStorageBucketsOpts struct {
	EndModifiedDt                    optional.Time
	ObjectStorageBucketIds           optional.Interface
	ObjectStorageBucketName          optional.String
	ObjectStorageBucketPurposes      optional.Interface
	ObjectStorageBucketState         optional.String
	ObjectStorageBucketStates        optional.Interface
	ObjectStorageBucketUserPurpose   optional.String
	ObjectStorageId                  optional.String
	ObjectStorageQuotaId             optional.String
	ObjectStorageSystemBucketEnabled optional.Bool
	ServiceZoneId                    optional.String
	StartModifiedDt                  optional.Time
	CreatedBy                        optional.String
	Page                             optional.Int32
	Size                             optional.Int32
	Sort                             optional.Interface
}

type ObjectStorageBucketVersionUpdateV4Request struct {
	ObjectStorageBucketVersionEnabled *bool `json:"objectStorageBucketVersionEnabled"`
}

type ObjectStorageIpsV4ControllerApiListObjectStorageBucketAccessControlOpts struct {
	RuleType optional.String
}

type ObjectStorageListV4Response struct {
	ProjectId                string       `json:"projectId,omitempty"`
	BlockId                  string       `json:"blockId,omitempty"`
	IsMultiAvailabilityZone  *bool        `json:"isMultiAvailabilityZone,omitempty"`
	ObjectStorageId          string       `json:"objectStorageId,omitempty"`
	ObjectStorageName        string       `json:"objectStorageName,omitempty"`
	ServiceZoneId            string       `json:"serviceZoneId,omitempty"`
	ObjectStorageDescription string       `json:"objectStorageDescription,omitempty"`
	CreatedBy                string       `json:"createdBy,omitempty"`
	CreatedDt                time.Time `json:"createdDt,omitempty"`
	ModifiedBy               string       `json:"modifiedBy,omitempty"`
	ModifiedDt               time.Time `json:"modifiedDt,omitempty"`
}

type ObjectStorageObjectListResponse struct {
	ObjectStorageBucketObjectFileSize string       `json:"objectStorageBucketObjectFileSize,omitempty"`
	ObjectStorageBucketObjectName     string       `json:"objectStorageBucketObjectName,omitempty"`
	ObjectStorageBucketObjectType     string       `json:"objectStorageBucketObjectType,omitempty"`
	ModifiedDt                        time.Time `json:"modifiedDt,omitempty"`
}

type ObjectStorageObjectMetadataResponse struct {
	ProjectId                             string       `json:"projectId,omitempty"`
	BlockId                               string       `json:"blockId,omitempty"`
	ObjectStorageBucketObjectName         string       `json:"objectStorageBucketObjectName,omitempty"`
	ObjectStorageBucketObjectPath         string       `json:"objectStorageBucketObjectPath,omitempty"`
	ObjectStorageBucketObjectSize         string       `json:"objectStorageBucketObjectSize,omitempty"`
	ObjectStorageBucketObjectTag          string       `json:"objectStorageBucketObjectTag,omitempty"`
	ObjectStorageBucketObjectType         string       `json:"objectStorageBucketObjectType,omitempty"`
	ObjectStorageBucketPrivateEndpointUrl string       `json:"objectStorageBucketPrivateEndpointUrl,omitempty"`
	ObjectStorageBucketPublicEndpointUrl  string       `json:"objectStorageBucketPublicEndpointUrl,omitempty"`
	ServiceZoneId                         string       `json:"serviceZoneId,omitempty"`
	CreatedBy                             string       `json:"createdBy,omitempty"`
	CreatedDt                             time.Time `json:"createdDt,omitempty"`
	ModifiedBy                            string       `json:"modifiedBy,omitempty"`
	ModifiedDt                            time.Time `json:"modifiedDt,omitempty"`
}

type ObjectStorageObjectV4ControllerApiListObjectStorageBucketObjectsOpts struct {
	Marker                        optional.String
	ObjectStorageBucketObjectName optional.String
	ObjectStorageBucketObjectPath optional.String
	ObjectStorageBucketObjectType optional.String
	Size                          optional.Int32
}

type ObjectStorageObjectV4ControllerApiSearchObjectStorageBucketPresignedUrlOpts struct {
	ObjectStorageBucketObjectVersionId optional.String
}

type ObjectStorageV4ControllerApiListObjectStorageOpts struct {
	IsMultiAvailabilityZone optional.Bool
	ObjectStorageName       optional.String
	Page                    optional.Int32
	Size                    optional.Int32
	Sort                    optional.Interface
}

type ReadBucketListRequest struct {
	EndModifiedDt                    optional.Time
	ObjectStorageBucketIds           optional.Interface
	ObjectStorageBucketName          optional.String
	ObjectStorageBucketPurposes      optional.Interface
	ObjectStorageBucketState         optional.String
	ObjectStorageBucketStates        optional.Interface
	ObjectStorageBucketUserPurpose   optional.String
	ObjectStorageId                  optional.String
	ObjectStorageQuotaId             optional.String
	ObjectStorageSystemBucketEnabled optional.Bool
	ServiceZoneId                    optional.String
	StartModifiedDt                  optional.Time
	CreatedBy                        optional.String
	Page                             optional.Int32
	Size                             optional.Int32
	Sort                             optional.Interface
}

type ReadObjectStorageListRequest struct {
	IsMultiAvailabilityZone optional.Bool
	ObjectStorageName       optional.String
	Page                    optional.Int32
	Size                    optional.Int32
	Sort                    optional.Interface
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
