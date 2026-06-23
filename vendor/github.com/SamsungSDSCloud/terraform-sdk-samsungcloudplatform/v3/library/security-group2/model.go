package securitygroup2

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

type AttachedObjectResponse struct {
	ObjectId    string `json:"objectId,omitempty"`
	ObjectName  string `json:"objectName,omitempty"`
	ObjectState string `json:"objectState,omitempty"`
	ObjectType  string `json:"objectType,omitempty"`
}

type BasicAuth struct {
	UserName string `json:"userName,omitempty"`
	Password string `json:"password,omitempty"`
}

type CheckResponse struct {
	Result *bool `json:"result,omitempty"`
}

type DetailSecurityGroupResponse struct {
	ProjectId                string                                     `json:"projectId,omitempty"`
	AttachedObjects          []AttachedObjectResponse `json:"attachedObjects,omitempty"`
	IsLoggable               *bool                                      `json:"isLoggable,omitempty"`
	RuleCount                int32                                      `json:"ruleCount,omitempty"`
	Scope                    string                                     `json:"scope,omitempty"`
	SecurityGroupId          string                                     `json:"securityGroupId,omitempty"`
	SecurityGroupName        string                                     `json:"securityGroupName,omitempty"`
	SecurityGroupState       string                                     `json:"securityGroupState,omitempty"`
	VendorObjectId           string                                     `json:"vendorObjectId,omitempty"`
	VpcId                    string                                     `json:"vpcId,omitempty"`
	ZoneId                   string                                     `json:"zoneId,omitempty"`
	SecurityGroupDescription string                                     `json:"securityGroupDescription,omitempty"`
	CreatedBy                string                                     `json:"createdBy,omitempty"`
	CreatedDt                time.Time                               `json:"createdDt,omitempty"`
	ModifiedBy               string                                     `json:"modifiedBy,omitempty"`
	ModifiedDt               time.Time                               `json:"modifiedDt,omitempty"`
}

type ListResponseSecurityGroupLogStorageDetailResponse struct {
	Contents   []SecurityGroupLogStorageDetailResponse `json:"contents,omitempty"`
	TotalCount int32                                                     `json:"totalCount,omitempty"`
}

type ListResponseSecurityGroupResponse struct {
	Contents   []SecurityGroupResponse `json:"contents,omitempty"`
	TotalCount int32                                     `json:"totalCount,omitempty"`
}

type ListResponseSecurityGroupRuleResponse struct {
	Contents   []SecurityGroupRuleResponse `json:"contents,omitempty"`
	TotalCount int32                                         `json:"totalCount,omitempty"`
}

type ListResponseSecurityGroupUserIpResponse struct {
	Contents   []SecurityGroupUserIpResponse `json:"contents,omitempty"`
	TotalCount int32                                           `json:"totalCount,omitempty"`
}

type SecurityGroupChangeLoggableRequest struct {
	IsLoggable *bool `json:"isLoggable"`
}

type SecurityGroupCreateBulkRuleRequest struct {
	Rules []SecurityGroupCreateRuleRequest `json:"rules"`
}

type SecurityGroupCreateRuleRequest struct {
	DestinationIpAddresses []string                      `json:"destinationIpAddresses,omitempty"`
	RuleDirection          string                        `json:"ruleDirection"`
	Services               []ServiceVo `json:"services"`
	SourceIpAddresses      []string                      `json:"sourceIpAddresses,omitempty"`
	RuleDescription        string                        `json:"ruleDescription,omitempty"`
}

type SecurityGroupCreateV3Request struct {
	Loggable                 *bool                          `json:"loggable"`
	SecurityGroupName        string                         `json:"securityGroupName"`
	ServiceZoneId            string                         `json:"serviceZoneId"`
	Tags                     []TagRequest `json:"tags,omitempty"`
	VpcId                    string                         `json:"vpcId"`
	SecurityGroupDescription string                         `json:"securityGroupDescription,omitempty"`
}

type SecurityGroupLogStorageCreateRequest struct {
	ObsBucketId string `json:"obsBucketId"`
	VpcId       string `json:"vpcId"`
}

type SecurityGroupLogStorageDetailResponse struct {
	ProjectId      string       `json:"projectId,omitempty"`
	LogStorageId   string       `json:"logStorageId"`
	LogStorageType string       `json:"logStorageType"`
	ObsBucketId    string       `json:"obsBucketId"`
	VpcId          string       `json:"vpcId"`
	CreatedBy      string       `json:"createdBy,omitempty"`
	CreatedDt      time.Time `json:"createdDt,omitempty"`
	ModifiedBy     string       `json:"modifiedBy,omitempty"`
	ModifiedDt     time.Time `json:"modifiedDt,omitempty"`
}

