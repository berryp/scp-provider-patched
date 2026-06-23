package kubernetes

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

type ClusterRoleBindingResponse struct {
	Annotations            map[string]string         `json:"annotations,omitempty"`
	ClusterId              string                    `json:"clusterId,omitempty"`
	ClusterName            string                    `json:"clusterName,omitempty"`
	ClusterRoleBindingName string                    `json:"clusterRoleBindingName,omitempty"`
	Labels                 map[string]string         `json:"labels,omitempty"`
	Region                 string                    `json:"region,omitempty"`
	Role                   *RoleRefVo  `json:"role,omitempty"`
	RunningAge             string                    `json:"runningAge,omitempty"`
	Subjects               []SubjectVo `json:"subjects,omitempty"`
	CreatedDt              time.Time              `json:"createdDt,omitempty"`
}

type ClusterRoleBindingServiceApiListClusterRoleBindingEventsV1Opts struct {
	SystemObject optional.Bool
	Page         optional.Int32
	Size         optional.Int32
}

type ClusterRoleBindingServiceApiListClusterRoleBindingsV1Opts struct {
	ClusterRoleBindingName optional.String
	SystemObject           optional.Bool
	Page                   optional.Int32
	Size                   optional.Int32
	Sort                   optional.String
}

type ClusterRoleBindingsResponse struct {
	ProjectId              string                   `json:"projectId,omitempty"`
	ClusterRoleBindingName string                   `json:"clusterRoleBindingName,omitempty"`
	Groups                 []string                 `json:"groups,omitempty"`
	Role                   *RoleRefVo `json:"role,omitempty"`
	RunningAge             string                   `json:"runningAge,omitempty"`
	ServiceAccounts        []string                 `json:"serviceAccounts,omitempty"`
	Users                  []string                 `json:"users,omitempty"`
	CreatedDt              time.Time             `json:"createdDt,omitempty"`
}

type ClusterRoleResponse struct {
	Annotations     map[string]string            `json:"annotations,omitempty"`
	ClusterId       string                       `json:"clusterId,omitempty"`
	ClusterName     string                       `json:"clusterName,omitempty"`
	ClusterRoleName string                       `json:"clusterRoleName,omitempty"`
	Labels          map[string]string            `json:"labels,omitempty"`
	PolicyRules     []PolicyRuleVo `json:"policyRules,omitempty"`
	Region          string                       `json:"region,omitempty"`
	RunningAge      string                       `json:"runningAge,omitempty"`
	CreatedDt       time.Time                 `json:"createdDt,omitempty"`
}

type ClusterRoleServiceApiListClusterRoleEventsV1Opts struct {
	SystemObject optional.Bool
	Page         optional.Int32
	Size         optional.Int32
}

type ClusterRoleServiceApiListClusterRolesV1Opts struct {
	ClusterRoleName optional.String
	SystemObject    optional.Bool
	Page            optional.Int32
	Size            optional.Int32
	Sort            optional.String
}

type ClusterRolesResponse struct {
	ProjectId       string       `json:"projectId,omitempty"`
	ClusterRoleName string       `json:"clusterRoleName,omitempty"`
	RunningAge      string       `json:"runningAge,omitempty"`
	CreatedDt       time.Time `json:"createdDt,omitempty"`
}

type CommonCodeResponse struct {
	ClassificationCode string `json:"classificationCode,omitempty"`
	Code               string `json:"code,omitempty"`
	Name               string `json:"name,omitempty"`
	Description        string `json:"description,omitempty"`
}

type ConfigMapResponse struct {
	ProjectId           string            `json:"projectId,omitempty"`
	Annotations         map[string]string `json:"annotations,omitempty"`
	ClusterId           string            `json:"clusterId,omitempty"`
	ClusterName         string            `json:"clusterName,omitempty"`
	ConfigMapBinaryData map[string]string `json:"configMapBinaryData,omitempty"`
	ConfigMapData       map[string]string `json:"configMapData,omitempty"`
	ConfigMapName       string            `json:"configMapName,omitempty"`
	Labels              map[string]string `json:"labels,omitempty"`
	NamespaceName       string            `json:"namespaceName,omitempty"`
	Region              string            `json:"region,omitempty"`
	RunningAge          string            `json:"runningAge,omitempty"`
	CreatedDt           time.Time      `json:"createdDt,omitempty"`
}

type ConfigMapServiceApiListConfigmapEventsV2Opts struct {
	SystemObject optional.Bool
	Page         optional.Int32
	Size         optional.Int32
}

type ConfigMapServiceApiListConfigmapsV2Opts struct {
	ConfigMapName optional.String
	SystemObject  optional.Bool
	Page          optional.Int32
	Size          optional.Int32
	Sort          optional.String
}

type ConfigMapsResponse struct {
	ProjectId          string       `json:"projectId,omitempty"`
	ConfigMapDataCount string       `json:"configMapDataCount,omitempty"`
	ConfigMapName      string       `json:"configMapName,omitempty"`
	RunningAge         string       `json:"runningAge,omitempty"`
	CreatedDt          time.Time `json:"createdDt,omitempty"`
}

type CronJobResponse struct {
	ProjectId                  string            `json:"projectId,omitempty"`
	ActiveJobCount             int32             `json:"activeJobCount,omitempty"`
	ActiveJobs                 []string          `json:"activeJobs,omitempty"`
	Annotations                map[string]string `json:"annotations,omitempty"`
	ClusterId                  string            `json:"clusterId,omitempty"`
	ClusterName                string            `json:"clusterName,omitempty"`
	Completions                int32             `json:"completions,omitempty"`
	ConcurrencyPolicy          string            `json:"concurrencyPolicy,omitempty"`
	CronJobName                string            `json:"cronJobName,omitempty"`
	FailedJobsHistoryLimit     int32             `json:"failedJobsHistoryLimit,omitempty"`
	Labels                     map[string]string `json:"labels,omitempty"`
	LastScheduleTime           time.Time      `json:"lastScheduleTime,omitempty"`
	NamespaceName              string            `json:"namespaceName,omitempty"`
	Parallelism                int32             `json:"parallelism,omitempty"`
	Region                     string            `json:"region,omitempty"`
	RunningAge                 string            `json:"runningAge,omitempty"`
	Schedule                   string            `json:"schedule,omitempty"`
	SelectorDetails            []string          `json:"selectorDetails,omitempty"`
	Selectors                  map[string]string `json:"selectors,omitempty"`
	StartingDeadlineSeconds    int64             `json:"startingDeadlineSeconds,omitempty"`
	SuccessfulJobsHistoryLimit int32             `json:"successfulJobsHistoryLimit,omitempty"`
	Suspended                  *bool             `json:"suspended,omitempty"`
	CreatedDt                  time.Time      `json:"createdDt,omitempty"`
}

