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

// Opportunities represent potential deals or sales in Lightfield. Each opportunity
// belongs to an account and can have tasks and notes associated with it.
//
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

// Creates a new opportunity record. The `$name` and `$stage` fields and the
// `$account` relationship are required.
//
// After creation, Lightfield automatically generates an opportunity summary in the
// background. The `$opportunityStatus` field is read-only and cannot be set via
// the API. The `$task` and `$note` relationships are also read-only — manage them
// via the `$opportunity` relationship on the task, or the
// `$account`/`$opportunity` note relationships instead.
//
// Supports idempotency via the `Idempotency-Key` header.
//
// **[Required scope](/using-the-api/scopes/):** `opportunities:create`
//
// **[Rate limit category](/using-the-api/rate-limits/):** Write
func (r *OpportunityService) New(ctx context.Context, body OpportunityNewParams, opts ...option.RequestOption) (res *OpportunityNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/opportunities"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Retrieves a single opportunity by its ID.
//
// **[Required scope](/using-the-api/scopes/):** `opportunities:read`
//
// **[Rate limit category](/using-the-api/rate-limits/):** Read
func (r *OpportunityService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *OpportunityGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/opportunities/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Updates an existing opportunity by ID. Only included fields and relationships
// are modified.
//
// The `$opportunityStatus` field is read-only and cannot be updated. The `$task`
// and `$note` relationships are also read-only — manage them via the
// `$opportunity` relationship on the task, or the `$account`/`$opportunity` note
// relationships instead.
//
// Supports idempotency via the `Idempotency-Key` header.
//
// **[Required scope](/using-the-api/scopes/):** `opportunities:update`
//
// **[Rate limit category](/using-the-api/rate-limits/):** Write
func (r *OpportunityService) Update(ctx context.Context, id string, body OpportunityUpdateParams, opts ...option.RequestOption) (res *OpportunityUpdateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/opportunities/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Returns a paginated list of opportunities. Use `offset` and `limit` to paginate
// through results. See <u>[List endpoints](/using-the-api/list-endpoints/)</u> for
// more information about pagination.
//
// **[Required scope](/using-the-api/scopes/):** `opportunities:read`
//
// **[Rate limit category](/using-the-api/rate-limits/):** Search
func (r *OpportunityService) List(ctx context.Context, query OpportunityListParams, opts ...option.RequestOption) (res *OpportunityListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/opportunities"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Returns the schema for all field and relationship definitions available on
// opportunities, including both system-defined and custom fields. Useful for
// understanding the shape of opportunity data before creating or updating records.
// See <u>[Fields and relationships](/using-the-api/fields-and-relationships/)</u>
// for more details.
//
// **[Required scope](/using-the-api/scopes/):** `opportunities:read`
//
// **[Rate limit category](/using-the-api/rate-limits/):** Read
func (r *OpportunityService) Definitions(ctx context.Context, opts ...option.RequestOption) (res *OpportunityDefinitionsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/opportunities/definitions"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

type OpportunityNewResponse struct {
	// Unique identifier for the entity.
	ID string `json:"id" api:"required"`
	// ISO 8601 timestamp of when the entity was created.
	CreatedAt string `json:"createdAt" api:"required"`
	// Map of field names to their typed values. System fields are prefixed with `$`
	// (e.g. `$name`, `$email`); custom attributes use their bare slug.
	Fields map[string]OpportunityNewResponseField `json:"fields" api:"required"`
	// URL to view the entity in the Lightfield web app, or null.
	HTTPLink string `json:"httpLink" api:"required"`
	// Map of relationship names to their associated entities. System relationships are
	// prefixed with `$` (e.g. `$owner`, `$contact`).
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
	// The field value, or null if unset.
	Value OpportunityNewResponseFieldValueUnion `json:"value" api:"required"`
	// The data type of the field.
	//
	// Any of "ADDRESS", "CHECKBOX", "CURRENCY", "DATETIME", "EMAIL", "FULL_NAME",
	// "MARKDOWN", "MULTI_SELECT", "NUMBER", "SINGLE_SELECT", "SOCIAL_HANDLE",
	// "TELEPHONE", "TEXT", "URL".
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
	// Whether the relationship is `has_one` or `has_many`.
	Cardinality string `json:"cardinality" api:"required"`
	// The type of the related object (e.g. `account`, `contact`).
	ObjectType string `json:"objectType" api:"required"`
	// IDs of the related entities.
	Values []string `json:"values" api:"required"`
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
	// Unique identifier for the entity.
	ID string `json:"id" api:"required"`
	// ISO 8601 timestamp of when the entity was created.
	CreatedAt string `json:"createdAt" api:"required"`
	// Map of field names to their typed values. System fields are prefixed with `$`
	// (e.g. `$name`, `$email`); custom attributes use their bare slug.
	Fields map[string]OpportunityGetResponseField `json:"fields" api:"required"`
	// URL to view the entity in the Lightfield web app, or null.
	HTTPLink string `json:"httpLink" api:"required"`
	// Map of relationship names to their associated entities. System relationships are
	// prefixed with `$` (e.g. `$owner`, `$contact`).
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
	// The field value, or null if unset.
	Value OpportunityGetResponseFieldValueUnion `json:"value" api:"required"`
	// The data type of the field.
	//
	// Any of "ADDRESS", "CHECKBOX", "CURRENCY", "DATETIME", "EMAIL", "FULL_NAME",
	// "MARKDOWN", "MULTI_SELECT", "NUMBER", "SINGLE_SELECT", "SOCIAL_HANDLE",
	// "TELEPHONE", "TEXT", "URL".
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
	// Whether the relationship is `has_one` or `has_many`.
	Cardinality string `json:"cardinality" api:"required"`
	// The type of the related object (e.g. `account`, `contact`).
	ObjectType string `json:"objectType" api:"required"`
	// IDs of the related entities.
	Values []string `json:"values" api:"required"`
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
	// Unique identifier for the entity.
	ID string `json:"id" api:"required"`
	// ISO 8601 timestamp of when the entity was created.
	CreatedAt string `json:"createdAt" api:"required"`
	// Map of field names to their typed values. System fields are prefixed with `$`
	// (e.g. `$name`, `$email`); custom attributes use their bare slug.
	Fields map[string]OpportunityUpdateResponseField `json:"fields" api:"required"`
	// URL to view the entity in the Lightfield web app, or null.
	HTTPLink string `json:"httpLink" api:"required"`
	// Map of relationship names to their associated entities. System relationships are
	// prefixed with `$` (e.g. `$owner`, `$contact`).
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
	// The field value, or null if unset.
	Value OpportunityUpdateResponseFieldValueUnion `json:"value" api:"required"`
	// The data type of the field.
	//
	// Any of "ADDRESS", "CHECKBOX", "CURRENCY", "DATETIME", "EMAIL", "FULL_NAME",
	// "MARKDOWN", "MULTI_SELECT", "NUMBER", "SINGLE_SELECT", "SOCIAL_HANDLE",
	// "TELEPHONE", "TEXT", "URL".
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
	// Whether the relationship is `has_one` or `has_many`.
	Cardinality string `json:"cardinality" api:"required"`
	// The type of the related object (e.g. `account`, `contact`).
	ObjectType string `json:"objectType" api:"required"`
	// IDs of the related entities.
	Values []string `json:"values" api:"required"`
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
	// Array of entity objects for the current page.
	Data []OpportunityListResponseData `json:"data" api:"required"`
	// The object type, always `"list"`.
	Object string `json:"object" api:"required"`
	// Total number of entities matching the query.
	TotalCount int64 `json:"totalCount" api:"required"`
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
	// Unique identifier for the entity.
	ID string `json:"id" api:"required"`
	// ISO 8601 timestamp of when the entity was created.
	CreatedAt string `json:"createdAt" api:"required"`
	// Map of field names to their typed values. System fields are prefixed with `$`
	// (e.g. `$name`, `$email`); custom attributes use their bare slug.
	Fields map[string]OpportunityListResponseDataField `json:"fields" api:"required"`
	// URL to view the entity in the Lightfield web app, or null.
	HTTPLink string `json:"httpLink" api:"required"`
	// Map of relationship names to their associated entities. System relationships are
	// prefixed with `$` (e.g. `$owner`, `$contact`).
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
	// The field value, or null if unset.
	Value OpportunityListResponseDataFieldValueUnion `json:"value" api:"required"`
	// The data type of the field.
	//
	// Any of "ADDRESS", "CHECKBOX", "CURRENCY", "DATETIME", "EMAIL", "FULL_NAME",
	// "MARKDOWN", "MULTI_SELECT", "NUMBER", "SINGLE_SELECT", "SOCIAL_HANDLE",
	// "TELEPHONE", "TEXT", "URL".
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
	// Whether the relationship is `has_one` or `has_many`.
	Cardinality string `json:"cardinality" api:"required"`
	// The type of the related object (e.g. `account`, `contact`).
	ObjectType string `json:"objectType" api:"required"`
	// IDs of the related entities.
	Values []string `json:"values" api:"required"`
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

type OpportunityDefinitionsResponse struct {
	// Map of field keys to their definitions, including both system and custom fields.
	FieldDefinitions map[string]OpportunityDefinitionsResponseFieldDefinition `json:"fieldDefinitions" api:"required"`
	// The object type these definitions belong to (e.g. `account`).
	ObjectType string `json:"objectType" api:"required"`
	// Map of relationship keys to their definitions.
	RelationshipDefinitions map[string]OpportunityDefinitionsResponseRelationshipDefinition `json:"relationshipDefinitions" api:"required"`
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
func (r OpportunityDefinitionsResponse) RawJSON() string { return r.JSON.raw }
func (r *OpportunityDefinitionsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OpportunityDefinitionsResponseFieldDefinition struct {
	// Description of the field, or null.
	Description string `json:"description" api:"required"`
	// Human-readable display name of the field.
	Label string `json:"label" api:"required"`
	// Type-specific configuration (e.g. select options, currency code).
	TypeConfiguration map[string]OpportunityDefinitionsResponseFieldDefinitionTypeConfigurationUnion `json:"typeConfiguration" api:"required"`
	// Data type of the field.
	//
	// Any of "ADDRESS", "CHECKBOX", "CURRENCY", "DATETIME", "EMAIL", "FULL_NAME",
	// "MARKDOWN", "MULTI_SELECT", "NUMBER", "SINGLE_SELECT", "SOCIAL_HANDLE",
	// "TELEPHONE", "TEXT", "URL".
	ValueType string `json:"valueType" api:"required"`
	// Unique identifier of the field definition.
	ID string `json:"id"`
	// `true` for fields that are not writable via the API (e.g. AI-generated
	// summaries). `false` or absent for writable fields.
	ReadOnly bool `json:"readOnly"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Description       respjson.Field
		Label             respjson.Field
		TypeConfiguration respjson.Field
		ValueType         respjson.Field
		ID                respjson.Field
		ReadOnly          respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OpportunityDefinitionsResponseFieldDefinition) RawJSON() string { return r.JSON.raw }
func (r *OpportunityDefinitionsResponseFieldDefinition) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// OpportunityDefinitionsResponseFieldDefinitionTypeConfigurationUnion contains all
// possible properties and values from [string], [float64], [bool],
// [[]OpportunityDefinitionsResponseFieldDefinitionTypeConfigurationArrayItemUnion],
// [map[string]OpportunityDefinitionsResponseFieldDefinitionTypeConfigurationMapItemUnion].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool
// OfOpportunityDefinitionsResponseFieldDefinitionTypeConfigurationArray OfAnyArray
// OfOpportunityDefinitionsResponseFieldDefinitionTypeConfigurationMapItemMapItem]
type OpportunityDefinitionsResponseFieldDefinitionTypeConfigurationUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a
	// [[]OpportunityDefinitionsResponseFieldDefinitionTypeConfigurationArrayItemUnion]
	// instead of an object.
	OfOpportunityDefinitionsResponseFieldDefinitionTypeConfigurationArray []OpportunityDefinitionsResponseFieldDefinitionTypeConfigurationArrayItemUnion `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	// This field will be present if the value is a [any] instead of an object.
	OfOpportunityDefinitionsResponseFieldDefinitionTypeConfigurationMapItemMapItem any `json:",inline"`
	JSON                                                                           struct {
		OfString                                                                       respjson.Field
		OfFloat                                                                        respjson.Field
		OfBool                                                                         respjson.Field
		OfOpportunityDefinitionsResponseFieldDefinitionTypeConfigurationArray          respjson.Field
		OfAnyArray                                                                     respjson.Field
		OfOpportunityDefinitionsResponseFieldDefinitionTypeConfigurationMapItemMapItem respjson.Field
		raw                                                                            string
	} `json:"-"`
}

func (u OpportunityDefinitionsResponseFieldDefinitionTypeConfigurationUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u OpportunityDefinitionsResponseFieldDefinitionTypeConfigurationUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u OpportunityDefinitionsResponseFieldDefinitionTypeConfigurationUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u OpportunityDefinitionsResponseFieldDefinitionTypeConfigurationUnion) AsOpportunityDefinitionsResponseFieldDefinitionTypeConfigurationArray() (v []OpportunityDefinitionsResponseFieldDefinitionTypeConfigurationArrayItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u OpportunityDefinitionsResponseFieldDefinitionTypeConfigurationUnion) AsOpportunityDefinitionsResponseFieldDefinitionTypeConfigurationMapMap() (v map[string]OpportunityDefinitionsResponseFieldDefinitionTypeConfigurationMapItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u OpportunityDefinitionsResponseFieldDefinitionTypeConfigurationUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *OpportunityDefinitionsResponseFieldDefinitionTypeConfigurationUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// OpportunityDefinitionsResponseFieldDefinitionTypeConfigurationArrayItemUnion
// contains all possible properties and values from [string], [float64], [bool],
// [[]any], [map[string]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray
// OfOpportunityDefinitionsResponseFieldDefinitionTypeConfigurationArrayItemMapItem]
type OpportunityDefinitionsResponseFieldDefinitionTypeConfigurationArrayItemUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	// This field will be present if the value is a [any] instead of an object.
	OfOpportunityDefinitionsResponseFieldDefinitionTypeConfigurationArrayItemMapItem any `json:",inline"`
	JSON                                                                             struct {
		OfString                                                                         respjson.Field
		OfFloat                                                                          respjson.Field
		OfBool                                                                           respjson.Field
		OfAnyArray                                                                       respjson.Field
		OfOpportunityDefinitionsResponseFieldDefinitionTypeConfigurationArrayItemMapItem respjson.Field
		raw                                                                              string
	} `json:"-"`
}

func (u OpportunityDefinitionsResponseFieldDefinitionTypeConfigurationArrayItemUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u OpportunityDefinitionsResponseFieldDefinitionTypeConfigurationArrayItemUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u OpportunityDefinitionsResponseFieldDefinitionTypeConfigurationArrayItemUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u OpportunityDefinitionsResponseFieldDefinitionTypeConfigurationArrayItemUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u OpportunityDefinitionsResponseFieldDefinitionTypeConfigurationArrayItemUnion) AsAnyMap() (v map[string]any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u OpportunityDefinitionsResponseFieldDefinitionTypeConfigurationArrayItemUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *OpportunityDefinitionsResponseFieldDefinitionTypeConfigurationArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// OpportunityDefinitionsResponseFieldDefinitionTypeConfigurationMapItemUnion
// contains all possible properties and values from [string], [float64], [bool],
// [[]any], [map[string]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray
// OfOpportunityDefinitionsResponseFieldDefinitionTypeConfigurationMapItemMapItem]
type OpportunityDefinitionsResponseFieldDefinitionTypeConfigurationMapItemUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	// This field will be present if the value is a [any] instead of an object.
	OfOpportunityDefinitionsResponseFieldDefinitionTypeConfigurationMapItemMapItem any `json:",inline"`
	JSON                                                                           struct {
		OfString                                                                       respjson.Field
		OfFloat                                                                        respjson.Field
		OfBool                                                                         respjson.Field
		OfAnyArray                                                                     respjson.Field
		OfOpportunityDefinitionsResponseFieldDefinitionTypeConfigurationMapItemMapItem respjson.Field
		raw                                                                            string
	} `json:"-"`
}

func (u OpportunityDefinitionsResponseFieldDefinitionTypeConfigurationMapItemUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u OpportunityDefinitionsResponseFieldDefinitionTypeConfigurationMapItemUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u OpportunityDefinitionsResponseFieldDefinitionTypeConfigurationMapItemUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u OpportunityDefinitionsResponseFieldDefinitionTypeConfigurationMapItemUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u OpportunityDefinitionsResponseFieldDefinitionTypeConfigurationMapItemUnion) AsAnyMap() (v map[string]any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u OpportunityDefinitionsResponseFieldDefinitionTypeConfigurationMapItemUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *OpportunityDefinitionsResponseFieldDefinitionTypeConfigurationMapItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OpportunityDefinitionsResponseRelationshipDefinition struct {
	// Whether this is a `has_one` or `has_many` relationship.
	//
	// Any of "HAS_ONE", "HAS_MANY".
	Cardinality string `json:"cardinality" api:"required"`
	// Description of the relationship, or null.
	Description string `json:"description" api:"required"`
	// Human-readable display name of the relationship.
	Label string `json:"label" api:"required"`
	// The type of the related object (e.g. `account`, `contact`).
	ObjectType string `json:"objectType" api:"required"`
	// Unique identifier of the relationship definition.
	ID string `json:"id"`
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
func (r OpportunityDefinitionsResponseRelationshipDefinition) RawJSON() string { return r.JSON.raw }
func (r *OpportunityDefinitionsResponseRelationshipDefinition) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OpportunityNewParams struct {
	// Field values for the new opportunity. System fields use a `$` prefix (e.g.
	// `$name`, `$stage`); custom attributes use their bare slug. Required: `$name`
	// (string) and `$stage` (option ID or label). Fields of type `SINGLE_SELECT` or
	// `MULTI_SELECT` accept either an option ID or label from the field's
	// `typeConfiguration.options` — call the
	// <u>[definitions endpoint](/api/resources/opportunity/methods/definitions)</u> to
	// discover available fields and options. See
	// <u>[Fields and relationships](/using-the-api/fields-and-relationships/)</u> for
	// value type details.
	Fields OpportunityNewParamsFields `json:"fields,omitzero" api:"required"`
	// Relationships to set on the new opportunity. System relationships use a `$`
	// prefix (e.g. `$account`, `$owner`); custom relationships use their bare slug.
	// `$account` is required. Each value is a single entity ID or an array of IDs.
	// Call the
	// <u>[definitions endpoint](/api/resources/opportunity/methods/definitions)</u> to
	// list available relationship keys.
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

// Field values for the new opportunity. System fields use a `$` prefix (e.g.
// `$name`, `$stage`); custom attributes use their bare slug. Required: `$name`
// (string) and `$stage` (option ID or label). Fields of type `SINGLE_SELECT` or
// `MULTI_SELECT` accept either an option ID or label from the field's
// `typeConfiguration.options` — call the
// <u>[definitions endpoint](/api/resources/opportunity/methods/definitions)</u> to
// discover available fields and options. See
// <u>[Fields and relationships](/using-the-api/fields-and-relationships/)</u> for
// value type details.
//
// The properties Name, Stage are required.
type OpportunityNewParamsFields struct {
	// Display name of the opportunity.
	Name string `json:"$name" api:"required"`
	// Pipeline stage (`SINGLE_SELECT`). Pass the option ID or label from the field
	// definition.
	Stage       string                                    `json:"$stage" api:"required"`
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

// Relationships to set on the new opportunity. System relationships use a `$`
// prefix (e.g. `$account`, `$owner`); custom relationships use their bare slug.
// `$account` is required. Each value is a single entity ID or an array of IDs.
// Call the
// <u>[definitions endpoint](/api/resources/opportunity/methods/definitions)</u> to
// list available relationship keys.
//
// The property Account is required.
type OpportunityNewParamsRelationships struct {
	// ID of the account this opportunity belongs to.
	Account OpportunityNewParamsRelationshipsAccountUnion `json:"$account,omitzero" api:"required"`
	// ID of the contact who is the internal champion.
	Champion OpportunityNewParamsRelationshipsChampionUnion `json:"$champion,omitzero"`
	// ID of the user who created this opportunity.
	CreatedBy OpportunityNewParamsRelationshipsCreatedByUnion `json:"$createdBy,omitzero"`
	// ID of the contact who is the evaluator.
	Evaluator OpportunityNewParamsRelationshipsEvaluatorUnion `json:"$evaluator,omitzero"`
	// ID of the user who owns this opportunity.
	Owner       OpportunityNewParamsRelationshipsOwnerUnion      `json:"$owner,omitzero"`
	ExtraFields map[string]OpportunityNewParamsRelationshipUnion `json:"-"`
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
type OpportunityNewParamsRelationshipsAccountUnion struct {
	OfString      param.Opt[string] `json:",omitzero,inline"`
	OfStringArray []string          `json:",omitzero,inline"`
	paramUnion
}

func (u OpportunityNewParamsRelationshipsAccountUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray)
}
func (u *OpportunityNewParamsRelationshipsAccountUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type OpportunityNewParamsRelationshipsChampionUnion struct {
	OfString      param.Opt[string] `json:",omitzero,inline"`
	OfStringArray []string          `json:",omitzero,inline"`
	paramUnion
}

func (u OpportunityNewParamsRelationshipsChampionUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray)
}
func (u *OpportunityNewParamsRelationshipsChampionUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type OpportunityNewParamsRelationshipsCreatedByUnion struct {
	OfString      param.Opt[string] `json:",omitzero,inline"`
	OfStringArray []string          `json:",omitzero,inline"`
	paramUnion
}

func (u OpportunityNewParamsRelationshipsCreatedByUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray)
}
func (u *OpportunityNewParamsRelationshipsCreatedByUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type OpportunityNewParamsRelationshipsEvaluatorUnion struct {
	OfString      param.Opt[string] `json:",omitzero,inline"`
	OfStringArray []string          `json:",omitzero,inline"`
	paramUnion
}

func (u OpportunityNewParamsRelationshipsEvaluatorUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray)
}
func (u *OpportunityNewParamsRelationshipsEvaluatorUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type OpportunityNewParamsRelationshipsOwnerUnion struct {
	OfString      param.Opt[string] `json:",omitzero,inline"`
	OfStringArray []string          `json:",omitzero,inline"`
	paramUnion
}

func (u OpportunityNewParamsRelationshipsOwnerUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray)
}
func (u *OpportunityNewParamsRelationshipsOwnerUnion) UnmarshalJSON(data []byte) error {
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
	// Field values to update — only provided fields are modified; omitted fields are
	// left unchanged. System fields use a `$` prefix (e.g. `$name`, `$stage`); custom
	// attributes use their bare slug. `SINGLE_SELECT` and `MULTI_SELECT` fields accept
	// an option ID or label — call the
	// <u>[definitions endpoint](/api/resources/opportunity/methods/definitions)</u>
	// for available options. See
	// <u>[Fields and relationships](/using-the-api/fields-and-relationships/)</u> for
	// value type details.
	Fields OpportunityUpdateParamsFields `json:"fields,omitzero"`
	// Relationship operations to apply. System relationships use a `$` prefix (e.g.
	// `$owner`, `$champion`). Each value is an operation object with `add`, `remove`,
	// or `replace`.
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

// Field values to update — only provided fields are modified; omitted fields are
// left unchanged. System fields use a `$` prefix (e.g. `$name`, `$stage`); custom
// attributes use their bare slug. `SINGLE_SELECT` and `MULTI_SELECT` fields accept
// an option ID or label — call the
// <u>[definitions endpoint](/api/resources/opportunity/methods/definitions)</u>
// for available options. See
// <u>[Fields and relationships](/using-the-api/fields-and-relationships/)</u> for
// value type details.
type OpportunityUpdateParamsFields struct {
	// Display name of the opportunity.
	Name param.Opt[string] `json:"$name,omitzero"`
	// Pipeline stage (`SINGLE_SELECT`). Pass the option ID or label from the field
	// definition.
	Stage       param.Opt[string]                            `json:"$stage,omitzero"`
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

// Relationship operations to apply. System relationships use a `$` prefix (e.g.
// `$owner`, `$champion`). Each value is an operation object with `add`, `remove`,
// or `replace`.
type OpportunityUpdateParamsRelationships struct {
	// Operation to modify the internal champion.
	Champion OpportunityUpdateParamsRelationshipsChampion `json:"$champion,omitzero"`
	// Operation to modify the evaluator.
	Evaluator OpportunityUpdateParamsRelationshipsEvaluator `json:"$evaluator,omitzero"`
	// Operation to modify the opportunity owner.
	Owner       OpportunityUpdateParamsRelationshipsOwner      `json:"$owner,omitzero"`
	ExtraFields map[string]OpportunityUpdateParamsRelationship `json:"-"`
	paramObj
}

func (r OpportunityUpdateParamsRelationships) MarshalJSON() (data []byte, err error) {
	type shadow OpportunityUpdateParamsRelationships
	return param.MarshalWithExtras(r, (*shadow)(&r), r.ExtraFields)
}
func (r *OpportunityUpdateParamsRelationships) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Operation to modify the internal champion.
type OpportunityUpdateParamsRelationshipsChampion struct {
	// Entity ID(s) to add to the relationship.
	Add OpportunityUpdateParamsRelationshipsChampionAddUnion `json:"add,omitzero"`
	// Entity ID(s) to remove from the relationship.
	Remove OpportunityUpdateParamsRelationshipsChampionRemoveUnion `json:"remove,omitzero"`
	// Entity ID(s) to set as the entire relationship, replacing all existing
	// associations.
	Replace OpportunityUpdateParamsRelationshipsChampionReplaceUnion `json:"replace,omitzero"`
	paramObj
}

func (r OpportunityUpdateParamsRelationshipsChampion) MarshalJSON() (data []byte, err error) {
	type shadow OpportunityUpdateParamsRelationshipsChampion
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *OpportunityUpdateParamsRelationshipsChampion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type OpportunityUpdateParamsRelationshipsChampionAddUnion struct {
	OfString      param.Opt[string] `json:",omitzero,inline"`
	OfStringArray []string          `json:",omitzero,inline"`
	paramUnion
}

func (u OpportunityUpdateParamsRelationshipsChampionAddUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray)
}
func (u *OpportunityUpdateParamsRelationshipsChampionAddUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type OpportunityUpdateParamsRelationshipsChampionRemoveUnion struct {
	OfString      param.Opt[string] `json:",omitzero,inline"`
	OfStringArray []string          `json:",omitzero,inline"`
	paramUnion
}

func (u OpportunityUpdateParamsRelationshipsChampionRemoveUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray)
}
func (u *OpportunityUpdateParamsRelationshipsChampionRemoveUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type OpportunityUpdateParamsRelationshipsChampionReplaceUnion struct {
	OfString      param.Opt[string] `json:",omitzero,inline"`
	OfStringArray []string          `json:",omitzero,inline"`
	paramUnion
}

func (u OpportunityUpdateParamsRelationshipsChampionReplaceUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray)
}
func (u *OpportunityUpdateParamsRelationshipsChampionReplaceUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Operation to modify the evaluator.
type OpportunityUpdateParamsRelationshipsEvaluator struct {
	// Entity ID(s) to add to the relationship.
	Add OpportunityUpdateParamsRelationshipsEvaluatorAddUnion `json:"add,omitzero"`
	// Entity ID(s) to remove from the relationship.
	Remove OpportunityUpdateParamsRelationshipsEvaluatorRemoveUnion `json:"remove,omitzero"`
	// Entity ID(s) to set as the entire relationship, replacing all existing
	// associations.
	Replace OpportunityUpdateParamsRelationshipsEvaluatorReplaceUnion `json:"replace,omitzero"`
	paramObj
}

func (r OpportunityUpdateParamsRelationshipsEvaluator) MarshalJSON() (data []byte, err error) {
	type shadow OpportunityUpdateParamsRelationshipsEvaluator
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *OpportunityUpdateParamsRelationshipsEvaluator) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type OpportunityUpdateParamsRelationshipsEvaluatorAddUnion struct {
	OfString      param.Opt[string] `json:",omitzero,inline"`
	OfStringArray []string          `json:",omitzero,inline"`
	paramUnion
}

func (u OpportunityUpdateParamsRelationshipsEvaluatorAddUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray)
}
func (u *OpportunityUpdateParamsRelationshipsEvaluatorAddUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type OpportunityUpdateParamsRelationshipsEvaluatorRemoveUnion struct {
	OfString      param.Opt[string] `json:",omitzero,inline"`
	OfStringArray []string          `json:",omitzero,inline"`
	paramUnion
}

func (u OpportunityUpdateParamsRelationshipsEvaluatorRemoveUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray)
}
func (u *OpportunityUpdateParamsRelationshipsEvaluatorRemoveUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type OpportunityUpdateParamsRelationshipsEvaluatorReplaceUnion struct {
	OfString      param.Opt[string] `json:",omitzero,inline"`
	OfStringArray []string          `json:",omitzero,inline"`
	paramUnion
}

func (u OpportunityUpdateParamsRelationshipsEvaluatorReplaceUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray)
}
func (u *OpportunityUpdateParamsRelationshipsEvaluatorReplaceUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Operation to modify the opportunity owner.
type OpportunityUpdateParamsRelationshipsOwner struct {
	// Entity ID(s) to add to the relationship.
	Add OpportunityUpdateParamsRelationshipsOwnerAddUnion `json:"add,omitzero"`
	// Entity ID(s) to remove from the relationship.
	Remove OpportunityUpdateParamsRelationshipsOwnerRemoveUnion `json:"remove,omitzero"`
	// Entity ID(s) to set as the entire relationship, replacing all existing
	// associations.
	Replace OpportunityUpdateParamsRelationshipsOwnerReplaceUnion `json:"replace,omitzero"`
	paramObj
}

func (r OpportunityUpdateParamsRelationshipsOwner) MarshalJSON() (data []byte, err error) {
	type shadow OpportunityUpdateParamsRelationshipsOwner
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *OpportunityUpdateParamsRelationshipsOwner) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type OpportunityUpdateParamsRelationshipsOwnerAddUnion struct {
	OfString      param.Opt[string] `json:",omitzero,inline"`
	OfStringArray []string          `json:",omitzero,inline"`
	paramUnion
}

func (u OpportunityUpdateParamsRelationshipsOwnerAddUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray)
}
func (u *OpportunityUpdateParamsRelationshipsOwnerAddUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type OpportunityUpdateParamsRelationshipsOwnerRemoveUnion struct {
	OfString      param.Opt[string] `json:",omitzero,inline"`
	OfStringArray []string          `json:",omitzero,inline"`
	paramUnion
}

func (u OpportunityUpdateParamsRelationshipsOwnerRemoveUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray)
}
func (u *OpportunityUpdateParamsRelationshipsOwnerRemoveUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type OpportunityUpdateParamsRelationshipsOwnerReplaceUnion struct {
	OfString      param.Opt[string] `json:",omitzero,inline"`
	OfStringArray []string          `json:",omitzero,inline"`
	paramUnion
}

func (u OpportunityUpdateParamsRelationshipsOwnerReplaceUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray)
}
func (u *OpportunityUpdateParamsRelationshipsOwnerReplaceUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// An operation to modify a relationship. Provide one of `add`, `remove`, or
// `replace`.
type OpportunityUpdateParamsRelationship struct {
	// Entity ID(s) to add to the relationship.
	Add OpportunityUpdateParamsRelationshipAddUnion `json:"add,omitzero"`
	// Entity ID(s) to remove from the relationship.
	Remove OpportunityUpdateParamsRelationshipRemoveUnion `json:"remove,omitzero"`
	// Entity ID(s) to set as the entire relationship, replacing all existing
	// associations.
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
	// Maximum number of records to return. Defaults to 25, maximum 25.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Number of records to skip for pagination. Defaults to 0.
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