type SecurityGroupLogStorageOpenApiControllerV2ApiListSecurityGroupLogStoragesV2Opts struct {
	ObsBucketId optional.String
}

type SecurityGroupLogStorageUpdateRequest struct {
	ObsBucketId string `json:"obsBucketId"`
}

type SecurityGroupModifyDescriptionRequest struct {
	SecurityGroupDescription string `json:"securityGroupDescription,omitempty"`
}

type SecurityGroupOpenApiControllerV2ApiListSecurityGroupRuleV2Opts struct {
	IpAddress       optional.String
	RuleDescription optional.String
	RuleDirection   optional.String
	RuleStates      optional.Interface
	Page            optional.Int32
	Size            optional.Int32
	Sort            optional.Interface
}

type SecurityGroupOpenApiControllerV2ApiListSecurityGroupV2Opts struct {
	IsLoggable          optional.Bool
	Scopes              optional.Interface
	SecurityGroupName   optional.String
	SecurityGroupStates optional.Interface
	VpcId               optional.String
	CreatedBy           optional.String
	Page                optional.Int32
	Size                optional.Int32
	Sort                optional.Interface
}

type SecurityGroupResponse struct {
	AttachedObjectCount int32        `json:"attachedObjectCount,omitempty"`
	IsLoggable          *bool        `json:"isLoggable,omitempty"`
	RuleCount           int32        `json:"ruleCount,omitempty"`
	Scope               string       `json:"scope,omitempty"`
	SecurityGroupId     string       `json:"securityGroupId,omitempty"`
	SecurityGroupName   string       `json:"securityGroupName,omitempty"`
	SecurityGroupState  string       `json:"securityGroupState,omitempty"`
	ZoneId              string       `json:"zoneId,omitempty"`
	CreatedBy           string       `json:"createdBy,omitempty"`
	CreatedDt           time.Time `json:"createdDt,omitempty"`
	ModifiedBy          string       `json:"modifiedBy,omitempty"`
	ModifiedDt          time.Time `json:"modifiedDt,omitempty"`
}

type SecurityGroupRuleDeleteRequest struct {
	RuleDeletionType string   `json:"ruleDeletionType,omitempty"`
	RuleIds          []string `json:"ruleIds,omitempty"`
}

type SecurityGroupRuleModifyRequest struct {
	DestinationIpAddresses []string                      `json:"destinationIpAddresses,omitempty"`
	RuleDirection          string                        `json:"ruleDirection"`
	Services               []ServiceVo `json:"services"`
	SourceIpAddresses      []string                      `json:"sourceIpAddresses,omitempty"`
	RuleDescription        string                        `json:"ruleDescription,omitempty"`
}

type SecurityGroupRuleResponse struct {
	IcmpServices    []string     `json:"icmpServices,omitempty"`
	IsAllService    *bool        `json:"isAllService,omitempty"`
	RuleAction      string       `json:"ruleAction,omitempty"`
	RuleDirection   string       `json:"ruleDirection,omitempty"`
	RuleId          string       `json:"ruleId,omitempty"`
	RuleOwnerId     string       `json:"ruleOwnerId,omitempty"`
	RuleOwnerType   string       `json:"ruleOwnerType,omitempty"`
	RuleState       string       `json:"ruleState,omitempty"`
	TargetNetworks  []string     `json:"targetNetworks,omitempty"`
	TcpServices     []string     `json:"tcpServices,omitempty"`
	UdpServices     []string     `json:"udpServices,omitempty"`
	VendorObjectId  string       `json:"vendorObjectId,omitempty"`
	VendorRuleId    int32        `json:"vendorRuleId,omitempty"`
	RuleDescription string       `json:"ruleDescription,omitempty"`
	CreatedBy       string       `json:"createdBy,omitempty"`
	CreatedDt       time.Time `json:"createdDt,omitempty"`
	ModifiedBy      string       `json:"modifiedBy,omitempty"`
	ModifiedDt      time.Time `json:"modifiedDt,omitempty"`
}

type SecurityGroupUserIpAttachRequest struct {
	UserIpAddress     string `json:"userIpAddress"`
	UserIpType        string `json:"userIpType"`
	UserIpDescription string `json:"userIpDescription"`
}

type SecurityGroupUserIpDetachRequest struct {
	UserIpAddress string `json:"userIpAddress"`
}

type SecurityGroupUserIpResponse struct {
	State             string `json:"state,omitempty"`
	UserIpAddress     string `json:"userIpAddress,omitempty"`
	UserIpType        string `json:"userIpType,omitempty"`
	UserIpDescription string `json:"userIpDescription,omitempty"`
}

type ServiceVo struct {
	ServiceType  string `json:"serviceType"`
	ServiceValue string `json:"serviceValue,omitempty"`
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
