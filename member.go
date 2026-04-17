// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package githubcomlightfldlightfieldgo

import (
	"context"
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

// Members represent users in your Lightfield workspace. Members can own accounts
// and opportunities, and are referenced in relationships like `$owner` and
// `$createdBy`.
//
// MemberService contains methods and other services that help with interacting
// with the Lightfield API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewMemberService] method instead.
type MemberService struct {
	Options []option.RequestOption
}

// NewMemberService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewMemberService(opts ...option.RequestOption) (r MemberService) {
	r = MemberService{}
	r.Options = opts
	return
}

// Retrieves a single member by their ID.
//
// **[Required scope](/using-the-api/scopes/):** `members:read`
//
// **[Rate limit category](/using-the-api/rate-limits/):** Read
func (r *MemberService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *MemberRetrieveResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/members/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Returns a paginated list of members in your workspace. Use `offset` and `limit`
// to paginate through results. See
// <u>[List endpoints](/using-the-api/list-endpoints/)</u> for more information
// about pagination.
//
// **[Required scope](/using-the-api/scopes/):** `members:read`
//
// **[Rate limit category](/using-the-api/rate-limits/):** Search
func (r *MemberService) List(ctx context.Context, query MemberListParams, opts ...option.RequestOption) (res *MemberListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/members"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

type MemberListResponse struct {
	// Array of member objects for the current page.
	Data []MemberListResponseData `json:"data" api:"required"`
	// The object type, always `"list"`.
	Object string `json:"object" api:"required"`
	// Total number of members in the workspace.
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
func (r MemberListResponse) RawJSON() string { return r.JSON.raw }
func (r *MemberListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MemberListResponseData struct {
	// Unique identifier for the member.
	ID string `json:"id" api:"required"`
	// ISO 8601 timestamp of when the member was created.
	CreatedAt string `json:"createdAt" api:"required"`
	// Map of field names to their typed values.
	Fields MemberListResponseDataFields `json:"fields" api:"required"`
	// URL to view the member in the Lightfield web app, or null.
	HTTPLink string `json:"httpLink" api:"required"`
	// Members do not expose writable or readable relationships in this API.
	Relationships any `json:"relationships" api:"required"`
	// ISO 8601 timestamp of when the member was last updated, or null.
	UpdatedAt string `json:"updatedAt" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID            respjson.Field
		CreatedAt     respjson.Field
		Fields        respjson.Field
		HTTPLink      respjson.Field
		Relationships respjson.Field
		UpdatedAt     respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MemberListResponseData) RawJSON() string { return r.JSON.raw }
func (r *MemberListResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Map of field names to their typed values.
type MemberListResponseDataFields struct {
	// The member's email address.
	Email MemberListResponseDataFieldsEmail `json:"$email" api:"required"`
	// The member's full name.
	Name MemberListResponseDataFieldsName `json:"$name" api:"required"`
	// URL of the member's profile image, or null if unset.
	ProfileImage MemberListResponseDataFieldsProfileImage `json:"$profileImage" api:"required"`
	// The member's workspace role.
	Role MemberListResponseDataFieldsRole `json:"$role" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Email        respjson.Field
		Name         respjson.Field
		ProfileImage respjson.Field
		Role         respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MemberListResponseDataFields) RawJSON() string { return r.JSON.raw }
func (r *MemberListResponseDataFields) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The member's email address.
type MemberListResponseDataFieldsEmail struct {
	// The field value.
	Value string `json:"value" api:"required"`
	// The data type of the field value.
	//
	// Any of "EMAIL".
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
func (r MemberListResponseDataFieldsEmail) RawJSON() string { return r.JSON.raw }
func (r *MemberListResponseDataFieldsEmail) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The member's full name.
type MemberListResponseDataFieldsName struct {
	Value MemberListResponseDataFieldsNameValue `json:"value" api:"required"`
	// The data type of the field value.
	//
	// Any of "FULL_NAME".
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
func (r MemberListResponseDataFieldsName) RawJSON() string { return r.JSON.raw }
func (r *MemberListResponseDataFieldsName) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MemberListResponseDataFieldsNameValue struct {
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
func (r MemberListResponseDataFieldsNameValue) RawJSON() string { return r.JSON.raw }
func (r *MemberListResponseDataFieldsNameValue) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URL of the member's profile image, or null if unset.
type MemberListResponseDataFieldsProfileImage struct {
	// The field value, or null if unset.
	Value string `json:"value" api:"required"`
	// The data type of the field value.
	//
	// Any of "URL".
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
func (r MemberListResponseDataFieldsProfileImage) RawJSON() string { return r.JSON.raw }
func (r *MemberListResponseDataFieldsProfileImage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The member's workspace role.
type MemberListResponseDataFieldsRole struct {
	// The field value.
	Value string `json:"value" api:"required"`
	// The data type of the field value.
	//
	// Any of "TEXT".
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
func (r MemberListResponseDataFieldsRole) RawJSON() string { return r.JSON.raw }
func (r *MemberListResponseDataFieldsRole) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MemberRetrieveResponse struct {
	// Unique identifier for the member.
	ID string `json:"id" api:"required"`
	// ISO 8601 timestamp of when the member was created.
	CreatedAt string `json:"createdAt" api:"required"`
	// Map of field names to their typed values.
	Fields MemberRetrieveResponseFields `json:"fields" api:"required"`
	// URL to view the member in the Lightfield web app, or null.
	HTTPLink string `json:"httpLink" api:"required"`
	// Members do not expose writable or readable relationships in this API.
	Relationships any `json:"relationships" api:"required"`
	// ISO 8601 timestamp of when the member was last updated, or null.
	UpdatedAt string `json:"updatedAt" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID            respjson.Field
		CreatedAt     respjson.Field
		Fields        respjson.Field
		HTTPLink      respjson.Field
		Relationships respjson.Field
		UpdatedAt     respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MemberRetrieveResponse) RawJSON() string { return r.JSON.raw }
func (r *MemberRetrieveResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Map of field names to their typed values.
type MemberRetrieveResponseFields struct {
	// The member's email address.
	Email MemberRetrieveResponseFieldsEmail `json:"$email" api:"required"`
	// The member's full name.
	Name MemberRetrieveResponseFieldsName `json:"$name" api:"required"`
	// URL of the member's profile image, or null if unset.
	ProfileImage MemberRetrieveResponseFieldsProfileImage `json:"$profileImage" api:"required"`
	// The member's workspace role.
	Role MemberRetrieveResponseFieldsRole `json:"$role" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Email        respjson.Field
		Name         respjson.Field
		ProfileImage respjson.Field
		Role         respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MemberRetrieveResponseFields) RawJSON() string { return r.JSON.raw }
func (r *MemberRetrieveResponseFields) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The member's email address.
type MemberRetrieveResponseFieldsEmail struct {
	// The field value.
	Value string `json:"value" api:"required"`
	// The data type of the field value.
	//
	// Any of "EMAIL".
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
func (r MemberRetrieveResponseFieldsEmail) RawJSON() string { return r.JSON.raw }
func (r *MemberRetrieveResponseFieldsEmail) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The member's full name.
type MemberRetrieveResponseFieldsName struct {
	Value MemberRetrieveResponseFieldsNameValue `json:"value" api:"required"`
	// The data type of the field value.
	//
	// Any of "FULL_NAME".
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
func (r MemberRetrieveResponseFieldsName) RawJSON() string { return r.JSON.raw }
func (r *MemberRetrieveResponseFieldsName) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MemberRetrieveResponseFieldsNameValue struct {
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
func (r MemberRetrieveResponseFieldsNameValue) RawJSON() string { return r.JSON.raw }
func (r *MemberRetrieveResponseFieldsNameValue) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URL of the member's profile image, or null if unset.
type MemberRetrieveResponseFieldsProfileImage struct {
	// The field value, or null if unset.
	Value string `json:"value" api:"required"`
	// The data type of the field value.
	//
	// Any of "URL".
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
func (r MemberRetrieveResponseFieldsProfileImage) RawJSON() string { return r.JSON.raw }
func (r *MemberRetrieveResponseFieldsProfileImage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The member's workspace role.
type MemberRetrieveResponseFieldsRole struct {
	// The field value.
	Value string `json:"value" api:"required"`
	// The data type of the field value.
	//
	// Any of "TEXT".
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
func (r MemberRetrieveResponseFieldsRole) RawJSON() string { return r.JSON.raw }
func (r *MemberRetrieveResponseFieldsRole) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MemberListParams struct {
	// Maximum number of records to return. Defaults to 25, maximum 25.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Number of records to skip for pagination. Defaults to 0.
	Offset param.Opt[int64] `query:"offset,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [MemberListParams]'s query parameters as `url.Values`.
func (r MemberListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
