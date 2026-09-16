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
	"github.com/Lightfld/lightfield-go/internal/requestconfig"
	"github.com/Lightfld/lightfield-go/option"
	"github.com/Lightfld/lightfield-go/packages/param"
	"github.com/Lightfld/lightfield-go/packages/respjson"
)

// Enrichment fills in missing information on a record from Lightfield's data
// providers.
//
// EnrichmentRunService contains methods and other services that help with
// interacting with the Lightfield API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewEnrichmentRunService] method instead.
type EnrichmentRunService struct {
	Options []option.RequestOption
}

// NewEnrichmentRunService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewEnrichmentRunService(opts ...option.RequestOption) (r EnrichmentRunService) {
	r = EnrichmentRunService{}
	r.Options = opts
	return
}

// Looks up missing information for a record from Lightfield's enrichment providers
// and writes it back to it. Accepts `contacts` and `accounts`.
//
// Enrichment runs in the background: this returns as soon as the run is created.
// Poll `GET /v1/enrichmentRun/{runId}` for its status. A record can only have one
// enrichment run at a time — if one is already in flight, this returns that run
// rather than starting a second.
//
// Which fields are enriched, and whether a provider answer overwrites an existing
// value or is raised as a suggestion, come from the workspace's enrichment
// settings. Naming fields explicitly overrides which fields are enriched, but not
// the write policy.
//
// **[Required scopes](/using-the-api/scopes/):** `contacts:update` to enrich a
// contact, `accounts:update` to enrich an account
//
// **[Rate limit category](/using-the-api/rate-limits/):** Write
func (r *EnrichmentRunService) Enrich(ctx context.Context, entityID string, params EnrichmentRunEnrichParams, opts ...option.RequestOption) (res *EnrichmentRunEnrichResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if entityID == "" {
		err = errors.New("missing required entityId parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/enrichmentRun/%v/%s", params.EntitySlug, url.PathEscape(entityID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Returns the status of an enrichment run by its ID.
//
// **[Required scopes](/using-the-api/scopes/):** `contacts:read` for a contact
// run, `accounts:read` for an account run
//
// **[Rate limit category](/using-the-api/rate-limits/):** Read
func (r *EnrichmentRunService) GetEnrichmentRun(ctx context.Context, runID string, opts ...option.RequestOption) (res *EnrichmentRunGetEnrichmentRunResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if runID == "" {
		err = errors.New("missing required runId parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/enrichmentRun/%s", url.PathEscape(runID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

type EnrichmentRunEnrichResponse struct {
	// The run, or null when the request was skipped.
	EnrichmentRun EnrichmentRunEnrichResponseEnrichmentRun `json:"enrichmentRun" api:"required"`
	// Why the request was skipped: `no_targets` when the workspace enriches nothing
	// this record is missing, `no_operations` when no provider can be scheduled for
	// the targeted fields — because none looks them up on request, or because the
	// record lacks the inputs (such as an email or domain) a lookup would need. Null
	// otherwise.
	//
	// Any of "no_targets", "no_operations".
	Reason EnrichmentRunEnrichResponseReason `json:"reason" api:"required"`
	// `started` when this request began a run, `already_running` when one was already
	// enriching the record, `skipped` when there was nothing to enrich.
	//
	// Any of "started", "already_running", "skipped".
	Status EnrichmentRunEnrichResponseStatus `json:"status" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		EnrichmentRun respjson.Field
		Reason        respjson.Field
		Status        respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EnrichmentRunEnrichResponse) RawJSON() string { return r.JSON.raw }
func (r *EnrichmentRunEnrichResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The run, or null when the request was skipped.
type EnrichmentRunEnrichResponseEnrichmentRun struct {
	// The enrichment run ID.
	ID string `json:"id" api:"required"`
	// When the run finished, or null while it is still open.
	CompletedAt string `json:"completedAt" api:"required"`
	CreatedAt   string `json:"createdAt" api:"required"`
	// The ID of the record being enriched.
	EntityID string `json:"entityId" api:"required"`
	// The type of record being enriched.
	//
	// Any of "contact", "account".
	EntityType string `json:"entityType" api:"required"`
	// When the run began executing, or null while it is queued.
	StartedAt string `json:"startedAt" api:"required"`
	// Where the run is: `queued` and `running` are in flight; `completed`, `failed`
	// and `timed_out` are final.
	//
	// Any of "queued", "running", "completed", "failed", "timed_out".
	Status string `json:"status" api:"required"`
	// The fields this run set out to fill.
	TargetFields []string `json:"targetFields" api:"required"`
	// What started the run: `creation` when the record was created, `manual` when
	// requested through the API.
	//
	// Any of "creation", "manual".
	Trigger   string `json:"trigger" api:"required"`
	UpdatedAt string `json:"updatedAt" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		CompletedAt  respjson.Field
		CreatedAt    respjson.Field
		EntityID     respjson.Field
		EntityType   respjson.Field
		StartedAt    respjson.Field
		Status       respjson.Field
		TargetFields respjson.Field
		Trigger      respjson.Field
		UpdatedAt    respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EnrichmentRunEnrichResponseEnrichmentRun) RawJSON() string { return r.JSON.raw }
func (r *EnrichmentRunEnrichResponseEnrichmentRun) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Why the request was skipped: `no_targets` when the workspace enriches nothing
// this record is missing, `no_operations` when no provider can be scheduled for
// the targeted fields — because none looks them up on request, or because the
// record lacks the inputs (such as an email or domain) a lookup would need. Null
// otherwise.
type EnrichmentRunEnrichResponseReason string

const (
	EnrichmentRunEnrichResponseReasonNoTargets    EnrichmentRunEnrichResponseReason = "no_targets"
	EnrichmentRunEnrichResponseReasonNoOperations EnrichmentRunEnrichResponseReason = "no_operations"
)

// `started` when this request began a run, `already_running` when one was already
// enriching the record, `skipped` when there was nothing to enrich.
type EnrichmentRunEnrichResponseStatus string

const (
	EnrichmentRunEnrichResponseStatusStarted        EnrichmentRunEnrichResponseStatus = "started"
	EnrichmentRunEnrichResponseStatusAlreadyRunning EnrichmentRunEnrichResponseStatus = "already_running"
	EnrichmentRunEnrichResponseStatusSkipped        EnrichmentRunEnrichResponseStatus = "skipped"
)

type EnrichmentRunGetEnrichmentRunResponse struct {
	// The enrichment run ID.
	ID string `json:"id" api:"required"`
	// When the run finished, or null while it is still open.
	CompletedAt string `json:"completedAt" api:"required"`
	CreatedAt   string `json:"createdAt" api:"required"`
	// The ID of the record being enriched.
	EntityID string `json:"entityId" api:"required"`
	// The type of record being enriched.
	//
	// Any of "contact", "account".
	EntityType EnrichmentRunGetEnrichmentRunResponseEntityType `json:"entityType" api:"required"`
	// When the run began executing, or null while it is queued.
	StartedAt string `json:"startedAt" api:"required"`
	// Where the run is: `queued` and `running` are in flight; `completed`, `failed`
	// and `timed_out` are final.
	//
	// Any of "queued", "running", "completed", "failed", "timed_out".
	Status EnrichmentRunGetEnrichmentRunResponseStatus `json:"status" api:"required"`
	// The fields this run set out to fill.
	TargetFields []string `json:"targetFields" api:"required"`
	// What started the run: `creation` when the record was created, `manual` when
	// requested through the API.
	//
	// Any of "creation", "manual".
	Trigger   EnrichmentRunGetEnrichmentRunResponseTrigger `json:"trigger" api:"required"`
	UpdatedAt string                                       `json:"updatedAt" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		CompletedAt  respjson.Field
		CreatedAt    respjson.Field
		EntityID     respjson.Field
		EntityType   respjson.Field
		StartedAt    respjson.Field
		Status       respjson.Field
		TargetFields respjson.Field
		Trigger      respjson.Field
		UpdatedAt    respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EnrichmentRunGetEnrichmentRunResponse) RawJSON() string { return r.JSON.raw }
func (r *EnrichmentRunGetEnrichmentRunResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The type of record being enriched.
type EnrichmentRunGetEnrichmentRunResponseEntityType string

const (
	EnrichmentRunGetEnrichmentRunResponseEntityTypeContact EnrichmentRunGetEnrichmentRunResponseEntityType = "contact"
	EnrichmentRunGetEnrichmentRunResponseEntityTypeAccount EnrichmentRunGetEnrichmentRunResponseEntityType = "account"
)

// Where the run is: `queued` and `running` are in flight; `completed`, `failed`
// and `timed_out` are final.
type EnrichmentRunGetEnrichmentRunResponseStatus string

const (
	EnrichmentRunGetEnrichmentRunResponseStatusQueued    EnrichmentRunGetEnrichmentRunResponseStatus = "queued"
	EnrichmentRunGetEnrichmentRunResponseStatusRunning   EnrichmentRunGetEnrichmentRunResponseStatus = "running"
	EnrichmentRunGetEnrichmentRunResponseStatusCompleted EnrichmentRunGetEnrichmentRunResponseStatus = "completed"
	EnrichmentRunGetEnrichmentRunResponseStatusFailed    EnrichmentRunGetEnrichmentRunResponseStatus = "failed"
	EnrichmentRunGetEnrichmentRunResponseStatusTimedOut  EnrichmentRunGetEnrichmentRunResponseStatus = "timed_out"
)

// What started the run: `creation` when the record was created, `manual` when
// requested through the API.
type EnrichmentRunGetEnrichmentRunResponseTrigger string

const (
	EnrichmentRunGetEnrichmentRunResponseTriggerCreation EnrichmentRunGetEnrichmentRunResponseTrigger = "creation"
	EnrichmentRunGetEnrichmentRunResponseTriggerManual   EnrichmentRunGetEnrichmentRunResponseTrigger = "manual"
)

type EnrichmentRunEnrichParams struct {
	// The type of record to enrich.
	//
	// Any of "contacts", "accounts".
	EntitySlug EnrichmentRunEnrichParamsEntitySlug `path:"entitySlug,omitzero" api:"required" json:"-"`
	// Fields to enrich, e.g. `["email", "title"]`. Named fields are refreshed even
	// when they already hold a value, so this is how a stale value is replaced. Omit
	// to fill only the fields the record is missing. Naming a field the workspace's
	// settings turn off enriches it anyway, and its answer is raised as a suggestion
	// rather than written over a value the record already holds. `profilePhotoUrl` has
	// no suggestion form, so it stays rejected while it is turned off. A named field
	// is not used to look the record up — the run derives it fresh from the record's
	// other identifiers — so naming every identifier (a contact's name and email
	// together) leaves nothing to search by and is skipped with `no_operations`.
	// `phone` is filled only when another field's lookup happens to return it, so
	// requesting it on its own is skipped the same way.
	Fields []string `json:"fields,omitzero"`
	paramObj
}

func (r EnrichmentRunEnrichParams) MarshalJSON() (data []byte, err error) {
	type shadow EnrichmentRunEnrichParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EnrichmentRunEnrichParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The type of record to enrich.
type EnrichmentRunEnrichParamsEntitySlug string

const (
	EnrichmentRunEnrichParamsEntitySlugContacts EnrichmentRunEnrichParamsEntitySlug = "contacts"
	EnrichmentRunEnrichParamsEntitySlugAccounts EnrichmentRunEnrichParamsEntitySlug = "accounts"
)
