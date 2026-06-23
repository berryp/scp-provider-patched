package firewall2

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

type FirewallChangeEnabledRequest struct {
	IsEnabled *bool `json:"isEnabled"`
}

type FirewallChangeLoggableRequest struct {
	IsLoggable *bool `json:"isLoggable"`
}

type FirewallCreateRuleRequest struct {
	DestinationIpAddresses []string                 `json:"destinationIpAddresses"`
	IsRuleEnabled          *bool                    `json:"isRuleEnabled"`
	RuleAction             string                   `json:"ruleAction"`
	RuleDirection          string                   `json:"ruleDirection"`
	RuleLocationId         string                   `json:"ruleLocationId,omitempty"`
	RuleLocationType       string                   `json:"ruleLocationType"`
	Services               []ServiceVo `json:"services"`
	SourceIpAddresses      []string                 `json:"sourceIpAddresses"`
	RuleDescription        string                   `json:"ruleDescription,omitempty"`
}

type FirewallDetailResponse struct {
	ProjectId                          string       `json:"projectId,omitempty"`
	FirewallId                         string       `json:"firewallId,omitempty"`
	FirewallName                       string       `json:"firewallName,omitempty"`
	FirewallState                      string       `json:"firewallState,omitempty"`
	IsEnabled                          *bool        `json:"isEnabled,omitempty"`
	IsLoggable                         *bool        `json:"isLoggable,omitempty"`
	ObjectId                           string       `json:"objectId,omitempty"`
	ObjectType                         string       `json:"objectType,omitempty"`
	ProjectVpcFirewallRuleCurrentCount int64        `json:"projectVpcFirewallRuleCurrentCount,omitempty"`
	ProjectVpcFirewallRuleMaxCount     int64        `json:"projectVpcFirewallRuleMaxCount,omitempty"`
	RuleCount                          int64        `json:"ruleCount,omitempty"`
	ServiceZoneId                      string       `json:"serviceZoneId,omitempty"`
	VpcId                              string       `json:"vpcId,omitempty"`
	CreatedBy                          string       `json:"createdBy,omitempty"`
	CreatedDt                          time.Time `json:"createdDt,omitempty"`
	ModifiedBy                         string       `json:"modifiedBy,omitempty"`
	ModifiedDt                         time.Time `json:"modifiedDt,omitempty"`
}

type FirewallListItemResponse struct {
	ProjectId     string       `json:"projectId,omitempty"`
	FirewallId    string       `json:"firewallId,omitempty"`
	FirewallName  string       `json:"firewallName,omitempty"`
	FirewallState string       `json:"firewallState,omitempty"`
	IsLoggable    *bool        `json:"isLoggable,omitempty"`
	ObjectId      string       `json:"objectId,omitempty"`
	ObjectType    string       `json:"objectType,omitempty"`
	RuleCount     int64        `json:"ruleCount,omitempty"`
	ServiceZoneId string       `json:"serviceZoneId,omitempty"`
	VpcId         string       `json:"vpcId,omitempty"`
	CreatedBy     string       `json:"createdBy,omitempty"`
	CreatedDt     time.Time `json:"createdDt,omitempty"`
	ModifiedBy    string       `json:"modifiedBy,omitempty"`
	ModifiedDt    time.Time `json:"modifiedDt,omitempty"`
}

type FirewallLogStorageCreatRequest struct {
	LogStorageType string `json:"logStorageType"`
	ObsBucketId    string `json:"obsBucketId"`
	VpcId          string `json:"vpcId"`
}

type FirewallLogStorageDetailResponse struct {
	ProjectId      string       `json:"projectId,omitempty"`
	LogStorageId   string       `json:"logStorageId"`
	LogStorageType string       `json:"logStorageType"`
	ObsBucketId    string       `json:"obsBucketId"`
	ServiceZoneId  string       `json:"serviceZoneId,omitempty"`
	VpcId          string       `json:"vpcId"`
	CreatedBy      string       `json:"createdBy,omitempty"`
	CreatedDt      time.Time `json:"createdDt,omitempty"`
	ModifiedBy     string       `json:"modifiedBy,omitempty"`
	ModifiedDt     time.Time `json:"modifiedDt,omitempty"`
}

type FirewallLogStorageUpdateRequest struct {
	ObsBucketId string `json:"obsBucketId"`
}

type FirewallLogStorageV2ApiListFirewallLogStoragesV2Opts struct {
	LogStorageType optional.String
	ObsBucketId    optional.String
}

type FirewallRuleChangeEnabledRequest struct {
	RuleEnabledChangeType string   `json:"ruleEnabledChangeType"`
	RuleIds               []string `json:"ruleIds,omitempty"`
}

type FirewallRuleChangeLocationRequest struct {
	LocationRuleId   string `json:"locationRuleId,omitempty"`
	RuleLocationType string `json:"ruleLocationType"`
}

type FirewallRuleCreateBulkRequest struct {
	BulkRuleLocationId   string                                   `json:"bulkRuleLocationId,omitempty"`
	BulkRuleLocationType string                                   `json:"bulkRuleLocationType"`
	BulkRules            []FirewallCreateRuleRequest `json:"bulkRules"`
}

type FirewallRuleDeleteRequest struct {
	RuleDeletionType string   `json:"ruleDeletionType"`
	RuleIds          []string `json:"ruleIds,omitempty"`
}

