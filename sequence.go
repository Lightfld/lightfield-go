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
	shimjson "github.com/Lightfld/lightfield-go/internal/encoding/json"
	"github.com/Lightfld/lightfield-go/internal/requestconfig"
	"github.com/Lightfld/lightfield-go/option"
	"github.com/Lightfld/lightfield-go/packages/param"
	"github.com/Lightfld/lightfield-go/packages/respjson"
)

// > [!WARNING] The Sequences API is in Preview. Endpoints, parameters, and
// > response schemas may change, including breaking changes, as we incorporate
// > feedback.
//
// Sequences are multi-step outreach cadences that send email and LinkedIn messages
// to enrolled contacts.
//
// SequenceService contains methods and other services that help with interacting
// with the Lightfield API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSequenceService] method instead.
type SequenceService struct {
	Options []option.RequestOption
}

// NewSequenceService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewSequenceService(opts ...option.RequestOption) (r SequenceService) {
	r = SequenceService{}
	r.Options = opts
	return
}

// Creates an unpublished sequence with its first draft and returns the draft.
// Unset send settings default to 9:00–17:00, Monday to Friday, UTC, with no daily
// limits. The sequence sends nothing until its draft is published and it is
// activated.
//
// Supports idempotency via the `Idempotency-Key` header.
//
// **[Required scope](/using-the-api/scopes/):** `sequences:update`
//
// **[Rate limit category](/using-the-api/rate-limits/):** Write
func (r *SequenceService) New(ctx context.Context, body SequenceNewParams, opts ...option.RequestOption) (res *SequenceCreateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/sequences"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Retrieves a sequence by its ID. Published sequences include their steps and send
// settings; a sequence that has never been published has a null definition and no
// channels. Pending edits live in its draft.
//
// **[Required scope](/using-the-api/scopes/):** `sequences:read`
//
// **[Rate limit category](/using-the-api/rate-limits/):** Read
func (r *SequenceService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *SequenceRetrieveResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/sequences/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Returns a paginated list of sequences, including ones never published, ordered
// by `createdAt` descending. A sequence that has never been published has a null
// definition and no channels. Use `offset` and `limit` to paginate, and bracketed
// filters for status, channels, and creation or activation dates.
//
// Supported filters: `status[equal]` (DRAFT, ACTIVE, DEACTIVATED),
// `channels[contains]` (EMAIL, LINKEDIN), and `createdAt`, `firstActivatedAt`, or
// `lastActivatedAt` with `[greaterThanOrEqual]` or `[lessThanOrEqual]`. Date
// values must be ISO 8601 timestamps with a timezone. Filters are combined with
// AND.
//
// **[Required scope](/using-the-api/scopes/):** `sequences:read`
//
// **[Rate limit category](/using-the-api/rate-limits/):** Search
func (r *SequenceService) List(ctx context.Context, query SequenceListParams, opts ...option.RequestOption) (res *SequenceListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/sequences"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Starts sending a published sequence to its enrolled contacts. Requires a
// published version and an available sender for each channel it uses; add senders
// through the draft. An already active sequence is returned unchanged. Edit steps
// and settings through the draft instead. The sequence's status changes
// immediately; its enrollments' statuses follow asynchronously.
//
// **[Required scope](/using-the-api/scopes/):** `sequences:update`
//
// **[Rate limit category](/using-the-api/rate-limits/):** Write
func (r *SequenceService) Activate(ctx context.Context, id string, body SequenceActivateParams, opts ...option.RequestOption) (res *SequenceActivateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/sequences/%s/activate", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Pauses sending for a published sequence; enrollments stay in place and resume on
// activation. An already paused sequence is returned unchanged. The sequence's
// status changes immediately; its enrollments' statuses follow asynchronously.
//
// **[Required scope](/using-the-api/scopes/):** `sequences:update`
//
// **[Rate limit category](/using-the-api/rate-limits/):** Write
func (r *SequenceService) Deactivate(ctx context.Context, id string, body SequenceDeactivateParams, opts ...option.RequestOption) (res *SequenceDeactivateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/sequences/%s/deactivate", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Discards the sequence's open draft at `expectedVersion`. A published sequence
// keeps its published version. A sequence that was never published is deleted with
// its draft, and `sequenceDeleted` is true. Returns 404 `draft_not_found` when no
// draft is open, and 409 when the open draft is not the one that was read.
//
// **[Required scope](/using-the-api/scopes/):** `sequences:update`
//
// **[Rate limit category](/using-the-api/rate-limits/):** Write
func (r *SequenceService) DiscardDraft(ctx context.Context, id string, body SequenceDiscardDraftParams, opts ...option.RequestOption) (res *SequenceDiscardDraftResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/sequences/%s/draft/discard", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Returns a paginated list of live enrollments in a sequence, including completed
// and cancelled history; a sequence that has never been published has none.
// Results are ordered by `enrolledAt` descending. Use `offset` and `limit` to
// paginate, and bracketed field filters to narrow results. Test-run enrollments
// are excluded.
//
// Supported filters: `status[equal]` (ACTIVE, AWAITING_REVIEW, NEEDS_ATTENTION,
// REPLIED, COMPLETED, PAUSED, CANCELLED, FAILED), `contactId[equal]`,
// `currentStepId[equal]`, and `enrolledAt`, `lastActivityAt`, or `repliedAt` with
// `[greaterThanOrEqual]` or `[lessThanOrEqual]`. Date values must be ISO 8601
// timestamps with a timezone. Filters are combined with AND.
//
// **[Required scope](/using-the-api/scopes/):** `sequences:read`
//
// **[Rate limit category](/using-the-api/rate-limits/):** Search
func (r *SequenceService) ListEnrollments(ctx context.Context, id string, query SequenceListEnrollmentsParams, opts ...option.RequestOption) (res *SequenceListEnrollmentsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/sequences/%s/enrollments", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Lists outbound messages and each enrollment's first incoming reply across email
// and LinkedIn in a sequence; a sequence that has never been published has none.
// Only the first reply per enrollment is returned, not every reply or one per
// channel. Superseded outbound attempts are excluded. Results have no message
// content.
//
// Results are ordered by `createdAt` descending. Supported equality filters:
// `contactId`, `enrollmentId`, `sequenceStepId`, `sequenceStepType` (EMAIL,
// LINKEDIN_MESSAGE, LINKEDIN_INVITE), `channel` (EMAIL, LINKEDIN), `direction`
// (OUTBOUND, INBOUND), and `status` (AWAITING_REVIEW, SCHEDULED, QUEUED, SENT,
// SKIPPED, FAILED, CANCELLED). Use `[equal]` or omit the operator. Date filters:
// `sentAt`, `openedAt`, `readAt`, `repliedAt`, `bouncedAt`, `acceptedAt`, and
// `receivedAt` support `[greaterThan]`, `[greaterThanOrEqual]`, `[lessThan]`, and
// `[lessThanOrEqual]`, with ISO 8601 timestamps including a timezone. Filters are
// combined with AND; outbound status filters exclude incoming replies.
//
// **[Required scope](/using-the-api/scopes/):** `sequences:read`
//
// **[Rate limit category](/using-the-api/rate-limits/):** Search
func (r *SequenceService) ListMessages(ctx context.Context, id string, query SequenceListMessagesParams, opts ...option.RequestOption) (res *SequenceListMessagesResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/sequences/%s/messages", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Publishes the sequence's open draft at `expectedVersion` and returns the
// sequence. A first publish leaves the sequence `DEACTIVATED`; activate it to
// start sending. Live enrollments move to the new version. On failure returns 400
// with `issues` and changes nothing; problems are reported per stage (content,
// reviewers, senders, entitlement), so fixing one stage may reveal the next. Two
// refusals come without `issues`: 400 `permissions_not_granted` when the draft has
// step permissions not yet granted (grant them on the sequence in Lightfield), and
// 400 `no_selectable_sender` when no sender is listed for a channel the steps use,
// or the sequence is active and none of its senders for that channel is available.
// Returns 404 `draft_not_found` when no draft is open, and 409 when the open draft
// is not the one that was read.
//
// **[Required scope](/using-the-api/scopes/):** `sequences:update`
//
// **[Rate limit category](/using-the-api/rate-limits/):** Write
func (r *SequenceService) PublishDraft(ctx context.Context, id string, body SequencePublishDraftParams, opts ...option.RequestOption) (res *SequencePublishDraftResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/sequences/%s/draft/publish", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Retrieves the sequence's open draft, the one shared with editors in Lightfield.
// Returns 404 `draft_not_found` when no draft is open. Its `version` is the
// `expectedVersion` the next save, publish or discard must send.
//
// **[Required scope](/using-the-api/scopes/):** `sequences:read`
//
// **[Rate limit category](/using-the-api/rate-limits/):** Read
func (r *SequenceService) GetDraft(ctx context.Context, id string, opts ...option.RequestOption) (res *SequenceRetrieveDraftResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/sequences/%s/draft", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Retrieves one contact's enrollment in a sequence: where it is in the cadence and
// its latest activity. Test-run enrollments are not returned. Status and
// remainingSends can be null until a current indexed value is available.
//
// **[Required scope](/using-the-api/scopes/):** `sequences:read`
//
// **[Rate limit category](/using-the-api/rate-limits/):** Read
func (r *SequenceService) GetEnrollment(ctx context.Context, enrollmentID string, query SequenceGetEnrollmentParams, opts ...option.RequestOption) (res *SequenceRetrieveEnrollmentResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if query.ID == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	if enrollmentID == "" {
		err = errors.New("missing required enrollmentId parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/sequences/%s/enrollments/%s", url.PathEscape(query.ID), url.PathEscape(enrollmentID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Retrieves an outbound message or the first incoming reply across email and
// LinkedIn for its enrollment, with any resolvable active same-workspace reference
// and plain-text content. Callers with `sequences:read` can read related content
// regardless of the sender or individual email/chat privacy. Content is null when
// unavailable or its captured source was deleted; a LinkedIn invitation without a
// note returns `content: { subject: null, body: null }`.
//
// **[Required scope](/using-the-api/scopes/):** `sequences:read`
//
// **[Rate limit category](/using-the-api/rate-limits/):** Read
func (r *SequenceService) GetMessage(ctx context.Context, messageID string, query SequenceGetMessageParams, opts ...option.RequestOption) (res *SequenceRetrieveMessageResponseUnion, err error) {
	opts = slices.Concat(r.Options, opts)
	if query.ID == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	if messageID == "" {
		err = errors.New("missing required messageId parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/sequences/%s/messages/%s", url.PathEscape(query.ID), url.PathEscape(messageID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Returns lifetime engagement and current step enrollment counts for a sequence.
// Includes live enrollments, excluding test enrollments and enrollments removed
// from the sequence. A sequence that has never been published reports zero counts
// and no steps. Rates are fractions; a rate is null when its denominator is zero.
//
// **[Required scope](/using-the-api/scopes/):** `sequences:read`
//
// **[Rate limit category](/using-the-api/rate-limits/):** Read
func (r *SequenceService) GetPerformance(ctx context.Context, id string, opts ...option.RequestOption) (res *SequenceRetrievePerformanceResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/sequences/%s/performance", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Saves changes to the sequence's open draft, opening one from the published
// version if none is open. Send the open draft’s `version` as `expectedVersion`;
// omit it only when no draft is open. A stale version returns 409: retrieve the
// draft and retry. `name` and `definition` are a JSON Merge Patch: omitted fields
// are unchanged, `messaging`, `sendWindow` and `pacing` change only the fields
// sent, `null` clears a messaging field or a `dailyLimit`, and `definition.steps`
// replaces the whole step list; keep a step by sending its `id`. Anything a draft
// read returns can be sent back. A value outside its field’s documented bounds
// returns 400 naming the field in `param`. Otherwise the save is all-or-nothing: a
// 400 lists its problems under `issues`, with step, settings, entitlement,
// reviewer and sender problems reported together, and changes nothing. A version
// conflict (409), a missing permission (403) or an unknown reference (400
// `referenced_resource_missing`) refuses on its own. A `steps` save returns 409
// while the draft has edits in Lightfield that are not valid yet. `reviewers` adds
// and removes members, and `senders` adds and removes senders. Changes take effect
// when the draft is published.
//
// **[Required scope](/using-the-api/scopes/):** `sequences:update`
//
// **[Rate limit category](/using-the-api/rate-limits/):** Write
func (r *SequenceService) UpdateDraft(ctx context.Context, id string, body SequenceUpdateDraftParams, opts ...option.RequestOption) (res *SequenceUpdateDraftResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/sequences/%s/draft", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Adds and removes contacts in one atomic request. If any new enrollment is
// invalid, nothing changes. Provide add, remove, or both, with at most 25 contact
// IDs total. Duplicate IDs within a list are ignored; an ID cannot appear in both
// lists. Published, paused sequences accept enrollment and wait for activation. A
// sequence that has never been published returns 400 `sequence_not_published`.
// Already-open enrollments remain unchanged, including paused enrollments.
// Contacts whose prior enrollment completed, was cancelled, or failed are enrolled
// again. Removing a contact cancels its open enrollment and stops its remaining
// steps; the enrollment stays readable with status `CANCELLED`, and adding the
// contact again starts a new enrollment. Removing a contact with no open
// enrollment returns no_active_enrollment. Execution starts asynchronously after
// enrollment.
//
// **[Required scope](/using-the-api/scopes/):** `sequences:update`
//
// **[Rate limit category](/using-the-api/rate-limits/):** Write
func (r *SequenceService) UpdateEnrollments(ctx context.Context, id string, body SequenceUpdateEnrollmentsParams, opts ...option.RequestOption) (res *SequenceUpdateEnrollmentsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/sequences/%s/enrollments", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

type SequenceActivateResponse struct {
	// Unique identifier of the sequence.
	ID string `json:"id" api:"required"`
	// Channels the published steps send on; empty before the first publish.
	//
	// Any of "EMAIL", "LINKEDIN".
	Channels []string `json:"channels" api:"required"`
	// ISO 8601 creation timestamp.
	CreatedAt string `json:"createdAt" api:"required"`
	// Published configuration, or null before the first publish.
	Definition SequenceActivateResponseDefinition `json:"definition" api:"required"`
	// When the sequence was first activated.
	FirstActivatedAt string `json:"firstActivatedAt" api:"required"`
	// When the sequence was most recently activated.
	LastActivatedAt string `json:"lastActivatedAt" api:"required"`
	// Name of the sequence.
	Name string `json:"name" api:"required"`
	// Always `sequence`.
	//
	// Any of "sequence".
	ObjectType SequenceActivateResponseObjectType `json:"objectType" api:"required"`
	// Extra members who may approve this sequence’s `FOR_REVIEW` messages before they
	// send. The sending member and the sequence’s creator always review, whether or
	// not they are listed. A change applies to sends once the draft is published. The
	// published list; empty before the first publish. Lists active members only.
	Reviewers []SequenceActivateResponseReviewer `json:"reviewers" api:"required"`
	// The senders this sequence sends from; new contacts rotate across them. A change
	// applies to sends once the draft is published. An archived sender stays listed
	// but never sends; any change to the list drops it.
	Senders []SequenceActivateResponseSender `json:"senders" api:"required"`
	// `DRAFT` has not been published; `ACTIVE` is enabled for sending; `DEACTIVATED`
	// is published but paused.
	//
	// Any of "DRAFT", "ACTIVE", "DEACTIVATED".
	Status SequenceActivateResponseStatus `json:"status" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID               respjson.Field
		Channels         respjson.Field
		CreatedAt        respjson.Field
		Definition       respjson.Field
		FirstActivatedAt respjson.Field
		LastActivatedAt  respjson.Field
		Name             respjson.Field
		ObjectType       respjson.Field
		Reviewers        respjson.Field
		Senders          respjson.Field
		Status           respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SequenceActivateResponse) RawJSON() string { return r.JSON.raw }
func (r *SequenceActivateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Published configuration, or null before the first publish.
type SequenceActivateResponseDefinition struct {
	// What this version of the sequence is for.
	Description string `json:"description" api:"required"`
	// Messaging inputs that frame every step.
	Messaging SequenceActivateResponseDefinitionMessaging `json:"messaging" api:"required"`
	// How the whole sequence paces each step type it sends, keyed by sequence step
	// type.
	Pacing SequenceActivateResponseDefinitionPacing `json:"pacing" api:"required"`
	// Send priority relative to other sequences.
	//
	// Any of "HIGH", "NORMAL".
	Priority string `json:"priority" api:"required"`
	// When the sequence may send.
	SendWindow SequenceActivateResponseDefinitionSendWindow `json:"sendWindow" api:"required"`
	// Email, LinkedIn and wait steps in execution order; call steps are not exposed in
	// this release.
	Steps []SequenceActivateResponseDefinitionStepUnion `json:"steps" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Description respjson.Field
		Messaging   respjson.Field
		Pacing      respjson.Field
		Priority    respjson.Field
		SendWindow  respjson.Field
		Steps       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SequenceActivateResponseDefinition) RawJSON() string { return r.JSON.raw }
func (r *SequenceActivateResponseDefinition) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Messaging inputs that frame every step.
type SequenceActivateResponseDefinitionMessaging struct {
	// What the recipient is asked to do.
	CallToAction string `json:"callToAction" api:"required"`
	// Who the sequence targets.
	Persona string `json:"persona" api:"required"`
	// Why the outreach is timely.
	Trigger string `json:"trigger" api:"required"`
	// The value proposition.
	ValueProp string `json:"valueProp" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CallToAction respjson.Field
		Persona      respjson.Field
		Trigger      respjson.Field
		ValueProp    respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SequenceActivateResponseDefinitionMessaging) RawJSON() string { return r.JSON.raw }
func (r *SequenceActivateResponseDefinitionMessaging) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// How the whole sequence paces each step type it sends, keyed by sequence step
// type.
type SequenceActivateResponseDefinitionPacing struct {
	// Email steps.
	Email SequenceActivateResponseDefinitionPacingEmail `json:"EMAIL" api:"required"`
	// LinkedIn connection requests.
	LinkedinInvite SequenceActivateResponseDefinitionPacingLinkedinInvite `json:"LINKEDIN_INVITE" api:"required"`
	// LinkedIn messages.
	LinkedinMessage SequenceActivateResponseDefinitionPacingLinkedinMessage `json:"LINKEDIN_MESSAGE" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Email           respjson.Field
		LinkedinInvite  respjson.Field
		LinkedinMessage respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SequenceActivateResponseDefinitionPacing) RawJSON() string { return r.JSON.raw }
func (r *SequenceActivateResponseDefinitionPacing) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Email steps.
type SequenceActivateResponseDefinitionPacingEmail struct {
	// Most sends of this step type across the whole sequence in any rolling 24 hours,
	// or null when the sequence sets no daily cap. Workspace and sender limits still
	// apply.
	DailyLimit int64 `json:"dailyLimit" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DailyLimit  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SequenceActivateResponseDefinitionPacingEmail) RawJSON() string { return r.JSON.raw }
func (r *SequenceActivateResponseDefinitionPacingEmail) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// LinkedIn connection requests.
type SequenceActivateResponseDefinitionPacingLinkedinInvite struct {
	// Most sends of this step type across the whole sequence in any rolling 24 hours,
	// or null when the sequence sets no daily cap. Workspace and sender limits still
	// apply.
	DailyLimit int64 `json:"dailyLimit" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DailyLimit  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SequenceActivateResponseDefinitionPacingLinkedinInvite) RawJSON() string { return r.JSON.raw }
func (r *SequenceActivateResponseDefinitionPacingLinkedinInvite) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// LinkedIn messages.
type SequenceActivateResponseDefinitionPacingLinkedinMessage struct {
	// Most sends of this step type across the whole sequence in any rolling 24 hours,
	// or null when the sequence sets no daily cap. Workspace and sender limits still
	// apply.
	DailyLimit int64 `json:"dailyLimit" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DailyLimit  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SequenceActivateResponseDefinitionPacingLinkedinMessage) RawJSON() string { return r.JSON.raw }
func (r *SequenceActivateResponseDefinitionPacingLinkedinMessage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// When the sequence may send.
type SequenceActivateResponseDefinitionSendWindow struct {
	// Days sends may go out, 0 = Sunday.
	DaysOfWeek []int64 `json:"daysOfWeek" api:"required"`
	// Local hour sends stop, exclusive (1–24).
	EndHour int64 `json:"endHour" api:"required"`
	// First local hour sends may go out (0–23).
	StartHour int64 `json:"startHour" api:"required"`
	// IANA timezone of the window.
	Timezone string `json:"timezone" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DaysOfWeek  respjson.Field
		EndHour     respjson.Field
		StartHour   respjson.Field
		Timezone    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SequenceActivateResponseDefinitionSendWindow) RawJSON() string { return r.JSON.raw }
func (r *SequenceActivateResponseDefinitionSendWindow) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// SequenceActivateResponseDefinitionStepUnion contains all possible properties and
// values from [SequenceActivateResponseDefinitionStepObject],
// [SequenceActivateResponseDefinitionStepObject2],
// [SequenceActivateResponseDefinitionStepObject3].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type SequenceActivateResponseDefinitionStepUnion struct {
	ID           string `json:"id"`
	DeliveryMode string `json:"deliveryMode"`
	Name         string `json:"name"`
	Prompt       string `json:"prompt"`
	Type         string `json:"type"`
	// This field is from variant [SequenceActivateResponseDefinitionStepObject2].
	IncludeNote bool `json:"includeNote"`
	// This field is from variant [SequenceActivateResponseDefinitionStepObject3].
	Duration int64 `json:"duration"`
	// This field is from variant [SequenceActivateResponseDefinitionStepObject3].
	Unit string `json:"unit"`
	JSON struct {
		ID           respjson.Field
		DeliveryMode respjson.Field
		Name         respjson.Field
		Prompt       respjson.Field
		Type         respjson.Field
		IncludeNote  respjson.Field
		Duration     respjson.Field
		Unit         respjson.Field
		raw          string
	} `json:"-"`
}

func (u SequenceActivateResponseDefinitionStepUnion) AsSequenceActivateResponseDefinitionStepObject() (v SequenceActivateResponseDefinitionStepObject) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u SequenceActivateResponseDefinitionStepUnion) AsSequenceActivateResponseDefinitionStepObject2() (v SequenceActivateResponseDefinitionStepObject2) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u SequenceActivateResponseDefinitionStepUnion) AsSequenceActivateResponseDefinitionStepObject3() (v SequenceActivateResponseDefinitionStepObject3) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u SequenceActivateResponseDefinitionStepUnion) RawJSON() string { return u.JSON.raw }

func (r *SequenceActivateResponseDefinitionStepUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SequenceActivateResponseDefinitionStepObject struct {
	// Step identifier, referenced by enrollment `currentStepId` and message
	// `sequenceStepId`.
	ID string `json:"id" api:"required"`
	// `AUTO` sends without manual approval, subject to the send window and limits;
	// `FOR_REVIEW` requires a person to approve the draft before sending.
	//
	// Any of "AUTO", "FOR_REVIEW".
	DeliveryMode string `json:"deliveryMode" api:"required"`
	// Display name of the step.
	Name string `json:"name" api:"required"`
	// Instructions the agent follows when composing this message.
	Prompt string `json:"prompt" api:"required"`
	// Step type.
	//
	// Any of "EMAIL", "LINKEDIN_MESSAGE".
	Type string `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		DeliveryMode respjson.Field
		Name         respjson.Field
		Prompt       respjson.Field
		Type         respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SequenceActivateResponseDefinitionStepObject) RawJSON() string { return r.JSON.raw }
func (r *SequenceActivateResponseDefinitionStepObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SequenceActivateResponseDefinitionStepObject2 struct {
	// Step identifier, referenced by enrollment `currentStepId` and message
	// `sequenceStepId`.
	ID string `json:"id" api:"required"`
	// `AUTO` sends without manual approval, subject to the send window and limits;
	// `FOR_REVIEW` requires a person to approve the draft before sending.
	//
	// Any of "AUTO", "FOR_REVIEW".
	DeliveryMode string `json:"deliveryMode" api:"required"`
	// Whether the connection request carries a note. A step authored before this
	// setting existed reads as true.
	IncludeNote bool `json:"includeNote" api:"required"`
	// Display name of the step.
	Name string `json:"name" api:"required"`
	// Instructions the agent follows when composing this message.
	Prompt string `json:"prompt" api:"required"`
	// Step type.
	//
	// Any of "LINKEDIN_INVITE".
	Type string `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		DeliveryMode respjson.Field
		IncludeNote  respjson.Field
		Name         respjson.Field
		Prompt       respjson.Field
		Type         respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SequenceActivateResponseDefinitionStepObject2) RawJSON() string { return r.JSON.raw }
func (r *SequenceActivateResponseDefinitionStepObject2) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SequenceActivateResponseDefinitionStepObject3 struct {
	// Step identifier.
	ID string `json:"id" api:"required"`
	// Length of the wait.
	Duration int64 `json:"duration" api:"required"`
	// `WAIT` pauses for a fixed time; `WAIT_FOR_CONNECTION` waits up to that long for
	// a LinkedIn invite to be accepted.
	//
	// Any of "WAIT", "WAIT_FOR_CONNECTION".
	Type string `json:"type" api:"required"`
	// Unit of `duration`.
	//
	// Any of "DAY", "MINUTE".
	Unit string `json:"unit" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Duration    respjson.Field
		Type        respjson.Field
		Unit        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SequenceActivateResponseDefinitionStepObject3) RawJSON() string { return r.JSON.raw }
func (r *SequenceActivateResponseDefinitionStepObject3) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Always `sequence`.
type SequenceActivateResponseObjectType string

const (
	SequenceActivateResponseObjectTypeSequence SequenceActivateResponseObjectType = "sequence"
)

type SequenceActivateResponseReviewer struct {
	// ID of the member.
	ID string `json:"id" api:"required"`
	// Always `member`.
	//
	// Any of "member".
	ObjectType string `json:"objectType" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		ObjectType  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SequenceActivateResponseReviewer) RawJSON() string { return r.JSON.raw }
func (r *SequenceActivateResponseReviewer) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SequenceActivateResponseSender struct {
	// ID of the sender.
	ID string `json:"id" api:"required"`
	// Always `sender`.
	//
	// Any of "sender".
	ObjectType string `json:"objectType" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		ObjectType  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SequenceActivateResponseSender) RawJSON() string { return r.JSON.raw }
func (r *SequenceActivateResponseSender) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// `DRAFT` has not been published; `ACTIVE` is enabled for sending; `DEACTIVATED`
// is published but paused.
type SequenceActivateResponseStatus string

const (
	SequenceActivateResponseStatusDraft       SequenceActivateResponseStatus = "DRAFT"
	SequenceActivateResponseStatusActive      SequenceActivateResponseStatus = "ACTIVE"
	SequenceActivateResponseStatusDeactivated SequenceActivateResponseStatus = "DEACTIVATED"
)

type SequenceCreateResponse struct {
	// Channels the draft's steps send on.
	//
	// Any of "EMAIL", "LINKEDIN".
	Channels []string `json:"channels" api:"required"`
	// The draft configuration, in the same shape as a published definition.
	Definition SequenceCreateResponseDefinition `json:"definition" api:"required"`
	// Name the sequence takes when the draft is published.
	Name string `json:"name" api:"required"`
	// What publishing this draft would change for contacts already in the sequence.
	// Before the first publish, only proposed senders it would leave off because they
	// were archived. Empty when nothing is affected.
	PublishWarnings []SequenceCreateResponsePublishWarning `json:"publishWarnings" api:"required"`
	// Extra members who may approve this sequence’s `FOR_REVIEW` messages before they
	// send. The sending member and the sequence’s creator always review, whether or
	// not they are listed. A change applies to sends once the draft is published.
	// Lists active members only.
	Reviewers []SequenceCreateResponseReviewer `json:"reviewers" api:"required"`
	// The senders this sequence sends from; new contacts rotate across them. A change
	// applies to sends once the draft is published. The list this draft publishes: the
	// live list until a save changes it. An archived sender stays listed but never
	// sends; any change to the list drops it.
	Senders []SequenceCreateResponseSender `json:"senders" api:"required"`
	// The sequence this draft belongs to.
	SequenceID string `json:"sequenceId" api:"required"`
	// When the draft last changed.
	UpdatedAt string `json:"updatedAt" api:"required"`
	// Opaque token for the current state of the draft. Send it as `expectedVersion` on
	// the next save, publish or discard; it changes whenever anyone edits the draft.
	Version string `json:"version" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Channels        respjson.Field
		Definition      respjson.Field
		Name            respjson.Field
		PublishWarnings respjson.Field
		Reviewers       respjson.Field
		Senders         respjson.Field
		SequenceID      respjson.Field
		UpdatedAt       respjson.Field
		Version         respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SequenceCreateResponse) RawJSON() string { return r.JSON.raw }
func (r *SequenceCreateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The draft configuration, in the same shape as a published definition.
type SequenceCreateResponseDefinition struct {
	// What this version of the sequence is for.
	Description string `json:"description" api:"required"`
	// Messaging inputs that frame every step.
	Messaging SequenceCreateResponseDefinitionMessaging `json:"messaging" api:"required"`
	// How the whole sequence paces each step type it sends, keyed by sequence step
	// type.
	Pacing SequenceCreateResponseDefinitionPacing `json:"pacing" api:"required"`
	// Send priority relative to other sequences.
	//
	// Any of "HIGH", "NORMAL".
	Priority string `json:"priority" api:"required"`
	// When the sequence may send.
	SendWindow SequenceCreateResponseDefinitionSendWindow `json:"sendWindow" api:"required"`
	// Email, LinkedIn and wait steps in execution order; call steps are not exposed in
	// this release.
	Steps []SequenceCreateResponseDefinitionStepUnion `json:"steps" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Description respjson.Field
		Messaging   respjson.Field
		Pacing      respjson.Field
		Priority    respjson.Field
		SendWindow  respjson.Field
		Steps       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SequenceCreateResponseDefinition) RawJSON() string { return r.JSON.raw }
func (r *SequenceCreateResponseDefinition) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Messaging inputs that frame every step.
type SequenceCreateResponseDefinitionMessaging struct {
	// What the recipient is asked to do.
	CallToAction string `json:"callToAction" api:"required"`
	// Who the sequence targets.
	Persona string `json:"persona" api:"required"`
	// Why the outreach is timely.
	Trigger string `json:"trigger" api:"required"`
	// The value proposition.
	ValueProp string `json:"valueProp" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CallToAction respjson.Field
		Persona      respjson.Field
		Trigger      respjson.Field
		ValueProp    respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SequenceCreateResponseDefinitionMessaging) RawJSON() string { return r.JSON.raw }
func (r *SequenceCreateResponseDefinitionMessaging) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// How the whole sequence paces each step type it sends, keyed by sequence step
// type.
type SequenceCreateResponseDefinitionPacing struct {
	// Email steps.
	Email SequenceCreateResponseDefinitionPacingEmail `json:"EMAIL" api:"required"`
	// LinkedIn connection requests.
	LinkedinInvite SequenceCreateResponseDefinitionPacingLinkedinInvite `json:"LINKEDIN_INVITE" api:"required"`
	// LinkedIn messages.
	LinkedinMessage SequenceCreateResponseDefinitionPacingLinkedinMessage `json:"LINKEDIN_MESSAGE" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Email           respjson.Field
		LinkedinInvite  respjson.Field
		LinkedinMessage respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SequenceCreateResponseDefinitionPacing) RawJSON() string { return r.JSON.raw }
func (r *SequenceCreateResponseDefinitionPacing) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Email steps.
type SequenceCreateResponseDefinitionPacingEmail struct {
	// Most sends of this step type across the whole sequence in any rolling 24 hours,
	// or null when the sequence sets no daily cap. Workspace and sender limits still
	// apply.
	DailyLimit int64 `json:"dailyLimit" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DailyLimit  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SequenceCreateResponseDefinitionPacingEmail) RawJSON() string { return r.JSON.raw }
func (r *SequenceCreateResponseDefinitionPacingEmail) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// LinkedIn connection requests.
type SequenceCreateResponseDefinitionPacingLinkedinInvite struct {
	// Most sends of this step type across the whole sequence in any rolling 24 hours,
	// or null when the sequence sets no daily cap. Workspace and sender limits still
	// apply.
	DailyLimit int64 `json:"dailyLimit" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DailyLimit  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SequenceCreateResponseDefinitionPacingLinkedinInvite) RawJSON() string { return r.JSON.raw }
func (r *SequenceCreateResponseDefinitionPacingLinkedinInvite) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// LinkedIn messages.
type SequenceCreateResponseDefinitionPacingLinkedinMessage struct {
	// Most sends of this step type across the whole sequence in any rolling 24 hours,
	// or null when the sequence sets no daily cap. Workspace and sender limits still
	// apply.
	DailyLimit int64 `json:"dailyLimit" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DailyLimit  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SequenceCreateResponseDefinitionPacingLinkedinMessage) RawJSON() string { return r.JSON.raw }
func (r *SequenceCreateResponseDefinitionPacingLinkedinMessage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// When the sequence may send.
type SequenceCreateResponseDefinitionSendWindow struct {
	// Days sends may go out, 0 = Sunday.
	DaysOfWeek []int64 `json:"daysOfWeek" api:"required"`
	// Local hour sends stop, exclusive (1–24).
	EndHour int64 `json:"endHour" api:"required"`
	// First local hour sends may go out (0–23).
	StartHour int64 `json:"startHour" api:"required"`
	// IANA timezone of the window.
	Timezone string `json:"timezone" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DaysOfWeek  respjson.Field
		EndHour     respjson.Field
		StartHour   respjson.Field
		Timezone    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SequenceCreateResponseDefinitionSendWindow) RawJSON() string { return r.JSON.raw }
func (r *SequenceCreateResponseDefinitionSendWindow) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// SequenceCreateResponseDefinitionStepUnion contains all possible properties and
// values from [SequenceCreateResponseDefinitionStepObject],
// [SequenceCreateResponseDefinitionStepObject2],
// [SequenceCreateResponseDefinitionStepObject3].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type SequenceCreateResponseDefinitionStepUnion struct {
	ID           string `json:"id"`
	DeliveryMode string `json:"deliveryMode"`
	Name         string `json:"name"`
	Prompt       string `json:"prompt"`
	Type         string `json:"type"`
	// This field is from variant [SequenceCreateResponseDefinitionStepObject2].
	IncludeNote bool `json:"includeNote"`
	// This field is from variant [SequenceCreateResponseDefinitionStepObject3].
	Duration int64 `json:"duration"`
	// This field is from variant [SequenceCreateResponseDefinitionStepObject3].
	Unit string `json:"unit"`
	JSON struct {
		ID           respjson.Field
		DeliveryMode respjson.Field
		Name         respjson.Field
		Prompt       respjson.Field
		Type         respjson.Field
		IncludeNote  respjson.Field
		Duration     respjson.Field
		Unit         respjson.Field
		raw          string
	} `json:"-"`
}

func (u SequenceCreateResponseDefinitionStepUnion) AsSequenceCreateResponseDefinitionStepObject() (v SequenceCreateResponseDefinitionStepObject) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u SequenceCreateResponseDefinitionStepUnion) AsSequenceCreateResponseDefinitionStepObject2() (v SequenceCreateResponseDefinitionStepObject2) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u SequenceCreateResponseDefinitionStepUnion) AsSequenceCreateResponseDefinitionStepObject3() (v SequenceCreateResponseDefinitionStepObject3) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u SequenceCreateResponseDefinitionStepUnion) RawJSON() string { return u.JSON.raw }

func (r *SequenceCreateResponseDefinitionStepUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SequenceCreateResponseDefinitionStepObject struct {
	// Step identifier, referenced by enrollment `currentStepId` and message
	// `sequenceStepId`.
	ID string `json:"id" api:"required"`
	// `AUTO` sends without manual approval, subject to the send window and limits;
	// `FOR_REVIEW` requires a person to approve the draft before sending.
	//
	// Any of "AUTO", "FOR_REVIEW".
	DeliveryMode string `json:"deliveryMode" api:"required"`
	// Display name of the step.
	Name string `json:"name" api:"required"`
	// Instructions the agent follows when composing this message.
	Prompt string `json:"prompt" api:"required"`
	// Step type.
	//
	// Any of "EMAIL", "LINKEDIN_MESSAGE".
	Type string `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		DeliveryMode respjson.Field
		Name         respjson.Field
		Prompt       respjson.Field
		Type         respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SequenceCreateResponseDefinitionStepObject) RawJSON() string { return r.JSON.raw }
func (r *SequenceCreateResponseDefinitionStepObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SequenceCreateResponseDefinitionStepObject2 struct {
	// Step identifier, referenced by enrollment `currentStepId` and message
	// `sequenceStepId`.
	ID string `json:"id" api:"required"`
	// `AUTO` sends without manual approval, subject to the send window and limits;
	// `FOR_REVIEW` requires a person to approve the draft before sending.
	//
	// Any of "AUTO", "FOR_REVIEW".
	DeliveryMode string `json:"deliveryMode" api:"required"`
	// Whether the connection request carries a note. A step authored before this
	// setting existed reads as true.
	IncludeNote bool `json:"includeNote" api:"required"`
	// Display name of the step.
	Name string `json:"name" api:"required"`
	// Instructions the agent follows when composing this message.
	Prompt string `json:"prompt" api:"required"`
	// Step type.
	//
	// Any of "LINKEDIN_INVITE".
	Type string `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		DeliveryMode respjson.Field
		IncludeNote  respjson.Field
		Name         respjson.Field
		Prompt       respjson.Field
		Type         respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SequenceCreateResponseDefinitionStepObject2) RawJSON() string { return r.JSON.raw }
func (r *SequenceCreateResponseDefinitionStepObject2) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SequenceCreateResponseDefinitionStepObject3 struct {
	// Step identifier.
	ID string `json:"id" api:"required"`
	// Length of the wait.
	Duration int64 `json:"duration" api:"required"`
	// `WAIT` pauses for a fixed time; `WAIT_FOR_CONNECTION` waits up to that long for
	// a LinkedIn invite to be accepted.
	//
	// Any of "WAIT", "WAIT_FOR_CONNECTION".
	Type string `json:"type" api:"required"`
	// Unit of `duration`.
	//
	// Any of "DAY", "MINUTE".
	Unit string `json:"unit" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Duration    respjson.Field
		Type        respjson.Field
		Unit        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SequenceCreateResponseDefinitionStepObject3) RawJSON() string { return r.JSON.raw }
func (r *SequenceCreateResponseDefinitionStepObject3) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SequenceCreateResponsePublishWarning struct {
	// What publishing would change.
	//
	// Any of "delay_shortened", "delay_lengthened", "connection_wait_shortened",
	// "connection_wait_lengthened", "send_schedule_changed",
	// "send_configuration_changed", "send_priority_changed", "senders_changed",
	// "email_sender_pool_emptied", "linkedin_account_pool_emptied",
	// "assigned_sender_removed", "delivery_became_automatic", "invite_note_enabled",
	// "invite_note_disabled", "step_inserted", "automatic_send_step_inserted",
	// "step_removed", "step_type_changed", "step_order_changed",
	// "full_identity_rewrite", "low_identity_overlap".
	Code string `json:"code" api:"required"`
	// What would change, in plain language.
	Message string `json:"message" api:"required"`
	// `destructive` when contacts already in the sequence may skip, repeat or receive
	// a step they would not have; `warning` otherwise.
	//
	// Any of "warning", "destructive".
	Severity string `json:"severity" api:"required"`
	// The steps involved; empty for a sequence-wide change.
	StepIDs []string `json:"stepIds" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Code        respjson.Field
		Message     respjson.Field
		Severity    respjson.Field
		StepIDs     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SequenceCreateResponsePublishWarning) RawJSON() string { return r.JSON.raw }
func (r *SequenceCreateResponsePublishWarning) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SequenceCreateResponseReviewer struct {
	// ID of the member.
	ID string `json:"id" api:"required"`
	// Always `member`.
	//
	// Any of "member".
	ObjectType string `json:"objectType" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		ObjectType  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SequenceCreateResponseReviewer) RawJSON() string { return r.JSON.raw }
func (r *SequenceCreateResponseReviewer) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SequenceCreateResponseSender struct {
	// ID of the sender.
	ID string `json:"id" api:"required"`
	// Always `sender`.
	//
	// Any of "sender".
	ObjectType string `json:"objectType" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		ObjectType  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SequenceCreateResponseSender) RawJSON() string { return r.JSON.raw }
func (r *SequenceCreateResponseSender) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SequenceDeactivateResponse struct {
	// Unique identifier of the sequence.
	ID string `json:"id" api:"required"`
	// Channels the published steps send on; empty before the first publish.
	//
	// Any of "EMAIL", "LINKEDIN".
	Channels []string `json:"channels" api:"required"`
	// ISO 8601 creation timestamp.
	CreatedAt string `json:"createdAt" api:"required"`
	// Published configuration, or null before the first publish.
	Definition SequenceDeactivateResponseDefinition `json:"definition" api:"required"`
	// When the sequence was first activated.
	FirstActivatedAt string `json:"firstActivatedAt" api:"required"`
	// When the sequence was most recently activated.
	LastActivatedAt string `json:"lastActivatedAt" api:"required"`
	// Name of the sequence.
	Name string `json:"name" api:"required"`
	// Always `sequence`.
	//
	// Any of "sequence".
	ObjectType SequenceDeactivateResponseObjectType `json:"objectType" api:"required"`
	// Extra members who may approve this sequence’s `FOR_REVIEW` messages before they
	// send. The sending member and the sequence’s creator always review, whether or
	// not they are listed. A change applies to sends once the draft is published. The
	// published list; empty before the first publish. Lists active members only.
	Reviewers []SequenceDeactivateResponseReviewer `json:"reviewers" api:"required"`
	// The senders this sequence sends from; new contacts rotate across them. A change
	// applies to sends once the draft is published. An archived sender stays listed
	// but never sends; any change to the list drops it.
	Senders []SequenceDeactivateResponseSender `json:"senders" api:"required"`
	// `DRAFT` has not been published; `ACTIVE` is enabled for sending; `DEACTIVATED`
	// is published but paused.
	//
	// Any of "DRAFT", "ACTIVE", "DEACTIVATED".
	Status SequenceDeactivateResponseStatus `json:"status" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID               respjson.Field
		Channels         respjson.Field
		CreatedAt        respjson.Field
		Definition       respjson.Field
		FirstActivatedAt respjson.Field
		LastActivatedAt  respjson.Field
		Name             respjson.Field
		ObjectType       respjson.Field
		Reviewers        respjson.Field
		Senders          respjson.Field
		Status           respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SequenceDeactivateResponse) RawJSON() string { return r.JSON.raw }
func (r *SequenceDeactivateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Published configuration, or null before the first publish.
type SequenceDeactivateResponseDefinition struct {
	// What this version of the sequence is for.
	Description string `json:"description" api:"required"`
	// Messaging inputs that frame every step.
	Messaging SequenceDeactivateResponseDefinitionMessaging `json:"messaging" api:"required"`
	// How the whole sequence paces each step type it sends, keyed by sequence step
	// type.
	Pacing SequenceDeactivateResponseDefinitionPacing `json:"pacing" api:"required"`
	// Send priority relative to other sequences.
	//
	// Any of "HIGH", "NORMAL".
	Priority string `json:"priority" api:"required"`
	// When the sequence may send.
	SendWindow SequenceDeactivateResponseDefinitionSendWindow `json:"sendWindow" api:"required"`
	// Email, LinkedIn and wait steps in execution order; call steps are not exposed in
	// this release.
	Steps []SequenceDeactivateResponseDefinitionStepUnion `json:"steps" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Description respjson.Field
		Messaging   respjson.Field
		Pacing      respjson.Field
		Priority    respjson.Field
		SendWindow  respjson.Field
		Steps       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SequenceDeactivateResponseDefinition) RawJSON() string { return r.JSON.raw }
func (r *SequenceDeactivateResponseDefinition) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Messaging inputs that frame every step.
type SequenceDeactivateResponseDefinitionMessaging struct {
	// What the recipient is asked to do.
	CallToAction string `json:"callToAction" api:"required"`
	// Who the sequence targets.
	Persona string `json:"persona" api:"required"`
	// Why the outreach is timely.
	Trigger string `json:"trigger" api:"required"`
	// The value proposition.
	ValueProp string `json:"valueProp" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CallToAction respjson.Field
		Persona      respjson.Field
		Trigger      respjson.Field
		ValueProp    respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SequenceDeactivateResponseDefinitionMessaging) RawJSON() string { return r.JSON.raw }
func (r *SequenceDeactivateResponseDefinitionMessaging) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// How the whole sequence paces each step type it sends, keyed by sequence step
// type.
type SequenceDeactivateResponseDefinitionPacing struct {
	// Email steps.
	Email SequenceDeactivateResponseDefinitionPacingEmail `json:"EMAIL" api:"required"`
	// LinkedIn connection requests.
	LinkedinInvite SequenceDeactivateResponseDefinitionPacingLinkedinInvite `json:"LINKEDIN_INVITE" api:"required"`
	// LinkedIn messages.
	LinkedinMessage SequenceDeactivateResponseDefinitionPacingLinkedinMessage `json:"LINKEDIN_MESSAGE" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Email           respjson.Field
		LinkedinInvite  respjson.Field
		LinkedinMessage respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SequenceDeactivateResponseDefinitionPacing) RawJSON() string { return r.JSON.raw }
func (r *SequenceDeactivateResponseDefinitionPacing) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Email steps.
type SequenceDeactivateResponseDefinitionPacingEmail struct {
	// Most sends of this step type across the whole sequence in any rolling 24 hours,
	// or null when the sequence sets no daily cap. Workspace and sender limits still
	// apply.
	DailyLimit int64 `json:"dailyLimit" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DailyLimit  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SequenceDeactivateResponseDefinitionPacingEmail) RawJSON() string { return r.JSON.raw }
func (r *SequenceDeactivateResponseDefinitionPacingEmail) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// LinkedIn connection requests.
type SequenceDeactivateResponseDefinitionPacingLinkedinInvite struct {
	// Most sends of this step type across the whole sequence in any rolling 24 hours,
	// or null when the sequence sets no daily cap. Workspace and sender limits still
	// apply.
	DailyLimit int64 `json:"dailyLimit" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DailyLimit  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SequenceDeactivateResponseDefinitionPacingLinkedinInvite) RawJSON() string { return r.JSON.raw }
func (r *SequenceDeactivateResponseDefinitionPacingLinkedinInvite) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// LinkedIn messages.
type SequenceDeactivateResponseDefinitionPacingLinkedinMessage struct {
	// Most sends of this step type across the whole sequence in any rolling 24 hours,
	// or null when the sequence sets no daily cap. Workspace and sender limits still
	// apply.
	DailyLimit int64 `json:"dailyLimit" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DailyLimit  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SequenceDeactivateResponseDefinitionPacingLinkedinMessage) RawJSON() string {
	return r.JSON.raw
}
func (r *SequenceDeactivateResponseDefinitionPacingLinkedinMessage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// When the sequence may send.
type SequenceDeactivateResponseDefinitionSendWindow struct {
	// Days sends may go out, 0 = Sunday.
	DaysOfWeek []int64 `json:"daysOfWeek" api:"required"`
	// Local hour sends stop, exclusive (1–24).
	EndHour int64 `json:"endHour" api:"required"`
	// First local hour sends may go out (0–23).
	StartHour int64 `json:"startHour" api:"required"`
	// IANA timezone of the window.
	Timezone string `json:"timezone" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DaysOfWeek  respjson.Field
		EndHour     respjson.Field
		StartHour   respjson.Field
		Timezone    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SequenceDeactivateResponseDefinitionSendWindow) RawJSON() string { return r.JSON.raw }
func (r *SequenceDeactivateResponseDefinitionSendWindow) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// SequenceDeactivateResponseDefinitionStepUnion contains all possible properties
// and values from [SequenceDeactivateResponseDefinitionStepObject],
// [SequenceDeactivateResponseDefinitionStepObject2],
// [SequenceDeactivateResponseDefinitionStepObject3].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type SequenceDeactivateResponseDefinitionStepUnion struct {
	ID           string `json:"id"`
	DeliveryMode string `json:"deliveryMode"`
	Name         string `json:"name"`
	Prompt       string `json:"prompt"`
	Type         string `json:"type"`
	// This field is from variant [SequenceDeactivateResponseDefinitionStepObject2].
	IncludeNote bool `json:"includeNote"`
	// This field is from variant [SequenceDeactivateResponseDefinitionStepObject3].
	Duration int64 `json:"duration"`
	// This field is from variant [SequenceDeactivateResponseDefinitionStepObject3].
	Unit string `json:"unit"`
	JSON struct {
		ID           respjson.Field
		DeliveryMode respjson.Field
		Name         respjson.Field
		Prompt       respjson.Field
		Type         respjson.Field
		IncludeNote  respjson.Field
		Duration     respjson.Field
		Unit         respjson.Field
		raw          string
	} `json:"-"`
}

func (u SequenceDeactivateResponseDefinitionStepUnion) AsSequenceDeactivateResponseDefinitionStepObject() (v SequenceDeactivateResponseDefinitionStepObject) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u SequenceDeactivateResponseDefinitionStepUnion) AsSequenceDeactivateResponseDefinitionStepObject2() (v SequenceDeactivateResponseDefinitionStepObject2) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u SequenceDeactivateResponseDefinitionStepUnion) AsSequenceDeactivateResponseDefinitionStepObject3() (v SequenceDeactivateResponseDefinitionStepObject3) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u SequenceDeactivateResponseDefinitionStepUnion) RawJSON() string { return u.JSON.raw }

func (r *SequenceDeactivateResponseDefinitionStepUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SequenceDeactivateResponseDefinitionStepObject struct {
	// Step identifier, referenced by enrollment `currentStepId` and message
	// `sequenceStepId`.
	ID string `json:"id" api:"required"`
	// `AUTO` sends without manual approval, subject to the send window and limits;
	// `FOR_REVIEW` requires a person to approve the draft before sending.
	//
	// Any of "AUTO", "FOR_REVIEW".
	DeliveryMode string `json:"deliveryMode" api:"required"`
	// Display name of the step.
	Name string `json:"name" api:"required"`
	// Instructions the agent follows when composing this message.
	Prompt string `json:"prompt" api:"required"`
	// Step type.
	//
	// Any of "EMAIL", "LINKEDIN_MESSAGE".
	Type string `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		DeliveryMode respjson.Field
		Name         respjson.Field
		Prompt       respjson.Field
		Type         respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SequenceDeactivateResponseDefinitionStepObject) RawJSON() string { return r.JSON.raw }
func (r *SequenceDeactivateResponseDefinitionStepObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SequenceDeactivateResponseDefinitionStepObject2 struct {
	// Step identifier, referenced by enrollment `currentStepId` and message
	// `sequenceStepId`.
	ID string `json:"id" api:"required"`
	// `AUTO` sends without manual approval, subject to the send window and limits;
	// `FOR_REVIEW` requires a person to approve the draft before sending.
	//
	// Any of "AUTO", "FOR_REVIEW".
	DeliveryMode string `json:"deliveryMode" api:"required"`
	// Whether the connection request carries a note. A step authored before this
	// setting existed reads as true.
	IncludeNote bool `json:"includeNote" api:"required"`
	// Display name of the step.
	Name string `json:"name" api:"required"`
	// Instructions the agent follows when composing this message.
	Prompt string `json:"prompt" api:"required"`
	// Step type.
	//
	// Any of "LINKEDIN_INVITE".
	Type string `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		DeliveryMode respjson.Field
		IncludeNote  respjson.Field
		Name         respjson.Field
		Prompt       respjson.Field
		Type         respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SequenceDeactivateResponseDefinitionStepObject2) RawJSON() string { return r.JSON.raw }
func (r *SequenceDeactivateResponseDefinitionStepObject2) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SequenceDeactivateResponseDefinitionStepObject3 struct {
	// Step identifier.
	ID string `json:"id" api:"required"`
	// Length of the wait.
	Duration int64 `json:"duration" api:"required"`
	// `WAIT` pauses for a fixed time; `WAIT_FOR_CONNECTION` waits up to that long for
	// a LinkedIn invite to be accepted.
	//
	// Any of "WAIT", "WAIT_FOR_CONNECTION".
	Type string `json:"type" api:"required"`
	// Unit of `duration`.
	//
	// Any of "DAY", "MINUTE".
	Unit string `json:"unit" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Duration    respjson.Field
		Type        respjson.Field
		Unit        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SequenceDeactivateResponseDefinitionStepObject3) RawJSON() string { return r.JSON.raw }
func (r *SequenceDeactivateResponseDefinitionStepObject3) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Always `sequence`.
type SequenceDeactivateResponseObjectType string

const (
	SequenceDeactivateResponseObjectTypeSequence SequenceDeactivateResponseObjectType = "sequence"
)

type SequenceDeactivateResponseReviewer struct {
	// ID of the member.
	ID string `json:"id" api:"required"`
	// Always `member`.
	//
	// Any of "member".
	ObjectType string `json:"objectType" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		ObjectType  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SequenceDeactivateResponseReviewer) RawJSON() string { return r.JSON.raw }
func (r *SequenceDeactivateResponseReviewer) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SequenceDeactivateResponseSender struct {
	// ID of the sender.
	ID string `json:"id" api:"required"`
	// Always `sender`.
	//
	// Any of "sender".
	ObjectType string `json:"objectType" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		ObjectType  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SequenceDeactivateResponseSender) RawJSON() string { return r.JSON.raw }
func (r *SequenceDeactivateResponseSender) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// `DRAFT` has not been published; `ACTIVE` is enabled for sending; `DEACTIVATED`
// is published but paused.
type SequenceDeactivateResponseStatus string

const (
	SequenceDeactivateResponseStatusDraft       SequenceDeactivateResponseStatus = "DRAFT"
	SequenceDeactivateResponseStatusActive      SequenceDeactivateResponseStatus = "ACTIVE"
	SequenceDeactivateResponseStatusDeactivated SequenceDeactivateResponseStatus = "DEACTIVATED"
)

type SequenceDiscardDraftResponse struct {
	// True when the sequence had never been published, so discarding its only draft
	// deleted it.
	SequenceDeleted bool `json:"sequenceDeleted" api:"required"`
	// The sequence whose draft was discarded.
	SequenceID string `json:"sequenceId" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		SequenceDeleted respjson.Field
		SequenceID      respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SequenceDiscardDraftResponse) RawJSON() string { return r.JSON.raw }
func (r *SequenceDiscardDraftResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SequenceListEnrollmentsResponse struct {
	Data []SequenceListEnrollmentsResponseData `json:"data" api:"required"`
	// Any of "list".
	Object     SequenceListEnrollmentsResponseObject `json:"object" api:"required"`
	TotalCount int64                                 `json:"totalCount" api:"required"`
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
func (r SequenceListEnrollmentsResponse) RawJSON() string { return r.JSON.raw }
func (r *SequenceListEnrollmentsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SequenceListEnrollmentsResponseData struct {
	// Unique identifier of the enrollment.
	ID string `json:"id" api:"required"`
	// When the enrollment was cancelled, or null if it was not.
	CancelledAt string `json:"cancelledAt" api:"required"`
	// When the enrollment completed, or null if it has not.
	CompletedAt string `json:"completedAt" api:"required"`
	// The enrolled contact.
	ContactID string `json:"contactId" api:"required"`
	// The current step ID from `definition.steps`, or null when no current step
	// exposed by this API can be resolved.
	CurrentStepID string `json:"currentStepId" api:"required"`
	// When the contact was enrolled.
	EnrolledAt string `json:"enrolledAt" api:"required"`
	// Member ID of the person who enrolled the contact, or null when no enrolling user
	// or workspace membership can be resolved.
	EnrolledBy string `json:"enrolledBy" api:"required"`
	// When the activity selected for display happened. A terminal reply or bounce can
	// take precedence over a later open.
	LastActivityAt string `json:"lastActivityAt" api:"required"`
	// The activity selected for display. A terminal reply or bounce can take
	// precedence over a later open.
	//
	// Any of "SENT", "OPENED", "REPLIED", "BOUNCED", "INVITED", "CONNECTED", "SEEN".
	LastActivityKind string `json:"lastActivityKind" api:"required"`
	// Before a send is queued, an estimate from the planned cadence and send window.
	// For a ready queued send, the earliest current eligibility based on its fixed
	// sender, pacing limits, and send window. Randomized gaps, jitter, other queued
	// work, or new sends may push it later. Null while awaiting review, sending,
	// paused, or when eligibility is unknown.
	NextSendAt string `json:"nextSendAt" api:"required"`
	// Unsent email and LinkedIn steps remaining, including connection requests; null
	// on retrieval when a current indexed value is unavailable.
	RemainingSends int64 `json:"remainingSends" api:"required"`
	// When this enrollment first received a reply across email or LinkedIn.
	RepliedAt string `json:"repliedAt" api:"required"`
	// Display status of the enrollment, or null on retrieval when the index is
	// unavailable or has not caught up with the enrollment lifecycle.
	//
	// Any of "ACTIVE", "AWAITING_REVIEW", "NEEDS_ATTENTION", "REPLIED", "COMPLETED",
	// "PAUSED", "CANCELLED", "FAILED".
	Status string `json:"status" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID               respjson.Field
		CancelledAt      respjson.Field
		CompletedAt      respjson.Field
		ContactID        respjson.Field
		CurrentStepID    respjson.Field
		EnrolledAt       respjson.Field
		EnrolledBy       respjson.Field
		LastActivityAt   respjson.Field
		LastActivityKind respjson.Field
		NextSendAt       respjson.Field
		RemainingSends   respjson.Field
		RepliedAt        respjson.Field
		Status           respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SequenceListEnrollmentsResponseData) RawJSON() string { return r.JSON.raw }
func (r *SequenceListEnrollmentsResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SequenceListEnrollmentsResponseObject string

const (
	SequenceListEnrollmentsResponseObjectList SequenceListEnrollmentsResponseObject = "list"
)

type SequenceListMessagesResponse struct {
	Data []SequenceListMessagesResponseDataUnion `json:"data" api:"required"`
	// Any of "list".
	Object     SequenceListMessagesResponseObject `json:"object" api:"required"`
	TotalCount int64                              `json:"totalCount" api:"required"`
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
func (r SequenceListMessagesResponse) RawJSON() string { return r.JSON.raw }
func (r *SequenceListMessagesResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// SequenceListMessagesResponseDataUnion contains all possible properties and
// values from [SequenceListMessagesResponseDataObject],
// [SequenceListMessagesResponseDataObject2].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type SequenceListMessagesResponseDataUnion struct {
	ID string `json:"id"`
	// This field is from variant [SequenceListMessagesResponseDataObject].
	AcceptedAt string `json:"acceptedAt"`
	// This field is from variant [SequenceListMessagesResponseDataObject].
	BouncedAt    string `json:"bouncedAt"`
	Channel      string `json:"channel"`
	ContactID    string `json:"contactId"`
	CreatedAt    string `json:"createdAt"`
	Direction    string `json:"direction"`
	EnrollmentID string `json:"enrollmentId"`
	// This field is from variant [SequenceListMessagesResponseDataObject].
	OpenedAt string `json:"openedAt"`
	// This field is from variant [SequenceListMessagesResponseDataObject].
	ReadAt string `json:"readAt"`
	// This field is a union of [SequenceListMessagesResponseDataObjectReferenceUnion],
	// [SequenceListMessagesResponseDataObject2ReferenceUnion]
	Reference SequenceListMessagesResponseDataUnionReference `json:"reference"`
	// This field is from variant [SequenceListMessagesResponseDataObject].
	RepliedAt string `json:"repliedAt"`
	// This field is from variant [SequenceListMessagesResponseDataObject].
	SentAt           string `json:"sentAt"`
	SequenceStepID   string `json:"sequenceStepId"`
	SequenceStepType string `json:"sequenceStepType"`
	// This field is from variant [SequenceListMessagesResponseDataObject].
	Status string `json:"status"`
	// This field is from variant [SequenceListMessagesResponseDataObject].
	StatusChangedAt string `json:"statusChangedAt"`
	// This field is from variant [SequenceListMessagesResponseDataObject2].
	ReceivedAt string `json:"receivedAt"`
	JSON       struct {
		ID               respjson.Field
		AcceptedAt       respjson.Field
		BouncedAt        respjson.Field
		Channel          respjson.Field
		ContactID        respjson.Field
		CreatedAt        respjson.Field
		Direction        respjson.Field
		EnrollmentID     respjson.Field
		OpenedAt         respjson.Field
		ReadAt           respjson.Field
		Reference        respjson.Field
		RepliedAt        respjson.Field
		SentAt           respjson.Field
		SequenceStepID   respjson.Field
		SequenceStepType respjson.Field
		Status           respjson.Field
		StatusChangedAt  respjson.Field
		ReceivedAt       respjson.Field
		raw              string
	} `json:"-"`
}

func (u SequenceListMessagesResponseDataUnion) AsSequenceListMessagesResponseDataObject() (v SequenceListMessagesResponseDataObject) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u SequenceListMessagesResponseDataUnion) AsSequenceListMessagesResponseDataObject2() (v SequenceListMessagesResponseDataObject2) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u SequenceListMessagesResponseDataUnion) RawJSON() string { return u.JSON.raw }

func (r *SequenceListMessagesResponseDataUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// SequenceListMessagesResponseDataUnionReference is an implicit subunion of
// [SequenceListMessagesResponseDataUnion].
// SequenceListMessagesResponseDataUnionReference provides convenient access to the
// sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [SequenceListMessagesResponseDataUnion].
type SequenceListMessagesResponseDataUnionReference struct {
	ID         string `json:"id"`
	ObjectType string `json:"objectType"`
	JSON       struct {
		ID         respjson.Field
		ObjectType respjson.Field
		raw        string
	} `json:"-"`
}

func (r *SequenceListMessagesResponseDataUnionReference) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SequenceListMessagesResponseDataObject struct {
	// Opaque sequence message ID beginning with `sqm_`; use the ID returned by list
	// when retrieving a message.
	ID         string `json:"id" api:"required"`
	AcceptedAt string `json:"acceptedAt" api:"required"`
	BouncedAt  string `json:"bouncedAt" api:"required"`
	// Any of "EMAIL", "LINKEDIN".
	Channel string `json:"channel" api:"required"`
	// Contact ID beginning with `con_`.
	ContactID string `json:"contactId" api:"required"`
	// ISO 8601 creation timestamp of the current outbound attempt, or receipt
	// timestamp for an incoming reply; lists sort newest first by this field.
	CreatedAt string `json:"createdAt" api:"required"`
	// Any of "OUTBOUND".
	Direction string `json:"direction" api:"required"`
	// Enrollment ID beginning with `sqe_`.
	EnrollmentID string `json:"enrollmentId" api:"required"`
	OpenedAt     string `json:"openedAt" api:"required"`
	ReadAt       string `json:"readAt" api:"required"`
	// Active same-workspace email or chat message reference, or null when unavailable;
	// individual mail/chat privacy does not restrict this sequence reference, but the
	// referenced endpoint retains its own scopes and access rules.
	Reference SequenceListMessagesResponseDataObjectReferenceUnion `json:"reference" api:"required"`
	RepliedAt string                                               `json:"repliedAt" api:"required"`
	SentAt    string                                               `json:"sentAt" api:"required"`
	// Authored step associated with this outbound message or with the outbound message
	// that received this reply.
	SequenceStepID string `json:"sequenceStepId" api:"required"`
	// Any of "EMAIL", "LINKEDIN_MESSAGE", "LINKEDIN_INVITE".
	SequenceStepType string `json:"sequenceStepType" api:"required"`
	// Any of "AWAITING_REVIEW", "SCHEDULED", "QUEUED", "SENT", "SKIPPED", "FAILED",
	// "CANCELLED".
	Status string `json:"status" api:"required"`
	// When the outbound status last changed; engagement timestamps do not necessarily
	// change this value.
	StatusChangedAt string `json:"statusChangedAt" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID               respjson.Field
		AcceptedAt       respjson.Field
		BouncedAt        respjson.Field
		Channel          respjson.Field
		ContactID        respjson.Field
		CreatedAt        respjson.Field
		Direction        respjson.Field
		EnrollmentID     respjson.Field
		OpenedAt         respjson.Field
		ReadAt           respjson.Field
		Reference        respjson.Field
		RepliedAt        respjson.Field
		SentAt           respjson.Field
		SequenceStepID   respjson.Field
		SequenceStepType respjson.Field
		Status           respjson.Field
		StatusChangedAt  respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SequenceListMessagesResponseDataObject) RawJSON() string { return r.JSON.raw }
func (r *SequenceListMessagesResponseDataObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// SequenceListMessagesResponseDataObjectReferenceUnion contains all possible
// properties and values from
// [SequenceListMessagesResponseDataObjectReferenceObject],
// [SequenceListMessagesResponseDataObjectReferenceObject2].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type SequenceListMessagesResponseDataObjectReferenceUnion struct {
	ID         string `json:"id"`
	ObjectType string `json:"objectType"`
	JSON       struct {
		ID         respjson.Field
		ObjectType respjson.Field
		raw        string
	} `json:"-"`
}

func (u SequenceListMessagesResponseDataObjectReferenceUnion) AsSequenceListMessagesResponseDataObjectReferenceObject() (v SequenceListMessagesResponseDataObjectReferenceObject) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u SequenceListMessagesResponseDataObjectReferenceUnion) AsSequenceListMessagesResponseDataObjectReferenceObject2() (v SequenceListMessagesResponseDataObjectReferenceObject2) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u SequenceListMessagesResponseDataObjectReferenceUnion) RawJSON() string { return u.JSON.raw }

func (r *SequenceListMessagesResponseDataObjectReferenceUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SequenceListMessagesResponseDataObjectReferenceObject struct {
	ID string `json:"id" api:"required"`
	// Any of "email".
	ObjectType string `json:"objectType" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		ObjectType  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SequenceListMessagesResponseDataObjectReferenceObject) RawJSON() string { return r.JSON.raw }
func (r *SequenceListMessagesResponseDataObjectReferenceObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SequenceListMessagesResponseDataObjectReferenceObject2 struct {
	ID string `json:"id" api:"required"`
	// Any of "message".
	ObjectType string `json:"objectType" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		ObjectType  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SequenceListMessagesResponseDataObjectReferenceObject2) RawJSON() string { return r.JSON.raw }
func (r *SequenceListMessagesResponseDataObjectReferenceObject2) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SequenceListMessagesResponseDataObject2 struct {
	// Opaque sequence message ID beginning with `sqm_`; use the ID returned by list
	// when retrieving a message.
	ID string `json:"id" api:"required"`
	// Any of "EMAIL", "LINKEDIN".
	Channel string `json:"channel" api:"required"`
	// Contact ID beginning with `con_`.
	ContactID string `json:"contactId" api:"required"`
	// ISO 8601 creation timestamp of the current outbound attempt, or receipt
	// timestamp for an incoming reply; lists sort newest first by this field.
	CreatedAt string `json:"createdAt" api:"required"`
	// Any of "INBOUND".
	Direction string `json:"direction" api:"required"`
	// Enrollment ID beginning with `sqe_`.
	EnrollmentID string `json:"enrollmentId" api:"required"`
	ReceivedAt   string `json:"receivedAt" api:"required"`
	// Active same-workspace email or chat message reference, or null when unavailable;
	// individual mail/chat privacy does not restrict this sequence reference, but the
	// referenced endpoint retains its own scopes and access rules.
	Reference SequenceListMessagesResponseDataObject2ReferenceUnion `json:"reference" api:"required"`
	// Authored step associated with this outbound message or with the outbound message
	// that received this reply.
	SequenceStepID string `json:"sequenceStepId" api:"required"`
	// Any of "EMAIL", "LINKEDIN_MESSAGE", "LINKEDIN_INVITE".
	SequenceStepType string `json:"sequenceStepType" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID               respjson.Field
		Channel          respjson.Field
		ContactID        respjson.Field
		CreatedAt        respjson.Field
		Direction        respjson.Field
		EnrollmentID     respjson.Field
		ReceivedAt       respjson.Field
		Reference        respjson.Field
		SequenceStepID   respjson.Field
		SequenceStepType respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SequenceListMessagesResponseDataObject2) RawJSON() string { return r.JSON.raw }
func (r *SequenceListMessagesResponseDataObject2) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// SequenceListMessagesResponseDataObject2ReferenceUnion contains all possible
// properties and values from
// [SequenceListMessagesResponseDataObject2ReferenceObject],
// [SequenceListMessagesResponseDataObject2ReferenceObject2].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type SequenceListMessagesResponseDataObject2ReferenceUnion struct {
	ID         string `json:"id"`
	ObjectType string `json:"objectType"`
	JSON       struct {
		ID         respjson.Field
		ObjectType respjson.Field
		raw        string
	} `json:"-"`
}

func (u SequenceListMessagesResponseDataObject2ReferenceUnion) AsSequenceListMessagesResponseDataObject2ReferenceObject() (v SequenceListMessagesResponseDataObject2ReferenceObject) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u SequenceListMessagesResponseDataObject2ReferenceUnion) AsSequenceListMessagesResponseDataObject2ReferenceObject2() (v SequenceListMessagesResponseDataObject2ReferenceObject2) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u SequenceListMessagesResponseDataObject2ReferenceUnion) RawJSON() string { return u.JSON.raw }

func (r *SequenceListMessagesResponseDataObject2ReferenceUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SequenceListMessagesResponseDataObject2ReferenceObject struct {
	ID string `json:"id" api:"required"`
	// Any of "email".
	ObjectType string `json:"objectType" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		ObjectType  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SequenceListMessagesResponseDataObject2ReferenceObject) RawJSON() string { return r.JSON.raw }
func (r *SequenceListMessagesResponseDataObject2ReferenceObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SequenceListMessagesResponseDataObject2ReferenceObject2 struct {
	ID string `json:"id" api:"required"`
	// Any of "message".
	ObjectType string `json:"objectType" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		ObjectType  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SequenceListMessagesResponseDataObject2ReferenceObject2) RawJSON() string { return r.JSON.raw }
func (r *SequenceListMessagesResponseDataObject2ReferenceObject2) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SequenceListMessagesResponseObject string

const (
	SequenceListMessagesResponseObjectList SequenceListMessagesResponseObject = "list"
)

type SequenceListResponse struct {
	Data []SequenceListResponseData `json:"data" api:"required"`
	// Any of "list".
	Object     SequenceListResponseObject `json:"object" api:"required"`
	TotalCount int64                      `json:"totalCount" api:"required"`
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
func (r SequenceListResponse) RawJSON() string { return r.JSON.raw }
func (r *SequenceListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SequenceListResponseData struct {
	// Unique identifier of the sequence.
	ID string `json:"id" api:"required"`
	// Channels the published steps send on; empty before the first publish.
	//
	// Any of "EMAIL", "LINKEDIN".
	Channels []string `json:"channels" api:"required"`
	// ISO 8601 creation timestamp.
	CreatedAt string `json:"createdAt" api:"required"`
	// Published configuration, or null before the first publish.
	Definition SequenceListResponseDataDefinition `json:"definition" api:"required"`
	// When the sequence was first activated.
	FirstActivatedAt string `json:"firstActivatedAt" api:"required"`
	// When the sequence was most recently activated.
	LastActivatedAt string `json:"lastActivatedAt" api:"required"`
	// Name of the sequence.
	Name string `json:"name" api:"required"`
	// Always `sequence`.
	//
	// Any of "sequence".
	ObjectType string `json:"objectType" api:"required"`
	// Extra members who may approve this sequence’s `FOR_REVIEW` messages before they
	// send. The sending member and the sequence’s creator always review, whether or
	// not they are listed. A change applies to sends once the draft is published. The
	// published list; empty before the first publish. Lists active members only.
	Reviewers []SequenceListResponseDataReviewer `json:"reviewers" api:"required"`
	// The senders this sequence sends from; new contacts rotate across them. A change
	// applies to sends once the draft is published. An archived sender stays listed
	// but never sends; any change to the list drops it.
	Senders []SequenceListResponseDataSender `json:"senders" api:"required"`
	// `DRAFT` has not been published; `ACTIVE` is enabled for sending; `DEACTIVATED`
	// is published but paused.
	//
	// Any of "DRAFT", "ACTIVE", "DEACTIVATED".
	Status string `json:"status" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID               respjson.Field
		Channels         respjson.Field
		CreatedAt        respjson.Field
		Definition       respjson.Field
		FirstActivatedAt respjson.Field
		LastActivatedAt  respjson.Field
		Name             respjson.Field
		ObjectType       respjson.Field
		Reviewers        respjson.Field
		Senders          respjson.Field
		Status           respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SequenceListResponseData) RawJSON() string { return r.JSON.raw }
func (r *SequenceListResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Published configuration, or null before the first publish.
type SequenceListResponseDataDefinition struct {
	// What this version of the sequence is for.
	Description string `json:"description" api:"required"`
	// Messaging inputs that frame every step.
	Messaging SequenceListResponseDataDefinitionMessaging `json:"messaging" api:"required"`
	// How the whole sequence paces each step type it sends, keyed by sequence step
	// type.
	Pacing SequenceListResponseDataDefinitionPacing `json:"pacing" api:"required"`
	// Send priority relative to other sequences.
	//
	// Any of "HIGH", "NORMAL".
	Priority string `json:"priority" api:"required"`
	// When the sequence may send.
	SendWindow SequenceListResponseDataDefinitionSendWindow `json:"sendWindow" api:"required"`
	// Email, LinkedIn and wait steps in execution order; call steps are not exposed in
	// this release.
	Steps []SequenceListResponseDataDefinitionStepUnion `json:"steps" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Description respjson.Field
		Messaging   respjson.Field
		Pacing      respjson.Field
		Priority    respjson.Field
		SendWindow  respjson.Field
		Steps       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SequenceListResponseDataDefinition) RawJSON() string { return r.JSON.raw }
func (r *SequenceListResponseDataDefinition) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Messaging inputs that frame every step.
type SequenceListResponseDataDefinitionMessaging struct {
	// What the recipient is asked to do.
	CallToAction string `json:"callToAction" api:"required"`
	// Who the sequence targets.
	Persona string `json:"persona" api:"required"`
	// Why the outreach is timely.
	Trigger string `json:"trigger" api:"required"`
	// The value proposition.
	ValueProp string `json:"valueProp" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CallToAction respjson.Field
		Persona      respjson.Field
		Trigger      respjson.Field
		ValueProp    respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SequenceListResponseDataDefinitionMessaging) RawJSON() string { return r.JSON.raw }
func (r *SequenceListResponseDataDefinitionMessaging) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// How the whole sequence paces each step type it sends, keyed by sequence step
// type.
type SequenceListResponseDataDefinitionPacing struct {
	// Email steps.
	Email SequenceListResponseDataDefinitionPacingEmail `json:"EMAIL" api:"required"`
	// LinkedIn connection requests.
	LinkedinInvite SequenceListResponseDataDefinitionPacingLinkedinInvite `json:"LINKEDIN_INVITE" api:"required"`
	// LinkedIn messages.
	LinkedinMessage SequenceListResponseDataDefinitionPacingLinkedinMessage `json:"LINKEDIN_MESSAGE" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Email           respjson.Field
		LinkedinInvite  respjson.Field
		LinkedinMessage respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SequenceListResponseDataDefinitionPacing) RawJSON() string { return r.JSON.raw }
func (r *SequenceListResponseDataDefinitionPacing) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Email steps.
type SequenceListResponseDataDefinitionPacingEmail struct {
	// Most sends of this step type across the whole sequence in any rolling 24 hours,
	// or null when the sequence sets no daily cap. Workspace and sender limits still
	// apply.
	DailyLimit int64 `json:"dailyLimit" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DailyLimit  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SequenceListResponseDataDefinitionPacingEmail) RawJSON() string { return r.JSON.raw }
func (r *SequenceListResponseDataDefinitionPacingEmail) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// LinkedIn connection requests.
type SequenceListResponseDataDefinitionPacingLinkedinInvite struct {
	// Most sends of this step type across the whole sequence in any rolling 24 hours,
	// or null when the sequence sets no daily cap. Workspace and sender limits still
	// apply.
	DailyLimit int64 `json:"dailyLimit" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DailyLimit  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SequenceListResponseDataDefinitionPacingLinkedinInvite) RawJSON() string { return r.JSON.raw }
func (r *SequenceListResponseDataDefinitionPacingLinkedinInvite) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// LinkedIn messages.
type SequenceListResponseDataDefinitionPacingLinkedinMessage struct {
	// Most sends of this step type across the whole sequence in any rolling 24 hours,
	// or null when the sequence sets no daily cap. Workspace and sender limits still
	// apply.
	DailyLimit int64 `json:"dailyLimit" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DailyLimit  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SequenceListResponseDataDefinitionPacingLinkedinMessage) RawJSON() string { return r.JSON.raw }
func (r *SequenceListResponseDataDefinitionPacingLinkedinMessage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// When the sequence may send.
type SequenceListResponseDataDefinitionSendWindow struct {
	// Days sends may go out, 0 = Sunday.
	DaysOfWeek []int64 `json:"daysOfWeek" api:"required"`
	// Local hour sends stop, exclusive (1–24).
	EndHour int64 `json:"endHour" api:"required"`
	// First local hour sends may go out (0–23).
	StartHour int64 `json:"startHour" api:"required"`
	// IANA timezone of the window.
	Timezone string `json:"timezone" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DaysOfWeek  respjson.Field
		EndHour     respjson.Field
		StartHour   respjson.Field
		Timezone    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SequenceListResponseDataDefinitionSendWindow) RawJSON() string { return r.JSON.raw }
func (r *SequenceListResponseDataDefinitionSendWindow) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// SequenceListResponseDataDefinitionStepUnion contains all possible properties and
// values from [SequenceListResponseDataDefinitionStepObject],
// [SequenceListResponseDataDefinitionStepObject2],
// [SequenceListResponseDataDefinitionStepObject3].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type SequenceListResponseDataDefinitionStepUnion struct {
	ID           string `json:"id"`
	DeliveryMode string `json:"deliveryMode"`
	Name         string `json:"name"`
	Prompt       string `json:"prompt"`
	Type         string `json:"type"`
	// This field is from variant [SequenceListResponseDataDefinitionStepObject2].
	IncludeNote bool `json:"includeNote"`
	// This field is from variant [SequenceListResponseDataDefinitionStepObject3].
	Duration int64 `json:"duration"`
	// This field is from variant [SequenceListResponseDataDefinitionStepObject3].
	Unit string `json:"unit"`
	JSON struct {
		ID           respjson.Field
		DeliveryMode respjson.Field
		Name         respjson.Field
		Prompt       respjson.Field
		Type         respjson.Field
		IncludeNote  respjson.Field
		Duration     respjson.Field
		Unit         respjson.Field
		raw          string
	} `json:"-"`
}

func (u SequenceListResponseDataDefinitionStepUnion) AsSequenceListResponseDataDefinitionStepObject() (v SequenceListResponseDataDefinitionStepObject) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u SequenceListResponseDataDefinitionStepUnion) AsSequenceListResponseDataDefinitionStepObject2() (v SequenceListResponseDataDefinitionStepObject2) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u SequenceListResponseDataDefinitionStepUnion) AsSequenceListResponseDataDefinitionStepObject3() (v SequenceListResponseDataDefinitionStepObject3) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u SequenceListResponseDataDefinitionStepUnion) RawJSON() string { return u.JSON.raw }

func (r *SequenceListResponseDataDefinitionStepUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SequenceListResponseDataDefinitionStepObject struct {
	// Step identifier, referenced by enrollment `currentStepId` and message
	// `sequenceStepId`.
	ID string `json:"id" api:"required"`
	// `AUTO` sends without manual approval, subject to the send window and limits;
	// `FOR_REVIEW` requires a person to approve the draft before sending.
	//
	// Any of "AUTO", "FOR_REVIEW".
	DeliveryMode string `json:"deliveryMode" api:"required"`
	// Display name of the step.
	Name string `json:"name" api:"required"`
	// Instructions the agent follows when composing this message.
	Prompt string `json:"prompt" api:"required"`
	// Step type.
	//
	// Any of "EMAIL", "LINKEDIN_MESSAGE".
	Type string `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		DeliveryMode respjson.Field
		Name         respjson.Field
		Prompt       respjson.Field
		Type         respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SequenceListResponseDataDefinitionStepObject) RawJSON() string { return r.JSON.raw }
func (r *SequenceListResponseDataDefinitionStepObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SequenceListResponseDataDefinitionStepObject2 struct {
	// Step identifier, referenced by enrollment `currentStepId` and message
	// `sequenceStepId`.
	ID string `json:"id" api:"required"`
	// `AUTO` sends without manual approval, subject to the send window and limits;
	// `FOR_REVIEW` requires a person to approve the draft before sending.
	//
	// Any of "AUTO", "FOR_REVIEW".
	DeliveryMode string `json:"deliveryMode" api:"required"`
	// Whether the connection request carries a note. A step authored before this
	// setting existed reads as true.
	IncludeNote bool `json:"includeNote" api:"required"`
	// Display name of the step.
	Name string `json:"name" api:"required"`
	// Instructions the agent follows when composing this message.
	Prompt string `json:"prompt" api:"required"`
	// Step type.
	//
	// Any of "LINKEDIN_INVITE".
	Type string `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		DeliveryMode respjson.Field
		IncludeNote  respjson.Field
		Name         respjson.Field
		Prompt       respjson.Field
		Type         respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SequenceListResponseDataDefinitionStepObject2) RawJSON() string { return r.JSON.raw }
func (r *SequenceListResponseDataDefinitionStepObject2) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SequenceListResponseDataDefinitionStepObject3 struct {
	// Step identifier.
	ID string `json:"id" api:"required"`
	// Length of the wait.
	Duration int64 `json:"duration" api:"required"`
	// `WAIT` pauses for a fixed time; `WAIT_FOR_CONNECTION` waits up to that long for
	// a LinkedIn invite to be accepted.
	//
	// Any of "WAIT", "WAIT_FOR_CONNECTION".
	Type string `json:"type" api:"required"`
	// Unit of `duration`.
	//
	// Any of "DAY", "MINUTE".
	Unit string `json:"unit" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Duration    respjson.Field
		Type        respjson.Field
		Unit        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SequenceListResponseDataDefinitionStepObject3) RawJSON() string { return r.JSON.raw }
func (r *SequenceListResponseDataDefinitionStepObject3) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SequenceListResponseDataReviewer struct {
	// ID of the member.
	ID string `json:"id" api:"required"`
	// Always `member`.
	//
	// Any of "member".
	ObjectType string `json:"objectType" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		ObjectType  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SequenceListResponseDataReviewer) RawJSON() string { return r.JSON.raw }
func (r *SequenceListResponseDataReviewer) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SequenceListResponseDataSender struct {
	// ID of the sender.
	ID string `json:"id" api:"required"`
	// Always `sender`.
	//
	// Any of "sender".
	ObjectType string `json:"objectType" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		ObjectType  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SequenceListResponseDataSender) RawJSON() string { return r.JSON.raw }
func (r *SequenceListResponseDataSender) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SequenceListResponseObject string

const (
	SequenceListResponseObjectList SequenceListResponseObject = "list"
)

type SequencePublishDraftResponse struct {
	// Unique identifier of the sequence.
	ID string `json:"id" api:"required"`
	// Channels the published steps send on; empty before the first publish.
	//
	// Any of "EMAIL", "LINKEDIN".
	Channels []string `json:"channels" api:"required"`
	// ISO 8601 creation timestamp.
	CreatedAt string `json:"createdAt" api:"required"`
	// Published configuration, or null before the first publish.
	Definition SequencePublishDraftResponseDefinition `json:"definition" api:"required"`
	// When the sequence was first activated.
	FirstActivatedAt string `json:"firstActivatedAt" api:"required"`
	// When the sequence was most recently activated.
	LastActivatedAt string `json:"lastActivatedAt" api:"required"`
	// Name of the sequence.
	Name string `json:"name" api:"required"`
	// Always `sequence`.
	//
	// Any of "sequence".
	ObjectType SequencePublishDraftResponseObjectType `json:"objectType" api:"required"`
	// Extra members who may approve this sequence’s `FOR_REVIEW` messages before they
	// send. The sending member and the sequence’s creator always review, whether or
	// not they are listed. A change applies to sends once the draft is published. The
	// published list; empty before the first publish. Lists active members only.
	Reviewers []SequencePublishDraftResponseReviewer `json:"reviewers" api:"required"`
	// The senders this sequence sends from; new contacts rotate across them. A change
	// applies to sends once the draft is published. An archived sender stays listed
	// but never sends; any change to the list drops it.
	Senders []SequencePublishDraftResponseSender `json:"senders" api:"required"`
	// `DRAFT` has not been published; `ACTIVE` is enabled for sending; `DEACTIVATED`
	// is published but paused.
	//
	// Any of "DRAFT", "ACTIVE", "DEACTIVATED".
	Status SequencePublishDraftResponseStatus `json:"status" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID               respjson.Field
		Channels         respjson.Field
		CreatedAt        respjson.Field
		Definition       respjson.Field
		FirstActivatedAt respjson.Field
		LastActivatedAt  respjson.Field
		Name             respjson.Field
		ObjectType       respjson.Field
		Reviewers        respjson.Field
		Senders          respjson.Field
		Status           respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SequencePublishDraftResponse) RawJSON() string { return r.JSON.raw }
func (r *SequencePublishDraftResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Published configuration, or null before the first publish.
type SequencePublishDraftResponseDefinition struct {
	// What this version of the sequence is for.
	Description string `json:"description" api:"required"`
	// Messaging inputs that frame every step.
	Messaging SequencePublishDraftResponseDefinitionMessaging `json:"messaging" api:"required"`
	// How the whole sequence paces each step type it sends, keyed by sequence step
	// type.
	Pacing SequencePublishDraftResponseDefinitionPacing `json:"pacing" api:"required"`
	// Send priority relative to other sequences.
	//
	// Any of "HIGH", "NORMAL".
	Priority string `json:"priority" api:"required"`
	// When the sequence may send.
	SendWindow SequencePublishDraftResponseDefinitionSendWindow `json:"sendWindow" api:"required"`
	// Email, LinkedIn and wait steps in execution order; call steps are not exposed in
	// this release.
	Steps []SequencePublishDraftResponseDefinitionStepUnion `json:"steps" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Description respjson.Field
		Messaging   respjson.Field
		Pacing      respjson.Field
		Priority    respjson.Field
		SendWindow  respjson.Field
		Steps       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SequencePublishDraftResponseDefinition) RawJSON() string { return r.JSON.raw }
func (r *SequencePublishDraftResponseDefinition) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Messaging inputs that frame every step.
type SequencePublishDraftResponseDefinitionMessaging struct {
	// What the recipient is asked to do.
	CallToAction string `json:"callToAction" api:"required"`
	// Who the sequence targets.
	Persona string `json:"persona" api:"required"`
	// Why the outreach is timely.
	Trigger string `json:"trigger" api:"required"`
	// The value proposition.
	ValueProp string `json:"valueProp" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CallToAction respjson.Field
		Persona      respjson.Field
		Trigger      respjson.Field
		ValueProp    respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SequencePublishDraftResponseDefinitionMessaging) RawJSON() string { return r.JSON.raw }
func (r *SequencePublishDraftResponseDefinitionMessaging) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// How the whole sequence paces each step type it sends, keyed by sequence step
// type.
type SequencePublishDraftResponseDefinitionPacing struct {
	// Email steps.
	Email SequencePublishDraftResponseDefinitionPacingEmail `json:"EMAIL" api:"required"`
	// LinkedIn connection requests.
	LinkedinInvite SequencePublishDraftResponseDefinitionPacingLinkedinInvite `json:"LINKEDIN_INVITE" api:"required"`
	// LinkedIn messages.
	LinkedinMessage SequencePublishDraftResponseDefinitionPacingLinkedinMessage `json:"LINKEDIN_MESSAGE" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Email           respjson.Field
		LinkedinInvite  respjson.Field
		LinkedinMessage respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SequencePublishDraftResponseDefinitionPacing) RawJSON() string { return r.JSON.raw }
func (r *SequencePublishDraftResponseDefinitionPacing) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Email steps.
type SequencePublishDraftResponseDefinitionPacingEmail struct {
	// Most sends of this step type across the whole sequence in any rolling 24 hours,
	// or null when the sequence sets no daily cap. Workspace and sender limits still
	// apply.
	DailyLimit int64 `json:"dailyLimit" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DailyLimit  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SequencePublishDraftResponseDefinitionPacingEmail) RawJSON() string { return r.JSON.raw }
func (r *SequencePublishDraftResponseDefinitionPacingEmail) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// LinkedIn connection requests.
type SequencePublishDraftResponseDefinitionPacingLinkedinInvite struct {
	// Most sends of this step type across the whole sequence in any rolling 24 hours,
	// or null when the sequence sets no daily cap. Workspace and sender limits still
	// apply.
	DailyLimit int64 `json:"dailyLimit" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DailyLimit  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SequencePublishDraftResponseDefinitionPacingLinkedinInvite) RawJSON() string {
	return r.JSON.raw
}
func (r *SequencePublishDraftResponseDefinitionPacingLinkedinInvite) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// LinkedIn messages.
type SequencePublishDraftResponseDefinitionPacingLinkedinMessage struct {
	// Most sends of this step type across the whole sequence in any rolling 24 hours,
	// or null when the sequence sets no daily cap. Workspace and sender limits still
	// apply.
	DailyLimit int64 `json:"dailyLimit" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DailyLimit  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SequencePublishDraftResponseDefinitionPacingLinkedinMessage) RawJSON() string {
	return r.JSON.raw
}
func (r *SequencePublishDraftResponseDefinitionPacingLinkedinMessage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// When the sequence may send.
type SequencePublishDraftResponseDefinitionSendWindow struct {
	// Days sends may go out, 0 = Sunday.
	DaysOfWeek []int64 `json:"daysOfWeek" api:"required"`
	// Local hour sends stop, exclusive (1–24).
	EndHour int64 `json:"endHour" api:"required"`
	// First local hour sends may go out (0–23).
	StartHour int64 `json:"startHour" api:"required"`
	// IANA timezone of the window.
	Timezone string `json:"timezone" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DaysOfWeek  respjson.Field
		EndHour     respjson.Field
		StartHour   respjson.Field
		Timezone    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SequencePublishDraftResponseDefinitionSendWindow) RawJSON() string { return r.JSON.raw }
func (r *SequencePublishDraftResponseDefinitionSendWindow) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// SequencePublishDraftResponseDefinitionStepUnion contains all possible properties
// and values from [SequencePublishDraftResponseDefinitionStepObject],
// [SequencePublishDraftResponseDefinitionStepObject2],
// [SequencePublishDraftResponseDefinitionStepObject3].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type SequencePublishDraftResponseDefinitionStepUnion struct {
	ID           string `json:"id"`
	DeliveryMode string `json:"deliveryMode"`
	Name         string `json:"name"`
	Prompt       string `json:"prompt"`
	Type         string `json:"type"`
	// This field is from variant [SequencePublishDraftResponseDefinitionStepObject2].
	IncludeNote bool `json:"includeNote"`
	// This field is from variant [SequencePublishDraftResponseDefinitionStepObject3].
	Duration int64 `json:"duration"`
	// This field is from variant [SequencePublishDraftResponseDefinitionStepObject3].
	Unit string `json:"unit"`
	JSON struct {
		ID           respjson.Field
		DeliveryMode respjson.Field
		Name         respjson.Field
		Prompt       respjson.Field
		Type         respjson.Field
		IncludeNote  respjson.Field
		Duration     respjson.Field
		Unit         respjson.Field
		raw          string
	} `json:"-"`
}

func (u SequencePublishDraftResponseDefinitionStepUnion) AsSequencePublishDraftResponseDefinitionStepObject() (v SequencePublishDraftResponseDefinitionStepObject) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u SequencePublishDraftResponseDefinitionStepUnion) AsSequencePublishDraftResponseDefinitionStepObject2() (v SequencePublishDraftResponseDefinitionStepObject2) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u SequencePublishDraftResponseDefinitionStepUnion) AsSequencePublishDraftResponseDefinitionStepObject3() (v SequencePublishDraftResponseDefinitionStepObject3) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u SequencePublishDraftResponseDefinitionStepUnion) RawJSON() string { return u.JSON.raw }

func (r *SequencePublishDraftResponseDefinitionStepUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SequencePublishDraftResponseDefinitionStepObject struct {
	// Step identifier, referenced by enrollment `currentStepId` and message
	// `sequenceStepId`.
	ID string `json:"id" api:"required"`
	// `AUTO` sends without manual approval, subject to the send window and limits;
	// `FOR_REVIEW` requires a person to approve the draft before sending.
	//
	// Any of "AUTO", "FOR_REVIEW".
	DeliveryMode string `json:"deliveryMode" api:"required"`
	// Display name of the step.
	Name string `json:"name" api:"required"`
	// Instructions the agent follows when composing this message.
	Prompt string `json:"prompt" api:"required"`
	// Step type.
	//
	// Any of "EMAIL", "LINKEDIN_MESSAGE".
	Type string `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		DeliveryMode respjson.Field
		Name         respjson.Field
		Prompt       respjson.Field
		Type         respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SequencePublishDraftResponseDefinitionStepObject) RawJSON() string { return r.JSON.raw }
func (r *SequencePublishDraftResponseDefinitionStepObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SequencePublishDraftResponseDefinitionStepObject2 struct {
	// Step identifier, referenced by enrollment `currentStepId` and message
	// `sequenceStepId`.
	ID string `json:"id" api:"required"`
	// `AUTO` sends without manual approval, subject to the send window and limits;
	// `FOR_REVIEW` requires a person to approve the draft before sending.
	//
	// Any of "AUTO", "FOR_REVIEW".
	DeliveryMode string `json:"deliveryMode" api:"required"`
	// Whether the connection request carries a note. A step authored before this
	// setting existed reads as true.
	IncludeNote bool `json:"includeNote" api:"required"`
	// Display name of the step.
	Name string `json:"name" api:"required"`
	// Instructions the agent follows when composing this message.
	Prompt string `json:"prompt" api:"required"`
	// Step type.
	//
	// Any of "LINKEDIN_INVITE".
	Type string `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		DeliveryMode respjson.Field
		IncludeNote  respjson.Field
		Name         respjson.Field
		Prompt       respjson.Field
		Type         respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SequencePublishDraftResponseDefinitionStepObject2) RawJSON() string { return r.JSON.raw }
func (r *SequencePublishDraftResponseDefinitionStepObject2) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SequencePublishDraftResponseDefinitionStepObject3 struct {
	// Step identifier.
	ID string `json:"id" api:"required"`
	// Length of the wait.
	Duration int64 `json:"duration" api:"required"`
	// `WAIT` pauses for a fixed time; `WAIT_FOR_CONNECTION` waits up to that long for
	// a LinkedIn invite to be accepted.
	//
	// Any of "WAIT", "WAIT_FOR_CONNECTION".
	Type string `json:"type" api:"required"`
	// Unit of `duration`.
	//
	// Any of "DAY", "MINUTE".
	Unit string `json:"unit" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Duration    respjson.Field
		Type        respjson.Field
		Unit        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SequencePublishDraftResponseDefinitionStepObject3) RawJSON() string { return r.JSON.raw }
func (r *SequencePublishDraftResponseDefinitionStepObject3) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Always `sequence`.
type SequencePublishDraftResponseObjectType string

const (
	SequencePublishDraftResponseObjectTypeSequence SequencePublishDraftResponseObjectType = "sequence"
)

type SequencePublishDraftResponseReviewer struct {
	// ID of the member.
	ID string `json:"id" api:"required"`
	// Always `member`.
	//
	// Any of "member".
	ObjectType string `json:"objectType" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		ObjectType  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SequencePublishDraftResponseReviewer) RawJSON() string { return r.JSON.raw }
func (r *SequencePublishDraftResponseReviewer) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SequencePublishDraftResponseSender struct {
	// ID of the sender.
	ID string `json:"id" api:"required"`
	// Always `sender`.
	//
	// Any of "sender".
	ObjectType string `json:"objectType" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		ObjectType  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SequencePublishDraftResponseSender) RawJSON() string { return r.JSON.raw }
func (r *SequencePublishDraftResponseSender) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// `DRAFT` has not been published; `ACTIVE` is enabled for sending; `DEACTIVATED`
// is published but paused.
type SequencePublishDraftResponseStatus string

const (
	SequencePublishDraftResponseStatusDraft       SequencePublishDraftResponseStatus = "DRAFT"
	SequencePublishDraftResponseStatusActive      SequencePublishDraftResponseStatus = "ACTIVE"
	SequencePublishDraftResponseStatusDeactivated SequencePublishDraftResponseStatus = "DEACTIVATED"
)

type SequenceRetrieveDraftResponse struct {
	// Channels the draft's steps send on.
	//
	// Any of "EMAIL", "LINKEDIN".
	Channels []string `json:"channels" api:"required"`
	// The draft configuration, in the same shape as a published definition.
	Definition SequenceRetrieveDraftResponseDefinition `json:"definition" api:"required"`
	// Name the sequence takes when the draft is published.
	Name string `json:"name" api:"required"`
	// What publishing this draft would change for contacts already in the sequence.
	// Before the first publish, only proposed senders it would leave off because they
	// were archived. Empty when nothing is affected.
	PublishWarnings []SequenceRetrieveDraftResponsePublishWarning `json:"publishWarnings" api:"required"`
	// Extra members who may approve this sequence’s `FOR_REVIEW` messages before they
	// send. The sending member and the sequence’s creator always review, whether or
	// not they are listed. A change applies to sends once the draft is published.
	// Lists active members only.
	Reviewers []SequenceRetrieveDraftResponseReviewer `json:"reviewers" api:"required"`
	// The senders this sequence sends from; new contacts rotate across them. A change
	// applies to sends once the draft is published. The list this draft publishes: the
	// live list until a save changes it. An archived sender stays listed but never
	// sends; any change to the list drops it.
	Senders []SequenceRetrieveDraftResponseSender `json:"senders" api:"required"`
	// The sequence this draft belongs to.
	SequenceID string `json:"sequenceId" api:"required"`
	// When the draft last changed.
	UpdatedAt string `json:"updatedAt" api:"required"`
	// Opaque token for the current state of the draft. Send it as `expectedVersion` on
	// the next save, publish or discard; it changes whenever anyone edits the draft.
	Version string `json:"version" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Channels        respjson.Field
		Definition      respjson.Field
		Name            respjson.Field
		PublishWarnings respjson.Field
		Reviewers       respjson.Field
		Senders         respjson.Field
		SequenceID      respjson.Field
		UpdatedAt       respjson.Field
		Version         respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SequenceRetrieveDraftResponse) RawJSON() string { return r.JSON.raw }
func (r *SequenceRetrieveDraftResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The draft configuration, in the same shape as a published definition.
type SequenceRetrieveDraftResponseDefinition struct {
	// What this version of the sequence is for.
	Description string `json:"description" api:"required"`
	// Messaging inputs that frame every step.
	Messaging SequenceRetrieveDraftResponseDefinitionMessaging `json:"messaging" api:"required"`
	// How the whole sequence paces each step type it sends, keyed by sequence step
	// type.
	Pacing SequenceRetrieveDraftResponseDefinitionPacing `json:"pacing" api:"required"`
	// Send priority relative to other sequences.
	//
	// Any of "HIGH", "NORMAL".
	Priority string `json:"priority" api:"required"`
	// When the sequence may send.
	SendWindow SequenceRetrieveDraftResponseDefinitionSendWindow `json:"sendWindow" api:"required"`
	// Email, LinkedIn and wait steps in execution order; call steps are not exposed in
	// this release.
	Steps []SequenceRetrieveDraftResponseDefinitionStepUnion `json:"steps" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Description respjson.Field
		Messaging   respjson.Field
		Pacing      respjson.Field
		Priority    respjson.Field
		SendWindow  respjson.Field
		Steps       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SequenceRetrieveDraftResponseDefinition) RawJSON() string { return r.JSON.raw }
func (r *SequenceRetrieveDraftResponseDefinition) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Messaging inputs that frame every step.
type SequenceRetrieveDraftResponseDefinitionMessaging struct {
	// What the recipient is asked to do.
	CallToAction string `json:"callToAction" api:"required"`
	// Who the sequence targets.
	Persona string `json:"persona" api:"required"`
	// Why the outreach is timely.
	Trigger string `json:"trigger" api:"required"`
	// The value proposition.
	ValueProp string `json:"valueProp" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CallToAction respjson.Field
		Persona      respjson.Field
		Trigger      respjson.Field
		ValueProp    respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SequenceRetrieveDraftResponseDefinitionMessaging) RawJSON() string { return r.JSON.raw }
func (r *SequenceRetrieveDraftResponseDefinitionMessaging) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// How the whole sequence paces each step type it sends, keyed by sequence step
// type.
type SequenceRetrieveDraftResponseDefinitionPacing struct {
	// Email steps.
	Email SequenceRetrieveDraftResponseDefinitionPacingEmail `json:"EMAIL" api:"required"`
	// LinkedIn connection requests.
	LinkedinInvite SequenceRetrieveDraftResponseDefinitionPacingLinkedinInvite `json:"LINKEDIN_INVITE" api:"required"`
	// LinkedIn messages.
	LinkedinMessage SequenceRetrieveDraftResponseDefinitionPacingLinkedinMessage `json:"LINKEDIN_MESSAGE" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Email           respjson.Field
		LinkedinInvite  respjson.Field
		LinkedinMessage respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SequenceRetrieveDraftResponseDefinitionPacing) RawJSON() string { return r.JSON.raw }
func (r *SequenceRetrieveDraftResponseDefinitionPacing) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Email steps.
type SequenceRetrieveDraftResponseDefinitionPacingEmail struct {
	// Most sends of this step type across the whole sequence in any rolling 24 hours,
	// or null when the sequence sets no daily cap. Workspace and sender limits still
	// apply.
	DailyLimit int64 `json:"dailyLimit" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DailyLimit  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SequenceRetrieveDraftResponseDefinitionPacingEmail) RawJSON() string { return r.JSON.raw }
func (r *SequenceRetrieveDraftResponseDefinitionPacingEmail) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// LinkedIn connection requests.
type SequenceRetrieveDraftResponseDefinitionPacingLinkedinInvite struct {
	// Most sends of this step type across the whole sequence in any rolling 24 hours,
	// or null when the sequence sets no daily cap. Workspace and sender limits still
	// apply.
	DailyLimit int64 `json:"dailyLimit" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DailyLimit  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SequenceRetrieveDraftResponseDefinitionPacingLinkedinInvite) RawJSON() string {
	return r.JSON.raw
}
func (r *SequenceRetrieveDraftResponseDefinitionPacingLinkedinInvite) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// LinkedIn messages.
type SequenceRetrieveDraftResponseDefinitionPacingLinkedinMessage struct {
	// Most sends of this step type across the whole sequence in any rolling 24 hours,
	// or null when the sequence sets no daily cap. Workspace and sender limits still
	// apply.
	DailyLimit int64 `json:"dailyLimit" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DailyLimit  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SequenceRetrieveDraftResponseDefinitionPacingLinkedinMessage) RawJSON() string {
	return r.JSON.raw
}
func (r *SequenceRetrieveDraftResponseDefinitionPacingLinkedinMessage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// When the sequence may send.
type SequenceRetrieveDraftResponseDefinitionSendWindow struct {
	// Days sends may go out, 0 = Sunday.
	DaysOfWeek []int64 `json:"daysOfWeek" api:"required"`
	// Local hour sends stop, exclusive (1–24).
	EndHour int64 `json:"endHour" api:"required"`
	// First local hour sends may go out (0–23).
	StartHour int64 `json:"startHour" api:"required"`
	// IANA timezone of the window.
	Timezone string `json:"timezone" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DaysOfWeek  respjson.Field
		EndHour     respjson.Field
		StartHour   respjson.Field
		Timezone    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SequenceRetrieveDraftResponseDefinitionSendWindow) RawJSON() string { return r.JSON.raw }
func (r *SequenceRetrieveDraftResponseDefinitionSendWindow) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// SequenceRetrieveDraftResponseDefinitionStepUnion contains all possible
// properties and values from [SequenceRetrieveDraftResponseDefinitionStepObject],
// [SequenceRetrieveDraftResponseDefinitionStepObject2],
// [SequenceRetrieveDraftResponseDefinitionStepObject3].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type SequenceRetrieveDraftResponseDefinitionStepUnion struct {
	ID           string `json:"id"`
	DeliveryMode string `json:"deliveryMode"`
	Name         string `json:"name"`
	Prompt       string `json:"prompt"`
	Type         string `json:"type"`
	// This field is from variant [SequenceRetrieveDraftResponseDefinitionStepObject2].
	IncludeNote bool `json:"includeNote"`
	// This field is from variant [SequenceRetrieveDraftResponseDefinitionStepObject3].
	Duration int64 `json:"duration"`
	// This field is from variant [SequenceRetrieveDraftResponseDefinitionStepObject3].
	Unit string `json:"unit"`
	JSON struct {
		ID           respjson.Field
		DeliveryMode respjson.Field
		Name         respjson.Field
		Prompt       respjson.Field
		Type         respjson.Field
		IncludeNote  respjson.Field
		Duration     respjson.Field
		Unit         respjson.Field
		raw          string
	} `json:"-"`
}

func (u SequenceRetrieveDraftResponseDefinitionStepUnion) AsSequenceRetrieveDraftResponseDefinitionStepObject() (v SequenceRetrieveDraftResponseDefinitionStepObject) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u SequenceRetrieveDraftResponseDefinitionStepUnion) AsSequenceRetrieveDraftResponseDefinitionStepObject2() (v SequenceRetrieveDraftResponseDefinitionStepObject2) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u SequenceRetrieveDraftResponseDefinitionStepUnion) AsSequenceRetrieveDraftResponseDefinitionStepObject3() (v SequenceRetrieveDraftResponseDefinitionStepObject3) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u SequenceRetrieveDraftResponseDefinitionStepUnion) RawJSON() string { return u.JSON.raw }

func (r *SequenceRetrieveDraftResponseDefinitionStepUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SequenceRetrieveDraftResponseDefinitionStepObject struct {
	// Step identifier, referenced by enrollment `currentStepId` and message
	// `sequenceStepId`.
	ID string `json:"id" api:"required"`
	// `AUTO` sends without manual approval, subject to the send window and limits;
	// `FOR_REVIEW` requires a person to approve the draft before sending.
	//
	// Any of "AUTO", "FOR_REVIEW".
	DeliveryMode string `json:"deliveryMode" api:"required"`
	// Display name of the step.
	Name string `json:"name" api:"required"`
	// Instructions the agent follows when composing this message.
	Prompt string `json:"prompt" api:"required"`
	// Step type.
	//
	// Any of "EMAIL", "LINKEDIN_MESSAGE".
	Type string `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		DeliveryMode respjson.Field
		Name         respjson.Field
		Prompt       respjson.Field
		Type         respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SequenceRetrieveDraftResponseDefinitionStepObject) RawJSON() string { return r.JSON.raw }
func (r *SequenceRetrieveDraftResponseDefinitionStepObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SequenceRetrieveDraftResponseDefinitionStepObject2 struct {
	// Step identifier, referenced by enrollment `currentStepId` and message
	// `sequenceStepId`.
	ID string `json:"id" api:"required"`
	// `AUTO` sends without manual approval, subject to the send window and limits;
	// `FOR_REVIEW` requires a person to approve the draft before sending.
	//
	// Any of "AUTO", "FOR_REVIEW".
	DeliveryMode string `json:"deliveryMode" api:"required"`
	// Whether the connection request carries a note. A step authored before this
	// setting existed reads as true.
	IncludeNote bool `json:"includeNote" api:"required"`
	// Display name of the step.
	Name string `json:"name" api:"required"`
	// Instructions the agent follows when composing this message.
	Prompt string `json:"prompt" api:"required"`
	// Step type.
	//
	// Any of "LINKEDIN_INVITE".
	Type string `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		DeliveryMode respjson.Field
		IncludeNote  respjson.Field
		Name         respjson.Field
		Prompt       respjson.Field
		Type         respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SequenceRetrieveDraftResponseDefinitionStepObject2) RawJSON() string { return r.JSON.raw }
func (r *SequenceRetrieveDraftResponseDefinitionStepObject2) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SequenceRetrieveDraftResponseDefinitionStepObject3 struct {
	// Step identifier.
	ID string `json:"id" api:"required"`
	// Length of the wait.
	Duration int64 `json:"duration" api:"required"`
	// `WAIT` pauses for a fixed time; `WAIT_FOR_CONNECTION` waits up to that long for
	// a LinkedIn invite to be accepted.
	//
	// Any of "WAIT", "WAIT_FOR_CONNECTION".
	Type string `json:"type" api:"required"`
	// Unit of `duration`.
	//
	// Any of "DAY", "MINUTE".
	Unit string `json:"unit" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Duration    respjson.Field
		Type        respjson.Field
		Unit        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SequenceRetrieveDraftResponseDefinitionStepObject3) RawJSON() string { return r.JSON.raw }
func (r *SequenceRetrieveDraftResponseDefinitionStepObject3) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SequenceRetrieveDraftResponsePublishWarning struct {
	// What publishing would change.
	//
	// Any of "delay_shortened", "delay_lengthened", "connection_wait_shortened",
	// "connection_wait_lengthened", "send_schedule_changed",
	// "send_configuration_changed", "send_priority_changed", "senders_changed",
	// "email_sender_pool_emptied", "linkedin_account_pool_emptied",
	// "assigned_sender_removed", "delivery_became_automatic", "invite_note_enabled",
	// "invite_note_disabled", "step_inserted", "automatic_send_step_inserted",
	// "step_removed", "step_type_changed", "step_order_changed",
	// "full_identity_rewrite", "low_identity_overlap".
	Code string `json:"code" api:"required"`
	// What would change, in plain language.
	Message string `json:"message" api:"required"`
	// `destructive` when contacts already in the sequence may skip, repeat or receive
	// a step they would not have; `warning` otherwise.
	//
	// Any of "warning", "destructive".
	Severity string `json:"severity" api:"required"`
	// The steps involved; empty for a sequence-wide change.
	StepIDs []string `json:"stepIds" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Code        respjson.Field
		Message     respjson.Field
		Severity    respjson.Field
		StepIDs     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SequenceRetrieveDraftResponsePublishWarning) RawJSON() string { return r.JSON.raw }
func (r *SequenceRetrieveDraftResponsePublishWarning) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SequenceRetrieveDraftResponseReviewer struct {
	// ID of the member.
	ID string `json:"id" api:"required"`
	// Always `member`.
	//
	// Any of "member".
	ObjectType string `json:"objectType" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		ObjectType  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SequenceRetrieveDraftResponseReviewer) RawJSON() string { return r.JSON.raw }
func (r *SequenceRetrieveDraftResponseReviewer) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SequenceRetrieveDraftResponseSender struct {
	// ID of the sender.
	ID string `json:"id" api:"required"`
	// Always `sender`.
	//
	// Any of "sender".
	ObjectType string `json:"objectType" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		ObjectType  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SequenceRetrieveDraftResponseSender) RawJSON() string { return r.JSON.raw }
func (r *SequenceRetrieveDraftResponseSender) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SequenceRetrieveEnrollmentResponse struct {
	// Unique identifier of the enrollment.
	ID string `json:"id" api:"required"`
	// When the enrollment was cancelled, or null if it was not.
	CancelledAt string `json:"cancelledAt" api:"required"`
	// When the enrollment completed, or null if it has not.
	CompletedAt string `json:"completedAt" api:"required"`
	// The enrolled contact.
	ContactID string `json:"contactId" api:"required"`
	// The current step ID from `definition.steps`, or null when no current step
	// exposed by this API can be resolved.
	CurrentStepID string `json:"currentStepId" api:"required"`
	// When the contact was enrolled.
	EnrolledAt string `json:"enrolledAt" api:"required"`
	// Member ID of the person who enrolled the contact, or null when no enrolling user
	// or workspace membership can be resolved.
	EnrolledBy string `json:"enrolledBy" api:"required"`
	// When the activity selected for display happened. A terminal reply or bounce can
	// take precedence over a later open.
	LastActivityAt string `json:"lastActivityAt" api:"required"`
	// The activity selected for display. A terminal reply or bounce can take
	// precedence over a later open.
	//
	// Any of "SENT", "OPENED", "REPLIED", "BOUNCED", "INVITED", "CONNECTED", "SEEN".
	LastActivityKind SequenceRetrieveEnrollmentResponseLastActivityKind `json:"lastActivityKind" api:"required"`
	// Before a send is queued, an estimate from the planned cadence and send window.
	// For a ready queued send, the earliest current eligibility based on its fixed
	// sender, pacing limits, and send window. Randomized gaps, jitter, other queued
	// work, or new sends may push it later. Null while awaiting review, sending,
	// paused, or when eligibility is unknown.
	NextSendAt string `json:"nextSendAt" api:"required"`
	// Unsent email and LinkedIn steps remaining, including connection requests; null
	// on retrieval when a current indexed value is unavailable.
	RemainingSends int64 `json:"remainingSends" api:"required"`
	// When this enrollment first received a reply across email or LinkedIn.
	RepliedAt string `json:"repliedAt" api:"required"`
	// Display status of the enrollment, or null on retrieval when the index is
	// unavailable or has not caught up with the enrollment lifecycle.
	//
	// Any of "ACTIVE", "AWAITING_REVIEW", "NEEDS_ATTENTION", "REPLIED", "COMPLETED",
	// "PAUSED", "CANCELLED", "FAILED".
	Status SequenceRetrieveEnrollmentResponseStatus `json:"status" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID               respjson.Field
		CancelledAt      respjson.Field
		CompletedAt      respjson.Field
		ContactID        respjson.Field
		CurrentStepID    respjson.Field
		EnrolledAt       respjson.Field
		EnrolledBy       respjson.Field
		LastActivityAt   respjson.Field
		LastActivityKind respjson.Field
		NextSendAt       respjson.Field
		RemainingSends   respjson.Field
		RepliedAt        respjson.Field
		Status           respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SequenceRetrieveEnrollmentResponse) RawJSON() string { return r.JSON.raw }
func (r *SequenceRetrieveEnrollmentResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The activity selected for display. A terminal reply or bounce can take
// precedence over a later open.
type SequenceRetrieveEnrollmentResponseLastActivityKind string

const (
	SequenceRetrieveEnrollmentResponseLastActivityKindSent      SequenceRetrieveEnrollmentResponseLastActivityKind = "SENT"
	SequenceRetrieveEnrollmentResponseLastActivityKindOpened    SequenceRetrieveEnrollmentResponseLastActivityKind = "OPENED"
	SequenceRetrieveEnrollmentResponseLastActivityKindReplied   SequenceRetrieveEnrollmentResponseLastActivityKind = "REPLIED"
	SequenceRetrieveEnrollmentResponseLastActivityKindBounced   SequenceRetrieveEnrollmentResponseLastActivityKind = "BOUNCED"
	SequenceRetrieveEnrollmentResponseLastActivityKindInvited   SequenceRetrieveEnrollmentResponseLastActivityKind = "INVITED"
	SequenceRetrieveEnrollmentResponseLastActivityKindConnected SequenceRetrieveEnrollmentResponseLastActivityKind = "CONNECTED"
	SequenceRetrieveEnrollmentResponseLastActivityKindSeen      SequenceRetrieveEnrollmentResponseLastActivityKind = "SEEN"
)

// Display status of the enrollment, or null on retrieval when the index is
// unavailable or has not caught up with the enrollment lifecycle.
type SequenceRetrieveEnrollmentResponseStatus string

const (
	SequenceRetrieveEnrollmentResponseStatusActive         SequenceRetrieveEnrollmentResponseStatus = "ACTIVE"
	SequenceRetrieveEnrollmentResponseStatusAwaitingReview SequenceRetrieveEnrollmentResponseStatus = "AWAITING_REVIEW"
	SequenceRetrieveEnrollmentResponseStatusNeedsAttention SequenceRetrieveEnrollmentResponseStatus = "NEEDS_ATTENTION"
	SequenceRetrieveEnrollmentResponseStatusReplied        SequenceRetrieveEnrollmentResponseStatus = "REPLIED"
	SequenceRetrieveEnrollmentResponseStatusCompleted      SequenceRetrieveEnrollmentResponseStatus = "COMPLETED"
	SequenceRetrieveEnrollmentResponseStatusPaused         SequenceRetrieveEnrollmentResponseStatus = "PAUSED"
	SequenceRetrieveEnrollmentResponseStatusCancelled      SequenceRetrieveEnrollmentResponseStatus = "CANCELLED"
	SequenceRetrieveEnrollmentResponseStatusFailed         SequenceRetrieveEnrollmentResponseStatus = "FAILED"
)

// SequenceRetrieveMessageResponseUnion contains all possible properties and values
// from [SequenceRetrieveMessageResponseObject],
// [SequenceRetrieveMessageResponseObject2].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type SequenceRetrieveMessageResponseUnion struct {
	ID string `json:"id"`
	// This field is from variant [SequenceRetrieveMessageResponseObject].
	AcceptedAt string `json:"acceptedAt"`
	// This field is from variant [SequenceRetrieveMessageResponseObject].
	BouncedAt string `json:"bouncedAt"`
	Channel   string `json:"channel"`
	ContactID string `json:"contactId"`
	// This field is a union of [SequenceRetrieveMessageResponseObjectContent],
	// [SequenceRetrieveMessageResponseObject2Content]
	Content      SequenceRetrieveMessageResponseUnionContent `json:"content"`
	CreatedAt    string                                      `json:"createdAt"`
	Direction    string                                      `json:"direction"`
	EnrollmentID string                                      `json:"enrollmentId"`
	// This field is from variant [SequenceRetrieveMessageResponseObject].
	OpenedAt string `json:"openedAt"`
	// This field is from variant [SequenceRetrieveMessageResponseObject].
	ReadAt string `json:"readAt"`
	// This field is a union of [SequenceRetrieveMessageResponseObjectReferenceUnion],
	// [SequenceRetrieveMessageResponseObject2ReferenceUnion]
	Reference SequenceRetrieveMessageResponseUnionReference `json:"reference"`
	// This field is from variant [SequenceRetrieveMessageResponseObject].
	RepliedAt string `json:"repliedAt"`
	// This field is from variant [SequenceRetrieveMessageResponseObject].
	SentAt           string `json:"sentAt"`
	SequenceStepID   string `json:"sequenceStepId"`
	SequenceStepType string `json:"sequenceStepType"`
	// This field is from variant [SequenceRetrieveMessageResponseObject].
	Status string `json:"status"`
	// This field is from variant [SequenceRetrieveMessageResponseObject].
	StatusChangedAt string `json:"statusChangedAt"`
	// This field is from variant [SequenceRetrieveMessageResponseObject2].
	ReceivedAt string `json:"receivedAt"`
	JSON       struct {
		ID               respjson.Field
		AcceptedAt       respjson.Field
		BouncedAt        respjson.Field
		Channel          respjson.Field
		ContactID        respjson.Field
		Content          respjson.Field
		CreatedAt        respjson.Field
		Direction        respjson.Field
		EnrollmentID     respjson.Field
		OpenedAt         respjson.Field
		ReadAt           respjson.Field
		Reference        respjson.Field
		RepliedAt        respjson.Field
		SentAt           respjson.Field
		SequenceStepID   respjson.Field
		SequenceStepType respjson.Field
		Status           respjson.Field
		StatusChangedAt  respjson.Field
		ReceivedAt       respjson.Field
		raw              string
	} `json:"-"`
}

func (u SequenceRetrieveMessageResponseUnion) AsSequenceRetrieveMessageResponseObject() (v SequenceRetrieveMessageResponseObject) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u SequenceRetrieveMessageResponseUnion) AsSequenceRetrieveMessageResponseObject2() (v SequenceRetrieveMessageResponseObject2) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u SequenceRetrieveMessageResponseUnion) RawJSON() string { return u.JSON.raw }

func (r *SequenceRetrieveMessageResponseUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// SequenceRetrieveMessageResponseUnionContent is an implicit subunion of
// [SequenceRetrieveMessageResponseUnion].
// SequenceRetrieveMessageResponseUnionContent provides convenient access to the
// sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [SequenceRetrieveMessageResponseUnion].
type SequenceRetrieveMessageResponseUnionContent struct {
	Body    string `json:"body"`
	Subject string `json:"subject"`
	JSON    struct {
		Body    respjson.Field
		Subject respjson.Field
		raw     string
	} `json:"-"`
}

func (r *SequenceRetrieveMessageResponseUnionContent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// SequenceRetrieveMessageResponseUnionReference is an implicit subunion of
// [SequenceRetrieveMessageResponseUnion].
// SequenceRetrieveMessageResponseUnionReference provides convenient access to the
// sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [SequenceRetrieveMessageResponseUnion].
type SequenceRetrieveMessageResponseUnionReference struct {
	ID         string `json:"id"`
	ObjectType string `json:"objectType"`
	JSON       struct {
		ID         respjson.Field
		ObjectType respjson.Field
		raw        string
	} `json:"-"`
}

func (r *SequenceRetrieveMessageResponseUnionReference) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SequenceRetrieveMessageResponseObject struct {
	// Opaque sequence message ID beginning with `sqm_`; use the ID returned by list
	// when retrieving a message.
	ID         string `json:"id" api:"required"`
	AcceptedAt string `json:"acceptedAt" api:"required"`
	BouncedAt  string `json:"bouncedAt" api:"required"`
	// Any of "EMAIL", "LINKEDIN".
	Channel string `json:"channel" api:"required"`
	// Contact ID beginning with `con_`.
	ContactID string `json:"contactId" api:"required"`
	// Message content available with `sequences:read` regardless of individual
	// mail/chat privacy; null when unavailable or the captured source was deleted, and
	// omitted entirely from list responses.
	Content SequenceRetrieveMessageResponseObjectContent `json:"content" api:"required"`
	// ISO 8601 creation timestamp of the current outbound attempt, or receipt
	// timestamp for an incoming reply; lists sort newest first by this field.
	CreatedAt string `json:"createdAt" api:"required"`
	// Any of "OUTBOUND".
	Direction string `json:"direction" api:"required"`
	// Enrollment ID beginning with `sqe_`.
	EnrollmentID string `json:"enrollmentId" api:"required"`
	OpenedAt     string `json:"openedAt" api:"required"`
	ReadAt       string `json:"readAt" api:"required"`
	// Active same-workspace email or chat message reference, or null when unavailable;
	// individual mail/chat privacy does not restrict this sequence reference, but the
	// referenced endpoint retains its own scopes and access rules.
	Reference SequenceRetrieveMessageResponseObjectReferenceUnion `json:"reference" api:"required"`
	RepliedAt string                                              `json:"repliedAt" api:"required"`
	SentAt    string                                              `json:"sentAt" api:"required"`
	// Authored step associated with this outbound message or with the outbound message
	// that received this reply.
	SequenceStepID string `json:"sequenceStepId" api:"required"`
	// Any of "EMAIL", "LINKEDIN_MESSAGE", "LINKEDIN_INVITE".
	SequenceStepType string `json:"sequenceStepType" api:"required"`
	// Any of "AWAITING_REVIEW", "SCHEDULED", "QUEUED", "SENT", "SKIPPED", "FAILED",
	// "CANCELLED".
	Status string `json:"status" api:"required"`
	// When the outbound status last changed; engagement timestamps do not necessarily
	// change this value.
	StatusChangedAt string `json:"statusChangedAt" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID               respjson.Field
		AcceptedAt       respjson.Field
		BouncedAt        respjson.Field
		Channel          respjson.Field
		ContactID        respjson.Field
		Content          respjson.Field
		CreatedAt        respjson.Field
		Direction        respjson.Field
		EnrollmentID     respjson.Field
		OpenedAt         respjson.Field
		ReadAt           respjson.Field
		Reference        respjson.Field
		RepliedAt        respjson.Field
		SentAt           respjson.Field
		SequenceStepID   respjson.Field
		SequenceStepType respjson.Field
		Status           respjson.Field
		StatusChangedAt  respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SequenceRetrieveMessageResponseObject) RawJSON() string { return r.JSON.raw }
func (r *SequenceRetrieveMessageResponseObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Message content available with `sequences:read` regardless of individual
// mail/chat privacy; null when unavailable or the captured source was deleted, and
// omitted entirely from list responses.
type SequenceRetrieveMessageResponseObjectContent struct {
	// Plain-text message body, or null for an invitation without a note.
	Body string `json:"body" api:"required"`
	// Email subject, or null for LinkedIn messages and invitations.
	Subject string `json:"subject" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Body        respjson.Field
		Subject     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SequenceRetrieveMessageResponseObjectContent) RawJSON() string { return r.JSON.raw }
func (r *SequenceRetrieveMessageResponseObjectContent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// SequenceRetrieveMessageResponseObjectReferenceUnion contains all possible
// properties and values from
// [SequenceRetrieveMessageResponseObjectReferenceObject],
// [SequenceRetrieveMessageResponseObjectReferenceObject2].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type SequenceRetrieveMessageResponseObjectReferenceUnion struct {
	ID         string `json:"id"`
	ObjectType string `json:"objectType"`
	JSON       struct {
		ID         respjson.Field
		ObjectType respjson.Field
		raw        string
	} `json:"-"`
}

func (u SequenceRetrieveMessageResponseObjectReferenceUnion) AsSequenceRetrieveMessageResponseObjectReferenceObject() (v SequenceRetrieveMessageResponseObjectReferenceObject) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u SequenceRetrieveMessageResponseObjectReferenceUnion) AsSequenceRetrieveMessageResponseObjectReferenceObject2() (v SequenceRetrieveMessageResponseObjectReferenceObject2) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u SequenceRetrieveMessageResponseObjectReferenceUnion) RawJSON() string { return u.JSON.raw }

func (r *SequenceRetrieveMessageResponseObjectReferenceUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SequenceRetrieveMessageResponseObjectReferenceObject struct {
	ID string `json:"id" api:"required"`
	// Any of "email".
	ObjectType string `json:"objectType" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		ObjectType  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SequenceRetrieveMessageResponseObjectReferenceObject) RawJSON() string { return r.JSON.raw }
func (r *SequenceRetrieveMessageResponseObjectReferenceObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SequenceRetrieveMessageResponseObjectReferenceObject2 struct {
	ID string `json:"id" api:"required"`
	// Any of "message".
	ObjectType string `json:"objectType" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		ObjectType  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SequenceRetrieveMessageResponseObjectReferenceObject2) RawJSON() string { return r.JSON.raw }
func (r *SequenceRetrieveMessageResponseObjectReferenceObject2) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SequenceRetrieveMessageResponseObject2 struct {
	// Opaque sequence message ID beginning with `sqm_`; use the ID returned by list
	// when retrieving a message.
	ID string `json:"id" api:"required"`
	// Any of "EMAIL", "LINKEDIN".
	Channel string `json:"channel" api:"required"`
	// Contact ID beginning with `con_`.
	ContactID string `json:"contactId" api:"required"`
	// Message content available with `sequences:read` regardless of individual
	// mail/chat privacy; null when unavailable or the captured source was deleted, and
	// omitted entirely from list responses.
	Content SequenceRetrieveMessageResponseObject2Content `json:"content" api:"required"`
	// ISO 8601 creation timestamp of the current outbound attempt, or receipt
	// timestamp for an incoming reply; lists sort newest first by this field.
	CreatedAt string `json:"createdAt" api:"required"`
	// Any of "INBOUND".
	Direction string `json:"direction" api:"required"`
	// Enrollment ID beginning with `sqe_`.
	EnrollmentID string `json:"enrollmentId" api:"required"`
	ReceivedAt   string `json:"receivedAt" api:"required"`
	// Active same-workspace email or chat message reference, or null when unavailable;
	// individual mail/chat privacy does not restrict this sequence reference, but the
	// referenced endpoint retains its own scopes and access rules.
	Reference SequenceRetrieveMessageResponseObject2ReferenceUnion `json:"reference" api:"required"`
	// Authored step associated with this outbound message or with the outbound message
	// that received this reply.
	SequenceStepID string `json:"sequenceStepId" api:"required"`
	// Any of "EMAIL", "LINKEDIN_MESSAGE", "LINKEDIN_INVITE".
	SequenceStepType string `json:"sequenceStepType" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID               respjson.Field
		Channel          respjson.Field
		ContactID        respjson.Field
		Content          respjson.Field
		CreatedAt        respjson.Field
		Direction        respjson.Field
		EnrollmentID     respjson.Field
		ReceivedAt       respjson.Field
		Reference        respjson.Field
		SequenceStepID   respjson.Field
		SequenceStepType respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SequenceRetrieveMessageResponseObject2) RawJSON() string { return r.JSON.raw }
func (r *SequenceRetrieveMessageResponseObject2) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Message content available with `sequences:read` regardless of individual
// mail/chat privacy; null when unavailable or the captured source was deleted, and
// omitted entirely from list responses.
type SequenceRetrieveMessageResponseObject2Content struct {
	// Plain-text message body, or null for an invitation without a note.
	Body string `json:"body" api:"required"`
	// Email subject, or null for LinkedIn messages and invitations.
	Subject string `json:"subject" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Body        respjson.Field
		Subject     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SequenceRetrieveMessageResponseObject2Content) RawJSON() string { return r.JSON.raw }
func (r *SequenceRetrieveMessageResponseObject2Content) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// SequenceRetrieveMessageResponseObject2ReferenceUnion contains all possible
// properties and values from
// [SequenceRetrieveMessageResponseObject2ReferenceObject],
// [SequenceRetrieveMessageResponseObject2ReferenceObject2].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type SequenceRetrieveMessageResponseObject2ReferenceUnion struct {
	ID         string `json:"id"`
	ObjectType string `json:"objectType"`
	JSON       struct {
		ID         respjson.Field
		ObjectType respjson.Field
		raw        string
	} `json:"-"`
}

func (u SequenceRetrieveMessageResponseObject2ReferenceUnion) AsSequenceRetrieveMessageResponseObject2ReferenceObject() (v SequenceRetrieveMessageResponseObject2ReferenceObject) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u SequenceRetrieveMessageResponseObject2ReferenceUnion) AsSequenceRetrieveMessageResponseObject2ReferenceObject2() (v SequenceRetrieveMessageResponseObject2ReferenceObject2) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u SequenceRetrieveMessageResponseObject2ReferenceUnion) RawJSON() string { return u.JSON.raw }

func (r *SequenceRetrieveMessageResponseObject2ReferenceUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SequenceRetrieveMessageResponseObject2ReferenceObject struct {
	ID string `json:"id" api:"required"`
	// Any of "email".
	ObjectType string `json:"objectType" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		ObjectType  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SequenceRetrieveMessageResponseObject2ReferenceObject) RawJSON() string { return r.JSON.raw }
func (r *SequenceRetrieveMessageResponseObject2ReferenceObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SequenceRetrieveMessageResponseObject2ReferenceObject2 struct {
	ID string `json:"id" api:"required"`
	// Any of "message".
	ObjectType string `json:"objectType" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		ObjectType  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SequenceRetrieveMessageResponseObject2ReferenceObject2) RawJSON() string { return r.JSON.raw }
func (r *SequenceRetrieveMessageResponseObject2ReferenceObject2) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SequenceRetrievePerformanceResponse struct {
	ByStepType SequenceRetrievePerformanceResponseByStepType `json:"byStepType" api:"required"`
	// Per-step statistics in published sequence order, including steps with no
	// activity. Steps no longer in the published definition are included if they have
	// sent messages.
	PerStep []SequenceRetrievePerformanceResponsePerStepUnion `json:"perStep" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ByStepType  respjson.Field
		PerStep     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SequenceRetrievePerformanceResponse) RawJSON() string { return r.JSON.raw }
func (r *SequenceRetrievePerformanceResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SequenceRetrievePerformanceResponseByStepType struct {
	// Sent email outcomes. Delivered is sent minus bounced, clamped at zero.
	Email             SequenceRetrievePerformanceResponseByStepTypeEmail             `json:"EMAIL" api:"required"`
	LinkedinInvite    SequenceRetrievePerformanceResponseByStepTypeLinkedinInvite    `json:"LINKEDIN_INVITE" api:"required"`
	LinkedinMessage   SequenceRetrievePerformanceResponseByStepTypeLinkedinMessage   `json:"LINKEDIN_MESSAGE" api:"required"`
	Wait              SequenceRetrievePerformanceResponseByStepTypeWait              `json:"WAIT" api:"required"`
	WaitForConnection SequenceRetrievePerformanceResponseByStepTypeWaitForConnection `json:"WAIT_FOR_CONNECTION" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Email             respjson.Field
		LinkedinInvite    respjson.Field
		LinkedinMessage   respjson.Field
		Wait              respjson.Field
		WaitForConnection respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SequenceRetrievePerformanceResponseByStepType) RawJSON() string { return r.JSON.raw }
func (r *SequenceRetrievePerformanceResponseByStepType) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Sent email outcomes. Delivered is sent minus bounced, clamped at zero.
type SequenceRetrievePerformanceResponseByStepTypeEmail struct {
	Bounced int64 `json:"bounced" api:"required"`
	// Bounced divided by sent; null if none sent.
	BounceRate float64 `json:"bounceRate" api:"required"`
	// Live enrollments currently on steps of this type.
	CurrentEnrollments int64 `json:"currentEnrollments" api:"required"`
	Delivered          int64 `json:"delivered" api:"required"`
	Opened             int64 `json:"opened" api:"required"`
	// Opened divided by delivered; null if none delivered.
	OpenRate float64 `json:"openRate" api:"required"`
	Replied  int64   `json:"replied" api:"required"`
	// Replied divided by delivered; null if none delivered.
	ReplyRate float64 `json:"replyRate" api:"required"`
	Sent      int64   `json:"sent" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Bounced            respjson.Field
		BounceRate         respjson.Field
		CurrentEnrollments respjson.Field
		Delivered          respjson.Field
		Opened             respjson.Field
		OpenRate           respjson.Field
		Replied            respjson.Field
		ReplyRate          respjson.Field
		Sent               respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SequenceRetrievePerformanceResponseByStepTypeEmail) RawJSON() string { return r.JSON.raw }
func (r *SequenceRetrievePerformanceResponseByStepTypeEmail) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SequenceRetrievePerformanceResponseByStepTypeLinkedinInvite struct {
	// Accepted divided by sent; null if none sent.
	AcceptanceRate float64 `json:"acceptanceRate" api:"required"`
	Accepted       int64   `json:"accepted" api:"required"`
	// Live enrollments currently on steps of this type.
	CurrentEnrollments int64 `json:"currentEnrollments" api:"required"`
	Replied            int64 `json:"replied" api:"required"`
	// Replied divided by sent; null if none sent.
	ReplyRate float64 `json:"replyRate" api:"required"`
	Sent      int64   `json:"sent" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AcceptanceRate     respjson.Field
		Accepted           respjson.Field
		CurrentEnrollments respjson.Field
		Replied            respjson.Field
		ReplyRate          respjson.Field
		Sent               respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SequenceRetrievePerformanceResponseByStepTypeLinkedinInvite) RawJSON() string {
	return r.JSON.raw
}
func (r *SequenceRetrievePerformanceResponseByStepTypeLinkedinInvite) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SequenceRetrievePerformanceResponseByStepTypeLinkedinMessage struct {
	// Live enrollments currently on steps of this type.
	CurrentEnrollments int64 `json:"currentEnrollments" api:"required"`
	Read               int64 `json:"read" api:"required"`
	// Read divided by sent; null if none sent.
	ReadRate float64 `json:"readRate" api:"required"`
	Replied  int64   `json:"replied" api:"required"`
	// Replied divided by sent; null if none sent.
	ReplyRate float64 `json:"replyRate" api:"required"`
	Sent      int64   `json:"sent" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CurrentEnrollments respjson.Field
		Read               respjson.Field
		ReadRate           respjson.Field
		Replied            respjson.Field
		ReplyRate          respjson.Field
		Sent               respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SequenceRetrievePerformanceResponseByStepTypeLinkedinMessage) RawJSON() string {
	return r.JSON.raw
}
func (r *SequenceRetrievePerformanceResponseByStepTypeLinkedinMessage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SequenceRetrievePerformanceResponseByStepTypeWait struct {
	// Live enrollments currently on steps of this type.
	CurrentEnrollments int64 `json:"currentEnrollments" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CurrentEnrollments respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SequenceRetrievePerformanceResponseByStepTypeWait) RawJSON() string { return r.JSON.raw }
func (r *SequenceRetrievePerformanceResponseByStepTypeWait) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SequenceRetrievePerformanceResponseByStepTypeWaitForConnection struct {
	// Live enrollments currently on steps of this type.
	CurrentEnrollments int64 `json:"currentEnrollments" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CurrentEnrollments respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SequenceRetrievePerformanceResponseByStepTypeWaitForConnection) RawJSON() string {
	return r.JSON.raw
}
func (r *SequenceRetrievePerformanceResponseByStepTypeWaitForConnection) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// SequenceRetrievePerformanceResponsePerStepUnion contains all possible properties
// and values from [SequenceRetrievePerformanceResponsePerStepEmailStep],
// [SequenceRetrievePerformanceResponsePerStepLinkedInMessageStep],
// [SequenceRetrievePerformanceResponsePerStepLinkedInInviteStep],
// [SequenceRetrievePerformanceResponsePerStepWaitStep],
// [SequenceRetrievePerformanceResponsePerStepWaitForConnectionStep].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type SequenceRetrievePerformanceResponsePerStepUnion struct {
	// This field is from variant
	// [SequenceRetrievePerformanceResponsePerStepEmailStep].
	Bounced int64 `json:"bounced"`
	// This field is from variant
	// [SequenceRetrievePerformanceResponsePerStepEmailStep].
	BounceRate         float64 `json:"bounceRate"`
	CurrentEnrollments int64   `json:"currentEnrollments"`
	// This field is from variant
	// [SequenceRetrievePerformanceResponsePerStepEmailStep].
	Delivered int64 `json:"delivered"`
	// This field is from variant
	// [SequenceRetrievePerformanceResponsePerStepEmailStep].
	Opened int64 `json:"opened"`
	// This field is from variant
	// [SequenceRetrievePerformanceResponsePerStepEmailStep].
	OpenRate         float64 `json:"openRate"`
	Replied          int64   `json:"replied"`
	ReplyRate        float64 `json:"replyRate"`
	Sent             int64   `json:"sent"`
	SequenceStepID   string  `json:"sequenceStepId"`
	SequenceStepType string  `json:"sequenceStepType"`
	// This field is from variant
	// [SequenceRetrievePerformanceResponsePerStepLinkedInMessageStep].
	Read int64 `json:"read"`
	// This field is from variant
	// [SequenceRetrievePerformanceResponsePerStepLinkedInMessageStep].
	ReadRate float64 `json:"readRate"`
	// This field is from variant
	// [SequenceRetrievePerformanceResponsePerStepLinkedInInviteStep].
	AcceptanceRate float64 `json:"acceptanceRate"`
	// This field is from variant
	// [SequenceRetrievePerformanceResponsePerStepLinkedInInviteStep].
	Accepted int64 `json:"accepted"`
	JSON     struct {
		Bounced            respjson.Field
		BounceRate         respjson.Field
		CurrentEnrollments respjson.Field
		Delivered          respjson.Field
		Opened             respjson.Field
		OpenRate           respjson.Field
		Replied            respjson.Field
		ReplyRate          respjson.Field
		Sent               respjson.Field
		SequenceStepID     respjson.Field
		SequenceStepType   respjson.Field
		Read               respjson.Field
		ReadRate           respjson.Field
		AcceptanceRate     respjson.Field
		Accepted           respjson.Field
		raw                string
	} `json:"-"`
}

func (u SequenceRetrievePerformanceResponsePerStepUnion) AsEmailStep() (v SequenceRetrievePerformanceResponsePerStepEmailStep) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u SequenceRetrievePerformanceResponsePerStepUnion) AsLinkedInMessageStep() (v SequenceRetrievePerformanceResponsePerStepLinkedInMessageStep) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u SequenceRetrievePerformanceResponsePerStepUnion) AsLinkedInInviteStep() (v SequenceRetrievePerformanceResponsePerStepLinkedInInviteStep) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u SequenceRetrievePerformanceResponsePerStepUnion) AsWaitStep() (v SequenceRetrievePerformanceResponsePerStepWaitStep) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u SequenceRetrievePerformanceResponsePerStepUnion) AsWaitForConnectionStep() (v SequenceRetrievePerformanceResponsePerStepWaitForConnectionStep) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u SequenceRetrievePerformanceResponsePerStepUnion) RawJSON() string { return u.JSON.raw }

func (r *SequenceRetrievePerformanceResponsePerStepUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SequenceRetrievePerformanceResponsePerStepEmailStep struct {
	Bounced int64 `json:"bounced" api:"required"`
	// Bounced divided by sent; null if none sent.
	BounceRate float64 `json:"bounceRate" api:"required"`
	// Live enrollments currently on this step in ACTIVE or PAUSED lifecycle state,
	// including review and attention states.
	CurrentEnrollments int64 `json:"currentEnrollments" api:"required"`
	Delivered          int64 `json:"delivered" api:"required"`
	Opened             int64 `json:"opened" api:"required"`
	// Opened divided by delivered; null if none delivered.
	OpenRate float64 `json:"openRate" api:"required"`
	Replied  int64   `json:"replied" api:"required"`
	// Replied divided by delivered; null if none delivered.
	ReplyRate      float64 `json:"replyRate" api:"required"`
	Sent           int64   `json:"sent" api:"required"`
	SequenceStepID string  `json:"sequenceStepId" api:"required"`
	// Any of "EMAIL".
	SequenceStepType string `json:"sequenceStepType" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Bounced            respjson.Field
		BounceRate         respjson.Field
		CurrentEnrollments respjson.Field
		Delivered          respjson.Field
		Opened             respjson.Field
		OpenRate           respjson.Field
		Replied            respjson.Field
		ReplyRate          respjson.Field
		Sent               respjson.Field
		SequenceStepID     respjson.Field
		SequenceStepType   respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SequenceRetrievePerformanceResponsePerStepEmailStep) RawJSON() string { return r.JSON.raw }
func (r *SequenceRetrievePerformanceResponsePerStepEmailStep) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SequenceRetrievePerformanceResponsePerStepLinkedInMessageStep struct {
	// Live enrollments currently on this step in ACTIVE or PAUSED lifecycle state,
	// including review and attention states.
	CurrentEnrollments int64 `json:"currentEnrollments" api:"required"`
	Read               int64 `json:"read" api:"required"`
	// Read divided by sent; null if none sent.
	ReadRate float64 `json:"readRate" api:"required"`
	Replied  int64   `json:"replied" api:"required"`
	// Replied divided by sent; null if none sent.
	ReplyRate      float64 `json:"replyRate" api:"required"`
	Sent           int64   `json:"sent" api:"required"`
	SequenceStepID string  `json:"sequenceStepId" api:"required"`
	// Any of "LINKEDIN_MESSAGE".
	SequenceStepType string `json:"sequenceStepType" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CurrentEnrollments respjson.Field
		Read               respjson.Field
		ReadRate           respjson.Field
		Replied            respjson.Field
		ReplyRate          respjson.Field
		Sent               respjson.Field
		SequenceStepID     respjson.Field
		SequenceStepType   respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SequenceRetrievePerformanceResponsePerStepLinkedInMessageStep) RawJSON() string {
	return r.JSON.raw
}
func (r *SequenceRetrievePerformanceResponsePerStepLinkedInMessageStep) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SequenceRetrievePerformanceResponsePerStepLinkedInInviteStep struct {
	// Accepted divided by sent; null if none sent.
	AcceptanceRate float64 `json:"acceptanceRate" api:"required"`
	Accepted       int64   `json:"accepted" api:"required"`
	// Live enrollments currently on this step in ACTIVE or PAUSED lifecycle state,
	// including review and attention states.
	CurrentEnrollments int64 `json:"currentEnrollments" api:"required"`
	Replied            int64 `json:"replied" api:"required"`
	// Replied divided by sent; null if none sent.
	ReplyRate      float64 `json:"replyRate" api:"required"`
	Sent           int64   `json:"sent" api:"required"`
	SequenceStepID string  `json:"sequenceStepId" api:"required"`
	// Any of "LINKEDIN_INVITE".
	SequenceStepType string `json:"sequenceStepType" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AcceptanceRate     respjson.Field
		Accepted           respjson.Field
		CurrentEnrollments respjson.Field
		Replied            respjson.Field
		ReplyRate          respjson.Field
		Sent               respjson.Field
		SequenceStepID     respjson.Field
		SequenceStepType   respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SequenceRetrievePerformanceResponsePerStepLinkedInInviteStep) RawJSON() string {
	return r.JSON.raw
}
func (r *SequenceRetrievePerformanceResponsePerStepLinkedInInviteStep) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SequenceRetrievePerformanceResponsePerStepWaitStep struct {
	// Live enrollments currently on this step in ACTIVE or PAUSED lifecycle state,
	// including review and attention states.
	CurrentEnrollments int64  `json:"currentEnrollments" api:"required"`
	SequenceStepID     string `json:"sequenceStepId" api:"required"`
	// Any of "WAIT".
	SequenceStepType string `json:"sequenceStepType" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CurrentEnrollments respjson.Field
		SequenceStepID     respjson.Field
		SequenceStepType   respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SequenceRetrievePerformanceResponsePerStepWaitStep) RawJSON() string { return r.JSON.raw }
func (r *SequenceRetrievePerformanceResponsePerStepWaitStep) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SequenceRetrievePerformanceResponsePerStepWaitForConnectionStep struct {
	// Live enrollments currently on this step in ACTIVE or PAUSED lifecycle state,
	// including review and attention states.
	CurrentEnrollments int64  `json:"currentEnrollments" api:"required"`
	SequenceStepID     string `json:"sequenceStepId" api:"required"`
	// Any of "WAIT_FOR_CONNECTION".
	SequenceStepType string `json:"sequenceStepType" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CurrentEnrollments respjson.Field
		SequenceStepID     respjson.Field
		SequenceStepType   respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SequenceRetrievePerformanceResponsePerStepWaitForConnectionStep) RawJSON() string {
	return r.JSON.raw
}
func (r *SequenceRetrievePerformanceResponsePerStepWaitForConnectionStep) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SequenceRetrieveResponse struct {
	// Unique identifier of the sequence.
	ID string `json:"id" api:"required"`
	// Channels the published steps send on; empty before the first publish.
	//
	// Any of "EMAIL", "LINKEDIN".
	Channels []string `json:"channels" api:"required"`
	// ISO 8601 creation timestamp.
	CreatedAt string `json:"createdAt" api:"required"`
	// Published configuration, or null before the first publish.
	Definition SequenceRetrieveResponseDefinition `json:"definition" api:"required"`
	// When the sequence was first activated.
	FirstActivatedAt string `json:"firstActivatedAt" api:"required"`
	// When the sequence was most recently activated.
	LastActivatedAt string `json:"lastActivatedAt" api:"required"`
	// Name of the sequence.
	Name string `json:"name" api:"required"`
	// Always `sequence`.
	//
	// Any of "sequence".
	ObjectType SequenceRetrieveResponseObjectType `json:"objectType" api:"required"`
	// Extra members who may approve this sequence’s `FOR_REVIEW` messages before they
	// send. The sending member and the sequence’s creator always review, whether or
	// not they are listed. A change applies to sends once the draft is published. The
	// published list; empty before the first publish. Lists active members only.
	Reviewers []SequenceRetrieveResponseReviewer `json:"reviewers" api:"required"`
	// The senders this sequence sends from; new contacts rotate across them. A change
	// applies to sends once the draft is published. An archived sender stays listed
	// but never sends; any change to the list drops it.
	Senders []SequenceRetrieveResponseSender `json:"senders" api:"required"`
	// `DRAFT` has not been published; `ACTIVE` is enabled for sending; `DEACTIVATED`
	// is published but paused.
	//
	// Any of "DRAFT", "ACTIVE", "DEACTIVATED".
	Status SequenceRetrieveResponseStatus `json:"status" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID               respjson.Field
		Channels         respjson.Field
		CreatedAt        respjson.Field
		Definition       respjson.Field
		FirstActivatedAt respjson.Field
		LastActivatedAt  respjson.Field
		Name             respjson.Field
		ObjectType       respjson.Field
		Reviewers        respjson.Field
		Senders          respjson.Field
		Status           respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SequenceRetrieveResponse) RawJSON() string { return r.JSON.raw }
func (r *SequenceRetrieveResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Published configuration, or null before the first publish.
type SequenceRetrieveResponseDefinition struct {
	// What this version of the sequence is for.
	Description string `json:"description" api:"required"`
	// Messaging inputs that frame every step.
	Messaging SequenceRetrieveResponseDefinitionMessaging `json:"messaging" api:"required"`
	// How the whole sequence paces each step type it sends, keyed by sequence step
	// type.
	Pacing SequenceRetrieveResponseDefinitionPacing `json:"pacing" api:"required"`
	// Send priority relative to other sequences.
	//
	// Any of "HIGH", "NORMAL".
	Priority string `json:"priority" api:"required"`
	// When the sequence may send.
	SendWindow SequenceRetrieveResponseDefinitionSendWindow `json:"sendWindow" api:"required"`
	// Email, LinkedIn and wait steps in execution order; call steps are not exposed in
	// this release.
	Steps []SequenceRetrieveResponseDefinitionStepUnion `json:"steps" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Description respjson.Field
		Messaging   respjson.Field
		Pacing      respjson.Field
		Priority    respjson.Field
		SendWindow  respjson.Field
		Steps       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SequenceRetrieveResponseDefinition) RawJSON() string { return r.JSON.raw }
func (r *SequenceRetrieveResponseDefinition) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Messaging inputs that frame every step.
type SequenceRetrieveResponseDefinitionMessaging struct {
	// What the recipient is asked to do.
	CallToAction string `json:"callToAction" api:"required"`
	// Who the sequence targets.
	Persona string `json:"persona" api:"required"`
	// Why the outreach is timely.
	Trigger string `json:"trigger" api:"required"`
	// The value proposition.
	ValueProp string `json:"valueProp" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CallToAction respjson.Field
		Persona      respjson.Field
		Trigger      respjson.Field
		ValueProp    respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SequenceRetrieveResponseDefinitionMessaging) RawJSON() string { return r.JSON.raw }
func (r *SequenceRetrieveResponseDefinitionMessaging) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// How the whole sequence paces each step type it sends, keyed by sequence step
// type.
type SequenceRetrieveResponseDefinitionPacing struct {
	// Email steps.
	Email SequenceRetrieveResponseDefinitionPacingEmail `json:"EMAIL" api:"required"`
	// LinkedIn connection requests.
	LinkedinInvite SequenceRetrieveResponseDefinitionPacingLinkedinInvite `json:"LINKEDIN_INVITE" api:"required"`
	// LinkedIn messages.
	LinkedinMessage SequenceRetrieveResponseDefinitionPacingLinkedinMessage `json:"LINKEDIN_MESSAGE" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Email           respjson.Field
		LinkedinInvite  respjson.Field
		LinkedinMessage respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SequenceRetrieveResponseDefinitionPacing) RawJSON() string { return r.JSON.raw }
func (r *SequenceRetrieveResponseDefinitionPacing) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Email steps.
type SequenceRetrieveResponseDefinitionPacingEmail struct {
	// Most sends of this step type across the whole sequence in any rolling 24 hours,
	// or null when the sequence sets no daily cap. Workspace and sender limits still
	// apply.
	DailyLimit int64 `json:"dailyLimit" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DailyLimit  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SequenceRetrieveResponseDefinitionPacingEmail) RawJSON() string { return r.JSON.raw }
func (r *SequenceRetrieveResponseDefinitionPacingEmail) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// LinkedIn connection requests.
type SequenceRetrieveResponseDefinitionPacingLinkedinInvite struct {
	// Most sends of this step type across the whole sequence in any rolling 24 hours,
	// or null when the sequence sets no daily cap. Workspace and sender limits still
	// apply.
	DailyLimit int64 `json:"dailyLimit" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DailyLimit  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SequenceRetrieveResponseDefinitionPacingLinkedinInvite) RawJSON() string { return r.JSON.raw }
func (r *SequenceRetrieveResponseDefinitionPacingLinkedinInvite) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// LinkedIn messages.
type SequenceRetrieveResponseDefinitionPacingLinkedinMessage struct {
	// Most sends of this step type across the whole sequence in any rolling 24 hours,
	// or null when the sequence sets no daily cap. Workspace and sender limits still
	// apply.
	DailyLimit int64 `json:"dailyLimit" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DailyLimit  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SequenceRetrieveResponseDefinitionPacingLinkedinMessage) RawJSON() string { return r.JSON.raw }
func (r *SequenceRetrieveResponseDefinitionPacingLinkedinMessage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// When the sequence may send.
type SequenceRetrieveResponseDefinitionSendWindow struct {
	// Days sends may go out, 0 = Sunday.
	DaysOfWeek []int64 `json:"daysOfWeek" api:"required"`
	// Local hour sends stop, exclusive (1–24).
	EndHour int64 `json:"endHour" api:"required"`
	// First local hour sends may go out (0–23).
	StartHour int64 `json:"startHour" api:"required"`
	// IANA timezone of the window.
	Timezone string `json:"timezone" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DaysOfWeek  respjson.Field
		EndHour     respjson.Field
		StartHour   respjson.Field
		Timezone    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SequenceRetrieveResponseDefinitionSendWindow) RawJSON() string { return r.JSON.raw }
func (r *SequenceRetrieveResponseDefinitionSendWindow) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// SequenceRetrieveResponseDefinitionStepUnion contains all possible properties and
// values from [SequenceRetrieveResponseDefinitionStepObject],
// [SequenceRetrieveResponseDefinitionStepObject2],
// [SequenceRetrieveResponseDefinitionStepObject3].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type SequenceRetrieveResponseDefinitionStepUnion struct {
	ID           string `json:"id"`
	DeliveryMode string `json:"deliveryMode"`
	Name         string `json:"name"`
	Prompt       string `json:"prompt"`
	Type         string `json:"type"`
	// This field is from variant [SequenceRetrieveResponseDefinitionStepObject2].
	IncludeNote bool `json:"includeNote"`
	// This field is from variant [SequenceRetrieveResponseDefinitionStepObject3].
	Duration int64 `json:"duration"`
	// This field is from variant [SequenceRetrieveResponseDefinitionStepObject3].
	Unit string `json:"unit"`
	JSON struct {
		ID           respjson.Field
		DeliveryMode respjson.Field
		Name         respjson.Field
		Prompt       respjson.Field
		Type         respjson.Field
		IncludeNote  respjson.Field
		Duration     respjson.Field
		Unit         respjson.Field
		raw          string
	} `json:"-"`
}

func (u SequenceRetrieveResponseDefinitionStepUnion) AsSequenceRetrieveResponseDefinitionStepObject() (v SequenceRetrieveResponseDefinitionStepObject) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u SequenceRetrieveResponseDefinitionStepUnion) AsSequenceRetrieveResponseDefinitionStepObject2() (v SequenceRetrieveResponseDefinitionStepObject2) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u SequenceRetrieveResponseDefinitionStepUnion) AsSequenceRetrieveResponseDefinitionStepObject3() (v SequenceRetrieveResponseDefinitionStepObject3) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u SequenceRetrieveResponseDefinitionStepUnion) RawJSON() string { return u.JSON.raw }

func (r *SequenceRetrieveResponseDefinitionStepUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SequenceRetrieveResponseDefinitionStepObject struct {
	// Step identifier, referenced by enrollment `currentStepId` and message
	// `sequenceStepId`.
	ID string `json:"id" api:"required"`
	// `AUTO` sends without manual approval, subject to the send window and limits;
	// `FOR_REVIEW` requires a person to approve the draft before sending.
	//
	// Any of "AUTO", "FOR_REVIEW".
	DeliveryMode string `json:"deliveryMode" api:"required"`
	// Display name of the step.
	Name string `json:"name" api:"required"`
	// Instructions the agent follows when composing this message.
	Prompt string `json:"prompt" api:"required"`
	// Step type.
	//
	// Any of "EMAIL", "LINKEDIN_MESSAGE".
	Type string `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		DeliveryMode respjson.Field
		Name         respjson.Field
		Prompt       respjson.Field
		Type         respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SequenceRetrieveResponseDefinitionStepObject) RawJSON() string { return r.JSON.raw }
func (r *SequenceRetrieveResponseDefinitionStepObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SequenceRetrieveResponseDefinitionStepObject2 struct {
	// Step identifier, referenced by enrollment `currentStepId` and message
	// `sequenceStepId`.
	ID string `json:"id" api:"required"`
	// `AUTO` sends without manual approval, subject to the send window and limits;
	// `FOR_REVIEW` requires a person to approve the draft before sending.
	//
	// Any of "AUTO", "FOR_REVIEW".
	DeliveryMode string `json:"deliveryMode" api:"required"`
	// Whether the connection request carries a note. A step authored before this
	// setting existed reads as true.
	IncludeNote bool `json:"includeNote" api:"required"`
	// Display name of the step.
	Name string `json:"name" api:"required"`
	// Instructions the agent follows when composing this message.
	Prompt string `json:"prompt" api:"required"`
	// Step type.
	//
	// Any of "LINKEDIN_INVITE".
	Type string `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		DeliveryMode respjson.Field
		IncludeNote  respjson.Field
		Name         respjson.Field
		Prompt       respjson.Field
		Type         respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SequenceRetrieveResponseDefinitionStepObject2) RawJSON() string { return r.JSON.raw }
func (r *SequenceRetrieveResponseDefinitionStepObject2) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SequenceRetrieveResponseDefinitionStepObject3 struct {
	// Step identifier.
	ID string `json:"id" api:"required"`
	// Length of the wait.
	Duration int64 `json:"duration" api:"required"`
	// `WAIT` pauses for a fixed time; `WAIT_FOR_CONNECTION` waits up to that long for
	// a LinkedIn invite to be accepted.
	//
	// Any of "WAIT", "WAIT_FOR_CONNECTION".
	Type string `json:"type" api:"required"`
	// Unit of `duration`.
	//
	// Any of "DAY", "MINUTE".
	Unit string `json:"unit" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Duration    respjson.Field
		Type        respjson.Field
		Unit        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SequenceRetrieveResponseDefinitionStepObject3) RawJSON() string { return r.JSON.raw }
func (r *SequenceRetrieveResponseDefinitionStepObject3) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Always `sequence`.
type SequenceRetrieveResponseObjectType string

const (
	SequenceRetrieveResponseObjectTypeSequence SequenceRetrieveResponseObjectType = "sequence"
)

type SequenceRetrieveResponseReviewer struct {
	// ID of the member.
	ID string `json:"id" api:"required"`
	// Always `member`.
	//
	// Any of "member".
	ObjectType string `json:"objectType" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		ObjectType  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SequenceRetrieveResponseReviewer) RawJSON() string { return r.JSON.raw }
func (r *SequenceRetrieveResponseReviewer) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SequenceRetrieveResponseSender struct {
	// ID of the sender.
	ID string `json:"id" api:"required"`
	// Always `sender`.
	//
	// Any of "sender".
	ObjectType string `json:"objectType" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		ObjectType  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SequenceRetrieveResponseSender) RawJSON() string { return r.JSON.raw }
func (r *SequenceRetrieveResponseSender) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// `DRAFT` has not been published; `ACTIVE` is enabled for sending; `DEACTIVATED`
// is published but paused.
type SequenceRetrieveResponseStatus string

const (
	SequenceRetrieveResponseStatusDraft       SequenceRetrieveResponseStatus = "DRAFT"
	SequenceRetrieveResponseStatusActive      SequenceRetrieveResponseStatus = "ACTIVE"
	SequenceRetrieveResponseStatusDeactivated SequenceRetrieveResponseStatus = "DEACTIVATED"
)

type SequenceUpdateDraftResponse struct {
	// Channels the draft's steps send on.
	//
	// Any of "EMAIL", "LINKEDIN".
	Channels []string `json:"channels" api:"required"`
	// The draft configuration, in the same shape as a published definition.
	Definition SequenceUpdateDraftResponseDefinition `json:"definition" api:"required"`
	// Name the sequence takes when the draft is published.
	Name string `json:"name" api:"required"`
	// What publishing this draft would change for contacts already in the sequence.
	// Before the first publish, only proposed senders it would leave off because they
	// were archived. Empty when nothing is affected.
	PublishWarnings []SequenceUpdateDraftResponsePublishWarning `json:"publishWarnings" api:"required"`
	// Extra members who may approve this sequence’s `FOR_REVIEW` messages before they
	// send. The sending member and the sequence’s creator always review, whether or
	// not they are listed. A change applies to sends once the draft is published.
	// Lists active members only.
	Reviewers []SequenceUpdateDraftResponseReviewer `json:"reviewers" api:"required"`
	// The senders this sequence sends from; new contacts rotate across them. A change
	// applies to sends once the draft is published. The list this draft publishes: the
	// live list until a save changes it. An archived sender stays listed but never
	// sends; any change to the list drops it.
	Senders []SequenceUpdateDraftResponseSender `json:"senders" api:"required"`
	// The sequence this draft belongs to.
	SequenceID string `json:"sequenceId" api:"required"`
	// When the draft last changed.
	UpdatedAt string `json:"updatedAt" api:"required"`
	// Opaque token for the current state of the draft. Send it as `expectedVersion` on
	// the next save, publish or discard; it changes whenever anyone edits the draft.
	Version string `json:"version" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Channels        respjson.Field
		Definition      respjson.Field
		Name            respjson.Field
		PublishWarnings respjson.Field
		Reviewers       respjson.Field
		Senders         respjson.Field
		SequenceID      respjson.Field
		UpdatedAt       respjson.Field
		Version         respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SequenceUpdateDraftResponse) RawJSON() string { return r.JSON.raw }
func (r *SequenceUpdateDraftResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The draft configuration, in the same shape as a published definition.
type SequenceUpdateDraftResponseDefinition struct {
	// What this version of the sequence is for.
	Description string `json:"description" api:"required"`
	// Messaging inputs that frame every step.
	Messaging SequenceUpdateDraftResponseDefinitionMessaging `json:"messaging" api:"required"`
	// How the whole sequence paces each step type it sends, keyed by sequence step
	// type.
	Pacing SequenceUpdateDraftResponseDefinitionPacing `json:"pacing" api:"required"`
	// Send priority relative to other sequences.
	//
	// Any of "HIGH", "NORMAL".
	Priority string `json:"priority" api:"required"`
	// When the sequence may send.
	SendWindow SequenceUpdateDraftResponseDefinitionSendWindow `json:"sendWindow" api:"required"`
	// Email, LinkedIn and wait steps in execution order; call steps are not exposed in
	// this release.
	Steps []SequenceUpdateDraftResponseDefinitionStepUnion `json:"steps" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Description respjson.Field
		Messaging   respjson.Field
		Pacing      respjson.Field
		Priority    respjson.Field
		SendWindow  respjson.Field
		Steps       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SequenceUpdateDraftResponseDefinition) RawJSON() string { return r.JSON.raw }
func (r *SequenceUpdateDraftResponseDefinition) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Messaging inputs that frame every step.
type SequenceUpdateDraftResponseDefinitionMessaging struct {
	// What the recipient is asked to do.
	CallToAction string `json:"callToAction" api:"required"`
	// Who the sequence targets.
	Persona string `json:"persona" api:"required"`
	// Why the outreach is timely.
	Trigger string `json:"trigger" api:"required"`
	// The value proposition.
	ValueProp string `json:"valueProp" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CallToAction respjson.Field
		Persona      respjson.Field
		Trigger      respjson.Field
		ValueProp    respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SequenceUpdateDraftResponseDefinitionMessaging) RawJSON() string { return r.JSON.raw }
func (r *SequenceUpdateDraftResponseDefinitionMessaging) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// How the whole sequence paces each step type it sends, keyed by sequence step
// type.
type SequenceUpdateDraftResponseDefinitionPacing struct {
	// Email steps.
	Email SequenceUpdateDraftResponseDefinitionPacingEmail `json:"EMAIL" api:"required"`
	// LinkedIn connection requests.
	LinkedinInvite SequenceUpdateDraftResponseDefinitionPacingLinkedinInvite `json:"LINKEDIN_INVITE" api:"required"`
	// LinkedIn messages.
	LinkedinMessage SequenceUpdateDraftResponseDefinitionPacingLinkedinMessage `json:"LINKEDIN_MESSAGE" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Email           respjson.Field
		LinkedinInvite  respjson.Field
		LinkedinMessage respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SequenceUpdateDraftResponseDefinitionPacing) RawJSON() string { return r.JSON.raw }
func (r *SequenceUpdateDraftResponseDefinitionPacing) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Email steps.
type SequenceUpdateDraftResponseDefinitionPacingEmail struct {
	// Most sends of this step type across the whole sequence in any rolling 24 hours,
	// or null when the sequence sets no daily cap. Workspace and sender limits still
	// apply.
	DailyLimit int64 `json:"dailyLimit" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DailyLimit  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SequenceUpdateDraftResponseDefinitionPacingEmail) RawJSON() string { return r.JSON.raw }
func (r *SequenceUpdateDraftResponseDefinitionPacingEmail) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// LinkedIn connection requests.
type SequenceUpdateDraftResponseDefinitionPacingLinkedinInvite struct {
	// Most sends of this step type across the whole sequence in any rolling 24 hours,
	// or null when the sequence sets no daily cap. Workspace and sender limits still
	// apply.
	DailyLimit int64 `json:"dailyLimit" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DailyLimit  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SequenceUpdateDraftResponseDefinitionPacingLinkedinInvite) RawJSON() string {
	return r.JSON.raw
}
func (r *SequenceUpdateDraftResponseDefinitionPacingLinkedinInvite) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// LinkedIn messages.
type SequenceUpdateDraftResponseDefinitionPacingLinkedinMessage struct {
	// Most sends of this step type across the whole sequence in any rolling 24 hours,
	// or null when the sequence sets no daily cap. Workspace and sender limits still
	// apply.
	DailyLimit int64 `json:"dailyLimit" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DailyLimit  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SequenceUpdateDraftResponseDefinitionPacingLinkedinMessage) RawJSON() string {
	return r.JSON.raw
}
func (r *SequenceUpdateDraftResponseDefinitionPacingLinkedinMessage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// When the sequence may send.
type SequenceUpdateDraftResponseDefinitionSendWindow struct {
	// Days sends may go out, 0 = Sunday.
	DaysOfWeek []int64 `json:"daysOfWeek" api:"required"`
	// Local hour sends stop, exclusive (1–24).
	EndHour int64 `json:"endHour" api:"required"`
	// First local hour sends may go out (0–23).
	StartHour int64 `json:"startHour" api:"required"`
	// IANA timezone of the window.
	Timezone string `json:"timezone" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DaysOfWeek  respjson.Field
		EndHour     respjson.Field
		StartHour   respjson.Field
		Timezone    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SequenceUpdateDraftResponseDefinitionSendWindow) RawJSON() string { return r.JSON.raw }
func (r *SequenceUpdateDraftResponseDefinitionSendWindow) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// SequenceUpdateDraftResponseDefinitionStepUnion contains all possible properties
// and values from [SequenceUpdateDraftResponseDefinitionStepObject],
// [SequenceUpdateDraftResponseDefinitionStepObject2],
// [SequenceUpdateDraftResponseDefinitionStepObject3].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type SequenceUpdateDraftResponseDefinitionStepUnion struct {
	ID           string `json:"id"`
	DeliveryMode string `json:"deliveryMode"`
	Name         string `json:"name"`
	Prompt       string `json:"prompt"`
	Type         string `json:"type"`
	// This field is from variant [SequenceUpdateDraftResponseDefinitionStepObject2].
	IncludeNote bool `json:"includeNote"`
	// This field is from variant [SequenceUpdateDraftResponseDefinitionStepObject3].
	Duration int64 `json:"duration"`
	// This field is from variant [SequenceUpdateDraftResponseDefinitionStepObject3].
	Unit string `json:"unit"`
	JSON struct {
		ID           respjson.Field
		DeliveryMode respjson.Field
		Name         respjson.Field
		Prompt       respjson.Field
		Type         respjson.Field
		IncludeNote  respjson.Field
		Duration     respjson.Field
		Unit         respjson.Field
		raw          string
	} `json:"-"`
}

func (u SequenceUpdateDraftResponseDefinitionStepUnion) AsSequenceUpdateDraftResponseDefinitionStepObject() (v SequenceUpdateDraftResponseDefinitionStepObject) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u SequenceUpdateDraftResponseDefinitionStepUnion) AsSequenceUpdateDraftResponseDefinitionStepObject2() (v SequenceUpdateDraftResponseDefinitionStepObject2) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u SequenceUpdateDraftResponseDefinitionStepUnion) AsSequenceUpdateDraftResponseDefinitionStepObject3() (v SequenceUpdateDraftResponseDefinitionStepObject3) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u SequenceUpdateDraftResponseDefinitionStepUnion) RawJSON() string { return u.JSON.raw }

func (r *SequenceUpdateDraftResponseDefinitionStepUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SequenceUpdateDraftResponseDefinitionStepObject struct {
	// Step identifier, referenced by enrollment `currentStepId` and message
	// `sequenceStepId`.
	ID string `json:"id" api:"required"`
	// `AUTO` sends without manual approval, subject to the send window and limits;
	// `FOR_REVIEW` requires a person to approve the draft before sending.
	//
	// Any of "AUTO", "FOR_REVIEW".
	DeliveryMode string `json:"deliveryMode" api:"required"`
	// Display name of the step.
	Name string `json:"name" api:"required"`
	// Instructions the agent follows when composing this message.
	Prompt string `json:"prompt" api:"required"`
	// Step type.
	//
	// Any of "EMAIL", "LINKEDIN_MESSAGE".
	Type string `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		DeliveryMode respjson.Field
		Name         respjson.Field
		Prompt       respjson.Field
		Type         respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SequenceUpdateDraftResponseDefinitionStepObject) RawJSON() string { return r.JSON.raw }
func (r *SequenceUpdateDraftResponseDefinitionStepObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SequenceUpdateDraftResponseDefinitionStepObject2 struct {
	// Step identifier, referenced by enrollment `currentStepId` and message
	// `sequenceStepId`.
	ID string `json:"id" api:"required"`
	// `AUTO` sends without manual approval, subject to the send window and limits;
	// `FOR_REVIEW` requires a person to approve the draft before sending.
	//
	// Any of "AUTO", "FOR_REVIEW".
	DeliveryMode string `json:"deliveryMode" api:"required"`
	// Whether the connection request carries a note. A step authored before this
	// setting existed reads as true.
	IncludeNote bool `json:"includeNote" api:"required"`
	// Display name of the step.
	Name string `json:"name" api:"required"`
	// Instructions the agent follows when composing this message.
	Prompt string `json:"prompt" api:"required"`
	// Step type.
	//
	// Any of "LINKEDIN_INVITE".
	Type string `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		DeliveryMode respjson.Field
		IncludeNote  respjson.Field
		Name         respjson.Field
		Prompt       respjson.Field
		Type         respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SequenceUpdateDraftResponseDefinitionStepObject2) RawJSON() string { return r.JSON.raw }
func (r *SequenceUpdateDraftResponseDefinitionStepObject2) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SequenceUpdateDraftResponseDefinitionStepObject3 struct {
	// Step identifier.
	ID string `json:"id" api:"required"`
	// Length of the wait.
	Duration int64 `json:"duration" api:"required"`
	// `WAIT` pauses for a fixed time; `WAIT_FOR_CONNECTION` waits up to that long for
	// a LinkedIn invite to be accepted.
	//
	// Any of "WAIT", "WAIT_FOR_CONNECTION".
	Type string `json:"type" api:"required"`
	// Unit of `duration`.
	//
	// Any of "DAY", "MINUTE".
	Unit string `json:"unit" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Duration    respjson.Field
		Type        respjson.Field
		Unit        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SequenceUpdateDraftResponseDefinitionStepObject3) RawJSON() string { return r.JSON.raw }
func (r *SequenceUpdateDraftResponseDefinitionStepObject3) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SequenceUpdateDraftResponsePublishWarning struct {
	// What publishing would change.
	//
	// Any of "delay_shortened", "delay_lengthened", "connection_wait_shortened",
	// "connection_wait_lengthened", "send_schedule_changed",
	// "send_configuration_changed", "send_priority_changed", "senders_changed",
	// "email_sender_pool_emptied", "linkedin_account_pool_emptied",
	// "assigned_sender_removed", "delivery_became_automatic", "invite_note_enabled",
	// "invite_note_disabled", "step_inserted", "automatic_send_step_inserted",
	// "step_removed", "step_type_changed", "step_order_changed",
	// "full_identity_rewrite", "low_identity_overlap".
	Code string `json:"code" api:"required"`
	// What would change, in plain language.
	Message string `json:"message" api:"required"`
	// `destructive` when contacts already in the sequence may skip, repeat or receive
	// a step they would not have; `warning` otherwise.
	//
	// Any of "warning", "destructive".
	Severity string `json:"severity" api:"required"`
	// The steps involved; empty for a sequence-wide change.
	StepIDs []string `json:"stepIds" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Code        respjson.Field
		Message     respjson.Field
		Severity    respjson.Field
		StepIDs     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SequenceUpdateDraftResponsePublishWarning) RawJSON() string { return r.JSON.raw }
func (r *SequenceUpdateDraftResponsePublishWarning) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SequenceUpdateDraftResponseReviewer struct {
	// ID of the member.
	ID string `json:"id" api:"required"`
	// Always `member`.
	//
	// Any of "member".
	ObjectType string `json:"objectType" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		ObjectType  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SequenceUpdateDraftResponseReviewer) RawJSON() string { return r.JSON.raw }
func (r *SequenceUpdateDraftResponseReviewer) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SequenceUpdateDraftResponseSender struct {
	// ID of the sender.
	ID string `json:"id" api:"required"`
	// Always `sender`.
	//
	// Any of "sender".
	ObjectType string `json:"objectType" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		ObjectType  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SequenceUpdateDraftResponseSender) RawJSON() string { return r.JSON.raw }
func (r *SequenceUpdateDraftResponseSender) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SequenceUpdateEnrollmentsResponse struct {
	// One result per distinct contact, additions first, then removals.
	Results []SequenceUpdateEnrollmentsResponseResultUnion `json:"results" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Results     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SequenceUpdateEnrollmentsResponse) RawJSON() string { return r.JSON.raw }
func (r *SequenceUpdateEnrollmentsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// SequenceUpdateEnrollmentsResponseResultUnion contains all possible properties
// and values from [SequenceUpdateEnrollmentsResponseResultObject],
// [SequenceUpdateEnrollmentsResponseResultObject2].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type SequenceUpdateEnrollmentsResponseResultUnion struct {
	ContactID string `json:"contactId"`
	// This field is a union of [string], [any]
	EnrollmentID SequenceUpdateEnrollmentsResponseResultUnionEnrollmentID `json:"enrollmentId"`
	Outcome      string                                                   `json:"outcome"`
	JSON         struct {
		ContactID    respjson.Field
		EnrollmentID respjson.Field
		Outcome      respjson.Field
		raw          string
	} `json:"-"`
}

func (u SequenceUpdateEnrollmentsResponseResultUnion) AsSequenceUpdateEnrollmentsResponseResultObject() (v SequenceUpdateEnrollmentsResponseResultObject) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u SequenceUpdateEnrollmentsResponseResultUnion) AsSequenceUpdateEnrollmentsResponseResultObject2() (v SequenceUpdateEnrollmentsResponseResultObject2) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u SequenceUpdateEnrollmentsResponseResultUnion) RawJSON() string { return u.JSON.raw }

func (r *SequenceUpdateEnrollmentsResponseResultUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// SequenceUpdateEnrollmentsResponseResultUnionEnrollmentID is an implicit subunion
// of [SequenceUpdateEnrollmentsResponseResultUnion].
// SequenceUpdateEnrollmentsResponseResultUnionEnrollmentID provides convenient
// access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [SequenceUpdateEnrollmentsResponseResultUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString
// OfSequenceUpdateEnrollmentsResponseResultObject2EnrollmentID]
type SequenceUpdateEnrollmentsResponseResultUnionEnrollmentID struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [any] instead of an object.
	OfSequenceUpdateEnrollmentsResponseResultObject2EnrollmentID any `json:",inline"`
	JSON                                                         struct {
		OfString                                                     respjson.Field
		OfSequenceUpdateEnrollmentsResponseResultObject2EnrollmentID respjson.Field
		raw                                                          string
	} `json:"-"`
}

func (r *SequenceUpdateEnrollmentsResponseResultUnionEnrollmentID) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SequenceUpdateEnrollmentsResponseResultObject struct {
	ContactID    string `json:"contactId" api:"required"`
	EnrollmentID string `json:"enrollmentId" api:"required"`
	// Any of "enrolled", "already_enrolled", "unenrolled".
	Outcome string `json:"outcome" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ContactID    respjson.Field
		EnrollmentID respjson.Field
		Outcome      respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SequenceUpdateEnrollmentsResponseResultObject) RawJSON() string { return r.JSON.raw }
func (r *SequenceUpdateEnrollmentsResponseResultObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SequenceUpdateEnrollmentsResponseResultObject2 struct {
	ContactID    string `json:"contactId" api:"required"`
	EnrollmentID any    `json:"enrollmentId" api:"required"`
	// Any of "no_active_enrollment".
	Outcome string `json:"outcome" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ContactID    respjson.Field
		EnrollmentID respjson.Field
		Outcome      respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SequenceUpdateEnrollmentsResponseResultObject2) RawJSON() string { return r.JSON.raw }
func (r *SequenceUpdateEnrollmentsResponseResultObject2) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SequenceNewParams struct {
	// Configuration of the first draft. `description` and `steps` are required.
	Definition SequenceNewParamsDefinition `json:"definition,omitzero" api:"required"`
	// Name of the sequence.
	Name string `json:"name" api:"required"`
	// Extra members who may approve this sequence’s `FOR_REVIEW` messages before they
	// send. The sending member and the sequence’s creator always review, whether or
	// not they are listed. A change applies to sends once the draft is published.
	// Member IDs of active members. Omit for none.
	Reviewers []string `json:"reviewers,omitzero"`
	// The senders this sequence sends from; new contacts rotate across them. A change
	// applies to sends once the draft is published. Sender IDs of unarchived senders.
	// Omit for none; activating needs a sender for each channel the steps use.
	Senders []string `json:"senders,omitzero"`
	paramObj
}

func (r SequenceNewParams) MarshalJSON() (data []byte, err error) {
	type shadow SequenceNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SequenceNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Configuration of the first draft. `description` and `steps` are required.
//
// The properties Description, Steps are required.
type SequenceNewParamsDefinition struct {
	// What this version of the sequence is for.
	Description string `json:"description" api:"required"`
	// The ordered step list.
	Steps []SequenceNewParamsDefinitionStepUnion `json:"steps,omitzero" api:"required"`
	// Messaging inputs that frame every step. Omitted fields are unchanged; null
	// clears a field.
	Messaging SequenceNewParamsDefinitionMessaging `json:"messaging,omitzero"`
	// How the whole sequence paces each step type it sends, keyed by sequence step
	// type. Omitted step types are unchanged.
	Pacing SequenceNewParamsDefinitionPacing `json:"pacing,omitzero"`
	// Send priority relative to other sequences.
	//
	// Any of "HIGH", "NORMAL".
	Priority string `json:"priority,omitzero"`
	// When the sequence may send. Omitted fields are unchanged; a new sequence
	// defaults to 9:00–17:00, Monday to Friday, UTC.
	SendWindow SequenceNewParamsDefinitionSendWindow `json:"sendWindow,omitzero"`
	paramObj
}

func (r SequenceNewParamsDefinition) MarshalJSON() (data []byte, err error) {
	type shadow SequenceNewParamsDefinition
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SequenceNewParamsDefinition) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[SequenceNewParamsDefinition](
		"priority", "HIGH", "NORMAL",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type SequenceNewParamsDefinitionStepUnion struct {
	OfSequenceNewsDefinitionStepObject  *SequenceNewParamsDefinitionStepObject  `json:",omitzero,inline"`
	OfSequenceNewsDefinitionStepObject2 *SequenceNewParamsDefinitionStepObject2 `json:",omitzero,inline"`
	OfSequenceNewsDefinitionStepObject3 *SequenceNewParamsDefinitionStepObject3 `json:",omitzero,inline"`
	paramUnion
}

func (u SequenceNewParamsDefinitionStepUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfSequenceNewsDefinitionStepObject, u.OfSequenceNewsDefinitionStepObject2, u.OfSequenceNewsDefinitionStepObject3)
}
func (u *SequenceNewParamsDefinitionStepUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// The properties DeliveryMode, Name, Prompt, Type are required.
type SequenceNewParamsDefinitionStepObject struct {
	// `AUTO` sends without manual approval, subject to the send window and limits;
	// `FOR_REVIEW` requires a person to approve the draft before sending.
	//
	// Any of "AUTO", "FOR_REVIEW".
	DeliveryMode string `json:"deliveryMode,omitzero" api:"required"`
	// Display name of the step.
	Name string `json:"name" api:"required"`
	// Instructions the agent follows when composing this message.
	Prompt string `json:"prompt" api:"required"`
	// Step type.
	//
	// Any of "EMAIL", "LINKEDIN_MESSAGE".
	Type string `json:"type,omitzero" api:"required"`
	paramObj
}

func (r SequenceNewParamsDefinitionStepObject) MarshalJSON() (data []byte, err error) {
	type shadow SequenceNewParamsDefinitionStepObject
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SequenceNewParamsDefinitionStepObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[SequenceNewParamsDefinitionStepObject](
		"deliveryMode", "AUTO", "FOR_REVIEW",
	)
	apijson.RegisterFieldValidator[SequenceNewParamsDefinitionStepObject](
		"type", "EMAIL", "LINKEDIN_MESSAGE",
	)
}

// The properties DeliveryMode, Name, Prompt, Type are required.
type SequenceNewParamsDefinitionStepObject2 struct {
	// `AUTO` sends without manual approval, subject to the send window and limits;
	// `FOR_REVIEW` requires a person to approve the draft before sending.
	//
	// Any of "AUTO", "FOR_REVIEW".
	DeliveryMode string `json:"deliveryMode,omitzero" api:"required"`
	// Display name of the step.
	Name string `json:"name" api:"required"`
	// Instructions the agent follows when composing this message.
	Prompt string `json:"prompt" api:"required"`
	// Step type.
	//
	// Any of "LINKEDIN_INVITE".
	Type string `json:"type,omitzero" api:"required"`
	// Whether the connection request carries a note. Omit to keep the current setting;
	// a new step defaults to false.
	IncludeNote param.Opt[bool] `json:"includeNote,omitzero"`
	paramObj
}

func (r SequenceNewParamsDefinitionStepObject2) MarshalJSON() (data []byte, err error) {
	type shadow SequenceNewParamsDefinitionStepObject2
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SequenceNewParamsDefinitionStepObject2) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[SequenceNewParamsDefinitionStepObject2](
		"deliveryMode", "AUTO", "FOR_REVIEW",
	)
	apijson.RegisterFieldValidator[SequenceNewParamsDefinitionStepObject2](
		"type", "LINKEDIN_INVITE",
	)
}

// The properties Duration, Type, Unit are required.
type SequenceNewParamsDefinitionStepObject3 struct {
	// Length of the wait.
	Duration int64 `json:"duration" api:"required"`
	// `WAIT` pauses for a fixed time; `WAIT_FOR_CONNECTION` waits up to that long for
	// a LinkedIn invite to be accepted.
	//
	// Any of "WAIT", "WAIT_FOR_CONNECTION".
	Type string `json:"type,omitzero" api:"required"`
	// Unit of `duration`.
	//
	// Any of "DAY", "MINUTE".
	Unit string `json:"unit,omitzero" api:"required"`
	paramObj
}

func (r SequenceNewParamsDefinitionStepObject3) MarshalJSON() (data []byte, err error) {
	type shadow SequenceNewParamsDefinitionStepObject3
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SequenceNewParamsDefinitionStepObject3) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[SequenceNewParamsDefinitionStepObject3](
		"type", "WAIT", "WAIT_FOR_CONNECTION",
	)
	apijson.RegisterFieldValidator[SequenceNewParamsDefinitionStepObject3](
		"unit", "DAY", "MINUTE",
	)
}

// Messaging inputs that frame every step. Omitted fields are unchanged; null
// clears a field.
type SequenceNewParamsDefinitionMessaging struct {
	// What the recipient is asked to do.
	CallToAction param.Opt[string] `json:"callToAction,omitzero"`
	// Who the sequence targets.
	Persona param.Opt[string] `json:"persona,omitzero"`
	// Why the outreach is timely.
	Trigger param.Opt[string] `json:"trigger,omitzero"`
	// The value proposition.
	ValueProp param.Opt[string] `json:"valueProp,omitzero"`
	paramObj
}

func (r SequenceNewParamsDefinitionMessaging) MarshalJSON() (data []byte, err error) {
	type shadow SequenceNewParamsDefinitionMessaging
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SequenceNewParamsDefinitionMessaging) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// How the whole sequence paces each step type it sends, keyed by sequence step
// type. Omitted step types are unchanged.
type SequenceNewParamsDefinitionPacing struct {
	// Email steps.
	Email SequenceNewParamsDefinitionPacingEmail `json:"EMAIL,omitzero"`
	// LinkedIn connection requests.
	LinkedinInvite SequenceNewParamsDefinitionPacingLinkedinInvite `json:"LINKEDIN_INVITE,omitzero"`
	// LinkedIn messages.
	LinkedinMessage SequenceNewParamsDefinitionPacingLinkedinMessage `json:"LINKEDIN_MESSAGE,omitzero"`
	paramObj
}

func (r SequenceNewParamsDefinitionPacing) MarshalJSON() (data []byte, err error) {
	type shadow SequenceNewParamsDefinitionPacing
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SequenceNewParamsDefinitionPacing) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Email steps.
//
// The property DailyLimit is required.
type SequenceNewParamsDefinitionPacingEmail struct {
	// Most sends of this step type across the whole sequence in any rolling 24 hours,
	// or null when the sequence sets no daily cap. Workspace and sender limits still
	// apply.
	DailyLimit param.Opt[int64] `json:"dailyLimit,omitzero" api:"required"`
	paramObj
}

func (r SequenceNewParamsDefinitionPacingEmail) MarshalJSON() (data []byte, err error) {
	type shadow SequenceNewParamsDefinitionPacingEmail
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SequenceNewParamsDefinitionPacingEmail) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// LinkedIn connection requests.
//
// The property DailyLimit is required.
type SequenceNewParamsDefinitionPacingLinkedinInvite struct {
	// Most sends of this step type across the whole sequence in any rolling 24 hours,
	// or null when the sequence sets no daily cap. Workspace and sender limits still
	// apply.
	DailyLimit param.Opt[int64] `json:"dailyLimit,omitzero" api:"required"`
	paramObj
}

func (r SequenceNewParamsDefinitionPacingLinkedinInvite) MarshalJSON() (data []byte, err error) {
	type shadow SequenceNewParamsDefinitionPacingLinkedinInvite
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SequenceNewParamsDefinitionPacingLinkedinInvite) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// LinkedIn messages.
//
// The property DailyLimit is required.
type SequenceNewParamsDefinitionPacingLinkedinMessage struct {
	// Most sends of this step type across the whole sequence in any rolling 24 hours,
	// or null when the sequence sets no daily cap. Workspace and sender limits still
	// apply.
	DailyLimit param.Opt[int64] `json:"dailyLimit,omitzero" api:"required"`
	paramObj
}

func (r SequenceNewParamsDefinitionPacingLinkedinMessage) MarshalJSON() (data []byte, err error) {
	type shadow SequenceNewParamsDefinitionPacingLinkedinMessage
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SequenceNewParamsDefinitionPacingLinkedinMessage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// When the sequence may send. Omitted fields are unchanged; a new sequence
// defaults to 9:00–17:00, Monday to Friday, UTC.
type SequenceNewParamsDefinitionSendWindow struct {
	// Local hour sends stop, exclusive (1–24).
	EndHour param.Opt[int64] `json:"endHour,omitzero"`
	// First local hour sends may go out (0–23).
	StartHour param.Opt[int64] `json:"startHour,omitzero"`
	// IANA timezone of the window.
	Timezone param.Opt[string] `json:"timezone,omitzero"`
	// Days sends may go out, 0 = Sunday.
	DaysOfWeek []int64 `json:"daysOfWeek,omitzero"`
	paramObj
}

func (r SequenceNewParamsDefinitionSendWindow) MarshalJSON() (data []byte, err error) {
	type shadow SequenceNewParamsDefinitionSendWindow
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SequenceNewParamsDefinitionSendWindow) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SequenceListParams struct {
	// Maximum number of records to return. Defaults to 25, maximum 25.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Number of records to skip for pagination. Defaults to 0.
	Offset param.Opt[int64] `query:"offset,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [SequenceListParams]'s query parameters as `url.Values`.
func (r SequenceListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type SequenceActivateParams struct {
	Body SequenceActivateParamsBody
	paramObj
}

func (r SequenceActivateParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.Body)
}
func (r *SequenceActivateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SequenceActivateParamsBody struct {
	paramObj
}

func (r SequenceActivateParamsBody) MarshalJSON() (data []byte, err error) {
	type shadow SequenceActivateParamsBody
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SequenceActivateParamsBody) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SequenceDeactivateParams struct {
	Body SequenceDeactivateParamsBody
	paramObj
}

func (r SequenceDeactivateParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.Body)
}
func (r *SequenceDeactivateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SequenceDeactivateParamsBody struct {
	paramObj
}

func (r SequenceDeactivateParamsBody) MarshalJSON() (data []byte, err error) {
	type shadow SequenceDeactivateParamsBody
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SequenceDeactivateParamsBody) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SequenceDiscardDraftParams struct {
	// The draft `version` this request was decided against.
	ExpectedVersion string `json:"expectedVersion" api:"required"`
	paramObj
}

func (r SequenceDiscardDraftParams) MarshalJSON() (data []byte, err error) {
	type shadow SequenceDiscardDraftParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SequenceDiscardDraftParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SequenceListEnrollmentsParams struct {
	// Maximum number of records to return. Defaults to 25, maximum 25.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Number of records to skip for pagination. Defaults to 0.
	Offset param.Opt[int64] `query:"offset,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [SequenceListEnrollmentsParams]'s query parameters as
// `url.Values`.
func (r SequenceListEnrollmentsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type SequenceListMessagesParams struct {
	// Maximum number of records to return. Defaults to 25, maximum 25.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Number of records to skip for pagination. Defaults to 0.
	Offset param.Opt[int64] `query:"offset,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [SequenceListMessagesParams]'s query parameters as
// `url.Values`.
func (r SequenceListMessagesParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type SequencePublishDraftParams struct {
	// The draft `version` this request was decided against.
	ExpectedVersion string `json:"expectedVersion" api:"required"`
	paramObj
}

func (r SequencePublishDraftParams) MarshalJSON() (data []byte, err error) {
	type shadow SequencePublishDraftParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SequencePublishDraftParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SequenceGetEnrollmentParams struct {
	// Unique identifier of the sequence.
	ID string `path:"id" api:"required" json:"-"`
	paramObj
}

type SequenceGetMessageParams struct {
	// Unique identifier of the sequence.
	ID string `path:"id" api:"required" json:"-"`
	paramObj
}

type SequenceUpdateDraftParams struct {
	// The open draft’s `version`, unchanged. Required when a draft is open; omit it
	// only when none is, to start one from the published version.
	ExpectedVersion param.Opt[string] `json:"expectedVersion,omitzero"`
	// Name of the sequence.
	Name param.Opt[string] `json:"name,omitzero"`
	// Fields to change. Omitted fields are unchanged.
	Definition SequenceUpdateDraftParamsDefinition `json:"definition,omitzero"`
	// Extra members who may approve this sequence’s `FOR_REVIEW` messages before they
	// send. The sending member and the sequence’s creator always review, whether or
	// not they are listed. A change applies to sends once the draft is published.
	// Changes the list by adding and removing members; omit to keep it.
	Reviewers SequenceUpdateDraftParamsReviewers `json:"reviewers,omitzero"`
	// The senders this sequence sends from; new contacts rotate across them. A change
	// applies to sends once the draft is published. Changes the list by adding and
	// removing senders; omit to keep it.
	Senders SequenceUpdateDraftParamsSenders `json:"senders,omitzero"`
	paramObj
}

func (r SequenceUpdateDraftParams) MarshalJSON() (data []byte, err error) {
	type shadow SequenceUpdateDraftParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SequenceUpdateDraftParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Fields to change. Omitted fields are unchanged.
type SequenceUpdateDraftParamsDefinition struct {
	// What this version of the sequence is for.
	Description param.Opt[string] `json:"description,omitzero"`
	// Messaging inputs that frame every step. Omitted fields are unchanged; null
	// clears a field.
	Messaging SequenceUpdateDraftParamsDefinitionMessaging `json:"messaging,omitzero"`
	// How the whole sequence paces each step type it sends, keyed by sequence step
	// type. Omitted step types are unchanged.
	Pacing SequenceUpdateDraftParamsDefinitionPacing `json:"pacing,omitzero"`
	// Send priority relative to other sequences.
	//
	// Any of "HIGH", "NORMAL".
	Priority string `json:"priority,omitzero"`
	// When the sequence may send. Omitted fields are unchanged; a new sequence
	// defaults to 9:00–17:00, Monday to Friday, UTC.
	SendWindow SequenceUpdateDraftParamsDefinitionSendWindow `json:"sendWindow,omitzero"`
	// The complete ordered step list, replacing the current one. Keep a step by
	// sending its `id`; omit `id` for a new step.
	Steps []SequenceUpdateDraftParamsDefinitionStepUnion `json:"steps,omitzero"`
	paramObj
}

func (r SequenceUpdateDraftParamsDefinition) MarshalJSON() (data []byte, err error) {
	type shadow SequenceUpdateDraftParamsDefinition
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SequenceUpdateDraftParamsDefinition) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[SequenceUpdateDraftParamsDefinition](
		"priority", "HIGH", "NORMAL",
	)
}

// Messaging inputs that frame every step. Omitted fields are unchanged; null
// clears a field.
type SequenceUpdateDraftParamsDefinitionMessaging struct {
	// What the recipient is asked to do.
	CallToAction param.Opt[string] `json:"callToAction,omitzero"`
	// Who the sequence targets.
	Persona param.Opt[string] `json:"persona,omitzero"`
	// Why the outreach is timely.
	Trigger param.Opt[string] `json:"trigger,omitzero"`
	// The value proposition.
	ValueProp param.Opt[string] `json:"valueProp,omitzero"`
	paramObj
}

func (r SequenceUpdateDraftParamsDefinitionMessaging) MarshalJSON() (data []byte, err error) {
	type shadow SequenceUpdateDraftParamsDefinitionMessaging
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SequenceUpdateDraftParamsDefinitionMessaging) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// How the whole sequence paces each step type it sends, keyed by sequence step
// type. Omitted step types are unchanged.
type SequenceUpdateDraftParamsDefinitionPacing struct {
	// Email steps.
	Email SequenceUpdateDraftParamsDefinitionPacingEmail `json:"EMAIL,omitzero"`
	// LinkedIn connection requests.
	LinkedinInvite SequenceUpdateDraftParamsDefinitionPacingLinkedinInvite `json:"LINKEDIN_INVITE,omitzero"`
	// LinkedIn messages.
	LinkedinMessage SequenceUpdateDraftParamsDefinitionPacingLinkedinMessage `json:"LINKEDIN_MESSAGE,omitzero"`
	paramObj
}

func (r SequenceUpdateDraftParamsDefinitionPacing) MarshalJSON() (data []byte, err error) {
	type shadow SequenceUpdateDraftParamsDefinitionPacing
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SequenceUpdateDraftParamsDefinitionPacing) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Email steps.
//
// The property DailyLimit is required.
type SequenceUpdateDraftParamsDefinitionPacingEmail struct {
	// Most sends of this step type across the whole sequence in any rolling 24 hours,
	// or null when the sequence sets no daily cap. Workspace and sender limits still
	// apply.
	DailyLimit param.Opt[int64] `json:"dailyLimit,omitzero" api:"required"`
	paramObj
}

func (r SequenceUpdateDraftParamsDefinitionPacingEmail) MarshalJSON() (data []byte, err error) {
	type shadow SequenceUpdateDraftParamsDefinitionPacingEmail
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SequenceUpdateDraftParamsDefinitionPacingEmail) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// LinkedIn connection requests.
//
// The property DailyLimit is required.
type SequenceUpdateDraftParamsDefinitionPacingLinkedinInvite struct {
	// Most sends of this step type across the whole sequence in any rolling 24 hours,
	// or null when the sequence sets no daily cap. Workspace and sender limits still
	// apply.
	DailyLimit param.Opt[int64] `json:"dailyLimit,omitzero" api:"required"`
	paramObj
}

func (r SequenceUpdateDraftParamsDefinitionPacingLinkedinInvite) MarshalJSON() (data []byte, err error) {
	type shadow SequenceUpdateDraftParamsDefinitionPacingLinkedinInvite
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SequenceUpdateDraftParamsDefinitionPacingLinkedinInvite) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// LinkedIn messages.
//
// The property DailyLimit is required.
type SequenceUpdateDraftParamsDefinitionPacingLinkedinMessage struct {
	// Most sends of this step type across the whole sequence in any rolling 24 hours,
	// or null when the sequence sets no daily cap. Workspace and sender limits still
	// apply.
	DailyLimit param.Opt[int64] `json:"dailyLimit,omitzero" api:"required"`
	paramObj
}

func (r SequenceUpdateDraftParamsDefinitionPacingLinkedinMessage) MarshalJSON() (data []byte, err error) {
	type shadow SequenceUpdateDraftParamsDefinitionPacingLinkedinMessage
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SequenceUpdateDraftParamsDefinitionPacingLinkedinMessage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// When the sequence may send. Omitted fields are unchanged; a new sequence
// defaults to 9:00–17:00, Monday to Friday, UTC.
type SequenceUpdateDraftParamsDefinitionSendWindow struct {
	// Local hour sends stop, exclusive (1–24).
	EndHour param.Opt[int64] `json:"endHour,omitzero"`
	// First local hour sends may go out (0–23).
	StartHour param.Opt[int64] `json:"startHour,omitzero"`
	// IANA timezone of the window.
	Timezone param.Opt[string] `json:"timezone,omitzero"`
	// Days sends may go out, 0 = Sunday.
	DaysOfWeek []int64 `json:"daysOfWeek,omitzero"`
	paramObj
}

func (r SequenceUpdateDraftParamsDefinitionSendWindow) MarshalJSON() (data []byte, err error) {
	type shadow SequenceUpdateDraftParamsDefinitionSendWindow
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SequenceUpdateDraftParamsDefinitionSendWindow) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type SequenceUpdateDraftParamsDefinitionStepUnion struct {
	OfSequenceUpdateDraftsDefinitionStepObject  *SequenceUpdateDraftParamsDefinitionStepObject  `json:",omitzero,inline"`
	OfSequenceUpdateDraftsDefinitionStepObject2 *SequenceUpdateDraftParamsDefinitionStepObject2 `json:",omitzero,inline"`
	OfSequenceUpdateDraftsDefinitionStepObject3 *SequenceUpdateDraftParamsDefinitionStepObject3 `json:",omitzero,inline"`
	paramUnion
}

func (u SequenceUpdateDraftParamsDefinitionStepUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfSequenceUpdateDraftsDefinitionStepObject, u.OfSequenceUpdateDraftsDefinitionStepObject2, u.OfSequenceUpdateDraftsDefinitionStepObject3)
}
func (u *SequenceUpdateDraftParamsDefinitionStepUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// The properties DeliveryMode, Name, Prompt, Type are required.
type SequenceUpdateDraftParamsDefinitionStepObject struct {
	// `AUTO` sends without manual approval, subject to the send window and limits;
	// `FOR_REVIEW` requires a person to approve the draft before sending.
	//
	// Any of "AUTO", "FOR_REVIEW".
	DeliveryMode string `json:"deliveryMode,omitzero" api:"required"`
	// Display name of the step.
	Name string `json:"name" api:"required"`
	// Instructions the agent follows when composing this message.
	Prompt string `json:"prompt" api:"required"`
	// Step type.
	//
	// Any of "EMAIL", "LINKEDIN_MESSAGE".
	Type string `json:"type,omitzero" api:"required"`
	// ID of a step in the current draft to keep. Omit for a new step. Kept steps must
	// stay in their relative order and keep their type.
	ID param.Opt[string] `json:"id,omitzero"`
	paramObj
}

func (r SequenceUpdateDraftParamsDefinitionStepObject) MarshalJSON() (data []byte, err error) {
	type shadow SequenceUpdateDraftParamsDefinitionStepObject
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SequenceUpdateDraftParamsDefinitionStepObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[SequenceUpdateDraftParamsDefinitionStepObject](
		"deliveryMode", "AUTO", "FOR_REVIEW",
	)
	apijson.RegisterFieldValidator[SequenceUpdateDraftParamsDefinitionStepObject](
		"type", "EMAIL", "LINKEDIN_MESSAGE",
	)
}

// The properties DeliveryMode, Name, Prompt, Type are required.
type SequenceUpdateDraftParamsDefinitionStepObject2 struct {
	// `AUTO` sends without manual approval, subject to the send window and limits;
	// `FOR_REVIEW` requires a person to approve the draft before sending.
	//
	// Any of "AUTO", "FOR_REVIEW".
	DeliveryMode string `json:"deliveryMode,omitzero" api:"required"`
	// Display name of the step.
	Name string `json:"name" api:"required"`
	// Instructions the agent follows when composing this message.
	Prompt string `json:"prompt" api:"required"`
	// Step type.
	//
	// Any of "LINKEDIN_INVITE".
	Type string `json:"type,omitzero" api:"required"`
	// ID of a step in the current draft to keep. Omit for a new step. Kept steps must
	// stay in their relative order and keep their type.
	ID param.Opt[string] `json:"id,omitzero"`
	// Whether the connection request carries a note. Omit to keep the current setting;
	// a new step defaults to false.
	IncludeNote param.Opt[bool] `json:"includeNote,omitzero"`
	paramObj
}

func (r SequenceUpdateDraftParamsDefinitionStepObject2) MarshalJSON() (data []byte, err error) {
	type shadow SequenceUpdateDraftParamsDefinitionStepObject2
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SequenceUpdateDraftParamsDefinitionStepObject2) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[SequenceUpdateDraftParamsDefinitionStepObject2](
		"deliveryMode", "AUTO", "FOR_REVIEW",
	)
	apijson.RegisterFieldValidator[SequenceUpdateDraftParamsDefinitionStepObject2](
		"type", "LINKEDIN_INVITE",
	)
}

// The properties Duration, Type, Unit are required.
type SequenceUpdateDraftParamsDefinitionStepObject3 struct {
	// Length of the wait.
	Duration int64 `json:"duration" api:"required"`
	// `WAIT` pauses for a fixed time; `WAIT_FOR_CONNECTION` waits up to that long for
	// a LinkedIn invite to be accepted.
	//
	// Any of "WAIT", "WAIT_FOR_CONNECTION".
	Type string `json:"type,omitzero" api:"required"`
	// Unit of `duration`.
	//
	// Any of "DAY", "MINUTE".
	Unit string `json:"unit,omitzero" api:"required"`
	// ID of a step in the current draft to keep. Omit for a new step. Kept steps must
	// stay in their relative order and keep their type.
	ID param.Opt[string] `json:"id,omitzero"`
	paramObj
}

func (r SequenceUpdateDraftParamsDefinitionStepObject3) MarshalJSON() (data []byte, err error) {
	type shadow SequenceUpdateDraftParamsDefinitionStepObject3
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SequenceUpdateDraftParamsDefinitionStepObject3) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[SequenceUpdateDraftParamsDefinitionStepObject3](
		"type", "WAIT", "WAIT_FOR_CONNECTION",
	)
	apijson.RegisterFieldValidator[SequenceUpdateDraftParamsDefinitionStepObject3](
		"unit", "DAY", "MINUTE",
	)
}

// Extra members who may approve this sequence’s `FOR_REVIEW` messages before they
// send. The sending member and the sequence’s creator always review, whether or
// not they are listed. A change applies to sends once the draft is published.
// Changes the list by adding and removing members; omit to keep it.
type SequenceUpdateDraftParamsReviewers struct {
	// Member IDs of active members to add. A member already listed is unchanged.
	Add []string `json:"add,omitzero"`
	// Member IDs to remove, including members who have left the workspace. A member
	// not listed is unchanged.
	Remove []string `json:"remove,omitzero"`
	paramObj
}

func (r SequenceUpdateDraftParamsReviewers) MarshalJSON() (data []byte, err error) {
	type shadow SequenceUpdateDraftParamsReviewers
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SequenceUpdateDraftParamsReviewers) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The senders this sequence sends from; new contacts rotate across them. A change
// applies to sends once the draft is published. Changes the list by adding and
// removing senders; omit to keep it.
type SequenceUpdateDraftParamsSenders struct {
	// Sender IDs of unarchived senders to add. A sender already listed is unchanged.
	Add []string `json:"add,omitzero"`
	// Sender IDs to remove, including archived senders. A sender not listed is
	// unchanged.
	Remove []string `json:"remove,omitzero"`
	paramObj
}

func (r SequenceUpdateDraftParamsSenders) MarshalJSON() (data []byte, err error) {
	type shadow SequenceUpdateDraftParamsSenders
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SequenceUpdateDraftParamsSenders) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SequenceUpdateEnrollmentsParams struct {
	// Contact IDs to enroll. Open enrollments are unchanged; contacts whose prior
	// enrollment ended are enrolled again.
	Add []string `json:"add,omitzero"`
	// Contact IDs to unenroll. Cancels open enrollments and preserves their history.
	Remove []string `json:"remove,omitzero"`
	paramObj
}

func (r SequenceUpdateEnrollmentsParams) MarshalJSON() (data []byte, err error) {
	type shadow SequenceUpdateEnrollmentsParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SequenceUpdateEnrollmentsParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
