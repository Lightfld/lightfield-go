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
	"github.com/Lightfld/lightfield-go/internal/requestconfig"
	"github.com/Lightfld/lightfield-go/option"
	"github.com/Lightfld/lightfield-go/packages/param"
	"github.com/Lightfld/lightfield-go/packages/respjson"
)

// Merge operations combine two records into one.
//
// MergeService contains methods and other services that help with interacting with
// the Lightfield API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewMergeService] method instead.
type MergeService struct {
	Options []option.RequestOption
}

// NewMergeService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewMergeService(opts ...option.RequestOption) (r MergeService) {
	r = MergeService{}
	r.Options = opts
	return
}

// Returns the status and details of a merge operation by its ID.
//
// **[Rate limit category](/using-the-api/rate-limits/):** Read
func (r *MergeService) GetMerge(ctx context.Context, id string, opts ...option.RequestOption) (res *MergeGetMergeResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/merges/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Merges two accounts into one. The primary account retains its ID; the duplicate
// is soft-deleted.
//
// **[Required scopes](/using-the-api/scopes/):** `accounts:update` +
// `accounts:delete`
//
// **[Rate limit category](/using-the-api/rate-limits/):** Write
func (r *MergeService) MergeAccounts(ctx context.Context, body MergeMergeAccountsParams, opts ...option.RequestOption) (res *MergeMergeAccountsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/accounts/merge"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Merges two contacts into one. The primary contact retains its ID; the duplicate
// is soft-deleted.
//
// **[Required scopes](/using-the-api/scopes/):** `contacts:update` +
// `contacts:delete`
//
// **[Rate limit category](/using-the-api/rate-limits/):** Write
func (r *MergeService) MergeContacts(ctx context.Context, body MergeMergeContactsParams, opts ...option.RequestOption) (res *MergeMergeContactsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/contacts/merge"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Merges two records of the specified custom object type into one. The primary
// record retains its ID; the duplicate is soft-deleted. Both records must belong
// to the custom object type named in the path.
//
// **[Required scopes](/using-the-api/scopes/):** `custom_objects:update` +
// `custom_objects:delete`
//
// **[Rate limit category](/using-the-api/rate-limits/):** Write
func (r *MergeService) MergeObjectValues(ctx context.Context, entitySlug string, body MergeMergeObjectValuesParams, opts ...option.RequestOption) (res *MergeMergeObjectValuesResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if entitySlug == "" {
		err = errors.New("missing required entitySlug parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/objects/%s/merge", url.PathEscape(entitySlug))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Merges two opportunities into one. The primary opportunity retains its ID; the
// duplicate is soft-deleted. Both opportunities must belong to the same account.
//
// **[Required scopes](/using-the-api/scopes/):** `opportunities:update` +
// `opportunities:delete`
//
// **[Rate limit category](/using-the-api/rate-limits/):** Write
func (r *MergeService) MergeOpportunities(ctx context.Context, body MergeMergeOpportunitiesParams, opts ...option.RequestOption) (res *MergeMergeOpportunitiesResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/opportunities/merge"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

type MergeGetMergeResponse struct {
	// Unique identifier for the merge operation.
	ID string `json:"id" api:"required"`
	// ISO 8601 timestamp of when the merge was created.
	CreatedAt string `json:"createdAt" api:"required"`
	// ID of the record that was merged into the primary and soft-deleted.
	DuplicateID string `json:"duplicateId" api:"required"`
	// The object type of the merged records (e.g. `account`, `contact`, `opportunity`,
	// or a custom object slug).
	EntityType string `json:"entityType" api:"required"`
	// ID of the record that was kept (the primary).
	PrimaryID string `json:"primaryId" api:"required"`
	// Current status of the merge: `cleanup_pending`, `done`, or `failed`.
	Status string `json:"status" api:"required"`
	// ISO 8601 timestamp of when the merge was last updated.
	UpdatedAt string `json:"updatedAt" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		DuplicateID respjson.Field
		EntityType  respjson.Field
		PrimaryID   respjson.Field
		Status      respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MergeGetMergeResponse) RawJSON() string { return r.JSON.raw }
func (r *MergeGetMergeResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MergeMergeAccountsResponse struct {
	Merge   MergeMergeAccountsResponseMerge   `json:"merge" api:"required"`
	Primary MergeMergeAccountsResponsePrimary `json:"primary" api:"required"`
	Summary MergeMergeAccountsResponseSummary `json:"summary" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Merge       respjson.Field
		Primary     respjson.Field
		Summary     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MergeMergeAccountsResponse) RawJSON() string { return r.JSON.raw }
func (r *MergeMergeAccountsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MergeMergeAccountsResponseMerge struct {
	// Unique identifier for the merge operation.
	ID string `json:"id" api:"required"`
	// Current status of the merge: `cleanup_pending`, `done`, or `failed`.
	Status string `json:"status" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MergeMergeAccountsResponseMerge) RawJSON() string { return r.JSON.raw }
func (r *MergeMergeAccountsResponseMerge) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MergeMergeAccountsResponsePrimary struct {
	// Unique identifier for the entity.
	ID string `json:"id" api:"required"`
	// ISO 8601 timestamp of when the entity was created.
	CreatedAt string `json:"createdAt" api:"required"`
	// Map of field names to their typed values. System fields are prefixed with `$`
	// (e.g. `$name`, `$email`); custom attributes use their bare slug.
	Fields map[string]MergeMergeAccountsResponsePrimaryField `json:"fields" api:"required"`
	// URL to view the entity in the Lightfield web app, or null.
	HTTPLink string `json:"httpLink" api:"required"`
	// Map of relationship names to their associated entities. System relationships are
	// prefixed with `$` (e.g. `$owner`, `$contact`).
	Relationships map[string]MergeMergeAccountsResponsePrimaryRelationship `json:"relationships" api:"required"`
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
func (r MergeMergeAccountsResponsePrimary) RawJSON() string { return r.JSON.raw }
func (r *MergeMergeAccountsResponsePrimary) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MergeMergeAccountsResponsePrimaryField struct {
	// The field value, or null if unset.
	Value MergeMergeAccountsResponsePrimaryFieldValueUnion `json:"value" api:"required"`
	// The data type of the field.
	//
	// Any of "ADDRESS", "CHECKBOX", "CURRENCY", "DATETIME", "EMAIL", "FULL_NAME",
	// "MARKDOWN", "MULTI_SELECT", "NUMBER", "SINGLE_SELECT", "SOCIAL_HANDLE",
	// "TELEPHONE", "TEXT", "URL", "HTML".
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
func (r MergeMergeAccountsResponsePrimaryField) RawJSON() string { return r.JSON.raw }
func (r *MergeMergeAccountsResponsePrimaryField) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// MergeMergeAccountsResponsePrimaryFieldValueUnion contains all possible
// properties and values from [string], [float64], [bool], [[]string],
// [MergeMergeAccountsResponsePrimaryFieldValueAddress],
// [MergeMergeAccountsResponsePrimaryFieldValueFullName].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfStringArray]
type MergeMergeAccountsResponsePrimaryFieldValueUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]string] instead of an object.
	OfStringArray []string `json:",inline"`
	// This field is from variant [MergeMergeAccountsResponsePrimaryFieldValueAddress].
	City string `json:"city"`
	// This field is from variant [MergeMergeAccountsResponsePrimaryFieldValueAddress].
	Country string `json:"country"`
	// This field is from variant [MergeMergeAccountsResponsePrimaryFieldValueAddress].
	Latitude float64 `json:"latitude"`
	// This field is from variant [MergeMergeAccountsResponsePrimaryFieldValueAddress].
	Longitude float64 `json:"longitude"`
	// This field is from variant [MergeMergeAccountsResponsePrimaryFieldValueAddress].
	PostalCode string `json:"postalCode"`
	// This field is from variant [MergeMergeAccountsResponsePrimaryFieldValueAddress].
	State string `json:"state"`
	// This field is from variant [MergeMergeAccountsResponsePrimaryFieldValueAddress].
	Street string `json:"street"`
	// This field is from variant [MergeMergeAccountsResponsePrimaryFieldValueAddress].
	Street2 string `json:"street2"`
	// This field is from variant
	// [MergeMergeAccountsResponsePrimaryFieldValueFullName].
	FirstName string `json:"firstName"`
	// This field is from variant
	// [MergeMergeAccountsResponsePrimaryFieldValueFullName].
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

func (u MergeMergeAccountsResponsePrimaryFieldValueUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MergeMergeAccountsResponsePrimaryFieldValueUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MergeMergeAccountsResponsePrimaryFieldValueUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MergeMergeAccountsResponsePrimaryFieldValueUnion) AsStringArray() (v []string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MergeMergeAccountsResponsePrimaryFieldValueUnion) AsAddress() (v MergeMergeAccountsResponsePrimaryFieldValueAddress) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MergeMergeAccountsResponsePrimaryFieldValueUnion) AsFullName() (v MergeMergeAccountsResponsePrimaryFieldValueFullName) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u MergeMergeAccountsResponsePrimaryFieldValueUnion) RawJSON() string { return u.JSON.raw }

func (r *MergeMergeAccountsResponsePrimaryFieldValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MergeMergeAccountsResponsePrimaryFieldValueAddress struct {
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
func (r MergeMergeAccountsResponsePrimaryFieldValueAddress) RawJSON() string { return r.JSON.raw }
func (r *MergeMergeAccountsResponsePrimaryFieldValueAddress) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MergeMergeAccountsResponsePrimaryFieldValueFullName struct {
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
func (r MergeMergeAccountsResponsePrimaryFieldValueFullName) RawJSON() string { return r.JSON.raw }
func (r *MergeMergeAccountsResponsePrimaryFieldValueFullName) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MergeMergeAccountsResponsePrimaryRelationship struct {
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
func (r MergeMergeAccountsResponsePrimaryRelationship) RawJSON() string { return r.JSON.raw }
func (r *MergeMergeAccountsResponsePrimaryRelationship) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MergeMergeAccountsResponseSummary struct {
	// Number of attribute fields written onto the primary record.
	FieldWriteCount int64 `json:"fieldWriteCount" api:"required"`
	// Number of related records re-pointed from the duplicate to the primary.
	SyncRepointedCount int64 `json:"syncRepointedCount" api:"required"`
	// Non-fatal warnings from the merge (e.g. skipped transfers).
	Warnings []string `json:"warnings" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FieldWriteCount    respjson.Field
		SyncRepointedCount respjson.Field
		Warnings           respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MergeMergeAccountsResponseSummary) RawJSON() string { return r.JSON.raw }
func (r *MergeMergeAccountsResponseSummary) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MergeMergeContactsResponse struct {
	Merge   MergeMergeContactsResponseMerge   `json:"merge" api:"required"`
	Primary MergeMergeContactsResponsePrimary `json:"primary" api:"required"`
	Summary MergeMergeContactsResponseSummary `json:"summary" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Merge       respjson.Field
		Primary     respjson.Field
		Summary     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MergeMergeContactsResponse) RawJSON() string { return r.JSON.raw }
func (r *MergeMergeContactsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MergeMergeContactsResponseMerge struct {
	// Unique identifier for the merge operation.
	ID string `json:"id" api:"required"`
	// Current status of the merge: `cleanup_pending`, `done`, or `failed`.
	Status string `json:"status" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MergeMergeContactsResponseMerge) RawJSON() string { return r.JSON.raw }
func (r *MergeMergeContactsResponseMerge) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MergeMergeContactsResponsePrimary struct {
	// Unique identifier for the entity.
	ID string `json:"id" api:"required"`
	// ISO 8601 timestamp of when the entity was created.
	CreatedAt string `json:"createdAt" api:"required"`
	// Map of field names to their typed values. System fields are prefixed with `$`
	// (e.g. `$name`, `$email`); custom attributes use their bare slug.
	Fields map[string]MergeMergeContactsResponsePrimaryField `json:"fields" api:"required"`
	// URL to view the entity in the Lightfield web app, or null.
	HTTPLink string `json:"httpLink" api:"required"`
	// Map of relationship names to their associated entities. System relationships are
	// prefixed with `$` (e.g. `$owner`, `$contact`).
	Relationships map[string]MergeMergeContactsResponsePrimaryRelationship `json:"relationships" api:"required"`
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
func (r MergeMergeContactsResponsePrimary) RawJSON() string { return r.JSON.raw }
func (r *MergeMergeContactsResponsePrimary) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MergeMergeContactsResponsePrimaryField struct {
	// The field value, or null if unset.
	Value MergeMergeContactsResponsePrimaryFieldValueUnion `json:"value" api:"required"`
	// The data type of the field.
	//
	// Any of "ADDRESS", "CHECKBOX", "CURRENCY", "DATETIME", "EMAIL", "FULL_NAME",
	// "MARKDOWN", "MULTI_SELECT", "NUMBER", "SINGLE_SELECT", "SOCIAL_HANDLE",
	// "TELEPHONE", "TEXT", "URL", "HTML".
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
func (r MergeMergeContactsResponsePrimaryField) RawJSON() string { return r.JSON.raw }
func (r *MergeMergeContactsResponsePrimaryField) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// MergeMergeContactsResponsePrimaryFieldValueUnion contains all possible
// properties and values from [string], [float64], [bool], [[]string],
// [MergeMergeContactsResponsePrimaryFieldValueAddress],
// [MergeMergeContactsResponsePrimaryFieldValueFullName].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfStringArray]
type MergeMergeContactsResponsePrimaryFieldValueUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]string] instead of an object.
	OfStringArray []string `json:",inline"`
	// This field is from variant [MergeMergeContactsResponsePrimaryFieldValueAddress].
	City string `json:"city"`
	// This field is from variant [MergeMergeContactsResponsePrimaryFieldValueAddress].
	Country string `json:"country"`
	// This field is from variant [MergeMergeContactsResponsePrimaryFieldValueAddress].
	Latitude float64 `json:"latitude"`
	// This field is from variant [MergeMergeContactsResponsePrimaryFieldValueAddress].
	Longitude float64 `json:"longitude"`
	// This field is from variant [MergeMergeContactsResponsePrimaryFieldValueAddress].
	PostalCode string `json:"postalCode"`
	// This field is from variant [MergeMergeContactsResponsePrimaryFieldValueAddress].
	State string `json:"state"`
	// This field is from variant [MergeMergeContactsResponsePrimaryFieldValueAddress].
	Street string `json:"street"`
	// This field is from variant [MergeMergeContactsResponsePrimaryFieldValueAddress].
	Street2 string `json:"street2"`
	// This field is from variant
	// [MergeMergeContactsResponsePrimaryFieldValueFullName].
	FirstName string `json:"firstName"`
	// This field is from variant
	// [MergeMergeContactsResponsePrimaryFieldValueFullName].
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

func (u MergeMergeContactsResponsePrimaryFieldValueUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MergeMergeContactsResponsePrimaryFieldValueUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MergeMergeContactsResponsePrimaryFieldValueUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MergeMergeContactsResponsePrimaryFieldValueUnion) AsStringArray() (v []string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MergeMergeContactsResponsePrimaryFieldValueUnion) AsAddress() (v MergeMergeContactsResponsePrimaryFieldValueAddress) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MergeMergeContactsResponsePrimaryFieldValueUnion) AsFullName() (v MergeMergeContactsResponsePrimaryFieldValueFullName) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u MergeMergeContactsResponsePrimaryFieldValueUnion) RawJSON() string { return u.JSON.raw }

func (r *MergeMergeContactsResponsePrimaryFieldValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MergeMergeContactsResponsePrimaryFieldValueAddress struct {
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
func (r MergeMergeContactsResponsePrimaryFieldValueAddress) RawJSON() string { return r.JSON.raw }
func (r *MergeMergeContactsResponsePrimaryFieldValueAddress) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MergeMergeContactsResponsePrimaryFieldValueFullName struct {
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
func (r MergeMergeContactsResponsePrimaryFieldValueFullName) RawJSON() string { return r.JSON.raw }
func (r *MergeMergeContactsResponsePrimaryFieldValueFullName) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MergeMergeContactsResponsePrimaryRelationship struct {
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
func (r MergeMergeContactsResponsePrimaryRelationship) RawJSON() string { return r.JSON.raw }
func (r *MergeMergeContactsResponsePrimaryRelationship) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MergeMergeContactsResponseSummary struct {
	// Number of attribute fields written onto the primary record.
	FieldWriteCount int64 `json:"fieldWriteCount" api:"required"`
	// Number of related records re-pointed from the duplicate to the primary.
	SyncRepointedCount int64 `json:"syncRepointedCount" api:"required"`
	// Non-fatal warnings from the merge (e.g. skipped transfers).
	Warnings []string `json:"warnings" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FieldWriteCount    respjson.Field
		SyncRepointedCount respjson.Field
		Warnings           respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MergeMergeContactsResponseSummary) RawJSON() string { return r.JSON.raw }
func (r *MergeMergeContactsResponseSummary) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MergeMergeObjectValuesResponse struct {
	Merge   MergeMergeObjectValuesResponseMerge   `json:"merge" api:"required"`
	Primary MergeMergeObjectValuesResponsePrimary `json:"primary" api:"required"`
	Summary MergeMergeObjectValuesResponseSummary `json:"summary" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Merge       respjson.Field
		Primary     respjson.Field
		Summary     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MergeMergeObjectValuesResponse) RawJSON() string { return r.JSON.raw }
func (r *MergeMergeObjectValuesResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MergeMergeObjectValuesResponseMerge struct {
	// Unique identifier for the merge operation.
	ID string `json:"id" api:"required"`
	// Current status of the merge: `cleanup_pending`, `done`, or `failed`.
	Status string `json:"status" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MergeMergeObjectValuesResponseMerge) RawJSON() string { return r.JSON.raw }
func (r *MergeMergeObjectValuesResponseMerge) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MergeMergeObjectValuesResponsePrimary struct {
	// Unique identifier for the entity.
	ID string `json:"id" api:"required"`
	// ISO 8601 timestamp of when the entity was created.
	CreatedAt string `json:"createdAt" api:"required"`
	// Map of field names to their typed values. System fields are prefixed with `$`
	// (e.g. `$name`, `$email`); custom attributes use their bare slug.
	Fields map[string]MergeMergeObjectValuesResponsePrimaryField `json:"fields" api:"required"`
	// URL to view the entity in the Lightfield web app, or null.
	HTTPLink string `json:"httpLink" api:"required"`
	// Map of relationship names to their associated entities. System relationships are
	// prefixed with `$` (e.g. `$owner`, `$contact`).
	Relationships map[string]MergeMergeObjectValuesResponsePrimaryRelationship `json:"relationships" api:"required"`
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
func (r MergeMergeObjectValuesResponsePrimary) RawJSON() string { return r.JSON.raw }
func (r *MergeMergeObjectValuesResponsePrimary) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MergeMergeObjectValuesResponsePrimaryField struct {
	// The field value, or null if unset.
	Value MergeMergeObjectValuesResponsePrimaryFieldValueUnion `json:"value" api:"required"`
	// The data type of the field.
	//
	// Any of "ADDRESS", "CHECKBOX", "CURRENCY", "DATETIME", "EMAIL", "FULL_NAME",
	// "MARKDOWN", "MULTI_SELECT", "NUMBER", "SINGLE_SELECT", "SOCIAL_HANDLE",
	// "TELEPHONE", "TEXT", "URL", "HTML".
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
func (r MergeMergeObjectValuesResponsePrimaryField) RawJSON() string { return r.JSON.raw }
func (r *MergeMergeObjectValuesResponsePrimaryField) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// MergeMergeObjectValuesResponsePrimaryFieldValueUnion contains all possible
// properties and values from [string], [float64], [bool], [[]string],
// [MergeMergeObjectValuesResponsePrimaryFieldValueAddress],
// [MergeMergeObjectValuesResponsePrimaryFieldValueFullName].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfStringArray]
type MergeMergeObjectValuesResponsePrimaryFieldValueUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]string] instead of an object.
	OfStringArray []string `json:",inline"`
	// This field is from variant
	// [MergeMergeObjectValuesResponsePrimaryFieldValueAddress].
	City string `json:"city"`
	// This field is from variant
	// [MergeMergeObjectValuesResponsePrimaryFieldValueAddress].
	Country string `json:"country"`
	// This field is from variant
	// [MergeMergeObjectValuesResponsePrimaryFieldValueAddress].
	Latitude float64 `json:"latitude"`
	// This field is from variant
	// [MergeMergeObjectValuesResponsePrimaryFieldValueAddress].
	Longitude float64 `json:"longitude"`
	// This field is from variant
	// [MergeMergeObjectValuesResponsePrimaryFieldValueAddress].
	PostalCode string `json:"postalCode"`
	// This field is from variant
	// [MergeMergeObjectValuesResponsePrimaryFieldValueAddress].
	State string `json:"state"`
	// This field is from variant
	// [MergeMergeObjectValuesResponsePrimaryFieldValueAddress].
	Street string `json:"street"`
	// This field is from variant
	// [MergeMergeObjectValuesResponsePrimaryFieldValueAddress].
	Street2 string `json:"street2"`
	// This field is from variant
	// [MergeMergeObjectValuesResponsePrimaryFieldValueFullName].
	FirstName string `json:"firstName"`
	// This field is from variant
	// [MergeMergeObjectValuesResponsePrimaryFieldValueFullName].
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

func (u MergeMergeObjectValuesResponsePrimaryFieldValueUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MergeMergeObjectValuesResponsePrimaryFieldValueUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MergeMergeObjectValuesResponsePrimaryFieldValueUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MergeMergeObjectValuesResponsePrimaryFieldValueUnion) AsStringArray() (v []string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MergeMergeObjectValuesResponsePrimaryFieldValueUnion) AsAddress() (v MergeMergeObjectValuesResponsePrimaryFieldValueAddress) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MergeMergeObjectValuesResponsePrimaryFieldValueUnion) AsFullName() (v MergeMergeObjectValuesResponsePrimaryFieldValueFullName) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u MergeMergeObjectValuesResponsePrimaryFieldValueUnion) RawJSON() string { return u.JSON.raw }

func (r *MergeMergeObjectValuesResponsePrimaryFieldValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MergeMergeObjectValuesResponsePrimaryFieldValueAddress struct {
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
func (r MergeMergeObjectValuesResponsePrimaryFieldValueAddress) RawJSON() string { return r.JSON.raw }
func (r *MergeMergeObjectValuesResponsePrimaryFieldValueAddress) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MergeMergeObjectValuesResponsePrimaryFieldValueFullName struct {
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
func (r MergeMergeObjectValuesResponsePrimaryFieldValueFullName) RawJSON() string { return r.JSON.raw }
func (r *MergeMergeObjectValuesResponsePrimaryFieldValueFullName) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MergeMergeObjectValuesResponsePrimaryRelationship struct {
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
func (r MergeMergeObjectValuesResponsePrimaryRelationship) RawJSON() string { return r.JSON.raw }
func (r *MergeMergeObjectValuesResponsePrimaryRelationship) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MergeMergeObjectValuesResponseSummary struct {
	// Number of attribute fields written onto the primary record.
	FieldWriteCount int64 `json:"fieldWriteCount" api:"required"`
	// Number of related records re-pointed from the duplicate to the primary.
	SyncRepointedCount int64 `json:"syncRepointedCount" api:"required"`
	// Non-fatal warnings from the merge (e.g. skipped transfers).
	Warnings []string `json:"warnings" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FieldWriteCount    respjson.Field
		SyncRepointedCount respjson.Field
		Warnings           respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MergeMergeObjectValuesResponseSummary) RawJSON() string { return r.JSON.raw }
func (r *MergeMergeObjectValuesResponseSummary) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MergeMergeOpportunitiesResponse struct {
	Merge   MergeMergeOpportunitiesResponseMerge   `json:"merge" api:"required"`
	Primary MergeMergeOpportunitiesResponsePrimary `json:"primary" api:"required"`
	Summary MergeMergeOpportunitiesResponseSummary `json:"summary" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Merge       respjson.Field
		Primary     respjson.Field
		Summary     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MergeMergeOpportunitiesResponse) RawJSON() string { return r.JSON.raw }
func (r *MergeMergeOpportunitiesResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MergeMergeOpportunitiesResponseMerge struct {
	// Unique identifier for the merge operation.
	ID string `json:"id" api:"required"`
	// Current status of the merge: `cleanup_pending`, `done`, or `failed`.
	Status string `json:"status" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MergeMergeOpportunitiesResponseMerge) RawJSON() string { return r.JSON.raw }
func (r *MergeMergeOpportunitiesResponseMerge) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MergeMergeOpportunitiesResponsePrimary struct {
	// Unique identifier for the entity.
	ID string `json:"id" api:"required"`
	// ISO 8601 timestamp of when the entity was created.
	CreatedAt string `json:"createdAt" api:"required"`
	// Map of field names to their typed values. System fields are prefixed with `$`
	// (e.g. `$name`, `$email`); custom attributes use their bare slug.
	Fields map[string]MergeMergeOpportunitiesResponsePrimaryField `json:"fields" api:"required"`
	// URL to view the entity in the Lightfield web app, or null.
	HTTPLink string `json:"httpLink" api:"required"`
	// Map of relationship names to their associated entities. System relationships are
	// prefixed with `$` (e.g. `$owner`, `$contact`).
	Relationships map[string]MergeMergeOpportunitiesResponsePrimaryRelationship `json:"relationships" api:"required"`
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
func (r MergeMergeOpportunitiesResponsePrimary) RawJSON() string { return r.JSON.raw }
func (r *MergeMergeOpportunitiesResponsePrimary) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MergeMergeOpportunitiesResponsePrimaryField struct {
	// The field value, or null if unset.
	Value MergeMergeOpportunitiesResponsePrimaryFieldValueUnion `json:"value" api:"required"`
	// The data type of the field.
	//
	// Any of "ADDRESS", "CHECKBOX", "CURRENCY", "DATETIME", "EMAIL", "FULL_NAME",
	// "MARKDOWN", "MULTI_SELECT", "NUMBER", "SINGLE_SELECT", "SOCIAL_HANDLE",
	// "TELEPHONE", "TEXT", "URL", "HTML".
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
func (r MergeMergeOpportunitiesResponsePrimaryField) RawJSON() string { return r.JSON.raw }
func (r *MergeMergeOpportunitiesResponsePrimaryField) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// MergeMergeOpportunitiesResponsePrimaryFieldValueUnion contains all possible
// properties and values from [string], [float64], [bool], [[]string],
// [MergeMergeOpportunitiesResponsePrimaryFieldValueAddress],
// [MergeMergeOpportunitiesResponsePrimaryFieldValueFullName].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfStringArray]
type MergeMergeOpportunitiesResponsePrimaryFieldValueUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]string] instead of an object.
	OfStringArray []string `json:",inline"`
	// This field is from variant
	// [MergeMergeOpportunitiesResponsePrimaryFieldValueAddress].
	City string `json:"city"`
	// This field is from variant
	// [MergeMergeOpportunitiesResponsePrimaryFieldValueAddress].
	Country string `json:"country"`
	// This field is from variant
	// [MergeMergeOpportunitiesResponsePrimaryFieldValueAddress].
	Latitude float64 `json:"latitude"`
	// This field is from variant
	// [MergeMergeOpportunitiesResponsePrimaryFieldValueAddress].
	Longitude float64 `json:"longitude"`
	// This field is from variant
	// [MergeMergeOpportunitiesResponsePrimaryFieldValueAddress].
	PostalCode string `json:"postalCode"`
	// This field is from variant
	// [MergeMergeOpportunitiesResponsePrimaryFieldValueAddress].
	State string `json:"state"`
	// This field is from variant
	// [MergeMergeOpportunitiesResponsePrimaryFieldValueAddress].
	Street string `json:"street"`
	// This field is from variant
	// [MergeMergeOpportunitiesResponsePrimaryFieldValueAddress].
	Street2 string `json:"street2"`
	// This field is from variant
	// [MergeMergeOpportunitiesResponsePrimaryFieldValueFullName].
	FirstName string `json:"firstName"`
	// This field is from variant
	// [MergeMergeOpportunitiesResponsePrimaryFieldValueFullName].
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

func (u MergeMergeOpportunitiesResponsePrimaryFieldValueUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MergeMergeOpportunitiesResponsePrimaryFieldValueUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MergeMergeOpportunitiesResponsePrimaryFieldValueUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MergeMergeOpportunitiesResponsePrimaryFieldValueUnion) AsStringArray() (v []string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MergeMergeOpportunitiesResponsePrimaryFieldValueUnion) AsAddress() (v MergeMergeOpportunitiesResponsePrimaryFieldValueAddress) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MergeMergeOpportunitiesResponsePrimaryFieldValueUnion) AsFullName() (v MergeMergeOpportunitiesResponsePrimaryFieldValueFullName) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u MergeMergeOpportunitiesResponsePrimaryFieldValueUnion) RawJSON() string { return u.JSON.raw }

func (r *MergeMergeOpportunitiesResponsePrimaryFieldValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MergeMergeOpportunitiesResponsePrimaryFieldValueAddress struct {
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
func (r MergeMergeOpportunitiesResponsePrimaryFieldValueAddress) RawJSON() string { return r.JSON.raw }
func (r *MergeMergeOpportunitiesResponsePrimaryFieldValueAddress) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MergeMergeOpportunitiesResponsePrimaryFieldValueFullName struct {
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
func (r MergeMergeOpportunitiesResponsePrimaryFieldValueFullName) RawJSON() string { return r.JSON.raw }
func (r *MergeMergeOpportunitiesResponsePrimaryFieldValueFullName) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MergeMergeOpportunitiesResponsePrimaryRelationship struct {
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
func (r MergeMergeOpportunitiesResponsePrimaryRelationship) RawJSON() string { return r.JSON.raw }
func (r *MergeMergeOpportunitiesResponsePrimaryRelationship) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MergeMergeOpportunitiesResponseSummary struct {
	// Number of attribute fields written onto the primary record.
	FieldWriteCount int64 `json:"fieldWriteCount" api:"required"`
	// Number of related records re-pointed from the duplicate to the primary.
	SyncRepointedCount int64 `json:"syncRepointedCount" api:"required"`
	// Non-fatal warnings from the merge (e.g. skipped transfers).
	Warnings []string `json:"warnings" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FieldWriteCount    respjson.Field
		SyncRepointedCount respjson.Field
		Warnings           respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MergeMergeOpportunitiesResponseSummary) RawJSON() string { return r.JSON.raw }
func (r *MergeMergeOpportunitiesResponseSummary) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MergeMergeAccountsParams struct {
	// ID of the duplicate record to merge into the primary and then discard.
	DuplicateID string `json:"duplicateId" api:"required"`
	// ID of the record to keep.
	PrimaryID string `json:"primaryId" api:"required"`
	// Per-field resolution overrides keyed by attribute slug.
	FieldResolutions map[string]MergeMergeAccountsParamsFieldResolutionUnion `json:"fieldResolutions,omitzero"`
	Options          MergeMergeAccountsParamsOptions                         `json:"options,omitzero"`
	paramObj
}

func (r MergeMergeAccountsParams) MarshalJSON() (data []byte, err error) {
	type shadow MergeMergeAccountsParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MergeMergeAccountsParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type MergeMergeAccountsParamsFieldResolutionUnion struct {
	// Check if union is this variant with
	// !param.IsOmitted(union.OfMergeMergeAccountssFieldResolutionString)
	OfMergeMergeAccountssFieldResolutionString param.Opt[string]                             `json:",omitzero,inline"`
	OfMergeMergeAccountssFieldResolutionValue  *MergeMergeAccountsParamsFieldResolutionValue `json:",omitzero,inline"`
	paramUnion
}

func (u MergeMergeAccountsParamsFieldResolutionUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfMergeMergeAccountssFieldResolutionString, u.OfMergeMergeAccountssFieldResolutionValue)
}
func (u *MergeMergeAccountsParamsFieldResolutionUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

type MergeMergeAccountsParamsFieldResolutionString string

const (
	MergeMergeAccountsParamsFieldResolutionStringPrimary   MergeMergeAccountsParamsFieldResolutionString = "primary"
	MergeMergeAccountsParamsFieldResolutionStringDuplicate MergeMergeAccountsParamsFieldResolutionString = "duplicate"
)

// The property Value is required.
type MergeMergeAccountsParamsFieldResolutionValue struct {
	Value MergeMergeAccountsParamsFieldResolutionValueValueUnion `json:"value,omitzero" api:"required"`
	paramObj
}

func (r MergeMergeAccountsParamsFieldResolutionValue) MarshalJSON() (data []byte, err error) {
	type shadow MergeMergeAccountsParamsFieldResolutionValue
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MergeMergeAccountsParamsFieldResolutionValue) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type MergeMergeAccountsParamsFieldResolutionValueValueUnion struct {
	OfString      param.Opt[string]                                          `json:",omitzero,inline"`
	OfFloat       param.Opt[float64]                                         `json:",omitzero,inline"`
	OfBool        param.Opt[bool]                                            `json:",omitzero,inline"`
	OfStringArray []string                                                   `json:",omitzero,inline"`
	OfAddress     *MergeMergeAccountsParamsFieldResolutionValueValueAddress  `json:",omitzero,inline"`
	OfFullName    *MergeMergeAccountsParamsFieldResolutionValueValueFullName `json:",omitzero,inline"`
	paramUnion
}

func (u MergeMergeAccountsParamsFieldResolutionValueValueUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString,
		u.OfFloat,
		u.OfBool,
		u.OfStringArray,
		u.OfAddress,
		u.OfFullName)
}
func (u *MergeMergeAccountsParamsFieldResolutionValueValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

type MergeMergeAccountsParamsFieldResolutionValueValueAddress struct {
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

func (r MergeMergeAccountsParamsFieldResolutionValueValueAddress) MarshalJSON() (data []byte, err error) {
	type shadow MergeMergeAccountsParamsFieldResolutionValueValueAddress
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MergeMergeAccountsParamsFieldResolutionValueValueAddress) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MergeMergeAccountsParamsFieldResolutionValueValueFullName struct {
	// The contact's first name.
	FirstName param.Opt[string] `json:"firstName,omitzero"`
	// The contact's last name.
	LastName param.Opt[string] `json:"lastName,omitzero"`
	paramObj
}

func (r MergeMergeAccountsParamsFieldResolutionValueValueFullName) MarshalJSON() (data []byte, err error) {
	type shadow MergeMergeAccountsParamsFieldResolutionValueValueFullName
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MergeMergeAccountsParamsFieldResolutionValueValueFullName) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MergeMergeAccountsParamsOptions struct {
	// When true, multi-select fields are merged by union rather than
	// primary-takes-all.
	MultiSelectUnion param.Opt[bool] `json:"multiSelectUnion,omitzero"`
	paramObj
}

func (r MergeMergeAccountsParamsOptions) MarshalJSON() (data []byte, err error) {
	type shadow MergeMergeAccountsParamsOptions
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MergeMergeAccountsParamsOptions) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MergeMergeContactsParams struct {
	// ID of the duplicate record to merge into the primary and then discard.
	DuplicateID string `json:"duplicateId" api:"required"`
	// ID of the record to keep.
	PrimaryID string `json:"primaryId" api:"required"`
	// Per-field resolution overrides keyed by attribute slug.
	FieldResolutions map[string]MergeMergeContactsParamsFieldResolutionUnion `json:"fieldResolutions,omitzero"`
	Options          MergeMergeContactsParamsOptions                         `json:"options,omitzero"`
	paramObj
}

func (r MergeMergeContactsParams) MarshalJSON() (data []byte, err error) {
	type shadow MergeMergeContactsParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MergeMergeContactsParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type MergeMergeContactsParamsFieldResolutionUnion struct {
	// Check if union is this variant with
	// !param.IsOmitted(union.OfMergeMergeContactssFieldResolutionString)
	OfMergeMergeContactssFieldResolutionString param.Opt[string]                             `json:",omitzero,inline"`
	OfMergeMergeContactssFieldResolutionValue  *MergeMergeContactsParamsFieldResolutionValue `json:",omitzero,inline"`
	paramUnion
}

func (u MergeMergeContactsParamsFieldResolutionUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfMergeMergeContactssFieldResolutionString, u.OfMergeMergeContactssFieldResolutionValue)
}
func (u *MergeMergeContactsParamsFieldResolutionUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

type MergeMergeContactsParamsFieldResolutionString string

const (
	MergeMergeContactsParamsFieldResolutionStringPrimary   MergeMergeContactsParamsFieldResolutionString = "primary"
	MergeMergeContactsParamsFieldResolutionStringDuplicate MergeMergeContactsParamsFieldResolutionString = "duplicate"
)

// The property Value is required.
type MergeMergeContactsParamsFieldResolutionValue struct {
	Value MergeMergeContactsParamsFieldResolutionValueValueUnion `json:"value,omitzero" api:"required"`
	paramObj
}

func (r MergeMergeContactsParamsFieldResolutionValue) MarshalJSON() (data []byte, err error) {
	type shadow MergeMergeContactsParamsFieldResolutionValue
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MergeMergeContactsParamsFieldResolutionValue) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type MergeMergeContactsParamsFieldResolutionValueValueUnion struct {
	OfString      param.Opt[string]                                          `json:",omitzero,inline"`
	OfFloat       param.Opt[float64]                                         `json:",omitzero,inline"`
	OfBool        param.Opt[bool]                                            `json:",omitzero,inline"`
	OfStringArray []string                                                   `json:",omitzero,inline"`
	OfAddress     *MergeMergeContactsParamsFieldResolutionValueValueAddress  `json:",omitzero,inline"`
	OfFullName    *MergeMergeContactsParamsFieldResolutionValueValueFullName `json:",omitzero,inline"`
	paramUnion
}

func (u MergeMergeContactsParamsFieldResolutionValueValueUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString,
		u.OfFloat,
		u.OfBool,
		u.OfStringArray,
		u.OfAddress,
		u.OfFullName)
}
func (u *MergeMergeContactsParamsFieldResolutionValueValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

type MergeMergeContactsParamsFieldResolutionValueValueAddress struct {
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

func (r MergeMergeContactsParamsFieldResolutionValueValueAddress) MarshalJSON() (data []byte, err error) {
	type shadow MergeMergeContactsParamsFieldResolutionValueValueAddress
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MergeMergeContactsParamsFieldResolutionValueValueAddress) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MergeMergeContactsParamsFieldResolutionValueValueFullName struct {
	// The contact's first name.
	FirstName param.Opt[string] `json:"firstName,omitzero"`
	// The contact's last name.
	LastName param.Opt[string] `json:"lastName,omitzero"`
	paramObj
}

func (r MergeMergeContactsParamsFieldResolutionValueValueFullName) MarshalJSON() (data []byte, err error) {
	type shadow MergeMergeContactsParamsFieldResolutionValueValueFullName
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MergeMergeContactsParamsFieldResolutionValueValueFullName) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MergeMergeContactsParamsOptions struct {
	// When true, multi-select fields are merged by union rather than
	// primary-takes-all.
	MultiSelectUnion param.Opt[bool] `json:"multiSelectUnion,omitzero"`
	paramObj
}

func (r MergeMergeContactsParamsOptions) MarshalJSON() (data []byte, err error) {
	type shadow MergeMergeContactsParamsOptions
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MergeMergeContactsParamsOptions) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MergeMergeObjectValuesParams struct {
	// ID of the duplicate record to merge into the primary and then discard.
	DuplicateID string `json:"duplicateId" api:"required"`
	// ID of the record to keep.
	PrimaryID string `json:"primaryId" api:"required"`
	// Per-field resolution overrides keyed by attribute slug.
	FieldResolutions map[string]MergeMergeObjectValuesParamsFieldResolutionUnion `json:"fieldResolutions,omitzero"`
	Options          MergeMergeObjectValuesParamsOptions                         `json:"options,omitzero"`
	paramObj
}

func (r MergeMergeObjectValuesParams) MarshalJSON() (data []byte, err error) {
	type shadow MergeMergeObjectValuesParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MergeMergeObjectValuesParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type MergeMergeObjectValuesParamsFieldResolutionUnion struct {
	// Check if union is this variant with
	// !param.IsOmitted(union.OfMergeMergeObjectValuessFieldResolutionString)
	OfMergeMergeObjectValuessFieldResolutionString param.Opt[string]                                 `json:",omitzero,inline"`
	OfMergeMergeObjectValuessFieldResolutionValue  *MergeMergeObjectValuesParamsFieldResolutionValue `json:",omitzero,inline"`
	paramUnion
}

func (u MergeMergeObjectValuesParamsFieldResolutionUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfMergeMergeObjectValuessFieldResolutionString, u.OfMergeMergeObjectValuessFieldResolutionValue)
}
func (u *MergeMergeObjectValuesParamsFieldResolutionUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

type MergeMergeObjectValuesParamsFieldResolutionString string

const (
	MergeMergeObjectValuesParamsFieldResolutionStringPrimary   MergeMergeObjectValuesParamsFieldResolutionString = "primary"
	MergeMergeObjectValuesParamsFieldResolutionStringDuplicate MergeMergeObjectValuesParamsFieldResolutionString = "duplicate"
)

// The property Value is required.
type MergeMergeObjectValuesParamsFieldResolutionValue struct {
	Value MergeMergeObjectValuesParamsFieldResolutionValueValueUnion `json:"value,omitzero" api:"required"`
	paramObj
}

func (r MergeMergeObjectValuesParamsFieldResolutionValue) MarshalJSON() (data []byte, err error) {
	type shadow MergeMergeObjectValuesParamsFieldResolutionValue
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MergeMergeObjectValuesParamsFieldResolutionValue) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type MergeMergeObjectValuesParamsFieldResolutionValueValueUnion struct {
	OfString      param.Opt[string]                                              `json:",omitzero,inline"`
	OfFloat       param.Opt[float64]                                             `json:",omitzero,inline"`
	OfBool        param.Opt[bool]                                                `json:",omitzero,inline"`
	OfStringArray []string                                                       `json:",omitzero,inline"`
	OfAddress     *MergeMergeObjectValuesParamsFieldResolutionValueValueAddress  `json:",omitzero,inline"`
	OfFullName    *MergeMergeObjectValuesParamsFieldResolutionValueValueFullName `json:",omitzero,inline"`
	paramUnion
}

func (u MergeMergeObjectValuesParamsFieldResolutionValueValueUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString,
		u.OfFloat,
		u.OfBool,
		u.OfStringArray,
		u.OfAddress,
		u.OfFullName)
}
func (u *MergeMergeObjectValuesParamsFieldResolutionValueValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

type MergeMergeObjectValuesParamsFieldResolutionValueValueAddress struct {
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

func (r MergeMergeObjectValuesParamsFieldResolutionValueValueAddress) MarshalJSON() (data []byte, err error) {
	type shadow MergeMergeObjectValuesParamsFieldResolutionValueValueAddress
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MergeMergeObjectValuesParamsFieldResolutionValueValueAddress) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MergeMergeObjectValuesParamsFieldResolutionValueValueFullName struct {
	// The contact's first name.
	FirstName param.Opt[string] `json:"firstName,omitzero"`
	// The contact's last name.
	LastName param.Opt[string] `json:"lastName,omitzero"`
	paramObj
}

func (r MergeMergeObjectValuesParamsFieldResolutionValueValueFullName) MarshalJSON() (data []byte, err error) {
	type shadow MergeMergeObjectValuesParamsFieldResolutionValueValueFullName
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MergeMergeObjectValuesParamsFieldResolutionValueValueFullName) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MergeMergeObjectValuesParamsOptions struct {
	// When true, multi-select fields are merged by union rather than
	// primary-takes-all.
	MultiSelectUnion param.Opt[bool] `json:"multiSelectUnion,omitzero"`
	paramObj
}

func (r MergeMergeObjectValuesParamsOptions) MarshalJSON() (data []byte, err error) {
	type shadow MergeMergeObjectValuesParamsOptions
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MergeMergeObjectValuesParamsOptions) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MergeMergeOpportunitiesParams struct {
	// ID of the duplicate record to merge into the primary and then discard.
	DuplicateID string `json:"duplicateId" api:"required"`
	// ID of the record to keep.
	PrimaryID string `json:"primaryId" api:"required"`
	// Per-field resolution overrides keyed by attribute slug.
	FieldResolutions map[string]MergeMergeOpportunitiesParamsFieldResolutionUnion `json:"fieldResolutions,omitzero"`
	Options          MergeMergeOpportunitiesParamsOptions                         `json:"options,omitzero"`
	paramObj
}

func (r MergeMergeOpportunitiesParams) MarshalJSON() (data []byte, err error) {
	type shadow MergeMergeOpportunitiesParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MergeMergeOpportunitiesParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type MergeMergeOpportunitiesParamsFieldResolutionUnion struct {
	// Check if union is this variant with
	// !param.IsOmitted(union.OfMergeMergeOpportunitiessFieldResolutionString)
	OfMergeMergeOpportunitiessFieldResolutionString param.Opt[string]                                  `json:",omitzero,inline"`
	OfMergeMergeOpportunitiessFieldResolutionValue  *MergeMergeOpportunitiesParamsFieldResolutionValue `json:",omitzero,inline"`
	paramUnion
}

func (u MergeMergeOpportunitiesParamsFieldResolutionUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfMergeMergeOpportunitiessFieldResolutionString, u.OfMergeMergeOpportunitiessFieldResolutionValue)
}
func (u *MergeMergeOpportunitiesParamsFieldResolutionUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

type MergeMergeOpportunitiesParamsFieldResolutionString string

const (
	MergeMergeOpportunitiesParamsFieldResolutionStringPrimary   MergeMergeOpportunitiesParamsFieldResolutionString = "primary"
	MergeMergeOpportunitiesParamsFieldResolutionStringDuplicate MergeMergeOpportunitiesParamsFieldResolutionString = "duplicate"
)

// The property Value is required.
type MergeMergeOpportunitiesParamsFieldResolutionValue struct {
	Value MergeMergeOpportunitiesParamsFieldResolutionValueValueUnion `json:"value,omitzero" api:"required"`
	paramObj
}

func (r MergeMergeOpportunitiesParamsFieldResolutionValue) MarshalJSON() (data []byte, err error) {
	type shadow MergeMergeOpportunitiesParamsFieldResolutionValue
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MergeMergeOpportunitiesParamsFieldResolutionValue) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type MergeMergeOpportunitiesParamsFieldResolutionValueValueUnion struct {
	OfString      param.Opt[string]                                               `json:",omitzero,inline"`
	OfFloat       param.Opt[float64]                                              `json:",omitzero,inline"`
	OfBool        param.Opt[bool]                                                 `json:",omitzero,inline"`
	OfStringArray []string                                                        `json:",omitzero,inline"`
	OfAddress     *MergeMergeOpportunitiesParamsFieldResolutionValueValueAddress  `json:",omitzero,inline"`
	OfFullName    *MergeMergeOpportunitiesParamsFieldResolutionValueValueFullName `json:",omitzero,inline"`
	paramUnion
}

func (u MergeMergeOpportunitiesParamsFieldResolutionValueValueUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString,
		u.OfFloat,
		u.OfBool,
		u.OfStringArray,
		u.OfAddress,
		u.OfFullName)
}
func (u *MergeMergeOpportunitiesParamsFieldResolutionValueValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

type MergeMergeOpportunitiesParamsFieldResolutionValueValueAddress struct {
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

func (r MergeMergeOpportunitiesParamsFieldResolutionValueValueAddress) MarshalJSON() (data []byte, err error) {
	type shadow MergeMergeOpportunitiesParamsFieldResolutionValueValueAddress
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MergeMergeOpportunitiesParamsFieldResolutionValueValueAddress) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MergeMergeOpportunitiesParamsFieldResolutionValueValueFullName struct {
	// The contact's first name.
	FirstName param.Opt[string] `json:"firstName,omitzero"`
	// The contact's last name.
	LastName param.Opt[string] `json:"lastName,omitzero"`
	paramObj
}

func (r MergeMergeOpportunitiesParamsFieldResolutionValueValueFullName) MarshalJSON() (data []byte, err error) {
	type shadow MergeMergeOpportunitiesParamsFieldResolutionValueValueFullName
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MergeMergeOpportunitiesParamsFieldResolutionValueValueFullName) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MergeMergeOpportunitiesParamsOptions struct {
	// When true, multi-select fields are merged by union rather than
	// primary-takes-all.
	MultiSelectUnion param.Opt[bool] `json:"multiSelectUnion,omitzero"`
	paramObj
}

func (r MergeMergeOpportunitiesParamsOptions) MarshalJSON() (data []byte, err error) {
	type shadow MergeMergeOpportunitiesParamsOptions
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MergeMergeOpportunitiesParamsOptions) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
