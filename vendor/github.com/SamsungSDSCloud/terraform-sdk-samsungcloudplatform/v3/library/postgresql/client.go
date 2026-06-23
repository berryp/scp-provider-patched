package postgresql

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
	PostgreSqlConfigurationControllerApi *PostgreSqlConfigurationControllerApiService
	PostgresqlBackupApi *PostgresqlBackupApiService
	PostgresqlInfraResourceApi *PostgresqlInfraResourceApiService
	PostgresqlNetworkApi *PostgresqlNetworkApiService
	PostgresqlOperationManagementApi *PostgresqlOperationManagementApiService
	PostgresqlPricingApi *PostgresqlPricingApiService
	PostgresqlProvisioningApi *PostgresqlProvisioningApiService
	PostgresqlSearchApi *PostgresqlSearchApiService
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
type PostgreSqlConfigurationControllerApiService service
type PostgresqlBackupApiService service
type PostgresqlInfraResourceApiService service
type PostgresqlNetworkApiService service
type PostgresqlOperationManagementApiService service
type PostgresqlPricingApiService service
type PostgresqlProvisioningApiService service
type PostgresqlSearchApiService service
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
	c.PostgreSqlConfigurationControllerApi = (*PostgreSqlConfigurationControllerApiService)(&c.common)
	c.PostgresqlBackupApi = (*PostgresqlBackupApiService)(&c.common)
	c.PostgresqlInfraResourceApi = (*PostgresqlInfraResourceApiService)(&c.common)
	c.PostgresqlNetworkApi = (*PostgresqlNetworkApiService)(&c.common)
	c.PostgresqlOperationManagementApi = (*PostgresqlOperationManagementApiService)(&c.common)
	c.PostgresqlPricingApi = (*PostgresqlPricingApiService)(&c.common)
	c.PostgresqlProvisioningApi = (*PostgresqlProvisioningApiService)(&c.common)
	c.PostgresqlSearchApi = (*PostgresqlSearchApiService)(&c.common)
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
