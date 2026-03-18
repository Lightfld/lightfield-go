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
	shimjson "github.com/Lightfld/lightfield-go/internal/encoding/json"
	"github.com/Lightfld/lightfield-go/internal/requestconfig"
	"github.com/Lightfld/lightfield-go/option"
	"github.com/Lightfld/lightfield-go/packages/param"
	"github.com/Lightfld/lightfield-go/packages/respjson"
)

// WebhookService contains methods and other services that help with interacting
// with the Lightfield API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWebhookService] method instead.
type WebhookService struct {
	Options []option.RequestOption
}

// NewWebhookService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewWebhookService(opts ...option.RequestOption) (r WebhookService) {
	r = WebhookService{}
	r.Options = opts
	return
}

func (r *WebhookService) Status(ctx context.Context, runID string, opts ...option.RequestOption) (res *WebhookStatusResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if runID == "" {
		err = errors.New("missing required runId parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/v1/workflowRun/%s/status", url.PathEscape(runID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

func (r *WebhookService) TriggerGet(ctx context.Context, webhookID string, opts ...option.RequestOption) (res *WebhookTriggerGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if webhookID == "" {
		err = errors.New("missing required webhookId parameter")
		return nil, err
	}
	path := fmt.Sprintf("webhooks/%s", url.PathEscape(webhookID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

func (r *WebhookService) TriggerPost(ctx context.Context, webhookID string, body WebhookTriggerPostParams, opts ...option.RequestOption) (res *WebhookTriggerPostResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if webhookID == "" {
		err = errors.New("missing required webhookId parameter")
		return nil, err
	}
	path := fmt.Sprintf("webhooks/%s", url.PathEscape(webhookID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

type WebhookStatusResponse struct {
	Status string `json:"status" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebhookStatusResponse) RawJSON() string { return r.JSON.raw }
func (r *WebhookStatusResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebhookTriggerGetResponse struct {
	ID     string `json:"id" api:"required"`
	RunID  string `json:"runId" api:"required"`
	Status string `json:"status" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		RunID       respjson.Field
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebhookTriggerGetResponse) RawJSON() string { return r.JSON.raw }
func (r *WebhookTriggerGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebhookTriggerPostResponse struct {
	ID     string `json:"id" api:"required"`
	RunID  string `json:"runId" api:"required"`
	Status string `json:"status" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		RunID       respjson.Field
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebhookTriggerPostResponse) RawJSON() string { return r.JSON.raw }
func (r *WebhookTriggerPostResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebhookTriggerPostParams struct {
	Body map[string]WebhookTriggerPostParamsBodyUnion
	paramObj
}

func (r WebhookTriggerPostParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.Body)
}
func (r *WebhookTriggerPostParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.Body)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type WebhookTriggerPostParamsBodyUnion struct {
	OfString                        param.Opt[string]                                   `json:",omitzero,inline"`
	OfFloat                         param.Opt[float64]                                  `json:",omitzero,inline"`
	OfBool                          param.Opt[bool]                                     `json:",omitzero,inline"`
	OfWebhookTriggerPostsBodyArray  []WebhookTriggerPostParamsBodyArrayItemUnion        `json:",omitzero,inline"`
	OfWebhookTriggerPostsBodyMapMap map[string]WebhookTriggerPostParamsBodyMapItemUnion `json:",omitzero,inline"`
	paramUnion
}

func (u WebhookTriggerPostParamsBodyUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString,
		u.OfFloat,
		u.OfBool,
		u.OfWebhookTriggerPostsBodyArray,
		u.OfWebhookTriggerPostsBodyMapMap)
}
func (u *WebhookTriggerPostParamsBodyUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type WebhookTriggerPostParamsBodyArrayItemUnion struct {
	OfString   param.Opt[string]  `json:",omitzero,inline"`
	OfFloat    param.Opt[float64] `json:",omitzero,inline"`
	OfBool     param.Opt[bool]    `json:",omitzero,inline"`
	OfAnyArray []any              `json:",omitzero,inline"`
	OfAnyMap   map[string]any     `json:",omitzero,inline"`
	paramUnion
}

func (u WebhookTriggerPostParamsBodyArrayItemUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString,
		u.OfFloat,
		u.OfBool,
		u.OfAnyArray,
		u.OfAnyMap)
}
func (u *WebhookTriggerPostParamsBodyArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type WebhookTriggerPostParamsBodyMapItemUnion struct {
	OfString   param.Opt[string]  `json:",omitzero,inline"`
	OfFloat    param.Opt[float64] `json:",omitzero,inline"`
	OfBool     param.Opt[bool]    `json:",omitzero,inline"`
	OfAnyArray []any              `json:",omitzero,inline"`
	OfAnyMap   map[string]any     `json:",omitzero,inline"`
	paramUnion
}

func (u WebhookTriggerPostParamsBodyMapItemUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString,
		u.OfFloat,
		u.OfBool,
		u.OfAnyArray,
		u.OfAnyMap)
}
func (u *WebhookTriggerPostParamsBodyMapItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}
