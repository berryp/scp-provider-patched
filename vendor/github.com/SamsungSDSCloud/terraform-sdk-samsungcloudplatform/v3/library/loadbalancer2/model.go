package loadbalancer2

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

type AutoScalingGroupInfo struct {
	AutoScalingGroupId   string                               `json:"autoScalingGroupId,omitempty"`
	AutoScalingGroupName string                               `json:"autoScalingGroupName,omitempty"`
	LbServerGroups       []LbServerGroupInfo `json:"lbServerGroups,omitempty"`
}

type BasicAuth struct {
	UserName string `json:"userName,omitempty"`
	Password string `json:"password,omitempty"`
}

type LbChangeRequest struct {
	LoadBalancerDescription string `json:"loadBalancerDescription,omitempty"`
}

type LbCreateV3Request struct {
	BlockId                 string                        `json:"blockId"`
	FirewallEnabled         *bool                         `json:"firewallEnabled,omitempty"`
	IsLoggable              *bool                         `json:"isLoggable"`
	LinkIpCidr              string                        `json:"linkIpCidr,omitempty"`
	LoadBalancerName        string                        `json:"loadBalancerName"`
	LoadBalancerSize        string                        `json:"loadBalancerSize"`
	ServiceIpCidr           string                        `json:"serviceIpCidr"`
	ServiceZoneId           string                        `json:"serviceZoneId"`
	Tags                    []TagRequest `json:"tags,omitempty"`
	VpcId                   string                        `json:"vpcId"`
	LoadBalancerDescription string                        `json:"loadBalancerDescription,omitempty"`
}

type LbDetailResponse struct {
	ProjectId               string       `json:"projectId,omitempty"`
	BlockId                 string       `json:"blockId,omitempty"`
	LinkIpAddress           string       `json:"linkIpAddress,omitempty"`
	LinkIpPoolId            string       `json:"linkIpPoolId,omitempty"`
	LoadBalancerId          string       `json:"loadBalancerId,omitempty"`
	LoadBalancerName        string       `json:"loadBalancerName,omitempty"`
	LoadBalancerSize        string       `json:"loadBalancerSize,omitempty"`
	LoadBalancerState       string       `json:"loadBalancerState,omitempty"`
	ServiceIpCidr           string       `json:"serviceIpCidr,omitempty"`
	ServiceZoneId           string       `json:"serviceZoneId,omitempty"`
	VpcId                   string       `json:"vpcId,omitempty"`
	LoadBalancerDescription string       `json:"loadBalancerDescription,omitempty"`
	CreatedBy               string       `json:"createdBy,omitempty"`
	CreatedDt               time.Time `json:"createdDt,omitempty"`
	ModifiedBy              string       `json:"modifiedBy,omitempty"`
	ModifiedDt              time.Time `json:"modifiedDt,omitempty"`
}

type LbLogStorageCreateRequest struct {
	LogStorageType string `json:"logStorageType"`
	ObsBucketId    string `json:"obsBucketId"`
	VpcId          string `json:"vpcId"`
}

type LbLogStorageOpenApiControllerApiGetLoadBalancerLogStorageListOpts struct {
	ObsBucketId optional.String
}

type LbLogStorageResponse struct {
	ProjectId      string       `json:"projectId,omitempty"`
	BlockId        string       `json:"blockId,omitempty"`
	LogStorageId   string       `json:"logStorageId,omitempty"`
	LogStorageType string       `json:"logStorageType,omitempty"`
	ObsBucketId    string       `json:"obsBucketId,omitempty"`
	ServiceZoneId  string       `json:"serviceZoneId,omitempty"`
	VpcId          string       `json:"vpcId,omitempty"`
	CreatedBy      string       `json:"createdBy,omitempty"`
	CreatedDt      time.Time `json:"createdDt,omitempty"`
	ModifiedBy     string       `json:"modifiedBy,omitempty"`
	ModifiedDt     time.Time `json:"modifiedDt,omitempty"`
}

type LbLogStorageUpdateRequest struct {
	ObsBucketId string `json:"obsBucketId"`
}

