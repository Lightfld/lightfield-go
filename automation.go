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
	"github.com/Lightfld/lightfield-go/shared"
)

// > [!WARNING] The Automations API is in beta and isn't available to every
// > workspace yet. Endpoints, parameters, and response schemas may change,
// > including breaking changes.
//
// Automations run their steps whenever one of their triggers fires. An automation
// created or updated through the API goes live immediately; pause it to stop new
// runs.
//
// AutomationService contains methods and other services that help with interacting
// with the Lightfield API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAutomationService] method instead.
type AutomationService struct {
	Options []option.RequestOption
}

// NewAutomationService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAutomationService(opts ...option.RequestOption) (r AutomationService) {
	r = AutomationService{}
	r.Options = opts
	return
}

// > [!WARNING] This endpoint is in beta and isn't available to every workspace
// > yet. Its parameters and response schemas may change, including breaking
// > changes.
//
// Creates an automation and makes it live at once: its triggers start runs as soon
// as the request returns. The permissions on its steps are granted by this
// request; nobody approves them in the app.
//
// Requires an API key whose user is a workspace owner or admin. Returns 422 when
// `triggers` is empty or the workspace cannot choose the requested `model`; omit
// `model` or send `auto`.
//
// **[Required scope](/using-the-api/scopes/):** `automations:create`
//
// **[Rate limit category](/using-the-api/rate-limits/):** Write
func (r *AutomationService) New(ctx context.Context, body AutomationNewParams, opts ...option.RequestOption) (res *AutomationCreateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/automations"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// > [!WARNING] This endpoint is in beta and isn't available to every workspace
// > yet. Its parameters and response schemas may change, including breaking
// > changes.
//
// Retrieves an automation’s live version, including relationship watches and their
// resolved directions. Inputs, skills and files set in the app are not part of the
// response.
//
// Requires an API key whose user is a workspace owner or admin.
//
// **[Required scope](/using-the-api/scopes/):** `automations:read`
//
// **[Rate limit category](/using-the-api/rate-limits/):** Read
func (r *AutomationService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *AutomationRetrieveResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/automations/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// > [!WARNING] This endpoint is in beta and isn't available to every workspace
// > yet. Its parameters and response schemas may change, including breaking
// > changes.
//
// Replaces the fields sent and leaves the rest as they are, then makes the result
// the live version. Each field is replaced whole: sending `steps` replaces every
// step, and sending `triggers` replaces every trigger, including relationship
// watches set in the app. A paused automation stays paused. A request that changes
// nothing returns the automation as it is.
//
// Requires an API key whose user is a workspace owner or admin. Returns 422 when
// `triggers` is empty or the workspace cannot choose the requested `model`; omit
// `model` or send `auto`. Returns 409 when the automation was changed while the
// request ran.
//
// **[Required scope](/using-the-api/scopes/):** `automations:update`
//
// **[Rate limit category](/using-the-api/rate-limits/):** Write
func (r *AutomationService) Update(ctx context.Context, id string, body AutomationUpdateParams, opts ...option.RequestOption) (res *AutomationUpdateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/automations/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// > [!WARNING] This endpoint is in beta and isn't available to every workspace
// > yet. Its parameters and response schemas may change, including breaking
// > changes.
//
// Returns a preview of each of the workspace’s live automations, newest first: its
// name, status, trigger types, step count and when it last ran and last failed.
// Retrieve an automation for its steps and trigger details. Automations never
// published from the app and personal assistant automations are not listed.
//
// Requires an API key whose user is a workspace owner or admin.
//
// **[Required scope](/using-the-api/scopes/):** `automations:read`
//
// **[Rate limit category](/using-the-api/rate-limits/):** Search
func (r *AutomationService) List(ctx context.Context, query AutomationListParams, opts ...option.RequestOption) (res *AutomationListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/automations"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// > [!WARNING] This endpoint is in beta and isn't available to every workspace
// > yet. Its parameters and response schemas may change, including breaking
// > changes.
//
// Deletes an automation, turns off its triggers and dismisses the runs it has
// waiting for review. This cannot be undone; deleting it again returns 404.
//
// Requires an API key whose user is a workspace owner or admin.
//
// **[Required scope](/using-the-api/scopes/):** `automations:delete`
//
// **[Rate limit category](/using-the-api/rate-limits/):** Write
func (r *AutomationService) Delete(ctx context.Context, id string, body AutomationDeleteParams, opts ...option.RequestOption) (res *AutomationDeleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/automations/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &res, opts...)
	return res, err
}

// > [!WARNING] This endpoint is in beta and isn't available to every workspace
// > yet. Its parameters and response schemas may change, including breaking
// > changes.
//
// Stops an automation from starting runs. A paused automation is returned
// unchanged.
//
// Requires an API key whose user is a workspace owner or admin.
//
// **[Required scope](/using-the-api/scopes/):** `automations:update`
//
// **[Rate limit category](/using-the-api/rate-limits/):** Write
func (r *AutomationService) Pause(ctx context.Context, id string, body AutomationPauseParams, opts ...option.RequestOption) (res *AutomationPauseResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/automations/%s/pause", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// > [!WARNING] This endpoint is in beta and isn't available to every workspace
// > yet. Its parameters and response schemas may change, including breaking
// > changes.
//
// Lets a paused automation start runs again. An active automation is returned
// unchanged.
//
// Requires an API key whose user is a workspace owner or admin.
//
// **[Required scope](/using-the-api/scopes/):** `automations:update`
//
// **[Rate limit category](/using-the-api/rate-limits/):** Write
func (r *AutomationService) Resume(ctx context.Context, id string, body AutomationResumeParams, opts ...option.RequestOption) (res *AutomationResumeResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/automations/%s/resume", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

type AutomationCreateResponse struct {
	// Unique identifier of the automation.
	ID string `json:"id" api:"required"`
	// When the automation was created.
	CreatedAt string `json:"createdAt" api:"required"`
	// What the automation is meant to accomplish.
	Description string `json:"description" api:"required"`
	// When its most recent failed run ended, test runs included; null if no run has
	// failed.
	LastFailureAt string `json:"lastFailureAt" api:"required"`
	// When its most recent run started, test runs included; null before the first run.
	LastRunAt string `json:"lastRunAt" api:"required"`
	// Model the agent steps run on.
	Model string `json:"model" api:"required"`
	// Display name of the automation.
	Name string `json:"name" api:"required"`
	// Always `automation`.
	//
	// Any of "automation".
	ObjectType AutomationCreateResponseObjectType `json:"objectType" api:"required"`
	// Who the automation acts as.
	RunAs AutomationCreateResponseRunAsUnion `json:"runAs" api:"required"`
	// `PAUSED` automations keep their triggers but start no runs.
	//
	// Any of "ACTIVE", "PAUSED".
	Status   AutomationCreateResponseStatus         `json:"status" api:"required"`
	Steps    []AutomationCreateResponseStepUnion    `json:"steps" api:"required"`
	Triggers []AutomationCreateResponseTriggerUnion `json:"triggers" api:"required"`
	// When the automation was last changed.
	UpdatedAt string `json:"updatedAt" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID            respjson.Field
		CreatedAt     respjson.Field
		Description   respjson.Field
		LastFailureAt respjson.Field
		LastRunAt     respjson.Field
		Model         respjson.Field
		Name          respjson.Field
		ObjectType    respjson.Field
		RunAs         respjson.Field
		Status        respjson.Field
		Steps         respjson.Field
		Triggers      respjson.Field
		UpdatedAt     respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutomationCreateResponse) RawJSON() string { return r.JSON.raw }
func (r *AutomationCreateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Always `automation`.
type AutomationCreateResponseObjectType string

const (
	AutomationCreateResponseObjectTypeAutomation AutomationCreateResponseObjectType = "automation"
)

// AutomationCreateResponseRunAsUnion contains all possible properties and values
// from [AutomationCreateResponseRunAsWorkspaceRunAs],
// [AutomationCreateResponseRunAsMemberRunAs],
// [AutomationCreateResponseRunAsFromInputRunAs].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type AutomationCreateResponseRunAsUnion struct {
	Type string `json:"type"`
	// This field is from variant [AutomationCreateResponseRunAsMemberRunAs].
	MemberID string `json:"memberId"`
	JSON     struct {
		Type     respjson.Field
		MemberID respjson.Field
		raw      string
	} `json:"-"`
}

func (u AutomationCreateResponseRunAsUnion) AsWorkspaceRunAs() (v AutomationCreateResponseRunAsWorkspaceRunAs) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomationCreateResponseRunAsUnion) AsMemberRunAs() (v AutomationCreateResponseRunAsMemberRunAs) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomationCreateResponseRunAsUnion) AsFromInputRunAs() (v AutomationCreateResponseRunAsFromInputRunAs) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AutomationCreateResponseRunAsUnion) RawJSON() string { return u.JSON.raw }

func (r *AutomationCreateResponseRunAsUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Runs with no acting member.
type AutomationCreateResponseRunAsWorkspaceRunAs struct {
	// Any of "WORKSPACE".
	Type string `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutomationCreateResponseRunAsWorkspaceRunAs) RawJSON() string { return r.JSON.raw }
func (r *AutomationCreateResponseRunAsWorkspaceRunAs) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AutomationCreateResponseRunAsMemberRunAs struct {
	// Member the automation runs as.
	MemberID string `json:"memberId" api:"required"`
	// Any of "MEMBER".
	Type string `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		MemberID    respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutomationCreateResponseRunAsMemberRunAs) RawJSON() string { return r.JSON.raw }
func (r *AutomationCreateResponseRunAsMemberRunAs) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The acting member is read from an input when a run starts. Set in the Lightfield
// app.
type AutomationCreateResponseRunAsFromInputRunAs struct {
	// Any of "FROM_INPUT".
	Type string `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutomationCreateResponseRunAsFromInputRunAs) RawJSON() string { return r.JSON.raw }
func (r *AutomationCreateResponseRunAsFromInputRunAs) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// `PAUSED` automations keep their triggers but start no runs.
type AutomationCreateResponseStatus string

const (
	AutomationCreateResponseStatusActive AutomationCreateResponseStatus = "ACTIVE"
	AutomationCreateResponseStatusPaused AutomationCreateResponseStatus = "PAUSED"
)

// AutomationCreateResponseStepUnion contains all possible properties and values
// from [AutomationCreateResponseStepAgentStep],
// [AutomationCreateResponseStepCodeStep], [AutomationCreateResponseStepSleepStep].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type AutomationCreateResponseStepUnion struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	// This field is from variant [AutomationCreateResponseStepAgentStep].
	Prompt string `json:"prompt"`
	Type   string `json:"type"`
	// This field is a union of
	// [[]AutomationCreateResponseStepAgentStepPermissionUnion],
	// [[]AutomationCreateResponseStepCodeStepPermissionUnion]
	Permissions AutomationCreateResponseStepUnionPermissions `json:"permissions"`
	// This field is from variant [AutomationCreateResponseStepCodeStep].
	Command string `json:"command"`
	// This field is from variant [AutomationCreateResponseStepCodeStep].
	Description string `json:"description"`
	// This field is from variant [AutomationCreateResponseStepSleepStep].
	Duration AutomationCreateResponseStepSleepStepDurationUnion `json:"duration"`
	// This field is from variant [AutomationCreateResponseStepSleepStep].
	Unit string `json:"unit"`
	JSON struct {
		ID          respjson.Field
		Name        respjson.Field
		Prompt      respjson.Field
		Type        respjson.Field
		Permissions respjson.Field
		Command     respjson.Field
		Description respjson.Field
		Duration    respjson.Field
		Unit        respjson.Field
		raw         string
	} `json:"-"`
}

func (u AutomationCreateResponseStepUnion) AsAgentStep() (v AutomationCreateResponseStepAgentStep) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomationCreateResponseStepUnion) AsCodeStep() (v AutomationCreateResponseStepCodeStep) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomationCreateResponseStepUnion) AsSleepStep() (v AutomationCreateResponseStepSleepStep) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AutomationCreateResponseStepUnion) RawJSON() string { return u.JSON.raw }

func (r *AutomationCreateResponseStepUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AutomationCreateResponseStepUnionPermissions is an implicit subunion of
// [AutomationCreateResponseStepUnion].
// AutomationCreateResponseStepUnionPermissions provides convenient access to the
// sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [AutomationCreateResponseStepUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfAutomationCreateResponseStepAgentStepPermissions
// OfAutomationCreateResponseStepCodeStepPermissions]
type AutomationCreateResponseStepUnionPermissions struct {
	// This field will be present if the value is a
	// [[]AutomationCreateResponseStepAgentStepPermissionUnion] instead of an object.
	OfAutomationCreateResponseStepAgentStepPermissions []AutomationCreateResponseStepAgentStepPermissionUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]AutomationCreateResponseStepCodeStepPermissionUnion] instead of an object.
	OfAutomationCreateResponseStepCodeStepPermissions []AutomationCreateResponseStepCodeStepPermissionUnion `json:",inline"`
	JSON                                              struct {
		OfAutomationCreateResponseStepAgentStepPermissions respjson.Field
		OfAutomationCreateResponseStepCodeStepPermissions  respjson.Field
		raw                                                string
	} `json:"-"`
}

func (r *AutomationCreateResponseStepUnionPermissions) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AutomationCreateResponseStepAgentStep struct {
	// Step identifier. Assigned by Lightfield.
	ID string `json:"id" api:"required"`
	// Display name of the step.
	Name string `json:"name" api:"required"`
	// Instructions the agent follows in this step.
	Prompt string `json:"prompt" api:"required"`
	// Any of "AGENT".
	Type string `json:"type" api:"required"`
	// What the agent may use in this step. Defaults to none.
	Permissions []AutomationCreateResponseStepAgentStepPermissionUnion `json:"permissions"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Name        respjson.Field
		Prompt      respjson.Field
		Type        respjson.Field
		Permissions respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutomationCreateResponseStepAgentStep) RawJSON() string { return r.JSON.raw }
func (r *AutomationCreateResponseStepAgentStep) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AutomationCreateResponseStepAgentStepPermissionUnion contains all possible
// properties and values from
// [AutomationCreateResponseStepAgentStepPermissionToolPermission],
// [AutomationCreateResponseStepAgentStepPermissionMcpPermission],
// [AutomationCreateResponseStepAgentStepPermissionAPIPermission],
// [AutomationCreateResponseStepAgentStepPermissionNetworkPermission].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type AutomationCreateResponseStepAgentStepPermissionUnion struct {
	// This field is from variant
	// [AutomationCreateResponseStepAgentStepPermissionToolPermission].
	Name   string `json:"name"`
	Type   string `json:"type"`
	Reason string `json:"reason"`
	// This field is from variant
	// [AutomationCreateResponseStepAgentStepPermissionMcpPermission].
	Connector string `json:"connector"`
	// This field is from variant
	// [AutomationCreateResponseStepAgentStepPermissionMcpPermission].
	Tools []string `json:"tools"`
	// This field is from variant
	// [AutomationCreateResponseStepAgentStepPermissionAPIPermission].
	Scopes []string `json:"scopes"`
	// This field is from variant
	// [AutomationCreateResponseStepAgentStepPermissionNetworkPermission].
	Domain string `json:"domain"`
	// This field is from variant
	// [AutomationCreateResponseStepAgentStepPermissionNetworkPermission].
	Headers map[string]string `json:"headers"`
	// This field is from variant
	// [AutomationCreateResponseStepAgentStepPermissionNetworkPermission].
	Method string `json:"method"`
	JSON   struct {
		Name      respjson.Field
		Type      respjson.Field
		Reason    respjson.Field
		Connector respjson.Field
		Tools     respjson.Field
		Scopes    respjson.Field
		Domain    respjson.Field
		Headers   respjson.Field
		Method    respjson.Field
		raw       string
	} `json:"-"`
}

func (u AutomationCreateResponseStepAgentStepPermissionUnion) AsToolPermission() (v AutomationCreateResponseStepAgentStepPermissionToolPermission) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomationCreateResponseStepAgentStepPermissionUnion) AsMcpPermission() (v AutomationCreateResponseStepAgentStepPermissionMcpPermission) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomationCreateResponseStepAgentStepPermissionUnion) AsAPIPermission() (v AutomationCreateResponseStepAgentStepPermissionAPIPermission) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomationCreateResponseStepAgentStepPermissionUnion) AsNetworkPermission() (v AutomationCreateResponseStepAgentStepPermissionNetworkPermission) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AutomationCreateResponseStepAgentStepPermissionUnion) RawJSON() string { return u.JSON.raw }

func (r *AutomationCreateResponseStepAgentStepPermissionUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AutomationCreateResponseStepAgentStepPermissionToolPermission struct {
	// Name of the tool, matching the Lightfield app.
	Name string `json:"name" api:"required"`
	// Any of "TOOL".
	Type string `json:"type" api:"required"`
	// Why the step needs this permission. Null when none was given. Returns 422 for an
	// empty or whitespace-only string.
	Reason string `json:"reason" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name        respjson.Field
		Type        respjson.Field
		Reason      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutomationCreateResponseStepAgentStepPermissionToolPermission) RawJSON() string {
	return r.JSON.raw
}
func (r *AutomationCreateResponseStepAgentStepPermissionToolPermission) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AutomationCreateResponseStepAgentStepPermissionMcpPermission struct {
	// Connector the tools are on. Must already be lowercase and trimmed.
	Connector string `json:"connector" api:"required"`
	// Unique, trimmed names of the connector tools the step may call.
	Tools []string `json:"tools" api:"required"`
	// Any of "MCP".
	Type string `json:"type" api:"required"`
	// Why the step needs this permission. Null when none was given. Returns 422 for an
	// empty or whitespace-only string.
	Reason string `json:"reason" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Connector   respjson.Field
		Tools       respjson.Field
		Type        respjson.Field
		Reason      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutomationCreateResponseStepAgentStepPermissionMcpPermission) RawJSON() string {
	return r.JSON.raw
}
func (r *AutomationCreateResponseStepAgentStepPermissionMcpPermission) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AutomationCreateResponseStepAgentStepPermissionAPIPermission struct {
	// Lightfield API scopes the step may use, for example `contacts:read`. The API key
	// must hold every scope it grants, and only a full-access key may grant `*`.
	Scopes []string `json:"scopes" api:"required"`
	// Any of "API".
	Type string `json:"type" api:"required"`
	// Why the step needs this permission. Null when none was given. Returns 422 for an
	// empty or whitespace-only string.
	Reason string `json:"reason" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Scopes      respjson.Field
		Type        respjson.Field
		Reason      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutomationCreateResponseStepAgentStepPermissionAPIPermission) RawJSON() string {
	return r.JSON.raw
}
func (r *AutomationCreateResponseStepAgentStepPermissionAPIPermission) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AutomationCreateResponseStepAgentStepPermissionNetworkPermission struct {
	// Lowercase host the step may call, optionally followed by a path glob. Returns
	// 422 for a scheme, query, fragment or surrounding whitespace.
	Domain string `json:"domain" api:"required"`
	// Any of "NETWORK".
	Type string `json:"type" api:"required"`
	// Headers added to every request this permission allows.
	Headers map[string]string `json:"headers"`
	// Restricts the permission to one HTTP method.
	//
	// Any of "GET", "POST", "PUT", "PATCH", "DELETE".
	Method string `json:"method"`
	// Why the step needs this permission. Null when none was given. Returns 422 for an
	// empty or whitespace-only string.
	Reason string `json:"reason" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Domain      respjson.Field
		Type        respjson.Field
		Headers     respjson.Field
		Method      respjson.Field
		Reason      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutomationCreateResponseStepAgentStepPermissionNetworkPermission) RawJSON() string {
	return r.JSON.raw
}
func (r *AutomationCreateResponseStepAgentStepPermissionNetworkPermission) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AutomationCreateResponseStepCodeStep struct {
	// Step identifier. Assigned by Lightfield.
	ID string `json:"id" api:"required"`
	// Bash command the step runs as written.
	Command string `json:"command" api:"required"`
	// Display name of the step.
	Name string `json:"name" api:"required"`
	// Any of "CODE".
	Type string `json:"type" api:"required"`
	// What the command does. Does not change what runs.
	Description string `json:"description"`
	// What the command may use. Defaults to none.
	Permissions []AutomationCreateResponseStepCodeStepPermissionUnion `json:"permissions"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Command     respjson.Field
		Name        respjson.Field
		Type        respjson.Field
		Description respjson.Field
		Permissions respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutomationCreateResponseStepCodeStep) RawJSON() string { return r.JSON.raw }
func (r *AutomationCreateResponseStepCodeStep) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AutomationCreateResponseStepCodeStepPermissionUnion contains all possible
// properties and values from
// [AutomationCreateResponseStepCodeStepPermissionAPIPermission],
// [AutomationCreateResponseStepCodeStepPermissionNetworkPermission].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type AutomationCreateResponseStepCodeStepPermissionUnion struct {
	// This field is from variant
	// [AutomationCreateResponseStepCodeStepPermissionAPIPermission].
	Scopes []string `json:"scopes"`
	Type   string   `json:"type"`
	Reason string   `json:"reason"`
	// This field is from variant
	// [AutomationCreateResponseStepCodeStepPermissionNetworkPermission].
	Domain string `json:"domain"`
	// This field is from variant
	// [AutomationCreateResponseStepCodeStepPermissionNetworkPermission].
	Headers map[string]string `json:"headers"`
	// This field is from variant
	// [AutomationCreateResponseStepCodeStepPermissionNetworkPermission].
	Method string `json:"method"`
	JSON   struct {
		Scopes  respjson.Field
		Type    respjson.Field
		Reason  respjson.Field
		Domain  respjson.Field
		Headers respjson.Field
		Method  respjson.Field
		raw     string
	} `json:"-"`
}

func (u AutomationCreateResponseStepCodeStepPermissionUnion) AsAPIPermission() (v AutomationCreateResponseStepCodeStepPermissionAPIPermission) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomationCreateResponseStepCodeStepPermissionUnion) AsNetworkPermission() (v AutomationCreateResponseStepCodeStepPermissionNetworkPermission) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AutomationCreateResponseStepCodeStepPermissionUnion) RawJSON() string { return u.JSON.raw }

func (r *AutomationCreateResponseStepCodeStepPermissionUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AutomationCreateResponseStepCodeStepPermissionAPIPermission struct {
	// Lightfield API scopes the step may use, for example `contacts:read`. The API key
	// must hold every scope it grants, and only a full-access key may grant `*`.
	Scopes []string `json:"scopes" api:"required"`
	// Any of "API".
	Type string `json:"type" api:"required"`
	// Why the step needs this permission. Null when none was given. Returns 422 for an
	// empty or whitespace-only string.
	Reason string `json:"reason" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Scopes      respjson.Field
		Type        respjson.Field
		Reason      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutomationCreateResponseStepCodeStepPermissionAPIPermission) RawJSON() string {
	return r.JSON.raw
}
func (r *AutomationCreateResponseStepCodeStepPermissionAPIPermission) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AutomationCreateResponseStepCodeStepPermissionNetworkPermission struct {
	// Lowercase host the step may call, optionally followed by a path glob. Returns
	// 422 for a scheme, query, fragment or surrounding whitespace.
	Domain string `json:"domain" api:"required"`
	// Any of "NETWORK".
	Type string `json:"type" api:"required"`
	// Headers added to every request this permission allows.
	Headers map[string]string `json:"headers"`
	// Restricts the permission to one HTTP method.
	//
	// Any of "GET", "POST", "PUT", "PATCH", "DELETE".
	Method string `json:"method"`
	// Why the step needs this permission. Null when none was given. Returns 422 for an
	// empty or whitespace-only string.
	Reason string `json:"reason" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Domain      respjson.Field
		Type        respjson.Field
		Headers     respjson.Field
		Method      respjson.Field
		Reason      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutomationCreateResponseStepCodeStepPermissionNetworkPermission) RawJSON() string {
	return r.JSON.raw
}
func (r *AutomationCreateResponseStepCodeStepPermissionNetworkPermission) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AutomationCreateResponseStepSleepStep struct {
	// Step identifier. Assigned by Lightfield.
	ID string `json:"id" api:"required"`
	// How long to sleep before the next step, at most one year.
	Duration AutomationCreateResponseStepSleepStepDurationUnion `json:"duration" api:"required"`
	// Display name of the step.
	Name string `json:"name" api:"required"`
	// Any of "SLEEP".
	Type string `json:"type" api:"required"`
	// Unit of `duration`.
	//
	// Any of "MILLISECOND", "SECOND", "MINUTE", "HOUR", "DAY".
	Unit string `json:"unit" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Duration    respjson.Field
		Name        respjson.Field
		Type        respjson.Field
		Unit        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutomationCreateResponseStepSleepStep) RawJSON() string { return r.JSON.raw }
func (r *AutomationCreateResponseStepSleepStep) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AutomationCreateResponseStepSleepStepDurationUnion contains all possible
// properties and values from [int64], [string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfInt OfString]
type AutomationCreateResponseStepSleepStepDurationUnion struct {
	// This field will be present if the value is a [int64] instead of an object.
	OfInt int64 `json:",inline"`
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	JSON     struct {
		OfInt    respjson.Field
		OfString respjson.Field
		raw      string
	} `json:"-"`
}

func (u AutomationCreateResponseStepSleepStepDurationUnion) AsInt() (v int64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomationCreateResponseStepSleepStepDurationUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AutomationCreateResponseStepSleepStepDurationUnion) RawJSON() string { return u.JSON.raw }

func (r *AutomationCreateResponseStepSleepStepDurationUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AutomationCreateResponseTriggerUnion contains all possible properties and values
// from [AutomationCreateResponseTriggerManualTrigger],
// [AutomationCreateResponseTriggerWebhookTrigger],
// [AutomationCreateResponseTriggerScheduledTrigger],
// [AutomationCreateResponseTriggerRecordCreatedTrigger],
// [AutomationCreateResponseTriggerRecordUpdatedTrigger],
// [AutomationCreateResponseTriggerSequenceReplyTrigger].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type AutomationCreateResponseTriggerUnion struct {
	Type string `json:"type"`
	// This field is from variant [AutomationCreateResponseTriggerWebhookTrigger].
	URL string `json:"url"`
	// This field is from variant [AutomationCreateResponseTriggerWebhookTrigger].
	ExampleResult AutomationCreateResponseTriggerWebhookTriggerExampleResult `json:"exampleResult"`
	// This field is from variant [AutomationCreateResponseTriggerScheduledTrigger].
	Schedule   AutomationCreateResponseTriggerScheduledTriggerSchedule `json:"schedule"`
	ObjectType string                                                  `json:"objectType"`
	// This field is from variant
	// [AutomationCreateResponseTriggerRecordUpdatedTrigger].
	WatchedFields []string `json:"watchedFields"`
	// This field is from variant
	// [AutomationCreateResponseTriggerRecordUpdatedTrigger].
	WatchedRelationships []AutomationCreateResponseTriggerRecordUpdatedTriggerWatchedRelationship `json:"watchedRelationships"`
	JSON                 struct {
		Type                 respjson.Field
		URL                  respjson.Field
		ExampleResult        respjson.Field
		Schedule             respjson.Field
		ObjectType           respjson.Field
		WatchedFields        respjson.Field
		WatchedRelationships respjson.Field
		raw                  string
	} `json:"-"`
}

func (u AutomationCreateResponseTriggerUnion) AsManualTrigger() (v AutomationCreateResponseTriggerManualTrigger) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomationCreateResponseTriggerUnion) AsWebhookTrigger() (v AutomationCreateResponseTriggerWebhookTrigger) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomationCreateResponseTriggerUnion) AsScheduledTrigger() (v AutomationCreateResponseTriggerScheduledTrigger) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomationCreateResponseTriggerUnion) AsRecordCreatedTrigger() (v AutomationCreateResponseTriggerRecordCreatedTrigger) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomationCreateResponseTriggerUnion) AsRecordUpdatedTrigger() (v AutomationCreateResponseTriggerRecordUpdatedTrigger) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomationCreateResponseTriggerUnion) AsSequenceReplyTrigger() (v AutomationCreateResponseTriggerSequenceReplyTrigger) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AutomationCreateResponseTriggerUnion) RawJSON() string { return u.JSON.raw }

func (r *AutomationCreateResponseTriggerUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AutomationCreateResponseTriggerManualTrigger struct {
	// Any of "MANUAL".
	Type string `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutomationCreateResponseTriggerManualTrigger) RawJSON() string { return r.JSON.raw }
func (r *AutomationCreateResponseTriggerManualTrigger) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AutomationCreateResponseTriggerWebhookTrigger struct {
	// Any of "WEBHOOK".
	Type string `json:"type" api:"required"`
	// URL that starts a run when it receives a POST request.
	URL string `json:"url" api:"required"`
	// A sample payload and the fields the steps read from it.
	ExampleResult AutomationCreateResponseTriggerWebhookTriggerExampleResult `json:"exampleResult"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type          respjson.Field
		URL           respjson.Field
		ExampleResult respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutomationCreateResponseTriggerWebhookTrigger) RawJSON() string { return r.JSON.raw }
func (r *AutomationCreateResponseTriggerWebhookTrigger) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A sample payload and the fields the steps read from it.
type AutomationCreateResponseTriggerWebhookTriggerExampleResult struct {
	Json   AutomationCreateResponseTriggerWebhookTriggerExampleResultJsonUnion `json:"json" api:"required"`
	Parsed []AutomationCreateResponseTriggerWebhookTriggerExampleResultParsed  `json:"parsed" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Json        respjson.Field
		Parsed      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutomationCreateResponseTriggerWebhookTriggerExampleResult) RawJSON() string {
	return r.JSON.raw
}
func (r *AutomationCreateResponseTriggerWebhookTriggerExampleResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AutomationCreateResponseTriggerWebhookTriggerExampleResultJsonUnion contains all
// possible properties and values from
// [map[string]AutomationCreateResponseTriggerWebhookTriggerExampleResultJsonMapItemUnion],
// [[]AutomationCreateResponseTriggerWebhookTriggerExampleResultJsonArrayItemUnion].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool
// OfAutomationCreateResponseTriggerWebhookTriggerExampleResultJsonMapItemArray
// OfJsonValueArray OfAnyArray
// OfAutomationCreateResponseTriggerWebhookTriggerExampleResultJsonArray]
type AutomationCreateResponseTriggerWebhookTriggerExampleResultJsonUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a
	// [[]AutomationCreateResponseTriggerWebhookTriggerExampleResultJsonMapItemArrayItemUnion]
	// instead of an object.
	OfAutomationCreateResponseTriggerWebhookTriggerExampleResultJsonMapItemArray []AutomationCreateResponseTriggerWebhookTriggerExampleResultJsonMapItemArrayItemUnion `json:",inline"`
	// This field will be present if the value is a [[]shared.JsonValueUnion] instead
	// of an object.
	OfJsonValueArray []shared.JsonValueUnion `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	// This field will be present if the value is a
	// [[]AutomationCreateResponseTriggerWebhookTriggerExampleResultJsonArrayItemUnion]
	// instead of an object.
	OfAutomationCreateResponseTriggerWebhookTriggerExampleResultJsonArray []AutomationCreateResponseTriggerWebhookTriggerExampleResultJsonArrayItemUnion `json:",inline"`
	JSON                                                                  struct {
		OfString                                                                     respjson.Field
		OfFloat                                                                      respjson.Field
		OfBool                                                                       respjson.Field
		OfAutomationCreateResponseTriggerWebhookTriggerExampleResultJsonMapItemArray respjson.Field
		OfJsonValueArray                                                             respjson.Field
		OfAnyArray                                                                   respjson.Field
		OfAutomationCreateResponseTriggerWebhookTriggerExampleResultJsonArray        respjson.Field
		raw                                                                          string
	} `json:"-"`
}

func (u AutomationCreateResponseTriggerWebhookTriggerExampleResultJsonUnion) AsAutomationCreateResponseTriggerWebhookTriggerExampleResultJsonMapMap() (v map[string]AutomationCreateResponseTriggerWebhookTriggerExampleResultJsonMapItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomationCreateResponseTriggerWebhookTriggerExampleResultJsonUnion) AsAutomationCreateResponseTriggerWebhookTriggerExampleResultJsonArray() (v []AutomationCreateResponseTriggerWebhookTriggerExampleResultJsonArrayItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AutomationCreateResponseTriggerWebhookTriggerExampleResultJsonUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *AutomationCreateResponseTriggerWebhookTriggerExampleResultJsonUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AutomationCreateResponseTriggerWebhookTriggerExampleResultJsonMapItemUnion
// contains all possible properties and values from [string], [float64], [bool],
// [[]AutomationCreateResponseTriggerWebhookTriggerExampleResultJsonMapItemArrayItemUnion],
// [map[string]AutomationCreateResponseTriggerWebhookTriggerExampleResultJsonMapItemMapItemUnion].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool
// OfAutomationCreateResponseTriggerWebhookTriggerExampleResultJsonMapItemArray
// OfJsonValueArray OfAnyArray]
type AutomationCreateResponseTriggerWebhookTriggerExampleResultJsonMapItemUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a
	// [[]AutomationCreateResponseTriggerWebhookTriggerExampleResultJsonMapItemArrayItemUnion]
	// instead of an object.
	OfAutomationCreateResponseTriggerWebhookTriggerExampleResultJsonMapItemArray []AutomationCreateResponseTriggerWebhookTriggerExampleResultJsonMapItemArrayItemUnion `json:",inline"`
	// This field will be present if the value is a [[]shared.JsonValueUnion] instead
	// of an object.
	OfJsonValueArray []shared.JsonValueUnion `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	JSON       struct {
		OfString                                                                     respjson.Field
		OfFloat                                                                      respjson.Field
		OfBool                                                                       respjson.Field
		OfAutomationCreateResponseTriggerWebhookTriggerExampleResultJsonMapItemArray respjson.Field
		OfJsonValueArray                                                             respjson.Field
		OfAnyArray                                                                   respjson.Field
		raw                                                                          string
	} `json:"-"`
}

func (u AutomationCreateResponseTriggerWebhookTriggerExampleResultJsonMapItemUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomationCreateResponseTriggerWebhookTriggerExampleResultJsonMapItemUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomationCreateResponseTriggerWebhookTriggerExampleResultJsonMapItemUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomationCreateResponseTriggerWebhookTriggerExampleResultJsonMapItemUnion) AsAutomationCreateResponseTriggerWebhookTriggerExampleResultJsonMapItemArray() (v []AutomationCreateResponseTriggerWebhookTriggerExampleResultJsonMapItemArrayItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomationCreateResponseTriggerWebhookTriggerExampleResultJsonMapItemUnion) AsAutomationCreateResponseTriggerWebhookTriggerExampleResultJsonMapItemMapMap() (v map[string]AutomationCreateResponseTriggerWebhookTriggerExampleResultJsonMapItemMapItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AutomationCreateResponseTriggerWebhookTriggerExampleResultJsonMapItemUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *AutomationCreateResponseTriggerWebhookTriggerExampleResultJsonMapItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AutomationCreateResponseTriggerWebhookTriggerExampleResultJsonMapItemArrayItemUnion
// contains all possible properties and values from [string], [float64], [bool],
// [[]shared.JsonValueUnion], [map[string]shared.JsonValueUnion].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfJsonValueArray OfAnyArray]
type AutomationCreateResponseTriggerWebhookTriggerExampleResultJsonMapItemArrayItemUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]shared.JsonValueUnion] instead
	// of an object.
	OfJsonValueArray []shared.JsonValueUnion `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	JSON       struct {
		OfString         respjson.Field
		OfFloat          respjson.Field
		OfBool           respjson.Field
		OfJsonValueArray respjson.Field
		OfAnyArray       respjson.Field
		raw              string
	} `json:"-"`
}

func (u AutomationCreateResponseTriggerWebhookTriggerExampleResultJsonMapItemArrayItemUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomationCreateResponseTriggerWebhookTriggerExampleResultJsonMapItemArrayItemUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomationCreateResponseTriggerWebhookTriggerExampleResultJsonMapItemArrayItemUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomationCreateResponseTriggerWebhookTriggerExampleResultJsonMapItemArrayItemUnion) AsJsonValueArray() (v []shared.JsonValueUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomationCreateResponseTriggerWebhookTriggerExampleResultJsonMapItemArrayItemUnion) AsJsonValueMap() (v map[string]shared.JsonValueUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AutomationCreateResponseTriggerWebhookTriggerExampleResultJsonMapItemArrayItemUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *AutomationCreateResponseTriggerWebhookTriggerExampleResultJsonMapItemArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AutomationCreateResponseTriggerWebhookTriggerExampleResultJsonMapItemMapItemUnion
// contains all possible properties and values from [string], [float64], [bool],
// [[]shared.JsonValueUnion], [map[string]shared.JsonValueUnion].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfJsonValueArray OfAnyArray]
type AutomationCreateResponseTriggerWebhookTriggerExampleResultJsonMapItemMapItemUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]shared.JsonValueUnion] instead
	// of an object.
	OfJsonValueArray []shared.JsonValueUnion `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	JSON       struct {
		OfString         respjson.Field
		OfFloat          respjson.Field
		OfBool           respjson.Field
		OfJsonValueArray respjson.Field
		OfAnyArray       respjson.Field
		raw              string
	} `json:"-"`
}

func (u AutomationCreateResponseTriggerWebhookTriggerExampleResultJsonMapItemMapItemUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomationCreateResponseTriggerWebhookTriggerExampleResultJsonMapItemMapItemUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomationCreateResponseTriggerWebhookTriggerExampleResultJsonMapItemMapItemUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomationCreateResponseTriggerWebhookTriggerExampleResultJsonMapItemMapItemUnion) AsJsonValueArray() (v []shared.JsonValueUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomationCreateResponseTriggerWebhookTriggerExampleResultJsonMapItemMapItemUnion) AsJsonValueMap() (v map[string]shared.JsonValueUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AutomationCreateResponseTriggerWebhookTriggerExampleResultJsonMapItemMapItemUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *AutomationCreateResponseTriggerWebhookTriggerExampleResultJsonMapItemMapItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AutomationCreateResponseTriggerWebhookTriggerExampleResultJsonArrayItemUnion
// contains all possible properties and values from [string], [float64], [bool],
// [[]AutomationCreateResponseTriggerWebhookTriggerExampleResultJsonArrayItemArrayItemUnion],
// [map[string]AutomationCreateResponseTriggerWebhookTriggerExampleResultJsonArrayItemMapItemUnion].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool
// OfAutomationCreateResponseTriggerWebhookTriggerExampleResultJsonArrayItemArray
// OfJsonValueArray OfAnyArray]
type AutomationCreateResponseTriggerWebhookTriggerExampleResultJsonArrayItemUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a
	// [[]AutomationCreateResponseTriggerWebhookTriggerExampleResultJsonArrayItemArrayItemUnion]
	// instead of an object.
	OfAutomationCreateResponseTriggerWebhookTriggerExampleResultJsonArrayItemArray []AutomationCreateResponseTriggerWebhookTriggerExampleResultJsonArrayItemArrayItemUnion `json:",inline"`
	// This field will be present if the value is a [[]shared.JsonValueUnion] instead
	// of an object.
	OfJsonValueArray []shared.JsonValueUnion `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	JSON       struct {
		OfString                                                                       respjson.Field
		OfFloat                                                                        respjson.Field
		OfBool                                                                         respjson.Field
		OfAutomationCreateResponseTriggerWebhookTriggerExampleResultJsonArrayItemArray respjson.Field
		OfJsonValueArray                                                               respjson.Field
		OfAnyArray                                                                     respjson.Field
		raw                                                                            string
	} `json:"-"`
}

func (u AutomationCreateResponseTriggerWebhookTriggerExampleResultJsonArrayItemUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomationCreateResponseTriggerWebhookTriggerExampleResultJsonArrayItemUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomationCreateResponseTriggerWebhookTriggerExampleResultJsonArrayItemUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomationCreateResponseTriggerWebhookTriggerExampleResultJsonArrayItemUnion) AsAutomationCreateResponseTriggerWebhookTriggerExampleResultJsonArrayItemArray() (v []AutomationCreateResponseTriggerWebhookTriggerExampleResultJsonArrayItemArrayItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomationCreateResponseTriggerWebhookTriggerExampleResultJsonArrayItemUnion) AsAutomationCreateResponseTriggerWebhookTriggerExampleResultJsonArrayItemMapMap() (v map[string]AutomationCreateResponseTriggerWebhookTriggerExampleResultJsonArrayItemMapItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AutomationCreateResponseTriggerWebhookTriggerExampleResultJsonArrayItemUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *AutomationCreateResponseTriggerWebhookTriggerExampleResultJsonArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AutomationCreateResponseTriggerWebhookTriggerExampleResultJsonArrayItemArrayItemUnion
// contains all possible properties and values from [string], [float64], [bool],
// [[]shared.JsonValueUnion], [map[string]shared.JsonValueUnion].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfJsonValueArray OfAnyArray]
type AutomationCreateResponseTriggerWebhookTriggerExampleResultJsonArrayItemArrayItemUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]shared.JsonValueUnion] instead
	// of an object.
	OfJsonValueArray []shared.JsonValueUnion `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	JSON       struct {
		OfString         respjson.Field
		OfFloat          respjson.Field
		OfBool           respjson.Field
		OfJsonValueArray respjson.Field
		OfAnyArray       respjson.Field
		raw              string
	} `json:"-"`
}

