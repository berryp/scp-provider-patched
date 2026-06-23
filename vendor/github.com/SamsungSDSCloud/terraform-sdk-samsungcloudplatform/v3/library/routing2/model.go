package routing2

import (
	"github.com/antihax/optional"
	"time"
)

type APIKey struct {
	Key    string
	Prefix string
}

type AsyncResponse struct {
	ProjectId  string `json:"projectId,omitempty"`
	RequestId  string `json:"requestId,omitempty"`
	ResourceId string `json:"resourceId,omitempty"`
}

type BasicAuth struct {
	UserName string `json:"userName,omitempty"`
	Password string `json:"password,omitempty"`
}

type CreateRoutingRulesRequest struct {
	RoutingRules []RoutingRule `json:"routingRules"`
}

type DcRoutingRuleListResponse struct {
	ProjectId                       string       `json:"projectId,omitempty"`
	BlockId                         string       `json:"blockId,omitempty"`
	DestinationNetworkCidr          string       `json:"destinationNetworkCidr,omitempty"`
	Editable                        *bool        `json:"editable,omitempty"`
	RoutingRuleId                   string       `json:"routingRuleId,omitempty"`
	RoutingRuleState                string       `json:"routingRuleState,omitempty"`
	ServiceZoneId                   string       `json:"serviceZoneId,omitempty"`
	SourceDirectConnectConnectionId string       `json:"sourceDirectConnectConnectionId,omitempty"`
	SourceServiceInterfaceId        string       `json:"sourceServiceInterfaceId,omitempty"`
	SourceServiceInterfaceName      string       `json:"sourceServiceInterfaceName,omitempty"`
	CreatedBy                       string       `json:"createdBy,omitempty"`
	CreatedDt                       time.Time `json:"createdDt,omitempty"`
	ModifiedBy                      string       `json:"modifiedBy,omitempty"`
	ModifiedDt                      time.Time `json:"modifiedDt,omitempty"`
}

type DcRoutingTableDetailResponse struct {
	ProjectId                 string       `json:"projectId,omitempty"`
	BlockId                   string       `json:"blockId,omitempty"`
	DirectConnectConnectionId string       `json:"directConnectConnectionId,omitempty"`
	RoutingRuleCount          int32        `json:"routingRuleCount,omitempty"`
	RoutingTableId            string       `json:"routingTableId,omitempty"`
	RoutingTableName          string       `json:"routingTableName,omitempty"`
	RoutingTableState         string       `json:"routingTableState,omitempty"`
	RoutingTableType          string       `json:"routingTableType,omitempty"`
	ServiceZoneId             string       `json:"serviceZoneId,omitempty"`
	T1RouterId                string       `json:"t1RouterId,omitempty"`
	CreatedBy                 string       `json:"createdBy,omitempty"`
	CreatedDt                 time.Time `json:"createdDt,omitempty"`
	ModifiedBy                string       `json:"modifiedBy,omitempty"`
	ModifiedDt                time.Time `json:"modifiedDt,omitempty"`
}

type DcRoutingTableListResponse struct {
	DirectConnectConnectionId string `json:"directConnectConnectionId,omitempty"`
	RoutingRuleCount          int64  `json:"routingRuleCount,omitempty"`
	RoutingTableId            string `json:"routingTableId,omitempty"`
	RoutingTableName          string `json:"routingTableName,omitempty"`
	RoutingTableType          string `json:"routingTableType,omitempty"`
	ServiceZoneId             string `json:"serviceZoneId,omitempty"`
	T1RouterId                string `json:"t1RouterId,omitempty"`
}

type DeleteRoutingRulesRequest struct {
	RoutingRuleIds []string `json:"routingRuleIds"`
}

type DirectConnectRoutingRuleOpenApiControllerApiListDcRoutingRulesOpts struct {
	DestinationNetworkCidr   optional.String
	Editable                 optional.Bool
	RoutingRuleId            optional.String
	SourceServiceInterfaceId optional.String
	Page                     optional.Int32
	Size                     optional.Int32
	Sort                     optional.Interface
}

type DirectConnectRoutingTableOpenApiControllerApiListDcRoutingTablesOpts struct {
	DirectConnectConnectionId optional.String
	RoutingTableId            optional.String
	RoutingTableName          optional.String
	CreatedBy                 optional.String
	Page                      optional.Int32
	Size                      optional.Int32
	Sort                      optional.Interface
}

