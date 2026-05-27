package mariadb

import (
	scpsdk "github.com/SamsungSDSCloud/terraform-sdk-samsungcloudplatform/v3/client"
)

type APIClient struct {
	cfg    *scpsdk.Configuration
	common service
	ConfigurationApi *ConfigurationApiService
	DatabaseArchiveApi *DatabaseArchiveApiService
	DatabaseBackupApi *DatabaseBackupApiService
	DatabaseBlockStorageApi *DatabaseBlockStorageApiService
	DatabaseContractApi *DatabaseContractApiService
	DatabaseDiskUsageApi *DatabaseDiskUsageApiService
	DatabaseEventControllerApi *DatabaseEventControllerApiService
	DatabaseManagementApi *DatabaseManagementApiService
	DatabaseOperateApi *DatabaseOperateApiService
	DatabaseParameterApi *DatabaseParameterApiService
	DatabaseReplicaApi *DatabaseReplicaApiService
	DatabaseUserApi *DatabaseUserApiService
	DbLogExportConfigApi *DbLogExportConfigApiService
	DbLogExportConfigUIApi *DbLogExportConfigUIApiService
	MariaDbConfigurationControllerApi *MariaDbConfigurationControllerApiService
	MariadbBackupApi *MariadbBackupApiService
	MariadbInfraResourceApi *MariadbInfraResourceApiService
	MariadbNetworkApi *MariadbNetworkApiService
	MariadbOperationManagementApi *MariadbOperationManagementApiService
	MariadbPricingApi *MariadbPricingApiService
	MariadbProvisioningApi *MariadbProvisioningApiService
	MariadbSearchApi *MariadbSearchApiService
	Client *Client
}

type service struct {
	client *APIClient
}

type ConfigurationApiService service
type DatabaseArchiveApiService service
type DatabaseBackupApiService service
type DatabaseBlockStorageApiService service
type DatabaseContractApiService service
type DatabaseDiskUsageApiService service
type DatabaseEventControllerApiService service
type DatabaseManagementApiService service
type DatabaseOperateApiService service
type DatabaseParameterApiService service
type DatabaseReplicaApiService service
type DatabaseUserApiService service
type DbLogExportConfigApiService service
type DbLogExportConfigUIApiService service
type MariaDbConfigurationControllerApiService service
type MariadbBackupApiService service
type MariadbInfraResourceApiService service
type MariadbNetworkApiService service
type MariadbOperationManagementApiService service
type MariadbPricingApiService service
type MariadbProvisioningApiService service
type MariadbSearchApiService service
type Client service

func NewAPIClient(cfg *scpsdk.Configuration) *APIClient {
	c := &APIClient{cfg: cfg}
	c.common.client = c
	c.ConfigurationApi = (*ConfigurationApiService)(&c.common)
	c.DatabaseArchiveApi = (*DatabaseArchiveApiService)(&c.common)
	c.DatabaseBackupApi = (*DatabaseBackupApiService)(&c.common)
	c.DatabaseBlockStorageApi = (*DatabaseBlockStorageApiService)(&c.common)
	c.DatabaseContractApi = (*DatabaseContractApiService)(&c.common)
	c.DatabaseDiskUsageApi = (*DatabaseDiskUsageApiService)(&c.common)
	c.DatabaseEventControllerApi = (*DatabaseEventControllerApiService)(&c.common)
	c.DatabaseManagementApi = (*DatabaseManagementApiService)(&c.common)
	c.DatabaseOperateApi = (*DatabaseOperateApiService)(&c.common)
	c.DatabaseParameterApi = (*DatabaseParameterApiService)(&c.common)
	c.DatabaseReplicaApi = (*DatabaseReplicaApiService)(&c.common)
	c.DatabaseUserApi = (*DatabaseUserApiService)(&c.common)
	c.DbLogExportConfigApi = (*DbLogExportConfigApiService)(&c.common)
	c.DbLogExportConfigUIApi = (*DbLogExportConfigUIApiService)(&c.common)
	c.MariaDbConfigurationControllerApi = (*MariaDbConfigurationControllerApiService)(&c.common)
	c.MariadbBackupApi = (*MariadbBackupApiService)(&c.common)
	c.MariadbInfraResourceApi = (*MariadbInfraResourceApiService)(&c.common)
	c.MariadbNetworkApi = (*MariadbNetworkApiService)(&c.common)
	c.MariadbOperationManagementApi = (*MariadbOperationManagementApiService)(&c.common)
	c.MariadbPricingApi = (*MariadbPricingApiService)(&c.common)
	c.MariadbProvisioningApi = (*MariadbProvisioningApiService)(&c.common)
	c.MariadbSearchApi = (*MariadbSearchApiService)(&c.common)
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
