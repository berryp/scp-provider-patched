package iam

import (
	"github.com/antihax/optional"
	"time"
)

type APIKey struct {
	Key    string
	Prefix string
}

type AccessKeyControllerApiCountAccessKeyOpts struct {
	ProjectId            optional.String
	AccessKeyProjectType optional.String
	AccessKeyState       optional.String
	ActiveYn             optional.Bool
	IssuerType           optional.String
	ProjectName          optional.String
	Page                 optional.Int32
	Size                 optional.Int32
	Sort                 optional.Interface
}

type AccessKeyControllerApiListAccessKeyBlockedProjectsOpts struct {
	ProjectId   optional.String
	ProjectName optional.String
	Page        optional.Int32
	Size        optional.Int32
	Sort        optional.Interface
}

type AccessKeyControllerApiListAccessKeys1Opts struct {
	ProjectId            optional.String
	AccessKeyProjectType optional.String
	AccessKeyState       optional.String
	ActiveYn             optional.Bool
	IssuerType           optional.String
	ProjectName          optional.String
	Page                 optional.Int32
	Size                 optional.Int32
	Sort                 optional.Interface
}

type AccessKeyControllerApiListAccessKeys2Opts struct {
	ProjectId            optional.String
	AccessKeyProjectType optional.String
	AccessKeyState       optional.String
	ActiveYn             optional.Bool
	IssuerType           optional.String
	ProjectName          optional.String
	Page                 optional.Int32
	Size                 optional.Int32
	Sort                 optional.Interface
}

type AccessKeyControllerApiListAccessKeysOpts struct {
	ProjectId            optional.String
	AccessKeyProjectType optional.String
	AccessKeyState       optional.String
	ActiveYn             optional.Bool
	IssuerType           optional.String
	ProjectName          optional.String
	Page                 optional.Int32
	Size                 optional.Int32
	Sort                 optional.Interface
}

type AccessKeyControllerApiListProjectsAccessKeysOpts struct {
	ParentId optional.String
	Page     optional.Int32
	Size     optional.Int32
	Sort     optional.Interface
}

type AccessKeyResponse struct {
	ProjectId          string       `json:"projectId,omitempty"`
	AccessKey          string       `json:"accessKey,omitempty"`
	AccessKeyActivated *bool        `json:"accessKeyActivated,omitempty"`
	AccessKeyId        string       `json:"accessKeyId,omitempty"`
	AccessKeyState     string       `json:"accessKeyState,omitempty"`
	AccessSecretKey    string       `json:"accessSecretKey,omitempty"`
	ExpiredDt          time.Time `json:"expiredDt,omitempty"`
	ProjectName        string       `json:"projectName,omitempty"`
	SecretVaultCount   int32        `json:"secretVaultCount,omitempty"`
	Description        string       `json:"description,omitempty"`
	CreatedBy          string       `json:"createdBy"`
	CreatedByName      string       `json:"createdByName"`
	CreatedByEmail     string       `json:"createdByEmail"`
	CreatedDt          time.Time `json:"createdDt"`
	ModifiedBy         string       `json:"modifiedBy"`
	ModifiedByName     string       `json:"modifiedByName"`
	ModifiedByEmail    string       `json:"modifiedByEmail"`
	ModifiedDt         time.Time `json:"modifiedDt"`
}

type AccessKeySecretCreateRequest struct {
	ProjectId    string `json:"projectId,omitempty"`
	DurationDays int32  `json:"durationDays,omitempty"`
	Description  string `json:"description,omitempty"`
}

type AccessKeySecretUpdateRequest struct {
	AccessKeyIds []string `json:"accessKeyIds"`
}

type AccessKeysResponse struct {
	ProjectId          string       `json:"projectId,omitempty"`
	AccessKey          string       `json:"accessKey,omitempty"`
	AccessKeyActivated *bool        `json:"accessKeyActivated,omitempty"`
	AccessKeyId        string       `json:"accessKeyId,omitempty"`
	AccessKeyState     string       `json:"accessKeyState,omitempty"`
	ExpiredDt          time.Time `json:"expiredDt,omitempty"`
	ProjectName        string       `json:"projectName,omitempty"`
	SecretVaultCount   int32        `json:"secretVaultCount,omitempty"`
	CreatedBy          string       `json:"createdBy"`
	CreatedByName      string       `json:"createdByName"`
	CreatedByEmail     string       `json:"createdByEmail"`
	CreatedDt          time.Time `json:"createdDt"`
	ModifiedBy         string       `json:"modifiedBy"`
	ModifiedByName     string       `json:"modifiedByName"`
	ModifiedByEmail    string       `json:"modifiedByEmail"`
	ModifiedDt         time.Time `json:"modifiedDt"`
}

type AssumeRoleUserResponse struct {
	RoleId  string `json:"roleId,omitempty"`
	RoleSrn string `json:"roleSrn,omitempty"`
}

type AssumeRoleWithSamlRequest struct {
	DurationMinutes int32  `json:"durationMinutes,omitempty"`
	PrincipalSrn    string `json:"principalSrn"`
	RoleSrn         string `json:"roleSrn"`
	SamlAssertion   string `json:"samlAssertion"`
}

