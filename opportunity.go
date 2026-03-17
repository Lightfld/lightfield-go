// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package lightfield

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/Lightfld/lightfield-go/internal/apijson"
	"github.com/Lightfld/lightfield-go/internal/apiquery"
	"github.com/Lightfld/lightfield-go/internal/requestconfig"
	"github.com/Lightfld/lightfield-go/option"
	"github.com/Lightfld/lightfield-go/packages/param"
	"github.com/Lightfld/lightfield-go/packages/respjson"
)

// OpportunityService contains methods and other services that help with
// interacting with the Lightfield API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewOpportunityService] method instead.
type OpportunityService struct {
	Options []option.RequestOption
}

// NewOpportunityService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewOpportunityService(opts ...option.RequestOption) (r OpportunityService) {
	r = OpportunityService{}
	r.Options = opts
	return
}

func (r *OpportunityService) New(ctx context.Context, body OpportunityNewParams, opts ...option.RequestOption) (res *OpportunityNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/opportunities"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

func (r *OpportunityService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *OpportunityGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("v1/opportunities/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

func (r *OpportunityService) Update(ctx context.Context, id string, body OpportunityUpdateParams, opts ...option.RequestOption) (res *OpportunityUpdateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("v1/opportunities/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

func (r *OpportunityService) List(ctx context.Context, query OpportunityListParams, opts ...option.RequestOption) (res *OpportunityListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/opportunities"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type OpportunityNewResponse struct {
	ID            string                                        `json:"id" api:"required"`
	CreatedAt     string                                        `json:"createdAt" api:"required"`
	Fields        map[string]OpportunityNewResponseField        `json:"fields" api:"required"`
	HTTPLink      string                                        `json:"httpLink" api:"required"`
	Relationships map[string]OpportunityNewResponseRelationship `json:"relationships" api:"required"`
	ExtraFields   map[string]OpportunityNewResponseUnion        `json:"" api:"extrafields"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID            respjson.Field
		CreatedAt     respjson.Field
		Fields        respjson.Field
		HTTPLink      respjson.Field
		Relationships respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OpportunityNewResponse) RawJSON() string { return r.JSON.raw }
func (r *OpportunityNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OpportunityNewResponseField struct {
	Value     OpportunityNewResponseFieldValueUnion `json:"value" api:"required"`
	ValueType string                                `json:"valueType" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Value       respjson.Field
		ValueType   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OpportunityNewResponseField) RawJSON() string { return r.JSON.raw }
func (r *OpportunityNewResponseField) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// OpportunityNewResponseFieldValueUnion contains all possible properties and
// values from [string], [float64], [bool],
// [[]OpportunityNewResponseFieldValueArrayItemUnion],
// [map[string]OpportunityNewResponseFieldValueMapItemUnion].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfOpportunityNewResponseFieldValueArray
// OfAnyArray OfOpportunityNewResponseFieldValueMapItemMapItem]
type OpportunityNewResponseFieldValueUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a
	// [[]OpportunityNewResponseFieldValueArrayItemUnion] instead of an object.
	OfOpportunityNewResponseFieldValueArray []OpportunityNewResponseFieldValueArrayItemUnion `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	// This field will be present if the value is a [any] instead of an object.
	OfOpportunityNewResponseFieldValueMapItemMapItem any `json:",inline"`
	JSON                                             struct {
		OfString                                         respjson.Field
		OfFloat                                          respjson.Field
		OfBool                                           respjson.Field
		OfOpportunityNewResponseFieldValueArray          respjson.Field
		OfAnyArray                                       respjson.Field
		OfOpportunityNewResponseFieldValueMapItemMapItem respjson.Field
		raw                                              string
	} `json:"-"`
}

func (u OpportunityNewResponseFieldValueUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u OpportunityNewResponseFieldValueUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u OpportunityNewResponseFieldValueUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u OpportunityNewResponseFieldValueUnion) AsOpportunityNewResponseFieldValueArray() (v []OpportunityNewResponseFieldValueArrayItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u OpportunityNewResponseFieldValueUnion) AsOpportunityNewResponseFieldValueMapMap() (v map[string]OpportunityNewResponseFieldValueMapItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u OpportunityNewResponseFieldValueUnion) RawJSON() string { return u.JSON.raw }

func (r *OpportunityNewResponseFieldValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// OpportunityNewResponseFieldValueArrayItemUnion contains all possible properties
// and values from [string], [float64], [bool], [[]any], [map[string]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray
// OfOpportunityNewResponseFieldValueArrayItemMapItem]
type OpportunityNewResponseFieldValueArrayItemUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	// This field will be present if the value is a [any] instead of an object.
	OfOpportunityNewResponseFieldValueArrayItemMapItem any `json:",inline"`
	JSON                                               struct {
		OfString                                           respjson.Field
		OfFloat                                            respjson.Field
		OfBool                                             respjson.Field
		OfAnyArray                                         respjson.Field
		OfOpportunityNewResponseFieldValueArrayItemMapItem respjson.Field
		raw                                                string
	} `json:"-"`
}

func (u OpportunityNewResponseFieldValueArrayItemUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u OpportunityNewResponseFieldValueArrayItemUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u OpportunityNewResponseFieldValueArrayItemUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u OpportunityNewResponseFieldValueArrayItemUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u OpportunityNewResponseFieldValueArrayItemUnion) AsAnyMap() (v map[string]any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u OpportunityNewResponseFieldValueArrayItemUnion) RawJSON() string { return u.JSON.raw }

func (r *OpportunityNewResponseFieldValueArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// OpportunityNewResponseFieldValueMapItemUnion contains all possible properties
// and values from [string], [float64], [bool], [[]any], [map[string]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray
// OfOpportunityNewResponseFieldValueMapItemMapItem]
type OpportunityNewResponseFieldValueMapItemUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	// This field will be present if the value is a [any] instead of an object.
	OfOpportunityNewResponseFieldValueMapItemMapItem any `json:",inline"`
	JSON                                             struct {
		OfString                                         respjson.Field
		OfFloat                                          respjson.Field
		OfBool                                           respjson.Field
		OfAnyArray                                       respjson.Field
		OfOpportunityNewResponseFieldValueMapItemMapItem respjson.Field
		raw                                              string
	} `json:"-"`
}

func (u OpportunityNewResponseFieldValueMapItemUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u OpportunityNewResponseFieldValueMapItemUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u OpportunityNewResponseFieldValueMapItemUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u OpportunityNewResponseFieldValueMapItemUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u OpportunityNewResponseFieldValueMapItemUnion) AsAnyMap() (v map[string]any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u OpportunityNewResponseFieldValueMapItemUnion) RawJSON() string { return u.JSON.raw }

func (r *OpportunityNewResponseFieldValueMapItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OpportunityNewResponseRelationship struct {
	Cardinality string   `json:"cardinality" api:"required"`
	ObjectType  string   `json:"objectType" api:"required"`
	Values      []string `json:"values" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Cardinality respjson.Field
		ObjectType  respjson.Field
		Values      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OpportunityNewResponseRelationship) RawJSON() string { return r.JSON.raw }
func (r *OpportunityNewResponseRelationship) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// OpportunityNewResponseUnion contains all possible properties and values from
// [string], [float64], [bool], [[]OpportunityNewResponseArrayItemUnion],
// [map[string]OpportunityNewResponseMapItemUnion].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfOpportunityNewResponseArray OfAnyArray
// OfOpportunityNewResponseMapItemMapItem]
type OpportunityNewResponseUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a
	// [[]OpportunityNewResponseArrayItemUnion] instead of an object.
	OfOpportunityNewResponseArray []OpportunityNewResponseArrayItemUnion `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	// This field will be present if the value is a [any] instead of an object.
	OfOpportunityNewResponseMapItemMapItem any `json:",inline"`
	JSON                                   struct {
		OfString                               respjson.Field
		OfFloat                                respjson.Field
		OfBool                                 respjson.Field
		OfOpportunityNewResponseArray          respjson.Field
		OfAnyArray                             respjson.Field
		OfOpportunityNewResponseMapItemMapItem respjson.Field
		raw                                    string
	} `json:"-"`
}

func (u OpportunityNewResponseUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u OpportunityNewResponseUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u OpportunityNewResponseUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u OpportunityNewResponseUnion) AsOpportunityNewResponseArray() (v []OpportunityNewResponseArrayItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u OpportunityNewResponseUnion) AsOpportunityNewResponseMapMap() (v map[string]OpportunityNewResponseMapItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u OpportunityNewResponseUnion) RawJSON() string { return u.JSON.raw }

func (r *OpportunityNewResponseUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// OpportunityNewResponseArrayItemUnion contains all possible properties and values
// from [string], [float64], [bool], [[]any], [map[string]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray
// OfOpportunityNewResponseArrayItemMapItem]
type OpportunityNewResponseArrayItemUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	// This field will be present if the value is a [any] instead of an object.
	OfOpportunityNewResponseArrayItemMapItem any `json:",inline"`
	JSON                                     struct {
		OfString                                 respjson.Field
		OfFloat                                  respjson.Field
		OfBool                                   respjson.Field
		OfAnyArray                               respjson.Field
		OfOpportunityNewResponseArrayItemMapItem respjson.Field
		raw                                      string
	} `json:"-"`
}

func (u OpportunityNewResponseArrayItemUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u OpportunityNewResponseArrayItemUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u OpportunityNewResponseArrayItemUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u OpportunityNewResponseArrayItemUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u OpportunityNewResponseArrayItemUnion) AsAnyMap() (v map[string]any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u OpportunityNewResponseArrayItemUnion) RawJSON() string { return u.JSON.raw }

func (r *OpportunityNewResponseArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// OpportunityNewResponseMapItemUnion contains all possible properties and values
// from [string], [float64], [bool], [[]any], [map[string]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray
// OfOpportunityNewResponseMapItemMapItem]
type OpportunityNewResponseMapItemUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	// This field will be present if the value is a [any] instead of an object.
	OfOpportunityNewResponseMapItemMapItem any `json:",inline"`
	JSON                                   struct {
		OfString                               respjson.Field
		OfFloat                                respjson.Field
		OfBool                                 respjson.Field
		OfAnyArray                             respjson.Field
		OfOpportunityNewResponseMapItemMapItem respjson.Field
		raw                                    string
	} `json:"-"`
}

func (u OpportunityNewResponseMapItemUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u OpportunityNewResponseMapItemUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u OpportunityNewResponseMapItemUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u OpportunityNewResponseMapItemUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u OpportunityNewResponseMapItemUnion) AsAnyMap() (v map[string]any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u OpportunityNewResponseMapItemUnion) RawJSON() string { return u.JSON.raw }

func (r *OpportunityNewResponseMapItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OpportunityGetResponse struct {
	ID            string                                        `json:"id" api:"required"`
	CreatedAt     string                                        `json:"createdAt" api:"required"`
	Fields        map[string]OpportunityGetResponseField        `json:"fields" api:"required"`
	HTTPLink      string                                        `json:"httpLink" api:"required"`
	Relationships map[string]OpportunityGetResponseRelationship `json:"relationships" api:"required"`
	ExtraFields   map[string]OpportunityGetResponseUnion        `json:"" api:"extrafields"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID            respjson.Field
		CreatedAt     respjson.Field
		Fields        respjson.Field
		HTTPLink      respjson.Field
		Relationships respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OpportunityGetResponse) RawJSON() string { return r.JSON.raw }
func (r *OpportunityGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OpportunityGetResponseField struct {
	Value     OpportunityGetResponseFieldValueUnion `json:"value" api:"required"`
	ValueType string                                `json:"valueType" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Value       respjson.Field
		ValueType   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OpportunityGetResponseField) RawJSON() string { return r.JSON.raw }
func (r *OpportunityGetResponseField) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// OpportunityGetResponseFieldValueUnion contains all possible properties and
// values from [string], [float64], [bool],
// [[]OpportunityGetResponseFieldValueArrayItemUnion],
// [map[string]OpportunityGetResponseFieldValueMapItemUnion].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfOpportunityGetResponseFieldValueArray
// OfAnyArray OfOpportunityGetResponseFieldValueMapItemMapItem]
type OpportunityGetResponseFieldValueUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a
	// [[]OpportunityGetResponseFieldValueArrayItemUnion] instead of an object.
	OfOpportunityGetResponseFieldValueArray []OpportunityGetResponseFieldValueArrayItemUnion `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	// This field will be present if the value is a [any] instead of an object.
	OfOpportunityGetResponseFieldValueMapItemMapItem any `json:",inline"`
	JSON                                             struct {
		OfString                                         respjson.Field
		OfFloat                                          respjson.Field
		OfBool                                           respjson.Field
		OfOpportunityGetResponseFieldValueArray          respjson.Field
		OfAnyArray                                       respjson.Field
		OfOpportunityGetResponseFieldValueMapItemMapItem respjson.Field
		raw                                              string
	} `json:"-"`
}

func (u OpportunityGetResponseFieldValueUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u OpportunityGetResponseFieldValueUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u OpportunityGetResponseFieldValueUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u OpportunityGetResponseFieldValueUnion) AsOpportunityGetResponseFieldValueArray() (v []OpportunityGetResponseFieldValueArrayItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u OpportunityGetResponseFieldValueUnion) AsOpportunityGetResponseFieldValueMapMap() (v map[string]OpportunityGetResponseFieldValueMapItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u OpportunityGetResponseFieldValueUnion) RawJSON() string { return u.JSON.raw }

func (r *OpportunityGetResponseFieldValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// OpportunityGetResponseFieldValueArrayItemUnion contains all possible properties
// and values from [string], [float64], [bool], [[]any], [map[string]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray
// OfOpportunityGetResponseFieldValueArrayItemMapItem]
type OpportunityGetResponseFieldValueArrayItemUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	// This field will be present if the value is a [any] instead of an object.
	OfOpportunityGetResponseFieldValueArrayItemMapItem any `json:",inline"`
	JSON                                               struct {
		OfString                                           respjson.Field
		OfFloat                                            respjson.Field
		OfBool                                             respjson.Field
		OfAnyArray                                         respjson.Field
		OfOpportunityGetResponseFieldValueArrayItemMapItem respjson.Field
		raw                                                string
	} `json:"-"`
}

func (u OpportunityGetResponseFieldValueArrayItemUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u OpportunityGetResponseFieldValueArrayItemUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u OpportunityGetResponseFieldValueArrayItemUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u OpportunityGetResponseFieldValueArrayItemUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u OpportunityGetResponseFieldValueArrayItemUnion) AsAnyMap() (v map[string]any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u OpportunityGetResponseFieldValueArrayItemUnion) RawJSON() string { return u.JSON.raw }

func (r *OpportunityGetResponseFieldValueArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// OpportunityGetResponseFieldValueMapItemUnion contains all possible properties
// and values from [string], [float64], [bool], [[]any], [map[string]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray
// OfOpportunityGetResponseFieldValueMapItemMapItem]
type OpportunityGetResponseFieldValueMapItemUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	// This field will be present if the value is a [any] instead of an object.
	OfOpportunityGetResponseFieldValueMapItemMapItem any `json:",inline"`
	JSON                                             struct {
		OfString                                         respjson.Field
		OfFloat                                          respjson.Field
		OfBool                                           respjson.Field
		OfAnyArray                                       respjson.Field
		OfOpportunityGetResponseFieldValueMapItemMapItem respjson.Field
		raw                                              string
	} `json:"-"`
}

func (u OpportunityGetResponseFieldValueMapItemUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u OpportunityGetResponseFieldValueMapItemUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u OpportunityGetResponseFieldValueMapItemUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u OpportunityGetResponseFieldValueMapItemUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u OpportunityGetResponseFieldValueMapItemUnion) AsAnyMap() (v map[string]any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u OpportunityGetResponseFieldValueMapItemUnion) RawJSON() string { return u.JSON.raw }

func (r *OpportunityGetResponseFieldValueMapItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OpportunityGetResponseRelationship struct {
	Cardinality string   `json:"cardinality" api:"required"`
	ObjectType  string   `json:"objectType" api:"required"`
	Values      []string `json:"values" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Cardinality respjson.Field
		ObjectType  respjson.Field
		Values      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OpportunityGetResponseRelationship) RawJSON() string { return r.JSON.raw }
func (r *OpportunityGetResponseRelationship) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// OpportunityGetResponseUnion contains all possible properties and values from
// [string], [float64], [bool], [[]OpportunityGetResponseArrayItemUnion],
// [map[string]OpportunityGetResponseMapItemUnion].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfOpportunityGetResponseArray OfAnyArray
// OfOpportunityGetResponseMapItemMapItem]
type OpportunityGetResponseUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a
	// [[]OpportunityGetResponseArrayItemUnion] instead of an object.
	OfOpportunityGetResponseArray []OpportunityGetResponseArrayItemUnion `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	// This field will be present if the value is a [any] instead of an object.
	OfOpportunityGetResponseMapItemMapItem any `json:",inline"`
	JSON                                   struct {
		OfString                               respjson.Field
		OfFloat                                respjson.Field
		OfBool                                 respjson.Field
		OfOpportunityGetResponseArray          respjson.Field
		OfAnyArray                             respjson.Field
		OfOpportunityGetResponseMapItemMapItem respjson.Field
		raw                                    string
	} `json:"-"`
}

func (u OpportunityGetResponseUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u OpportunityGetResponseUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u OpportunityGetResponseUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u OpportunityGetResponseUnion) AsOpportunityGetResponseArray() (v []OpportunityGetResponseArrayItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u OpportunityGetResponseUnion) AsOpportunityGetResponseMapMap() (v map[string]OpportunityGetResponseMapItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u OpportunityGetResponseUnion) RawJSON() string { return u.JSON.raw }

func (r *OpportunityGetResponseUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// OpportunityGetResponseArrayItemUnion contains all possible properties and values
// from [string], [float64], [bool], [[]any], [map[string]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray
// OfOpportunityGetResponseArrayItemMapItem]
type OpportunityGetResponseArrayItemUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	// This field will be present if the value is a [any] instead of an object.
	OfOpportunityGetResponseArrayItemMapItem any `json:",inline"`
	JSON                                     struct {
		OfString                                 respjson.Field
		OfFloat                                  respjson.Field
		OfBool                                   respjson.Field
		OfAnyArray                               respjson.Field
		OfOpportunityGetResponseArrayItemMapItem respjson.Field
		raw                                      string
	} `json:"-"`
}

func (u OpportunityGetResponseArrayItemUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u OpportunityGetResponseArrayItemUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u OpportunityGetResponseArrayItemUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u OpportunityGetResponseArrayItemUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u OpportunityGetResponseArrayItemUnion) AsAnyMap() (v map[string]any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u OpportunityGetResponseArrayItemUnion) RawJSON() string { return u.JSON.raw }

func (r *OpportunityGetResponseArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// OpportunityGetResponseMapItemUnion contains all possible properties and values
// from [string], [float64], [bool], [[]any], [map[string]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray
// OfOpportunityGetResponseMapItemMapItem]
type OpportunityGetResponseMapItemUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	// This field will be present if the value is a [any] instead of an object.
	OfOpportunityGetResponseMapItemMapItem any `json:",inline"`
	JSON                                   struct {
		OfString                               respjson.Field
		OfFloat                                respjson.Field
		OfBool                                 respjson.Field
		OfAnyArray                             respjson.Field
		OfOpportunityGetResponseMapItemMapItem respjson.Field
		raw                                    string
	} `json:"-"`
}

func (u OpportunityGetResponseMapItemUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u OpportunityGetResponseMapItemUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u OpportunityGetResponseMapItemUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u OpportunityGetResponseMapItemUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u OpportunityGetResponseMapItemUnion) AsAnyMap() (v map[string]any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u OpportunityGetResponseMapItemUnion) RawJSON() string { return u.JSON.raw }

func (r *OpportunityGetResponseMapItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OpportunityUpdateResponse struct {
	ID            string                                           `json:"id" api:"required"`
	CreatedAt     string                                           `json:"createdAt" api:"required"`
	Fields        map[string]OpportunityUpdateResponseField        `json:"fields" api:"required"`
	HTTPLink      string                                           `json:"httpLink" api:"required"`
	Relationships map[string]OpportunityUpdateResponseRelationship `json:"relationships" api:"required"`
	ExtraFields   map[string]OpportunityUpdateResponseUnion        `json:"" api:"extrafields"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID            respjson.Field
		CreatedAt     respjson.Field
		Fields        respjson.Field
		HTTPLink      respjson.Field
		Relationships respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OpportunityUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *OpportunityUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OpportunityUpdateResponseField struct {
	Value     OpportunityUpdateResponseFieldValueUnion `json:"value" api:"required"`
	ValueType string                                   `json:"valueType" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Value       respjson.Field
		ValueType   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OpportunityUpdateResponseField) RawJSON() string { return r.JSON.raw }
func (r *OpportunityUpdateResponseField) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// OpportunityUpdateResponseFieldValueUnion contains all possible properties and
// values from [string], [float64], [bool],
// [[]OpportunityUpdateResponseFieldValueArrayItemUnion],
// [map[string]OpportunityUpdateResponseFieldValueMapItemUnion].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool
// OfOpportunityUpdateResponseFieldValueArray OfAnyArray
// OfOpportunityUpdateResponseFieldValueMapItemMapItem]
type OpportunityUpdateResponseFieldValueUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a
	// [[]OpportunityUpdateResponseFieldValueArrayItemUnion] instead of an object.
	OfOpportunityUpdateResponseFieldValueArray []OpportunityUpdateResponseFieldValueArrayItemUnion `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	// This field will be present if the value is a [any] instead of an object.
	OfOpportunityUpdateResponseFieldValueMapItemMapItem any `json:",inline"`
	JSON                                                struct {
		OfString                                            respjson.Field
		OfFloat                                             respjson.Field
		OfBool                                              respjson.Field
		OfOpportunityUpdateResponseFieldValueArray          respjson.Field
		OfAnyArray                                          respjson.Field
		OfOpportunityUpdateResponseFieldValueMapItemMapItem respjson.Field
		raw                                                 string
	} `json:"-"`
}

func (u OpportunityUpdateResponseFieldValueUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u OpportunityUpdateResponseFieldValueUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u OpportunityUpdateResponseFieldValueUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u OpportunityUpdateResponseFieldValueUnion) AsOpportunityUpdateResponseFieldValueArray() (v []OpportunityUpdateResponseFieldValueArrayItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u OpportunityUpdateResponseFieldValueUnion) AsOpportunityUpdateResponseFieldValueMapMap() (v map[string]OpportunityUpdateResponseFieldValueMapItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u OpportunityUpdateResponseFieldValueUnion) RawJSON() string { return u.JSON.raw }

func (r *OpportunityUpdateResponseFieldValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// OpportunityUpdateResponseFieldValueArrayItemUnion contains all possible
// properties and values from [string], [float64], [bool], [[]any],
// [map[string]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray
// OfOpportunityUpdateResponseFieldValueArrayItemMapItem]
type OpportunityUpdateResponseFieldValueArrayItemUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	// This field will be present if the value is a [any] instead of an object.
	OfOpportunityUpdateResponseFieldValueArrayItemMapItem any `json:",inline"`
	JSON                                                  struct {
		OfString                                              respjson.Field
		OfFloat                                               respjson.Field
		OfBool                                                respjson.Field
		OfAnyArray                                            respjson.Field
		OfOpportunityUpdateResponseFieldValueArrayItemMapItem respjson.Field
		raw                                                   string
	} `json:"-"`
}

func (u OpportunityUpdateResponseFieldValueArrayItemUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u OpportunityUpdateResponseFieldValueArrayItemUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u OpportunityUpdateResponseFieldValueArrayItemUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u OpportunityUpdateResponseFieldValueArrayItemUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u OpportunityUpdateResponseFieldValueArrayItemUnion) AsAnyMap() (v map[string]any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u OpportunityUpdateResponseFieldValueArrayItemUnion) RawJSON() string { return u.JSON.raw }

func (r *OpportunityUpdateResponseFieldValueArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// OpportunityUpdateResponseFieldValueMapItemUnion contains all possible properties
// and values from [string], [float64], [bool], [[]any], [map[string]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray
// OfOpportunityUpdateResponseFieldValueMapItemMapItem]
type OpportunityUpdateResponseFieldValueMapItemUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	// This field will be present if the value is a [any] instead of an object.
	OfOpportunityUpdateResponseFieldValueMapItemMapItem any `json:",inline"`
	JSON                                                struct {
		OfString                                            respjson.Field
		OfFloat                                             respjson.Field
		OfBool                                              respjson.Field
		OfAnyArray                                          respjson.Field
		OfOpportunityUpdateResponseFieldValueMapItemMapItem respjson.Field
		raw                                                 string
	} `json:"-"`
}

func (u OpportunityUpdateResponseFieldValueMapItemUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u OpportunityUpdateResponseFieldValueMapItemUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u OpportunityUpdateResponseFieldValueMapItemUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u OpportunityUpdateResponseFieldValueMapItemUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u OpportunityUpdateResponseFieldValueMapItemUnion) AsAnyMap() (v map[string]any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u OpportunityUpdateResponseFieldValueMapItemUnion) RawJSON() string { return u.JSON.raw }

func (r *OpportunityUpdateResponseFieldValueMapItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OpportunityUpdateResponseRelationship struct {
	Cardinality string   `json:"cardinality" api:"required"`
	ObjectType  string   `json:"objectType" api:"required"`
	Values      []string `json:"values" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Cardinality respjson.Field
		ObjectType  respjson.Field
		Values      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OpportunityUpdateResponseRelationship) RawJSON() string { return r.JSON.raw }
func (r *OpportunityUpdateResponseRelationship) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// OpportunityUpdateResponseUnion contains all possible properties and values from
// [string], [float64], [bool], [[]OpportunityUpdateResponseArrayItemUnion],
// [map[string]OpportunityUpdateResponseMapItemUnion].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfOpportunityUpdateResponseArray
// OfAnyArray OfOpportunityUpdateResponseMapItemMapItem]
type OpportunityUpdateResponseUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a
	// [[]OpportunityUpdateResponseArrayItemUnion] instead of an object.
	OfOpportunityUpdateResponseArray []OpportunityUpdateResponseArrayItemUnion `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	// This field will be present if the value is a [any] instead of an object.
	OfOpportunityUpdateResponseMapItemMapItem any `json:",inline"`
	JSON                                      struct {
		OfString                                  respjson.Field
		OfFloat                                   respjson.Field
		OfBool                                    respjson.Field
		OfOpportunityUpdateResponseArray          respjson.Field
		OfAnyArray                                respjson.Field
		OfOpportunityUpdateResponseMapItemMapItem respjson.Field
		raw                                       string
	} `json:"-"`
}

func (u OpportunityUpdateResponseUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u OpportunityUpdateResponseUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u OpportunityUpdateResponseUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u OpportunityUpdateResponseUnion) AsOpportunityUpdateResponseArray() (v []OpportunityUpdateResponseArrayItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u OpportunityUpdateResponseUnion) AsOpportunityUpdateResponseMapMap() (v map[string]OpportunityUpdateResponseMapItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u OpportunityUpdateResponseUnion) RawJSON() string { return u.JSON.raw }

func (r *OpportunityUpdateResponseUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// OpportunityUpdateResponseArrayItemUnion contains all possible properties and
// values from [string], [float64], [bool], [[]any], [map[string]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray
// OfOpportunityUpdateResponseArrayItemMapItem]
type OpportunityUpdateResponseArrayItemUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	// This field will be present if the value is a [any] instead of an object.
	OfOpportunityUpdateResponseArrayItemMapItem any `json:",inline"`
	JSON                                        struct {
		OfString                                    respjson.Field
		OfFloat                                     respjson.Field
		OfBool                                      respjson.Field
		OfAnyArray                                  respjson.Field
		OfOpportunityUpdateResponseArrayItemMapItem respjson.Field
		raw                                         string
	} `json:"-"`
}

func (u OpportunityUpdateResponseArrayItemUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u OpportunityUpdateResponseArrayItemUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u OpportunityUpdateResponseArrayItemUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u OpportunityUpdateResponseArrayItemUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u OpportunityUpdateResponseArrayItemUnion) AsAnyMap() (v map[string]any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u OpportunityUpdateResponseArrayItemUnion) RawJSON() string { return u.JSON.raw }

func (r *OpportunityUpdateResponseArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// OpportunityUpdateResponseMapItemUnion contains all possible properties and
// values from [string], [float64], [bool], [[]any], [map[string]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray
// OfOpportunityUpdateResponseMapItemMapItem]
type OpportunityUpdateResponseMapItemUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	// This field will be present if the value is a [any] instead of an object.
	OfOpportunityUpdateResponseMapItemMapItem any `json:",inline"`
	JSON                                      struct {
		OfString                                  respjson.Field
		OfFloat                                   respjson.Field
		OfBool                                    respjson.Field
		OfAnyArray                                respjson.Field
		OfOpportunityUpdateResponseMapItemMapItem respjson.Field
		raw                                       string
	} `json:"-"`
}

func (u OpportunityUpdateResponseMapItemUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u OpportunityUpdateResponseMapItemUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u OpportunityUpdateResponseMapItemUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u OpportunityUpdateResponseMapItemUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u OpportunityUpdateResponseMapItemUnion) AsAnyMap() (v map[string]any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u OpportunityUpdateResponseMapItemUnion) RawJSON() string { return u.JSON.raw }

func (r *OpportunityUpdateResponseMapItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OpportunityListResponse struct {
	Data       []OpportunityListResponseData `json:"data" api:"required"`
	Object     string                        `json:"object" api:"required"`
	TotalCount int64                         `json:"totalCount" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Object      respjson.Field
		TotalCount  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OpportunityListResponse) RawJSON() string { return r.JSON.raw }
func (r *OpportunityListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OpportunityListResponseData struct {
	ID            string                                             `json:"id" api:"required"`
	CreatedAt     string                                             `json:"createdAt" api:"required"`
	Fields        map[string]OpportunityListResponseDataField        `json:"fields" api:"required"`
	HTTPLink      string                                             `json:"httpLink" api:"required"`
	Relationships map[string]OpportunityListResponseDataRelationship `json:"relationships" api:"required"`
	ExtraFields   map[string]OpportunityListResponseDataUnion        `json:"" api:"extrafields"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID            respjson.Field
		CreatedAt     respjson.Field
		Fields        respjson.Field
		HTTPLink      respjson.Field
		Relationships respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OpportunityListResponseData) RawJSON() string { return r.JSON.raw }
func (r *OpportunityListResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OpportunityListResponseDataField struct {
	Value     OpportunityListResponseDataFieldValueUnion `json:"value" api:"required"`
	ValueType string                                     `json:"valueType" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Value       respjson.Field
		ValueType   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OpportunityListResponseDataField) RawJSON() string { return r.JSON.raw }
func (r *OpportunityListResponseDataField) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// OpportunityListResponseDataFieldValueUnion contains all possible properties and
// values from [string], [float64], [bool],
// [[]OpportunityListResponseDataFieldValueArrayItemUnion],
// [map[string]OpportunityListResponseDataFieldValueMapItemUnion].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool
// OfOpportunityListResponseDataFieldValueArray OfAnyArray
// OfOpportunityListResponseDataFieldValueMapItemMapItem]
type OpportunityListResponseDataFieldValueUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a
	// [[]OpportunityListResponseDataFieldValueArrayItemUnion] instead of an object.
	OfOpportunityListResponseDataFieldValueArray []OpportunityListResponseDataFieldValueArrayItemUnion `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	// This field will be present if the value is a [any] instead of an object.
	OfOpportunityListResponseDataFieldValueMapItemMapItem any `json:",inline"`
	JSON                                                  struct {
		OfString                                              respjson.Field
		OfFloat                                               respjson.Field
		OfBool                                                respjson.Field
		OfOpportunityListResponseDataFieldValueArray          respjson.Field
		OfAnyArray                                            respjson.Field
		OfOpportunityListResponseDataFieldValueMapItemMapItem respjson.Field
		raw                                                   string
	} `json:"-"`
}

func (u OpportunityListResponseDataFieldValueUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u OpportunityListResponseDataFieldValueUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u OpportunityListResponseDataFieldValueUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u OpportunityListResponseDataFieldValueUnion) AsOpportunityListResponseDataFieldValueArray() (v []OpportunityListResponseDataFieldValueArrayItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u OpportunityListResponseDataFieldValueUnion) AsOpportunityListResponseDataFieldValueMapMap() (v map[string]OpportunityListResponseDataFieldValueMapItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u OpportunityListResponseDataFieldValueUnion) RawJSON() string { return u.JSON.raw }

func (r *OpportunityListResponseDataFieldValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// OpportunityListResponseDataFieldValueArrayItemUnion contains all possible
// properties and values from [string], [float64], [bool], [[]any],
// [map[string]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray
// OfOpportunityListResponseDataFieldValueArrayItemMapItem]
type OpportunityListResponseDataFieldValueArrayItemUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	// This field will be present if the value is a [any] instead of an object.
	OfOpportunityListResponseDataFieldValueArrayItemMapItem any `json:",inline"`
	JSON                                                    struct {
		OfString                                                respjson.Field
		OfFloat                                                 respjson.Field
		OfBool                                                  respjson.Field
		OfAnyArray                                              respjson.Field
		OfOpportunityListResponseDataFieldValueArrayItemMapItem respjson.Field
		raw                                                     string
	} `json:"-"`
}

func (u OpportunityListResponseDataFieldValueArrayItemUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u OpportunityListResponseDataFieldValueArrayItemUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u OpportunityListResponseDataFieldValueArrayItemUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u OpportunityListResponseDataFieldValueArrayItemUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u OpportunityListResponseDataFieldValueArrayItemUnion) AsAnyMap() (v map[string]any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u OpportunityListResponseDataFieldValueArrayItemUnion) RawJSON() string { return u.JSON.raw }

func (r *OpportunityListResponseDataFieldValueArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// OpportunityListResponseDataFieldValueMapItemUnion contains all possible
// properties and values from [string], [float64], [bool], [[]any],
// [map[string]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray
// OfOpportunityListResponseDataFieldValueMapItemMapItem]
type OpportunityListResponseDataFieldValueMapItemUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	// This field will be present if the value is a [any] instead of an object.
	OfOpportunityListResponseDataFieldValueMapItemMapItem any `json:",inline"`
	JSON                                                  struct {
		OfString                                              respjson.Field
		OfFloat                                               respjson.Field
		OfBool                                                respjson.Field
		OfAnyArray                                            respjson.Field
		OfOpportunityListResponseDataFieldValueMapItemMapItem respjson.Field
		raw                                                   string
	} `json:"-"`
}

func (u OpportunityListResponseDataFieldValueMapItemUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u OpportunityListResponseDataFieldValueMapItemUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u OpportunityListResponseDataFieldValueMapItemUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u OpportunityListResponseDataFieldValueMapItemUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u OpportunityListResponseDataFieldValueMapItemUnion) AsAnyMap() (v map[string]any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u OpportunityListResponseDataFieldValueMapItemUnion) RawJSON() string { return u.JSON.raw }

func (r *OpportunityListResponseDataFieldValueMapItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OpportunityListResponseDataRelationship struct {
	Cardinality string   `json:"cardinality" api:"required"`
	ObjectType  string   `json:"objectType" api:"required"`
	Values      []string `json:"values" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Cardinality respjson.Field
		ObjectType  respjson.Field
		Values      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OpportunityListResponseDataRelationship) RawJSON() string { return r.JSON.raw }
func (r *OpportunityListResponseDataRelationship) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// OpportunityListResponseDataUnion contains all possible properties and values
// from [string], [float64], [bool], [[]OpportunityListResponseDataArrayItemUnion],
// [map[string]OpportunityListResponseDataMapItemUnion].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfOpportunityListResponseDataArray
// OfAnyArray OfOpportunityListResponseDataMapItemMapItem]
type OpportunityListResponseDataUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a
	// [[]OpportunityListResponseDataArrayItemUnion] instead of an object.
	OfOpportunityListResponseDataArray []OpportunityListResponseDataArrayItemUnion `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	// This field will be present if the value is a [any] instead of an object.
	OfOpportunityListResponseDataMapItemMapItem any `json:",inline"`
	JSON                                        struct {
		OfString                                    respjson.Field
		OfFloat                                     respjson.Field
		OfBool                                      respjson.Field
		OfOpportunityListResponseDataArray          respjson.Field
		OfAnyArray                                  respjson.Field
		OfOpportunityListResponseDataMapItemMapItem respjson.Field
		raw                                         string
	} `json:"-"`
}

func (u OpportunityListResponseDataUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u OpportunityListResponseDataUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u OpportunityListResponseDataUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u OpportunityListResponseDataUnion) AsOpportunityListResponseDataArray() (v []OpportunityListResponseDataArrayItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u OpportunityListResponseDataUnion) AsOpportunityListResponseDataMapMap() (v map[string]OpportunityListResponseDataMapItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u OpportunityListResponseDataUnion) RawJSON() string { return u.JSON.raw }

func (r *OpportunityListResponseDataUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// OpportunityListResponseDataArrayItemUnion contains all possible properties and
// values from [string], [float64], [bool], [[]any], [map[string]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray
// OfOpportunityListResponseDataArrayItemMapItem]
type OpportunityListResponseDataArrayItemUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	// This field will be present if the value is a [any] instead of an object.
	OfOpportunityListResponseDataArrayItemMapItem any `json:",inline"`
	JSON                                          struct {
		OfString                                      respjson.Field
		OfFloat                                       respjson.Field
		OfBool                                        respjson.Field
		OfAnyArray                                    respjson.Field
		OfOpportunityListResponseDataArrayItemMapItem respjson.Field
		raw                                           string
	} `json:"-"`
}

func (u OpportunityListResponseDataArrayItemUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u OpportunityListResponseDataArrayItemUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u OpportunityListResponseDataArrayItemUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u OpportunityListResponseDataArrayItemUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u OpportunityListResponseDataArrayItemUnion) AsAnyMap() (v map[string]any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u OpportunityListResponseDataArrayItemUnion) RawJSON() string { return u.JSON.raw }

func (r *OpportunityListResponseDataArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// OpportunityListResponseDataMapItemUnion contains all possible properties and
// values from [string], [float64], [bool], [[]any], [map[string]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray
// OfOpportunityListResponseDataMapItemMapItem]
type OpportunityListResponseDataMapItemUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	// This field will be present if the value is a [any] instead of an object.
	OfOpportunityListResponseDataMapItemMapItem any `json:",inline"`
	JSON                                        struct {
		OfString                                    respjson.Field
		OfFloat                                     respjson.Field
		OfBool                                      respjson.Field
		OfAnyArray                                  respjson.Field
		OfOpportunityListResponseDataMapItemMapItem respjson.Field
		raw                                         string
	} `json:"-"`
}

func (u OpportunityListResponseDataMapItemUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u OpportunityListResponseDataMapItemUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u OpportunityListResponseDataMapItemUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u OpportunityListResponseDataMapItemUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u OpportunityListResponseDataMapItemUnion) AsAnyMap() (v map[string]any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u OpportunityListResponseDataMapItemUnion) RawJSON() string { return u.JSON.raw }

func (r *OpportunityListResponseDataMapItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OpportunityNewParams struct {
	Fields        OpportunityNewParamsFields        `json:"fields,omitzero" api:"required"`
	Relationships OpportunityNewParamsRelationships `json:"relationships,omitzero" api:"required"`
	paramObj
}

func (r OpportunityNewParams) MarshalJSON() (data []byte, err error) {
	type shadow OpportunityNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *OpportunityNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties SystemName, SystemStage are required.
type OpportunityNewParamsFields struct {
	SystemName  string                                    `json:"system_name" api:"required"`
	SystemStage string                                    `json:"system_stage" api:"required"`
	ExtraFields map[string]OpportunityNewParamsFieldUnion `json:"-"`
	paramObj
}

func (r OpportunityNewParamsFields) MarshalJSON() (data []byte, err error) {
	type shadow OpportunityNewParamsFields
	return param.MarshalWithExtras(r, (*shadow)(&r), r.ExtraFields)
}
func (r *OpportunityNewParamsFields) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type OpportunityNewParamsFieldUnion struct {
	OfString                     param.Opt[string]                                `json:",omitzero,inline"`
	OfFloat                      param.Opt[float64]                               `json:",omitzero,inline"`
	OfBool                       param.Opt[bool]                                  `json:",omitzero,inline"`
	OfOpportunityNewsFieldArray  []OpportunityNewParamsFieldArrayItemUnion        `json:",omitzero,inline"`
	OfOpportunityNewsFieldMapMap map[string]OpportunityNewParamsFieldMapItemUnion `json:",omitzero,inline"`
	paramUnion
}

func (u OpportunityNewParamsFieldUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString,
		u.OfFloat,
		u.OfBool,
		u.OfOpportunityNewsFieldArray,
		u.OfOpportunityNewsFieldMapMap)
}
func (u *OpportunityNewParamsFieldUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type OpportunityNewParamsFieldArrayItemUnion struct {
	OfString   param.Opt[string]  `json:",omitzero,inline"`
	OfFloat    param.Opt[float64] `json:",omitzero,inline"`
	OfBool     param.Opt[bool]    `json:",omitzero,inline"`
	OfAnyArray []any              `json:",omitzero,inline"`
	OfAnyMap   map[string]any     `json:",omitzero,inline"`
	paramUnion
}

func (u OpportunityNewParamsFieldArrayItemUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString,
		u.OfFloat,
		u.OfBool,
		u.OfAnyArray,
		u.OfAnyMap)
}
func (u *OpportunityNewParamsFieldArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type OpportunityNewParamsFieldMapItemUnion struct {
	OfString   param.Opt[string]  `json:",omitzero,inline"`
	OfFloat    param.Opt[float64] `json:",omitzero,inline"`
	OfBool     param.Opt[bool]    `json:",omitzero,inline"`
	OfAnyArray []any              `json:",omitzero,inline"`
	OfAnyMap   map[string]any     `json:",omitzero,inline"`
	paramUnion
}

func (u OpportunityNewParamsFieldMapItemUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString,
		u.OfFloat,
		u.OfBool,
		u.OfAnyArray,
		u.OfAnyMap)
}
func (u *OpportunityNewParamsFieldMapItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// The property SystemAccount is required.
type OpportunityNewParamsRelationships struct {
	SystemAccount   OpportunityNewParamsRelationshipsSystemAccountUnion   `json:"system_account,omitzero" api:"required"`
	SystemChampion  OpportunityNewParamsRelationshipsSystemChampionUnion  `json:"system_champion,omitzero"`
	SystemCreatedBy OpportunityNewParamsRelationshipsSystemCreatedByUnion `json:"system_createdBy,omitzero"`
	SystemEvaluator OpportunityNewParamsRelationshipsSystemEvaluatorUnion `json:"system_evaluator,omitzero"`
	SystemOwner     OpportunityNewParamsRelationshipsSystemOwnerUnion     `json:"system_owner,omitzero"`
	ExtraFields     map[string]OpportunityNewParamsRelationshipUnion      `json:"-"`
	paramObj
}

func (r OpportunityNewParamsRelationships) MarshalJSON() (data []byte, err error) {
	type shadow OpportunityNewParamsRelationships
	return param.MarshalWithExtras(r, (*shadow)(&r), r.ExtraFields)
}
func (r *OpportunityNewParamsRelationships) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type OpportunityNewParamsRelationshipsSystemAccountUnion struct {
	OfString      param.Opt[string] `json:",omitzero,inline"`
	OfStringArray []string          `json:",omitzero,inline"`
	paramUnion
}

func (u OpportunityNewParamsRelationshipsSystemAccountUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray)
}
func (u *OpportunityNewParamsRelationshipsSystemAccountUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type OpportunityNewParamsRelationshipsSystemChampionUnion struct {
	OfString      param.Opt[string] `json:",omitzero,inline"`
	OfStringArray []string          `json:",omitzero,inline"`
	paramUnion
}

func (u OpportunityNewParamsRelationshipsSystemChampionUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray)
}
func (u *OpportunityNewParamsRelationshipsSystemChampionUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type OpportunityNewParamsRelationshipsSystemCreatedByUnion struct {
	OfString      param.Opt[string] `json:",omitzero,inline"`
	OfStringArray []string          `json:",omitzero,inline"`
	paramUnion
}

func (u OpportunityNewParamsRelationshipsSystemCreatedByUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray)
}
func (u *OpportunityNewParamsRelationshipsSystemCreatedByUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type OpportunityNewParamsRelationshipsSystemEvaluatorUnion struct {
	OfString      param.Opt[string] `json:",omitzero,inline"`
	OfStringArray []string          `json:",omitzero,inline"`
	paramUnion
}

func (u OpportunityNewParamsRelationshipsSystemEvaluatorUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray)
}
func (u *OpportunityNewParamsRelationshipsSystemEvaluatorUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type OpportunityNewParamsRelationshipsSystemOwnerUnion struct {
	OfString      param.Opt[string] `json:",omitzero,inline"`
	OfStringArray []string          `json:",omitzero,inline"`
	paramUnion
}

func (u OpportunityNewParamsRelationshipsSystemOwnerUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray)
}
func (u *OpportunityNewParamsRelationshipsSystemOwnerUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type OpportunityNewParamsRelationshipUnion struct {
	OfString      param.Opt[string] `json:",omitzero,inline"`
	OfStringArray []string          `json:",omitzero,inline"`
	paramUnion
}

func (u OpportunityNewParamsRelationshipUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray)
}
func (u *OpportunityNewParamsRelationshipUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

type OpportunityUpdateParams struct {
	Fields        OpportunityUpdateParamsFields        `json:"fields,omitzero"`
	Relationships OpportunityUpdateParamsRelationships `json:"relationships,omitzero"`
	paramObj
}

func (r OpportunityUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow OpportunityUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *OpportunityUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OpportunityUpdateParamsFields struct {
	SystemName  param.Opt[string]                            `json:"system_name,omitzero"`
	SystemStage param.Opt[string]                            `json:"system_stage,omitzero"`
	ExtraFields map[string]OpportunityUpdateParamsFieldUnion `json:"-"`
	paramObj
}

func (r OpportunityUpdateParamsFields) MarshalJSON() (data []byte, err error) {
	type shadow OpportunityUpdateParamsFields
	return param.MarshalWithExtras(r, (*shadow)(&r), r.ExtraFields)
}
func (r *OpportunityUpdateParamsFields) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type OpportunityUpdateParamsFieldUnion struct {
	OfString                        param.Opt[string]                                   `json:",omitzero,inline"`
	OfFloat                         param.Opt[float64]                                  `json:",omitzero,inline"`
	OfBool                          param.Opt[bool]                                     `json:",omitzero,inline"`
	OfOpportunityUpdatesFieldArray  []OpportunityUpdateParamsFieldArrayItemUnion        `json:",omitzero,inline"`
	OfOpportunityUpdatesFieldMapMap map[string]OpportunityUpdateParamsFieldMapItemUnion `json:",omitzero,inline"`
	paramUnion
}

func (u OpportunityUpdateParamsFieldUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString,
		u.OfFloat,
		u.OfBool,
		u.OfOpportunityUpdatesFieldArray,
		u.OfOpportunityUpdatesFieldMapMap)
}
func (u *OpportunityUpdateParamsFieldUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type OpportunityUpdateParamsFieldArrayItemUnion struct {
	OfString   param.Opt[string]  `json:",omitzero,inline"`
	OfFloat    param.Opt[float64] `json:",omitzero,inline"`
	OfBool     param.Opt[bool]    `json:",omitzero,inline"`
	OfAnyArray []any              `json:",omitzero,inline"`
	OfAnyMap   map[string]any     `json:",omitzero,inline"`
	paramUnion
}

func (u OpportunityUpdateParamsFieldArrayItemUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString,
		u.OfFloat,
		u.OfBool,
		u.OfAnyArray,
		u.OfAnyMap)
}
func (u *OpportunityUpdateParamsFieldArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type OpportunityUpdateParamsFieldMapItemUnion struct {
	OfString   param.Opt[string]  `json:",omitzero,inline"`
	OfFloat    param.Opt[float64] `json:",omitzero,inline"`
	OfBool     param.Opt[bool]    `json:",omitzero,inline"`
	OfAnyArray []any              `json:",omitzero,inline"`
	OfAnyMap   map[string]any     `json:",omitzero,inline"`
	paramUnion
}

func (u OpportunityUpdateParamsFieldMapItemUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString,
		u.OfFloat,
		u.OfBool,
		u.OfAnyArray,
		u.OfAnyMap)
}
func (u *OpportunityUpdateParamsFieldMapItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

type OpportunityUpdateParamsRelationships struct {
	SystemAccount   OpportunityUpdateParamsRelationshipsSystemAccount   `json:"system_account,omitzero"`
	SystemChampion  OpportunityUpdateParamsRelationshipsSystemChampion  `json:"system_champion,omitzero"`
	SystemCreatedBy OpportunityUpdateParamsRelationshipsSystemCreatedBy `json:"system_createdBy,omitzero"`
	SystemEvaluator OpportunityUpdateParamsRelationshipsSystemEvaluator `json:"system_evaluator,omitzero"`
	SystemOwner     OpportunityUpdateParamsRelationshipsSystemOwner     `json:"system_owner,omitzero"`
	ExtraFields     map[string]OpportunityUpdateParamsRelationship      `json:"-"`
	paramObj
}

func (r OpportunityUpdateParamsRelationships) MarshalJSON() (data []byte, err error) {
	type shadow OpportunityUpdateParamsRelationships
	return param.MarshalWithExtras(r, (*shadow)(&r), r.ExtraFields)
}
func (r *OpportunityUpdateParamsRelationships) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OpportunityUpdateParamsRelationshipsSystemAccount struct {
	Add     OpportunityUpdateParamsRelationshipsSystemAccountAddUnion     `json:"add,omitzero"`
	Remove  OpportunityUpdateParamsRelationshipsSystemAccountRemoveUnion  `json:"remove,omitzero"`
	Replace OpportunityUpdateParamsRelationshipsSystemAccountReplaceUnion `json:"replace,omitzero"`
	paramObj
}

func (r OpportunityUpdateParamsRelationshipsSystemAccount) MarshalJSON() (data []byte, err error) {
	type shadow OpportunityUpdateParamsRelationshipsSystemAccount
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *OpportunityUpdateParamsRelationshipsSystemAccount) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type OpportunityUpdateParamsRelationshipsSystemAccountAddUnion struct {
	OfString      param.Opt[string] `json:",omitzero,inline"`
	OfStringArray []string          `json:",omitzero,inline"`
	paramUnion
}

func (u OpportunityUpdateParamsRelationshipsSystemAccountAddUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray)
}
func (u *OpportunityUpdateParamsRelationshipsSystemAccountAddUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type OpportunityUpdateParamsRelationshipsSystemAccountRemoveUnion struct {
	OfString      param.Opt[string] `json:",omitzero,inline"`
	OfStringArray []string          `json:",omitzero,inline"`
	paramUnion
}

func (u OpportunityUpdateParamsRelationshipsSystemAccountRemoveUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray)
}
func (u *OpportunityUpdateParamsRelationshipsSystemAccountRemoveUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type OpportunityUpdateParamsRelationshipsSystemAccountReplaceUnion struct {
	OfString      param.Opt[string] `json:",omitzero,inline"`
	OfStringArray []string          `json:",omitzero,inline"`
	paramUnion
}

func (u OpportunityUpdateParamsRelationshipsSystemAccountReplaceUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray)
}
func (u *OpportunityUpdateParamsRelationshipsSystemAccountReplaceUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

type OpportunityUpdateParamsRelationshipsSystemChampion struct {
	Add     OpportunityUpdateParamsRelationshipsSystemChampionAddUnion     `json:"add,omitzero"`
	Remove  OpportunityUpdateParamsRelationshipsSystemChampionRemoveUnion  `json:"remove,omitzero"`
	Replace OpportunityUpdateParamsRelationshipsSystemChampionReplaceUnion `json:"replace,omitzero"`
	paramObj
}

func (r OpportunityUpdateParamsRelationshipsSystemChampion) MarshalJSON() (data []byte, err error) {
	type shadow OpportunityUpdateParamsRelationshipsSystemChampion
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *OpportunityUpdateParamsRelationshipsSystemChampion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type OpportunityUpdateParamsRelationshipsSystemChampionAddUnion struct {
	OfString      param.Opt[string] `json:",omitzero,inline"`
	OfStringArray []string          `json:",omitzero,inline"`
	paramUnion
}

func (u OpportunityUpdateParamsRelationshipsSystemChampionAddUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray)
}
func (u *OpportunityUpdateParamsRelationshipsSystemChampionAddUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type OpportunityUpdateParamsRelationshipsSystemChampionRemoveUnion struct {
	OfString      param.Opt[string] `json:",omitzero,inline"`
	OfStringArray []string          `json:",omitzero,inline"`
	paramUnion
}

func (u OpportunityUpdateParamsRelationshipsSystemChampionRemoveUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray)
}
func (u *OpportunityUpdateParamsRelationshipsSystemChampionRemoveUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type OpportunityUpdateParamsRelationshipsSystemChampionReplaceUnion struct {
	OfString      param.Opt[string] `json:",omitzero,inline"`
	OfStringArray []string          `json:",omitzero,inline"`
	paramUnion
}

func (u OpportunityUpdateParamsRelationshipsSystemChampionReplaceUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray)
}
func (u *OpportunityUpdateParamsRelationshipsSystemChampionReplaceUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

type OpportunityUpdateParamsRelationshipsSystemCreatedBy struct {
	Add     OpportunityUpdateParamsRelationshipsSystemCreatedByAddUnion     `json:"add,omitzero"`
	Remove  OpportunityUpdateParamsRelationshipsSystemCreatedByRemoveUnion  `json:"remove,omitzero"`
	Replace OpportunityUpdateParamsRelationshipsSystemCreatedByReplaceUnion `json:"replace,omitzero"`
	paramObj
}

func (r OpportunityUpdateParamsRelationshipsSystemCreatedBy) MarshalJSON() (data []byte, err error) {
	type shadow OpportunityUpdateParamsRelationshipsSystemCreatedBy
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *OpportunityUpdateParamsRelationshipsSystemCreatedBy) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type OpportunityUpdateParamsRelationshipsSystemCreatedByAddUnion struct {
	OfString      param.Opt[string] `json:",omitzero,inline"`
	OfStringArray []string          `json:",omitzero,inline"`
	paramUnion
}

func (u OpportunityUpdateParamsRelationshipsSystemCreatedByAddUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray)
}
func (u *OpportunityUpdateParamsRelationshipsSystemCreatedByAddUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type OpportunityUpdateParamsRelationshipsSystemCreatedByRemoveUnion struct {
	OfString      param.Opt[string] `json:",omitzero,inline"`
	OfStringArray []string          `json:",omitzero,inline"`
	paramUnion
}

func (u OpportunityUpdateParamsRelationshipsSystemCreatedByRemoveUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray)
}
func (u *OpportunityUpdateParamsRelationshipsSystemCreatedByRemoveUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type OpportunityUpdateParamsRelationshipsSystemCreatedByReplaceUnion struct {
	OfString      param.Opt[string] `json:",omitzero,inline"`
	OfStringArray []string          `json:",omitzero,inline"`
	paramUnion
}

func (u OpportunityUpdateParamsRelationshipsSystemCreatedByReplaceUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray)
}
func (u *OpportunityUpdateParamsRelationshipsSystemCreatedByReplaceUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

type OpportunityUpdateParamsRelationshipsSystemEvaluator struct {
	Add     OpportunityUpdateParamsRelationshipsSystemEvaluatorAddUnion     `json:"add,omitzero"`
	Remove  OpportunityUpdateParamsRelationshipsSystemEvaluatorRemoveUnion  `json:"remove,omitzero"`
	Replace OpportunityUpdateParamsRelationshipsSystemEvaluatorReplaceUnion `json:"replace,omitzero"`
	paramObj
}

func (r OpportunityUpdateParamsRelationshipsSystemEvaluator) MarshalJSON() (data []byte, err error) {
	type shadow OpportunityUpdateParamsRelationshipsSystemEvaluator
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *OpportunityUpdateParamsRelationshipsSystemEvaluator) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type OpportunityUpdateParamsRelationshipsSystemEvaluatorAddUnion struct {
	OfString      param.Opt[string] `json:",omitzero,inline"`
	OfStringArray []string          `json:",omitzero,inline"`
	paramUnion
}

func (u OpportunityUpdateParamsRelationshipsSystemEvaluatorAddUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray)
}
func (u *OpportunityUpdateParamsRelationshipsSystemEvaluatorAddUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type OpportunityUpdateParamsRelationshipsSystemEvaluatorRemoveUnion struct {
	OfString      param.Opt[string] `json:",omitzero,inline"`
	OfStringArray []string          `json:",omitzero,inline"`
	paramUnion
}

func (u OpportunityUpdateParamsRelationshipsSystemEvaluatorRemoveUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray)
}
func (u *OpportunityUpdateParamsRelationshipsSystemEvaluatorRemoveUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type OpportunityUpdateParamsRelationshipsSystemEvaluatorReplaceUnion struct {
	OfString      param.Opt[string] `json:",omitzero,inline"`
	OfStringArray []string          `json:",omitzero,inline"`
	paramUnion
}

func (u OpportunityUpdateParamsRelationshipsSystemEvaluatorReplaceUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray)
}
func (u *OpportunityUpdateParamsRelationshipsSystemEvaluatorReplaceUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

type OpportunityUpdateParamsRelationshipsSystemOwner struct {
	Add     OpportunityUpdateParamsRelationshipsSystemOwnerAddUnion     `json:"add,omitzero"`
	Remove  OpportunityUpdateParamsRelationshipsSystemOwnerRemoveUnion  `json:"remove,omitzero"`
	Replace OpportunityUpdateParamsRelationshipsSystemOwnerReplaceUnion `json:"replace,omitzero"`
	paramObj
}

func (r OpportunityUpdateParamsRelationshipsSystemOwner) MarshalJSON() (data []byte, err error) {
	type shadow OpportunityUpdateParamsRelationshipsSystemOwner
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *OpportunityUpdateParamsRelationshipsSystemOwner) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type OpportunityUpdateParamsRelationshipsSystemOwnerAddUnion struct {
	OfString      param.Opt[string] `json:",omitzero,inline"`
	OfStringArray []string          `json:",omitzero,inline"`
	paramUnion
}

func (u OpportunityUpdateParamsRelationshipsSystemOwnerAddUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray)
}
func (u *OpportunityUpdateParamsRelationshipsSystemOwnerAddUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type OpportunityUpdateParamsRelationshipsSystemOwnerRemoveUnion struct {
	OfString      param.Opt[string] `json:",omitzero,inline"`
	OfStringArray []string          `json:",omitzero,inline"`
	paramUnion
}

func (u OpportunityUpdateParamsRelationshipsSystemOwnerRemoveUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray)
}
func (u *OpportunityUpdateParamsRelationshipsSystemOwnerRemoveUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type OpportunityUpdateParamsRelationshipsSystemOwnerReplaceUnion struct {
	OfString      param.Opt[string] `json:",omitzero,inline"`
	OfStringArray []string          `json:",omitzero,inline"`
	paramUnion
}

func (u OpportunityUpdateParamsRelationshipsSystemOwnerReplaceUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray)
}
func (u *OpportunityUpdateParamsRelationshipsSystemOwnerReplaceUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

type OpportunityUpdateParamsRelationship struct {
	Add     OpportunityUpdateParamsRelationshipAddUnion     `json:"add,omitzero"`
	Remove  OpportunityUpdateParamsRelationshipRemoveUnion  `json:"remove,omitzero"`
	Replace OpportunityUpdateParamsRelationshipReplaceUnion `json:"replace,omitzero"`
	paramObj
}

func (r OpportunityUpdateParamsRelationship) MarshalJSON() (data []byte, err error) {
	type shadow OpportunityUpdateParamsRelationship
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *OpportunityUpdateParamsRelationship) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type OpportunityUpdateParamsRelationshipAddUnion struct {
	OfString      param.Opt[string] `json:",omitzero,inline"`
	OfStringArray []string          `json:",omitzero,inline"`
	paramUnion
}

func (u OpportunityUpdateParamsRelationshipAddUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray)
}
func (u *OpportunityUpdateParamsRelationshipAddUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type OpportunityUpdateParamsRelationshipRemoveUnion struct {
	OfString      param.Opt[string] `json:",omitzero,inline"`
	OfStringArray []string          `json:",omitzero,inline"`
	paramUnion
}

func (u OpportunityUpdateParamsRelationshipRemoveUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray)
}
func (u *OpportunityUpdateParamsRelationshipRemoveUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type OpportunityUpdateParamsRelationshipReplaceUnion struct {
	OfString      param.Opt[string] `json:",omitzero,inline"`
	OfStringArray []string          `json:",omitzero,inline"`
	paramUnion
}

func (u OpportunityUpdateParamsRelationshipReplaceUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray)
}
func (u *OpportunityUpdateParamsRelationshipReplaceUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

type OpportunityListParams struct {
	Limit  param.Opt[int64] `query:"limit,omitzero" json:"-"`
	Offset param.Opt[int64] `query:"offset,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [OpportunityListParams]'s query parameters as `url.Values`.
func (r OpportunityListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
