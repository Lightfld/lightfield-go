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

// Contacts represent individual people in Lightfield. Contacts can be associated
// with one or more accounts.
//
// ContactService contains methods and other services that help with interacting
// with the Lightfield API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewContactService] method instead.
type ContactService struct {
	Options []option.RequestOption
}

// NewContactService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewContactService(opts ...option.RequestOption) (r ContactService) {
	r = ContactService{}
	r.Options = opts
	return
}

// Creates a new contact record.
//
// After creation, Lightfield automatically enriches the contact in the background.
//
// Supports idempotency via the `Idempotency-Key` header.
//
// **[Required scope](/using-the-api/scopes/):** `contacts:create`
//
// **[Rate limit category](/using-the-api/rate-limits/):** Write
func (r *ContactService) New(ctx context.Context, body ContactNewParams, opts ...option.RequestOption) (res *ContactNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/contacts"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Retrieves a single contact by its ID.
//
// **[Required scope](/using-the-api/scopes/):** `contacts:read`
//
// **[Rate limit category](/using-the-api/rate-limits/):** Read
func (r *ContactService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *ContactGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/contacts/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Updates an existing contact by ID. Only included fields and relationships are
// modified.
//
// Supports idempotency via the `Idempotency-Key` header.
//
// **[Required scope](/using-the-api/scopes/):** `contacts:update`
//
// **[Rate limit category](/using-the-api/rate-limits/):** Write
func (r *ContactService) Update(ctx context.Context, id string, body ContactUpdateParams, opts ...option.RequestOption) (res *ContactUpdateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/contacts/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Returns a paginated list of contacts. Use `offset` and `limit` to paginate
// through results. See <u>[List endpoints](/using-the-api/list-endpoints/)</u> for
// more information about pagination.
//
// **[Required scope](/using-the-api/scopes/):** `contacts:read`
//
// **[Rate limit category](/using-the-api/rate-limits/):** Search
func (r *ContactService) List(ctx context.Context, query ContactListParams, opts ...option.RequestOption) (res *ContactListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/contacts"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Returns the schema for all field and relationship definitions available on
// contacts, including both system-defined and custom fields. Useful for
// understanding the shape of contact data before creating or updating records. See
// <u>[Fields and relationships](/using-the-api/fields-and-relationships/)</u> for
// more details.
//
// **[Required scope](/using-the-api/scopes/):** `contacts:read`
//
// **[Rate limit category](/using-the-api/rate-limits/):** Read
func (r *ContactService) Definitions(ctx context.Context, opts ...option.RequestOption) (res *ContactDefinitionsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/contacts/definitions"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

type ContactNewResponse struct {
	ID            string                                    `json:"id" api:"required"`
	CreatedAt     string                                    `json:"createdAt" api:"required"`
	Fields        map[string]ContactNewResponseField        `json:"fields" api:"required"`
	HTTPLink      string                                    `json:"httpLink" api:"required"`
	Relationships map[string]ContactNewResponseRelationship `json:"relationships" api:"required"`
	ExtraFields   map[string]ContactNewResponseUnion        `json:"" api:"extrafields"`
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
func (r ContactNewResponse) RawJSON() string { return r.JSON.raw }
func (r *ContactNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ContactNewResponseField struct {
	Value     ContactNewResponseFieldValueUnion `json:"value" api:"required"`
	ValueType string                            `json:"valueType" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Value       respjson.Field
		ValueType   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ContactNewResponseField) RawJSON() string { return r.JSON.raw }
func (r *ContactNewResponseField) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ContactNewResponseFieldValueUnion contains all possible properties and values
// from [string], [float64], [bool],
// [[]ContactNewResponseFieldValueArrayItemUnion],
// [map[string]ContactNewResponseFieldValueMapItemUnion].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfContactNewResponseFieldValueArray
// OfAnyArray OfContactNewResponseFieldValueMapItemMapItem]
type ContactNewResponseFieldValueUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a
	// [[]ContactNewResponseFieldValueArrayItemUnion] instead of an object.
	OfContactNewResponseFieldValueArray []ContactNewResponseFieldValueArrayItemUnion `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	// This field will be present if the value is a [any] instead of an object.
	OfContactNewResponseFieldValueMapItemMapItem any `json:",inline"`
	JSON                                         struct {
		OfString                                     respjson.Field
		OfFloat                                      respjson.Field
		OfBool                                       respjson.Field
		OfContactNewResponseFieldValueArray          respjson.Field
		OfAnyArray                                   respjson.Field
		OfContactNewResponseFieldValueMapItemMapItem respjson.Field
		raw                                          string
	} `json:"-"`
}

func (u ContactNewResponseFieldValueUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ContactNewResponseFieldValueUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ContactNewResponseFieldValueUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ContactNewResponseFieldValueUnion) AsContactNewResponseFieldValueArray() (v []ContactNewResponseFieldValueArrayItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ContactNewResponseFieldValueUnion) AsContactNewResponseFieldValueMapMap() (v map[string]ContactNewResponseFieldValueMapItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ContactNewResponseFieldValueUnion) RawJSON() string { return u.JSON.raw }

func (r *ContactNewResponseFieldValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ContactNewResponseFieldValueArrayItemUnion contains all possible properties and
// values from [string], [float64], [bool], [[]any], [map[string]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray
// OfContactNewResponseFieldValueArrayItemMapItem]
type ContactNewResponseFieldValueArrayItemUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	// This field will be present if the value is a [any] instead of an object.
	OfContactNewResponseFieldValueArrayItemMapItem any `json:",inline"`
	JSON                                           struct {
		OfString                                       respjson.Field
		OfFloat                                        respjson.Field
		OfBool                                         respjson.Field
		OfAnyArray                                     respjson.Field
		OfContactNewResponseFieldValueArrayItemMapItem respjson.Field
		raw                                            string
	} `json:"-"`
}

func (u ContactNewResponseFieldValueArrayItemUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ContactNewResponseFieldValueArrayItemUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ContactNewResponseFieldValueArrayItemUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ContactNewResponseFieldValueArrayItemUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ContactNewResponseFieldValueArrayItemUnion) AsAnyMap() (v map[string]any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ContactNewResponseFieldValueArrayItemUnion) RawJSON() string { return u.JSON.raw }

func (r *ContactNewResponseFieldValueArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ContactNewResponseFieldValueMapItemUnion contains all possible properties and
// values from [string], [float64], [bool], [[]any], [map[string]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray
// OfContactNewResponseFieldValueMapItemMapItem]
type ContactNewResponseFieldValueMapItemUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	// This field will be present if the value is a [any] instead of an object.
	OfContactNewResponseFieldValueMapItemMapItem any `json:",inline"`
	JSON                                         struct {
		OfString                                     respjson.Field
		OfFloat                                      respjson.Field
		OfBool                                       respjson.Field
		OfAnyArray                                   respjson.Field
		OfContactNewResponseFieldValueMapItemMapItem respjson.Field
		raw                                          string
	} `json:"-"`
}

func (u ContactNewResponseFieldValueMapItemUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ContactNewResponseFieldValueMapItemUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ContactNewResponseFieldValueMapItemUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ContactNewResponseFieldValueMapItemUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ContactNewResponseFieldValueMapItemUnion) AsAnyMap() (v map[string]any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ContactNewResponseFieldValueMapItemUnion) RawJSON() string { return u.JSON.raw }

func (r *ContactNewResponseFieldValueMapItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ContactNewResponseRelationship struct {
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
func (r ContactNewResponseRelationship) RawJSON() string { return r.JSON.raw }
func (r *ContactNewResponseRelationship) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ContactNewResponseUnion contains all possible properties and values from
// [string], [float64], [bool], [[]ContactNewResponseArrayItemUnion],
// [map[string]ContactNewResponseMapItemUnion].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfContactNewResponseArray OfAnyArray
// OfContactNewResponseMapItemMapItem]
type ContactNewResponseUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a
	// [[]ContactNewResponseArrayItemUnion] instead of an object.
	OfContactNewResponseArray []ContactNewResponseArrayItemUnion `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	// This field will be present if the value is a [any] instead of an object.
	OfContactNewResponseMapItemMapItem any `json:",inline"`
	JSON                               struct {
		OfString                           respjson.Field
		OfFloat                            respjson.Field
		OfBool                             respjson.Field
		OfContactNewResponseArray          respjson.Field
		OfAnyArray                         respjson.Field
		OfContactNewResponseMapItemMapItem respjson.Field
		raw                                string
	} `json:"-"`
}

func (u ContactNewResponseUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ContactNewResponseUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ContactNewResponseUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ContactNewResponseUnion) AsContactNewResponseArray() (v []ContactNewResponseArrayItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ContactNewResponseUnion) AsContactNewResponseMapMap() (v map[string]ContactNewResponseMapItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ContactNewResponseUnion) RawJSON() string { return u.JSON.raw }

func (r *ContactNewResponseUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ContactNewResponseArrayItemUnion contains all possible properties and values
// from [string], [float64], [bool], [[]any], [map[string]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray
// OfContactNewResponseArrayItemMapItem]
type ContactNewResponseArrayItemUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	// This field will be present if the value is a [any] instead of an object.
	OfContactNewResponseArrayItemMapItem any `json:",inline"`
	JSON                                 struct {
		OfString                             respjson.Field
		OfFloat                              respjson.Field
		OfBool                               respjson.Field
		OfAnyArray                           respjson.Field
		OfContactNewResponseArrayItemMapItem respjson.Field
		raw                                  string
	} `json:"-"`
}

func (u ContactNewResponseArrayItemUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ContactNewResponseArrayItemUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ContactNewResponseArrayItemUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ContactNewResponseArrayItemUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ContactNewResponseArrayItemUnion) AsAnyMap() (v map[string]any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ContactNewResponseArrayItemUnion) RawJSON() string { return u.JSON.raw }

func (r *ContactNewResponseArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ContactNewResponseMapItemUnion contains all possible properties and values from
// [string], [float64], [bool], [[]any], [map[string]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray
// OfContactNewResponseMapItemMapItem]
type ContactNewResponseMapItemUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	// This field will be present if the value is a [any] instead of an object.
	OfContactNewResponseMapItemMapItem any `json:",inline"`
	JSON                               struct {
		OfString                           respjson.Field
		OfFloat                            respjson.Field
		OfBool                             respjson.Field
		OfAnyArray                         respjson.Field
		OfContactNewResponseMapItemMapItem respjson.Field
		raw                                string
	} `json:"-"`
}

func (u ContactNewResponseMapItemUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ContactNewResponseMapItemUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ContactNewResponseMapItemUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ContactNewResponseMapItemUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ContactNewResponseMapItemUnion) AsAnyMap() (v map[string]any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ContactNewResponseMapItemUnion) RawJSON() string { return u.JSON.raw }

func (r *ContactNewResponseMapItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ContactGetResponse struct {
	ID            string                                    `json:"id" api:"required"`
	CreatedAt     string                                    `json:"createdAt" api:"required"`
	Fields        map[string]ContactGetResponseField        `json:"fields" api:"required"`
	HTTPLink      string                                    `json:"httpLink" api:"required"`
	Relationships map[string]ContactGetResponseRelationship `json:"relationships" api:"required"`
	ExtraFields   map[string]ContactGetResponseUnion        `json:"" api:"extrafields"`
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
func (r ContactGetResponse) RawJSON() string { return r.JSON.raw }
func (r *ContactGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ContactGetResponseField struct {
	Value     ContactGetResponseFieldValueUnion `json:"value" api:"required"`
	ValueType string                            `json:"valueType" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Value       respjson.Field
		ValueType   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ContactGetResponseField) RawJSON() string { return r.JSON.raw }
func (r *ContactGetResponseField) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ContactGetResponseFieldValueUnion contains all possible properties and values
// from [string], [float64], [bool],
// [[]ContactGetResponseFieldValueArrayItemUnion],
// [map[string]ContactGetResponseFieldValueMapItemUnion].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfContactGetResponseFieldValueArray
// OfAnyArray OfContactGetResponseFieldValueMapItemMapItem]
type ContactGetResponseFieldValueUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a
	// [[]ContactGetResponseFieldValueArrayItemUnion] instead of an object.
	OfContactGetResponseFieldValueArray []ContactGetResponseFieldValueArrayItemUnion `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	// This field will be present if the value is a [any] instead of an object.
	OfContactGetResponseFieldValueMapItemMapItem any `json:",inline"`
	JSON                                         struct {
		OfString                                     respjson.Field
		OfFloat                                      respjson.Field
		OfBool                                       respjson.Field
		OfContactGetResponseFieldValueArray          respjson.Field
		OfAnyArray                                   respjson.Field
		OfContactGetResponseFieldValueMapItemMapItem respjson.Field
		raw                                          string
	} `json:"-"`
}

func (u ContactGetResponseFieldValueUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ContactGetResponseFieldValueUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ContactGetResponseFieldValueUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ContactGetResponseFieldValueUnion) AsContactGetResponseFieldValueArray() (v []ContactGetResponseFieldValueArrayItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ContactGetResponseFieldValueUnion) AsContactGetResponseFieldValueMapMap() (v map[string]ContactGetResponseFieldValueMapItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ContactGetResponseFieldValueUnion) RawJSON() string { return u.JSON.raw }

func (r *ContactGetResponseFieldValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ContactGetResponseFieldValueArrayItemUnion contains all possible properties and
// values from [string], [float64], [bool], [[]any], [map[string]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray
// OfContactGetResponseFieldValueArrayItemMapItem]
type ContactGetResponseFieldValueArrayItemUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	// This field will be present if the value is a [any] instead of an object.
	OfContactGetResponseFieldValueArrayItemMapItem any `json:",inline"`
	JSON                                           struct {
		OfString                                       respjson.Field
		OfFloat                                        respjson.Field
		OfBool                                         respjson.Field
		OfAnyArray                                     respjson.Field
		OfContactGetResponseFieldValueArrayItemMapItem respjson.Field
		raw                                            string
	} `json:"-"`
}

func (u ContactGetResponseFieldValueArrayItemUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ContactGetResponseFieldValueArrayItemUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ContactGetResponseFieldValueArrayItemUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ContactGetResponseFieldValueArrayItemUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ContactGetResponseFieldValueArrayItemUnion) AsAnyMap() (v map[string]any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ContactGetResponseFieldValueArrayItemUnion) RawJSON() string { return u.JSON.raw }

func (r *ContactGetResponseFieldValueArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ContactGetResponseFieldValueMapItemUnion contains all possible properties and
// values from [string], [float64], [bool], [[]any], [map[string]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray
// OfContactGetResponseFieldValueMapItemMapItem]
type ContactGetResponseFieldValueMapItemUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	// This field will be present if the value is a [any] instead of an object.
	OfContactGetResponseFieldValueMapItemMapItem any `json:",inline"`
	JSON                                         struct {
		OfString                                     respjson.Field
		OfFloat                                      respjson.Field
		OfBool                                       respjson.Field
		OfAnyArray                                   respjson.Field
		OfContactGetResponseFieldValueMapItemMapItem respjson.Field
		raw                                          string
	} `json:"-"`
}

func (u ContactGetResponseFieldValueMapItemUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ContactGetResponseFieldValueMapItemUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ContactGetResponseFieldValueMapItemUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ContactGetResponseFieldValueMapItemUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ContactGetResponseFieldValueMapItemUnion) AsAnyMap() (v map[string]any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ContactGetResponseFieldValueMapItemUnion) RawJSON() string { return u.JSON.raw }

func (r *ContactGetResponseFieldValueMapItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ContactGetResponseRelationship struct {
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
func (r ContactGetResponseRelationship) RawJSON() string { return r.JSON.raw }
func (r *ContactGetResponseRelationship) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ContactGetResponseUnion contains all possible properties and values from
// [string], [float64], [bool], [[]ContactGetResponseArrayItemUnion],
// [map[string]ContactGetResponseMapItemUnion].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfContactGetResponseArray OfAnyArray
// OfContactGetResponseMapItemMapItem]
type ContactGetResponseUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a
	// [[]ContactGetResponseArrayItemUnion] instead of an object.
	OfContactGetResponseArray []ContactGetResponseArrayItemUnion `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	// This field will be present if the value is a [any] instead of an object.
	OfContactGetResponseMapItemMapItem any `json:",inline"`
	JSON                               struct {
		OfString                           respjson.Field
		OfFloat                            respjson.Field
		OfBool                             respjson.Field
		OfContactGetResponseArray          respjson.Field
		OfAnyArray                         respjson.Field
		OfContactGetResponseMapItemMapItem respjson.Field
		raw                                string
	} `json:"-"`
}

func (u ContactGetResponseUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ContactGetResponseUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ContactGetResponseUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ContactGetResponseUnion) AsContactGetResponseArray() (v []ContactGetResponseArrayItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ContactGetResponseUnion) AsContactGetResponseMapMap() (v map[string]ContactGetResponseMapItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ContactGetResponseUnion) RawJSON() string { return u.JSON.raw }

func (r *ContactGetResponseUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ContactGetResponseArrayItemUnion contains all possible properties and values
// from [string], [float64], [bool], [[]any], [map[string]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray
// OfContactGetResponseArrayItemMapItem]
type ContactGetResponseArrayItemUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	// This field will be present if the value is a [any] instead of an object.
	OfContactGetResponseArrayItemMapItem any `json:",inline"`
	JSON                                 struct {
		OfString                             respjson.Field
		OfFloat                              respjson.Field
		OfBool                               respjson.Field
		OfAnyArray                           respjson.Field
		OfContactGetResponseArrayItemMapItem respjson.Field
		raw                                  string
	} `json:"-"`
}

func (u ContactGetResponseArrayItemUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ContactGetResponseArrayItemUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ContactGetResponseArrayItemUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ContactGetResponseArrayItemUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ContactGetResponseArrayItemUnion) AsAnyMap() (v map[string]any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ContactGetResponseArrayItemUnion) RawJSON() string { return u.JSON.raw }

func (r *ContactGetResponseArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ContactGetResponseMapItemUnion contains all possible properties and values from
// [string], [float64], [bool], [[]any], [map[string]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray
// OfContactGetResponseMapItemMapItem]
type ContactGetResponseMapItemUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	// This field will be present if the value is a [any] instead of an object.
	OfContactGetResponseMapItemMapItem any `json:",inline"`
	JSON                               struct {
		OfString                           respjson.Field
		OfFloat                            respjson.Field
		OfBool                             respjson.Field
		OfAnyArray                         respjson.Field
		OfContactGetResponseMapItemMapItem respjson.Field
		raw                                string
	} `json:"-"`
}

func (u ContactGetResponseMapItemUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ContactGetResponseMapItemUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ContactGetResponseMapItemUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ContactGetResponseMapItemUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ContactGetResponseMapItemUnion) AsAnyMap() (v map[string]any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ContactGetResponseMapItemUnion) RawJSON() string { return u.JSON.raw }

func (r *ContactGetResponseMapItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ContactUpdateResponse struct {
	ID            string                                       `json:"id" api:"required"`
	CreatedAt     string                                       `json:"createdAt" api:"required"`
	Fields        map[string]ContactUpdateResponseField        `json:"fields" api:"required"`
	HTTPLink      string                                       `json:"httpLink" api:"required"`
	Relationships map[string]ContactUpdateResponseRelationship `json:"relationships" api:"required"`
	ExtraFields   map[string]ContactUpdateResponseUnion        `json:"" api:"extrafields"`
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
func (r ContactUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *ContactUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ContactUpdateResponseField struct {
	Value     ContactUpdateResponseFieldValueUnion `json:"value" api:"required"`
	ValueType string                               `json:"valueType" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Value       respjson.Field
		ValueType   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ContactUpdateResponseField) RawJSON() string { return r.JSON.raw }
func (r *ContactUpdateResponseField) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ContactUpdateResponseFieldValueUnion contains all possible properties and values
// from [string], [float64], [bool],
// [[]ContactUpdateResponseFieldValueArrayItemUnion],
// [map[string]ContactUpdateResponseFieldValueMapItemUnion].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfContactUpdateResponseFieldValueArray
// OfAnyArray OfContactUpdateResponseFieldValueMapItemMapItem]
type ContactUpdateResponseFieldValueUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a
	// [[]ContactUpdateResponseFieldValueArrayItemUnion] instead of an object.
	OfContactUpdateResponseFieldValueArray []ContactUpdateResponseFieldValueArrayItemUnion `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	// This field will be present if the value is a [any] instead of an object.
	OfContactUpdateResponseFieldValueMapItemMapItem any `json:",inline"`
	JSON                                            struct {
		OfString                                        respjson.Field
		OfFloat                                         respjson.Field
		OfBool                                          respjson.Field
		OfContactUpdateResponseFieldValueArray          respjson.Field
		OfAnyArray                                      respjson.Field
		OfContactUpdateResponseFieldValueMapItemMapItem respjson.Field
		raw                                             string
	} `json:"-"`
}

func (u ContactUpdateResponseFieldValueUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ContactUpdateResponseFieldValueUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ContactUpdateResponseFieldValueUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ContactUpdateResponseFieldValueUnion) AsContactUpdateResponseFieldValueArray() (v []ContactUpdateResponseFieldValueArrayItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ContactUpdateResponseFieldValueUnion) AsContactUpdateResponseFieldValueMapMap() (v map[string]ContactUpdateResponseFieldValueMapItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ContactUpdateResponseFieldValueUnion) RawJSON() string { return u.JSON.raw }

func (r *ContactUpdateResponseFieldValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ContactUpdateResponseFieldValueArrayItemUnion contains all possible properties
// and values from [string], [float64], [bool], [[]any], [map[string]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray
// OfContactUpdateResponseFieldValueArrayItemMapItem]
type ContactUpdateResponseFieldValueArrayItemUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	// This field will be present if the value is a [any] instead of an object.
	OfContactUpdateResponseFieldValueArrayItemMapItem any `json:",inline"`
	JSON                                              struct {
		OfString                                          respjson.Field
		OfFloat                                           respjson.Field
		OfBool                                            respjson.Field
		OfAnyArray                                        respjson.Field
		OfContactUpdateResponseFieldValueArrayItemMapItem respjson.Field
		raw                                               string
	} `json:"-"`
}

func (u ContactUpdateResponseFieldValueArrayItemUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ContactUpdateResponseFieldValueArrayItemUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ContactUpdateResponseFieldValueArrayItemUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ContactUpdateResponseFieldValueArrayItemUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ContactUpdateResponseFieldValueArrayItemUnion) AsAnyMap() (v map[string]any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ContactUpdateResponseFieldValueArrayItemUnion) RawJSON() string { return u.JSON.raw }

func (r *ContactUpdateResponseFieldValueArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ContactUpdateResponseFieldValueMapItemUnion contains all possible properties and
// values from [string], [float64], [bool], [[]any], [map[string]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray
// OfContactUpdateResponseFieldValueMapItemMapItem]
type ContactUpdateResponseFieldValueMapItemUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	// This field will be present if the value is a [any] instead of an object.
	OfContactUpdateResponseFieldValueMapItemMapItem any `json:",inline"`
	JSON                                            struct {
		OfString                                        respjson.Field
		OfFloat                                         respjson.Field
		OfBool                                          respjson.Field
		OfAnyArray                                      respjson.Field
		OfContactUpdateResponseFieldValueMapItemMapItem respjson.Field
		raw                                             string
	} `json:"-"`
}

func (u ContactUpdateResponseFieldValueMapItemUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ContactUpdateResponseFieldValueMapItemUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ContactUpdateResponseFieldValueMapItemUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ContactUpdateResponseFieldValueMapItemUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ContactUpdateResponseFieldValueMapItemUnion) AsAnyMap() (v map[string]any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ContactUpdateResponseFieldValueMapItemUnion) RawJSON() string { return u.JSON.raw }

func (r *ContactUpdateResponseFieldValueMapItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ContactUpdateResponseRelationship struct {
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
func (r ContactUpdateResponseRelationship) RawJSON() string { return r.JSON.raw }
func (r *ContactUpdateResponseRelationship) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ContactUpdateResponseUnion contains all possible properties and values from
// [string], [float64], [bool], [[]ContactUpdateResponseArrayItemUnion],
// [map[string]ContactUpdateResponseMapItemUnion].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfContactUpdateResponseArray OfAnyArray
// OfContactUpdateResponseMapItemMapItem]
type ContactUpdateResponseUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a
	// [[]ContactUpdateResponseArrayItemUnion] instead of an object.
	OfContactUpdateResponseArray []ContactUpdateResponseArrayItemUnion `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	// This field will be present if the value is a [any] instead of an object.
	OfContactUpdateResponseMapItemMapItem any `json:",inline"`
	JSON                                  struct {
		OfString                              respjson.Field
		OfFloat                               respjson.Field
		OfBool                                respjson.Field
		OfContactUpdateResponseArray          respjson.Field
		OfAnyArray                            respjson.Field
		OfContactUpdateResponseMapItemMapItem respjson.Field
		raw                                   string
	} `json:"-"`
}

func (u ContactUpdateResponseUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ContactUpdateResponseUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ContactUpdateResponseUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ContactUpdateResponseUnion) AsContactUpdateResponseArray() (v []ContactUpdateResponseArrayItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ContactUpdateResponseUnion) AsContactUpdateResponseMapMap() (v map[string]ContactUpdateResponseMapItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ContactUpdateResponseUnion) RawJSON() string { return u.JSON.raw }

func (r *ContactUpdateResponseUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ContactUpdateResponseArrayItemUnion contains all possible properties and values
// from [string], [float64], [bool], [[]any], [map[string]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray
// OfContactUpdateResponseArrayItemMapItem]
type ContactUpdateResponseArrayItemUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	// This field will be present if the value is a [any] instead of an object.
	OfContactUpdateResponseArrayItemMapItem any `json:",inline"`
	JSON                                    struct {
		OfString                                respjson.Field
		OfFloat                                 respjson.Field
		OfBool                                  respjson.Field
		OfAnyArray                              respjson.Field
		OfContactUpdateResponseArrayItemMapItem respjson.Field
		raw                                     string
	} `json:"-"`
}

func (u ContactUpdateResponseArrayItemUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ContactUpdateResponseArrayItemUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ContactUpdateResponseArrayItemUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ContactUpdateResponseArrayItemUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ContactUpdateResponseArrayItemUnion) AsAnyMap() (v map[string]any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ContactUpdateResponseArrayItemUnion) RawJSON() string { return u.JSON.raw }

func (r *ContactUpdateResponseArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ContactUpdateResponseMapItemUnion contains all possible properties and values
// from [string], [float64], [bool], [[]any], [map[string]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray
// OfContactUpdateResponseMapItemMapItem]
type ContactUpdateResponseMapItemUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	// This field will be present if the value is a [any] instead of an object.
	OfContactUpdateResponseMapItemMapItem any `json:",inline"`
	JSON                                  struct {
		OfString                              respjson.Field
		OfFloat                               respjson.Field
		OfBool                                respjson.Field
		OfAnyArray                            respjson.Field
		OfContactUpdateResponseMapItemMapItem respjson.Field
		raw                                   string
	} `json:"-"`
}

func (u ContactUpdateResponseMapItemUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ContactUpdateResponseMapItemUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ContactUpdateResponseMapItemUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ContactUpdateResponseMapItemUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ContactUpdateResponseMapItemUnion) AsAnyMap() (v map[string]any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ContactUpdateResponseMapItemUnion) RawJSON() string { return u.JSON.raw }

func (r *ContactUpdateResponseMapItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ContactListResponse struct {
	Data       []ContactListResponseData `json:"data" api:"required"`
	Object     string                    `json:"object" api:"required"`
	TotalCount int64                     `json:"totalCount" api:"required"`
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
func (r ContactListResponse) RawJSON() string { return r.JSON.raw }
func (r *ContactListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ContactListResponseData struct {
	ID            string                                         `json:"id" api:"required"`
	CreatedAt     string                                         `json:"createdAt" api:"required"`
	Fields        map[string]ContactListResponseDataField        `json:"fields" api:"required"`
	HTTPLink      string                                         `json:"httpLink" api:"required"`
	Relationships map[string]ContactListResponseDataRelationship `json:"relationships" api:"required"`
	ExtraFields   map[string]ContactListResponseDataUnion        `json:"" api:"extrafields"`
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
func (r ContactListResponseData) RawJSON() string { return r.JSON.raw }
func (r *ContactListResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ContactListResponseDataField struct {
	Value     ContactListResponseDataFieldValueUnion `json:"value" api:"required"`
	ValueType string                                 `json:"valueType" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Value       respjson.Field
		ValueType   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ContactListResponseDataField) RawJSON() string { return r.JSON.raw }
func (r *ContactListResponseDataField) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ContactListResponseDataFieldValueUnion contains all possible properties and
// values from [string], [float64], [bool],
// [[]ContactListResponseDataFieldValueArrayItemUnion],
// [map[string]ContactListResponseDataFieldValueMapItemUnion].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfContactListResponseDataFieldValueArray
// OfAnyArray OfContactListResponseDataFieldValueMapItemMapItem]
type ContactListResponseDataFieldValueUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a
	// [[]ContactListResponseDataFieldValueArrayItemUnion] instead of an object.
	OfContactListResponseDataFieldValueArray []ContactListResponseDataFieldValueArrayItemUnion `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	// This field will be present if the value is a [any] instead of an object.
	OfContactListResponseDataFieldValueMapItemMapItem any `json:",inline"`
	JSON                                              struct {
		OfString                                          respjson.Field
		OfFloat                                           respjson.Field
		OfBool                                            respjson.Field
		OfContactListResponseDataFieldValueArray          respjson.Field
		OfAnyArray                                        respjson.Field
		OfContactListResponseDataFieldValueMapItemMapItem respjson.Field
		raw                                               string
	} `json:"-"`
}

func (u ContactListResponseDataFieldValueUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ContactListResponseDataFieldValueUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ContactListResponseDataFieldValueUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ContactListResponseDataFieldValueUnion) AsContactListResponseDataFieldValueArray() (v []ContactListResponseDataFieldValueArrayItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ContactListResponseDataFieldValueUnion) AsContactListResponseDataFieldValueMapMap() (v map[string]ContactListResponseDataFieldValueMapItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ContactListResponseDataFieldValueUnion) RawJSON() string { return u.JSON.raw }

func (r *ContactListResponseDataFieldValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ContactListResponseDataFieldValueArrayItemUnion contains all possible properties
// and values from [string], [float64], [bool], [[]any], [map[string]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray
// OfContactListResponseDataFieldValueArrayItemMapItem]
type ContactListResponseDataFieldValueArrayItemUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	// This field will be present if the value is a [any] instead of an object.
	OfContactListResponseDataFieldValueArrayItemMapItem any `json:",inline"`
	JSON                                                struct {
		OfString                                            respjson.Field
		OfFloat                                             respjson.Field
		OfBool                                              respjson.Field
		OfAnyArray                                          respjson.Field
		OfContactListResponseDataFieldValueArrayItemMapItem respjson.Field
		raw                                                 string
	} `json:"-"`
}

func (u ContactListResponseDataFieldValueArrayItemUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ContactListResponseDataFieldValueArrayItemUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ContactListResponseDataFieldValueArrayItemUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ContactListResponseDataFieldValueArrayItemUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ContactListResponseDataFieldValueArrayItemUnion) AsAnyMap() (v map[string]any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ContactListResponseDataFieldValueArrayItemUnion) RawJSON() string { return u.JSON.raw }

func (r *ContactListResponseDataFieldValueArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ContactListResponseDataFieldValueMapItemUnion contains all possible properties
// and values from [string], [float64], [bool], [[]any], [map[string]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray
// OfContactListResponseDataFieldValueMapItemMapItem]
type ContactListResponseDataFieldValueMapItemUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	// This field will be present if the value is a [any] instead of an object.
	OfContactListResponseDataFieldValueMapItemMapItem any `json:",inline"`
	JSON                                              struct {
		OfString                                          respjson.Field
		OfFloat                                           respjson.Field
		OfBool                                            respjson.Field
		OfAnyArray                                        respjson.Field
		OfContactListResponseDataFieldValueMapItemMapItem respjson.Field
		raw                                               string
	} `json:"-"`
}

func (u ContactListResponseDataFieldValueMapItemUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ContactListResponseDataFieldValueMapItemUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ContactListResponseDataFieldValueMapItemUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ContactListResponseDataFieldValueMapItemUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ContactListResponseDataFieldValueMapItemUnion) AsAnyMap() (v map[string]any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ContactListResponseDataFieldValueMapItemUnion) RawJSON() string { return u.JSON.raw }

func (r *ContactListResponseDataFieldValueMapItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ContactListResponseDataRelationship struct {
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
func (r ContactListResponseDataRelationship) RawJSON() string { return r.JSON.raw }
func (r *ContactListResponseDataRelationship) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ContactListResponseDataUnion contains all possible properties and values from
// [string], [float64], [bool], [[]ContactListResponseDataArrayItemUnion],
// [map[string]ContactListResponseDataMapItemUnion].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfContactListResponseDataArray OfAnyArray
// OfContactListResponseDataMapItemMapItem]
type ContactListResponseDataUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a
	// [[]ContactListResponseDataArrayItemUnion] instead of an object.
	OfContactListResponseDataArray []ContactListResponseDataArrayItemUnion `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	// This field will be present if the value is a [any] instead of an object.
	OfContactListResponseDataMapItemMapItem any `json:",inline"`
	JSON                                    struct {
		OfString                                respjson.Field
		OfFloat                                 respjson.Field
		OfBool                                  respjson.Field
		OfContactListResponseDataArray          respjson.Field
		OfAnyArray                              respjson.Field
		OfContactListResponseDataMapItemMapItem respjson.Field
		raw                                     string
	} `json:"-"`
}

func (u ContactListResponseDataUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ContactListResponseDataUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ContactListResponseDataUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ContactListResponseDataUnion) AsContactListResponseDataArray() (v []ContactListResponseDataArrayItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ContactListResponseDataUnion) AsContactListResponseDataMapMap() (v map[string]ContactListResponseDataMapItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ContactListResponseDataUnion) RawJSON() string { return u.JSON.raw }

func (r *ContactListResponseDataUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ContactListResponseDataArrayItemUnion contains all possible properties and
// values from [string], [float64], [bool], [[]any], [map[string]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray
// OfContactListResponseDataArrayItemMapItem]
type ContactListResponseDataArrayItemUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	// This field will be present if the value is a [any] instead of an object.
	OfContactListResponseDataArrayItemMapItem any `json:",inline"`
	JSON                                      struct {
		OfString                                  respjson.Field
		OfFloat                                   respjson.Field
		OfBool                                    respjson.Field
		OfAnyArray                                respjson.Field
		OfContactListResponseDataArrayItemMapItem respjson.Field
		raw                                       string
	} `json:"-"`
}

func (u ContactListResponseDataArrayItemUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ContactListResponseDataArrayItemUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ContactListResponseDataArrayItemUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ContactListResponseDataArrayItemUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ContactListResponseDataArrayItemUnion) AsAnyMap() (v map[string]any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ContactListResponseDataArrayItemUnion) RawJSON() string { return u.JSON.raw }

func (r *ContactListResponseDataArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ContactListResponseDataMapItemUnion contains all possible properties and values
// from [string], [float64], [bool], [[]any], [map[string]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray
// OfContactListResponseDataMapItemMapItem]
type ContactListResponseDataMapItemUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	// This field will be present if the value is a [any] instead of an object.
	OfContactListResponseDataMapItemMapItem any `json:",inline"`
	JSON                                    struct {
		OfString                                respjson.Field
		OfFloat                                 respjson.Field
		OfBool                                  respjson.Field
		OfAnyArray                              respjson.Field
		OfContactListResponseDataMapItemMapItem respjson.Field
		raw                                     string
	} `json:"-"`
}

func (u ContactListResponseDataMapItemUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ContactListResponseDataMapItemUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ContactListResponseDataMapItemUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ContactListResponseDataMapItemUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ContactListResponseDataMapItemUnion) AsAnyMap() (v map[string]any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ContactListResponseDataMapItemUnion) RawJSON() string { return u.JSON.raw }

func (r *ContactListResponseDataMapItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ContactDefinitionsResponse struct {
	FieldDefinitions        map[string]ContactDefinitionsResponseFieldDefinition        `json:"fieldDefinitions" api:"required"`
	ObjectType              string                                                      `json:"objectType" api:"required"`
	RelationshipDefinitions map[string]ContactDefinitionsResponseRelationshipDefinition `json:"relationshipDefinitions" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FieldDefinitions        respjson.Field
		ObjectType              respjson.Field
		RelationshipDefinitions respjson.Field
		ExtraFields             map[string]respjson.Field
		raw                     string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ContactDefinitionsResponse) RawJSON() string { return r.JSON.raw }
func (r *ContactDefinitionsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ContactDefinitionsResponseFieldDefinition struct {
	Description       string                                                                     `json:"description" api:"required"`
	Label             string                                                                     `json:"label" api:"required"`
	TypeConfiguration map[string]ContactDefinitionsResponseFieldDefinitionTypeConfigurationUnion `json:"typeConfiguration" api:"required"`
	ValueType         string                                                                     `json:"valueType" api:"required"`
	ID                string                                                                     `json:"id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Description       respjson.Field
		Label             respjson.Field
		TypeConfiguration respjson.Field
		ValueType         respjson.Field
		ID                respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ContactDefinitionsResponseFieldDefinition) RawJSON() string { return r.JSON.raw }
func (r *ContactDefinitionsResponseFieldDefinition) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ContactDefinitionsResponseFieldDefinitionTypeConfigurationUnion contains all
// possible properties and values from [string], [float64], [bool],
// [[]ContactDefinitionsResponseFieldDefinitionTypeConfigurationArrayItemUnion],
// [map[string]ContactDefinitionsResponseFieldDefinitionTypeConfigurationMapItemUnion].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool
// OfContactDefinitionsResponseFieldDefinitionTypeConfigurationArray OfAnyArray
// OfContactDefinitionsResponseFieldDefinitionTypeConfigurationMapItemMapItem]
type ContactDefinitionsResponseFieldDefinitionTypeConfigurationUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a
	// [[]ContactDefinitionsResponseFieldDefinitionTypeConfigurationArrayItemUnion]
	// instead of an object.
	OfContactDefinitionsResponseFieldDefinitionTypeConfigurationArray []ContactDefinitionsResponseFieldDefinitionTypeConfigurationArrayItemUnion `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	// This field will be present if the value is a [any] instead of an object.
	OfContactDefinitionsResponseFieldDefinitionTypeConfigurationMapItemMapItem any `json:",inline"`
	JSON                                                                       struct {
		OfString                                                                   respjson.Field
		OfFloat                                                                    respjson.Field
		OfBool                                                                     respjson.Field
		OfContactDefinitionsResponseFieldDefinitionTypeConfigurationArray          respjson.Field
		OfAnyArray                                                                 respjson.Field
		OfContactDefinitionsResponseFieldDefinitionTypeConfigurationMapItemMapItem respjson.Field
		raw                                                                        string
	} `json:"-"`
}

func (u ContactDefinitionsResponseFieldDefinitionTypeConfigurationUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ContactDefinitionsResponseFieldDefinitionTypeConfigurationUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ContactDefinitionsResponseFieldDefinitionTypeConfigurationUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ContactDefinitionsResponseFieldDefinitionTypeConfigurationUnion) AsContactDefinitionsResponseFieldDefinitionTypeConfigurationArray() (v []ContactDefinitionsResponseFieldDefinitionTypeConfigurationArrayItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ContactDefinitionsResponseFieldDefinitionTypeConfigurationUnion) AsContactDefinitionsResponseFieldDefinitionTypeConfigurationMapMap() (v map[string]ContactDefinitionsResponseFieldDefinitionTypeConfigurationMapItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ContactDefinitionsResponseFieldDefinitionTypeConfigurationUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ContactDefinitionsResponseFieldDefinitionTypeConfigurationUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ContactDefinitionsResponseFieldDefinitionTypeConfigurationArrayItemUnion
// contains all possible properties and values from [string], [float64], [bool],
// [[]any], [map[string]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray
// OfContactDefinitionsResponseFieldDefinitionTypeConfigurationArrayItemMapItem]
type ContactDefinitionsResponseFieldDefinitionTypeConfigurationArrayItemUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	// This field will be present if the value is a [any] instead of an object.
	OfContactDefinitionsResponseFieldDefinitionTypeConfigurationArrayItemMapItem any `json:",inline"`
	JSON                                                                         struct {
		OfString                                                                     respjson.Field
		OfFloat                                                                      respjson.Field
		OfBool                                                                       respjson.Field
		OfAnyArray                                                                   respjson.Field
		OfContactDefinitionsResponseFieldDefinitionTypeConfigurationArrayItemMapItem respjson.Field
		raw                                                                          string
	} `json:"-"`
}

func (u ContactDefinitionsResponseFieldDefinitionTypeConfigurationArrayItemUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ContactDefinitionsResponseFieldDefinitionTypeConfigurationArrayItemUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ContactDefinitionsResponseFieldDefinitionTypeConfigurationArrayItemUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ContactDefinitionsResponseFieldDefinitionTypeConfigurationArrayItemUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ContactDefinitionsResponseFieldDefinitionTypeConfigurationArrayItemUnion) AsAnyMap() (v map[string]any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ContactDefinitionsResponseFieldDefinitionTypeConfigurationArrayItemUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ContactDefinitionsResponseFieldDefinitionTypeConfigurationArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ContactDefinitionsResponseFieldDefinitionTypeConfigurationMapItemUnion contains
// all possible properties and values from [string], [float64], [bool], [[]any],
// [map[string]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray
// OfContactDefinitionsResponseFieldDefinitionTypeConfigurationMapItemMapItem]
type ContactDefinitionsResponseFieldDefinitionTypeConfigurationMapItemUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	// This field will be present if the value is a [any] instead of an object.
	OfContactDefinitionsResponseFieldDefinitionTypeConfigurationMapItemMapItem any `json:",inline"`
	JSON                                                                       struct {
		OfString                                                                   respjson.Field
		OfFloat                                                                    respjson.Field
		OfBool                                                                     respjson.Field
		OfAnyArray                                                                 respjson.Field
		OfContactDefinitionsResponseFieldDefinitionTypeConfigurationMapItemMapItem respjson.Field
		raw                                                                        string
	} `json:"-"`
}

func (u ContactDefinitionsResponseFieldDefinitionTypeConfigurationMapItemUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ContactDefinitionsResponseFieldDefinitionTypeConfigurationMapItemUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ContactDefinitionsResponseFieldDefinitionTypeConfigurationMapItemUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ContactDefinitionsResponseFieldDefinitionTypeConfigurationMapItemUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ContactDefinitionsResponseFieldDefinitionTypeConfigurationMapItemUnion) AsAnyMap() (v map[string]any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ContactDefinitionsResponseFieldDefinitionTypeConfigurationMapItemUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ContactDefinitionsResponseFieldDefinitionTypeConfigurationMapItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ContactDefinitionsResponseRelationshipDefinition struct {
	// Any of "HAS_ONE", "HAS_MANY".
	Cardinality string `json:"cardinality" api:"required"`
	Description string `json:"description" api:"required"`
	Label       string `json:"label" api:"required"`
	ObjectType  string `json:"objectType" api:"required"`
	ID          string `json:"id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Cardinality respjson.Field
		Description respjson.Field
		Label       respjson.Field
		ObjectType  respjson.Field
		ID          respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ContactDefinitionsResponseRelationshipDefinition) RawJSON() string { return r.JSON.raw }
func (r *ContactDefinitionsResponseRelationshipDefinition) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ContactNewParams struct {
	Fields        ContactNewParamsFields        `json:"fields,omitzero" api:"required"`
	Relationships ContactNewParamsRelationships `json:"relationships,omitzero"`
	paramObj
}

func (r ContactNewParams) MarshalJSON() (data []byte, err error) {
	type shadow ContactNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ContactNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ContactNewParamsFields struct {
	ProfilePhotoURL param.Opt[string]                     `json:"$profilePhotoUrl,omitzero"`
	Email           []string                              `json:"$email,omitzero"`
	Name            ContactNewParamsFieldsName            `json:"$name,omitzero"`
	ExtraFields     map[string]ContactNewParamsFieldUnion `json:"-"`
	paramObj
}

func (r ContactNewParamsFields) MarshalJSON() (data []byte, err error) {
	type shadow ContactNewParamsFields
	return param.MarshalWithExtras(r, (*shadow)(&r), r.ExtraFields)
}
func (r *ContactNewParamsFields) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ContactNewParamsFieldsName struct {
	FirstName param.Opt[string] `json:"firstName,omitzero"`
	LastName  param.Opt[string] `json:"lastName,omitzero"`
	paramObj
}

func (r ContactNewParamsFieldsName) MarshalJSON() (data []byte, err error) {
	type shadow ContactNewParamsFieldsName
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ContactNewParamsFieldsName) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ContactNewParamsFieldUnion struct {
	OfString                 param.Opt[string]                            `json:",omitzero,inline"`
	OfFloat                  param.Opt[float64]                           `json:",omitzero,inline"`
	OfBool                   param.Opt[bool]                              `json:",omitzero,inline"`
	OfContactNewsFieldArray  []ContactNewParamsFieldArrayItemUnion        `json:",omitzero,inline"`
	OfContactNewsFieldMapMap map[string]ContactNewParamsFieldMapItemUnion `json:",omitzero,inline"`
	paramUnion
}

func (u ContactNewParamsFieldUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString,
		u.OfFloat,
		u.OfBool,
		u.OfContactNewsFieldArray,
		u.OfContactNewsFieldMapMap)
}
func (u *ContactNewParamsFieldUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ContactNewParamsFieldArrayItemUnion struct {
	OfString   param.Opt[string]  `json:",omitzero,inline"`
	OfFloat    param.Opt[float64] `json:",omitzero,inline"`
	OfBool     param.Opt[bool]    `json:",omitzero,inline"`
	OfAnyArray []any              `json:",omitzero,inline"`
	OfAnyMap   map[string]any     `json:",omitzero,inline"`
	paramUnion
}

func (u ContactNewParamsFieldArrayItemUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString,
		u.OfFloat,
		u.OfBool,
		u.OfAnyArray,
		u.OfAnyMap)
}
func (u *ContactNewParamsFieldArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ContactNewParamsFieldMapItemUnion struct {
	OfString   param.Opt[string]  `json:",omitzero,inline"`
	OfFloat    param.Opt[float64] `json:",omitzero,inline"`
	OfBool     param.Opt[bool]    `json:",omitzero,inline"`
	OfAnyArray []any              `json:",omitzero,inline"`
	OfAnyMap   map[string]any     `json:",omitzero,inline"`
	paramUnion
}

func (u ContactNewParamsFieldMapItemUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString,
		u.OfFloat,
		u.OfBool,
		u.OfAnyArray,
		u.OfAnyMap)
}
func (u *ContactNewParamsFieldMapItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

type ContactNewParamsRelationships struct {
	Accounts    ContactNewParamsRelationshipsAccountsUnion   `json:"$accounts,omitzero"`
	ExtraFields map[string]ContactNewParamsRelationshipUnion `json:"-"`
	paramObj
}

func (r ContactNewParamsRelationships) MarshalJSON() (data []byte, err error) {
	type shadow ContactNewParamsRelationships
	return param.MarshalWithExtras(r, (*shadow)(&r), r.ExtraFields)
}
func (r *ContactNewParamsRelationships) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ContactNewParamsRelationshipsAccountsUnion struct {
	OfString      param.Opt[string] `json:",omitzero,inline"`
	OfStringArray []string          `json:",omitzero,inline"`
	paramUnion
}

func (u ContactNewParamsRelationshipsAccountsUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray)
}
func (u *ContactNewParamsRelationshipsAccountsUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ContactNewParamsRelationshipUnion struct {
	OfString      param.Opt[string] `json:",omitzero,inline"`
	OfStringArray []string          `json:",omitzero,inline"`
	paramUnion
}

func (u ContactNewParamsRelationshipUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray)
}
func (u *ContactNewParamsRelationshipUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

type ContactUpdateParams struct {
	Fields        ContactUpdateParamsFields        `json:"fields,omitzero"`
	Relationships ContactUpdateParamsRelationships `json:"relationships,omitzero"`
	paramObj
}

func (r ContactUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow ContactUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ContactUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ContactUpdateParamsFields struct {
	ProfilePhotoURL param.Opt[string]                        `json:"$profilePhotoUrl,omitzero"`
	Email           []string                                 `json:"$email,omitzero"`
	Name            ContactUpdateParamsFieldsName            `json:"$name,omitzero"`
	ExtraFields     map[string]ContactUpdateParamsFieldUnion `json:"-"`
	paramObj
}

func (r ContactUpdateParamsFields) MarshalJSON() (data []byte, err error) {
	type shadow ContactUpdateParamsFields
	return param.MarshalWithExtras(r, (*shadow)(&r), r.ExtraFields)
}
func (r *ContactUpdateParamsFields) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ContactUpdateParamsFieldsName struct {
	FirstName param.Opt[string] `json:"firstName,omitzero"`
	LastName  param.Opt[string] `json:"lastName,omitzero"`
	paramObj
}

func (r ContactUpdateParamsFieldsName) MarshalJSON() (data []byte, err error) {
	type shadow ContactUpdateParamsFieldsName
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ContactUpdateParamsFieldsName) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ContactUpdateParamsFieldUnion struct {
	OfString                    param.Opt[string]                               `json:",omitzero,inline"`
	OfFloat                     param.Opt[float64]                              `json:",omitzero,inline"`
	OfBool                      param.Opt[bool]                                 `json:",omitzero,inline"`
	OfContactUpdatesFieldArray  []ContactUpdateParamsFieldArrayItemUnion        `json:",omitzero,inline"`
	OfContactUpdatesFieldMapMap map[string]ContactUpdateParamsFieldMapItemUnion `json:",omitzero,inline"`
	paramUnion
}

func (u ContactUpdateParamsFieldUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString,
		u.OfFloat,
		u.OfBool,
		u.OfContactUpdatesFieldArray,
		u.OfContactUpdatesFieldMapMap)
}
func (u *ContactUpdateParamsFieldUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ContactUpdateParamsFieldArrayItemUnion struct {
	OfString   param.Opt[string]  `json:",omitzero,inline"`
	OfFloat    param.Opt[float64] `json:",omitzero,inline"`
	OfBool     param.Opt[bool]    `json:",omitzero,inline"`
	OfAnyArray []any              `json:",omitzero,inline"`
	OfAnyMap   map[string]any     `json:",omitzero,inline"`
	paramUnion
}

func (u ContactUpdateParamsFieldArrayItemUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString,
		u.OfFloat,
		u.OfBool,
		u.OfAnyArray,
		u.OfAnyMap)
}
func (u *ContactUpdateParamsFieldArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ContactUpdateParamsFieldMapItemUnion struct {
	OfString   param.Opt[string]  `json:",omitzero,inline"`
	OfFloat    param.Opt[float64] `json:",omitzero,inline"`
	OfBool     param.Opt[bool]    `json:",omitzero,inline"`
	OfAnyArray []any              `json:",omitzero,inline"`
	OfAnyMap   map[string]any     `json:",omitzero,inline"`
	paramUnion
}

func (u ContactUpdateParamsFieldMapItemUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString,
		u.OfFloat,
		u.OfBool,
		u.OfAnyArray,
		u.OfAnyMap)
}
func (u *ContactUpdateParamsFieldMapItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

type ContactUpdateParamsRelationships struct {
	Accounts    ContactUpdateParamsRelationshipsAccounts   `json:"$accounts,omitzero"`
	ExtraFields map[string]ContactUpdateParamsRelationship `json:"-"`
	paramObj
}

func (r ContactUpdateParamsRelationships) MarshalJSON() (data []byte, err error) {
	type shadow ContactUpdateParamsRelationships
	return param.MarshalWithExtras(r, (*shadow)(&r), r.ExtraFields)
}
func (r *ContactUpdateParamsRelationships) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ContactUpdateParamsRelationshipsAccounts struct {
	Add     ContactUpdateParamsRelationshipsAccountsAddUnion     `json:"add,omitzero"`
	Remove  ContactUpdateParamsRelationshipsAccountsRemoveUnion  `json:"remove,omitzero"`
	Replace ContactUpdateParamsRelationshipsAccountsReplaceUnion `json:"replace,omitzero"`
	paramObj
}

func (r ContactUpdateParamsRelationshipsAccounts) MarshalJSON() (data []byte, err error) {
	type shadow ContactUpdateParamsRelationshipsAccounts
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ContactUpdateParamsRelationshipsAccounts) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ContactUpdateParamsRelationshipsAccountsAddUnion struct {
	OfString      param.Opt[string] `json:",omitzero,inline"`
	OfStringArray []string          `json:",omitzero,inline"`
	paramUnion
}

func (u ContactUpdateParamsRelationshipsAccountsAddUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray)
}
func (u *ContactUpdateParamsRelationshipsAccountsAddUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ContactUpdateParamsRelationshipsAccountsRemoveUnion struct {
	OfString      param.Opt[string] `json:",omitzero,inline"`
	OfStringArray []string          `json:",omitzero,inline"`
	paramUnion
}

func (u ContactUpdateParamsRelationshipsAccountsRemoveUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray)
}
func (u *ContactUpdateParamsRelationshipsAccountsRemoveUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ContactUpdateParamsRelationshipsAccountsReplaceUnion struct {
	OfString      param.Opt[string] `json:",omitzero,inline"`
	OfStringArray []string          `json:",omitzero,inline"`
	paramUnion
}

func (u ContactUpdateParamsRelationshipsAccountsReplaceUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray)
}
func (u *ContactUpdateParamsRelationshipsAccountsReplaceUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

type ContactUpdateParamsRelationship struct {
	Add     ContactUpdateParamsRelationshipAddUnion     `json:"add,omitzero"`
	Remove  ContactUpdateParamsRelationshipRemoveUnion  `json:"remove,omitzero"`
	Replace ContactUpdateParamsRelationshipReplaceUnion `json:"replace,omitzero"`
	paramObj
}

func (r ContactUpdateParamsRelationship) MarshalJSON() (data []byte, err error) {
	type shadow ContactUpdateParamsRelationship
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ContactUpdateParamsRelationship) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ContactUpdateParamsRelationshipAddUnion struct {
	OfString      param.Opt[string] `json:",omitzero,inline"`
	OfStringArray []string          `json:",omitzero,inline"`
	paramUnion
}

func (u ContactUpdateParamsRelationshipAddUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray)
}
func (u *ContactUpdateParamsRelationshipAddUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ContactUpdateParamsRelationshipRemoveUnion struct {
	OfString      param.Opt[string] `json:",omitzero,inline"`
	OfStringArray []string          `json:",omitzero,inline"`
	paramUnion
}

func (u ContactUpdateParamsRelationshipRemoveUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray)
}
func (u *ContactUpdateParamsRelationshipRemoveUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ContactUpdateParamsRelationshipReplaceUnion struct {
	OfString      param.Opt[string] `json:",omitzero,inline"`
	OfStringArray []string          `json:",omitzero,inline"`
	paramUnion
}

func (u ContactUpdateParamsRelationshipReplaceUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray)
}
func (u *ContactUpdateParamsRelationshipReplaceUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

type ContactListParams struct {
	Limit  param.Opt[int64] `query:"limit,omitzero" json:"-"`
	Offset param.Opt[int64] `query:"offset,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ContactListParams]'s query parameters as `url.Values`.
func (r ContactListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
