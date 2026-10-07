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
	shimjson "github.com/Lightfld/lightfield-go/internal/encoding/json"
	"github.com/Lightfld/lightfield-go/internal/requestconfig"
	"github.com/Lightfld/lightfield-go/option"
	"github.com/Lightfld/lightfield-go/packages/param"
	"github.com/Lightfld/lightfield-go/packages/respjson"
)

// > [!WARNING] The Senders API is in Preview. Endpoints, parameters, and response
// > schemas may change, including breaking changes, as we incorporate feedback.
//
// Senders are the email mailboxes and LinkedIn accounts sequences send from.
// Creating, connecting and archiving a sender happen in the Lightfield app.
//
// SenderService contains methods and other services that help with interacting
// with the Lightfield API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSenderService] method instead.
type SenderService struct {
	Options []option.RequestOption
}

// NewSenderService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewSenderService(opts ...option.RequestOption) (r SenderService) {
	r = SenderService{}
	r.Options = opts
	return
}

// Retrieves a sender by its ID, archived ones included.
//
// **[Required scope](/using-the-api/scopes/):** `senders:read`
//
// **[Rate limit category](/using-the-api/rate-limits/):** Read
func (r *SenderService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *SenderRetrieveResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/senders/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Replaces a sender's pacing and returns the updated sender. Send exactly the
// sender's step types, with every value: read the sender first to change only some
// of them. Concurrent updates are last-write-wins.
//
// Each step type must stay within the ceilings the workspace's own limits have,
// and `maxGapSeconds` can't be below `minGapSeconds`: `EMAIL` on an OAuth sender
// `dailyLimit` at least 1 and gaps 60 to 3600 seconds; `EMAIL` on an SMTP sender
// `dailyLimit` 1 to 100 and gaps 60 to 3600 seconds; `LINKEDIN_INVITE`
// `dailyLimit` 1 to 30 and gaps 120 to 3600 seconds; `LINKEDIN_MESSAGE`
// `dailyLimit` 1 to 60 and gaps 120 to 3600 seconds. Out-of-range values, or step
// types the sender doesn't send, return 400.
//
// Returns 409 if the sender is archived; archived senders can be read but not
// updated.
//
// **[Required scope](/using-the-api/scopes/):** `senders:update`
//
// **[Rate limit category](/using-the-api/rate-limits/):** Write
func (r *SenderService) Update(ctx context.Context, id string, body SenderUpdateParams, opts ...option.RequestOption) (res *SenderUpdateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/senders/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Returns a paginated list of the email and LinkedIn senders that sequences send
// from, archived ones included, ordered by `createdAt` descending. Use `offset`
// and `limit` to paginate, and bracketed filters to narrow the list.
//
// Supported filters: `channel[equal]` (EMAIL, LINKEDIN), `health[equal]`
// (UNVERIFIED, HEALTHY, UNHEALTHY), `owner[equal]` (a member ID), and
// `archived[equal]` (true, false). Filters are combined with AND.
//
// **[Required scope](/using-the-api/scopes/):** `senders:read`
//
// **[Rate limit category](/using-the-api/rate-limits/):** Search
func (r *SenderService) List(ctx context.Context, query SenderListParams, opts ...option.RequestOption) (res *SenderListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/senders"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Clears a health reason that sending recorded, so the sender can be picked for
// new sends again, and returns the updated sender. A sender with nothing to clear
// is returned unchanged.
//
// Clears `healthReason` `BOUNCE_THRESHOLD`, `RATE_LIMIT`, `INVITE_LIMIT` on an
// email sender, and `RATE_LIMIT`, `INVITE_LIMIT`, `UNKNOWN` on a LinkedIn sender.
// Clearing `BOUNCE_THRESHOLD` turns bounce protection off until hard bounces
// recur. Contacts the sender paused stay paused until they are resumed, and a
// LinkedIn limit may also have deactivated sequences that list the sender;
// clearing it does not reactivate them.
//
// Returns 409 for any other health reason, which clears only once its cause is
// fixed, for example reconnecting the account or fixing DNS; for an archived
// sender; or when the sender's health changed while the request ran.
//
// **[Required scope](/using-the-api/scopes/):** `senders:update`
//
// **[Rate limit category](/using-the-api/rate-limits/):** Write
func (r *SenderService) MarkHealthy(ctx context.Context, id string, body SenderMarkHealthyParams, opts ...option.RequestOption) (res *SenderMarkHealthyResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/senders/%s/mark_healthy", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

type SenderListResponse struct {
	Data []SenderListResponseData `json:"data" api:"required"`
	// Any of "list".
	Object     SenderListResponseObject `json:"object" api:"required"`
	TotalCount int64                    `json:"totalCount" api:"required"`
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
func (r SenderListResponse) RawJSON() string { return r.JSON.raw }
func (r *SenderListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SenderListResponseData struct {
	// Unique identifier of the sender.
	ID string `json:"id" api:"required"`
	// When the sender was archived, or null if it is active. Archived senders never
	// send.
	ArchivedAt string `json:"archivedAt" api:"required"`
	// Whether this sender sends email or LinkedIn messages.
	//
	// Any of "EMAIL", "LINKEDIN".
	Channel string `json:"channel" api:"required"`
	// When the sender was created.
	CreatedAt string `json:"createdAt" api:"required"`
	// The address email from this sender is sent from. Null for LinkedIn senders, or
	// when the connected account has no address.
	EmailAddress string `json:"emailAddress" api:"required"`
	// `HEALTHY` can send. `UNHEALTHY` cannot send until `healthReason` is resolved.
	// `UNVERIFIED` means no health verdict yet: an email sender in this state cannot
	// send until it passes a health check, while a LinkedIn sender is never
	// health-checked, so a working LinkedIn sender is usually `UNVERIFIED` and can
	// send.
	//
	// Any of "UNVERIFIED", "HEALTHY", "UNHEALTHY".
	Health string `json:"health" api:"required"`
	// Why the sender is `UNHEALTHY` or `UNVERIFIED`, for example
	// `ACCOUNT_DISCONNECTED`, `BOUNCE_THRESHOLD` or `INVITE_LIMIT`; null when healthy.
	HealthReason string `json:"healthReason" api:"required"`
	// Always `sender`.
	//
	// Any of "sender".
	ObjectType string `json:"objectType" api:"required"`
	// Member ID of the person who owns this sender, or null when no workspace
	// membership can be resolved.
	Owner string `json:"owner" api:"required"`
	// How this sender paces each step type it sends, keyed by sequence step type:
	// `EMAIL` for an email sender; `LINKEDIN_INVITE` and `LINKEDIN_MESSAGE` for a
	// LinkedIn sender.
	Pacing SenderListResponseDataPacing `json:"pacing" api:"required"`
	// SMTP senders only: the connected mailbox replies are sent to and read from. Null
	// for other senders.
	ReplyToAddress string `json:"replyToAddress" api:"required"`
	// How an email sender sends: `SMTP` through its own mail server, `OAUTH` through
	// its connected Google or Microsoft account. Null for LinkedIn senders.
	//
	// Any of "SMTP", "OAUTH".
	Transport string `json:"transport" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID             respjson.Field
		ArchivedAt     respjson.Field
		Channel        respjson.Field
		CreatedAt      respjson.Field
		EmailAddress   respjson.Field
		Health         respjson.Field
		HealthReason   respjson.Field
		ObjectType     respjson.Field
		Owner          respjson.Field
		Pacing         respjson.Field
		ReplyToAddress respjson.Field
		Transport      respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SenderListResponseData) RawJSON() string { return r.JSON.raw }
func (r *SenderListResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// How this sender paces each step type it sends, keyed by sequence step type:
// `EMAIL` for an email sender; `LINKEDIN_INVITE` and `LINKEDIN_MESSAGE` for a
// LinkedIn sender.
type SenderListResponseDataPacing struct {
	// Email steps. Present only on an email sender.
	Email SenderListResponseDataPacingEmail `json:"EMAIL"`
	// LinkedIn connection requests. Present only on a LinkedIn sender.
	LinkedinInvite SenderListResponseDataPacingLinkedinInvite `json:"LINKEDIN_INVITE"`
	// LinkedIn messages. Present only on a LinkedIn sender.
	LinkedinMessage SenderListResponseDataPacingLinkedinMessage `json:"LINKEDIN_MESSAGE"`
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
func (r SenderListResponseDataPacing) RawJSON() string { return r.JSON.raw }
func (r *SenderListResponseDataPacing) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Email steps. Present only on an email sender.
type SenderListResponseDataPacingEmail struct {
	// Most sends of this step type in any rolling 24 hours.
	DailyLimit int64 `json:"dailyLimit" api:"required"`
	// Upper bound of the minimum spacing, in seconds. Before each send, a minimum
	// spacing is drawn at random between `minGapSeconds` and this value, and the send
	// waits until that long after the previous one. Sends can be further apart, for
	// example outside the send window or once `dailyLimit` is reached. Null means the
	// minimum spacing is always `minGapSeconds`.
	MaxGapSeconds int64 `json:"maxGapSeconds" api:"required"`
	// Lower bound of the minimum spacing between two of its sends, in seconds.
	MinGapSeconds int64 `json:"minGapSeconds" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DailyLimit    respjson.Field
		MaxGapSeconds respjson.Field
		MinGapSeconds respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SenderListResponseDataPacingEmail) RawJSON() string { return r.JSON.raw }
func (r *SenderListResponseDataPacingEmail) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// LinkedIn connection requests. Present only on a LinkedIn sender.
type SenderListResponseDataPacingLinkedinInvite struct {
	// Most sends of this step type in any rolling 24 hours.
	DailyLimit int64 `json:"dailyLimit" api:"required"`
	// Upper bound of the minimum spacing, in seconds. Before each send, a minimum
	// spacing is drawn at random between `minGapSeconds` and this value, and the send
	// waits until that long after the previous one. Sends can be further apart, for
	// example outside the send window or once `dailyLimit` is reached. Null means the
	// minimum spacing is always `minGapSeconds`.
	MaxGapSeconds int64 `json:"maxGapSeconds" api:"required"`
	// Lower bound of the minimum spacing between two of its sends, in seconds.
	MinGapSeconds int64 `json:"minGapSeconds" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DailyLimit    respjson.Field
		MaxGapSeconds respjson.Field
		MinGapSeconds respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SenderListResponseDataPacingLinkedinInvite) RawJSON() string { return r.JSON.raw }
func (r *SenderListResponseDataPacingLinkedinInvite) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// LinkedIn messages. Present only on a LinkedIn sender.
type SenderListResponseDataPacingLinkedinMessage struct {
	// Most sends of this step type in any rolling 24 hours.
	DailyLimit int64 `json:"dailyLimit" api:"required"`
	// Upper bound of the minimum spacing, in seconds. Before each send, a minimum
	// spacing is drawn at random between `minGapSeconds` and this value, and the send
	// waits until that long after the previous one. Sends can be further apart, for
	// example outside the send window or once `dailyLimit` is reached. Null means the
	// minimum spacing is always `minGapSeconds`.
	MaxGapSeconds int64 `json:"maxGapSeconds" api:"required"`
	// Lower bound of the minimum spacing between two of its sends, in seconds.
	MinGapSeconds int64 `json:"minGapSeconds" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DailyLimit    respjson.Field
		MaxGapSeconds respjson.Field
		MinGapSeconds respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SenderListResponseDataPacingLinkedinMessage) RawJSON() string { return r.JSON.raw }
func (r *SenderListResponseDataPacingLinkedinMessage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SenderListResponseObject string

const (
	SenderListResponseObjectList SenderListResponseObject = "list"
)

type SenderMarkHealthyResponse struct {
	// Unique identifier of the sender.
	ID string `json:"id" api:"required"`
	// When the sender was archived, or null if it is active. Archived senders never
	// send.
	ArchivedAt string `json:"archivedAt" api:"required"`
	// Whether this sender sends email or LinkedIn messages.
	//
	// Any of "EMAIL", "LINKEDIN".
	Channel SenderMarkHealthyResponseChannel `json:"channel" api:"required"`
	// When the sender was created.
	CreatedAt string `json:"createdAt" api:"required"`
	// The address email from this sender is sent from. Null for LinkedIn senders, or
	// when the connected account has no address.
	EmailAddress string `json:"emailAddress" api:"required"`
	// `HEALTHY` can send. `UNHEALTHY` cannot send until `healthReason` is resolved.
	// `UNVERIFIED` means no health verdict yet: an email sender in this state cannot
	// send until it passes a health check, while a LinkedIn sender is never
	// health-checked, so a working LinkedIn sender is usually `UNVERIFIED` and can
	// send.
	//
	// Any of "UNVERIFIED", "HEALTHY", "UNHEALTHY".
	Health SenderMarkHealthyResponseHealth `json:"health" api:"required"`
	// Why the sender is `UNHEALTHY` or `UNVERIFIED`, for example
	// `ACCOUNT_DISCONNECTED`, `BOUNCE_THRESHOLD` or `INVITE_LIMIT`; null when healthy.
	HealthReason string `json:"healthReason" api:"required"`
	// Always `sender`.
	//
	// Any of "sender".
	ObjectType SenderMarkHealthyResponseObjectType `json:"objectType" api:"required"`
	// Member ID of the person who owns this sender, or null when no workspace
	// membership can be resolved.
	Owner string `json:"owner" api:"required"`
	// How this sender paces each step type it sends, keyed by sequence step type:
	// `EMAIL` for an email sender; `LINKEDIN_INVITE` and `LINKEDIN_MESSAGE` for a
	// LinkedIn sender.
	Pacing SenderMarkHealthyResponsePacing `json:"pacing" api:"required"`
	// SMTP senders only: the connected mailbox replies are sent to and read from. Null
	// for other senders.
	ReplyToAddress string `json:"replyToAddress" api:"required"`
	// How an email sender sends: `SMTP` through its own mail server, `OAUTH` through
	// its connected Google or Microsoft account. Null for LinkedIn senders.
	//
	// Any of "SMTP", "OAUTH".
	Transport SenderMarkHealthyResponseTransport `json:"transport" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID             respjson.Field
		ArchivedAt     respjson.Field
		Channel        respjson.Field
		CreatedAt      respjson.Field
		EmailAddress   respjson.Field
		Health         respjson.Field
		HealthReason   respjson.Field
		ObjectType     respjson.Field
		Owner          respjson.Field
		Pacing         respjson.Field
		ReplyToAddress respjson.Field
		Transport      respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SenderMarkHealthyResponse) RawJSON() string { return r.JSON.raw }
func (r *SenderMarkHealthyResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Whether this sender sends email or LinkedIn messages.
type SenderMarkHealthyResponseChannel string

const (
	SenderMarkHealthyResponseChannelEmail    SenderMarkHealthyResponseChannel = "EMAIL"
	SenderMarkHealthyResponseChannelLinkedin SenderMarkHealthyResponseChannel = "LINKEDIN"
)

// `HEALTHY` can send. `UNHEALTHY` cannot send until `healthReason` is resolved.
// `UNVERIFIED` means no health verdict yet: an email sender in this state cannot
// send until it passes a health check, while a LinkedIn sender is never
// health-checked, so a working LinkedIn sender is usually `UNVERIFIED` and can
// send.
type SenderMarkHealthyResponseHealth string

const (
	SenderMarkHealthyResponseHealthUnverified SenderMarkHealthyResponseHealth = "UNVERIFIED"
	SenderMarkHealthyResponseHealthHealthy    SenderMarkHealthyResponseHealth = "HEALTHY"
	SenderMarkHealthyResponseHealthUnhealthy  SenderMarkHealthyResponseHealth = "UNHEALTHY"
)

// Always `sender`.
type SenderMarkHealthyResponseObjectType string

const (
	SenderMarkHealthyResponseObjectTypeSender SenderMarkHealthyResponseObjectType = "sender"
)

// How this sender paces each step type it sends, keyed by sequence step type:
// `EMAIL` for an email sender; `LINKEDIN_INVITE` and `LINKEDIN_MESSAGE` for a
// LinkedIn sender.
type SenderMarkHealthyResponsePacing struct {
	// Email steps. Present only on an email sender.
	Email SenderMarkHealthyResponsePacingEmail `json:"EMAIL"`
	// LinkedIn connection requests. Present only on a LinkedIn sender.
	LinkedinInvite SenderMarkHealthyResponsePacingLinkedinInvite `json:"LINKEDIN_INVITE"`
	// LinkedIn messages. Present only on a LinkedIn sender.
	LinkedinMessage SenderMarkHealthyResponsePacingLinkedinMessage `json:"LINKEDIN_MESSAGE"`
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
func (r SenderMarkHealthyResponsePacing) RawJSON() string { return r.JSON.raw }
func (r *SenderMarkHealthyResponsePacing) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Email steps. Present only on an email sender.
type SenderMarkHealthyResponsePacingEmail struct {
	// Most sends of this step type in any rolling 24 hours.
	DailyLimit int64 `json:"dailyLimit" api:"required"`
	// Upper bound of the minimum spacing, in seconds. Before each send, a minimum
	// spacing is drawn at random between `minGapSeconds` and this value, and the send
	// waits until that long after the previous one. Sends can be further apart, for
	// example outside the send window or once `dailyLimit` is reached. Null means the
	// minimum spacing is always `minGapSeconds`.
	MaxGapSeconds int64 `json:"maxGapSeconds" api:"required"`
	// Lower bound of the minimum spacing between two of its sends, in seconds.
	MinGapSeconds int64 `json:"minGapSeconds" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DailyLimit    respjson.Field
		MaxGapSeconds respjson.Field
		MinGapSeconds respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SenderMarkHealthyResponsePacingEmail) RawJSON() string { return r.JSON.raw }
func (r *SenderMarkHealthyResponsePacingEmail) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// LinkedIn connection requests. Present only on a LinkedIn sender.
type SenderMarkHealthyResponsePacingLinkedinInvite struct {
	// Most sends of this step type in any rolling 24 hours.
	DailyLimit int64 `json:"dailyLimit" api:"required"`
	// Upper bound of the minimum spacing, in seconds. Before each send, a minimum
	// spacing is drawn at random between `minGapSeconds` and this value, and the send
	// waits until that long after the previous one. Sends can be further apart, for
	// example outside the send window or once `dailyLimit` is reached. Null means the
	// minimum spacing is always `minGapSeconds`.
	MaxGapSeconds int64 `json:"maxGapSeconds" api:"required"`
	// Lower bound of the minimum spacing between two of its sends, in seconds.
	MinGapSeconds int64 `json:"minGapSeconds" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DailyLimit    respjson.Field
		MaxGapSeconds respjson.Field
		MinGapSeconds respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SenderMarkHealthyResponsePacingLinkedinInvite) RawJSON() string { return r.JSON.raw }
func (r *SenderMarkHealthyResponsePacingLinkedinInvite) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// LinkedIn messages. Present only on a LinkedIn sender.
type SenderMarkHealthyResponsePacingLinkedinMessage struct {
	// Most sends of this step type in any rolling 24 hours.
	DailyLimit int64 `json:"dailyLimit" api:"required"`
	// Upper bound of the minimum spacing, in seconds. Before each send, a minimum
	// spacing is drawn at random between `minGapSeconds` and this value, and the send
	// waits until that long after the previous one. Sends can be further apart, for
	// example outside the send window or once `dailyLimit` is reached. Null means the
	// minimum spacing is always `minGapSeconds`.
	MaxGapSeconds int64 `json:"maxGapSeconds" api:"required"`
	// Lower bound of the minimum spacing between two of its sends, in seconds.
	MinGapSeconds int64 `json:"minGapSeconds" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DailyLimit    respjson.Field
		MaxGapSeconds respjson.Field
		MinGapSeconds respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SenderMarkHealthyResponsePacingLinkedinMessage) RawJSON() string { return r.JSON.raw }
func (r *SenderMarkHealthyResponsePacingLinkedinMessage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// How an email sender sends: `SMTP` through its own mail server, `OAUTH` through
// its connected Google or Microsoft account. Null for LinkedIn senders.
type SenderMarkHealthyResponseTransport string

const (
	SenderMarkHealthyResponseTransportSmtp  SenderMarkHealthyResponseTransport = "SMTP"
	SenderMarkHealthyResponseTransportOAuth SenderMarkHealthyResponseTransport = "OAUTH"
)

type SenderRetrieveResponse struct {
	// Unique identifier of the sender.
	ID string `json:"id" api:"required"`
	// When the sender was archived, or null if it is active. Archived senders never
	// send.
	ArchivedAt string `json:"archivedAt" api:"required"`
	// Whether this sender sends email or LinkedIn messages.
	//
	// Any of "EMAIL", "LINKEDIN".
	Channel SenderRetrieveResponseChannel `json:"channel" api:"required"`
	// When the sender was created.
	CreatedAt string `json:"createdAt" api:"required"`
	// The address email from this sender is sent from. Null for LinkedIn senders, or
	// when the connected account has no address.
	EmailAddress string `json:"emailAddress" api:"required"`
	// `HEALTHY` can send. `UNHEALTHY` cannot send until `healthReason` is resolved.
	// `UNVERIFIED` means no health verdict yet: an email sender in this state cannot
	// send until it passes a health check, while a LinkedIn sender is never
	// health-checked, so a working LinkedIn sender is usually `UNVERIFIED` and can
	// send.
	//
	// Any of "UNVERIFIED", "HEALTHY", "UNHEALTHY".
	Health SenderRetrieveResponseHealth `json:"health" api:"required"`
	// Why the sender is `UNHEALTHY` or `UNVERIFIED`, for example
	// `ACCOUNT_DISCONNECTED`, `BOUNCE_THRESHOLD` or `INVITE_LIMIT`; null when healthy.
	HealthReason string `json:"healthReason" api:"required"`
	// Always `sender`.
	//
	// Any of "sender".
	ObjectType SenderRetrieveResponseObjectType `json:"objectType" api:"required"`
	// Member ID of the person who owns this sender, or null when no workspace
	// membership can be resolved.
	Owner string `json:"owner" api:"required"`
	// How this sender paces each step type it sends, keyed by sequence step type:
	// `EMAIL` for an email sender; `LINKEDIN_INVITE` and `LINKEDIN_MESSAGE` for a
	// LinkedIn sender.
	Pacing SenderRetrieveResponsePacing `json:"pacing" api:"required"`
	// SMTP senders only: the connected mailbox replies are sent to and read from. Null
	// for other senders.
	ReplyToAddress string `json:"replyToAddress" api:"required"`
	// How an email sender sends: `SMTP` through its own mail server, `OAUTH` through
	// its connected Google or Microsoft account. Null for LinkedIn senders.
	//
	// Any of "SMTP", "OAUTH".
	Transport SenderRetrieveResponseTransport `json:"transport" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID             respjson.Field
		ArchivedAt     respjson.Field
		Channel        respjson.Field
		CreatedAt      respjson.Field
		EmailAddress   respjson.Field
		Health         respjson.Field
		HealthReason   respjson.Field
		ObjectType     respjson.Field
		Owner          respjson.Field
		Pacing         respjson.Field
		ReplyToAddress respjson.Field
		Transport      respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SenderRetrieveResponse) RawJSON() string { return r.JSON.raw }
func (r *SenderRetrieveResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Whether this sender sends email or LinkedIn messages.
type SenderRetrieveResponseChannel string

const (
	SenderRetrieveResponseChannelEmail    SenderRetrieveResponseChannel = "EMAIL"
	SenderRetrieveResponseChannelLinkedin SenderRetrieveResponseChannel = "LINKEDIN"
)

// `HEALTHY` can send. `UNHEALTHY` cannot send until `healthReason` is resolved.
// `UNVERIFIED` means no health verdict yet: an email sender in this state cannot
// send until it passes a health check, while a LinkedIn sender is never
// health-checked, so a working LinkedIn sender is usually `UNVERIFIED` and can
// send.
type SenderRetrieveResponseHealth string

const (
	SenderRetrieveResponseHealthUnverified SenderRetrieveResponseHealth = "UNVERIFIED"
	SenderRetrieveResponseHealthHealthy    SenderRetrieveResponseHealth = "HEALTHY"
	SenderRetrieveResponseHealthUnhealthy  SenderRetrieveResponseHealth = "UNHEALTHY"
)

// Always `sender`.
type SenderRetrieveResponseObjectType string

const (
	SenderRetrieveResponseObjectTypeSender SenderRetrieveResponseObjectType = "sender"
)

// How this sender paces each step type it sends, keyed by sequence step type:
// `EMAIL` for an email sender; `LINKEDIN_INVITE` and `LINKEDIN_MESSAGE` for a
// LinkedIn sender.
type SenderRetrieveResponsePacing struct {
	// Email steps. Present only on an email sender.
	Email SenderRetrieveResponsePacingEmail `json:"EMAIL"`
	// LinkedIn connection requests. Present only on a LinkedIn sender.
	LinkedinInvite SenderRetrieveResponsePacingLinkedinInvite `json:"LINKEDIN_INVITE"`
	// LinkedIn messages. Present only on a LinkedIn sender.
	LinkedinMessage SenderRetrieveResponsePacingLinkedinMessage `json:"LINKEDIN_MESSAGE"`
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
func (r SenderRetrieveResponsePacing) RawJSON() string { return r.JSON.raw }
func (r *SenderRetrieveResponsePacing) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Email steps. Present only on an email sender.
type SenderRetrieveResponsePacingEmail struct {
	// Most sends of this step type in any rolling 24 hours.
	DailyLimit int64 `json:"dailyLimit" api:"required"`
	// Upper bound of the minimum spacing, in seconds. Before each send, a minimum
	// spacing is drawn at random between `minGapSeconds` and this value, and the send
	// waits until that long after the previous one. Sends can be further apart, for
	// example outside the send window or once `dailyLimit` is reached. Null means the
	// minimum spacing is always `minGapSeconds`.
	MaxGapSeconds int64 `json:"maxGapSeconds" api:"required"`
	// Lower bound of the minimum spacing between two of its sends, in seconds.
	MinGapSeconds int64 `json:"minGapSeconds" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DailyLimit    respjson.Field
		MaxGapSeconds respjson.Field
		MinGapSeconds respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SenderRetrieveResponsePacingEmail) RawJSON() string { return r.JSON.raw }
func (r *SenderRetrieveResponsePacingEmail) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// LinkedIn connection requests. Present only on a LinkedIn sender.
type SenderRetrieveResponsePacingLinkedinInvite struct {
	// Most sends of this step type in any rolling 24 hours.
	DailyLimit int64 `json:"dailyLimit" api:"required"`
	// Upper bound of the minimum spacing, in seconds. Before each send, a minimum
	// spacing is drawn at random between `minGapSeconds` and this value, and the send
	// waits until that long after the previous one. Sends can be further apart, for
	// example outside the send window or once `dailyLimit` is reached. Null means the
	// minimum spacing is always `minGapSeconds`.
	MaxGapSeconds int64 `json:"maxGapSeconds" api:"required"`
	// Lower bound of the minimum spacing between two of its sends, in seconds.
	MinGapSeconds int64 `json:"minGapSeconds" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DailyLimit    respjson.Field
		MaxGapSeconds respjson.Field
		MinGapSeconds respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SenderRetrieveResponsePacingLinkedinInvite) RawJSON() string { return r.JSON.raw }
func (r *SenderRetrieveResponsePacingLinkedinInvite) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// LinkedIn messages. Present only on a LinkedIn sender.
type SenderRetrieveResponsePacingLinkedinMessage struct {
	// Most sends of this step type in any rolling 24 hours.
	DailyLimit int64 `json:"dailyLimit" api:"required"`
	// Upper bound of the minimum spacing, in seconds. Before each send, a minimum
	// spacing is drawn at random between `minGapSeconds` and this value, and the send
	// waits until that long after the previous one. Sends can be further apart, for
	// example outside the send window or once `dailyLimit` is reached. Null means the
	// minimum spacing is always `minGapSeconds`.
	MaxGapSeconds int64 `json:"maxGapSeconds" api:"required"`
	// Lower bound of the minimum spacing between two of its sends, in seconds.
	MinGapSeconds int64 `json:"minGapSeconds" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DailyLimit    respjson.Field
		MaxGapSeconds respjson.Field
		MinGapSeconds respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SenderRetrieveResponsePacingLinkedinMessage) RawJSON() string { return r.JSON.raw }
func (r *SenderRetrieveResponsePacingLinkedinMessage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// How an email sender sends: `SMTP` through its own mail server, `OAUTH` through
// its connected Google or Microsoft account. Null for LinkedIn senders.
type SenderRetrieveResponseTransport string

const (
	SenderRetrieveResponseTransportSmtp  SenderRetrieveResponseTransport = "SMTP"
	SenderRetrieveResponseTransportOAuth SenderRetrieveResponseTransport = "OAUTH"
)

type SenderUpdateResponse struct {
	// Unique identifier of the sender.
	ID string `json:"id" api:"required"`
	// When the sender was archived, or null if it is active. Archived senders never
	// send.
	ArchivedAt string `json:"archivedAt" api:"required"`
	// Whether this sender sends email or LinkedIn messages.
	//
	// Any of "EMAIL", "LINKEDIN".
	Channel SenderUpdateResponseChannel `json:"channel" api:"required"`
	// When the sender was created.
	CreatedAt string `json:"createdAt" api:"required"`
	// The address email from this sender is sent from. Null for LinkedIn senders, or
	// when the connected account has no address.
	EmailAddress string `json:"emailAddress" api:"required"`
	// `HEALTHY` can send. `UNHEALTHY` cannot send until `healthReason` is resolved.
	// `UNVERIFIED` means no health verdict yet: an email sender in this state cannot
	// send until it passes a health check, while a LinkedIn sender is never
	// health-checked, so a working LinkedIn sender is usually `UNVERIFIED` and can
	// send.
	//
	// Any of "UNVERIFIED", "HEALTHY", "UNHEALTHY".
	Health SenderUpdateResponseHealth `json:"health" api:"required"`
	// Why the sender is `UNHEALTHY` or `UNVERIFIED`, for example
	// `ACCOUNT_DISCONNECTED`, `BOUNCE_THRESHOLD` or `INVITE_LIMIT`; null when healthy.
	HealthReason string `json:"healthReason" api:"required"`
	// Always `sender`.
	//
	// Any of "sender".
	ObjectType SenderUpdateResponseObjectType `json:"objectType" api:"required"`
	// Member ID of the person who owns this sender, or null when no workspace
	// membership can be resolved.
	Owner string `json:"owner" api:"required"`
	// How this sender paces each step type it sends, keyed by sequence step type:
	// `EMAIL` for an email sender; `LINKEDIN_INVITE` and `LINKEDIN_MESSAGE` for a
	// LinkedIn sender.
	Pacing SenderUpdateResponsePacing `json:"pacing" api:"required"`
	// SMTP senders only: the connected mailbox replies are sent to and read from. Null
	// for other senders.
	ReplyToAddress string `json:"replyToAddress" api:"required"`
	// How an email sender sends: `SMTP` through its own mail server, `OAUTH` through
	// its connected Google or Microsoft account. Null for LinkedIn senders.
	//
	// Any of "SMTP", "OAUTH".
	Transport SenderUpdateResponseTransport `json:"transport" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID             respjson.Field
		ArchivedAt     respjson.Field
		Channel        respjson.Field
		CreatedAt      respjson.Field
		EmailAddress   respjson.Field
		Health         respjson.Field
		HealthReason   respjson.Field
		ObjectType     respjson.Field
		Owner          respjson.Field
		Pacing         respjson.Field
		ReplyToAddress respjson.Field
		Transport      respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SenderUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *SenderUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Whether this sender sends email or LinkedIn messages.
type SenderUpdateResponseChannel string

const (
	SenderUpdateResponseChannelEmail    SenderUpdateResponseChannel = "EMAIL"
	SenderUpdateResponseChannelLinkedin SenderUpdateResponseChannel = "LINKEDIN"
)

// `HEALTHY` can send. `UNHEALTHY` cannot send until `healthReason` is resolved.
// `UNVERIFIED` means no health verdict yet: an email sender in this state cannot
// send until it passes a health check, while a LinkedIn sender is never
// health-checked, so a working LinkedIn sender is usually `UNVERIFIED` and can
// send.
type SenderUpdateResponseHealth string

const (
	SenderUpdateResponseHealthUnverified SenderUpdateResponseHealth = "UNVERIFIED"
	SenderUpdateResponseHealthHealthy    SenderUpdateResponseHealth = "HEALTHY"
	SenderUpdateResponseHealthUnhealthy  SenderUpdateResponseHealth = "UNHEALTHY"
)

// Always `sender`.
type SenderUpdateResponseObjectType string

const (
	SenderUpdateResponseObjectTypeSender SenderUpdateResponseObjectType = "sender"
)

// How this sender paces each step type it sends, keyed by sequence step type:
// `EMAIL` for an email sender; `LINKEDIN_INVITE` and `LINKEDIN_MESSAGE` for a
// LinkedIn sender.
type SenderUpdateResponsePacing struct {
	// Email steps. Present only on an email sender.
	Email SenderUpdateResponsePacingEmail `json:"EMAIL"`
	// LinkedIn connection requests. Present only on a LinkedIn sender.
	LinkedinInvite SenderUpdateResponsePacingLinkedinInvite `json:"LINKEDIN_INVITE"`
	// LinkedIn messages. Present only on a LinkedIn sender.
	LinkedinMessage SenderUpdateResponsePacingLinkedinMessage `json:"LINKEDIN_MESSAGE"`
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
func (r SenderUpdateResponsePacing) RawJSON() string { return r.JSON.raw }
func (r *SenderUpdateResponsePacing) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Email steps. Present only on an email sender.
type SenderUpdateResponsePacingEmail struct {
	// Most sends of this step type in any rolling 24 hours.
	DailyLimit int64 `json:"dailyLimit" api:"required"`
	// Upper bound of the minimum spacing, in seconds. Before each send, a minimum
	// spacing is drawn at random between `minGapSeconds` and this value, and the send
	// waits until that long after the previous one. Sends can be further apart, for
	// example outside the send window or once `dailyLimit` is reached. Null means the
	// minimum spacing is always `minGapSeconds`.
	MaxGapSeconds int64 `json:"maxGapSeconds" api:"required"`
	// Lower bound of the minimum spacing between two of its sends, in seconds.
	MinGapSeconds int64 `json:"minGapSeconds" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DailyLimit    respjson.Field
		MaxGapSeconds respjson.Field
		MinGapSeconds respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SenderUpdateResponsePacingEmail) RawJSON() string { return r.JSON.raw }
func (r *SenderUpdateResponsePacingEmail) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// LinkedIn connection requests. Present only on a LinkedIn sender.
type SenderUpdateResponsePacingLinkedinInvite struct {
	// Most sends of this step type in any rolling 24 hours.
	DailyLimit int64 `json:"dailyLimit" api:"required"`
	// Upper bound of the minimum spacing, in seconds. Before each send, a minimum
	// spacing is drawn at random between `minGapSeconds` and this value, and the send
	// waits until that long after the previous one. Sends can be further apart, for
	// example outside the send window or once `dailyLimit` is reached. Null means the
	// minimum spacing is always `minGapSeconds`.
	MaxGapSeconds int64 `json:"maxGapSeconds" api:"required"`
	// Lower bound of the minimum spacing between two of its sends, in seconds.
	MinGapSeconds int64 `json:"minGapSeconds" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DailyLimit    respjson.Field
		MaxGapSeconds respjson.Field
		MinGapSeconds respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SenderUpdateResponsePacingLinkedinInvite) RawJSON() string { return r.JSON.raw }
func (r *SenderUpdateResponsePacingLinkedinInvite) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// LinkedIn messages. Present only on a LinkedIn sender.
type SenderUpdateResponsePacingLinkedinMessage struct {
	// Most sends of this step type in any rolling 24 hours.
	DailyLimit int64 `json:"dailyLimit" api:"required"`
	// Upper bound of the minimum spacing, in seconds. Before each send, a minimum
	// spacing is drawn at random between `minGapSeconds` and this value, and the send
	// waits until that long after the previous one. Sends can be further apart, for
	// example outside the send window or once `dailyLimit` is reached. Null means the
	// minimum spacing is always `minGapSeconds`.
	MaxGapSeconds int64 `json:"maxGapSeconds" api:"required"`
	// Lower bound of the minimum spacing between two of its sends, in seconds.
	MinGapSeconds int64 `json:"minGapSeconds" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DailyLimit    respjson.Field
		MaxGapSeconds respjson.Field
		MinGapSeconds respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SenderUpdateResponsePacingLinkedinMessage) RawJSON() string { return r.JSON.raw }
func (r *SenderUpdateResponsePacingLinkedinMessage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// How an email sender sends: `SMTP` through its own mail server, `OAUTH` through
// its connected Google or Microsoft account. Null for LinkedIn senders.
type SenderUpdateResponseTransport string

const (
	SenderUpdateResponseTransportSmtp  SenderUpdateResponseTransport = "SMTP"
	SenderUpdateResponseTransportOAuth SenderUpdateResponseTransport = "OAUTH"
)

type SenderUpdateParams struct {
	// The sender's new pacing, replacing all of its current values. Send exactly the
	// step types the sender's channel sends: `EMAIL` for an email sender;
	// `LINKEDIN_INVITE` and `LINKEDIN_MESSAGE` for a LinkedIn sender.
	Pacing SenderUpdateParamsPacing `json:"pacing,omitzero" api:"required"`
	paramObj
}

func (r SenderUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow SenderUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SenderUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The sender's new pacing, replacing all of its current values. Send exactly the
// step types the sender's channel sends: `EMAIL` for an email sender;
// `LINKEDIN_INVITE` and `LINKEDIN_MESSAGE` for a LinkedIn sender.
type SenderUpdateParamsPacing struct {
	// Email steps. Present only on an email sender.
	Email SenderUpdateParamsPacingEmail `json:"EMAIL,omitzero"`
	// LinkedIn connection requests. Present only on a LinkedIn sender.
	LinkedinInvite SenderUpdateParamsPacingLinkedinInvite `json:"LINKEDIN_INVITE,omitzero"`
	// LinkedIn messages. Present only on a LinkedIn sender.
	LinkedinMessage SenderUpdateParamsPacingLinkedinMessage `json:"LINKEDIN_MESSAGE,omitzero"`
	paramObj
}

func (r SenderUpdateParamsPacing) MarshalJSON() (data []byte, err error) {
	type shadow SenderUpdateParamsPacing
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SenderUpdateParamsPacing) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Email steps. Present only on an email sender.
//
// The properties DailyLimit, MaxGapSeconds, MinGapSeconds are required.
type SenderUpdateParamsPacingEmail struct {
	// Upper bound of the minimum spacing, in seconds. Before each send, a minimum
	// spacing is drawn at random between `minGapSeconds` and this value, and the send
	// waits until that long after the previous one. Sends can be further apart, for
	// example outside the send window or once `dailyLimit` is reached. Null means the
	// minimum spacing is always `minGapSeconds`.
	MaxGapSeconds param.Opt[int64] `json:"maxGapSeconds,omitzero" api:"required"`
	// Most sends of this step type in any rolling 24 hours.
	DailyLimit int64 `json:"dailyLimit" api:"required"`
	// Lower bound of the minimum spacing between two of its sends, in seconds.
	MinGapSeconds int64 `json:"minGapSeconds" api:"required"`
	paramObj
}

func (r SenderUpdateParamsPacingEmail) MarshalJSON() (data []byte, err error) {
	type shadow SenderUpdateParamsPacingEmail
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SenderUpdateParamsPacingEmail) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// LinkedIn connection requests. Present only on a LinkedIn sender.
//
// The properties DailyLimit, MaxGapSeconds, MinGapSeconds are required.
type SenderUpdateParamsPacingLinkedinInvite struct {
	// Upper bound of the minimum spacing, in seconds. Before each send, a minimum
	// spacing is drawn at random between `minGapSeconds` and this value, and the send
	// waits until that long after the previous one. Sends can be further apart, for
	// example outside the send window or once `dailyLimit` is reached. Null means the
	// minimum spacing is always `minGapSeconds`.
	MaxGapSeconds param.Opt[int64] `json:"maxGapSeconds,omitzero" api:"required"`
	// Most sends of this step type in any rolling 24 hours.
	DailyLimit int64 `json:"dailyLimit" api:"required"`
	// Lower bound of the minimum spacing between two of its sends, in seconds.
	MinGapSeconds int64 `json:"minGapSeconds" api:"required"`
	paramObj
}

func (r SenderUpdateParamsPacingLinkedinInvite) MarshalJSON() (data []byte, err error) {
	type shadow SenderUpdateParamsPacingLinkedinInvite
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SenderUpdateParamsPacingLinkedinInvite) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// LinkedIn messages. Present only on a LinkedIn sender.
//
// The properties DailyLimit, MaxGapSeconds, MinGapSeconds are required.
type SenderUpdateParamsPacingLinkedinMessage struct {
	// Upper bound of the minimum spacing, in seconds. Before each send, a minimum
	// spacing is drawn at random between `minGapSeconds` and this value, and the send
	// waits until that long after the previous one. Sends can be further apart, for
	// example outside the send window or once `dailyLimit` is reached. Null means the
	// minimum spacing is always `minGapSeconds`.
	MaxGapSeconds param.Opt[int64] `json:"maxGapSeconds,omitzero" api:"required"`
	// Most sends of this step type in any rolling 24 hours.
	DailyLimit int64 `json:"dailyLimit" api:"required"`
	// Lower bound of the minimum spacing between two of its sends, in seconds.
	MinGapSeconds int64 `json:"minGapSeconds" api:"required"`
	paramObj
}

func (r SenderUpdateParamsPacingLinkedinMessage) MarshalJSON() (data []byte, err error) {
	type shadow SenderUpdateParamsPacingLinkedinMessage
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SenderUpdateParamsPacingLinkedinMessage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SenderListParams struct {
	// Maximum number of records to return. Defaults to 25, maximum 25.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Number of records to skip for pagination. Defaults to 0.
	Offset param.Opt[int64] `query:"offset,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [SenderListParams]'s query parameters as `url.Values`.
func (r SenderListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type SenderMarkHealthyParams struct {
	Body SenderMarkHealthyParamsBody
	paramObj
}

func (r SenderMarkHealthyParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.Body)
}
func (r *SenderMarkHealthyParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SenderMarkHealthyParamsBody struct {
	paramObj
}

func (r SenderMarkHealthyParamsBody) MarshalJSON() (data []byte, err error) {
	type shadow SenderMarkHealthyParamsBody
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SenderMarkHealthyParamsBody) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
