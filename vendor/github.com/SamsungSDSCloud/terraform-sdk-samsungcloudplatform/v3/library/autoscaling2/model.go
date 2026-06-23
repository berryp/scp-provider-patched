package autoscaling2

import (
	"github.com/antihax/optional"
	"time"
)

type APIKey struct {
	Key    string
	Prefix string
}

type AsgDrainingDetailResponse struct {
	AsgDrainingEnabled       *bool  `json:"asgDrainingEnabled,omitempty"`
	AsgDrainingTimeoutSecond *int32 `json:"asgDrainingTimeoutSecond,omitempty"`
}

type AsgDrainingUpdateRequest struct {
	AsgDrainingEnabled       *bool  `json:"asgDrainingEnabled"`
	AsgDrainingTimeoutSecond *int32 `json:"asgDrainingTimeoutSecond,omitempty"`
}

type AsgEventDetailResponse struct {
	ProjectId        string                                    `json:"projectId,omitempty"`
	ActionHistories  []AsgEventHistoryResponse `json:"actionHistories,omitempty"`
	AsgId            string                                    `json:"asgId,omitempty"`
	BlockId          string                                    `json:"blockId,omitempty"`
	ChildEventIds    []string                                  `json:"childEventIds,omitempty"`
	EventAfterValue  string                                    `json:"eventAfterValue,omitempty"`
	EventBeforeValue string                                    `json:"eventBeforeValue,omitempty"`
	EventId          string                                    `json:"eventId,omitempty"`
	EventName        string                                    `json:"eventName,omitempty"`
	EventState       string                                    `json:"eventState,omitempty"`
	EventType        string                                    `json:"eventType,omitempty"`
	FailCount        *int32                                    `json:"failCount,omitempty"`
	ParentEventId    string                                    `json:"parentEventId,omitempty"`
	RequestEventId   string                                    `json:"requestEventId,omitempty"`
	ResultEventId    string                                    `json:"resultEventId,omitempty"`
	ServiceZoneId    string                                    `json:"serviceZoneId,omitempty"`
	CreatedBy        string                                    `json:"createdBy,omitempty"`
	CreatedDt        time.Time                              `json:"createdDt,omitempty"`
	ModifiedBy       string                                    `json:"modifiedBy,omitempty"`
	ModifiedDt       time.Time                              `json:"modifiedDt,omitempty"`
}

type AsgEventHistoryResponse struct {
	ProjectId            string       `json:"projectId,omitempty"`
	ActionName           string       `json:"actionName,omitempty"`
	ActionObjectId       string       `json:"actionObjectId,omitempty"`
	ActionObjectName     string       `json:"actionObjectName,omitempty"`
	ActionState          string       `json:"actionState,omitempty"`
	AvailabilityZoneName string       `json:"availabilityZoneName,omitempty"`
	BlockId              string       `json:"blockId,omitempty"`
	FailCount            *int32       `json:"failCount,omitempty"`
	ServiceZoneId        string       `json:"serviceZoneId,omitempty"`
	Description          string       `json:"description,omitempty"`
	CreatedBy            string       `json:"createdBy,omitempty"`
	CreatedDt            time.Time `json:"createdDt,omitempty"`
	ModifiedBy           string       `json:"modifiedBy,omitempty"`
	ModifiedDt           time.Time `json:"modifiedDt,omitempty"`
}

type AsgEventListItemResponse struct {
	ProjectId        string                                    `json:"projectId,omitempty"`
	ActionHistories  []AsgEventHistoryResponse `json:"actionHistories,omitempty"`
	BlockId          string                                    `json:"blockId,omitempty"`
	ChildEventIds    []string                                  `json:"childEventIds,omitempty"`
	EventAfterValue  string                                    `json:"eventAfterValue,omitempty"`
	EventBeforeValue string                                    `json:"eventBeforeValue,omitempty"`
	EventId          string                                    `json:"eventId,omitempty"`
	EventName        string                                    `json:"eventName,omitempty"`
	EventState       string                                    `json:"eventState,omitempty"`
	EventType        string                                    `json:"eventType,omitempty"`
	ParentEventId    string                                    `json:"parentEventId,omitempty"`
	RequestEventId   string                                    `json:"requestEventId,omitempty"`
	ResultEventId    string                                    `json:"resultEventId,omitempty"`
	ServiceZoneId    string                                    `json:"serviceZoneId,omitempty"`
	CreatedBy        string                                    `json:"createdBy,omitempty"`
	CreatedDt        time.Time                              `json:"createdDt,omitempty"`
	ModifiedBy       string                                    `json:"modifiedBy,omitempty"`
	ModifiedDt       time.Time                              `json:"modifiedDt,omitempty"`
}