type FirewallRuleDetailResponse struct {
	ProjectId              string       `json:"projectId,omitempty"`
	DestinationIpAddresses []string     `json:"destinationIpAddresses,omitempty"`
	IcmpServices           []string     `json:"icmpServices,omitempty"`
	IsAllService           *bool        `json:"isAllService,omitempty"`
	IsRuleEnabled          *bool        `json:"isRuleEnabled,omitempty"`
	Revision               int32        `json:"revision,omitempty"`
	RuleAction             string       `json:"ruleAction,omitempty"`
	RuleDirection          string       `json:"ruleDirection,omitempty"`
	RuleId                 string       `json:"ruleId,omitempty"`
	RuleOwnerId            string       `json:"ruleOwnerId,omitempty"`
	RuleOwnerType          string       `json:"ruleOwnerType,omitempty"`
	RuleState              string       `json:"ruleState,omitempty"`
	Sequence               int32        `json:"sequence,omitempty"`
	ServiceZoneId          string       `json:"serviceZoneId,omitempty"`
	SourceIpAddresses      []string     `json:"sourceIpAddresses,omitempty"`
	TcpServices            []string     `json:"tcpServices,omitempty"`
	UdpServices            []string     `json:"udpServices,omitempty"`
	VendorObjectId         string       `json:"vendorObjectId,omitempty"`
	VendorRuleId           string       `json:"vendorRuleId,omitempty"`
	RuleDescription        string       `json:"ruleDescription,omitempty"`
	CreatedBy              string       `json:"createdBy,omitempty"`
	CreatedDt              time.Time `json:"createdDt,omitempty"`
	ModifiedBy             string       `json:"modifiedBy,omitempty"`
	ModifiedDt             time.Time `json:"modifiedDt,omitempty"`
}

type FirewallRuleListItemResponse struct {
	ProjectId              string       `json:"projectId,omitempty"`
	DestinationIpAddresses []string     `json:"destinationIpAddresses,omitempty"`
	IcmpServices           []string     `json:"icmpServices,omitempty"`
	IsAllService           *bool        `json:"isAllService,omitempty"`
	IsRuleEnabled          *bool        `json:"isRuleEnabled,omitempty"`
	RuleAction             string       `json:"ruleAction,omitempty"`
	RuleDirection          string       `json:"ruleDirection,omitempty"`
	RuleId                 string       `json:"ruleId,omitempty"`
	RuleOwnerId            string       `json:"ruleOwnerId,omitempty"`
	RuleOwnerType          string       `json:"ruleOwnerType,omitempty"`
	RuleState              string       `json:"ruleState,omitempty"`
	Sequence               int32        `json:"sequence,omitempty"`
	ServiceZoneId          string       `json:"serviceZoneId,omitempty"`
	SourceIpAddresses      []string     `json:"sourceIpAddresses,omitempty"`
	TcpServices            []string     `json:"tcpServices,omitempty"`
	UdpServices            []string     `json:"udpServices,omitempty"`
	VendorObjectId         string       `json:"vendorObjectId,omitempty"`
	VendorRuleId           string       `json:"vendorRuleId,omitempty"`
	RuleDescription        string       `json:"ruleDescription,omitempty"`
	CreatedBy              string       `json:"createdBy,omitempty"`
	CreatedDt              time.Time `json:"createdDt,omitempty"`
	ModifiedBy             string       `json:"modifiedBy,omitempty"`
	ModifiedDt             time.Time `json:"modifiedDt,omitempty"`
}

type FirewallRuleUpdateRequest struct {
	DestinationIpAddresses []string                 `json:"destinationIpAddresses"`
	RuleAction             string                   `json:"ruleAction"`
	RuleDirection          string                   `json:"ruleDirection"`
	Services               []ServiceVo `json:"services"`
	SourceIpAddresses      []string                 `json:"sourceIpAddresses"`
	RuleDescription        string                   `json:"ruleDescription,omitempty"`
}

type FirewallRuleV2ApiListFirewallRulesV2Opts struct {
	DestinationIpAddress optional.String
	IsRuleEnabled        optional.Bool
	RuleDescription      optional.String
	RuleDirections       optional.Interface
	RuleStates           optional.Interface
	SourceIpAddress      optional.String
	CreatedBy            optional.String
	Page                 optional.Int32
	Size                 optional.Int32
	Sort                 optional.Interface
}

type FirewallV2ApiListFirewallsV2Opts struct {
	FirewallName   optional.String
	FirewallStates optional.Interface
	IsLoggable     optional.Bool
	ObjectId       optional.String
	ObjectTypes    optional.Interface
	VpcId          optional.String
	CreatedBy      optional.String
	Page           optional.Int32
	Size           optional.Int32
	Sort           optional.Interface
}

type ListResponseFirewallListItemResponse struct {
	Contents   []FirewallListItemResponse `json:"contents,omitempty"`
	TotalCount int32                                   `json:"totalCount,omitempty"`
}

type ListResponseFirewallLogStorageDetailResponse struct {
	Contents   []FirewallLogStorageDetailResponse `json:"contents,omitempty"`
	TotalCount int32                                           `json:"totalCount,omitempty"`
}

type ListResponseFirewallRuleListItemResponse struct {
	Contents   []FirewallRuleListItemResponse `json:"contents,omitempty"`
	TotalCount int32                                       `json:"totalCount,omitempty"`
}

type ServiceVo struct {
	ServiceType  string `json:"serviceType"`
	ServiceValue string `json:"serviceValue,omitempty"`
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
