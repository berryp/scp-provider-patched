package configinspection

import (
	"github.com/antihax/optional"
	"time"
)

type APIKey struct {
	Key    string
	Prefix string
}

type AuthKeyRequest struct {
	AuthKeyId string `json:"authKeyId,omitempty"`
}

type AuthKeyResponse struct {
	AuthKeyCreateDt time.Time `json:"authKeyCreateDt,omitempty"`
	AuthKeyExpireDt time.Time `json:"authKeyExpireDt,omitempty"`
	AuthKeyId       string       `json:"authKeyId,omitempty"`
	AuthKeyState    string       `json:"authKeyState,omitempty"`
	UserId          string       `json:"userId,omitempty"`
}

type BasicAuth struct {
	UserName string `json:"userName,omitempty"`
	Password string `json:"password,omitempty"`
}

type CheckResponse struct {
	Result *bool `json:"result,omitempty"`
}

type ConfigInspectionDashboardOpenApiControllerApiGetDiagnosisResultListOpts struct {
	DiagnosisName  optional.String
	DiagnosisState optional.String
	EndDate        optional.String
	StartDate      optional.String
	Page           optional.Int32
	Size           optional.Int32
	Sort           optional.Interface
}

type ConfigInspectionListResponse struct {
	Contents   []ConfigInspectionSummaryResponse `json:"contents,omitempty"`
	TotalCount int32                                                 `json:"totalCount,omitempty"`
}

type ConfigInspectionOpenApiControllerApiConfigInspectionListOpts struct {
	CspType              optional.Interface
	DiagnosisAccountId   optional.String
	DiagnosisName        optional.String
	EndDate              optional.String
	RecentDiagnosisState optional.Interface
	StartDate            optional.String
	Page                 optional.Int32
	Size                 optional.Int32
	Sort                 optional.Interface
}

type ConfigInspectionSummaryResponse struct {
	CspType              string       `json:"cspType,omitempty"`
	DiagnosisAccount     string       `json:"diagnosisAccount,omitempty"`
	DiagnosisCheckType   string       `json:"diagnosisCheckType,omitempty"`
	DiagnosisId          string       `json:"diagnosisId,omitempty"`
	DiagnosisName        string       `json:"diagnosisName,omitempty"`
	DiagnosisType        string       `json:"diagnosisType,omitempty"`
	ErrorState           string       `json:"errorState,omitempty"`
	PlanType             string       `json:"planType,omitempty"`
	RecentDiagnosisDt    time.Time `json:"recentDiagnosisDt,omitempty"`
	RecentDiagnosisState string       `json:"recentDiagnosisState,omitempty"`
	CreatedDt            time.Time `json:"createdDt,omitempty"`
}

type DiagnosisCreateResponse struct {
	DiagnosisId string `json:"diagnosisId,omitempty"`
}

type DiagnosisObjectDetailResponse struct {
	AuthKeyResponse  *AuthKeyResponse                 `json:"authKeyResponse,omitempty"`
	ScheduleResponse *DiagnosisScheduleResponse       `json:"scheduleResponse,omitempty"`
	SummaryResponse  *ConfigInspectionSummaryResponse `json:"summaryResponse,omitempty"`
}

type DiagnosisObjectListRequest struct {
	AuthKeyRequest             *AuthKeyRequest           `json:"authKeyRequest,omitempty"`
	CspType                    string                                        `json:"cspType,omitempty"`
	DiagnosisCheckType         string                                        `json:"diagnosisCheckType,omitempty"`
	DiagnosisObjectRequestList []DiagnosisObjectRequest  `json:"diagnosisObjectRequestList,omitempty"`
	DiagnosisType              string                                        `json:"diagnosisType,omitempty"`
	PlanType                   string                                        `json:"planType,omitempty"`
	ScheduleRequest            *DiagnosisScheduleRequest `json:"scheduleRequest,omitempty"`
	Tags                       []TagRequest              `json:"tags,omitempty"`
}

type DiagnosisObjectRequest struct {
	DiagnosisAccountId string `json:"diagnosisAccountId,omitempty"`
	DiagnosisId        string `json:"diagnosisId,omitempty"`
	DiagnosisName      string `json:"diagnosisName,omitempty"`
}

type DiagnosisRequest struct {
	ProjectId          string `json:"projectId"`
	AccessKey          string `json:"accessKey"`
	DiagnosisCheckType string `json:"diagnosisCheckType"`
	DiagnosisId        string `json:"diagnosisId"`
	SecretKey          string `json:"secretKey,omitempty"`
	TenantId           string `json:"tenantId,omitempty"`
}

