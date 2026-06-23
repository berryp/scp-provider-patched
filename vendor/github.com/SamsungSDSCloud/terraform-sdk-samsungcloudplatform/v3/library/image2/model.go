package image2

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

type CloneCustomImageResponse struct {
	AcceptedDt         time.Time `json:"acceptedDt,omitempty"`
	ApprovalState      string       `json:"approvalState,omitempty"`
	ApproverObjectId   string       `json:"approverObjectId,omitempty"`
	ApproverProjectId  string       `json:"approverProjectId,omitempty"`
	Contents           string       `json:"contents,omitempty"`
	RequestedDt        time.Time `json:"requestedDt,omitempty"`
	RequesterObjectId  string       `json:"requesterObjectId,omitempty"`
	RequesterProjectId string       `json:"requesterProjectId,omitempty"`
	RequesterUserId    string       `json:"requesterUserId,omitempty"`
}

type CustomImageApprovalResponse struct {
	ImageId string `json:"imageId,omitempty"`
	Success *bool  `json:"success,omitempty"`
}

type CustomImageCloneApproveRequest struct {
	AvailabilityZoneName string                 `json:"availabilityZoneName,omitempty"`
	ImageName            string                 `json:"imageName"`
	Tags                 []TagRequest `json:"tags,omitempty"`
	Description          string                 `json:"description,omitempty"`
}

type CustomImageCloneRequest struct {
	ApproverProjectId string `json:"approverProjectId"`
	ImageId           string `json:"imageId"`
	Description       string `json:"description,omitempty"`
}

type CustomImageCreateRequest struct {
	ImageName        string                 `json:"imageName"`
	Tags             []TagRequest `json:"tags,omitempty"`
	VirtualServerId  string                 `json:"virtualServerId"`
	ImageDescription string                 `json:"imageDescription,omitempty"`
}

type CustomImageResponse struct {
	ProjectId             string                                `json:"projectId,omitempty"`
	AvailabilityZoneName  string                                `json:"availabilityZoneName,omitempty"`
	BaseImage             string                                `json:"baseImage,omitempty"`
	BlockId               string                                `json:"blockId,omitempty"`
	Category              string                                `json:"category,omitempty"`
	DefaultDiskSize       int32                                 `json:"defaultDiskSize,omitempty"`
	DiskSize              int32                                 `json:"diskSize,omitempty"`
	Disks                 []ImageDiskSnapshotResponse `json:"disks,omitempty"`
	Icon                  *ImageIconResponse          `json:"icon,omitempty"`
	ImageId               string                                `json:"imageId,omitempty"`
	ImageName             string                                `json:"imageName,omitempty"`
	ImageState            string                                `json:"imageState,omitempty"`
	ImageType             string                                `json:"imageType,omitempty"`
	OriginImageId         string                                `json:"originImageId,omitempty"`
	OriginImageName       string                                `json:"originImageName,omitempty"`
	OriginVirtualServerId string                                `json:"originVirtualServerId,omitempty"`
	OsType                string                                `json:"osType,omitempty"`
	ProductGroupId        string                                `json:"productGroupId,omitempty"`
	Products              []ImageProductAssocResponse `json:"products,omitempty"`
	Properties            map[string]string                     `json:"properties,omitempty"`
	ServiceZoneId         string                                `json:"serviceZoneId,omitempty"`
	ImageDescription      string                                `json:"imageDescription,omitempty"`
	CreatedBy             string                                `json:"createdBy,omitempty"`
	CreatedDt             time.Time                          `json:"createdDt,omitempty"`
	ModifiedBy            string                                `json:"modifiedBy,omitempty"`
	ModifiedDt            time.Time                          `json:"modifiedDt,omitempty"`
}

type CustomImageUpdateRequest struct {
	ImageDescription string `json:"imageDescription,omitempty"`
}

