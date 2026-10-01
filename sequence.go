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

// Retrieves a sequence by its ID. Published sequences include their steps and send
// settings; an unpublished draft has a null definition.
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

// Returns a paginated list of sequences, including unpublished drafts, ordered by
// `createdAt` descending. An unpublished draft has a null definition. List results
// update asynchronously and may lag recent changes. Use `offset` and `limit` to
// paginate, and bracketed filters for status, channels, and creation or activation
// dates.
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

// Returns a paginated list of live enrollments in a published sequence, including
// completed and cancelled history. Results are ordered by `enrolledAt` descending
// and update asynchronously. Use `offset` and `limit` to paginate, and bracketed
// field filters to narrow results. Test-run enrollments are excluded.
//
// Supported filters: `status[equal]` (ACTIVE, NEEDS_REVIEW, NEEDS_ATTENTION,
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
// and LinkedIn in a published sequence. Only the first reply per enrollment is
// returned, not every reply or one per channel. Superseded outbound attempts are
// excluded. Results have no message content and update asynchronously.
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

// Retrieves one contact's enrollment in a published sequence: where it is in the
// cadence and its latest activity. Test-run enrollments are not returned. Status
// and remainingSends can be null until a current indexed value is available.
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
// note returns `content: { subject: null, body: null }`. Message visibility and
// status update asynchronously.
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

// Adds and removes contacts in one atomic request. If any new enrollment is
// invalid, nothing changes. Provide add, remove, or both, with at most 25 contact
// IDs total. Duplicate IDs within a list are ignored; an ID cannot appear in both
// lists. Published, paused sequences accept enrollment and wait for activation.
// Draft sequences return 404. Already-open enrollments remain unchanged, including
// paused enrollments. Contacts whose prior enrollment completed, was cancelled, or
// failed are enrolled again. Removing a contact with no open enrollment returns
// no_active_enrollment. Execution starts asynchronously after enrollment.
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
	// Always `sequenceEnrollment`.
	//
	// Any of "sequenceEnrollment".
	ObjectType string `json:"objectType" api:"required"`
	// Unsent email and LinkedIn steps remaining, including connection requests; null
	// on retrieval when a current indexed value is unavailable.
	RemainingSends int64 `json:"remainingSends" api:"required"`
	// When this enrollment first received a reply across email or LinkedIn.
	RepliedAt string `json:"repliedAt" api:"required"`
	// Display status of the enrollment, or null on retrieval when the index is
	// unavailable or has not caught up with the enrollment lifecycle.
	//
	// Any of "ACTIVE", "NEEDS_REVIEW", "NEEDS_ATTENTION", "REPLIED", "COMPLETED",
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
		ObjectType       respjson.Field
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
	ObjectType   string `json:"objectType"`
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
		ObjectType       respjson.Field
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
	// Always `sequenceMessage`.
	//
	// Any of "sequenceMessage".
	ObjectType string `json:"objectType" api:"required"`
	OpenedAt   string `json:"openedAt" api:"required"`
	ReadAt     string `json:"readAt" api:"required"`
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
		ObjectType       respjson.Field
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
	// Always `sequenceMessage`.
	//
	// Any of "sequenceMessage".
	ObjectType string `json:"objectType" api:"required"`
	ReceivedAt string `json:"receivedAt" api:"required"`
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
		ObjectType       respjson.Field
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
	// `DRAFT` has not been published; `ACTIVE` is enabled for sending; `DEACTIVATED`
	// is published but paused.
	//
	// Any of "DRAFT", "ACTIVE", "DEACTIVATED".
	Status string `json:"status" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID               respjson.Field
		CreatedAt        respjson.Field
		Definition       respjson.Field
		FirstActivatedAt respjson.Field
		LastActivatedAt  respjson.Field
		Name             respjson.Field
		ObjectType       respjson.Field
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
	// Channels the sequence sends on.
	//
	// Any of "EMAIL", "LINKEDIN".
	Channels []string `json:"channels" api:"required"`
	// Per-day send caps across the whole sequence, by step type.
	DailyLimits SequenceListResponseDataDefinitionDailyLimits `json:"dailyLimits" api:"required"`
	// Description of the published version.
	Description string `json:"description" api:"required"`
	// Messaging inputs that frame every step.
	Messaging SequenceListResponseDataDefinitionMessaging `json:"messaging" api:"required"`
	// Send priority relative to other sequences.
	//
	// Any of "HIGH", "NORMAL".
	Priority string `json:"priority" api:"required"`
	// When the sequence may send.
	SendWindow SequenceListResponseDataDefinitionSendWindow `json:"sendWindow" api:"required"`
	// Published email, LinkedIn and wait steps in execution order; call steps are not
	// exposed in this release.
	Steps []SequenceListResponseDataDefinitionStepUnion `json:"steps" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Channels    respjson.Field
		DailyLimits respjson.Field
		Description respjson.Field
		Messaging   respjson.Field
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

