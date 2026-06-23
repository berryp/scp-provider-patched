package redis

import (
	scpsdk "github.com/SamsungSDSCloud/terraform-sdk-samsungcloudplatform/v3/client"
)

type APIClient struct {
	cfg    *scpsdk.Configuration
	common service
	ConfigurationApi *ConfigurationApiService
	DatabaseBackupApi *DatabaseBackupApiService
	DatabaseBlockStorageApi *DatabaseBlockStorageApiService
	DatabaseContractApi *DatabaseContractApiService
	DatabaseEventControllerApi *DatabaseEventControllerApiService
	DatabaseManagementApi *DatabaseManagementApiService
	DatabaseOperateApi *DatabaseOperateApiService
	DatabaseParameterApi *DatabaseParameterApiService
	DatabaseUserApi *DatabaseUserApiService
	RedisBackupApi *RedisBackupApiService
	RedisClusterBackupApi *RedisClusterBackupApiService
	RedisClusterInfraResourceApi *RedisClusterInfraResourceApiService
	RedisClusterNetworkApi *RedisClusterNetworkApiService
	RedisClusterOperationManagementApi *RedisClusterOperationManagementApiService
	RedisClusterPricingApi *RedisClusterPricingApiService
	RedisClusterProvisioningApi *RedisClusterProvisioningApiService
	RedisClusterSearchApi *RedisClusterSearchApiService
	RedisConfigurationControllerApi *RedisConfigurationControllerApiService
	RedisInfraResourceApi *RedisInfraResourceApiService
	RedisNetworkApi *RedisNetworkApiService
	RedisOperationManagementApi *RedisOperationManagementApiService
	RedisPricingApi *RedisPricingApiService
	RedisProvisioningApi *RedisProvisioningApiService
	RedisRenameCommandControllerApi *RedisRenameCommandControllerApiService
	RedisSearchApi *RedisSearchApiService
	Client *Client
}

type service struct {
	client *APIClient
}

type ConfigurationApiService service
type DatabaseBackupApiService service
type DatabaseBlockStorageApiService service
type DatabaseContractApiService service
type DatabaseEventControllerApiService service
type DatabaseManagementApiService service
type DatabaseOperateApiService service
type DatabaseParameterApiService service
type DatabaseUserApiService service
type RedisBackupApiService service
type RedisClusterBackupApiService service
type RedisClusterInfraResourceApiService service
type RedisClusterNetworkApiService service
type RedisClusterOperationManagementApiService service
type RedisClusterPricingApiService service
type RedisClusterProvisioningApiService service
type RedisClusterSearchApiService service
type RedisConfigurationControllerApiService service
type RedisInfraResourceApiService service
type RedisNetworkApiService service
type RedisOperationManagementApiService service
type RedisPricingApiService service
type RedisProvisioningApiService service
type RedisRenameCommandControllerApiService service
type RedisSearchApiService service
type Client service

func NewAPIClient(cfg *scpsdk.Configuration) *APIClient {
	c := &APIClient{cfg: cfg}
	c.common.client = c
	c.ConfigurationApi = (*ConfigurationApiService)(&c.common)
	c.DatabaseBackupApi = (*DatabaseBackupApiService)(&c.common)
	c.DatabaseBlockStorageApi = (*DatabaseBlockStorageApiService)(&c.common)
	c.DatabaseContractApi = (*DatabaseContractApiService)(&c.common)
	c.DatabaseEventControllerApi = (*DatabaseEventControllerApiService)(&c.common)
	c.DatabaseManagementApi = (*DatabaseManagementApiService)(&c.common)
	c.DatabaseOperateApi = (*DatabaseOperateApiService)(&c.common)
	c.DatabaseParameterApi = (*DatabaseParameterApiService)(&c.common)
	c.DatabaseUserApi = (*DatabaseUserApiService)(&c.common)
	c.RedisBackupApi = (*RedisBackupApiService)(&c.common)
	c.RedisClusterBackupApi = (*RedisClusterBackupApiService)(&c.common)
	c.RedisClusterInfraResourceApi = (*RedisClusterInfraResourceApiService)(&c.common)
	c.RedisClusterNetworkApi = (*RedisClusterNetworkApiService)(&c.common)
	c.RedisClusterOperationManagementApi = (*RedisClusterOperationManagementApiService)(&c.common)
	c.RedisClusterPricingApi = (*RedisClusterPricingApiService)(&c.common)
	c.RedisClusterProvisioningApi = (*RedisClusterProvisioningApiService)(&c.common)
	c.RedisClusterSearchApi = (*RedisClusterSearchApiService)(&c.common)
	c.RedisConfigurationControllerApi = (*RedisConfigurationControllerApiService)(&c.common)
	c.RedisInfraResourceApi = (*RedisInfraResourceApiService)(&c.common)
	c.RedisNetworkApi = (*RedisNetworkApiService)(&c.common)
	c.RedisOperationManagementApi = (*RedisOperationManagementApiService)(&c.common)
	c.RedisPricingApi = (*RedisPricingApiService)(&c.common)
	c.RedisProvisioningApi = (*RedisProvisioningApiService)(&c.common)
	c.RedisRenameCommandControllerApi = (*RedisRenameCommandControllerApiService)(&c.common)
	c.RedisSearchApi = (*RedisSearchApiService)(&c.common)
	c.Client = (*Client)(&c.common)
	return c
}

func (c *APIClient) ChangeBasePath(p string) { c.cfg.BasePath = p }

func (c *APIClient) ToTagRequestList(tags map[string]interface{}) []TagRequest {
	if len(tags) == 0 {
		return nil
	}
	out := make([]TagRequest, 0, len(tags))
	for k, v := range tags {
		t := TagRequest{TagKey: k}
		if s, ok := v.(string); ok {
			t.TagValue = s
		}
		out = append(out, t)
	}
	return out
}
