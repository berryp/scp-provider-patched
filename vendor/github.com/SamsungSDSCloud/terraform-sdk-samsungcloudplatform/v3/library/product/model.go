package product

import (
	"github.com/antihax/optional"
	"time"
)

type APIKey struct {
	Key    string
	Prefix string
}

type BasicAuth struct {
	UserName string `json:"userName,omitempty"`
	Password string `json:"password,omitempty"`
}

type CommonCodeResponse struct {
	ClassificationCode string `json:"classificationCode,omitempty"`
	Code               string `json:"code,omitempty"`
	Name               string `json:"name,omitempty"`
	Description        string `json:"description,omitempty"`
}

type ItemDomain struct {
	Code          string `json:"code,omitempty"`
	Cost          string `json:"cost,omitempty"`
	CostUnitYn    string `json:"costUnitYn,omitempty"`
	Name          string `json:"name,omitempty"`
	ProductId     string `json:"productId,omitempty"`
	ProductIdList string `json:"productIdList,omitempty"`
	Properties    string `json:"properties,omitempty"`
	ServicedFor   string `json:"servicedFor,omitempty"`
	State         string `json:"state,omitempty"`
	Type_         string `json:"type,omitempty"`
	Unit          string `json:"unit,omitempty"`
	Value         string `json:"value,omitempty"`
	Version       string `json:"version,omitempty"`
	Description   string `json:"description,omitempty"`
}

type ItemForCalculatorResponse struct {
	ItemName  string `json:"itemName,omitempty"`
	ItemType  string `json:"itemType,omitempty"`
	ItemValue string `json:"itemValue,omitempty"`
}

type ItemVo struct {
	Code          string            `json:"code,omitempty"`
	Cost          string            `json:"cost,omitempty"`
	Name          string            `json:"name,omitempty"`
	ProductId     string            `json:"productId,omitempty"`
	ProductIdList string            `json:"productIdList,omitempty"`
	Properties    map[string]string `json:"properties,omitempty"`
	ServicedFor   string            `json:"servicedFor,omitempty"`
	State         string            `json:"state,omitempty"`
	Type_         string            `json:"type,omitempty"`
	Unit          string            `json:"unit,omitempty"`
	Value         string            `json:"value,omitempty"`
	Version       string            `json:"version,omitempty"`
	Description   string            `json:"description,omitempty"`
}

type ListCategoriesRequest struct {
	CategoryId    string
	CategoryState string
	ExposureScope string
	ProductId     string
	ProductState  string
	LanguageCode  string
}

type ListMenusRequest struct {
	CategoryId    string
	ExposureType  string
	ExposureScope string
	ProductId     string
	ZoneIds       string
}

type ListResponseV2MenuResponse struct {
	Contents   []MenuResponse `json:"contents,omitempty"`
	TotalCount int64                     `json:"totalCount,omitempty"`
}

type ListResponseV2ProductCategoryResponse struct {
	Contents   []ProductCategoryResponse `json:"contents,omitempty"`
	TotalCount int64                                `json:"totalCount,omitempty"`
}

type ListResponseV2ProductGroupForCalculatorResponse struct {
	Contents   []ProductGroupForCalculatorResponse `json:"contents,omitempty"`
	TotalCount int64                                          `json:"totalCount,omitempty"`
}

type ListResponseV2ProductGroupsResponse struct {
	Contents   []ProductGroupsResponse `json:"contents,omitempty"`
	TotalCount int64                              `json:"totalCount,omitempty"`
}

type ListResponseV2ProductsResponse1 struct {
	Contents   []ProductsResponse1 `json:"contents,omitempty"`
	TotalCount int64                          `json:"totalCount,omitempty"`
}