type AsgEventV2ApiGetAsgEventListV2Opts struct {
	EndDate    optional.String
	EventNames optional.Interface
	EventState optional.String
	EventType  optional.String
	StartDate  optional.String
	Page       optional.Int32
	Size       optional.Int32
	Sort       optional.String
}

type AsgLaunchConfigurationV2ApiGetLaunchConfigListV2Opts struct {
	ImageId       optional.String
	LcName        optional.String
	ServiceZoneId optional.String
	CreatedBy     optional.String
	Page          optional.Int32
	Size          optional.Int32
	Sort          optional.String
}

type AsgLoadBalancersUpdateRequest struct {
	AttachLbRuleIds []string `json:"attachLbRuleIds,omitempty"`
	DetachLbRuleIds []string `json:"detachLbRuleIds,omitempty"`
}

type AsgNotificationCreateRequest struct {
	AsgNotificationPoints []string `json:"asgNotificationPoints"`
	AsgNotificationTypes  []string `json:"asgNotificationTypes"`
	RecipientUserIds      []string `json:"recipientUserIds"`
}

type AsgNotificationResponse struct {
	ProjectId             string       `json:"projectId,omitempty"`
	AsgId                 string       `json:"asgId,omitempty"`
	AsgNotificationId     string       `json:"asgNotificationId,omitempty"`
	AsgNotificationPoints []string     `json:"asgNotificationPoints,omitempty"`
	AsgNotificationState  string       `json:"asgNotificationState,omitempty"`
	AsgNotificationTypes  []string     `json:"asgNotificationTypes,omitempty"`
	BlockId               string       `json:"blockId,omitempty"`
	RecipientDepartment   string       `json:"recipientDepartment,omitempty"`
	RecipientEmail        string       `json:"recipientEmail,omitempty"`
	RecipientName         string       `json:"recipientName,omitempty"`
	RecipientUserId       string       `json:"recipientUserId,omitempty"`
	ServiceZoneId         string       `json:"serviceZoneId,omitempty"`
	CreatedBy             string       `json:"createdBy,omitempty"`
	CreatedDt             time.Time `json:"createdDt,omitempty"`
	ModifiedBy            string       `json:"modifiedBy,omitempty"`
	ModifiedDt            time.Time `json:"modifiedDt,omitempty"`
}

type AsgNotificationUpdateRequest struct {
	AsgNotificationPoints []string `json:"asgNotificationPoints,omitempty"`
	AsgNotificationState  string   `json:"asgNotificationState,omitempty"`
	AsgNotificationTypes  []string `json:"asgNotificationTypes,omitempty"`
}

type AsgNotificationV2ApiGetAsgNotificationListV2Opts struct {
	Page optional.Int32
	Size optional.Int32
	Sort optional.String
}

type AsgPolicyCreateRequest struct {
	ComparisonOperator string `json:"comparisonOperator,omitempty"`
	CooldownSeconds    *int32 `json:"cooldownSeconds,omitempty"`
	EvaluationMinutes  *int32 `json:"evaluationMinutes,omitempty"`
	MetricMethod       string `json:"metricMethod,omitempty"`
	MetricType         string `json:"metricType,omitempty"`
	PolicyName         string `json:"policyName"`
	ScaleMethod        string `json:"scaleMethod,omitempty"`
	ScaleType          string `json:"scaleType,omitempty"`
	ScaleValue         *int32 `json:"scaleValue,omitempty"`
	Threshold          string `json:"threshold,omitempty"`
}