type CustomImageV2ApiListCustomImagesOpts struct {
	ImageName        optional.String
	ImageState       optional.String
	IsSkeCustomImage optional.Bool
	OriginImageName  optional.String
	ProductGroupId   optional.String
	ServiceZoneId    optional.String
	ServicedFor      optional.String
	ServicedGroupFor optional.String
	CreatedBy        optional.String
	Page             optional.Int32
	Size             optional.Int32
	Sort             optional.Interface
}

type DetailImageTypeResponse struct {
	ImageId   string `json:"imageId,omitempty"`
	ImageType string `json:"imageType,omitempty"`
}

type ImageDiskSnapshotResponse struct {
	BootEnabled    *bool        `json:"bootEnabled,omitempty"`
	DeviceNode     string       `json:"deviceNode,omitempty"`
	DiskSize       int32        `json:"diskSize,omitempty"`
	EncryptEnabled *bool        `json:"encryptEnabled,omitempty"`
	ImageId        string       `json:"imageId,omitempty"`
	ProductId      string       `json:"productId,omitempty"`
	Seq            int32        `json:"seq,omitempty"`
	CreatedBy      string       `json:"createdBy,omitempty"`
	CreatedDt      time.Time `json:"createdDt,omitempty"`
	ModifiedBy     string       `json:"modifiedBy,omitempty"`
	ModifiedDt     time.Time `json:"modifiedDt,omitempty"`
}

type ImageIconResponse struct {
	IconFile     string `json:"iconFile,omitempty"`
	IconFileName string `json:"iconFileName,omitempty"`
	ImageId      string `json:"imageId,omitempty"`
	Description  string `json:"description,omitempty"`
}

type ImageProductAssocResponse struct {
	ImageId      string       `json:"imageId,omitempty"`
	ProductId    string       `json:"productId,omitempty"`
	ProductName  string       `json:"productName,omitempty"`
	ProductType  string       `json:"productType,omitempty"`
	ProductValue string       `json:"productValue,omitempty"`
	Seq          int32        `json:"seq,omitempty"`
	CreatedDt    time.Time `json:"createdDt,omitempty"`
}

type ListResponseCloneCustomImageResponse struct {
	Contents   []CloneCustomImageResponse `json:"contents,omitempty"`
	TotalCount int32                                `json:"totalCount,omitempty"`
}

type ListResponseCustomImageResponse struct {
	Contents   []CustomImageResponse `json:"contents,omitempty"`
	TotalCount int32                           `json:"totalCount,omitempty"`
}

type ListResponseMigrationImageResponse struct {
	Contents   []MigrationImageResponse `json:"contents,omitempty"`
	TotalCount int32                              `json:"totalCount,omitempty"`
}

type ListResponseStandardImageResponse struct {
	Contents   []StandardImageResponse `json:"contents,omitempty"`
	TotalCount int32                             `json:"totalCount,omitempty"`
}

type MigrationImageCreateRequest struct {
	AccessKey            string                                            `json:"accessKey"`
	AvailabilityZoneName string                                            `json:"availabilityZoneName,omitempty"`
	ImageName            string                                            `json:"imageName"`
	OriginalImageId      string                                            `json:"originalImageId"`
	OsAdminCredential    *VirtualServerCreateOsCredentialRequest `json:"osAdminCredential"`
	OvaUrl               string                                            `json:"ovaUrl"`
	SecretKey            string                                            `json:"secretKey"`
	ServiceZoneId        string                                            `json:"serviceZoneId"`
	Tags                 []TagRequest                            `json:"tags,omitempty"`
	ImageDescription     string                                            `json:"imageDescription,omitempty"`
}