type CronJobServiceApiListChildJobsV2Opts struct {
	SystemObject optional.Bool
	Page         optional.Int32
	Size         optional.Int32
}

type CronJobServiceApiListCronJobEventsV2Opts struct {
	SystemObject optional.Bool
	Page         optional.Int32
	Size         optional.Int32
}

type CronJobServiceApiListCronJobsV2Opts struct {
	CronJobName  optional.String
	SystemObject optional.Bool
	Page         optional.Int32
	Size         optional.Int32
	Sort         optional.String
}

type DaemonSetResponse struct {
	ProjectId              string            `json:"projectId,omitempty"`
	ActiveDaemonSet        int32             `json:"activeDaemonSet,omitempty"`
	Annotations            map[string]string `json:"annotations,omitempty"`
	ClusterId              string            `json:"clusterId,omitempty"`
	ClusterName            string            `json:"clusterName,omitempty"`
	CurrentNumberScheduled int32             `json:"currentNumberScheduled,omitempty"`
	DaemonSetName          string            `json:"daemonSetName,omitempty"`
	DesiredNumberScheduled int32             `json:"desiredNumberScheduled,omitempty"`
	Labels                 map[string]string `json:"labels,omitempty"`
	LastScheduleTime       string            `json:"lastScheduleTime,omitempty"`
	NamespaceName          string            `json:"namespaceName,omitempty"`
	NodeSelector           map[string]string `json:"nodeSelector,omitempty"`
	NumberAvailable        int32             `json:"numberAvailable,omitempty"`
	NumberMisscheduled     int32             `json:"numberMisscheduled,omitempty"`
	NumberReady            int32             `json:"numberReady,omitempty"`
	Region                 string            `json:"region,omitempty"`
	RunningAge             string            `json:"runningAge,omitempty"`
	Schedule               string            `json:"schedule,omitempty"`
	SelectorDetails        []string          `json:"selectorDetails,omitempty"`
	Selectors              map[string]string `json:"selectors,omitempty"`
	UpdatedNumberScheduled int32             `json:"updatedNumberScheduled,omitempty"`
	CreatedDt              time.Time      `json:"createdDt,omitempty"`
}

type DaemonSetServiceApiListDaemonSetEventsV2Opts struct {
	SystemObject optional.Bool
	Page         optional.Int32
	Size         optional.Int32
}

type DaemonSetServiceApiListDaemonSetPodsV2Opts struct {
	SystemObject optional.Bool
	Page         optional.Int32
	Size         optional.Int32
}

type DaemonSetServiceApiListDaemonSetsV2Opts struct {
	DaemonSetName optional.String
	SystemObject  optional.Bool
	Page          optional.Int32
	Size          optional.Int32
	Sort          optional.String
}

type DaemonSetsResponse struct {
	ProjectId              string            `json:"projectId,omitempty"`
	ActiveDaemonSet        int32             `json:"activeDaemonSet,omitempty"`
	CurrentNumberScheduled int32             `json:"currentNumberScheduled,omitempty"`
	DaemonSetName          string            `json:"daemonSetName,omitempty"`
	DesiredNumberScheduled int32             `json:"desiredNumberScheduled,omitempty"`
	LastScheduleTime       string            `json:"lastScheduleTime,omitempty"`
	NamespaceName          string            `json:"namespaceName,omitempty"`
	NodeSelector           map[string]string `json:"nodeSelector,omitempty"`
	NumberAvailable        int32             `json:"numberAvailable,omitempty"`
	NumberReady            int32             `json:"numberReady,omitempty"`
	RunningAge             string            `json:"runningAge,omitempty"`
	Schedule               string            `json:"schedule,omitempty"`
	UpdatedNumberScheduled int32             `json:"updatedNumberScheduled,omitempty"`
	CreatedDt              time.Time      `json:"createdDt,omitempty"`
}

type DeploymentResponse struct {
	Annotations         map[string]string `json:"annotations,omitempty"`
	AvailableReplicas   string            `json:"availableReplicas,omitempty"`
	ClusterId           string            `json:"clusterId,omitempty"`
	ClusterName         string            `json:"clusterName,omitempty"`
	DeploymentName      string            `json:"deploymentName,omitempty"`
	Labels              map[string]string `json:"labels,omitempty"`
	MaxSurge            string            `json:"maxSurge,omitempty"`
	MaxUnavailable      string            `json:"maxUnavailable,omitempty"`
	MinReadySeconds     string            `json:"minReadySeconds,omitempty"`
	Namespace           string            `json:"namespace,omitempty"`
	ReadyReplicas       string            `json:"readyReplicas,omitempty"`
	Region              string            `json:"region,omitempty"`
	RunningAge          string            `json:"runningAge,omitempty"`
	Selector            map[string]string `json:"selector,omitempty"`
	SelectorDetails     []string          `json:"selectorDetails,omitempty"`
	SpecReplicas        string            `json:"specReplicas,omitempty"`
	StatusReplicas      string            `json:"statusReplicas,omitempty"`
	StrategyType        string            `json:"strategyType,omitempty"`
	UnavailableReplicas string            `json:"unavailableReplicas,omitempty"`
	UpdatedReplicas     string            `json:"updatedReplicas,omitempty"`
	CreatedDt           time.Time      `json:"createdDt,omitempty"`
}

type DeploymentServiceApiListDeploymentEventsV2Opts struct {
	SystemObject optional.Bool
	Page         optional.Int32
	Size         optional.Int32
}

type DeploymentServiceApiListDeploymentPodsV2Opts struct {
	SystemObject optional.Bool
	Page         optional.Int32
	Size         optional.Int32
}

type DeploymentServiceApiListDeploymentsV2Opts struct {
	DeploymentName optional.String
	SystemObject   optional.Bool
	Page           optional.Int32
	Size           optional.Int32
	Sort           optional.String
}

type DeploymentsResponse struct {
	AvailableReplicas string       `json:"availableReplicas,omitempty"`
	DeploymentName    string       `json:"deploymentName,omitempty"`
	Namespace         string       `json:"namespace,omitempty"`
	ReadyReplicas     string       `json:"readyReplicas,omitempty"`
	RunningAge        string       `json:"runningAge,omitempty"`
	SpecReplicas      string       `json:"specReplicas,omitempty"`
	UpdatedReplicas   string       `json:"updatedReplicas,omitempty"`
	CreatedDt         time.Time `json:"createdDt,omitempty"`
}