func (u AutomationCreateResponseTriggerWebhookTriggerExampleResultJsonArrayItemArrayItemUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomationCreateResponseTriggerWebhookTriggerExampleResultJsonArrayItemArrayItemUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomationCreateResponseTriggerWebhookTriggerExampleResultJsonArrayItemArrayItemUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomationCreateResponseTriggerWebhookTriggerExampleResultJsonArrayItemArrayItemUnion) AsJsonValueArray() (v []shared.JsonValueUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomationCreateResponseTriggerWebhookTriggerExampleResultJsonArrayItemArrayItemUnion) AsJsonValueMap() (v map[string]shared.JsonValueUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AutomationCreateResponseTriggerWebhookTriggerExampleResultJsonArrayItemArrayItemUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *AutomationCreateResponseTriggerWebhookTriggerExampleResultJsonArrayItemArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AutomationCreateResponseTriggerWebhookTriggerExampleResultJsonArrayItemMapItemUnion
// contains all possible properties and values from [string], [float64], [bool],
// [[]shared.JsonValueUnion], [map[string]shared.JsonValueUnion].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfJsonValueArray OfAnyArray]
type AutomationCreateResponseTriggerWebhookTriggerExampleResultJsonArrayItemMapItemUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]shared.JsonValueUnion] instead
	// of an object.
	OfJsonValueArray []shared.JsonValueUnion `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	JSON       struct {
		OfString         respjson.Field
		OfFloat          respjson.Field
		OfBool           respjson.Field
		OfJsonValueArray respjson.Field
		OfAnyArray       respjson.Field
		raw              string
	} `json:"-"`
}

func (u AutomationCreateResponseTriggerWebhookTriggerExampleResultJsonArrayItemMapItemUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomationCreateResponseTriggerWebhookTriggerExampleResultJsonArrayItemMapItemUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomationCreateResponseTriggerWebhookTriggerExampleResultJsonArrayItemMapItemUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomationCreateResponseTriggerWebhookTriggerExampleResultJsonArrayItemMapItemUnion) AsJsonValueArray() (v []shared.JsonValueUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomationCreateResponseTriggerWebhookTriggerExampleResultJsonArrayItemMapItemUnion) AsJsonValueMap() (v map[string]shared.JsonValueUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AutomationCreateResponseTriggerWebhookTriggerExampleResultJsonArrayItemMapItemUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *AutomationCreateResponseTriggerWebhookTriggerExampleResultJsonArrayItemMapItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AutomationCreateResponseTriggerWebhookTriggerExampleResultParsed struct {
	Path  []string                                                                   `json:"path" api:"required"`
	Label string                                                                     `json:"label"`
	Value AutomationCreateResponseTriggerWebhookTriggerExampleResultParsedValueUnion `json:"value" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Path        respjson.Field
		Label       respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutomationCreateResponseTriggerWebhookTriggerExampleResultParsed) RawJSON() string {
	return r.JSON.raw
}
func (r *AutomationCreateResponseTriggerWebhookTriggerExampleResultParsed) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AutomationCreateResponseTriggerWebhookTriggerExampleResultParsedValueUnion
// contains all possible properties and values from [string], [float64], [bool],
// [[]AutomationCreateResponseTriggerWebhookTriggerExampleResultParsedValueArrayItemUnion],
// [map[string]AutomationCreateResponseTriggerWebhookTriggerExampleResultParsedValueMapItemUnion].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool
// OfAutomationCreateResponseTriggerWebhookTriggerExampleResultParsedValueArray
// OfJsonValueArray OfAnyArray]
type AutomationCreateResponseTriggerWebhookTriggerExampleResultParsedValueUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a
	// [[]AutomationCreateResponseTriggerWebhookTriggerExampleResultParsedValueArrayItemUnion]
	// instead of an object.
	OfAutomationCreateResponseTriggerWebhookTriggerExampleResultParsedValueArray []AutomationCreateResponseTriggerWebhookTriggerExampleResultParsedValueArrayItemUnion `json:",inline"`
	// This field will be present if the value is a [[]shared.JsonValueUnion] instead
	// of an object.
	OfJsonValueArray []shared.JsonValueUnion `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	JSON       struct {
		OfString                                                                     respjson.Field
		OfFloat                                                                      respjson.Field
		OfBool                                                                       respjson.Field
		OfAutomationCreateResponseTriggerWebhookTriggerExampleResultParsedValueArray respjson.Field
		OfJsonValueArray                                                             respjson.Field
		OfAnyArray                                                                   respjson.Field
		raw                                                                          string
	} `json:"-"`
}

func (u AutomationCreateResponseTriggerWebhookTriggerExampleResultParsedValueUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomationCreateResponseTriggerWebhookTriggerExampleResultParsedValueUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomationCreateResponseTriggerWebhookTriggerExampleResultParsedValueUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomationCreateResponseTriggerWebhookTriggerExampleResultParsedValueUnion) AsAutomationCreateResponseTriggerWebhookTriggerExampleResultParsedValueArray() (v []AutomationCreateResponseTriggerWebhookTriggerExampleResultParsedValueArrayItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomationCreateResponseTriggerWebhookTriggerExampleResultParsedValueUnion) AsAutomationCreateResponseTriggerWebhookTriggerExampleResultParsedValueMapMap() (v map[string]AutomationCreateResponseTriggerWebhookTriggerExampleResultParsedValueMapItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AutomationCreateResponseTriggerWebhookTriggerExampleResultParsedValueUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *AutomationCreateResponseTriggerWebhookTriggerExampleResultParsedValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AutomationCreateResponseTriggerWebhookTriggerExampleResultParsedValueArrayItemUnion
// contains all possible properties and values from [string], [float64], [bool],
// [[]shared.JsonValueUnion], [map[string]shared.JsonValueUnion].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfJsonValueArray OfAnyArray]
type AutomationCreateResponseTriggerWebhookTriggerExampleResultParsedValueArrayItemUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]shared.JsonValueUnion] instead
	// of an object.
	OfJsonValueArray []shared.JsonValueUnion `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	JSON       struct {
		OfString         respjson.Field
		OfFloat          respjson.Field
		OfBool           respjson.Field
		OfJsonValueArray respjson.Field
		OfAnyArray       respjson.Field
		raw              string
	} `json:"-"`
}

func (u AutomationCreateResponseTriggerWebhookTriggerExampleResultParsedValueArrayItemUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomationCreateResponseTriggerWebhookTriggerExampleResultParsedValueArrayItemUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomationCreateResponseTriggerWebhookTriggerExampleResultParsedValueArrayItemUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomationCreateResponseTriggerWebhookTriggerExampleResultParsedValueArrayItemUnion) AsJsonValueArray() (v []shared.JsonValueUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomationCreateResponseTriggerWebhookTriggerExampleResultParsedValueArrayItemUnion) AsJsonValueMap() (v map[string]shared.JsonValueUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AutomationCreateResponseTriggerWebhookTriggerExampleResultParsedValueArrayItemUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *AutomationCreateResponseTriggerWebhookTriggerExampleResultParsedValueArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AutomationCreateResponseTriggerWebhookTriggerExampleResultParsedValueMapItemUnion
// contains all possible properties and values from [string], [float64], [bool],
// [[]shared.JsonValueUnion], [map[string]shared.JsonValueUnion].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfJsonValueArray OfAnyArray]
type AutomationCreateResponseTriggerWebhookTriggerExampleResultParsedValueMapItemUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]shared.JsonValueUnion] instead
	// of an object.
	OfJsonValueArray []shared.JsonValueUnion `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	JSON       struct {
		OfString         respjson.Field
		OfFloat          respjson.Field
		OfBool           respjson.Field
		OfJsonValueArray respjson.Field
		OfAnyArray       respjson.Field
		raw              string
	} `json:"-"`
}

func (u AutomationCreateResponseTriggerWebhookTriggerExampleResultParsedValueMapItemUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomationCreateResponseTriggerWebhookTriggerExampleResultParsedValueMapItemUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomationCreateResponseTriggerWebhookTriggerExampleResultParsedValueMapItemUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomationCreateResponseTriggerWebhookTriggerExampleResultParsedValueMapItemUnion) AsJsonValueArray() (v []shared.JsonValueUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomationCreateResponseTriggerWebhookTriggerExampleResultParsedValueMapItemUnion) AsJsonValueMap() (v map[string]shared.JsonValueUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AutomationCreateResponseTriggerWebhookTriggerExampleResultParsedValueMapItemUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *AutomationCreateResponseTriggerWebhookTriggerExampleResultParsedValueMapItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AutomationCreateResponseTriggerScheduledTrigger struct {
	Schedule AutomationCreateResponseTriggerScheduledTriggerSchedule `json:"schedule" api:"required"`
	// Any of "SCHEDULED".
	Type string `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Schedule    respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutomationCreateResponseTriggerScheduledTrigger) RawJSON() string { return r.JSON.raw }
func (r *AutomationCreateResponseTriggerScheduledTrigger) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AutomationCreateResponseTriggerScheduledTriggerSchedule struct {
	// 5-field cron expression, for example `0 9 * * 1-5`.
	Cron string `json:"cron" api:"required"`
	// IANA timezone the cron is evaluated in.
	Timezone string `json:"timezone" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Cron        respjson.Field
		Timezone    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutomationCreateResponseTriggerScheduledTriggerSchedule) RawJSON() string { return r.JSON.raw }
func (r *AutomationCreateResponseTriggerScheduledTriggerSchedule) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AutomationCreateResponseTriggerRecordCreatedTrigger struct {
	ObjectType string `json:"objectType" api:"required"`
	// Any of "RECORD_CREATED".
	Type string `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ObjectType  respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutomationCreateResponseTriggerRecordCreatedTrigger) RawJSON() string { return r.JSON.raw }
func (r *AutomationCreateResponseTriggerRecordCreatedTrigger) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AutomationCreateResponseTriggerRecordUpdatedTrigger struct {
	ObjectType string `json:"objectType" api:"required"`
	// Any of "RECORD_UPDATED".
	Type                 string                                                                   `json:"type" api:"required"`
	WatchedFields        []string                                                                 `json:"watchedFields" api:"required"`
	WatchedRelationships []AutomationCreateResponseTriggerRecordUpdatedTriggerWatchedRelationship `json:"watchedRelationships"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ObjectType           respjson.Field
		Type                 respjson.Field
		WatchedFields        respjson.Field
		WatchedRelationships respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutomationCreateResponseTriggerRecordUpdatedTrigger) RawJSON() string { return r.JSON.raw }
func (r *AutomationCreateResponseTriggerRecordUpdatedTrigger) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AutomationCreateResponseTriggerRecordUpdatedTriggerWatchedRelationship struct {
	// The endpoint occupied by the record whose changes start the run.
	//
	// Any of "SOURCE", "DESTINATION".
	Direction string `json:"direction" api:"required"`
	// Public relationship key, such as `$contact` or a custom relationship slug.
	Key string `json:"key" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Direction   respjson.Field
		Key         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutomationCreateResponseTriggerRecordUpdatedTriggerWatchedRelationship) RawJSON() string {
	return r.JSON.raw
}
func (r *AutomationCreateResponseTriggerRecordUpdatedTriggerWatchedRelationship) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AutomationCreateResponseTriggerSequenceReplyTrigger struct {
	// Any of "SEQUENCE_REPLY".
	Type string `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutomationCreateResponseTriggerSequenceReplyTrigger) RawJSON() string { return r.JSON.raw }
func (r *AutomationCreateResponseTriggerSequenceReplyTrigger) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AutomationDeleteResponse struct {
	ID string `json:"id" api:"required"`
	// Always `true`.
	Deleted bool `json:"deleted" api:"required"`
	// Any of "automation".
	ObjectType AutomationDeleteResponseObjectType `json:"objectType" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Deleted     respjson.Field
		ObjectType  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutomationDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *AutomationDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AutomationDeleteResponseObjectType string

const (
	AutomationDeleteResponseObjectTypeAutomation AutomationDeleteResponseObjectType = "automation"
)

type AutomationListResponse struct {
	Data []AutomationListResponseData `json:"data" api:"required"`
	// Any of "list".
	Object     AutomationListResponseObject `json:"object" api:"required"`
	TotalCount int64                        `json:"totalCount" api:"required"`
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
func (r AutomationListResponse) RawJSON() string { return r.JSON.raw }
func (r *AutomationListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AutomationListResponseData struct {
	// Unique identifier of the automation.
	ID string `json:"id" api:"required"`
	// When the automation was created.
	CreatedAt string `json:"createdAt" api:"required"`
	// What the automation is meant to accomplish.
	Description string `json:"description" api:"required"`
	// When its most recent failed run ended, test runs included; null if no run has
	// failed.
	LastFailureAt string `json:"lastFailureAt" api:"required"`
	// When its most recent run started, test runs included; null before the first run.
	LastRunAt string `json:"lastRunAt" api:"required"`
	// Display name of the automation.
	Name string `json:"name" api:"required"`
	// Always `automation`.
	//
	// Any of "automation".
	ObjectType string `json:"objectType" api:"required"`
	// `PAUSED` automations keep their triggers but start no runs.
	//
	// Any of "ACTIVE", "PAUSED".
	Status string `json:"status" api:"required"`
	// How many steps a run has.
	StepCount int64 `json:"stepCount" api:"required"`
	// The kinds of trigger that start a run, each listed once.
	//
	// Any of "MANUAL", "WEBHOOK", "SCHEDULED", "RECORD_CREATED", "RECORD_UPDATED",
	// "SEQUENCE_REPLY".
	TriggerTypes []string `json:"triggerTypes" api:"required"`
	// When the automation was last changed.
	UpdatedAt string `json:"updatedAt" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID            respjson.Field
		CreatedAt     respjson.Field
		Description   respjson.Field
		LastFailureAt respjson.Field
		LastRunAt     respjson.Field
		Name          respjson.Field
		ObjectType    respjson.Field
		Status        respjson.Field
		StepCount     respjson.Field
		TriggerTypes  respjson.Field
		UpdatedAt     respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutomationListResponseData) RawJSON() string { return r.JSON.raw }
func (r *AutomationListResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AutomationListResponseObject string

const (
	AutomationListResponseObjectList AutomationListResponseObject = "list"
)

type AutomationPauseResponse struct {
	// Unique identifier of the automation.
	ID string `json:"id" api:"required"`
	// When the automation was created.
	CreatedAt string `json:"createdAt" api:"required"`
	// What the automation is meant to accomplish.
	Description string `json:"description" api:"required"`
	// When its most recent failed run ended, test runs included; null if no run has
	// failed.
	LastFailureAt string `json:"lastFailureAt" api:"required"`
	// When its most recent run started, test runs included; null before the first run.
	LastRunAt string `json:"lastRunAt" api:"required"`
	// Model the agent steps run on.
	Model string `json:"model" api:"required"`
	// Display name of the automation.
	Name string `json:"name" api:"required"`
	// Always `automation`.
	//
	// Any of "automation".
	ObjectType AutomationPauseResponseObjectType `json:"objectType" api:"required"`
	// Who the automation acts as.
	RunAs AutomationPauseResponseRunAsUnion `json:"runAs" api:"required"`
	// `PAUSED` automations keep their triggers but start no runs.
	//
	// Any of "ACTIVE", "PAUSED".
	Status   AutomationPauseResponseStatus         `json:"status" api:"required"`
	Steps    []AutomationPauseResponseStepUnion    `json:"steps" api:"required"`
	Triggers []AutomationPauseResponseTriggerUnion `json:"triggers" api:"required"`
	// When the automation was last changed.
	UpdatedAt string `json:"updatedAt" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID            respjson.Field
		CreatedAt     respjson.Field
		Description   respjson.Field
		LastFailureAt respjson.Field
		LastRunAt     respjson.Field
		Model         respjson.Field
		Name          respjson.Field
		ObjectType    respjson.Field
		RunAs         respjson.Field
		Status        respjson.Field
		Steps         respjson.Field
		Triggers      respjson.Field
		UpdatedAt     respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutomationPauseResponse) RawJSON() string { return r.JSON.raw }
func (r *AutomationPauseResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Always `automation`.
type AutomationPauseResponseObjectType string

const (
	AutomationPauseResponseObjectTypeAutomation AutomationPauseResponseObjectType = "automation"
)

// AutomationPauseResponseRunAsUnion contains all possible properties and values
// from [AutomationPauseResponseRunAsWorkspaceRunAs],
// [AutomationPauseResponseRunAsMemberRunAs],
// [AutomationPauseResponseRunAsFromInputRunAs].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type AutomationPauseResponseRunAsUnion struct {
	Type string `json:"type"`
	// This field is from variant [AutomationPauseResponseRunAsMemberRunAs].
	MemberID string `json:"memberId"`
	JSON     struct {
		Type     respjson.Field
		MemberID respjson.Field
		raw      string
	} `json:"-"`
}

func (u AutomationPauseResponseRunAsUnion) AsWorkspaceRunAs() (v AutomationPauseResponseRunAsWorkspaceRunAs) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomationPauseResponseRunAsUnion) AsMemberRunAs() (v AutomationPauseResponseRunAsMemberRunAs) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomationPauseResponseRunAsUnion) AsFromInputRunAs() (v AutomationPauseResponseRunAsFromInputRunAs) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AutomationPauseResponseRunAsUnion) RawJSON() string { return u.JSON.raw }

func (r *AutomationPauseResponseRunAsUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Runs with no acting member.
type AutomationPauseResponseRunAsWorkspaceRunAs struct {
	// Any of "WORKSPACE".
	Type string `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutomationPauseResponseRunAsWorkspaceRunAs) RawJSON() string { return r.JSON.raw }
func (r *AutomationPauseResponseRunAsWorkspaceRunAs) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AutomationPauseResponseRunAsMemberRunAs struct {
	// Member the automation runs as.
	MemberID string `json:"memberId" api:"required"`
	// Any of "MEMBER".
	Type string `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		MemberID    respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutomationPauseResponseRunAsMemberRunAs) RawJSON() string { return r.JSON.raw }
func (r *AutomationPauseResponseRunAsMemberRunAs) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The acting member is read from an input when a run starts. Set in the Lightfield
// app.
type AutomationPauseResponseRunAsFromInputRunAs struct {
	// Any of "FROM_INPUT".
	Type string `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutomationPauseResponseRunAsFromInputRunAs) RawJSON() string { return r.JSON.raw }
func (r *AutomationPauseResponseRunAsFromInputRunAs) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// `PAUSED` automations keep their triggers but start no runs.
type AutomationPauseResponseStatus string

const (
	AutomationPauseResponseStatusActive AutomationPauseResponseStatus = "ACTIVE"
	AutomationPauseResponseStatusPaused AutomationPauseResponseStatus = "PAUSED"
)

// AutomationPauseResponseStepUnion contains all possible properties and values
// from [AutomationPauseResponseStepAgentStep],
// [AutomationPauseResponseStepCodeStep], [AutomationPauseResponseStepSleepStep].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type AutomationPauseResponseStepUnion struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	// This field is from variant [AutomationPauseResponseStepAgentStep].
	Prompt string `json:"prompt"`
	Type   string `json:"type"`
	// This field is a union of
	// [[]AutomationPauseResponseStepAgentStepPermissionUnion],
	// [[]AutomationPauseResponseStepCodeStepPermissionUnion]
	Permissions AutomationPauseResponseStepUnionPermissions `json:"permissions"`
	// This field is from variant [AutomationPauseResponseStepCodeStep].
	Command string `json:"command"`
	// This field is from variant [AutomationPauseResponseStepCodeStep].
	Description string `json:"description"`
	// This field is from variant [AutomationPauseResponseStepSleepStep].
	Duration AutomationPauseResponseStepSleepStepDurationUnion `json:"duration"`
	// This field is from variant [AutomationPauseResponseStepSleepStep].
	Unit string `json:"unit"`
	JSON struct {
		ID          respjson.Field
		Name        respjson.Field
		Prompt      respjson.Field
		Type        respjson.Field
		Permissions respjson.Field
		Command     respjson.Field
		Description respjson.Field
		Duration    respjson.Field
		Unit        respjson.Field
		raw         string
	} `json:"-"`
}

func (u AutomationPauseResponseStepUnion) AsAgentStep() (v AutomationPauseResponseStepAgentStep) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomationPauseResponseStepUnion) AsCodeStep() (v AutomationPauseResponseStepCodeStep) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomationPauseResponseStepUnion) AsSleepStep() (v AutomationPauseResponseStepSleepStep) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AutomationPauseResponseStepUnion) RawJSON() string { return u.JSON.raw }

func (r *AutomationPauseResponseStepUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AutomationPauseResponseStepUnionPermissions is an implicit subunion of
// [AutomationPauseResponseStepUnion]. AutomationPauseResponseStepUnionPermissions
// provides convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [AutomationPauseResponseStepUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfAutomationPauseResponseStepAgentStepPermissions
// OfAutomationPauseResponseStepCodeStepPermissions]
type AutomationPauseResponseStepUnionPermissions struct {
	// This field will be present if the value is a
	// [[]AutomationPauseResponseStepAgentStepPermissionUnion] instead of an object.
	OfAutomationPauseResponseStepAgentStepPermissions []AutomationPauseResponseStepAgentStepPermissionUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]AutomationPauseResponseStepCodeStepPermissionUnion] instead of an object.
	OfAutomationPauseResponseStepCodeStepPermissions []AutomationPauseResponseStepCodeStepPermissionUnion `json:",inline"`
	JSON                                             struct {
		OfAutomationPauseResponseStepAgentStepPermissions respjson.Field
		OfAutomationPauseResponseStepCodeStepPermissions  respjson.Field
		raw                                               string
	} `json:"-"`
}

func (r *AutomationPauseResponseStepUnionPermissions) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AutomationPauseResponseStepAgentStep struct {
	// Step identifier. Assigned by Lightfield.
	ID string `json:"id" api:"required"`
	// Display name of the step.
	Name string `json:"name" api:"required"`
	// Instructions the agent follows in this step.
	Prompt string `json:"prompt" api:"required"`
	// Any of "AGENT".
	Type string `json:"type" api:"required"`
	// What the agent may use in this step. Defaults to none.
	Permissions []AutomationPauseResponseStepAgentStepPermissionUnion `json:"permissions"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Name        respjson.Field
		Prompt      respjson.Field
		Type        respjson.Field
		Permissions respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutomationPauseResponseStepAgentStep) RawJSON() string { return r.JSON.raw }
func (r *AutomationPauseResponseStepAgentStep) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AutomationPauseResponseStepAgentStepPermissionUnion contains all possible
// properties and values from
// [AutomationPauseResponseStepAgentStepPermissionToolPermission],
// [AutomationPauseResponseStepAgentStepPermissionMcpPermission],
// [AutomationPauseResponseStepAgentStepPermissionAPIPermission],
// [AutomationPauseResponseStepAgentStepPermissionNetworkPermission].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type AutomationPauseResponseStepAgentStepPermissionUnion struct {
	// This field is from variant
	// [AutomationPauseResponseStepAgentStepPermissionToolPermission].
	Name   string `json:"name"`
	Type   string `json:"type"`
	Reason string `json:"reason"`
	// This field is from variant
	// [AutomationPauseResponseStepAgentStepPermissionMcpPermission].
	Connector string `json:"connector"`
	// This field is from variant
	// [AutomationPauseResponseStepAgentStepPermissionMcpPermission].
	Tools []string `json:"tools"`
	// This field is from variant
	// [AutomationPauseResponseStepAgentStepPermissionAPIPermission].
	Scopes []string `json:"scopes"`
	// This field is from variant
	// [AutomationPauseResponseStepAgentStepPermissionNetworkPermission].
	Domain string `json:"domain"`
	// This field is from variant
	// [AutomationPauseResponseStepAgentStepPermissionNetworkPermission].
	Headers map[string]string `json:"headers"`
	// This field is from variant
	// [AutomationPauseResponseStepAgentStepPermissionNetworkPermission].
	Method string `json:"method"`
	JSON   struct {
		Name      respjson.Field
		Type      respjson.Field
		Reason    respjson.Field
		Connector respjson.Field
		Tools     respjson.Field
		Scopes    respjson.Field
		Domain    respjson.Field
		Headers   respjson.Field
		Method    respjson.Field
		raw       string
	} `json:"-"`
}

func (u AutomationPauseResponseStepAgentStepPermissionUnion) AsToolPermission() (v AutomationPauseResponseStepAgentStepPermissionToolPermission) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomationPauseResponseStepAgentStepPermissionUnion) AsMcpPermission() (v AutomationPauseResponseStepAgentStepPermissionMcpPermission) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomationPauseResponseStepAgentStepPermissionUnion) AsAPIPermission() (v AutomationPauseResponseStepAgentStepPermissionAPIPermission) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomationPauseResponseStepAgentStepPermissionUnion) AsNetworkPermission() (v AutomationPauseResponseStepAgentStepPermissionNetworkPermission) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AutomationPauseResponseStepAgentStepPermissionUnion) RawJSON() string { return u.JSON.raw }

func (r *AutomationPauseResponseStepAgentStepPermissionUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AutomationPauseResponseStepAgentStepPermissionToolPermission struct {
	// Name of the tool, matching the Lightfield app.
	Name string `json:"name" api:"required"`
	// Any of "TOOL".
	Type string `json:"type" api:"required"`
	// Why the step needs this permission. Null when none was given. Returns 422 for an
	// empty or whitespace-only string.
	Reason string `json:"reason" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name        respjson.Field
		Type        respjson.Field
		Reason      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutomationPauseResponseStepAgentStepPermissionToolPermission) RawJSON() string {
	return r.JSON.raw
}
func (r *AutomationPauseResponseStepAgentStepPermissionToolPermission) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AutomationPauseResponseStepAgentStepPermissionMcpPermission struct {
	// Connector the tools are on. Must already be lowercase and trimmed.
	Connector string `json:"connector" api:"required"`
	// Unique, trimmed names of the connector tools the step may call.
	Tools []string `json:"tools" api:"required"`
	// Any of "MCP".
	Type string `json:"type" api:"required"`
	// Why the step needs this permission. Null when none was given. Returns 422 for an
	// empty or whitespace-only string.
	Reason string `json:"reason" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Connector   respjson.Field
		Tools       respjson.Field
		Type        respjson.Field
		Reason      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutomationPauseResponseStepAgentStepPermissionMcpPermission) RawJSON() string {
	return r.JSON.raw
}
func (r *AutomationPauseResponseStepAgentStepPermissionMcpPermission) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AutomationPauseResponseStepAgentStepPermissionAPIPermission struct {
	// Lightfield API scopes the step may use, for example `contacts:read`. The API key
	// must hold every scope it grants, and only a full-access key may grant `*`.
	Scopes []string `json:"scopes" api:"required"`
	// Any of "API".
	Type string `json:"type" api:"required"`
	// Why the step needs this permission. Null when none was given. Returns 422 for an
	// empty or whitespace-only string.
	Reason string `json:"reason" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Scopes      respjson.Field
		Type        respjson.Field
		Reason      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutomationPauseResponseStepAgentStepPermissionAPIPermission) RawJSON() string {
	return r.JSON.raw
}
func (r *AutomationPauseResponseStepAgentStepPermissionAPIPermission) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AutomationPauseResponseStepAgentStepPermissionNetworkPermission struct {
	// Lowercase host the step may call, optionally followed by a path glob. Returns
	// 422 for a scheme, query, fragment or surrounding whitespace.
	Domain string `json:"domain" api:"required"`
	// Any of "NETWORK".
	Type string `json:"type" api:"required"`
	// Headers added to every request this permission allows.
	Headers map[string]string `json:"headers"`
	// Restricts the permission to one HTTP method.
	//
	// Any of "GET", "POST", "PUT", "PATCH", "DELETE".
	Method string `json:"method"`
	// Why the step needs this permission. Null when none was given. Returns 422 for an
	// empty or whitespace-only string.
	Reason string `json:"reason" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Domain      respjson.Field
		Type        respjson.Field
		Headers     respjson.Field
		Method      respjson.Field
		Reason      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutomationPauseResponseStepAgentStepPermissionNetworkPermission) RawJSON() string {
	return r.JSON.raw
}
func (r *AutomationPauseResponseStepAgentStepPermissionNetworkPermission) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AutomationPauseResponseStepCodeStep struct {
	// Step identifier. Assigned by Lightfield.
	ID string `json:"id" api:"required"`
	// Bash command the step runs as written.
	Command string `json:"command" api:"required"`
	// Display name of the step.
	Name string `json:"name" api:"required"`
	// Any of "CODE".
	Type string `json:"type" api:"required"`
	// What the command does. Does not change what runs.
	Description string `json:"description"`
	// What the command may use. Defaults to none.
	Permissions []AutomationPauseResponseStepCodeStepPermissionUnion `json:"permissions"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Command     respjson.Field
		Name        respjson.Field
		Type        respjson.Field
		Description respjson.Field
		Permissions respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutomationPauseResponseStepCodeStep) RawJSON() string { return r.JSON.raw }
func (r *AutomationPauseResponseStepCodeStep) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AutomationPauseResponseStepCodeStepPermissionUnion contains all possible
// properties and values from
// [AutomationPauseResponseStepCodeStepPermissionAPIPermission],
// [AutomationPauseResponseStepCodeStepPermissionNetworkPermission].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type AutomationPauseResponseStepCodeStepPermissionUnion struct {
	// This field is from variant
	// [AutomationPauseResponseStepCodeStepPermissionAPIPermission].
	Scopes []string `json:"scopes"`
	Type   string   `json:"type"`
	Reason string   `json:"reason"`
	// This field is from variant
	// [AutomationPauseResponseStepCodeStepPermissionNetworkPermission].
	Domain string `json:"domain"`
	// This field is from variant
	// [AutomationPauseResponseStepCodeStepPermissionNetworkPermission].
	Headers map[string]string `json:"headers"`
	// This field is from variant
	// [AutomationPauseResponseStepCodeStepPermissionNetworkPermission].
	Method string `json:"method"`
	JSON   struct {
		Scopes  respjson.Field
		Type    respjson.Field
		Reason  respjson.Field
		Domain  respjson.Field
		Headers respjson.Field
		Method  respjson.Field
		raw     string
	} `json:"-"`
}

func (u AutomationPauseResponseStepCodeStepPermissionUnion) AsAPIPermission() (v AutomationPauseResponseStepCodeStepPermissionAPIPermission) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomationPauseResponseStepCodeStepPermissionUnion) AsNetworkPermission() (v AutomationPauseResponseStepCodeStepPermissionNetworkPermission) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AutomationPauseResponseStepCodeStepPermissionUnion) RawJSON() string { return u.JSON.raw }

func (r *AutomationPauseResponseStepCodeStepPermissionUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AutomationPauseResponseStepCodeStepPermissionAPIPermission struct {
	// Lightfield API scopes the step may use, for example `contacts:read`. The API key
	// must hold every scope it grants, and only a full-access key may grant `*`.
	Scopes []string `json:"scopes" api:"required"`
	// Any of "API".
	Type string `json:"type" api:"required"`
	// Why the step needs this permission. Null when none was given. Returns 422 for an
	// empty or whitespace-only string.
	Reason string `json:"reason" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Scopes      respjson.Field
		Type        respjson.Field
		Reason      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutomationPauseResponseStepCodeStepPermissionAPIPermission) RawJSON() string {
	return r.JSON.raw
}
func (r *AutomationPauseResponseStepCodeStepPermissionAPIPermission) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AutomationPauseResponseStepCodeStepPermissionNetworkPermission struct {
	// Lowercase host the step may call, optionally followed by a path glob. Returns
	// 422 for a scheme, query, fragment or surrounding whitespace.
	Domain string `json:"domain" api:"required"`
	// Any of "NETWORK".
	Type string `json:"type" api:"required"`
	// Headers added to every request this permission allows.
	Headers map[string]string `json:"headers"`
	// Restricts the permission to one HTTP method.
	//
	// Any of "GET", "POST", "PUT", "PATCH", "DELETE".
	Method string `json:"method"`
	// Why the step needs this permission. Null when none was given. Returns 422 for an
	// empty or whitespace-only string.
	Reason string `json:"reason" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Domain      respjson.Field
		Type        respjson.Field
		Headers     respjson.Field
		Method      respjson.Field
		Reason      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutomationPauseResponseStepCodeStepPermissionNetworkPermission) RawJSON() string {
	return r.JSON.raw
}
func (r *AutomationPauseResponseStepCodeStepPermissionNetworkPermission) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AutomationPauseResponseStepSleepStep struct {
	// Step identifier. Assigned by Lightfield.
	ID string `json:"id" api:"required"`
	// How long to sleep before the next step, at most one year.
	Duration AutomationPauseResponseStepSleepStepDurationUnion `json:"duration" api:"required"`
	// Display name of the step.
	Name string `json:"name" api:"required"`
	// Any of "SLEEP".
	Type string `json:"type" api:"required"`
	// Unit of `duration`.
	//
	// Any of "MILLISECOND", "SECOND", "MINUTE", "HOUR", "DAY".
	Unit string `json:"unit" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Duration    respjson.Field
		Name        respjson.Field
		Type        respjson.Field
		Unit        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutomationPauseResponseStepSleepStep) RawJSON() string { return r.JSON.raw }
func (r *AutomationPauseResponseStepSleepStep) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AutomationPauseResponseStepSleepStepDurationUnion contains all possible
// properties and values from [int64], [string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfInt OfString]
type AutomationPauseResponseStepSleepStepDurationUnion struct {
	// This field will be present if the value is a [int64] instead of an object.
	OfInt int64 `json:",inline"`
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	JSON     struct {
		OfInt    respjson.Field
		OfString respjson.Field
		raw      string
	} `json:"-"`
}

func (u AutomationPauseResponseStepSleepStepDurationUnion) AsInt() (v int64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomationPauseResponseStepSleepStepDurationUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AutomationPauseResponseStepSleepStepDurationUnion) RawJSON() string { return u.JSON.raw }

func (r *AutomationPauseResponseStepSleepStepDurationUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AutomationPauseResponseTriggerUnion contains all possible properties and values
// from [AutomationPauseResponseTriggerManualTrigger],
// [AutomationPauseResponseTriggerWebhookTrigger],
// [AutomationPauseResponseTriggerScheduledTrigger],
// [AutomationPauseResponseTriggerRecordCreatedTrigger],
// [AutomationPauseResponseTriggerRecordUpdatedTrigger],
// [AutomationPauseResponseTriggerSequenceReplyTrigger].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type AutomationPauseResponseTriggerUnion struct {
	Type string `json:"type"`
	// This field is from variant [AutomationPauseResponseTriggerWebhookTrigger].
	URL string `json:"url"`
	// This field is from variant [AutomationPauseResponseTriggerWebhookTrigger].
	ExampleResult AutomationPauseResponseTriggerWebhookTriggerExampleResult `json:"exampleResult"`
	// This field is from variant [AutomationPauseResponseTriggerScheduledTrigger].
	Schedule   AutomationPauseResponseTriggerScheduledTriggerSchedule `json:"schedule"`
	ObjectType string                                                 `json:"objectType"`
	// This field is from variant [AutomationPauseResponseTriggerRecordUpdatedTrigger].
	WatchedFields []string `json:"watchedFields"`
	// This field is from variant [AutomationPauseResponseTriggerRecordUpdatedTrigger].
	WatchedRelationships []AutomationPauseResponseTriggerRecordUpdatedTriggerWatchedRelationship `json:"watchedRelationships"`
	JSON                 struct {
		Type                 respjson.Field
		URL                  respjson.Field
		ExampleResult        respjson.Field
		Schedule             respjson.Field
		ObjectType           respjson.Field
		WatchedFields        respjson.Field
		WatchedRelationships respjson.Field
		raw                  string
	} `json:"-"`
}

func (u AutomationPauseResponseTriggerUnion) AsManualTrigger() (v AutomationPauseResponseTriggerManualTrigger) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomationPauseResponseTriggerUnion) AsWebhookTrigger() (v AutomationPauseResponseTriggerWebhookTrigger) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomationPauseResponseTriggerUnion) AsScheduledTrigger() (v AutomationPauseResponseTriggerScheduledTrigger) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomationPauseResponseTriggerUnion) AsRecordCreatedTrigger() (v AutomationPauseResponseTriggerRecordCreatedTrigger) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomationPauseResponseTriggerUnion) AsRecordUpdatedTrigger() (v AutomationPauseResponseTriggerRecordUpdatedTrigger) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomationPauseResponseTriggerUnion) AsSequenceReplyTrigger() (v AutomationPauseResponseTriggerSequenceReplyTrigger) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AutomationPauseResponseTriggerUnion) RawJSON() string { return u.JSON.raw }

func (r *AutomationPauseResponseTriggerUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AutomationPauseResponseTriggerManualTrigger struct {
	// Any of "MANUAL".
	Type string `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutomationPauseResponseTriggerManualTrigger) RawJSON() string { return r.JSON.raw }
func (r *AutomationPauseResponseTriggerManualTrigger) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AutomationPauseResponseTriggerWebhookTrigger struct {
	// Any of "WEBHOOK".
	Type string `json:"type" api:"required"`
	// URL that starts a run when it receives a POST request.
	URL string `json:"url" api:"required"`
	// A sample payload and the fields the steps read from it.
	ExampleResult AutomationPauseResponseTriggerWebhookTriggerExampleResult `json:"exampleResult"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type          respjson.Field
		URL           respjson.Field
		ExampleResult respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutomationPauseResponseTriggerWebhookTrigger) RawJSON() string { return r.JSON.raw }
func (r *AutomationPauseResponseTriggerWebhookTrigger) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A sample payload and the fields the steps read from it.
type AutomationPauseResponseTriggerWebhookTriggerExampleResult struct {
	Json   AutomationPauseResponseTriggerWebhookTriggerExampleResultJsonUnion `json:"json" api:"required"`
	Parsed []AutomationPauseResponseTriggerWebhookTriggerExampleResultParsed  `json:"parsed" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Json        respjson.Field
		Parsed      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutomationPauseResponseTriggerWebhookTriggerExampleResult) RawJSON() string {
	return r.JSON.raw
}
func (r *AutomationPauseResponseTriggerWebhookTriggerExampleResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AutomationPauseResponseTriggerWebhookTriggerExampleResultJsonUnion contains all
// possible properties and values from
// [map[string]AutomationPauseResponseTriggerWebhookTriggerExampleResultJsonMapItemUnion],
// [[]AutomationPauseResponseTriggerWebhookTriggerExampleResultJsonArrayItemUnion].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool
// OfAutomationPauseResponseTriggerWebhookTriggerExampleResultJsonMapItemArray
// OfJsonValueArray OfAnyArray
// OfAutomationPauseResponseTriggerWebhookTriggerExampleResultJsonArray]
type AutomationPauseResponseTriggerWebhookTriggerExampleResultJsonUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a
	// [[]AutomationPauseResponseTriggerWebhookTriggerExampleResultJsonMapItemArrayItemUnion]
	// instead of an object.
	OfAutomationPauseResponseTriggerWebhookTriggerExampleResultJsonMapItemArray []AutomationPauseResponseTriggerWebhookTriggerExampleResultJsonMapItemArrayItemUnion `json:",inline"`
	// This field will be present if the value is a [[]shared.JsonValueUnion] instead
	// of an object.
	OfJsonValueArray []shared.JsonValueUnion `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	// This field will be present if the value is a
	// [[]AutomationPauseResponseTriggerWebhookTriggerExampleResultJsonArrayItemUnion]
	// instead of an object.
	OfAutomationPauseResponseTriggerWebhookTriggerExampleResultJsonArray []AutomationPauseResponseTriggerWebhookTriggerExampleResultJsonArrayItemUnion `json:",inline"`
	JSON                                                                 struct {
		OfString                                                                    respjson.Field
		OfFloat                                                                     respjson.Field
		OfBool                                                                      respjson.Field
		OfAutomationPauseResponseTriggerWebhookTriggerExampleResultJsonMapItemArray respjson.Field
		OfJsonValueArray                                                            respjson.Field
		OfAnyArray                                                                  respjson.Field
		OfAutomationPauseResponseTriggerWebhookTriggerExampleResultJsonArray        respjson.Field
		raw                                                                         string
	} `json:"-"`
}

func (u AutomationPauseResponseTriggerWebhookTriggerExampleResultJsonUnion) AsAutomationPauseResponseTriggerWebhookTriggerExampleResultJsonMapMap() (v map[string]AutomationPauseResponseTriggerWebhookTriggerExampleResultJsonMapItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomationPauseResponseTriggerWebhookTriggerExampleResultJsonUnion) AsAutomationPauseResponseTriggerWebhookTriggerExampleResultJsonArray() (v []AutomationPauseResponseTriggerWebhookTriggerExampleResultJsonArrayItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AutomationPauseResponseTriggerWebhookTriggerExampleResultJsonUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *AutomationPauseResponseTriggerWebhookTriggerExampleResultJsonUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AutomationPauseResponseTriggerWebhookTriggerExampleResultJsonMapItemUnion
// contains all possible properties and values from [string], [float64], [bool],
// [[]AutomationPauseResponseTriggerWebhookTriggerExampleResultJsonMapItemArrayItemUnion],
// [map[string]AutomationPauseResponseTriggerWebhookTriggerExampleResultJsonMapItemMapItemUnion].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool
// OfAutomationPauseResponseTriggerWebhookTriggerExampleResultJsonMapItemArray
// OfJsonValueArray OfAnyArray]
type AutomationPauseResponseTriggerWebhookTriggerExampleResultJsonMapItemUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a
	// [[]AutomationPauseResponseTriggerWebhookTriggerExampleResultJsonMapItemArrayItemUnion]
	// instead of an object.
	OfAutomationPauseResponseTriggerWebhookTriggerExampleResultJsonMapItemArray []AutomationPauseResponseTriggerWebhookTriggerExampleResultJsonMapItemArrayItemUnion `json:",inline"`
	// This field will be present if the value is a [[]shared.JsonValueUnion] instead
	// of an object.
	OfJsonValueArray []shared.JsonValueUnion `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	JSON       struct {
		OfString                                                                    respjson.Field
		OfFloat                                                                     respjson.Field
		OfBool                                                                      respjson.Field
		OfAutomationPauseResponseTriggerWebhookTriggerExampleResultJsonMapItemArray respjson.Field
		OfJsonValueArray                                                            respjson.Field
		OfAnyArray                                                                  respjson.Field
		raw                                                                         string
	} `json:"-"`
}

func (u AutomationPauseResponseTriggerWebhookTriggerExampleResultJsonMapItemUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomationPauseResponseTriggerWebhookTriggerExampleResultJsonMapItemUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomationPauseResponseTriggerWebhookTriggerExampleResultJsonMapItemUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomationPauseResponseTriggerWebhookTriggerExampleResultJsonMapItemUnion) AsAutomationPauseResponseTriggerWebhookTriggerExampleResultJsonMapItemArray() (v []AutomationPauseResponseTriggerWebhookTriggerExampleResultJsonMapItemArrayItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomationPauseResponseTriggerWebhookTriggerExampleResultJsonMapItemUnion) AsAutomationPauseResponseTriggerWebhookTriggerExampleResultJsonMapItemMapMap() (v map[string]AutomationPauseResponseTriggerWebhookTriggerExampleResultJsonMapItemMapItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AutomationPauseResponseTriggerWebhookTriggerExampleResultJsonMapItemUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *AutomationPauseResponseTriggerWebhookTriggerExampleResultJsonMapItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AutomationPauseResponseTriggerWebhookTriggerExampleResultJsonMapItemArrayItemUnion
// contains all possible properties and values from [string], [float64], [bool],
// [[]shared.JsonValueUnion], [map[string]shared.JsonValueUnion].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfJsonValueArray OfAnyArray]
type AutomationPauseResponseTriggerWebhookTriggerExampleResultJsonMapItemArrayItemUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]shared.JsonValueUnion] instead
	// of an object.
	OfJsonValueArray []shared.JsonValueUnion `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	JSON       struct {
		OfString         respjson.Field
		OfFloat          respjson.Field
		OfBool           respjson.Field
		OfJsonValueArray respjson.Field
		OfAnyArray       respjson.Field
		raw              string
	} `json:"-"`
}

func (u AutomationPauseResponseTriggerWebhookTriggerExampleResultJsonMapItemArrayItemUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomationPauseResponseTriggerWebhookTriggerExampleResultJsonMapItemArrayItemUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomationPauseResponseTriggerWebhookTriggerExampleResultJsonMapItemArrayItemUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomationPauseResponseTriggerWebhookTriggerExampleResultJsonMapItemArrayItemUnion) AsJsonValueArray() (v []shared.JsonValueUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomationPauseResponseTriggerWebhookTriggerExampleResultJsonMapItemArrayItemUnion) AsJsonValueMap() (v map[string]shared.JsonValueUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AutomationPauseResponseTriggerWebhookTriggerExampleResultJsonMapItemArrayItemUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *AutomationPauseResponseTriggerWebhookTriggerExampleResultJsonMapItemArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AutomationPauseResponseTriggerWebhookTriggerExampleResultJsonMapItemMapItemUnion
// contains all possible properties and values from [string], [float64], [bool],
// [[]shared.JsonValueUnion], [map[string]shared.JsonValueUnion].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfJsonValueArray OfAnyArray]
type AutomationPauseResponseTriggerWebhookTriggerExampleResultJsonMapItemMapItemUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]shared.JsonValueUnion] instead
	// of an object.
	OfJsonValueArray []shared.JsonValueUnion `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	JSON       struct {
		OfString         respjson.Field
		OfFloat          respjson.Field
		OfBool           respjson.Field
		OfJsonValueArray respjson.Field
		OfAnyArray       respjson.Field
		raw              string
	} `json:"-"`
}

func (u AutomationPauseResponseTriggerWebhookTriggerExampleResultJsonMapItemMapItemUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomationPauseResponseTriggerWebhookTriggerExampleResultJsonMapItemMapItemUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomationPauseResponseTriggerWebhookTriggerExampleResultJsonMapItemMapItemUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomationPauseResponseTriggerWebhookTriggerExampleResultJsonMapItemMapItemUnion) AsJsonValueArray() (v []shared.JsonValueUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomationPauseResponseTriggerWebhookTriggerExampleResultJsonMapItemMapItemUnion) AsJsonValueMap() (v map[string]shared.JsonValueUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AutomationPauseResponseTriggerWebhookTriggerExampleResultJsonMapItemMapItemUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *AutomationPauseResponseTriggerWebhookTriggerExampleResultJsonMapItemMapItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AutomationPauseResponseTriggerWebhookTriggerExampleResultJsonArrayItemUnion
// contains all possible properties and values from [string], [float64], [bool],
// [[]AutomationPauseResponseTriggerWebhookTriggerExampleResultJsonArrayItemArrayItemUnion],
// [map[string]AutomationPauseResponseTriggerWebhookTriggerExampleResultJsonArrayItemMapItemUnion].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool
// OfAutomationPauseResponseTriggerWebhookTriggerExampleResultJsonArrayItemArray
// OfJsonValueArray OfAnyArray]
type AutomationPauseResponseTriggerWebhookTriggerExampleResultJsonArrayItemUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a
	// [[]AutomationPauseResponseTriggerWebhookTriggerExampleResultJsonArrayItemArrayItemUnion]
	// instead of an object.
	OfAutomationPauseResponseTriggerWebhookTriggerExampleResultJsonArrayItemArray []AutomationPauseResponseTriggerWebhookTriggerExampleResultJsonArrayItemArrayItemUnion `json:",inline"`
	// This field will be present if the value is a [[]shared.JsonValueUnion] instead
	// of an object.
	OfJsonValueArray []shared.JsonValueUnion `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	JSON       struct {
		OfString                                                                      respjson.Field
		OfFloat                                                                       respjson.Field
		OfBool                                                                        respjson.Field
		OfAutomationPauseResponseTriggerWebhookTriggerExampleResultJsonArrayItemArray respjson.Field
		OfJsonValueArray                                                              respjson.Field
		OfAnyArray                                                                    respjson.Field
		raw                                                                           string
	} `json:"-"`
}

func (u AutomationPauseResponseTriggerWebhookTriggerExampleResultJsonArrayItemUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomationPauseResponseTriggerWebhookTriggerExampleResultJsonArrayItemUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomationPauseResponseTriggerWebhookTriggerExampleResultJsonArrayItemUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomationPauseResponseTriggerWebhookTriggerExampleResultJsonArrayItemUnion) AsAutomationPauseResponseTriggerWebhookTriggerExampleResultJsonArrayItemArray() (v []AutomationPauseResponseTriggerWebhookTriggerExampleResultJsonArrayItemArrayItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomationPauseResponseTriggerWebhookTriggerExampleResultJsonArrayItemUnion) AsAutomationPauseResponseTriggerWebhookTriggerExampleResultJsonArrayItemMapMap() (v map[string]AutomationPauseResponseTriggerWebhookTriggerExampleResultJsonArrayItemMapItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AutomationPauseResponseTriggerWebhookTriggerExampleResultJsonArrayItemUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *AutomationPauseResponseTriggerWebhookTriggerExampleResultJsonArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AutomationPauseResponseTriggerWebhookTriggerExampleResultJsonArrayItemArrayItemUnion
// contains all possible properties and values from [string], [float64], [bool],
// [[]shared.JsonValueUnion], [map[string]shared.JsonValueUnion].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfJsonValueArray OfAnyArray]
type AutomationPauseResponseTriggerWebhookTriggerExampleResultJsonArrayItemArrayItemUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]shared.JsonValueUnion] instead
	// of an object.
	OfJsonValueArray []shared.JsonValueUnion `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	JSON       struct {
		OfString         respjson.Field
		OfFloat          respjson.Field
		OfBool           respjson.Field
		OfJsonValueArray respjson.Field
		OfAnyArray       respjson.Field
		raw              string
	} `json:"-"`
}

func (u AutomationPauseResponseTriggerWebhookTriggerExampleResultJsonArrayItemArrayItemUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomationPauseResponseTriggerWebhookTriggerExampleResultJsonArrayItemArrayItemUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomationPauseResponseTriggerWebhookTriggerExampleResultJsonArrayItemArrayItemUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomationPauseResponseTriggerWebhookTriggerExampleResultJsonArrayItemArrayItemUnion) AsJsonValueArray() (v []shared.JsonValueUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomationPauseResponseTriggerWebhookTriggerExampleResultJsonArrayItemArrayItemUnion) AsJsonValueMap() (v map[string]shared.JsonValueUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AutomationPauseResponseTriggerWebhookTriggerExampleResultJsonArrayItemArrayItemUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *AutomationPauseResponseTriggerWebhookTriggerExampleResultJsonArrayItemArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AutomationPauseResponseTriggerWebhookTriggerExampleResultJsonArrayItemMapItemUnion
// contains all possible properties and values from [string], [float64], [bool],
// [[]shared.JsonValueUnion], [map[string]shared.JsonValueUnion].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfJsonValueArray OfAnyArray]
type AutomationPauseResponseTriggerWebhookTriggerExampleResultJsonArrayItemMapItemUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]shared.JsonValueUnion] instead
	// of an object.
	OfJsonValueArray []shared.JsonValueUnion `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	JSON       struct {
		OfString         respjson.Field
		OfFloat          respjson.Field
		OfBool           respjson.Field
		OfJsonValueArray respjson.Field
		OfAnyArray       respjson.Field
		raw              string
	} `json:"-"`
}

func (u AutomationPauseResponseTriggerWebhookTriggerExampleResultJsonArrayItemMapItemUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomationPauseResponseTriggerWebhookTriggerExampleResultJsonArrayItemMapItemUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomationPauseResponseTriggerWebhookTriggerExampleResultJsonArrayItemMapItemUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomationPauseResponseTriggerWebhookTriggerExampleResultJsonArrayItemMapItemUnion) AsJsonValueArray() (v []shared.JsonValueUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomationPauseResponseTriggerWebhookTriggerExampleResultJsonArrayItemMapItemUnion) AsJsonValueMap() (v map[string]shared.JsonValueUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AutomationPauseResponseTriggerWebhookTriggerExampleResultJsonArrayItemMapItemUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *AutomationPauseResponseTriggerWebhookTriggerExampleResultJsonArrayItemMapItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AutomationPauseResponseTriggerWebhookTriggerExampleResultParsed struct {
	Path  []string                                                                  `json:"path" api:"required"`
	Label string                                                                    `json:"label"`
	Value AutomationPauseResponseTriggerWebhookTriggerExampleResultParsedValueUnion `json:"value" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Path        respjson.Field
		Label       respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutomationPauseResponseTriggerWebhookTriggerExampleResultParsed) RawJSON() string {
	return r.JSON.raw
}
func (r *AutomationPauseResponseTriggerWebhookTriggerExampleResultParsed) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AutomationPauseResponseTriggerWebhookTriggerExampleResultParsedValueUnion
// contains all possible properties and values from [string], [float64], [bool],
// [[]AutomationPauseResponseTriggerWebhookTriggerExampleResultParsedValueArrayItemUnion],
// [map[string]AutomationPauseResponseTriggerWebhookTriggerExampleResultParsedValueMapItemUnion].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool
// OfAutomationPauseResponseTriggerWebhookTriggerExampleResultParsedValueArray
// OfJsonValueArray OfAnyArray]
type AutomationPauseResponseTriggerWebhookTriggerExampleResultParsedValueUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a
	// [[]AutomationPauseResponseTriggerWebhookTriggerExampleResultParsedValueArrayItemUnion]
	// instead of an object.
	OfAutomationPauseResponseTriggerWebhookTriggerExampleResultParsedValueArray []AutomationPauseResponseTriggerWebhookTriggerExampleResultParsedValueArrayItemUnion `json:",inline"`
	// This field will be present if the value is a [[]shared.JsonValueUnion] instead
	// of an object.
	OfJsonValueArray []shared.JsonValueUnion `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	JSON       struct {
		OfString                                                                    respjson.Field
		OfFloat                                                                     respjson.Field
		OfBool                                                                      respjson.Field
		OfAutomationPauseResponseTriggerWebhookTriggerExampleResultParsedValueArray respjson.Field
		OfJsonValueArray                                                            respjson.Field
		OfAnyArray                                                                  respjson.Field
		raw                                                                         string
	} `json:"-"`
}

func (u AutomationPauseResponseTriggerWebhookTriggerExampleResultParsedValueUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomationPauseResponseTriggerWebhookTriggerExampleResultParsedValueUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomationPauseResponseTriggerWebhookTriggerExampleResultParsedValueUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomationPauseResponseTriggerWebhookTriggerExampleResultParsedValueUnion) AsAutomationPauseResponseTriggerWebhookTriggerExampleResultParsedValueArray() (v []AutomationPauseResponseTriggerWebhookTriggerExampleResultParsedValueArrayItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomationPauseResponseTriggerWebhookTriggerExampleResultParsedValueUnion) AsAutomationPauseResponseTriggerWebhookTriggerExampleResultParsedValueMapMap() (v map[string]AutomationPauseResponseTriggerWebhookTriggerExampleResultParsedValueMapItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AutomationPauseResponseTriggerWebhookTriggerExampleResultParsedValueUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *AutomationPauseResponseTriggerWebhookTriggerExampleResultParsedValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AutomationPauseResponseTriggerWebhookTriggerExampleResultParsedValueArrayItemUnion
// contains all possible properties and values from [string], [float64], [bool],
// [[]shared.JsonValueUnion], [map[string]shared.JsonValueUnion].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfJsonValueArray OfAnyArray]
type AutomationPauseResponseTriggerWebhookTriggerExampleResultParsedValueArrayItemUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]shared.JsonValueUnion] instead
	// of an object.
	OfJsonValueArray []shared.JsonValueUnion `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	JSON       struct {
		OfString         respjson.Field
		OfFloat          respjson.Field
		OfBool           respjson.Field
		OfJsonValueArray respjson.Field
		OfAnyArray       respjson.Field
		raw              string
	} `json:"-"`
}

func (u AutomationPauseResponseTriggerWebhookTriggerExampleResultParsedValueArrayItemUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomationPauseResponseTriggerWebhookTriggerExampleResultParsedValueArrayItemUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomationPauseResponseTriggerWebhookTriggerExampleResultParsedValueArrayItemUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomationPauseResponseTriggerWebhookTriggerExampleResultParsedValueArrayItemUnion) AsJsonValueArray() (v []shared.JsonValueUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomationPauseResponseTriggerWebhookTriggerExampleResultParsedValueArrayItemUnion) AsJsonValueMap() (v map[string]shared.JsonValueUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AutomationPauseResponseTriggerWebhookTriggerExampleResultParsedValueArrayItemUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *AutomationPauseResponseTriggerWebhookTriggerExampleResultParsedValueArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AutomationPauseResponseTriggerWebhookTriggerExampleResultParsedValueMapItemUnion
// contains all possible properties and values from [string], [float64], [bool],
// [[]shared.JsonValueUnion], [map[string]shared.JsonValueUnion].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfJsonValueArray OfAnyArray]
type AutomationPauseResponseTriggerWebhookTriggerExampleResultParsedValueMapItemUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]shared.JsonValueUnion] instead
	// of an object.
	OfJsonValueArray []shared.JsonValueUnion `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	JSON       struct {
		OfString         respjson.Field
		OfFloat          respjson.Field
		OfBool           respjson.Field
		OfJsonValueArray respjson.Field
		OfAnyArray       respjson.Field
		raw              string
	} `json:"-"`
}

func (u AutomationPauseResponseTriggerWebhookTriggerExampleResultParsedValueMapItemUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomationPauseResponseTriggerWebhookTriggerExampleResultParsedValueMapItemUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomationPauseResponseTriggerWebhookTriggerExampleResultParsedValueMapItemUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomationPauseResponseTriggerWebhookTriggerExampleResultParsedValueMapItemUnion) AsJsonValueArray() (v []shared.JsonValueUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomationPauseResponseTriggerWebhookTriggerExampleResultParsedValueMapItemUnion) AsJsonValueMap() (v map[string]shared.JsonValueUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AutomationPauseResponseTriggerWebhookTriggerExampleResultParsedValueMapItemUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *AutomationPauseResponseTriggerWebhookTriggerExampleResultParsedValueMapItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AutomationPauseResponseTriggerScheduledTrigger struct {
	Schedule AutomationPauseResponseTriggerScheduledTriggerSchedule `json:"schedule" api:"required"`
	// Any of "SCHEDULED".
	Type string `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Schedule    respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutomationPauseResponseTriggerScheduledTrigger) RawJSON() string { return r.JSON.raw }
func (r *AutomationPauseResponseTriggerScheduledTrigger) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AutomationPauseResponseTriggerScheduledTriggerSchedule struct {
	// 5-field cron expression, for example `0 9 * * 1-5`.
	Cron string `json:"cron" api:"required"`
	// IANA timezone the cron is evaluated in.
	Timezone string `json:"timezone" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Cron        respjson.Field
		Timezone    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutomationPauseResponseTriggerScheduledTriggerSchedule) RawJSON() string { return r.JSON.raw }
func (r *AutomationPauseResponseTriggerScheduledTriggerSchedule) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AutomationPauseResponseTriggerRecordCreatedTrigger struct {
	ObjectType string `json:"objectType" api:"required"`
	// Any of "RECORD_CREATED".
	Type string `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ObjectType  respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutomationPauseResponseTriggerRecordCreatedTrigger) RawJSON() string { return r.JSON.raw }
func (r *AutomationPauseResponseTriggerRecordCreatedTrigger) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AutomationPauseResponseTriggerRecordUpdatedTrigger struct {
	ObjectType string `json:"objectType" api:"required"`
	// Any of "RECORD_UPDATED".
	Type                 string                                                                  `json:"type" api:"required"`
	WatchedFields        []string                                                                `json:"watchedFields" api:"required"`
	WatchedRelationships []AutomationPauseResponseTriggerRecordUpdatedTriggerWatchedRelationship `json:"watchedRelationships"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ObjectType           respjson.Field
		Type                 respjson.Field
		WatchedFields        respjson.Field
		WatchedRelationships respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutomationPauseResponseTriggerRecordUpdatedTrigger) RawJSON() string { return r.JSON.raw }
func (r *AutomationPauseResponseTriggerRecordUpdatedTrigger) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AutomationPauseResponseTriggerRecordUpdatedTriggerWatchedRelationship struct {
	// The endpoint occupied by the record whose changes start the run.
	//
	// Any of "SOURCE", "DESTINATION".
	Direction string `json:"direction" api:"required"`
	// Public relationship key, such as `$contact` or a custom relationship slug.
	Key string `json:"key" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Direction   respjson.Field
		Key         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutomationPauseResponseTriggerRecordUpdatedTriggerWatchedRelationship) RawJSON() string {
	return r.JSON.raw
}
func (r *AutomationPauseResponseTriggerRecordUpdatedTriggerWatchedRelationship) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AutomationPauseResponseTriggerSequenceReplyTrigger struct {
	// Any of "SEQUENCE_REPLY".
	Type string `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutomationPauseResponseTriggerSequenceReplyTrigger) RawJSON() string { return r.JSON.raw }
func (r *AutomationPauseResponseTriggerSequenceReplyTrigger) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AutomationResumeResponse struct {
	// Unique identifier of the automation.
	ID string `json:"id" api:"required"`
	// When the automation was created.
	CreatedAt string `json:"createdAt" api:"required"`
	// What the automation is meant to accomplish.
	Description string `json:"description" api:"required"`
	// When its most recent failed run ended, test runs included; null if no run has
	// failed.
	LastFailureAt string `json:"lastFailureAt" api:"required"`
	// When its most recent run started, test runs included; null before the first run.
	LastRunAt string `json:"lastRunAt" api:"required"`
	// Model the agent steps run on.
	Model string `json:"model" api:"required"`
	// Display name of the automation.
	Name string `json:"name" api:"required"`
	// Always `automation`.
	//
	// Any of "automation".
	ObjectType AutomationResumeResponseObjectType `json:"objectType" api:"required"`
	// Who the automation acts as.
	RunAs AutomationResumeResponseRunAsUnion `json:"runAs" api:"required"`
	// `PAUSED` automations keep their triggers but start no runs.
	//
	// Any of "ACTIVE", "PAUSED".
	Status   AutomationResumeResponseStatus         `json:"status" api:"required"`
	Steps    []AutomationResumeResponseStepUnion    `json:"steps" api:"required"`
	Triggers []AutomationResumeResponseTriggerUnion `json:"triggers" api:"required"`
	// When the automation was last changed.
	UpdatedAt string `json:"updatedAt" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID            respjson.Field
		CreatedAt     respjson.Field
		Description   respjson.Field
		LastFailureAt respjson.Field
		LastRunAt     respjson.Field
		Model         respjson.Field
		Name          respjson.Field
		ObjectType    respjson.Field
		RunAs         respjson.Field
		Status        respjson.Field
		Steps         respjson.Field
		Triggers      respjson.Field
		UpdatedAt     respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutomationResumeResponse) RawJSON() string { return r.JSON.raw }
func (r *AutomationResumeResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Always `automation`.
type AutomationResumeResponseObjectType string

const (
	AutomationResumeResponseObjectTypeAutomation AutomationResumeResponseObjectType = "automation"
)

// AutomationResumeResponseRunAsUnion contains all possible properties and values
// from [AutomationResumeResponseRunAsWorkspaceRunAs],
// [AutomationResumeResponseRunAsMemberRunAs],
// [AutomationResumeResponseRunAsFromInputRunAs].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type AutomationResumeResponseRunAsUnion struct {
	Type string `json:"type"`
	// This field is from variant [AutomationResumeResponseRunAsMemberRunAs].
	MemberID string `json:"memberId"`
	JSON     struct {
		Type     respjson.Field
		MemberID respjson.Field
		raw      string
	} `json:"-"`
}

func (u AutomationResumeResponseRunAsUnion) AsWorkspaceRunAs() (v AutomationResumeResponseRunAsWorkspaceRunAs) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomationResumeResponseRunAsUnion) AsMemberRunAs() (v AutomationResumeResponseRunAsMemberRunAs) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomationResumeResponseRunAsUnion) AsFromInputRunAs() (v AutomationResumeResponseRunAsFromInputRunAs) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AutomationResumeResponseRunAsUnion) RawJSON() string { return u.JSON.raw }

func (r *AutomationResumeResponseRunAsUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Runs with no acting member.
type AutomationResumeResponseRunAsWorkspaceRunAs struct {
	// Any of "WORKSPACE".
	Type string `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutomationResumeResponseRunAsWorkspaceRunAs) RawJSON() string { return r.JSON.raw }
func (r *AutomationResumeResponseRunAsWorkspaceRunAs) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AutomationResumeResponseRunAsMemberRunAs struct {
	// Member the automation runs as.
	MemberID string `json:"memberId" api:"required"`
	// Any of "MEMBER".
	Type string `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		MemberID    respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutomationResumeResponseRunAsMemberRunAs) RawJSON() string { return r.JSON.raw }
func (r *AutomationResumeResponseRunAsMemberRunAs) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The acting member is read from an input when a run starts. Set in the Lightfield
// app.
type AutomationResumeResponseRunAsFromInputRunAs struct {
	// Any of "FROM_INPUT".
	Type string `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutomationResumeResponseRunAsFromInputRunAs) RawJSON() string { return r.JSON.raw }
func (r *AutomationResumeResponseRunAsFromInputRunAs) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// `PAUSED` automations keep their triggers but start no runs.
type AutomationResumeResponseStatus string

const (
	AutomationResumeResponseStatusActive AutomationResumeResponseStatus = "ACTIVE"
	AutomationResumeResponseStatusPaused AutomationResumeResponseStatus = "PAUSED"
)

// AutomationResumeResponseStepUnion contains all possible properties and values
// from [AutomationResumeResponseStepAgentStep],
// [AutomationResumeResponseStepCodeStep], [AutomationResumeResponseStepSleepStep].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type AutomationResumeResponseStepUnion struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	// This field is from variant [AutomationResumeResponseStepAgentStep].
	Prompt string `json:"prompt"`
	Type   string `json:"type"`
	// This field is a union of
	// [[]AutomationResumeResponseStepAgentStepPermissionUnion],
	// [[]AutomationResumeResponseStepCodeStepPermissionUnion]
	Permissions AutomationResumeResponseStepUnionPermissions `json:"permissions"`
	// This field is from variant [AutomationResumeResponseStepCodeStep].
	Command string `json:"command"`
	// This field is from variant [AutomationResumeResponseStepCodeStep].
	Description string `json:"description"`
	// This field is from variant [AutomationResumeResponseStepSleepStep].
	Duration AutomationResumeResponseStepSleepStepDurationUnion `json:"duration"`
	// This field is from variant [AutomationResumeResponseStepSleepStep].
	Unit string `json:"unit"`
	JSON struct {
		ID          respjson.Field
		Name        respjson.Field
		Prompt      respjson.Field
		Type        respjson.Field
		Permissions respjson.Field
		Command     respjson.Field
		Description respjson.Field
		Duration    respjson.Field
		Unit        respjson.Field
		raw         string
	} `json:"-"`
}

func (u AutomationResumeResponseStepUnion) AsAgentStep() (v AutomationResumeResponseStepAgentStep) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomationResumeResponseStepUnion) AsCodeStep() (v AutomationResumeResponseStepCodeStep) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomationResumeResponseStepUnion) AsSleepStep() (v AutomationResumeResponseStepSleepStep) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AutomationResumeResponseStepUnion) RawJSON() string { return u.JSON.raw }

func (r *AutomationResumeResponseStepUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AutomationResumeResponseStepUnionPermissions is an implicit subunion of
// [AutomationResumeResponseStepUnion].
// AutomationResumeResponseStepUnionPermissions provides convenient access to the
// sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [AutomationResumeResponseStepUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfAutomationResumeResponseStepAgentStepPermissions
// OfAutomationResumeResponseStepCodeStepPermissions]
type AutomationResumeResponseStepUnionPermissions struct {
	// This field will be present if the value is a
	// [[]AutomationResumeResponseStepAgentStepPermissionUnion] instead of an object.
	OfAutomationResumeResponseStepAgentStepPermissions []AutomationResumeResponseStepAgentStepPermissionUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]AutomationResumeResponseStepCodeStepPermissionUnion] instead of an object.
	OfAutomationResumeResponseStepCodeStepPermissions []AutomationResumeResponseStepCodeStepPermissionUnion `json:",inline"`
	JSON                                              struct {
		OfAutomationResumeResponseStepAgentStepPermissions respjson.Field
		OfAutomationResumeResponseStepCodeStepPermissions  respjson.Field
		raw                                                string
	} `json:"-"`
}

func (r *AutomationResumeResponseStepUnionPermissions) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AutomationResumeResponseStepAgentStep struct {
	// Step identifier. Assigned by Lightfield.
	ID string `json:"id" api:"required"`
	// Display name of the step.
	Name string `json:"name" api:"required"`
	// Instructions the agent follows in this step.
	Prompt string `json:"prompt" api:"required"`
	// Any of "AGENT".
	Type string `json:"type" api:"required"`
	// What the agent may use in this step. Defaults to none.
	Permissions []AutomationResumeResponseStepAgentStepPermissionUnion `json:"permissions"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Name        respjson.Field
		Prompt      respjson.Field
		Type        respjson.Field
		Permissions respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutomationResumeResponseStepAgentStep) RawJSON() string { return r.JSON.raw }
func (r *AutomationResumeResponseStepAgentStep) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AutomationResumeResponseStepAgentStepPermissionUnion contains all possible
// properties and values from
// [AutomationResumeResponseStepAgentStepPermissionToolPermission],
// [AutomationResumeResponseStepAgentStepPermissionMcpPermission],
// [AutomationResumeResponseStepAgentStepPermissionAPIPermission],
// [AutomationResumeResponseStepAgentStepPermissionNetworkPermission].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type AutomationResumeResponseStepAgentStepPermissionUnion struct {
	// This field is from variant
	// [AutomationResumeResponseStepAgentStepPermissionToolPermission].
	Name   string `json:"name"`
	Type   string `json:"type"`
	Reason string `json:"reason"`
	// This field is from variant
	// [AutomationResumeResponseStepAgentStepPermissionMcpPermission].
	Connector string `json:"connector"`
	// This field is from variant
	// [AutomationResumeResponseStepAgentStepPermissionMcpPermission].
	Tools []string `json:"tools"`
	// This field is from variant
	// [AutomationResumeResponseStepAgentStepPermissionAPIPermission].
	Scopes []string `json:"scopes"`
	// This field is from variant
	// [AutomationResumeResponseStepAgentStepPermissionNetworkPermission].
	Domain string `json:"domain"`
	// This field is from variant
	// [AutomationResumeResponseStepAgentStepPermissionNetworkPermission].
	Headers map[string]string `json:"headers"`
	// This field is from variant
	// [AutomationResumeResponseStepAgentStepPermissionNetworkPermission].
	Method string `json:"method"`
	JSON   struct {
		Name      respjson.Field
		Type      respjson.Field
		Reason    respjson.Field
		Connector respjson.Field
		Tools     respjson.Field
		Scopes    respjson.Field
		Domain    respjson.Field
		Headers   respjson.Field
		Method    respjson.Field
		raw       string
	} `json:"-"`
}

func (u AutomationResumeResponseStepAgentStepPermissionUnion) AsToolPermission() (v AutomationResumeResponseStepAgentStepPermissionToolPermission) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomationResumeResponseStepAgentStepPermissionUnion) AsMcpPermission() (v AutomationResumeResponseStepAgentStepPermissionMcpPermission) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomationResumeResponseStepAgentStepPermissionUnion) AsAPIPermission() (v AutomationResumeResponseStepAgentStepPermissionAPIPermission) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomationResumeResponseStepAgentStepPermissionUnion) AsNetworkPermission() (v AutomationResumeResponseStepAgentStepPermissionNetworkPermission) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AutomationResumeResponseStepAgentStepPermissionUnion) RawJSON() string { return u.JSON.raw }

func (r *AutomationResumeResponseStepAgentStepPermissionUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AutomationResumeResponseStepAgentStepPermissionToolPermission struct {
	// Name of the tool, matching the Lightfield app.
	Name string `json:"name" api:"required"`
	// Any of "TOOL".
	Type string `json:"type" api:"required"`
	// Why the step needs this permission. Null when none was given. Returns 422 for an
	// empty or whitespace-only string.
	Reason string `json:"reason" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name        respjson.Field
		Type        respjson.Field
		Reason      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutomationResumeResponseStepAgentStepPermissionToolPermission) RawJSON() string {
	return r.JSON.raw
}
func (r *AutomationResumeResponseStepAgentStepPermissionToolPermission) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AutomationResumeResponseStepAgentStepPermissionMcpPermission struct {
	// Connector the tools are on. Must already be lowercase and trimmed.
	Connector string `json:"connector" api:"required"`
	// Unique, trimmed names of the connector tools the step may call.
	Tools []string `json:"tools" api:"required"`
	// Any of "MCP".
	Type string `json:"type" api:"required"`
	// Why the step needs this permission. Null when none was given. Returns 422 for an
	// empty or whitespace-only string.
	Reason string `json:"reason" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Connector   respjson.Field
		Tools       respjson.Field
		Type        respjson.Field
		Reason      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutomationResumeResponseStepAgentStepPermissionMcpPermission) RawJSON() string {
	return r.JSON.raw
}
func (r *AutomationResumeResponseStepAgentStepPermissionMcpPermission) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AutomationResumeResponseStepAgentStepPermissionAPIPermission struct {
	// Lightfield API scopes the step may use, for example `contacts:read`. The API key
	// must hold every scope it grants, and only a full-access key may grant `*`.
	Scopes []string `json:"scopes" api:"required"`
	// Any of "API".
	Type string `json:"type" api:"required"`
	// Why the step needs this permission. Null when none was given. Returns 422 for an
	// empty or whitespace-only string.
	Reason string `json:"reason" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Scopes      respjson.Field
		Type        respjson.Field
		Reason      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutomationResumeResponseStepAgentStepPermissionAPIPermission) RawJSON() string {
	return r.JSON.raw
}
func (r *AutomationResumeResponseStepAgentStepPermissionAPIPermission) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AutomationResumeResponseStepAgentStepPermissionNetworkPermission struct {
	// Lowercase host the step may call, optionally followed by a path glob. Returns
	// 422 for a scheme, query, fragment or surrounding whitespace.
	Domain string `json:"domain" api:"required"`
	// Any of "NETWORK".
	Type string `json:"type" api:"required"`
	// Headers added to every request this permission allows.
	Headers map[string]string `json:"headers"`
	// Restricts the permission to one HTTP method.
	//
	// Any of "GET", "POST", "PUT", "PATCH", "DELETE".
	Method string `json:"method"`
	// Why the step needs this permission. Null when none was given. Returns 422 for an
	// empty or whitespace-only string.
	Reason string `json:"reason" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Domain      respjson.Field
		Type        respjson.Field
		Headers     respjson.Field
		Method      respjson.Field
		Reason      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutomationResumeResponseStepAgentStepPermissionNetworkPermission) RawJSON() string {
	return r.JSON.raw
}
func (r *AutomationResumeResponseStepAgentStepPermissionNetworkPermission) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AutomationResumeResponseStepCodeStep struct {
	// Step identifier. Assigned by Lightfield.
	ID string `json:"id" api:"required"`
	// Bash command the step runs as written.
	Command string `json:"command" api:"required"`
	// Display name of the step.
	Name string `json:"name" api:"required"`
	// Any of "CODE".
	Type string `json:"type" api:"required"`
	// What the command does. Does not change what runs.
	Description string `json:"description"`
	// What the command may use. Defaults to none.
	Permissions []AutomationResumeResponseStepCodeStepPermissionUnion `json:"permissions"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Command     respjson.Field
		Name        respjson.Field
		Type        respjson.Field
		Description respjson.Field
		Permissions respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutomationResumeResponseStepCodeStep) RawJSON() string { return r.JSON.raw }
func (r *AutomationResumeResponseStepCodeStep) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AutomationResumeResponseStepCodeStepPermissionUnion contains all possible
// properties and values from
// [AutomationResumeResponseStepCodeStepPermissionAPIPermission],
// [AutomationResumeResponseStepCodeStepPermissionNetworkPermission].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type AutomationResumeResponseStepCodeStepPermissionUnion struct {
	// This field is from variant
	// [AutomationResumeResponseStepCodeStepPermissionAPIPermission].
	Scopes []string `json:"scopes"`
	Type   string   `json:"type"`
	Reason string   `json:"reason"`
	// This field is from variant
	// [AutomationResumeResponseStepCodeStepPermissionNetworkPermission].
	Domain string `json:"domain"`
	// This field is from variant
	// [AutomationResumeResponseStepCodeStepPermissionNetworkPermission].
	Headers map[string]string `json:"headers"`
	// This field is from variant
	// [AutomationResumeResponseStepCodeStepPermissionNetworkPermission].
	Method string `json:"method"`
	JSON   struct {
		Scopes  respjson.Field
		Type    respjson.Field
		Reason  respjson.Field
		Domain  respjson.Field
		Headers respjson.Field
		Method  respjson.Field
		raw     string
	} `json:"-"`
}