// Per-day send caps across the whole sequence, by step type.
type SequenceListResponseDataDefinitionDailyLimits struct {
	// Sequence-level send cap for a one-day window, or null when no daily cap is
	// configured; workspace and sender limits still apply.
	Email int64 `json:"EMAIL" api:"required"`
	// Sequence-level send cap for a one-day window, or null when no daily cap is
	// configured; workspace and sender limits still apply.
	LinkedinInvite int64 `json:"LINKEDIN_INVITE" api:"required"`
	// Sequence-level send cap for a one-day window, or null when no daily cap is
	// configured; workspace and sender limits still apply.
	LinkedinMessage int64 `json:"LINKEDIN_MESSAGE" api:"required"`
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
func (r SequenceListResponseDataDefinitionDailyLimits) RawJSON() string { return r.JSON.raw }
func (r *SequenceListResponseDataDefinitionDailyLimits) UnmarshalJSON(data []byte) error {
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
	// Whether the connection request carries a note. Null for sequences authored
	// before this setting existed.
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

type SequenceListResponseObject string

const (
	SequenceListResponseObjectList SequenceListResponseObject = "list"
)

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
	// Always `sequenceEnrollment`.
	//
	// Any of "sequenceEnrollment".
	ObjectType SequenceRetrieveEnrollmentResponseObjectType `json:"objectType" api:"required"`
	// Unsent email and LinkedIn steps remaining, including connection requests; null
	// on retrieval when a current indexed value is unavailable.
	RemainingSends int64 `json:"remainingSends" api:"required"`
	// When this enrollment first received a reply across email or LinkedIn.
	RepliedAt string `json:"repliedAt" api:"required"`
	// Display status of the enrollment, or null on retrieval when the index is
	// unavailable or has not caught up with the enrollment lifecycle.
	//
	// Any of "ACTIVE", "NEEDS_REVIEW", "NEEDS_ATTENTION", "REPLIED", "COMPLETED",
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
		ObjectType       respjson.Field
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

// Always `sequenceEnrollment`.
type SequenceRetrieveEnrollmentResponseObjectType string

const (
	SequenceRetrieveEnrollmentResponseObjectTypeSequenceEnrollment SequenceRetrieveEnrollmentResponseObjectType = "sequenceEnrollment"
)

// Display status of the enrollment, or null on retrieval when the index is
// unavailable or has not caught up with the enrollment lifecycle.
type SequenceRetrieveEnrollmentResponseStatus string

const (
	SequenceRetrieveEnrollmentResponseStatusActive         SequenceRetrieveEnrollmentResponseStatus = "ACTIVE"
	SequenceRetrieveEnrollmentResponseStatusNeedsReview    SequenceRetrieveEnrollmentResponseStatus = "NEEDS_REVIEW"
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
	ObjectType   string                                      `json:"objectType"`
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
		ObjectType       respjson.Field
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
	// Always `sequenceMessage`.
	//
	// Any of "sequenceMessage".
	ObjectType string `json:"objectType" api:"required"`
	OpenedAt   string `json:"openedAt" api:"required"`
	ReadAt     string `json:"readAt" api:"required"`
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
		ObjectType       respjson.Field
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
	// Always `sequenceMessage`.
	//
	// Any of "sequenceMessage".
	ObjectType string `json:"objectType" api:"required"`
	ReceivedAt string `json:"receivedAt" api:"required"`
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
		ObjectType       respjson.Field
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

type SequenceRetrieveResponse struct {
	// Unique identifier of the sequence.
	ID string `json:"id" api:"required"`
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
	// `DRAFT` has not been published; `ACTIVE` is enabled for sending; `DEACTIVATED`
	// is published but paused.
	//
	// Any of "DRAFT", "ACTIVE", "DEACTIVATED".
	Status SequenceRetrieveResponseStatus `json:"status" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID               respjson.Field
		CreatedAt        respjson.Field
		Definition       respjson.Field
		FirstActivatedAt respjson.Field
		LastActivatedAt  respjson.Field
		Name             respjson.Field
		ObjectType       respjson.Field
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
	// Channels the sequence sends on.
	//
	// Any of "EMAIL", "LINKEDIN".
	Channels []string `json:"channels" api:"required"`
	// Per-day send caps across the whole sequence, by step type.
	DailyLimits SequenceRetrieveResponseDefinitionDailyLimits `json:"dailyLimits" api:"required"`
	// Description of the published version.
	Description string `json:"description" api:"required"`
	// Messaging inputs that frame every step.
	Messaging SequenceRetrieveResponseDefinitionMessaging `json:"messaging" api:"required"`
	// Send priority relative to other sequences.
	//
	// Any of "HIGH", "NORMAL".
	Priority string `json:"priority" api:"required"`
	// When the sequence may send.
	SendWindow SequenceRetrieveResponseDefinitionSendWindow `json:"sendWindow" api:"required"`
	// Published email, LinkedIn and wait steps in execution order; call steps are not
	// exposed in this release.
	Steps []SequenceRetrieveResponseDefinitionStepUnion `json:"steps" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Channels    respjson.Field
		DailyLimits respjson.Field
		Description respjson.Field
		Messaging   respjson.Field
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

// Per-day send caps across the whole sequence, by step type.
type SequenceRetrieveResponseDefinitionDailyLimits struct {
	// Sequence-level send cap for a one-day window, or null when no daily cap is
	// configured; workspace and sender limits still apply.
	Email int64 `json:"EMAIL" api:"required"`
	// Sequence-level send cap for a one-day window, or null when no daily cap is
	// configured; workspace and sender limits still apply.
	LinkedinInvite int64 `json:"LINKEDIN_INVITE" api:"required"`
	// Sequence-level send cap for a one-day window, or null when no daily cap is
	// configured; workspace and sender limits still apply.
	LinkedinMessage int64 `json:"LINKEDIN_MESSAGE" api:"required"`
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
func (r SequenceRetrieveResponseDefinitionDailyLimits) RawJSON() string { return r.JSON.raw }
func (r *SequenceRetrieveResponseDefinitionDailyLimits) UnmarshalJSON(data []byte) error {
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
	// Whether the connection request carries a note. Null for sequences authored
	// before this setting existed.
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

// `DRAFT` has not been published; `ACTIVE` is enabled for sending; `DEACTIVATED`
// is published but paused.
type SequenceRetrieveResponseStatus string

const (
	SequenceRetrieveResponseStatusDraft       SequenceRetrieveResponseStatus = "DRAFT"
	SequenceRetrieveResponseStatusActive      SequenceRetrieveResponseStatus = "ACTIVE"
	SequenceRetrieveResponseStatusDeactivated SequenceRetrieveResponseStatus = "DEACTIVATED"
)

type SequenceUpdateEnrollmentsResponse struct {
	// One result per distinct contact, additions first, then removals; enrollment IDs
	// can be retrieved immediately, but indexed status, remaining sends and list
	// visibility may lag.
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