type EndpointPortVo struct {
	Port     string `json:"port,omitempty"`
	PortName string `json:"portName,omitempty"`
	Protocol string `json:"protocol,omitempty"`
}

type EndpointResponse struct {
	ProjectId     string                         `json:"projectId,omitempty"`
	Annotations   map[string]string              `json:"annotations,omitempty"`
	EndpointName  string                         `json:"endpointName,omitempty"`
	Labels        map[string]string              `json:"labels,omitempty"`
	NamespaceName string                         `json:"namespaceName,omitempty"`
	Region        string                         `json:"region,omitempty"`
	Subsets       []SubsetResponse `json:"subsets,omitempty"`
	CreatedDt     time.Time                   `json:"createdDt,omitempty"`
}

type EndpointServiceApiListEndpointEventsV1Opts struct {
	SystemObject optional.Bool
	Page         optional.Int32
	Size         optional.Int32
}

type EndpointServiceApiListEndpointV1Opts struct {
	EndpointName optional.String
	SystemObject optional.Bool
	Page         optional.Int32
	Size         optional.Int32
	Sort         optional.String
}

type EndpointsResponse struct {
	EndpointName string                         `json:"endpointName,omitempty"`
	RunningAge   string                         `json:"runningAge,omitempty"`
	Subsets      []SubsetResponse `json:"subsets,omitempty"`
	CreatedDt    time.Time                   `json:"createdDt,omitempty"`
}

type EventsResponse struct {
	ProjectId          string       `json:"projectId,omitempty"`
	Age                string       `json:"age,omitempty"`
	EventComponent     string       `json:"eventComponent,omitempty"`
	EventCount         int32        `json:"eventCount,omitempty"`
	EventFirstSeenDate time.Time `json:"eventFirstSeenDate,omitempty"`
	EventHost          string       `json:"eventHost,omitempty"`
	EventKind          string       `json:"eventKind,omitempty"`
	EventLastSeenDate  time.Time `json:"eventLastSeenDate,omitempty"`
	EventMessage       string       `json:"eventMessage,omitempty"`
	EventName          string       `json:"eventName,omitempty"`
	EventReason        string       `json:"eventReason,omitempty"`
	EventType          string       `json:"eventType,omitempty"`
	NamespaceName      string       `json:"namespaceName,omitempty"`
}

type IngressClassResponse struct {
	ProjectId        string                           `json:"projectId,omitempty"`
	Annotations      map[string]string                `json:"annotations,omitempty"`
	Controller       string                           `json:"controller,omitempty"`
	IngressClassName string                           `json:"ingressClassName,omitempty"`
	IsDefaultClass   string                           `json:"isDefaultClass,omitempty"`
	Labels           map[string]string                `json:"labels,omitempty"`
	Parameter        *ParameterResponse `json:"parameter,omitempty"`
	Region           string                           `json:"region,omitempty"`
	CreatedDt        time.Time                     `json:"createdDt,omitempty"`
}

type IngressClassServiceApiListIngressClassEventsV1Opts struct {
	SystemObject optional.Bool
	Page         optional.Int32
	Size         optional.Int32
}

type IngressClassServiceApiListIngressClassV1Opts struct {
	IngressClassName optional.String
	SystemObject     optional.Bool
	Page             optional.Int32
	Size             optional.Int32
	Sort             optional.String
}

type IngressClassesResponse struct {
	Controller       string                           `json:"controller,omitempty"`
	IngressClassName string                           `json:"ingressClassName,omitempty"`
	IsDefaultClass   string                           `json:"isDefaultClass,omitempty"`
	Parameter        *ParameterResponse `json:"parameter,omitempty"`
	RunningAge       string                           `json:"runningAge,omitempty"`
	CreatedDt        time.Time                     `json:"createdDt,omitempty"`
}

type IngressResponse struct {
	ProjectId             string                              `json:"projectId,omitempty"`
	Annotations           map[string]string                   `json:"annotations,omitempty"`
	ClusterId             string                              `json:"clusterId,omitempty"`
	ClusterName           string                              `json:"clusterName,omitempty"`
	IngressClass          string                              `json:"ingressClass,omitempty"`
	IngressDefaultBackend string                              `json:"ingressDefaultBackend,omitempty"`
	IngressIp             []string                            `json:"ingressIp,omitempty"`
	IngressName           string                              `json:"ingressName,omitempty"`
	IngressRules          []IngressRuleResponse `json:"ingressRules,omitempty"`
	IngressTlses          []IngressTlsResponse  `json:"ingressTlses,omitempty"`
	Labels                map[string]string                   `json:"labels,omitempty"`
	NamespaceName         string                              `json:"namespaceName,omitempty"`
	Region                string                              `json:"region,omitempty"`
	RunningAge            string                              `json:"runningAge,omitempty"`
	CreatedDt             time.Time                        `json:"createdDt,omitempty"`
}

type IngressRulePathResponse struct {
	K8sServiceName     string `json:"k8sServiceName,omitempty"`
	K8sServicePort     string `json:"k8sServicePort,omitempty"`
	K8sServicePortName string `json:"k8sServicePortName,omitempty"`
	Path               string `json:"path,omitempty"`
}

type IngressRuleResponse struct {
	Host             string                                  `json:"host,omitempty"`
	IngressRulePaths []IngressRulePathResponse `json:"ingressRulePaths,omitempty"`
}

type IngressServiceApiListIngressEventsV2Opts struct {
	SystemObject optional.Bool
	Page         optional.Int32
	Size         optional.Int32
}

type IngressServiceApiListIngressV2Opts struct {
	IngressName  optional.String
	SystemObject optional.Bool
	Page         optional.Int32
	Size         optional.Int32
	Sort         optional.String
}

type IngressTlsResponse struct {
	Hosts      []string `json:"hosts,omitempty"`
	SecretName string   `json:"secretName,omitempty"`
}

type IngressesResponse struct {
	ProjectId    string       `json:"projectId,omitempty"`
	IngressClass string       `json:"ingressClass,omitempty"`
	IngressHost  []string     `json:"ingressHost,omitempty"`
	IngressIP    []string     `json:"ingressIP,omitempty"`
	IngressName  string       `json:"ingressName,omitempty"`
	IngressPort  string       `json:"ingressPort,omitempty"`
	RunningAge   string       `json:"runningAge,omitempty"`
	CreatedDt    time.Time `json:"createdDt,omitempty"`
}