func (u AutomationResumeResponseStepCodeStepPermissionUnion) AsAPIPermission() (v AutomationResumeResponseStepCodeStepPermissionAPIPermission) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomationResumeResponseStepCodeStepPermissionUnion) AsNetworkPermission() (v AutomationResumeResponseStepCodeStepPermissionNetworkPermission) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AutomationResumeResponseStepCodeStepPermissionUnion) RawJSON() string { return u.JSON.raw }

func (r *AutomationResumeResponseStepCodeStepPermissionUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AutomationResumeResponseStepCodeStepPermissionAPIPermission struct {
	// Lightfield API scopes the step may use, for example `contacts:read`. The API key
	// must hold every scope it grants, and only a full-access key may grant `*`.
	Scopes []string `json:"scopes" api:"required"`
	// Any of "API".
	Type string `json:"type" api:"required"`
	// Why the step needs this permission. Null when none was given. Returns 422 for an
	// empty or whitespace-only string.
	Reason string `json:"reason" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Scopes      respjson.Field
		Type        respjson.Field
		Reason      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutomationResumeResponseStepCodeStepPermissionAPIPermission) RawJSON() string {
	return r.JSON.raw
}
func (r *AutomationResumeResponseStepCodeStepPermissionAPIPermission) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AutomationResumeResponseStepCodeStepPermissionNetworkPermission struct {
	// Lowercase host the step may call, optionally followed by a path glob. Returns
	// 422 for a scheme, query, fragment or surrounding whitespace.
	Domain string `json:"domain" api:"required"`
	// Any of "NETWORK".
	Type string `json:"type" api:"required"`
	// Headers added to every request this permission allows.
	Headers map[string]string `json:"headers"`
	// Restricts the permission to one HTTP method.
	//
	// Any of "GET", "POST", "PUT", "PATCH", "DELETE".
	Method string `json:"method"`
	// Why the step needs this permission. Null when none was given. Returns 422 for an
	// empty or whitespace-only string.
	Reason string `json:"reason" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Domain      respjson.Field
		Type        respjson.Field
		Headers     respjson.Field
		Method      respjson.Field
		Reason      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutomationResumeResponseStepCodeStepPermissionNetworkPermission) RawJSON() string {
	return r.JSON.raw
}
func (r *AutomationResumeResponseStepCodeStepPermissionNetworkPermission) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AutomationResumeResponseStepSleepStep struct {
	// Step identifier. Assigned by Lightfield.
	ID string `json:"id" api:"required"`
	// How long to sleep before the next step, at most one year.
	Duration AutomationResumeResponseStepSleepStepDurationUnion `json:"duration" api:"required"`
	// Display name of the step.
	Name string `json:"name" api:"required"`
	// Any of "SLEEP".
	Type string `json:"type" api:"required"`
	// Unit of `duration`.
	//
	// Any of "MILLISECOND", "SECOND", "MINUTE", "HOUR", "DAY".
	Unit string `json:"unit" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Duration    respjson.Field
		Name        respjson.Field
		Type        respjson.Field
		Unit        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutomationResumeResponseStepSleepStep) RawJSON() string { return r.JSON.raw }
func (r *AutomationResumeResponseStepSleepStep) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AutomationResumeResponseStepSleepStepDurationUnion contains all possible
// properties and values from [int64], [string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfInt OfString]
type AutomationResumeResponseStepSleepStepDurationUnion struct {
	// This field will be present if the value is a [int64] instead of an object.
	OfInt int64 `json:",inline"`
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	JSON     struct {
		OfInt    respjson.Field
		OfString respjson.Field
		raw      string
	} `json:"-"`
}

func (u AutomationResumeResponseStepSleepStepDurationUnion) AsInt() (v int64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomationResumeResponseStepSleepStepDurationUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AutomationResumeResponseStepSleepStepDurationUnion) RawJSON() string { return u.JSON.raw }

func (r *AutomationResumeResponseStepSleepStepDurationUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AutomationResumeResponseTriggerUnion contains all possible properties and values
// from [AutomationResumeResponseTriggerManualTrigger],
// [AutomationResumeResponseTriggerWebhookTrigger],
// [AutomationResumeResponseTriggerScheduledTrigger],
// [AutomationResumeResponseTriggerRecordCreatedTrigger],
// [AutomationResumeResponseTriggerRecordUpdatedTrigger],
// [AutomationResumeResponseTriggerSequenceReplyTrigger].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type AutomationResumeResponseTriggerUnion struct {
	Type string `json:"type"`
	// This field is from variant [AutomationResumeResponseTriggerWebhookTrigger].
	URL string `json:"url"`
	// This field is from variant [AutomationResumeResponseTriggerWebhookTrigger].
	ExampleResult AutomationResumeResponseTriggerWebhookTriggerExampleResult `json:"exampleResult"`
	// This field is from variant [AutomationResumeResponseTriggerScheduledTrigger].
	Schedule   AutomationResumeResponseTriggerScheduledTriggerSchedule `json:"schedule"`
	ObjectType string                                                  `json:"objectType"`
	// This field is from variant
	// [AutomationResumeResponseTriggerRecordUpdatedTrigger].
	WatchedFields []string `json:"watchedFields"`
	// This field is from variant
	// [AutomationResumeResponseTriggerRecordUpdatedTrigger].
	WatchedRelationships []AutomationResumeResponseTriggerRecordUpdatedTriggerWatchedRelationship `json:"watchedRelationships"`
	JSON                 struct {
		Type                 respjson.Field
		URL                  respjson.Field
		ExampleResult        respjson.Field
		Schedule             respjson.Field
		ObjectType           respjson.Field
		WatchedFields        respjson.Field
		WatchedRelationships respjson.Field
		raw                  string
	} `json:"-"`
}

func (u AutomationResumeResponseTriggerUnion) AsManualTrigger() (v AutomationResumeResponseTriggerManualTrigger) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomationResumeResponseTriggerUnion) AsWebhookTrigger() (v AutomationResumeResponseTriggerWebhookTrigger) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomationResumeResponseTriggerUnion) AsScheduledTrigger() (v AutomationResumeResponseTriggerScheduledTrigger) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomationResumeResponseTriggerUnion) AsRecordCreatedTrigger() (v AutomationResumeResponseTriggerRecordCreatedTrigger) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomationResumeResponseTriggerUnion) AsRecordUpdatedTrigger() (v AutomationResumeResponseTriggerRecordUpdatedTrigger) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomationResumeResponseTriggerUnion) AsSequenceReplyTrigger() (v AutomationResumeResponseTriggerSequenceReplyTrigger) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AutomationResumeResponseTriggerUnion) RawJSON() string { return u.JSON.raw }

func (r *AutomationResumeResponseTriggerUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AutomationResumeResponseTriggerManualTrigger struct {
	// Any of "MANUAL".
	Type string `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutomationResumeResponseTriggerManualTrigger) RawJSON() string { return r.JSON.raw }
func (r *AutomationResumeResponseTriggerManualTrigger) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AutomationResumeResponseTriggerWebhookTrigger struct {
	// Any of "WEBHOOK".
	Type string `json:"type" api:"required"`
	// URL that starts a run when it receives a POST request.
	URL string `json:"url" api:"required"`
	// A sample payload and the fields the steps read from it.
	ExampleResult AutomationResumeResponseTriggerWebhookTriggerExampleResult `json:"exampleResult"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type          respjson.Field
		URL           respjson.Field
		ExampleResult respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutomationResumeResponseTriggerWebhookTrigger) RawJSON() string { return r.JSON.raw }
func (r *AutomationResumeResponseTriggerWebhookTrigger) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A sample payload and the fields the steps read from it.
type AutomationResumeResponseTriggerWebhookTriggerExampleResult struct {
	Json   AutomationResumeResponseTriggerWebhookTriggerExampleResultJsonUnion `json:"json" api:"required"`
	Parsed []AutomationResumeResponseTriggerWebhookTriggerExampleResultParsed  `json:"parsed" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Json        respjson.Field
		Parsed      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutomationResumeResponseTriggerWebhookTriggerExampleResult) RawJSON() string {
	return r.JSON.raw
}
func (r *AutomationResumeResponseTriggerWebhookTriggerExampleResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AutomationResumeResponseTriggerWebhookTriggerExampleResultJsonUnion contains all
// possible properties and values from
// [map[string]AutomationResumeResponseTriggerWebhookTriggerExampleResultJsonMapItemUnion],
// [[]AutomationResumeResponseTriggerWebhookTriggerExampleResultJsonArrayItemUnion].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool
// OfAutomationResumeResponseTriggerWebhookTriggerExampleResultJsonMapItemArray
// OfJsonValueArray OfAnyArray
// OfAutomationResumeResponseTriggerWebhookTriggerExampleResultJsonArray]
type AutomationResumeResponseTriggerWebhookTriggerExampleResultJsonUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a
	// [[]AutomationResumeResponseTriggerWebhookTriggerExampleResultJsonMapItemArrayItemUnion]
	// instead of an object.
	OfAutomationResumeResponseTriggerWebhookTriggerExampleResultJsonMapItemArray []AutomationResumeResponseTriggerWebhookTriggerExampleResultJsonMapItemArrayItemUnion `json:",inline"`
	// This field will be present if the value is a [[]shared.JsonValueUnion] instead
	// of an object.
	OfJsonValueArray []shared.JsonValueUnion `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	// This field will be present if the value is a
	// [[]AutomationResumeResponseTriggerWebhookTriggerExampleResultJsonArrayItemUnion]
	// instead of an object.
	OfAutomationResumeResponseTriggerWebhookTriggerExampleResultJsonArray []AutomationResumeResponseTriggerWebhookTriggerExampleResultJsonArrayItemUnion `json:",inline"`
	JSON                                                                  struct {
		OfString                                                                     respjson.Field
		OfFloat                                                                      respjson.Field
		OfBool                                                                       respjson.Field
		OfAutomationResumeResponseTriggerWebhookTriggerExampleResultJsonMapItemArray respjson.Field
		OfJsonValueArray                                                             respjson.Field
		OfAnyArray                                                                   respjson.Field
		OfAutomationResumeResponseTriggerWebhookTriggerExampleResultJsonArray        respjson.Field
		raw                                                                          string
	} `json:"-"`
}

func (u AutomationResumeResponseTriggerWebhookTriggerExampleResultJsonUnion) AsAutomationResumeResponseTriggerWebhookTriggerExampleResultJsonMapMap() (v map[string]AutomationResumeResponseTriggerWebhookTriggerExampleResultJsonMapItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomationResumeResponseTriggerWebhookTriggerExampleResultJsonUnion) AsAutomationResumeResponseTriggerWebhookTriggerExampleResultJsonArray() (v []AutomationResumeResponseTriggerWebhookTriggerExampleResultJsonArrayItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AutomationResumeResponseTriggerWebhookTriggerExampleResultJsonUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *AutomationResumeResponseTriggerWebhookTriggerExampleResultJsonUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AutomationResumeResponseTriggerWebhookTriggerExampleResultJsonMapItemUnion
// contains all possible properties and values from [string], [float64], [bool],
// [[]AutomationResumeResponseTriggerWebhookTriggerExampleResultJsonMapItemArrayItemUnion],
// [map[string]AutomationResumeResponseTriggerWebhookTriggerExampleResultJsonMapItemMapItemUnion].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool
// OfAutomationResumeResponseTriggerWebhookTriggerExampleResultJsonMapItemArray
// OfJsonValueArray OfAnyArray]
type AutomationResumeResponseTriggerWebhookTriggerExampleResultJsonMapItemUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a
	// [[]AutomationResumeResponseTriggerWebhookTriggerExampleResultJsonMapItemArrayItemUnion]
	// instead of an object.
	OfAutomationResumeResponseTriggerWebhookTriggerExampleResultJsonMapItemArray []AutomationResumeResponseTriggerWebhookTriggerExampleResultJsonMapItemArrayItemUnion `json:",inline"`
	// This field will be present if the value is a [[]shared.JsonValueUnion] instead
	// of an object.
	OfJsonValueArray []shared.JsonValueUnion `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	JSON       struct {
		OfString                                                                     respjson.Field
		OfFloat                                                                      respjson.Field
		OfBool                                                                       respjson.Field
		OfAutomationResumeResponseTriggerWebhookTriggerExampleResultJsonMapItemArray respjson.Field
		OfJsonValueArray                                                             respjson.Field
		OfAnyArray                                                                   respjson.Field
		raw                                                                          string
	} `json:"-"`
}

func (u AutomationResumeResponseTriggerWebhookTriggerExampleResultJsonMapItemUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomationResumeResponseTriggerWebhookTriggerExampleResultJsonMapItemUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomationResumeResponseTriggerWebhookTriggerExampleResultJsonMapItemUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomationResumeResponseTriggerWebhookTriggerExampleResultJsonMapItemUnion) AsAutomationResumeResponseTriggerWebhookTriggerExampleResultJsonMapItemArray() (v []AutomationResumeResponseTriggerWebhookTriggerExampleResultJsonMapItemArrayItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomationResumeResponseTriggerWebhookTriggerExampleResultJsonMapItemUnion) AsAutomationResumeResponseTriggerWebhookTriggerExampleResultJsonMapItemMapMap() (v map[string]AutomationResumeResponseTriggerWebhookTriggerExampleResultJsonMapItemMapItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AutomationResumeResponseTriggerWebhookTriggerExampleResultJsonMapItemUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *AutomationResumeResponseTriggerWebhookTriggerExampleResultJsonMapItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AutomationResumeResponseTriggerWebhookTriggerExampleResultJsonMapItemArrayItemUnion
// contains all possible properties and values from [string], [float64], [bool],
// [[]shared.JsonValueUnion], [map[string]shared.JsonValueUnion].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfJsonValueArray OfAnyArray]
type AutomationResumeResponseTriggerWebhookTriggerExampleResultJsonMapItemArrayItemUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]shared.JsonValueUnion] instead
	// of an object.
	OfJsonValueArray []shared.JsonValueUnion `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	JSON       struct {
		OfString         respjson.Field
		OfFloat          respjson.Field
		OfBool           respjson.Field
		OfJsonValueArray respjson.Field
		OfAnyArray       respjson.Field
		raw              string
	} `json:"-"`
}

func (u AutomationResumeResponseTriggerWebhookTriggerExampleResultJsonMapItemArrayItemUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomationResumeResponseTriggerWebhookTriggerExampleResultJsonMapItemArrayItemUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomationResumeResponseTriggerWebhookTriggerExampleResultJsonMapItemArrayItemUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomationResumeResponseTriggerWebhookTriggerExampleResultJsonMapItemArrayItemUnion) AsJsonValueArray() (v []shared.JsonValueUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomationResumeResponseTriggerWebhookTriggerExampleResultJsonMapItemArrayItemUnion) AsJsonValueMap() (v map[string]shared.JsonValueUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AutomationResumeResponseTriggerWebhookTriggerExampleResultJsonMapItemArrayItemUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *AutomationResumeResponseTriggerWebhookTriggerExampleResultJsonMapItemArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AutomationResumeResponseTriggerWebhookTriggerExampleResultJsonMapItemMapItemUnion
// contains all possible properties and values from [string], [float64], [bool],
// [[]shared.JsonValueUnion], [map[string]shared.JsonValueUnion].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfJsonValueArray OfAnyArray]
type AutomationResumeResponseTriggerWebhookTriggerExampleResultJsonMapItemMapItemUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]shared.JsonValueUnion] instead
	// of an object.
	OfJsonValueArray []shared.JsonValueUnion `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	JSON       struct {
		OfString         respjson.Field
		OfFloat          respjson.Field
		OfBool           respjson.Field
		OfJsonValueArray respjson.Field
		OfAnyArray       respjson.Field
		raw              string
	} `json:"-"`
}

func (u AutomationResumeResponseTriggerWebhookTriggerExampleResultJsonMapItemMapItemUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomationResumeResponseTriggerWebhookTriggerExampleResultJsonMapItemMapItemUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomationResumeResponseTriggerWebhookTriggerExampleResultJsonMapItemMapItemUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomationResumeResponseTriggerWebhookTriggerExampleResultJsonMapItemMapItemUnion) AsJsonValueArray() (v []shared.JsonValueUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomationResumeResponseTriggerWebhookTriggerExampleResultJsonMapItemMapItemUnion) AsJsonValueMap() (v map[string]shared.JsonValueUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AutomationResumeResponseTriggerWebhookTriggerExampleResultJsonMapItemMapItemUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *AutomationResumeResponseTriggerWebhookTriggerExampleResultJsonMapItemMapItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AutomationResumeResponseTriggerWebhookTriggerExampleResultJsonArrayItemUnion
// contains all possible properties and values from [string], [float64], [bool],
// [[]AutomationResumeResponseTriggerWebhookTriggerExampleResultJsonArrayItemArrayItemUnion],
// [map[string]AutomationResumeResponseTriggerWebhookTriggerExampleResultJsonArrayItemMapItemUnion].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool
// OfAutomationResumeResponseTriggerWebhookTriggerExampleResultJsonArrayItemArray
// OfJsonValueArray OfAnyArray]
type AutomationResumeResponseTriggerWebhookTriggerExampleResultJsonArrayItemUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a
	// [[]AutomationResumeResponseTriggerWebhookTriggerExampleResultJsonArrayItemArrayItemUnion]
	// instead of an object.
	OfAutomationResumeResponseTriggerWebhookTriggerExampleResultJsonArrayItemArray []AutomationResumeResponseTriggerWebhookTriggerExampleResultJsonArrayItemArrayItemUnion `json:",inline"`
	// This field will be present if the value is a [[]shared.JsonValueUnion] instead
	// of an object.
	OfJsonValueArray []shared.JsonValueUnion `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	JSON       struct {
		OfString                                                                       respjson.Field
		OfFloat                                                                        respjson.Field
		OfBool                                                                         respjson.Field
		OfAutomationResumeResponseTriggerWebhookTriggerExampleResultJsonArrayItemArray respjson.Field
		OfJsonValueArray                                                               respjson.Field
		OfAnyArray                                                                     respjson.Field
		raw                                                                            string
	} `json:"-"`
}

func (u AutomationResumeResponseTriggerWebhookTriggerExampleResultJsonArrayItemUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomationResumeResponseTriggerWebhookTriggerExampleResultJsonArrayItemUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomationResumeResponseTriggerWebhookTriggerExampleResultJsonArrayItemUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomationResumeResponseTriggerWebhookTriggerExampleResultJsonArrayItemUnion) AsAutomationResumeResponseTriggerWebhookTriggerExampleResultJsonArrayItemArray() (v []AutomationResumeResponseTriggerWebhookTriggerExampleResultJsonArrayItemArrayItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomationResumeResponseTriggerWebhookTriggerExampleResultJsonArrayItemUnion) AsAutomationResumeResponseTriggerWebhookTriggerExampleResultJsonArrayItemMapMap() (v map[string]AutomationResumeResponseTriggerWebhookTriggerExampleResultJsonArrayItemMapItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AutomationResumeResponseTriggerWebhookTriggerExampleResultJsonArrayItemUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *AutomationResumeResponseTriggerWebhookTriggerExampleResultJsonArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AutomationResumeResponseTriggerWebhookTriggerExampleResultJsonArrayItemArrayItemUnion
// contains all possible properties and values from [string], [float64], [bool],
// [[]shared.JsonValueUnion], [map[string]shared.JsonValueUnion].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfJsonValueArray OfAnyArray]
type AutomationResumeResponseTriggerWebhookTriggerExampleResultJsonArrayItemArrayItemUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]shared.JsonValueUnion] instead
	// of an object.
	OfJsonValueArray []shared.JsonValueUnion `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	JSON       struct {
		OfString         respjson.Field
		OfFloat          respjson.Field
		OfBool           respjson.Field
		OfJsonValueArray respjson.Field
		OfAnyArray       respjson.Field
		raw              string
	} `json:"-"`
}

func (u AutomationResumeResponseTriggerWebhookTriggerExampleResultJsonArrayItemArrayItemUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomationResumeResponseTriggerWebhookTriggerExampleResultJsonArrayItemArrayItemUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomationResumeResponseTriggerWebhookTriggerExampleResultJsonArrayItemArrayItemUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomationResumeResponseTriggerWebhookTriggerExampleResultJsonArrayItemArrayItemUnion) AsJsonValueArray() (v []shared.JsonValueUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomationResumeResponseTriggerWebhookTriggerExampleResultJsonArrayItemArrayItemUnion) AsJsonValueMap() (v map[string]shared.JsonValueUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AutomationResumeResponseTriggerWebhookTriggerExampleResultJsonArrayItemArrayItemUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *AutomationResumeResponseTriggerWebhookTriggerExampleResultJsonArrayItemArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AutomationResumeResponseTriggerWebhookTriggerExampleResultJsonArrayItemMapItemUnion
// contains all possible properties and values from [string], [float64], [bool],
// [[]shared.JsonValueUnion], [map[string]shared.JsonValueUnion].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfJsonValueArray OfAnyArray]
type AutomationResumeResponseTriggerWebhookTriggerExampleResultJsonArrayItemMapItemUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]shared.JsonValueUnion] instead
	// of an object.
	OfJsonValueArray []shared.JsonValueUnion `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	JSON       struct {
		OfString         respjson.Field
		OfFloat          respjson.Field
		OfBool           respjson.Field
		OfJsonValueArray respjson.Field
		OfAnyArray       respjson.Field
		raw              string
	} `json:"-"`
}

func (u AutomationResumeResponseTriggerWebhookTriggerExampleResultJsonArrayItemMapItemUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomationResumeResponseTriggerWebhookTriggerExampleResultJsonArrayItemMapItemUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomationResumeResponseTriggerWebhookTriggerExampleResultJsonArrayItemMapItemUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomationResumeResponseTriggerWebhookTriggerExampleResultJsonArrayItemMapItemUnion) AsJsonValueArray() (v []shared.JsonValueUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomationResumeResponseTriggerWebhookTriggerExampleResultJsonArrayItemMapItemUnion) AsJsonValueMap() (v map[string]shared.JsonValueUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AutomationResumeResponseTriggerWebhookTriggerExampleResultJsonArrayItemMapItemUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *AutomationResumeResponseTriggerWebhookTriggerExampleResultJsonArrayItemMapItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AutomationResumeResponseTriggerWebhookTriggerExampleResultParsed struct {
	Path  []string                                                                   `json:"path" api:"required"`
	Label string                                                                     `json:"label"`
	Value AutomationResumeResponseTriggerWebhookTriggerExampleResultParsedValueUnion `json:"value" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Path        respjson.Field
		Label       respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutomationResumeResponseTriggerWebhookTriggerExampleResultParsed) RawJSON() string {
	return r.JSON.raw
}
func (r *AutomationResumeResponseTriggerWebhookTriggerExampleResultParsed) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AutomationResumeResponseTriggerWebhookTriggerExampleResultParsedValueUnion
// contains all possible properties and values from [string], [float64], [bool],
// [[]AutomationResumeResponseTriggerWebhookTriggerExampleResultParsedValueArrayItemUnion],
// [map[string]AutomationResumeResponseTriggerWebhookTriggerExampleResultParsedValueMapItemUnion].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool
// OfAutomationResumeResponseTriggerWebhookTriggerExampleResultParsedValueArray
// OfJsonValueArray OfAnyArray]
type AutomationResumeResponseTriggerWebhookTriggerExampleResultParsedValueUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a
	// [[]AutomationResumeResponseTriggerWebhookTriggerExampleResultParsedValueArrayItemUnion]
	// instead of an object.
	OfAutomationResumeResponseTriggerWebhookTriggerExampleResultParsedValueArray []AutomationResumeResponseTriggerWebhookTriggerExampleResultParsedValueArrayItemUnion `json:",inline"`
	// This field will be present if the value is a [[]shared.JsonValueUnion] instead
	// of an object.
	OfJsonValueArray []shared.JsonValueUnion `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	JSON       struct {
		OfString                                                                     respjson.Field
		OfFloat                                                                      respjson.Field
		OfBool                                                                       respjson.Field
		OfAutomationResumeResponseTriggerWebhookTriggerExampleResultParsedValueArray respjson.Field
		OfJsonValueArray                                                             respjson.Field
		OfAnyArray                                                                   respjson.Field
		raw                                                                          string
	} `json:"-"`
}

func (u AutomationResumeResponseTriggerWebhookTriggerExampleResultParsedValueUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomationResumeResponseTriggerWebhookTriggerExampleResultParsedValueUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomationResumeResponseTriggerWebhookTriggerExampleResultParsedValueUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomationResumeResponseTriggerWebhookTriggerExampleResultParsedValueUnion) AsAutomationResumeResponseTriggerWebhookTriggerExampleResultParsedValueArray() (v []AutomationResumeResponseTriggerWebhookTriggerExampleResultParsedValueArrayItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomationResumeResponseTriggerWebhookTriggerExampleResultParsedValueUnion) AsAutomationResumeResponseTriggerWebhookTriggerExampleResultParsedValueMapMap() (v map[string]AutomationResumeResponseTriggerWebhookTriggerExampleResultParsedValueMapItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AutomationResumeResponseTriggerWebhookTriggerExampleResultParsedValueUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *AutomationResumeResponseTriggerWebhookTriggerExampleResultParsedValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AutomationResumeResponseTriggerWebhookTriggerExampleResultParsedValueArrayItemUnion
// contains all possible properties and values from [string], [float64], [bool],
// [[]shared.JsonValueUnion], [map[string]shared.JsonValueUnion].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfJsonValueArray OfAnyArray]
type AutomationResumeResponseTriggerWebhookTriggerExampleResultParsedValueArrayItemUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]shared.JsonValueUnion] instead
	// of an object.
	OfJsonValueArray []shared.JsonValueUnion `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	JSON       struct {
		OfString         respjson.Field
		OfFloat          respjson.Field
		OfBool           respjson.Field
		OfJsonValueArray respjson.Field
		OfAnyArray       respjson.Field
		raw              string
	} `json:"-"`
}

func (u AutomationResumeResponseTriggerWebhookTriggerExampleResultParsedValueArrayItemUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomationResumeResponseTriggerWebhookTriggerExampleResultParsedValueArrayItemUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomationResumeResponseTriggerWebhookTriggerExampleResultParsedValueArrayItemUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomationResumeResponseTriggerWebhookTriggerExampleResultParsedValueArrayItemUnion) AsJsonValueArray() (v []shared.JsonValueUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomationResumeResponseTriggerWebhookTriggerExampleResultParsedValueArrayItemUnion) AsJsonValueMap() (v map[string]shared.JsonValueUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AutomationResumeResponseTriggerWebhookTriggerExampleResultParsedValueArrayItemUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *AutomationResumeResponseTriggerWebhookTriggerExampleResultParsedValueArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AutomationResumeResponseTriggerWebhookTriggerExampleResultParsedValueMapItemUnion
// contains all possible properties and values from [string], [float64], [bool],
// [[]shared.JsonValueUnion], [map[string]shared.JsonValueUnion].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfJsonValueArray OfAnyArray]
type AutomationResumeResponseTriggerWebhookTriggerExampleResultParsedValueMapItemUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]shared.JsonValueUnion] instead
	// of an object.
	OfJsonValueArray []shared.JsonValueUnion `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	JSON       struct {
		OfString         respjson.Field
		OfFloat          respjson.Field
		OfBool           respjson.Field
		OfJsonValueArray respjson.Field
		OfAnyArray       respjson.Field
		raw              string
	} `json:"-"`
}

func (u AutomationResumeResponseTriggerWebhookTriggerExampleResultParsedValueMapItemUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomationResumeResponseTriggerWebhookTriggerExampleResultParsedValueMapItemUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomationResumeResponseTriggerWebhookTriggerExampleResultParsedValueMapItemUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomationResumeResponseTriggerWebhookTriggerExampleResultParsedValueMapItemUnion) AsJsonValueArray() (v []shared.JsonValueUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomationResumeResponseTriggerWebhookTriggerExampleResultParsedValueMapItemUnion) AsJsonValueMap() (v map[string]shared.JsonValueUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AutomationResumeResponseTriggerWebhookTriggerExampleResultParsedValueMapItemUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *AutomationResumeResponseTriggerWebhookTriggerExampleResultParsedValueMapItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AutomationResumeResponseTriggerScheduledTrigger struct {
	Schedule AutomationResumeResponseTriggerScheduledTriggerSchedule `json:"schedule" api:"required"`
	// Any of "SCHEDULED".
	Type string `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Schedule    respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutomationResumeResponseTriggerScheduledTrigger) RawJSON() string { return r.JSON.raw }
func (r *AutomationResumeResponseTriggerScheduledTrigger) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AutomationResumeResponseTriggerScheduledTriggerSchedule struct {
	// 5-field cron expression, for example `0 9 * * 1-5`.
	Cron string `json:"cron" api:"required"`
	// IANA timezone the cron is evaluated in.
	Timezone string `json:"timezone" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Cron        respjson.Field
		Timezone    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutomationResumeResponseTriggerScheduledTriggerSchedule) RawJSON() string { return r.JSON.raw }
func (r *AutomationResumeResponseTriggerScheduledTriggerSchedule) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AutomationResumeResponseTriggerRecordCreatedTrigger struct {
	ObjectType string `json:"objectType" api:"required"`
	// Any of "RECORD_CREATED".
	Type string `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ObjectType  respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutomationResumeResponseTriggerRecordCreatedTrigger) RawJSON() string { return r.JSON.raw }
func (r *AutomationResumeResponseTriggerRecordCreatedTrigger) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AutomationResumeResponseTriggerRecordUpdatedTrigger struct {
	ObjectType string `json:"objectType" api:"required"`
	// Any of "RECORD_UPDATED".
	Type                 string                                                                   `json:"type" api:"required"`
	WatchedFields        []string                                                                 `json:"watchedFields" api:"required"`
	WatchedRelationships []AutomationResumeResponseTriggerRecordUpdatedTriggerWatchedRelationship `json:"watchedRelationships"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ObjectType           respjson.Field
		Type                 respjson.Field
		WatchedFields        respjson.Field
		WatchedRelationships respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutomationResumeResponseTriggerRecordUpdatedTrigger) RawJSON() string { return r.JSON.raw }
func (r *AutomationResumeResponseTriggerRecordUpdatedTrigger) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AutomationResumeResponseTriggerRecordUpdatedTriggerWatchedRelationship struct {
	// The endpoint occupied by the record whose changes start the run.
	//
	// Any of "SOURCE", "DESTINATION".
	Direction string `json:"direction" api:"required"`
	// Public relationship key, such as `$contact` or a custom relationship slug.
	Key string `json:"key" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Direction   respjson.Field
		Key         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutomationResumeResponseTriggerRecordUpdatedTriggerWatchedRelationship) RawJSON() string {
	return r.JSON.raw
}
func (r *AutomationResumeResponseTriggerRecordUpdatedTriggerWatchedRelationship) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AutomationResumeResponseTriggerSequenceReplyTrigger struct {
	// Any of "SEQUENCE_REPLY".
	Type string `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutomationResumeResponseTriggerSequenceReplyTrigger) RawJSON() string { return r.JSON.raw }
func (r *AutomationResumeResponseTriggerSequenceReplyTrigger) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AutomationRetrieveResponse struct {
	// Unique identifier of the automation.
	ID string `json:"id" api:"required"`
	// When the automation was created.
	CreatedAt string `json:"createdAt" api:"required"`
	// What the automation is meant to accomplish.
	Description string `json:"description" api:"required"`
	// When its most recent failed run ended, test runs included; null if no run has
	// failed.
	LastFailureAt string `json:"lastFailureAt" api:"required"`
	// When its most recent run started, test runs included; null before the first run.
	LastRunAt string `json:"lastRunAt" api:"required"`
	// Model the agent steps run on.
	Model string `json:"model" api:"required"`
	// Display name of the automation.
	Name string `json:"name" api:"required"`
	// Always `automation`.
	//
	// Any of "automation".
	ObjectType AutomationRetrieveResponseObjectType `json:"objectType" api:"required"`
	// Who the automation acts as.
	RunAs AutomationRetrieveResponseRunAsUnion `json:"runAs" api:"required"`
	// `PAUSED` automations keep their triggers but start no runs.
	//
	// Any of "ACTIVE", "PAUSED".
	Status   AutomationRetrieveResponseStatus         `json:"status" api:"required"`
	Steps    []AutomationRetrieveResponseStepUnion    `json:"steps" api:"required"`
	Triggers []AutomationRetrieveResponseTriggerUnion `json:"triggers" api:"required"`
	// When the automation was last changed.
	UpdatedAt string `json:"updatedAt" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID            respjson.Field
		CreatedAt     respjson.Field
		Description   respjson.Field
		LastFailureAt respjson.Field
		LastRunAt     respjson.Field
		Model         respjson.Field
		Name          respjson.Field
		ObjectType    respjson.Field
		RunAs         respjson.Field
		Status        respjson.Field
		Steps         respjson.Field
		Triggers      respjson.Field
		UpdatedAt     respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutomationRetrieveResponse) RawJSON() string { return r.JSON.raw }
func (r *AutomationRetrieveResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Always `automation`.
type AutomationRetrieveResponseObjectType string

const (
	AutomationRetrieveResponseObjectTypeAutomation AutomationRetrieveResponseObjectType = "automation"
)

// AutomationRetrieveResponseRunAsUnion contains all possible properties and values
// from [AutomationRetrieveResponseRunAsWorkspaceRunAs],
// [AutomationRetrieveResponseRunAsMemberRunAs],
// [AutomationRetrieveResponseRunAsFromInputRunAs].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type AutomationRetrieveResponseRunAsUnion struct {
	Type string `json:"type"`
	// This field is from variant [AutomationRetrieveResponseRunAsMemberRunAs].
	MemberID string `json:"memberId"`
	JSON     struct {
		Type     respjson.Field
		MemberID respjson.Field
		raw      string
	} `json:"-"`
}

func (u AutomationRetrieveResponseRunAsUnion) AsWorkspaceRunAs() (v AutomationRetrieveResponseRunAsWorkspaceRunAs) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomationRetrieveResponseRunAsUnion) AsMemberRunAs() (v AutomationRetrieveResponseRunAsMemberRunAs) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomationRetrieveResponseRunAsUnion) AsFromInputRunAs() (v AutomationRetrieveResponseRunAsFromInputRunAs) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AutomationRetrieveResponseRunAsUnion) RawJSON() string { return u.JSON.raw }

func (r *AutomationRetrieveResponseRunAsUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Runs with no acting member.
type AutomationRetrieveResponseRunAsWorkspaceRunAs struct {
	// Any of "WORKSPACE".
	Type string `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutomationRetrieveResponseRunAsWorkspaceRunAs) RawJSON() string { return r.JSON.raw }
func (r *AutomationRetrieveResponseRunAsWorkspaceRunAs) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AutomationRetrieveResponseRunAsMemberRunAs struct {
	// Member the automation runs as.
	MemberID string `json:"memberId" api:"required"`
	// Any of "MEMBER".
	Type string `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		MemberID    respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutomationRetrieveResponseRunAsMemberRunAs) RawJSON() string { return r.JSON.raw }
func (r *AutomationRetrieveResponseRunAsMemberRunAs) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The acting member is read from an input when a run starts. Set in the Lightfield
// app.
type AutomationRetrieveResponseRunAsFromInputRunAs struct {
	// Any of "FROM_INPUT".
	Type string `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutomationRetrieveResponseRunAsFromInputRunAs) RawJSON() string { return r.JSON.raw }
func (r *AutomationRetrieveResponseRunAsFromInputRunAs) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// `PAUSED` automations keep their triggers but start no runs.
type AutomationRetrieveResponseStatus string

const (
	AutomationRetrieveResponseStatusActive AutomationRetrieveResponseStatus = "ACTIVE"
	AutomationRetrieveResponseStatusPaused AutomationRetrieveResponseStatus = "PAUSED"
)

// AutomationRetrieveResponseStepUnion contains all possible properties and values
// from [AutomationRetrieveResponseStepAgentStep],
// [AutomationRetrieveResponseStepCodeStep],
// [AutomationRetrieveResponseStepSleepStep].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type AutomationRetrieveResponseStepUnion struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	// This field is from variant [AutomationRetrieveResponseStepAgentStep].
	Prompt string `json:"prompt"`
	Type   string `json:"type"`
	// This field is a union of
	// [[]AutomationRetrieveResponseStepAgentStepPermissionUnion],
	// [[]AutomationRetrieveResponseStepCodeStepPermissionUnion]
	Permissions AutomationRetrieveResponseStepUnionPermissions `json:"permissions"`
	// This field is from variant [AutomationRetrieveResponseStepCodeStep].
	Command string `json:"command"`
	// This field is from variant [AutomationRetrieveResponseStepCodeStep].
	Description string `json:"description"`
	// This field is from variant [AutomationRetrieveResponseStepSleepStep].
	Duration AutomationRetrieveResponseStepSleepStepDurationUnion `json:"duration"`
	// This field is from variant [AutomationRetrieveResponseStepSleepStep].
	Unit string `json:"unit"`
	JSON struct {
		ID          respjson.Field
		Name        respjson.Field
		Prompt      respjson.Field
		Type        respjson.Field
		Permissions respjson.Field
		Command     respjson.Field
		Description respjson.Field
		Duration    respjson.Field
		Unit        respjson.Field
		raw         string
	} `json:"-"`
}

func (u AutomationRetrieveResponseStepUnion) AsAgentStep() (v AutomationRetrieveResponseStepAgentStep) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomationRetrieveResponseStepUnion) AsCodeStep() (v AutomationRetrieveResponseStepCodeStep) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomationRetrieveResponseStepUnion) AsSleepStep() (v AutomationRetrieveResponseStepSleepStep) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AutomationRetrieveResponseStepUnion) RawJSON() string { return u.JSON.raw }

func (r *AutomationRetrieveResponseStepUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AutomationRetrieveResponseStepUnionPermissions is an implicit subunion of
// [AutomationRetrieveResponseStepUnion].
// AutomationRetrieveResponseStepUnionPermissions provides convenient access to the
// sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [AutomationRetrieveResponseStepUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfAutomationRetrieveResponseStepAgentStepPermissions
// OfAutomationRetrieveResponseStepCodeStepPermissions]
type AutomationRetrieveResponseStepUnionPermissions struct {
	// This field will be present if the value is a
	// [[]AutomationRetrieveResponseStepAgentStepPermissionUnion] instead of an object.
	OfAutomationRetrieveResponseStepAgentStepPermissions []AutomationRetrieveResponseStepAgentStepPermissionUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]AutomationRetrieveResponseStepCodeStepPermissionUnion] instead of an object.
	OfAutomationRetrieveResponseStepCodeStepPermissions []AutomationRetrieveResponseStepCodeStepPermissionUnion `json:",inline"`
	JSON                                                struct {
		OfAutomationRetrieveResponseStepAgentStepPermissions respjson.Field
		OfAutomationRetrieveResponseStepCodeStepPermissions  respjson.Field
		raw                                                  string
	} `json:"-"`
}

func (r *AutomationRetrieveResponseStepUnionPermissions) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AutomationRetrieveResponseStepAgentStep struct {
	// Step identifier. Assigned by Lightfield.
	ID string `json:"id" api:"required"`
	// Display name of the step.
	Name string `json:"name" api:"required"`
	// Instructions the agent follows in this step.
	Prompt string `json:"prompt" api:"required"`
	// Any of "AGENT".
	Type string `json:"type" api:"required"`
	// What the agent may use in this step. Defaults to none.
	Permissions []AutomationRetrieveResponseStepAgentStepPermissionUnion `json:"permissions"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Name        respjson.Field
		Prompt      respjson.Field
		Type        respjson.Field
		Permissions respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutomationRetrieveResponseStepAgentStep) RawJSON() string { return r.JSON.raw }
func (r *AutomationRetrieveResponseStepAgentStep) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AutomationRetrieveResponseStepAgentStepPermissionUnion contains all possible
// properties and values from
// [AutomationRetrieveResponseStepAgentStepPermissionToolPermission],
// [AutomationRetrieveResponseStepAgentStepPermissionMcpPermission],
// [AutomationRetrieveResponseStepAgentStepPermissionAPIPermission],
// [AutomationRetrieveResponseStepAgentStepPermissionNetworkPermission].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type AutomationRetrieveResponseStepAgentStepPermissionUnion struct {
	// This field is from variant
	// [AutomationRetrieveResponseStepAgentStepPermissionToolPermission].
	Name   string `json:"name"`
	Type   string `json:"type"`
	Reason string `json:"reason"`
	// This field is from variant
	// [AutomationRetrieveResponseStepAgentStepPermissionMcpPermission].
	Connector string `json:"connector"`
	// This field is from variant
	// [AutomationRetrieveResponseStepAgentStepPermissionMcpPermission].
	Tools []string `json:"tools"`
	// This field is from variant
	// [AutomationRetrieveResponseStepAgentStepPermissionAPIPermission].
	Scopes []string `json:"scopes"`
	// This field is from variant
	// [AutomationRetrieveResponseStepAgentStepPermissionNetworkPermission].
	Domain string `json:"domain"`
	// This field is from variant
	// [AutomationRetrieveResponseStepAgentStepPermissionNetworkPermission].
	Headers map[string]string `json:"headers"`
	// This field is from variant
	// [AutomationRetrieveResponseStepAgentStepPermissionNetworkPermission].
	Method string `json:"method"`
	JSON   struct {
		Name      respjson.Field
		Type      respjson.Field
		Reason    respjson.Field
		Connector respjson.Field
		Tools     respjson.Field
		Scopes    respjson.Field
		Domain    respjson.Field
		Headers   respjson.Field
		Method    respjson.Field
		raw       string
	} `json:"-"`
}