type LbMonitorRequest struct {
	HttpMethod        string `json:"httpMethod,omitempty"`
	HttpVersion       string `json:"httpVersion,omitempty"`
	LbMonitorCount    int32  `json:"lbMonitorCount"`
	LbMonitorInterval int32  `json:"lbMonitorInterval"`
	LbMonitorPort     int32  `json:"lbMonitorPort,omitempty"`
	LbMonitorTimeout  int32  `json:"lbMonitorTimeout"`
	LbMonitorUrl      string `json:"lbMonitorUrl,omitempty"`
	Protocol          string `json:"protocol"`
	RequestBody       string `json:"requestBody,omitempty"`
	ResponseBody      string `json:"responseBody,omitempty"`
}

type LbMonitorResponse struct {
	HttpMethod        string `json:"httpMethod,omitempty"`
	HttpVersion       string `json:"httpVersion,omitempty"`
	LbMonitorCount    int32  `json:"lbMonitorCount"`
	LbMonitorId       string `json:"lbMonitorId,omitempty"`
	LbMonitorInterval int32  `json:"lbMonitorInterval"`
	LbMonitorPort     int32  `json:"lbMonitorPort"`
	LbMonitorState    string `json:"lbMonitorState,omitempty"`
	LbMonitorTimeout  int32  `json:"lbMonitorTimeout"`
	LbMonitorUrl      string `json:"lbMonitorUrl,omitempty"`
	Protocol          string `json:"protocol,omitempty"`
	RequestBody       string `json:"requestBody,omitempty"`
	ResponseBody      string `json:"responseBody,omitempty"`
}

type LbProfileAttrCreateRequest struct {
	RedirectType       string `json:"redirectType,omitempty"`
	RequestHeaderSize  int32  `json:"requestHeaderSize,omitempty"`
	ResponseHeaderSize int32  `json:"responseHeaderSize,omitempty"`
	ResponseTimeout    int32  `json:"responseTimeout,omitempty"`
	SessionTimeout     int32  `json:"sessionTimeout,omitempty"`
	XforwardedFor      string `json:"xforwardedFor,omitempty"`
}

type LbProfileAttrModifyRequest struct {
	RedirectType       string `json:"redirectType,omitempty"`
	RequestHeaderSize  int32  `json:"requestHeaderSize,omitempty"`
	ResponseHeaderSize int32  `json:"responseHeaderSize,omitempty"`
	ResponseTimeout    int32  `json:"responseTimeout,omitempty"`
	SessionTimeout     int32  `json:"sessionTimeout,omitempty"`
	XforwardedFor      string `json:"xforwardedFor,omitempty"`
}

type LbProfileAttrResponse struct {
	RedirectType       string `json:"redirectType,omitempty"`
	RequestHeaderSize  int32  `json:"requestHeaderSize,omitempty"`
	ResponseHeaderSize int32  `json:"responseHeaderSize,omitempty"`
	ResponseTimeout    int32  `json:"responseTimeout,omitempty"`
	SessionTimeout     int32  `json:"sessionTimeout,omitempty"`
	XforwardedFor      string `json:"xforwardedFor,omitempty"`
}

type LbProfileChangeRequest struct {
	LbProfileAttrs *LbProfileAttrModifyRequest `json:"lbProfileAttrs,omitempty"`
}

type LbProfileCreateRequest struct {
	LayerType         string                                       `json:"layerType,omitempty"`
	LbProfileAttrs    *LbProfileAttrCreateRequest `json:"lbProfileAttrs,omitempty"`
	LbProfileCategory string                                       `json:"lbProfileCategory"`
	LbProfileName     string                                       `json:"lbProfileName"`
	LbProfileType     string                                       `json:"lbProfileType"`
	Protocol          string                                       `json:"protocol,omitempty"`
	Tags              []TagRequest                `json:"tags,omitempty"`
}

type LbProfileDetailResponse struct {
	ProjectId         string                                     `json:"projectId,omitempty"`
	BlockId           string                                     `json:"blockId,omitempty"`
	LayerType         string                                     `json:"layerType,omitempty"`
	LbId              string                                     `json:"lbId,omitempty"`
	LbProfileAttrs    *LbProfileAttrResponse    `json:"lbProfileAttrs,omitempty"`
	LbProfileCategory string                                     `json:"lbProfileCategory,omitempty"`
	LbProfileId       string                                     `json:"lbProfileId,omitempty"`
	LbProfileName     string                                     `json:"lbProfileName,omitempty"`
	LbProfileState    string                                     `json:"lbProfileState,omitempty"`
	LbProfileType     string                                     `json:"lbProfileType,omitempty"`
	LbServiceNames    []string                                   `json:"lbServiceNames,omitempty"`
	LbServices        []LbServiceIdNameResponse `json:"lbServices,omitempty"`
	Protocol          string                                     `json:"protocol,omitempty"`
	ServiceZoneId     string                                     `json:"serviceZoneId,omitempty"`
	CreatedBy         string                                     `json:"createdBy,omitempty"`
	CreatedDt         time.Time                               `json:"createdDt,omitempty"`
	ModifiedBy        string                                     `json:"modifiedBy,omitempty"`
	ModifiedDt        time.Time                               `json:"modifiedDt,omitempty"`
}