type MenuCategory struct {
	DefaultMenuId             string                          `json:"defaultMenuId,omitempty"`
	DelYn                     string                          `json:"delYn,omitempty"`
	FilePath                  string                          `json:"filePath,omitempty"`
	FunctionDescriptionsEn    []MenuDetailFuncDesc `json:"functionDescriptionsEn,omitempty"`
	FunctionDescriptionsKo    []MenuDetailFuncDesc `json:"functionDescriptionsKo,omitempty"`
	GuidePath                 string                          `json:"guidePath,omitempty"`
	IconFileName              string                          `json:"iconFileName,omitempty"`
	IconFilePath              string                          `json:"iconFilePath,omitempty"`
	Id                        string                          `json:"id,omitempty"`
	Key                       string                          `json:"key,omitempty"`
	Lang                      string                          `json:"lang,omitempty"`
	MenuExposureFor           string                          `json:"menuExposureFor,omitempty"`
	MessageKeyDesc            string                          `json:"messageKeyDesc,omitempty"`
	MessageKeyDescEn          string                          `json:"messageKeyDescEn,omitempty"`
	MessageKeyDescKo          string                          `json:"messageKeyDescKo,omitempty"`
	MessageKeyDetail          string                          `json:"messageKeyDetail,omitempty"`
	MessageKeyDetailEn        string                          `json:"messageKeyDetailEn,omitempty"`
	MessageKeyDetailKo        string                          `json:"messageKeyDetailKo,omitempty"`
	MessageKeyName            string                          `json:"messageKeyName,omitempty"`
	MessageKeySummary         string                          `json:"messageKeySummary,omitempty"`
	MessageKeySummaryEn       string                          `json:"messageKeySummaryEn,omitempty"`
	MessageKeySummaryKo       string                          `json:"messageKeySummaryKo,omitempty"`
	Name                      string                          `json:"name,omitempty"`
	NameEn                    string                          `json:"nameEn,omitempty"`
	ParentId                  string                          `json:"parentId,omitempty"`
	ProductConsolePath        string                          `json:"productConsolePath,omitempty"`
	ProductConsoleRequestPath string                          `json:"productConsoleRequestPath,omitempty"`
	ProductOfferingId         string                          `json:"productOfferingId,omitempty"`
	Productset                string                          `json:"productset,omitempty"`
	ResourceType              string                          `json:"resourceType,omitempty"`
	ResourceTypeId            string                          `json:"resourceTypeId,omitempty"`
	Route                     string                          `json:"route,omitempty"`
	RouteConsole              string                          `json:"routeConsole,omitempty"`
	RouteConsoleRequest       string                          `json:"routeConsoleRequest,omitempty"`
	Scope                     string                          `json:"scope,omitempty"`
	Seq                       int32                           `json:"seq,omitempty"`
	ShowQuery                 string                          `json:"showQuery,omitempty"`
	SrnName                   string                          `json:"srnName,omitempty"`
	State                     string                          `json:"state,omitempty"`
	Type_                     string                          `json:"type,omitempty"`
	Value                     string                          `json:"value,omitempty"`
	CreatedBy                 string                          `json:"createdBy,omitempty"`
	CreatedDt                 time.Time                    `json:"createdDt,omitempty"`
	ModifiedBy                string                          `json:"modifiedBy,omitempty"`
	ModifiedDt                time.Time                    `json:"modifiedDt,omitempty"`
}

type MenuDetailFuncDesc struct {
	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
}

type MenuDetailResponse struct {
	MenuDetailGuideFilePath string `json:"menuDetailGuideFilePath,omitempty"`
	MenuDetailGuidePath     string `json:"menuDetailGuidePath,omitempty"`
	MenuDetailId            string `json:"menuDetailId,omitempty"`
	MenuDetailParentId      string `json:"menuDetailParentId,omitempty"`
	MenuDetailType          string `json:"menuDetailType,omitempty"`
	MenuDetailDescription   string `json:"menuDetailDescription,omitempty"`
}

type MenuPriorProductResponse struct {
	PriorProductConsolePath  string `json:"priorProductConsolePath,omitempty"`
	PriorProductIconFilePath string `json:"priorProductIconFilePath,omitempty"`
	PriorProductId           string `json:"priorProductId,omitempty"`
	PriorProductName         string `json:"priorProductName,omitempty"`
	PriorProductOfferingId   string `json:"priorProductOfferingId,omitempty"`
	PriorProductParentId     string `json:"priorProductParentId,omitempty"`
	PriorProductRequestPath  string `json:"priorProductRequestPath,omitempty"`
	PriorProductSeq          string `json:"priorProductSeq,omitempty"`
	PriorProductState        string `json:"priorProductState,omitempty"`
	PriorProductDescription  string `json:"priorProductDescription,omitempty"`
}

type MenuResponse struct {
	AvailableButton        *bool                                 `json:"availableButton,omitempty"`
	ButtonName             string                                `json:"buttonName,omitempty"`
	DefaultMenuId          string                                `json:"defaultMenuId,omitempty"`
	MenuConsolePath        string                                `json:"menuConsolePath,omitempty"`
	MenuConsoleRequestPath string                                `json:"menuConsoleRequestPath,omitempty"`
	MenuDetails            []MenuDetailResponse       `json:"menuDetails,omitempty"`
	MenuExposureFor        string                                `json:"menuExposureFor,omitempty"`
	MenuId                 string                                `json:"menuId,omitempty"`
	MenuInformationPath    string                                `json:"menuInformationPath,omitempty"`
	MenuName               string                                `json:"menuName,omitempty"`
	MenuPriorProducts      []MenuPriorProductResponse `json:"menuPriorProducts,omitempty"`
	MenuResourceType       string                                `json:"menuResourceType,omitempty"`
	MenuState              string                                `json:"menuState,omitempty"`
	MenuType               string                                `json:"menuType,omitempty"`
	ParentId               string                                `json:"parentId,omitempty"`
}

