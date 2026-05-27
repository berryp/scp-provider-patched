package project

import (
	"github.com/antihax/optional"
	"time"
)

type APIKey struct {
	Key    string
	Prefix string
}

type AccountDetailResponseV3 struct {
	AccountAdminEmail           string                             `json:"accountAdminEmail,omitempty"`
	AccountAdminId              string                             `json:"accountAdminId,omitempty"`
	AccountAdminLoginId         string                             `json:"accountAdminLoginId,omitempty"`
	AccountAdminName            string                             `json:"accountAdminName,omitempty"`
	AccountCreateRequestId      string                             `json:"accountCreateRequestId,omitempty"`
	AccountId                   string                             `json:"accountId,omitempty"`
	AccountName                 string                             `json:"accountName,omitempty"`
	AccountType                 string                             `json:"accountType,omitempty"`
	AccountUsers                []AccountMemberResponse `json:"accountUsers,omitempty"`
	BillingCorporationCode      string                             `json:"billingCorporationCode,omitempty"`
	BillingCorporationName      string                             `json:"billingCorporationName,omitempty"`
	BusinessCategoryId          string                             `json:"businessCategoryId,omitempty"`
	BusinessCategoryName        string                             `json:"businessCategoryName,omitempty"`
	BusinessCategoryUsers       []BusinessCategoryUser  `json:"businessCategoryUsers,omitempty"`
	BusinessGroupId             string                             `json:"businessGroupId,omitempty"`
	BusinessLicenseFileId       string                             `json:"businessLicenseFileId,omitempty"`
	BusinessLicenseFileName     string                             `json:"businessLicenseFileName,omitempty"`
	BusinessRegistrationNumber  string                             `json:"businessRegistrationNumber,omitempty"`
	CompanyId                   string                             `json:"companyId,omitempty"`
	CompanyName                 string                             `json:"companyName,omitempty"`
	CustomerCode                string                             `json:"customerCode,omitempty"`
	IndustryType                string                             `json:"industryType,omitempty"`
	IsCcbsSend                  string                             `json:"isCcbsSend,omitempty"`
	IsDeleted                   string                             `json:"isDeleted,omitempty"`
	IsInternalProject           *bool                              `json:"isInternalProject,omitempty"`
	IsProjectCreatable          *bool                              `json:"isProjectCreatable,omitempty"`
	NetworkType                 string                             `json:"networkType,omitempty"`
	PaymentAccountId            string                             `json:"paymentAccountId,omitempty"`
	ProductContractInfos        []ProductContractInfoV3 `json:"productContractInfos,omitempty"`
	SalesRepresentativeKnoxId   string                             `json:"salesRepresentativeKnoxId,omitempty"`
	SalesRepresentativeKnoxName string                             `json:"salesRepresentativeKnoxName,omitempty"`
	SapId                       string                             `json:"sapId,omitempty"`
	ServiceZones                []ZoneResponseV3        `json:"serviceZones,omitempty"`
	AccountDescription          string                             `json:"accountDescription,omitempty"`
	CreatedBy                   string                             `json:"createdBy"`
	CreatedByName               string                             `json:"createdByName"`
	CreatedByEmail              string                             `json:"createdByEmail"`
	CreatedDt                   time.Time                       `json:"createdDt"`
	ModifiedBy                  string                             `json:"modifiedBy"`
	ModifiedByName              string                             `json:"modifiedByName"`
	ModifiedByEmail             string                             `json:"modifiedByEmail"`
	ModifiedDt                  time.Time                       `json:"modifiedDt"`
}