type LbProfileOpenApiControllerApiGetLoadBalancerProfileListOpts struct {
	LbProfileCategory optional.String
	LbProfileName     optional.String
	LbServiceName     optional.String
	LoadBalancerName  optional.String
	CreatedBy         optional.String
	Page              optional.Int32
	Size              optional.Int32
	Sort              optional.Interface
}

type LbProfileResponse struct {
	ProjectId         string       `json:"projectId,omitempty"`
	BlockId           string       `json:"blockId,omitempty"`
	LbProfileCategory string       `json:"lbProfileCategory,omitempty"`
	LbProfileName     string       `json:"lbProfileName,omitempty"`
	LbProfileState    string       `json:"lbProfileState,omitempty"`
	LbProfileType     string       `json:"lbProfileType,omitempty"`
	LbServiceNames    []string     `json:"lbServiceNames,omitempty"`
	ServiceZoneId     string       `json:"serviceZoneId,omitempty"`
	CreatedBy         string       `json:"createdBy,omitempty"`
	CreatedDt         time.Time `json:"createdDt,omitempty"`
	ModifiedBy        string       `json:"modifiedBy,omitempty"`
	ModifiedDt        time.Time `json:"modifiedDt,omitempty"`
}

type LbRequest struct {
	BlockId                 string                        `json:"blockId"`
	FirewallEnabled         *bool                         `json:"firewallEnabled,omitempty"`
	LinkIpCidr              string                        `json:"linkIpCidr,omitempty"`
	LoadBalancerName        string                        `json:"loadBalancerName"`
	LoadBalancerSize        string                        `json:"loadBalancerSize"`
	ProductGroupId          string                        `json:"productGroupId"`
	ProductId               string                        `json:"productId"`
	ServiceIpCidr           string                        `json:"serviceIpCidr"`
	ServiceZoneId           string                        `json:"serviceZoneId"`
	Tags                    []TagRequest `json:"tags,omitempty"`
	VpcId                   string                        `json:"vpcId"`
	LoadBalancerDescription string                        `json:"loadBalancerDescription,omitempty"`
}

type LbResponse struct {
	ProjectId         string       `json:"projectId,omitempty"`
	BlockId           string       `json:"blockId,omitempty"`
	LoadBalancerId    string       `json:"loadBalancerId,omitempty"`
	LoadBalancerName  string       `json:"loadBalancerName,omitempty"`
	LoadBalancerSize  string       `json:"loadBalancerSize,omitempty"`
	LoadBalancerState string       `json:"loadBalancerState,omitempty"`
	ServiceZoneId     string       `json:"serviceZoneId,omitempty"`
	VpcId             string       `json:"vpcId,omitempty"`
	CreatedBy         string       `json:"createdBy,omitempty"`
	CreatedDt         time.Time `json:"createdDt,omitempty"`
	ModifiedBy        string       `json:"modifiedBy,omitempty"`
	ModifiedDt        time.Time `json:"modifiedDt,omitempty"`
}

type LbRuleConsoleResponse struct {
	LbRuleId             string                                                `json:"lbRuleId,omitempty"`
	LbRuleSeq            int32                                                 `json:"lbRuleSeq,omitempty"`
	LbServerGroupId      string                                                `json:"lbServerGroupId,omitempty"`
	LbServerGroupMembers []LbServerGroupMemberConsoleResponse `json:"lbServerGroupMembers,omitempty"`
	LbServerGroupName    string                                                `json:"lbServerGroupName,omitempty"`
	PatternUrl           string                                                `json:"patternUrl,omitempty"`
}

type LbRuleForAsgResponse struct {
	AutoScaleGroupIds []string `json:"autoScaleGroupIds,omitempty"`
	LbRuleId          string   `json:"lbRuleId,omitempty"`
	LbServerGroupId   string   `json:"lbServerGroupId,omitempty"`
	PatternUrl        string   `json:"patternUrl,omitempty"`
	Seq               int32    `json:"seq,omitempty"`
}