type MenuTree struct {
	DefaultName string `json:"defaultName,omitempty"`
	Depth       int32  `json:"depth,omitempty"`
	Id          string `json:"id,omitempty"`
	IsLeaf      *bool  `json:"isLeaf,omitempty"`
	IsOpened    *bool  `json:"isOpened,omitempty"`
	IsRoot      *bool  `json:"isRoot,omitempty"`
	MenuKey     string `json:"menuKey,omitempty"`
	MessageKey  string `json:"messageKey,omitempty"`
	Name        string `json:"name,omitempty"`
	ParentId    string `json:"parentId,omitempty"`
	ParentKey   string `json:"parentKey,omitempty"`
	Seq         int32  `json:"seq,omitempty"`
	Type_       string `json:"type,omitempty"`
}

type PlatformControllerApiDetailApiDocsOpts struct {
	ApiScopes            optional.Interface
	BuildApiResponse     optional.Bool
	DisplayFieldRefDesc  optional.Bool
	ForceGenMsgCode      optional.Bool
	IgnoreInternalHeader optional.Bool
	MsgCodePrefix        optional.String
	State                optional.Interface
	UriPrefix            optional.Interface
	UseCliMsgCode        optional.Bool
	Version              optional.Interface
}

type ProductAdminMenuControllerApiGetDefaultProductsOpts struct {
	ProductType optional.String
}

type ProductCategoryResponse struct {
	IconFileName               string                               `json:"iconFileName,omitempty"`
	Product                    []ProductOfferingResponse `json:"product,omitempty"`
	ProductCategoryId          string                               `json:"productCategoryId,omitempty"`
	ProductCategoryName        string                               `json:"productCategoryName,omitempty"`
	ProductCategoryPath        string                               `json:"productCategoryPath,omitempty"`
	ProductCategoryState       string                               `json:"productCategoryState,omitempty"`
	ProductSet                 string                               `json:"productSet,omitempty"`
	ProductCategoryDescription string                               `json:"productCategoryDescription,omitempty"`
}

type ProductDomain struct {
	Id              string                  `json:"id,omitempty"`
	ImageProductSeq string                  `json:"imageProductSeq,omitempty"`
	ItemJson        string                  `json:"itemJson,omitempty"`
	Items           []ItemDomain `json:"items,omitempty"`
	Name            string                  `json:"name,omitempty"`
	Propertie       map[string]string       `json:"propertie,omitempty"`
	Properties      map[string]string       `json:"properties,omitempty"`
	Seq             string                  `json:"seq,omitempty"`
	State           string                  `json:"state,omitempty"`
	Type_           string                  `json:"type,omitempty"`
	Value           string                  `json:"value,omitempty"`
	Description     string                  `json:"description,omitempty"`
	CreatedBy       string                  `json:"createdBy,omitempty"`
	CreatedDt       time.Time            `json:"createdDt,omitempty"`
	ModifiedBy      string                  `json:"modifiedBy,omitempty"`
	ModifiedDt      time.Time            `json:"modifiedDt,omitempty"`
}

type ProductForCalculatorResponse struct {
	Item               []ItemForCalculatorResponse `json:"item,omitempty"`
	ProductId          string                                 `json:"productId,omitempty"`
	ProductName        string                                 `json:"productName,omitempty"`
	ProductState       string                                 `json:"productState,omitempty"`
	ProductType        string                                 `json:"productType,omitempty"`
	ProductDescription string                                 `json:"productDescription,omitempty"`
}

type ProductGroupDetailResponse struct {
	Products map[string][]ProductForCalculatorResponse `json:"products,omitempty"`
}

type ProductGroupForCalculatorResponse struct{
	ProductGroup []interface{}	`json:"productGroup,omitempty"`
	ProductGroupId string	`json:"productGroupId,omitempty"`
	ProductGroupName string	`json:"productGroupName,omitempty"`
	ProductGroupSequence string	`json:"productGroupSequence,omitempty"`
	ProductGroupType string	`json:"productGroupType,omitempty"`
	TargetProduct string	`json:"targetProduct,omitempty"`
	TargetProductGroup string	`json:"targetProductGroup,omitempty"`
}