func (u AutomationRetrieveResponseStepAgentStepPermissionUnion) AsToolPermission() (v AutomationRetrieveResponseStepAgentStepPermissionToolPermission) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomationRetrieveResponseStepAgentStepPermissionUnion) AsMcpPermission() (v AutomationRetrieveResponseStepAgentStepPermissionMcpPermission) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomationRetrieveResponseStepAgentStepPermissionUnion) AsAPIPermission() (v AutomationRetrieveResponseStepAgentStepPermissionAPIPermission) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomationRetrieveResponseStepAgentStepPermissionUnion) AsNetworkPermission() (v AutomationRetrieveResponseStepAgentStepPermissionNetworkPermission) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AutomationRetrieveResponseStepAgentStepPermissionUnion) RawJSON() string { return u.JSON.raw }

func (r *AutomationRetrieveResponseStepAgentStepPermissionUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AutomationRetrieveResponseStepAgentStepPermissionToolPermission struct {
	// Name of the tool, matching the Lightfield app.
	Name string `json:"name" api:"required"`
	// Any of "TOOL".
	Type string `json:"type" api:"required"`
	// Why the step needs this permission. Null when none was given. Returns 422 for an
	// empty or whitespace-only string.
	Reason string `json:"reason" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name        respjson.Field
		Type        respjson.Field
		Reason      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutomationRetrieveResponseStepAgentStepPermissionToolPermission) RawJSON() string {
	return r.JSON.raw
}
func (r *AutomationRetrieveResponseStepAgentStepPermissionToolPermission) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AutomationRetrieveResponseStepAgentStepPermissionMcpPermission struct {
	// Connector the tools are on. Must already be lowercase and trimmed.
	Connector string `json:"connector" api:"required"`
	// Unique, trimmed names of the connector tools the step may call.
	Tools []string `json:"tools" api:"required"`
	// Any of "MCP".
	Type string `json:"type" api:"required"`
	// Why the step needs this permission. Null when none was given. Returns 422 for an
	// empty or whitespace-only string.
	Reason string `json:"reason" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Connector   respjson.Field
		Tools       respjson.Field
		Type        respjson.Field
		Reason      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutomationRetrieveResponseStepAgentStepPermissionMcpPermission) RawJSON() string {
	return r.JSON.raw
}
func (r *AutomationRetrieveResponseStepAgentStepPermissionMcpPermission) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AutomationRetrieveResponseStepAgentStepPermissionAPIPermission struct {
	// Lightfield API scopes the step may use, for example `contacts:read`. The API key
	// must hold every scope it grants, and only a full-access key may grant `*`.
	Scopes []string `json:"scopes" api:"required"`
	// Any of "API".
	Type string `json:"type" api:"required"`
	// Why the step needs this permission. Null when none was given. Returns 422 for an
	// empty or whitespace-only string.
	Reason string `json:"reason" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Scopes      respjson.Field
		Type        respjson.Field
		Reason      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutomationRetrieveResponseStepAgentStepPermissionAPIPermission) RawJSON() string {
	return r.JSON.raw
}
func (r *AutomationRetrieveResponseStepAgentStepPermissionAPIPermission) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AutomationRetrieveResponseStepAgentStepPermissionNetworkPermission struct {
	// Lowercase host the step may call, optionally followed by a path glob. Returns
	// 422 for a scheme, query, fragment or surrounding whitespace.
	Domain string `json:"domain" api:"required"`
	// Any of "NETWORK".
	Type string `json:"type" api:"required"`
	// Headers added to every request this permission allows.
	Headers map[string]string `json:"headers"`
	// Restricts the permission to one HTTP method.
	//
	// Any of "GET", "POST", "PUT", "PATCH", "DELETE".
	Method string `json:"method"`
	// Why the step needs this permission. Null when none was given. Returns 422 for an
	// empty or whitespace-only string.
	Reason string `json:"reason" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Domain      respjson.Field
		Type        respjson.Field
		Headers     respjson.Field
		Method      respjson.Field
		Reason      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutomationRetrieveResponseStepAgentStepPermissionNetworkPermission) RawJSON() string {
	return r.JSON.raw
}
func (r *AutomationRetrieveResponseStepAgentStepPermissionNetworkPermission) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AutomationRetrieveResponseStepCodeStep struct {
	// Step identifier. Assigned by Lightfield.
	ID string `json:"id" api:"required"`
	// Bash command the step runs as written.
	Command string `json:"command" api:"required"`
	// Display name of the step.
	Name string `json:"name" api:"required"`
	// Any of "CODE".
	Type string `json:"type" api:"required"`
	// What the command does. Does not change what runs.
	Description string `json:"description"`
	// What the command may use. Defaults to none.
	Permissions []AutomationRetrieveResponseStepCodeStepPermissionUnion `json:"permissions"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Command     respjson.Field
		Name        respjson.Field
		Type        respjson.Field
		Description respjson.Field
		Permissions respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutomationRetrieveResponseStepCodeStep) RawJSON() string { return r.JSON.raw }
func (r *AutomationRetrieveResponseStepCodeStep) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AutomationRetrieveResponseStepCodeStepPermissionUnion contains all possible
// properties and values from
// [AutomationRetrieveResponseStepCodeStepPermissionAPIPermission],
// [AutomationRetrieveResponseStepCodeStepPermissionNetworkPermission].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type AutomationRetrieveResponseStepCodeStepPermissionUnion struct {
	// This field is from variant
	// [AutomationRetrieveResponseStepCodeStepPermissionAPIPermission].
	Scopes []string `json:"scopes"`
	Type   string   `json:"type"`
	Reason string   `json:"reason"`
	// This field is from variant
	// [AutomationRetrieveResponseStepCodeStepPermissionNetworkPermission].
	Domain string `json:"domain"`
	// This field is from variant
	// [AutomationRetrieveResponseStepCodeStepPermissionNetworkPermission].
	Headers map[string]string `json:"headers"`
	// This field is from variant
	// [AutomationRetrieveResponseStepCodeStepPermissionNetworkPermission].
	Method string `json:"method"`
	JSON   struct {
		Scopes  respjson.Field
		Type    respjson.Field
		Reason  respjson.Field
		Domain  respjson.Field
		Headers respjson.Field
		Method  respjson.Field
		raw     string
	} `json:"-"`
}

func (u AutomationRetrieveResponseStepCodeStepPermissionUnion) AsAPIPermission() (v AutomationRetrieveResponseStepCodeStepPermissionAPIPermission) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomationRetrieveResponseStepCodeStepPermissionUnion) AsNetworkPermission() (v AutomationRetrieveResponseStepCodeStepPermissionNetworkPermission) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AutomationRetrieveResponseStepCodeStepPermissionUnion) RawJSON() string { return u.JSON.raw }

func (r *AutomationRetrieveResponseStepCodeStepPermissionUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AutomationRetrieveResponseStepCodeStepPermissionAPIPermission struct {
	// Lightfield API scopes the step may use, for example `contacts:read`. The API key
	// must hold every scope it grants, and only a full-access key may grant `*`.
	Scopes []string `json:"scopes" api:"required"`
	// Any of "API".
	Type string `json:"type" api:"required"`
	// Why the step needs this permission. Null when none was given. Returns 422 for an
	// empty or whitespace-only string.
	Reason string `json:"reason" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Scopes      respjson.Field
		Type        respjson.Field
		Reason      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutomationRetrieveResponseStepCodeStepPermissionAPIPermission) RawJSON() string {
	return r.JSON.raw
}
func (r *AutomationRetrieveResponseStepCodeStepPermissionAPIPermission) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AutomationRetrieveResponseStepCodeStepPermissionNetworkPermission struct {
	// Lowercase host the step may call, optionally followed by a path glob. Returns
	// 422 for a scheme, query, fragment or surrounding whitespace.
	Domain string `json:"domain" api:"required"`
	// Any of "NETWORK".
	Type string `json:"type" api:"required"`
	// Headers added to every request this permission allows.
	Headers map[string]string `json:"headers"`
	// Restricts the permission to one HTTP method.
	//
	// Any of "GET", "POST", "PUT", "PATCH", "DELETE".
	Method string `json:"method"`
	// Why the step needs this permission. Null when none was given. Returns 422 for an
	// empty or whitespace-only string.
	Reason string `json:"reason" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Domain      respjson.Field
		Type        respjson.Field
		Headers     respjson.Field
		Method      respjson.Field
		Reason      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutomationRetrieveResponseStepCodeStepPermissionNetworkPermission) RawJSON() string {
	return r.JSON.raw
}
func (r *AutomationRetrieveResponseStepCodeStepPermissionNetworkPermission) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AutomationRetrieveResponseStepSleepStep struct {
	// Step identifier. Assigned by Lightfield.
	ID string `json:"id" api:"required"`
	// How long to sleep before the next step, at most one year.
	Duration AutomationRetrieveResponseStepSleepStepDurationUnion `json:"duration" api:"required"`
	// Display name of the step.
	Name string `json:"name" api:"required"`
	// Any of "SLEEP".
	Type string `json:"type" api:"required"`
	// Unit of `duration`.
	//
	// Any of "MILLISECOND", "SECOND", "MINUTE", "HOUR", "DAY".
	Unit string `json:"unit" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Duration    respjson.Field
		Name        respjson.Field
		Type        respjson.Field
		Unit        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutomationRetrieveResponseStepSleepStep) RawJSON() string { return r.JSON.raw }
func (r *AutomationRetrieveResponseStepSleepStep) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AutomationRetrieveResponseStepSleepStepDurationUnion contains all possible
// properties and values from [int64], [string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfInt OfString]
type AutomationRetrieveResponseStepSleepStepDurationUnion struct {
	// This field will be present if the value is a [int64] instead of an object.
	OfInt int64 `json:",inline"`
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	JSON     struct {
		OfInt    respjson.Field
		OfString respjson.Field
		raw      string
	} `json:"-"`
}

func (u AutomationRetrieveResponseStepSleepStepDurationUnion) AsInt() (v int64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomationRetrieveResponseStepSleepStepDurationUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AutomationRetrieveResponseStepSleepStepDurationUnion) RawJSON() string { return u.JSON.raw }

func (r *AutomationRetrieveResponseStepSleepStepDurationUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AutomationRetrieveResponseTriggerUnion contains all possible properties and
// values from [AutomationRetrieveResponseTriggerManualTrigger],
// [AutomationRetrieveResponseTriggerWebhookTrigger],
// [AutomationRetrieveResponseTriggerScheduledTrigger],
// [AutomationRetrieveResponseTriggerRecordCreatedTrigger],
// [AutomationRetrieveResponseTriggerRecordUpdatedTrigger],
// [AutomationRetrieveResponseTriggerSequenceReplyTrigger].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type AutomationRetrieveResponseTriggerUnion struct {
	Type string `json:"type"`
	// This field is from variant [AutomationRetrieveResponseTriggerWebhookTrigger].
	URL string `json:"url"`
	// This field is from variant [AutomationRetrieveResponseTriggerWebhookTrigger].
	ExampleResult AutomationRetrieveResponseTriggerWebhookTriggerExampleResult `json:"exampleResult"`
	// This field is from variant [AutomationRetrieveResponseTriggerScheduledTrigger].
	Schedule   AutomationRetrieveResponseTriggerScheduledTriggerSchedule `json:"schedule"`
	ObjectType string                                                    `json:"objectType"`
	// This field is from variant
	// [AutomationRetrieveResponseTriggerRecordUpdatedTrigger].
	WatchedFields []string `json:"watchedFields"`
	// This field is from variant
	// [AutomationRetrieveResponseTriggerRecordUpdatedTrigger].
	WatchedRelationships []AutomationRetrieveResponseTriggerRecordUpdatedTriggerWatchedRelationship `json:"watchedRelationships"`
	JSON                 struct {
		Type                 respjson.Field
		URL                  respjson.Field
		ExampleResult        respjson.Field
		Schedule             respjson.Field
		ObjectType           respjson.Field
		WatchedFields        respjson.Field
		WatchedRelationships respjson.Field
		raw                  string
	} `json:"-"`
}

func (u AutomationRetrieveResponseTriggerUnion) AsManualTrigger() (v AutomationRetrieveResponseTriggerManualTrigger) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomationRetrieveResponseTriggerUnion) AsWebhookTrigger() (v AutomationRetrieveResponseTriggerWebhookTrigger) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomationRetrieveResponseTriggerUnion) AsScheduledTrigger() (v AutomationRetrieveResponseTriggerScheduledTrigger) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomationRetrieveResponseTriggerUnion) AsRecordCreatedTrigger() (v AutomationRetrieveResponseTriggerRecordCreatedTrigger) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomationRetrieveResponseTriggerUnion) AsRecordUpdatedTrigger() (v AutomationRetrieveResponseTriggerRecordUpdatedTrigger) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomationRetrieveResponseTriggerUnion) AsSequenceReplyTrigger() (v AutomationRetrieveResponseTriggerSequenceReplyTrigger) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AutomationRetrieveResponseTriggerUnion) RawJSON() string { return u.JSON.raw }

func (r *AutomationRetrieveResponseTriggerUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AutomationRetrieveResponseTriggerManualTrigger struct {
	// Any of "MANUAL".
	Type string `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutomationRetrieveResponseTriggerManualTrigger) RawJSON() string { return r.JSON.raw }
func (r *AutomationRetrieveResponseTriggerManualTrigger) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AutomationRetrieveResponseTriggerWebhookTrigger struct {
	// Any of "WEBHOOK".
	Type string `json:"type" api:"required"`
	// URL that starts a run when it receives a POST request.
	URL string `json:"url" api:"required"`
	// A sample payload and the fields the steps read from it.
	ExampleResult AutomationRetrieveResponseTriggerWebhookTriggerExampleResult `json:"exampleResult"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type          respjson.Field
		URL           respjson.Field
		ExampleResult respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutomationRetrieveResponseTriggerWebhookTrigger) RawJSON() string { return r.JSON.raw }
func (r *AutomationRetrieveResponseTriggerWebhookTrigger) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A sample payload and the fields the steps read from it.
type AutomationRetrieveResponseTriggerWebhookTriggerExampleResult struct {
	Json   AutomationRetrieveResponseTriggerWebhookTriggerExampleResultJsonUnion `json:"json" api:"required"`
	Parsed []AutomationRetrieveResponseTriggerWebhookTriggerExampleResultParsed  `json:"parsed" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Json        respjson.Field
		Parsed      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutomationRetrieveResponseTriggerWebhookTriggerExampleResult) RawJSON() string {
	return r.JSON.raw
}
func (r *AutomationRetrieveResponseTriggerWebhookTriggerExampleResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AutomationRetrieveResponseTriggerWebhookTriggerExampleResultJsonUnion contains
// all possible properties and values from
// [map[string]AutomationRetrieveResponseTriggerWebhookTriggerExampleResultJsonMapItemUnion],
// [[]AutomationRetrieveResponseTriggerWebhookTriggerExampleResultJsonArrayItemUnion].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool
// OfAutomationRetrieveResponseTriggerWebhookTriggerExampleResultJsonMapItemArray
// OfJsonValueArray OfAnyArray
// OfAutomationRetrieveResponseTriggerWebhookTriggerExampleResultJsonArray]
type AutomationRetrieveResponseTriggerWebhookTriggerExampleResultJsonUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a
	// [[]AutomationRetrieveResponseTriggerWebhookTriggerExampleResultJsonMapItemArrayItemUnion]
	// instead of an object.
	OfAutomationRetrieveResponseTriggerWebhookTriggerExampleResultJsonMapItemArray []AutomationRetrieveResponseTriggerWebhookTriggerExampleResultJsonMapItemArrayItemUnion `json:",inline"`
	// This field will be present if the value is a [[]shared.JsonValueUnion] instead
	// of an object.
	OfJsonValueArray []shared.JsonValueUnion `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	// This field will be present if the value is a
	// [[]AutomationRetrieveResponseTriggerWebhookTriggerExampleResultJsonArrayItemUnion]
	// instead of an object.
	OfAutomationRetrieveResponseTriggerWebhookTriggerExampleResultJsonArray []AutomationRetrieveResponseTriggerWebhookTriggerExampleResultJsonArrayItemUnion `json:",inline"`
	JSON                                                                    struct {
		OfString                                                                       respjson.Field
		OfFloat                                                                        respjson.Field
		OfBool                                                                         respjson.Field
		OfAutomationRetrieveResponseTriggerWebhookTriggerExampleResultJsonMapItemArray respjson.Field
		OfJsonValueArray                                                               respjson.Field
		OfAnyArray                                                                     respjson.Field
		OfAutomationRetrieveResponseTriggerWebhookTriggerExampleResultJsonArray        respjson.Field
		raw                                                                            string
	} `json:"-"`
}

func (u AutomationRetrieveResponseTriggerWebhookTriggerExampleResultJsonUnion) AsAutomationRetrieveResponseTriggerWebhookTriggerExampleResultJsonMapMap() (v map[string]AutomationRetrieveResponseTriggerWebhookTriggerExampleResultJsonMapItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomationRetrieveResponseTriggerWebhookTriggerExampleResultJsonUnion) AsAutomationRetrieveResponseTriggerWebhookTriggerExampleResultJsonArray() (v []AutomationRetrieveResponseTriggerWebhookTriggerExampleResultJsonArrayItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AutomationRetrieveResponseTriggerWebhookTriggerExampleResultJsonUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *AutomationRetrieveResponseTriggerWebhookTriggerExampleResultJsonUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AutomationRetrieveResponseTriggerWebhookTriggerExampleResultJsonMapItemUnion
// contains all possible properties and values from [string], [float64], [bool],
// [[]AutomationRetrieveResponseTriggerWebhookTriggerExampleResultJsonMapItemArrayItemUnion],
// [map[string]AutomationRetrieveResponseTriggerWebhookTriggerExampleResultJsonMapItemMapItemUnion].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool
// OfAutomationRetrieveResponseTriggerWebhookTriggerExampleResultJsonMapItemArray
// OfJsonValueArray OfAnyArray]
type AutomationRetrieveResponseTriggerWebhookTriggerExampleResultJsonMapItemUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a
	// [[]AutomationRetrieveResponseTriggerWebhookTriggerExampleResultJsonMapItemArrayItemUnion]
	// instead of an object.
	OfAutomationRetrieveResponseTriggerWebhookTriggerExampleResultJsonMapItemArray []AutomationRetrieveResponseTriggerWebhookTriggerExampleResultJsonMapItemArrayItemUnion `json:",inline"`
	// This field will be present if the value is a [[]shared.JsonValueUnion] instead
	// of an object.
	OfJsonValueArray []shared.JsonValueUnion `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	JSON       struct {
		OfString                                                                       respjson.Field
		OfFloat                                                                        respjson.Field
		OfBool                                                                         respjson.Field
		OfAutomationRetrieveResponseTriggerWebhookTriggerExampleResultJsonMapItemArray respjson.Field
		OfJsonValueArray                                                               respjson.Field
		OfAnyArray                                                                     respjson.Field
		raw                                                                            string
	} `json:"-"`
}

func (u AutomationRetrieveResponseTriggerWebhookTriggerExampleResultJsonMapItemUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomationRetrieveResponseTriggerWebhookTriggerExampleResultJsonMapItemUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomationRetrieveResponseTriggerWebhookTriggerExampleResultJsonMapItemUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomationRetrieveResponseTriggerWebhookTriggerExampleResultJsonMapItemUnion) AsAutomationRetrieveResponseTriggerWebhookTriggerExampleResultJsonMapItemArray() (v []AutomationRetrieveResponseTriggerWebhookTriggerExampleResultJsonMapItemArrayItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomationRetrieveResponseTriggerWebhookTriggerExampleResultJsonMapItemUnion) AsAutomationRetrieveResponseTriggerWebhookTriggerExampleResultJsonMapItemMapMap() (v map[string]AutomationRetrieveResponseTriggerWebhookTriggerExampleResultJsonMapItemMapItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AutomationRetrieveResponseTriggerWebhookTriggerExampleResultJsonMapItemUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *AutomationRetrieveResponseTriggerWebhookTriggerExampleResultJsonMapItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AutomationRetrieveResponseTriggerWebhookTriggerExampleResultJsonMapItemArrayItemUnion
// contains all possible properties and values from [string], [float64], [bool],
// [[]shared.JsonValueUnion], [map[string]shared.JsonValueUnion].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfJsonValueArray OfAnyArray]
type AutomationRetrieveResponseTriggerWebhookTriggerExampleResultJsonMapItemArrayItemUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]shared.JsonValueUnion] instead
	// of an object.
	OfJsonValueArray []shared.JsonValueUnion `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	JSON       struct {
		OfString         respjson.Field
		OfFloat          respjson.Field
		OfBool           respjson.Field
		OfJsonValueArray respjson.Field
		OfAnyArray       respjson.Field
		raw              string
	} `json:"-"`
}

func (u AutomationRetrieveResponseTriggerWebhookTriggerExampleResultJsonMapItemArrayItemUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomationRetrieveResponseTriggerWebhookTriggerExampleResultJsonMapItemArrayItemUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomationRetrieveResponseTriggerWebhookTriggerExampleResultJsonMapItemArrayItemUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomationRetrieveResponseTriggerWebhookTriggerExampleResultJsonMapItemArrayItemUnion) AsJsonValueArray() (v []shared.JsonValueUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomationRetrieveResponseTriggerWebhookTriggerExampleResultJsonMapItemArrayItemUnion) AsJsonValueMap() (v map[string]shared.JsonValueUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AutomationRetrieveResponseTriggerWebhookTriggerExampleResultJsonMapItemArrayItemUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *AutomationRetrieveResponseTriggerWebhookTriggerExampleResultJsonMapItemArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AutomationRetrieveResponseTriggerWebhookTriggerExampleResultJsonMapItemMapItemUnion
// contains all possible properties and values from [string], [float64], [bool],
// [[]shared.JsonValueUnion], [map[string]shared.JsonValueUnion].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfJsonValueArray OfAnyArray]
type AutomationRetrieveResponseTriggerWebhookTriggerExampleResultJsonMapItemMapItemUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]shared.JsonValueUnion] instead
	// of an object.
	OfJsonValueArray []shared.JsonValueUnion `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	JSON       struct {
		OfString         respjson.Field
		OfFloat          respjson.Field
		OfBool           respjson.Field
		OfJsonValueArray respjson.Field
		OfAnyArray       respjson.Field
		raw              string
	} `json:"-"`
}

func (u AutomationRetrieveResponseTriggerWebhookTriggerExampleResultJsonMapItemMapItemUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomationRetrieveResponseTriggerWebhookTriggerExampleResultJsonMapItemMapItemUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomationRetrieveResponseTriggerWebhookTriggerExampleResultJsonMapItemMapItemUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomationRetrieveResponseTriggerWebhookTriggerExampleResultJsonMapItemMapItemUnion) AsJsonValueArray() (v []shared.JsonValueUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomationRetrieveResponseTriggerWebhookTriggerExampleResultJsonMapItemMapItemUnion) AsJsonValueMap() (v map[string]shared.JsonValueUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AutomationRetrieveResponseTriggerWebhookTriggerExampleResultJsonMapItemMapItemUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *AutomationRetrieveResponseTriggerWebhookTriggerExampleResultJsonMapItemMapItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AutomationRetrieveResponseTriggerWebhookTriggerExampleResultJsonArrayItemUnion
// contains all possible properties and values from [string], [float64], [bool],
// [[]AutomationRetrieveResponseTriggerWebhookTriggerExampleResultJsonArrayItemArrayItemUnion],
// [map[string]AutomationRetrieveResponseTriggerWebhookTriggerExampleResultJsonArrayItemMapItemUnion].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool
// OfAutomationRetrieveResponseTriggerWebhookTriggerExampleResultJsonArrayItemArray
// OfJsonValueArray OfAnyArray]
type AutomationRetrieveResponseTriggerWebhookTriggerExampleResultJsonArrayItemUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a
	// [[]AutomationRetrieveResponseTriggerWebhookTriggerExampleResultJsonArrayItemArrayItemUnion]
	// instead of an object.
	OfAutomationRetrieveResponseTriggerWebhookTriggerExampleResultJsonArrayItemArray []AutomationRetrieveResponseTriggerWebhookTriggerExampleResultJsonArrayItemArrayItemUnion `json:",inline"`
	// This field will be present if the value is a [[]shared.JsonValueUnion] instead
	// of an object.
	OfJsonValueArray []shared.JsonValueUnion `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	JSON       struct {
		OfString                                                                         respjson.Field
		OfFloat                                                                          respjson.Field
		OfBool                                                                           respjson.Field
		OfAutomationRetrieveResponseTriggerWebhookTriggerExampleResultJsonArrayItemArray respjson.Field
		OfJsonValueArray                                                                 respjson.Field
		OfAnyArray                                                                       respjson.Field
		raw                                                                              string
	} `json:"-"`
}

func (u AutomationRetrieveResponseTriggerWebhookTriggerExampleResultJsonArrayItemUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomationRetrieveResponseTriggerWebhookTriggerExampleResultJsonArrayItemUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomationRetrieveResponseTriggerWebhookTriggerExampleResultJsonArrayItemUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomationRetrieveResponseTriggerWebhookTriggerExampleResultJsonArrayItemUnion) AsAutomationRetrieveResponseTriggerWebhookTriggerExampleResultJsonArrayItemArray() (v []AutomationRetrieveResponseTriggerWebhookTriggerExampleResultJsonArrayItemArrayItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomationRetrieveResponseTriggerWebhookTriggerExampleResultJsonArrayItemUnion) AsAutomationRetrieveResponseTriggerWebhookTriggerExampleResultJsonArrayItemMapMap() (v map[string]AutomationRetrieveResponseTriggerWebhookTriggerExampleResultJsonArrayItemMapItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AutomationRetrieveResponseTriggerWebhookTriggerExampleResultJsonArrayItemUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *AutomationRetrieveResponseTriggerWebhookTriggerExampleResultJsonArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AutomationRetrieveResponseTriggerWebhookTriggerExampleResultJsonArrayItemArrayItemUnion
// contains all possible properties and values from [string], [float64], [bool],
// [[]shared.JsonValueUnion], [map[string]shared.JsonValueUnion].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfJsonValueArray OfAnyArray]
type AutomationRetrieveResponseTriggerWebhookTriggerExampleResultJsonArrayItemArrayItemUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]shared.JsonValueUnion] instead
	// of an object.
	OfJsonValueArray []shared.JsonValueUnion `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	JSON       struct {
		OfString         respjson.Field
		OfFloat          respjson.Field
		OfBool           respjson.Field
		OfJsonValueArray respjson.Field
		OfAnyArray       respjson.Field
		raw              string
	} `json:"-"`
}

func (u AutomationRetrieveResponseTriggerWebhookTriggerExampleResultJsonArrayItemArrayItemUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomationRetrieveResponseTriggerWebhookTriggerExampleResultJsonArrayItemArrayItemUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomationRetrieveResponseTriggerWebhookTriggerExampleResultJsonArrayItemArrayItemUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomationRetrieveResponseTriggerWebhookTriggerExampleResultJsonArrayItemArrayItemUnion) AsJsonValueArray() (v []shared.JsonValueUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomationRetrieveResponseTriggerWebhookTriggerExampleResultJsonArrayItemArrayItemUnion) AsJsonValueMap() (v map[string]shared.JsonValueUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AutomationRetrieveResponseTriggerWebhookTriggerExampleResultJsonArrayItemArrayItemUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *AutomationRetrieveResponseTriggerWebhookTriggerExampleResultJsonArrayItemArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AutomationRetrieveResponseTriggerWebhookTriggerExampleResultJsonArrayItemMapItemUnion
// contains all possible properties and values from [string], [float64], [bool],
// [[]shared.JsonValueUnion], [map[string]shared.JsonValueUnion].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfJsonValueArray OfAnyArray]
type AutomationRetrieveResponseTriggerWebhookTriggerExampleResultJsonArrayItemMapItemUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]shared.JsonValueUnion] instead
	// of an object.
	OfJsonValueArray []shared.JsonValueUnion `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	JSON       struct {
		OfString         respjson.Field
		OfFloat          respjson.Field
		OfBool           respjson.Field
		OfJsonValueArray respjson.Field
		OfAnyArray       respjson.Field
		raw              string
	} `json:"-"`
}

func (u AutomationRetrieveResponseTriggerWebhookTriggerExampleResultJsonArrayItemMapItemUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomationRetrieveResponseTriggerWebhookTriggerExampleResultJsonArrayItemMapItemUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomationRetrieveResponseTriggerWebhookTriggerExampleResultJsonArrayItemMapItemUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomationRetrieveResponseTriggerWebhookTriggerExampleResultJsonArrayItemMapItemUnion) AsJsonValueArray() (v []shared.JsonValueUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomationRetrieveResponseTriggerWebhookTriggerExampleResultJsonArrayItemMapItemUnion) AsJsonValueMap() (v map[string]shared.JsonValueUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AutomationRetrieveResponseTriggerWebhookTriggerExampleResultJsonArrayItemMapItemUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *AutomationRetrieveResponseTriggerWebhookTriggerExampleResultJsonArrayItemMapItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AutomationRetrieveResponseTriggerWebhookTriggerExampleResultParsed struct {
	Path  []string                                                                     `json:"path" api:"required"`
	Label string                                                                       `json:"label"`
	Value AutomationRetrieveResponseTriggerWebhookTriggerExampleResultParsedValueUnion `json:"value" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Path        respjson.Field
		Label       respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutomationRetrieveResponseTriggerWebhookTriggerExampleResultParsed) RawJSON() string {
	return r.JSON.raw
}
func (r *AutomationRetrieveResponseTriggerWebhookTriggerExampleResultParsed) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AutomationRetrieveResponseTriggerWebhookTriggerExampleResultParsedValueUnion
// contains all possible properties and values from [string], [float64], [bool],
// [[]AutomationRetrieveResponseTriggerWebhookTriggerExampleResultParsedValueArrayItemUnion],
// [map[string]AutomationRetrieveResponseTriggerWebhookTriggerExampleResultParsedValueMapItemUnion].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool
// OfAutomationRetrieveResponseTriggerWebhookTriggerExampleResultParsedValueArray
// OfJsonValueArray OfAnyArray]
type AutomationRetrieveResponseTriggerWebhookTriggerExampleResultParsedValueUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a
	// [[]AutomationRetrieveResponseTriggerWebhookTriggerExampleResultParsedValueArrayItemUnion]
	// instead of an object.
	OfAutomationRetrieveResponseTriggerWebhookTriggerExampleResultParsedValueArray []AutomationRetrieveResponseTriggerWebhookTriggerExampleResultParsedValueArrayItemUnion `json:",inline"`
	// This field will be present if the value is a [[]shared.JsonValueUnion] instead
	// of an object.
	OfJsonValueArray []shared.JsonValueUnion `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	JSON       struct {
		OfString                                                                       respjson.Field
		OfFloat                                                                        respjson.Field
		OfBool                                                                         respjson.Field
		OfAutomationRetrieveResponseTriggerWebhookTriggerExampleResultParsedValueArray respjson.Field
		OfJsonValueArray                                                               respjson.Field
		OfAnyArray                                                                     respjson.Field
		raw                                                                            string
	} `json:"-"`
}

func (u AutomationRetrieveResponseTriggerWebhookTriggerExampleResultParsedValueUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomationRetrieveResponseTriggerWebhookTriggerExampleResultParsedValueUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomationRetrieveResponseTriggerWebhookTriggerExampleResultParsedValueUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomationRetrieveResponseTriggerWebhookTriggerExampleResultParsedValueUnion) AsAutomationRetrieveResponseTriggerWebhookTriggerExampleResultParsedValueArray() (v []AutomationRetrieveResponseTriggerWebhookTriggerExampleResultParsedValueArrayItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomationRetrieveResponseTriggerWebhookTriggerExampleResultParsedValueUnion) AsAutomationRetrieveResponseTriggerWebhookTriggerExampleResultParsedValueMapMap() (v map[string]AutomationRetrieveResponseTriggerWebhookTriggerExampleResultParsedValueMapItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AutomationRetrieveResponseTriggerWebhookTriggerExampleResultParsedValueUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *AutomationRetrieveResponseTriggerWebhookTriggerExampleResultParsedValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AutomationRetrieveResponseTriggerWebhookTriggerExampleResultParsedValueArrayItemUnion
// contains all possible properties and values from [string], [float64], [bool],
// [[]shared.JsonValueUnion], [map[string]shared.JsonValueUnion].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfJsonValueArray OfAnyArray]
type AutomationRetrieveResponseTriggerWebhookTriggerExampleResultParsedValueArrayItemUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]shared.JsonValueUnion] instead
	// of an object.
	OfJsonValueArray []shared.JsonValueUnion `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	JSON       struct {
		OfString         respjson.Field
		OfFloat          respjson.Field
		OfBool           respjson.Field
		OfJsonValueArray respjson.Field
		OfAnyArray       respjson.Field
		raw              string
	} `json:"-"`
}

func (u AutomationRetrieveResponseTriggerWebhookTriggerExampleResultParsedValueArrayItemUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomationRetrieveResponseTriggerWebhookTriggerExampleResultParsedValueArrayItemUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomationRetrieveResponseTriggerWebhookTriggerExampleResultParsedValueArrayItemUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomationRetrieveResponseTriggerWebhookTriggerExampleResultParsedValueArrayItemUnion) AsJsonValueArray() (v []shared.JsonValueUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomationRetrieveResponseTriggerWebhookTriggerExampleResultParsedValueArrayItemUnion) AsJsonValueMap() (v map[string]shared.JsonValueUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AutomationRetrieveResponseTriggerWebhookTriggerExampleResultParsedValueArrayItemUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *AutomationRetrieveResponseTriggerWebhookTriggerExampleResultParsedValueArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AutomationRetrieveResponseTriggerWebhookTriggerExampleResultParsedValueMapItemUnion
// contains all possible properties and values from [string], [float64], [bool],
// [[]shared.JsonValueUnion], [map[string]shared.JsonValueUnion].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfJsonValueArray OfAnyArray]
type AutomationRetrieveResponseTriggerWebhookTriggerExampleResultParsedValueMapItemUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]shared.JsonValueUnion] instead
	// of an object.
	OfJsonValueArray []shared.JsonValueUnion `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	JSON       struct {
		OfString         respjson.Field
		OfFloat          respjson.Field
		OfBool           respjson.Field
		OfJsonValueArray respjson.Field
		OfAnyArray       respjson.Field
		raw              string
	} `json:"-"`
}