type LbRuleRequest struct {
	LbRuleId        string `json:"lbRuleId,omitempty"`
	LbRuleSeq       int32  `json:"lbRuleSeq"`
	LbServerGroupId string `json:"lbServerGroupId,omitempty"`
	PatternUrl      string `json:"patternUrl,omitempty"`
}

type LbRuleResponse struct {
	LbRuleId          string `json:"lbRuleId,omitempty"`
	LbRuleSeq         int32  `json:"lbRuleSeq,omitempty"`
	LbServerGroupId   string `json:"lbServerGroupId,omitempty"`
	LbServerGroupName string `json:"lbServerGroupName,omitempty"`
	PatternUrl        string `json:"patternUrl,omitempty"`
}

type LbServerGroupChangeRequest struct {
	LbAlgorithm            string                                        `json:"lbAlgorithm"`
	LbMonitor              *LbMonitorRequest            `json:"lbMonitor,omitempty"`
	LbServerGroupMembers   []LbServerGroupMemberRequest `json:"lbServerGroupMembers,omitempty"`
	TcpMultiplexingEnabled *bool                                         `json:"tcpMultiplexingEnabled,omitempty"`
}

type LbServerGroupCreateRequest struct {
	LbAlgorithm            string                                        `json:"lbAlgorithm"`
	LbMonitor              *LbMonitorRequest            `json:"lbMonitor,omitempty"`
	LbServerGroupMembers   []LbServerGroupMemberRequest `json:"lbServerGroupMembers,omitempty"`
	LbServerGroupName      string                                        `json:"lbServerGroupName"`
	Tags                   []TagRequest                 `json:"tags,omitempty"`
	TcpMultiplexingEnabled *bool                                         `json:"tcpMultiplexingEnabled,omitempty"`
}

type LbServerGroupDetailResponse struct {
	ProjectId              string                                         `json:"projectId,omitempty"`
	BlockId                string                                         `json:"blockId,omitempty"`
	LbMonitor              *LbMonitorResponse            `json:"lbMonitor,omitempty"`
	LbServerGroupAlgorithm string                                         `json:"lbServerGroupAlgorithm,omitempty"`
	LbServerGroupId        string                                         `json:"lbServerGroupId,omitempty"`
	LbServerGroupMembers   []LbServerGroupMemberResponse `json:"lbServerGroupMembers,omitempty"`
	LbServerGroupName      string                                         `json:"lbServerGroupName,omitempty"`
	LbServerGroupState     string                                         `json:"lbServerGroupState,omitempty"`
	LbServerGroupType      string                                         `json:"lbServerGroupType,omitempty"`
	LbServices             []LbServiceSubResponse        `json:"lbServices,omitempty"`
	LoadBalancerId         string                                         `json:"loadBalancerId,omitempty"`
	LoadBalancerName       string                                         `json:"loadBalancerName,omitempty"`
	LoadBalancerState      string                                         `json:"loadBalancerState,omitempty"`
	Persistence            string                                         `json:"persistence,omitempty"`
	ServiceZoneId          string                                         `json:"serviceZoneId,omitempty"`
	TcpMultiplexingEnabled *bool                                          `json:"tcpMultiplexingEnabled,omitempty"`
	CreatedBy              string                                         `json:"createdBy,omitempty"`
	CreatedDt              time.Time                                   `json:"createdDt,omitempty"`
	ModifiedBy             string                                         `json:"modifiedBy,omitempty"`
	ModifiedDt             time.Time                                   `json:"modifiedDt,omitempty"`
}

type LbServerGroupInfo struct {
	LbServerGroupId   string `json:"lbServerGroupId,omitempty"`
	LbServerGroupName string `json:"lbServerGroupName,omitempty"`
}

type LbServerGroupMemberConsoleResponse struct {
	LbServerGroupMemberNsxStatus string `json:"lbServerGroupMemberNsxStatus,omitempty"`
	ObjectId                     string `json:"objectId,omitempty"`
	ObjectIpAddress              string `json:"objectIpAddress,omitempty"`
	ObjectName                   string `json:"objectName,omitempty"`
}

type LbServerGroupMemberRequest struct {
	JoinState       string `json:"joinState"`
	ObjectId        string `json:"objectId,omitempty"`
	ObjectIpAddress string `json:"objectIpAddress,omitempty"`
	ObjectPort      int32  `json:"objectPort,omitempty"`
	ObjectType      string `json:"objectType"`
	Weight          int32  `json:"weight,omitempty"`
}

