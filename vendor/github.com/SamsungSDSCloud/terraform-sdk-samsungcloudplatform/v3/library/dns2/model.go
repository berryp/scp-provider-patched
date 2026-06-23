package dns2

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

type ChangeAutoExtensionRequest struct {
	AutoExtension *bool `json:"autoExtension"`
}

type ChangeDnsDomainDescriptionRequest struct {
	Description string `json:"description,omitempty"`
}

type ChangeDnsRecordOpenApiV3Request struct {
	DnsRecordMapping []DnsRecordMappingRequest `json:"dnsRecordMapping"`
	Ttl              int32                             `json:"ttl"`
}

type ChangeExpiredDateRequest struct {
	ExpiredDate string `json:"expiredDate"`
}

type CreateDnsDomainOpenApiV3Request struct {
	CountryType       string               `json:"countryType"`
	DnsDomainName     string               `json:"dnsDomainName"`
	DnsEnvUsage       string               `json:"dnsEnvUsage"`
	EnCompanyName     string               `json:"enCompanyName"`
	EnDetailAddress   string               `json:"enDetailAddress"`
	ExpiredDate       string               `json:"expiredDate,omitempty"`
	FirstEnAddress    string               `json:"firstEnAddress"`
	FirstKoAddress    string               `json:"firstKoAddress"`
	KoCompanyName     string               `json:"koCompanyName"`
	KoDetailAddress   string               `json:"koDetailAddress"`
	PostalCode        string               `json:"postalCode"`
	RegisteredByEmail string               `json:"registeredByEmail"`
	RegisteredByTelno string               `json:"registeredByTelno"`
	SecondEnAddress   string               `json:"secondEnAddress,omitempty"`
	SecondKoAddress   string               `json:"secondKoAddress,omitempty"`
	Tags              []TagRequest `json:"tags,omitempty"`
	DnsDescription    string               `json:"dnsDescription,omitempty"`
}

type CreateDnsRecordOpenApiV3Request struct {
	DnsRecordMapping []DnsRecordMappingRequest `json:"dnsRecordMapping"`
	DnsRecordName    string                            `json:"dnsRecordName"`
	DnsRecordType    string                            `json:"dnsRecordType"`
	Ttl              int32                             `json:"ttl"`
}

type DnsDomainRecordDetailResponse struct {
	ProjectId        string                                   `json:"projectId,omitempty"`
	BlockId          string                                   `json:"blockId,omitempty"`
	DnsRecordId      string                                   `json:"dnsRecordId,omitempty"`
	DnsRecordMapping []DnsDomainRecordMappingResponse `json:"dnsRecordMapping,omitempty"`
	DnsRecordName    string                                   `json:"dnsRecordName,omitempty"`
	DnsRecordType    string                                   `json:"dnsRecordType,omitempty"`
	DnsState         string                                   `json:"dnsState,omitempty"`
	ServiceZoneId    string                                   `json:"serviceZoneId,omitempty"`
	Ttl              int32                                    `json:"ttl,omitempty"`
	CreatedBy        string                                   `json:"createdBy,omitempty"`
	CreatedDt        time.Time                             `json:"createdDt,omitempty"`
	ModifiedBy       string                                   `json:"modifiedBy,omitempty"`
	ModifiedDt       time.Time                             `json:"modifiedDt,omitempty"`
}

type DnsDomainRecordListItemResponse struct {
	DnsRecordId        string   `json:"dnsRecordId,omitempty"`
	DnsRecordName      string   `json:"dnsRecordName,omitempty"`
	DnsRecordType      string   `json:"dnsRecordType,omitempty"`
	DnsState           string   `json:"dnsState,omitempty"`
	RecordDestinations []string `json:"recordDestinations,omitempty"`
	Ttl                int32    `json:"ttl,omitempty"`
}

type DnsDomainRecordMappingResponse struct {
	DnsRecordMapId    string `json:"dnsRecordMapId,omitempty"`
	Preference        int32  `json:"preference,omitempty"`
	RecordDestination string `json:"recordDestination,omitempty"`
}