type JobResponse struct {
	ProjectId       string            `json:"projectId,omitempty"`
	Age             string            `json:"age,omitempty"`
	Annotations     map[string]string `json:"annotations,omitempty"`
	BackOffLimit    int32             `json:"backOffLimit,omitempty"`
	ClusterId       string            `json:"clusterId,omitempty"`
	ClusterName     string            `json:"clusterName,omitempty"`
	Completion      string            `json:"completion,omitempty"`
	CompletionMode  string            `json:"completionMode,omitempty"`
	CompletionTime  time.Time      `json:"completionTime,omitempty"`
	Completions     int32             `json:"completions,omitempty"`
	EndDt           time.Time      `json:"endDt,omitempty"`
	JobName         string            `json:"jobName,omitempty"`
	JobState        string            `json:"jobState,omitempty"`
	Labels          map[string]string `json:"labels,omitempty"`
	NamespaceName   string            `json:"namespaceName,omitempty"`
	Parallelism     int32             `json:"parallelism,omitempty"`
	Region          string            `json:"region,omitempty"`
	RunningAge      string            `json:"runningAge,omitempty"`
	SelectorDetails []string          `json:"selectorDetails,omitempty"`
	Selectors       map[string]string `json:"selectors,omitempty"`
	StartTime       time.Time      `json:"startTime,omitempty"`
	CreatedDt       time.Time      `json:"createdDt,omitempty"`
}

type JobServiceApiListJobEventsV2Opts struct {
	SystemObject optional.Bool
	Page         optional.Int32
	Size         optional.Int32
}

type JobServiceApiListJobPodsV2Opts struct {
	SystemObject optional.Bool
	Page         optional.Int32
	Size         optional.Int32
}

type JobServiceApiListJobsV2Opts struct {
	JobName      optional.String
	SystemObject optional.Bool
	Page         optional.Int32
	Size         optional.Int32
	Sort         optional.String
}

type JobsResponse struct {
	ProjectId     string            `json:"projectId,omitempty"`
	Age           string            `json:"age,omitempty"`
	Annotations   map[string]string `json:"annotations,omitempty"`
	EndDt         time.Time      `json:"endDt,omitempty"`
	JobCompletion string            `json:"jobCompletion,omitempty"`
	JobName       string            `json:"jobName,omitempty"`
	Labels        map[string]string `json:"labels,omitempty"`
	NamespaceName string            `json:"namespaceName,omitempty"`
	RunningAge    string            `json:"runningAge,omitempty"`
	Selectors     map[string]string `json:"selectors,omitempty"`
	CreatedDt     time.Time      `json:"createdDt,omitempty"`
}

type K8SSvcPortVo struct {
	NodePort   string `json:"nodePort,omitempty"`
	Port       string `json:"port,omitempty"`
	PortName   string `json:"portName,omitempty"`
	Protocol   string `json:"protocol,omitempty"`
	TargetPort string `json:"targetPort,omitempty"`
}

type K8sObjectCreateRequest struct {
	Yaml string `json:"yaml"`
}

type K8sObjectDeleteRequest struct {
	Names []string `json:"names"`
}

type K8sObjectDetailResponse struct {
	ProjectId     string `json:"projectId,omitempty"`
	KubectlResult string `json:"kubectlResult,omitempty"`
}

type K8sObjectResponse struct {
	ProjectId     string `json:"projectId,omitempty"`
	KubectlResult string `json:"kubectlResult,omitempty"`
	Yaml          string `json:"yaml,omitempty"`
}

type K8sObjectUpdateRequest struct {
	Yaml string `json:"yaml"`
}

type K8sServiceApiListK8sServiceEventsV2Opts struct {
	SystemObject optional.Bool
	Page         optional.Int32
	Size         optional.Int32
}

type K8sServiceApiListK8sServiceV2Opts struct {
	K8sServiceName optional.String
	SystemObject   optional.Bool
	Page           optional.Int32
	Size           optional.Int32
	Sort           optional.String
}

type K8sSvcResponse struct {
	Annotations           map[string]string            `json:"annotations,omitempty"`
	ClusterIP             string                       `json:"clusterIP,omitempty"`
	ClusterId             string                       `json:"clusterId,omitempty"`
	ClusterName           string                       `json:"clusterName,omitempty"`
	ExternalIPs           []string                     `json:"externalIPs,omitempty"`
	ExternalTrafficPolicy string                       `json:"externalTrafficPolicy,omitempty"`
	Labels                map[string]string            `json:"labels,omitempty"`
	Ports                 []K8SSvcPortVo `json:"ports,omitempty"`
	Region                string                       `json:"region,omitempty"`
	Selectors             map[string]string            `json:"selectors,omitempty"`
	ServiceName           string                       `json:"serviceName,omitempty"`
	ServiceType           string                       `json:"serviceType,omitempty"`
	SessionAffinity       string                       `json:"sessionAffinity,omitempty"`
	CreatedDt             time.Time                 `json:"createdDt,omitempty"`
}

type K8sSvcsResponse struct {
	ClusterIP   string                       `json:"clusterIP,omitempty"`
	ExternalIPs []string                     `json:"externalIPs,omitempty"`
	Ports       []K8SSvcPortVo `json:"ports,omitempty"`
	RunningAge  string                       `json:"runningAge,omitempty"`
	ServiceName string                       `json:"serviceName,omitempty"`
	ServiceType string                       `json:"serviceType,omitempty"`
	CreatedDt   time.Time                 `json:"createdDt,omitempty"`
}

type NamespaceResponse struct {
	ProjectId       string       `json:"projectId,omitempty"`
	Age             string       `json:"age,omitempty"`
	Annotations     []string     `json:"annotations,omitempty"`
	ClusterId       string       `json:"clusterId,omitempty"`
	ClusterName     string       `json:"clusterName,omitempty"`
	Labels          []string     `json:"labels,omitempty"`
	NamespaceName   string       `json:"namespaceName,omitempty"`
	NamespaceStatus string       `json:"namespaceStatus,omitempty"`
	Region          string       `json:"region,omitempty"`
	CreatedDt       time.Time `json:"createdDt,omitempty"`
}

type NamespaceServiceApiListNamespaceEventsV2Opts struct {
	SystemObject optional.Bool
	Page         optional.Int32
	Size         optional.Int32
}

type NamespaceServiceApiListNamespaceV2Opts struct {
	NamespaceName optional.String
	SystemObject  optional.Bool
	Page          optional.Int32
	Size          optional.Int32
	Sort          optional.String
}

