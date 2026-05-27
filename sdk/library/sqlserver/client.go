package sqlserver

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
	DatabaseReplicaApi *DatabaseReplicaApiService
	DatabaseUserApi *DatabaseUserApiService
	DbLogExportConfigApi *DbLogExportConfigApiService
	DbLogExportConfigUIApi *DbLogExportConfigUIApiService
	MsSqlConfigurationControllerApi *MsSqlConfigurationControllerApiService
	SqlserverBackupApi *SqlserverBackupApiService
	SqlserverChangeBackupMethodControllerApi *SqlserverChangeBackupMethodControllerApiService
	SqlserverInfraResourceApi *SqlserverInfraResourceApiService
	SqlserverNetworkApi *SqlserverNetworkApiService
	SqlserverOperationManagementApi *SqlserverOperationManagementApiService
	SqlserverPricingApi *SqlserverPricingApiService
	SqlserverProvisioningApi *SqlserverProvisioningApiService
	SqlserverSearchApi *SqlserverSearchApiService
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
type DatabaseReplicaApiService service
type DatabaseUserApiService service
type DbLogExportConfigApiService service
type DbLogExportConfigUIApiService service
type MsSqlConfigurationControllerApiService service
type SqlserverBackupApiService service
type SqlserverChangeBackupMethodControllerApiService service
type SqlserverInfraResourceApiService service
type SqlserverNetworkApiService service
type SqlserverOperationManagementApiService service
type SqlserverPricingApiService service
type SqlserverProvisioningApiService service
type SqlserverSearchApiService service
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
	c.DatabaseReplicaApi = (*DatabaseReplicaApiService)(&c.common)
	c.DatabaseUserApi = (*DatabaseUserApiService)(&c.common)
	c.DbLogExportConfigApi = (*DbLogExportConfigApiService)(&c.common)
	c.DbLogExportConfigUIApi = (*DbLogExportConfigUIApiService)(&c.common)
	c.MsSqlConfigurationControllerApi = (*MsSqlConfigurationControllerApiService)(&c.common)
	c.SqlserverBackupApi = (*SqlserverBackupApiService)(&c.common)
	c.SqlserverChangeBackupMethodControllerApi = (*SqlserverChangeBackupMethodControllerApiService)(&c.common)
	c.SqlserverInfraResourceApi = (*SqlserverInfraResourceApiService)(&c.common)
	c.SqlserverNetworkApi = (*SqlserverNetworkApiService)(&c.common)
	c.SqlserverOperationManagementApi = (*SqlserverOperationManagementApiService)(&c.common)
	c.SqlserverPricingApi = (*SqlserverPricingApiService)(&c.common)
	c.SqlserverProvisioningApi = (*SqlserverProvisioningApiService)(&c.common)
	c.SqlserverSearchApi = (*SqlserverSearchApiService)(&c.common)
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