type MigrationImageResponse struct {
	ProjectId             string                                `json:"projectId,omitempty"`
	AvailabilityZoneName  string                                `json:"availabilityZoneName,omitempty"`
	BaseImage             string                                `json:"baseImage,omitempty"`
	BlockId               string                                `json:"blockId,omitempty"`
	Category              string                                `json:"category,omitempty"`
	DefaultDiskSize       int32                                 `json:"defaultDiskSize,omitempty"`
	DiskSize              int32                                 `json:"diskSize,omitempty"`
	Disks                 []ImageDiskSnapshotResponse `json:"disks,omitempty"`
	Icon                  *ImageIconResponse          `json:"icon,omitempty"`
	ImageId               string                                `json:"imageId,omitempty"`
	ImageName             string                                `json:"imageName,omitempty"`
	ImageState            string                                `json:"imageState,omitempty"`
	ImageType             string                                `json:"imageType,omitempty"`
	OriginImageId         string                                `json:"originImageId,omitempty"`
	OriginImageName       string                                `json:"originImageName,omitempty"`
	OriginVirtualServerId string                                `json:"originVirtualServerId,omitempty"`
	OsType                string                                `json:"osType,omitempty"`
	ProductGroupId        string                                `json:"productGroupId,omitempty"`
	Products              []ImageProductAssocResponse `json:"products,omitempty"`
	Properties            map[string]string                     `json:"properties,omitempty"`
	ServiceZoneId         string                                `json:"serviceZoneId,omitempty"`
	ImageDescription      string                                `json:"imageDescription,omitempty"`
	CreatedBy             string                                `json:"createdBy,omitempty"`
	CreatedDt             time.Time                          `json:"createdDt,omitempty"`
	ModifiedBy            string                                `json:"modifiedBy,omitempty"`
	ModifiedDt            time.Time                          `json:"modifiedDt,omitempty"`
}

type MigrationImageUpdateRequest struct {
	ImageDescription string `json:"imageDescription,omitempty"`
}

type MigrationImageV2ApiListMigrationImagesOpts struct {
	ImageName        optional.String
	ImageState       optional.String
	IsSkeCustomImage optional.Bool
	OriginImageName  optional.String
	ProductGroupId   optional.String
	ServiceZoneId    optional.String
	ServicedFor      optional.String
	ServicedGroupFor optional.String
	CreatedBy        optional.String
	Page             optional.Int32
	Size             optional.Int32
	Sort             optional.Interface
}

type StandardImageResponse struct {
	ProjectId        string                                `json:"projectId,omitempty"`
	BaseImage        string                                `json:"baseImage,omitempty"`
	BlockId          string                                `json:"blockId,omitempty"`
	Category         string                                `json:"category,omitempty"`
	Icon             *ImageIconResponse          `json:"icon,omitempty"`
	ImageId          string                                `json:"imageId,omitempty"`
	ImageName        string                                `json:"imageName,omitempty"`
	ImageState       string                                `json:"imageState,omitempty"`
	ImageType        string                                `json:"imageType,omitempty"`
	OsType           string                                `json:"osType,omitempty"`
	ProductGroupId   string                                `json:"productGroupId,omitempty"`
	Products         []ImageProductAssocResponse `json:"products,omitempty"`
	Properties       map[string]string                     `json:"properties,omitempty"`
	ServiceZoneId    string                                `json:"serviceZoneId,omitempty"`
	ImageDescription string                                `json:"imageDescription,omitempty"`
	CreatedBy        string                                `json:"createdBy,omitempty"`
	CreatedDt        time.Time                          `json:"createdDt,omitempty"`
	ModifiedBy       string                                `json:"modifiedBy,omitempty"`
	ModifiedDt       time.Time                          `json:"modifiedDt,omitempty"`
}

type StandardImageV2ApiListStandardImagesOpts struct {
	ImageName        optional.String
	ImageState       optional.String
	ImageType        optional.String
	ProductGroupId   optional.String
	ServicedFor      optional.String
	ServicedGroupFor optional.String
	Page             optional.Int32
	Size             optional.Int32
	Sort             optional.Interface
}

type TagRequest struct {
	TagKey   string `json:"tagKey"`
	TagValue string `json:"tagValue,omitempty"`
}

type VirtualServerCreateOsCredentialRequest struct {
	OsUserId       string `json:"osUserId,omitempty"`
	OsUserPassword string `json:"osUserPassword"`
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
