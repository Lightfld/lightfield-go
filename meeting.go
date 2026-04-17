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

// Meetings represent synced or manually created interactions in Lightfield. Read
// responses are privacy-aware and may be redacted based on the caller. For
// transcript uploads and attachment flows, see
// <u>[Uploading meeting transcripts](/using-the-api/uploading-meeting-transcripts/)</u>.
//
// MeetingService contains methods and other services that help with interacting
// with the Lightfield API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewMeetingService] method instead.
type MeetingService struct {
	Options []option.RequestOption
}

// NewMeetingService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewMeetingService(opts ...option.RequestOption) (r MeetingService) {
	r = MeetingService{}
	r.Options = opts
	return
}

// Creates a new meeting record. This endpoint only supports creation of meetings
// in the past. The `$title`, `$startDate`, and `$endDate` fields are required.
// Only the `$transcript` relationship is writable on create; all other meeting
// relationships are system-managed. The response is privacy-aware and includes a
// read-only `accessLevel`. See
// <u>[Uploading meeting transcripts](/using-the-api/uploading-meeting-transcripts/)</u>
// for the full file upload and transcript attachment flow.
//
// Supports idempotency via the `Idempotency-Key` header.
//
// **[Required scope](/using-the-api/scopes/):** `meetings:create`
//
// **[Rate limit category](/using-the-api/rate-limits/):** Write
func (r *MeetingService) New(ctx context.Context, body MeetingNewParams, opts ...option.RequestOption) (res *MeetingCreateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/meetings"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Retrieves a single meeting by its ID. Meeting fields and transcript visibility
// are redacted based on the caller-specific privacy resolution, and the response
// includes a read-only `accessLevel`.
//
// **[Required scope](/using-the-api/scopes/):** `meetings:read`
//
// **[Rate limit category](/using-the-api/rate-limits/):** Read
func (r *MeetingService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *MeetingRetrieveResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/meetings/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Updates an existing meeting by ID. Only included fields and relationships are
// modified.
//
// Only `fields.$privacySetting` and `relationships.$transcript.replace` are
// writable. Use `$transcript.replace` to set the meeting transcript. Clearing or
// removing `$transcript` is not supported. The response is privacy-aware and
// includes a read-only `accessLevel`. See
// <u>[Uploading meeting transcripts](/using-the-api/uploading-meeting-transcripts/)</u>
// for the full file upload and transcript attachment flow.
//
// Supports idempotency via the `Idempotency-Key` header.
//
// **[Required scope](/using-the-api/scopes/):** `meetings:update`
//
// **[Rate limit category](/using-the-api/rate-limits/):** Write
func (r *MeetingService) Update(ctx context.Context, id string, body MeetingUpdateParams, opts ...option.RequestOption) (res *MeetingUpdateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/meetings/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Returns a paginated list of meetings. Use `offset` and `limit` to paginate
// through results. Each meeting is privacy-filtered per caller, includes a
// read-only `accessLevel`, and may redact transcript or content fields based on
// the caller-specific privacy resolution. See
// <u>[List endpoints](/using-the-api/list-endpoints/)</u> for more information
// about pagination.
//
// **[Required scope](/using-the-api/scopes/):** `meetings:read`
//
// **[Rate limit category](/using-the-api/rate-limits/):** Search
func (r *MeetingService) List(ctx context.Context, query MeetingListParams, opts ...option.RequestOption) (res *MeetingListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/meetings"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

type MeetingCreateResponse struct {
	// Unique identifier for the entity.
	ID string `json:"id" api:"required"`
	// The caller's resolved access level for this meeting.
	//
	// Any of "FULL", "METADATA".
	AccessLevel MeetingCreateResponseAccessLevel `json:"accessLevel" api:"required"`
	// ISO 8601 timestamp of when the entity was created.
	CreatedAt string `json:"createdAt" api:"required"`
	// Map of field names to their typed values. System fields are prefixed with `$`
	// (e.g. `$name`, `$email`); custom attributes use their bare slug.
	Fields map[string]MeetingCreateResponseField `json:"fields" api:"required"`
	// URL to view the entity in the Lightfield web app, or null.
	HTTPLink string `json:"httpLink" api:"required"`
	// Always `meeting`.
	//
	// Any of "meeting".
	ObjectType MeetingCreateResponseObjectType `json:"objectType" api:"required"`
	// Map of relationship names to their associated entities. System relationships are
	// prefixed with `$` (e.g. `$owner`, `$contact`).
	Relationships map[string]MeetingCreateResponseRelationship `json:"relationships" api:"required"`
	// ISO 8601 timestamp of when the entity was last updated, or null.
	UpdatedAt string `json:"updatedAt" api:"required"`
	// External identifier for the entity, or null if unset.
	ExternalID string `json:"externalId" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID            respjson.Field
		AccessLevel   respjson.Field
		CreatedAt     respjson.Field
		Fields        respjson.Field
		HTTPLink      respjson.Field
		ObjectType    respjson.Field
		Relationships respjson.Field
		UpdatedAt     respjson.Field
		ExternalID    respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MeetingCreateResponse) RawJSON() string { return r.JSON.raw }
func (r *MeetingCreateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The caller's resolved access level for this meeting.
type MeetingCreateResponseAccessLevel string

const (
	MeetingCreateResponseAccessLevelFull     MeetingCreateResponseAccessLevel = "FULL"
	MeetingCreateResponseAccessLevelMetadata MeetingCreateResponseAccessLevel = "METADATA"
)

type MeetingCreateResponseField struct {
	// The field value, or null if unset.
	Value MeetingCreateResponseFieldValueUnion `json:"value" api:"required"`
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
func (r MeetingCreateResponseField) RawJSON() string { return r.JSON.raw }
func (r *MeetingCreateResponseField) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// MeetingCreateResponseFieldValueUnion contains all possible properties and values
// from [string], [float64], [bool], [[]string],
// [MeetingCreateResponseFieldValueAddress],
// [MeetingCreateResponseFieldValueFullName].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfStringArray]
type MeetingCreateResponseFieldValueUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]string] instead of an object.
	OfStringArray []string `json:",inline"`
	// This field is from variant [MeetingCreateResponseFieldValueAddress].
	City string `json:"city"`
	// This field is from variant [MeetingCreateResponseFieldValueAddress].
	Country string `json:"country"`
	// This field is from variant [MeetingCreateResponseFieldValueAddress].
	Latitude float64 `json:"latitude"`
	// This field is from variant [MeetingCreateResponseFieldValueAddress].
	Longitude float64 `json:"longitude"`
	// This field is from variant [MeetingCreateResponseFieldValueAddress].
	PostalCode string `json:"postalCode"`
	// This field is from variant [MeetingCreateResponseFieldValueAddress].
	State string `json:"state"`
	// This field is from variant [MeetingCreateResponseFieldValueAddress].
	Street string `json:"street"`
	// This field is from variant [MeetingCreateResponseFieldValueAddress].
	Street2 string `json:"street2"`
	// This field is from variant [MeetingCreateResponseFieldValueFullName].
	FirstName string `json:"firstName"`
	// This field is from variant [MeetingCreateResponseFieldValueFullName].
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

func (u MeetingCreateResponseFieldValueUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MeetingCreateResponseFieldValueUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MeetingCreateResponseFieldValueUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MeetingCreateResponseFieldValueUnion) AsStringArray() (v []string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MeetingCreateResponseFieldValueUnion) AsAddress() (v MeetingCreateResponseFieldValueAddress) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MeetingCreateResponseFieldValueUnion) AsFullName() (v MeetingCreateResponseFieldValueFullName) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u MeetingCreateResponseFieldValueUnion) RawJSON() string { return u.JSON.raw }

func (r *MeetingCreateResponseFieldValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MeetingCreateResponseFieldValueAddress struct {
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
func (r MeetingCreateResponseFieldValueAddress) RawJSON() string { return r.JSON.raw }
func (r *MeetingCreateResponseFieldValueAddress) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MeetingCreateResponseFieldValueFullName struct {
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
func (r MeetingCreateResponseFieldValueFullName) RawJSON() string { return r.JSON.raw }
func (r *MeetingCreateResponseFieldValueFullName) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Always `meeting`.
type MeetingCreateResponseObjectType string

const (
	MeetingCreateResponseObjectTypeMeeting MeetingCreateResponseObjectType = "meeting"
)

type MeetingCreateResponseRelationship struct {
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
func (r MeetingCreateResponseRelationship) RawJSON() string { return r.JSON.raw }
func (r *MeetingCreateResponseRelationship) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MeetingListResponse struct {
	// Array of meeting objects for the current page.
	Data []MeetingListResponseData `json:"data" api:"required"`
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
func (r MeetingListResponse) RawJSON() string { return r.JSON.raw }
func (r *MeetingListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MeetingListResponseData struct {
	// Unique identifier for the entity.
	ID string `json:"id" api:"required"`
	// The caller's resolved access level for this meeting.
	//
	// Any of "FULL", "METADATA".
	AccessLevel string `json:"accessLevel" api:"required"`
	// ISO 8601 timestamp of when the entity was created.
	CreatedAt string `json:"createdAt" api:"required"`
	// Map of field names to their typed values. System fields are prefixed with `$`
	// (e.g. `$name`, `$email`); custom attributes use their bare slug.
	Fields map[string]MeetingListResponseDataField `json:"fields" api:"required"`
	// URL to view the entity in the Lightfield web app, or null.
	HTTPLink string `json:"httpLink" api:"required"`
	// Always `meeting`.
	//
	// Any of "meeting".
	ObjectType string `json:"objectType" api:"required"`
	// Map of relationship names to their associated entities. System relationships are
	// prefixed with `$` (e.g. `$owner`, `$contact`).
	Relationships map[string]MeetingListResponseDataRelationship `json:"relationships" api:"required"`
	// ISO 8601 timestamp of when the entity was last updated, or null.
	UpdatedAt string `json:"updatedAt" api:"required"`
	// External identifier for the entity, or null if unset.
	ExternalID string `json:"externalId" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID            respjson.Field
		AccessLevel   respjson.Field
		CreatedAt     respjson.Field
		Fields        respjson.Field
		HTTPLink      respjson.Field
		ObjectType    respjson.Field
		Relationships respjson.Field
		UpdatedAt     respjson.Field
		ExternalID    respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MeetingListResponseData) RawJSON() string { return r.JSON.raw }
func (r *MeetingListResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MeetingListResponseDataField struct {
	// The field value, or null if unset.
	Value MeetingListResponseDataFieldValueUnion `json:"value" api:"required"`
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
func (r MeetingListResponseDataField) RawJSON() string { return r.JSON.raw }
func (r *MeetingListResponseDataField) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// MeetingListResponseDataFieldValueUnion contains all possible properties and
// values from [string], [float64], [bool], [[]string],
// [MeetingListResponseDataFieldValueAddress],
// [MeetingListResponseDataFieldValueFullName].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfStringArray]
type MeetingListResponseDataFieldValueUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]string] instead of an object.
	OfStringArray []string `json:",inline"`
	// This field is from variant [MeetingListResponseDataFieldValueAddress].
	City string `json:"city"`
	// This field is from variant [MeetingListResponseDataFieldValueAddress].
	Country string `json:"country"`
	// This field is from variant [MeetingListResponseDataFieldValueAddress].
	Latitude float64 `json:"latitude"`
	// This field is from variant [MeetingListResponseDataFieldValueAddress].
	Longitude float64 `json:"longitude"`
	// This field is from variant [MeetingListResponseDataFieldValueAddress].
	PostalCode string `json:"postalCode"`
	// This field is from variant [MeetingListResponseDataFieldValueAddress].
	State string `json:"state"`
	// This field is from variant [MeetingListResponseDataFieldValueAddress].
	Street string `json:"street"`
	// This field is from variant [MeetingListResponseDataFieldValueAddress].
	Street2 string `json:"street2"`
	// This field is from variant [MeetingListResponseDataFieldValueFullName].
	FirstName string `json:"firstName"`
	// This field is from variant [MeetingListResponseDataFieldValueFullName].
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

func (u MeetingListResponseDataFieldValueUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MeetingListResponseDataFieldValueUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MeetingListResponseDataFieldValueUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MeetingListResponseDataFieldValueUnion) AsStringArray() (v []string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MeetingListResponseDataFieldValueUnion) AsAddress() (v MeetingListResponseDataFieldValueAddress) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MeetingListResponseDataFieldValueUnion) AsFullName() (v MeetingListResponseDataFieldValueFullName) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u MeetingListResponseDataFieldValueUnion) RawJSON() string { return u.JSON.raw }

func (r *MeetingListResponseDataFieldValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MeetingListResponseDataFieldValueAddress struct {
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
func (r MeetingListResponseDataFieldValueAddress) RawJSON() string { return r.JSON.raw }
func (r *MeetingListResponseDataFieldValueAddress) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MeetingListResponseDataFieldValueFullName struct {
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
func (r MeetingListResponseDataFieldValueFullName) RawJSON() string { return r.JSON.raw }
func (r *MeetingListResponseDataFieldValueFullName) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MeetingListResponseDataRelationship struct {
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
func (r MeetingListResponseDataRelationship) RawJSON() string { return r.JSON.raw }
func (r *MeetingListResponseDataRelationship) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MeetingRetrieveResponse struct {
	// Unique identifier for the entity.
	ID string `json:"id" api:"required"`
	// The caller's resolved access level for this meeting.
	//
	// Any of "FULL", "METADATA".
	AccessLevel MeetingRetrieveResponseAccessLevel `json:"accessLevel" api:"required"`
	// ISO 8601 timestamp of when the entity was created.
	CreatedAt string `json:"createdAt" api:"required"`
	// Map of field names to their typed values. System fields are prefixed with `$`
	// (e.g. `$name`, `$email`); custom attributes use their bare slug.
	Fields map[string]MeetingRetrieveResponseField `json:"fields" api:"required"`
	// URL to view the entity in the Lightfield web app, or null.
	HTTPLink string `json:"httpLink" api:"required"`
	// Always `meeting`.
	//
	// Any of "meeting".
	ObjectType MeetingRetrieveResponseObjectType `json:"objectType" api:"required"`
	// Map of relationship names to their associated entities. System relationships are
	// prefixed with `$` (e.g. `$owner`, `$contact`).
	Relationships map[string]MeetingRetrieveResponseRelationship `json:"relationships" api:"required"`
	// ISO 8601 timestamp of when the entity was last updated, or null.
	UpdatedAt string `json:"updatedAt" api:"required"`
	// External identifier for the entity, or null if unset.
	ExternalID string `json:"externalId" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID            respjson.Field
		AccessLevel   respjson.Field
		CreatedAt     respjson.Field
		Fields        respjson.Field
		HTTPLink      respjson.Field
		ObjectType    respjson.Field
		Relationships respjson.Field
		UpdatedAt     respjson.Field
		ExternalID    respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MeetingRetrieveResponse) RawJSON() string { return r.JSON.raw }
func (r *MeetingRetrieveResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The caller's resolved access level for this meeting.
type MeetingRetrieveResponseAccessLevel string

const (
	MeetingRetrieveResponseAccessLevelFull     MeetingRetrieveResponseAccessLevel = "FULL"
	MeetingRetrieveResponseAccessLevelMetadata MeetingRetrieveResponseAccessLevel = "METADATA"
)

type MeetingRetrieveResponseField struct {
	// The field value, or null if unset.
	Value MeetingRetrieveResponseFieldValueUnion `json:"value" api:"required"`
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
func (r MeetingRetrieveResponseField) RawJSON() string { return r.JSON.raw }
func (r *MeetingRetrieveResponseField) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// MeetingRetrieveResponseFieldValueUnion contains all possible properties and
// values from [string], [float64], [bool], [[]string],
// [MeetingRetrieveResponseFieldValueAddress],
// [MeetingRetrieveResponseFieldValueFullName].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfStringArray]
type MeetingRetrieveResponseFieldValueUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]string] instead of an object.
	OfStringArray []string `json:",inline"`
	// This field is from variant [MeetingRetrieveResponseFieldValueAddress].
	City string `json:"city"`
	// This field is from variant [MeetingRetrieveResponseFieldValueAddress].
	Country string `json:"country"`
	// This field is from variant [MeetingRetrieveResponseFieldValueAddress].
	Latitude float64 `json:"latitude"`
	// This field is from variant [MeetingRetrieveResponseFieldValueAddress].
	Longitude float64 `json:"longitude"`
	// This field is from variant [MeetingRetrieveResponseFieldValueAddress].
	PostalCode string `json:"postalCode"`
	// This field is from variant [MeetingRetrieveResponseFieldValueAddress].
	State string `json:"state"`
	// This field is from variant [MeetingRetrieveResponseFieldValueAddress].
	Street string `json:"street"`
	// This field is from variant [MeetingRetrieveResponseFieldValueAddress].
	Street2 string `json:"street2"`
	// This field is from variant [MeetingRetrieveResponseFieldValueFullName].
	FirstName string `json:"firstName"`
	// This field is from variant [MeetingRetrieveResponseFieldValueFullName].
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

func (u MeetingRetrieveResponseFieldValueUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MeetingRetrieveResponseFieldValueUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MeetingRetrieveResponseFieldValueUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MeetingRetrieveResponseFieldValueUnion) AsStringArray() (v []string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MeetingRetrieveResponseFieldValueUnion) AsAddress() (v MeetingRetrieveResponseFieldValueAddress) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MeetingRetrieveResponseFieldValueUnion) AsFullName() (v MeetingRetrieveResponseFieldValueFullName) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u MeetingRetrieveResponseFieldValueUnion) RawJSON() string { return u.JSON.raw }

func (r *MeetingRetrieveResponseFieldValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MeetingRetrieveResponseFieldValueAddress struct {
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
func (r MeetingRetrieveResponseFieldValueAddress) RawJSON() string { return r.JSON.raw }
func (r *MeetingRetrieveResponseFieldValueAddress) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MeetingRetrieveResponseFieldValueFullName struct {
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
func (r MeetingRetrieveResponseFieldValueFullName) RawJSON() string { return r.JSON.raw }
func (r *MeetingRetrieveResponseFieldValueFullName) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Always `meeting`.
type MeetingRetrieveResponseObjectType string

const (
	MeetingRetrieveResponseObjectTypeMeeting MeetingRetrieveResponseObjectType = "meeting"
)

type MeetingRetrieveResponseRelationship struct {
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
func (r MeetingRetrieveResponseRelationship) RawJSON() string { return r.JSON.raw }
func (r *MeetingRetrieveResponseRelationship) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MeetingUpdateResponse struct {
	// Unique identifier for the entity.
	ID string `json:"id" api:"required"`
	// The caller's resolved access level for this meeting.
	//
	// Any of "FULL", "METADATA".
	AccessLevel MeetingUpdateResponseAccessLevel `json:"accessLevel" api:"required"`
	// ISO 8601 timestamp of when the entity was created.
	CreatedAt string `json:"createdAt" api:"required"`
	// Map of field names to their typed values. System fields are prefixed with `$`
	// (e.g. `$name`, `$email`); custom attributes use their bare slug.
	Fields map[string]MeetingUpdateResponseField `json:"fields" api:"required"`
	// URL to view the entity in the Lightfield web app, or null.
	HTTPLink string `json:"httpLink" api:"required"`
	// Always `meeting`.
	//
	// Any of "meeting".
	ObjectType MeetingUpdateResponseObjectType `json:"objectType" api:"required"`
	// Map of relationship names to their associated entities. System relationships are
	// prefixed with `$` (e.g. `$owner`, `$contact`).
	Relationships map[string]MeetingUpdateResponseRelationship `json:"relationships" api:"required"`
	// ISO 8601 timestamp of when the entity was last updated, or null.
	UpdatedAt string `json:"updatedAt" api:"required"`
	// External identifier for the entity, or null if unset.
	ExternalID string `json:"externalId" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID            respjson.Field
		AccessLevel   respjson.Field
		CreatedAt     respjson.Field
		Fields        respjson.Field
		HTTPLink      respjson.Field
		ObjectType    respjson.Field
		Relationships respjson.Field
		UpdatedAt     respjson.Field
		ExternalID    respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MeetingUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *MeetingUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The caller's resolved access level for this meeting.
type MeetingUpdateResponseAccessLevel string

const (
	MeetingUpdateResponseAccessLevelFull     MeetingUpdateResponseAccessLevel = "FULL"
	MeetingUpdateResponseAccessLevelMetadata MeetingUpdateResponseAccessLevel = "METADATA"
)

type MeetingUpdateResponseField struct {
	// The field value, or null if unset.
	Value MeetingUpdateResponseFieldValueUnion `json:"value" api:"required"`
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
func (r MeetingUpdateResponseField) RawJSON() string { return r.JSON.raw }
func (r *MeetingUpdateResponseField) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// MeetingUpdateResponseFieldValueUnion contains all possible properties and values
// from [string], [float64], [bool], [[]string],
// [MeetingUpdateResponseFieldValueAddress],
// [MeetingUpdateResponseFieldValueFullName].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfStringArray]
type MeetingUpdateResponseFieldValueUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]string] instead of an object.
	OfStringArray []string `json:",inline"`
	// This field is from variant [MeetingUpdateResponseFieldValueAddress].
	City string `json:"city"`
	// This field is from variant [MeetingUpdateResponseFieldValueAddress].
	Country string `json:"country"`
	// This field is from variant [MeetingUpdateResponseFieldValueAddress].
	Latitude float64 `json:"latitude"`
	// This field is from variant [MeetingUpdateResponseFieldValueAddress].
	Longitude float64 `json:"longitude"`
	// This field is from variant [MeetingUpdateResponseFieldValueAddress].
	PostalCode string `json:"postalCode"`
	// This field is from variant [MeetingUpdateResponseFieldValueAddress].
	State string `json:"state"`
	// This field is from variant [MeetingUpdateResponseFieldValueAddress].
	Street string `json:"street"`
	// This field is from variant [MeetingUpdateResponseFieldValueAddress].
	Street2 string `json:"street2"`
	// This field is from variant [MeetingUpdateResponseFieldValueFullName].
	FirstName string `json:"firstName"`
	// This field is from variant [MeetingUpdateResponseFieldValueFullName].
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

func (u MeetingUpdateResponseFieldValueUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MeetingUpdateResponseFieldValueUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MeetingUpdateResponseFieldValueUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MeetingUpdateResponseFieldValueUnion) AsStringArray() (v []string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MeetingUpdateResponseFieldValueUnion) AsAddress() (v MeetingUpdateResponseFieldValueAddress) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MeetingUpdateResponseFieldValueUnion) AsFullName() (v MeetingUpdateResponseFieldValueFullName) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u MeetingUpdateResponseFieldValueUnion) RawJSON() string { return u.JSON.raw }

func (r *MeetingUpdateResponseFieldValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MeetingUpdateResponseFieldValueAddress struct {
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
func (r MeetingUpdateResponseFieldValueAddress) RawJSON() string { return r.JSON.raw }
func (r *MeetingUpdateResponseFieldValueAddress) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MeetingUpdateResponseFieldValueFullName struct {
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
func (r MeetingUpdateResponseFieldValueFullName) RawJSON() string { return r.JSON.raw }
func (r *MeetingUpdateResponseFieldValueFullName) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Always `meeting`.
type MeetingUpdateResponseObjectType string

const (
	MeetingUpdateResponseObjectTypeMeeting MeetingUpdateResponseObjectType = "meeting"
)

type MeetingUpdateResponseRelationship struct {
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
func (r MeetingUpdateResponseRelationship) RawJSON() string { return r.JSON.raw }
func (r *MeetingUpdateResponseRelationship) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MeetingNewParams struct {
	// Field values for the new MANUAL meeting. System fields use a `$` prefix (for
	// example `$title`, `$startDate`, `$endDate`). Required: `$title`, `$startDate`,
	// and `$endDate`. `$organizerEmail` accepts a single email address when provided;
	// `$attendeeEmails` accepts an array of email addresses. See
	// <u>[Fields and relationships](/using-the-api/fields-and-relationships/)</u> for
	// value type details.
	Fields MeetingNewParamsFields `json:"fields,omitzero" api:"required"`
	// When true, the initial post-create meeting sync may auto-create account and
	// contact records for external attendees.
	AutoCreateRecords param.Opt[bool] `json:"autoCreateRecords,omitzero"`
	// Relationships to set on the new meeting. Only `$transcript` is writable on
	// create; all other meeting relationships are system-managed.
	Relationships MeetingNewParamsRelationships `json:"relationships,omitzero"`
	paramObj
}

func (r MeetingNewParams) MarshalJSON() (data []byte, err error) {
	type shadow MeetingNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MeetingNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Field values for the new MANUAL meeting. System fields use a `$` prefix (for
// example `$title`, `$startDate`, `$endDate`). Required: `$title`, `$startDate`,
// and `$endDate`. `$organizerEmail` accepts a single email address when provided;
// `$attendeeEmails` accepts an array of email addresses. See
// <u>[Fields and relationships](/using-the-api/fields-and-relationships/)</u> for
// value type details.
//
// The properties EndDate, StartDate, Title are required.
type MeetingNewParamsFields struct {
	// The end time of the meeting in ISO 8601 format. Must be in the past.
	EndDate string `json:"$endDate" api:"required"`
	// The start time of the meeting in ISO 8601 format. Must be in the past.
	StartDate string `json:"$startDate" api:"required"`
	// The title of the meeting.
	Title string `json:"$title" api:"required"`
	// A description of the meeting.
	Description param.Opt[string] `json:"$description,omitzero"`
	// The URL for the meeting.
	MeetingURL param.Opt[string] `json:"$meetingUrl,omitzero"`
	// The email address of the meeting organizer. This field accepts a single email
	// address.
	OrganizerEmail param.Opt[string] `json:"$organizerEmail,omitzero"`
	// The privacy setting for the meeting (`FULL` or `METADATA`).
	//
	// Any of "FULL", "METADATA".
	PrivacySetting string `json:"$privacySetting,omitzero"`
	// A list of attendee email addresses.
	AttendeeEmails []string `json:"$attendeeEmails,omitzero"`
	paramObj
}

func (r MeetingNewParamsFields) MarshalJSON() (data []byte, err error) {
	type shadow MeetingNewParamsFields
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MeetingNewParamsFields) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[MeetingNewParamsFields](
		"$privacySetting", "FULL", "METADATA",
	)
}

// Relationships to set on the new meeting. Only `$transcript` is writable on
// create; all other meeting relationships are system-managed.
//
// The property Transcript is required.
type MeetingNewParamsRelationships struct {
	// The ID of the file to attach as the meeting transcript when creating the
	// meeting. Only one transcript can be attached to a meeting.
	Transcript MeetingNewParamsRelationshipsTranscriptUnion `json:"$transcript,omitzero" api:"required"`
	paramObj
}

func (r MeetingNewParamsRelationships) MarshalJSON() (data []byte, err error) {
	type shadow MeetingNewParamsRelationships
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MeetingNewParamsRelationships) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type MeetingNewParamsRelationshipsTranscriptUnion struct {
	OfString      param.Opt[string] `json:",omitzero,inline"`
	OfStringArray []string          `json:",omitzero,inline"`
	paramUnion
}

func (u MeetingNewParamsRelationshipsTranscriptUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray)
}
func (u *MeetingNewParamsRelationshipsTranscriptUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

type MeetingUpdateParams struct {
	// Field values to update. Only `$privacySetting` is writable, and omitted fields
	// are left unchanged.
	Fields MeetingUpdateParamsFields `json:"fields,omitzero"`
	// Relationship operations to apply. Only `$transcript.replace` is supported;
	// removing or clearing `$transcript` is not supported.
	Relationships MeetingUpdateParamsRelationships `json:"relationships,omitzero"`
	paramObj
}

func (r MeetingUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow MeetingUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MeetingUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Field values to update. Only `$privacySetting` is writable, and omitted fields
// are left unchanged.
//
// The property PrivacySetting is required.
type MeetingUpdateParamsFields struct {
	// The privacy setting for the meeting.
	//
	// Any of "FULL", "METADATA".
	PrivacySetting string `json:"$privacySetting,omitzero" api:"required"`
	paramObj
}

func (r MeetingUpdateParamsFields) MarshalJSON() (data []byte, err error) {
	type shadow MeetingUpdateParamsFields
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MeetingUpdateParamsFields) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[MeetingUpdateParamsFields](
		"$privacySetting", "FULL", "METADATA",
	)
}

// Relationship operations to apply. Only `$transcript.replace` is supported;
// removing or clearing `$transcript` is not supported.
//
// The property Transcript is required.
type MeetingUpdateParamsRelationships struct {
	Transcript MeetingUpdateParamsRelationshipsTranscript `json:"$transcript,omitzero" api:"required"`
	paramObj
}

func (r MeetingUpdateParamsRelationships) MarshalJSON() (data []byte, err error) {
	type shadow MeetingUpdateParamsRelationships
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MeetingUpdateParamsRelationships) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Replace is required.
type MeetingUpdateParamsRelationshipsTranscript struct {
	// The file ID to set as the meeting transcript.
	Replace string `json:"replace" api:"required"`
	paramObj
}

func (r MeetingUpdateParamsRelationshipsTranscript) MarshalJSON() (data []byte, err error) {
	type shadow MeetingUpdateParamsRelationshipsTranscript
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MeetingUpdateParamsRelationshipsTranscript) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MeetingListParams struct {
	// Maximum number of records to return. Defaults to 25, maximum 25.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Number of records to skip for pagination. Defaults to 0.
	Offset param.Opt[int64] `query:"offset,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [MeetingListParams]'s query parameters as `url.Values`.
func (r MeetingListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
