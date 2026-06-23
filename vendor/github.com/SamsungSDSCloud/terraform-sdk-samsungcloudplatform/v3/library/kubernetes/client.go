package kubernetes

import (
	scpsdk "github.com/SamsungSDSCloud/terraform-sdk-samsungcloudplatform/v3/client"
)

type APIClient struct {
	cfg    *scpsdk.Configuration
	common service
	ClusterRoleBindingServiceApi *ClusterRoleBindingServiceApiService
	ClusterRoleServiceApi *ClusterRoleServiceApiService
	CommonCodeControllerApi *CommonCodeControllerApiService
	ConfigMapServiceApi *ConfigMapServiceApiService
	CronJobServiceApi *CronJobServiceApiService
	DaemonSetServiceApi *DaemonSetServiceApiService
	DeploymentServiceApi *DeploymentServiceApiService
	EndpointServiceApi *EndpointServiceApiService
	IngressClassServiceApi *IngressClassServiceApiService
	IngressServiceApi *IngressServiceApiService
	JobServiceApi *JobServiceApiService
	K8sObjectYamlServiceApi *K8sObjectYamlServiceApiService
	K8sServiceApi *K8sServiceApiService
	NamespaceServiceApi *NamespaceServiceApiService
	NodeServiceApi *NodeServiceApiService
	PersistentVolumeClaimServiceApi *PersistentVolumeClaimServiceApiService
	PersistentVolumeServiceApi *PersistentVolumeServiceApiService
	PlatformControllerApi *PlatformControllerApiService
	PodServiceApi *PodServiceApiService
	RoleBindingServiceApi *RoleBindingServiceApiService
	RoleServiceApi *RoleServiceApiService
	SecretServiceApi *SecretServiceApiService
	StatefulSetServiceApi *StatefulSetServiceApiService
	StorageClassServiceApi *StorageClassServiceApiService
	Client *Client
}

type service struct {
	client *APIClient
}

type ClusterRoleBindingServiceApiService service
type ClusterRoleServiceApiService service
type CommonCodeControllerApiService service
type ConfigMapServiceApiService service
type CronJobServiceApiService service
type DaemonSetServiceApiService service
type DeploymentServiceApiService service
type EndpointServiceApiService service
type IngressClassServiceApiService service
type IngressServiceApiService service
type JobServiceApiService service
type K8sObjectYamlServiceApiService service
type K8sServiceApiService service
type NamespaceServiceApiService service
type NodeServiceApiService service
type PersistentVolumeClaimServiceApiService service
type PersistentVolumeServiceApiService service
type PlatformControllerApiService service
type PodServiceApiService service
type RoleBindingServiceApiService service
type RoleServiceApiService service
type SecretServiceApiService service
type StatefulSetServiceApiService service
type StorageClassServiceApiService service
type Client service

func NewAPIClient(cfg *scpsdk.Configuration) *APIClient {
	c := &APIClient{cfg: cfg}
	c.common.client = c
	c.ClusterRoleBindingServiceApi = (*ClusterRoleBindingServiceApiService)(&c.common)
	c.ClusterRoleServiceApi = (*ClusterRoleServiceApiService)(&c.common)
	c.CommonCodeControllerApi = (*CommonCodeControllerApiService)(&c.common)
	c.ConfigMapServiceApi = (*ConfigMapServiceApiService)(&c.common)
	c.CronJobServiceApi = (*CronJobServiceApiService)(&c.common)
	c.DaemonSetServiceApi = (*DaemonSetServiceApiService)(&c.common)
	c.DeploymentServiceApi = (*DeploymentServiceApiService)(&c.common)
	c.EndpointServiceApi = (*EndpointServiceApiService)(&c.common)
	c.IngressClassServiceApi = (*IngressClassServiceApiService)(&c.common)
	c.IngressServiceApi = (*IngressServiceApiService)(&c.common)
	c.JobServiceApi = (*JobServiceApiService)(&c.common)
	c.K8sObjectYamlServiceApi = (*K8sObjectYamlServiceApiService)(&c.common)
	c.K8sServiceApi = (*K8sServiceApiService)(&c.common)
	c.NamespaceServiceApi = (*NamespaceServiceApiService)(&c.common)
	c.NodeServiceApi = (*NodeServiceApiService)(&c.common)
	c.PersistentVolumeClaimServiceApi = (*PersistentVolumeClaimServiceApiService)(&c.common)
	c.PersistentVolumeServiceApi = (*PersistentVolumeServiceApiService)(&c.common)
	c.PlatformControllerApi = (*PlatformControllerApiService)(&c.common)
	c.PodServiceApi = (*PodServiceApiService)(&c.common)
	c.RoleBindingServiceApi = (*RoleBindingServiceApiService)(&c.common)
	c.RoleServiceApi = (*RoleServiceApiService)(&c.common)
	c.SecretServiceApi = (*SecretServiceApiService)(&c.common)
	c.StatefulSetServiceApi = (*StatefulSetServiceApiService)(&c.common)
	c.StorageClassServiceApi = (*StorageClassServiceApiService)(&c.common)
	c.Client = (*Client)(&c.common)
	return c
}

func (c *APIClient) ChangeBasePath(p string) { c.cfg.BasePath = p }