type AsgPolicyResponse struct {
	ProjectId          string       `json:"projectId,omitempty"`
	AsgId              string       `json:"asgId,omitempty"`
	BlockId            string       `json:"blockId,omitempty"`
	ComparisonOperator string       `json:"comparisonOperator,omitempty"`
	CooldownSeconds    *int32       `json:"cooldownSeconds,omitempty"`
	EvaluationMinutes  *int32       `json:"evaluationMinutes,omitempty"`
	MetricMethod       string       `json:"metricMethod,omitempty"`
	MetricType         string       `json:"metricType,omitempty"`
	PolicyId           string       `json:"policyId,omitempty"`
	PolicyName         string       `json:"policyName,omitempty"`
	PolicyState        string       `json:"policyState,omitempty"`
	ScaleMethod        string       `json:"scaleMethod,omitempty"`
	ScaleType          string       `json:"scaleType,omitempty"`
	ScaleValue         *int32       `json:"scaleValue,omitempty"`
	ServiceId          string       `json:"serviceId,omitempty"`
	ServiceZoneId      string       `json:"serviceZoneId,omitempty"`
	Threshold          string       `json:"threshold,omitempty"`
	ThresholdUnit      string       `json:"thresholdUnit,omitempty"`
	CreatedBy          string       `json:"createdBy,omitempty"`
	CreatedDt          time.Time `json:"createdDt,omitempty"`
	ModifiedBy         string       `json:"modifiedBy,omitempty"`
	ModifiedDt         time.Time `json:"modifiedDt,omitempty"`
}

type AsgPolicyUpdateRequest struct {
	ComparisonOperator string `json:"comparisonOperator,omitempty"`
	CooldownSeconds    *int32 `json:"cooldownSeconds,omitempty"`
	EvaluationMinutes  *int32 `json:"evaluationMinutes,omitempty"`
	MetricMethod       string `json:"metricMethod,omitempty"`
	MetricType         string `json:"metricType,omitempty"`
	PolicyName         string `json:"policyName"`
	ScaleMethod        string `json:"scaleMethod,omitempty"`
	ScaleValue         *int32 `json:"scaleValue,omitempty"`
	Threshold          string `json:"threshold,omitempty"`
}

type AsgPolicyV2ApiGetAsgPolicyListV2Opts struct {
	MetricMethod optional.String
	MetricType   optional.String
	PolicyName   optional.String
	ScaleType    optional.String
	Page         optional.Int32
	Size         optional.Int32
	Sort         optional.String
}

type AsgScheduleCreateRequest struct {
	AsgScheduleName       string `json:"asgScheduleName"`
	DesiredServerCount    *int32 `json:"desiredServerCount,omitempty"`
	MaxServerCount        *int32 `json:"maxServerCount,omitempty"`
	MinServerCount        *int32 `json:"minServerCount,omitempty"`
	ScheduleDay           string `json:"scheduleDay,omitempty"`
	ScheduleDayOfWeek     string `json:"scheduleDayOfWeek,omitempty"`
	ScheduleEndDate       string `json:"scheduleEndDate,omitempty"`
	ScheduleFrequencyType string `json:"scheduleFrequencyType"`
	ScheduleHour          *int32 `json:"scheduleHour"`
	ScheduleMinute        *int32 `json:"scheduleMinute"`
	ScheduleStartDate     string `json:"scheduleStartDate"`
	Timezone              string `json:"timezone"`
}

type AsgScheduleDetailResponse struct {
	ProjectId             string       `json:"projectId,omitempty"`
	AsgScheduleId         string       `json:"asgScheduleId,omitempty"`
	AsgScheduleName       string       `json:"asgScheduleName,omitempty"`
	AsgScheduleState      string       `json:"asgScheduleState,omitempty"`
	BlockId               string       `json:"blockId,omitempty"`
	DesiredServerCount    *int32       `json:"desiredServerCount,omitempty"`
	MaxServerCount        *int32       `json:"maxServerCount,omitempty"`
	MinServerCount        *int32       `json:"minServerCount,omitempty"`
	ScheduleDay           string       `json:"scheduleDay,omitempty"`
	ScheduleDayOfWeek     string       `json:"scheduleDayOfWeek,omitempty"`
	ScheduleEndDate       string       `json:"scheduleEndDate,omitempty"`
	ScheduleEndDt         time.Time `json:"scheduleEndDt,omitempty"`
	ScheduleFrequencyType string       `json:"scheduleFrequencyType,omitempty"`
	ScheduleHour          *int32       `json:"scheduleHour,omitempty"`
	ScheduleMinute        *int32       `json:"scheduleMinute,omitempty"`
	ScheduleStartDate     string       `json:"scheduleStartDate,omitempty"`
	ScheduleStartDt       time.Time `json:"scheduleStartDt,omitempty"`
	ServiceZoneId         string       `json:"serviceZoneId,omitempty"`
	Timezone              string       `json:"timezone,omitempty"`
	CreatedBy             string       `json:"createdBy,omitempty"`
	CreatedDt             time.Time `json:"createdDt,omitempty"`
	ModifiedBy            string       `json:"modifiedBy,omitempty"`
	ModifiedDt            time.Time `json:"modifiedDt,omitempty"`
}

