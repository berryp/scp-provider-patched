package keypair

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

type CreateRequest struct {
	KeyPairName string
	Tags        map[string]interface{}
}

type KeyPairCreateV1Request struct {
	KeyPairName string                  `json:"keyPairName"`
	Tags        []TagRequest `json:"tags,omitempty"`
}

type KeyPairCreateV1Response struct {
	ProjectId    string       `json:"projectId,omitempty"`
	KeyPairId    string       `json:"keyPairId,omitempty"`
	KeyPairName  string       `json:"keyPairName,omitempty"`
	KeyPairState string       `json:"keyPairState,omitempty"`
	PrivateKey   string       `json:"privateKey,omitempty"`
	CreatedBy    string       `json:"createdBy,omitempty"`
	CreatedDt    time.Time `json:"createdDt,omitempty"`
	ModifiedBy   string       `json:"modifiedBy,omitempty"`
	ModifiedDt   time.Time `json:"modifiedDt,omitempty"`
}

type KeyPairV1ApiListKeyPairsOpts struct {
	KeyPairName optional.String
	CreatedBy   optional.String
	Page        optional.Int32
	Size        optional.Int32
	Sort        optional.Interface
}

type KeyPairV1Response struct {
	ProjectId                 string       `json:"projectId,omitempty"`
	KeyPairId                 string       `json:"keyPairId,omitempty"`
	KeyPairName               string       `json:"keyPairName,omitempty"`
	KeyPairState              string       `json:"keyPairState,omitempty"`
	LaunchConfigurationIdList []string     `json:"launchConfigurationIdList,omitempty"`
	VirtualServerIdList       []string     `json:"virtualServerIdList,omitempty"`
	CreatedBy                 string       `json:"createdBy,omitempty"`
	CreatedDt                 time.Time `json:"createdDt,omitempty"`
	ModifiedBy                string       `json:"modifiedBy,omitempty"`
	ModifiedDt                time.Time `json:"modifiedDt,omitempty"`
}

type ListKeyPairsRequestParam struct {
	KeyPairName string
	CreatedBy   string
	Page        int32
	Size        int32
	Sort        string
}

type ListResponseKeyPairV1Response struct {
	Contents   []KeyPairV1Response `json:"contents,omitempty"`
	TotalCount int32                          `json:"totalCount,omitempty"`
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
