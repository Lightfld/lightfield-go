// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package githubcomlightfldlightfieldgo

import (
	"context"
	"net/http"
	"slices"

	"github.com/Lightfld/lightfield-go/internal/apijson"
	"github.com/Lightfld/lightfield-go/internal/requestconfig"
	"github.com/Lightfld/lightfield-go/option"
	"github.com/Lightfld/lightfield-go/packages/respjson"
)

// Auth helper to validate the current token and describe the access it grants.
// This endpoint is currently only available for API keys.
//
// AuthService contains methods and other services that help with interacting with
// the Lightfield API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAuthService] method instead.
type AuthService struct {
	Options []option.RequestOption
}

// NewAuthService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewAuthService(opts ...option.RequestOption) (r AuthService) {
	r = AuthService{}
	r.Options = opts
	return
}

// Returns metadata for the current API key, including the subject type and granted
// public scopes. Use this endpoint to confirm a key is active before making scoped
// API requests.
//
// **[Required scope](/using-the-api/scopes/):** None
//
// **[Rate limit category](/using-the-api/rate-limits/):** Read
func (r *AuthService) Validate(ctx context.Context, opts ...option.RequestOption) (res *AuthValidateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/auth/validate"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

type AuthValidateResponse struct {
	// Whether the current API key is valid. Always `true` on successful responses.
	Active bool `json:"active" api:"required"`
	// Granted public scopes for the current API key. Empty when the key has full
	// access.
	Scopes []string `json:"scopes" api:"required"`
	// Whether the API key belongs to a `user` or `workspace`.
	//
	// Any of "user", "workspace".
	SubjectType AuthValidateResponseSubjectType `json:"subjectType" api:"required"`
	// Credential family, always `api_key`.
	//
	// Any of "api_key".
	TokenType AuthValidateResponseTokenType `json:"tokenType" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Active      respjson.Field
		Scopes      respjson.Field
		SubjectType respjson.Field
		TokenType   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AuthValidateResponse) RawJSON() string { return r.JSON.raw }
func (r *AuthValidateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API key belongs to a `user` or `workspace`.
type AuthValidateResponseSubjectType string

const (
	AuthValidateResponseSubjectTypeUser      AuthValidateResponseSubjectType = "user"
	AuthValidateResponseSubjectTypeWorkspace AuthValidateResponseSubjectType = "workspace"
)

// Credential family, always `api_key`.
type AuthValidateResponseTokenType string

const (
	AuthValidateResponseTokenTypeAPIKey AuthValidateResponseTokenType = "api_key"
)
