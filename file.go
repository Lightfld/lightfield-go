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

// Files are used to upload documents via presigned URLs. After uploading and
// completing a file, link it to resources through their own APIs (e.g. attach a
// transcript to a meeting). See
// <u>[File uploads](/using-the-api/file-uploads/)</u> for the full upload flow and
// supported purposes. For meeting transcript attachments, see
// <u>[Uploading meeting transcripts](/using-the-api/uploading-meeting-transcripts/)</u>.
//
// FileService contains methods and other services that help with interacting with
// the Lightfield API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewFileService] method instead.
type FileService struct {
	Options []option.RequestOption
}

// NewFileService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewFileService(opts ...option.RequestOption) (r FileService) {
	r = FileService{}
	r.Options = opts
	return
}

// Creates a new file upload session and returns an upload URL.
//
// After uploading the file bytes to `uploadUrl`, call
// `POST /v1/files/{id}/complete` to finalize the upload. Optionally pass `purpose`
// to validate MIME type and size constraints at creation time. See
// <u>[File uploads](/using-the-api/file-uploads/)</u> for the full upload flow,
// supported purposes, and size limits. If you are uploading a meeting transcript,
// see
// <u>[Uploading meeting transcripts](/using-the-api/uploading-meeting-transcripts/)</u>
// for the follow-up meeting attachment flow.
//
// **[Required scope](/using-the-api/scopes/):** `files:create`
//
// **[Rate limit category](/using-the-api/rate-limits/):** Write
func (r *FileService) New(ctx context.Context, body FileNewParams, opts ...option.RequestOption) (res *FileCreateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/files"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Retrieves a single file by its ID.
//
// **[Required scope](/using-the-api/scopes/):** `files:read`
//
// **[Rate limit category](/using-the-api/rate-limits/):** Read
func (r *FileService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *FileRetrieveResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/files/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Returns a paginated list of files in your workspace. Use `offset` and `limit` to
// paginate through results. See
// <u>[List endpoints](/using-the-api/list-endpoints/)</u> for more information
// about pagination.
//
// **[Required scope](/using-the-api/scopes/):** `files:read`
//
// **[Rate limit category](/using-the-api/rate-limits/):** Search
func (r *FileService) List(ctx context.Context, query FileListParams, opts ...option.RequestOption) (res *FileListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/files"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Cancels a pending upload by transitioning the file to `CANCELLED`. Only files in
// `PENDING` status can be cancelled. **[Required scope](/using-the-api/scopes/):**
// `files:create`
//
// **[Rate limit category](/using-the-api/rate-limits/):** Write
func (r *FileService) Cancel(ctx context.Context, id string, body FileCancelParams, opts ...option.RequestOption) (res *FileCancelResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/files/%s/cancel", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Finalizes an upload after the file bytes have been uploaded.
//
// If an optional `md5` hex digest is provided, the server validates the checksum
// before marking the file as completed.
//
// **[Required scope](/using-the-api/scopes/):** `files:create`
//
// **[Rate limit category](/using-the-api/rate-limits/):** Write
func (r *FileService) Complete(ctx context.Context, id string, body FileCompleteParams, opts ...option.RequestOption) (res *FileCompleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/files/%s/complete", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Returns a temporary download URL for the file. Only available for files in
// `COMPLETED` status.
//
// **[Required scope](/using-the-api/scopes/):** `files:read`
//
// **[Rate limit category](/using-the-api/rate-limits/):** Read
func (r *FileService) URL(ctx context.Context, id string, opts ...option.RequestOption) (res *FileURLResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/files/%s/url", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

type FileCancelResponse struct {
	// Unique identifier for the file.
	ID string `json:"id" api:"required"`
	// When the file upload was completed.
	CompletedAt string `json:"completedAt" api:"required"`
	// When the file upload session was created.
	CreatedAt string `json:"createdAt" api:"required"`
	// When the upload session expires. Null once completed, cancelled, or expired.
	ExpiresAt string `json:"expiresAt" api:"required"`
	// Original filename.
	Filename string `json:"filename" api:"required"`
	// MIME type of the file.
	MimeType string `json:"mimeType" api:"required"`
	// File size in bytes.
	SizeBytes int64 `json:"sizeBytes" api:"required"`
	// Current upload status of the file.
	//
	// Any of "PENDING", "COMPLETED", "CANCELLED", "EXPIRED".
	Status FileCancelResponseStatus `json:"status" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CompletedAt respjson.Field
		CreatedAt   respjson.Field
		ExpiresAt   respjson.Field
		Filename    respjson.Field
		MimeType    respjson.Field
		SizeBytes   respjson.Field
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FileCancelResponse) RawJSON() string { return r.JSON.raw }
func (r *FileCancelResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Current upload status of the file.
type FileCancelResponseStatus string

const (
	FileCancelResponseStatusPending   FileCancelResponseStatus = "PENDING"
	FileCancelResponseStatusCompleted FileCancelResponseStatus = "COMPLETED"
	FileCancelResponseStatusCancelled FileCancelResponseStatus = "CANCELLED"
	FileCancelResponseStatusExpired   FileCancelResponseStatus = "EXPIRED"
)

type FileCompleteResponse struct {
	// Unique identifier for the file.
	ID string `json:"id" api:"required"`
	// When the file upload was completed.
	CompletedAt string `json:"completedAt" api:"required"`
	// When the file upload session was created.
	CreatedAt string `json:"createdAt" api:"required"`
	// When the upload session expires. Null once completed, cancelled, or expired.
	ExpiresAt string `json:"expiresAt" api:"required"`
	// Original filename.
	Filename string `json:"filename" api:"required"`
	// MIME type of the file.
	MimeType string `json:"mimeType" api:"required"`
	// File size in bytes.
	SizeBytes int64 `json:"sizeBytes" api:"required"`
	// Current upload status of the file.
	//
	// Any of "PENDING", "COMPLETED", "CANCELLED", "EXPIRED".
	Status FileCompleteResponseStatus `json:"status" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CompletedAt respjson.Field
		CreatedAt   respjson.Field
		ExpiresAt   respjson.Field
		Filename    respjson.Field
		MimeType    respjson.Field
		SizeBytes   respjson.Field
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FileCompleteResponse) RawJSON() string { return r.JSON.raw }
func (r *FileCompleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Current upload status of the file.
type FileCompleteResponseStatus string

const (
	FileCompleteResponseStatusPending   FileCompleteResponseStatus = "PENDING"
	FileCompleteResponseStatusCompleted FileCompleteResponseStatus = "COMPLETED"
	FileCompleteResponseStatusCancelled FileCompleteResponseStatus = "CANCELLED"
	FileCompleteResponseStatusExpired   FileCompleteResponseStatus = "EXPIRED"
)

type FileCreateResponse struct {
	// Unique identifier for the file.
	ID string `json:"id" api:"required"`
	// When the file upload was completed.
	CompletedAt string `json:"completedAt" api:"required"`
	// When the file upload session was created.
	CreatedAt string `json:"createdAt" api:"required"`
	// When the upload session expires. Null once completed, cancelled, or expired.
	ExpiresAt string `json:"expiresAt" api:"required"`
	// Original filename.
	Filename string `json:"filename" api:"required"`
	// MIME type of the file.
	MimeType string `json:"mimeType" api:"required"`
	// File size in bytes.
	SizeBytes int64 `json:"sizeBytes" api:"required"`
	// Current upload status of the file.
	//
	// Any of "PENDING", "COMPLETED", "CANCELLED", "EXPIRED".
	Status FileCreateResponseStatus `json:"status" api:"required"`
	// Headers to include in the upload request.
	UploadHeaders map[string]string `json:"uploadHeaders" api:"required"`
	// Upload URL. Upload the file bytes directly to this URL.
	UploadURL string `json:"uploadUrl" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID            respjson.Field
		CompletedAt   respjson.Field
		CreatedAt     respjson.Field
		ExpiresAt     respjson.Field
		Filename      respjson.Field
		MimeType      respjson.Field
		SizeBytes     respjson.Field
		Status        respjson.Field
		UploadHeaders respjson.Field
		UploadURL     respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FileCreateResponse) RawJSON() string { return r.JSON.raw }
func (r *FileCreateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Current upload status of the file.
type FileCreateResponseStatus string

const (
	FileCreateResponseStatusPending   FileCreateResponseStatus = "PENDING"
	FileCreateResponseStatusCompleted FileCreateResponseStatus = "COMPLETED"
	FileCreateResponseStatusCancelled FileCreateResponseStatus = "CANCELLED"
	FileCreateResponseStatusExpired   FileCreateResponseStatus = "EXPIRED"
)

type FileListResponse struct {
	// Array of file objects for the current page.
	Data []FileListResponseData `json:"data" api:"required"`
	// The object type, always `"list"`.
	Object string `json:"object" api:"required"`
	// Total number of matching files.
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
func (r FileListResponse) RawJSON() string { return r.JSON.raw }
func (r *FileListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FileListResponseData struct {
	// Unique identifier for the file.
	ID string `json:"id" api:"required"`
	// When the file upload was completed.
	CompletedAt string `json:"completedAt" api:"required"`
	// When the file upload session was created.
	CreatedAt string `json:"createdAt" api:"required"`
	// When the upload session expires. Null once completed, cancelled, or expired.
	ExpiresAt string `json:"expiresAt" api:"required"`
	// Original filename.
	Filename string `json:"filename" api:"required"`
	// MIME type of the file.
	MimeType string `json:"mimeType" api:"required"`
	// File size in bytes.
	SizeBytes int64 `json:"sizeBytes" api:"required"`
	// Current upload status of the file.
	//
	// Any of "PENDING", "COMPLETED", "CANCELLED", "EXPIRED".
	Status string `json:"status" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CompletedAt respjson.Field
		CreatedAt   respjson.Field
		ExpiresAt   respjson.Field
		Filename    respjson.Field
		MimeType    respjson.Field
		SizeBytes   respjson.Field
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FileListResponseData) RawJSON() string { return r.JSON.raw }
func (r *FileListResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FileRetrieveResponse struct {
	// Unique identifier for the file.
	ID string `json:"id" api:"required"`
	// When the file upload was completed.
	CompletedAt string `json:"completedAt" api:"required"`
	// When the file upload session was created.
	CreatedAt string `json:"createdAt" api:"required"`
	// When the upload session expires. Null once completed, cancelled, or expired.
	ExpiresAt string `json:"expiresAt" api:"required"`
	// Original filename.
	Filename string `json:"filename" api:"required"`
	// MIME type of the file.
	MimeType string `json:"mimeType" api:"required"`
	// File size in bytes.
	SizeBytes int64 `json:"sizeBytes" api:"required"`
	// Current upload status of the file.
	//
	// Any of "PENDING", "COMPLETED", "CANCELLED", "EXPIRED".
	Status FileRetrieveResponseStatus `json:"status" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CompletedAt respjson.Field
		CreatedAt   respjson.Field
		ExpiresAt   respjson.Field
		Filename    respjson.Field
		MimeType    respjson.Field
		SizeBytes   respjson.Field
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FileRetrieveResponse) RawJSON() string { return r.JSON.raw }
func (r *FileRetrieveResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Current upload status of the file.
type FileRetrieveResponseStatus string

const (
	FileRetrieveResponseStatusPending   FileRetrieveResponseStatus = "PENDING"
	FileRetrieveResponseStatusCompleted FileRetrieveResponseStatus = "COMPLETED"
	FileRetrieveResponseStatusCancelled FileRetrieveResponseStatus = "CANCELLED"
	FileRetrieveResponseStatusExpired   FileRetrieveResponseStatus = "EXPIRED"
)

type FileURLResponse struct {
	// When the download URL expires.
	ExpiresAt string `json:"expiresAt" api:"required"`
	// Temporary download URL for the file.
	URL string `json:"url" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ExpiresAt   respjson.Field
		URL         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FileURLResponse) RawJSON() string { return r.JSON.raw }
func (r *FileURLResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FileNewParams struct {
	// Original filename.
	Filename string `json:"filename" api:"required"`
	// MIME type of the file. Must be allowed for the given purpose (if specified).
	MimeType string `json:"mimeType" api:"required"`
	// Expected file size in bytes. Maximum 512 MB.
	SizeBytes int64 `json:"sizeBytes" api:"required"`
	// Optional validation hint. When provided, the server enforces purpose-specific
	// MIME type and file size constraints. Use `meeting_transcript` for files that
	// will be attached to a meeting as its transcript. Use `knowledge_user` or
	// `knowledge_workspace` to add the file to the authenticated user's or workspace's
	// Knowledge, making it available to the AI assistant. Not persisted or returned in
	// responses.
	//
	// Any of "meeting_transcript", "knowledge_user", "knowledge_workspace".
	Purpose FileNewParamsPurpose `json:"purpose,omitzero"`
	paramObj
}

func (r FileNewParams) MarshalJSON() (data []byte, err error) {
	type shadow FileNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FileNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Optional validation hint. When provided, the server enforces purpose-specific
// MIME type and file size constraints. Use `meeting_transcript` for files that
// will be attached to a meeting as its transcript. Use `knowledge_user` or
// `knowledge_workspace` to add the file to the authenticated user's or workspace's
// Knowledge, making it available to the AI assistant. Not persisted or returned in
// responses.
type FileNewParamsPurpose string

const (
	FileNewParamsPurposeMeetingTranscript  FileNewParamsPurpose = "meeting_transcript"
	FileNewParamsPurposeKnowledgeUser      FileNewParamsPurpose = "knowledge_user"
	FileNewParamsPurposeKnowledgeWorkspace FileNewParamsPurpose = "knowledge_workspace"
)

type FileListParams struct {
	// Maximum number of records to return. Defaults to 25, maximum 25.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Number of records to skip for pagination. Defaults to 0.
	Offset param.Opt[int64] `query:"offset,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [FileListParams]'s query parameters as `url.Values`.
func (r FileListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type FileCancelParams struct {
	Body FileCancelParamsBody
	paramObj
}

func (r FileCancelParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.Body)
}
func (r *FileCancelParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FileCancelParamsBody struct {
	paramObj
}

func (r FileCancelParamsBody) MarshalJSON() (data []byte, err error) {
	type shadow FileCancelParamsBody
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FileCancelParamsBody) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FileCompleteParams struct {
	// Optional MD5 hex digest of the uploaded file for checksum verification.
	Md5 param.Opt[string] `json:"md5,omitzero"`
	paramObj
}

func (r FileCompleteParams) MarshalJSON() (data []byte, err error) {
	type shadow FileCompleteParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FileCompleteParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
