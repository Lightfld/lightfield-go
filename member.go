// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package githubcomlightfldlightfieldgo

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

// Members represent users in your Lightfield workspace. Members can own accounts
// and opportunities, and are referenced in relationships like `$owner` and
// `$createdBy`.
//
// MemberService contains methods and other services that help with interacting
// with the Lightfield API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewMemberService] method instead.
type MemberService struct {
	Options []option.RequestOption
}

// NewMemberService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewMemberService(opts ...option.RequestOption) (r MemberService) {
	r = MemberService{}
	r.Options = opts
	return
}

// Retrieves a single member by their ID.
//
// **[Required scope](/using-the-api/scopes/):** `members:read`
//
// **[Rate limit category](/using-the-api/rate-limits/):** Read
func (r *MemberService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *MemberGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/members/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Returns a paginated list of members in your workspace. Use `offset` and `limit`
// to paginate through results. See
// <u>[List endpoints](/using-the-api/list-endpoints/)</u> for more information
// about pagination.
//
// **[Required scope](/using-the-api/scopes/):** `members:read`
//
// **[Rate limit category](/using-the-api/rate-limits/):** Search
func (r *MemberService) List(ctx context.Context, query MemberListParams, opts ...option.RequestOption) (res *MemberListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/members"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

type MemberGetResponse struct {
	ID        string                            `json:"id" api:"required"`
	CreatedAt string                            `json:"createdAt" api:"required"`
	Fields    map[string]MemberGetResponseField `json:"fields" api:"required"`
	HTTPLink  string                            `json:"httpLink" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		Fields      respjson.Field
		HTTPLink    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MemberGetResponse) RawJSON() string { return r.JSON.raw }
func (r *MemberGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MemberGetResponseField struct {
	Value     MemberGetResponseFieldValueUnion `json:"value" api:"required"`
	ValueType string                           `json:"valueType" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Value       respjson.Field
		ValueType   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MemberGetResponseField) RawJSON() string { return r.JSON.raw }
func (r *MemberGetResponseField) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// MemberGetResponseFieldValueUnion contains all possible properties and values
// from [string], [float64], [bool], [[]MemberGetResponseFieldValueArrayItemUnion],
// [map[string]MemberGetResponseFieldValueMapItemUnion].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfMemberGetResponseFieldValueArray
// OfAnyArray OfMemberGetResponseFieldValueMapItemMapItem]
type MemberGetResponseFieldValueUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a
	// [[]MemberGetResponseFieldValueArrayItemUnion] instead of an object.
	OfMemberGetResponseFieldValueArray []MemberGetResponseFieldValueArrayItemUnion `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	// This field will be present if the value is a [any] instead of an object.
	OfMemberGetResponseFieldValueMapItemMapItem any `json:",inline"`
	JSON                                        struct {
		OfString                                    respjson.Field
		OfFloat                                     respjson.Field
		OfBool                                      respjson.Field
		OfMemberGetResponseFieldValueArray          respjson.Field
		OfAnyArray                                  respjson.Field
		OfMemberGetResponseFieldValueMapItemMapItem respjson.Field
		raw                                         string
	} `json:"-"`
}

func (u MemberGetResponseFieldValueUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MemberGetResponseFieldValueUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MemberGetResponseFieldValueUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MemberGetResponseFieldValueUnion) AsMemberGetResponseFieldValueArray() (v []MemberGetResponseFieldValueArrayItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MemberGetResponseFieldValueUnion) AsMemberGetResponseFieldValueMapMap() (v map[string]MemberGetResponseFieldValueMapItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u MemberGetResponseFieldValueUnion) RawJSON() string { return u.JSON.raw }

func (r *MemberGetResponseFieldValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// MemberGetResponseFieldValueArrayItemUnion contains all possible properties and
// values from [string], [float64], [bool], [[]any], [map[string]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray
// OfMemberGetResponseFieldValueArrayItemMapItem]
type MemberGetResponseFieldValueArrayItemUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	// This field will be present if the value is a [any] instead of an object.
	OfMemberGetResponseFieldValueArrayItemMapItem any `json:",inline"`
	JSON                                          struct {
		OfString                                      respjson.Field
		OfFloat                                       respjson.Field
		OfBool                                        respjson.Field
		OfAnyArray                                    respjson.Field
		OfMemberGetResponseFieldValueArrayItemMapItem respjson.Field
		raw                                           string
	} `json:"-"`
}

func (u MemberGetResponseFieldValueArrayItemUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MemberGetResponseFieldValueArrayItemUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MemberGetResponseFieldValueArrayItemUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MemberGetResponseFieldValueArrayItemUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MemberGetResponseFieldValueArrayItemUnion) AsAnyMap() (v map[string]any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u MemberGetResponseFieldValueArrayItemUnion) RawJSON() string { return u.JSON.raw }

func (r *MemberGetResponseFieldValueArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// MemberGetResponseFieldValueMapItemUnion contains all possible properties and
// values from [string], [float64], [bool], [[]any], [map[string]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray
// OfMemberGetResponseFieldValueMapItemMapItem]
type MemberGetResponseFieldValueMapItemUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	// This field will be present if the value is a [any] instead of an object.
	OfMemberGetResponseFieldValueMapItemMapItem any `json:",inline"`
	JSON                                        struct {
		OfString                                    respjson.Field
		OfFloat                                     respjson.Field
		OfBool                                      respjson.Field
		OfAnyArray                                  respjson.Field
		OfMemberGetResponseFieldValueMapItemMapItem respjson.Field
		raw                                         string
	} `json:"-"`
}

func (u MemberGetResponseFieldValueMapItemUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MemberGetResponseFieldValueMapItemUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MemberGetResponseFieldValueMapItemUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MemberGetResponseFieldValueMapItemUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MemberGetResponseFieldValueMapItemUnion) AsAnyMap() (v map[string]any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u MemberGetResponseFieldValueMapItemUnion) RawJSON() string { return u.JSON.raw }

func (r *MemberGetResponseFieldValueMapItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MemberListResponse struct {
	Data       []MemberListResponseData `json:"data" api:"required"`
	Object     string                   `json:"object" api:"required"`
	TotalCount int64                    `json:"totalCount" api:"required"`
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
func (r MemberListResponse) RawJSON() string { return r.JSON.raw }
func (r *MemberListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MemberListResponseData struct {
	ID        string                                 `json:"id" api:"required"`
	CreatedAt string                                 `json:"createdAt" api:"required"`
	Fields    map[string]MemberListResponseDataField `json:"fields" api:"required"`
	HTTPLink  string                                 `json:"httpLink" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		Fields      respjson.Field
		HTTPLink    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MemberListResponseData) RawJSON() string { return r.JSON.raw }
func (r *MemberListResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MemberListResponseDataField struct {
	Value     MemberListResponseDataFieldValueUnion `json:"value" api:"required"`
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
func (r MemberListResponseDataField) RawJSON() string { return r.JSON.raw }
func (r *MemberListResponseDataField) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// MemberListResponseDataFieldValueUnion contains all possible properties and
// values from [string], [float64], [bool],
// [[]MemberListResponseDataFieldValueArrayItemUnion],
// [map[string]MemberListResponseDataFieldValueMapItemUnion].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfMemberListResponseDataFieldValueArray
// OfAnyArray OfMemberListResponseDataFieldValueMapItemMapItem]
type MemberListResponseDataFieldValueUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a
	// [[]MemberListResponseDataFieldValueArrayItemUnion] instead of an object.
	OfMemberListResponseDataFieldValueArray []MemberListResponseDataFieldValueArrayItemUnion `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	// This field will be present if the value is a [any] instead of an object.
	OfMemberListResponseDataFieldValueMapItemMapItem any `json:",inline"`
	JSON                                             struct {
		OfString                                         respjson.Field
		OfFloat                                          respjson.Field
		OfBool                                           respjson.Field
		OfMemberListResponseDataFieldValueArray          respjson.Field
		OfAnyArray                                       respjson.Field
		OfMemberListResponseDataFieldValueMapItemMapItem respjson.Field
		raw                                              string
	} `json:"-"`
}

func (u MemberListResponseDataFieldValueUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MemberListResponseDataFieldValueUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MemberListResponseDataFieldValueUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MemberListResponseDataFieldValueUnion) AsMemberListResponseDataFieldValueArray() (v []MemberListResponseDataFieldValueArrayItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MemberListResponseDataFieldValueUnion) AsMemberListResponseDataFieldValueMapMap() (v map[string]MemberListResponseDataFieldValueMapItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u MemberListResponseDataFieldValueUnion) RawJSON() string { return u.JSON.raw }

func (r *MemberListResponseDataFieldValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// MemberListResponseDataFieldValueArrayItemUnion contains all possible properties
// and values from [string], [float64], [bool], [[]any], [map[string]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray
// OfMemberListResponseDataFieldValueArrayItemMapItem]
type MemberListResponseDataFieldValueArrayItemUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	// This field will be present if the value is a [any] instead of an object.
	OfMemberListResponseDataFieldValueArrayItemMapItem any `json:",inline"`
	JSON                                               struct {
		OfString                                           respjson.Field
		OfFloat                                            respjson.Field
		OfBool                                             respjson.Field
		OfAnyArray                                         respjson.Field
		OfMemberListResponseDataFieldValueArrayItemMapItem respjson.Field
		raw                                                string
	} `json:"-"`
}

func (u MemberListResponseDataFieldValueArrayItemUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MemberListResponseDataFieldValueArrayItemUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MemberListResponseDataFieldValueArrayItemUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MemberListResponseDataFieldValueArrayItemUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MemberListResponseDataFieldValueArrayItemUnion) AsAnyMap() (v map[string]any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u MemberListResponseDataFieldValueArrayItemUnion) RawJSON() string { return u.JSON.raw }

func (r *MemberListResponseDataFieldValueArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// MemberListResponseDataFieldValueMapItemUnion contains all possible properties
// and values from [string], [float64], [bool], [[]any], [map[string]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray
// OfMemberListResponseDataFieldValueMapItemMapItem]
type MemberListResponseDataFieldValueMapItemUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	// This field will be present if the value is a [any] instead of an object.
	OfMemberListResponseDataFieldValueMapItemMapItem any `json:",inline"`
	JSON                                             struct {
		OfString                                         respjson.Field
		OfFloat                                          respjson.Field
		OfBool                                           respjson.Field
		OfAnyArray                                       respjson.Field
		OfMemberListResponseDataFieldValueMapItemMapItem respjson.Field
		raw                                              string
	} `json:"-"`
}

func (u MemberListResponseDataFieldValueMapItemUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MemberListResponseDataFieldValueMapItemUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MemberListResponseDataFieldValueMapItemUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MemberListResponseDataFieldValueMapItemUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MemberListResponseDataFieldValueMapItemUnion) AsAnyMap() (v map[string]any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u MemberListResponseDataFieldValueMapItemUnion) RawJSON() string { return u.JSON.raw }

func (r *MemberListResponseDataFieldValueMapItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MemberListParams struct {
	Limit  param.Opt[int64] `query:"limit,omitzero" json:"-"`
	Offset param.Opt[int64] `query:"offset,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [MemberListParams]'s query parameters as `url.Values`.
func (r MemberListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
