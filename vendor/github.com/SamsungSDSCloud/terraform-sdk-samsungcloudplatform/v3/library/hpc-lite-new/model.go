package hpclitenew

import (
	"time"
)

type APIKey struct {
	Key    string
	Prefix string
}

type AsyncListResponse struct {
	ProjectId      string   `json:"projectId,omitempty"`
	RequestId      string   `json:"requestId,omitempty"`
	ResourceIdList []string `json:"resourceIdList,omitempty"`
}

type BasicAuth struct {
	UserName string `json:"userName,omitempty"`
	Password string `json:"password,omitempty"`
}

type HpcLiteNewCreateRequest struct {
	CoServiceZoneId       string
	Contract              string
	HyperThreadingEnabled string
	ImageId               string
	InitScript            string
	OsUserId              string
	OsUserPassword        string
	ProductGroupId        string
	ResourcePoolId        string
	ServerDetails         []ServerDetailRequest
	ServerType            string
	ServiceZoneId         string
	Tags                  map[string]interface{}
	VlanPoolCidr          string
}

type HpcLiteNewDeleteRequest struct {
	ServerIds     []string
	ServiceZoneId string
}

type HpcLitePlusOpenApiCreateRequestVo struct {
	CoServiceZoneId       string                                `json:"coServiceZoneId,omitempty"`
	Contract              string                                `json:"contract,omitempty"`
	HyperThreadingEnabled string                                `json:"hyperThreadingEnabled,omitempty"`
	ImageId               string                                `json:"imageId,omitempty"`
	InitScript            string                                `json:"initScript,omitempty"`
	OsUserId              string                                `json:"osUserId,omitempty"`
	OsUserPassword        string                                `json:"osUserPassword,omitempty"`
	ProductGroupId        string                                `json:"productGroupId,omitempty"`
	ResourcePoolId        string                                `json:"resourcePoolId,omitempty"`
	ServerDetails         []ServerDetailRequestVo `json:"serverDetails,omitempty"`
	ServerType            string                                `json:"serverType,omitempty"`
	ServiceZoneId         string                                `json:"serviceZoneId,omitempty"`
	Tags                  []TagRequest            `json:"tags,omitempty"`
	VlanPoolCidr          string                                `json:"vlanPoolCidr,omitempty"`
}

type HpcLitePlusOpenApiDeleteRequestVo struct {
	ServerIds     []string `json:"serverIds,omitempty"`
	ServiceZoneId string   `json:"serviceZoneId,omitempty"`
}

type HpcLitePlusOpenApiDetailResponseDto struct {
	ProjectId      string       `json:"projectId,omitempty"`
	BlockId        string       `json:"blockId,omitempty"`
	CoServiceZone  string       `json:"coServiceZone,omitempty"`
	Contract       string       `json:"contract,omitempty"`
	ContractId     string       `json:"contractId,omitempty"`
	HardwareName   string       `json:"hardwareName,omitempty"`
	HyperThreading string       `json:"hyperThreading,omitempty"`
	ImageId        string       `json:"imageId,omitempty"`
	ImageName      string       `json:"imageName,omitempty"`
	InitScript     string       `json:"initScript,omitempty"`
	IpAddress      string       `json:"ipAddress,omitempty"`
	RegionName     string       `json:"regionName,omitempty"`
	ServerId       string       `json:"serverId,omitempty"`
	ServerName     string       `json:"serverName,omitempty"`
	ServerState    string       `json:"serverState,omitempty"`
	ServerType     string       `json:"serverType,omitempty"`
	ServerTypeId   string       `json:"serverTypeId,omitempty"`
	Timezone       string       `json:"timezone,omitempty"`
	VlanId         string       `json:"vlanId,omitempty"`
	ZoneId         string       `json:"zoneId,omitempty"`
	CreatedBy      string       `json:"createdBy,omitempty"`
	CreatedDt      time.Time `json:"createdDt,omitempty"`
	ModifiedBy     string       `json:"modifiedBy,omitempty"`
	ModifiedDt     time.Time `json:"modifiedDt,omitempty"`
}

type ServerDetailRequest struct {
	IpAddress  string
	ServerName string
}

type ServerDetailRequestVo struct {
	IpAddress  string `json:"ipAddress,omitempty"`
	ServerName string `json:"serverName,omitempty"`
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
