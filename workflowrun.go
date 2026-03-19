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
	"github.com/Lightfld/lightfield-go/packages/respjson"
)

// Workflow runs represent executions of automated workflows.
//
// WorkflowRunService contains methods and other services that help with
// interacting with the Lightfield API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWorkflowRunService] method instead.
type WorkflowRunService struct {
	Options []option.RequestOption
}

// NewWorkflowRunService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewWorkflowRunService(opts ...option.RequestOption) (r WorkflowRunService) {
	r = WorkflowRunService{}
	r.Options = opts
	return
}

// Returns the current status of a workflow run.
func (r *WorkflowRunService) Status(ctx context.Context, runID string, opts ...option.RequestOption) (res *WorkflowRunStatusResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if runID == "" {
		err = errors.New("missing required runId parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/workflowRun/%s/status", url.PathEscape(runID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

type WorkflowRunStatusResponse struct {
	Status string `json:"status" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkflowRunStatusResponse) RawJSON() string { return r.JSON.raw }
func (r *WorkflowRunStatusResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
