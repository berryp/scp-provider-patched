package loggingaudit

import (
	scpsdk "github.com/SamsungSDSCloud/terraform-sdk-samsungcloudplatform/v3/client"
)

type APIClient struct {
	cfg    *scpsdk.Configuration
	common service
	CommonCodeControllerApi *CommonCodeControllerApiService
	LoggingControllerApi *LoggingControllerApiService
	PlatformControllerApi *PlatformControllerApiService
	TrailControllerApi *TrailControllerApiService
	Client *Client
}

type service struct {
	client *APIClient
}

type CommonCodeControllerApiService service
type LoggingControllerApiService service
type PlatformControllerApiService service
type TrailControllerApiService service
type Client service

func NewAPIClient(cfg *scpsdk.Configuration) *APIClient {
	c := &APIClient{cfg: cfg}
	c.common.client = c
	c.CommonCodeControllerApi = (*CommonCodeControllerApiService)(&c.common)
	c.LoggingControllerApi = (*LoggingControllerApiService)(&c.common)
	c.PlatformControllerApi = (*PlatformControllerApiService)(&c.common)
	c.TrailControllerApi = (*TrailControllerApiService)(&c.common)
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