type AsgScheduleListItemResponse struct {
	ProjectId             string       `json:"projectId,omitempty"`
	AsgScheduleId         string       `json:"asgScheduleId,omitempty"`
	AsgScheduleName       string       `json:"asgScheduleName,omitempty"`
	AsgScheduleState      string       `json:"asgScheduleState,omitempty"`
	BlockId               string       `json:"blockId,omitempty"`
	DesiredServerCount    *int32       `json:"desiredServerCount,omitempty"`
	MaxServerCount        *int32       `json:"maxServerCount,omitempty"`
	MinServerCount        *int32       `json:"minServerCount,omitempty"`
	ScheduleDay           string       `json:"scheduleDay,omitempty"`
	ScheduleDayOfWeek     string       `json:"scheduleDayOfWeek,omitempty"`
	ScheduleEndDate       string       `json:"scheduleEndDate,omitempty"`
	ScheduleEndDt         time.Time `json:"scheduleEndDt,omitempty"`
	ScheduleFrequencyType string       `json:"scheduleFrequencyType,omitempty"`
	ScheduleHour          *int32       `json:"scheduleHour,omitempty"`
	ScheduleMinute        *int32       `json:"scheduleMinute,omitempty"`
	ScheduleStartDate     string       `json:"scheduleStartDate,omitempty"`
	ScheduleStartDt       time.Time `json:"scheduleStartDt,omitempty"`
	ServiceZoneId         string       `json:"serviceZoneId,omitempty"`
	Timezone              string       `json:"timezone,omitempty"`
	CreatedBy             string       `json:"createdBy,omitempty"`
	CreatedDt             time.Time `json:"createdDt,omitempty"`
	ModifiedBy            string       `json:"modifiedBy,omitempty"`
	ModifiedDt            time.Time `json:"modifiedDt,omitempty"`
}

type AsgScheduleUpdateRequest struct {
	AsgScheduleName       string `json:"asgScheduleName"`
	AsgScheduleState      string `json:"asgScheduleState"`
	DesiredServerCount    *int32 `json:"desiredServerCount,omitempty"`
	MaxServerCount        *int32 `json:"maxServerCount,omitempty"`
	MinServerCount        *int32 `json:"minServerCount,omitempty"`
	ScheduleDay           string `json:"scheduleDay,omitempty"`
	ScheduleDayOfWeek     string `json:"scheduleDayOfWeek,omitempty"`
	ScheduleEndDate       string `json:"scheduleEndDate,omitempty"`
	ScheduleFrequencyType string `json:"scheduleFrequencyType"`
	ScheduleHour          *int32 `json:"scheduleHour"`
	ScheduleMinute        *int32 `json:"scheduleMinute"`
	ScheduleStartDate     string `json:"scheduleStartDate"`
	Timezone              string `json:"timezone"`
}

type AsgScheduleV2ApiGetAsgScheduleListV2Opts struct {
	Page optional.Int32
	Size optional.Int32
	Sort optional.String
}

type AsgServerCountUpdateRequest struct {
	DesiredServerCount *int32 `json:"desiredServerCount,omitempty"`
	MaxServerCount     *int32 `json:"maxServerCount,omitempty"`
	MinServerCount     *int32 `json:"minServerCount,omitempty"`
}