func (u AutomationRetrieveResponseTriggerWebhookTriggerExampleResultParsedValueMapItemUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomationRetrieveResponseTriggerWebhookTriggerExampleResultParsedValueMapItemUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomationRetrieveResponseTriggerWebhookTriggerExampleResultParsedValueMapItemUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomationRetrieveResponseTriggerWebhookTriggerExampleResultParsedValueMapItemUnion) AsJsonValueArray() (v []shared.JsonValueUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomationRetrieveResponseTriggerWebhookTriggerExampleResultParsedValueMapItemUnion) AsJsonValueMap() (v map[string]shared.JsonValueUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AutomationRetrieveResponseTriggerWebhookTriggerExampleResultParsedValueMapItemUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *AutomationRetrieveResponseTriggerWebhookTriggerExampleResultParsedValueMapItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AutomationRetrieveResponseTriggerScheduledTrigger struct {
	Schedule AutomationRetrieveResponseTriggerScheduledTriggerSchedule `json:"schedule" api:"required"`
	// Any of "SCHEDULED".
	Type string `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Schedule    respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutomationRetrieveResponseTriggerScheduledTrigger) RawJSON() string { return r.JSON.raw }
func (r *AutomationRetrieveResponseTriggerScheduledTrigger) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AutomationRetrieveResponseTriggerScheduledTriggerSchedule struct {
	// 5-field cron expression, for example `0 9 * * 1-5`.
	Cron string `json:"cron" api:"required"`
	// IANA timezone the cron is evaluated in.
	Timezone string `json:"timezone" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Cron        respjson.Field
		Timezone    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutomationRetrieveResponseTriggerScheduledTriggerSchedule) RawJSON() string {
	return r.JSON.raw
}
func (r *AutomationRetrieveResponseTriggerScheduledTriggerSchedule) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AutomationRetrieveResponseTriggerRecordCreatedTrigger struct {
	ObjectType string `json:"objectType" api:"required"`
	// Any of "RECORD_CREATED".
	Type string `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ObjectType  respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutomationRetrieveResponseTriggerRecordCreatedTrigger) RawJSON() string { return r.JSON.raw }
func (r *AutomationRetrieveResponseTriggerRecordCreatedTrigger) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AutomationRetrieveResponseTriggerRecordUpdatedTrigger struct {
	ObjectType string `json:"objectType" api:"required"`
	// Any of "RECORD_UPDATED".
	Type                 string                                                                     `json:"type" api:"required"`
	WatchedFields        []string                                                                   `json:"watchedFields" api:"required"`
	WatchedRelationships []AutomationRetrieveResponseTriggerRecordUpdatedTriggerWatchedRelationship `json:"watchedRelationships"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ObjectType           respjson.Field
		Type                 respjson.Field
		WatchedFields        respjson.Field
		WatchedRelationships respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutomationRetrieveResponseTriggerRecordUpdatedTrigger) RawJSON() string { return r.JSON.raw }
func (r *AutomationRetrieveResponseTriggerRecordUpdatedTrigger) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AutomationRetrieveResponseTriggerRecordUpdatedTriggerWatchedRelationship struct {
	// The endpoint occupied by the record whose changes start the run.
	//
	// Any of "SOURCE", "DESTINATION".
	Direction string `json:"direction" api:"required"`
	// Public relationship key, such as `$contact` or a custom relationship slug.
	Key string `json:"key" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Direction   respjson.Field
		Key         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutomationRetrieveResponseTriggerRecordUpdatedTriggerWatchedRelationship) RawJSON() string {
	return r.JSON.raw
}
func (r *AutomationRetrieveResponseTriggerRecordUpdatedTriggerWatchedRelationship) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AutomationRetrieveResponseTriggerSequenceReplyTrigger struct {
	// Any of "SEQUENCE_REPLY".
	Type string `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutomationRetrieveResponseTriggerSequenceReplyTrigger) RawJSON() string { return r.JSON.raw }
func (r *AutomationRetrieveResponseTriggerSequenceReplyTrigger) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AutomationUpdateResponse struct {
	// Unique identifier of the automation.
	ID string `json:"id" api:"required"`
	// When the automation was created.
	CreatedAt string `json:"createdAt" api:"required"`
	// What the automation is meant to accomplish.
	Description string `json:"description" api:"required"`
	// When its most recent failed run ended, test runs included; null if no run has
	// failed.
	LastFailureAt string `json:"lastFailureAt" api:"required"`
	// When its most recent run started, test runs included; null before the first run.
	LastRunAt string `json:"lastRunAt" api:"required"`
	// Model the agent steps run on.
	Model string `json:"model" api:"required"`
	// Display name of the automation.
	Name string `json:"name" api:"required"`
	// Always `automation`.
	//
	// Any of "automation".
	ObjectType AutomationUpdateResponseObjectType `json:"objectType" api:"required"`
	// Who the automation acts as.
	RunAs AutomationUpdateResponseRunAsUnion `json:"runAs" api:"required"`
	// `PAUSED` automations keep their triggers but start no runs.
	//
	// Any of "ACTIVE", "PAUSED".
	Status   AutomationUpdateResponseStatus         `json:"status" api:"required"`
	Steps    []AutomationUpdateResponseStepUnion    `json:"steps" api:"required"`
	Triggers []AutomationUpdateResponseTriggerUnion `json:"triggers" api:"required"`
	// When the automation was last changed.
	UpdatedAt string `json:"updatedAt" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID            respjson.Field
		CreatedAt     respjson.Field
		Description   respjson.Field
		LastFailureAt respjson.Field
		LastRunAt     respjson.Field
		Model         respjson.Field
		Name          respjson.Field
		ObjectType    respjson.Field
		RunAs         respjson.Field
		Status        respjson.Field
		Steps         respjson.Field
		Triggers      respjson.Field
		UpdatedAt     respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutomationUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *AutomationUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Always `automation`.
type AutomationUpdateResponseObjectType string

const (
	AutomationUpdateResponseObjectTypeAutomation AutomationUpdateResponseObjectType = "automation"
)

// AutomationUpdateResponseRunAsUnion contains all possible properties and values
// from [AutomationUpdateResponseRunAsWorkspaceRunAs],
// [AutomationUpdateResponseRunAsMemberRunAs],
// [AutomationUpdateResponseRunAsFromInputRunAs].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type AutomationUpdateResponseRunAsUnion struct {
	Type string `json:"type"`
	// This field is from variant [AutomationUpdateResponseRunAsMemberRunAs].
	MemberID string `json:"memberId"`
	JSON     struct {
		Type     respjson.Field
		MemberID respjson.Field
		raw      string
	} `json:"-"`
}

func (u AutomationUpdateResponseRunAsUnion) AsWorkspaceRunAs() (v AutomationUpdateResponseRunAsWorkspaceRunAs) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomationUpdateResponseRunAsUnion) AsMemberRunAs() (v AutomationUpdateResponseRunAsMemberRunAs) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomationUpdateResponseRunAsUnion) AsFromInputRunAs() (v AutomationUpdateResponseRunAsFromInputRunAs) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AutomationUpdateResponseRunAsUnion) RawJSON() string { return u.JSON.raw }

func (r *AutomationUpdateResponseRunAsUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Runs with no acting member.
type AutomationUpdateResponseRunAsWorkspaceRunAs struct {
	// Any of "WORKSPACE".
	Type string `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutomationUpdateResponseRunAsWorkspaceRunAs) RawJSON() string { return r.JSON.raw }
func (r *AutomationUpdateResponseRunAsWorkspaceRunAs) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AutomationUpdateResponseRunAsMemberRunAs struct {
	// Member the automation runs as.
	MemberID string `json:"memberId" api:"required"`
	// Any of "MEMBER".
	Type string `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		MemberID    respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutomationUpdateResponseRunAsMemberRunAs) RawJSON() string { return r.JSON.raw }
func (r *AutomationUpdateResponseRunAsMemberRunAs) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The acting member is read from an input when a run starts. Set in the Lightfield
// app.
type AutomationUpdateResponseRunAsFromInputRunAs struct {
	// Any of "FROM_INPUT".
	Type string `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutomationUpdateResponseRunAsFromInputRunAs) RawJSON() string { return r.JSON.raw }
func (r *AutomationUpdateResponseRunAsFromInputRunAs) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// `PAUSED` automations keep their triggers but start no runs.
type AutomationUpdateResponseStatus string

const (
	AutomationUpdateResponseStatusActive AutomationUpdateResponseStatus = "ACTIVE"
	AutomationUpdateResponseStatusPaused AutomationUpdateResponseStatus = "PAUSED"
)

// AutomationUpdateResponseStepUnion contains all possible properties and values
// from [AutomationUpdateResponseStepAgentStep],
// [AutomationUpdateResponseStepCodeStep], [AutomationUpdateResponseStepSleepStep].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type AutomationUpdateResponseStepUnion struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	// This field is from variant [AutomationUpdateResponseStepAgentStep].
	Prompt string `json:"prompt"`
	Type   string `json:"type"`
	// This field is a union of
	// [[]AutomationUpdateResponseStepAgentStepPermissionUnion],
	// [[]AutomationUpdateResponseStepCodeStepPermissionUnion]
	Permissions AutomationUpdateResponseStepUnionPermissions `json:"permissions"`
	// This field is from variant [AutomationUpdateResponseStepCodeStep].
	Command string `json:"command"`
	// This field is from variant [AutomationUpdateResponseStepCodeStep].
	Description string `json:"description"`
	// This field is from variant [AutomationUpdateResponseStepSleepStep].
	Duration AutomationUpdateResponseStepSleepStepDurationUnion `json:"duration"`
	// This field is from variant [AutomationUpdateResponseStepSleepStep].
	Unit string `json:"unit"`
	JSON struct {
		ID          respjson.Field
		Name        respjson.Field
		Prompt      respjson.Field
		Type        respjson.Field
		Permissions respjson.Field
		Command     respjson.Field
		Description respjson.Field
		Duration    respjson.Field
		Unit        respjson.Field
		raw         string
	} `json:"-"`
}

func (u AutomationUpdateResponseStepUnion) AsAgentStep() (v AutomationUpdateResponseStepAgentStep) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomationUpdateResponseStepUnion) AsCodeStep() (v AutomationUpdateResponseStepCodeStep) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomationUpdateResponseStepUnion) AsSleepStep() (v AutomationUpdateResponseStepSleepStep) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AutomationUpdateResponseStepUnion) RawJSON() string { return u.JSON.raw }

func (r *AutomationUpdateResponseStepUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AutomationUpdateResponseStepUnionPermissions is an implicit subunion of
// [AutomationUpdateResponseStepUnion].
// AutomationUpdateResponseStepUnionPermissions provides convenient access to the
// sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [AutomationUpdateResponseStepUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfAutomationUpdateResponseStepAgentStepPermissions
// OfAutomationUpdateResponseStepCodeStepPermissions]
type AutomationUpdateResponseStepUnionPermissions struct {
	// This field will be present if the value is a
	// [[]AutomationUpdateResponseStepAgentStepPermissionUnion] instead of an object.
	OfAutomationUpdateResponseStepAgentStepPermissions []AutomationUpdateResponseStepAgentStepPermissionUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]AutomationUpdateResponseStepCodeStepPermissionUnion] instead of an object.
	OfAutomationUpdateResponseStepCodeStepPermissions []AutomationUpdateResponseStepCodeStepPermissionUnion `json:",inline"`
	JSON                                              struct {
		OfAutomationUpdateResponseStepAgentStepPermissions respjson.Field
		OfAutomationUpdateResponseStepCodeStepPermissions  respjson.Field
		raw                                                string
	} `json:"-"`
}

func (r *AutomationUpdateResponseStepUnionPermissions) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AutomationUpdateResponseStepAgentStep struct {
	// Step identifier. Assigned by Lightfield.
	ID string `json:"id" api:"required"`
	// Display name of the step.
	Name string `json:"name" api:"required"`
	// Instructions the agent follows in this step.
	Prompt string `json:"prompt" api:"required"`
	// Any of "AGENT".
	Type string `json:"type" api:"required"`
	// What the agent may use in this step. Defaults to none.
	Permissions []AutomationUpdateResponseStepAgentStepPermissionUnion `json:"permissions"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Name        respjson.Field
		Prompt      respjson.Field
		Type        respjson.Field
		Permissions respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutomationUpdateResponseStepAgentStep) RawJSON() string { return r.JSON.raw }
func (r *AutomationUpdateResponseStepAgentStep) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AutomationUpdateResponseStepAgentStepPermissionUnion contains all possible
// properties and values from
// [AutomationUpdateResponseStepAgentStepPermissionToolPermission],
// [AutomationUpdateResponseStepAgentStepPermissionMcpPermission],
// [AutomationUpdateResponseStepAgentStepPermissionAPIPermission],
// [AutomationUpdateResponseStepAgentStepPermissionNetworkPermission].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type AutomationUpdateResponseStepAgentStepPermissionUnion struct {
	// This field is from variant
	// [AutomationUpdateResponseStepAgentStepPermissionToolPermission].
	Name   string `json:"name"`
	Type   string `json:"type"`
	Reason string `json:"reason"`
	// This field is from variant
	// [AutomationUpdateResponseStepAgentStepPermissionMcpPermission].
	Connector string `json:"connector"`
	// This field is from variant
	// [AutomationUpdateResponseStepAgentStepPermissionMcpPermission].
	Tools []string `json:"tools"`
	// This field is from variant
	// [AutomationUpdateResponseStepAgentStepPermissionAPIPermission].
	Scopes []string `json:"scopes"`
	// This field is from variant
	// [AutomationUpdateResponseStepAgentStepPermissionNetworkPermission].
	Domain string `json:"domain"`
	// This field is from variant
	// [AutomationUpdateResponseStepAgentStepPermissionNetworkPermission].
	Headers map[string]string `json:"headers"`
	// This field is from variant
	// [AutomationUpdateResponseStepAgentStepPermissionNetworkPermission].
	Method string `json:"method"`
	JSON   struct {
		Name      respjson.Field
		Type      respjson.Field
		Reason    respjson.Field
		Connector respjson.Field
		Tools     respjson.Field
		Scopes    respjson.Field
		Domain    respjson.Field
		Headers   respjson.Field
		Method    respjson.Field
		raw       string
	} `json:"-"`
}

func (u AutomationUpdateResponseStepAgentStepPermissionUnion) AsToolPermission() (v AutomationUpdateResponseStepAgentStepPermissionToolPermission) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomationUpdateResponseStepAgentStepPermissionUnion) AsMcpPermission() (v AutomationUpdateResponseStepAgentStepPermissionMcpPermission) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomationUpdateResponseStepAgentStepPermissionUnion) AsAPIPermission() (v AutomationUpdateResponseStepAgentStepPermissionAPIPermission) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomationUpdateResponseStepAgentStepPermissionUnion) AsNetworkPermission() (v AutomationUpdateResponseStepAgentStepPermissionNetworkPermission) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AutomationUpdateResponseStepAgentStepPermissionUnion) RawJSON() string { return u.JSON.raw }

func (r *AutomationUpdateResponseStepAgentStepPermissionUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AutomationUpdateResponseStepAgentStepPermissionToolPermission struct {
	// Name of the tool, matching the Lightfield app.
	Name string `json:"name" api:"required"`
	// Any of "TOOL".
	Type string `json:"type" api:"required"`
	// Why the step needs this permission. Null when none was given. Returns 422 for an
	// empty or whitespace-only string.
	Reason string `json:"reason" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name        respjson.Field
		Type        respjson.Field
		Reason      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutomationUpdateResponseStepAgentStepPermissionToolPermission) RawJSON() string {
	return r.JSON.raw
}
func (r *AutomationUpdateResponseStepAgentStepPermissionToolPermission) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AutomationUpdateResponseStepAgentStepPermissionMcpPermission struct {
	// Connector the tools are on. Must already be lowercase and trimmed.
	Connector string `json:"connector" api:"required"`
	// Unique, trimmed names of the connector tools the step may call.
	Tools []string `json:"tools" api:"required"`
	// Any of "MCP".
	Type string `json:"type" api:"required"`
	// Why the step needs this permission. Null when none was given. Returns 422 for an
	// empty or whitespace-only string.
	Reason string `json:"reason" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Connector   respjson.Field
		Tools       respjson.Field
		Type        respjson.Field
		Reason      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutomationUpdateResponseStepAgentStepPermissionMcpPermission) RawJSON() string {
	return r.JSON.raw
}
func (r *AutomationUpdateResponseStepAgentStepPermissionMcpPermission) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AutomationUpdateResponseStepAgentStepPermissionAPIPermission struct {
	// Lightfield API scopes the step may use, for example `contacts:read`. The API key
	// must hold every scope it grants, and only a full-access key may grant `*`.
	Scopes []string `json:"scopes" api:"required"`
	// Any of "API".
	Type string `json:"type" api:"required"`
	// Why the step needs this permission. Null when none was given. Returns 422 for an
	// empty or whitespace-only string.
	Reason string `json:"reason" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Scopes      respjson.Field
		Type        respjson.Field
		Reason      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutomationUpdateResponseStepAgentStepPermissionAPIPermission) RawJSON() string {
	return r.JSON.raw
}
func (r *AutomationUpdateResponseStepAgentStepPermissionAPIPermission) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AutomationUpdateResponseStepAgentStepPermissionNetworkPermission struct {
	// Lowercase host the step may call, optionally followed by a path glob. Returns
	// 422 for a scheme, query, fragment or surrounding whitespace.
	Domain string `json:"domain" api:"required"`
	// Any of "NETWORK".
	Type string `json:"type" api:"required"`
	// Headers added to every request this permission allows.
	Headers map[string]string `json:"headers"`
	// Restricts the permission to one HTTP method.
	//
	// Any of "GET", "POST", "PUT", "PATCH", "DELETE".
	Method string `json:"method"`
	// Why the step needs this permission. Null when none was given. Returns 422 for an
	// empty or whitespace-only string.
	Reason string `json:"reason" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Domain      respjson.Field
		Type        respjson.Field
		Headers     respjson.Field
		Method      respjson.Field
		Reason      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutomationUpdateResponseStepAgentStepPermissionNetworkPermission) RawJSON() string {
	return r.JSON.raw
}
func (r *AutomationUpdateResponseStepAgentStepPermissionNetworkPermission) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AutomationUpdateResponseStepCodeStep struct {
	// Step identifier. Assigned by Lightfield.
	ID string `json:"id" api:"required"`
	// Bash command the step runs as written.
	Command string `json:"command" api:"required"`
	// Display name of the step.
	Name string `json:"name" api:"required"`
	// Any of "CODE".
	Type string `json:"type" api:"required"`
	// What the command does. Does not change what runs.
	Description string `json:"description"`
	// What the command may use. Defaults to none.
	Permissions []AutomationUpdateResponseStepCodeStepPermissionUnion `json:"permissions"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Command     respjson.Field
		Name        respjson.Field
		Type        respjson.Field
		Description respjson.Field
		Permissions respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutomationUpdateResponseStepCodeStep) RawJSON() string { return r.JSON.raw }
func (r *AutomationUpdateResponseStepCodeStep) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AutomationUpdateResponseStepCodeStepPermissionUnion contains all possible
// properties and values from
// [AutomationUpdateResponseStepCodeStepPermissionAPIPermission],
// [AutomationUpdateResponseStepCodeStepPermissionNetworkPermission].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type AutomationUpdateResponseStepCodeStepPermissionUnion struct {
	// This field is from variant
	// [AutomationUpdateResponseStepCodeStepPermissionAPIPermission].
	Scopes []string `json:"scopes"`
	Type   string   `json:"type"`
	Reason string   `json:"reason"`
	// This field is from variant
	// [AutomationUpdateResponseStepCodeStepPermissionNetworkPermission].
	Domain string `json:"domain"`
	// This field is from variant
	// [AutomationUpdateResponseStepCodeStepPermissionNetworkPermission].
	Headers map[string]string `json:"headers"`
	// This field is from variant
	// [AutomationUpdateResponseStepCodeStepPermissionNetworkPermission].
	Method string `json:"method"`
	JSON   struct {
		Scopes  respjson.Field
		Type    respjson.Field
		Reason  respjson.Field
		Domain  respjson.Field
		Headers respjson.Field
		Method  respjson.Field
		raw     string
	} `json:"-"`
}

func (u AutomationUpdateResponseStepCodeStepPermissionUnion) AsAPIPermission() (v AutomationUpdateResponseStepCodeStepPermissionAPIPermission) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomationUpdateResponseStepCodeStepPermissionUnion) AsNetworkPermission() (v AutomationUpdateResponseStepCodeStepPermissionNetworkPermission) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AutomationUpdateResponseStepCodeStepPermissionUnion) RawJSON() string { return u.JSON.raw }

func (r *AutomationUpdateResponseStepCodeStepPermissionUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AutomationUpdateResponseStepCodeStepPermissionAPIPermission struct {
	// Lightfield API scopes the step may use, for example `contacts:read`. The API key
	// must hold every scope it grants, and only a full-access key may grant `*`.
	Scopes []string `json:"scopes" api:"required"`
	// Any of "API".
	Type string `json:"type" api:"required"`
	// Why the step needs this permission. Null when none was given. Returns 422 for an
	// empty or whitespace-only string.
	Reason string `json:"reason" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Scopes      respjson.Field
		Type        respjson.Field
		Reason      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutomationUpdateResponseStepCodeStepPermissionAPIPermission) RawJSON() string {
	return r.JSON.raw
}
func (r *AutomationUpdateResponseStepCodeStepPermissionAPIPermission) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AutomationUpdateResponseStepCodeStepPermissionNetworkPermission struct {
	// Lowercase host the step may call, optionally followed by a path glob. Returns
	// 422 for a scheme, query, fragment or surrounding whitespace.
	Domain string `json:"domain" api:"required"`
	// Any of "NETWORK".
	Type string `json:"type" api:"required"`
	// Headers added to every request this permission allows.
	Headers map[string]string `json:"headers"`
	// Restricts the permission to one HTTP method.
	//
	// Any of "GET", "POST", "PUT", "PATCH", "DELETE".
	Method string `json:"method"`
	// Why the step needs this permission. Null when none was given. Returns 422 for an
	// empty or whitespace-only string.
	Reason string `json:"reason" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Domain      respjson.Field
		Type        respjson.Field
		Headers     respjson.Field
		Method      respjson.Field
		Reason      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutomationUpdateResponseStepCodeStepPermissionNetworkPermission) RawJSON() string {
	return r.JSON.raw
}
func (r *AutomationUpdateResponseStepCodeStepPermissionNetworkPermission) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AutomationUpdateResponseStepSleepStep struct {
	// Step identifier. Assigned by Lightfield.
	ID string `json:"id" api:"required"`
	// How long to sleep before the next step, at most one year.
	Duration AutomationUpdateResponseStepSleepStepDurationUnion `json:"duration" api:"required"`
	// Display name of the step.
	Name string `json:"name" api:"required"`
	// Any of "SLEEP".
	Type string `json:"type" api:"required"`
	// Unit of `duration`.
	//
	// Any of "MILLISECOND", "SECOND", "MINUTE", "HOUR", "DAY".
	Unit string `json:"unit" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Duration    respjson.Field
		Name        respjson.Field
		Type        respjson.Field
		Unit        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutomationUpdateResponseStepSleepStep) RawJSON() string { return r.JSON.raw }
func (r *AutomationUpdateResponseStepSleepStep) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AutomationUpdateResponseStepSleepStepDurationUnion contains all possible
// properties and values from [int64], [string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfInt OfString]
type AutomationUpdateResponseStepSleepStepDurationUnion struct {
	// This field will be present if the value is a [int64] instead of an object.
	OfInt int64 `json:",inline"`
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	JSON     struct {
		OfInt    respjson.Field
		OfString respjson.Field
		raw      string
	} `json:"-"`
}

func (u AutomationUpdateResponseStepSleepStepDurationUnion) AsInt() (v int64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomationUpdateResponseStepSleepStepDurationUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AutomationUpdateResponseStepSleepStepDurationUnion) RawJSON() string { return u.JSON.raw }

func (r *AutomationUpdateResponseStepSleepStepDurationUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AutomationUpdateResponseTriggerUnion contains all possible properties and values
// from [AutomationUpdateResponseTriggerManualTrigger],
// [AutomationUpdateResponseTriggerWebhookTrigger],
// [AutomationUpdateResponseTriggerScheduledTrigger],
// [AutomationUpdateResponseTriggerRecordCreatedTrigger],
// [AutomationUpdateResponseTriggerRecordUpdatedTrigger],
// [AutomationUpdateResponseTriggerSequenceReplyTrigger].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type AutomationUpdateResponseTriggerUnion struct {
	Type string `json:"type"`
	// This field is from variant [AutomationUpdateResponseTriggerWebhookTrigger].
	URL string `json:"url"`
	// This field is from variant [AutomationUpdateResponseTriggerWebhookTrigger].
	ExampleResult AutomationUpdateResponseTriggerWebhookTriggerExampleResult `json:"exampleResult"`
	// This field is from variant [AutomationUpdateResponseTriggerScheduledTrigger].
	Schedule   AutomationUpdateResponseTriggerScheduledTriggerSchedule `json:"schedule"`
	ObjectType string                                                  `json:"objectType"`
	// This field is from variant
	// [AutomationUpdateResponseTriggerRecordUpdatedTrigger].
	WatchedFields []string `json:"watchedFields"`
	// This field is from variant
	// [AutomationUpdateResponseTriggerRecordUpdatedTrigger].
	WatchedRelationships []AutomationUpdateResponseTriggerRecordUpdatedTriggerWatchedRelationship `json:"watchedRelationships"`
	JSON                 struct {
		Type                 respjson.Field
		URL                  respjson.Field
		ExampleResult        respjson.Field
		Schedule             respjson.Field
		ObjectType           respjson.Field
		WatchedFields        respjson.Field
		WatchedRelationships respjson.Field
		raw                  string
	} `json:"-"`
}

func (u AutomationUpdateResponseTriggerUnion) AsManualTrigger() (v AutomationUpdateResponseTriggerManualTrigger) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomationUpdateResponseTriggerUnion) AsWebhookTrigger() (v AutomationUpdateResponseTriggerWebhookTrigger) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomationUpdateResponseTriggerUnion) AsScheduledTrigger() (v AutomationUpdateResponseTriggerScheduledTrigger) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomationUpdateResponseTriggerUnion) AsRecordCreatedTrigger() (v AutomationUpdateResponseTriggerRecordCreatedTrigger) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomationUpdateResponseTriggerUnion) AsRecordUpdatedTrigger() (v AutomationUpdateResponseTriggerRecordUpdatedTrigger) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomationUpdateResponseTriggerUnion) AsSequenceReplyTrigger() (v AutomationUpdateResponseTriggerSequenceReplyTrigger) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AutomationUpdateResponseTriggerUnion) RawJSON() string { return u.JSON.raw }

func (r *AutomationUpdateResponseTriggerUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AutomationUpdateResponseTriggerManualTrigger struct {
	// Any of "MANUAL".
	Type string `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutomationUpdateResponseTriggerManualTrigger) RawJSON() string { return r.JSON.raw }
func (r *AutomationUpdateResponseTriggerManualTrigger) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AutomationUpdateResponseTriggerWebhookTrigger struct {
	// Any of "WEBHOOK".
	Type string `json:"type" api:"required"`
	// URL that starts a run when it receives a POST request.
	URL string `json:"url" api:"required"`
	// A sample payload and the fields the steps read from it.
	ExampleResult AutomationUpdateResponseTriggerWebhookTriggerExampleResult `json:"exampleResult"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type          respjson.Field
		URL           respjson.Field
		ExampleResult respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutomationUpdateResponseTriggerWebhookTrigger) RawJSON() string { return r.JSON.raw }
func (r *AutomationUpdateResponseTriggerWebhookTrigger) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A sample payload and the fields the steps read from it.
type AutomationUpdateResponseTriggerWebhookTriggerExampleResult struct {
	Json   AutomationUpdateResponseTriggerWebhookTriggerExampleResultJsonUnion `json:"json" api:"required"`
	Parsed []AutomationUpdateResponseTriggerWebhookTriggerExampleResultParsed  `json:"parsed" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Json        respjson.Field
		Parsed      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutomationUpdateResponseTriggerWebhookTriggerExampleResult) RawJSON() string {
	return r.JSON.raw
}
func (r *AutomationUpdateResponseTriggerWebhookTriggerExampleResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AutomationUpdateResponseTriggerWebhookTriggerExampleResultJsonUnion contains all
// possible properties and values from
// [map[string]AutomationUpdateResponseTriggerWebhookTriggerExampleResultJsonMapItemUnion],
// [[]AutomationUpdateResponseTriggerWebhookTriggerExampleResultJsonArrayItemUnion].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool
// OfAutomationUpdateResponseTriggerWebhookTriggerExampleResultJsonMapItemArray
// OfJsonValueArray OfAnyArray
// OfAutomationUpdateResponseTriggerWebhookTriggerExampleResultJsonArray]
type AutomationUpdateResponseTriggerWebhookTriggerExampleResultJsonUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a
	// [[]AutomationUpdateResponseTriggerWebhookTriggerExampleResultJsonMapItemArrayItemUnion]
	// instead of an object.
	OfAutomationUpdateResponseTriggerWebhookTriggerExampleResultJsonMapItemArray []AutomationUpdateResponseTriggerWebhookTriggerExampleResultJsonMapItemArrayItemUnion `json:",inline"`
	// This field will be present if the value is a [[]shared.JsonValueUnion] instead
	// of an object.
	OfJsonValueArray []shared.JsonValueUnion `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	// This field will be present if the value is a
	// [[]AutomationUpdateResponseTriggerWebhookTriggerExampleResultJsonArrayItemUnion]
	// instead of an object.
	OfAutomationUpdateResponseTriggerWebhookTriggerExampleResultJsonArray []AutomationUpdateResponseTriggerWebhookTriggerExampleResultJsonArrayItemUnion `json:",inline"`
	JSON                                                                  struct {
		OfString                                                                     respjson.Field
		OfFloat                                                                      respjson.Field
		OfBool                                                                       respjson.Field
		OfAutomationUpdateResponseTriggerWebhookTriggerExampleResultJsonMapItemArray respjson.Field
		OfJsonValueArray                                                             respjson.Field
		OfAnyArray                                                                   respjson.Field
		OfAutomationUpdateResponseTriggerWebhookTriggerExampleResultJsonArray        respjson.Field
		raw                                                                          string
	} `json:"-"`
}

func (u AutomationUpdateResponseTriggerWebhookTriggerExampleResultJsonUnion) AsAutomationUpdateResponseTriggerWebhookTriggerExampleResultJsonMapMap() (v map[string]AutomationUpdateResponseTriggerWebhookTriggerExampleResultJsonMapItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomationUpdateResponseTriggerWebhookTriggerExampleResultJsonUnion) AsAutomationUpdateResponseTriggerWebhookTriggerExampleResultJsonArray() (v []AutomationUpdateResponseTriggerWebhookTriggerExampleResultJsonArrayItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AutomationUpdateResponseTriggerWebhookTriggerExampleResultJsonUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *AutomationUpdateResponseTriggerWebhookTriggerExampleResultJsonUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AutomationUpdateResponseTriggerWebhookTriggerExampleResultJsonMapItemUnion
// contains all possible properties and values from [string], [float64], [bool],
// [[]AutomationUpdateResponseTriggerWebhookTriggerExampleResultJsonMapItemArrayItemUnion],
// [map[string]AutomationUpdateResponseTriggerWebhookTriggerExampleResultJsonMapItemMapItemUnion].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool
// OfAutomationUpdateResponseTriggerWebhookTriggerExampleResultJsonMapItemArray
// OfJsonValueArray OfAnyArray]
type AutomationUpdateResponseTriggerWebhookTriggerExampleResultJsonMapItemUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a
	// [[]AutomationUpdateResponseTriggerWebhookTriggerExampleResultJsonMapItemArrayItemUnion]
	// instead of an object.
	OfAutomationUpdateResponseTriggerWebhookTriggerExampleResultJsonMapItemArray []AutomationUpdateResponseTriggerWebhookTriggerExampleResultJsonMapItemArrayItemUnion `json:",inline"`
	// This field will be present if the value is a [[]shared.JsonValueUnion] instead
	// of an object.
	OfJsonValueArray []shared.JsonValueUnion `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	JSON       struct {
		OfString                                                                     respjson.Field
		OfFloat                                                                      respjson.Field
		OfBool                                                                       respjson.Field
		OfAutomationUpdateResponseTriggerWebhookTriggerExampleResultJsonMapItemArray respjson.Field
		OfJsonValueArray                                                             respjson.Field
		OfAnyArray                                                                   respjson.Field
		raw                                                                          string
	} `json:"-"`
}

func (u AutomationUpdateResponseTriggerWebhookTriggerExampleResultJsonMapItemUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomationUpdateResponseTriggerWebhookTriggerExampleResultJsonMapItemUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomationUpdateResponseTriggerWebhookTriggerExampleResultJsonMapItemUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomationUpdateResponseTriggerWebhookTriggerExampleResultJsonMapItemUnion) AsAutomationUpdateResponseTriggerWebhookTriggerExampleResultJsonMapItemArray() (v []AutomationUpdateResponseTriggerWebhookTriggerExampleResultJsonMapItemArrayItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomationUpdateResponseTriggerWebhookTriggerExampleResultJsonMapItemUnion) AsAutomationUpdateResponseTriggerWebhookTriggerExampleResultJsonMapItemMapMap() (v map[string]AutomationUpdateResponseTriggerWebhookTriggerExampleResultJsonMapItemMapItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AutomationUpdateResponseTriggerWebhookTriggerExampleResultJsonMapItemUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *AutomationUpdateResponseTriggerWebhookTriggerExampleResultJsonMapItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AutomationUpdateResponseTriggerWebhookTriggerExampleResultJsonMapItemArrayItemUnion
// contains all possible properties and values from [string], [float64], [bool],
// [[]shared.JsonValueUnion], [map[string]shared.JsonValueUnion].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfJsonValueArray OfAnyArray]
type AutomationUpdateResponseTriggerWebhookTriggerExampleResultJsonMapItemArrayItemUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]shared.JsonValueUnion] instead
	// of an object.
	OfJsonValueArray []shared.JsonValueUnion `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	JSON       struct {
		OfString         respjson.Field
		OfFloat          respjson.Field
		OfBool           respjson.Field
		OfJsonValueArray respjson.Field
		OfAnyArray       respjson.Field
		raw              string
	} `json:"-"`
}

func (u AutomationUpdateResponseTriggerWebhookTriggerExampleResultJsonMapItemArrayItemUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomationUpdateResponseTriggerWebhookTriggerExampleResultJsonMapItemArrayItemUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomationUpdateResponseTriggerWebhookTriggerExampleResultJsonMapItemArrayItemUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomationUpdateResponseTriggerWebhookTriggerExampleResultJsonMapItemArrayItemUnion) AsJsonValueArray() (v []shared.JsonValueUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomationUpdateResponseTriggerWebhookTriggerExampleResultJsonMapItemArrayItemUnion) AsJsonValueMap() (v map[string]shared.JsonValueUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AutomationUpdateResponseTriggerWebhookTriggerExampleResultJsonMapItemArrayItemUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *AutomationUpdateResponseTriggerWebhookTriggerExampleResultJsonMapItemArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AutomationUpdateResponseTriggerWebhookTriggerExampleResultJsonMapItemMapItemUnion
// contains all possible properties and values from [string], [float64], [bool],
// [[]shared.JsonValueUnion], [map[string]shared.JsonValueUnion].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfJsonValueArray OfAnyArray]
type AutomationUpdateResponseTriggerWebhookTriggerExampleResultJsonMapItemMapItemUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]shared.JsonValueUnion] instead
	// of an object.
	OfJsonValueArray []shared.JsonValueUnion `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	JSON       struct {
		OfString         respjson.Field
		OfFloat          respjson.Field
		OfBool           respjson.Field
		OfJsonValueArray respjson.Field
		OfAnyArray       respjson.Field
		raw              string
	} `json:"-"`
}

func (u AutomationUpdateResponseTriggerWebhookTriggerExampleResultJsonMapItemMapItemUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomationUpdateResponseTriggerWebhookTriggerExampleResultJsonMapItemMapItemUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomationUpdateResponseTriggerWebhookTriggerExampleResultJsonMapItemMapItemUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomationUpdateResponseTriggerWebhookTriggerExampleResultJsonMapItemMapItemUnion) AsJsonValueArray() (v []shared.JsonValueUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomationUpdateResponseTriggerWebhookTriggerExampleResultJsonMapItemMapItemUnion) AsJsonValueMap() (v map[string]shared.JsonValueUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AutomationUpdateResponseTriggerWebhookTriggerExampleResultJsonMapItemMapItemUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *AutomationUpdateResponseTriggerWebhookTriggerExampleResultJsonMapItemMapItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AutomationUpdateResponseTriggerWebhookTriggerExampleResultJsonArrayItemUnion
// contains all possible properties and values from [string], [float64], [bool],
// [[]AutomationUpdateResponseTriggerWebhookTriggerExampleResultJsonArrayItemArrayItemUnion],
// [map[string]AutomationUpdateResponseTriggerWebhookTriggerExampleResultJsonArrayItemMapItemUnion].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool
// OfAutomationUpdateResponseTriggerWebhookTriggerExampleResultJsonArrayItemArray
// OfJsonValueArray OfAnyArray]
type AutomationUpdateResponseTriggerWebhookTriggerExampleResultJsonArrayItemUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a
	// [[]AutomationUpdateResponseTriggerWebhookTriggerExampleResultJsonArrayItemArrayItemUnion]
	// instead of an object.
	OfAutomationUpdateResponseTriggerWebhookTriggerExampleResultJsonArrayItemArray []AutomationUpdateResponseTriggerWebhookTriggerExampleResultJsonArrayItemArrayItemUnion `json:",inline"`
	// This field will be present if the value is a [[]shared.JsonValueUnion] instead
	// of an object.
	OfJsonValueArray []shared.JsonValueUnion `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	JSON       struct {
		OfString                                                                       respjson.Field
		OfFloat                                                                        respjson.Field
		OfBool                                                                         respjson.Field
		OfAutomationUpdateResponseTriggerWebhookTriggerExampleResultJsonArrayItemArray respjson.Field
		OfJsonValueArray                                                               respjson.Field
		OfAnyArray                                                                     respjson.Field
		raw                                                                            string
	} `json:"-"`
}

func (u AutomationUpdateResponseTriggerWebhookTriggerExampleResultJsonArrayItemUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomationUpdateResponseTriggerWebhookTriggerExampleResultJsonArrayItemUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomationUpdateResponseTriggerWebhookTriggerExampleResultJsonArrayItemUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomationUpdateResponseTriggerWebhookTriggerExampleResultJsonArrayItemUnion) AsAutomationUpdateResponseTriggerWebhookTriggerExampleResultJsonArrayItemArray() (v []AutomationUpdateResponseTriggerWebhookTriggerExampleResultJsonArrayItemArrayItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomationUpdateResponseTriggerWebhookTriggerExampleResultJsonArrayItemUnion) AsAutomationUpdateResponseTriggerWebhookTriggerExampleResultJsonArrayItemMapMap() (v map[string]AutomationUpdateResponseTriggerWebhookTriggerExampleResultJsonArrayItemMapItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AutomationUpdateResponseTriggerWebhookTriggerExampleResultJsonArrayItemUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *AutomationUpdateResponseTriggerWebhookTriggerExampleResultJsonArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AutomationUpdateResponseTriggerWebhookTriggerExampleResultJsonArrayItemArrayItemUnion
// contains all possible properties and values from [string], [float64], [bool],
// [[]shared.JsonValueUnion], [map[string]shared.JsonValueUnion].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfJsonValueArray OfAnyArray]
type AutomationUpdateResponseTriggerWebhookTriggerExampleResultJsonArrayItemArrayItemUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]shared.JsonValueUnion] instead
	// of an object.
	OfJsonValueArray []shared.JsonValueUnion `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	JSON       struct {
		OfString         respjson.Field
		OfFloat          respjson.Field
		OfBool           respjson.Field
		OfJsonValueArray respjson.Field
		OfAnyArray       respjson.Field
		raw              string
	} `json:"-"`
}

