package certificate

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

type CertControlApiCreateCertificateOpts struct {
	XCmpTimezone optional.String
}

type CertControlApiCreateSelfSignedCertificateOpts struct {
	XCmpTimezone optional.String
}

type CertControlApiListCertificatesOpts struct {
	CertificateName optional.String
	CommonName      optional.String
	IsMine          optional.Bool
	Mine            optional.Bool
	Page            optional.Int32
	Size            optional.Int32
	Sort            optional.Interface
}

type CertControlApiListRecipientsOpts struct {
	CompanyName optional.String
	Email       optional.String
	UserName    optional.String
	Page        optional.Int32
	Size        optional.Int32
	Sort        optional.Interface
}

type CertControlApiValidateCertificationPrdOpts struct {
	CertificateChain       optional.String
	IsNeedCertificateChain optional.String
}

type CertificateListResponse struct {
	CertificateExpirationDate time.Time `json:"certificateExpirationDate,omitempty"`
	CertificateId             string       `json:"certificateId,omitempty"`
	CertificateName           string       `json:"certificateName,omitempty"`
	CertificateStartDate      time.Time `json:"certificateStartDate,omitempty"`
	CertificateState          string       `json:"certificateState,omitempty"`
	CommonName                string       `json:"commonName,omitempty"`
	PurposeType               string       `json:"purposeType,omitempty"`
	CreatedBy                 string       `json:"createdBy,omitempty"`
	CreatedDt                 time.Time `json:"createdDt,omitempty"`
	ModifiedBy                string       `json:"modifiedBy,omitempty"`
	ModifiedDt                time.Time `json:"modifiedDt,omitempty"`
}

type CertificateResponse struct {
	ProjectId                 string       `json:"projectId,omitempty"`
	CertificateChain          string       `json:"certificateChain,omitempty"`
	CertificateExpirationDate time.Time `json:"certificateExpirationDate,omitempty"`
	CertificateId             string       `json:"certificateId,omitempty"`
	CertificateName           string       `json:"certificateName,omitempty"`
	CertificateStartDate      time.Time `json:"certificateStartDate,omitempty"`
	CertificateState          string       `json:"certificateState,omitempty"`
	CertificateType           string       `json:"certificateType,omitempty"`
	CertificateVersion        string       `json:"certificateVersion,omitempty"`
	CommonName                string       `json:"commonName,omitempty"`
	KeyBitSize                int32        `json:"keyBitSize,omitempty"`
	OrganizationName          string       `json:"organizationName,omitempty"`
	PrivateKey                string       `json:"privateKey,omitempty"`
	PublicCertificate         string       `json:"publicCertificate,omitempty"`
	PurposeType               string       `json:"purposeType,omitempty"`
	UsedResourceCount         int32        `json:"usedResourceCount,omitempty"`
	CreatedBy                 string       `json:"createdBy,omitempty"`
	CreatedByName             string       `json:"createdByName,omitempty"`
	CreatedByEmail            string       `json:"createdByEmail,omitempty"`
	CreatedDt                 time.Time `json:"createdDt,omitempty"`
	ModifiedBy                string       `json:"modifiedBy,omitempty"`
	ModifiedDt                time.Time `json:"modifiedDt,omitempty"`
}

type CheckResponse struct {
	Result *bool `json:"result,omitempty"`
}

type CommonCodeResponse struct {
	ClassificationCode string `json:"classificationCode,omitempty"`
	Code               string `json:"code,omitempty"`
	Name               string `json:"name,omitempty"`
	Description        string `json:"description,omitempty"`
}

type CreateDevCertificateRequest struct {
	CertificateExpirationDate string                      `json:"certificateExpirationDate"`
	CertificateName           string                      `json:"certificateName"`
	CertificateStartDate      string                      `json:"certificateStartDate"`
	CommonName                string                      `json:"commonName"`
	OrganizationName          string                      `json:"organizationName"`
	Recipients                []Recipient  `json:"recipients,omitempty"`
	Tags                      []TagRequest `json:"tags,omitempty"`
}

type CreatePrdCertificateRequest struct {
	CertificateChain  string                      `json:"certificateChain,omitempty"`
	CertificateName   string                      `json:"certificateName"`
	PrivateKey        string                      `json:"privateKey"`
	PublicCertificate string                      `json:"publicCertificate"`
	Recipients        []Recipient  `json:"recipients,omitempty"`
	Tags              []TagRequest `json:"tags,omitempty"`
}

type PageResponseV2CertificateListResponse struct {
	Contents   []CertificateListResponse `json:"contents,omitempty"`
	Page       int32                                    `json:"page,omitempty"`
	Size       int32                                    `json:"size,omitempty"`
	Sort       []string                                 `json:"sort,omitempty"`
	TotalCount int64                                    `json:"totalCount,omitempty"`
}

type PageResponseV2RecipientResponse struct {
	Contents   []RecipientResponse `json:"contents,omitempty"`
	Page       int32                              `json:"page,omitempty"`
	Size       int32                              `json:"size,omitempty"`
	Sort       []string                           `json:"sort,omitempty"`
	TotalCount int64                              `json:"totalCount,omitempty"`
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

type Recipient struct {
	Email    string `json:"email,omitempty"`
	UserId   string `json:"userId"`
	UserName string `json:"userName,omitempty"`
}

type RecipientResponse struct {
	CompanyName string `json:"companyName"`
	Email       string `json:"email"`
	UserId      string `json:"userId"`
	UserName    string `json:"userName"`
}

type RecipientsRequest struct {
	Recipients []Recipient `json:"recipients"`
}

type StatusResponse struct {
	Status string `json:"status,omitempty"`
}

type TagRequest struct {
	TagKey   string `json:"tagKey"`
	TagValue string `json:"tagValue,omitempty"`
}

type ValidateNameRequest struct {
	CertificateName string `json:"certificateName"`
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