type AsgVirtualServerListItemResponse struct {
	FileStorageLinkedState string            `json:"fileStorageLinkedState,omitempty"`
	ImageId                string            `json:"imageId,omitempty"`
	Ip                     string            `json:"ip,omitempty"`
	Properties             map[string]string `json:"properties,omitempty"`
	ServerGroupId          string            `json:"serverGroupId,omitempty"`
	ServicedFor            string            `json:"servicedFor,omitempty"`
	ServicedGroupFor       string            `json:"servicedGroupFor,omitempty"`
	VirtualServerId        string            `json:"virtualServerId,omitempty"`
	VirtualServerName      string            `json:"virtualServerName,omitempty"`
	VirtualServerState     string            `json:"virtualServerState,omitempty"`
	CreatedBy              string            `json:"createdBy,omitempty"`
	CreatedDt              time.Time      `json:"createdDt,omitempty"`
	ModifiedBy             string            `json:"modifiedBy,omitempty"`
	ModifiedDt             time.Time      `json:"modifiedDt,omitempty"`
}

type AsgVirtualServerUpdateRequest struct {
	LbConnectionState string `json:"lbConnectionState"`
}

type AsgVirtualServerV2ApiGetAsgVirtualServerListV2Opts struct {
	Page optional.Int32
	Size optional.Int32
	Sort optional.Interface
}

type AsgVpcInfoRequest struct {
	LocalSubnetId string `json:"localSubnetId,omitempty"`
	NatEnabled    *bool  `json:"natEnabled"`
	SubnetId      string `json:"subnetId"`
	VpcId         string `json:"vpcId"`
}

type AsgVpcInfoResponse struct {
	LocalSubnetId string `json:"localSubnetId,omitempty"`
	NatEnabled    *bool  `json:"natEnabled,omitempty"`
	SubnetId      string `json:"subnetId,omitempty"`
	VpcId         string `json:"vpcId,omitempty"`
}

type AsyncResponse struct {
	ProjectId  string `json:"projectId,omitempty"`
	RequestId  string `json:"requestId,omitempty"`
	ResourceId string `json:"resourceId,omitempty"`
}

type AutoScalingGroupCreateV4Request struct {
	AsgName                      string                             `json:"asgName"`
	AvailabilityZoneName         string                             `json:"availabilityZoneName,omitempty"`
	DeploymentEnvType            string                             `json:"deploymentEnvType,omitempty"`
	DesiredServerCount           *int32                             `json:"desiredServerCount"`
	DesiredServerCountEditable   *bool                              `json:"desiredServerCountEditable"`
	FileStorageId                string                             `json:"fileStorageId,omitempty"`
	LcId                         string                             `json:"lcId"`
	MaxServerCount               *int32                             `json:"maxServerCount"`
	MinServerCount               *int32                             `json:"minServerCount"`
	MultiAvailabilityZoneEnabled *bool                              `json:"multiAvailabilityZoneEnabled,omitempty"`
	SecurityGroupIds             []string                           `json:"securityGroupIds"`
	ServerNamePrefix             string                             `json:"serverNamePrefix"`
	ServiceLevelProductId        string                             `json:"serviceLevelProductId,omitempty"`
	Tags                         []TagRequest       `json:"tags,omitempty"`
	VpcInfo                      *AsgVpcInfoRequest `json:"vpcInfo"`
}