type LbServerGroupMemberResponse struct {
	AutoScalingGroupId    string `json:"autoScalingGroupId,omitempty"`
	JoinState             string `json:"joinState,omitempty"`
	LbServerGroupMemberId string `json:"lbServerGroupMemberId,omitempty"`
	MemberIpAddress       string `json:"memberIpAddress,omitempty"`
	MemberPort            int32  `json:"memberPort,omitempty"`
	MemberWeight          int32  `json:"memberWeight,omitempty"`
	ObjectId              string `json:"objectId,omitempty"`
	ObjectName            string `json:"objectName,omitempty"`
	ObjectType            string `json:"objectType,omitempty"`
}

type LbServerGroupOpenApiControllerApiGetLoadBalancerServerGroupListOpts struct {
	LbServerGroupName optional.String
	LbServiceName     optional.String
	LoadBalancerName  optional.String
	MemberIpAddress   optional.String
	CreatedBy         optional.String
	Page              optional.Int32
	Size              optional.Int32
	Sort              optional.Interface
}

type LbServerGroupResponse struct {
	ProjectId            string                                         `json:"projectId,omitempty"`
	BlockId              string                                         `json:"blockId,omitempty"`
	LbServerGroupId      string                                         `json:"lbServerGroupId,omitempty"`
	LbServerGroupMembers []LbServerGroupMemberResponse `json:"lbServerGroupMembers,omitempty"`
	LbServerGroupName    string                                         `json:"lbServerGroupName,omitempty"`
	LbServerGroupState   string                                         `json:"lbServerGroupState,omitempty"`
	LbServerGroupType    string                                         `json:"lbServerGroupType,omitempty"`
	LbServices           []LbServiceSubResponse        `json:"lbServices,omitempty"`
	LoadBalancerState    string                                         `json:"loadBalancerState,omitempty"`
	Persistence          string                                         `json:"persistence,omitempty"`
	ServiceZoneId        string                                         `json:"serviceZoneId,omitempty"`
	CreatedBy            string                                         `json:"createdBy,omitempty"`
	CreatedDt            time.Time                                   `json:"createdDt,omitempty"`
	ModifiedBy           string                                         `json:"modifiedBy,omitempty"`
	ModifiedDt           time.Time                                   `json:"modifiedDt,omitempty"`
}

type LbServiceChangeRequest struct {
	ApplicationProfileId   string                           `json:"applicationProfileId,omitempty"`
	ClientCertificateId    string                           `json:"clientCertificateId,omitempty"`
	ClientSslSecurityLevel string                           `json:"clientSslSecurityLevel,omitempty"`
	DefaultForwardingPorts string                           `json:"defaultForwardingPorts,omitempty"`
	LbRules                []LbRuleRequest `json:"lbRules,omitempty"`
	Persistence            string                           `json:"persistence,omitempty"`
	PersistenceProfileId   string                           `json:"persistenceProfileId,omitempty"`
	ServerCertificateId    string                           `json:"serverCertificateId,omitempty"`
	ServerSslSecurityLevel string                           `json:"serverSslSecurityLevel,omitempty"`
	ServicePorts           string                           `json:"servicePorts,omitempty"`
	UseAccessLog           *bool                            `json:"useAccessLog,omitempty"`
}

