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

// Tasks represent action items in Lightfield. Each task belongs to an account, is
// assigned to a member, and can optionally be associated with an opportunity.
//
// TaskService contains methods and other services that help with interacting with
// the Lightfield API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTaskService] method instead.
type TaskService struct {
	Options []option.RequestOption
}

// NewTaskService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewTaskService(opts ...option.RequestOption) (r TaskService) {
	r = TaskService{}
	r.Options = opts
	return
}

// Creates a new task record. The `$title` and `$status` fields and the
// `$assignedTo` relationship are required.
//
// If `$createdBy` is omitted it defaults to the authenticated user. The `$note`
// relationship is read-only — manage notes via their own relationships.
//
// Supports idempotency via the `Idempotency-Key` header.
//
// **[Required scope](/using-the-api/scopes/):** `tasks:create`
//
// **[Rate limit category](/using-the-api/rate-limits/):** Write
func (r *TaskService) New(ctx context.Context, body TaskNewParams, opts ...option.RequestOption) (res *TaskCreateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/tasks"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Retrieves a single task by its ID.
//
// **[Required scope](/using-the-api/scopes/):** `tasks:read`
//
// **[Rate limit category](/using-the-api/rate-limits/):** Read
func (r *TaskService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *TaskRetrieveResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/tasks/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Updates an existing task by ID. Only included fields and relationships are
// modified.
//
// The `$note` relationship is read-only — manage notes via their own
// relationships.
//
// Supports idempotency via the `Idempotency-Key` header.
//
// **[Required scope](/using-the-api/scopes/):** `tasks:update`
//
// **[Rate limit category](/using-the-api/rate-limits/):** Write
func (r *TaskService) Update(ctx context.Context, id string, body TaskUpdateParams, opts ...option.RequestOption) (res *TaskUpdateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/tasks/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Returns a paginated list of tasks. Use `offset` and `limit` to paginate through
// results. See <u>[List endpoints](/using-the-api/list-endpoints/)</u> for more
// information about pagination.
//
// **[Required scope](/using-the-api/scopes/):** `tasks:read`
//
// **[Rate limit category](/using-the-api/rate-limits/):** Search
func (r *TaskService) List(ctx context.Context, query TaskListParams, opts ...option.RequestOption) (res *TaskListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/tasks"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Returns the schema for the field and relationship definitions available on
// tasks. Useful for understanding the shape of task data before creating or
// updating records. See
// <u>[Fields and relationships](/using-the-api/fields-and-relationships/)</u> for
// more details.
//
// **[Required scope](/using-the-api/scopes/):** `tasks:read`
//
// **[Rate limit category](/using-the-api/rate-limits/):** Read
func (r *TaskService) Definitions(ctx context.Context, opts ...option.RequestOption) (res *TaskDefinitionsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/tasks/definitions"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

type TaskCreateResponse struct {
	// Unique identifier for the entity.
	ID string `json:"id" api:"required"`
	// ISO 8601 timestamp of when the entity was created.
	CreatedAt string `json:"createdAt" api:"required"`
	// Map of field names to their typed values. System fields are prefixed with `$`
	// (e.g. `$name`, `$email`); custom attributes use their bare slug.
	Fields map[string]TaskCreateResponseField `json:"fields" api:"required"`
	// URL to view the entity in the Lightfield web app, or null.
	HTTPLink string `json:"httpLink" api:"required"`
	// Map of relationship names to their associated entities. System relationships are
	// prefixed with `$` (e.g. `$owner`, `$contact`).
	Relationships map[string]TaskCreateResponseRelationship `json:"relationships" api:"required"`
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
func (r TaskCreateResponse) RawJSON() string { return r.JSON.raw }
func (r *TaskCreateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TaskCreateResponseField struct {
	// The field value, or null if unset.
	Value TaskCreateResponseFieldValueUnion `json:"value" api:"required"`
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
func (r TaskCreateResponseField) RawJSON() string { return r.JSON.raw }
func (r *TaskCreateResponseField) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// TaskCreateResponseFieldValueUnion contains all possible properties and values
// from [string], [float64], [bool], [[]string],
// [TaskCreateResponseFieldValueAddress], [TaskCreateResponseFieldValueFullName].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfStringArray]
type TaskCreateResponseFieldValueUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]string] instead of an object.
	OfStringArray []string `json:",inline"`
	// This field is from variant [TaskCreateResponseFieldValueAddress].
	City string `json:"city"`
	// This field is from variant [TaskCreateResponseFieldValueAddress].
	Country string `json:"country"`
	// This field is from variant [TaskCreateResponseFieldValueAddress].
	Latitude float64 `json:"latitude"`
	// This field is from variant [TaskCreateResponseFieldValueAddress].
	Longitude float64 `json:"longitude"`
	// This field is from variant [TaskCreateResponseFieldValueAddress].
	PostalCode string `json:"postalCode"`
	// This field is from variant [TaskCreateResponseFieldValueAddress].
	State string `json:"state"`
	// This field is from variant [TaskCreateResponseFieldValueAddress].
	Street string `json:"street"`
	// This field is from variant [TaskCreateResponseFieldValueAddress].
	Street2 string `json:"street2"`
	// This field is from variant [TaskCreateResponseFieldValueFullName].
	FirstName string `json:"firstName"`
	// This field is from variant [TaskCreateResponseFieldValueFullName].
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

func (u TaskCreateResponseFieldValueUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u TaskCreateResponseFieldValueUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u TaskCreateResponseFieldValueUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u TaskCreateResponseFieldValueUnion) AsStringArray() (v []string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u TaskCreateResponseFieldValueUnion) AsAddress() (v TaskCreateResponseFieldValueAddress) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u TaskCreateResponseFieldValueUnion) AsFullName() (v TaskCreateResponseFieldValueFullName) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u TaskCreateResponseFieldValueUnion) RawJSON() string { return u.JSON.raw }

func (r *TaskCreateResponseFieldValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TaskCreateResponseFieldValueAddress struct {
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
func (r TaskCreateResponseFieldValueAddress) RawJSON() string { return r.JSON.raw }
func (r *TaskCreateResponseFieldValueAddress) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TaskCreateResponseFieldValueFullName struct {
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
func (r TaskCreateResponseFieldValueFullName) RawJSON() string { return r.JSON.raw }
func (r *TaskCreateResponseFieldValueFullName) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TaskCreateResponseRelationship struct {
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
func (r TaskCreateResponseRelationship) RawJSON() string { return r.JSON.raw }
func (r *TaskCreateResponseRelationship) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TaskDefinitionsResponse struct {
	// Map of field keys to their definitions, including both system and custom fields.
	FieldDefinitions map[string]TaskDefinitionsResponseFieldDefinition `json:"fieldDefinitions" api:"required"`
	// The object type these definitions belong to (e.g. `account`).
	ObjectType string `json:"objectType" api:"required"`
	// Map of relationship keys to their definitions.
	RelationshipDefinitions map[string]TaskDefinitionsResponseRelationshipDefinition `json:"relationshipDefinitions" api:"required"`
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
func (r TaskDefinitionsResponse) RawJSON() string { return r.JSON.raw }
func (r *TaskDefinitionsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TaskDefinitionsResponseFieldDefinition struct {
	// Description of the field, or null.
	Description string `json:"description" api:"required"`
	// Human-readable display name of the field.
	Label string `json:"label" api:"required"`
	// Type-specific configuration (e.g. select options, currency code).
	TypeConfiguration TaskDefinitionsResponseFieldDefinitionTypeConfiguration `json:"typeConfiguration" api:"required"`
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
func (r TaskDefinitionsResponseFieldDefinition) RawJSON() string { return r.JSON.raw }
func (r *TaskDefinitionsResponseFieldDefinition) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Type-specific configuration (e.g. select options, currency code).
type TaskDefinitionsResponseFieldDefinitionTypeConfiguration struct {
	// ISO 4217 3-letter currency code.
	Currency string `json:"currency"`
	// Social platform associated with this handle field.
	//
	// Any of "TWITTER", "LINKEDIN", "FACEBOOK", "INSTAGRAM".
	HandleService string `json:"handleService"`
	// Whether this field accepts multiple values.
	MultipleValues bool `json:"multipleValues"`
	// Available options for select fields.
	Options []TaskDefinitionsResponseFieldDefinitionTypeConfigurationOption `json:"options"`
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
func (r TaskDefinitionsResponseFieldDefinitionTypeConfiguration) RawJSON() string { return r.JSON.raw }
func (r *TaskDefinitionsResponseFieldDefinitionTypeConfiguration) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TaskDefinitionsResponseFieldDefinitionTypeConfigurationOption struct {
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
func (r TaskDefinitionsResponseFieldDefinitionTypeConfigurationOption) RawJSON() string {
	return r.JSON.raw
}
func (r *TaskDefinitionsResponseFieldDefinitionTypeConfigurationOption) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TaskDefinitionsResponseRelationshipDefinition struct {
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
func (r TaskDefinitionsResponseRelationshipDefinition) RawJSON() string { return r.JSON.raw }
func (r *TaskDefinitionsResponseRelationshipDefinition) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TaskListResponse struct {
	// Array of entity objects for the current page.
	Data []TaskListResponseData `json:"data" api:"required"`
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
func (r TaskListResponse) RawJSON() string { return r.JSON.raw }
func (r *TaskListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TaskListResponseData struct {
	// Unique identifier for the entity.
	ID string `json:"id" api:"required"`
	// ISO 8601 timestamp of when the entity was created.
	CreatedAt string `json:"createdAt" api:"required"`
	// Map of field names to their typed values. System fields are prefixed with `$`
	// (e.g. `$name`, `$email`); custom attributes use their bare slug.
	Fields map[string]TaskListResponseDataField `json:"fields" api:"required"`
	// URL to view the entity in the Lightfield web app, or null.
	HTTPLink string `json:"httpLink" api:"required"`
	// Map of relationship names to their associated entities. System relationships are
	// prefixed with `$` (e.g. `$owner`, `$contact`).
	Relationships map[string]TaskListResponseDataRelationship `json:"relationships" api:"required"`
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
func (r TaskListResponseData) RawJSON() string { return r.JSON.raw }
func (r *TaskListResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TaskListResponseDataField struct {
	// The field value, or null if unset.
	Value TaskListResponseDataFieldValueUnion `json:"value" api:"required"`
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
func (r TaskListResponseDataField) RawJSON() string { return r.JSON.raw }
func (r *TaskListResponseDataField) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// TaskListResponseDataFieldValueUnion contains all possible properties and values
// from [string], [float64], [bool], [[]string],
// [TaskListResponseDataFieldValueAddress],
// [TaskListResponseDataFieldValueFullName].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfStringArray]
type TaskListResponseDataFieldValueUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]string] instead of an object.
	OfStringArray []string `json:",inline"`
	// This field is from variant [TaskListResponseDataFieldValueAddress].
	City string `json:"city"`
	// This field is from variant [TaskListResponseDataFieldValueAddress].
	Country string `json:"country"`
	// This field is from variant [TaskListResponseDataFieldValueAddress].
	Latitude float64 `json:"latitude"`
	// This field is from variant [TaskListResponseDataFieldValueAddress].
	Longitude float64 `json:"longitude"`
	// This field is from variant [TaskListResponseDataFieldValueAddress].
	PostalCode string `json:"postalCode"`
	// This field is from variant [TaskListResponseDataFieldValueAddress].
	State string `json:"state"`
	// This field is from variant [TaskListResponseDataFieldValueAddress].
	Street string `json:"street"`
	// This field is from variant [TaskListResponseDataFieldValueAddress].
	Street2 string `json:"street2"`
	// This field is from variant [TaskListResponseDataFieldValueFullName].
	FirstName string `json:"firstName"`
	// This field is from variant [TaskListResponseDataFieldValueFullName].
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

func (u TaskListResponseDataFieldValueUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u TaskListResponseDataFieldValueUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u TaskListResponseDataFieldValueUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u TaskListResponseDataFieldValueUnion) AsStringArray() (v []string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u TaskListResponseDataFieldValueUnion) AsAddress() (v TaskListResponseDataFieldValueAddress) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u TaskListResponseDataFieldValueUnion) AsFullName() (v TaskListResponseDataFieldValueFullName) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u TaskListResponseDataFieldValueUnion) RawJSON() string { return u.JSON.raw }

func (r *TaskListResponseDataFieldValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TaskListResponseDataFieldValueAddress struct {
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
func (r TaskListResponseDataFieldValueAddress) RawJSON() string { return r.JSON.raw }
func (r *TaskListResponseDataFieldValueAddress) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TaskListResponseDataFieldValueFullName struct {
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
func (r TaskListResponseDataFieldValueFullName) RawJSON() string { return r.JSON.raw }
func (r *TaskListResponseDataFieldValueFullName) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TaskListResponseDataRelationship struct {
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
func (r TaskListResponseDataRelationship) RawJSON() string { return r.JSON.raw }
func (r *TaskListResponseDataRelationship) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TaskRetrieveResponse struct {
	// Unique identifier for the entity.
	ID string `json:"id" api:"required"`
	// ISO 8601 timestamp of when the entity was created.
	CreatedAt string `json:"createdAt" api:"required"`
	// Map of field names to their typed values. System fields are prefixed with `$`
	// (e.g. `$name`, `$email`); custom attributes use their bare slug.
	Fields map[string]TaskRetrieveResponseField `json:"fields" api:"required"`
	// URL to view the entity in the Lightfield web app, or null.
	HTTPLink string `json:"httpLink" api:"required"`
	// Map of relationship names to their associated entities. System relationships are
	// prefixed with `$` (e.g. `$owner`, `$contact`).
	Relationships map[string]TaskRetrieveResponseRelationship `json:"relationships" api:"required"`
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
func (r TaskRetrieveResponse) RawJSON() string { return r.JSON.raw }
func (r *TaskRetrieveResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TaskRetrieveResponseField struct {
	// The field value, or null if unset.
	Value TaskRetrieveResponseFieldValueUnion `json:"value" api:"required"`
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
func (r TaskRetrieveResponseField) RawJSON() string { return r.JSON.raw }
func (r *TaskRetrieveResponseField) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// TaskRetrieveResponseFieldValueUnion contains all possible properties and values
// from [string], [float64], [bool], [[]string],
// [TaskRetrieveResponseFieldValueAddress],
// [TaskRetrieveResponseFieldValueFullName].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfStringArray]
type TaskRetrieveResponseFieldValueUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]string] instead of an object.
	OfStringArray []string `json:",inline"`
	// This field is from variant [TaskRetrieveResponseFieldValueAddress].
	City string `json:"city"`
	// This field is from variant [TaskRetrieveResponseFieldValueAddress].
	Country string `json:"country"`
	// This field is from variant [TaskRetrieveResponseFieldValueAddress].
	Latitude float64 `json:"latitude"`
	// This field is from variant [TaskRetrieveResponseFieldValueAddress].
	Longitude float64 `json:"longitude"`
	// This field is from variant [TaskRetrieveResponseFieldValueAddress].
	PostalCode string `json:"postalCode"`
	// This field is from variant [TaskRetrieveResponseFieldValueAddress].
	State string `json:"state"`
	// This field is from variant [TaskRetrieveResponseFieldValueAddress].
	Street string `json:"street"`
	// This field is from variant [TaskRetrieveResponseFieldValueAddress].
	Street2 string `json:"street2"`
	// This field is from variant [TaskRetrieveResponseFieldValueFullName].
	FirstName string `json:"firstName"`
	// This field is from variant [TaskRetrieveResponseFieldValueFullName].
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

func (u TaskRetrieveResponseFieldValueUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u TaskRetrieveResponseFieldValueUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u TaskRetrieveResponseFieldValueUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u TaskRetrieveResponseFieldValueUnion) AsStringArray() (v []string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u TaskRetrieveResponseFieldValueUnion) AsAddress() (v TaskRetrieveResponseFieldValueAddress) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u TaskRetrieveResponseFieldValueUnion) AsFullName() (v TaskRetrieveResponseFieldValueFullName) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u TaskRetrieveResponseFieldValueUnion) RawJSON() string { return u.JSON.raw }

func (r *TaskRetrieveResponseFieldValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TaskRetrieveResponseFieldValueAddress struct {
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
func (r TaskRetrieveResponseFieldValueAddress) RawJSON() string { return r.JSON.raw }
func (r *TaskRetrieveResponseFieldValueAddress) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TaskRetrieveResponseFieldValueFullName struct {
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
func (r TaskRetrieveResponseFieldValueFullName) RawJSON() string { return r.JSON.raw }
func (r *TaskRetrieveResponseFieldValueFullName) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TaskRetrieveResponseRelationship struct {
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
func (r TaskRetrieveResponseRelationship) RawJSON() string { return r.JSON.raw }
func (r *TaskRetrieveResponseRelationship) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TaskUpdateResponse struct {
	// Unique identifier for the entity.
	ID string `json:"id" api:"required"`
	// ISO 8601 timestamp of when the entity was created.
	CreatedAt string `json:"createdAt" api:"required"`
	// Map of field names to their typed values. System fields are prefixed with `$`
	// (e.g. `$name`, `$email`); custom attributes use their bare slug.
	Fields map[string]TaskUpdateResponseField `json:"fields" api:"required"`
	// URL to view the entity in the Lightfield web app, or null.
	HTTPLink string `json:"httpLink" api:"required"`
	// Map of relationship names to their associated entities. System relationships are
	// prefixed with `$` (e.g. `$owner`, `$contact`).
	Relationships map[string]TaskUpdateResponseRelationship `json:"relationships" api:"required"`
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
func (r TaskUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *TaskUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TaskUpdateResponseField struct {
	// The field value, or null if unset.
	Value TaskUpdateResponseFieldValueUnion `json:"value" api:"required"`
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
func (r TaskUpdateResponseField) RawJSON() string { return r.JSON.raw }
func (r *TaskUpdateResponseField) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// TaskUpdateResponseFieldValueUnion contains all possible properties and values
// from [string], [float64], [bool], [[]string],
// [TaskUpdateResponseFieldValueAddress], [TaskUpdateResponseFieldValueFullName].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfStringArray]
type TaskUpdateResponseFieldValueUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]string] instead of an object.
	OfStringArray []string `json:",inline"`
	// This field is from variant [TaskUpdateResponseFieldValueAddress].
	City string `json:"city"`
	// This field is from variant [TaskUpdateResponseFieldValueAddress].
	Country string `json:"country"`
	// This field is from variant [TaskUpdateResponseFieldValueAddress].
	Latitude float64 `json:"latitude"`
	// This field is from variant [TaskUpdateResponseFieldValueAddress].
	Longitude float64 `json:"longitude"`
	// This field is from variant [TaskUpdateResponseFieldValueAddress].
	PostalCode string `json:"postalCode"`
	// This field is from variant [TaskUpdateResponseFieldValueAddress].
	State string `json:"state"`
	// This field is from variant [TaskUpdateResponseFieldValueAddress].
	Street string `json:"street"`
	// This field is from variant [TaskUpdateResponseFieldValueAddress].
	Street2 string `json:"street2"`
	// This field is from variant [TaskUpdateResponseFieldValueFullName].
	FirstName string `json:"firstName"`
	// This field is from variant [TaskUpdateResponseFieldValueFullName].
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

func (u TaskUpdateResponseFieldValueUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u TaskUpdateResponseFieldValueUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u TaskUpdateResponseFieldValueUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u TaskUpdateResponseFieldValueUnion) AsStringArray() (v []string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u TaskUpdateResponseFieldValueUnion) AsAddress() (v TaskUpdateResponseFieldValueAddress) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u TaskUpdateResponseFieldValueUnion) AsFullName() (v TaskUpdateResponseFieldValueFullName) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u TaskUpdateResponseFieldValueUnion) RawJSON() string { return u.JSON.raw }

func (r *TaskUpdateResponseFieldValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TaskUpdateResponseFieldValueAddress struct {
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
func (r TaskUpdateResponseFieldValueAddress) RawJSON() string { return r.JSON.raw }
func (r *TaskUpdateResponseFieldValueAddress) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TaskUpdateResponseFieldValueFullName struct {
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
func (r TaskUpdateResponseFieldValueFullName) RawJSON() string { return r.JSON.raw }
func (r *TaskUpdateResponseFieldValueFullName) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TaskUpdateResponseRelationship struct {
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
func (r TaskUpdateResponseRelationship) RawJSON() string { return r.JSON.raw }
func (r *TaskUpdateResponseRelationship) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TaskNewParams struct {
	// Field values for the new task. Tasks only support the documented system fields,
	// all prefixed with `$` (e.g. `$title`, `$status`). Required: `$title` (string)
	// and `$status` (one of `TODO`, `IN_PROGRESS`, `COMPLETE`, `CANCELLED`). Call the
	// <u>[definitions endpoint](/api/resources/task/methods/definitions)</u> to
	// discover the available fields. See
	// <u>[Fields and relationships](/using-the-api/fields-and-relationships/)</u> for
	// value type details.
	Fields TaskNewParamsFields `json:"fields,omitzero" api:"required"`
	// Relationships to set on the new task. System relationships use a `$` prefix
	// (e.g. `$account`, `$assignedTo`); custom relationships use their bare slug.
	// `$assignedTo` is required. Each value is a single entity ID or an array of IDs.
	// Call the <u>[definitions endpoint](/api/resources/task/methods/definitions)</u>
	// to list available relationship keys.
	Relationships map[string]TaskNewParamsRelationshipUnion `json:"relationships,omitzero" api:"required"`
	paramObj
}

func (r TaskNewParams) MarshalJSON() (data []byte, err error) {
	type shadow TaskNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TaskNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Field values for the new task. Tasks only support the documented system fields,
// all prefixed with `$` (e.g. `$title`, `$status`). Required: `$title` (string)
// and `$status` (one of `TODO`, `IN_PROGRESS`, `COMPLETE`, `CANCELLED`). Call the
// <u>[definitions endpoint](/api/resources/task/methods/definitions)</u> to
// discover the available fields. See
// <u>[Fields and relationships](/using-the-api/fields-and-relationships/)</u> for
// value type details.
//
// The properties Status, Title are required.
type TaskNewParamsFields struct {
	// Task status. One of: `TODO`, `IN_PROGRESS`, `COMPLETE`, `CANCELLED`.
	Status string `json:"$status" api:"required"`
	// Title of the task.
	Title string `json:"$title" api:"required"`
	// Description of the task in markdown format.
	Description param.Opt[string] `json:"$description,omitzero"`
	// Due date as an ISO 8601 datetime string.
	DueAt param.Opt[string] `json:"$dueAt,omitzero"`
	paramObj
}

func (r TaskNewParamsFields) MarshalJSON() (data []byte, err error) {
	type shadow TaskNewParamsFields
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TaskNewParamsFields) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type TaskNewParamsRelationshipUnion struct {
	OfString      param.Opt[string] `json:",omitzero,inline"`
	OfStringArray []string          `json:",omitzero,inline"`
	paramUnion
}

func (u TaskNewParamsRelationshipUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray)
}
func (u *TaskNewParamsRelationshipUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

type TaskUpdateParams struct {
	// Field values to update — only provided fields are modified; omitted fields are
	// left unchanged. Tasks only support the documented system fields, all prefixed
	// with `$` (e.g. `$title`, `$status`). Call the
	// <u>[definitions endpoint](/api/resources/task/methods/definitions)</u> for
	// available fields. See
	// <u>[Fields and relationships](/using-the-api/fields-and-relationships/)</u> for
	// value type details.
	Fields TaskUpdateParamsFields `json:"fields,omitzero"`
	// Relationship operations to apply. System relationships use a `$` prefix (e.g.
	// `$account`, `$assignedTo`). Each value is an operation object with `add`,
	// `remove`, or `replace`.
	Relationships map[string]TaskUpdateParamsRelationship `json:"relationships,omitzero"`
	paramObj
}

func (r TaskUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow TaskUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TaskUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Field values to update — only provided fields are modified; omitted fields are
// left unchanged. Tasks only support the documented system fields, all prefixed
// with `$` (e.g. `$title`, `$status`). Call the
// <u>[definitions endpoint](/api/resources/task/methods/definitions)</u> for
// available fields. See
// <u>[Fields and relationships](/using-the-api/fields-and-relationships/)</u> for
// value type details.
type TaskUpdateParamsFields struct {
	// Description of the task in markdown format.
	Description param.Opt[string] `json:"$description,omitzero"`
	// Due date as an ISO 8601 datetime string.
	DueAt param.Opt[string] `json:"$dueAt,omitzero"`
	// Task status. One of: `TODO`, `IN_PROGRESS`, `COMPLETE`, `CANCELLED`.
	Status param.Opt[string] `json:"$status,omitzero"`
	// Title of the task.
	Title param.Opt[string] `json:"$title,omitzero"`
	paramObj
}

func (r TaskUpdateParamsFields) MarshalJSON() (data []byte, err error) {
	type shadow TaskUpdateParamsFields
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TaskUpdateParamsFields) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An operation to modify a relationship. Provide one of `add`, `remove`, or
// `replace`.
type TaskUpdateParamsRelationship struct {
	// Entity ID(s) to add to the relationship.
	Add TaskUpdateParamsRelationshipAddUnion `json:"add,omitzero"`
	// Entity ID(s) to remove from the relationship.
	Remove TaskUpdateParamsRelationshipRemoveUnion `json:"remove,omitzero"`
	// Entity ID(s) to set as the entire relationship, replacing all existing
	// associations.
	Replace TaskUpdateParamsRelationshipReplaceUnion `json:"replace,omitzero"`
	paramObj
}

func (r TaskUpdateParamsRelationship) MarshalJSON() (data []byte, err error) {
	type shadow TaskUpdateParamsRelationship
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TaskUpdateParamsRelationship) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type TaskUpdateParamsRelationshipAddUnion struct {
	OfString      param.Opt[string] `json:",omitzero,inline"`
	OfStringArray []string          `json:",omitzero,inline"`
	paramUnion
}

func (u TaskUpdateParamsRelationshipAddUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray)
}
func (u *TaskUpdateParamsRelationshipAddUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type TaskUpdateParamsRelationshipRemoveUnion struct {
	OfString      param.Opt[string] `json:",omitzero,inline"`
	OfStringArray []string          `json:",omitzero,inline"`
	paramUnion
}

func (u TaskUpdateParamsRelationshipRemoveUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray)
}
func (u *TaskUpdateParamsRelationshipRemoveUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type TaskUpdateParamsRelationshipReplaceUnion struct {
	OfString      param.Opt[string] `json:",omitzero,inline"`
	OfStringArray []string          `json:",omitzero,inline"`
	paramUnion
}

func (u TaskUpdateParamsRelationshipReplaceUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray)
}
func (u *TaskUpdateParamsRelationshipReplaceUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

type TaskListParams struct {
	// Maximum number of records to return. Defaults to 25, maximum 25.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Number of records to skip for pagination. Defaults to 0.
	Offset param.Opt[int64] `query:"offset,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [TaskListParams]'s query parameters as `url.Values`.
func (r TaskListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
