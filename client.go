// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package githubcomlightfldlightfieldgo

import (
	"context"
	"net/http"
	"os"
	"slices"

	"github.com/Lightfld/lightfield-go/internal/requestconfig"
	"github.com/Lightfld/lightfield-go/option"
)

// Client creates a struct with services and top level methods that help with
// interacting with the Lightfield API. You should not instantiate this client
// directly, and instead use the [NewClient] method instead.
type Client struct {
	Options []option.RequestOption
	// Accounts represent companies or organizations in Lightfield. Each account can
	// have contacts, opportunities, tasks, and notes associated with it.
	Account AccountService
	// Contacts represent individual people in Lightfield. Contacts can be associated
	// with one or more accounts.
	Contact ContactService
	// Lists are curated collections of accounts, contacts, or opportunities in
	// Lightfield. Each list contains entities of a single type.
	List ListService
	// Meetings represent synced or manually created interactions in Lightfield. Read
	// responses are privacy-aware and may be redacted based on the caller. For
	// transcript uploads and attachment flows, see
	// <u>[Uploading meeting transcripts](/using-the-api/uploading-meeting-transcripts/)</u>.
	Meeting MeetingService
	// Notes represent free-form text records in Lightfield. Each note can be
	// associated with zero or more accounts and opportunities.
	Note NoteService
	// Opportunities represent potential deals or sales in Lightfield. Each opportunity
	// belongs to an account and can have tasks and notes associated with it.
	Opportunity OpportunityService
	// Tasks represent action items in Lightfield. Each task belongs to an account, is
	// assigned to a member, and can optionally be associated with an opportunity.
	Task TaskService
	// Members represent users in your Lightfield workspace. Members can own accounts
	// and opportunities, and are referenced in relationships like `$owner` and
	// `$createdBy`.
	Member MemberService
	// Workflow runs represent executions of automated workflows.
	WorkflowRun WorkflowRunService
	// Files are used to upload documents via presigned URLs. After uploading and
	// completing a file, link it to resources through their own APIs (e.g. attach a
	// transcript to a meeting). See
	// <u>[File uploads](/using-the-api/file-uploads/)</u> for the full upload flow and
	// supported purposes. For meeting transcript attachments, see
	// <u>[Uploading meeting transcripts](/using-the-api/uploading-meeting-transcripts/)</u>.
	File FileService
}

// DefaultClientOptions read from the environment (LIGHTFIELD_BASE_URL). This
// should be used to initialize new clients.
func DefaultClientOptions() []option.RequestOption {
	defaults := []option.RequestOption{option.WithEnvironmentProduction()}
	if o, ok := os.LookupEnv("LIGHTFIELD_BASE_URL"); ok {
		defaults = append(defaults, option.WithBaseURL(o))
	}
	return defaults
}

// NewClient generates a new client with the default option read from the
// environment (LIGHTFIELD_BASE_URL). The option passed in as arguments are applied
// after these default arguments, and all option will be passed down to the
// services and requests that this client makes.
func NewClient(opts ...option.RequestOption) (r Client) {
	opts = append(DefaultClientOptions(), opts...)

	r = Client{Options: opts}

	r.Account = NewAccountService(opts...)
	r.Contact = NewContactService(opts...)
	r.List = NewListService(opts...)
	r.Meeting = NewMeetingService(opts...)
	r.Note = NewNoteService(opts...)
	r.Opportunity = NewOpportunityService(opts...)
	r.Task = NewTaskService(opts...)
	r.Member = NewMemberService(opts...)
	r.WorkflowRun = NewWorkflowRunService(opts...)
	r.File = NewFileService(opts...)

	return
}

// Execute makes a request with the given context, method, URL, request params,
// response, and request options. This is useful for hitting undocumented endpoints
// while retaining the base URL, auth, retries, and other options from the client.
//
// If a byte slice or an [io.Reader] is supplied to params, it will be used as-is
// for the request body.
//
// The params is by default serialized into the body using [encoding/json]. If your
// type implements a MarshalJSON function, it will be used instead to serialize the
// request. If a URLQuery method is implemented, the returned [url.Values] will be
// used as query strings to the url.
//
// If your params struct uses [param.Field], you must provide either [MarshalJSON],
// [URLQuery], and/or [MarshalForm] functions. It is undefined behavior to use a
// struct uses [param.Field] without specifying how it is serialized.
//
// Any "…Params" object defined in this library can be used as the request
// argument. Note that 'path' arguments will not be forwarded into the url.
//
// The response body will be deserialized into the res variable, depending on its
// type:
//
//   - A pointer to a [*http.Response] is populated by the raw response.
//   - A pointer to a byte array will be populated with the contents of the request
//     body.
//   - A pointer to any other type uses this library's default JSON decoding, which
//     respects UnmarshalJSON if it is defined on the type.
//   - A nil value will not read the response body.
//
// For even greater flexibility, see [option.WithResponseInto] and
// [option.WithResponseBodyInto].
func (r *Client) Execute(ctx context.Context, method string, path string, params any, res any, opts ...option.RequestOption) error {
	opts = slices.Concat(r.Options, opts)
	return requestconfig.ExecuteNewRequest(ctx, method, path, params, res, opts...)
}

// Get makes a GET request with the given URL, params, and optionally deserializes
// to a response. See [Execute] documentation on the params and response.
func (r *Client) Get(ctx context.Context, path string, params any, res any, opts ...option.RequestOption) error {
	return r.Execute(ctx, http.MethodGet, path, params, res, opts...)
}

// Post makes a POST request with the given URL, params, and optionally
// deserializes to a response. See [Execute] documentation on the params and
// response.
func (r *Client) Post(ctx context.Context, path string, params any, res any, opts ...option.RequestOption) error {
	return r.Execute(ctx, http.MethodPost, path, params, res, opts...)
}

// Put makes a PUT request with the given URL, params, and optionally deserializes
// to a response. See [Execute] documentation on the params and response.
func (r *Client) Put(ctx context.Context, path string, params any, res any, opts ...option.RequestOption) error {
	return r.Execute(ctx, http.MethodPut, path, params, res, opts...)
}

// Patch makes a PATCH request with the given URL, params, and optionally
// deserializes to a response. See [Execute] documentation on the params and
// response.
func (r *Client) Patch(ctx context.Context, path string, params any, res any, opts ...option.RequestOption) error {
	return r.Execute(ctx, http.MethodPatch, path, params, res, opts...)
}

// Delete makes a DELETE request with the given URL, params, and optionally
// deserializes to a response. See [Execute] documentation on the params and
// response.
func (r *Client) Delete(ctx context.Context, path string, params any, res any, opts ...option.RequestOption) error {
	return r.Execute(ctx, http.MethodDelete, path, params, res, opts...)
}