func (u AutomationUpdateResponseTriggerWebhookTriggerExampleResultJsonArrayItemArrayItemUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomationUpdateResponseTriggerWebhookTriggerExampleResultJsonArrayItemArrayItemUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomationUpdateResponseTriggerWebhookTriggerExampleResultJsonArrayItemArrayItemUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomationUpdateResponseTriggerWebhookTriggerExampleResultJsonArrayItemArrayItemUnion) AsJsonValueArray() (v []shared.JsonValueUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomationUpdateResponseTriggerWebhookTriggerExampleResultJsonArrayItemArrayItemUnion) AsJsonValueMap() (v map[string]shared.JsonValueUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AutomationUpdateResponseTriggerWebhookTriggerExampleResultJsonArrayItemArrayItemUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *AutomationUpdateResponseTriggerWebhookTriggerExampleResultJsonArrayItemArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AutomationUpdateResponseTriggerWebhookTriggerExampleResultJsonArrayItemMapItemUnion
// contains all possible properties and values from [string], [float64], [bool],
// [[]shared.JsonValueUnion], [map[string]shared.JsonValueUnion].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfJsonValueArray OfAnyArray]
type AutomationUpdateResponseTriggerWebhookTriggerExampleResultJsonArrayItemMapItemUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]shared.JsonValueUnion] instead
	// of an object.
	OfJsonValueArray []shared.JsonValueUnion `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	JSON       struct {
		OfString         respjson.Field
		OfFloat          respjson.Field
		OfBool           respjson.Field
		OfJsonValueArray respjson.Field
		OfAnyArray       respjson.Field
		raw              string
	} `json:"-"`
}

func (u AutomationUpdateResponseTriggerWebhookTriggerExampleResultJsonArrayItemMapItemUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomationUpdateResponseTriggerWebhookTriggerExampleResultJsonArrayItemMapItemUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomationUpdateResponseTriggerWebhookTriggerExampleResultJsonArrayItemMapItemUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomationUpdateResponseTriggerWebhookTriggerExampleResultJsonArrayItemMapItemUnion) AsJsonValueArray() (v []shared.JsonValueUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomationUpdateResponseTriggerWebhookTriggerExampleResultJsonArrayItemMapItemUnion) AsJsonValueMap() (v map[string]shared.JsonValueUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AutomationUpdateResponseTriggerWebhookTriggerExampleResultJsonArrayItemMapItemUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *AutomationUpdateResponseTriggerWebhookTriggerExampleResultJsonArrayItemMapItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AutomationUpdateResponseTriggerWebhookTriggerExampleResultParsed struct {
	Path  []string                                                                   `json:"path" api:"required"`
	Label string                                                                     `json:"label"`
	Value AutomationUpdateResponseTriggerWebhookTriggerExampleResultParsedValueUnion `json:"value" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Path        respjson.Field
		Label       respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutomationUpdateResponseTriggerWebhookTriggerExampleResultParsed) RawJSON() string {
	return r.JSON.raw
}
func (r *AutomationUpdateResponseTriggerWebhookTriggerExampleResultParsed) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AutomationUpdateResponseTriggerWebhookTriggerExampleResultParsedValueUnion
// contains all possible properties and values from [string], [float64], [bool],
// [[]AutomationUpdateResponseTriggerWebhookTriggerExampleResultParsedValueArrayItemUnion],
// [map[string]AutomationUpdateResponseTriggerWebhookTriggerExampleResultParsedValueMapItemUnion].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool
// OfAutomationUpdateResponseTriggerWebhookTriggerExampleResultParsedValueArray
// OfJsonValueArray OfAnyArray]
type AutomationUpdateResponseTriggerWebhookTriggerExampleResultParsedValueUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a
	// [[]AutomationUpdateResponseTriggerWebhookTriggerExampleResultParsedValueArrayItemUnion]
	// instead of an object.
	OfAutomationUpdateResponseTriggerWebhookTriggerExampleResultParsedValueArray []AutomationUpdateResponseTriggerWebhookTriggerExampleResultParsedValueArrayItemUnion `json:",inline"`
	// This field will be present if the value is a [[]shared.JsonValueUnion] instead
	// of an object.
	OfJsonValueArray []shared.JsonValueUnion `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	JSON       struct {
		OfString                                                                     respjson.Field
		OfFloat                                                                      respjson.Field
		OfBool                                                                       respjson.Field
		OfAutomationUpdateResponseTriggerWebhookTriggerExampleResultParsedValueArray respjson.Field
		OfJsonValueArray                                                             respjson.Field
		OfAnyArray                                                                   respjson.Field
		raw                                                                          string
	} `json:"-"`
}

func (u AutomationUpdateResponseTriggerWebhookTriggerExampleResultParsedValueUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomationUpdateResponseTriggerWebhookTriggerExampleResultParsedValueUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomationUpdateResponseTriggerWebhookTriggerExampleResultParsedValueUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomationUpdateResponseTriggerWebhookTriggerExampleResultParsedValueUnion) AsAutomationUpdateResponseTriggerWebhookTriggerExampleResultParsedValueArray() (v []AutomationUpdateResponseTriggerWebhookTriggerExampleResultParsedValueArrayItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomationUpdateResponseTriggerWebhookTriggerExampleResultParsedValueUnion) AsAutomationUpdateResponseTriggerWebhookTriggerExampleResultParsedValueMapMap() (v map[string]AutomationUpdateResponseTriggerWebhookTriggerExampleResultParsedValueMapItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AutomationUpdateResponseTriggerWebhookTriggerExampleResultParsedValueUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *AutomationUpdateResponseTriggerWebhookTriggerExampleResultParsedValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AutomationUpdateResponseTriggerWebhookTriggerExampleResultParsedValueArrayItemUnion
// contains all possible properties and values from [string], [float64], [bool],
// [[]shared.JsonValueUnion], [map[string]shared.JsonValueUnion].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfJsonValueArray OfAnyArray]
type AutomationUpdateResponseTriggerWebhookTriggerExampleResultParsedValueArrayItemUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]shared.JsonValueUnion] instead
	// of an object.
	OfJsonValueArray []shared.JsonValueUnion `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	JSON       struct {
		OfString         respjson.Field
		OfFloat          respjson.Field
		OfBool           respjson.Field
		OfJsonValueArray respjson.Field
		OfAnyArray       respjson.Field
		raw              string
	} `json:"-"`
}

func (u AutomationUpdateResponseTriggerWebhookTriggerExampleResultParsedValueArrayItemUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomationUpdateResponseTriggerWebhookTriggerExampleResultParsedValueArrayItemUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomationUpdateResponseTriggerWebhookTriggerExampleResultParsedValueArrayItemUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomationUpdateResponseTriggerWebhookTriggerExampleResultParsedValueArrayItemUnion) AsJsonValueArray() (v []shared.JsonValueUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomationUpdateResponseTriggerWebhookTriggerExampleResultParsedValueArrayItemUnion) AsJsonValueMap() (v map[string]shared.JsonValueUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AutomationUpdateResponseTriggerWebhookTriggerExampleResultParsedValueArrayItemUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *AutomationUpdateResponseTriggerWebhookTriggerExampleResultParsedValueArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AutomationUpdateResponseTriggerWebhookTriggerExampleResultParsedValueMapItemUnion
// contains all possible properties and values from [string], [float64], [bool],
// [[]shared.JsonValueUnion], [map[string]shared.JsonValueUnion].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfJsonValueArray OfAnyArray]
type AutomationUpdateResponseTriggerWebhookTriggerExampleResultParsedValueMapItemUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]shared.JsonValueUnion] instead
	// of an object.
	OfJsonValueArray []shared.JsonValueUnion `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	JSON       struct {
		OfString         respjson.Field
		OfFloat          respjson.Field
		OfBool           respjson.Field
		OfJsonValueArray respjson.Field
		OfAnyArray       respjson.Field
		raw              string
	} `json:"-"`
}

func (u AutomationUpdateResponseTriggerWebhookTriggerExampleResultParsedValueMapItemUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomationUpdateResponseTriggerWebhookTriggerExampleResultParsedValueMapItemUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomationUpdateResponseTriggerWebhookTriggerExampleResultParsedValueMapItemUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomationUpdateResponseTriggerWebhookTriggerExampleResultParsedValueMapItemUnion) AsJsonValueArray() (v []shared.JsonValueUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomationUpdateResponseTriggerWebhookTriggerExampleResultParsedValueMapItemUnion) AsJsonValueMap() (v map[string]shared.JsonValueUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AutomationUpdateResponseTriggerWebhookTriggerExampleResultParsedValueMapItemUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *AutomationUpdateResponseTriggerWebhookTriggerExampleResultParsedValueMapItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AutomationUpdateResponseTriggerScheduledTrigger struct {
	Schedule AutomationUpdateResponseTriggerScheduledTriggerSchedule `json:"schedule" api:"required"`
	// Any of "SCHEDULED".
	Type string `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Schedule    respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutomationUpdateResponseTriggerScheduledTrigger) RawJSON() string { return r.JSON.raw }
func (r *AutomationUpdateResponseTriggerScheduledTrigger) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AutomationUpdateResponseTriggerScheduledTriggerSchedule struct {
	// 5-field cron expression, for example `0 9 * * 1-5`.
	Cron string `json:"cron" api:"required"`
	// IANA timezone the cron is evaluated in.
	Timezone string `json:"timezone" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Cron        respjson.Field
		Timezone    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutomationUpdateResponseTriggerScheduledTriggerSchedule) RawJSON() string { return r.JSON.raw }
func (r *AutomationUpdateResponseTriggerScheduledTriggerSchedule) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AutomationUpdateResponseTriggerRecordCreatedTrigger struct {
	ObjectType string `json:"objectType" api:"required"`
	// Any of "RECORD_CREATED".
	Type string `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ObjectType  respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutomationUpdateResponseTriggerRecordCreatedTrigger) RawJSON() string { return r.JSON.raw }