type AssumeRoleWithWebIdentityRequest struct {
	DurationSeconds int32  `json:"durationSeconds,omitempty"`
	IdToken         string `json:"idToken"`
	PrincipalSrn    string `json:"principalSrn"`
	RoleSessionName string `json:"roleSessionName"`
	RoleSrn         string `json:"roleSrn"`
}

type AuthorizationCheckRequest struct {
	AccessLevel          string `json:"accessLevel"`
	ActionName           string `json:"actionName"`
	CmpServiceName       string `json:"cmpServiceName"`
	IsCausePolicyDisplay *bool  `json:"isCausePolicyDisplay,omitempty"`
	ResourceId           string `json:"resourceId,omitempty"`
	ResourceProjectId    string `json:"resourceProjectId,omitempty"`
	ResourceSrn          string `json:"resourceSrn,omitempty"`
}

type AuthorizationCheckResponse struct {
	ActionName               string               `json:"actionName,omitempty"`
	AuthorizationCheckResult string               `json:"authorizationCheckResult,omitempty"`
	CausePolicys             []CausePolicy `json:"causePolicys,omitempty"`
	CmpServiceName           string               `json:"cmpServiceName,omitempty"`
	ResourceSrn              string               `json:"resourceSrn,omitempty"`
}

type AuthorizationDashboardResponse struct {
	DefaultSystemGroups []GroupsResponse  `json:"defaultSystemGroups,omitempty"`
	GroupCount          int32                    `json:"groupCount,omitempty"`
	Groups              []GroupsResponse  `json:"groups,omitempty"`
	MemberCount         int32                    `json:"memberCount,omitempty"`
	PolicyCount         int32                    `json:"policyCount,omitempty"`
	Policys             []PolicysResponse `json:"policys,omitempty"`
}

type BasicAuth struct {
	UserName string `json:"userName,omitempty"`
	Password string `json:"password,omitempty"`
}

type BlockedAccessKeyProjectResponse struct {
	ProjectId          string       `json:"projectId,omitempty"`
	AccessKeyId        string       `json:"accessKeyId,omitempty"`
	BlockedAccessKeyId string       `json:"blockedAccessKeyId,omitempty"`
	ProjectName        string       `json:"projectName,omitempty"`
	CreatedBy          string       `json:"createdBy"`
	CreatedByName      string       `json:"createdByName"`
	CreatedByEmail     string       `json:"createdByEmail"`
	CreatedDt          time.Time `json:"createdDt"`
	ModifiedBy         string       `json:"modifiedBy"`
	ModifiedByName     string       `json:"modifiedByName"`
	ModifiedByEmail    string       `json:"modifiedByEmail"`
	ModifiedDt         time.Time `json:"modifiedDt"`
}