type ProductGroupsResponse struct {
	ProductGroupId     string `json:"productGroupId,omitempty"`
	TargetProduct      string `json:"targetProduct,omitempty"`
	TargetProductGroup string `json:"targetProductGroup,omitempty"`
	CreatedBy          string `json:"createdBy,omitempty"`
	CreatedDt          string `json:"createdDt,omitempty"`
	ModifiedBy         string `json:"modifiedBy,omitempty"`
	ModifiedDt         string `json:"modifiedDt,omitempty"`
}

type ProductOfferingResponse struct {
	Menu                              []MenuResponse `json:"menu,omitempty"`
	NewBadgeEndDate                   string                    `json:"newBadgeEndDate,omitempty"`
	OfferingExposureFor               string                    `json:"offeringExposureFor,omitempty"`
	OfferingResourceType              string                    `json:"offeringResourceType,omitempty"`
	ProductOfferingConsolePath        string                    `json:"productOfferingConsolePath,omitempty"`
	ProductOfferingConsoleRequestPath string                    `json:"productOfferingConsoleRequestPath,omitempty"`
	ProductOfferingDetailInfo         string                    `json:"productOfferingDetailInfo,omitempty"`
	ProductOfferingId                 string                    `json:"productOfferingId,omitempty"`
	ProductOfferingName               string                    `json:"productOfferingName,omitempty"`
	ProductOfferingPath               string                    `json:"productOfferingPath,omitempty"`
	ProductOfferingState              string                    `json:"productOfferingState,omitempty"`
	Visible                           *bool                     `json:"visible,omitempty"`
	ProductOfferingDescription        string                    `json:"productOfferingDescription,omitempty"`
}

type ProductResponse struct {
	Id          string                   `json:"id,omitempty"`
	Items       []ItemVo      `json:"items,omitempty"`
	ItemsMap    []map[string]interface{} `json:"itemsMap,omitempty"`
	ItemsString string                   `json:"itemsString,omitempty"`
	Name        string                   `json:"name,omitempty"`
	Properties  map[string]string        `json:"properties,omitempty"`
	RateId      string                   `json:"rateId,omitempty"`
	Seq         string                   `json:"seq,omitempty"`
	State       string                   `json:"state,omitempty"`
	Type_       string                   `json:"type,omitempty"`
	Description string                   `json:"description,omitempty"`
	CreatedBy   string                   `json:"createdBy,omitempty"`
	CreatedDt   time.Time             `json:"createdDt,omitempty"`
	ModifiedBy  string                   `json:"modifiedBy,omitempty"`
	ModifiedDt  time.Time             `json:"modifiedDt,omitempty"`
}

type ProductV2ControllerApiDetailProductGroupOpts struct {
	ForCalculator optional.String
}

type ProductV2ControllerApiDetailProductOpts struct {
	ItemState optional.String
}

type ProductV2ControllerApiListCategoriesOpts struct {
	CategoryId    optional.String
	CategoryState optional.String
	ExposureScope optional.String
	ProductId     optional.String
	ProductState  optional.String
}

type ProductV2ControllerApiListMenusOpts struct {
	CategoryId   optional.String
	ExposureType optional.String
	ProductId    optional.String
	ZoneIds      optional.String
}

type ProductV2ControllerApiListProducsByZoneIdOpts struct {
	ProductGroupId optional.String
	ProductType    optional.String
}

type ProductV2ControllerApiListProductGroupsByZoneIdOpts struct {
	ObjectId           optional.String
	TargetProduct      optional.String
	TargetProductGroup optional.String
}

type ProductV2ControllerApiListProductGroupsOpts struct {
	ForCalculator      optional.String
	ProductGroupType   optional.String
	TargetProduct      optional.String
	TargetProductGroup optional.String
}

type ProductsResponse struct {
	Content []ProductDomain `json:"content,omitempty"`
	Total   int32                      `json:"total,omitempty"`
}

type ProductsResponse1 struct {
	ProductId   string `json:"productId,omitempty"`
	ProductName string `json:"productName,omitempty"`
	ProductType string `json:"productType,omitempty"`
	CreatedBy   string `json:"createdBy,omitempty"`
	CreatedDt   string `json:"createdDt,omitempty"`
	ModifiedBy  string `json:"modifiedBy,omitempty"`
	ModifiedDt  string `json:"modifiedDt,omitempty"`
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
