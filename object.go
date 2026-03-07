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

	"github.com/stainless-sdks/lightfield-go/internal/apijson"
	"github.com/stainless-sdks/lightfield-go/internal/apiquery"
	"github.com/stainless-sdks/lightfield-go/internal/requestconfig"
	"github.com/stainless-sdks/lightfield-go/option"
	"github.com/stainless-sdks/lightfield-go/packages/param"
	"github.com/stainless-sdks/lightfield-go/packages/respjson"
)

// ObjectService contains methods and other services that help with interacting
// with the Lightfield API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewObjectService] method instead.
type ObjectService struct {
	Options []option.RequestOption
}

// NewObjectService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewObjectService(opts ...option.RequestOption) (r ObjectService) {
	r = ObjectService{}
	r.Options = opts
	return
}

func (r *ObjectService) New(ctx context.Context, entityType ObjectNewParamsEntityType, body ObjectNewParams, opts ...option.RequestOption) (res *ObjectNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("v1/objects/%v", entityType)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

func (r *ObjectService) Get(ctx context.Context, id string, query ObjectGetParams, opts ...option.RequestOption) (res *ObjectGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("v1/objects/%v/%s", query.EntityType, url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

func (r *ObjectService) Update(ctx context.Context, id string, params ObjectUpdateParams, opts ...option.RequestOption) (res *ObjectUpdateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("v1/objects/%v/%s", params.EntityType, url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return
}

func (r *ObjectService) List(ctx context.Context, entityType ObjectListParamsEntityType, query ObjectListParams, opts ...option.RequestOption) (res *ObjectListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("v1/objects/%v", entityType)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type ObjectNewResponse = any

type ObjectGetResponse struct {
	ID                  string                                         `json:"id" api:"required"`
	CreatedAt           string                                         `json:"createdAt" api:"required"`
	CustomFields        map[string]ObjectGetResponseCustomField        `json:"customFields" api:"required"`
	CustomRelationships map[string]ObjectGetResponseCustomRelationship `json:"customRelationships" api:"required"`
	Object              string                                         `json:"object" api:"required"`
	UpdatedAt           string                                         `json:"updatedAt" api:"required"`
	ExtraFields         map[string]ObjectGetResponseUnion              `json:"" api:"extrafields"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                  respjson.Field
		CreatedAt           respjson.Field
		CustomFields        respjson.Field
		CustomRelationships respjson.Field
		Object              respjson.Field
		UpdatedAt           respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ObjectGetResponse) RawJSON() string { return r.JSON.raw }
func (r *ObjectGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ObjectGetResponseCustomField struct {
	Value ObjectGetResponseCustomFieldValueUnion `json:"value" api:"required"`
	// Any of "ADDRESS", "CHECKBOX", "CURRENCY", "DATETIME", "EMAIL", "FULL_NAME",
	// "MULTI_SELECT", "NUMBER", "SINGLE_SELECT", "SOCIAL_HANDLE", "TELEPHONE", "TEXT",
	// "TOME_MARKDOWN", "URL".
	ValueType string `json:"valueType" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Value       respjson.Field
		ValueType   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ObjectGetResponseCustomField) RawJSON() string { return r.JSON.raw }
func (r *ObjectGetResponseCustomField) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ObjectGetResponseCustomFieldValueUnion contains all possible properties and
// values from [string], [float64], [bool],
// [[]ObjectGetResponseCustomFieldValueArrayItemUnion],
// [map[string]ObjectGetResponseCustomFieldValueMapItemUnion].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfObjectGetResponseCustomFieldValueArray
// OfAnyArray OfObjectGetResponseCustomFieldValueMapItemMapItem]
type ObjectGetResponseCustomFieldValueUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a
	// [[]ObjectGetResponseCustomFieldValueArrayItemUnion] instead of an object.
	OfObjectGetResponseCustomFieldValueArray []ObjectGetResponseCustomFieldValueArrayItemUnion `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	// This field will be present if the value is a [any] instead of an object.
	OfObjectGetResponseCustomFieldValueMapItemMapItem any `json:",inline"`
	JSON                                              struct {
		OfString                                          respjson.Field
		OfFloat                                           respjson.Field
		OfBool                                            respjson.Field
		OfObjectGetResponseCustomFieldValueArray          respjson.Field
		OfAnyArray                                        respjson.Field
		OfObjectGetResponseCustomFieldValueMapItemMapItem respjson.Field
		raw                                               string
	} `json:"-"`
}

func (u ObjectGetResponseCustomFieldValueUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ObjectGetResponseCustomFieldValueUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ObjectGetResponseCustomFieldValueUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ObjectGetResponseCustomFieldValueUnion) AsObjectGetResponseCustomFieldValueArray() (v []ObjectGetResponseCustomFieldValueArrayItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ObjectGetResponseCustomFieldValueUnion) AsObjectGetResponseCustomFieldValueMapMap() (v map[string]ObjectGetResponseCustomFieldValueMapItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ObjectGetResponseCustomFieldValueUnion) RawJSON() string { return u.JSON.raw }

func (r *ObjectGetResponseCustomFieldValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ObjectGetResponseCustomFieldValueArrayItemUnion contains all possible properties
// and values from [string], [float64], [bool], [[]any], [map[string]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray
// OfObjectGetResponseCustomFieldValueArrayItemMapItem]
type ObjectGetResponseCustomFieldValueArrayItemUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	// This field will be present if the value is a [any] instead of an object.
	OfObjectGetResponseCustomFieldValueArrayItemMapItem any `json:",inline"`
	JSON                                                struct {
		OfString                                            respjson.Field
		OfFloat                                             respjson.Field
		OfBool                                              respjson.Field
		OfAnyArray                                          respjson.Field
		OfObjectGetResponseCustomFieldValueArrayItemMapItem respjson.Field
		raw                                                 string
	} `json:"-"`
}

func (u ObjectGetResponseCustomFieldValueArrayItemUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ObjectGetResponseCustomFieldValueArrayItemUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ObjectGetResponseCustomFieldValueArrayItemUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ObjectGetResponseCustomFieldValueArrayItemUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ObjectGetResponseCustomFieldValueArrayItemUnion) AsAnyMap() (v map[string]any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ObjectGetResponseCustomFieldValueArrayItemUnion) RawJSON() string { return u.JSON.raw }

func (r *ObjectGetResponseCustomFieldValueArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ObjectGetResponseCustomFieldValueMapItemUnion contains all possible properties
// and values from [string], [float64], [bool], [[]any], [map[string]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray
// OfObjectGetResponseCustomFieldValueMapItemMapItem]
type ObjectGetResponseCustomFieldValueMapItemUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	// This field will be present if the value is a [any] instead of an object.
	OfObjectGetResponseCustomFieldValueMapItemMapItem any `json:",inline"`
	JSON                                              struct {
		OfString                                          respjson.Field
		OfFloat                                           respjson.Field
		OfBool                                            respjson.Field
		OfAnyArray                                        respjson.Field
		OfObjectGetResponseCustomFieldValueMapItemMapItem respjson.Field
		raw                                               string
	} `json:"-"`
}

func (u ObjectGetResponseCustomFieldValueMapItemUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ObjectGetResponseCustomFieldValueMapItemUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ObjectGetResponseCustomFieldValueMapItemUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ObjectGetResponseCustomFieldValueMapItemUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ObjectGetResponseCustomFieldValueMapItemUnion) AsAnyMap() (v map[string]any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ObjectGetResponseCustomFieldValueMapItemUnion) RawJSON() string { return u.JSON.raw }

func (r *ObjectGetResponseCustomFieldValueMapItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ObjectGetResponseCustomRelationship struct {
	Entity           string   `json:"entity" api:"required"`
	RelationshipType string   `json:"relationshipType" api:"required"`
	Values           []string `json:"values" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Entity           respjson.Field
		RelationshipType respjson.Field
		Values           respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ObjectGetResponseCustomRelationship) RawJSON() string { return r.JSON.raw }
func (r *ObjectGetResponseCustomRelationship) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ObjectGetResponseUnion contains all possible properties and values from
// [string], [float64], [bool], [[]ObjectGetResponseArrayItemUnion],
// [map[string]ObjectGetResponseMapItemUnion].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfObjectGetResponseArray OfAnyArray
// OfObjectGetResponseMapItemMapItem]
type ObjectGetResponseUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]ObjectGetResponseArrayItemUnion]
	// instead of an object.
	OfObjectGetResponseArray []ObjectGetResponseArrayItemUnion `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	// This field will be present if the value is a [any] instead of an object.
	OfObjectGetResponseMapItemMapItem any `json:",inline"`
	JSON                              struct {
		OfString                          respjson.Field
		OfFloat                           respjson.Field
		OfBool                            respjson.Field
		OfObjectGetResponseArray          respjson.Field
		OfAnyArray                        respjson.Field
		OfObjectGetResponseMapItemMapItem respjson.Field
		raw                               string
	} `json:"-"`
}

func (u ObjectGetResponseUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ObjectGetResponseUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ObjectGetResponseUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ObjectGetResponseUnion) AsObjectGetResponseArray() (v []ObjectGetResponseArrayItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ObjectGetResponseUnion) AsObjectGetResponseMapMap() (v map[string]ObjectGetResponseMapItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ObjectGetResponseUnion) RawJSON() string { return u.JSON.raw }

func (r *ObjectGetResponseUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ObjectGetResponseArrayItemUnion contains all possible properties and values from
// [string], [float64], [bool], [[]any], [map[string]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray
// OfObjectGetResponseArrayItemMapItem]
type ObjectGetResponseArrayItemUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	// This field will be present if the value is a [any] instead of an object.
	OfObjectGetResponseArrayItemMapItem any `json:",inline"`
	JSON                                struct {
		OfString                            respjson.Field
		OfFloat                             respjson.Field
		OfBool                              respjson.Field
		OfAnyArray                          respjson.Field
		OfObjectGetResponseArrayItemMapItem respjson.Field
		raw                                 string
	} `json:"-"`
}

func (u ObjectGetResponseArrayItemUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ObjectGetResponseArrayItemUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ObjectGetResponseArrayItemUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ObjectGetResponseArrayItemUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ObjectGetResponseArrayItemUnion) AsAnyMap() (v map[string]any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ObjectGetResponseArrayItemUnion) RawJSON() string { return u.JSON.raw }

func (r *ObjectGetResponseArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ObjectGetResponseMapItemUnion contains all possible properties and values from
// [string], [float64], [bool], [[]any], [map[string]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray
// OfObjectGetResponseMapItemMapItem]
type ObjectGetResponseMapItemUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	// This field will be present if the value is a [any] instead of an object.
	OfObjectGetResponseMapItemMapItem any `json:",inline"`
	JSON                              struct {
		OfString                          respjson.Field
		OfFloat                           respjson.Field
		OfBool                            respjson.Field
		OfAnyArray                        respjson.Field
		OfObjectGetResponseMapItemMapItem respjson.Field
		raw                               string
	} `json:"-"`
}

func (u ObjectGetResponseMapItemUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ObjectGetResponseMapItemUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ObjectGetResponseMapItemUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ObjectGetResponseMapItemUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ObjectGetResponseMapItemUnion) AsAnyMap() (v map[string]any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ObjectGetResponseMapItemUnion) RawJSON() string { return u.JSON.raw }

func (r *ObjectGetResponseMapItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ObjectUpdateResponse = any

type ObjectListResponse struct {
	Data       []ObjectListResponseData `json:"data" api:"required"`
	Object     string                   `json:"object" api:"required"`
	TotalCount int64                    `json:"total_count" api:"required"`
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
func (r ObjectListResponse) RawJSON() string { return r.JSON.raw }
func (r *ObjectListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ObjectListResponseData struct {
	ID                  string                                              `json:"id" api:"required"`
	CreatedAt           string                                              `json:"createdAt" api:"required"`
	CustomFields        map[string]ObjectListResponseDataCustomField        `json:"customFields" api:"required"`
	CustomRelationships map[string]ObjectListResponseDataCustomRelationship `json:"customRelationships" api:"required"`
	Object              string                                              `json:"object" api:"required"`
	UpdatedAt           string                                              `json:"updatedAt" api:"required"`
	ExtraFields         map[string]ObjectListResponseDataUnion              `json:"" api:"extrafields"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                  respjson.Field
		CreatedAt           respjson.Field
		CustomFields        respjson.Field
		CustomRelationships respjson.Field
		Object              respjson.Field
		UpdatedAt           respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ObjectListResponseData) RawJSON() string { return r.JSON.raw }
func (r *ObjectListResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ObjectListResponseDataCustomField struct {
	Value ObjectListResponseDataCustomFieldValueUnion `json:"value" api:"required"`
	// Any of "ADDRESS", "CHECKBOX", "CURRENCY", "DATETIME", "EMAIL", "FULL_NAME",
	// "MULTI_SELECT", "NUMBER", "SINGLE_SELECT", "SOCIAL_HANDLE", "TELEPHONE", "TEXT",
	// "TOME_MARKDOWN", "URL".
	ValueType string `json:"valueType" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Value       respjson.Field
		ValueType   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ObjectListResponseDataCustomField) RawJSON() string { return r.JSON.raw }
func (r *ObjectListResponseDataCustomField) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ObjectListResponseDataCustomFieldValueUnion contains all possible properties and
// values from [string], [float64], [bool],
// [[]ObjectListResponseDataCustomFieldValueArrayItemUnion],
// [map[string]ObjectListResponseDataCustomFieldValueMapItemUnion].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool
// OfObjectListResponseDataCustomFieldValueArray OfAnyArray
// OfObjectListResponseDataCustomFieldValueMapItemMapItem]
type ObjectListResponseDataCustomFieldValueUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a
	// [[]ObjectListResponseDataCustomFieldValueArrayItemUnion] instead of an object.
	OfObjectListResponseDataCustomFieldValueArray []ObjectListResponseDataCustomFieldValueArrayItemUnion `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	// This field will be present if the value is a [any] instead of an object.
	OfObjectListResponseDataCustomFieldValueMapItemMapItem any `json:",inline"`
	JSON                                                   struct {
		OfString                                               respjson.Field
		OfFloat                                                respjson.Field
		OfBool                                                 respjson.Field
		OfObjectListResponseDataCustomFieldValueArray          respjson.Field
		OfAnyArray                                             respjson.Field
		OfObjectListResponseDataCustomFieldValueMapItemMapItem respjson.Field
		raw                                                    string
	} `json:"-"`
}

func (u ObjectListResponseDataCustomFieldValueUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ObjectListResponseDataCustomFieldValueUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ObjectListResponseDataCustomFieldValueUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ObjectListResponseDataCustomFieldValueUnion) AsObjectListResponseDataCustomFieldValueArray() (v []ObjectListResponseDataCustomFieldValueArrayItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ObjectListResponseDataCustomFieldValueUnion) AsObjectListResponseDataCustomFieldValueMapMap() (v map[string]ObjectListResponseDataCustomFieldValueMapItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ObjectListResponseDataCustomFieldValueUnion) RawJSON() string { return u.JSON.raw }

func (r *ObjectListResponseDataCustomFieldValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ObjectListResponseDataCustomFieldValueArrayItemUnion contains all possible
// properties and values from [string], [float64], [bool], [[]any],
// [map[string]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray
// OfObjectListResponseDataCustomFieldValueArrayItemMapItem]
type ObjectListResponseDataCustomFieldValueArrayItemUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	// This field will be present if the value is a [any] instead of an object.
	OfObjectListResponseDataCustomFieldValueArrayItemMapItem any `json:",inline"`
	JSON                                                     struct {
		OfString                                                 respjson.Field
		OfFloat                                                  respjson.Field
		OfBool                                                   respjson.Field
		OfAnyArray                                               respjson.Field
		OfObjectListResponseDataCustomFieldValueArrayItemMapItem respjson.Field
		raw                                                      string
	} `json:"-"`
}

func (u ObjectListResponseDataCustomFieldValueArrayItemUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ObjectListResponseDataCustomFieldValueArrayItemUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ObjectListResponseDataCustomFieldValueArrayItemUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ObjectListResponseDataCustomFieldValueArrayItemUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ObjectListResponseDataCustomFieldValueArrayItemUnion) AsAnyMap() (v map[string]any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ObjectListResponseDataCustomFieldValueArrayItemUnion) RawJSON() string { return u.JSON.raw }

func (r *ObjectListResponseDataCustomFieldValueArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ObjectListResponseDataCustomFieldValueMapItemUnion contains all possible
// properties and values from [string], [float64], [bool], [[]any],
// [map[string]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray
// OfObjectListResponseDataCustomFieldValueMapItemMapItem]
type ObjectListResponseDataCustomFieldValueMapItemUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	// This field will be present if the value is a [any] instead of an object.
	OfObjectListResponseDataCustomFieldValueMapItemMapItem any `json:",inline"`
	JSON                                                   struct {
		OfString                                               respjson.Field
		OfFloat                                                respjson.Field
		OfBool                                                 respjson.Field
		OfAnyArray                                             respjson.Field
		OfObjectListResponseDataCustomFieldValueMapItemMapItem respjson.Field
		raw                                                    string
	} `json:"-"`
}

func (u ObjectListResponseDataCustomFieldValueMapItemUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ObjectListResponseDataCustomFieldValueMapItemUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ObjectListResponseDataCustomFieldValueMapItemUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ObjectListResponseDataCustomFieldValueMapItemUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ObjectListResponseDataCustomFieldValueMapItemUnion) AsAnyMap() (v map[string]any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ObjectListResponseDataCustomFieldValueMapItemUnion) RawJSON() string { return u.JSON.raw }

func (r *ObjectListResponseDataCustomFieldValueMapItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ObjectListResponseDataCustomRelationship struct {
	Entity           string   `json:"entity" api:"required"`
	RelationshipType string   `json:"relationshipType" api:"required"`
	Values           []string `json:"values" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Entity           respjson.Field
		RelationshipType respjson.Field
		Values           respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ObjectListResponseDataCustomRelationship) RawJSON() string { return r.JSON.raw }
func (r *ObjectListResponseDataCustomRelationship) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ObjectListResponseDataUnion contains all possible properties and values from
// [string], [float64], [bool], [[]ObjectListResponseDataArrayItemUnion],
// [map[string]ObjectListResponseDataMapItemUnion].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfObjectListResponseDataArray OfAnyArray
// OfObjectListResponseDataMapItemMapItem]
type ObjectListResponseDataUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a
	// [[]ObjectListResponseDataArrayItemUnion] instead of an object.
	OfObjectListResponseDataArray []ObjectListResponseDataArrayItemUnion `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	// This field will be present if the value is a [any] instead of an object.
	OfObjectListResponseDataMapItemMapItem any `json:",inline"`
	JSON                                   struct {
		OfString                               respjson.Field
		OfFloat                                respjson.Field
		OfBool                                 respjson.Field
		OfObjectListResponseDataArray          respjson.Field
		OfAnyArray                             respjson.Field
		OfObjectListResponseDataMapItemMapItem respjson.Field
		raw                                    string
	} `json:"-"`
}

func (u ObjectListResponseDataUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ObjectListResponseDataUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ObjectListResponseDataUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ObjectListResponseDataUnion) AsObjectListResponseDataArray() (v []ObjectListResponseDataArrayItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ObjectListResponseDataUnion) AsObjectListResponseDataMapMap() (v map[string]ObjectListResponseDataMapItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ObjectListResponseDataUnion) RawJSON() string { return u.JSON.raw }

func (r *ObjectListResponseDataUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ObjectListResponseDataArrayItemUnion contains all possible properties and values
// from [string], [float64], [bool], [[]any], [map[string]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray
// OfObjectListResponseDataArrayItemMapItem]
type ObjectListResponseDataArrayItemUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	// This field will be present if the value is a [any] instead of an object.
	OfObjectListResponseDataArrayItemMapItem any `json:",inline"`
	JSON                                     struct {
		OfString                                 respjson.Field
		OfFloat                                  respjson.Field
		OfBool                                   respjson.Field
		OfAnyArray                               respjson.Field
		OfObjectListResponseDataArrayItemMapItem respjson.Field
		raw                                      string
	} `json:"-"`
}

func (u ObjectListResponseDataArrayItemUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ObjectListResponseDataArrayItemUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ObjectListResponseDataArrayItemUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ObjectListResponseDataArrayItemUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ObjectListResponseDataArrayItemUnion) AsAnyMap() (v map[string]any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ObjectListResponseDataArrayItemUnion) RawJSON() string { return u.JSON.raw }

func (r *ObjectListResponseDataArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ObjectListResponseDataMapItemUnion contains all possible properties and values
// from [string], [float64], [bool], [[]any], [map[string]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray
// OfObjectListResponseDataMapItemMapItem]
type ObjectListResponseDataMapItemUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	// This field will be present if the value is a [any] instead of an object.
	OfObjectListResponseDataMapItemMapItem any `json:",inline"`
	JSON                                   struct {
		OfString                               respjson.Field
		OfFloat                                respjson.Field
		OfBool                                 respjson.Field
		OfAnyArray                             respjson.Field
		OfObjectListResponseDataMapItemMapItem respjson.Field
		raw                                    string
	} `json:"-"`
}

func (u ObjectListResponseDataMapItemUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ObjectListResponseDataMapItemUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ObjectListResponseDataMapItemUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ObjectListResponseDataMapItemUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ObjectListResponseDataMapItemUnion) AsAnyMap() (v map[string]any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ObjectListResponseDataMapItemUnion) RawJSON() string { return u.JSON.raw }

func (r *ObjectListResponseDataMapItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ObjectNewParams struct {
	Fields        map[string]ObjectNewParamsFieldUnion        `json:"fields,omitzero" api:"required"`
	Relationships map[string]ObjectNewParamsRelationshipUnion `json:"relationships,omitzero"`
	paramObj
}

func (r ObjectNewParams) MarshalJSON() (data []byte, err error) {
	type shadow ObjectNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ObjectNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ObjectNewParamsEntityType string

const (
	ObjectNewParamsEntityTypeAccounts      ObjectNewParamsEntityType = "accounts"
	ObjectNewParamsEntityTypeContacts      ObjectNewParamsEntityType = "contacts"
	ObjectNewParamsEntityTypeOpportunities ObjectNewParamsEntityType = "opportunities"
	ObjectNewParamsEntityTypeTasks         ObjectNewParamsEntityType = "tasks"
	ObjectNewParamsEntityTypeNotes         ObjectNewParamsEntityType = "notes"
	ObjectNewParamsEntityTypeMeetings      ObjectNewParamsEntityType = "meetings"
)

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ObjectNewParamsFieldUnion struct {
	OfString                param.Opt[string]                           `json:",omitzero,inline"`
	OfFloat                 param.Opt[float64]                          `json:",omitzero,inline"`
	OfBool                  param.Opt[bool]                             `json:",omitzero,inline"`
	OfObjectNewsFieldArray  []ObjectNewParamsFieldArrayItemUnion        `json:",omitzero,inline"`
	OfObjectNewsFieldMapMap map[string]ObjectNewParamsFieldMapItemUnion `json:",omitzero,inline"`
	paramUnion
}

func (u ObjectNewParamsFieldUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString,
		u.OfFloat,
		u.OfBool,
		u.OfObjectNewsFieldArray,
		u.OfObjectNewsFieldMapMap)
}
func (u *ObjectNewParamsFieldUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ObjectNewParamsFieldArrayItemUnion struct {
	OfString   param.Opt[string]  `json:",omitzero,inline"`
	OfFloat    param.Opt[float64] `json:",omitzero,inline"`
	OfBool     param.Opt[bool]    `json:",omitzero,inline"`
	OfAnyArray []any              `json:",omitzero,inline"`
	OfAnyMap   map[string]any     `json:",omitzero,inline"`
	paramUnion
}

func (u ObjectNewParamsFieldArrayItemUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString,
		u.OfFloat,
		u.OfBool,
		u.OfAnyArray,
		u.OfAnyMap)
}
func (u *ObjectNewParamsFieldArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ObjectNewParamsFieldMapItemUnion struct {
	OfString   param.Opt[string]  `json:",omitzero,inline"`
	OfFloat    param.Opt[float64] `json:",omitzero,inline"`
	OfBool     param.Opt[bool]    `json:",omitzero,inline"`
	OfAnyArray []any              `json:",omitzero,inline"`
	OfAnyMap   map[string]any     `json:",omitzero,inline"`
	paramUnion
}

func (u ObjectNewParamsFieldMapItemUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString,
		u.OfFloat,
		u.OfBool,
		u.OfAnyArray,
		u.OfAnyMap)
}
func (u *ObjectNewParamsFieldMapItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ObjectNewParamsRelationshipUnion struct {
	OfString      param.Opt[string] `json:",omitzero,inline"`
	OfStringArray []string          `json:",omitzero,inline"`
	paramUnion
}

func (u ObjectNewParamsRelationshipUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray)
}
func (u *ObjectNewParamsRelationshipUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

type ObjectGetParams struct {
	// Any of "accounts", "contacts", "opportunities", "tasks", "notes", "meetings".
	EntityType ObjectGetParamsEntityType `path:"entityType,omitzero" api:"required" json:"-"`
	paramObj
}

type ObjectGetParamsEntityType string

const (
	ObjectGetParamsEntityTypeAccounts      ObjectGetParamsEntityType = "accounts"
	ObjectGetParamsEntityTypeContacts      ObjectGetParamsEntityType = "contacts"
	ObjectGetParamsEntityTypeOpportunities ObjectGetParamsEntityType = "opportunities"
	ObjectGetParamsEntityTypeTasks         ObjectGetParamsEntityType = "tasks"
	ObjectGetParamsEntityTypeNotes         ObjectGetParamsEntityType = "notes"
	ObjectGetParamsEntityTypeMeetings      ObjectGetParamsEntityType = "meetings"
)

type ObjectUpdateParams struct {
	// Any of "accounts", "contacts", "opportunities", "tasks", "notes", "meetings".
	EntityType    ObjectUpdateParamsEntityType              `path:"entityType,omitzero" api:"required" json:"-"`
	Fields        map[string]ObjectUpdateParamsFieldUnion   `json:"fields,omitzero" api:"required"`
	Relationships map[string]ObjectUpdateParamsRelationship `json:"relationships,omitzero"`
	paramObj
}

func (r ObjectUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow ObjectUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ObjectUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ObjectUpdateParamsEntityType string

const (
	ObjectUpdateParamsEntityTypeAccounts      ObjectUpdateParamsEntityType = "accounts"
	ObjectUpdateParamsEntityTypeContacts      ObjectUpdateParamsEntityType = "contacts"
	ObjectUpdateParamsEntityTypeOpportunities ObjectUpdateParamsEntityType = "opportunities"
	ObjectUpdateParamsEntityTypeTasks         ObjectUpdateParamsEntityType = "tasks"
	ObjectUpdateParamsEntityTypeNotes         ObjectUpdateParamsEntityType = "notes"
	ObjectUpdateParamsEntityTypeMeetings      ObjectUpdateParamsEntityType = "meetings"
)

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ObjectUpdateParamsFieldUnion struct {
	OfString                   param.Opt[string]                              `json:",omitzero,inline"`
	OfFloat                    param.Opt[float64]                             `json:",omitzero,inline"`
	OfBool                     param.Opt[bool]                                `json:",omitzero,inline"`
	OfObjectUpdatesFieldArray  []ObjectUpdateParamsFieldArrayItemUnion        `json:",omitzero,inline"`
	OfObjectUpdatesFieldMapMap map[string]ObjectUpdateParamsFieldMapItemUnion `json:",omitzero,inline"`
	paramUnion
}

func (u ObjectUpdateParamsFieldUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString,
		u.OfFloat,
		u.OfBool,
		u.OfObjectUpdatesFieldArray,
		u.OfObjectUpdatesFieldMapMap)
}
func (u *ObjectUpdateParamsFieldUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ObjectUpdateParamsFieldArrayItemUnion struct {
	OfString   param.Opt[string]  `json:",omitzero,inline"`
	OfFloat    param.Opt[float64] `json:",omitzero,inline"`
	OfBool     param.Opt[bool]    `json:",omitzero,inline"`
	OfAnyArray []any              `json:",omitzero,inline"`
	OfAnyMap   map[string]any     `json:",omitzero,inline"`
	paramUnion
}

func (u ObjectUpdateParamsFieldArrayItemUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString,
		u.OfFloat,
		u.OfBool,
		u.OfAnyArray,
		u.OfAnyMap)
}
func (u *ObjectUpdateParamsFieldArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ObjectUpdateParamsFieldMapItemUnion struct {
	OfString   param.Opt[string]  `json:",omitzero,inline"`
	OfFloat    param.Opt[float64] `json:",omitzero,inline"`
	OfBool     param.Opt[bool]    `json:",omitzero,inline"`
	OfAnyArray []any              `json:",omitzero,inline"`
	OfAnyMap   map[string]any     `json:",omitzero,inline"`
	paramUnion
}

func (u ObjectUpdateParamsFieldMapItemUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString,
		u.OfFloat,
		u.OfBool,
		u.OfAnyArray,
		u.OfAnyMap)
}
func (u *ObjectUpdateParamsFieldMapItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

type ObjectUpdateParamsRelationship struct {
	Add    ObjectUpdateParamsRelationshipAddUnion    `json:"add,omitzero"`
	Remove ObjectUpdateParamsRelationshipRemoveUnion `json:"remove,omitzero"`
	paramObj
}

func (r ObjectUpdateParamsRelationship) MarshalJSON() (data []byte, err error) {
	type shadow ObjectUpdateParamsRelationship
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ObjectUpdateParamsRelationship) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ObjectUpdateParamsRelationshipAddUnion struct {
	OfString      param.Opt[string] `json:",omitzero,inline"`
	OfStringArray []string          `json:",omitzero,inline"`
	paramUnion
}

func (u ObjectUpdateParamsRelationshipAddUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray)
}
func (u *ObjectUpdateParamsRelationshipAddUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ObjectUpdateParamsRelationshipRemoveUnion struct {
	OfString      param.Opt[string] `json:",omitzero,inline"`
	OfStringArray []string          `json:",omitzero,inline"`
	paramUnion
}

func (u ObjectUpdateParamsRelationshipRemoveUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray)
}
func (u *ObjectUpdateParamsRelationshipRemoveUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

type ObjectListParams struct {
	Limit  param.Opt[int64] `query:"limit,omitzero" json:"-"`
	Offset param.Opt[int64] `query:"offset,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ObjectListParams]'s query parameters as `url.Values`.
func (r ObjectListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ObjectListParamsEntityType string

const (
	ObjectListParamsEntityTypeAccounts      ObjectListParamsEntityType = "accounts"
	ObjectListParamsEntityTypeContacts      ObjectListParamsEntityType = "contacts"
	ObjectListParamsEntityTypeOpportunities ObjectListParamsEntityType = "opportunities"
	ObjectListParamsEntityTypeTasks         ObjectListParamsEntityType = "tasks"
	ObjectListParamsEntityTypeNotes         ObjectListParamsEntityType = "notes"
	ObjectListParamsEntityTypeMeetings      ObjectListParamsEntityType = "meetings"
)