type NamespacesResponse struct {
	ProjectId       string       `json:"projectId,omitempty"`
	Age             string       `json:"age,omitempty"`
	NamespaceName   string       `json:"namespaceName,omitempty"`
	NamespaceStatus string       `json:"namespaceStatus,omitempty"`
	CreatedDt       time.Time `json:"createdDt,omitempty"`
}

type NodeCapacityResponse struct {
	Cpu              string `json:"cpu,omitempty"`
	EphemeralStorage string `json:"ephemeralStorage,omitempty"`
	Gpu              string `json:"gpu,omitempty"`
	Hugepages1Gi     string `json:"hugepages1Gi,omitempty"`
	Hugepages2Mi     string `json:"hugepages2Mi,omitempty"`
	Memory           string `json:"memory,omitempty"`
	NodeCapacityType string `json:"nodeCapacityType,omitempty"`
	PodCount         string `json:"podCount,omitempty"`
}

type NodeConditionResponse struct {
	LastHeartbeatTime    time.Time `json:"lastHeartbeatTime,omitempty"`
	LastTransitionTime   time.Time `json:"lastTransitionTime,omitempty"`
	NodeConditionMessage string       `json:"nodeConditionMessage,omitempty"`
	NodeConditionReason  string       `json:"nodeConditionReason,omitempty"`
	NodeConditionStatus  string       `json:"nodeConditionStatus,omitempty"`
	NodeConditionType    string       `json:"nodeConditionType,omitempty"`
}

type NodeResponse struct {
	ProjectId         string                                `json:"projectId,omitempty"`
	Age               string                                `json:"age,omitempty"`
	Annotations       []string                              `json:"annotations,omitempty"`
	Architecture      string                                `json:"architecture,omitempty"`
	BootId            string                                `json:"bootId,omitempty"`
	ClusterId         string                                `json:"clusterId,omitempty"`
	ClusterName       string                                `json:"clusterName,omitempty"`
	ContainerVersion  string                                `json:"containerVersion,omitempty"`
	ExternalIpAddress string                                `json:"externalIpAddress,omitempty"`
	HostName          string                                `json:"hostName,omitempty"`
	Ip                string                                `json:"ip,omitempty"`
	K8sVersion        string                                `json:"k8sVersion,omitempty"`
	KernelVersion     string                                `json:"kernelVersion,omitempty"`
	KubeProxyVersion  string                                `json:"kubeProxyVersion,omitempty"`
	KubeletVersion    string                                `json:"kubeletVersion,omitempty"`
	Labels            []string                              `json:"labels,omitempty"`
	MachineId         string                                `json:"machineId,omitempty"`
	NodeCapacities    []NodeCapacityResponse  `json:"nodeCapacities,omitempty"`
	NodeConditions    []NodeConditionResponse `json:"nodeConditions,omitempty"`
	NodeName          string                                `json:"nodeName,omitempty"`
	NodeStatus        string                                `json:"nodeStatus,omitempty"`
	OperatingSystem   string                                `json:"operatingSystem,omitempty"`
	OsImage           string                                `json:"osImage,omitempty"`
	PodCidr           string                                `json:"podCidr,omitempty"`
	Region            string                                `json:"region,omitempty"`
	SystemUuid        string                                `json:"systemUuid,omitempty"`
	Taints            []string                              `json:"taints,omitempty"`
	CreatedBy         string                                `json:"createdBy,omitempty"`
	CreatedDt         time.Time                          `json:"createdDt,omitempty"`
}

type NodeServiceApiListNodeEventsV2Opts struct {
	SystemObject optional.Bool
	Page         optional.Int32
	Size         optional.Int32
}

type NodeServiceApiListNodesV2Opts struct {
	NodeName     optional.String
	SystemObject optional.Bool
	Page         optional.Int32
	Size         optional.Int32
	Sort         optional.String
}

type NodeServiceApiListPodsInNodeV2Opts struct {
	SystemObject optional.Bool
	Page         optional.Int32
	Size         optional.Int32
}

type NodesResponse struct {
	ProjectId         string       `json:"projectId,omitempty"`
	Age               string       `json:"age,omitempty"`
	ContainerVersion  string       `json:"containerVersion,omitempty"`
	ExternalIpAddress string       `json:"externalIpAddress,omitempty"`
	Ip                string       `json:"ip,omitempty"`
	K8sVersion        string       `json:"k8sVersion,omitempty"`
	KernelVersion     string       `json:"kernelVersion,omitempty"`
	NodeName          string       `json:"nodeName,omitempty"`
	NodeStatus        string       `json:"nodeStatus,omitempty"`
	OsImage           string       `json:"osImage,omitempty"`
	CreatedDt         time.Time `json:"createdDt,omitempty"`
}

type PageResponse struct {
	Contents   []interface{} `json:"contents,omitempty"`
	TotalCount int64         `json:"totalCount,omitempty"`
}

type PageResponseClusterRoleBindingsResponse struct {
	Contents   []ClusterRoleBindingsResponse `json:"contents,omitempty"`
	TotalCount int64                                       `json:"totalCount,omitempty"`
}

type PageResponseClusterRolesResponse struct {
	Contents   []ClusterRolesResponse `json:"contents,omitempty"`
	TotalCount int64                                `json:"totalCount,omitempty"`
}

type PageResponseConfigMapsResponse struct {
	Contents   []ConfigMapsResponse `json:"contents,omitempty"`
	TotalCount int64                              `json:"totalCount,omitempty"`
}

type PageResponseCronJobResponse struct {
	Contents   []CronJobResponse `json:"contents,omitempty"`
	TotalCount int64                           `json:"totalCount,omitempty"`
}

type PageResponseDaemonSetsResponse struct {
	Contents   []DaemonSetsResponse `json:"contents,omitempty"`
	TotalCount int64                              `json:"totalCount,omitempty"`
}

type PageResponseDeploymentsResponse struct {
	Contents   []DeploymentsResponse `json:"contents,omitempty"`
	TotalCount int64                               `json:"totalCount,omitempty"`
}

type PageResponseEndpointsResponse struct {
	Contents   []EndpointsResponse `json:"contents,omitempty"`
	TotalCount int64                             `json:"totalCount,omitempty"`
}

type PageResponseEventsResponse struct {
	Contents   []EventsResponse `json:"contents,omitempty"`
	TotalCount int64                          `json:"totalCount,omitempty"`
}