type AccountDetailV2 struct {
	AccountAdminId             string                            `json:"accountAdminId,omitempty"`
	AccountCreateRequestId     string                            `json:"accountCreateRequestId,omitempty"`
	AccountId                  string                            `json:"accountId,omitempty"`
	AccountName                string                            `json:"accountName,omitempty"`
	AccountType                string                            `json:"accountType,omitempty"`
	AccountUserEmail           string                            `json:"accountUserEmail,omitempty"`
	AccountUserId              string                            `json:"accountUserId,omitempty"`
	AccountUserLoginId         string                            `json:"accountUserLoginId,omitempty"`
	AccountUserName            string                            `json:"accountUserName,omitempty"`
	BillingCorporationCode     string                            `json:"billingCorporationCode,omitempty"`
	BillingCorporationName     string                            `json:"billingCorporationName,omitempty"`
	BusinessCategoryId         string                            `json:"businessCategoryId,omitempty"`
	BusinessCategoryName       string                            `json:"businessCategoryName,omitempty"`
	BusinessCategoryUsers      []BusinessCategoryUser `json:"businessCategoryUsers,omitempty"`
	BusinessGroupId            string                            `json:"businessGroupId,omitempty"`
	BusinessLicenseFileId      string                            `json:"businessLicenseFileId,omitempty"`
	BusinessLicenseFileName    string                            `json:"businessLicenseFileName,omitempty"`
	BusinessRegistrationNumber string                            `json:"businessRegistrationNumber,omitempty"`
	CompanyId                  string                            `json:"companyId,omitempty"`
	CompanyName                string                            `json:"companyName,omitempty"`
	CreatedByLoginId           string                            `json:"createdByLoginId,omitempty"`
	CreatedDtStr               string                            `json:"createdDtStr,omitempty"`
	CustomerCode               string                            `json:"customerCode,omitempty"`
	ExchangeRate               string                            `json:"exchangeRate,omitempty"`
	FixedExchangeRate          string                            `json:"fixedExchangeRate,omitempty"`
	IndustryType               string                            `json:"industryType,omitempty"`
	IrpCode                    string                            `json:"irpCode,omitempty"`
	IsCcbsSend                 string                            `json:"isCcbsSend,omitempty"`
	IsDeleted                  string                            `json:"isDeleted,omitempty"`
	IsInternalProject          *bool                             `json:"isInternalProject,omitempty"`
	IsProjectCreatable         *bool                             `json:"isProjectCreatable,omitempty"`
	LegalReviewNumber          string                            `json:"legalReviewNumber,omitempty"`
	ModifiedByLoginId          string                            `json:"modifiedByLoginId,omitempty"`
	ModifiedDtStr              string                            `json:"modifiedDtStr,omitempty"`
	NetworkType                string                            `json:"networkType,omitempty"`
	PaymentAccountId           string                            `json:"paymentAccountId,omitempty"`
	ProductContractInfos       []ProductContractInfo  `json:"productContractInfos,omitempty"`
	RoundPlaces                string                            `json:"roundPlaces,omitempty"`
	RoundType                  string                            `json:"roundType,omitempty"`
	SalesDepartmentId          string                            `json:"salesDepartmentId,omitempty"`
	SalesOrganizationId        string                            `json:"salesOrganizationId,omitempty"`
	SapContractCode            string                            `json:"sapContractCode,omitempty"`
	SapContractName            string                            `json:"sapContractName,omitempty"`
	SapContractNumber          string                            `json:"sapContractNumber,omitempty"`
	SapId                      string                            `json:"sapId,omitempty"`
	ServiceZones               []AccountZoneV2        `json:"serviceZones,omitempty"`
	AccountDescription         string                            `json:"accountDescription,omitempty"`
	CreatedBy                  string                            `json:"createdBy,omitempty"`
	CreatedByName              string                            `json:"createdByName,omitempty"`
	CreatedDt                  time.Time                      `json:"createdDt,omitempty"`
	ModifiedBy                 string                            `json:"modifiedBy,omitempty"`
	ModifiedByName             string                            `json:"modifiedByName,omitempty"`
	ModifiedDt                 time.Time                      `json:"modifiedDt,omitempty"`
}

type AccountMemberResponse struct {
	Company    string `json:"company,omitempty"`
	Email      string `json:"email,omitempty"`
	Id         string `json:"id,omitempty"`
	IsKnoxUser *bool  `json:"isKnoxUser,omitempty"`
	MemberType string `json:"memberType,omitempty"`
	Name       string `json:"name,omitempty"`
}

type AccountResponseV3 struct {
	AccountAdminId     string       `json:"accountAdminId,omitempty"`
	AccountId          string       `json:"accountId,omitempty"`
	AccountName        string       `json:"accountName,omitempty"`
	AccountType        string       `json:"accountType,omitempty"`
	IsProjectCreatable *bool        `json:"isProjectCreatable,omitempty"`
	NetworkType        string       `json:"networkType,omitempty"`
	CreatedBy          string       `json:"createdBy"`
	CreatedByName      string       `json:"createdByName"`
	CreatedByEmail     string       `json:"createdByEmail"`
	CreatedDt          time.Time `json:"createdDt"`
	ModifiedBy         string       `json:"modifiedBy"`
	ModifiedByName     string       `json:"modifiedByName"`
	ModifiedByEmail    string       `json:"modifiedByEmail"`
	ModifiedDt         time.Time `json:"modifiedDt"`
}

