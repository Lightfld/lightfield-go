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
// To avoid duplicates, we recommend a find-or-create pattern — use
// <u>[list filtering](/using-the-api/list-endpoints/#filtering)</u> to check if a
// record exists before creating.
//
// **[Required scope](/using-the-api/scopes/):** `opportunities:create`
//
// **[Rate limit category](/using-the-api/rate-limits/):** Write
func (r *OpportunityService) New(ctx context.Context, body OpportunityNewParams, opts ...option.RequestOption) (res *OpportunityCreateResponse, err error) {
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
func (r *OpportunityService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *OpportunityRetrieveResponse, err error) {
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
// through results, and `$field` query parameters to filter. See
// <u>[List endpoints](/using-the-api/list-endpoints/)</u> for more information
// about <u>[pagination](/using-the-api/list-endpoints/#pagination)</u> and
// <u>[filtering](/using-the-api/list-endpoints/#filtering)</u>.
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

type OpportunityCreateResponse struct {
	// Unique identifier for the entity.
	ID string `json:"id" api:"required"`
	// ISO 8601 timestamp of when the entity was created.
	CreatedAt string `json:"createdAt" api:"required"`
	// Map of field names to their typed values. System fields are prefixed with `$`
	// (e.g. `$name`, `$email`); custom attributes use their bare slug.
	Fields map[string]OpportunityCreateResponseField `json:"fields" api:"required"`
	// URL to view the entity in the Lightfield web app, or null.
	HTTPLink string `json:"httpLink" api:"required"`
	// Map of relationship names to their associated entities. System relationships are
	// prefixed with `$` (e.g. `$owner`, `$contact`).
	Relationships map[string]OpportunityCreateResponseRelationship `json:"relationships" api:"required"`
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
func (r OpportunityCreateResponse) RawJSON() string { return r.JSON.raw }
func (r *OpportunityCreateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OpportunityCreateResponseField struct {
	// The field value, or null if unset.
	Value OpportunityCreateResponseFieldValueUnion `json:"value" api:"required"`
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
func (r OpportunityCreateResponseField) RawJSON() string { return r.JSON.raw }
func (r *OpportunityCreateResponseField) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// OpportunityCreateResponseFieldValueUnion contains all possible properties and
// values from [string], [float64], [bool], [[]string],
// [OpportunityCreateResponseFieldValueAddress],
// [OpportunityCreateResponseFieldValueFullName].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfStringArray]
type OpportunityCreateResponseFieldValueUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]string] instead of an object.
	OfStringArray []string `json:",inline"`
	// This field is from variant [OpportunityCreateResponseFieldValueAddress].
	City string `json:"city"`
	// This field is from variant [OpportunityCreateResponseFieldValueAddress].
	Country string `json:"country"`
	// This field is from variant [OpportunityCreateResponseFieldValueAddress].
	Latitude float64 `json:"latitude"`
	// This field is from variant [OpportunityCreateResponseFieldValueAddress].
	Longitude float64 `json:"longitude"`
	// This field is from variant [OpportunityCreateResponseFieldValueAddress].
	PostalCode string `json:"postalCode"`
	// This field is from variant [OpportunityCreateResponseFieldValueAddress].
	State string `json:"state"`
	// This field is from variant [OpportunityCreateResponseFieldValueAddress].
	Street string `json:"street"`
	// This field is from variant [OpportunityCreateResponseFieldValueAddress].
	Street2 string `json:"street2"`
	// This field is from variant [OpportunityCreateResponseFieldValueFullName].
	FirstName string `json:"firstName"`
	// This field is from variant [OpportunityCreateResponseFieldValueFullName].
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

func (u OpportunityCreateResponseFieldValueUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u OpportunityCreateResponseFieldValueUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u OpportunityCreateResponseFieldValueUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u OpportunityCreateResponseFieldValueUnion) AsStringArray() (v []string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u OpportunityCreateResponseFieldValueUnion) AsAddress() (v OpportunityCreateResponseFieldValueAddress) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u OpportunityCreateResponseFieldValueUnion) AsFullName() (v OpportunityCreateResponseFieldValueFullName) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u OpportunityCreateResponseFieldValueUnion) RawJSON() string { return u.JSON.raw }

func (r *OpportunityCreateResponseFieldValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OpportunityCreateResponseFieldValueAddress struct {
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
func (r OpportunityCreateResponseFieldValueAddress) RawJSON() string { return r.JSON.raw }
func (r *OpportunityCreateResponseFieldValueAddress) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OpportunityCreateResponseFieldValueFullName struct {
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
func (r OpportunityCreateResponseFieldValueFullName) RawJSON() string { return r.JSON.raw }
func (r *OpportunityCreateResponseFieldValueFullName) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OpportunityCreateResponseRelationship struct {
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
func (r OpportunityCreateResponseRelationship) RawJSON() string { return r.JSON.raw }
func (r *OpportunityCreateResponseRelationship) UnmarshalJSON(data []byte) error {
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
	TypeConfiguration OpportunityDefinitionsResponseFieldDefinitionTypeConfiguration `json:"typeConfiguration" api:"required"`
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

// Type-specific configuration (e.g. select options, currency code).
type OpportunityDefinitionsResponseFieldDefinitionTypeConfiguration struct {
	// ISO 4217 3-letter currency code.
	Currency string `json:"currency"`
	// Social platform associated with this handle field.
	//
	// Any of "TWITTER", "LINKEDIN", "FACEBOOK", "INSTAGRAM".
	HandleService string `json:"handleService"`
	// Whether this field accepts multiple values.
	MultipleValues bool `json:"multipleValues"`
	// Available options for select fields.
	Options []OpportunityDefinitionsResponseFieldDefinitionTypeConfigurationOption `json:"options"`
	// Whether values for this field must be unique.
	Unique bool `json:"unique"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Currency       respjson.Field
		HandleService  respjson.Field
		MultipleValues respjson.Field
		Options        respjson.Field
		Unique         respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OpportunityDefinitionsResponseFieldDefinitionTypeConfiguration) RawJSON() string {
	return r.JSON.raw
}
func (r *OpportunityDefinitionsResponseFieldDefinitionTypeConfiguration) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OpportunityDefinitionsResponseFieldDefinitionTypeConfigurationOption struct {
	// Unique identifier of the select option.
	ID string `json:"id" api:"required"`
	// Human-readable display name of the option.
	Label string `json:"label" api:"required"`
	// Description of the option, or null.
	Description string `json:"description" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Label       respjson.Field
		Description respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OpportunityDefinitionsResponseFieldDefinitionTypeConfigurationOption) RawJSON() string {
	return r.JSON.raw
}
func (r *OpportunityDefinitionsResponseFieldDefinitionTypeConfigurationOption) UnmarshalJSON(data []byte) error {
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
// values from [string], [float64], [bool], [[]string],
// [OpportunityListResponseDataFieldValueAddress],
// [OpportunityListResponseDataFieldValueFullName].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfStringArray]
type OpportunityListResponseDataFieldValueUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]string] instead of an object.
	OfStringArray []string `json:",inline"`
	// This field is from variant [OpportunityListResponseDataFieldValueAddress].
	City string `json:"city"`
	// This field is from variant [OpportunityListResponseDataFieldValueAddress].
	Country string `json:"country"`
	// This field is from variant [OpportunityListResponseDataFieldValueAddress].
	Latitude float64 `json:"latitude"`
	// This field is from variant [OpportunityListResponseDataFieldValueAddress].
	Longitude float64 `json:"longitude"`
	// This field is from variant [OpportunityListResponseDataFieldValueAddress].
	PostalCode string `json:"postalCode"`
	// This field is from variant [OpportunityListResponseDataFieldValueAddress].
	State string `json:"state"`
	// This field is from variant [OpportunityListResponseDataFieldValueAddress].
	Street string `json:"street"`
	// This field is from variant [OpportunityListResponseDataFieldValueAddress].
	Street2 string `json:"street2"`
	// This field is from variant [OpportunityListResponseDataFieldValueFullName].
	FirstName string `json:"firstName"`
	// This field is from variant [OpportunityListResponseDataFieldValueFullName].
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

func (u OpportunityListResponseDataFieldValueUnion) AsStringArray() (v []string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u OpportunityListResponseDataFieldValueUnion) AsAddress() (v OpportunityListResponseDataFieldValueAddress) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u OpportunityListResponseDataFieldValueUnion) AsFullName() (v OpportunityListResponseDataFieldValueFullName) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u OpportunityListResponseDataFieldValueUnion) RawJSON() string { return u.JSON.raw }

func (r *OpportunityListResponseDataFieldValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OpportunityListResponseDataFieldValueAddress struct {
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
func (r OpportunityListResponseDataFieldValueAddress) RawJSON() string { return r.JSON.raw }
func (r *OpportunityListResponseDataFieldValueAddress) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OpportunityListResponseDataFieldValueFullName struct {
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
func (r OpportunityListResponseDataFieldValueFullName) RawJSON() string { return r.JSON.raw }
func (r *OpportunityListResponseDataFieldValueFullName) UnmarshalJSON(data []byte) error {
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

type OpportunityRetrieveResponse struct {
	// Unique identifier for the entity.
	ID string `json:"id" api:"required"`
	// ISO 8601 timestamp of when the entity was created.
	CreatedAt string `json:"createdAt" api:"required"`
	// Map of field names to their typed values. System fields are prefixed with `$`
	// (e.g. `$name`, `$email`); custom attributes use their bare slug.
	Fields map[string]OpportunityRetrieveResponseField `json:"fields" api:"required"`
	// URL to view the entity in the Lightfield web app, or null.
	HTTPLink string `json:"httpLink" api:"required"`
	// Map of relationship names to their associated entities. System relationships are
	// prefixed with `$` (e.g. `$owner`, `$contact`).
	Relationships map[string]OpportunityRetrieveResponseRelationship `json:"relationships" api:"required"`
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
func (r OpportunityRetrieveResponse) RawJSON() string { return r.JSON.raw }
func (r *OpportunityRetrieveResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OpportunityRetrieveResponseField struct {
	// The field value, or null if unset.
	Value OpportunityRetrieveResponseFieldValueUnion `json:"value" api:"required"`
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
func (r OpportunityRetrieveResponseField) RawJSON() string { return r.JSON.raw }
func (r *OpportunityRetrieveResponseField) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// OpportunityRetrieveResponseFieldValueUnion contains all possible properties and
// values from [string], [float64], [bool], [[]string],
// [OpportunityRetrieveResponseFieldValueAddress],
// [OpportunityRetrieveResponseFieldValueFullName].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfStringArray]
type OpportunityRetrieveResponseFieldValueUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]string] instead of an object.
	OfStringArray []string `json:",inline"`
	// This field is from variant [OpportunityRetrieveResponseFieldValueAddress].
	City string `json:"city"`
	// This field is from variant [OpportunityRetrieveResponseFieldValueAddress].
	Country string `json:"country"`
	// This field is from variant [OpportunityRetrieveResponseFieldValueAddress].
	Latitude float64 `json:"latitude"`
	// This field is from variant [OpportunityRetrieveResponseFieldValueAddress].
	Longitude float64 `json:"longitude"`
	// This field is from variant [OpportunityRetrieveResponseFieldValueAddress].
	PostalCode string `json:"postalCode"`
	// This field is from variant [OpportunityRetrieveResponseFieldValueAddress].
	State string `json:"state"`
	// This field is from variant [OpportunityRetrieveResponseFieldValueAddress].
	Street string `json:"street"`
	// This field is from variant [OpportunityRetrieveResponseFieldValueAddress].
	Street2 string `json:"street2"`
	// This field is from variant [OpportunityRetrieveResponseFieldValueFullName].
	FirstName string `json:"firstName"`
	// This field is from variant [OpportunityRetrieveResponseFieldValueFullName].
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

func (u OpportunityRetrieveResponseFieldValueUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u OpportunityRetrieveResponseFieldValueUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u OpportunityRetrieveResponseFieldValueUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u OpportunityRetrieveResponseFieldValueUnion) AsStringArray() (v []string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u OpportunityRetrieveResponseFieldValueUnion) AsAddress() (v OpportunityRetrieveResponseFieldValueAddress) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u OpportunityRetrieveResponseFieldValueUnion) AsFullName() (v OpportunityRetrieveResponseFieldValueFullName) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u OpportunityRetrieveResponseFieldValueUnion) RawJSON() string { return u.JSON.raw }

func (r *OpportunityRetrieveResponseFieldValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OpportunityRetrieveResponseFieldValueAddress struct {
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
func (r OpportunityRetrieveResponseFieldValueAddress) RawJSON() string { return r.JSON.raw }
func (r *OpportunityRetrieveResponseFieldValueAddress) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OpportunityRetrieveResponseFieldValueFullName struct {
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
func (r OpportunityRetrieveResponseFieldValueFullName) RawJSON() string { return r.JSON.raw }
func (r *OpportunityRetrieveResponseFieldValueFullName) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OpportunityRetrieveResponseRelationship struct {
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
func (r OpportunityRetrieveResponseRelationship) RawJSON() string { return r.JSON.raw }
func (r *OpportunityRetrieveResponseRelationship) UnmarshalJSON(data []byte) error {
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
// values from [string], [float64], [bool], [[]string],
// [OpportunityUpdateResponseFieldValueAddress],
// [OpportunityUpdateResponseFieldValueFullName].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfStringArray]
type OpportunityUpdateResponseFieldValueUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]string] instead of an object.
	OfStringArray []string `json:",inline"`
	// This field is from variant [OpportunityUpdateResponseFieldValueAddress].
	City string `json:"city"`
	// This field is from variant [OpportunityUpdateResponseFieldValueAddress].
	Country string `json:"country"`
	// This field is from variant [OpportunityUpdateResponseFieldValueAddress].
	Latitude float64 `json:"latitude"`
	// This field is from variant [OpportunityUpdateResponseFieldValueAddress].
	Longitude float64 `json:"longitude"`
	// This field is from variant [OpportunityUpdateResponseFieldValueAddress].
	PostalCode string `json:"postalCode"`
	// This field is from variant [OpportunityUpdateResponseFieldValueAddress].
	State string `json:"state"`
	// This field is from variant [OpportunityUpdateResponseFieldValueAddress].
	Street string `json:"street"`
	// This field is from variant [OpportunityUpdateResponseFieldValueAddress].
	Street2 string `json:"street2"`
	// This field is from variant [OpportunityUpdateResponseFieldValueFullName].
	FirstName string `json:"firstName"`
	// This field is from variant [OpportunityUpdateResponseFieldValueFullName].
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

func (u OpportunityUpdateResponseFieldValueUnion) AsStringArray() (v []string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u OpportunityUpdateResponseFieldValueUnion) AsAddress() (v OpportunityUpdateResponseFieldValueAddress) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u OpportunityUpdateResponseFieldValueUnion) AsFullName() (v OpportunityUpdateResponseFieldValueFullName) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u OpportunityUpdateResponseFieldValueUnion) RawJSON() string { return u.JSON.raw }

func (r *OpportunityUpdateResponseFieldValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OpportunityUpdateResponseFieldValueAddress struct {
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
func (r OpportunityUpdateResponseFieldValueAddress) RawJSON() string { return r.JSON.raw }
func (r *OpportunityUpdateResponseFieldValueAddress) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OpportunityUpdateResponseFieldValueFullName struct {
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
func (r OpportunityUpdateResponseFieldValueFullName) RawJSON() string { return r.JSON.raw }
func (r *OpportunityUpdateResponseFieldValueFullName) UnmarshalJSON(data []byte) error {
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
	Fields map[string]OpportunityNewParamsFieldUnion `json:"fields,omitzero" api:"required"`
	// Relationships to set on the new opportunity. System relationships use a `$`
	// prefix (e.g. `$account`, `$owner`); custom relationships use their bare slug.
	// `$account` is required. Each value is a single entity ID or an array of IDs.
	// Call the
	// <u>[definitions endpoint](/api/resources/opportunity/methods/definitions)</u> to
	// list available relationship keys.
	Relationships map[string]OpportunityNewParamsRelationshipUnion `json:"relationships,omitzero" api:"required"`
	paramObj
}

func (r OpportunityNewParams) MarshalJSON() (data []byte, err error) {
	type shadow OpportunityNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *OpportunityNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type OpportunityNewParamsFieldUnion struct {
	OfString      param.Opt[string]                  `json:",omitzero,inline"`
	OfFloat       param.Opt[float64]                 `json:",omitzero,inline"`
	OfBool        param.Opt[bool]                    `json:",omitzero,inline"`
	OfStringArray []string                           `json:",omitzero,inline"`
	OfAddress     *OpportunityNewParamsFieldAddress  `json:",omitzero,inline"`
	OfFullName    *OpportunityNewParamsFieldFullName `json:",omitzero,inline"`
	paramUnion
}

func (u OpportunityNewParamsFieldUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString,
		u.OfFloat,
		u.OfBool,
		u.OfStringArray,
		u.OfAddress,
		u.OfFullName)
}
func (u *OpportunityNewParamsFieldUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

type OpportunityNewParamsFieldAddress struct {
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

func (r OpportunityNewParamsFieldAddress) MarshalJSON() (data []byte, err error) {
	type shadow OpportunityNewParamsFieldAddress
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *OpportunityNewParamsFieldAddress) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OpportunityNewParamsFieldFullName struct {
	// The contact's first name.
	FirstName param.Opt[string] `json:"firstName,omitzero"`
	// The contact's last name.
	LastName param.Opt[string] `json:"lastName,omitzero"`
	paramObj
}

func (r OpportunityNewParamsFieldFullName) MarshalJSON() (data []byte, err error) {
	type shadow OpportunityNewParamsFieldFullName
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *OpportunityNewParamsFieldFullName) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
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
	Fields map[string]OpportunityUpdateParamsFieldUnion `json:"fields,omitzero"`
	// Relationship operations to apply. System relationships use a `$` prefix (e.g.
	// `$owner`, `$champion`). Each value is an operation object with `add`, `remove`,
	// or `replace`.
	Relationships map[string]OpportunityUpdateParamsRelationship `json:"relationships,omitzero"`
	paramObj
}

func (r OpportunityUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow OpportunityUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *OpportunityUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type OpportunityUpdateParamsFieldUnion struct {
	OfString      param.Opt[string]                     `json:",omitzero,inline"`
	OfFloat       param.Opt[float64]                    `json:",omitzero,inline"`
	OfBool        param.Opt[bool]                       `json:",omitzero,inline"`
	OfStringArray []string                              `json:",omitzero,inline"`
	OfAddress     *OpportunityUpdateParamsFieldAddress  `json:",omitzero,inline"`
	OfFullName    *OpportunityUpdateParamsFieldFullName `json:",omitzero,inline"`
	paramUnion
}

func (u OpportunityUpdateParamsFieldUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString,
		u.OfFloat,
		u.OfBool,
		u.OfStringArray,
		u.OfAddress,
		u.OfFullName)
}
func (u *OpportunityUpdateParamsFieldUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

type OpportunityUpdateParamsFieldAddress struct {
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

func (r OpportunityUpdateParamsFieldAddress) MarshalJSON() (data []byte, err error) {
	type shadow OpportunityUpdateParamsFieldAddress
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *OpportunityUpdateParamsFieldAddress) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OpportunityUpdateParamsFieldFullName struct {
	// The contact's first name.
	FirstName param.Opt[string] `json:"firstName,omitzero"`
	// The contact's last name.
	LastName param.Opt[string] `json:"lastName,omitzero"`
	paramObj
}

func (r OpportunityUpdateParamsFieldFullName) MarshalJSON() (data []byte, err error) {
	type shadow OpportunityUpdateParamsFieldFullName
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *OpportunityUpdateParamsFieldFullName) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
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
