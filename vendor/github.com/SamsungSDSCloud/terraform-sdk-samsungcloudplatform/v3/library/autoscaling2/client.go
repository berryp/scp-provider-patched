package autoscaling2

import (
	scpsdk "github.com/SamsungSDSCloud/terraform-sdk-samsungcloudplatform/v3/client"
)

type APIClient struct {
	cfg    *scpsdk.Configuration
	common service
	AsgEventV2Api *AsgEventV2ApiService
	AsgLaunchConfigurationV2Api *AsgLaunchConfigurationV2ApiService
	AsgLaunchConfigurationV4Api *AsgLaunchConfigurationV4ApiService
	AsgLaunchConfigurationV5Api *AsgLaunchConfigurationV5ApiService
	AsgLaunchConfigurationV6Api *AsgLaunchConfigurationV6ApiService
	AsgLoadBalancerV2Api *AsgLoadBalancerV2ApiService
	AsgNotificationV2Api *AsgNotificationV2ApiService
	AsgPolicyV2Api *AsgPolicyV2ApiService
	AsgScheduleV2Api *AsgScheduleV2ApiService
	AsgVirtualServerV2Api *AsgVirtualServerV2ApiService
	AutoScalingGroupV2Api *AutoScalingGroupV2ApiService
	AutoScalingGroupV3Api *AutoScalingGroupV3ApiService
	AutoScalingGroupV4Api *AutoScalingGroupV4ApiService
}

type service struct {
	client *APIClient
}

type AsgEventV2ApiService service
type AsgLaunchConfigurationV2ApiService service
type AsgLaunchConfigurationV4ApiService service
type AsgLaunchConfigurationV5ApiService service
type AsgLaunchConfigurationV6ApiService service
type AsgLoadBalancerV2ApiService service
type AsgNotificationV2ApiService service
type AsgPolicyV2ApiService service
type AsgScheduleV2ApiService service
type AsgVirtualServerV2ApiService service
type AutoScalingGroupV2ApiService service
type AutoScalingGroupV3ApiService service
type AutoScalingGroupV4ApiService service

func NewAPIClient(cfg *scpsdk.Configuration) *APIClient {
	c := &APIClient{cfg: cfg}
	c.common.client = c
	c.AsgEventV2Api = (*AsgEventV2ApiService)(&c.common)
	c.AsgLaunchConfigurationV2Api = (*AsgLaunchConfigurationV2ApiService)(&c.common)
	c.AsgLaunchConfigurationV4Api = (*AsgLaunchConfigurationV4ApiService)(&c.common)
	c.AsgLaunchConfigurationV5Api = (*AsgLaunchConfigurationV5ApiService)(&c.common)
	c.AsgLaunchConfigurationV6Api = (*AsgLaunchConfigurationV6ApiService)(&c.common)
	c.AsgLoadBalancerV2Api = (*AsgLoadBalancerV2ApiService)(&c.common)
	c.AsgNotificationV2Api = (*AsgNotificationV2ApiService)(&c.common)
	c.AsgPolicyV2Api = (*AsgPolicyV2ApiService)(&c.common)
	c.AsgScheduleV2Api = (*AsgScheduleV2ApiService)(&c.common)
	c.AsgVirtualServerV2Api = (*AsgVirtualServerV2ApiService)(&c.common)
	c.AutoScalingGroupV2Api = (*AutoScalingGroupV2ApiService)(&c.common)
	c.AutoScalingGroupV3Api = (*AutoScalingGroupV3ApiService)(&c.common)
	c.AutoScalingGroupV4Api = (*AutoScalingGroupV4ApiService)(&c.common)
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