type PageResponseIngressClassesResponse struct {
	Contents   []IngressClassesResponse `json:"contents,omitempty"`
	TotalCount int64                                  `json:"totalCount,omitempty"`
}

type PageResponseIngressesResponse struct {
	Contents   []IngressesResponse `json:"contents,omitempty"`
	TotalCount int64                             `json:"totalCount,omitempty"`
}

type PageResponseJobsResponse struct {
	Contents   []JobsResponse `json:"contents,omitempty"`
	TotalCount int64                        `json:"totalCount,omitempty"`
}

type PageResponseK8sSvcsResponse struct {
	Contents   []K8sSvcsResponse `json:"contents,omitempty"`
	TotalCount int64                           `json:"totalCount,omitempty"`
}

type PageResponseNamespacesResponse struct {
	Contents   []NamespacesResponse `json:"contents,omitempty"`
	TotalCount int64                              `json:"totalCount,omitempty"`
}

type PageResponseNodesResponse struct {
	Contents   []NodesResponse `json:"contents,omitempty"`
	TotalCount int64                         `json:"totalCount,omitempty"`
}

type PageResponsePodsResponse struct {
	Contents   []PodsResponse `json:"contents,omitempty"`
	TotalCount int64                        `json:"totalCount,omitempty"`
}

type PageResponsePvcsResponse struct {
	Contents   []PvcsResponse `json:"contents,omitempty"`
	TotalCount int64                        `json:"totalCount,omitempty"`
}

type PageResponsePvsResponse struct {
	Contents   []PvsResponse `json:"contents,omitempty"`
	TotalCount int64                       `json:"totalCount,omitempty"`
}

type PageResponseRoleBindingsResponse struct {
	Contents   []RoleBindingsResponse `json:"contents,omitempty"`
	TotalCount int64                                `json:"totalCount,omitempty"`
}

type PageResponseRolesResponse struct {
	Contents   []RolesResponse `json:"contents,omitempty"`
	TotalCount int64                         `json:"totalCount,omitempty"`
}

type PageResponseSecretsResponse struct {
	Contents   []SecretsResponse `json:"contents,omitempty"`
	TotalCount int64                           `json:"totalCount,omitempty"`
}

type PageResponseStatefulSetsResponse struct {
	Contents   []StatefulSetsResponse `json:"contents,omitempty"`
	TotalCount int64                                `json:"totalCount,omitempty"`
}

type PageResponseStorageClassesResponse struct {
	Contents   []StorageClassesResponse `json:"contents,omitempty"`
	TotalCount int64                                  `json:"totalCount,omitempty"`
}

type ParameterResponse struct {
	ApiGroup string `json:"apiGroup,omitempty"`
	Kind     string `json:"kind,omitempty"`
	Name     string `json:"name,omitempty"`
	Scope    string `json:"scope,omitempty"`
}

type PersistentVolumeClaimServiceApiListPersistentVolumeClaimEventsV2Opts struct {
	SystemObject optional.Bool
	Page         optional.Int32
	Size         optional.Int32
}

type PersistentVolumeClaimServiceApiListPersistentVolumeClaimsV2Opts struct {
	PvcName      optional.String
	SystemObject optional.Bool
	Page         optional.Int32
	Size         optional.Int32
	Sort         optional.String
}

type PersistentVolumeServiceApiListPersistentVolumeEventsV2Opts struct {
	SystemObject optional.Bool
	Page         optional.Int32
	Size         optional.Int32
}

