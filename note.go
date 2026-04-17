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

// Notes represent free-form text records in Lightfield. Each note can be
// associated with zero or more accounts and opportunities.
//
// NoteService contains methods and other services that help with interacting with
// the Lightfield API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewNoteService] method instead.
type NoteService struct {
	Options []option.RequestOption
}

// NewNoteService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewNoteService(opts ...option.RequestOption) (r NoteService) {
	r = NoteService{}
	r.Options = opts
	return
}

// Creates a new note record.
//
// The note author is automatically set to the API key owner.
//
// Supports idempotency via the `Idempotency-Key` header.
//
// **[Required scope](/using-the-api/scopes/):** `notes:create`
//
// **[Rate limit category](/using-the-api/rate-limits/):** Write
func (r *NoteService) New(ctx context.Context, body NoteNewParams, opts ...option.RequestOption) (res *NoteCreateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/notes"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Retrieves a single note by its ID.
//
// **[Required scope](/using-the-api/scopes/):** `notes:read`
//
// **[Rate limit category](/using-the-api/rate-limits/):** Read
func (r *NoteService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *NoteRetrieveResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/notes/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Updates an existing note by ID. Only included fields and relationships are
// modified.
//
// Both `$account` and `$opportunity` relationships can be modified via `add` or
// `remove` operations.
//
// Supports idempotency via the `Idempotency-Key` header.
//
// **[Required scope](/using-the-api/scopes/):** `notes:update`
//
// **[Rate limit category](/using-the-api/rate-limits/):** Write
func (r *NoteService) Update(ctx context.Context, id string, body NoteUpdateParams, opts ...option.RequestOption) (res *NoteUpdateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/notes/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Returns a paginated list of notes. Use `offset` and `limit` to paginate through
// results. See <u>[List endpoints](/using-the-api/list-endpoints/)</u> for more
// information about pagination.
//
// **[Required scope](/using-the-api/scopes/):** `notes:read`
//
// **[Rate limit category](/using-the-api/rate-limits/):** Search
func (r *NoteService) List(ctx context.Context, query NoteListParams, opts ...option.RequestOption) (res *NoteListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/notes"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

type NoteCreateResponse struct {
	// Unique identifier for the entity.
	ID string `json:"id" api:"required"`
	// ISO 8601 timestamp of when the entity was created.
	CreatedAt string `json:"createdAt" api:"required"`
	// Map of field names to their typed values. System fields are prefixed with `$`
	// (e.g. `$name`, `$email`); custom attributes use their bare slug.
	Fields map[string]NoteCreateResponseField `json:"fields" api:"required"`
	// URL to view the entity in the Lightfield web app, or null.
	HTTPLink string `json:"httpLink" api:"required"`
	// Map of relationship names to their associated entities. System relationships are
	// prefixed with `$` (e.g. `$owner`, `$contact`).
	Relationships map[string]NoteCreateResponseRelationship `json:"relationships" api:"required"`
	// ISO 8601 timestamp of when the entity was last updated, or null.
	UpdatedAt string `json:"updatedAt" api:"required"`
	// External identifier for the entity, or null if unset.
	ExternalID string `json:"externalId" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID            respjson.Field
		CreatedAt     respjson.Field
		Fields        respjson.Field
		HTTPLink      respjson.Field
		Relationships respjson.Field
		UpdatedAt     respjson.Field
		ExternalID    respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NoteCreateResponse) RawJSON() string { return r.JSON.raw }
func (r *NoteCreateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NoteCreateResponseField struct {
	// The field value, or null if unset.
	Value NoteCreateResponseFieldValueUnion `json:"value" api:"required"`
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
func (r NoteCreateResponseField) RawJSON() string { return r.JSON.raw }
func (r *NoteCreateResponseField) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// NoteCreateResponseFieldValueUnion contains all possible properties and values
// from [string], [float64], [bool], [[]string],
// [NoteCreateResponseFieldValueAddress], [NoteCreateResponseFieldValueFullName].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfStringArray]
type NoteCreateResponseFieldValueUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]string] instead of an object.
	OfStringArray []string `json:",inline"`
	// This field is from variant [NoteCreateResponseFieldValueAddress].
	City string `json:"city"`
	// This field is from variant [NoteCreateResponseFieldValueAddress].
	Country string `json:"country"`
	// This field is from variant [NoteCreateResponseFieldValueAddress].
	Latitude float64 `json:"latitude"`
	// This field is from variant [NoteCreateResponseFieldValueAddress].
	Longitude float64 `json:"longitude"`
	// This field is from variant [NoteCreateResponseFieldValueAddress].
	PostalCode string `json:"postalCode"`
	// This field is from variant [NoteCreateResponseFieldValueAddress].
	State string `json:"state"`
	// This field is from variant [NoteCreateResponseFieldValueAddress].
	Street string `json:"street"`
	// This field is from variant [NoteCreateResponseFieldValueAddress].
	Street2 string `json:"street2"`
	// This field is from variant [NoteCreateResponseFieldValueFullName].
	FirstName string `json:"firstName"`
	// This field is from variant [NoteCreateResponseFieldValueFullName].
	LastName string `json:"lastName"`
	JSON     struct {
		OfString      respjson.Field
		OfFloat       respjson.Field
		OfBool        respjson.Field
		OfStringArray respjson.Field
		City          respjson.Field
		Country       respjson.Field
		Latitude      respjson.Field
		Longitude     respjson.Field
		PostalCode    respjson.Field
		State         respjson.Field
		Street        respjson.Field
		Street2       respjson.Field
		FirstName     respjson.Field
		LastName      respjson.Field
		raw           string
	} `json:"-"`
}

func (u NoteCreateResponseFieldValueUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u NoteCreateResponseFieldValueUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u NoteCreateResponseFieldValueUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u NoteCreateResponseFieldValueUnion) AsStringArray() (v []string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u NoteCreateResponseFieldValueUnion) AsAddress() (v NoteCreateResponseFieldValueAddress) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u NoteCreateResponseFieldValueUnion) AsFullName() (v NoteCreateResponseFieldValueFullName) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u NoteCreateResponseFieldValueUnion) RawJSON() string { return u.JSON.raw }

func (r *NoteCreateResponseFieldValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NoteCreateResponseFieldValueAddress struct {
	// City name.
	City string `json:"city" api:"nullable"`
	// 2-letter ISO 3166-1 alpha-2 country code.
	Country string `json:"country" api:"nullable"`
	// Latitude coordinate.
	Latitude float64 `json:"latitude" api:"nullable"`
	// Longitude coordinate.
	Longitude float64 `json:"longitude" api:"nullable"`
	// Postal or ZIP code.
	PostalCode string `json:"postalCode" api:"nullable"`
	// State or province.
	State string `json:"state" api:"nullable"`
	// Street address line 1.
	Street string `json:"street" api:"nullable"`
	// Street address line 2.
	Street2 string `json:"street2" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		City        respjson.Field
		Country     respjson.Field
		Latitude    respjson.Field
		Longitude   respjson.Field
		PostalCode  respjson.Field
		State       respjson.Field
		Street      respjson.Field
		Street2     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NoteCreateResponseFieldValueAddress) RawJSON() string { return r.JSON.raw }
func (r *NoteCreateResponseFieldValueAddress) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NoteCreateResponseFieldValueFullName struct {
	// The contact's first name.
	FirstName string `json:"firstName" api:"nullable"`
	// The contact's last name.
	LastName string `json:"lastName" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FirstName   respjson.Field
		LastName    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NoteCreateResponseFieldValueFullName) RawJSON() string { return r.JSON.raw }
func (r *NoteCreateResponseFieldValueFullName) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NoteCreateResponseRelationship struct {
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
func (r NoteCreateResponseRelationship) RawJSON() string { return r.JSON.raw }
func (r *NoteCreateResponseRelationship) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NoteListResponse struct {
	// Array of entity objects for the current page.
	Data []NoteListResponseData `json:"data" api:"required"`
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
func (r NoteListResponse) RawJSON() string { return r.JSON.raw }
func (r *NoteListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NoteListResponseData struct {
	// Unique identifier for the entity.
	ID string `json:"id" api:"required"`
	// ISO 8601 timestamp of when the entity was created.
	CreatedAt string `json:"createdAt" api:"required"`
	// Map of field names to their typed values. System fields are prefixed with `$`
	// (e.g. `$name`, `$email`); custom attributes use their bare slug.
	Fields map[string]NoteListResponseDataField `json:"fields" api:"required"`
	// URL to view the entity in the Lightfield web app, or null.
	HTTPLink string `json:"httpLink" api:"required"`
	// Map of relationship names to their associated entities. System relationships are
	// prefixed with `$` (e.g. `$owner`, `$contact`).
	Relationships map[string]NoteListResponseDataRelationship `json:"relationships" api:"required"`
	// ISO 8601 timestamp of when the entity was last updated, or null.
	UpdatedAt string `json:"updatedAt" api:"required"`
	// External identifier for the entity, or null if unset.
	ExternalID string `json:"externalId" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID            respjson.Field
		CreatedAt     respjson.Field
		Fields        respjson.Field
		HTTPLink      respjson.Field
		Relationships respjson.Field
		UpdatedAt     respjson.Field
		ExternalID    respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NoteListResponseData) RawJSON() string { return r.JSON.raw }
func (r *NoteListResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NoteListResponseDataField struct {
	// The field value, or null if unset.
	Value NoteListResponseDataFieldValueUnion `json:"value" api:"required"`
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
func (r NoteListResponseDataField) RawJSON() string { return r.JSON.raw }
func (r *NoteListResponseDataField) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// NoteListResponseDataFieldValueUnion contains all possible properties and values
// from [string], [float64], [bool], [[]string],
// [NoteListResponseDataFieldValueAddress],
// [NoteListResponseDataFieldValueFullName].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfStringArray]
type NoteListResponseDataFieldValueUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]string] instead of an object.
	OfStringArray []string `json:",inline"`
	// This field is from variant [NoteListResponseDataFieldValueAddress].
	City string `json:"city"`
	// This field is from variant [NoteListResponseDataFieldValueAddress].
	Country string `json:"country"`
	// This field is from variant [NoteListResponseDataFieldValueAddress].
	Latitude float64 `json:"latitude"`
	// This field is from variant [NoteListResponseDataFieldValueAddress].
	Longitude float64 `json:"longitude"`
	// This field is from variant [NoteListResponseDataFieldValueAddress].
	PostalCode string `json:"postalCode"`
	// This field is from variant [NoteListResponseDataFieldValueAddress].
	State string `json:"state"`
	// This field is from variant [NoteListResponseDataFieldValueAddress].
	Street string `json:"street"`
	// This field is from variant [NoteListResponseDataFieldValueAddress].
	Street2 string `json:"street2"`
	// This field is from variant [NoteListResponseDataFieldValueFullName].
	FirstName string `json:"firstName"`
	// This field is from variant [NoteListResponseDataFieldValueFullName].
	LastName string `json:"lastName"`
	JSON     struct {
		OfString      respjson.Field
		OfFloat       respjson.Field
		OfBool        respjson.Field
		OfStringArray respjson.Field
		City          respjson.Field
		Country       respjson.Field
		Latitude      respjson.Field
		Longitude     respjson.Field
		PostalCode    respjson.Field
		State         respjson.Field
		Street        respjson.Field
		Street2       respjson.Field
		FirstName     respjson.Field
		LastName      respjson.Field
		raw           string
	} `json:"-"`
}

func (u NoteListResponseDataFieldValueUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u NoteListResponseDataFieldValueUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u NoteListResponseDataFieldValueUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u NoteListResponseDataFieldValueUnion) AsStringArray() (v []string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u NoteListResponseDataFieldValueUnion) AsAddress() (v NoteListResponseDataFieldValueAddress) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u NoteListResponseDataFieldValueUnion) AsFullName() (v NoteListResponseDataFieldValueFullName) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u NoteListResponseDataFieldValueUnion) RawJSON() string { return u.JSON.raw }

func (r *NoteListResponseDataFieldValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NoteListResponseDataFieldValueAddress struct {
	// City name.
	City string `json:"city" api:"nullable"`
	// 2-letter ISO 3166-1 alpha-2 country code.
	Country string `json:"country" api:"nullable"`
	// Latitude coordinate.
	Latitude float64 `json:"latitude" api:"nullable"`
	// Longitude coordinate.
	Longitude float64 `json:"longitude" api:"nullable"`
	// Postal or ZIP code.
	PostalCode string `json:"postalCode" api:"nullable"`
	// State or province.
	State string `json:"state" api:"nullable"`
	// Street address line 1.
	Street string `json:"street" api:"nullable"`
	// Street address line 2.
	Street2 string `json:"street2" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		City        respjson.Field
		Country     respjson.Field
		Latitude    respjson.Field
		Longitude   respjson.Field
		PostalCode  respjson.Field
		State       respjson.Field
		Street      respjson.Field
		Street2     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NoteListResponseDataFieldValueAddress) RawJSON() string { return r.JSON.raw }
func (r *NoteListResponseDataFieldValueAddress) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NoteListResponseDataFieldValueFullName struct {
	// The contact's first name.
	FirstName string `json:"firstName" api:"nullable"`
	// The contact's last name.
	LastName string `json:"lastName" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FirstName   respjson.Field
		LastName    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NoteListResponseDataFieldValueFullName) RawJSON() string { return r.JSON.raw }
func (r *NoteListResponseDataFieldValueFullName) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NoteListResponseDataRelationship struct {
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
func (r NoteListResponseDataRelationship) RawJSON() string { return r.JSON.raw }
func (r *NoteListResponseDataRelationship) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NoteRetrieveResponse struct {
	// Unique identifier for the entity.
	ID string `json:"id" api:"required"`
	// ISO 8601 timestamp of when the entity was created.
	CreatedAt string `json:"createdAt" api:"required"`
	// Map of field names to their typed values. System fields are prefixed with `$`
	// (e.g. `$name`, `$email`); custom attributes use their bare slug.
	Fields map[string]NoteRetrieveResponseField `json:"fields" api:"required"`
	// URL to view the entity in the Lightfield web app, or null.
	HTTPLink string `json:"httpLink" api:"required"`
	// Map of relationship names to their associated entities. System relationships are
	// prefixed with `$` (e.g. `$owner`, `$contact`).
	Relationships map[string]NoteRetrieveResponseRelationship `json:"relationships" api:"required"`
	// ISO 8601 timestamp of when the entity was last updated, or null.
	UpdatedAt string `json:"updatedAt" api:"required"`
	// External identifier for the entity, or null if unset.
	ExternalID string `json:"externalId" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID            respjson.Field
		CreatedAt     respjson.Field
		Fields        respjson.Field
		HTTPLink      respjson.Field
		Relationships respjson.Field
		UpdatedAt     respjson.Field
		ExternalID    respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NoteRetrieveResponse) RawJSON() string { return r.JSON.raw }
func (r *NoteRetrieveResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NoteRetrieveResponseField struct {
	// The field value, or null if unset.
	Value NoteRetrieveResponseFieldValueUnion `json:"value" api:"required"`
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
func (r NoteRetrieveResponseField) RawJSON() string { return r.JSON.raw }
func (r *NoteRetrieveResponseField) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// NoteRetrieveResponseFieldValueUnion contains all possible properties and values
// from [string], [float64], [bool], [[]string],
// [NoteRetrieveResponseFieldValueAddress],
// [NoteRetrieveResponseFieldValueFullName].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfStringArray]
type NoteRetrieveResponseFieldValueUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]string] instead of an object.
	OfStringArray []string `json:",inline"`
	// This field is from variant [NoteRetrieveResponseFieldValueAddress].
	City string `json:"city"`
	// This field is from variant [NoteRetrieveResponseFieldValueAddress].
	Country string `json:"country"`
	// This field is from variant [NoteRetrieveResponseFieldValueAddress].
	Latitude float64 `json:"latitude"`
	// This field is from variant [NoteRetrieveResponseFieldValueAddress].
	Longitude float64 `json:"longitude"`
	// This field is from variant [NoteRetrieveResponseFieldValueAddress].
	PostalCode string `json:"postalCode"`
	// This field is from variant [NoteRetrieveResponseFieldValueAddress].
	State string `json:"state"`
	// This field is from variant [NoteRetrieveResponseFieldValueAddress].
	Street string `json:"street"`
	// This field is from variant [NoteRetrieveResponseFieldValueAddress].
	Street2 string `json:"street2"`
	// This field is from variant [NoteRetrieveResponseFieldValueFullName].
	FirstName string `json:"firstName"`
	// This field is from variant [NoteRetrieveResponseFieldValueFullName].
	LastName string `json:"lastName"`
	JSON     struct {
		OfString      respjson.Field
		OfFloat       respjson.Field
		OfBool        respjson.Field
		OfStringArray respjson.Field
		City          respjson.Field
		Country       respjson.Field
		Latitude      respjson.Field
		Longitude     respjson.Field
		PostalCode    respjson.Field
		State         respjson.Field
		Street        respjson.Field
		Street2       respjson.Field
		FirstName     respjson.Field
		LastName      respjson.Field
		raw           string
	} `json:"-"`
}

func (u NoteRetrieveResponseFieldValueUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u NoteRetrieveResponseFieldValueUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u NoteRetrieveResponseFieldValueUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u NoteRetrieveResponseFieldValueUnion) AsStringArray() (v []string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u NoteRetrieveResponseFieldValueUnion) AsAddress() (v NoteRetrieveResponseFieldValueAddress) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u NoteRetrieveResponseFieldValueUnion) AsFullName() (v NoteRetrieveResponseFieldValueFullName) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u NoteRetrieveResponseFieldValueUnion) RawJSON() string { return u.JSON.raw }

func (r *NoteRetrieveResponseFieldValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NoteRetrieveResponseFieldValueAddress struct {
	// City name.
	City string `json:"city" api:"nullable"`
	// 2-letter ISO 3166-1 alpha-2 country code.
	Country string `json:"country" api:"nullable"`
	// Latitude coordinate.
	Latitude float64 `json:"latitude" api:"nullable"`
	// Longitude coordinate.
	Longitude float64 `json:"longitude" api:"nullable"`
	// Postal or ZIP code.
	PostalCode string `json:"postalCode" api:"nullable"`
	// State or province.
	State string `json:"state" api:"nullable"`
	// Street address line 1.
	Street string `json:"street" api:"nullable"`
	// Street address line 2.
	Street2 string `json:"street2" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		City        respjson.Field
		Country     respjson.Field
		Latitude    respjson.Field
		Longitude   respjson.Field
		PostalCode  respjson.Field
		State       respjson.Field
		Street      respjson.Field
		Street2     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NoteRetrieveResponseFieldValueAddress) RawJSON() string { return r.JSON.raw }
func (r *NoteRetrieveResponseFieldValueAddress) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NoteRetrieveResponseFieldValueFullName struct {
	// The contact's first name.
	FirstName string `json:"firstName" api:"nullable"`
	// The contact's last name.
	LastName string `json:"lastName" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FirstName   respjson.Field
		LastName    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NoteRetrieveResponseFieldValueFullName) RawJSON() string { return r.JSON.raw }
func (r *NoteRetrieveResponseFieldValueFullName) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NoteRetrieveResponseRelationship struct {
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
func (r NoteRetrieveResponseRelationship) RawJSON() string { return r.JSON.raw }
func (r *NoteRetrieveResponseRelationship) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NoteUpdateResponse struct {
	// Unique identifier for the entity.
	ID string `json:"id" api:"required"`
	// ISO 8601 timestamp of when the entity was created.
	CreatedAt string `json:"createdAt" api:"required"`
	// Map of field names to their typed values. System fields are prefixed with `$`
	// (e.g. `$name`, `$email`); custom attributes use their bare slug.
	Fields map[string]NoteUpdateResponseField `json:"fields" api:"required"`
	// URL to view the entity in the Lightfield web app, or null.
	HTTPLink string `json:"httpLink" api:"required"`
	// Map of relationship names to their associated entities. System relationships are
	// prefixed with `$` (e.g. `$owner`, `$contact`).
	Relationships map[string]NoteUpdateResponseRelationship `json:"relationships" api:"required"`
	// ISO 8601 timestamp of when the entity was last updated, or null.
	UpdatedAt string `json:"updatedAt" api:"required"`
	// External identifier for the entity, or null if unset.
	ExternalID string `json:"externalId" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID            respjson.Field
		CreatedAt     respjson.Field
		Fields        respjson.Field
		HTTPLink      respjson.Field
		Relationships respjson.Field
		UpdatedAt     respjson.Field
		ExternalID    respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NoteUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *NoteUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NoteUpdateResponseField struct {
	// The field value, or null if unset.
	Value NoteUpdateResponseFieldValueUnion `json:"value" api:"required"`
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
func (r NoteUpdateResponseField) RawJSON() string { return r.JSON.raw }
func (r *NoteUpdateResponseField) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// NoteUpdateResponseFieldValueUnion contains all possible properties and values
// from [string], [float64], [bool], [[]string],
// [NoteUpdateResponseFieldValueAddress], [NoteUpdateResponseFieldValueFullName].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfStringArray]
type NoteUpdateResponseFieldValueUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]string] instead of an object.
	OfStringArray []string `json:",inline"`
	// This field is from variant [NoteUpdateResponseFieldValueAddress].
	City string `json:"city"`
	// This field is from variant [NoteUpdateResponseFieldValueAddress].
	Country string `json:"country"`
	// This field is from variant [NoteUpdateResponseFieldValueAddress].
	Latitude float64 `json:"latitude"`
	// This field is from variant [NoteUpdateResponseFieldValueAddress].
	Longitude float64 `json:"longitude"`
	// This field is from variant [NoteUpdateResponseFieldValueAddress].
	PostalCode string `json:"postalCode"`
	// This field is from variant [NoteUpdateResponseFieldValueAddress].
	State string `json:"state"`
	// This field is from variant [NoteUpdateResponseFieldValueAddress].
	Street string `json:"street"`
	// This field is from variant [NoteUpdateResponseFieldValueAddress].
	Street2 string `json:"street2"`
	// This field is from variant [NoteUpdateResponseFieldValueFullName].
	FirstName string `json:"firstName"`
	// This field is from variant [NoteUpdateResponseFieldValueFullName].
	LastName string `json:"lastName"`
	JSON     struct {
		OfString      respjson.Field
		OfFloat       respjson.Field
		OfBool        respjson.Field
		OfStringArray respjson.Field
		City          respjson.Field
		Country       respjson.Field
		Latitude      respjson.Field
		Longitude     respjson.Field
		PostalCode    respjson.Field
		State         respjson.Field
		Street        respjson.Field
		Street2       respjson.Field
		FirstName     respjson.Field
		LastName      respjson.Field
		raw           string
	} `json:"-"`
}

func (u NoteUpdateResponseFieldValueUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u NoteUpdateResponseFieldValueUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u NoteUpdateResponseFieldValueUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u NoteUpdateResponseFieldValueUnion) AsStringArray() (v []string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u NoteUpdateResponseFieldValueUnion) AsAddress() (v NoteUpdateResponseFieldValueAddress) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u NoteUpdateResponseFieldValueUnion) AsFullName() (v NoteUpdateResponseFieldValueFullName) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u NoteUpdateResponseFieldValueUnion) RawJSON() string { return u.JSON.raw }

func (r *NoteUpdateResponseFieldValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NoteUpdateResponseFieldValueAddress struct {
	// City name.
	City string `json:"city" api:"nullable"`
	// 2-letter ISO 3166-1 alpha-2 country code.
	Country string `json:"country" api:"nullable"`
	// Latitude coordinate.
	Latitude float64 `json:"latitude" api:"nullable"`
	// Longitude coordinate.
	Longitude float64 `json:"longitude" api:"nullable"`
	// Postal or ZIP code.
	PostalCode string `json:"postalCode" api:"nullable"`
	// State or province.
	State string `json:"state" api:"nullable"`
	// Street address line 1.
	Street string `json:"street" api:"nullable"`
	// Street address line 2.
	Street2 string `json:"street2" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		City        respjson.Field
		Country     respjson.Field
		Latitude    respjson.Field
		Longitude   respjson.Field
		PostalCode  respjson.Field
		State       respjson.Field
		Street      respjson.Field
		Street2     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NoteUpdateResponseFieldValueAddress) RawJSON() string { return r.JSON.raw }
func (r *NoteUpdateResponseFieldValueAddress) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NoteUpdateResponseFieldValueFullName struct {
	// The contact's first name.
	FirstName string `json:"firstName" api:"nullable"`
	// The contact's last name.
	LastName string `json:"lastName" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FirstName   respjson.Field
		LastName    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NoteUpdateResponseFieldValueFullName) RawJSON() string { return r.JSON.raw }
func (r *NoteUpdateResponseFieldValueFullName) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NoteUpdateResponseRelationship struct {
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
func (r NoteUpdateResponseRelationship) RawJSON() string { return r.JSON.raw }
func (r *NoteUpdateResponseRelationship) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NoteNewParams struct {
	// Field values for the new note. `$title` is required; `$content` is optional. See
	// **[Fields and relationships](/using-the-api/fields-and-relationships/)** for
	// value type details.
	Fields NoteNewParamsFields `json:"fields,omitzero" api:"required"`
	// Relationships to set on the new note. System relationships use a `$` prefix
	// (e.g. `$account`, `$opportunity`). Each value is a single entity ID or an array
	// of IDs. The note author is automatically set to the API key owner.
	Relationships NoteNewParamsRelationships `json:"relationships,omitzero"`
	paramObj
}

func (r NoteNewParams) MarshalJSON() (data []byte, err error) {
	type shadow NoteNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *NoteNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Field values for the new note. `$title` is required; `$content` is optional. See
// **[Fields and relationships](/using-the-api/fields-and-relationships/)** for
// value type details.
//
// The property Title is required.
type NoteNewParamsFields struct {
	// Title of the note.
	Title string `json:"$title" api:"required"`
	// Content of the note as markdown formatted text.
	Content param.Opt[string] `json:"$content,omitzero"`
	paramObj
}

func (r NoteNewParamsFields) MarshalJSON() (data []byte, err error) {
	type shadow NoteNewParamsFields
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *NoteNewParamsFields) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Relationships to set on the new note. System relationships use a `$` prefix
// (e.g. `$account`, `$opportunity`). Each value is a single entity ID or an array
// of IDs. The note author is automatically set to the API key owner.
type NoteNewParamsRelationships struct {
	// ID(s) of accounts to associate with this note.
	Account NoteNewParamsRelationshipsAccountUnion `json:"$account,omitzero"`
	// ID(s) of opportunities to associate with this note.
	Opportunity NoteNewParamsRelationshipsOpportunityUnion `json:"$opportunity,omitzero"`
	paramObj
}

func (r NoteNewParamsRelationships) MarshalJSON() (data []byte, err error) {
	type shadow NoteNewParamsRelationships
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *NoteNewParamsRelationships) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type NoteNewParamsRelationshipsAccountUnion struct {
	OfString      param.Opt[string] `json:",omitzero,inline"`
	OfStringArray []string          `json:",omitzero,inline"`
	paramUnion
}

func (u NoteNewParamsRelationshipsAccountUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray)
}
func (u *NoteNewParamsRelationshipsAccountUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type NoteNewParamsRelationshipsOpportunityUnion struct {
	OfString      param.Opt[string] `json:",omitzero,inline"`
	OfStringArray []string          `json:",omitzero,inline"`
	paramUnion
}

func (u NoteNewParamsRelationshipsOpportunityUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray)
}
func (u *NoteNewParamsRelationshipsOpportunityUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

type NoteUpdateParams struct {
	// Field values to update — only provided fields are modified; omitted fields are
	// left unchanged. See
	// **[Fields and relationships](/using-the-api/fields-and-relationships/)** for
	// value type details.
	Fields NoteUpdateParamsFields `json:"fields,omitzero"`
	// Relationship operations to apply. System relationships use a `$` prefix (e.g.
	// `$account`, `$opportunity`). Each value is an operation object with `add` or
	// `remove`.
	Relationships NoteUpdateParamsRelationships `json:"relationships,omitzero"`
	paramObj
}

func (r NoteUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow NoteUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *NoteUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Field values to update — only provided fields are modified; omitted fields are
// left unchanged. See
// **[Fields and relationships](/using-the-api/fields-and-relationships/)** for
// value type details.
type NoteUpdateParamsFields struct {
	// Content of the note as markdown formatted text.
	Content param.Opt[string] `json:"$content,omitzero"`
	// Title of the note.
	Title param.Opt[string] `json:"$title,omitzero"`
	paramObj
}

func (r NoteUpdateParamsFields) MarshalJSON() (data []byte, err error) {
	type shadow NoteUpdateParamsFields
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *NoteUpdateParamsFields) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Relationship operations to apply. System relationships use a `$` prefix (e.g.
// `$account`, `$opportunity`). Each value is an operation object with `add` or
// `remove`.
type NoteUpdateParamsRelationships struct {
	// Operation to modify associated accounts.
	Account NoteUpdateParamsRelationshipsAccountUnion `json:"$account,omitzero"`
	// Operation to modify associated opportunities.
	Opportunity NoteUpdateParamsRelationshipsOpportunityUnion `json:"$opportunity,omitzero"`
	paramObj
}

func (r NoteUpdateParamsRelationships) MarshalJSON() (data []byte, err error) {
	type shadow NoteUpdateParamsRelationships
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *NoteUpdateParamsRelationships) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type NoteUpdateParamsRelationshipsAccountUnion struct {
	OfNoteUpdatesRelationshipsAccountAdd    *NoteUpdateParamsRelationshipsAccountAdd    `json:",omitzero,inline"`
	OfNoteUpdatesRelationshipsAccountRemove *NoteUpdateParamsRelationshipsAccountRemove `json:",omitzero,inline"`
	paramUnion
}

func (u NoteUpdateParamsRelationshipsAccountUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfNoteUpdatesRelationshipsAccountAdd, u.OfNoteUpdatesRelationshipsAccountRemove)
}
func (u *NoteUpdateParamsRelationshipsAccountUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// The property Add is required.
type NoteUpdateParamsRelationshipsAccountAdd struct {
	// Entity ID(s) to add to the relationship.
	Add NoteUpdateParamsRelationshipsAccountAddAddUnion `json:"add,omitzero" api:"required"`
	paramObj
}

func (r NoteUpdateParamsRelationshipsAccountAdd) MarshalJSON() (data []byte, err error) {
	type shadow NoteUpdateParamsRelationshipsAccountAdd
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *NoteUpdateParamsRelationshipsAccountAdd) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type NoteUpdateParamsRelationshipsAccountAddAddUnion struct {
	OfString      param.Opt[string] `json:",omitzero,inline"`
	OfStringArray []string          `json:",omitzero,inline"`
	paramUnion
}

func (u NoteUpdateParamsRelationshipsAccountAddAddUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray)
}
func (u *NoteUpdateParamsRelationshipsAccountAddAddUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// The property Remove is required.
type NoteUpdateParamsRelationshipsAccountRemove struct {
	// Entity ID(s) to remove from the relationship.
	Remove NoteUpdateParamsRelationshipsAccountRemoveRemoveUnion `json:"remove,omitzero" api:"required"`
	paramObj
}

func (r NoteUpdateParamsRelationshipsAccountRemove) MarshalJSON() (data []byte, err error) {
	type shadow NoteUpdateParamsRelationshipsAccountRemove
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *NoteUpdateParamsRelationshipsAccountRemove) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type NoteUpdateParamsRelationshipsAccountRemoveRemoveUnion struct {
	OfString      param.Opt[string] `json:",omitzero,inline"`
	OfStringArray []string          `json:",omitzero,inline"`
	paramUnion
}

func (u NoteUpdateParamsRelationshipsAccountRemoveRemoveUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray)
}
func (u *NoteUpdateParamsRelationshipsAccountRemoveRemoveUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type NoteUpdateParamsRelationshipsOpportunityUnion struct {
	OfNoteUpdatesRelationshipsOpportunityAdd    *NoteUpdateParamsRelationshipsOpportunityAdd    `json:",omitzero,inline"`
	OfNoteUpdatesRelationshipsOpportunityRemove *NoteUpdateParamsRelationshipsOpportunityRemove `json:",omitzero,inline"`
	paramUnion
}

func (u NoteUpdateParamsRelationshipsOpportunityUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfNoteUpdatesRelationshipsOpportunityAdd, u.OfNoteUpdatesRelationshipsOpportunityRemove)
}
func (u *NoteUpdateParamsRelationshipsOpportunityUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// The property Add is required.
type NoteUpdateParamsRelationshipsOpportunityAdd struct {
	// Entity ID(s) to add to the relationship.
	Add NoteUpdateParamsRelationshipsOpportunityAddAddUnion `json:"add,omitzero" api:"required"`
	paramObj
}

func (r NoteUpdateParamsRelationshipsOpportunityAdd) MarshalJSON() (data []byte, err error) {
	type shadow NoteUpdateParamsRelationshipsOpportunityAdd
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *NoteUpdateParamsRelationshipsOpportunityAdd) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type NoteUpdateParamsRelationshipsOpportunityAddAddUnion struct {
	OfString      param.Opt[string] `json:",omitzero,inline"`
	OfStringArray []string          `json:",omitzero,inline"`
	paramUnion
}

func (u NoteUpdateParamsRelationshipsOpportunityAddAddUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray)
}
func (u *NoteUpdateParamsRelationshipsOpportunityAddAddUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// The property Remove is required.
type NoteUpdateParamsRelationshipsOpportunityRemove struct {
	// Entity ID(s) to remove from the relationship.
	Remove NoteUpdateParamsRelationshipsOpportunityRemoveRemoveUnion `json:"remove,omitzero" api:"required"`
	paramObj
}

func (r NoteUpdateParamsRelationshipsOpportunityRemove) MarshalJSON() (data []byte, err error) {
	type shadow NoteUpdateParamsRelationshipsOpportunityRemove
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *NoteUpdateParamsRelationshipsOpportunityRemove) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type NoteUpdateParamsRelationshipsOpportunityRemoveRemoveUnion struct {
	OfString      param.Opt[string] `json:",omitzero,inline"`
	OfStringArray []string          `json:",omitzero,inline"`
	paramUnion
}

func (u NoteUpdateParamsRelationshipsOpportunityRemoveRemoveUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray)
}
func (u *NoteUpdateParamsRelationshipsOpportunityRemoveRemoveUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

type NoteListParams struct {
	// Maximum number of records to return. Defaults to 25, maximum 25.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Number of records to skip for pagination. Defaults to 0.
	Offset param.Opt[int64] `query:"offset,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [NoteListParams]'s query parameters as `url.Values`.
func (r NoteListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
