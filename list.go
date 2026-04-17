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

// Lists are curated collections of accounts, contacts, or opportunities in
// Lightfield. Each list contains entities of a single type.
//
// ListService contains methods and other services that help with interacting with
// the Lightfield API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewListService] method instead.
type ListService struct {
	Options []option.RequestOption
}

// NewListService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewListService(opts ...option.RequestOption) (r ListService) {
	r = ListService{}
	r.Options = opts
	return
}

// Creates a new list. The `$name` and `$objectType` fields are required.
//
// Supports idempotency via the `Idempotency-Key` header.
//
// **[Required scope](/using-the-api/scopes/):** `lists:create`
//
// **[Rate limit category](/using-the-api/rate-limits/):** Write
func (r *ListService) New(ctx context.Context, body ListNewParams, opts ...option.RequestOption) (res *ListCreateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/lists"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Retrieves a single list by its ID.
//
// **[Required scope](/using-the-api/scopes/):** `lists:read`
//
// **[Rate limit category](/using-the-api/rate-limits/):** Read
func (r *ListService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *ListRetrieveResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/lists/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Updates an existing list by ID. Only included fields are modified.
//
// Supports idempotency via the `Idempotency-Key` header.
//
// **[Required scope](/using-the-api/scopes/):** `lists:update`
//
// **[Rate limit category](/using-the-api/rate-limits/):** Write
func (r *ListService) Update(ctx context.Context, id string, body ListUpdateParams, opts ...option.RequestOption) (res *ListUpdateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/lists/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Returns a paginated list of lists. Use `offset` and `limit` to paginate through
// results. See <u>[List endpoints](/using-the-api/list-endpoints/)</u> for more
// information about pagination.
//
// **[Required scope](/using-the-api/scopes/):** `lists:read`
//
// **[Rate limit category](/using-the-api/rate-limits/):** Search
func (r *ListService) List(ctx context.Context, query ListListParams, opts ...option.RequestOption) (res *ListListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/lists"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Returns a paginated list of accounts that belong to the specified list.
//
// **[Required scopes](/using-the-api/scopes/):** `lists:read` and `accounts:read`
//
// **[Rate limit category](/using-the-api/rate-limits/):** Search
func (r *ListService) ListAccounts(ctx context.Context, listID string, query ListListAccountsParams, opts ...option.RequestOption) (res *ListListAccountsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if listID == "" {
		err = errors.New("missing required listId parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/lists/%s/accounts", url.PathEscape(listID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Returns a paginated list of contacts that belong to the specified list.
//
// **[Required scopes](/using-the-api/scopes/):** `lists:read` and `contacts:read`
//
// **[Rate limit category](/using-the-api/rate-limits/):** Search
func (r *ListService) ListContacts(ctx context.Context, listID string, query ListListContactsParams, opts ...option.RequestOption) (res *ListListContactsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if listID == "" {
		err = errors.New("missing required listId parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/lists/%s/contacts", url.PathEscape(listID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Returns a paginated list of opportunities that belong to the specified list.
//
// **[Required scopes](/using-the-api/scopes/):** `lists:read` and
// `opportunities:read`
//
// **[Rate limit category](/using-the-api/rate-limits/):** Search
func (r *ListService) ListOpportunities(ctx context.Context, listID string, query ListListOpportunitiesParams, opts ...option.RequestOption) (res *ListListOpportunitiesResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if listID == "" {
		err = errors.New("missing required listId parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/lists/%s/opportunities", url.PathEscape(listID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

type ListCreateResponse struct {
	// Unique identifier for the list.
	ID string `json:"id" api:"required"`
	// ISO 8601 timestamp of when the list was created.
	CreatedAt string `json:"createdAt" api:"required"`
	// Map of field names to their typed values. System fields are prefixed with `$`
	// (e.g. `$name`, `$objectType`).
	Fields map[string]ListCreateResponseField `json:"fields" api:"required"`
	// URL to view the list in the Lightfield web app, or null.
	HTTPLink string `json:"httpLink" api:"required"`
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
func (r ListCreateResponse) RawJSON() string { return r.JSON.raw }
func (r *ListCreateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ListCreateResponseField struct {
	// The field value, or null if unset.
	Value ListCreateResponseFieldValueUnion `json:"value" api:"required"`
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
func (r ListCreateResponseField) RawJSON() string { return r.JSON.raw }
func (r *ListCreateResponseField) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ListCreateResponseFieldValueUnion contains all possible properties and values
// from [string], [float64], [bool], [[]string],
// [ListCreateResponseFieldValueAddress], [ListCreateResponseFieldValueFullName].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfStringArray]
type ListCreateResponseFieldValueUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]string] instead of an object.
	OfStringArray []string `json:",inline"`
	// This field is from variant [ListCreateResponseFieldValueAddress].
	City string `json:"city"`
	// This field is from variant [ListCreateResponseFieldValueAddress].
	Country string `json:"country"`
	// This field is from variant [ListCreateResponseFieldValueAddress].
	Latitude float64 `json:"latitude"`
	// This field is from variant [ListCreateResponseFieldValueAddress].
	Longitude float64 `json:"longitude"`
	// This field is from variant [ListCreateResponseFieldValueAddress].
	PostalCode string `json:"postalCode"`
	// This field is from variant [ListCreateResponseFieldValueAddress].
	State string `json:"state"`
	// This field is from variant [ListCreateResponseFieldValueAddress].
	Street string `json:"street"`
	// This field is from variant [ListCreateResponseFieldValueAddress].
	Street2 string `json:"street2"`
	// This field is from variant [ListCreateResponseFieldValueFullName].
	FirstName string `json:"firstName"`
	// This field is from variant [ListCreateResponseFieldValueFullName].
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

func (u ListCreateResponseFieldValueUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ListCreateResponseFieldValueUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ListCreateResponseFieldValueUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ListCreateResponseFieldValueUnion) AsStringArray() (v []string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ListCreateResponseFieldValueUnion) AsAddress() (v ListCreateResponseFieldValueAddress) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ListCreateResponseFieldValueUnion) AsFullName() (v ListCreateResponseFieldValueFullName) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ListCreateResponseFieldValueUnion) RawJSON() string { return u.JSON.raw }

func (r *ListCreateResponseFieldValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ListCreateResponseFieldValueAddress struct {
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
func (r ListCreateResponseFieldValueAddress) RawJSON() string { return r.JSON.raw }
func (r *ListCreateResponseFieldValueAddress) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ListCreateResponseFieldValueFullName struct {
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
func (r ListCreateResponseFieldValueFullName) RawJSON() string { return r.JSON.raw }
func (r *ListCreateResponseFieldValueFullName) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ListListAccountsResponse struct {
	// Array of entity objects for the current page.
	Data []ListListAccountsResponseData `json:"data" api:"required"`
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
func (r ListListAccountsResponse) RawJSON() string { return r.JSON.raw }
func (r *ListListAccountsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ListListAccountsResponseData struct {
	// Unique identifier for the entity.
	ID string `json:"id" api:"required"`
	// ISO 8601 timestamp of when the entity was created.
	CreatedAt string `json:"createdAt" api:"required"`
	// Map of field names to their typed values. System fields are prefixed with `$`
	// (e.g. `$name`, `$email`); custom attributes use their bare slug.
	Fields map[string]ListListAccountsResponseDataField `json:"fields" api:"required"`
	// URL to view the entity in the Lightfield web app, or null.
	HTTPLink string `json:"httpLink" api:"required"`
	// Map of relationship names to their associated entities. System relationships are
	// prefixed with `$` (e.g. `$owner`, `$contact`).
	Relationships map[string]ListListAccountsResponseDataRelationship `json:"relationships" api:"required"`
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
func (r ListListAccountsResponseData) RawJSON() string { return r.JSON.raw }
func (r *ListListAccountsResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ListListAccountsResponseDataField struct {
	// The field value, or null if unset.
	Value ListListAccountsResponseDataFieldValueUnion `json:"value" api:"required"`
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
func (r ListListAccountsResponseDataField) RawJSON() string { return r.JSON.raw }
func (r *ListListAccountsResponseDataField) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ListListAccountsResponseDataFieldValueUnion contains all possible properties and
// values from [string], [float64], [bool], [[]string],
// [ListListAccountsResponseDataFieldValueAddress],
// [ListListAccountsResponseDataFieldValueFullName].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfStringArray]
type ListListAccountsResponseDataFieldValueUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]string] instead of an object.
	OfStringArray []string `json:",inline"`
	// This field is from variant [ListListAccountsResponseDataFieldValueAddress].
	City string `json:"city"`
	// This field is from variant [ListListAccountsResponseDataFieldValueAddress].
	Country string `json:"country"`
	// This field is from variant [ListListAccountsResponseDataFieldValueAddress].
	Latitude float64 `json:"latitude"`
	// This field is from variant [ListListAccountsResponseDataFieldValueAddress].
	Longitude float64 `json:"longitude"`
	// This field is from variant [ListListAccountsResponseDataFieldValueAddress].
	PostalCode string `json:"postalCode"`
	// This field is from variant [ListListAccountsResponseDataFieldValueAddress].
	State string `json:"state"`
	// This field is from variant [ListListAccountsResponseDataFieldValueAddress].
	Street string `json:"street"`
	// This field is from variant [ListListAccountsResponseDataFieldValueAddress].
	Street2 string `json:"street2"`
	// This field is from variant [ListListAccountsResponseDataFieldValueFullName].
	FirstName string `json:"firstName"`
	// This field is from variant [ListListAccountsResponseDataFieldValueFullName].
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

func (u ListListAccountsResponseDataFieldValueUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ListListAccountsResponseDataFieldValueUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ListListAccountsResponseDataFieldValueUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ListListAccountsResponseDataFieldValueUnion) AsStringArray() (v []string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ListListAccountsResponseDataFieldValueUnion) AsAddress() (v ListListAccountsResponseDataFieldValueAddress) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ListListAccountsResponseDataFieldValueUnion) AsFullName() (v ListListAccountsResponseDataFieldValueFullName) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ListListAccountsResponseDataFieldValueUnion) RawJSON() string { return u.JSON.raw }

func (r *ListListAccountsResponseDataFieldValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ListListAccountsResponseDataFieldValueAddress struct {
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
func (r ListListAccountsResponseDataFieldValueAddress) RawJSON() string { return r.JSON.raw }
func (r *ListListAccountsResponseDataFieldValueAddress) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ListListAccountsResponseDataFieldValueFullName struct {
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
func (r ListListAccountsResponseDataFieldValueFullName) RawJSON() string { return r.JSON.raw }
func (r *ListListAccountsResponseDataFieldValueFullName) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ListListAccountsResponseDataRelationship struct {
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
func (r ListListAccountsResponseDataRelationship) RawJSON() string { return r.JSON.raw }
func (r *ListListAccountsResponseDataRelationship) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ListListContactsResponse struct {
	// Array of entity objects for the current page.
	Data []ListListContactsResponseData `json:"data" api:"required"`
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
func (r ListListContactsResponse) RawJSON() string { return r.JSON.raw }
func (r *ListListContactsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ListListContactsResponseData struct {
	// Unique identifier for the entity.
	ID string `json:"id" api:"required"`
	// ISO 8601 timestamp of when the entity was created.
	CreatedAt string `json:"createdAt" api:"required"`
	// Map of field names to their typed values. System fields are prefixed with `$`
	// (e.g. `$name`, `$email`); custom attributes use their bare slug.
	Fields map[string]ListListContactsResponseDataField `json:"fields" api:"required"`
	// URL to view the entity in the Lightfield web app, or null.
	HTTPLink string `json:"httpLink" api:"required"`
	// Map of relationship names to their associated entities. System relationships are
	// prefixed with `$` (e.g. `$owner`, `$contact`).
	Relationships map[string]ListListContactsResponseDataRelationship `json:"relationships" api:"required"`
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
func (r ListListContactsResponseData) RawJSON() string { return r.JSON.raw }
func (r *ListListContactsResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ListListContactsResponseDataField struct {
	// The field value, or null if unset.
	Value ListListContactsResponseDataFieldValueUnion `json:"value" api:"required"`
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
func (r ListListContactsResponseDataField) RawJSON() string { return r.JSON.raw }
func (r *ListListContactsResponseDataField) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ListListContactsResponseDataFieldValueUnion contains all possible properties and
// values from [string], [float64], [bool], [[]string],
// [ListListContactsResponseDataFieldValueAddress],
// [ListListContactsResponseDataFieldValueFullName].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfStringArray]
type ListListContactsResponseDataFieldValueUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]string] instead of an object.
	OfStringArray []string `json:",inline"`
	// This field is from variant [ListListContactsResponseDataFieldValueAddress].
	City string `json:"city"`
	// This field is from variant [ListListContactsResponseDataFieldValueAddress].
	Country string `json:"country"`
	// This field is from variant [ListListContactsResponseDataFieldValueAddress].
	Latitude float64 `json:"latitude"`
	// This field is from variant [ListListContactsResponseDataFieldValueAddress].
	Longitude float64 `json:"longitude"`
	// This field is from variant [ListListContactsResponseDataFieldValueAddress].
	PostalCode string `json:"postalCode"`
	// This field is from variant [ListListContactsResponseDataFieldValueAddress].
	State string `json:"state"`
	// This field is from variant [ListListContactsResponseDataFieldValueAddress].
	Street string `json:"street"`
	// This field is from variant [ListListContactsResponseDataFieldValueAddress].
	Street2 string `json:"street2"`
	// This field is from variant [ListListContactsResponseDataFieldValueFullName].
	FirstName string `json:"firstName"`
	// This field is from variant [ListListContactsResponseDataFieldValueFullName].
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

func (u ListListContactsResponseDataFieldValueUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ListListContactsResponseDataFieldValueUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ListListContactsResponseDataFieldValueUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ListListContactsResponseDataFieldValueUnion) AsStringArray() (v []string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ListListContactsResponseDataFieldValueUnion) AsAddress() (v ListListContactsResponseDataFieldValueAddress) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ListListContactsResponseDataFieldValueUnion) AsFullName() (v ListListContactsResponseDataFieldValueFullName) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ListListContactsResponseDataFieldValueUnion) RawJSON() string { return u.JSON.raw }

func (r *ListListContactsResponseDataFieldValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ListListContactsResponseDataFieldValueAddress struct {
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
func (r ListListContactsResponseDataFieldValueAddress) RawJSON() string { return r.JSON.raw }
func (r *ListListContactsResponseDataFieldValueAddress) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ListListContactsResponseDataFieldValueFullName struct {
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
func (r ListListContactsResponseDataFieldValueFullName) RawJSON() string { return r.JSON.raw }
func (r *ListListContactsResponseDataFieldValueFullName) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ListListContactsResponseDataRelationship struct {
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
func (r ListListContactsResponseDataRelationship) RawJSON() string { return r.JSON.raw }
func (r *ListListContactsResponseDataRelationship) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ListListOpportunitiesResponse struct {
	// Array of entity objects for the current page.
	Data []ListListOpportunitiesResponseData `json:"data" api:"required"`
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
func (r ListListOpportunitiesResponse) RawJSON() string { return r.JSON.raw }
func (r *ListListOpportunitiesResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ListListOpportunitiesResponseData struct {
	// Unique identifier for the entity.
	ID string `json:"id" api:"required"`
	// ISO 8601 timestamp of when the entity was created.
	CreatedAt string `json:"createdAt" api:"required"`
	// Map of field names to their typed values. System fields are prefixed with `$`
	// (e.g. `$name`, `$email`); custom attributes use their bare slug.
	Fields map[string]ListListOpportunitiesResponseDataField `json:"fields" api:"required"`
	// URL to view the entity in the Lightfield web app, or null.
	HTTPLink string `json:"httpLink" api:"required"`
	// Map of relationship names to their associated entities. System relationships are
	// prefixed with `$` (e.g. `$owner`, `$contact`).
	Relationships map[string]ListListOpportunitiesResponseDataRelationship `json:"relationships" api:"required"`
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
func (r ListListOpportunitiesResponseData) RawJSON() string { return r.JSON.raw }
func (r *ListListOpportunitiesResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ListListOpportunitiesResponseDataField struct {
	// The field value, or null if unset.
	Value ListListOpportunitiesResponseDataFieldValueUnion `json:"value" api:"required"`
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
func (r ListListOpportunitiesResponseDataField) RawJSON() string { return r.JSON.raw }
func (r *ListListOpportunitiesResponseDataField) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ListListOpportunitiesResponseDataFieldValueUnion contains all possible
// properties and values from [string], [float64], [bool], [[]string],
// [ListListOpportunitiesResponseDataFieldValueAddress],
// [ListListOpportunitiesResponseDataFieldValueFullName].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfStringArray]
type ListListOpportunitiesResponseDataFieldValueUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]string] instead of an object.
	OfStringArray []string `json:",inline"`
	// This field is from variant [ListListOpportunitiesResponseDataFieldValueAddress].
	City string `json:"city"`
	// This field is from variant [ListListOpportunitiesResponseDataFieldValueAddress].
	Country string `json:"country"`
	// This field is from variant [ListListOpportunitiesResponseDataFieldValueAddress].
	Latitude float64 `json:"latitude"`
	// This field is from variant [ListListOpportunitiesResponseDataFieldValueAddress].
	Longitude float64 `json:"longitude"`
	// This field is from variant [ListListOpportunitiesResponseDataFieldValueAddress].
	PostalCode string `json:"postalCode"`
	// This field is from variant [ListListOpportunitiesResponseDataFieldValueAddress].
	State string `json:"state"`
	// This field is from variant [ListListOpportunitiesResponseDataFieldValueAddress].
	Street string `json:"street"`
	// This field is from variant [ListListOpportunitiesResponseDataFieldValueAddress].
	Street2 string `json:"street2"`
	// This field is from variant
	// [ListListOpportunitiesResponseDataFieldValueFullName].
	FirstName string `json:"firstName"`
	// This field is from variant
	// [ListListOpportunitiesResponseDataFieldValueFullName].
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

func (u ListListOpportunitiesResponseDataFieldValueUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ListListOpportunitiesResponseDataFieldValueUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ListListOpportunitiesResponseDataFieldValueUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ListListOpportunitiesResponseDataFieldValueUnion) AsStringArray() (v []string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ListListOpportunitiesResponseDataFieldValueUnion) AsAddress() (v ListListOpportunitiesResponseDataFieldValueAddress) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ListListOpportunitiesResponseDataFieldValueUnion) AsFullName() (v ListListOpportunitiesResponseDataFieldValueFullName) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ListListOpportunitiesResponseDataFieldValueUnion) RawJSON() string { return u.JSON.raw }

func (r *ListListOpportunitiesResponseDataFieldValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ListListOpportunitiesResponseDataFieldValueAddress struct {
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
func (r ListListOpportunitiesResponseDataFieldValueAddress) RawJSON() string { return r.JSON.raw }
func (r *ListListOpportunitiesResponseDataFieldValueAddress) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ListListOpportunitiesResponseDataFieldValueFullName struct {
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
func (r ListListOpportunitiesResponseDataFieldValueFullName) RawJSON() string { return r.JSON.raw }
func (r *ListListOpportunitiesResponseDataFieldValueFullName) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ListListOpportunitiesResponseDataRelationship struct {
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
func (r ListListOpportunitiesResponseDataRelationship) RawJSON() string { return r.JSON.raw }
func (r *ListListOpportunitiesResponseDataRelationship) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ListListResponse struct {
	// Array of list objects for the current page.
	Data []ListListResponseData `json:"data" api:"required"`
	// The object type, always `"list"`.
	Object string `json:"object" api:"required"`
	// Total number of lists matching the query.
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
func (r ListListResponse) RawJSON() string { return r.JSON.raw }
func (r *ListListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ListListResponseData struct {
	// Unique identifier for the list.
	ID string `json:"id" api:"required"`
	// ISO 8601 timestamp of when the list was created.
	CreatedAt string `json:"createdAt" api:"required"`
	// Map of field names to their typed values. System fields are prefixed with `$`
	// (e.g. `$name`, `$objectType`).
	Fields map[string]ListListResponseDataField `json:"fields" api:"required"`
	// URL to view the list in the Lightfield web app, or null.
	HTTPLink string `json:"httpLink" api:"required"`
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
func (r ListListResponseData) RawJSON() string { return r.JSON.raw }
func (r *ListListResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ListListResponseDataField struct {
	// The field value, or null if unset.
	Value ListListResponseDataFieldValueUnion `json:"value" api:"required"`
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
func (r ListListResponseDataField) RawJSON() string { return r.JSON.raw }
func (r *ListListResponseDataField) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ListListResponseDataFieldValueUnion contains all possible properties and values
// from [string], [float64], [bool], [[]string],
// [ListListResponseDataFieldValueAddress],
// [ListListResponseDataFieldValueFullName].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfStringArray]
type ListListResponseDataFieldValueUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]string] instead of an object.
	OfStringArray []string `json:",inline"`
	// This field is from variant [ListListResponseDataFieldValueAddress].
	City string `json:"city"`
	// This field is from variant [ListListResponseDataFieldValueAddress].
	Country string `json:"country"`
	// This field is from variant [ListListResponseDataFieldValueAddress].
	Latitude float64 `json:"latitude"`
	// This field is from variant [ListListResponseDataFieldValueAddress].
	Longitude float64 `json:"longitude"`
	// This field is from variant [ListListResponseDataFieldValueAddress].
	PostalCode string `json:"postalCode"`
	// This field is from variant [ListListResponseDataFieldValueAddress].
	State string `json:"state"`
	// This field is from variant [ListListResponseDataFieldValueAddress].
	Street string `json:"street"`
	// This field is from variant [ListListResponseDataFieldValueAddress].
	Street2 string `json:"street2"`
	// This field is from variant [ListListResponseDataFieldValueFullName].
	FirstName string `json:"firstName"`
	// This field is from variant [ListListResponseDataFieldValueFullName].
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

func (u ListListResponseDataFieldValueUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ListListResponseDataFieldValueUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ListListResponseDataFieldValueUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ListListResponseDataFieldValueUnion) AsStringArray() (v []string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ListListResponseDataFieldValueUnion) AsAddress() (v ListListResponseDataFieldValueAddress) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ListListResponseDataFieldValueUnion) AsFullName() (v ListListResponseDataFieldValueFullName) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ListListResponseDataFieldValueUnion) RawJSON() string { return u.JSON.raw }

func (r *ListListResponseDataFieldValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ListListResponseDataFieldValueAddress struct {
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
func (r ListListResponseDataFieldValueAddress) RawJSON() string { return r.JSON.raw }
func (r *ListListResponseDataFieldValueAddress) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ListListResponseDataFieldValueFullName struct {
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
func (r ListListResponseDataFieldValueFullName) RawJSON() string { return r.JSON.raw }
func (r *ListListResponseDataFieldValueFullName) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ListRetrieveResponse struct {
	// Unique identifier for the list.
	ID string `json:"id" api:"required"`
	// ISO 8601 timestamp of when the list was created.
	CreatedAt string `json:"createdAt" api:"required"`
	// Map of field names to their typed values. System fields are prefixed with `$`
	// (e.g. `$name`, `$objectType`).
	Fields map[string]ListRetrieveResponseField `json:"fields" api:"required"`
	// URL to view the list in the Lightfield web app, or null.
	HTTPLink string `json:"httpLink" api:"required"`
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
func (r ListRetrieveResponse) RawJSON() string { return r.JSON.raw }
func (r *ListRetrieveResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ListRetrieveResponseField struct {
	// The field value, or null if unset.
	Value ListRetrieveResponseFieldValueUnion `json:"value" api:"required"`
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
func (r ListRetrieveResponseField) RawJSON() string { return r.JSON.raw }
func (r *ListRetrieveResponseField) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ListRetrieveResponseFieldValueUnion contains all possible properties and values
// from [string], [float64], [bool], [[]string],
// [ListRetrieveResponseFieldValueAddress],
// [ListRetrieveResponseFieldValueFullName].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfStringArray]
type ListRetrieveResponseFieldValueUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]string] instead of an object.
	OfStringArray []string `json:",inline"`
	// This field is from variant [ListRetrieveResponseFieldValueAddress].
	City string `json:"city"`
	// This field is from variant [ListRetrieveResponseFieldValueAddress].
	Country string `json:"country"`
	// This field is from variant [ListRetrieveResponseFieldValueAddress].
	Latitude float64 `json:"latitude"`
	// This field is from variant [ListRetrieveResponseFieldValueAddress].
	Longitude float64 `json:"longitude"`
	// This field is from variant [ListRetrieveResponseFieldValueAddress].
	PostalCode string `json:"postalCode"`
	// This field is from variant [ListRetrieveResponseFieldValueAddress].
	State string `json:"state"`
	// This field is from variant [ListRetrieveResponseFieldValueAddress].
	Street string `json:"street"`
	// This field is from variant [ListRetrieveResponseFieldValueAddress].
	Street2 string `json:"street2"`
	// This field is from variant [ListRetrieveResponseFieldValueFullName].
	FirstName string `json:"firstName"`
	// This field is from variant [ListRetrieveResponseFieldValueFullName].
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

func (u ListRetrieveResponseFieldValueUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ListRetrieveResponseFieldValueUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ListRetrieveResponseFieldValueUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ListRetrieveResponseFieldValueUnion) AsStringArray() (v []string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ListRetrieveResponseFieldValueUnion) AsAddress() (v ListRetrieveResponseFieldValueAddress) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ListRetrieveResponseFieldValueUnion) AsFullName() (v ListRetrieveResponseFieldValueFullName) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ListRetrieveResponseFieldValueUnion) RawJSON() string { return u.JSON.raw }

func (r *ListRetrieveResponseFieldValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ListRetrieveResponseFieldValueAddress struct {
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
func (r ListRetrieveResponseFieldValueAddress) RawJSON() string { return r.JSON.raw }
func (r *ListRetrieveResponseFieldValueAddress) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ListRetrieveResponseFieldValueFullName struct {
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
func (r ListRetrieveResponseFieldValueFullName) RawJSON() string { return r.JSON.raw }
func (r *ListRetrieveResponseFieldValueFullName) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ListUpdateResponse struct {
	// Unique identifier for the list.
	ID string `json:"id" api:"required"`
	// ISO 8601 timestamp of when the list was created.
	CreatedAt string `json:"createdAt" api:"required"`
	// Map of field names to their typed values. System fields are prefixed with `$`
	// (e.g. `$name`, `$objectType`).
	Fields map[string]ListUpdateResponseField `json:"fields" api:"required"`
	// URL to view the list in the Lightfield web app, or null.
	HTTPLink string `json:"httpLink" api:"required"`
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
func (r ListUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *ListUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ListUpdateResponseField struct {
	// The field value, or null if unset.
	Value ListUpdateResponseFieldValueUnion `json:"value" api:"required"`
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
func (r ListUpdateResponseField) RawJSON() string { return r.JSON.raw }
func (r *ListUpdateResponseField) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ListUpdateResponseFieldValueUnion contains all possible properties and values
// from [string], [float64], [bool], [[]string],
// [ListUpdateResponseFieldValueAddress], [ListUpdateResponseFieldValueFullName].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfStringArray]
type ListUpdateResponseFieldValueUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]string] instead of an object.
	OfStringArray []string `json:",inline"`
	// This field is from variant [ListUpdateResponseFieldValueAddress].
	City string `json:"city"`
	// This field is from variant [ListUpdateResponseFieldValueAddress].
	Country string `json:"country"`
	// This field is from variant [ListUpdateResponseFieldValueAddress].
	Latitude float64 `json:"latitude"`
	// This field is from variant [ListUpdateResponseFieldValueAddress].
	Longitude float64 `json:"longitude"`
	// This field is from variant [ListUpdateResponseFieldValueAddress].
	PostalCode string `json:"postalCode"`
	// This field is from variant [ListUpdateResponseFieldValueAddress].
	State string `json:"state"`
	// This field is from variant [ListUpdateResponseFieldValueAddress].
	Street string `json:"street"`
	// This field is from variant [ListUpdateResponseFieldValueAddress].
	Street2 string `json:"street2"`
	// This field is from variant [ListUpdateResponseFieldValueFullName].
	FirstName string `json:"firstName"`
	// This field is from variant [ListUpdateResponseFieldValueFullName].
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

func (u ListUpdateResponseFieldValueUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ListUpdateResponseFieldValueUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ListUpdateResponseFieldValueUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ListUpdateResponseFieldValueUnion) AsStringArray() (v []string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ListUpdateResponseFieldValueUnion) AsAddress() (v ListUpdateResponseFieldValueAddress) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ListUpdateResponseFieldValueUnion) AsFullName() (v ListUpdateResponseFieldValueFullName) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ListUpdateResponseFieldValueUnion) RawJSON() string { return u.JSON.raw }

func (r *ListUpdateResponseFieldValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ListUpdateResponseFieldValueAddress struct {
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
func (r ListUpdateResponseFieldValueAddress) RawJSON() string { return r.JSON.raw }
func (r *ListUpdateResponseFieldValueAddress) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ListUpdateResponseFieldValueFullName struct {
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
func (r ListUpdateResponseFieldValueFullName) RawJSON() string { return r.JSON.raw }
func (r *ListUpdateResponseFieldValueFullName) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ListNewParams struct {
	// Field values for the new list. Required: `$name` (string) and `$objectType`.
	Fields ListNewParamsFields `json:"fields,omitzero" api:"required"`
	// Relationships to set on the new list.
	Relationships ListNewParamsRelationshipsUnion `json:"relationships,omitzero"`
	paramObj
}

func (r ListNewParams) MarshalJSON() (data []byte, err error) {
	type shadow ListNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ListNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Field values for the new list. Required: `$name` (string) and `$objectType`.
//
// The properties Name, ObjectType are required.
type ListNewParamsFields struct {
	// Display name of the list.
	Name string `json:"$name" api:"required"`
	// The type of entities this list contains. One of `account`, `contact`, or
	// `opportunity`.
	//
	// Any of "account", "contact", "opportunity".
	ObjectType string `json:"$objectType,omitzero" api:"required"`
	paramObj
}

func (r ListNewParamsFields) MarshalJSON() (data []byte, err error) {
	type shadow ListNewParamsFields
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ListNewParamsFields) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ListNewParamsFields](
		"$objectType", "account", "contact", "opportunity",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ListNewParamsRelationshipsUnion struct {
	OfListNewsRelationshipsAccounts      *ListNewParamsRelationshipsAccounts      `json:",omitzero,inline"`
	OfListNewsRelationshipsContacts      *ListNewParamsRelationshipsContacts      `json:",omitzero,inline"`
	OfListNewsRelationshipsOpportunities *ListNewParamsRelationshipsOpportunities `json:",omitzero,inline"`
	paramUnion
}

func (u ListNewParamsRelationshipsUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfListNewsRelationshipsAccounts, u.OfListNewsRelationshipsContacts, u.OfListNewsRelationshipsOpportunities)
}
func (u *ListNewParamsRelationshipsUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// The property Accounts is required.
type ListNewParamsRelationshipsAccounts struct {
	// Account ID(s) to add as initial members. List `$objectType` must be `account`.
	Accounts ListNewParamsRelationshipsAccountsAccountsUnion `json:"$accounts,omitzero" api:"required"`
	paramObj
}

func (r ListNewParamsRelationshipsAccounts) MarshalJSON() (data []byte, err error) {
	type shadow ListNewParamsRelationshipsAccounts
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ListNewParamsRelationshipsAccounts) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ListNewParamsRelationshipsAccountsAccountsUnion struct {
	OfString      param.Opt[string] `json:",omitzero,inline"`
	OfStringArray []string          `json:",omitzero,inline"`
	paramUnion
}

func (u ListNewParamsRelationshipsAccountsAccountsUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray)
}
func (u *ListNewParamsRelationshipsAccountsAccountsUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// The property Contacts is required.
type ListNewParamsRelationshipsContacts struct {
	// Contact ID(s) to add as initial members. List `$objectType` must be `contact`.
	Contacts ListNewParamsRelationshipsContactsContactsUnion `json:"$contacts,omitzero" api:"required"`
	paramObj
}

func (r ListNewParamsRelationshipsContacts) MarshalJSON() (data []byte, err error) {
	type shadow ListNewParamsRelationshipsContacts
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ListNewParamsRelationshipsContacts) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ListNewParamsRelationshipsContactsContactsUnion struct {
	OfString      param.Opt[string] `json:",omitzero,inline"`
	OfStringArray []string          `json:",omitzero,inline"`
	paramUnion
}

func (u ListNewParamsRelationshipsContactsContactsUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray)
}
func (u *ListNewParamsRelationshipsContactsContactsUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// The property Opportunities is required.
type ListNewParamsRelationshipsOpportunities struct {
	// Opportunity ID(s) to add as initial members. List `$objectType` must be
	// `opportunity`.
	Opportunities ListNewParamsRelationshipsOpportunitiesOpportunitiesUnion `json:"$opportunities,omitzero" api:"required"`
	paramObj
}

func (r ListNewParamsRelationshipsOpportunities) MarshalJSON() (data []byte, err error) {
	type shadow ListNewParamsRelationshipsOpportunities
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ListNewParamsRelationshipsOpportunities) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ListNewParamsRelationshipsOpportunitiesOpportunitiesUnion struct {
	OfString      param.Opt[string] `json:",omitzero,inline"`
	OfStringArray []string          `json:",omitzero,inline"`
	paramUnion
}

func (u ListNewParamsRelationshipsOpportunitiesOpportunitiesUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray)
}
func (u *ListNewParamsRelationshipsOpportunitiesOpportunitiesUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

type ListUpdateParams struct {
	// Field values to update — only provided fields are modified; omitted fields are
	// left unchanged.
	Fields ListUpdateParamsFields `json:"fields,omitzero"`
	// Relationship operations. Use the key matching the list's `$objectType` (e.g.
	// `$accounts` for an account list).
	Relationships ListUpdateParamsRelationshipsUnion `json:"relationships,omitzero"`
	paramObj
}

func (r ListUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow ListUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ListUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Field values to update — only provided fields are modified; omitted fields are
// left unchanged.
type ListUpdateParamsFields struct {
	// Display name of the list.
	Name param.Opt[string] `json:"$name,omitzero"`
	paramObj
}

func (r ListUpdateParamsFields) MarshalJSON() (data []byte, err error) {
	type shadow ListUpdateParamsFields
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ListUpdateParamsFields) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ListUpdateParamsRelationshipsUnion struct {
	OfListUpdatesRelationshipsAccounts      *ListUpdateParamsRelationshipsAccounts      `json:",omitzero,inline"`
	OfListUpdatesRelationshipsContacts      *ListUpdateParamsRelationshipsContacts      `json:",omitzero,inline"`
	OfListUpdatesRelationshipsOpportunities *ListUpdateParamsRelationshipsOpportunities `json:",omitzero,inline"`
	paramUnion
}

func (u ListUpdateParamsRelationshipsUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfListUpdatesRelationshipsAccounts, u.OfListUpdatesRelationshipsContacts, u.OfListUpdatesRelationshipsOpportunities)
}
func (u *ListUpdateParamsRelationshipsUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// The property Accounts is required.
type ListUpdateParamsRelationshipsAccounts struct {
	// Add/remove accounts. List `$objectType` must be `account`.
	Accounts ListUpdateParamsRelationshipsAccountsAccounts `json:"$accounts,omitzero" api:"required"`
	paramObj
}

func (r ListUpdateParamsRelationshipsAccounts) MarshalJSON() (data []byte, err error) {
	type shadow ListUpdateParamsRelationshipsAccounts
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ListUpdateParamsRelationshipsAccounts) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Add/remove accounts. List `$objectType` must be `account`.
type ListUpdateParamsRelationshipsAccountsAccounts struct {
	// Entity ID(s) to add to the list.
	Add ListUpdateParamsRelationshipsAccountsAccountsAddUnion `json:"add,omitzero"`
	// Entity ID(s) to remove from the list.
	Remove ListUpdateParamsRelationshipsAccountsAccountsRemoveUnion `json:"remove,omitzero"`
	paramObj
}

func (r ListUpdateParamsRelationshipsAccountsAccounts) MarshalJSON() (data []byte, err error) {
	type shadow ListUpdateParamsRelationshipsAccountsAccounts
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ListUpdateParamsRelationshipsAccountsAccounts) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ListUpdateParamsRelationshipsAccountsAccountsAddUnion struct {
	OfString      param.Opt[string] `json:",omitzero,inline"`
	OfStringArray []string          `json:",omitzero,inline"`
	paramUnion
}

func (u ListUpdateParamsRelationshipsAccountsAccountsAddUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray)
}
func (u *ListUpdateParamsRelationshipsAccountsAccountsAddUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ListUpdateParamsRelationshipsAccountsAccountsRemoveUnion struct {
	OfString      param.Opt[string] `json:",omitzero,inline"`
	OfStringArray []string          `json:",omitzero,inline"`
	paramUnion
}

func (u ListUpdateParamsRelationshipsAccountsAccountsRemoveUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray)
}
func (u *ListUpdateParamsRelationshipsAccountsAccountsRemoveUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// The property Contacts is required.
type ListUpdateParamsRelationshipsContacts struct {
	// Add/remove contacts. List `$objectType` must be `contact`.
	Contacts ListUpdateParamsRelationshipsContactsContacts `json:"$contacts,omitzero" api:"required"`
	paramObj
}

func (r ListUpdateParamsRelationshipsContacts) MarshalJSON() (data []byte, err error) {
	type shadow ListUpdateParamsRelationshipsContacts
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ListUpdateParamsRelationshipsContacts) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Add/remove contacts. List `$objectType` must be `contact`.
type ListUpdateParamsRelationshipsContactsContacts struct {
	// Entity ID(s) to add to the list.
	Add ListUpdateParamsRelationshipsContactsContactsAddUnion `json:"add,omitzero"`
	// Entity ID(s) to remove from the list.
	Remove ListUpdateParamsRelationshipsContactsContactsRemoveUnion `json:"remove,omitzero"`
	paramObj
}

func (r ListUpdateParamsRelationshipsContactsContacts) MarshalJSON() (data []byte, err error) {
	type shadow ListUpdateParamsRelationshipsContactsContacts
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ListUpdateParamsRelationshipsContactsContacts) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ListUpdateParamsRelationshipsContactsContactsAddUnion struct {
	OfString      param.Opt[string] `json:",omitzero,inline"`
	OfStringArray []string          `json:",omitzero,inline"`
	paramUnion
}

func (u ListUpdateParamsRelationshipsContactsContactsAddUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray)
}
func (u *ListUpdateParamsRelationshipsContactsContactsAddUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ListUpdateParamsRelationshipsContactsContactsRemoveUnion struct {
	OfString      param.Opt[string] `json:",omitzero,inline"`
	OfStringArray []string          `json:",omitzero,inline"`
	paramUnion
}

func (u ListUpdateParamsRelationshipsContactsContactsRemoveUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray)
}
func (u *ListUpdateParamsRelationshipsContactsContactsRemoveUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// The property Opportunities is required.
type ListUpdateParamsRelationshipsOpportunities struct {
	// Add/remove opportunities. List `$objectType` must be `opportunity`.
	Opportunities ListUpdateParamsRelationshipsOpportunitiesOpportunities `json:"$opportunities,omitzero" api:"required"`
	paramObj
}

func (r ListUpdateParamsRelationshipsOpportunities) MarshalJSON() (data []byte, err error) {
	type shadow ListUpdateParamsRelationshipsOpportunities
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ListUpdateParamsRelationshipsOpportunities) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Add/remove opportunities. List `$objectType` must be `opportunity`.
type ListUpdateParamsRelationshipsOpportunitiesOpportunities struct {
	// Entity ID(s) to add to the list.
	Add ListUpdateParamsRelationshipsOpportunitiesOpportunitiesAddUnion `json:"add,omitzero"`
	// Entity ID(s) to remove from the list.
	Remove ListUpdateParamsRelationshipsOpportunitiesOpportunitiesRemoveUnion `json:"remove,omitzero"`
	paramObj
}

func (r ListUpdateParamsRelationshipsOpportunitiesOpportunities) MarshalJSON() (data []byte, err error) {
	type shadow ListUpdateParamsRelationshipsOpportunitiesOpportunities
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ListUpdateParamsRelationshipsOpportunitiesOpportunities) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ListUpdateParamsRelationshipsOpportunitiesOpportunitiesAddUnion struct {
	OfString      param.Opt[string] `json:",omitzero,inline"`
	OfStringArray []string          `json:",omitzero,inline"`
	paramUnion
}

func (u ListUpdateParamsRelationshipsOpportunitiesOpportunitiesAddUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray)
}
func (u *ListUpdateParamsRelationshipsOpportunitiesOpportunitiesAddUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ListUpdateParamsRelationshipsOpportunitiesOpportunitiesRemoveUnion struct {
	OfString      param.Opt[string] `json:",omitzero,inline"`
	OfStringArray []string          `json:",omitzero,inline"`
	paramUnion
}

func (u ListUpdateParamsRelationshipsOpportunitiesOpportunitiesRemoveUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray)
}
func (u *ListUpdateParamsRelationshipsOpportunitiesOpportunitiesRemoveUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

type ListListParams struct {
	// Maximum number of records to return. Defaults to 25, maximum 25.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Number of records to skip for pagination. Defaults to 0.
	Offset param.Opt[int64] `query:"offset,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ListListParams]'s query parameters as `url.Values`.
func (r ListListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ListListAccountsParams struct {
	// Maximum number of records to return. Defaults to 25, maximum 25.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Number of records to skip for pagination. Defaults to 0.
	Offset param.Opt[int64] `query:"offset,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ListListAccountsParams]'s query parameters as `url.Values`.
func (r ListListAccountsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ListListContactsParams struct {
	// Maximum number of records to return. Defaults to 25, maximum 25.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Number of records to skip for pagination. Defaults to 0.
	Offset param.Opt[int64] `query:"offset,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ListListContactsParams]'s query parameters as `url.Values`.
func (r ListListContactsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ListListOpportunitiesParams struct {
	// Maximum number of records to return. Defaults to 25, maximum 25.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Number of records to skip for pagination. Defaults to 0.
	Offset param.Opt[int64] `query:"offset,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ListListOpportunitiesParams]'s query parameters as
// `url.Values`.
func (r ListListOpportunitiesParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