type DiagnosisResultDetail struct {
	ActionGuide        string `json:"actionGuide,omitempty"`
	Changed            *bool  `json:"changed,omitempty"`
	DiagnosisCheckType string `json:"diagnosisCheckType,omitempty"`
	DiagnosisCriteria  string `json:"diagnosisCriteria,omitempty"`
	DiagnosisItem      string `json:"diagnosisItem,omitempty"`
	DiagnosisLayer     string `json:"diagnosisLayer,omitempty"`
	DiagnosisMethod    string `json:"diagnosisMethod,omitempty"`
	DiagnosisResult    string `json:"diagnosisResult,omitempty"`
	ResultContents     string `json:"resultContents,omitempty"`
	SubCategory        string `json:"subCategory,omitempty"`
}

type DiagnosisResultDetailResponse struct {
	ChecklistName            string                                      `json:"checklistName,omitempty"`
	DiagnosisAccount         string                                      `json:"diagnosisAccount,omitempty"`
	DiagnosisCheckType       string                                      `json:"diagnosisCheckType,omitempty"`
	DiagnosisId              string                                      `json:"diagnosisId,omitempty"`
	DiagnosisName            string                                      `json:"diagnosisName,omitempty"`
	DiagnosisRequestSequence string                                      `json:"diagnosisRequestSequence,omitempty"`
	ProceedDate              time.Time                                `json:"proceedDate,omitempty"`
	ResultDetailList         []DiagnosisResultDetail `json:"resultDetailList,omitempty"`
	TotalCount               int32                                       `json:"totalCount,omitempty"`
}

type DiagnosisResultResponse struct {
	Check                    int32        `json:"check,omitempty"`
	CspType                  string       `json:"cspType,omitempty"`
	DiagnosisAccountId       string       `json:"diagnosisAccountId,omitempty"`
	DiagnosisCheckType       string       `json:"diagnosisCheckType,omitempty"`
	DiagnosisId              string       `json:"diagnosisId,omitempty"`
	DiagnosisName            string       `json:"diagnosisName,omitempty"`
	DiagnosisRequestSequence string       `json:"diagnosisRequestSequence,omitempty"`
	DiagnosisResult          string       `json:"diagnosisResult,omitempty"`
	DiagnosisTotalCnt        int32        `json:"diagnosisTotalCnt,omitempty"`
	Error_                   int32        `json:"error,omitempty"`
	Fail                     int32        `json:"fail,omitempty"`
	Na                       int32        `json:"na,omitempty"`
	Pass                     int32        `json:"pass,omitempty"`
	ProceedDate              time.Time `json:"proceedDate,omitempty"`
	Total                    int32        `json:"total,omitempty"`
}

type DiagnosisScheduleRequest struct {
	DiagnosisStartTimePattern string `json:"diagnosisStartTimePattern,omitempty"`
	FrequencyType             string `json:"frequencyType,omitempty"`
	FrequencyValue            string `json:"frequencyValue,omitempty"`
	UseDiagnosisCheckTypeBP   string `json:"useDiagnosisCheckTypeBP,omitempty"`
	UseDiagnosisCheckTypeSSI  string `json:"useDiagnosisCheckTypeSSI,omitempty"`
}

type DiagnosisScheduleResponse struct {
	DiagnosisId               string `json:"diagnosisId,omitempty"`
	DiagnosisStartTimePattern string `json:"diagnosisStartTimePattern,omitempty"`
	FrequencyType             string `json:"frequencyType,omitempty"`
	FrequencyValue            string `json:"frequencyValue,omitempty"`
	UseDiagnosisCheckTypeBP   string `json:"useDiagnosisCheckTypeBP,omitempty"`
	UseDiagnosisCheckTypeSSI  string `json:"useDiagnosisCheckTypeSSI,omitempty"`
}

type PageResponseDiagnosisResultResponse struct {
	Contents   []DiagnosisResultResponse `json:"contents,omitempty"`
	TotalCount int64                                         `json:"totalCount,omitempty"`
}

type TagRequest struct {
	TagKey   string `json:"tagKey"`
	TagValue string `json:"tagValue,omitempty"`
}

type TerminateResponse struct {
	DeleteFlag *bool  `json:"deleteFlag,omitempty"`
	Msg        string `json:"msg,omitempty"`
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