type AccountUserUpdateRequest struct {
	UserEmails []string `json:"userEmails,omitempty"`
}

type AccountV2 struct {
	AccountAdminId     string       `json:"accountAdminId,omitempty"`
	AccountId          string       `json:"accountId,omitempty"`
	AccountName        string       `json:"accountName,omitempty"`
	AccountType        string       `json:"accountType,omitempty"`
	IsProjectCreatable *bool        `json:"isProjectCreatable,omitempty"`
	NetworkType        string       `json:"networkType,omitempty"`
	CreatedBy          string       `json:"createdBy,omitempty"`
	CreatedDt          time.Time `json:"createdDt,omitempty"`
	ModifiedBy         string       `json:"modifiedBy,omitempty"`
	ModifiedDt         time.Time `json:"modifiedDt,omitempty"`
}

type AccountV3ControllerApiCheckAccountMemberOpts struct {
	MemberType optional.String
}

type AccountV3ControllerApiListAccountMembersOpts struct {
	MemberType optional.String
}

type AccountZoneV2 struct {
	BlockId             string `json:"blockId,omitempty"`
	BlockName           string `json:"blockName,omitempty"`
	ServiceZoneId       string `json:"serviceZoneId,omitempty"`
	ServiceZoneLocation string `json:"serviceZoneLocation,omitempty"`
	ServiceZoneName     string `json:"serviceZoneName,omitempty"`
}

type AvailabilityZone struct {
	AvailabilityZoneName string `json:"availabilityZoneName,omitempty"`
}

type BasicAuth struct {
	UserName string `json:"userName,omitempty"`
	Password string `json:"password,omitempty"`
}

