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

// Accounts represent companies or organizations in Lightfield. Each account can
// have contacts, opportunities, tasks, and notes associated with it.
//
// AccountService contains methods and other services that help with interacting
// with the Lightfield API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountService] method instead.
type AccountService struct {
	Options []option.RequestOption
}

// NewAccountService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewAccountService(opts ...option.RequestOption) (r AccountService) {
	r = AccountService{}
	r.Options = opts
	return
}

// Creates a new account record. The `$name` field is required.
//
// If a `$website` is provided, Lightfield automatically enriches the account in
// the background. The `$howTheyMakeMoney` and `$accountStatus` fields are
// read-only and cannot be set via the API. The `$opportunity`, `$task`, and
// `$note` relationships are also read-only — manage them via the `$account`
// relationship on the opportunity or task, or the `$account`/`$opportunity` note
// relationships instead.
//
// Supports idempotency via the `Idempotency-Key` header.
//
// **[Required scope](/using-the-api/scopes/):** `accounts:create`
//
// **[Rate limit category](/using-the-api/rate-limits/):** Write
func (r *AccountService) New(ctx context.Context, body AccountNewParams, opts ...option.RequestOption) (res *AccountCreateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/accounts"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Retrieves a single account by its ID.
//
// **[Required scope](/using-the-api/scopes/):** `accounts:read`
//
// **[Rate limit category](/using-the-api/rate-limits/):** Read
func (r *AccountService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *AccountRetrieveResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/accounts/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Updates an existing account by ID. Only included fields and relationships are
// modified.
//
// The `$howTheyMakeMoney` and `$accountStatus` fields are read-only and cannot be
// updated. The `$opportunity`, `$task`, and `$note` relationships are also
// read-only — manage them via the `$account` relationship on the opportunity or
// task, or the `$account`/`$opportunity` note relationships instead.
//
// Supports idempotency via the `Idempotency-Key` header.
//
// **[Required scope](/using-the-api/scopes/):** `accounts:update`
//
// **[Rate limit category](/using-the-api/rate-limits/):** Write
func (r *AccountService) Update(ctx context.Context, id string, body AccountUpdateParams, opts ...option.RequestOption) (res *AccountUpdateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/accounts/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Returns a paginated list of accounts. Use `offset` and `limit` to paginate
// through results. See <u>[List endpoints](/using-the-api/list-endpoints/)</u> for
// more information about pagination.
//
// **[Required scope](/using-the-api/scopes/):** `accounts:read`
//
// **[Rate limit category](/using-the-api/rate-limits/):** Search
func (r *AccountService) List(ctx context.Context, query AccountListParams, opts ...option.RequestOption) (res *AccountListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/accounts"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Returns the schema for all field and relationship definitions available on
// accounts, including both system-defined and custom fields. Useful for
// understanding the shape of account data before creating or updating records. See
// <u>[Fields and relationships](/using-the-api/fields-and-relationships/)</u> for
// more details.
//
// **[Required scope](/using-the-api/scopes/):** `accounts:read`
//
// **[Rate limit category](/using-the-api/rate-limits/):** Read
func (r *AccountService) Definitions(ctx context.Context, opts ...option.RequestOption) (res *AccountDefinitionsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/accounts/definitions"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

type AccountCreateResponse struct {
	// Unique identifier for the entity.
	ID string `json:"id" api:"required"`
	// ISO 8601 timestamp of when the entity was created.
	CreatedAt string `json:"createdAt" api:"required"`
	// Map of field names to their typed values. System fields are prefixed with `$`
	// (e.g. `$name`, `$email`); custom attributes use their bare slug.
	Fields map[string]AccountCreateResponseField `json:"fields" api:"required"`
	// URL to view the entity in the Lightfield web app, or null.
	HTTPLink string `json:"httpLink" api:"required"`
	// Map of relationship names to their associated entities. System relationships are
	// prefixed with `$` (e.g. `$owner`, `$contact`).
	Relationships map[string]AccountCreateResponseRelationship `json:"relationships" api:"required"`
	ExtraFields   map[string]AccountCreateResponseUnion        `json:"" api:"extrafields"`
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
func (r AccountCreateResponse) RawJSON() string { return r.JSON.raw }
func (r *AccountCreateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AccountCreateResponseField struct {
	// The field value, or null if unset.
	Value AccountCreateResponseFieldValueUnion `json:"value" api:"required"`
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
func (r AccountCreateResponseField) RawJSON() string { return r.JSON.raw }
func (r *AccountCreateResponseField) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AccountCreateResponseFieldValueUnion contains all possible properties and values
// from [string], [float64], [bool],
// [[]AccountCreateResponseFieldValueArrayItemUnion],
// [map[string]AccountCreateResponseFieldValueMapItemUnion].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAccountCreateResponseFieldValueArray
// OfAnyArray OfAccountCreateResponseFieldValueMapItemMapItem]
type AccountCreateResponseFieldValueUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a
	// [[]AccountCreateResponseFieldValueArrayItemUnion] instead of an object.
	OfAccountCreateResponseFieldValueArray []AccountCreateResponseFieldValueArrayItemUnion `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	// This field will be present if the value is a [any] instead of an object.
	OfAccountCreateResponseFieldValueMapItemMapItem any `json:",inline"`
	JSON                                            struct {
		OfString                                        respjson.Field
		OfFloat                                         respjson.Field
		OfBool                                          respjson.Field
		OfAccountCreateResponseFieldValueArray          respjson.Field
		OfAnyArray                                      respjson.Field
		OfAccountCreateResponseFieldValueMapItemMapItem respjson.Field
		raw                                             string
	} `json:"-"`
}

func (u AccountCreateResponseFieldValueUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AccountCreateResponseFieldValueUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AccountCreateResponseFieldValueUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AccountCreateResponseFieldValueUnion) AsAccountCreateResponseFieldValueArray() (v []AccountCreateResponseFieldValueArrayItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AccountCreateResponseFieldValueUnion) AsAccountCreateResponseFieldValueMapMap() (v map[string]AccountCreateResponseFieldValueMapItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AccountCreateResponseFieldValueUnion) RawJSON() string { return u.JSON.raw }

func (r *AccountCreateResponseFieldValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AccountCreateResponseFieldValueArrayItemUnion contains all possible properties
// and values from [string], [float64], [bool], [[]any], [map[string]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray
// OfAccountCreateResponseFieldValueArrayItemMapItem]
type AccountCreateResponseFieldValueArrayItemUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	// This field will be present if the value is a [any] instead of an object.
	OfAccountCreateResponseFieldValueArrayItemMapItem any `json:",inline"`
	JSON                                              struct {
		OfString                                          respjson.Field
		OfFloat                                           respjson.Field
		OfBool                                            respjson.Field
		OfAnyArray                                        respjson.Field
		OfAccountCreateResponseFieldValueArrayItemMapItem respjson.Field
		raw                                               string
	} `json:"-"`
}

func (u AccountCreateResponseFieldValueArrayItemUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AccountCreateResponseFieldValueArrayItemUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AccountCreateResponseFieldValueArrayItemUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AccountCreateResponseFieldValueArrayItemUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AccountCreateResponseFieldValueArrayItemUnion) AsAnyMap() (v map[string]any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AccountCreateResponseFieldValueArrayItemUnion) RawJSON() string { return u.JSON.raw }

func (r *AccountCreateResponseFieldValueArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AccountCreateResponseFieldValueMapItemUnion contains all possible properties and
// values from [string], [float64], [bool], [[]any], [map[string]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray
// OfAccountCreateResponseFieldValueMapItemMapItem]
type AccountCreateResponseFieldValueMapItemUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	// This field will be present if the value is a [any] instead of an object.
	OfAccountCreateResponseFieldValueMapItemMapItem any `json:",inline"`
	JSON                                            struct {
		OfString                                        respjson.Field
		OfFloat                                         respjson.Field
		OfBool                                          respjson.Field
		OfAnyArray                                      respjson.Field
		OfAccountCreateResponseFieldValueMapItemMapItem respjson.Field
		raw                                             string
	} `json:"-"`
}

func (u AccountCreateResponseFieldValueMapItemUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AccountCreateResponseFieldValueMapItemUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AccountCreateResponseFieldValueMapItemUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AccountCreateResponseFieldValueMapItemUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AccountCreateResponseFieldValueMapItemUnion) AsAnyMap() (v map[string]any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AccountCreateResponseFieldValueMapItemUnion) RawJSON() string { return u.JSON.raw }

func (r *AccountCreateResponseFieldValueMapItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AccountCreateResponseRelationship struct {
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
func (r AccountCreateResponseRelationship) RawJSON() string { return r.JSON.raw }
func (r *AccountCreateResponseRelationship) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AccountCreateResponseUnion contains all possible properties and values from
// [string], [float64], [bool], [[]AccountCreateResponseArrayItemUnion],
// [map[string]AccountCreateResponseMapItemUnion].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAccountCreateResponseArray OfAnyArray
// OfAccountCreateResponseMapItemMapItem]
type AccountCreateResponseUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a
	// [[]AccountCreateResponseArrayItemUnion] instead of an object.
	OfAccountCreateResponseArray []AccountCreateResponseArrayItemUnion `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	// This field will be present if the value is a [any] instead of an object.
	OfAccountCreateResponseMapItemMapItem any `json:",inline"`
	JSON                                  struct {
		OfString                              respjson.Field
		OfFloat                               respjson.Field
		OfBool                                respjson.Field
		OfAccountCreateResponseArray          respjson.Field
		OfAnyArray                            respjson.Field
		OfAccountCreateResponseMapItemMapItem respjson.Field
		raw                                   string
	} `json:"-"`
}

func (u AccountCreateResponseUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AccountCreateResponseUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AccountCreateResponseUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AccountCreateResponseUnion) AsAccountCreateResponseArray() (v []AccountCreateResponseArrayItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AccountCreateResponseUnion) AsAccountCreateResponseMapMap() (v map[string]AccountCreateResponseMapItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AccountCreateResponseUnion) RawJSON() string { return u.JSON.raw }

func (r *AccountCreateResponseUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AccountCreateResponseArrayItemUnion contains all possible properties and values
// from [string], [float64], [bool], [[]any], [map[string]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray
// OfAccountCreateResponseArrayItemMapItem]
type AccountCreateResponseArrayItemUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	// This field will be present if the value is a [any] instead of an object.
	OfAccountCreateResponseArrayItemMapItem any `json:",inline"`
	JSON                                    struct {
		OfString                                respjson.Field
		OfFloat                                 respjson.Field
		OfBool                                  respjson.Field
		OfAnyArray                              respjson.Field
		OfAccountCreateResponseArrayItemMapItem respjson.Field
		raw                                     string
	} `json:"-"`
}

func (u AccountCreateResponseArrayItemUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AccountCreateResponseArrayItemUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AccountCreateResponseArrayItemUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AccountCreateResponseArrayItemUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AccountCreateResponseArrayItemUnion) AsAnyMap() (v map[string]any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AccountCreateResponseArrayItemUnion) RawJSON() string { return u.JSON.raw }

func (r *AccountCreateResponseArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AccountCreateResponseMapItemUnion contains all possible properties and values
// from [string], [float64], [bool], [[]any], [map[string]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray
// OfAccountCreateResponseMapItemMapItem]
type AccountCreateResponseMapItemUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	// This field will be present if the value is a [any] instead of an object.
	OfAccountCreateResponseMapItemMapItem any `json:",inline"`
	JSON                                  struct {
		OfString                              respjson.Field
		OfFloat                               respjson.Field
		OfBool                                respjson.Field
		OfAnyArray                            respjson.Field
		OfAccountCreateResponseMapItemMapItem respjson.Field
		raw                                   string
	} `json:"-"`
}

func (u AccountCreateResponseMapItemUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AccountCreateResponseMapItemUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AccountCreateResponseMapItemUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AccountCreateResponseMapItemUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AccountCreateResponseMapItemUnion) AsAnyMap() (v map[string]any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AccountCreateResponseMapItemUnion) RawJSON() string { return u.JSON.raw }

func (r *AccountCreateResponseMapItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDefinitionsResponse struct {
	// Map of field keys to their definitions, including both system and custom fields.
	FieldDefinitions map[string]AccountDefinitionsResponseFieldDefinition `json:"fieldDefinitions" api:"required"`
	// The object type these definitions belong to (e.g. `account`).
	ObjectType string `json:"objectType" api:"required"`
	// Map of relationship keys to their definitions.
	RelationshipDefinitions map[string]AccountDefinitionsResponseRelationshipDefinition `json:"relationshipDefinitions" api:"required"`
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
func (r AccountDefinitionsResponse) RawJSON() string { return r.JSON.raw }
func (r *AccountDefinitionsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDefinitionsResponseFieldDefinition struct {
	// Description of the field, or null.
	Description string `json:"description" api:"required"`
	// Human-readable display name of the field.
	Label string `json:"label" api:"required"`
	// Type-specific configuration (e.g. select options, currency code).
	TypeConfiguration map[string]AccountDefinitionsResponseFieldDefinitionTypeConfigurationUnion `json:"typeConfiguration" api:"required"`
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
func (r AccountDefinitionsResponseFieldDefinition) RawJSON() string { return r.JSON.raw }
func (r *AccountDefinitionsResponseFieldDefinition) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AccountDefinitionsResponseFieldDefinitionTypeConfigurationUnion contains all
// possible properties and values from [string], [float64], [bool],
// [[]AccountDefinitionsResponseFieldDefinitionTypeConfigurationArrayItemUnion],
// [map[string]AccountDefinitionsResponseFieldDefinitionTypeConfigurationMapItemUnion].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool
// OfAccountDefinitionsResponseFieldDefinitionTypeConfigurationArray OfAnyArray
// OfAccountDefinitionsResponseFieldDefinitionTypeConfigurationMapItemMapItem]
type AccountDefinitionsResponseFieldDefinitionTypeConfigurationUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a
	// [[]AccountDefinitionsResponseFieldDefinitionTypeConfigurationArrayItemUnion]
	// instead of an object.
	OfAccountDefinitionsResponseFieldDefinitionTypeConfigurationArray []AccountDefinitionsResponseFieldDefinitionTypeConfigurationArrayItemUnion `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	// This field will be present if the value is a [any] instead of an object.
	OfAccountDefinitionsResponseFieldDefinitionTypeConfigurationMapItemMapItem any `json:",inline"`
	JSON                                                                       struct {
		OfString                                                                   respjson.Field
		OfFloat                                                                    respjson.Field
		OfBool                                                                     respjson.Field
		OfAccountDefinitionsResponseFieldDefinitionTypeConfigurationArray          respjson.Field
		OfAnyArray                                                                 respjson.Field
		OfAccountDefinitionsResponseFieldDefinitionTypeConfigurationMapItemMapItem respjson.Field
		raw                                                                        string
	} `json:"-"`
}

func (u AccountDefinitionsResponseFieldDefinitionTypeConfigurationUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AccountDefinitionsResponseFieldDefinitionTypeConfigurationUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AccountDefinitionsResponseFieldDefinitionTypeConfigurationUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AccountDefinitionsResponseFieldDefinitionTypeConfigurationUnion) AsAccountDefinitionsResponseFieldDefinitionTypeConfigurationArray() (v []AccountDefinitionsResponseFieldDefinitionTypeConfigurationArrayItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AccountDefinitionsResponseFieldDefinitionTypeConfigurationUnion) AsAccountDefinitionsResponseFieldDefinitionTypeConfigurationMapMap() (v map[string]AccountDefinitionsResponseFieldDefinitionTypeConfigurationMapItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AccountDefinitionsResponseFieldDefinitionTypeConfigurationUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *AccountDefinitionsResponseFieldDefinitionTypeConfigurationUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AccountDefinitionsResponseFieldDefinitionTypeConfigurationArrayItemUnion
// contains all possible properties and values from [string], [float64], [bool],
// [[]any], [map[string]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray
// OfAccountDefinitionsResponseFieldDefinitionTypeConfigurationArrayItemMapItem]
type AccountDefinitionsResponseFieldDefinitionTypeConfigurationArrayItemUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	// This field will be present if the value is a [any] instead of an object.
	OfAccountDefinitionsResponseFieldDefinitionTypeConfigurationArrayItemMapItem any `json:",inline"`
	JSON                                                                         struct {
		OfString                                                                     respjson.Field
		OfFloat                                                                      respjson.Field
		OfBool                                                                       respjson.Field
		OfAnyArray                                                                   respjson.Field
		OfAccountDefinitionsResponseFieldDefinitionTypeConfigurationArrayItemMapItem respjson.Field
		raw                                                                          string
	} `json:"-"`
}

func (u AccountDefinitionsResponseFieldDefinitionTypeConfigurationArrayItemUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AccountDefinitionsResponseFieldDefinitionTypeConfigurationArrayItemUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AccountDefinitionsResponseFieldDefinitionTypeConfigurationArrayItemUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AccountDefinitionsResponseFieldDefinitionTypeConfigurationArrayItemUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AccountDefinitionsResponseFieldDefinitionTypeConfigurationArrayItemUnion) AsAnyMap() (v map[string]any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AccountDefinitionsResponseFieldDefinitionTypeConfigurationArrayItemUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *AccountDefinitionsResponseFieldDefinitionTypeConfigurationArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AccountDefinitionsResponseFieldDefinitionTypeConfigurationMapItemUnion contains
// all possible properties and values from [string], [float64], [bool], [[]any],
// [map[string]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray
// OfAccountDefinitionsResponseFieldDefinitionTypeConfigurationMapItemMapItem]
type AccountDefinitionsResponseFieldDefinitionTypeConfigurationMapItemUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	// This field will be present if the value is a [any] instead of an object.
	OfAccountDefinitionsResponseFieldDefinitionTypeConfigurationMapItemMapItem any `json:",inline"`
	JSON                                                                       struct {
		OfString                                                                   respjson.Field
		OfFloat                                                                    respjson.Field
		OfBool                                                                     respjson.Field
		OfAnyArray                                                                 respjson.Field
		OfAccountDefinitionsResponseFieldDefinitionTypeConfigurationMapItemMapItem respjson.Field
		raw                                                                        string
	} `json:"-"`
}

func (u AccountDefinitionsResponseFieldDefinitionTypeConfigurationMapItemUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AccountDefinitionsResponseFieldDefinitionTypeConfigurationMapItemUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AccountDefinitionsResponseFieldDefinitionTypeConfigurationMapItemUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AccountDefinitionsResponseFieldDefinitionTypeConfigurationMapItemUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AccountDefinitionsResponseFieldDefinitionTypeConfigurationMapItemUnion) AsAnyMap() (v map[string]any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AccountDefinitionsResponseFieldDefinitionTypeConfigurationMapItemUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *AccountDefinitionsResponseFieldDefinitionTypeConfigurationMapItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDefinitionsResponseRelationshipDefinition struct {
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
func (r AccountDefinitionsResponseRelationshipDefinition) RawJSON() string { return r.JSON.raw }
func (r *AccountDefinitionsResponseRelationshipDefinition) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AccountListResponse struct {
	// Array of entity objects for the current page.
	Data []AccountListResponseData `json:"data" api:"required"`
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
func (r AccountListResponse) RawJSON() string { return r.JSON.raw }
func (r *AccountListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AccountListResponseData struct {
	// Unique identifier for the entity.
	ID string `json:"id" api:"required"`
	// ISO 8601 timestamp of when the entity was created.
	CreatedAt string `json:"createdAt" api:"required"`
	// Map of field names to their typed values. System fields are prefixed with `$`
	// (e.g. `$name`, `$email`); custom attributes use their bare slug.
	Fields map[string]AccountListResponseDataField `json:"fields" api:"required"`
	// URL to view the entity in the Lightfield web app, or null.
	HTTPLink string `json:"httpLink" api:"required"`
	// Map of relationship names to their associated entities. System relationships are
	// prefixed with `$` (e.g. `$owner`, `$contact`).
	Relationships map[string]AccountListResponseDataRelationship `json:"relationships" api:"required"`
	ExtraFields   map[string]AccountListResponseDataUnion        `json:"" api:"extrafields"`
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
func (r AccountListResponseData) RawJSON() string { return r.JSON.raw }
func (r *AccountListResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AccountListResponseDataField struct {
	// The field value, or null if unset.
	Value AccountListResponseDataFieldValueUnion `json:"value" api:"required"`
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
func (r AccountListResponseDataField) RawJSON() string { return r.JSON.raw }
func (r *AccountListResponseDataField) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AccountListResponseDataFieldValueUnion contains all possible properties and
// values from [string], [float64], [bool],
// [[]AccountListResponseDataFieldValueArrayItemUnion],
// [map[string]AccountListResponseDataFieldValueMapItemUnion].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAccountListResponseDataFieldValueArray
// OfAnyArray OfAccountListResponseDataFieldValueMapItemMapItem]
type AccountListResponseDataFieldValueUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a
	// [[]AccountListResponseDataFieldValueArrayItemUnion] instead of an object.
	OfAccountListResponseDataFieldValueArray []AccountListResponseDataFieldValueArrayItemUnion `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	// This field will be present if the value is a [any] instead of an object.
	OfAccountListResponseDataFieldValueMapItemMapItem any `json:",inline"`
	JSON                                              struct {
		OfString                                          respjson.Field
		OfFloat                                           respjson.Field
		OfBool                                            respjson.Field
		OfAccountListResponseDataFieldValueArray          respjson.Field
		OfAnyArray                                        respjson.Field
		OfAccountListResponseDataFieldValueMapItemMapItem respjson.Field
		raw                                               string
	} `json:"-"`
}

func (u AccountListResponseDataFieldValueUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AccountListResponseDataFieldValueUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AccountListResponseDataFieldValueUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AccountListResponseDataFieldValueUnion) AsAccountListResponseDataFieldValueArray() (v []AccountListResponseDataFieldValueArrayItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AccountListResponseDataFieldValueUnion) AsAccountListResponseDataFieldValueMapMap() (v map[string]AccountListResponseDataFieldValueMapItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AccountListResponseDataFieldValueUnion) RawJSON() string { return u.JSON.raw }

func (r *AccountListResponseDataFieldValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AccountListResponseDataFieldValueArrayItemUnion contains all possible properties
// and values from [string], [float64], [bool], [[]any], [map[string]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray
// OfAccountListResponseDataFieldValueArrayItemMapItem]
type AccountListResponseDataFieldValueArrayItemUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	// This field will be present if the value is a [any] instead of an object.
	OfAccountListResponseDataFieldValueArrayItemMapItem any `json:",inline"`
	JSON                                                struct {
		OfString                                            respjson.Field
		OfFloat                                             respjson.Field
		OfBool                                              respjson.Field
		OfAnyArray                                          respjson.Field
		OfAccountListResponseDataFieldValueArrayItemMapItem respjson.Field
		raw                                                 string
	} `json:"-"`
}

func (u AccountListResponseDataFieldValueArrayItemUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AccountListResponseDataFieldValueArrayItemUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AccountListResponseDataFieldValueArrayItemUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AccountListResponseDataFieldValueArrayItemUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AccountListResponseDataFieldValueArrayItemUnion) AsAnyMap() (v map[string]any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AccountListResponseDataFieldValueArrayItemUnion) RawJSON() string { return u.JSON.raw }

func (r *AccountListResponseDataFieldValueArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AccountListResponseDataFieldValueMapItemUnion contains all possible properties
// and values from [string], [float64], [bool], [[]any], [map[string]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray
// OfAccountListResponseDataFieldValueMapItemMapItem]
type AccountListResponseDataFieldValueMapItemUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	// This field will be present if the value is a [any] instead of an object.
	OfAccountListResponseDataFieldValueMapItemMapItem any `json:",inline"`
	JSON                                              struct {
		OfString                                          respjson.Field
		OfFloat                                           respjson.Field
		OfBool                                            respjson.Field
		OfAnyArray                                        respjson.Field
		OfAccountListResponseDataFieldValueMapItemMapItem respjson.Field
		raw                                               string
	} `json:"-"`
}

func (u AccountListResponseDataFieldValueMapItemUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AccountListResponseDataFieldValueMapItemUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AccountListResponseDataFieldValueMapItemUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AccountListResponseDataFieldValueMapItemUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AccountListResponseDataFieldValueMapItemUnion) AsAnyMap() (v map[string]any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AccountListResponseDataFieldValueMapItemUnion) RawJSON() string { return u.JSON.raw }

func (r *AccountListResponseDataFieldValueMapItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AccountListResponseDataRelationship struct {
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
func (r AccountListResponseDataRelationship) RawJSON() string { return r.JSON.raw }
func (r *AccountListResponseDataRelationship) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AccountListResponseDataUnion contains all possible properties and values from
// [string], [float64], [bool], [[]AccountListResponseDataArrayItemUnion],
// [map[string]AccountListResponseDataMapItemUnion].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAccountListResponseDataArray OfAnyArray
// OfAccountListResponseDataMapItemMapItem]
type AccountListResponseDataUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a
	// [[]AccountListResponseDataArrayItemUnion] instead of an object.
	OfAccountListResponseDataArray []AccountListResponseDataArrayItemUnion `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	// This field will be present if the value is a [any] instead of an object.
	OfAccountListResponseDataMapItemMapItem any `json:",inline"`
	JSON                                    struct {
		OfString                                respjson.Field
		OfFloat                                 respjson.Field
		OfBool                                  respjson.Field
		OfAccountListResponseDataArray          respjson.Field
		OfAnyArray                              respjson.Field
		OfAccountListResponseDataMapItemMapItem respjson.Field
		raw                                     string
	} `json:"-"`
}

func (u AccountListResponseDataUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AccountListResponseDataUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AccountListResponseDataUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AccountListResponseDataUnion) AsAccountListResponseDataArray() (v []AccountListResponseDataArrayItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AccountListResponseDataUnion) AsAccountListResponseDataMapMap() (v map[string]AccountListResponseDataMapItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AccountListResponseDataUnion) RawJSON() string { return u.JSON.raw }

func (r *AccountListResponseDataUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AccountListResponseDataArrayItemUnion contains all possible properties and
// values from [string], [float64], [bool], [[]any], [map[string]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray
// OfAccountListResponseDataArrayItemMapItem]
type AccountListResponseDataArrayItemUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	// This field will be present if the value is a [any] instead of an object.
	OfAccountListResponseDataArrayItemMapItem any `json:",inline"`
	JSON                                      struct {
		OfString                                  respjson.Field
		OfFloat                                   respjson.Field
		OfBool                                    respjson.Field
		OfAnyArray                                respjson.Field
		OfAccountListResponseDataArrayItemMapItem respjson.Field
		raw                                       string
	} `json:"-"`
}

func (u AccountListResponseDataArrayItemUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AccountListResponseDataArrayItemUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AccountListResponseDataArrayItemUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AccountListResponseDataArrayItemUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AccountListResponseDataArrayItemUnion) AsAnyMap() (v map[string]any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AccountListResponseDataArrayItemUnion) RawJSON() string { return u.JSON.raw }

func (r *AccountListResponseDataArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AccountListResponseDataMapItemUnion contains all possible properties and values
// from [string], [float64], [bool], [[]any], [map[string]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray
// OfAccountListResponseDataMapItemMapItem]
type AccountListResponseDataMapItemUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	// This field will be present if the value is a [any] instead of an object.
	OfAccountListResponseDataMapItemMapItem any `json:",inline"`
	JSON                                    struct {
		OfString                                respjson.Field
		OfFloat                                 respjson.Field
		OfBool                                  respjson.Field
		OfAnyArray                              respjson.Field
		OfAccountListResponseDataMapItemMapItem respjson.Field
		raw                                     string
	} `json:"-"`
}

func (u AccountListResponseDataMapItemUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AccountListResponseDataMapItemUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AccountListResponseDataMapItemUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AccountListResponseDataMapItemUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AccountListResponseDataMapItemUnion) AsAnyMap() (v map[string]any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AccountListResponseDataMapItemUnion) RawJSON() string { return u.JSON.raw }

func (r *AccountListResponseDataMapItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AccountRetrieveResponse struct {
	// Unique identifier for the entity.
	ID string `json:"id" api:"required"`
	// ISO 8601 timestamp of when the entity was created.
	CreatedAt string `json:"createdAt" api:"required"`
	// Map of field names to their typed values. System fields are prefixed with `$`
	// (e.g. `$name`, `$email`); custom attributes use their bare slug.
	Fields map[string]AccountRetrieveResponseField `json:"fields" api:"required"`
	// URL to view the entity in the Lightfield web app, or null.
	HTTPLink string `json:"httpLink" api:"required"`
	// Map of relationship names to their associated entities. System relationships are
	// prefixed with `$` (e.g. `$owner`, `$contact`).
	Relationships map[string]AccountRetrieveResponseRelationship `json:"relationships" api:"required"`
	ExtraFields   map[string]AccountRetrieveResponseUnion        `json:"" api:"extrafields"`
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
func (r AccountRetrieveResponse) RawJSON() string { return r.JSON.raw }
func (r *AccountRetrieveResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AccountRetrieveResponseField struct {
	// The field value, or null if unset.
	Value AccountRetrieveResponseFieldValueUnion `json:"value" api:"required"`
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
func (r AccountRetrieveResponseField) RawJSON() string { return r.JSON.raw }
func (r *AccountRetrieveResponseField) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AccountRetrieveResponseFieldValueUnion contains all possible properties and
// values from [string], [float64], [bool],
// [[]AccountRetrieveResponseFieldValueArrayItemUnion],
// [map[string]AccountRetrieveResponseFieldValueMapItemUnion].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAccountRetrieveResponseFieldValueArray
// OfAnyArray OfAccountRetrieveResponseFieldValueMapItemMapItem]
type AccountRetrieveResponseFieldValueUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a
	// [[]AccountRetrieveResponseFieldValueArrayItemUnion] instead of an object.
	OfAccountRetrieveResponseFieldValueArray []AccountRetrieveResponseFieldValueArrayItemUnion `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	// This field will be present if the value is a [any] instead of an object.
	OfAccountRetrieveResponseFieldValueMapItemMapItem any `json:",inline"`
	JSON                                              struct {
		OfString                                          respjson.Field
		OfFloat                                           respjson.Field
		OfBool                                            respjson.Field
		OfAccountRetrieveResponseFieldValueArray          respjson.Field
		OfAnyArray                                        respjson.Field
		OfAccountRetrieveResponseFieldValueMapItemMapItem respjson.Field
		raw                                               string
	} `json:"-"`
}

func (u AccountRetrieveResponseFieldValueUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AccountRetrieveResponseFieldValueUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AccountRetrieveResponseFieldValueUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AccountRetrieveResponseFieldValueUnion) AsAccountRetrieveResponseFieldValueArray() (v []AccountRetrieveResponseFieldValueArrayItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AccountRetrieveResponseFieldValueUnion) AsAccountRetrieveResponseFieldValueMapMap() (v map[string]AccountRetrieveResponseFieldValueMapItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AccountRetrieveResponseFieldValueUnion) RawJSON() string { return u.JSON.raw }

func (r *AccountRetrieveResponseFieldValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AccountRetrieveResponseFieldValueArrayItemUnion contains all possible properties
// and values from [string], [float64], [bool], [[]any], [map[string]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray
// OfAccountRetrieveResponseFieldValueArrayItemMapItem]
type AccountRetrieveResponseFieldValueArrayItemUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	// This field will be present if the value is a [any] instead of an object.
	OfAccountRetrieveResponseFieldValueArrayItemMapItem any `json:",inline"`
	JSON                                                struct {
		OfString                                            respjson.Field
		OfFloat                                             respjson.Field
		OfBool                                              respjson.Field
		OfAnyArray                                          respjson.Field
		OfAccountRetrieveResponseFieldValueArrayItemMapItem respjson.Field
		raw                                                 string
	} `json:"-"`
}

func (u AccountRetrieveResponseFieldValueArrayItemUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AccountRetrieveResponseFieldValueArrayItemUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AccountRetrieveResponseFieldValueArrayItemUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AccountRetrieveResponseFieldValueArrayItemUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AccountRetrieveResponseFieldValueArrayItemUnion) AsAnyMap() (v map[string]any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AccountRetrieveResponseFieldValueArrayItemUnion) RawJSON() string { return u.JSON.raw }

func (r *AccountRetrieveResponseFieldValueArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AccountRetrieveResponseFieldValueMapItemUnion contains all possible properties
// and values from [string], [float64], [bool], [[]any], [map[string]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray
// OfAccountRetrieveResponseFieldValueMapItemMapItem]
type AccountRetrieveResponseFieldValueMapItemUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	// This field will be present if the value is a [any] instead of an object.
	OfAccountRetrieveResponseFieldValueMapItemMapItem any `json:",inline"`
	JSON                                              struct {
		OfString                                          respjson.Field
		OfFloat                                           respjson.Field
		OfBool                                            respjson.Field
		OfAnyArray                                        respjson.Field
		OfAccountRetrieveResponseFieldValueMapItemMapItem respjson.Field
		raw                                               string
	} `json:"-"`
}

func (u AccountRetrieveResponseFieldValueMapItemUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AccountRetrieveResponseFieldValueMapItemUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AccountRetrieveResponseFieldValueMapItemUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AccountRetrieveResponseFieldValueMapItemUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AccountRetrieveResponseFieldValueMapItemUnion) AsAnyMap() (v map[string]any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AccountRetrieveResponseFieldValueMapItemUnion) RawJSON() string { return u.JSON.raw }

func (r *AccountRetrieveResponseFieldValueMapItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AccountRetrieveResponseRelationship struct {
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
func (r AccountRetrieveResponseRelationship) RawJSON() string { return r.JSON.raw }
func (r *AccountRetrieveResponseRelationship) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AccountRetrieveResponseUnion contains all possible properties and values from
// [string], [float64], [bool], [[]AccountRetrieveResponseArrayItemUnion],
// [map[string]AccountRetrieveResponseMapItemUnion].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAccountRetrieveResponseArray OfAnyArray
// OfAccountRetrieveResponseMapItemMapItem]
type AccountRetrieveResponseUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a
	// [[]AccountRetrieveResponseArrayItemUnion] instead of an object.
	OfAccountRetrieveResponseArray []AccountRetrieveResponseArrayItemUnion `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	// This field will be present if the value is a [any] instead of an object.
	OfAccountRetrieveResponseMapItemMapItem any `json:",inline"`
	JSON                                    struct {
		OfString                                respjson.Field
		OfFloat                                 respjson.Field
		OfBool                                  respjson.Field
		OfAccountRetrieveResponseArray          respjson.Field
		OfAnyArray                              respjson.Field
		OfAccountRetrieveResponseMapItemMapItem respjson.Field
		raw                                     string
	} `json:"-"`
}

func (u AccountRetrieveResponseUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AccountRetrieveResponseUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AccountRetrieveResponseUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AccountRetrieveResponseUnion) AsAccountRetrieveResponseArray() (v []AccountRetrieveResponseArrayItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AccountRetrieveResponseUnion) AsAccountRetrieveResponseMapMap() (v map[string]AccountRetrieveResponseMapItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AccountRetrieveResponseUnion) RawJSON() string { return u.JSON.raw }

func (r *AccountRetrieveResponseUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AccountRetrieveResponseArrayItemUnion contains all possible properties and
// values from [string], [float64], [bool], [[]any], [map[string]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray
// OfAccountRetrieveResponseArrayItemMapItem]
type AccountRetrieveResponseArrayItemUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	// This field will be present if the value is a [any] instead of an object.
	OfAccountRetrieveResponseArrayItemMapItem any `json:",inline"`
	JSON                                      struct {
		OfString                                  respjson.Field
		OfFloat                                   respjson.Field
		OfBool                                    respjson.Field
		OfAnyArray                                respjson.Field
		OfAccountRetrieveResponseArrayItemMapItem respjson.Field
		raw                                       string
	} `json:"-"`
}

func (u AccountRetrieveResponseArrayItemUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AccountRetrieveResponseArrayItemUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AccountRetrieveResponseArrayItemUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AccountRetrieveResponseArrayItemUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AccountRetrieveResponseArrayItemUnion) AsAnyMap() (v map[string]any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AccountRetrieveResponseArrayItemUnion) RawJSON() string { return u.JSON.raw }

func (r *AccountRetrieveResponseArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AccountRetrieveResponseMapItemUnion contains all possible properties and values
// from [string], [float64], [bool], [[]any], [map[string]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray
// OfAccountRetrieveResponseMapItemMapItem]
type AccountRetrieveResponseMapItemUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	// This field will be present if the value is a [any] instead of an object.
	OfAccountRetrieveResponseMapItemMapItem any `json:",inline"`
	JSON                                    struct {
		OfString                                respjson.Field
		OfFloat                                 respjson.Field
		OfBool                                  respjson.Field
		OfAnyArray                              respjson.Field
		OfAccountRetrieveResponseMapItemMapItem respjson.Field
		raw                                     string
	} `json:"-"`
}

func (u AccountRetrieveResponseMapItemUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AccountRetrieveResponseMapItemUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AccountRetrieveResponseMapItemUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AccountRetrieveResponseMapItemUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AccountRetrieveResponseMapItemUnion) AsAnyMap() (v map[string]any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AccountRetrieveResponseMapItemUnion) RawJSON() string { return u.JSON.raw }

func (r *AccountRetrieveResponseMapItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AccountUpdateResponse struct {
	// Unique identifier for the entity.
	ID string `json:"id" api:"required"`
	// ISO 8601 timestamp of when the entity was created.
	CreatedAt string `json:"createdAt" api:"required"`
	// Map of field names to their typed values. System fields are prefixed with `$`
	// (e.g. `$name`, `$email`); custom attributes use their bare slug.
	Fields map[string]AccountUpdateResponseField `json:"fields" api:"required"`
	// URL to view the entity in the Lightfield web app, or null.
	HTTPLink string `json:"httpLink" api:"required"`
	// Map of relationship names to their associated entities. System relationships are
	// prefixed with `$` (e.g. `$owner`, `$contact`).
	Relationships map[string]AccountUpdateResponseRelationship `json:"relationships" api:"required"`
	ExtraFields   map[string]AccountUpdateResponseUnion        `json:"" api:"extrafields"`
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
func (r AccountUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *AccountUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AccountUpdateResponseField struct {
	// The field value, or null if unset.
	Value AccountUpdateResponseFieldValueUnion `json:"value" api:"required"`
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
func (r AccountUpdateResponseField) RawJSON() string { return r.JSON.raw }
func (r *AccountUpdateResponseField) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AccountUpdateResponseFieldValueUnion contains all possible properties and values
// from [string], [float64], [bool],
// [[]AccountUpdateResponseFieldValueArrayItemUnion],
// [map[string]AccountUpdateResponseFieldValueMapItemUnion].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAccountUpdateResponseFieldValueArray
// OfAnyArray OfAccountUpdateResponseFieldValueMapItemMapItem]
type AccountUpdateResponseFieldValueUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a
	// [[]AccountUpdateResponseFieldValueArrayItemUnion] instead of an object.
	OfAccountUpdateResponseFieldValueArray []AccountUpdateResponseFieldValueArrayItemUnion `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	// This field will be present if the value is a [any] instead of an object.
	OfAccountUpdateResponseFieldValueMapItemMapItem any `json:",inline"`
	JSON                                            struct {
		OfString                                        respjson.Field
		OfFloat                                         respjson.Field
		OfBool                                          respjson.Field
		OfAccountUpdateResponseFieldValueArray          respjson.Field
		OfAnyArray                                      respjson.Field
		OfAccountUpdateResponseFieldValueMapItemMapItem respjson.Field
		raw                                             string
	} `json:"-"`
}

func (u AccountUpdateResponseFieldValueUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AccountUpdateResponseFieldValueUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AccountUpdateResponseFieldValueUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AccountUpdateResponseFieldValueUnion) AsAccountUpdateResponseFieldValueArray() (v []AccountUpdateResponseFieldValueArrayItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AccountUpdateResponseFieldValueUnion) AsAccountUpdateResponseFieldValueMapMap() (v map[string]AccountUpdateResponseFieldValueMapItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AccountUpdateResponseFieldValueUnion) RawJSON() string { return u.JSON.raw }

func (r *AccountUpdateResponseFieldValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AccountUpdateResponseFieldValueArrayItemUnion contains all possible properties
// and values from [string], [float64], [bool], [[]any], [map[string]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray
// OfAccountUpdateResponseFieldValueArrayItemMapItem]
type AccountUpdateResponseFieldValueArrayItemUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	// This field will be present if the value is a [any] instead of an object.
	OfAccountUpdateResponseFieldValueArrayItemMapItem any `json:",inline"`
	JSON                                              struct {
		OfString                                          respjson.Field
		OfFloat                                           respjson.Field
		OfBool                                            respjson.Field
		OfAnyArray                                        respjson.Field
		OfAccountUpdateResponseFieldValueArrayItemMapItem respjson.Field
		raw                                               string
	} `json:"-"`
}

func (u AccountUpdateResponseFieldValueArrayItemUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AccountUpdateResponseFieldValueArrayItemUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AccountUpdateResponseFieldValueArrayItemUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AccountUpdateResponseFieldValueArrayItemUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AccountUpdateResponseFieldValueArrayItemUnion) AsAnyMap() (v map[string]any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AccountUpdateResponseFieldValueArrayItemUnion) RawJSON() string { return u.JSON.raw }

func (r *AccountUpdateResponseFieldValueArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AccountUpdateResponseFieldValueMapItemUnion contains all possible properties and
// values from [string], [float64], [bool], [[]any], [map[string]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray
// OfAccountUpdateResponseFieldValueMapItemMapItem]
type AccountUpdateResponseFieldValueMapItemUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	// This field will be present if the value is a [any] instead of an object.
	OfAccountUpdateResponseFieldValueMapItemMapItem any `json:",inline"`
	JSON                                            struct {
		OfString                                        respjson.Field
		OfFloat                                         respjson.Field
		OfBool                                          respjson.Field
		OfAnyArray                                      respjson.Field
		OfAccountUpdateResponseFieldValueMapItemMapItem respjson.Field
		raw                                             string
	} `json:"-"`
}

func (u AccountUpdateResponseFieldValueMapItemUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AccountUpdateResponseFieldValueMapItemUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AccountUpdateResponseFieldValueMapItemUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AccountUpdateResponseFieldValueMapItemUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AccountUpdateResponseFieldValueMapItemUnion) AsAnyMap() (v map[string]any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AccountUpdateResponseFieldValueMapItemUnion) RawJSON() string { return u.JSON.raw }

func (r *AccountUpdateResponseFieldValueMapItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AccountUpdateResponseRelationship struct {
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
func (r AccountUpdateResponseRelationship) RawJSON() string { return r.JSON.raw }
func (r *AccountUpdateResponseRelationship) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AccountUpdateResponseUnion contains all possible properties and values from
// [string], [float64], [bool], [[]AccountUpdateResponseArrayItemUnion],
// [map[string]AccountUpdateResponseMapItemUnion].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAccountUpdateResponseArray OfAnyArray
// OfAccountUpdateResponseMapItemMapItem]
type AccountUpdateResponseUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a
	// [[]AccountUpdateResponseArrayItemUnion] instead of an object.
	OfAccountUpdateResponseArray []AccountUpdateResponseArrayItemUnion `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	// This field will be present if the value is a [any] instead of an object.
	OfAccountUpdateResponseMapItemMapItem any `json:",inline"`
	JSON                                  struct {
		OfString                              respjson.Field
		OfFloat                               respjson.Field
		OfBool                                respjson.Field
		OfAccountUpdateResponseArray          respjson.Field
		OfAnyArray                            respjson.Field
		OfAccountUpdateResponseMapItemMapItem respjson.Field
		raw                                   string
	} `json:"-"`
}

func (u AccountUpdateResponseUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AccountUpdateResponseUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AccountUpdateResponseUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AccountUpdateResponseUnion) AsAccountUpdateResponseArray() (v []AccountUpdateResponseArrayItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AccountUpdateResponseUnion) AsAccountUpdateResponseMapMap() (v map[string]AccountUpdateResponseMapItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AccountUpdateResponseUnion) RawJSON() string { return u.JSON.raw }

func (r *AccountUpdateResponseUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AccountUpdateResponseArrayItemUnion contains all possible properties and values
// from [string], [float64], [bool], [[]any], [map[string]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray
// OfAccountUpdateResponseArrayItemMapItem]
type AccountUpdateResponseArrayItemUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	// This field will be present if the value is a [any] instead of an object.
	OfAccountUpdateResponseArrayItemMapItem any `json:",inline"`
	JSON                                    struct {
		OfString                                respjson.Field
		OfFloat                                 respjson.Field
		OfBool                                  respjson.Field
		OfAnyArray                              respjson.Field
		OfAccountUpdateResponseArrayItemMapItem respjson.Field
		raw                                     string
	} `json:"-"`
}

func (u AccountUpdateResponseArrayItemUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AccountUpdateResponseArrayItemUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AccountUpdateResponseArrayItemUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AccountUpdateResponseArrayItemUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AccountUpdateResponseArrayItemUnion) AsAnyMap() (v map[string]any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AccountUpdateResponseArrayItemUnion) RawJSON() string { return u.JSON.raw }

func (r *AccountUpdateResponseArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AccountUpdateResponseMapItemUnion contains all possible properties and values
// from [string], [float64], [bool], [[]any], [map[string]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray
// OfAccountUpdateResponseMapItemMapItem]
type AccountUpdateResponseMapItemUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	// This field will be present if the value is a [any] instead of an object.
	OfAccountUpdateResponseMapItemMapItem any `json:",inline"`
	JSON                                  struct {
		OfString                              respjson.Field
		OfFloat                               respjson.Field
		OfBool                                respjson.Field
		OfAnyArray                            respjson.Field
		OfAccountUpdateResponseMapItemMapItem respjson.Field
		raw                                   string
	} `json:"-"`
}

func (u AccountUpdateResponseMapItemUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AccountUpdateResponseMapItemUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AccountUpdateResponseMapItemUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AccountUpdateResponseMapItemUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AccountUpdateResponseMapItemUnion) AsAnyMap() (v map[string]any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AccountUpdateResponseMapItemUnion) RawJSON() string { return u.JSON.raw }

func (r *AccountUpdateResponseMapItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AccountNewParams struct {
	// Field values for the new account. System fields use a `$` prefix (e.g. `$name`,
	// `$website`); custom attributes use their bare slug (e.g. `tier`, `renewalDate`).
	// Required: `$name` (string). Fields of type `SINGLE_SELECT` or `MULTI_SELECT`
	// accept either an option ID or label from the field's `typeConfiguration.options`
	// — call the
	// <u>[definitions endpoint](/api/resources/account/methods/definitions)</u> to
	// discover available fields and options. See
	// <u>[Fields and relationships](/using-the-api/fields-and-relationships/)</u> for
	// value type details.
	Fields AccountNewParamsFields `json:"fields,omitzero" api:"required"`
	// Relationships to set on the new account. System relationships use a `$` prefix
	// (e.g. `$owner`, `$contact`); custom relationships use their bare slug. Each
	// value is a single entity ID or an array of IDs. Call the
	// <u>[definitions endpoint](/api/resources/account/methods/definitions)</u> to
	// list available relationship keys.
	Relationships AccountNewParamsRelationships `json:"relationships,omitzero"`
	paramObj
}

func (r AccountNewParams) MarshalJSON() (data []byte, err error) {
	type shadow AccountNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AccountNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Field values for the new account. System fields use a `$` prefix (e.g. `$name`,
// `$website`); custom attributes use their bare slug (e.g. `tier`, `renewalDate`).
// Required: `$name` (string). Fields of type `SINGLE_SELECT` or `MULTI_SELECT`
// accept either an option ID or label from the field's `typeConfiguration.options`
// — call the
// <u>[definitions endpoint](/api/resources/account/methods/definitions)</u> to
// discover available fields and options. See
// <u>[Fields and relationships](/using-the-api/fields-and-relationships/)</u> for
// value type details.
//
// The property Name is required.
type AccountNewParamsFields struct {
	// Display name of the account.
	Name string `json:"$name" api:"required"`
	// Facebook handle or profile identifier (`SOCIAL_HANDLE`).
	Facebook param.Opt[string] `json:"$facebook,omitzero"`
	// Employee count range (`SINGLE_SELECT`). Pass the option ID or label from the
	// field definition.
	Headcount param.Opt[string] `json:"$headcount,omitzero"`
	// Instagram handle or profile identifier (`SOCIAL_HANDLE`).
	Instagram param.Opt[string] `json:"$instagram,omitzero"`
	// Most recent funding round type (`SINGLE_SELECT`). Pass the option ID or label
	// from the field definition.
	LastFundingType param.Opt[string] `json:"$lastFundingType,omitzero"`
	// LinkedIn handle or profile identifier (`SOCIAL_HANDLE`).
	LinkedIn param.Opt[string] `json:"$linkedIn,omitzero"`
	// Twitter/X handle (`SOCIAL_HANDLE`).
	Twitter param.Opt[string] `json:"$twitter,omitzero"`
	// Industries the account operates in (`MULTI_SELECT`). Pass option IDs or labels
	// from the field definition.
	Industry []string `json:"$industry,omitzero"`
	// Primary address (`ADDRESS`).
	PrimaryAddress AccountNewParamsFieldsPrimaryAddress `json:"$primaryAddress,omitzero"`
	// Website URLs associated with the account (`URL`, multi-value).
	Website     []string                              `json:"$website,omitzero"`
	ExtraFields map[string]AccountNewParamsFieldUnion `json:"-"`
	paramObj
}

func (r AccountNewParamsFields) MarshalJSON() (data []byte, err error) {
	type shadow AccountNewParamsFields
	return param.MarshalWithExtras(r, (*shadow)(&r), r.ExtraFields)
}
func (r *AccountNewParamsFields) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Primary address (`ADDRESS`).
type AccountNewParamsFieldsPrimaryAddress struct {
	// City name.
	City param.Opt[string] `json:"city,omitzero"`
	// 2-letter ISO 3166-1 alpha-2 country code.
	Country param.Opt[string] `json:"country,omitzero"`
	// Latitude coordinate.
	Latitude param.Opt[float64] `json:"latitude,omitzero"`
	// Longitude coordinate.
	Longitude param.Opt[float64] `json:"longitude,omitzero"`
	// Postal or ZIP code.
	PostalCode param.Opt[string] `json:"postalCode,omitzero"`
	// State or province.
	State param.Opt[string] `json:"state,omitzero"`
	// Street address line 1.
	Street param.Opt[string] `json:"street,omitzero"`
	// Street address line 2.
	Street2 param.Opt[string] `json:"street2,omitzero"`
	paramObj
}

func (r AccountNewParamsFieldsPrimaryAddress) MarshalJSON() (data []byte, err error) {
	type shadow AccountNewParamsFieldsPrimaryAddress
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AccountNewParamsFieldsPrimaryAddress) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type AccountNewParamsFieldUnion struct {
	OfString                 param.Opt[string]                            `json:",omitzero,inline"`
	OfFloat                  param.Opt[float64]                           `json:",omitzero,inline"`
	OfBool                   param.Opt[bool]                              `json:",omitzero,inline"`
	OfAccountNewsFieldArray  []AccountNewParamsFieldArrayItemUnion        `json:",omitzero,inline"`
	OfAccountNewsFieldMapMap map[string]AccountNewParamsFieldMapItemUnion `json:",omitzero,inline"`
	paramUnion
}

func (u AccountNewParamsFieldUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString,
		u.OfFloat,
		u.OfBool,
		u.OfAccountNewsFieldArray,
		u.OfAccountNewsFieldMapMap)
}
func (u *AccountNewParamsFieldUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type AccountNewParamsFieldArrayItemUnion struct {
	OfString   param.Opt[string]  `json:",omitzero,inline"`
	OfFloat    param.Opt[float64] `json:",omitzero,inline"`
	OfBool     param.Opt[bool]    `json:",omitzero,inline"`
	OfAnyArray []any              `json:",omitzero,inline"`
	OfAnyMap   map[string]any     `json:",omitzero,inline"`
	paramUnion
}

func (u AccountNewParamsFieldArrayItemUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString,
		u.OfFloat,
		u.OfBool,
		u.OfAnyArray,
		u.OfAnyMap)
}
func (u *AccountNewParamsFieldArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type AccountNewParamsFieldMapItemUnion struct {
	OfString   param.Opt[string]  `json:",omitzero,inline"`
	OfFloat    param.Opt[float64] `json:",omitzero,inline"`
	OfBool     param.Opt[bool]    `json:",omitzero,inline"`
	OfAnyArray []any              `json:",omitzero,inline"`
	OfAnyMap   map[string]any     `json:",omitzero,inline"`
	paramUnion
}

func (u AccountNewParamsFieldMapItemUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString,
		u.OfFloat,
		u.OfBool,
		u.OfAnyArray,
		u.OfAnyMap)
}
func (u *AccountNewParamsFieldMapItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Relationships to set on the new account. System relationships use a `$` prefix
// (e.g. `$owner`, `$contact`); custom relationships use their bare slug. Each
// value is a single entity ID or an array of IDs. Call the
// <u>[definitions endpoint](/api/resources/account/methods/definitions)</u> to
// list available relationship keys.
type AccountNewParamsRelationships struct {
	// ID(s) of contacts to associate with this account.
	Contact AccountNewParamsRelationshipsContactUnion `json:"$contact,omitzero"`
	// ID of the user who owns this account.
	Owner       AccountNewParamsRelationshipsOwnerUnion      `json:"$owner,omitzero"`
	ExtraFields map[string]AccountNewParamsRelationshipUnion `json:"-"`
	paramObj
}

func (r AccountNewParamsRelationships) MarshalJSON() (data []byte, err error) {
	type shadow AccountNewParamsRelationships
	return param.MarshalWithExtras(r, (*shadow)(&r), r.ExtraFields)
}
func (r *AccountNewParamsRelationships) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type AccountNewParamsRelationshipsContactUnion struct {
	OfString      param.Opt[string] `json:",omitzero,inline"`
	OfStringArray []string          `json:",omitzero,inline"`
	paramUnion
}

func (u AccountNewParamsRelationshipsContactUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray)
}
func (u *AccountNewParamsRelationshipsContactUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type AccountNewParamsRelationshipsOwnerUnion struct {
	OfString      param.Opt[string] `json:",omitzero,inline"`
	OfStringArray []string          `json:",omitzero,inline"`
	paramUnion
}

func (u AccountNewParamsRelationshipsOwnerUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray)
}
func (u *AccountNewParamsRelationshipsOwnerUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type AccountNewParamsRelationshipUnion struct {
	OfString      param.Opt[string] `json:",omitzero,inline"`
	OfStringArray []string          `json:",omitzero,inline"`
	paramUnion
}

func (u AccountNewParamsRelationshipUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray)
}
func (u *AccountNewParamsRelationshipUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

type AccountUpdateParams struct {
	// Field values to update — only provided fields are modified; omitted fields are
	// left unchanged. System fields use a `$` prefix (e.g. `$name`); custom attributes
	// use their bare slug. `SINGLE_SELECT` and `MULTI_SELECT` fields accept an option
	// ID or label — call the
	// <u>[definitions endpoint](/api/resources/account/methods/definitions)</u> for
	// available options. See
	// <u>[Fields and relationships](/using-the-api/fields-and-relationships/)</u> for
	// value type details.
	Fields AccountUpdateParamsFields `json:"fields,omitzero"`
	// Relationship operations to apply. System relationships use a `$` prefix (e.g.
	// `$owner`, `$contact`). Each value is an operation object with `add`, `remove`,
	// or `replace`.
	Relationships AccountUpdateParamsRelationships `json:"relationships,omitzero"`
	paramObj
}

func (r AccountUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow AccountUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AccountUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Field values to update — only provided fields are modified; omitted fields are
// left unchanged. System fields use a `$` prefix (e.g. `$name`); custom attributes
// use their bare slug. `SINGLE_SELECT` and `MULTI_SELECT` fields accept an option
// ID or label — call the
// <u>[definitions endpoint](/api/resources/account/methods/definitions)</u> for
// available options. See
// <u>[Fields and relationships](/using-the-api/fields-and-relationships/)</u> for
// value type details.
type AccountUpdateParamsFields struct {
	// Facebook handle or profile identifier (`SOCIAL_HANDLE`).
	Facebook param.Opt[string] `json:"$facebook,omitzero"`
	// Employee count range (`SINGLE_SELECT`). Pass the option ID or label from the
	// field definition.
	Headcount param.Opt[string] `json:"$headcount,omitzero"`
	// Instagram handle or profile identifier (`SOCIAL_HANDLE`).
	Instagram param.Opt[string] `json:"$instagram,omitzero"`
	// Most recent funding round type (`SINGLE_SELECT`). Pass the option ID or label
	// from the field definition.
	LastFundingType param.Opt[string] `json:"$lastFundingType,omitzero"`
	// LinkedIn handle or profile identifier (`SOCIAL_HANDLE`).
	LinkedIn param.Opt[string] `json:"$linkedIn,omitzero"`
	// Display name of the account.
	Name param.Opt[string] `json:"$name,omitzero"`
	// Twitter/X handle (`SOCIAL_HANDLE`).
	Twitter param.Opt[string] `json:"$twitter,omitzero"`
	// Industries the account operates in (`MULTI_SELECT`). Pass option IDs or labels
	// from the field definition.
	Industry []string `json:"$industry,omitzero"`
	// Primary address (`ADDRESS`).
	PrimaryAddress AccountUpdateParamsFieldsPrimaryAddress `json:"$primaryAddress,omitzero"`
	// Website URLs associated with the account (`URL`, multi-value).
	Website     []string                                 `json:"$website,omitzero"`
	ExtraFields map[string]AccountUpdateParamsFieldUnion `json:"-"`
	paramObj
}

func (r AccountUpdateParamsFields) MarshalJSON() (data []byte, err error) {
	type shadow AccountUpdateParamsFields
	return param.MarshalWithExtras(r, (*shadow)(&r), r.ExtraFields)
}
func (r *AccountUpdateParamsFields) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Primary address (`ADDRESS`).
type AccountUpdateParamsFieldsPrimaryAddress struct {
	// City name.
	City param.Opt[string] `json:"city,omitzero"`
	// 2-letter ISO 3166-1 alpha-2 country code.
	Country param.Opt[string] `json:"country,omitzero"`
	// Latitude coordinate.
	Latitude param.Opt[float64] `json:"latitude,omitzero"`
	// Longitude coordinate.
	Longitude param.Opt[float64] `json:"longitude,omitzero"`
	// Postal or ZIP code.
	PostalCode param.Opt[string] `json:"postalCode,omitzero"`
	// State or province.
	State param.Opt[string] `json:"state,omitzero"`
	// Street address line 1.
	Street param.Opt[string] `json:"street,omitzero"`
	// Street address line 2.
	Street2 param.Opt[string] `json:"street2,omitzero"`
	paramObj
}

func (r AccountUpdateParamsFieldsPrimaryAddress) MarshalJSON() (data []byte, err error) {
	type shadow AccountUpdateParamsFieldsPrimaryAddress
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AccountUpdateParamsFieldsPrimaryAddress) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type AccountUpdateParamsFieldUnion struct {
	OfString                    param.Opt[string]                               `json:",omitzero,inline"`
	OfFloat                     param.Opt[float64]                              `json:",omitzero,inline"`
	OfBool                      param.Opt[bool]                                 `json:",omitzero,inline"`
	OfAccountUpdatesFieldArray  []AccountUpdateParamsFieldArrayItemUnion        `json:",omitzero,inline"`
	OfAccountUpdatesFieldMapMap map[string]AccountUpdateParamsFieldMapItemUnion `json:",omitzero,inline"`
	paramUnion
}

func (u AccountUpdateParamsFieldUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString,
		u.OfFloat,
		u.OfBool,
		u.OfAccountUpdatesFieldArray,
		u.OfAccountUpdatesFieldMapMap)
}
func (u *AccountUpdateParamsFieldUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type AccountUpdateParamsFieldArrayItemUnion struct {
	OfString   param.Opt[string]  `json:",omitzero,inline"`
	OfFloat    param.Opt[float64] `json:",omitzero,inline"`
	OfBool     param.Opt[bool]    `json:",omitzero,inline"`
	OfAnyArray []any              `json:",omitzero,inline"`
	OfAnyMap   map[string]any     `json:",omitzero,inline"`
	paramUnion
}

func (u AccountUpdateParamsFieldArrayItemUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString,
		u.OfFloat,
		u.OfBool,
		u.OfAnyArray,
		u.OfAnyMap)
}
func (u *AccountUpdateParamsFieldArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type AccountUpdateParamsFieldMapItemUnion struct {
	OfString   param.Opt[string]  `json:",omitzero,inline"`
	OfFloat    param.Opt[float64] `json:",omitzero,inline"`
	OfBool     param.Opt[bool]    `json:",omitzero,inline"`
	OfAnyArray []any              `json:",omitzero,inline"`
	OfAnyMap   map[string]any     `json:",omitzero,inline"`
	paramUnion
}

func (u AccountUpdateParamsFieldMapItemUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString,
		u.OfFloat,
		u.OfBool,
		u.OfAnyArray,
		u.OfAnyMap)
}
func (u *AccountUpdateParamsFieldMapItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Relationship operations to apply. System relationships use a `$` prefix (e.g.
// `$owner`, `$contact`). Each value is an operation object with `add`, `remove`,
// or `replace`.
type AccountUpdateParamsRelationships struct {
	// Operation to modify associated contacts.
	Contact AccountUpdateParamsRelationshipsContact `json:"$contact,omitzero"`
	// Operation to modify the account owner.
	Owner       AccountUpdateParamsRelationshipsOwner      `json:"$owner,omitzero"`
	ExtraFields map[string]AccountUpdateParamsRelationship `json:"-"`
	paramObj
}

func (r AccountUpdateParamsRelationships) MarshalJSON() (data []byte, err error) {
	type shadow AccountUpdateParamsRelationships
	return param.MarshalWithExtras(r, (*shadow)(&r), r.ExtraFields)
}
func (r *AccountUpdateParamsRelationships) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Operation to modify associated contacts.
type AccountUpdateParamsRelationshipsContact struct {
	// Entity ID(s) to add to the relationship.
	Add AccountUpdateParamsRelationshipsContactAddUnion `json:"add,omitzero"`
	// Entity ID(s) to remove from the relationship.
	Remove AccountUpdateParamsRelationshipsContactRemoveUnion `json:"remove,omitzero"`
	// Entity ID(s) to set as the entire relationship, replacing all existing
	// associations.
	Replace AccountUpdateParamsRelationshipsContactReplaceUnion `json:"replace,omitzero"`
	paramObj
}

func (r AccountUpdateParamsRelationshipsContact) MarshalJSON() (data []byte, err error) {
	type shadow AccountUpdateParamsRelationshipsContact
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AccountUpdateParamsRelationshipsContact) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type AccountUpdateParamsRelationshipsContactAddUnion struct {
	OfString      param.Opt[string] `json:",omitzero,inline"`
	OfStringArray []string          `json:",omitzero,inline"`
	paramUnion
}

func (u AccountUpdateParamsRelationshipsContactAddUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray)
}
func (u *AccountUpdateParamsRelationshipsContactAddUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type AccountUpdateParamsRelationshipsContactRemoveUnion struct {
	OfString      param.Opt[string] `json:",omitzero,inline"`
	OfStringArray []string          `json:",omitzero,inline"`
	paramUnion
}

func (u AccountUpdateParamsRelationshipsContactRemoveUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray)
}
func (u *AccountUpdateParamsRelationshipsContactRemoveUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type AccountUpdateParamsRelationshipsContactReplaceUnion struct {
	OfString      param.Opt[string] `json:",omitzero,inline"`
	OfStringArray []string          `json:",omitzero,inline"`
	paramUnion
}

func (u AccountUpdateParamsRelationshipsContactReplaceUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray)
}
func (u *AccountUpdateParamsRelationshipsContactReplaceUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Operation to modify the account owner.
type AccountUpdateParamsRelationshipsOwner struct {
	// Entity ID(s) to add to the relationship.
	Add AccountUpdateParamsRelationshipsOwnerAddUnion `json:"add,omitzero"`
	// Entity ID(s) to remove from the relationship.
	Remove AccountUpdateParamsRelationshipsOwnerRemoveUnion `json:"remove,omitzero"`
	// Entity ID(s) to set as the entire relationship, replacing all existing
	// associations.
	Replace AccountUpdateParamsRelationshipsOwnerReplaceUnion `json:"replace,omitzero"`
	paramObj
}

func (r AccountUpdateParamsRelationshipsOwner) MarshalJSON() (data []byte, err error) {
	type shadow AccountUpdateParamsRelationshipsOwner
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AccountUpdateParamsRelationshipsOwner) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type AccountUpdateParamsRelationshipsOwnerAddUnion struct {
	OfString      param.Opt[string] `json:",omitzero,inline"`
	OfStringArray []string          `json:",omitzero,inline"`
	paramUnion
}

func (u AccountUpdateParamsRelationshipsOwnerAddUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray)
}
func (u *AccountUpdateParamsRelationshipsOwnerAddUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type AccountUpdateParamsRelationshipsOwnerRemoveUnion struct {
	OfString      param.Opt[string] `json:",omitzero,inline"`
	OfStringArray []string          `json:",omitzero,inline"`
	paramUnion
}

func (u AccountUpdateParamsRelationshipsOwnerRemoveUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray)
}
func (u *AccountUpdateParamsRelationshipsOwnerRemoveUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type AccountUpdateParamsRelationshipsOwnerReplaceUnion struct {
	OfString      param.Opt[string] `json:",omitzero,inline"`
	OfStringArray []string          `json:",omitzero,inline"`
	paramUnion
}

func (u AccountUpdateParamsRelationshipsOwnerReplaceUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray)
}
func (u *AccountUpdateParamsRelationshipsOwnerReplaceUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// An operation to modify a relationship. Provide one of `add`, `remove`, or
// `replace`.
type AccountUpdateParamsRelationship struct {
	// Entity ID(s) to add to the relationship.
	Add AccountUpdateParamsRelationshipAddUnion `json:"add,omitzero"`
	// Entity ID(s) to remove from the relationship.
	Remove AccountUpdateParamsRelationshipRemoveUnion `json:"remove,omitzero"`
	// Entity ID(s) to set as the entire relationship, replacing all existing
	// associations.
	Replace AccountUpdateParamsRelationshipReplaceUnion `json:"replace,omitzero"`
	paramObj
}

func (r AccountUpdateParamsRelationship) MarshalJSON() (data []byte, err error) {
	type shadow AccountUpdateParamsRelationship
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AccountUpdateParamsRelationship) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type AccountUpdateParamsRelationshipAddUnion struct {
	OfString      param.Opt[string] `json:",omitzero,inline"`
	OfStringArray []string          `json:",omitzero,inline"`
	paramUnion
}

func (u AccountUpdateParamsRelationshipAddUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray)
}
func (u *AccountUpdateParamsRelationshipAddUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type AccountUpdateParamsRelationshipRemoveUnion struct {
	OfString      param.Opt[string] `json:",omitzero,inline"`
	OfStringArray []string          `json:",omitzero,inline"`
	paramUnion
}

func (u AccountUpdateParamsRelationshipRemoveUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray)
}
func (u *AccountUpdateParamsRelationshipRemoveUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type AccountUpdateParamsRelationshipReplaceUnion struct {
	OfString      param.Opt[string] `json:",omitzero,inline"`
	OfStringArray []string          `json:",omitzero,inline"`
	paramUnion
}

func (u AccountUpdateParamsRelationshipReplaceUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray)
}
func (u *AccountUpdateParamsRelationshipReplaceUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

type AccountListParams struct {
	// Maximum number of records to return. Defaults to 25, maximum 25.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Number of records to skip for pagination. Defaults to 0.
	Offset param.Opt[int64] `query:"offset,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [AccountListParams]'s query parameters as `url.Values`.
func (r AccountListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