type DnsDomainServiceDetailResponse struct {
	ProjectId             string       `json:"projectId,omitempty"`
	AutoExtension         *bool        `json:"autoExtension,omitempty"`
	BlockId               string       `json:"blockId,omitempty"`
	CountryType           string       `json:"countryType,omitempty"`
	DnsDomainId           string       `json:"dnsDomainId,omitempty"`
	DnsDomainName         string       `json:"dnsDomainName,omitempty"`
	DnsDomainType         string       `json:"dnsDomainType,omitempty"`
	DnsEnvUsage           string       `json:"dnsEnvUsage,omitempty"`
	DnsState              string       `json:"dnsState,omitempty"`
	EnCompanyName         string       `json:"enCompanyName,omitempty"`
	EnDetailAddress       string       `json:"enDetailAddress,omitempty"`
	ExpiredDate           time.Time `json:"expiredDate,omitempty"`
	FirstEnAddress        string       `json:"firstEnAddress,omitempty"`
	FirstKoAddress        string       `json:"firstKoAddress,omitempty"`
	KoCompanyName         string       `json:"koCompanyName,omitempty"`
	KoDetailAddress       string       `json:"koDetailAddress,omitempty"`
	LinkedRecordCount     int32        `json:"linkedRecordCount,omitempty"`
	PostalCode            string       `json:"postalCode,omitempty"`
	RegisteredByEmail     string       `json:"registeredByEmail,omitempty"`
	RegisteredByFirstName string       `json:"registeredByFirstName,omitempty"`
	RegisteredByLastName  string       `json:"registeredByLastName,omitempty"`
	RegisteredByTelno     string       `json:"registeredByTelno,omitempty"`
	SecondEnAddress       string       `json:"secondEnAddress,omitempty"`
	SecondKoAddress       string       `json:"secondKoAddress,omitempty"`
	ServiceZoneId         string       `json:"serviceZoneId,omitempty"`
	StartDate             time.Time `json:"startDate,omitempty"`
	DnsDescription        string       `json:"dnsDescription,omitempty"`
	CreatedBy             string       `json:"createdBy,omitempty"`
	CreatedDt             time.Time `json:"createdDt,omitempty"`
	ModifiedBy            string       `json:"modifiedBy,omitempty"`
	ModifiedDt            time.Time `json:"modifiedDt,omitempty"`
}

type DnsDomainServiceListItemResponse struct {
	AutoExtension     *bool        `json:"autoExtension,omitempty"`
	DnsDomainId       string       `json:"dnsDomainId,omitempty"`
	DnsDomainName     string       `json:"dnsDomainName,omitempty"`
	DnsDomainType     string       `json:"dnsDomainType,omitempty"`
	DnsEnvUsage       string       `json:"dnsEnvUsage,omitempty"`
	DnsState          string       `json:"dnsState,omitempty"`
	ExpiredDate       time.Time `json:"expiredDate,omitempty"`
	LinkedRecordCount int32        `json:"linkedRecordCount,omitempty"`
	Region            string       `json:"region,omitempty"`
	StartDate         time.Time `json:"startDate,omitempty"`
}

type DnsOpenApiV2ControllerApiListDnsDomainOpts struct {
	DnsDomainName optional.String
	DnsEnvUsage   optional.String
	CreatedBy     optional.String
	Page          optional.Int32
	Size          optional.Int32
	Sort          optional.Interface
}

type DnsOpenApiV2ControllerApiListDnsRecordOpts struct {
	DnsRecordName     optional.String
	DnsRecordType     optional.String
	RecordDestination optional.String
	CreatedBy         optional.String
	Page              optional.Int32
	Size              optional.Int32
	Sort              optional.Interface
}

type DnsRecordMappingRequest struct {
	Preference        int32  `json:"preference,omitempty"`
	RecordDestination string `json:"recordDestination"`
}

type ListResponseDnsDomainRecordListItemResponse struct {
	Contents   []DnsDomainRecordListItemResponse `json:"contents,omitempty"`
	TotalCount int32                                     `json:"totalCount,omitempty"`
}

type ListResponseDnsDomainServiceListItemResponse struct {
	Contents   []DnsDomainServiceListItemResponse `json:"contents,omitempty"`
	TotalCount int32                                      `json:"totalCount,omitempty"`
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