type ListResponseDcRoutingRuleListResponse struct {
	Contents   []DcRoutingRuleListResponse `json:"contents,omitempty"`
	TotalCount int32                                   `json:"totalCount,omitempty"`
}

type ListResponseDcRoutingTableListResponse struct {
	Contents   []DcRoutingTableListResponse `json:"contents,omitempty"`
	TotalCount int32                                    `json:"totalCount,omitempty"`
}

type ListResponseRoutingRuleRouteListResponse struct {
	Contents   []RoutingRuleRouteListResponse `json:"contents,omitempty"`
	TotalCount int32                                      `json:"totalCount,omitempty"`
}

type ListResponseTgwRoutingRuleListResponse struct {
	Contents   []TgwRoutingRuleListResponse `json:"contents,omitempty"`
	TotalCount int32                                    `json:"totalCount,omitempty"`
}

type ListResponseTgwRoutingTableListResponse struct {
	Contents   []TgwRoutingTableListResponse `json:"contents,omitempty"`
	TotalCount int32                                     `json:"totalCount,omitempty"`
}

type ListResponseVpcRoutingRuleListResponse struct {
	Contents   []VpcRoutingRuleListResponse `json:"contents,omitempty"`
	TotalCount int32                                    `json:"totalCount,omitempty"`
}

type ListResponseVpcRoutingTableListResponse struct {
	Contents   []VpcRoutingTableListResponse `json:"contents,omitempty"`
	TotalCount int32                                     `json:"totalCount,omitempty"`
}

type RoutingRule struct {
	DestinationNetworkCidr     string `json:"destinationNetworkCidr"`
	SourceServiceInterfaceId   string `json:"sourceServiceInterfaceId"`
	SourceServiceInterfaceName string `json:"sourceServiceInterfaceName"`
}

type RoutingRuleRouteListResponse struct {
	SourceServiceInterfaceId   string `json:"sourceServiceInterfaceId,omitempty"`
	SourceServiceInterfaceName string `json:"sourceServiceInterfaceName,omitempty"`
}

type TgwRoutingRuleListResponse struct {
	ProjectId                        string       `json:"projectId,omitempty"`
	BlockId                          string       `json:"blockId,omitempty"`
	DestinationNetworkCidr           string       `json:"destinationNetworkCidr,omitempty"`
	Editable                         *bool        `json:"editable,omitempty"`
	RoutingRuleId                    string       `json:"routingRuleId,omitempty"`
	RoutingRuleState                 string       `json:"routingRuleState,omitempty"`
	ServiceZoneId                    string       `json:"serviceZoneId,omitempty"`
	SourceServiceInterfaceId         string       `json:"sourceServiceInterfaceId,omitempty"`
	SourceServiceInterfaceName       string       `json:"sourceServiceInterfaceName,omitempty"`
	SourceTransitGatewayConnectionId string       `json:"sourceTransitGatewayConnectionId,omitempty"`
	CreatedBy                        string       `json:"createdBy,omitempty"`
	CreatedDt                        time.Time `json:"createdDt,omitempty"`
	ModifiedBy                       string       `json:"modifiedBy,omitempty"`
	ModifiedDt                       time.Time `json:"modifiedDt,omitempty"`
}

type TgwRoutingTableDetailResponse struct {
	ProjectId                  string       `json:"projectId,omitempty"`
	BlockId                    string       `json:"blockId,omitempty"`
	RoutingRuleCount           int32        `json:"routingRuleCount,omitempty"`
	RoutingTableId             string       `json:"routingTableId,omitempty"`
	RoutingTableName           string       `json:"routingTableName,omitempty"`
	RoutingTableState          string       `json:"routingTableState,omitempty"`
	RoutingTableType           string       `json:"routingTableType,omitempty"`
	ServiceZoneId              string       `json:"serviceZoneId,omitempty"`
	T1RouterId                 string       `json:"t1RouterId,omitempty"`
	TransitGatewayConnectionId string       `json:"transitGatewayConnectionId,omitempty"`
	CreatedBy                  string       `json:"createdBy,omitempty"`
	CreatedDt                  time.Time `json:"createdDt,omitempty"`
	ModifiedBy                 string       `json:"modifiedBy,omitempty"`
	ModifiedDt                 time.Time `json:"modifiedDt,omitempty"`
}

