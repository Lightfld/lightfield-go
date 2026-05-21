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

// Emails represent messages synced from connected email accounts in Lightfield.
// Read responses are privacy-aware and may be redacted based on the caller.
//
// EmailService contains methods and other services that help with interacting with
// the Lightfield API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewEmailService] method instead.
type EmailService struct {
	Options []option.RequestOption
}

// NewEmailService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewEmailService(opts ...option.RequestOption) (r EmailService) {
	r = EmailService{}
	r.Options = opts
	return
}

// Retrieves a single email by its ID. Email fields are redacted based on the
// caller-specific privacy resolution, and the response includes a read-only
// `accessLevel`.
//
// **[Required scope](/using-the-api/scopes/):** `emails:read`
//
// **[Rate limit category](/using-the-api/rate-limits/):** Read
func (r *EmailService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *EmailRetrieveResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/emails/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Sends an email via the connected email account that owns the `from` address.
// Currently supports new sends only; replies and forwards are not yet supported.
//
// Supports idempotency via the `Idempotency-Key` header.
//
// **[Required scope](/using-the-api/scopes/):** `emails:create`
//
// **[Rate limit category](/using-the-api/rate-limits/):** Write
func (r *EmailService) Send(ctx context.Context, body EmailSendParams, opts ...option.RequestOption) (res *EmailSendResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/emails/send"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

type EmailRetrieveResponse struct {
	// Unique identifier for the entity.
	ID string `json:"id" api:"required"`
	// The caller's resolved access level for this email.
	//
	// Any of "FULL", "METADATA".
	AccessLevel EmailRetrieveResponseAccessLevel `json:"accessLevel" api:"required"`
	// ISO 8601 timestamp of when the entity was created.
	CreatedAt string `json:"createdAt" api:"required"`
	// Map of field names to their typed values. System fields are prefixed with `$`
	// (e.g. `$name`, `$email`); custom attributes use their bare slug.
	Fields map[string]EmailRetrieveResponseField `json:"fields" api:"required"`
	// URL to view the entity in the Lightfield web app, or null.
	HTTPLink string `json:"httpLink" api:"required"`
	// Always `email`.
	//
	// Any of "email".
	ObjectType EmailRetrieveResponseObjectType `json:"objectType" api:"required"`
	// Map of relationship names to their associated entities. System relationships are
	// prefixed with `$` (e.g. `$owner`, `$contact`).
	Relationships map[string]EmailRetrieveResponseRelationship `json:"relationships" api:"required"`
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
func (r EmailRetrieveResponse) RawJSON() string { return r.JSON.raw }
func (r *EmailRetrieveResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The caller's resolved access level for this email.
type EmailRetrieveResponseAccessLevel string

const (
	EmailRetrieveResponseAccessLevelFull     EmailRetrieveResponseAccessLevel = "FULL"
	EmailRetrieveResponseAccessLevelMetadata EmailRetrieveResponseAccessLevel = "METADATA"
)

type EmailRetrieveResponseField struct {
	// The field value, or null if unset.
	Value EmailRetrieveResponseFieldValueUnion `json:"value" api:"required"`
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
func (r EmailRetrieveResponseField) RawJSON() string { return r.JSON.raw }
func (r *EmailRetrieveResponseField) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// EmailRetrieveResponseFieldValueUnion contains all possible properties and values
// from [string], [float64], [bool], [[]string],
// [EmailRetrieveResponseFieldValueAddress],
// [EmailRetrieveResponseFieldValueFullName].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfStringArray]
type EmailRetrieveResponseFieldValueUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]string] instead of an object.
	OfStringArray []string `json:",inline"`
	// This field is from variant [EmailRetrieveResponseFieldValueAddress].
	City string `json:"city"`
	// This field is from variant [EmailRetrieveResponseFieldValueAddress].
	Country string `json:"country"`
	// This field is from variant [EmailRetrieveResponseFieldValueAddress].
	Latitude float64 `json:"latitude"`
	// This field is from variant [EmailRetrieveResponseFieldValueAddress].
	Longitude float64 `json:"longitude"`
	// This field is from variant [EmailRetrieveResponseFieldValueAddress].
	PostalCode string `json:"postalCode"`
	// This field is from variant [EmailRetrieveResponseFieldValueAddress].
	State string `json:"state"`
	// This field is from variant [EmailRetrieveResponseFieldValueAddress].
	Street string `json:"street"`
	// This field is from variant [EmailRetrieveResponseFieldValueAddress].
	Street2 string `json:"street2"`
	// This field is from variant [EmailRetrieveResponseFieldValueFullName].
	FirstName string `json:"firstName"`
	// This field is from variant [EmailRetrieveResponseFieldValueFullName].
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

func (u EmailRetrieveResponseFieldValueUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u EmailRetrieveResponseFieldValueUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u EmailRetrieveResponseFieldValueUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u EmailRetrieveResponseFieldValueUnion) AsStringArray() (v []string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u EmailRetrieveResponseFieldValueUnion) AsAddress() (v EmailRetrieveResponseFieldValueAddress) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u EmailRetrieveResponseFieldValueUnion) AsFullName() (v EmailRetrieveResponseFieldValueFullName) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u EmailRetrieveResponseFieldValueUnion) RawJSON() string { return u.JSON.raw }

func (r *EmailRetrieveResponseFieldValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmailRetrieveResponseFieldValueAddress struct {
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
func (r EmailRetrieveResponseFieldValueAddress) RawJSON() string { return r.JSON.raw }
func (r *EmailRetrieveResponseFieldValueAddress) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmailRetrieveResponseFieldValueFullName struct {
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
func (r EmailRetrieveResponseFieldValueFullName) RawJSON() string { return r.JSON.raw }
func (r *EmailRetrieveResponseFieldValueFullName) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Always `email`.
type EmailRetrieveResponseObjectType string

const (
	EmailRetrieveResponseObjectTypeEmail EmailRetrieveResponseObjectType = "email"
)

type EmailRetrieveResponseRelationship struct {
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
func (r EmailRetrieveResponseRelationship) RawJSON() string { return r.JSON.raw }
func (r *EmailRetrieveResponseRelationship) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmailSendResponse struct {
	// ISO 8601 timestamp of when the send completed.
	SentAt string `json:"sentAt" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		SentAt      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmailSendResponse) RawJSON() string { return r.JSON.raw }
func (r *EmailSendResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmailSendParams struct {
	Body EmailSendParamsBody `json:"body,omitzero" api:"required"`
	// Bare email address (no display name). Must match a connected email account owned
	// by the API key user. Compared case-insensitively. Used as the From header when
	// sending.
	From string `json:"from" api:"required"`
	// Email subject. Cannot be empty.
	Subject string `json:"subject" api:"required"`
	// Recipient email addresses (bare, no display names). At least 1, at most 500.
	To []string `json:"to,omitzero" api:"required"`
	// Optional list of file IDs (uploaded via the Files API) to attach to the email.
	// Maximum 5 attachments per email. Each attachment must be ≤ 3MB and the total
	// across all attachments must be ≤ 25MB.
	Attachments []string `json:"attachments,omitzero"`
	// Bcc recipients (same shape as `to`).
	Bcc []string `json:"bcc,omitzero"`
	// Cc recipients (same shape as `to`).
	Cc []string `json:"cc,omitzero"`
	paramObj
}

func (r EmailSendParams) MarshalJSON() (data []byte, err error) {
	type shadow EmailSendParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EmailSendParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Content is required.
type EmailSendParamsBody struct {
	// Email body content.
	Content string `json:"content" api:"required"`
	// Defaults to `HTML`.
	//
	// Any of "HTML", "TEXT".
	ContentType string `json:"contentType,omitzero"`
	paramObj
}

func (r EmailSendParamsBody) MarshalJSON() (data []byte, err error) {
	type shadow EmailSendParamsBody
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EmailSendParamsBody) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[EmailSendParamsBody](
		"contentType", "HTML", "TEXT",
	)
}