type AutoScalingGroupResponse struct {
	ProjectId                    string                              `json:"projectId,omitempty"`
	AsgId                        string                              `json:"asgId,omitempty"`
	AsgName                      string                              `json:"asgName,omitempty"`
	AsgState                     string                              `json:"asgState,omitempty"`
	AvailabilityZoneName         string                              `json:"availabilityZoneName,omitempty"`
	BlockId                      string                              `json:"blockId,omitempty"`
	DeploymentEnvType            string                              `json:"deploymentEnvType,omitempty"`
	DesiredServerCount           *int32                              `json:"desiredServerCount,omitempty"`
	DesiredServerCountEditable   *bool                               `json:"desiredServerCountEditable,omitempty"`
	DnsEnabled                   *bool                               `json:"dnsEnabled,omitempty"`
	FileStorageId                string                              `json:"fileStorageId,omitempty"`
	IsTerminating                *bool                               `json:"isTerminating,omitempty"`
	LcId                         string                              `json:"lcId,omitempty"`
	LcName                       string                              `json:"lcName,omitempty"`
	LocalSubnetId                string                              `json:"localSubnetId,omitempty"`
	MaxServerCount               *int32                              `json:"maxServerCount,omitempty"`
	MinServerCount               *int32                              `json:"minServerCount,omitempty"`
	MultiAvailabilityZoneEnabled *bool                               `json:"multiAvailabilityZoneEnabled,omitempty"`
	SecurityGroupIds             []string                            `json:"securityGroupIds,omitempty"`
	ServerNamePrefix             string                              `json:"serverNamePrefix,omitempty"`
	ServiceId                    string                              `json:"serviceId,omitempty"`
	ServiceLevelProductId        string                              `json:"serviceLevelProductId,omitempty"`
	ServiceZoneId                string                              `json:"serviceZoneId,omitempty"`
	SubnetId                     string                              `json:"subnetId,omitempty"`
	VpcId                        string                              `json:"vpcId,omitempty"`
	VpcInfo                      *AsgVpcInfoResponse `json:"vpcInfo,omitempty"`
	CreatedBy                    string                              `json:"createdBy,omitempty"`
	CreatedDt                    time.Time                        `json:"createdDt,omitempty"`
	ModifiedBy                   string                              `json:"modifiedBy,omitempty"`
	ModifiedDt                   time.Time                        `json:"modifiedDt,omitempty"`
}

type AutoScalingGroupUpdateRequest struct {
	DeploymentEnvType          string   `json:"deploymentEnvType,omitempty"`
	DesiredServerCountEditable *bool    `json:"desiredServerCountEditable,omitempty"`
	LcId                       string   `json:"lcId,omitempty"`
	SecurityGroupIds           []string `json:"securityGroupIds,omitempty"`
}

type AutoScalingGroupV2ApiGetAsgListV2Opts struct {
	AsgName         optional.String
	AsgState        optional.String
	LcName          optional.String
	LocalSubnetId   optional.String
	SecurityGroupId optional.String
	ServiceId       optional.String
	ServiceZoneId   optional.String
	SubnetId        optional.String
	VpcId           optional.String
	CreatedBy       optional.String
	Page            optional.Int32
	Size            optional.Int32
	Sort            optional.String
}

type BasicAuth struct {
	UserName string `json:"userName,omitempty"`
	Password string `json:"password,omitempty"`
}

type CheckResponse struct {
	Result *bool `json:"result,omitempty"`
}

type LaunchConfigBlockStorageRequest struct {
	BlockStorageSize  *int32 `json:"blockStorageSize"`
	EncryptionEnabled *bool  `json:"encryptionEnabled"`
	IsBootDisk        *bool  `json:"isBootDisk"`
	ProductId         string `json:"productId"`
}

type LaunchConfigBlockStorageResponse struct {
	BlockStorageSize  *int32 `json:"blockStorageSize,omitempty"`
	DiskType          string `json:"diskType,omitempty"`
	EncryptionEnabled *bool  `json:"encryptionEnabled,omitempty"`
	IsBootDisk        *bool  `json:"isBootDisk,omitempty"`
	ProductId         string `json:"productId,omitempty"`
}

type LaunchConfigBlockStorageV2Request struct {
	BlockStorageSize  *int32 `json:"blockStorageSize"`
	DiskType          string `json:"diskType"`
	EncryptionEnabled *bool  `json:"encryptionEnabled"`
	IsBootDisk        *bool  `json:"isBootDisk"`
}

type LaunchConfigCreateV5Request struct {
	BlockStorages  []LaunchConfigBlockStorageRequest `json:"blockStorages"`
	ImageId        string                                            `json:"imageId"`
	InitialScript  string                                            `json:"initialScript,omitempty"`
	LcName         string                                            `json:"lcName"`
	ScaleProductId string                                            `json:"scaleProductId"`
	ServiceZoneId  string                                            `json:"serviceZoneId"`
	Tags           []TagRequest                      `json:"tags,omitempty"`
}

type LaunchConfigCreateV6Request struct {
	BlockStorages []LaunchConfigBlockStorageV2Request `json:"blockStorages"`
	ImageId       string                                              `json:"imageId"`
	InitialScript string                                              `json:"initialScript,omitempty"`
	KeyPairId     string                                              `json:"keyPairId"`
	LcName        string                                              `json:"lcName"`
	ServerType    string                                              `json:"serverType"`
	ServiceZoneId string                                              `json:"serviceZoneId"`
	Tags          []TagRequest                        `json:"tags,omitempty"`
}