type PersistentVolumeServiceApiListPersistentVolumesV2Opts struct {
	PvName       optional.String
	SystemObject optional.Bool
	Page         optional.Int32
	Size         optional.Int32
	Sort         optional.String
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

type PodContainerResponse struct {
	ContainerName string `json:"containerName,omitempty"`
	Restarts      string `json:"restarts,omitempty"`
	Status        string `json:"status,omitempty"`
}

type PodLogResponse struct {
	PodLog string `json:"podLog,omitempty"`
}

type PodResponse struct {
	ProjectId         string                               `json:"projectId,omitempty"`
	Annotations       map[string]string                    `json:"annotations,omitempty"`
	ClusterId         string                               `json:"clusterId,omitempty"`
	ClusterName       string                               `json:"clusterName,omitempty"`
	ContainerList     []PodContainerResponse `json:"containerList,omitempty"`
	InitContainerList []PodContainerResponse `json:"initContainerList,omitempty"`
	Labels            map[string]string                    `json:"labels,omitempty"`
	Namespace         string                               `json:"namespace,omitempty"`
	NodeIp            string                               `json:"nodeIp,omitempty"`
	NodeName          string                               `json:"nodeName,omitempty"`
	NodeSelector      map[string]string                    `json:"nodeSelector,omitempty"`
	PodIp             string                               `json:"podIp,omitempty"`
	PodName           string                               `json:"podName,omitempty"`
	PodStatus         string                               `json:"podStatus,omitempty"`
	Priority          int32                                `json:"priority,omitempty"`
	PriorityClassName string                               `json:"priorityClassName,omitempty"`
	QosClass          string                               `json:"qosClass,omitempty"`
	Region            string                               `json:"region,omitempty"`
	StartTime         time.Time                         `json:"startTime,omitempty"`
	CreatedDt         time.Time                         `json:"createdDt,omitempty"`
}

type PodServiceApiGetPodLogDownloadOpts struct {
	ContainerName optional.String
	Previous      optional.Bool
}

type PodServiceApiGetPodLogOpts struct {
	ContainerName optional.String
	Previous      optional.Bool
}

type PodServiceApiListPodEventsV2Opts struct {
	SystemObject optional.Bool
	Page         optional.Int32
	Size         optional.Int32
}

type PodServiceApiListPodsV2Opts struct {
	PodName      optional.String
	SystemObject optional.Bool
	Page         optional.Int32
	Size         optional.Int32
	Sort         optional.String
}

type PodsResponse struct {
	ProjectId       string `json:"projectId,omitempty"`
	Age             string `json:"age,omitempty"`
	ContainerStatus string `json:"containerStatus,omitempty"`
	Namespace       string `json:"namespace,omitempty"`
	NodeName        string `json:"nodeName,omitempty"`
	PodIp           string `json:"podIp,omitempty"`
	PodName         string `json:"podName,omitempty"`
	PodStatus       string `json:"podStatus,omitempty"`
	Restarts        string `json:"restarts,omitempty"`
}

type PolicyRuleVo struct {
	ApiGroupForSort string   `json:"apiGroupForSort,omitempty"`
	NonResourceUrls []string `json:"nonResourceUrls,omitempty"`
	Resource        string   `json:"resource,omitempty"`
	ResourceNames   []string `json:"resourceNames,omitempty"`
	Verbs           []string `json:"verbs,omitempty"`
}

type PvResponse struct {
	ProjectId        string            `json:"projectId,omitempty"`
	AccessModes      []string          `json:"accessModes,omitempty"`
	Annotations      map[string]string `json:"annotations,omitempty"`
	ClaimRefName     string            `json:"claimRefName,omitempty"`
	ClusterId        string            `json:"clusterId,omitempty"`
	ClusterName      string            `json:"clusterName,omitempty"`
	Labels           map[string]string `json:"labels,omitempty"`
	PvName           string            `json:"pvName,omitempty"`
	PvReclaimPolicy  string            `json:"pvReclaimPolicy,omitempty"`
	PvSize           string            `json:"pvSize,omitempty"`
	PvStatus         string            `json:"pvStatus,omitempty"`
	PvVolumeMode     string            `json:"pvVolumeMode,omitempty"`
	Region           string            `json:"region,omitempty"`
	StorageClassName string            `json:"storageClassName,omitempty"`
	CreatedDt        time.Time      `json:"createdDt,omitempty"`
}

type PvcResponse struct {
	ProjectId        string            `json:"projectId,omitempty"`
	AccessModes      []string          `json:"accessModes,omitempty"`
	Annotations      map[string]string `json:"annotations,omitempty"`
	ClusterId        string            `json:"clusterId,omitempty"`
	ClusterName      string            `json:"clusterName,omitempty"`
	Labels           map[string]string `json:"labels,omitempty"`
	NamespaceName    string            `json:"namespaceName,omitempty"`
	PvName           string            `json:"pvName,omitempty"`
	PvVolumeMode     string            `json:"pvVolumeMode,omitempty"`
	PvcName          string            `json:"pvcName,omitempty"`
	PvcSize          string            `json:"pvcSize,omitempty"`
	PvcStatus        string            `json:"pvcStatus,omitempty"`
	Region           string            `json:"region,omitempty"`
	StorageClassName string            `json:"storageClassName,omitempty"`
	CreatedDt        time.Time      `json:"createdDt,omitempty"`
}

type PvcsResponse struct {
	ProjectId        string   `json:"projectId,omitempty"`
	AccessModes      []string `json:"accessModes,omitempty"`
	PvName           string   `json:"pvName,omitempty"`
	PvcName          string   `json:"pvcName,omitempty"`
	PvcSize          string   `json:"pvcSize,omitempty"`
	PvcStatus        string   `json:"pvcStatus,omitempty"`
	RunningAge       string   `json:"runningAge,omitempty"`
	StorageClassName string   `json:"storageClassName,omitempty"`
}

type PvsResponse struct {
	ProjectId        string   `json:"projectId,omitempty"`
	AccessModes      []string `json:"accessModes,omitempty"`
	ClaimRefName     string   `json:"claimRefName,omitempty"`
	PvName           string   `json:"pvName,omitempty"`
	PvReason         string   `json:"pvReason,omitempty"`
	PvReclaimPolicy  string   `json:"pvReclaimPolicy,omitempty"`
	PvSize           string   `json:"pvSize,omitempty"`
	PvStatus         string   `json:"pvStatus,omitempty"`
	RunningAge       string   `json:"runningAge,omitempty"`
	StorageClassName string   `json:"storageClassName,omitempty"`
}

type RoleBindingResponse struct {
	Annotations     map[string]string         `json:"annotations,omitempty"`
	ClusterId       string                    `json:"clusterId,omitempty"`
	ClusterName     string                    `json:"clusterName,omitempty"`
	Labels          map[string]string         `json:"labels,omitempty"`
	NamespaceName   string                    `json:"namespaceName,omitempty"`
	Region          string                    `json:"region,omitempty"`
	Role            *RoleRefVo  `json:"role,omitempty"`
	RoleBindingName string                    `json:"roleBindingName,omitempty"`
	RunningAge      string                    `json:"runningAge,omitempty"`
	Subjects        []SubjectVo `json:"subjects,omitempty"`
	CreatedDt       time.Time              `json:"createdDt,omitempty"`
}

type RoleBindingServiceApiListRoleBindingEventsV1Opts struct {
	SystemObject optional.Bool
	Page         optional.Int32
	Size         optional.Int32
}

type RoleBindingServiceApiListRoleBindingsV1Opts struct {
	RoleBindingName optional.String
	SystemObject    optional.Bool
	Page            optional.Int32
	Size            optional.Int32
	Sort            optional.String
}

type RoleBindingsResponse struct {
	ProjectId       string                   `json:"projectId,omitempty"`
	Groups          []string                 `json:"groups,omitempty"`
	NamespaceName   string                   `json:"namespaceName,omitempty"`
	Role            *RoleRefVo `json:"role,omitempty"`
	RoleBindingName string                   `json:"roleBindingName,omitempty"`
	RunningAge      string                   `json:"runningAge,omitempty"`
	ServiceAccounts []string                 `json:"serviceAccounts,omitempty"`
	Users           []string                 `json:"users,omitempty"`
	CreatedDt       time.Time             `json:"createdDt,omitempty"`
}

type RoleRefVo struct {
	Kind string `json:"kind,omitempty"`
	Name string `json:"name,omitempty"`
}

type RoleResponse struct {
	Annotations   map[string]string            `json:"annotations,omitempty"`
	ClusterId     string                       `json:"clusterId,omitempty"`
	ClusterName   string                       `json:"clusterName,omitempty"`
	Labels        map[string]string            `json:"labels,omitempty"`
	NamespaceName string                       `json:"namespaceName,omitempty"`
	PolicyRules   []PolicyRuleVo `json:"policyRules,omitempty"`
	Region        string                       `json:"region,omitempty"`
	RoleName      string                       `json:"roleName,omitempty"`
	RunningAge    string                       `json:"runningAge,omitempty"`
	CreatedDt     time.Time                 `json:"createdDt,omitempty"`
}

type RoleServiceApiListRoleEventsV1Opts struct {
	SystemObject optional.Bool
	Page         optional.Int32
	Size         optional.Int32
}

type RoleServiceApiListRolesV1Opts struct {
	RoleName     optional.String
	SystemObject optional.Bool
	Page         optional.Int32
	Size         optional.Int32
	Sort         optional.String
}

type RolesResponse struct {
	ProjectId     string       `json:"projectId,omitempty"`
	NamespaceName string       `json:"namespaceName,omitempty"`
	RoleName      string       `json:"roleName,omitempty"`
	RunningAge    string       `json:"runningAge,omitempty"`
	CreatedDt     time.Time `json:"createdDt,omitempty"`
}

type SecretResponse struct {
	ProjectId     string                 `json:"projectId,omitempty"`
	Annotations   map[string]interface{} `json:"annotations,omitempty"`
	ClusterId     string                 `json:"clusterId,omitempty"`
	ClusterName   string                 `json:"clusterName,omitempty"`
	Labels        map[string]interface{} `json:"labels,omitempty"`
	NamespaceName string                 `json:"namespaceName,omitempty"`
	Region        string                 `json:"region,omitempty"`
	RunningAge    string                 `json:"runningAge,omitempty"`
	SecretData    map[string]string      `json:"secretData,omitempty"`
	SecretName    string                 `json:"secretName,omitempty"`
	SecretType    string                 `json:"secretType,omitempty"`
	CreatedDt     time.Time           `json:"createdDt,omitempty"`
}

type SecretServiceApiListSecretEventsV2Opts struct {
	SystemObject optional.Bool
	Page         optional.Int32
	Size         optional.Int32
}

type SecretServiceApiListSecretsV2Opts struct {
	SecretName   optional.String
	SystemObject optional.Bool
	Page         optional.Int32
	Size         optional.Int32
	Sort         optional.String
}

type SecretsResponse struct {
	ProjectId  string            `json:"projectId,omitempty"`
	RunningAge string            `json:"runningAge,omitempty"`
	SecretData map[string]string `json:"secretData,omitempty"`
	SecretName string            `json:"secretName,omitempty"`
	SecretType string            `json:"secretType,omitempty"`
}

type StatefulSetResponse struct {
	ProjectId       string            `json:"projectId,omitempty"`
	Annotations     map[string]string `json:"annotations,omitempty"`
	ClusterId       string            `json:"clusterId,omitempty"`
	ClusterName     string            `json:"clusterName,omitempty"`
	Labels          map[string]string `json:"labels,omitempty"`
	NamespaceName   string            `json:"namespaceName,omitempty"`
	ReadyReplicas   int32             `json:"readyReplicas,omitempty"`
	Region          string            `json:"region,omitempty"`
	RunningAge      string            `json:"runningAge,omitempty"`
	SelectorDetails []string          `json:"selectorDetails,omitempty"`
	Selectors       map[string]string `json:"selectors,omitempty"`
	SpecReplicas    int32             `json:"specReplicas,omitempty"`
	StatefulSetName string            `json:"statefulSetName,omitempty"`
	StatusReplicas  int32             `json:"statusReplicas,omitempty"`
	UpdateStrategy  string            `json:"updateStrategy,omitempty"`
	CreatedDt       time.Time      `json:"createdDt,omitempty"`
}

type StatefulSetServiceApiListStatefulSetEventsV2Opts struct {
	SystemObject optional.Bool
	Page         optional.Int32
	Size         optional.Int32
}

type StatefulSetServiceApiListStatefulSetPodsV2Opts struct {
	SystemObject optional.Bool
	Page         optional.Int32
	Size         optional.Int32
}

type StatefulSetServiceApiListStatefulSetsV2Opts struct {
	StatefulSetName optional.String
	SystemObject    optional.Bool
	Page            optional.Int32
	Size            optional.Int32
	Sort            optional.String
}

type StatefulSetsResponse struct {
	ProjectId       string       `json:"projectId,omitempty"`
	NamespaceName   string       `json:"namespaceName,omitempty"`
	ReadyReplicas   int32        `json:"readyReplicas,omitempty"`
	RunningAge      string       `json:"runningAge,omitempty"`
	SpecReplicas    int32        `json:"specReplicas,omitempty"`
	StatefulSetName string       `json:"statefulSetName,omitempty"`
	StatusReplicas  int32        `json:"statusReplicas,omitempty"`
	CreatedDt       time.Time `json:"createdDt,omitempty"`
}

type StorageClassResponse struct {
	ProjectId              string            `json:"projectId,omitempty"`
	AllowVolumeExpansion   string            `json:"allowVolumeExpansion,omitempty"`
	Annotations            map[string]string `json:"annotations,omitempty"`
	ClusterId              string            `json:"clusterId,omitempty"`
	ClusterName            string            `json:"clusterName,omitempty"`
	DefaultClass           string            `json:"defaultClass,omitempty"`
	Labels                 map[string]string `json:"labels,omitempty"`
	MountOptions           []string          `json:"mountOptions,omitempty"`
	Provisioner            string            `json:"provisioner,omitempty"`
	PvReclaimPolicy        string            `json:"pvReclaimPolicy,omitempty"`
	Region                 string            `json:"region,omitempty"`
	StorageClassName       string            `json:"storageClassName,omitempty"`
	StorageClassParameters map[string]string `json:"storageClassParameters,omitempty"`
	VolumeBindingMode      string            `json:"volumeBindingMode,omitempty"`
	CreatedDt              time.Time      `json:"createdDt,omitempty"`
}

type StorageClassServiceApiListStorageClassEventsV2Opts struct {
	SystemObject optional.Bool
	Page         optional.Int32
	Size         optional.Int32
}

type StorageClassServiceApiListStorageClassesV2Opts struct {
	StorageClassName optional.String
	SystemObject     optional.Bool
	Page             optional.Int32
	Size             optional.Int32
	Sort             optional.String
}

type StorageClassesResponse struct {
	ProjectId            string `json:"projectId,omitempty"`
	AllowVolumeExpansion string `json:"allowVolumeExpansion,omitempty"`
	DefaultClass         string `json:"defaultClass,omitempty"`
	Provisioner          string `json:"provisioner,omitempty"`
	PvReclaimPolicy      string `json:"pvReclaimPolicy,omitempty"`
	RunningAge           string `json:"runningAge,omitempty"`
	StorageClassName     string `json:"storageClassName,omitempty"`
	VolumeBindingMode    string `json:"volumeBindingMode,omitempty"`
}

type SubjectVo struct {
	Kind      string `json:"kind,omitempty"`
	Name      string `json:"name,omitempty"`
	Namespace string `json:"namespace,omitempty"`
}

type SubsetResponse struct {
	Addresses         []string                       `json:"addresses,omitempty"`
	NotReadyAddresses []string                       `json:"notReadyAddresses,omitempty"`
	Ports             []EndpointPortVo `json:"ports,omitempty"`
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
