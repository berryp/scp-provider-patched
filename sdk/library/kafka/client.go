package kafka

import (
	scpsdk "github.com/SamsungSDSCloud/terraform-sdk-samsungcloudplatform/v3/client"
)

type APIClient struct {
	cfg    *scpsdk.Configuration
	common service
	ConfigurationApi *ConfigurationApiService
	DatabaseBlockStorageApi *DatabaseBlockStorageApiService
	DatabaseContractApi *DatabaseContractApiService
	DatabaseEventControllerApi *DatabaseEventControllerApiService
	DatabaseManagementApi *DatabaseManagementApiService
	DatabaseOperateApi *DatabaseOperateApiService
	DatabaseParameterApi *DatabaseParameterApiService
	KafkaConfigurationControllerApi *KafkaConfigurationControllerApiService
	KafkaInfraResourceApi *KafkaInfraResourceApiService
	KafkaNetworkApi *KafkaNetworkApiService
	KafkaOperationManagementApi *KafkaOperationManagementApiService
	KafkaPricingApi *KafkaPricingApiService
	KafkaProvisioningApi *KafkaProvisioningApiService
	KafkaSearchApi *KafkaSearchApiService
	Client *Client
}

type service struct {
	client *APIClient
}

type ConfigurationApiService service
type DatabaseBlockStorageApiService service
type DatabaseContractApiService service
type DatabaseEventControllerApiService service
type DatabaseManagementApiService service
type DatabaseOperateApiService service
type DatabaseParameterApiService service
type KafkaConfigurationControllerApiService service
type KafkaInfraResourceApiService service
type KafkaNetworkApiService service
type KafkaOperationManagementApiService service
type KafkaPricingApiService service
type KafkaProvisioningApiService service
type KafkaSearchApiService service
type Client service

func NewAPIClient(cfg *scpsdk.Configuration) *APIClient {
	c := &APIClient{cfg: cfg}
	c.common.client = c
	c.ConfigurationApi = (*ConfigurationApiService)(&c.common)
	c.DatabaseBlockStorageApi = (*DatabaseBlockStorageApiService)(&c.common)
	c.DatabaseContractApi = (*DatabaseContractApiService)(&c.common)
	c.DatabaseEventControllerApi = (*DatabaseEventControllerApiService)(&c.common)
	c.DatabaseManagementApi = (*DatabaseManagementApiService)(&c.common)
	c.DatabaseOperateApi = (*DatabaseOperateApiService)(&c.common)
	c.DatabaseParameterApi = (*DatabaseParameterApiService)(&c.common)
	c.KafkaConfigurationControllerApi = (*KafkaConfigurationControllerApiService)(&c.common)
	c.KafkaInfraResourceApi = (*KafkaInfraResourceApiService)(&c.common)
	c.KafkaNetworkApi = (*KafkaNetworkApiService)(&c.common)
	c.KafkaOperationManagementApi = (*KafkaOperationManagementApiService)(&c.common)
	c.KafkaPricingApi = (*KafkaPricingApiService)(&c.common)
	c.KafkaProvisioningApi = (*KafkaProvisioningApiService)(&c.common)
	c.KafkaSearchApi = (*KafkaSearchApiService)(&c.common)
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