type LaunchConfigDetailV4Response struct {
	ProjectId         string                                             `json:"projectId,omitempty"`
	AsgIds            []string                                           `json:"asgIds,omitempty"`
	BlockId           string                                             `json:"blockId,omitempty"`
	BlockStorages     []LaunchConfigBlockStorageResponse `json:"blockStorages,omitempty"`
	ContractProductId string                                             `json:"contractProductId,omitempty"`
	ImageId           string                                             `json:"imageId,omitempty"`
	InitialScript     string                                             `json:"initialScript,omitempty"`
	KeyPairId         string                                             `json:"keyPairId,omitempty"`
	LcId              string                                             `json:"lcId,omitempty"`
	LcName            string                                             `json:"lcName,omitempty"`
	OsProductId       string                                             `json:"osProductId,omitempty"`
	OsType            string                                             `json:"osType,omitempty"`
	ProductGroupId    string                                             `json:"productGroupId,omitempty"`
	ScaleProductId    string                                             `json:"scaleProductId,omitempty"`
	ServerType        string                                             `json:"serverType,omitempty"`
	ServiceZoneId     string                                             `json:"serviceZoneId,omitempty"`
	CreatedBy         string                                             `json:"createdBy,omitempty"`
	CreatedDt         time.Time                                       `json:"createdDt,omitempty"`
	ModifiedBy        string                                             `json:"modifiedBy,omitempty"`
	ModifiedDt        time.Time                                       `json:"modifiedDt,omitempty"`
}

type LaunchConfigListItemResponse struct {
	ProjectId     string       `json:"projectId,omitempty"`
	AsgCount      *int32       `json:"asgCount,omitempty"`
	BlockId       string       `json:"blockId,omitempty"`
	ImageId       string       `json:"imageId,omitempty"`
	KeyPairId     string       `json:"keyPairId,omitempty"`
	LcId          string       `json:"lcId,omitempty"`
	LcName        string       `json:"lcName,omitempty"`
	ServiceZoneId string       `json:"serviceZoneId,omitempty"`
	CreatedBy     string       `json:"createdBy,omitempty"`
	CreatedDt     time.Time `json:"createdDt,omitempty"`
	ModifiedBy    string       `json:"modifiedBy,omitempty"`
	ModifiedDt    time.Time `json:"modifiedDt,omitempty"`
}

type ListResponseAsgEventListItemResponse struct {
	Contents   []AsgEventListItemResponse `json:"contents,omitempty"`
	TotalCount *int32                                     `json:"totalCount,omitempty"`
}

type ListResponseAsgNotificationResponse struct {
	Contents   []AsgNotificationResponse `json:"contents,omitempty"`
	TotalCount *int32                                    `json:"totalCount,omitempty"`
}

type ListResponseAsgPolicyResponse struct {
	Contents   []AsgPolicyResponse `json:"contents,omitempty"`
	TotalCount *int32                              `json:"totalCount,omitempty"`
}

type ListResponseAsgScheduleListItemResponse struct {
	Contents   []AsgScheduleListItemResponse `json:"contents,omitempty"`
	TotalCount *int32                                        `json:"totalCount,omitempty"`
}

type ListResponseAutoScalingGroupResponse struct {
	Contents   []AutoScalingGroupResponse `json:"contents,omitempty"`
	TotalCount *int32                                     `json:"totalCount,omitempty"`
}

type ListResponseLaunchConfigListItemResponse struct {
	Contents   []LaunchConfigListItemResponse `json:"contents,omitempty"`
	TotalCount *int32                                         `json:"totalCount,omitempty"`
}

type PageResponseV2AsgVirtualServerListItemResponse struct {
	Contents   []AsgVirtualServerListItemResponse `json:"contents,omitempty"`
	Page       *int32                                             `json:"page,omitempty"`
	Size       *int32                                             `json:"size,omitempty"`
	Sort       []string                                           `json:"sort,omitempty"`
	TotalCount int64                                              `json:"totalCount,omitempty"`
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