type LbServiceDetailConsoleResponse struct {
	ProjectId              string                                   `json:"projectId,omitempty"`
	ApplicationProfileId   string                                   `json:"applicationProfileId,omitempty"`
	ApplicationProfileName string                                   `json:"applicationProfileName,omitempty"`
	AutoScalingGroups      []AutoScalingGroupInfo  `json:"autoScalingGroups,omitempty"`
	BlockId                string                                   `json:"blockId,omitempty"`
	ClientCertificateId    string                                   `json:"clientCertificateId,omitempty"`
	ClientCertificateName  string                                   `json:"clientCertificateName,omitempty"`
	ClientSslSecurityLevel string                                   `json:"clientSslSecurityLevel,omitempty"`
	DefaultForwardingPorts string                                   `json:"defaultForwardingPorts,omitempty"`
	LayerType              string                                   `json:"layerType,omitempty"`
	LbRules                []LbRuleConsoleResponse `json:"lbRules,omitempty"`
	LbServiceId            string                                   `json:"lbServiceId,omitempty"`
	LbServiceIpId          string                                   `json:"lbServiceIpId,omitempty"`
	LbServiceIpState       string                                   `json:"lbServiceIpState,omitempty"`
	LbServiceName          string                                   `json:"lbServiceName,omitempty"`
	LbServiceNsxStatus     string                                   `json:"lbServiceNsxStatus,omitempty"`
	LbServiceState         string                                   `json:"lbServiceState,omitempty"`
	LoadBalancerId         string                                   `json:"loadBalancerId,omitempty"`
	LoadBalancerName       string                                   `json:"loadBalancerName,omitempty"`
	LoadBalancerState      string                                   `json:"loadBalancerState,omitempty"`
	NatIpAddress           string                                   `json:"natIpAddress,omitempty"`
	Persistence            string                                   `json:"persistence,omitempty"`
	PersistenceProfileId   string                                   `json:"persistenceProfileId,omitempty"`
	PersistenceProfileName string                                   `json:"persistenceProfileName,omitempty"`
	Protocol               string                                   `json:"protocol,omitempty"`
	ServerCertificateId    string                                   `json:"serverCertificateId,omitempty"`
	ServerCertificateName  string                                   `json:"serverCertificateName,omitempty"`
	ServerSslSecurityLevel string                                   `json:"serverSslSecurityLevel,omitempty"`
	ServiceIpAddress       string                                   `json:"serviceIpAddress,omitempty"`
	ServicePorts           string                                   `json:"servicePorts,omitempty"`
	ServiceZoneId          string                                   `json:"serviceZoneId,omitempty"`
	UseAccessLog           *bool                                    `json:"useAccessLog,omitempty"`
	UseAutoscaling         *bool                                    `json:"useAutoscaling,omitempty"`
	CreatedBy              string                                   `json:"createdBy,omitempty"`
	CreatedDt              time.Time                             `json:"createdDt,omitempty"`
	ModifiedBy             string                                   `json:"modifiedBy,omitempty"`
	ModifiedDt             time.Time                             `json:"modifiedDt,omitempty"`
}

type LbServiceDetailResponse struct {
	ProjectId              string                            `json:"projectId,omitempty"`
	ApplicationProfileId   string                            `json:"applicationProfileId,omitempty"`
	ApplicationProfileName string                            `json:"applicationProfileName,omitempty"`
	BlockId                string                            `json:"blockId,omitempty"`
	ClientCertificateId    string                            `json:"clientCertificateId,omitempty"`
	ClientCertificateName  string                            `json:"clientCertificateName,omitempty"`
	ClientSslSecurityLevel string                            `json:"clientSslSecurityLevel,omitempty"`
	DefaultForwardingPorts string                            `json:"defaultForwardingPorts,omitempty"`
	LayerType              string                            `json:"layerType,omitempty"`
	LbRules                []LbRuleResponse `json:"lbRules,omitempty"`
	LbServiceId            string                            `json:"lbServiceId,omitempty"`
	LbServiceIpId          string                            `json:"lbServiceIpId,omitempty"`
	LbServiceName          string                            `json:"lbServiceName,omitempty"`
	LbServiceState         string                            `json:"lbServiceState,omitempty"`
	NatIpAddress           string                            `json:"natIpAddress,omitempty"`
	Persistence            string                            `json:"persistence,omitempty"`
	PersistenceProfileId   string                            `json:"persistenceProfileId,omitempty"`
	PersistenceProfileName string                            `json:"persistenceProfileName,omitempty"`
	Protocol               string                            `json:"protocol,omitempty"`
	ServerCertificateId    string                            `json:"serverCertificateId,omitempty"`
	ServerCertificateName  string                            `json:"serverCertificateName,omitempty"`
	ServerSslSecurityLevel string                            `json:"serverSslSecurityLevel,omitempty"`
	ServiceIpAddress       string                            `json:"serviceIpAddress,omitempty"`
	ServicePorts           string                            `json:"servicePorts,omitempty"`
	ServiceZoneId          string                            `json:"serviceZoneId,omitempty"`
	UseAccessLog           *bool                             `json:"useAccessLog,omitempty"`
	CreatedBy              string                            `json:"createdBy,omitempty"`
	CreatedDt              time.Time                      `json:"createdDt,omitempty"`
	ModifiedBy             string                            `json:"modifiedBy,omitempty"`
	ModifiedDt             time.Time                      `json:"modifiedDt,omitempty"`
}