func (r *AutomationUpdateResponseTriggerRecordCreatedTrigger) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AutomationUpdateResponseTriggerRecordUpdatedTrigger struct {
	ObjectType string `json:"objectType" api:"required"`
	// Any of "RECORD_UPDATED".
	Type                 string                                                                   `json:"type" api:"required"`
	WatchedFields        []string                                                                 `json:"watchedFields" api:"required"`
	WatchedRelationships []AutomationUpdateResponseTriggerRecordUpdatedTriggerWatchedRelationship `json:"watchedRelationships"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ObjectType           respjson.Field
		Type                 respjson.Field
		WatchedFields        respjson.Field
		WatchedRelationships respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutomationUpdateResponseTriggerRecordUpdatedTrigger) RawJSON() string { return r.JSON.raw }
func (r *AutomationUpdateResponseTriggerRecordUpdatedTrigger) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AutomationUpdateResponseTriggerRecordUpdatedTriggerWatchedRelationship struct {
	// The endpoint occupied by the record whose changes start the run.
	//
	// Any of "SOURCE", "DESTINATION".
	Direction string `json:"direction" api:"required"`
	// Public relationship key, such as `$contact` or a custom relationship slug.
	Key string `json:"key" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Direction   respjson.Field
		Key         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutomationUpdateResponseTriggerRecordUpdatedTriggerWatchedRelationship) RawJSON() string {
	return r.JSON.raw
}
func (r *AutomationUpdateResponseTriggerRecordUpdatedTriggerWatchedRelationship) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AutomationUpdateResponseTriggerSequenceReplyTrigger struct {
	// Any of "SEQUENCE_REPLY".
	Type string `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutomationUpdateResponseTriggerSequenceReplyTrigger) RawJSON() string { return r.JSON.raw }
func (r *AutomationUpdateResponseTriggerSequenceReplyTrigger) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AutomationNewParams struct {
	// Display name of the automation.
	Name string `json:"name" api:"required"`
	// What a run does, in order.
	Steps []AutomationNewParamsStepUnion `json:"steps,omitzero" api:"required"`
	// What starts a run.
	Triggers []AutomationNewParamsTriggerUnion `json:"triggers,omitzero" api:"required"`
	// What the automation is meant to accomplish.
	Description param.Opt[string] `json:"description,omitzero"`
	// Model the agent steps run on. Returns 400 for an unsupported model and 422 when
	// the workspace cannot choose the requested model; omit `model` or send `auto`.
	//
	// Any of "auto", "auto_low", "auto_high", "sonnet_5_5", "haiku_5_5",
	// "gpt_6_1_terra", "gpt_6_luna".
	Model AutomationNewParamsModel `json:"model,omitzero"`
	// Who the automation acts as. Workspace admins may name only themselves, a member,
	// or `WORKSPACE`; only a workspace owner may name another admin or an owner.
	RunAs AutomationNewParamsRunAsUnion `json:"runAs,omitzero"`
	paramObj
}

func (r AutomationNewParams) MarshalJSON() (data []byte, err error) {
	type shadow AutomationNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AutomationNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type AutomationNewParamsStepUnion struct {
	OfAgentStep *AutomationNewParamsStepAgentStep `json:",omitzero,inline"`
	OfCodeStep  *AutomationNewParamsStepCodeStep  `json:",omitzero,inline"`
	OfSleepStep *AutomationNewParamsStepSleepStep `json:",omitzero,inline"`
	paramUnion
}

func (u AutomationNewParamsStepUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfAgentStep, u.OfCodeStep, u.OfSleepStep)
}
func (u *AutomationNewParamsStepUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// The properties Name, Prompt, Type are required.
type AutomationNewParamsStepAgentStep struct {
	// Display name of the step.
	Name string `json:"name" api:"required"`
	// Instructions the agent follows in this step.
	Prompt string `json:"prompt" api:"required"`
	// Any of "AGENT".
	Type string `json:"type,omitzero" api:"required"`
	// What the agent may use in this step. Defaults to none.
	Permissions []AutomationNewParamsStepAgentStepPermissionUnion `json:"permissions,omitzero"`
	paramObj
}

func (r AutomationNewParamsStepAgentStep) MarshalJSON() (data []byte, err error) {
	type shadow AutomationNewParamsStepAgentStep
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AutomationNewParamsStepAgentStep) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[AutomationNewParamsStepAgentStep](
		"type", "AGENT",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type AutomationNewParamsStepAgentStepPermissionUnion struct {
	OfToolPermission    *AutomationNewParamsStepAgentStepPermissionToolPermission    `json:",omitzero,inline"`
	OfMcpPermission     *AutomationNewParamsStepAgentStepPermissionMcpPermission     `json:",omitzero,inline"`
	OfAPIPermission     *AutomationNewParamsStepAgentStepPermissionAPIPermission     `json:",omitzero,inline"`
	OfNetworkPermission *AutomationNewParamsStepAgentStepPermissionNetworkPermission `json:",omitzero,inline"`
	paramUnion
}

func (u AutomationNewParamsStepAgentStepPermissionUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfToolPermission, u.OfMcpPermission, u.OfAPIPermission, u.OfNetworkPermission)
}
func (u *AutomationNewParamsStepAgentStepPermissionUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// The properties Name, Type are required.
type AutomationNewParamsStepAgentStepPermissionToolPermission struct {
	// Name of the tool, matching the Lightfield app.
	Name string `json:"name" api:"required"`
	// Any of "TOOL".
	Type string `json:"type,omitzero" api:"required"`
	// Why the step needs this permission. Null when none was given. Returns 422 for an
	// empty or whitespace-only string.
	Reason param.Opt[string] `json:"reason,omitzero"`
	paramObj
}

func (r AutomationNewParamsStepAgentStepPermissionToolPermission) MarshalJSON() (data []byte, err error) {
	type shadow AutomationNewParamsStepAgentStepPermissionToolPermission
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AutomationNewParamsStepAgentStepPermissionToolPermission) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[AutomationNewParamsStepAgentStepPermissionToolPermission](
		"type", "TOOL",
	)
}

// The properties Connector, Tools, Type are required.
type AutomationNewParamsStepAgentStepPermissionMcpPermission struct {
	// Connector the tools are on. Must already be lowercase and trimmed.
	Connector string `json:"connector" api:"required"`
	// Unique, trimmed names of the connector tools the step may call.
	Tools []string `json:"tools,omitzero" api:"required"`
	// Any of "MCP".
	Type string `json:"type,omitzero" api:"required"`
	// Why the step needs this permission. Null when none was given. Returns 422 for an
	// empty or whitespace-only string.
	Reason param.Opt[string] `json:"reason,omitzero"`
	paramObj
}

func (r AutomationNewParamsStepAgentStepPermissionMcpPermission) MarshalJSON() (data []byte, err error) {
	type shadow AutomationNewParamsStepAgentStepPermissionMcpPermission
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AutomationNewParamsStepAgentStepPermissionMcpPermission) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[AutomationNewParamsStepAgentStepPermissionMcpPermission](
		"type", "MCP",
	)
}

// The properties Scopes, Type are required.
type AutomationNewParamsStepAgentStepPermissionAPIPermission struct {
	// Lightfield API scopes the step may use, for example `contacts:read`. The API key
	// must hold every scope it grants, and only a full-access key may grant `*`.
	Scopes []string `json:"scopes,omitzero" api:"required"`
	// Any of "API".
	Type string `json:"type,omitzero" api:"required"`
	// Why the step needs this permission. Null when none was given. Returns 422 for an
	// empty or whitespace-only string.
	Reason param.Opt[string] `json:"reason,omitzero"`
	paramObj
}

func (r AutomationNewParamsStepAgentStepPermissionAPIPermission) MarshalJSON() (data []byte, err error) {
	type shadow AutomationNewParamsStepAgentStepPermissionAPIPermission
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AutomationNewParamsStepAgentStepPermissionAPIPermission) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[AutomationNewParamsStepAgentStepPermissionAPIPermission](
		"type", "API",
	)
}

// The properties Domain, Type are required.
type AutomationNewParamsStepAgentStepPermissionNetworkPermission struct {
	// Lowercase host the step may call, optionally followed by a path glob. Returns
	// 422 for a scheme, query, fragment or surrounding whitespace.
	Domain string `json:"domain" api:"required"`
	// Any of "NETWORK".
	Type string `json:"type,omitzero" api:"required"`
	// Why the step needs this permission. Null when none was given. Returns 422 for an
	// empty or whitespace-only string.
	Reason param.Opt[string] `json:"reason,omitzero"`
	// Headers added to every request this permission allows.
	Headers map[string]string `json:"headers,omitzero"`
	// Restricts the permission to one HTTP method.
	//
	// Any of "GET", "POST", "PUT", "PATCH", "DELETE".
	Method string `json:"method,omitzero"`
	paramObj
}

func (r AutomationNewParamsStepAgentStepPermissionNetworkPermission) MarshalJSON() (data []byte, err error) {
	type shadow AutomationNewParamsStepAgentStepPermissionNetworkPermission
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AutomationNewParamsStepAgentStepPermissionNetworkPermission) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[AutomationNewParamsStepAgentStepPermissionNetworkPermission](
		"type", "NETWORK",
	)
	apijson.RegisterFieldValidator[AutomationNewParamsStepAgentStepPermissionNetworkPermission](
		"method", "GET", "POST", "PUT", "PATCH", "DELETE",
	)
}

// The properties Command, Name, Type are required.
type AutomationNewParamsStepCodeStep struct {
	// Bash command the step runs as written.
	Command string `json:"command" api:"required"`
	// Display name of the step.
	Name string `json:"name" api:"required"`
	// Any of "CODE".
	Type string `json:"type,omitzero" api:"required"`
	// What the command does. Does not change what runs.
	Description param.Opt[string] `json:"description,omitzero"`
	// What the command may use. Defaults to none.
	Permissions []AutomationNewParamsStepCodeStepPermissionUnion `json:"permissions,omitzero"`
	paramObj
}

func (r AutomationNewParamsStepCodeStep) MarshalJSON() (data []byte, err error) {
	type shadow AutomationNewParamsStepCodeStep
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AutomationNewParamsStepCodeStep) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[AutomationNewParamsStepCodeStep](
		"type", "CODE",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type AutomationNewParamsStepCodeStepPermissionUnion struct {
	OfAPIPermission     *AutomationNewParamsStepCodeStepPermissionAPIPermission     `json:",omitzero,inline"`
	OfNetworkPermission *AutomationNewParamsStepCodeStepPermissionNetworkPermission `json:",omitzero,inline"`
	paramUnion
}

func (u AutomationNewParamsStepCodeStepPermissionUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfAPIPermission, u.OfNetworkPermission)
}
func (u *AutomationNewParamsStepCodeStepPermissionUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// The properties Scopes, Type are required.
type AutomationNewParamsStepCodeStepPermissionAPIPermission struct {
	// Lightfield API scopes the step may use, for example `contacts:read`. The API key
	// must hold every scope it grants, and only a full-access key may grant `*`.
	Scopes []string `json:"scopes,omitzero" api:"required"`
	// Any of "API".
	Type string `json:"type,omitzero" api:"required"`
	// Why the step needs this permission. Null when none was given. Returns 422 for an
	// empty or whitespace-only string.
	Reason param.Opt[string] `json:"reason,omitzero"`
	paramObj
}

func (r AutomationNewParamsStepCodeStepPermissionAPIPermission) MarshalJSON() (data []byte, err error) {
	type shadow AutomationNewParamsStepCodeStepPermissionAPIPermission
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AutomationNewParamsStepCodeStepPermissionAPIPermission) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[AutomationNewParamsStepCodeStepPermissionAPIPermission](
		"type", "API",
	)
}

// The properties Domain, Type are required.
type AutomationNewParamsStepCodeStepPermissionNetworkPermission struct {
	// Lowercase host the step may call, optionally followed by a path glob. Returns
	// 422 for a scheme, query, fragment or surrounding whitespace.
	Domain string `json:"domain" api:"required"`
	// Any of "NETWORK".
	Type string `json:"type,omitzero" api:"required"`
	// Why the step needs this permission. Null when none was given. Returns 422 for an
	// empty or whitespace-only string.
	Reason param.Opt[string] `json:"reason,omitzero"`
	// Headers added to every request this permission allows.
	Headers map[string]string `json:"headers,omitzero"`
	// Restricts the permission to one HTTP method.
	//
	// Any of "GET", "POST", "PUT", "PATCH", "DELETE".
	Method string `json:"method,omitzero"`
	paramObj
}

func (r AutomationNewParamsStepCodeStepPermissionNetworkPermission) MarshalJSON() (data []byte, err error) {
	type shadow AutomationNewParamsStepCodeStepPermissionNetworkPermission
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AutomationNewParamsStepCodeStepPermissionNetworkPermission) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[AutomationNewParamsStepCodeStepPermissionNetworkPermission](
		"type", "NETWORK",
	)
	apijson.RegisterFieldValidator[AutomationNewParamsStepCodeStepPermissionNetworkPermission](
		"method", "GET", "POST", "PUT", "PATCH", "DELETE",
	)
}

// The properties Duration, Name, Type, Unit are required.
type AutomationNewParamsStepSleepStep struct {
	// How long to sleep before the next step, at most one year.
	Duration AutomationNewParamsStepSleepStepDurationUnion `json:"duration,omitzero" api:"required"`
	// Display name of the step.
	Name string `json:"name" api:"required"`
	// Any of "SLEEP".
	Type string `json:"type,omitzero" api:"required"`
	// Unit of `duration`.
	//
	// Any of "MILLISECOND", "SECOND", "MINUTE", "HOUR", "DAY".
	Unit string `json:"unit,omitzero" api:"required"`
	paramObj
}

func (r AutomationNewParamsStepSleepStep) MarshalJSON() (data []byte, err error) {
	type shadow AutomationNewParamsStepSleepStep
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AutomationNewParamsStepSleepStep) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[AutomationNewParamsStepSleepStep](
		"type", "SLEEP",
	)
	apijson.RegisterFieldValidator[AutomationNewParamsStepSleepStep](
		"unit", "MILLISECOND", "SECOND", "MINUTE", "HOUR", "DAY",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type AutomationNewParamsStepSleepStepDurationUnion struct {
	OfInt    param.Opt[int64]  `json:",omitzero,inline"`
	OfString param.Opt[string] `json:",omitzero,inline"`
	paramUnion
}

func (u AutomationNewParamsStepSleepStepDurationUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfInt, u.OfString)
}
func (u *AutomationNewParamsStepSleepStepDurationUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type AutomationNewParamsTriggerUnion struct {
	OfManualTrigger        *AutomationNewParamsTriggerManualTrigger        `json:",omitzero,inline"`
	OfWebhookTrigger       *AutomationNewParamsTriggerWebhookTrigger       `json:",omitzero,inline"`
	OfScheduledTrigger     *AutomationNewParamsTriggerScheduledTrigger     `json:",omitzero,inline"`
	OfRecordCreatedTrigger *AutomationNewParamsTriggerRecordCreatedTrigger `json:",omitzero,inline"`
	OfRecordUpdatedTrigger *AutomationNewParamsTriggerRecordUpdatedTrigger `json:",omitzero,inline"`
	OfSequenceReplyTrigger *AutomationNewParamsTriggerSequenceReplyTrigger `json:",omitzero,inline"`
	paramUnion
}

func (u AutomationNewParamsTriggerUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfManualTrigger,
		u.OfWebhookTrigger,
		u.OfScheduledTrigger,
		u.OfRecordCreatedTrigger,
		u.OfRecordUpdatedTrigger,
		u.OfSequenceReplyTrigger)
}
func (u *AutomationNewParamsTriggerUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// The property Type is required.
type AutomationNewParamsTriggerManualTrigger struct {
	// Any of "MANUAL".
	Type string `json:"type,omitzero" api:"required"`
	paramObj
}

func (r AutomationNewParamsTriggerManualTrigger) MarshalJSON() (data []byte, err error) {
	type shadow AutomationNewParamsTriggerManualTrigger
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AutomationNewParamsTriggerManualTrigger) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[AutomationNewParamsTriggerManualTrigger](
		"type", "MANUAL",
	)
}

// The property Type is required.
type AutomationNewParamsTriggerWebhookTrigger struct {
	// Any of "WEBHOOK".
	Type string `json:"type,omitzero" api:"required"`
	// A sample payload and the fields the steps read from it.
	ExampleResult AutomationNewParamsTriggerWebhookTriggerExampleResult `json:"exampleResult,omitzero"`
	paramObj
}

func (r AutomationNewParamsTriggerWebhookTrigger) MarshalJSON() (data []byte, err error) {
	type shadow AutomationNewParamsTriggerWebhookTrigger
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AutomationNewParamsTriggerWebhookTrigger) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[AutomationNewParamsTriggerWebhookTrigger](
		"type", "WEBHOOK",
	)
}

// A sample payload and the fields the steps read from it.
//
// The properties Json, Parsed are required.
type AutomationNewParamsTriggerWebhookTriggerExampleResult struct {
	Json   AutomationNewParamsTriggerWebhookTriggerExampleResultJsonUnion `json:"json,omitzero" api:"required"`
	Parsed []AutomationNewParamsTriggerWebhookTriggerExampleResultParsed  `json:"parsed,omitzero" api:"required"`
	paramObj
}

func (r AutomationNewParamsTriggerWebhookTriggerExampleResult) MarshalJSON() (data []byte, err error) {
	type shadow AutomationNewParamsTriggerWebhookTriggerExampleResult
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AutomationNewParamsTriggerWebhookTriggerExampleResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type AutomationNewParamsTriggerWebhookTriggerExampleResultJsonUnion struct {
	OfAutomationNewsTriggerWebhookTriggerExampleResultJsonMapMap map[string]AutomationNewParamsTriggerWebhookTriggerExampleResultJsonMapItemUnion `json:",omitzero,inline"`
	OfAutomationNewsTriggerWebhookTriggerExampleResultJsonArray  []AutomationNewParamsTriggerWebhookTriggerExampleResultJsonArrayItemUnion        `json:",omitzero,inline"`
	paramUnion
}

func (u AutomationNewParamsTriggerWebhookTriggerExampleResultJsonUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfAutomationNewsTriggerWebhookTriggerExampleResultJsonMapMap, u.OfAutomationNewsTriggerWebhookTriggerExampleResultJsonArray)
}
func (u *AutomationNewParamsTriggerWebhookTriggerExampleResultJsonUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type AutomationNewParamsTriggerWebhookTriggerExampleResultJsonMapItemUnion struct {
	OfString                                                            param.Opt[string]                                                                       `json:",omitzero,inline"`
	OfFloat                                                             param.Opt[float64]                                                                      `json:",omitzero,inline"`
	OfBool                                                              param.Opt[bool]                                                                         `json:",omitzero,inline"`
	OfAutomationNewsTriggerWebhookTriggerExampleResultJsonMapItemArray  []AutomationNewParamsTriggerWebhookTriggerExampleResultJsonMapItemArrayItemUnion        `json:",omitzero,inline"`
	OfAutomationNewsTriggerWebhookTriggerExampleResultJsonMapItemMapMap map[string]AutomationNewParamsTriggerWebhookTriggerExampleResultJsonMapItemMapItemUnion `json:",omitzero,inline"`
	paramUnion
}

func (u AutomationNewParamsTriggerWebhookTriggerExampleResultJsonMapItemUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString,
		u.OfFloat,
		u.OfBool,
		u.OfAutomationNewsTriggerWebhookTriggerExampleResultJsonMapItemArray,
		u.OfAutomationNewsTriggerWebhookTriggerExampleResultJsonMapItemMapMap)
}
func (u *AutomationNewParamsTriggerWebhookTriggerExampleResultJsonMapItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type AutomationNewParamsTriggerWebhookTriggerExampleResultJsonMapItemArrayItemUnion struct {
	OfString         param.Opt[string]                     `json:",omitzero,inline"`
	OfFloat          param.Opt[float64]                    `json:",omitzero,inline"`
	OfBool           param.Opt[bool]                       `json:",omitzero,inline"`
	OfJsonValueArray []shared.JsonValueUnionParam          `json:",omitzero,inline"`
	OfJsonValueMap   map[string]shared.JsonValueUnionParam `json:",omitzero,inline"`
	paramUnion
}

func (u AutomationNewParamsTriggerWebhookTriggerExampleResultJsonMapItemArrayItemUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString,
		u.OfFloat,
		u.OfBool,
		u.OfJsonValueArray,
		u.OfJsonValueMap)
}
func (u *AutomationNewParamsTriggerWebhookTriggerExampleResultJsonMapItemArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type AutomationNewParamsTriggerWebhookTriggerExampleResultJsonMapItemMapItemUnion struct {
	OfString         param.Opt[string]                     `json:",omitzero,inline"`
	OfFloat          param.Opt[float64]                    `json:",omitzero,inline"`
	OfBool           param.Opt[bool]                       `json:",omitzero,inline"`
	OfJsonValueArray []shared.JsonValueUnionParam          `json:",omitzero,inline"`
	OfJsonValueMap   map[string]shared.JsonValueUnionParam `json:",omitzero,inline"`
	paramUnion
}

func (u AutomationNewParamsTriggerWebhookTriggerExampleResultJsonMapItemMapItemUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString,
		u.OfFloat,
		u.OfBool,
		u.OfJsonValueArray,
		u.OfJsonValueMap)
}
func (u *AutomationNewParamsTriggerWebhookTriggerExampleResultJsonMapItemMapItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type AutomationNewParamsTriggerWebhookTriggerExampleResultJsonArrayItemUnion struct {
	OfString                                                              param.Opt[string]                                                                         `json:",omitzero,inline"`
	OfFloat                                                               param.Opt[float64]                                                                        `json:",omitzero,inline"`
	OfBool                                                                param.Opt[bool]                                                                           `json:",omitzero,inline"`
	OfAutomationNewsTriggerWebhookTriggerExampleResultJsonArrayItemArray  []AutomationNewParamsTriggerWebhookTriggerExampleResultJsonArrayItemArrayItemUnion        `json:",omitzero,inline"`
	OfAutomationNewsTriggerWebhookTriggerExampleResultJsonArrayItemMapMap map[string]AutomationNewParamsTriggerWebhookTriggerExampleResultJsonArrayItemMapItemUnion `json:",omitzero,inline"`
	paramUnion
}

func (u AutomationNewParamsTriggerWebhookTriggerExampleResultJsonArrayItemUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString,
		u.OfFloat,
		u.OfBool,
		u.OfAutomationNewsTriggerWebhookTriggerExampleResultJsonArrayItemArray,
		u.OfAutomationNewsTriggerWebhookTriggerExampleResultJsonArrayItemMapMap)
}
func (u *AutomationNewParamsTriggerWebhookTriggerExampleResultJsonArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type AutomationNewParamsTriggerWebhookTriggerExampleResultJsonArrayItemArrayItemUnion struct {
	OfString         param.Opt[string]                     `json:",omitzero,inline"`
	OfFloat          param.Opt[float64]                    `json:",omitzero,inline"`
	OfBool           param.Opt[bool]                       `json:",omitzero,inline"`
	OfJsonValueArray []shared.JsonValueUnionParam          `json:",omitzero,inline"`
	OfJsonValueMap   map[string]shared.JsonValueUnionParam `json:",omitzero,inline"`
	paramUnion
}

func (u AutomationNewParamsTriggerWebhookTriggerExampleResultJsonArrayItemArrayItemUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString,
		u.OfFloat,
		u.OfBool,
		u.OfJsonValueArray,
		u.OfJsonValueMap)
}
func (u *AutomationNewParamsTriggerWebhookTriggerExampleResultJsonArrayItemArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type AutomationNewParamsTriggerWebhookTriggerExampleResultJsonArrayItemMapItemUnion struct {
	OfString         param.Opt[string]                     `json:",omitzero,inline"`
	OfFloat          param.Opt[float64]                    `json:",omitzero,inline"`
	OfBool           param.Opt[bool]                       `json:",omitzero,inline"`
	OfJsonValueArray []shared.JsonValueUnionParam          `json:",omitzero,inline"`
	OfJsonValueMap   map[string]shared.JsonValueUnionParam `json:",omitzero,inline"`
	paramUnion
}

func (u AutomationNewParamsTriggerWebhookTriggerExampleResultJsonArrayItemMapItemUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString,
		u.OfFloat,
		u.OfBool,
		u.OfJsonValueArray,
		u.OfJsonValueMap)
}
func (u *AutomationNewParamsTriggerWebhookTriggerExampleResultJsonArrayItemMapItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// The property Path is required.
type AutomationNewParamsTriggerWebhookTriggerExampleResultParsed struct {
	Path  []string                                                              `json:"path,omitzero" api:"required"`
	Label param.Opt[string]                                                     `json:"label,omitzero"`
	Value AutomationNewParamsTriggerWebhookTriggerExampleResultParsedValueUnion `json:"value,omitzero"`
	paramObj
}

func (r AutomationNewParamsTriggerWebhookTriggerExampleResultParsed) MarshalJSON() (data []byte, err error) {
	type shadow AutomationNewParamsTriggerWebhookTriggerExampleResultParsed
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AutomationNewParamsTriggerWebhookTriggerExampleResultParsed) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type AutomationNewParamsTriggerWebhookTriggerExampleResultParsedValueUnion struct {
	OfString                                                            param.Opt[string]                                                                       `json:",omitzero,inline"`
	OfFloat                                                             param.Opt[float64]                                                                      `json:",omitzero,inline"`
	OfBool                                                              param.Opt[bool]                                                                         `json:",omitzero,inline"`
	OfAutomationNewsTriggerWebhookTriggerExampleResultParsedValueArray  []AutomationNewParamsTriggerWebhookTriggerExampleResultParsedValueArrayItemUnion        `json:",omitzero,inline"`
	OfAutomationNewsTriggerWebhookTriggerExampleResultParsedValueMapMap map[string]AutomationNewParamsTriggerWebhookTriggerExampleResultParsedValueMapItemUnion `json:",omitzero,inline"`
	paramUnion
}

func (u AutomationNewParamsTriggerWebhookTriggerExampleResultParsedValueUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString,
		u.OfFloat,
		u.OfBool,
		u.OfAutomationNewsTriggerWebhookTriggerExampleResultParsedValueArray,
		u.OfAutomationNewsTriggerWebhookTriggerExampleResultParsedValueMapMap)
}
func (u *AutomationNewParamsTriggerWebhookTriggerExampleResultParsedValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type AutomationNewParamsTriggerWebhookTriggerExampleResultParsedValueArrayItemUnion struct {
	OfString         param.Opt[string]                     `json:",omitzero,inline"`
	OfFloat          param.Opt[float64]                    `json:",omitzero,inline"`
	OfBool           param.Opt[bool]                       `json:",omitzero,inline"`
	OfJsonValueArray []shared.JsonValueUnionParam          `json:",omitzero,inline"`
	OfJsonValueMap   map[string]shared.JsonValueUnionParam `json:",omitzero,inline"`
	paramUnion
}

func (u AutomationNewParamsTriggerWebhookTriggerExampleResultParsedValueArrayItemUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString,
		u.OfFloat,
		u.OfBool,
		u.OfJsonValueArray,
		u.OfJsonValueMap)
}
func (u *AutomationNewParamsTriggerWebhookTriggerExampleResultParsedValueArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type AutomationNewParamsTriggerWebhookTriggerExampleResultParsedValueMapItemUnion struct {
	OfString         param.Opt[string]                     `json:",omitzero,inline"`
	OfFloat          param.Opt[float64]                    `json:",omitzero,inline"`
	OfBool           param.Opt[bool]                       `json:",omitzero,inline"`
	OfJsonValueArray []shared.JsonValueUnionParam          `json:",omitzero,inline"`
	OfJsonValueMap   map[string]shared.JsonValueUnionParam `json:",omitzero,inline"`
	paramUnion
}

func (u AutomationNewParamsTriggerWebhookTriggerExampleResultParsedValueMapItemUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString,
		u.OfFloat,
		u.OfBool,
		u.OfJsonValueArray,
		u.OfJsonValueMap)
}
func (u *AutomationNewParamsTriggerWebhookTriggerExampleResultParsedValueMapItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// The properties Schedule, Type are required.
type AutomationNewParamsTriggerScheduledTrigger struct {
	Schedule AutomationNewParamsTriggerScheduledTriggerSchedule `json:"schedule,omitzero" api:"required"`
	// Any of "SCHEDULED".
	Type string `json:"type,omitzero" api:"required"`
	paramObj
}

func (r AutomationNewParamsTriggerScheduledTrigger) MarshalJSON() (data []byte, err error) {
	type shadow AutomationNewParamsTriggerScheduledTrigger
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AutomationNewParamsTriggerScheduledTrigger) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[AutomationNewParamsTriggerScheduledTrigger](
		"type", "SCHEDULED",
	)
}

// The properties Cron, Timezone are required.
type AutomationNewParamsTriggerScheduledTriggerSchedule struct {
	// 5-field cron expression, for example `0 9 * * 1-5`.
	Cron string `json:"cron" api:"required"`
	// IANA timezone the cron is evaluated in.
	Timezone string `json:"timezone" api:"required"`
	paramObj
}

func (r AutomationNewParamsTriggerScheduledTriggerSchedule) MarshalJSON() (data []byte, err error) {
	type shadow AutomationNewParamsTriggerScheduledTriggerSchedule
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AutomationNewParamsTriggerScheduledTriggerSchedule) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties ObjectType, Type are required.
type AutomationNewParamsTriggerRecordCreatedTrigger struct {
	// Record type to watch: `$account`, `$contact`, `$opportunity`, `$meeting`,
	// `$task`, `$note`, `$email`, or the slug of a custom object. The `$` marks a
	// built-in; a bare name is always a custom object slug.
	ObjectType string `json:"objectType" api:"required"`
	// Any of "RECORD_CREATED".
	Type string `json:"type,omitzero" api:"required"`
	paramObj
}

func (r AutomationNewParamsTriggerRecordCreatedTrigger) MarshalJSON() (data []byte, err error) {
	type shadow AutomationNewParamsTriggerRecordCreatedTrigger
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AutomationNewParamsTriggerRecordCreatedTrigger) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[AutomationNewParamsTriggerRecordCreatedTrigger](
		"type", "RECORD_CREATED",
	)
}

// The properties ObjectType, Type are required.
type AutomationNewParamsTriggerRecordUpdatedTrigger struct {
	ObjectType string `json:"objectType" api:"required"`
	// Any of "RECORD_UPDATED".
	Type string `json:"type,omitzero" api:"required"`
	// Fields whose changes start a run, as field keys (`$name` for a built-in field,
	// the slug for a custom one) or attribute definition IDs. With no field or
	// relationship watches, runs on any change.
	WatchedFields []string `json:"watchedFields,omitzero"`
	// Relationships whose additions or removals start a run.
	WatchedRelationships []AutomationNewParamsTriggerRecordUpdatedTriggerWatchedRelationship `json:"watchedRelationships,omitzero"`
	paramObj
}

func (r AutomationNewParamsTriggerRecordUpdatedTrigger) MarshalJSON() (data []byte, err error) {
	type shadow AutomationNewParamsTriggerRecordUpdatedTrigger
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AutomationNewParamsTriggerRecordUpdatedTrigger) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[AutomationNewParamsTriggerRecordUpdatedTrigger](
		"type", "RECORD_UPDATED",
	)
}

// The property Key is required.
type AutomationNewParamsTriggerRecordUpdatedTriggerWatchedRelationship struct {
	// Public relationship key, such as `$contact` or a custom relationship slug.
	Key string `json:"key" api:"required"`
	// Optional when the endpoint object types differ. Required for a
	// self-relationship. A contradictory or ambiguous direction returns 422.
	//
	// Any of "SOURCE", "DESTINATION".
	Direction string `json:"direction,omitzero"`
	paramObj
}

func (r AutomationNewParamsTriggerRecordUpdatedTriggerWatchedRelationship) MarshalJSON() (data []byte, err error) {
	type shadow AutomationNewParamsTriggerRecordUpdatedTriggerWatchedRelationship
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AutomationNewParamsTriggerRecordUpdatedTriggerWatchedRelationship) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[AutomationNewParamsTriggerRecordUpdatedTriggerWatchedRelationship](
		"direction", "SOURCE", "DESTINATION",
	)
}

// The property Type is required.
type AutomationNewParamsTriggerSequenceReplyTrigger struct {
	// Any of "SEQUENCE_REPLY".
	Type string `json:"type,omitzero" api:"required"`
	paramObj
}

func (r AutomationNewParamsTriggerSequenceReplyTrigger) MarshalJSON() (data []byte, err error) {
	type shadow AutomationNewParamsTriggerSequenceReplyTrigger
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AutomationNewParamsTriggerSequenceReplyTrigger) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[AutomationNewParamsTriggerSequenceReplyTrigger](
		"type", "SEQUENCE_REPLY",
	)
}

// Model the agent steps run on. Returns 400 for an unsupported model and 422 when
// the workspace cannot choose the requested model; omit `model` or send `auto`.
type AutomationNewParamsModel string

const (
	AutomationNewParamsModelAuto        AutomationNewParamsModel = "auto"
	AutomationNewParamsModelAutoLow     AutomationNewParamsModel = "auto_low"
	AutomationNewParamsModelAutoHigh    AutomationNewParamsModel = "auto_high"
	AutomationNewParamsModelSonnet5_5   AutomationNewParamsModel = "sonnet_5_5"
	AutomationNewParamsModelHaiku5_5    AutomationNewParamsModel = "haiku_5_5"
	AutomationNewParamsModelGpt6_1Terra AutomationNewParamsModel = "gpt_6_1_terra"
	AutomationNewParamsModelGpt6Luna    AutomationNewParamsModel = "gpt_6_luna"
)

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type AutomationNewParamsRunAsUnion struct {
	OfWorkspaceRunAs *AutomationNewParamsRunAsWorkspaceRunAs `json:",omitzero,inline"`
	OfMemberRunAs    *AutomationNewParamsRunAsMemberRunAs    `json:",omitzero,inline"`
	paramUnion
}

func (u AutomationNewParamsRunAsUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfWorkspaceRunAs, u.OfMemberRunAs)
}
func (u *AutomationNewParamsRunAsUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Runs with no acting member.
//
// The property Type is required.
type AutomationNewParamsRunAsWorkspaceRunAs struct {
	// Any of "WORKSPACE".
	Type string `json:"type,omitzero" api:"required"`
	paramObj
}

func (r AutomationNewParamsRunAsWorkspaceRunAs) MarshalJSON() (data []byte, err error) {
	type shadow AutomationNewParamsRunAsWorkspaceRunAs
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AutomationNewParamsRunAsWorkspaceRunAs) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[AutomationNewParamsRunAsWorkspaceRunAs](
		"type", "WORKSPACE",
	)
}

// The properties MemberID, Type are required.
type AutomationNewParamsRunAsMemberRunAs struct {
	// Member the automation runs as.
	MemberID string `json:"memberId" api:"required"`
	// Any of "MEMBER".
	Type string `json:"type,omitzero" api:"required"`
	paramObj
}

func (r AutomationNewParamsRunAsMemberRunAs) MarshalJSON() (data []byte, err error) {
	type shadow AutomationNewParamsRunAsMemberRunAs
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AutomationNewParamsRunAsMemberRunAs) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[AutomationNewParamsRunAsMemberRunAs](
		"type", "MEMBER",
	)
}

type AutomationUpdateParams struct {
	// What the automation is meant to accomplish.
	Description param.Opt[string] `json:"description,omitzero"`
	// Display name of the automation.
	Name param.Opt[string] `json:"name,omitzero"`
	// Model the agent steps run on. Returns 400 for an unsupported model and 422 when
	// the workspace cannot choose the requested model; omit `model` or send `auto`.
	//
	// Any of "auto", "auto_low", "auto_high", "sonnet_5_5", "haiku_5_5",
	// "gpt_6_1_terra", "gpt_6_luna".
	Model AutomationUpdateParamsModel `json:"model,omitzero"`
	// Who the automation acts as. Workspace admins may name only themselves, a member,
	// or `WORKSPACE`; only a workspace owner may name another admin or an owner.
	RunAs AutomationUpdateParamsRunAsUnion `json:"runAs,omitzero"`
	// What a run does, in order.
	Steps []AutomationUpdateParamsStepUnion `json:"steps,omitzero"`
	// What starts a run.
	Triggers []AutomationUpdateParamsTriggerUnion `json:"triggers,omitzero"`
	paramObj
}

func (r AutomationUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow AutomationUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AutomationUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Model the agent steps run on. Returns 400 for an unsupported model and 422 when
// the workspace cannot choose the requested model; omit `model` or send `auto`.
type AutomationUpdateParamsModel string

const (
	AutomationUpdateParamsModelAuto        AutomationUpdateParamsModel = "auto"
	AutomationUpdateParamsModelAutoLow     AutomationUpdateParamsModel = "auto_low"
	AutomationUpdateParamsModelAutoHigh    AutomationUpdateParamsModel = "auto_high"
	AutomationUpdateParamsModelSonnet5_5   AutomationUpdateParamsModel = "sonnet_5_5"
	AutomationUpdateParamsModelHaiku5_5    AutomationUpdateParamsModel = "haiku_5_5"
	AutomationUpdateParamsModelGpt6_1Terra AutomationUpdateParamsModel = "gpt_6_1_terra"
	AutomationUpdateParamsModelGpt6Luna    AutomationUpdateParamsModel = "gpt_6_luna"
)

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type AutomationUpdateParamsRunAsUnion struct {
	OfWorkspaceRunAs *AutomationUpdateParamsRunAsWorkspaceRunAs `json:",omitzero,inline"`
	OfMemberRunAs    *AutomationUpdateParamsRunAsMemberRunAs    `json:",omitzero,inline"`
	paramUnion
}

func (u AutomationUpdateParamsRunAsUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfWorkspaceRunAs, u.OfMemberRunAs)
}
func (u *AutomationUpdateParamsRunAsUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Runs with no acting member.
//
// The property Type is required.
type AutomationUpdateParamsRunAsWorkspaceRunAs struct {
	// Any of "WORKSPACE".
	Type string `json:"type,omitzero" api:"required"`
	paramObj
}

func (r AutomationUpdateParamsRunAsWorkspaceRunAs) MarshalJSON() (data []byte, err error) {
	type shadow AutomationUpdateParamsRunAsWorkspaceRunAs
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AutomationUpdateParamsRunAsWorkspaceRunAs) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[AutomationUpdateParamsRunAsWorkspaceRunAs](
		"type", "WORKSPACE",
	)
}

// The properties MemberID, Type are required.
type AutomationUpdateParamsRunAsMemberRunAs struct {
	// Member the automation runs as.
	MemberID string `json:"memberId" api:"required"`
	// Any of "MEMBER".
	Type string `json:"type,omitzero" api:"required"`
	paramObj
}

func (r AutomationUpdateParamsRunAsMemberRunAs) MarshalJSON() (data []byte, err error) {
	type shadow AutomationUpdateParamsRunAsMemberRunAs
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AutomationUpdateParamsRunAsMemberRunAs) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[AutomationUpdateParamsRunAsMemberRunAs](
		"type", "MEMBER",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type AutomationUpdateParamsStepUnion struct {
	OfAgentStep *AutomationUpdateParamsStepAgentStep `json:",omitzero,inline"`
	OfCodeStep  *AutomationUpdateParamsStepCodeStep  `json:",omitzero,inline"`
	OfSleepStep *AutomationUpdateParamsStepSleepStep `json:",omitzero,inline"`
	paramUnion
}

func (u AutomationUpdateParamsStepUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfAgentStep, u.OfCodeStep, u.OfSleepStep)
}
func (u *AutomationUpdateParamsStepUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// The properties Name, Prompt, Type are required.
type AutomationUpdateParamsStepAgentStep struct {
	// Display name of the step.
	Name string `json:"name" api:"required"`
	// Instructions the agent follows in this step.
	Prompt string `json:"prompt" api:"required"`
	// Any of "AGENT".
	Type string `json:"type,omitzero" api:"required"`
	// What the agent may use in this step. Defaults to none.
	Permissions []AutomationUpdateParamsStepAgentStepPermissionUnion `json:"permissions,omitzero"`
	paramObj
}

func (r AutomationUpdateParamsStepAgentStep) MarshalJSON() (data []byte, err error) {
	type shadow AutomationUpdateParamsStepAgentStep
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AutomationUpdateParamsStepAgentStep) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[AutomationUpdateParamsStepAgentStep](
		"type", "AGENT",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type AutomationUpdateParamsStepAgentStepPermissionUnion struct {
	OfToolPermission    *AutomationUpdateParamsStepAgentStepPermissionToolPermission    `json:",omitzero,inline"`
	OfMcpPermission     *AutomationUpdateParamsStepAgentStepPermissionMcpPermission     `json:",omitzero,inline"`
	OfAPIPermission     *AutomationUpdateParamsStepAgentStepPermissionAPIPermission     `json:",omitzero,inline"`
	OfNetworkPermission *AutomationUpdateParamsStepAgentStepPermissionNetworkPermission `json:",omitzero,inline"`
	paramUnion
}

func (u AutomationUpdateParamsStepAgentStepPermissionUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfToolPermission, u.OfMcpPermission, u.OfAPIPermission, u.OfNetworkPermission)
}
func (u *AutomationUpdateParamsStepAgentStepPermissionUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// The properties Name, Type are required.
type AutomationUpdateParamsStepAgentStepPermissionToolPermission struct {
	// Name of the tool, matching the Lightfield app.
	Name string `json:"name" api:"required"`
	// Any of "TOOL".
	Type string `json:"type,omitzero" api:"required"`
	// Why the step needs this permission. Null when none was given. Returns 422 for an
	// empty or whitespace-only string.
	Reason param.Opt[string] `json:"reason,omitzero"`
	paramObj
}

func (r AutomationUpdateParamsStepAgentStepPermissionToolPermission) MarshalJSON() (data []byte, err error) {
	type shadow AutomationUpdateParamsStepAgentStepPermissionToolPermission
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AutomationUpdateParamsStepAgentStepPermissionToolPermission) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[AutomationUpdateParamsStepAgentStepPermissionToolPermission](
		"type", "TOOL",
	)
}

// The properties Connector, Tools, Type are required.
type AutomationUpdateParamsStepAgentStepPermissionMcpPermission struct {
	// Connector the tools are on. Must already be lowercase and trimmed.
	Connector string `json:"connector" api:"required"`
	// Unique, trimmed names of the connector tools the step may call.
	Tools []string `json:"tools,omitzero" api:"required"`
	// Any of "MCP".
	Type string `json:"type,omitzero" api:"required"`
	// Why the step needs this permission. Null when none was given. Returns 422 for an
	// empty or whitespace-only string.
	Reason param.Opt[string] `json:"reason,omitzero"`
	paramObj
}

func (r AutomationUpdateParamsStepAgentStepPermissionMcpPermission) MarshalJSON() (data []byte, err error) {
	type shadow AutomationUpdateParamsStepAgentStepPermissionMcpPermission
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AutomationUpdateParamsStepAgentStepPermissionMcpPermission) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[AutomationUpdateParamsStepAgentStepPermissionMcpPermission](
		"type", "MCP",
	)
}

// The properties Scopes, Type are required.
type AutomationUpdateParamsStepAgentStepPermissionAPIPermission struct {
	// Lightfield API scopes the step may use, for example `contacts:read`. The API key
	// must hold every scope it grants, and only a full-access key may grant `*`.
	Scopes []string `json:"scopes,omitzero" api:"required"`
	// Any of "API".
	Type string `json:"type,omitzero" api:"required"`
	// Why the step needs this permission. Null when none was given. Returns 422 for an
	// empty or whitespace-only string.
	Reason param.Opt[string] `json:"reason,omitzero"`
	paramObj
}

func (r AutomationUpdateParamsStepAgentStepPermissionAPIPermission) MarshalJSON() (data []byte, err error) {
	type shadow AutomationUpdateParamsStepAgentStepPermissionAPIPermission
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AutomationUpdateParamsStepAgentStepPermissionAPIPermission) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[AutomationUpdateParamsStepAgentStepPermissionAPIPermission](
		"type", "API",
	)
}

// The properties Domain, Type are required.
type AutomationUpdateParamsStepAgentStepPermissionNetworkPermission struct {
	// Lowercase host the step may call, optionally followed by a path glob. Returns
	// 422 for a scheme, query, fragment or surrounding whitespace.
	Domain string `json:"domain" api:"required"`
	// Any of "NETWORK".
	Type string `json:"type,omitzero" api:"required"`
	// Why the step needs this permission. Null when none was given. Returns 422 for an
	// empty or whitespace-only string.
	Reason param.Opt[string] `json:"reason,omitzero"`
	// Headers added to every request this permission allows.
	Headers map[string]string `json:"headers,omitzero"`
	// Restricts the permission to one HTTP method.
	//
	// Any of "GET", "POST", "PUT", "PATCH", "DELETE".
	Method string `json:"method,omitzero"`
	paramObj
}

func (r AutomationUpdateParamsStepAgentStepPermissionNetworkPermission) MarshalJSON() (data []byte, err error) {
	type shadow AutomationUpdateParamsStepAgentStepPermissionNetworkPermission
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AutomationUpdateParamsStepAgentStepPermissionNetworkPermission) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[AutomationUpdateParamsStepAgentStepPermissionNetworkPermission](
		"type", "NETWORK",
	)
	apijson.RegisterFieldValidator[AutomationUpdateParamsStepAgentStepPermissionNetworkPermission](
		"method", "GET", "POST", "PUT", "PATCH", "DELETE",
	)
}

// The properties Command, Name, Type are required.
type AutomationUpdateParamsStepCodeStep struct {
	// Bash command the step runs as written.
	Command string `json:"command" api:"required"`
	// Display name of the step.
	Name string `json:"name" api:"required"`
	// Any of "CODE".
	Type string `json:"type,omitzero" api:"required"`
	// What the command does. Does not change what runs.
	Description param.Opt[string] `json:"description,omitzero"`
	// What the command may use. Defaults to none.
	Permissions []AutomationUpdateParamsStepCodeStepPermissionUnion `json:"permissions,omitzero"`
	paramObj
}

func (r AutomationUpdateParamsStepCodeStep) MarshalJSON() (data []byte, err error) {
	type shadow AutomationUpdateParamsStepCodeStep
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AutomationUpdateParamsStepCodeStep) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[AutomationUpdateParamsStepCodeStep](
		"type", "CODE",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type AutomationUpdateParamsStepCodeStepPermissionUnion struct {
	OfAPIPermission     *AutomationUpdateParamsStepCodeStepPermissionAPIPermission     `json:",omitzero,inline"`
	OfNetworkPermission *AutomationUpdateParamsStepCodeStepPermissionNetworkPermission `json:",omitzero,inline"`
	paramUnion
}

func (u AutomationUpdateParamsStepCodeStepPermissionUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfAPIPermission, u.OfNetworkPermission)
}
func (u *AutomationUpdateParamsStepCodeStepPermissionUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// The properties Scopes, Type are required.
type AutomationUpdateParamsStepCodeStepPermissionAPIPermission struct {
	// Lightfield API scopes the step may use, for example `contacts:read`. The API key
	// must hold every scope it grants, and only a full-access key may grant `*`.
	Scopes []string `json:"scopes,omitzero" api:"required"`
	// Any of "API".
	Type string `json:"type,omitzero" api:"required"`
	// Why the step needs this permission. Null when none was given. Returns 422 for an
	// empty or whitespace-only string.
	Reason param.Opt[string] `json:"reason,omitzero"`
	paramObj
}

func (r AutomationUpdateParamsStepCodeStepPermissionAPIPermission) MarshalJSON() (data []byte, err error) {
	type shadow AutomationUpdateParamsStepCodeStepPermissionAPIPermission
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AutomationUpdateParamsStepCodeStepPermissionAPIPermission) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[AutomationUpdateParamsStepCodeStepPermissionAPIPermission](
		"type", "API",
	)
}

// The properties Domain, Type are required.
type AutomationUpdateParamsStepCodeStepPermissionNetworkPermission struct {
	// Lowercase host the step may call, optionally followed by a path glob. Returns
	// 422 for a scheme, query, fragment or surrounding whitespace.
	Domain string `json:"domain" api:"required"`
	// Any of "NETWORK".
	Type string `json:"type,omitzero" api:"required"`
	// Why the step needs this permission. Null when none was given. Returns 422 for an
	// empty or whitespace-only string.
	Reason param.Opt[string] `json:"reason,omitzero"`
	// Headers added to every request this permission allows.
	Headers map[string]string `json:"headers,omitzero"`
	// Restricts the permission to one HTTP method.
	//
	// Any of "GET", "POST", "PUT", "PATCH", "DELETE".
	Method string `json:"method,omitzero"`
	paramObj
}

func (r AutomationUpdateParamsStepCodeStepPermissionNetworkPermission) MarshalJSON() (data []byte, err error) {
	type shadow AutomationUpdateParamsStepCodeStepPermissionNetworkPermission
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AutomationUpdateParamsStepCodeStepPermissionNetworkPermission) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[AutomationUpdateParamsStepCodeStepPermissionNetworkPermission](
		"type", "NETWORK",
	)
	apijson.RegisterFieldValidator[AutomationUpdateParamsStepCodeStepPermissionNetworkPermission](
		"method", "GET", "POST", "PUT", "PATCH", "DELETE",
	)
}

// The properties Duration, Name, Type, Unit are required.
type AutomationUpdateParamsStepSleepStep struct {
	// How long to sleep before the next step, at most one year.
	Duration AutomationUpdateParamsStepSleepStepDurationUnion `json:"duration,omitzero" api:"required"`
	// Display name of the step.
	Name string `json:"name" api:"required"`
	// Any of "SLEEP".
	Type string `json:"type,omitzero" api:"required"`
	// Unit of `duration`.
	//
	// Any of "MILLISECOND", "SECOND", "MINUTE", "HOUR", "DAY".
	Unit string `json:"unit,omitzero" api:"required"`
	paramObj
}

func (r AutomationUpdateParamsStepSleepStep) MarshalJSON() (data []byte, err error) {
	type shadow AutomationUpdateParamsStepSleepStep
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AutomationUpdateParamsStepSleepStep) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[AutomationUpdateParamsStepSleepStep](
		"type", "SLEEP",
	)
	apijson.RegisterFieldValidator[AutomationUpdateParamsStepSleepStep](
		"unit", "MILLISECOND", "SECOND", "MINUTE", "HOUR", "DAY",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type AutomationUpdateParamsStepSleepStepDurationUnion struct {
	OfInt    param.Opt[int64]  `json:",omitzero,inline"`
	OfString param.Opt[string] `json:",omitzero,inline"`
	paramUnion
}

func (u AutomationUpdateParamsStepSleepStepDurationUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfInt, u.OfString)
}
func (u *AutomationUpdateParamsStepSleepStepDurationUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type AutomationUpdateParamsTriggerUnion struct {
	OfManualTrigger        *AutomationUpdateParamsTriggerManualTrigger        `json:",omitzero,inline"`
	OfWebhookTrigger       *AutomationUpdateParamsTriggerWebhookTrigger       `json:",omitzero,inline"`
	OfScheduledTrigger     *AutomationUpdateParamsTriggerScheduledTrigger     `json:",omitzero,inline"`
	OfRecordCreatedTrigger *AutomationUpdateParamsTriggerRecordCreatedTrigger `json:",omitzero,inline"`
	OfRecordUpdatedTrigger *AutomationUpdateParamsTriggerRecordUpdatedTrigger `json:",omitzero,inline"`
	OfSequenceReplyTrigger *AutomationUpdateParamsTriggerSequenceReplyTrigger `json:",omitzero,inline"`
	paramUnion
}

func (u AutomationUpdateParamsTriggerUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfManualTrigger,
		u.OfWebhookTrigger,
		u.OfScheduledTrigger,
		u.OfRecordCreatedTrigger,
		u.OfRecordUpdatedTrigger,
		u.OfSequenceReplyTrigger)
}
func (u *AutomationUpdateParamsTriggerUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// The property Type is required.
type AutomationUpdateParamsTriggerManualTrigger struct {
	// Any of "MANUAL".
	Type string `json:"type,omitzero" api:"required"`
	paramObj
}

func (r AutomationUpdateParamsTriggerManualTrigger) MarshalJSON() (data []byte, err error) {
	type shadow AutomationUpdateParamsTriggerManualTrigger
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AutomationUpdateParamsTriggerManualTrigger) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[AutomationUpdateParamsTriggerManualTrigger](
		"type", "MANUAL",
	)
}

// The property Type is required.
type AutomationUpdateParamsTriggerWebhookTrigger struct {
	// Any of "WEBHOOK".
	Type string `json:"type,omitzero" api:"required"`
	// A sample payload and the fields the steps read from it.
	ExampleResult AutomationUpdateParamsTriggerWebhookTriggerExampleResult `json:"exampleResult,omitzero"`
	paramObj
}

func (r AutomationUpdateParamsTriggerWebhookTrigger) MarshalJSON() (data []byte, err error) {
	type shadow AutomationUpdateParamsTriggerWebhookTrigger
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AutomationUpdateParamsTriggerWebhookTrigger) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[AutomationUpdateParamsTriggerWebhookTrigger](
		"type", "WEBHOOK",
	)
}

// A sample payload and the fields the steps read from it.
//
// The properties Json, Parsed are required.
type AutomationUpdateParamsTriggerWebhookTriggerExampleResult struct {
	Json   AutomationUpdateParamsTriggerWebhookTriggerExampleResultJsonUnion `json:"json,omitzero" api:"required"`
	Parsed []AutomationUpdateParamsTriggerWebhookTriggerExampleResultParsed  `json:"parsed,omitzero" api:"required"`
	paramObj
}

func (r AutomationUpdateParamsTriggerWebhookTriggerExampleResult) MarshalJSON() (data []byte, err error) {
	type shadow AutomationUpdateParamsTriggerWebhookTriggerExampleResult
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AutomationUpdateParamsTriggerWebhookTriggerExampleResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type AutomationUpdateParamsTriggerWebhookTriggerExampleResultJsonUnion struct {
	OfAutomationUpdatesTriggerWebhookTriggerExampleResultJsonMapMap map[string]AutomationUpdateParamsTriggerWebhookTriggerExampleResultJsonMapItemUnion `json:",omitzero,inline"`
	OfAutomationUpdatesTriggerWebhookTriggerExampleResultJsonArray  []AutomationUpdateParamsTriggerWebhookTriggerExampleResultJsonArrayItemUnion        `json:",omitzero,inline"`
	paramUnion
}

func (u AutomationUpdateParamsTriggerWebhookTriggerExampleResultJsonUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfAutomationUpdatesTriggerWebhookTriggerExampleResultJsonMapMap, u.OfAutomationUpdatesTriggerWebhookTriggerExampleResultJsonArray)
}
func (u *AutomationUpdateParamsTriggerWebhookTriggerExampleResultJsonUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type AutomationUpdateParamsTriggerWebhookTriggerExampleResultJsonMapItemUnion struct {
	OfString                                                               param.Opt[string]                                                                          `json:",omitzero,inline"`
	OfFloat                                                                param.Opt[float64]                                                                         `json:",omitzero,inline"`
	OfBool                                                                 param.Opt[bool]                                                                            `json:",omitzero,inline"`
	OfAutomationUpdatesTriggerWebhookTriggerExampleResultJsonMapItemArray  []AutomationUpdateParamsTriggerWebhookTriggerExampleResultJsonMapItemArrayItemUnion        `json:",omitzero,inline"`
	OfAutomationUpdatesTriggerWebhookTriggerExampleResultJsonMapItemMapMap map[string]AutomationUpdateParamsTriggerWebhookTriggerExampleResultJsonMapItemMapItemUnion `json:",omitzero,inline"`
	paramUnion
}

func (u AutomationUpdateParamsTriggerWebhookTriggerExampleResultJsonMapItemUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString,
		u.OfFloat,
		u.OfBool,
		u.OfAutomationUpdatesTriggerWebhookTriggerExampleResultJsonMapItemArray,
		u.OfAutomationUpdatesTriggerWebhookTriggerExampleResultJsonMapItemMapMap)
}
func (u *AutomationUpdateParamsTriggerWebhookTriggerExampleResultJsonMapItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type AutomationUpdateParamsTriggerWebhookTriggerExampleResultJsonMapItemArrayItemUnion struct {
	OfString         param.Opt[string]                     `json:",omitzero,inline"`
	OfFloat          param.Opt[float64]                    `json:",omitzero,inline"`
	OfBool           param.Opt[bool]                       `json:",omitzero,inline"`
	OfJsonValueArray []shared.JsonValueUnionParam          `json:",omitzero,inline"`
	OfJsonValueMap   map[string]shared.JsonValueUnionParam `json:",omitzero,inline"`
	paramUnion
}

func (u AutomationUpdateParamsTriggerWebhookTriggerExampleResultJsonMapItemArrayItemUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString,
		u.OfFloat,
		u.OfBool,
		u.OfJsonValueArray,
		u.OfJsonValueMap)
}
func (u *AutomationUpdateParamsTriggerWebhookTriggerExampleResultJsonMapItemArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type AutomationUpdateParamsTriggerWebhookTriggerExampleResultJsonMapItemMapItemUnion struct {
	OfString         param.Opt[string]                     `json:",omitzero,inline"`
	OfFloat          param.Opt[float64]                    `json:",omitzero,inline"`
	OfBool           param.Opt[bool]                       `json:",omitzero,inline"`
	OfJsonValueArray []shared.JsonValueUnionParam          `json:",omitzero,inline"`
	OfJsonValueMap   map[string]shared.JsonValueUnionParam `json:",omitzero,inline"`
	paramUnion
}

func (u AutomationUpdateParamsTriggerWebhookTriggerExampleResultJsonMapItemMapItemUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString,
		u.OfFloat,
		u.OfBool,
		u.OfJsonValueArray,
		u.OfJsonValueMap)
}
func (u *AutomationUpdateParamsTriggerWebhookTriggerExampleResultJsonMapItemMapItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type AutomationUpdateParamsTriggerWebhookTriggerExampleResultJsonArrayItemUnion struct {
	OfString                                                                 param.Opt[string]                                                                            `json:",omitzero,inline"`
	OfFloat                                                                  param.Opt[float64]                                                                           `json:",omitzero,inline"`
	OfBool                                                                   param.Opt[bool]                                                                              `json:",omitzero,inline"`
	OfAutomationUpdatesTriggerWebhookTriggerExampleResultJsonArrayItemArray  []AutomationUpdateParamsTriggerWebhookTriggerExampleResultJsonArrayItemArrayItemUnion        `json:",omitzero,inline"`
	OfAutomationUpdatesTriggerWebhookTriggerExampleResultJsonArrayItemMapMap map[string]AutomationUpdateParamsTriggerWebhookTriggerExampleResultJsonArrayItemMapItemUnion `json:",omitzero,inline"`
	paramUnion
}

func (u AutomationUpdateParamsTriggerWebhookTriggerExampleResultJsonArrayItemUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString,
		u.OfFloat,
		u.OfBool,
		u.OfAutomationUpdatesTriggerWebhookTriggerExampleResultJsonArrayItemArray,
		u.OfAutomationUpdatesTriggerWebhookTriggerExampleResultJsonArrayItemMapMap)
}
func (u *AutomationUpdateParamsTriggerWebhookTriggerExampleResultJsonArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type AutomationUpdateParamsTriggerWebhookTriggerExampleResultJsonArrayItemArrayItemUnion struct {
	OfString         param.Opt[string]                     `json:",omitzero,inline"`
	OfFloat          param.Opt[float64]                    `json:",omitzero,inline"`
	OfBool           param.Opt[bool]                       `json:",omitzero,inline"`
	OfJsonValueArray []shared.JsonValueUnionParam          `json:",omitzero,inline"`
	OfJsonValueMap   map[string]shared.JsonValueUnionParam `json:",omitzero,inline"`
	paramUnion
}

func (u AutomationUpdateParamsTriggerWebhookTriggerExampleResultJsonArrayItemArrayItemUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString,
		u.OfFloat,
		u.OfBool,
		u.OfJsonValueArray,
		u.OfJsonValueMap)
}
func (u *AutomationUpdateParamsTriggerWebhookTriggerExampleResultJsonArrayItemArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type AutomationUpdateParamsTriggerWebhookTriggerExampleResultJsonArrayItemMapItemUnion struct {
	OfString         param.Opt[string]                     `json:",omitzero,inline"`
	OfFloat          param.Opt[float64]                    `json:",omitzero,inline"`
	OfBool           param.Opt[bool]                       `json:",omitzero,inline"`
	OfJsonValueArray []shared.JsonValueUnionParam          `json:",omitzero,inline"`
	OfJsonValueMap   map[string]shared.JsonValueUnionParam `json:",omitzero,inline"`
	paramUnion
}

func (u AutomationUpdateParamsTriggerWebhookTriggerExampleResultJsonArrayItemMapItemUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString,
		u.OfFloat,
		u.OfBool,
		u.OfJsonValueArray,
		u.OfJsonValueMap)
}
func (u *AutomationUpdateParamsTriggerWebhookTriggerExampleResultJsonArrayItemMapItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// The property Path is required.
type AutomationUpdateParamsTriggerWebhookTriggerExampleResultParsed struct {
	Path  []string                                                                 `json:"path,omitzero" api:"required"`
	Label param.Opt[string]                                                        `json:"label,omitzero"`
	Value AutomationUpdateParamsTriggerWebhookTriggerExampleResultParsedValueUnion `json:"value,omitzero"`
	paramObj
}

func (r AutomationUpdateParamsTriggerWebhookTriggerExampleResultParsed) MarshalJSON() (data []byte, err error) {
	type shadow AutomationUpdateParamsTriggerWebhookTriggerExampleResultParsed
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AutomationUpdateParamsTriggerWebhookTriggerExampleResultParsed) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type AutomationUpdateParamsTriggerWebhookTriggerExampleResultParsedValueUnion struct {
	OfString                                                               param.Opt[string]                                                                          `json:",omitzero,inline"`
	OfFloat                                                                param.Opt[float64]                                                                         `json:",omitzero,inline"`
	OfBool                                                                 param.Opt[bool]                                                                            `json:",omitzero,inline"`
	OfAutomationUpdatesTriggerWebhookTriggerExampleResultParsedValueArray  []AutomationUpdateParamsTriggerWebhookTriggerExampleResultParsedValueArrayItemUnion        `json:",omitzero,inline"`
	OfAutomationUpdatesTriggerWebhookTriggerExampleResultParsedValueMapMap map[string]AutomationUpdateParamsTriggerWebhookTriggerExampleResultParsedValueMapItemUnion `json:",omitzero,inline"`
	paramUnion
}

func (u AutomationUpdateParamsTriggerWebhookTriggerExampleResultParsedValueUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString,
		u.OfFloat,
		u.OfBool,
		u.OfAutomationUpdatesTriggerWebhookTriggerExampleResultParsedValueArray,
		u.OfAutomationUpdatesTriggerWebhookTriggerExampleResultParsedValueMapMap)
}
func (u *AutomationUpdateParamsTriggerWebhookTriggerExampleResultParsedValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type AutomationUpdateParamsTriggerWebhookTriggerExampleResultParsedValueArrayItemUnion struct {
	OfString         param.Opt[string]                     `json:",omitzero,inline"`
	OfFloat          param.Opt[float64]                    `json:",omitzero,inline"`
	OfBool           param.Opt[bool]                       `json:",omitzero,inline"`
	OfJsonValueArray []shared.JsonValueUnionParam          `json:",omitzero,inline"`
	OfJsonValueMap   map[string]shared.JsonValueUnionParam `json:",omitzero,inline"`
	paramUnion
}

func (u AutomationUpdateParamsTriggerWebhookTriggerExampleResultParsedValueArrayItemUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString,
		u.OfFloat,
		u.OfBool,
		u.OfJsonValueArray,
		u.OfJsonValueMap)
}
func (u *AutomationUpdateParamsTriggerWebhookTriggerExampleResultParsedValueArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type AutomationUpdateParamsTriggerWebhookTriggerExampleResultParsedValueMapItemUnion struct {
	OfString         param.Opt[string]                     `json:",omitzero,inline"`
	OfFloat          param.Opt[float64]                    `json:",omitzero,inline"`
	OfBool           param.Opt[bool]                       `json:",omitzero,inline"`
	OfJsonValueArray []shared.JsonValueUnionParam          `json:",omitzero,inline"`
	OfJsonValueMap   map[string]shared.JsonValueUnionParam `json:",omitzero,inline"`
	paramUnion
}

func (u AutomationUpdateParamsTriggerWebhookTriggerExampleResultParsedValueMapItemUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString,
		u.OfFloat,
		u.OfBool,
		u.OfJsonValueArray,
		u.OfJsonValueMap)
}
func (u *AutomationUpdateParamsTriggerWebhookTriggerExampleResultParsedValueMapItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// The properties Schedule, Type are required.
type AutomationUpdateParamsTriggerScheduledTrigger struct {
	Schedule AutomationUpdateParamsTriggerScheduledTriggerSchedule `json:"schedule,omitzero" api:"required"`
	// Any of "SCHEDULED".
	Type string `json:"type,omitzero" api:"required"`
	paramObj
}

func (r AutomationUpdateParamsTriggerScheduledTrigger) MarshalJSON() (data []byte, err error) {
	type shadow AutomationUpdateParamsTriggerScheduledTrigger
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AutomationUpdateParamsTriggerScheduledTrigger) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[AutomationUpdateParamsTriggerScheduledTrigger](
		"type", "SCHEDULED",
	)
}

// The properties Cron, Timezone are required.
type AutomationUpdateParamsTriggerScheduledTriggerSchedule struct {
	// 5-field cron expression, for example `0 9 * * 1-5`.
	Cron string `json:"cron" api:"required"`
	// IANA timezone the cron is evaluated in.
	Timezone string `json:"timezone" api:"required"`
	paramObj
}

func (r AutomationUpdateParamsTriggerScheduledTriggerSchedule) MarshalJSON() (data []byte, err error) {
	type shadow AutomationUpdateParamsTriggerScheduledTriggerSchedule
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AutomationUpdateParamsTriggerScheduledTriggerSchedule) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties ObjectType, Type are required.
type AutomationUpdateParamsTriggerRecordCreatedTrigger struct {
	// Record type to watch: `$account`, `$contact`, `$opportunity`, `$meeting`,
	// `$task`, `$note`, `$email`, or the slug of a custom object. The `$` marks a
	// built-in; a bare name is always a custom object slug.
	ObjectType string `json:"objectType" api:"required"`
	// Any of "RECORD_CREATED".
	Type string `json:"type,omitzero" api:"required"`
	paramObj
}

func (r AutomationUpdateParamsTriggerRecordCreatedTrigger) MarshalJSON() (data []byte, err error) {
	type shadow AutomationUpdateParamsTriggerRecordCreatedTrigger
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AutomationUpdateParamsTriggerRecordCreatedTrigger) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[AutomationUpdateParamsTriggerRecordCreatedTrigger](
		"type", "RECORD_CREATED",
	)
}

// The properties ObjectType, Type are required.
type AutomationUpdateParamsTriggerRecordUpdatedTrigger struct {
	ObjectType string `json:"objectType" api:"required"`
	// Any of "RECORD_UPDATED".
	Type string `json:"type,omitzero" api:"required"`
	// Fields whose changes start a run, as field keys (`$name` for a built-in field,
	// the slug for a custom one) or attribute definition IDs. With no field or
	// relationship watches, runs on any change.
	WatchedFields []string `json:"watchedFields,omitzero"`
	// Relationships whose additions or removals start a run.
	WatchedRelationships []AutomationUpdateParamsTriggerRecordUpdatedTriggerWatchedRelationship `json:"watchedRelationships,omitzero"`
	paramObj
}

func (r AutomationUpdateParamsTriggerRecordUpdatedTrigger) MarshalJSON() (data []byte, err error) {
	type shadow AutomationUpdateParamsTriggerRecordUpdatedTrigger
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AutomationUpdateParamsTriggerRecordUpdatedTrigger) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[AutomationUpdateParamsTriggerRecordUpdatedTrigger](
		"type", "RECORD_UPDATED",
	)
}

// The property Key is required.
type AutomationUpdateParamsTriggerRecordUpdatedTriggerWatchedRelationship struct {
	// Public relationship key, such as `$contact` or a custom relationship slug.
	Key string `json:"key" api:"required"`
	// Optional when the endpoint object types differ. Required for a
	// self-relationship. A contradictory or ambiguous direction returns 422.
	//
	// Any of "SOURCE", "DESTINATION".
	Direction string `json:"direction,omitzero"`
	paramObj
}

func (r AutomationUpdateParamsTriggerRecordUpdatedTriggerWatchedRelationship) MarshalJSON() (data []byte, err error) {
	type shadow AutomationUpdateParamsTriggerRecordUpdatedTriggerWatchedRelationship
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AutomationUpdateParamsTriggerRecordUpdatedTriggerWatchedRelationship) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[AutomationUpdateParamsTriggerRecordUpdatedTriggerWatchedRelationship](
		"direction", "SOURCE", "DESTINATION",
	)
}

// The property Type is required.
type AutomationUpdateParamsTriggerSequenceReplyTrigger struct {
	// Any of "SEQUENCE_REPLY".
	Type string `json:"type,omitzero" api:"required"`
	paramObj
}

func (r AutomationUpdateParamsTriggerSequenceReplyTrigger) MarshalJSON() (data []byte, err error) {
	type shadow AutomationUpdateParamsTriggerSequenceReplyTrigger
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AutomationUpdateParamsTriggerSequenceReplyTrigger) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[AutomationUpdateParamsTriggerSequenceReplyTrigger](
		"type", "SEQUENCE_REPLY",
	)
}

type AutomationListParams struct {
	// Maximum number of records to return. Defaults to 25, maximum 25.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Number of records to skip for pagination. Defaults to 0.
	Offset param.Opt[int64] `query:"offset,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [AutomationListParams]'s query parameters as `url.Values`.
func (r AutomationListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AutomationDeleteParams struct {
	Body AutomationDeleteParamsBody
	paramObj
}

func (r AutomationDeleteParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.Body)
}
func (r *AutomationDeleteParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AutomationDeleteParamsBody struct {
	paramObj
}

func (r AutomationDeleteParamsBody) MarshalJSON() (data []byte, err error) {
	type shadow AutomationDeleteParamsBody
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AutomationDeleteParamsBody) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AutomationPauseParams struct {
	Body AutomationPauseParamsBody
	paramObj
}

func (r AutomationPauseParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.Body)
}
func (r *AutomationPauseParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AutomationPauseParamsBody struct {
	paramObj
}

func (r AutomationPauseParamsBody) MarshalJSON() (data []byte, err error) {
	type shadow AutomationPauseParamsBody
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AutomationPauseParamsBody) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AutomationResumeParams struct {
	Body AutomationResumeParamsBody
	paramObj
}

func (r AutomationResumeParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.Body)
}
func (r *AutomationResumeParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AutomationResumeParamsBody struct {
	paramObj
}

func (r AutomationResumeParamsBody) MarshalJSON() (data []byte, err error) {
	type shadow AutomationResumeParamsBody
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AutomationResumeParamsBody) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