type BusinessCategoryUser struct {
	BusinessCategoryUserEmail string `json:"businessCategoryUserEmail,omitempty"`
	BusinessCategoryUserName  string `json:"businessCategoryUserName,omitempty"`
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

type ListProjectRequest struct {
	AccountName          string
	BillYearMonth        string
	IsBillingInfoDemand  bool
	IsResourceInfoDemand bool
	IsUserInfoDemand     bool
	ProjectName          string
	CreatedByEmail       string
}

type ListResponseV2AccountMemberResponse struct {
	Contents   []AccountMemberResponse `json:"contents,omitempty"`
	TotalCount int64                              `json:"totalCount,omitempty"`
}

type ListResponseV2AccountResponseV3 struct {
	Contents   []AccountResponseV3 `json:"contents,omitempty"`
	TotalCount int64                          `json:"totalCount,omitempty"`
}

type ListResponseV2ProjectSummaryResponseV3 struct {
	Contents   []ProjectSummaryResponseV3 `json:"contents,omitempty"`
	TotalCount int64                                 `json:"totalCount,omitempty"`
}

type ListResponseV2ZoneResponseV3 struct {
	Contents   []ZoneResponseV3 `json:"contents,omitempty"`
	TotalCount int64                       `json:"totalCount,omitempty"`
}

type PageResponseV2ProjectDetailV2 struct {
	Contents   []ProjectDetailV2 `json:"contents,omitempty"`
	Page       int32                        `json:"page,omitempty"`
	Size       int32                        `json:"size,omitempty"`
	Sort       []string                     `json:"sort,omitempty"`
	TotalCount int64                        `json:"totalCount,omitempty"`
}

type PageResponseV2ProjectResponseV3 struct {
	Contents   []ProjectResponseV3 `json:"contents,omitempty"`
	Page       int32                          `json:"page,omitempty"`
	Size       int32                          `json:"size,omitempty"`
	Sort       []string                       `json:"sort,omitempty"`
	TotalCount int64                          `json:"totalCount,omitempty"`
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

type ProductCategoryResource struct {
	ProductCategoryName string                         `json:"productCategoryName,omitempty"`
	ProductResources    []ProductResourceV2 `json:"productResources,omitempty"`
}

type ProductCategoryV2 struct {
	ProductCategoryId          string                         `json:"productCategoryId,omitempty"`
	ProductCategoryName        string                         `json:"productCategoryName,omitempty"`
	ProductCategoryState       string                         `json:"productCategoryState,omitempty"`
	ProductSet                 string                         `json:"productSet,omitempty"`
	Products                   []ProductOfferingV2 `json:"products,omitempty"`
	ProductCategoryDescription string                         `json:"productCategoryDescription,omitempty"`
}

type ProductContractInfo struct {
	DivideRate            string `json:"divideRate,omitempty"`
	ItemNumber            string `json:"itemNumber,omitempty"`
	L7Product             string `json:"l7Product,omitempty"`
	L7ProductSerialNumber int32  `json:"l7ProductSerialNumber,omitempty"`
	WbsCode               string `json:"wbsCode,omitempty"`
}

type ProductContractInfoV3 struct {
	L3ProductCode string `json:"l3ProductCode,omitempty"`
	L3ProductName string `json:"l3ProductName,omitempty"`
	WbsCode       string `json:"wbsCode,omitempty"`
	WbsName       string `json:"wbsName,omitempty"`
}

type ProductOfferingV2 struct {
	IsProductCreatable         string `json:"isProductCreatable,omitempty"`
	ProductOfferingDetailInfo  string `json:"productOfferingDetailInfo,omitempty"`
	ProductOfferingId          string `json:"productOfferingId,omitempty"`
	ProductOfferingName        string `json:"productOfferingName,omitempty"`
	ProductOfferingState       string `json:"productOfferingState,omitempty"`
	ProductOfferingDescription string `json:"productOfferingDescription,omitempty"`
}

type ProductResourceV2 struct {
	ProductOfferingName          string `json:"productOfferingName,omitempty"`
	ProductOfferingResourceCount int64  `json:"productOfferingResourceCount,omitempty"`
}

type ProjectBudgetCreateRequestV3 struct {
	AlarmThresholds        []int32 `json:"alarmThresholds,omitempty"`
	BudgetAmount           float64 `json:"budgetAmount,omitempty"`
	BudgetStartYymm        string  `json:"budgetStartYymm,omitempty"`
	BudgetType             string  `json:"budgetType,omitempty"`
	CreatePreventThreshold int32   `json:"createPreventThreshold,omitempty"`
	IsBudgetUsed           *bool   `json:"isBudgetUsed,omitempty"`
	IsCreatePrevent        *bool   `json:"isCreatePrevent,omitempty"`
	RequestGuide           string  `json:"requestGuide,omitempty"`
}

type ProjectBudgetResponse struct {
	AlarmThresholds        []int32                                          `json:"alarmThresholds,omitempty"`
	BudgetAmount           float64                                          `json:"budgetAmount,omitempty"`
	BudgetStartYymm        *ProjectBudgetResponseBudgetStartYymm `json:"budgetStartYymm,omitempty"`
	BudgetType             string                                           `json:"budgetType,omitempty"`
	BudgetTypeChangeDt     time.Time                                     `json:"budgetTypeChangeDt,omitempty"`
	CreatePreventThreshold int32                                            `json:"createPreventThreshold,omitempty"`
	CreatePreventYn        string                                           `json:"createPreventYn,omitempty"`
	EmailAlarmSendYn       string                                           `json:"emailAlarmSendYn,omitempty"`
	RequestGuide           string                                           `json:"requestGuide,omitempty"`
	UseYn                  string                                           `json:"useYn,omitempty"`
}

type ProjectBudgetResponseBudgetStartYymm struct {
	Year       int32  `json:"year,omitempty"`
	Month      string `json:"month,omitempty"`
	MonthValue int32  `json:"monthValue,omitempty"`
	LeapYear   *bool  `json:"leapYear,omitempty"`
}

type ProjectBudgetResponseV3 struct {
	AlarmThresholds        []int32 `json:"alarmThresholds,omitempty"`
	BudgetAmount           float64 `json:"budgetAmount,omitempty"`
	BudgetStartYymm        string  `json:"budgetStartYymm,omitempty"`
	CreatePreventThreshold int32   `json:"createPreventThreshold,omitempty"`
	IsBudgetUsed           *bool   `json:"isBudgetUsed,omitempty"`
	IsCreatePrevent        *bool   `json:"isCreatePrevent,omitempty"`
	RequestGuide           string  `json:"requestGuide,omitempty"`
}

type ProjectControllerV2ApiListAccountsByMyProject1Opts struct {
	AccessLevel         optional.String
	ActionName          optional.String
	CmpServiceName      optional.String
	IsUserAuthorization optional.Bool
	MyProject           optional.Bool
}

type ProjectControllerV2ApiListMyProjectsOpts struct {
	AccountName          optional.String
	BillYearMonth        optional.String
	IsBillingInfoDemand  optional.Bool
	IsResourceInfoDemand optional.Bool
	IsUserInfoDemand     optional.Bool
	ProjectName          optional.String
	CreatedByEmail       optional.String
	Page                 optional.Int32
	Size                 optional.Int32
	Sort                 optional.Interface
}

type ProjectControllerV2ApiListProductResourcesOpts struct {
	ProductCategoryId optional.String
}

type ProjectControllerV2ApiListProjectProductResourcesOpts struct {
	ProductCategoryId optional.String
}

type ProjectControllerV2ApiListProjectProductsOpts struct {
	LanguageCode optional.String
}

type ProjectControllerV2ApiListProjects1Opts struct {
	AccessLevel         optional.String
	ActionName          optional.String
	CmpServiceName      optional.String
	IsUserAuthorization optional.Bool
}

type ProjectCreateMemberRequest struct {
	AdministratorGroupMemberEmails     []string `json:"administratorGroupMemberEmails"`
	OperatorGroupMemberEmails          []string `json:"operatorGroupMemberEmails,omitempty"`
	ProjectManagementGroupMemberEmails []string `json:"projectManagementGroupMemberEmails,omitempty"`
}

type ProjectCreateRequest struct {
	AccountId          string `json:"accountId"`
	ProjectName        string `json:"projectName"`
	ProjectDescription string `json:"projectDescription,omitempty"`
}

type ProjectCreateRequestV3 struct {
	AccountId          string                                   `json:"accountId"`
	Budget             *ProjectBudgetCreateRequestV3 `json:"budget,omitempty"`
	DefaultZoneId      string                                   `json:"defaultZoneId"`
	IgwCreateYn        string                                   `json:"igwCreateYn"`
	MemberList         *ProjectCreateMemberRequest   `json:"memberList"`
	ProjectName        string                                   `json:"projectName"`
	ProjectDescription string                                   `json:"projectDescription,omitempty"`
}

type ProjectDetailResponseV3 struct {
	ProjectId               string                              `json:"projectId,omitempty"`
	AccountAdminEmail       string                              `json:"accountAdminEmail,omitempty"`
	AccountAdminName        string                              `json:"accountAdminName,omitempty"`
	AccountId               string                              `json:"accountId,omitempty"`
	AccountName             string                              `json:"accountName,omitempty"`
	AccountType             string                              `json:"accountType,omitempty"`
	BillingOrganizationCode string                              `json:"billingOrganizationCode,omitempty"`
	Budget                  *ProjectBudgetResponseV3 `json:"budget,omitempty"`
	BusinessCategoryId      string                              `json:"businessCategoryId,omitempty"`
	BusinessCategoryName    string                              `json:"businessCategoryName,omitempty"`
	BusinessCategoryUsers   []BusinessCategoryUser   `json:"businessCategoryUsers,omitempty"`
	BusinessGroupId         string                              `json:"businessGroupId,omitempty"`
	CompanyId               string                              `json:"companyId,omitempty"`
	CompanyName             string                              `json:"companyName,omitempty"`
	DefaultZoneId           string                              `json:"defaultZoneId,omitempty"`
	FreeTrialExpiredDate    string                              `json:"freeTrialExpiredDate,omitempty"`
	FreeTrialExpiredDday    string                              `json:"freeTrialExpiredDday,omitempty"`
	FreeTrialStartDate      string                              `json:"freeTrialStartDate,omitempty"`
	IgwCreateYn             string                              `json:"igwCreateYn,omitempty"`
	NetworkType             string                              `json:"networkType,omitempty"`
	PriceSystemYear         string                              `json:"priceSystemYear,omitempty"`
	ProjectName             string                              `json:"projectName"`
	ProjectState            string                              `json:"projectState,omitempty"`
	ServiceZones            []ZoneResponseV3         `json:"serviceZones,omitempty"`
	VpcVersion              string                              `json:"vpcVersion,omitempty"`
	ProjectDescription      string                              `json:"projectDescription,omitempty"`
	CreatedBy               string                              `json:"createdBy"`
	CreatedByName           string                              `json:"createdByName"`
	CreatedByEmail          string                              `json:"createdByEmail"`
	CreatedDt               time.Time                        `json:"createdDt"`
	ModifiedBy              string                              `json:"modifiedBy"`
	ModifiedByName          string                              `json:"modifiedByName"`
	ModifiedByEmail         string                              `json:"modifiedByEmail"`
	ModifiedDt              time.Time                        `json:"modifiedDt"`
}

type ProjectDetailV2 struct {
	ProjectId               string                            `json:"projectId,omitempty"`
	AccountCode             string                            `json:"accountCode,omitempty"`
	AccountId               string                            `json:"accountId,omitempty"`
	AccountName             string                            `json:"accountName,omitempty"`
	AccountType             string                            `json:"accountType,omitempty"`
	AccountUserEmail        string                            `json:"accountUserEmail,omitempty"`
	AccountUserName         string                            `json:"accountUserName,omitempty"`
	BillingOrganizationCode string                            `json:"billingOrganizationCode,omitempty"`
	BillingUnit             string                            `json:"billingUnit,omitempty"`
	Budget                  *ProjectBudgetResponse `json:"budget,omitempty"`
	BusinessCategoryId      string                            `json:"businessCategoryId,omitempty"`
	BusinessCategoryName    string                            `json:"businessCategoryName,omitempty"`
	BusinessCategoryUsers   []BusinessCategoryUser `json:"businessCategoryUsers,omitempty"`
	BusinessGroupId         string                            `json:"businessGroupId,omitempty"`
	CompanyId               string                            `json:"companyId,omitempty"`
	CompanyName             string                            `json:"companyName,omitempty"`
	CreatedDtStr            string                            `json:"createdDtStr,omitempty"`
	CurrentMonthBillAmount  float64                           `json:"currentMonthBillAmount,omitempty"`
	DefaultZoneId           string                            `json:"defaultZoneId,omitempty"`
	EstimatedUsedAmount     float64                           `json:"estimatedUsedAmount,omitempty"`
	FixedCostAmount         string                            `json:"fixedCostAmount,omitempty"`
	FixedExchangeRate       string                            `json:"fixedExchangeRate,omitempty"`
	FreeTrialExpiredDate    string                            `json:"freeTrialExpiredDate,omitempty"`
	FreeTrialExpiredDday    string                            `json:"freeTrialExpiredDday,omitempty"`
	FreeTrialStartDate      string                            `json:"freeTrialStartDate,omitempty"`
	IgwCreateYn             string                            `json:"igwCreateYn,omitempty"`
	IsFixedExchangeRate     string                            `json:"isFixedExchangeRate,omitempty"`
	LastMonthBillAmount     float64                           `json:"lastMonthBillAmount,omitempty"`
	LogArchiveProjectYn     string                            `json:"logArchiveProjectYn,omitempty"`
	ModifiedDtStr           string                            `json:"modifiedDtStr,omitempty"`
	NetworkType             string                            `json:"networkType,omitempty"`
	OrganizationId          string                            `json:"organizationId,omitempty"`
	OrganizationName        string                            `json:"organizationName,omitempty"`
	OwnerId                 string                            `json:"ownerId,omitempty"`
	OwnerName               string                            `json:"ownerName,omitempty"`
	PriceSystemYear         string                            `json:"priceSystemYear,omitempty"`
	ProjectAttributes       map[string]string                 `json:"projectAttributes,omitempty"`
	ProjectBudget           int64                             `json:"projectBudget,omitempty"`
	ProjectMemberCount      int32                             `json:"projectMemberCount,omitempty"`
	ProjectName             string                            `json:"projectName"`
	ProjectResourceCount    int32                             `json:"projectResourceCount,omitempty"`
	ProjectServiceCount     int32                             `json:"projectServiceCount,omitempty"`
	ProjectState            string                            `json:"projectState,omitempty"`
	ServiceZones            []ProjectZoneV2        `json:"serviceZones,omitempty"`
	StartDt                 string                            `json:"startDt,omitempty"`
	StartDtStr              string                            `json:"startDtStr,omitempty"`
	VpcVersion              string                            `json:"vpcVersion,omitempty"`
	ProjectDescription      string                            `json:"projectDescription,omitempty"`
	CreatedBy               string                            `json:"createdBy,omitempty"`
	CreatedByName           string                            `json:"createdByName,omitempty"`
	CreatedByEmail          string                            `json:"createdByEmail,omitempty"`
	CreatedDt               time.Time                      `json:"createdDt,omitempty"`
	ModifiedBy              string                            `json:"modifiedBy,omitempty"`
	ModifiedByName          string                            `json:"modifiedByName,omitempty"`
	ModifiedByEmail         string                            `json:"modifiedByEmail,omitempty"`
	ModifiedDt              time.Time                      `json:"modifiedDt,omitempty"`
}

type ProjectResponseAccountV2 struct {
	Contents   []AccountV2 `json:"contents,omitempty"`
	TotalCount int32                  `json:"totalCount"`
}

type ProjectResponseProductCategoryResource struct {
	Contents   []ProductCategoryResource `json:"contents,omitempty"`
	TotalCount int32                                `json:"totalCount"`
}

type ProjectResponseProductCategoryV2 struct {
	Contents   []ProductCategoryV2 `json:"contents,omitempty"`
	TotalCount int32                          `json:"totalCount"`
}

type ProjectResponseProjectV2 struct {
	Contents   []ProjectV2 `json:"contents,omitempty"`
	TotalCount int32                  `json:"totalCount"`
}

type ProjectResponseV3 struct {
	ProjectId               string                            `json:"projectId,omitempty"`
	AccountAdminEmail       string                            `json:"accountAdminEmail,omitempty"`
	AccountAdminName        string                            `json:"accountAdminName,omitempty"`
	AccountId               string                            `json:"accountId,omitempty"`
	AccountName             string                            `json:"accountName,omitempty"`
	AccountType             string                            `json:"accountType,omitempty"`
	BillingOrganizationCode string                            `json:"billingOrganizationCode,omitempty"`
	BusinessCategoryId      string                            `json:"businessCategoryId,omitempty"`
	BusinessCategoryName    string                            `json:"businessCategoryName,omitempty"`
	BusinessCategoryUsers   []BusinessCategoryUser `json:"businessCategoryUsers,omitempty"`
	BusinessGroupId         string                            `json:"businessGroupId,omitempty"`
	CompanyId               string                            `json:"companyId,omitempty"`
	CompanyName             string                            `json:"companyName,omitempty"`
	CurrentMonthBillAmount  float64                           `json:"currentMonthBillAmount,omitempty"`
	DefaultZoneId           string                            `json:"defaultZoneId,omitempty"`
	EstimatedUsedAmount     float64                           `json:"estimatedUsedAmount,omitempty"`
	FreeTrialExpiredDate    string                            `json:"freeTrialExpiredDate,omitempty"`
	FreeTrialExpiredDday    string                            `json:"freeTrialExpiredDday,omitempty"`
	FreeTrialStartDate      string                            `json:"freeTrialStartDate,omitempty"`
	IgwCreateYn             string                            `json:"igwCreateYn,omitempty"`
	LastMonthBillAmount     float64                           `json:"lastMonthBillAmount,omitempty"`
	NetworkType             string                            `json:"networkType,omitempty"`
	PriceSystemYear         string                            `json:"priceSystemYear,omitempty"`
	ProjectMemberCount      int32                             `json:"projectMemberCount,omitempty"`
	ProjectName             string                            `json:"projectName"`
	ProjectResourceCount    int32                             `json:"projectResourceCount,omitempty"`
	ProjectServiceCount     int32                             `json:"projectServiceCount,omitempty"`
	ProjectState            string                            `json:"projectState,omitempty"`
	ServiceZones            []ZoneResponseV3       `json:"serviceZones,omitempty"`
	VpcVersion              string                            `json:"vpcVersion,omitempty"`
	ProjectDescription      string                            `json:"projectDescription,omitempty"`
	CreatedBy               string                            `json:"createdBy"`
	CreatedByName           string                            `json:"createdByName"`
	CreatedByEmail          string                            `json:"createdByEmail"`
	CreatedDt               time.Time                      `json:"createdDt"`
	ModifiedBy              string                            `json:"modifiedBy"`
	ModifiedByName          string                            `json:"modifiedByName"`
	ModifiedByEmail         string                            `json:"modifiedByEmail"`
	ModifiedDt              time.Time                      `json:"modifiedDt"`
}

type ProjectSummaryResponseV3 struct {
	ProjectId               string       `json:"projectId,omitempty"`
	AccountId               string       `json:"accountId,omitempty"`
	AccountName             string       `json:"accountName,omitempty"`
	AccountType             string       `json:"accountType,omitempty"`
	BillingOrganizationCode string       `json:"billingOrganizationCode,omitempty"`
	BusinessCategoryId      string       `json:"businessCategoryId,omitempty"`
	BusinessGroupId         string       `json:"businessGroupId,omitempty"`
	CompanyId               string       `json:"companyId,omitempty"`
	ProjectName             string       `json:"projectName"`
	ProjectState            string       `json:"projectState,omitempty"`
	ProjectDescription      string       `json:"projectDescription,omitempty"`
	CreatedBy               string       `json:"createdBy"`
	CreatedByName           string       `json:"createdByName"`
	CreatedByEmail          string       `json:"createdByEmail"`
	CreatedDt               time.Time `json:"createdDt"`
	ModifiedBy              string       `json:"modifiedBy"`
	ModifiedByName          string       `json:"modifiedByName"`
	ModifiedByEmail         string       `json:"modifiedByEmail"`
	ModifiedDt              time.Time `json:"modifiedDt"`
}

type ProjectUpdateRequest struct {
	ProjectName        string `json:"projectName"`
	ProjectDescription string `json:"projectDescription,omitempty"`
}

type ProjectV2 struct {
	ProjectId               string            `json:"projectId,omitempty"`
	AccountCode             string            `json:"accountCode,omitempty"`
	AccountId               string            `json:"accountId,omitempty"`
	AccountName             string            `json:"accountName,omitempty"`
	AccountType             string            `json:"accountType,omitempty"`
	BillingOrganizationCode string            `json:"billingOrganizationCode,omitempty"`
	BillingUnit             string            `json:"billingUnit,omitempty"`
	BusinessCategoryId      string            `json:"businessCategoryId,omitempty"`
	BusinessGroupId         string            `json:"businessGroupId,omitempty"`
	CompanyId               string            `json:"companyId,omitempty"`
	CreatedDtStr            string            `json:"createdDtStr,omitempty"`
	FixedCostAmount         string            `json:"fixedCostAmount,omitempty"`
	FixedExchangeRate       string            `json:"fixedExchangeRate,omitempty"`
	IsFixedExchangeRate     string            `json:"isFixedExchangeRate,omitempty"`
	ModifiedDtStr           string            `json:"modifiedDtStr,omitempty"`
	OrganizationId          string            `json:"organizationId,omitempty"`
	OwnerId                 string            `json:"ownerId,omitempty"`
	OwnerName               string            `json:"ownerName,omitempty"`
	ProjectAttributes       map[string]string `json:"projectAttributes,omitempty"`
	ProjectBudget           int64             `json:"projectBudget,omitempty"`
	ProjectName             string            `json:"projectName"`
	ProjectState            string            `json:"projectState,omitempty"`
	StartDt                 string            `json:"startDt,omitempty"`
	StartDtStr              string            `json:"startDtStr,omitempty"`
	ProjectDescription      string            `json:"projectDescription,omitempty"`
	CreatedBy               string            `json:"createdBy,omitempty"`
	CreatedDt               time.Time      `json:"createdDt,omitempty"`
	ModifiedBy              string            `json:"modifiedBy,omitempty"`
	ModifiedDt              time.Time      `json:"modifiedDt,omitempty"`
}

type ProjectV3ControllerApiListProjectsOpts struct {
	AccountName          optional.String
	BillYearMonth        optional.String
	IsBillingInfoDemand  optional.Bool
	IsResourceInfoDemand optional.Bool
	IsUserInfoDemand     optional.Bool
	ProjectName          optional.String
	CreatedByEmail       optional.String
	Page                 optional.Int32
	Size                 optional.Int32
	Sort                 optional.Interface
}

type ProjectZoneResponse struct {
	BlockId      string `json:"blockId,omitempty"`
	InstanceType string `json:"instanceType,omitempty"`
	Region       string `json:"region,omitempty"`
	VpcZone      *bool  `json:"vpcZone,omitempty"`
	ZoneId       string `json:"zoneId,omitempty"`
	ZoneName     string `json:"zoneName,omitempty"`
}

type ProjectZoneV2 struct {
	BlockId             string `json:"blockId,omitempty"`
	ServiceZoneId       string `json:"serviceZoneId,omitempty"`
	ServiceZoneLocation string `json:"serviceZoneLocation,omitempty"`
	ServiceZoneName     string `json:"serviceZoneName,omitempty"`
}

type Response struct {
	Content *interface{} `json:"content,omitempty"`
	Total   int32        `json:"total,omitempty"`
}

type ValidationResponse struct {
	Result *bool `json:"result,omitempty"`
}

type ZoneControllerV2ApiListZonesOfProjectOpts struct {
	InstanceTypes optional.Interface
}

type ZoneResponseV3 struct {
	AvailabilityZones       []AvailabilityZone `json:"availabilityZones,omitempty"`
	BlockId                 string                        `json:"blockId,omitempty"`
	IsMultiAvailabilityZone *bool                         `json:"isMultiAvailabilityZone,omitempty"`
	ServiceZoneId           string                        `json:"serviceZoneId,omitempty"`
	ServiceZoneLocation     string                        `json:"serviceZoneLocation,omitempty"`
	ServiceZoneName         string                        `json:"serviceZoneName,omitempty"`
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