type LbServiceForAsgResponse struct {
	AutoScaleGroupIds      []string                                `json:"autoScaleGroupIds,omitempty"`
	DefaultForwardingPorts string                                  `json:"defaultForwardingPorts,omitempty"`
	LayerType              string                                  `json:"layerType,omitempty"`
	LbRules                []LbRuleForAsgResponse `json:"lbRules,omitempty"`
	LbServiceId            string                                  `json:"lbServiceId,omitempty"`
	LbServiceIpAddress     string                                  `json:"lbServiceIpAddress,omitempty"`
	LbServiceName          string                                  `json:"lbServiceName,omitempty"`
	LbServiceState         string                                  `json:"lbServiceState,omitempty"`
	LoadBalancerId         string                                  `json:"loadBalancerId,omitempty"`
	LoadBalancerName       string                                  `json:"loadBalancerName,omitempty"`
	NatIpAddress           string                                  `json:"natIpAddress,omitempty"`
	Persistence            string                                  `json:"persistence,omitempty"`
	Protocol               string                                  `json:"protocol,omitempty"`
	ServicePorts           string                                  `json:"servicePorts,omitempty"`
}

type LbServiceIdNameResponse struct {
	LbServiceId   string `json:"lbServiceId,omitempty"`
	LbServiceName string `json:"lbServiceName,omitempty"`
}

type LbServiceIpRequest struct {
	LbServiceIpId string `json:"lbServiceIpId"`
	NatActive     *bool  `json:"natActive,omitempty"`
	PublicIpId    string `json:"publicIpId,omitempty"`
}

type LbServiceIpResponse struct {
	LbServiceIpId    string                            `json:"lbServiceIpId,omitempty"`
	LbServiceIpState string                            `json:"lbServiceIpState,omitempty"`
	LbServices       []LbServiceSubSo `json:"lbServices,omitempty"`
	NatIpAddress     string                            `json:"natIpAddress,omitempty"`
	NatIpId          string                            `json:"natIpId,omitempty"`
	ServiceIpAddress string                            `json:"serviceIpAddress,omitempty"`
	ServiceIpCidr    string                            `json:"serviceIpCidr,omitempty"`
	ServiceIpId      string                            `json:"serviceIpId,omitempty"`
	ServiceIpPoolId  string                            `json:"serviceIpPoolId,omitempty"`
}

type LbServiceOpenApiControllerApiGetLoadBalancerServiceIncludingServerGroupMembersInfoOpts struct {
	StatusCheck optional.Bool
}

type LbServiceOpenApiControllerApiGetLoadBalancerServiceIpListOpts struct {
	LbServiceName    optional.String
	NatIpAddress     optional.String
	ServiceIpAddress optional.String
	Page             optional.Int32
	Size             optional.Int32
	Sort             optional.Interface
}

type LbServiceOpenApiControllerApiGetLoadBalancerServiceListOpts struct {
	LayerType        optional.String
	LbServiceName    optional.String
	LoadBalancerName optional.String
	NatIpAddress     optional.String
	Protocol         optional.String
	ServiceIpAddress optional.String
	StatusCheck      optional.Bool
	CreatedBy        optional.String
	Page             optional.Int32
	Size             optional.Int32
	Sort             optional.Interface
}

type LbServiceRequest struct {
	ApplicationProfileId   string                           `json:"applicationProfileId"`
	ClientCertificateId    string                           `json:"clientCertificateId,omitempty"`
	ClientSslSecurityLevel string                           `json:"clientSslSecurityLevel,omitempty"`
	DefaultForwardingPorts string                           `json:"defaultForwardingPorts,omitempty"`
	LayerType              string                           `json:"layerType"`
	LbRules                []LbRuleRequest `json:"lbRules,omitempty"`
	LbServiceIpId          string                           `json:"lbServiceIpId,omitempty"`
	LbServiceName          string                           `json:"lbServiceName"`
	NatActive              *bool                            `json:"natActive,omitempty"`
	Persistence            string                           `json:"persistence"`
	PersistenceProfileId   string                           `json:"persistenceProfileId,omitempty"`
	Protocol               string                           `json:"protocol"`
	PublicIpId             string                           `json:"publicIpId,omitempty"`
	ServerCertificateId    string                           `json:"serverCertificateId,omitempty"`
	ServerSslSecurityLevel string                           `json:"serverSslSecurityLevel,omitempty"`
	ServiceIpAddress       string                           `json:"serviceIpAddress"`
	ServicePorts           string                           `json:"servicePorts"`
	Tags                   []TagRequest    `json:"tags,omitempty"`
	UseAccessLog           *bool                            `json:"useAccessLog,omitempty"`
}