type CausePolicy struct {
	PolicyId                string  `json:"policyId,omitempty"`
	PolicyStatementLocation []int32 `json:"policyStatementLocation,omitempty"`
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

type CountResponse struct {
	Count int64 `json:"count,omitempty"`
}

type CredentialResponse struct {
	AccessKey       string       `json:"accessKey,omitempty"`
	AccessKeyId     string       `json:"accessKeyId,omitempty"`
	AccessSecretKey string       `json:"accessSecretKey,omitempty"`
	ExpiredDt       time.Time `json:"expiredDt,omitempty"`
}

type EmailSearchUserResponse struct {
	EmailSearchResult string                `json:"emailSearchResult,omitempty"`
	User              *UsersResponse `json:"user,omitempty"`
}

type GroupControllerApiListGroupMembersOpts struct {
	CompanyName optional.String
	Email       optional.String
	UserName    optional.String
	Page        optional.Int32
	Size        optional.Int32
	Sort        optional.Interface
}

type GroupControllerApiListGroupPolicysOpts struct {
	PolicyName optional.String
	PolicyType optional.String
	Page       optional.Int32
	Size       optional.Int32
	Sort       optional.Interface
}

type GroupControllerApiListGroupsOpts struct {
	GroupName       optional.String
	ModifiedByEmail optional.String
	Page            optional.Int32
	Size            optional.Int32
	Sort            optional.Interface
}

type GroupCreateRequest struct {
	GroupName   string `json:"groupName"`
	Description string `json:"description,omitempty"`
}

type GroupLinkResponse struct {
	GroupId     string `json:"groupId,omitempty"`
	GroupName   string `json:"groupName,omitempty"`
	GroupType   string `json:"groupType,omitempty"`
	UserGroupId string `json:"userGroupId,omitempty"`
}

type GroupMembersAddRequest struct {
	UserIds []string `json:"userIds"`
}

type GroupMembersRemoveRequest struct {
	UserGroupIds []string `json:"userGroupIds"`
}

type GroupMembersResponse struct {
	CompanyName     string       `json:"companyName,omitempty"`
	Email           string       `json:"email,omitempty"`
	LastAccessDt    time.Time `json:"lastAccessDt,omitempty"`
	UserGroupId     string       `json:"userGroupId,omitempty"`
	UserId          string       `json:"userId,omitempty"`
	UserName        string       `json:"userName,omitempty"`
	CreatedBy       string       `json:"createdBy"`
	CreatedByName   string       `json:"createdByName"`
	CreatedByEmail  string       `json:"createdByEmail"`
	CreatedDt       time.Time `json:"createdDt"`
	ModifiedBy      string       `json:"modifiedBy"`
	ModifiedByName  string       `json:"modifiedByName"`
	ModifiedByEmail string       `json:"modifiedByEmail"`
	ModifiedDt      time.Time `json:"modifiedDt"`
}

type GroupPolicysAddRequest struct {
	PolicyIds []string `json:"policyIds"`
}

type GroupPolicysRemoveRequest struct {
	PrincipalPolicyIds []string `json:"principalPolicyIds"`
}

type GroupPolicysResponse struct {
	PolicyId          string       `json:"policyId,omitempty"`
	PolicyName        string       `json:"policyName,omitempty"`
	PolicyType        string       `json:"policyType,omitempty"`
	PrincipalPolicyId string       `json:"principalPolicyId,omitempty"`
	Description       string       `json:"description,omitempty"`
	CreatedBy         string       `json:"createdBy"`
	CreatedByName     string       `json:"createdByName"`
	CreatedByEmail    string       `json:"createdByEmail"`
	CreatedDt         time.Time `json:"createdDt"`
	ModifiedBy        string       `json:"modifiedBy"`
	ModifiedByName    string       `json:"modifiedByName"`
	ModifiedByEmail   string       `json:"modifiedByEmail"`
	ModifiedDt        time.Time `json:"modifiedDt"`
}

type GroupResponse struct {
	ProjectId        string       `json:"projectId,omitempty"`
	GroupId          string       `json:"groupId,omitempty"`
	GroupName        string       `json:"groupName,omitempty"`
	GroupPolicyCount int32        `json:"groupPolicyCount,omitempty"`
	GroupSrn         string       `json:"groupSrn,omitempty"`
	GroupType        string       `json:"groupType,omitempty"`
	GroupUserCount   int32        `json:"groupUserCount,omitempty"`
	Description      string       `json:"description,omitempty"`
	CreatedBy        string       `json:"createdBy"`
	CreatedByName    string       `json:"createdByName"`
	CreatedByEmail   string       `json:"createdByEmail"`
	CreatedDt        time.Time `json:"createdDt"`
	ModifiedBy       string       `json:"modifiedBy"`
	ModifiedByName   string       `json:"modifiedByName"`
	ModifiedByEmail  string       `json:"modifiedByEmail"`
	ModifiedDt       time.Time `json:"modifiedDt"`
}

type GroupUpdateRequest struct {
	GroupName   string `json:"groupName"`
	Description string `json:"description,omitempty"`
}

type GroupsResponse struct {
	ProjectId       string       `json:"projectId,omitempty"`
	GroupId         string       `json:"groupId,omitempty"`
	GroupName       string       `json:"groupName,omitempty"`
	GroupType       string       `json:"groupType,omitempty"`
	GroupUserCount  int32        `json:"groupUserCount,omitempty"`
	Description     string       `json:"description,omitempty"`
	CreatedBy       string       `json:"createdBy"`
	CreatedByName   string       `json:"createdByName"`
	CreatedByEmail  string       `json:"createdByEmail"`
	CreatedDt       time.Time `json:"createdDt"`
	ModifiedBy      string       `json:"modifiedBy"`
	ModifiedByName  string       `json:"modifiedByName"`
	ModifiedByEmail string       `json:"modifiedByEmail"`
	ModifiedDt      time.Time `json:"modifiedDt"`
}

type ListMemberRequest struct {
	CompanyName string
	Email       string
	UserName    string
}

type ListPolicyRequest struct {
	PolicyName string
	PolicyType string
}

type ListResponseV2GroupLinkResponse struct {
	Contents   []GroupLinkResponse `json:"contents,omitempty"`
	TotalCount int64                      `json:"totalCount,omitempty"`
}

type ListResponseV2ProjectMemberCountResponse struct {
	Contents   []ProjectMemberCountResponse `json:"contents,omitempty"`
	TotalCount int64                               `json:"totalCount,omitempty"`
}

type MemberControllerApiListMemberGroupsOpts struct {
	GroupName optional.String
	Page      optional.Int32
	Size      optional.Int32
	Sort      optional.Interface
}

type MemberControllerApiListMembersOpts struct {
	CompanyName optional.String
	Email       optional.String
	UserName    optional.String
	Page        optional.Int32
	Size        optional.Int32
	Sort        optional.Interface
}

type MemberGroupsAddRequest struct {
	GroupIds []string `json:"groupIds"`
}

type MemberGroupsRemoveRequest struct {
	UserGroupIds []string `json:"userGroupIds"`
}

type MemberGroupsResponse struct {
	GroupId         string       `json:"groupId,omitempty"`
	GroupName       string       `json:"groupName,omitempty"`
	GroupType       string       `json:"groupType,omitempty"`
	UserGroupId     string       `json:"userGroupId,omitempty"`
	CreatedBy       string       `json:"createdBy"`
	CreatedByName   string       `json:"createdByName"`
	CreatedByEmail  string       `json:"createdByEmail"`
	CreatedDt       time.Time `json:"createdDt"`
	ModifiedBy      string       `json:"modifiedBy"`
	ModifiedByName  string       `json:"modifiedByName"`
	ModifiedByEmail string       `json:"modifiedByEmail"`
	ModifiedDt      time.Time `json:"modifiedDt"`
}

type MemberResponse struct {
	ProjectId                   string               `json:"projectId,omitempty"`
	CompanyName                 string               `json:"companyName,omitempty"`
	Email                       string               `json:"email,omitempty"`
	LastAccessDt                time.Time         `json:"lastAccessDt,omitempty"`
	OrganizationId              string               `json:"organizationId,omitempty"`
	PositionName                string               `json:"positionName,omitempty"`
	RegisteredBy                string               `json:"registeredBy,omitempty"`
	RegisteredDt                time.Time         `json:"registeredDt,omitempty"`
	Tags                        []TagResponse `json:"tags,omitempty"`
	UserGroupCount              int32                `json:"userGroupCount,omitempty"`
	UserId                      string               `json:"userId,omitempty"`
	UserName                    string               `json:"userName,omitempty"`
	UserPasswordReuseLimitCount string               `json:"userPasswordReuseLimitCount,omitempty"`
	UserSrn                     string               `json:"userSrn,omitempty"`
	CreatedBy                   string               `json:"createdBy"`
	CreatedByName               string               `json:"createdByName"`
	CreatedByEmail              string               `json:"createdByEmail"`
	CreatedDt                   time.Time         `json:"createdDt"`
	ModifiedBy                  string               `json:"modifiedBy"`
	ModifiedByName              string               `json:"modifiedByName"`
	ModifiedByEmail             string               `json:"modifiedByEmail"`
	ModifiedDt                  time.Time         `json:"modifiedDt"`
}

type MembersAddRequest struct {
	GroupIds   []string            `json:"groupIds"`
	Tags       []TagRequest `json:"tags,omitempty"`
	UserEmails []string            `json:"userEmails"`
}

type MembersRemoveRequest struct {
	UserIds []string `json:"userIds"`
}

type MembersResponse struct {
	CompanyName                 string       `json:"companyName,omitempty"`
	Email                       string       `json:"email,omitempty"`
	LastAccessDt                time.Time `json:"lastAccessDt,omitempty"`
	OrganizationId              string       `json:"organizationId,omitempty"`
	PositionName                string       `json:"positionName,omitempty"`
	UserGroupCount              int32        `json:"userGroupCount,omitempty"`
	UserId                      string       `json:"userId,omitempty"`
	UserName                    string       `json:"userName,omitempty"`
	UserPasswordReuseLimitCount string       `json:"userPasswordReuseLimitCount,omitempty"`
	CreatedBy                   string       `json:"createdBy"`
	CreatedByName               string       `json:"createdByName"`
	CreatedByEmail              string       `json:"createdByEmail"`
	CreatedDt                   time.Time `json:"createdDt"`
	ModifiedBy                  string       `json:"modifiedBy"`
	ModifiedByName              string       `json:"modifiedByName"`
	ModifiedByEmail             string       `json:"modifiedByEmail"`
	ModifiedDt                  time.Time `json:"modifiedDt"`
}

type OidcRoleAccessKeyResponse struct {
	ProjectId      string                         `json:"projectId,omitempty"`
	AssumeRoleUser *AssumeRoleUserResponse `json:"assumeRoleUser,omitempty"`
	Audiences      []string                       `json:"audiences,omitempty"`
	Credential     *CredentialResponse     `json:"credential,omitempty"`
	ProviderId     string                         `json:"providerId,omitempty"`
}

type PageResponseV2AccessKeysResponse struct {
	Contents   []AccessKeysResponse `json:"contents,omitempty"`
	Page       int32                       `json:"page,omitempty"`
	Size       int32                       `json:"size,omitempty"`
	Sort       []string                    `json:"sort,omitempty"`
	TotalCount int64                       `json:"totalCount,omitempty"`
}

type PageResponseV2BlockedAccessKeyProjectResponse struct {
	Contents   []BlockedAccessKeyProjectResponse `json:"contents,omitempty"`
	Page       int32                                    `json:"page,omitempty"`
	Size       int32                                    `json:"size,omitempty"`
	Sort       []string                                 `json:"sort,omitempty"`
	TotalCount int64                                    `json:"totalCount,omitempty"`
}

type PageResponseV2GroupMembersResponse struct {
	Contents   []GroupMembersResponse `json:"contents,omitempty"`
	Page       int32                         `json:"page,omitempty"`
	Size       int32                         `json:"size,omitempty"`
	Sort       []string                      `json:"sort,omitempty"`
	TotalCount int64                         `json:"totalCount,omitempty"`
}

type PageResponseV2GroupPolicysResponse struct {
	Contents   []GroupPolicysResponse `json:"contents,omitempty"`
	Page       int32                         `json:"page,omitempty"`
	Size       int32                         `json:"size,omitempty"`
	Sort       []string                      `json:"sort,omitempty"`
	TotalCount int64                         `json:"totalCount,omitempty"`
}

type PageResponseV2GroupsResponse struct {
	Contents   []GroupsResponse `json:"contents,omitempty"`
	Page       int32                   `json:"page,omitempty"`
	Size       int32                   `json:"size,omitempty"`
	Sort       []string                `json:"sort,omitempty"`
	TotalCount int64                   `json:"totalCount,omitempty"`
}

type PageResponseV2MemberGroupsResponse struct {
	Contents   []MemberGroupsResponse `json:"contents,omitempty"`
	Page       int32                         `json:"page,omitempty"`
	Size       int32                         `json:"size,omitempty"`
	Sort       []string                      `json:"sort,omitempty"`
	TotalCount int64                         `json:"totalCount,omitempty"`
}

type PageResponseV2MembersResponse struct {
	Contents   []MembersResponse `json:"contents,omitempty"`
	Page       int32                    `json:"page,omitempty"`
	Size       int32                    `json:"size,omitempty"`
	Sort       []string                 `json:"sort,omitempty"`
	TotalCount int64                    `json:"totalCount,omitempty"`
}

type PageResponseV2PolicyGroupsResponse struct {
	Contents   []PolicyGroupsResponse `json:"contents,omitempty"`
	Page       int32                         `json:"page,omitempty"`
	Size       int32                         `json:"size,omitempty"`
	Sort       []string                      `json:"sort,omitempty"`
	TotalCount int64                         `json:"totalCount,omitempty"`
}

type PageResponseV2PolicyRolesResponse struct {
	Contents   []PolicyRolesResponse `json:"contents,omitempty"`
	Page       int32                        `json:"page,omitempty"`
	Size       int32                        `json:"size,omitempty"`
	Sort       []string                     `json:"sort,omitempty"`
	TotalCount int64                        `json:"totalCount,omitempty"`
}

type PageResponseV2PolicysResponse struct {
	Contents   []PolicysResponse `json:"contents,omitempty"`
	Page       int32                    `json:"page,omitempty"`
	Size       int32                    `json:"size,omitempty"`
	Sort       []string                 `json:"sort,omitempty"`
	TotalCount int64                    `json:"totalCount,omitempty"`
}

type PageResponseV2ProjectsAccessKeysResponse struct {
	Contents   []ProjectsAccessKeysResponse `json:"contents,omitempty"`
	Page       int32                               `json:"page,omitempty"`
	Size       int32                               `json:"size,omitempty"`
	Sort       []string                            `json:"sort,omitempty"`
	TotalCount int64                               `json:"totalCount,omitempty"`
}

type PageResponseV2RolePolicysResponse struct {
	Contents   []RolePolicysResponse `json:"contents,omitempty"`
	Page       int32                        `json:"page,omitempty"`
	Size       int32                        `json:"size,omitempty"`
	Sort       []string                     `json:"sort,omitempty"`
	TotalCount int64                        `json:"totalCount,omitempty"`
}

type PageResponseV2RolesResponse struct {
	Contents   []RolesResponse `json:"contents,omitempty"`
	Page       int32                  `json:"page,omitempty"`
	Size       int32                  `json:"size,omitempty"`
	Sort       []string               `json:"sort,omitempty"`
	TotalCount int64                  `json:"totalCount,omitempty"`
}

type PageResponseV2SamlProvidersResponse struct {
	Contents   []SamlProvidersResponse `json:"contents,omitempty"`
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

type PolicyControllerApiListPolicyGroupsOpts struct {
	GroupName optional.String
	Page      optional.Int32
	Size      optional.Int32
	Sort      optional.Interface
}

type PolicyControllerApiListPolicyRolesOpts struct {
	RoleName optional.String
	Page     optional.Int32
	Size     optional.Int32
	Sort     optional.Interface
}

type PolicyControllerApiListPolicysOpts struct {
	ModifiedByEmail optional.String
	PolicyName      optional.String
	Page            optional.Int32
	Size            optional.Int32
	Sort            optional.Interface
}

type PolicyCreateRequest struct {
	PolicyJson  string                          `json:"policyJson"`
	PolicyName  string                          `json:"policyName"`
	Principals  []PolicyPrincipalRequest `json:"principals,omitempty"`
	Tags        []TagRequest             `json:"tags,omitempty"`
	Description string                          `json:"description,omitempty"`
}

type PolicyGroupsResponse struct {
	GroupId           string       `json:"groupId,omitempty"`
	GroupName         string       `json:"groupName,omitempty"`
	GroupType         string       `json:"groupType,omitempty"`
	PrincipalPolicyId string       `json:"principalPolicyId,omitempty"`
	CreatedBy         string       `json:"createdBy"`
	CreatedByName     string       `json:"createdByName"`
	CreatedByEmail    string       `json:"createdByEmail"`
	CreatedDt         time.Time `json:"createdDt"`
	ModifiedBy        string       `json:"modifiedBy"`
	ModifiedByName    string       `json:"modifiedByName"`
	ModifiedByEmail   string       `json:"modifiedByEmail"`
	ModifiedDt        time.Time `json:"modifiedDt"`
}

type PolicyPrincipalRequest struct {
	PrincipalId   string `json:"principalId"`
	PrincipalType string `json:"principalType"`
}

type PolicyResponse struct {
	ProjectId            string               `json:"projectId,omitempty"`
	PolicyId             string               `json:"policyId,omitempty"`
	PolicyJson           string               `json:"policyJson,omitempty"`
	PolicyName           string               `json:"policyName,omitempty"`
	PolicyPrincipalCount int32                `json:"policyPrincipalCount,omitempty"`
	PolicySrn            string               `json:"policySrn,omitempty"`
	PolicyType           string               `json:"policyType,omitempty"`
	PolicyVersion        string               `json:"policyVersion,omitempty"`
	Tags                 []TagResponse `json:"tags,omitempty"`
	Description          string               `json:"description,omitempty"`
	CreatedBy            string               `json:"createdBy"`
	CreatedByName        string               `json:"createdByName"`
	CreatedByEmail       string               `json:"createdByEmail"`
	CreatedDt            time.Time         `json:"createdDt"`
	ModifiedBy           string               `json:"modifiedBy"`
	ModifiedByName       string               `json:"modifiedByName"`
	ModifiedByEmail      string               `json:"modifiedByEmail"`
	ModifiedDt           time.Time         `json:"modifiedDt"`
}

type PolicyRolesResponse struct {
	PrincipalPolicyId string       `json:"principalPolicyId,omitempty"`
	RoleId            string       `json:"roleId,omitempty"`
	RoleName          string       `json:"roleName,omitempty"`
	Description       string       `json:"description,omitempty"`
	CreatedBy         string       `json:"createdBy"`
	CreatedByName     string       `json:"createdByName"`
	CreatedByEmail    string       `json:"createdByEmail"`
	CreatedDt         time.Time `json:"createdDt"`
	ModifiedBy        string       `json:"modifiedBy"`
	ModifiedByName    string       `json:"modifiedByName"`
	ModifiedByEmail   string       `json:"modifiedByEmail"`
	ModifiedDt        time.Time `json:"modifiedDt"`
}

type PolicyUpdateRequest struct {
	PolicyJson  string                          `json:"policyJson"`
	PolicyName  string                          `json:"policyName"`
	Principals  []PolicyPrincipalRequest `json:"principals,omitempty"`
	Description string                          `json:"description,omitempty"`
}

type PolicyValidationError struct {
	PolicyValidationErrorLocation []string `json:"policyValidationErrorLocation,omitempty"`
	PolicyValidationErrorMessage  string   `json:"policyValidationErrorMessage,omitempty"`
}

type PolicyValidationRequest struct {
	PolicyJson string `json:"policyJson"`
}

type PolicyValidityResponse struct {
	PolicyValidationErrors []PolicyValidationError `json:"policyValidationErrors,omitempty"`
}

type PolicysResponse struct {
	PolicyId        string       `json:"policyId,omitempty"`
	PolicyName      string       `json:"policyName,omitempty"`
	PolicyType      string       `json:"policyType,omitempty"`
	Description     string       `json:"description,omitempty"`
	CreatedBy       string       `json:"createdBy"`
	CreatedByName   string       `json:"createdByName"`
	CreatedByEmail  string       `json:"createdByEmail"`
	CreatedDt       time.Time `json:"createdDt"`
	ModifiedBy      string       `json:"modifiedBy"`
	ModifiedByName  string       `json:"modifiedByName"`
	ModifiedByEmail string       `json:"modifiedByEmail"`
	ModifiedDt      time.Time `json:"modifiedDt"`
}

type ProjectControllerApiListProjectsAccessKeys1Opts struct {
	ParentId optional.String
	Page     optional.Int32
	Size     optional.Int32
	Sort     optional.Interface
}

type ProjectMemberCountResponse struct {
	ProjectId   string `json:"projectId,omitempty"`
	MemberCount int32  `json:"memberCount,omitempty"`
}

type ProjectsAccessKeysResponse struct {
	AccessKey            string       `json:"accessKey,omitempty"`
	AccessKeyActivated   *bool        `json:"accessKeyActivated,omitempty"`
	AccessKeyId          string       `json:"accessKeyId,omitempty"`
	AccessKeyState       string       `json:"accessKeyState,omitempty"`
	AccessKeyType        string       `json:"accessKeyType,omitempty"`
	ExpiredDt            time.Time `json:"expiredDt,omitempty"`
	IssuerType           string       `json:"issuerType,omitempty"`
	LastAccessDt         time.Time `json:"lastAccessDt,omitempty"`
	ProjectAccessKeyType string       `json:"projectAccessKeyType,omitempty"`
	SecretVaultCount     int32        `json:"secretVaultCount,omitempty"`
	Description          string       `json:"description,omitempty"`
	CreatedBy            string       `json:"createdBy"`
	CreatedByName        string       `json:"createdByName"`
	CreatedByEmail       string       `json:"createdByEmail"`
	CreatedDt            time.Time `json:"createdDt"`
	ModifiedBy           string       `json:"modifiedBy"`
	ModifiedByName       string       `json:"modifiedByName"`
	ModifiedByEmail      string       `json:"modifiedByEmail"`
	ModifiedDt           time.Time `json:"modifiedDt"`
}

type ResultResponse struct {
	Result string `json:"result,omitempty"`
}

type RoleAccessKeyResponse struct {
	ProjectId         string       `json:"projectId,omitempty"`
	AccessKey         string       `json:"accessKey,omitempty"`
	AccessKeyId       string       `json:"accessKeyId,omitempty"`
	AccessSecretKey   string       `json:"accessSecretKey,omitempty"`
	ExpiredDt         time.Time `json:"expiredDt,omitempty"`
	FederationType    string       `json:"federationType,omitempty"`
	LastAccessDt      time.Time `json:"lastAccessDt,omitempty"`
	PrincipalIdentity string       `json:"principalIdentity,omitempty"`
	ProviderId        string       `json:"providerId,omitempty"`
	RoleId            string       `json:"roleId,omitempty"`
}

type RoleControllerApiListRolePolicysOpts struct {
	PolicyName optional.String
	PolicyType optional.String
	Page       optional.Int32
	Size       optional.Int32
	Sort       optional.Interface
}

type RoleControllerApiListRolesOpts struct {
	ModifiedByEmail           optional.String
	RoleName                  optional.String
	TrustPrincipalServiceName optional.String
	Page                      optional.Int32
	Size                      optional.Int32
	Sort                      optional.Interface
}

type RoleCreateRequest struct {
	RoleName        string                          `json:"roleName"`
	Tags            []TagRequest             `json:"tags,omitempty"`
	TrustPrincipals *TrustPrincipalsResponse `json:"trustPrincipals,omitempty"`
	Description     string                          `json:"description,omitempty"`
}

type RolePolicysAddRequest struct {
	PolicyIds []string `json:"policyIds"`
}

type RolePolicysRemoveRequest struct {
	PolicyIds []string `json:"policyIds"`
}

type RolePolicysResponse struct {
	PolicyId        string       `json:"policyId,omitempty"`
	PolicyName      string       `json:"policyName,omitempty"`
	PolicyType      string       `json:"policyType,omitempty"`
	Description     string       `json:"description,omitempty"`
	CreatedBy       string       `json:"createdBy"`
	CreatedByName   string       `json:"createdByName"`
	CreatedByEmail  string       `json:"createdByEmail"`
	CreatedDt       time.Time `json:"createdDt"`
	ModifiedBy      string       `json:"modifiedBy"`
	ModifiedByName  string       `json:"modifiedByName"`
	ModifiedByEmail string       `json:"modifiedByEmail"`
	ModifiedDt      time.Time `json:"modifiedDt"`
}

type RoleResponse struct {
	ProjectId       string                          `json:"projectId,omitempty"`
	RoleId          string                          `json:"roleId,omitempty"`
	RoleName        string                          `json:"roleName,omitempty"`
	RolePolicyCount int32                           `json:"rolePolicyCount,omitempty"`
	RoleSrn         string                          `json:"roleSrn,omitempty"`
	SessionTime     int32                           `json:"sessionTime,omitempty"`
	Tags            []TagResponse            `json:"tags,omitempty"`
	TrustPrincipals *TrustPrincipalsResponse `json:"trustPrincipals,omitempty"`
	Description     string                          `json:"description,omitempty"`
	CreatedBy       string                          `json:"createdBy"`
	CreatedByName   string                          `json:"createdByName"`
	CreatedByEmail  string                          `json:"createdByEmail"`
	CreatedDt       time.Time                    `json:"createdDt"`
	ModifiedBy      string                          `json:"modifiedBy"`
	ModifiedByName  string                          `json:"modifiedByName"`
	ModifiedByEmail string                          `json:"modifiedByEmail"`
	ModifiedDt      time.Time                    `json:"modifiedDt"`
}

type RoleUpdateRequest struct {
	RoleName        string                          `json:"roleName"`
	TrustPrincipals *TrustPrincipalsResponse `json:"trustPrincipals,omitempty"`
	Description     string                          `json:"description,omitempty"`
}

type RolesResponse struct {
	ProjectId       string       `json:"projectId,omitempty"`
	RoleId          string       `json:"roleId,omitempty"`
	RoleName        string       `json:"roleName,omitempty"`
	Description     string       `json:"description,omitempty"`
	CreatedBy       string       `json:"createdBy"`
	CreatedByName   string       `json:"createdByName"`
	CreatedByEmail  string       `json:"createdByEmail"`
	CreatedDt       time.Time `json:"createdDt"`
	ModifiedBy      string       `json:"modifiedBy"`
	ModifiedByName  string       `json:"modifiedByName"`
	ModifiedByEmail string       `json:"modifiedByEmail"`
	ModifiedDt      time.Time `json:"modifiedDt"`
}

type SamlControllerApiListSamlProvidersOpts struct {
	FederationType   optional.String
	SamlProviderName optional.String
	Page             optional.Int32
	Size             optional.Int32
	Sort             optional.Interface
}

type SamlProviderCreateRequest struct {
	IdpMetadataXml         string              `json:"idpMetadataXml"`
	IdpMetadataXmlFileName string              `json:"idpMetadataXmlFileName"`
	SamlProviderName       string              `json:"samlProviderName"`
	Tags                   []TagRequest `json:"tags,omitempty"`
}

type SamlProviderResponse struct {
	ProjectId               string               `json:"projectId,omitempty"`
	FederationType          string               `json:"federationType,omitempty"`
	IdpMetadataFileContents string               `json:"idpMetadataFileContents,omitempty"`
	IdpMetadataFileName     string               `json:"idpMetadataFileName,omitempty"`
	IdpMetadataFileSizeByte int64                `json:"idpMetadataFileSizeByte,omitempty"`
	SamlProviderId          string               `json:"samlProviderId,omitempty"`
	SamlProviderName        string               `json:"samlProviderName,omitempty"`
	Tags                    []TagResponse `json:"tags,omitempty"`
	CreatedBy               string               `json:"createdBy"`
	CreatedDt               time.Time         `json:"createdDt"`
	ModifiedBy              string               `json:"modifiedBy"`
	ModifiedDt              time.Time         `json:"modifiedDt"`
}

type SamlProviderUpdateRequest struct {
	AddOrUpdateTags []TagRequest `json:"addOrUpdateTags,omitempty"`
	DeleteTags      []string            `json:"deleteTags,omitempty"`
}

type SamlProvidersResponse struct {
	ProjectId        string       `json:"projectId,omitempty"`
	FederationType   string       `json:"federationType,omitempty"`
	SamlProviderId   string       `json:"samlProviderId,omitempty"`
	SamlProviderName string       `json:"samlProviderName,omitempty"`
	CreatedBy        string       `json:"createdBy"`
	CreatedDt        time.Time `json:"createdDt"`
	ModifiedBy       string       `json:"modifiedBy"`
	ModifiedDt       time.Time `json:"modifiedDt"`
}

type SecurityInfoResponse struct {
	IpAclActivated *bool        `json:"ipAclActivated,omitempty"`
	IpAddresses    []string     `json:"ipAddresses,omitempty"`
	MfaActivated   *bool        `json:"mfaActivated,omitempty"`
	CreatedBy      string       `json:"createdBy"`
	CreatedDt      time.Time `json:"createdDt"`
	ModifiedBy     string       `json:"modifiedBy"`
	ModifiedDt     time.Time `json:"modifiedDt"`
}

type SecurityInfoUpdateRequest struct {
	IpAclActivated *bool    `json:"ipAclActivated"`
	IpAddresses    []string `json:"ipAddresses,omitempty"`
	MfaActivated   *bool    `json:"mfaActivated"`
}

type TagRequest struct {
	TagKey   string `json:"tagKey"`
	TagValue string `json:"tagValue,omitempty"`
}

type TagResponse struct {
	TagKey   string `json:"tagKey"`
	TagValue string `json:"tagValue,omitempty"`
}

type TemporaryAccessKeyControllerApiCountTemporaryAccessKeyOpts struct {
	ProjectId            optional.String
	AccessKeyProjectType optional.String
	AccessKeyState       optional.String
	ActiveYn             optional.Bool
	IssuerType           optional.String
	ProjectName          optional.String
	Page                 optional.Int32
	Size                 optional.Int32
	Sort                 optional.Interface
}

type TemporaryAccessKeyControllerApiListTemporaryAccessKeys1Opts struct {
	ProjectId            optional.String
	AccessKeyProjectType optional.String
	AccessKeyState       optional.String
	ActiveYn             optional.Bool
	IssuerType           optional.String
	ProjectName          optional.String
	Page                 optional.Int32
	Size                 optional.Int32
	Sort                 optional.Interface
}

type TemporaryAccessKeyControllerApiListTemporaryAccessKeysOpts struct {
	ProjectId            optional.String
	AccessKeyProjectType optional.String
	AccessKeyState       optional.String
	ActiveYn             optional.Bool
	IssuerType           optional.String
	ProjectName          optional.String
	Page                 optional.Int32
	Size                 optional.Int32
	Sort                 optional.Interface
}

type TemporaryAccessKeySecretCreateRequest struct {
	DurationMinutes int32  `json:"durationMinutes"`
	Otp             string `json:"otp"`
	Description     string `json:"description,omitempty"`
}

type TrustPrincipalsResponse struct {
	OidcProviderNames []string `json:"oidcProviderNames,omitempty"`
	ProjectIds        []string `json:"projectIds,omitempty"`
	SamlProviderNames []string `json:"samlProviderNames,omitempty"`
	ServiceNames      []string `json:"serviceNames,omitempty"`
	UserSrns          []string `json:"userSrns,omitempty"`
}

type UsersResponse struct {
	CompanyName     string       `json:"companyName,omitempty"`
	Email           string       `json:"email,omitempty"`
	LastAccessDt    time.Time `json:"lastAccessDt,omitempty"`
	LoginId         string       `json:"loginId,omitempty"`
	UserId          string       `json:"userId,omitempty"`
	UserName        string       `json:"userName,omitempty"`
	CreatedBy       string       `json:"createdBy"`
	CreatedByName   string       `json:"createdByName"`
	CreatedByEmail  string       `json:"createdByEmail"`
	CreatedDt       time.Time `json:"createdDt"`
	ModifiedBy      string       `json:"modifiedBy"`
	ModifiedByName  string       `json:"modifiedByName"`
	ModifiedByEmail string       `json:"modifiedByEmail"`
	ModifiedDt      time.Time `json:"modifiedDt"`
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