type TgwRoutingTableListResponse struct {
	RoutingRuleCount           int64  `json:"routingRuleCount,omitempty"`
	RoutingTableId             string `json:"routingTableId,omitempty"`
	RoutingTableName           string `json:"routingTableName,omitempty"`
	RoutingTableType           string `json:"routingTableType,omitempty"`
	ServiceZoneId              string `json:"serviceZoneId,omitempty"`
	T1RouterId                 string `json:"t1RouterId,omitempty"`
	TransitGatewayConnectionId string `json:"transitGatewayConnectionId,omitempty"`
}

type TransitGatewayRoutingRuleOpenApiControllerApiListTgwRoutingRulesOpts struct {
	DestinationNetworkCidr   optional.String
	Editable                 optional.Bool
	RoutingRuleId            optional.String
	SourceServiceInterfaceId optional.String
	Page                     optional.Int32
	Size                     optional.Int32
	Sort                     optional.Interface
}

type TransitGatewayRoutingTableOpenApiControllerApiListTgwRoutingTablesOpts struct {
	RoutingTableId             optional.String
	RoutingTableName           optional.String
	TransitGatewayConnectionId optional.String
	CreatedBy                  optional.String
	Page                       optional.Int32
	Size                       optional.Int32
	Sort                       optional.Interface
}

type VpcRoutingRuleListResponse struct {
	ProjectId                  string       `json:"projectId,omitempty"`
	BlockId                    string       `json:"blockId,omitempty"`
	DestinationNetworkCidr     string       `json:"destinationNetworkCidr,omitempty"`
	Editable                   *bool        `json:"editable,omitempty"`
	RoutingRuleId              string       `json:"routingRuleId,omitempty"`
	RoutingRuleState           string       `json:"routingRuleState,omitempty"`
	ServiceZoneId              string       `json:"serviceZoneId,omitempty"`
	SourceServiceInterfaceId   string       `json:"sourceServiceInterfaceId,omitempty"`
	SourceServiceInterfaceName string       `json:"sourceServiceInterfaceName,omitempty"`
	SourceVpcId                string       `json:"sourceVpcId,omitempty"`
	CreatedBy                  string       `json:"createdBy,omitempty"`
	CreatedDt                  time.Time `json:"createdDt,omitempty"`
	ModifiedBy                 string       `json:"modifiedBy,omitempty"`
	ModifiedDt                 time.Time `json:"modifiedDt,omitempty"`
}

type VpcRoutingRuleOpenApiControllerApiListVpcRoutingRulesOpts struct {
	DestinationNetworkCidr   optional.String
	Editable                 optional.Bool
	RoutingRuleId            optional.String
	SourceServiceInterfaceId optional.String
	Page                     optional.Int32
	Size                     optional.Int32
	Sort                     optional.Interface
}

type VpcRoutingTableDetailResponse struct {
	ProjectId         string       `json:"projectId,omitempty"`
	BlockId           string       `json:"blockId,omitempty"`
	RoutingRuleCount  int32        `json:"routingRuleCount,omitempty"`
	RoutingTableId    string       `json:"routingTableId,omitempty"`
	RoutingTableName  string       `json:"routingTableName,omitempty"`
	RoutingTableState string       `json:"routingTableState,omitempty"`
	RoutingTableType  string       `json:"routingTableType,omitempty"`
	ServiceZoneId     string       `json:"serviceZoneId,omitempty"`
	T1RouterId        string       `json:"t1RouterId,omitempty"`
	VpcId             string       `json:"vpcId,omitempty"`
	CreatedBy         string       `json:"createdBy,omitempty"`
	CreatedDt         time.Time `json:"createdDt,omitempty"`
	ModifiedBy        string       `json:"modifiedBy,omitempty"`
	ModifiedDt        time.Time `json:"modifiedDt,omitempty"`
}

type VpcRoutingTableListResponse struct {
	RoutingRuleCount int64  `json:"routingRuleCount,omitempty"`
	RoutingTableId   string `json:"routingTableId,omitempty"`
	RoutingTableName string `json:"routingTableName,omitempty"`
	RoutingTableType string `json:"routingTableType,omitempty"`
	ServiceZoneId    string `json:"serviceZoneId,omitempty"`
	T1RouterId       string `json:"t1RouterId,omitempty"`
	VpcId            string `json:"vpcId,omitempty"`
}

type VpcRoutingTableOpenApiControllerApiListVpcRoutingTablesOpts struct {
	RoutingTableId   optional.String
	RoutingTableName optional.String
	VpcId            optional.String
	CreatedBy        optional.String
	Page             optional.Int32
	Size             optional.Int32
	Sort             optional.Interface
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