type LbServiceResponse struct {
	ProjectId              string       `json:"projectId,omitempty"`
	BlockId                string       `json:"blockId,omitempty"`
	DefaultForwardingPorts string       `json:"defaultForwardingPorts,omitempty"`
	LayerType              string       `json:"layerType,omitempty"`
	LbServiceId            string       `json:"lbServiceId,omitempty"`
	LbServiceIpId          string       `json:"lbServiceIpId,omitempty"`
	LbServiceName          string       `json:"lbServiceName,omitempty"`
	LbServiceState         string       `json:"lbServiceState,omitempty"`
	NatIpAddress           string       `json:"natIpAddress,omitempty"`
	Protocol               string       `json:"protocol,omitempty"`
	ServiceIpAddress       string       `json:"serviceIpAddress,omitempty"`
	ServicePorts           string       `json:"servicePorts,omitempty"`
	ServiceZoneId          string       `json:"serviceZoneId,omitempty"`
	CreatedBy              string       `json:"createdBy,omitempty"`
	CreatedDt              time.Time `json:"createdDt,omitempty"`
	ModifiedBy             string       `json:"modifiedBy,omitempty"`
	ModifiedDt             time.Time `json:"modifiedDt,omitempty"`
}

type LbServiceSubResponse struct {
	LbServiceId    string `json:"lbServiceId,omitempty"`
	LbServiceName  string `json:"lbServiceName,omitempty"`
	LbServiceState string `json:"lbServiceState,omitempty"`
	Persistence    string `json:"persistence,omitempty"`
}

type LbServiceSubSo struct {
	LbServiceId    string `json:"lbServiceId,omitempty"`
	LbServiceName  string `json:"lbServiceName,omitempty"`
	LbServiceState string `json:"lbServiceState,omitempty"`
	Persistence    string `json:"persistence,omitempty"`
	VendorObjectId string `json:"vendorObjectId,omitempty"`
}

type LbUsageResponse struct {
	Limit int64 `json:"limit,omitempty"`
	Usage int64 `json:"usage,omitempty"`
}

type ListResponseLbLogStorageResponse struct {
	Contents   []LbLogStorageResponse `json:"contents,omitempty"`
	TotalCount int32                                   `json:"totalCount,omitempty"`
}

type ListResponseLbProfileResponse struct {
	Contents   []LbProfileResponse `json:"contents,omitempty"`
	TotalCount int32                                `json:"totalCount,omitempty"`
}

type ListResponseLbResponse struct {
	Contents   []LbResponse `json:"contents,omitempty"`
	TotalCount int32                         `json:"totalCount,omitempty"`
}

type ListResponseLbServerGroupResponse struct {
	Contents   []LbServerGroupResponse `json:"contents,omitempty"`
	TotalCount int32                                    `json:"totalCount,omitempty"`
}

type ListResponseLbServiceForAsgResponse struct {
	Contents   []LbServiceForAsgResponse `json:"contents,omitempty"`
	TotalCount int32                                      `json:"totalCount,omitempty"`
}

type ListResponseLbServiceIpResponse struct {
	Contents   []LbServiceIpResponse `json:"contents,omitempty"`
	TotalCount int32                                  `json:"totalCount,omitempty"`
}

type ListResponseLbServiceResponse struct {
	Contents   []LbServiceResponse `json:"contents,omitempty"`
	TotalCount int32                                `json:"totalCount,omitempty"`
}

type ListResponseLbServiceSubSo struct {
	Contents   []LbServiceSubSo `json:"contents,omitempty"`
	TotalCount int32                             `json:"totalCount,omitempty"`
}

type ListResponseString struct {
	Contents   []string `json:"contents,omitempty"`
	TotalCount int32    `json:"totalCount,omitempty"`
}

type LoadBalancerOpenApiControllerApiGetLoadBalancerListOpts struct {
	LoadBalancerName  optional.String
	LoadBalancerSize  optional.String
	LoadBalancerState optional.String
	Region            optional.String
	VpcName           optional.String
	CreatedBy         optional.String
	Page              optional.Int32
	Size              optional.Int32
	Sort              optional.Interface
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
