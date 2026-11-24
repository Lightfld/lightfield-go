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
	"time"

	"github.com/Lightfld/lightfield-go/internal/apijson"
	"github.com/Lightfld/lightfield-go/internal/apiquery"
	shimjson "github.com/Lightfld/lightfield-go/internal/encoding/json"
	"github.com/Lightfld/lightfield-go/internal/requestconfig"
	"github.com/Lightfld/lightfield-go/option"
	"github.com/Lightfld/lightfield-go/packages/param"
	"github.com/Lightfld/lightfield-go/packages/respjson"
)

// Custom objects and relationships are available on Pro and Growth plans. Records
// can be fetched and manipulated via these endpoints, and definitions can be
// fetched to discover the available custom object types in the system.
//
// ObjectService contains methods and other services that help with interacting
// with the Lightfield API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewObjectService] method instead.
type ObjectService struct {
	Options []option.RequestOption
}

// NewObjectService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewObjectService(opts ...option.RequestOption) (r ObjectService) {
	r = ObjectService{}
	r.Options = opts
	return
}

// Creates a new record for the specified custom object type.
//
// Read definitions for the object type’s configured `$pipeline` and `$stage`
// fields and options. See
// <u>[Pipelines and stages](/using-the-api/fields-and-relationships/#pipelines-and-stages)</u>
// for write rules.
func (r *ObjectService) New(ctx context.Context, entitySlug string, body ObjectNewParams, opts ...option.RequestOption) (res *ObjectCreateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if entitySlug == "" {
		err = errors.New("missing required entitySlug parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/objects/%s/values", url.PathEscape(entitySlug))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Retrieves a single record by ID for the specified custom object type. A record
// you may see by name only is returned as `{ id, restricted: { name, reason } }`,
// with no `fields` or `relationships`; check for `restricted` before reading them.
func (r *ObjectService) Get(ctx context.Context, id string, query ObjectGetParams, opts ...option.RequestOption) (res *ObjectRetrieveResponseUnion, err error) {
	opts = slices.Concat(r.Options, opts)
	if query.EntitySlug == "" {
		err = errors.New("missing required entitySlug parameter")
		return nil, err
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/objects/%s/values/%s", url.PathEscape(query.EntitySlug), url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Updates an existing record by ID for the specified custom object type. Only
// included fields and relationships are modified.
//
// Read definitions for the object type’s configured `$pipeline` and `$stage`
// fields and options. See
// <u>[Pipelines and stages](/using-the-api/fields-and-relationships/#pipelines-and-stages)</u>
// for write rules.
func (r *ObjectService) Update(ctx context.Context, id string, params ObjectUpdateParams, opts ...option.RequestOption) (res *ObjectUpdateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.EntitySlug == "" {
		err = errors.New("missing required entitySlug parameter")
		return nil, err
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/objects/%s/values/%s", url.PathEscape(params.EntitySlug), url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Returns a paginated list of records for the specified custom object type.
// Results are sorted by when the record was created in Lightfield (`createdAt`),
// newest first, unless you set `sort`; sort by `$updatedAt` to poll for changes.
// Filter with `field[operator]=value` parameters, which must all match, or with a
// JSON `filter` for `or` logic and nested groups. Use `fields` to return only the
// fields you need, which makes large pulls much smaller. Each page holds only the
// records you can currently read, so a page can contain fewer records than
// `limit`, or none, while more remain, and `totalCount` can include records you
// can no longer read. To read every record, keep advancing `offset` by `limit`
// until it reaches `totalCount`; do not stop at a short or empty page. Use
// `offset` and `limit` to paginate through results, and `$field` query parameters
// to filter. See <u>[List endpoints](/using-the-api/list-endpoints/)</u> for more
// information about <u>[pagination](/using-the-api/list-endpoints/#pagination)</u>
// and <u>[filtering](/using-the-api/list-endpoints/#filtering)</u>.
func (r *ObjectService) List(ctx context.Context, entitySlug string, query ObjectListParams, opts ...option.RequestOption) (res *ObjectListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if entitySlug == "" {
		err = errors.New("missing required entitySlug parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/objects/%s", url.PathEscape(entitySlug))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Moves a custom object record to the trash. The record is soft-deleted and may be
// restored with `POST /v1/objects/{entitySlug}/values/{id}/restore`.
//
// Calling delete on an already-trashed record is a no-op and returns the existing
// record.
//
// **[Required scope](/using-the-api/scopes/):** `custom_objects:delete`
//
// **[Rate limit category](/using-the-api/rate-limits/):** Write
func (r *ObjectService) Delete(ctx context.Context, id string, params ObjectDeleteParams, opts ...option.RequestOption) (res *ObjectDeleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.EntitySlug == "" {
		err = errors.New("missing required entitySlug parameter")
		return nil, err
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/objects/%s/values/%s", url.PathEscape(params.EntitySlug), url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, params, &res, opts...)
	return res, err
}

// > [!WARNING] This endpoint is in beta and isn't available to every workspace
// > yet. Its parameters and response schemas may change, including breaking
// > changes.
//
// Archives a custom object definition and its active lists while retaining its
// records. Returns 409 if dependencies prevent archival; resolve them in the
// Lightfield app. Restoring the definition does not restore its lists.
//
// **[Required scope](/using-the-api/scopes/):** `config:delete`
//
// **[Rate limit category](/using-the-api/rate-limits/):** Write
func (r *ObjectService) ArchiveType(ctx context.Context, entitySlug string, body ObjectArchiveTypeParams, opts ...option.RequestOption) (res *ObjectArchiveTypeResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if entitySlug == "" {
		err = errors.New("missing required entitySlug parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/objects/%s", url.PathEscape(entitySlug))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &res, opts...)
	return res, err
}

// > [!WARNING] This endpoint is in beta and isn't available to every workspace
// > yet. Its parameters and response schemas may change, including breaking
// > changes.
//
// Creates a pipeline on the object type. Each stage receives a UUIDv4 unless the
// caller supplies its ID.
//
// A name, slug or stage ID that another pipeline on the object type has is refused
// with 409 and a `code` of `crm.pipeline.name_taken`, `crm.pipeline.slug_taken` or
// `crm.pipeline.stage_id_taken`. Stage IDs on an object type must also be distinct
// after ignoring case, hyphens, and underscores. A new collision returns 409
// `crm.pipeline.stage_id_taken`. A stage ID that newly conflicts with another
// stage’s label, in either its stored or published spelling, is refused with 409
// `crm.pipeline.stage_id_taken`. A name or stage ID can be taken once the other
// pipeline gives it up, so the request can be sent again then. A workspace without
// multiple pipelines is refused with 403 `crm.pipeline.configuration_disabled`.
//
// Requires a workspace owner or admin.
//
// **[Required scopes](/using-the-api/scopes/):** `custom_objects:read` and
// `config:create`
//
// **[Rate limit category](/using-the-api/rate-limits/):** Write
func (r *ObjectService) NewPipeline(ctx context.Context, entitySlug string, body ObjectNewPipelineParams, opts ...option.RequestOption) (res *ObjectCreatePipelineResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if entitySlug == "" {
		err = errors.New("missing required entitySlug parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/objects/%s/pipelines", url.PathEscape(entitySlug))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// > [!WARNING] This endpoint is in beta and isn't available to every workspace
// > yet. Its parameters and response schemas may change, including breaking
// > changes.
//
// Creates a custom object type with its default attributes and relationships.
// Returns its public metadata. Requires the custom objects entitlement.
//
// When the slug belongs to an archived custom object type, restores that type with
// its records and applies the request as if creating it: omitted optional fields
// take their defaults. Its lists stay archived. Returns 409 when the slug belongs
// to an active custom object type in this workspace.
//
// **[Required scope](/using-the-api/scopes/):** `config:create`
//
// **[Rate limit category](/using-the-api/rate-limits/):** Write
func (r *ObjectService) NewType(ctx context.Context, body ObjectNewTypeParams, opts ...option.RequestOption) (res *ObjectCreateTypeResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/objects"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Returns field and relationship definitions for the specified custom object type.
func (r *ObjectService) Definitions(ctx context.Context, entitySlug string, query ObjectDefinitionsParams, opts ...option.RequestOption) (res *ObjectDefinitionsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if entitySlug == "" {
		err = errors.New("missing required entitySlug parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/objects/%s/definitions", url.PathEscape(entitySlug))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// > [!WARNING] This endpoint is in beta and isn't available to every workspace
// > yet. Its parameters and response schemas may change, including breaking
// > changes.
//
// Deletes the pipeline and its stages, clearing the stage of every record in one.
// Automations and saved filters that reference it are left as they are. With
// `dryRun=true`, runs every check and returns the count without deleting anything.
// A workspace without multiple pipelines is refused with 403
// `crm.pipeline.configuration_disabled`.
//
// Requires a workspace owner or admin.
//
// **[Required scopes](/using-the-api/scopes/):** `custom_objects:read` and
// `config:delete`
//
// **[Rate limit category](/using-the-api/rate-limits/):** Write
func (r *ObjectService) DeletePipeline(ctx context.Context, id string, params ObjectDeletePipelineParams, opts ...option.RequestOption) (res *ObjectDeletePipelineResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.EntitySlug == "" {
		err = errors.New("missing required entitySlug parameter")
		return nil, err
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/objects/%s/pipelines/%s", url.PathEscape(params.EntitySlug), url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, params, &res, opts...)
	return res, err
}

// Returns the value-change history for a single field on a record, newest first.
// Consecutive identical values are collapsed. History is cursor-paginated: pass
// `nextCursor` from the previous response as `after` to page through older values.
// Only attribute-backed fields (custom attributes and attribute-backed system
// fields) have history — column-backed system fields return an error. A record you
// may see by name only has no visible history and returns 404, like a record that
// does not exist.
//
// **[Required scope](/using-the-api/scopes/):** `custom_objects:read`
//
// **[Rate limit category](/using-the-api/rate-limits/):** Read
func (r *ObjectService) FieldHistory(ctx context.Context, fieldKey string, params ObjectFieldHistoryParams, opts ...option.RequestOption) (res *ObjectFieldHistoryResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.EntitySlug == "" {
		err = errors.New("missing required entitySlug parameter")
		return nil, err
	}
	if params.ID == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	if fieldKey == "" {
		err = errors.New("missing required fieldKey parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/objects/%s/values/%s/fields/%s/history", url.PathEscape(params.EntitySlug), url.PathEscape(params.ID), url.PathEscape(fieldKey))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

// Returns all custom object types available to the caller, sorted alphabetically
// by `objectType`.
func (r *ObjectService) ListDefinitions(ctx context.Context, opts ...option.RequestOption) (res *ObjectListDefinitionsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/objects"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// > [!WARNING] This endpoint is in beta and isn't available to every workspace
// > yet. Its parameters and response schemas may change, including breaking
// > changes.
//
// Returns every pipeline on the object type, oldest first, each with its stages in
// display order. Empty when the object type has none.
//
// **[Required scope](/using-the-api/scopes/):** `custom_objects:read`
//
// **[Rate limit category](/using-the-api/rate-limits/):** Read
func (r *ObjectService) ListPipelines(ctx context.Context, entitySlug string, opts ...option.RequestOption) (res *ObjectListPipelinesResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if entitySlug == "" {
		err = errors.New("missing required entitySlug parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/objects/%s/pipelines", url.PathEscape(entitySlug))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Restores a trashed custom object record, bringing it back to active.
//
// Calling restore on a record that is not trashed is a no-op and returns the
// existing record. A record of a different custom object type than `entitySlug` is
// reported as not found.
//
// **[Required scope](/using-the-api/scopes/):** `custom_objects:restore`
//
// **[Rate limit category](/using-the-api/rate-limits/):** Write
func (r *ObjectService) Restore(ctx context.Context, id string, params ObjectRestoreParams, opts ...option.RequestOption) (res *ObjectRestoreResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.EntitySlug == "" {
		err = errors.New("missing required entitySlug parameter")
		return nil, err
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/objects/%s/values/%s/restore", url.PathEscape(params.EntitySlug), url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// > [!WARNING] This endpoint is in beta and isn't available to every workspace
// > yet. Its parameters and response schemas may change, including breaking
// > changes.
//
// Returns one pipeline on the object type, with its stages in display order.
//
// **[Required scope](/using-the-api/scopes/):** `custom_objects:read`
//
// **[Rate limit category](/using-the-api/rate-limits/):** Read
func (r *ObjectService) GetPipeline(ctx context.Context, id string, query ObjectGetPipelineParams, opts ...option.RequestOption) (res *ObjectRetrievePipelineResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if query.EntitySlug == "" {
		err = errors.New("missing required entitySlug parameter")
		return nil, err
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/objects/%s/pipelines/%s", url.PathEscape(query.EntitySlug), url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// > [!WARNING] This endpoint is in beta and isn't available to every workspace
// > yet. Its parameters and response schemas may change, including breaking
// > changes.
//
// Retrieves the public metadata of a custom object type in this workspace. Returns
// 404 for unknown slugs, system object types, and archived definitions unless
// includeArchived is true.
//
// **[Required scope](/using-the-api/scopes/):** `config:read`
//
// **[Rate limit category](/using-the-api/rate-limits/):** Read
func (r *ObjectService) GetSettings(ctx context.Context, entitySlug string, query ObjectGetSettingsParams, opts ...option.RequestOption) (res *ObjectRetrieveSettingsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if entitySlug == "" {
		err = errors.New("missing required entitySlug parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/objects/%s/settings", url.PathEscape(entitySlug))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// > [!WARNING] This endpoint is in beta and isn't available to every workspace
// > yet. Its parameters and response schemas may change, including breaking
// > changes.
//
// Requires a workspace owner or admin. Creates, changes, archives and restores
// definitions, keyed the way `GET` returns them. A key that names no definition
// creates one; a key that names one changes only the properties sent, and restores
// it first if it is archived. Sending the same request again changes nothing.
//
// Every definition is checked before any is written. If any check fails, nothing
// is changed and the response is 422, or 409 when a definition that already exists
// disagrees with the request, with every problem listed under `issues`. The writes
// themselves are independent: in the rare case one fails after the checks pass,
// the others are still written, and the response is 409 with the written
// definitions listed under `applied`. If another request is writing definitions at
// the same time, the response can be 429 with nothing changed; send the request
// again.
//
// Pipeline and stage fields cannot be changed through this endpoint. Labels are
// stored without leading or trailing whitespace. Labels, new keys and new option
// IDs can be at most 255 characters.
//
// **[Required scopes](/using-the-api/scopes/):** `custom_objects:read`, plus
// `config:create` to create, `config:update` to change or restore, `config:delete`
// to archive
//
// **[Rate limit category](/using-the-api/rate-limits/):** Write
func (r *ObjectService) UpdateDefinitions(ctx context.Context, entitySlug string, body ObjectUpdateDefinitionsParams, opts ...option.RequestOption) (res *ObjectUpdateDefinitionsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if entitySlug == "" {
		err = errors.New("missing required entitySlug parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/objects/%s/definitions", url.PathEscape(entitySlug))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// > [!WARNING] This endpoint is in beta and isn't available to every workspace
// > yet. Its parameters and response schemas may change, including breaking
// > changes.
//
// Changes only the properties sent. A stage entry without an ID creates a new
// stage on each request. Use the returned IDs on subsequent updates to keep those
// stages and their record assignments. If the list omits an existing stage, the
// update returns 409 and lists the stages it would delete. Repeat with
// `force=true` to confirm. With `dryRun=true`, runs every check and returns
// `{ affectedRecordCount }`, the records whose stage the update would clear,
// without changing anything.
//
// A name, slug or stage ID that another pipeline on the object type has is refused
// with 409 and a `code` of `crm.pipeline.name_taken`, `crm.pipeline.slug_taken` or
// `crm.pipeline.stage_id_taken`. Stage IDs on an object type must also be distinct
// after ignoring case, hyphens, and underscores. A new collision returns 409
// `crm.pipeline.stage_id_taken`. A stage ID that newly conflicts with another
// stage’s label, in either its stored or published spelling, is refused with 409
// `crm.pipeline.stage_id_taken`. A name or stage ID can be taken once the other
// pipeline gives it up, so the request can be sent again then. A workspace without
// multiple pipelines is refused with 403 `crm.pipeline.configuration_disabled`.
//
// Requires a workspace owner or admin.
//
// **[Required scopes](/using-the-api/scopes/):** `custom_objects:read` and
// `config:update`
//
// **[Rate limit category](/using-the-api/rate-limits/):** Write
func (r *ObjectService) UpdatePipeline(ctx context.Context, id string, params ObjectUpdatePipelineParams, opts ...option.RequestOption) (res *ObjectUpdatePipelineResponseUnion, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.EntitySlug == "" {
		err = errors.New("missing required entitySlug parameter")
		return nil, err
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/objects/%s/pipelines/%s", url.PathEscape(params.EntitySlug), url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// > [!WARNING] This endpoint is in beta and isn't available to every workspace
// > yet. Its parameters and response schemas may change, including breaking
// > changes.
//
// Updates display labels, description, and icon for an active custom object type.
// Omitted fields remain unchanged. Set labelPlural, description, or icon to null
// to clear them. The slug is immutable. Missing, system, and archived object types
// return 404.
//
// **[Required scope](/using-the-api/scopes/):** `config:update`
//
// **[Rate limit category](/using-the-api/rate-limits/):** Write
func (r *ObjectService) UpdateSettings(ctx context.Context, entitySlug string, body ObjectUpdateSettingsParams, opts ...option.RequestOption) (res *ObjectUpdateSettingsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if entitySlug == "" {
		err = errors.New("missing required entitySlug parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/objects/%s/settings", url.PathEscape(entitySlug))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

type ObjectArchiveTypeResponse struct {
	// Stable public ID of the custom object definition.
	ID string `json:"id" api:"required"`
	// Whether this custom object type is archived.
	Archived bool `json:"archived" api:"required"`
	// ISO 8601 archive timestamp, or null when active.
	ArchivedAt string `json:"archivedAt" api:"required"`
	// ISO 8601 creation timestamp.
	CreatedAt string `json:"createdAt" api:"required"`
	// Description of the custom object type, or null.
	Description string `json:"description" api:"required"`
	// Display icon, or null when no icon is configured.
	Icon string `json:"icon" api:"required"`
	// Plural display label, or null.
	LabelPlural string `json:"labelPlural" api:"required"`
	// Singular display label.
	LabelSingular string `json:"labelSingular" api:"required"`
	// Custom object slug used in API paths.
	Slug string `json:"slug" api:"required"`
	// ISO 8601 last-update timestamp.
	UpdatedAt string `json:"updatedAt" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID            respjson.Field
		Archived      respjson.Field
		ArchivedAt    respjson.Field
		CreatedAt     respjson.Field
		Description   respjson.Field
		Icon          respjson.Field
		LabelPlural   respjson.Field
		LabelSingular respjson.Field
		Slug          respjson.Field
		UpdatedAt     respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ObjectArchiveTypeResponse) RawJSON() string { return r.JSON.raw }
func (r *ObjectArchiveTypeResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ObjectCreatePipelineResponse struct {
	// Unique identifier of the pipeline.
	ID string `json:"id" api:"required"`
	// Description of the pipeline, or null.
	Description string `json:"description" api:"required"`
	// Whether stages carry a `forecastPercentage`. Turning it off clears every stage
	// forecast.
	ForecastEnabled bool `json:"forecastEnabled" api:"required"`
	// Display name of the pipeline.
	Name string `json:"name" api:"required"`
	// Key derived from the name when the pipeline was created, unique on the object
	// type. Pipelines provisioned by Lightfield use a `$` prefix (e.g.
	// `$default-pipeline`).
	Slug string `json:"slug" api:"required"`
	// Stages in display order. Each `id` is the option ID records hold in `$stage`.
	Stages []ObjectCreatePipelineResponseStage `json:"stages" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID              respjson.Field
		Description     respjson.Field
		ForecastEnabled respjson.Field
		Name            respjson.Field
		Slug            respjson.Field
		Stages          respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ObjectCreatePipelineResponse) RawJSON() string { return r.JSON.raw }
func (r *ObjectCreatePipelineResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ObjectCreatePipelineResponseStage struct {
	// Unique identifier of the select option.
	ID string `json:"id" api:"required"`
	// Human-readable display name of the option.
	Label string `json:"label" api:"required"`
	// Lifecycle group of the option, in-progress (`ACTIVE`) or done (`COMPLETED`),
	// `ACTIVE` when none was set. Absent on selects whose options are not grouped.
	//
	// Any of "ACTIVE", "COMPLETED".
	Category string `json:"category"`
	// Color the option is displayed in, `GRAY` when none was chosen. Absent on selects
	// whose options take no color.
	//
	// Any of "RED", "ORANGE", "YELLOW", "LIME", "GREEN", "BLUE", "INDIGO", "PURPLE",
	// "COPPER", "GRAY".
	Color string `json:"color"`
	// Description of the option, or null.
	Description string `json:"description" api:"nullable"`
	// Forecast weighting of the option from 0 to 100, or null when none is set. Absent
	// on selects whose options take no forecast.
	ForecastPercentage float64 `json:"forecastPercentage" api:"nullable"`
	// Icon the option is displayed with, `BACKLOG` when none was chosen. Present only
	// on $stage.
	//
	// Any of "BACKLOG", "QUALIFICATION", "DISCOVERY", "ACTIVE", "NEGOTIATION",
	// "CLOSED_WON", "CLOSED_LOST".
	Icon string `json:"icon"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                 respjson.Field
		Label              respjson.Field
		Category           respjson.Field
		Color              respjson.Field
		Description        respjson.Field
		ForecastPercentage respjson.Field
		Icon               respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ObjectCreatePipelineResponseStage) RawJSON() string { return r.JSON.raw }
func (r *ObjectCreatePipelineResponseStage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ObjectCreateResponse struct {
	// Unique identifier for the entity.
	ID string `json:"id" api:"required"`
	// ISO 8601 timestamp of when the entity was created.
	CreatedAt string `json:"createdAt" api:"required"`
	// Map of field names to their typed values. System fields are prefixed with `$`
	// (e.g. `$name`, `$email`); custom attributes use their bare slug.
	Fields map[string]ObjectCreateResponseField `json:"fields" api:"required"`
	// URL to view the entity in the Lightfield web app, or null.
	HTTPLink string `json:"httpLink" api:"required"`
	// Map of relationship names to their associated entities. System relationships are
	// prefixed with `$` (e.g. `$owner`, `$contact`). A self-referential relationship
	// keeps its merged values and includes optional directional values. List filters
	// select the current record's role with `:asSource` or `:asDestination`, for
	// example `?parent:asSource[contains]=account_id`.
	Relationships map[string]ObjectCreateResponseRelationship `json:"relationships" api:"required"`
	// ISO 8601 timestamp of when the entity was last updated, or null.
	UpdatedAt string `json:"updatedAt" api:"required"`
	// External identifier for the entity, or null if unset.
	ExternalID string `json:"externalId" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID            respjson.Field
		CreatedAt     respjson.Field
		Fields        respjson.Field
		HTTPLink      respjson.Field
		Relationships respjson.Field
		UpdatedAt     respjson.Field
		ExternalID    respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ObjectCreateResponse) RawJSON() string { return r.JSON.raw }
func (r *ObjectCreateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ObjectCreateResponseField struct {
	// The field value, or null if unset.
	Value ObjectCreateResponseFieldValueUnion `json:"value" api:"required"`
	// The data type of the field.
	//
	// Any of "ADDRESS", "CHECKBOX", "CURRENCY", "DATE", "DATETIME", "EMAIL",
	// "FULL_NAME", "MARKDOWN", "MULTI_SELECT", "NUMBER", "SINGLE_SELECT",
	// "SOCIAL_HANDLE", "TELEPHONE", "TEXT", "URL", "HTML".
	ValueType string `json:"valueType" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Value       respjson.Field
		ValueType   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ObjectCreateResponseField) RawJSON() string { return r.JSON.raw }
func (r *ObjectCreateResponseField) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ObjectCreateResponseFieldValueUnion contains all possible properties and values
// from [string], [float64], [bool], [[]string],
// [ObjectCreateResponseFieldValueAddress],
// [ObjectCreateResponseFieldValueFullName].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfStringArray]
type ObjectCreateResponseFieldValueUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]string] instead of an object.
	OfStringArray []string `json:",inline"`
	// This field is from variant [ObjectCreateResponseFieldValueAddress].
	City string `json:"city"`
	// This field is from variant [ObjectCreateResponseFieldValueAddress].
	Country string `json:"country"`
	// This field is from variant [ObjectCreateResponseFieldValueAddress].
	Latitude float64 `json:"latitude"`
	// This field is from variant [ObjectCreateResponseFieldValueAddress].
	Longitude float64 `json:"longitude"`
	// This field is from variant [ObjectCreateResponseFieldValueAddress].
	PostalCode string `json:"postalCode"`
	// This field is from variant [ObjectCreateResponseFieldValueAddress].
	State string `json:"state"`
	// This field is from variant [ObjectCreateResponseFieldValueAddress].
	Street string `json:"street"`
	// This field is from variant [ObjectCreateResponseFieldValueAddress].
	Street2 string `json:"street2"`
	// This field is from variant [ObjectCreateResponseFieldValueFullName].
	FirstName string `json:"firstName"`
	// This field is from variant [ObjectCreateResponseFieldValueFullName].
	LastName string `json:"lastName"`
	JSON     struct {
		OfString      respjson.Field
		OfFloat       respjson.Field
		OfBool        respjson.Field
		OfStringArray respjson.Field
		City          respjson.Field
		Country       respjson.Field
		Latitude      respjson.Field
		Longitude     respjson.Field
		PostalCode    respjson.Field
		State         respjson.Field
		Street        respjson.Field
		Street2       respjson.Field
		FirstName     respjson.Field
		LastName      respjson.Field
		raw           string
	} `json:"-"`
}

func (u ObjectCreateResponseFieldValueUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ObjectCreateResponseFieldValueUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ObjectCreateResponseFieldValueUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ObjectCreateResponseFieldValueUnion) AsStringArray() (v []string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ObjectCreateResponseFieldValueUnion) AsAddress() (v ObjectCreateResponseFieldValueAddress) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ObjectCreateResponseFieldValueUnion) AsFullName() (v ObjectCreateResponseFieldValueFullName) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ObjectCreateResponseFieldValueUnion) RawJSON() string { return u.JSON.raw }

func (r *ObjectCreateResponseFieldValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ObjectCreateResponseFieldValueAddress struct {
	// City name.
	City string `json:"city" api:"nullable"`
	// 2-letter ISO 3166-1 alpha-2 country code.
	Country string `json:"country" api:"nullable"`
	// Latitude coordinate.
	Latitude float64 `json:"latitude" api:"nullable"`
	// Longitude coordinate.
	Longitude float64 `json:"longitude" api:"nullable"`
	// Postal or ZIP code.
	PostalCode string `json:"postalCode" api:"nullable"`
	// State or province.
	State string `json:"state" api:"nullable"`
	// Street address line 1.
	Street string `json:"street" api:"nullable"`
	// Street address line 2.
	Street2 string `json:"street2" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		City        respjson.Field
		Country     respjson.Field
		Latitude    respjson.Field
		Longitude   respjson.Field
		PostalCode  respjson.Field
		State       respjson.Field
		Street      respjson.Field
		Street2     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ObjectCreateResponseFieldValueAddress) RawJSON() string { return r.JSON.raw }
func (r *ObjectCreateResponseFieldValueAddress) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ObjectCreateResponseFieldValueFullName struct {
	// The contact's first name.
	FirstName string `json:"firstName" api:"nullable"`
	// The contact's last name.
	LastName string `json:"lastName" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FirstName   respjson.Field
		LastName    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ObjectCreateResponseFieldValueFullName) RawJSON() string { return r.JSON.raw }
func (r *ObjectCreateResponseFieldValueFullName) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ObjectCreateResponseRelationship struct {
	// Whether the relationship is `has_one` or `has_many`.
	Cardinality string `json:"cardinality" api:"required"`
	// The type of the related object (e.g. `account`, `contact`).
	ObjectType string `json:"objectType" api:"required"`
	// IDs of the related entities.
	Values []string `json:"values" api:"required"`
	// Directional values for a self-referential relationship. List filters use the
	// relationship key with `:asSource` or `:asDestination`; writes continue to use
	// the bare key.
	Directions ObjectCreateResponseRelationshipDirections `json:"directions"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Cardinality respjson.Field
		ObjectType  respjson.Field
		Values      respjson.Field
		Directions  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ObjectCreateResponseRelationship) RawJSON() string { return r.JSON.raw }
func (r *ObjectCreateResponseRelationship) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Directional values for a self-referential relationship. List filters use the
// relationship key with `:asSource` or `:asDestination`; writes continue to use
// the bare key.
type ObjectCreateResponseRelationshipDirections struct {
	// The current record is the destination (incoming) endpoint.
	AsDestination ObjectCreateResponseRelationshipDirectionsAsDestination `json:"asDestination" api:"required"`
	// The current record is the source (outgoing) endpoint.
	AsSource ObjectCreateResponseRelationshipDirectionsAsSource `json:"asSource" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AsDestination respjson.Field
		AsSource      respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ObjectCreateResponseRelationshipDirections) RawJSON() string { return r.JSON.raw }
func (r *ObjectCreateResponseRelationshipDirections) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The current record is the destination (incoming) endpoint.
type ObjectCreateResponseRelationshipDirectionsAsDestination struct {
	// Whether this endpoint direction is `HAS_ONE` or `HAS_MANY`.
	//
	// Any of "HAS_ONE", "HAS_MANY".
	Cardinality string `json:"cardinality" api:"required"`
	// IDs of the related entities in this direction.
	Values []string `json:"values" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Cardinality respjson.Field
		Values      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ObjectCreateResponseRelationshipDirectionsAsDestination) RawJSON() string { return r.JSON.raw }
func (r *ObjectCreateResponseRelationshipDirectionsAsDestination) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The current record is the source (outgoing) endpoint.
type ObjectCreateResponseRelationshipDirectionsAsSource struct {
	// Whether this endpoint direction is `HAS_ONE` or `HAS_MANY`.
	//
	// Any of "HAS_ONE", "HAS_MANY".
	Cardinality string `json:"cardinality" api:"required"`
	// IDs of the related entities in this direction.
	Values []string `json:"values" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Cardinality respjson.Field
		Values      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ObjectCreateResponseRelationshipDirectionsAsSource) RawJSON() string { return r.JSON.raw }
func (r *ObjectCreateResponseRelationshipDirectionsAsSource) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ObjectCreateTypeResponse struct {
	// Stable public ID of the custom object definition.
	ID string `json:"id" api:"required"`
	// Whether this custom object type is archived.
	Archived bool `json:"archived" api:"required"`
	// ISO 8601 archive timestamp, or null when active.
	ArchivedAt string `json:"archivedAt" api:"required"`
	// ISO 8601 creation timestamp.
	CreatedAt string `json:"createdAt" api:"required"`
	// Description of the custom object type, or null.
	Description string `json:"description" api:"required"`
	// Display icon, or null when no icon is configured.
	Icon string `json:"icon" api:"required"`
	// Plural display label, or null.
	LabelPlural string `json:"labelPlural" api:"required"`
	// Singular display label.
	LabelSingular string `json:"labelSingular" api:"required"`
	// Custom object slug used in API paths.
	Slug string `json:"slug" api:"required"`
	// ISO 8601 last-update timestamp.
	UpdatedAt string `json:"updatedAt" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID            respjson.Field
		Archived      respjson.Field
		ArchivedAt    respjson.Field
		CreatedAt     respjson.Field
		Description   respjson.Field
		Icon          respjson.Field
		LabelPlural   respjson.Field
		LabelSingular respjson.Field
		Slug          respjson.Field
		UpdatedAt     respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ObjectCreateTypeResponse) RawJSON() string { return r.JSON.raw }
func (r *ObjectCreateTypeResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ObjectDefinitionsResponse struct {
	// Map of field keys to their definitions, including both system and custom fields.
	FieldDefinitions map[string]ObjectDefinitionsResponseFieldDefinition `json:"fieldDefinitions" api:"required"`
	// The object type these definitions belong to (e.g. `account`).
	ObjectType string `json:"objectType" api:"required"`
	// Map of relationship keys to their definitions. Self-referential relationships
	// include optional directional metadata on the bare definition.
	RelationshipDefinitions map[string]ObjectDefinitionsResponseRelationshipDefinition `json:"relationshipDefinitions" api:"required"`
	// Map of relationship group keys to their definitions. A group collects
	// relationships between the same two object types, such as the roles a contact can
	// hold on an opportunity. Relationships name their group with `groupKey`.
	RelationshipGroupDefinitions map[string]ObjectDefinitionsResponseRelationshipGroupDefinition `json:"relationshipGroupDefinitions" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FieldDefinitions             respjson.Field
		ObjectType                   respjson.Field
		RelationshipDefinitions      respjson.Field
		RelationshipGroupDefinitions respjson.Field
		ExtraFields                  map[string]respjson.Field
		raw                          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ObjectDefinitionsResponse) RawJSON() string { return r.JSON.raw }
func (r *ObjectDefinitionsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ObjectDefinitionsResponseFieldDefinition struct {
	// Whether the field definition is archived.
	Archived bool `json:"archived" api:"required"`
	// When the field definition was archived, or null if active.
	ArchivedAt time.Time `json:"archivedAt" api:"required" format:"date-time"`
	// Description of the field, or null.
	Description string `json:"description" api:"required"`
	// Human-readable display name of the field.
	Label string `json:"label" api:"required"`
	// Type-specific configuration (e.g. select options, currency code). A key the
	// field type does not support is absent; a supported key with nothing set is null.
	TypeConfiguration ObjectDefinitionsResponseFieldDefinitionTypeConfiguration `json:"typeConfiguration" api:"required"`
	// Data type of the field.
	//
	// Any of "ADDRESS", "CHECKBOX", "CURRENCY", "DATE", "DATETIME", "EMAIL",
	// "FULL_NAME", "MARKDOWN", "MULTI_SELECT", "NUMBER", "SINGLE_SELECT",
	// "SOCIAL_HANDLE", "TELEPHONE", "TEXT", "URL", "HTML".
	ValueType string `json:"valueType" api:"required"`
	// Unique identifier of the field definition.
	ID string `json:"id"`
	// Whether AI fills this field automatically (`ON`), only suggests values
	// (`SUGGEST`), or leaves it alone (`OFF`). Absent on built-in fields that have no
	// `id`.
	//
	// Any of "ON", "OFF", "SUGGEST".
	AIFillable string `json:"aiFillable"`
	// Instructions AI follows when filling this field, or null when none are set.
	// Absent on built-in fields that have no `id`.
	AIFillablePrompt string `json:"aiFillablePrompt" api:"nullable"`
	// Formula a calculated field evaluates, with fields and relationships referenced
	// by key. Null for fields that are not calculated. Absent on built-in fields that
	// have no `id`.
	Formula string `json:"formula" api:"nullable"`
	// `true` for fields that are not writable via the API (e.g. AI-generated
	// summaries, calculated fields). `false` or absent for writable fields.
	ReadOnly bool `json:"readOnly"`
	// Whether values are entered (`MANUAL`) or calculated from a formula (`COMPUTED`).
	// Absent on built-in fields that have no `id`.
	//
	// Any of "MANUAL", "COMPUTED".
	ValueSource string `json:"valueSource"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Archived          respjson.Field
		ArchivedAt        respjson.Field
		Description       respjson.Field
		Label             respjson.Field
		TypeConfiguration respjson.Field
		ValueType         respjson.Field
		ID                respjson.Field
		AIFillable        respjson.Field
		AIFillablePrompt  respjson.Field
		Formula           respjson.Field
		ReadOnly          respjson.Field
		ValueSource       respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ObjectDefinitionsResponseFieldDefinition) RawJSON() string { return r.JSON.raw }
func (r *ObjectDefinitionsResponseFieldDefinition) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Type-specific configuration (e.g. select options, currency code). A key the
// field type does not support is absent; a supported key with nothing set is null.
type ObjectDefinitionsResponseFieldDefinitionTypeConfiguration struct {
	// ISO 4217 3-letter currency code. Present on `CURRENCY` fields, absent on every
	// other type.
	Currency string `json:"currency"`
	// Number of decimal places values carry, or null when none is set. Present on
	// `NUMBER` and `CURRENCY` fields, absent on every other type.
	DecimalPlaces int64 `json:"decimalPlaces" api:"nullable"`
	// Social platform associated with this handle field, or null when none is set.
	// Present on `SOCIAL_HANDLE` fields, absent on every other type.
	//
	// Any of "TWITTER", "LINKEDIN", "FACEBOOK", "INSTAGRAM".
	HandleService string `json:"handleService" api:"nullable"`
	// Whether this field accepts multiple values. Present on `EMAIL`, `TELEPHONE` and
	// `URL` fields, absent on every other type.
	MultipleValues bool `json:"multipleValues"`
	// Available options. Always present on `SINGLE_SELECT` and `MULTI_SELECT` fields,
	// as an empty list when there are none. Absent on every other type.
	Options []ObjectDefinitionsResponseFieldDefinitionTypeConfigurationOption `json:"options"`
	// For a dependent select, the field key whose selected option determines which of
	// this field's options apply. Absent for ordinary selects.
	ParentFieldKey string `json:"parentFieldKey"`
	// Whether values for this field must be unique, or null when not set. Present on
	// `EMAIL`, `TELEPHONE`, `URL` and `SOCIAL_HANDLE` fields, absent on every other
	// type.
	Unique bool `json:"unique" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Currency       respjson.Field
		DecimalPlaces  respjson.Field
		HandleService  respjson.Field
		MultipleValues respjson.Field
		Options        respjson.Field
		ParentFieldKey respjson.Field
		Unique         respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ObjectDefinitionsResponseFieldDefinitionTypeConfiguration) RawJSON() string {
	return r.JSON.raw
}
func (r *ObjectDefinitionsResponseFieldDefinitionTypeConfiguration) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ObjectDefinitionsResponseFieldDefinitionTypeConfigurationOption struct {
	// Unique identifier of the select option.
	ID string `json:"id" api:"required"`
	// Human-readable display name of the option.
	Label string `json:"label" api:"required"`
	// Lifecycle group of the option, in-progress (`ACTIVE`) or done (`COMPLETED`),
	// `ACTIVE` when none was set. Absent on selects whose options are not grouped.
	//
	// Any of "ACTIVE", "COMPLETED".
	Category string `json:"category"`
	// Color the option is displayed in, `GRAY` when none was chosen. Absent on selects
	// whose options take no color.
	//
	// Any of "RED", "ORANGE", "YELLOW", "LIME", "GREEN", "BLUE", "INDIGO", "PURPLE",
	// "COPPER", "GRAY".
	Color string `json:"color"`
	// Description of the option, or null.
	Description string `json:"description" api:"nullable"`
	// Forecast weighting of the option from 0 to 100, or null when none is set. Absent
	// on selects whose options take no forecast.
	ForecastPercentage float64 `json:"forecastPercentage" api:"nullable"`
	// Icon the option is displayed with, `BACKLOG` when none was chosen. Present only
	// on $stage.
	//
	// Any of "BACKLOG", "QUALIFICATION", "DISCOVERY", "ACTIVE", "NEGOTIATION",
	// "CLOSED_WON", "CLOSED_LOST".
	Icon string `json:"icon"`
	// For a dependent select, the ID of the option in the parent field named by
	// `typeConfiguration.parentFieldKey` that owns this option. Absent for ordinary
	// selects.
	ParentID string `json:"parentId"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                 respjson.Field
		Label              respjson.Field
		Category           respjson.Field
		Color              respjson.Field
		Description        respjson.Field
		ForecastPercentage respjson.Field
		Icon               respjson.Field
		ParentID           respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ObjectDefinitionsResponseFieldDefinitionTypeConfigurationOption) RawJSON() string {
	return r.JSON.raw
}
func (r *ObjectDefinitionsResponseFieldDefinitionTypeConfigurationOption) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ObjectDefinitionsResponseRelationshipDefinition struct {
	// Whether the relationship definition is archived.
	Archived bool `json:"archived" api:"required"`
	// When the relationship definition was archived, or null if active.
	ArchivedAt time.Time `json:"archivedAt" api:"required" format:"date-time"`
	// Whether this is a `has_one` or `has_many` relationship.
	//
	// Any of "HAS_ONE", "HAS_MANY".
	Cardinality string `json:"cardinality" api:"required"`
	// Description of the relationship, or null.
	Description string `json:"description" api:"required"`
	// Human-readable display name of the relationship, as it reads from this object
	// type.
	Label string `json:"label" api:"required"`
	// The type of the related object (e.g. `account`, `contact`).
	ObjectType string `json:"objectType" api:"required"`
	// Unique identifier of the relationship definition.
	ID string `json:"id"`
	// Directional metadata for a self-referential relationship. List filters use the
	// relationship key with `:asSource` or `:asDestination`; writes continue to use
	// the bare key.
	Directions ObjectDefinitionsResponseRelationshipDefinitionDirections `json:"directions"`
	// Key of the relationship group this relationship belongs to, as listed under
	// `relationshipGroupDefinitions`. Absent when the relationship is not in a group.
	GroupKey string `json:"groupKey"`
	// How the same relationship reads from the related object type.
	Inverse ObjectDefinitionsResponseRelationshipDefinitionInverse `json:"inverse"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Archived    respjson.Field
		ArchivedAt  respjson.Field
		Cardinality respjson.Field
		Description respjson.Field
		Label       respjson.Field
		ObjectType  respjson.Field
		ID          respjson.Field
		Directions  respjson.Field
		GroupKey    respjson.Field
		Inverse     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ObjectDefinitionsResponseRelationshipDefinition) RawJSON() string { return r.JSON.raw }
func (r *ObjectDefinitionsResponseRelationshipDefinition) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Directional metadata for a self-referential relationship. List filters use the
// relationship key with `:asSource` or `:asDestination`; writes continue to use
// the bare key.
type ObjectDefinitionsResponseRelationshipDefinitionDirections struct {
	// The current record is the destination (incoming) endpoint.
	AsDestination ObjectDefinitionsResponseRelationshipDefinitionDirectionsAsDestination `json:"asDestination" api:"required"`
	// The current record is the source (outgoing) endpoint.
	AsSource ObjectDefinitionsResponseRelationshipDefinitionDirectionsAsSource `json:"asSource" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AsDestination respjson.Field
		AsSource      respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ObjectDefinitionsResponseRelationshipDefinitionDirections) RawJSON() string {
	return r.JSON.raw
}
func (r *ObjectDefinitionsResponseRelationshipDefinitionDirections) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The current record is the destination (incoming) endpoint.
type ObjectDefinitionsResponseRelationshipDefinitionDirectionsAsDestination struct {
	// Whether this endpoint direction is `HAS_ONE` or `HAS_MANY`.
	//
	// Any of "HAS_ONE", "HAS_MANY".
	Cardinality string `json:"cardinality" api:"required"`
	// Human-readable display name for this endpoint direction.
	Label string `json:"label" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Cardinality respjson.Field
		Label       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ObjectDefinitionsResponseRelationshipDefinitionDirectionsAsDestination) RawJSON() string {
	return r.JSON.raw
}
func (r *ObjectDefinitionsResponseRelationshipDefinitionDirectionsAsDestination) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The current record is the source (outgoing) endpoint.
type ObjectDefinitionsResponseRelationshipDefinitionDirectionsAsSource struct {
	// Whether this endpoint direction is `HAS_ONE` or `HAS_MANY`.
	//
	// Any of "HAS_ONE", "HAS_MANY".
	Cardinality string `json:"cardinality" api:"required"`
	// Human-readable display name for this endpoint direction.
	Label string `json:"label" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Cardinality respjson.Field
		Label       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ObjectDefinitionsResponseRelationshipDefinitionDirectionsAsSource) RawJSON() string {
	return r.JSON.raw
}
func (r *ObjectDefinitionsResponseRelationshipDefinitionDirectionsAsSource) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// How the same relationship reads from the related object type.
type ObjectDefinitionsResponseRelationshipDefinitionInverse struct {
	// Whether this endpoint direction is `HAS_ONE` or `HAS_MANY`.
	//
	// Any of "HAS_ONE", "HAS_MANY".
	Cardinality string `json:"cardinality" api:"required"`
	// Human-readable display name for this endpoint direction.
	Label string `json:"label" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Cardinality respjson.Field
		Label       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ObjectDefinitionsResponseRelationshipDefinitionInverse) RawJSON() string { return r.JSON.raw }
func (r *ObjectDefinitionsResponseRelationshipDefinitionInverse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ObjectDefinitionsResponseRelationshipGroupDefinition struct {
	// Unique identifier of the relationship group.
	ID string `json:"id" api:"required"`
	// Whether the relationship group is archived.
	Archived bool `json:"archived" api:"required"`
	// When the relationship group was archived, or null if active.
	ArchivedAt time.Time `json:"archivedAt" api:"required" format:"date-time"`
	// How the same group reads from the related object type.
	Inverse ObjectDefinitionsResponseRelationshipGroupDefinitionInverse `json:"inverse" api:"required"`
	// Human-readable display name of the group, as it reads from this object type.
	Label string `json:"label" api:"required"`
	// The type of the related object every relationship in the group points to.
	ObjectType string `json:"objectType" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Archived    respjson.Field
		ArchivedAt  respjson.Field
		Inverse     respjson.Field
		Label       respjson.Field
		ObjectType  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ObjectDefinitionsResponseRelationshipGroupDefinition) RawJSON() string { return r.JSON.raw }
func (r *ObjectDefinitionsResponseRelationshipGroupDefinition) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// How the same group reads from the related object type.
type ObjectDefinitionsResponseRelationshipGroupDefinitionInverse struct {
	// Human-readable display name of the group, as it reads from the related object
	// type.
	Label string `json:"label" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Label       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ObjectDefinitionsResponseRelationshipGroupDefinitionInverse) RawJSON() string {
	return r.JSON.raw
}
func (r *ObjectDefinitionsResponseRelationshipGroupDefinitionInverse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ObjectDeletePipelineResponse struct {
	// Number of records whose stage is, or would be, cleared.
	AffectedRecordCount int64 `json:"affectedRecordCount" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AffectedRecordCount respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ObjectDeletePipelineResponse) RawJSON() string { return r.JSON.raw }
func (r *ObjectDeletePipelineResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ObjectDeleteResponse struct {
	// Unique identifier for the entity.
	ID string `json:"id" api:"required"`
	// ISO 8601 timestamp of when the entity was created.
	CreatedAt string `json:"createdAt" api:"required"`
	// Map of field names to their typed values. System fields are prefixed with `$`
	// (e.g. `$name`, `$email`); custom attributes use their bare slug.
	Fields map[string]ObjectDeleteResponseField `json:"fields" api:"required"`
	// URL to view the entity in the Lightfield web app, or null.
	HTTPLink string `json:"httpLink" api:"required"`
	// Map of relationship names to their associated entities. System relationships are
	// prefixed with `$` (e.g. `$owner`, `$contact`). A self-referential relationship
	// keeps its merged values and includes optional directional values. List filters
	// select the current record's role with `:asSource` or `:asDestination`, for
	// example `?parent:asSource[contains]=account_id`.
	Relationships map[string]ObjectDeleteResponseRelationship `json:"relationships" api:"required"`
	// ISO 8601 timestamp of when the entity was last updated, or null.
	UpdatedAt string `json:"updatedAt" api:"required"`
	// External identifier for the entity, or null if unset.
	ExternalID string `json:"externalId" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID            respjson.Field
		CreatedAt     respjson.Field
		Fields        respjson.Field
		HTTPLink      respjson.Field
		Relationships respjson.Field
		UpdatedAt     respjson.Field
		ExternalID    respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ObjectDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *ObjectDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ObjectDeleteResponseField struct {
	// The field value, or null if unset.
	Value ObjectDeleteResponseFieldValueUnion `json:"value" api:"required"`
	// The data type of the field.
	//
	// Any of "ADDRESS", "CHECKBOX", "CURRENCY", "DATE", "DATETIME", "EMAIL",
	// "FULL_NAME", "MARKDOWN", "MULTI_SELECT", "NUMBER", "SINGLE_SELECT",
	// "SOCIAL_HANDLE", "TELEPHONE", "TEXT", "URL", "HTML".
	ValueType string `json:"valueType" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Value       respjson.Field
		ValueType   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ObjectDeleteResponseField) RawJSON() string { return r.JSON.raw }
func (r *ObjectDeleteResponseField) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ObjectDeleteResponseFieldValueUnion contains all possible properties and values
// from [string], [float64], [bool], [[]string],
// [ObjectDeleteResponseFieldValueAddress],
// [ObjectDeleteResponseFieldValueFullName].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfStringArray]
type ObjectDeleteResponseFieldValueUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]string] instead of an object.
	OfStringArray []string `json:",inline"`
	// This field is from variant [ObjectDeleteResponseFieldValueAddress].
	City string `json:"city"`
	// This field is from variant [ObjectDeleteResponseFieldValueAddress].
	Country string `json:"country"`
	// This field is from variant [ObjectDeleteResponseFieldValueAddress].
	Latitude float64 `json:"latitude"`
	// This field is from variant [ObjectDeleteResponseFieldValueAddress].
	Longitude float64 `json:"longitude"`
	// This field is from variant [ObjectDeleteResponseFieldValueAddress].
	PostalCode string `json:"postalCode"`
	// This field is from variant [ObjectDeleteResponseFieldValueAddress].
	State string `json:"state"`
	// This field is from variant [ObjectDeleteResponseFieldValueAddress].
	Street string `json:"street"`
	// This field is from variant [ObjectDeleteResponseFieldValueAddress].
	Street2 string `json:"street2"`
	// This field is from variant [ObjectDeleteResponseFieldValueFullName].
	FirstName string `json:"firstName"`
	// This field is from variant [ObjectDeleteResponseFieldValueFullName].
	LastName string `json:"lastName"`
	JSON     struct {
		OfString      respjson.Field
		OfFloat       respjson.Field
		OfBool        respjson.Field
		OfStringArray respjson.Field
		City          respjson.Field
		Country       respjson.Field
		Latitude      respjson.Field
		Longitude     respjson.Field
		PostalCode    respjson.Field
		State         respjson.Field
		Street        respjson.Field
		Street2       respjson.Field
		FirstName     respjson.Field
		LastName      respjson.Field
		raw           string
	} `json:"-"`
}

func (u ObjectDeleteResponseFieldValueUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ObjectDeleteResponseFieldValueUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ObjectDeleteResponseFieldValueUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ObjectDeleteResponseFieldValueUnion) AsStringArray() (v []string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ObjectDeleteResponseFieldValueUnion) AsAddress() (v ObjectDeleteResponseFieldValueAddress) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ObjectDeleteResponseFieldValueUnion) AsFullName() (v ObjectDeleteResponseFieldValueFullName) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ObjectDeleteResponseFieldValueUnion) RawJSON() string { return u.JSON.raw }

func (r *ObjectDeleteResponseFieldValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ObjectDeleteResponseFieldValueAddress struct {
	// City name.
	City string `json:"city" api:"nullable"`
	// 2-letter ISO 3166-1 alpha-2 country code.
	Country string `json:"country" api:"nullable"`
	// Latitude coordinate.
	Latitude float64 `json:"latitude" api:"nullable"`
	// Longitude coordinate.
	Longitude float64 `json:"longitude" api:"nullable"`
	// Postal or ZIP code.
	PostalCode string `json:"postalCode" api:"nullable"`
	// State or province.
	State string `json:"state" api:"nullable"`
	// Street address line 1.
	Street string `json:"street" api:"nullable"`
	// Street address line 2.
	Street2 string `json:"street2" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		City        respjson.Field
		Country     respjson.Field
		Latitude    respjson.Field
		Longitude   respjson.Field
		PostalCode  respjson.Field
		State       respjson.Field
		Street      respjson.Field
		Street2     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ObjectDeleteResponseFieldValueAddress) RawJSON() string { return r.JSON.raw }
func (r *ObjectDeleteResponseFieldValueAddress) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ObjectDeleteResponseFieldValueFullName struct {
	// The contact's first name.
	FirstName string `json:"firstName" api:"nullable"`
	// The contact's last name.
	LastName string `json:"lastName" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FirstName   respjson.Field
		LastName    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ObjectDeleteResponseFieldValueFullName) RawJSON() string { return r.JSON.raw }
func (r *ObjectDeleteResponseFieldValueFullName) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ObjectDeleteResponseRelationship struct {
	// Whether the relationship is `has_one` or `has_many`.
	Cardinality string `json:"cardinality" api:"required"`
	// The type of the related object (e.g. `account`, `contact`).
	ObjectType string `json:"objectType" api:"required"`
	// IDs of the related entities.
	Values []string `json:"values" api:"required"`
	// Directional values for a self-referential relationship. List filters use the
	// relationship key with `:asSource` or `:asDestination`; writes continue to use
	// the bare key.
	Directions ObjectDeleteResponseRelationshipDirections `json:"directions"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Cardinality respjson.Field
		ObjectType  respjson.Field
		Values      respjson.Field
		Directions  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ObjectDeleteResponseRelationship) RawJSON() string { return r.JSON.raw }
func (r *ObjectDeleteResponseRelationship) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Directional values for a self-referential relationship. List filters use the
// relationship key with `:asSource` or `:asDestination`; writes continue to use
// the bare key.
type ObjectDeleteResponseRelationshipDirections struct {
	// The current record is the destination (incoming) endpoint.
	AsDestination ObjectDeleteResponseRelationshipDirectionsAsDestination `json:"asDestination" api:"required"`
	// The current record is the source (outgoing) endpoint.
	AsSource ObjectDeleteResponseRelationshipDirectionsAsSource `json:"asSource" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AsDestination respjson.Field
		AsSource      respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ObjectDeleteResponseRelationshipDirections) RawJSON() string { return r.JSON.raw }
func (r *ObjectDeleteResponseRelationshipDirections) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The current record is the destination (incoming) endpoint.
type ObjectDeleteResponseRelationshipDirectionsAsDestination struct {
	// Whether this endpoint direction is `HAS_ONE` or `HAS_MANY`.
	//
	// Any of "HAS_ONE", "HAS_MANY".
	Cardinality string `json:"cardinality" api:"required"`
	// IDs of the related entities in this direction.
	Values []string `json:"values" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Cardinality respjson.Field
		Values      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ObjectDeleteResponseRelationshipDirectionsAsDestination) RawJSON() string { return r.JSON.raw }
func (r *ObjectDeleteResponseRelationshipDirectionsAsDestination) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The current record is the source (outgoing) endpoint.
type ObjectDeleteResponseRelationshipDirectionsAsSource struct {
	// Whether this endpoint direction is `HAS_ONE` or `HAS_MANY`.
	//
	// Any of "HAS_ONE", "HAS_MANY".
	Cardinality string `json:"cardinality" api:"required"`
	// IDs of the related entities in this direction.
	Values []string `json:"values" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Cardinality respjson.Field
		Values      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ObjectDeleteResponseRelationshipDirectionsAsSource) RawJSON() string { return r.JSON.raw }
func (r *ObjectDeleteResponseRelationshipDirectionsAsSource) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ObjectFieldHistoryResponse struct {
	// Recorded values for the field, newest first.
	Data []ObjectFieldHistoryResponseData `json:"data" api:"required"`
	// Whether more history exists beyond this page.
	HasMore bool `json:"hasMore" api:"required"`
	// Cursor to pass as `after` to fetch the next page, or null when there are no more
	// entries.
	NextCursor string `json:"nextCursor" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		HasMore     respjson.Field
		NextCursor  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ObjectFieldHistoryResponse) RawJSON() string { return r.JSON.raw }
func (r *ObjectFieldHistoryResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A single recorded version of a field value.
type ObjectFieldHistoryResponseData struct {
	// Human-readable rendering of the value (e.g. a select option label), suitable for
	// display.
	DisplayValue string `json:"displayValue" api:"required"`
	// True for the record’s original value. Only set when the full history fits in the
	// response (never on a truncated/paginated page).
	IsCreate bool `json:"isCreate" api:"required"`
	// ISO 8601 timestamp of when this value was recorded.
	RecordedAt string `json:"recordedAt" api:"required"`
	// The field value, or null if unset.
	Value ObjectFieldHistoryResponseDataValueUnion `json:"value" api:"required"`
	// The data type of the field.
	//
	// Any of "ADDRESS", "CHECKBOX", "CURRENCY", "DATE", "DATETIME", "EMAIL",
	// "FULL_NAME", "MARKDOWN", "MULTI_SELECT", "NUMBER", "SINGLE_SELECT",
	// "SOCIAL_HANDLE", "TELEPHONE", "TEXT", "URL", "HTML".
	ValueType string `json:"valueType" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DisplayValue respjson.Field
		IsCreate     respjson.Field
		RecordedAt   respjson.Field
		Value        respjson.Field
		ValueType    respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ObjectFieldHistoryResponseData) RawJSON() string { return r.JSON.raw }
func (r *ObjectFieldHistoryResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ObjectFieldHistoryResponseDataValueUnion contains all possible properties and
// values from [string], [float64], [bool], [[]string],
// [ObjectFieldHistoryResponseDataValueAddress],
// [ObjectFieldHistoryResponseDataValueFullName].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfStringArray]
type ObjectFieldHistoryResponseDataValueUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]string] instead of an object.
	OfStringArray []string `json:",inline"`
	// This field is from variant [ObjectFieldHistoryResponseDataValueAddress].
	City string `json:"city"`
	// This field is from variant [ObjectFieldHistoryResponseDataValueAddress].
	Country string `json:"country"`
	// This field is from variant [ObjectFieldHistoryResponseDataValueAddress].
	Latitude float64 `json:"latitude"`
	// This field is from variant [ObjectFieldHistoryResponseDataValueAddress].
	Longitude float64 `json:"longitude"`
	// This field is from variant [ObjectFieldHistoryResponseDataValueAddress].
	PostalCode string `json:"postalCode"`
	// This field is from variant [ObjectFieldHistoryResponseDataValueAddress].
	State string `json:"state"`
	// This field is from variant [ObjectFieldHistoryResponseDataValueAddress].
	Street string `json:"street"`
	// This field is from variant [ObjectFieldHistoryResponseDataValueAddress].
	Street2 string `json:"street2"`
	// This field is from variant [ObjectFieldHistoryResponseDataValueFullName].
	FirstName string `json:"firstName"`
	// This field is from variant [ObjectFieldHistoryResponseDataValueFullName].
	LastName string `json:"lastName"`
	JSON     struct {
		OfString      respjson.Field
		OfFloat       respjson.Field
		OfBool        respjson.Field
		OfStringArray respjson.Field
		City          respjson.Field
		Country       respjson.Field
		Latitude      respjson.Field
		Longitude     respjson.Field
		PostalCode    respjson.Field
		State         respjson.Field
		Street        respjson.Field
		Street2       respjson.Field
		FirstName     respjson.Field
		LastName      respjson.Field
		raw           string
	} `json:"-"`
}

func (u ObjectFieldHistoryResponseDataValueUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ObjectFieldHistoryResponseDataValueUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ObjectFieldHistoryResponseDataValueUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ObjectFieldHistoryResponseDataValueUnion) AsStringArray() (v []string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ObjectFieldHistoryResponseDataValueUnion) AsAddress() (v ObjectFieldHistoryResponseDataValueAddress) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ObjectFieldHistoryResponseDataValueUnion) AsFullName() (v ObjectFieldHistoryResponseDataValueFullName) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ObjectFieldHistoryResponseDataValueUnion) RawJSON() string { return u.JSON.raw }

func (r *ObjectFieldHistoryResponseDataValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ObjectFieldHistoryResponseDataValueAddress struct {
	// City name.
	City string `json:"city" api:"nullable"`
	// 2-letter ISO 3166-1 alpha-2 country code.
	Country string `json:"country" api:"nullable"`
	// Latitude coordinate.
	Latitude float64 `json:"latitude" api:"nullable"`
	// Longitude coordinate.
	Longitude float64 `json:"longitude" api:"nullable"`
	// Postal or ZIP code.
	PostalCode string `json:"postalCode" api:"nullable"`
	// State or province.
	State string `json:"state" api:"nullable"`
	// Street address line 1.
	Street string `json:"street" api:"nullable"`
	// Street address line 2.
	Street2 string `json:"street2" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		City        respjson.Field
		Country     respjson.Field
		Latitude    respjson.Field
		Longitude   respjson.Field
		PostalCode  respjson.Field
		State       respjson.Field
		Street      respjson.Field
		Street2     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ObjectFieldHistoryResponseDataValueAddress) RawJSON() string { return r.JSON.raw }
func (r *ObjectFieldHistoryResponseDataValueAddress) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ObjectFieldHistoryResponseDataValueFullName struct {
	// The contact's first name.
	FirstName string `json:"firstName" api:"nullable"`
	// The contact's last name.
	LastName string `json:"lastName" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FirstName   respjson.Field
		LastName    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ObjectFieldHistoryResponseDataValueFullName) RawJSON() string { return r.JSON.raw }
func (r *ObjectFieldHistoryResponseDataValueFullName) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ObjectListDefinitionsResponse struct {
	// All object types available to the caller.
	Data []ObjectListDefinitionsResponseData `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ObjectListDefinitionsResponse) RawJSON() string { return r.JSON.raw }
func (r *ObjectListDefinitionsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ObjectListDefinitionsResponseData struct {
	// Human-readable display name.
	Label string `json:"label" api:"required"`
	// The slug used to reference this object type in the API.
	ObjectType string `json:"objectType" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Label       respjson.Field
		ObjectType  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ObjectListDefinitionsResponseData) RawJSON() string { return r.JSON.raw }
func (r *ObjectListDefinitionsResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ObjectListPipelinesResponse struct {
	Data []ObjectListPipelinesResponseData `json:"data" api:"required"`
	// Any of "list".
	Object ObjectListPipelinesResponseObject `json:"object" api:"required"`
	// Number of pipelines on the object type.
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
func (r ObjectListPipelinesResponse) RawJSON() string { return r.JSON.raw }
func (r *ObjectListPipelinesResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ObjectListPipelinesResponseData struct {
	// Unique identifier of the pipeline.
	ID string `json:"id" api:"required"`
	// Description of the pipeline, or null.
	Description string `json:"description" api:"required"`
	// Whether stages carry a `forecastPercentage`. Turning it off clears every stage
	// forecast.
	ForecastEnabled bool `json:"forecastEnabled" api:"required"`
	// Display name of the pipeline.
	Name string `json:"name" api:"required"`
	// Key derived from the name when the pipeline was created, unique on the object
	// type. Pipelines provisioned by Lightfield use a `$` prefix (e.g.
	// `$default-pipeline`).
	Slug string `json:"slug" api:"required"`
	// Stages in display order. Each `id` is the option ID records hold in `$stage`.
	Stages []ObjectListPipelinesResponseDataStage `json:"stages" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID              respjson.Field
		Description     respjson.Field
		ForecastEnabled respjson.Field
		Name            respjson.Field
		Slug            respjson.Field
		Stages          respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ObjectListPipelinesResponseData) RawJSON() string { return r.JSON.raw }
func (r *ObjectListPipelinesResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ObjectListPipelinesResponseDataStage struct {
	// Unique identifier of the select option.
	ID string `json:"id" api:"required"`
	// Human-readable display name of the option.
	Label string `json:"label" api:"required"`
	// Lifecycle group of the option, in-progress (`ACTIVE`) or done (`COMPLETED`),
	// `ACTIVE` when none was set. Absent on selects whose options are not grouped.
	//
	// Any of "ACTIVE", "COMPLETED".
	Category string `json:"category"`
	// Color the option is displayed in, `GRAY` when none was chosen. Absent on selects
	// whose options take no color.
	//
	// Any of "RED", "ORANGE", "YELLOW", "LIME", "GREEN", "BLUE", "INDIGO", "PURPLE",
	// "COPPER", "GRAY".
	Color string `json:"color"`
	// Description of the option, or null.
	Description string `json:"description" api:"nullable"`
	// Forecast weighting of the option from 0 to 100, or null when none is set. Absent
	// on selects whose options take no forecast.
	ForecastPercentage float64 `json:"forecastPercentage" api:"nullable"`
	// Icon the option is displayed with, `BACKLOG` when none was chosen. Present only
	// on $stage.
	//
	// Any of "BACKLOG", "QUALIFICATION", "DISCOVERY", "ACTIVE", "NEGOTIATION",
	// "CLOSED_WON", "CLOSED_LOST".
	Icon string `json:"icon"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                 respjson.Field
		Label              respjson.Field
		Category           respjson.Field
		Color              respjson.Field
		Description        respjson.Field
		ForecastPercentage respjson.Field
		Icon               respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ObjectListPipelinesResponseDataStage) RawJSON() string { return r.JSON.raw }
func (r *ObjectListPipelinesResponseDataStage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ObjectListPipelinesResponseObject string

const (
	ObjectListPipelinesResponseObjectList ObjectListPipelinesResponseObject = "list"
)

type ObjectListResponse struct {
	// Array of entity objects for the current page.
	Data []ObjectListResponseData `json:"data" api:"required"`
	// The object type, always `"list"`.
	Object string `json:"object" api:"required"`
	// Total number of entities matching the query.
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
func (r ObjectListResponse) RawJSON() string { return r.JSON.raw }
func (r *ObjectListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ObjectListResponseData struct {
	// Unique identifier for the entity.
	ID string `json:"id" api:"required"`
	// ISO 8601 timestamp of when the entity was created.
	CreatedAt string `json:"createdAt" api:"required"`
	// Map of field names to their typed values. System fields are prefixed with `$`
	// (e.g. `$name`, `$email`); custom attributes use their bare slug.
	Fields map[string]ObjectListResponseDataField `json:"fields" api:"required"`
	// URL to view the entity in the Lightfield web app, or null.
	HTTPLink string `json:"httpLink" api:"required"`
	// Map of relationship names to their associated entities. System relationships are
	// prefixed with `$` (e.g. `$owner`, `$contact`). A self-referential relationship
	// keeps its merged values and includes optional directional values. List filters
	// select the current record's role with `:asSource` or `:asDestination`, for
	// example `?parent:asSource[contains]=account_id`.
	Relationships map[string]ObjectListResponseDataRelationship `json:"relationships" api:"required"`
	// ISO 8601 timestamp of when the entity was last updated, or null.
	UpdatedAt string `json:"updatedAt" api:"required"`
	// External identifier for the entity, or null if unset.
	ExternalID string `json:"externalId" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID            respjson.Field
		CreatedAt     respjson.Field
		Fields        respjson.Field
		HTTPLink      respjson.Field
		Relationships respjson.Field
		UpdatedAt     respjson.Field
		ExternalID    respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ObjectListResponseData) RawJSON() string { return r.JSON.raw }
func (r *ObjectListResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ObjectListResponseDataField struct {
	// The field value, or null if unset.
	Value ObjectListResponseDataFieldValueUnion `json:"value" api:"required"`
	// The data type of the field.
	//
	// Any of "ADDRESS", "CHECKBOX", "CURRENCY", "DATE", "DATETIME", "EMAIL",
	// "FULL_NAME", "MARKDOWN", "MULTI_SELECT", "NUMBER", "SINGLE_SELECT",
	// "SOCIAL_HANDLE", "TELEPHONE", "TEXT", "URL", "HTML".
	ValueType string `json:"valueType" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Value       respjson.Field
		ValueType   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ObjectListResponseDataField) RawJSON() string { return r.JSON.raw }
func (r *ObjectListResponseDataField) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ObjectListResponseDataFieldValueUnion contains all possible properties and
// values from [string], [float64], [bool], [[]string],
// [ObjectListResponseDataFieldValueAddress],
// [ObjectListResponseDataFieldValueFullName].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfStringArray]
type ObjectListResponseDataFieldValueUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]string] instead of an object.
	OfStringArray []string `json:",inline"`
	// This field is from variant [ObjectListResponseDataFieldValueAddress].
	City string `json:"city"`
	// This field is from variant [ObjectListResponseDataFieldValueAddress].
	Country string `json:"country"`
	// This field is from variant [ObjectListResponseDataFieldValueAddress].
	Latitude float64 `json:"latitude"`
	// This field is from variant [ObjectListResponseDataFieldValueAddress].
	Longitude float64 `json:"longitude"`
	// This field is from variant [ObjectListResponseDataFieldValueAddress].
	PostalCode string `json:"postalCode"`
	// This field is from variant [ObjectListResponseDataFieldValueAddress].
	State string `json:"state"`
	// This field is from variant [ObjectListResponseDataFieldValueAddress].
	Street string `json:"street"`
	// This field is from variant [ObjectListResponseDataFieldValueAddress].
	Street2 string `json:"street2"`
	// This field is from variant [ObjectListResponseDataFieldValueFullName].
	FirstName string `json:"firstName"`
	// This field is from variant [ObjectListResponseDataFieldValueFullName].
	LastName string `json:"lastName"`
	JSON     struct {
		OfString      respjson.Field
		OfFloat       respjson.Field
		OfBool        respjson.Field
		OfStringArray respjson.Field
		City          respjson.Field
		Country       respjson.Field
		Latitude      respjson.Field
		Longitude     respjson.Field
		PostalCode    respjson.Field
		State         respjson.Field
		Street        respjson.Field
		Street2       respjson.Field
		FirstName     respjson.Field
		LastName      respjson.Field
		raw           string
	} `json:"-"`
}

func (u ObjectListResponseDataFieldValueUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ObjectListResponseDataFieldValueUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ObjectListResponseDataFieldValueUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ObjectListResponseDataFieldValueUnion) AsStringArray() (v []string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ObjectListResponseDataFieldValueUnion) AsAddress() (v ObjectListResponseDataFieldValueAddress) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ObjectListResponseDataFieldValueUnion) AsFullName() (v ObjectListResponseDataFieldValueFullName) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ObjectListResponseDataFieldValueUnion) RawJSON() string { return u.JSON.raw }

func (r *ObjectListResponseDataFieldValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ObjectListResponseDataFieldValueAddress struct {
	// City name.
	City string `json:"city" api:"nullable"`
	// 2-letter ISO 3166-1 alpha-2 country code.
	Country string `json:"country" api:"nullable"`
	// Latitude coordinate.
	Latitude float64 `json:"latitude" api:"nullable"`
	// Longitude coordinate.
	Longitude float64 `json:"longitude" api:"nullable"`
	// Postal or ZIP code.
	PostalCode string `json:"postalCode" api:"nullable"`
	// State or province.
	State string `json:"state" api:"nullable"`
	// Street address line 1.
	Street string `json:"street" api:"nullable"`
	// Street address line 2.
	Street2 string `json:"street2" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		City        respjson.Field
		Country     respjson.Field
		Latitude    respjson.Field
		Longitude   respjson.Field
		PostalCode  respjson.Field
		State       respjson.Field
		Street      respjson.Field
		Street2     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ObjectListResponseDataFieldValueAddress) RawJSON() string { return r.JSON.raw }
func (r *ObjectListResponseDataFieldValueAddress) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ObjectListResponseDataFieldValueFullName struct {
	// The contact's first name.
	FirstName string `json:"firstName" api:"nullable"`
	// The contact's last name.
	LastName string `json:"lastName" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FirstName   respjson.Field
		LastName    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ObjectListResponseDataFieldValueFullName) RawJSON() string { return r.JSON.raw }
func (r *ObjectListResponseDataFieldValueFullName) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ObjectListResponseDataRelationship struct {
	// Whether the relationship is `has_one` or `has_many`.
	Cardinality string `json:"cardinality" api:"required"`
	// The type of the related object (e.g. `account`, `contact`).
	ObjectType string `json:"objectType" api:"required"`
	// IDs of the related entities.
	Values []string `json:"values" api:"required"`
	// Directional values for a self-referential relationship. List filters use the
	// relationship key with `:asSource` or `:asDestination`; writes continue to use
	// the bare key.
	Directions ObjectListResponseDataRelationshipDirections `json:"directions"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Cardinality respjson.Field
		ObjectType  respjson.Field
		Values      respjson.Field
		Directions  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ObjectListResponseDataRelationship) RawJSON() string { return r.JSON.raw }
func (r *ObjectListResponseDataRelationship) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Directional values for a self-referential relationship. List filters use the
// relationship key with `:asSource` or `:asDestination`; writes continue to use
// the bare key.
type ObjectListResponseDataRelationshipDirections struct {
	// The current record is the destination (incoming) endpoint.
	AsDestination ObjectListResponseDataRelationshipDirectionsAsDestination `json:"asDestination" api:"required"`
	// The current record is the source (outgoing) endpoint.
	AsSource ObjectListResponseDataRelationshipDirectionsAsSource `json:"asSource" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AsDestination respjson.Field
		AsSource      respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ObjectListResponseDataRelationshipDirections) RawJSON() string { return r.JSON.raw }
func (r *ObjectListResponseDataRelationshipDirections) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The current record is the destination (incoming) endpoint.
type ObjectListResponseDataRelationshipDirectionsAsDestination struct {
	// Whether this endpoint direction is `HAS_ONE` or `HAS_MANY`.
	//
	// Any of "HAS_ONE", "HAS_MANY".
	Cardinality string `json:"cardinality" api:"required"`
	// IDs of the related entities in this direction.
	Values []string `json:"values" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Cardinality respjson.Field
		Values      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ObjectListResponseDataRelationshipDirectionsAsDestination) RawJSON() string {
	return r.JSON.raw
}
func (r *ObjectListResponseDataRelationshipDirectionsAsDestination) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The current record is the source (outgoing) endpoint.
type ObjectListResponseDataRelationshipDirectionsAsSource struct {
	// Whether this endpoint direction is `HAS_ONE` or `HAS_MANY`.
	//
	// Any of "HAS_ONE", "HAS_MANY".
	Cardinality string `json:"cardinality" api:"required"`
	// IDs of the related entities in this direction.
	Values []string `json:"values" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Cardinality respjson.Field
		Values      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ObjectListResponseDataRelationshipDirectionsAsSource) RawJSON() string { return r.JSON.raw }
func (r *ObjectListResponseDataRelationshipDirectionsAsSource) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ObjectRestoreResponse struct {
	// Unique identifier for the entity.
	ID string `json:"id" api:"required"`
	// ISO 8601 timestamp of when the entity was created.
	CreatedAt string `json:"createdAt" api:"required"`
	// Map of field names to their typed values. System fields are prefixed with `$`
	// (e.g. `$name`, `$email`); custom attributes use their bare slug.
	Fields map[string]ObjectRestoreResponseField `json:"fields" api:"required"`
	// URL to view the entity in the Lightfield web app, or null.
	HTTPLink string `json:"httpLink" api:"required"`
	// Map of relationship names to their associated entities. System relationships are
	// prefixed with `$` (e.g. `$owner`, `$contact`). A self-referential relationship
	// keeps its merged values and includes optional directional values. List filters
	// select the current record's role with `:asSource` or `:asDestination`, for
	// example `?parent:asSource[contains]=account_id`.
	Relationships map[string]ObjectRestoreResponseRelationship `json:"relationships" api:"required"`
	// ISO 8601 timestamp of when the entity was last updated, or null.
	UpdatedAt string `json:"updatedAt" api:"required"`
	// External identifier for the entity, or null if unset.
	ExternalID string `json:"externalId" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID            respjson.Field
		CreatedAt     respjson.Field
		Fields        respjson.Field
		HTTPLink      respjson.Field
		Relationships respjson.Field
		UpdatedAt     respjson.Field
		ExternalID    respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ObjectRestoreResponse) RawJSON() string { return r.JSON.raw }
func (r *ObjectRestoreResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ObjectRestoreResponseField struct {
	// The field value, or null if unset.
	Value ObjectRestoreResponseFieldValueUnion `json:"value" api:"required"`
	// The data type of the field.
	//
	// Any of "ADDRESS", "CHECKBOX", "CURRENCY", "DATE", "DATETIME", "EMAIL",
	// "FULL_NAME", "MARKDOWN", "MULTI_SELECT", "NUMBER", "SINGLE_SELECT",
	// "SOCIAL_HANDLE", "TELEPHONE", "TEXT", "URL", "HTML".
	ValueType string `json:"valueType" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Value       respjson.Field
		ValueType   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ObjectRestoreResponseField) RawJSON() string { return r.JSON.raw }
func (r *ObjectRestoreResponseField) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ObjectRestoreResponseFieldValueUnion contains all possible properties and values
// from [string], [float64], [bool], [[]string],
// [ObjectRestoreResponseFieldValueAddress],
// [ObjectRestoreResponseFieldValueFullName].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfStringArray]
type ObjectRestoreResponseFieldValueUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]string] instead of an object.
	OfStringArray []string `json:",inline"`
	// This field is from variant [ObjectRestoreResponseFieldValueAddress].
	City string `json:"city"`
	// This field is from variant [ObjectRestoreResponseFieldValueAddress].
	Country string `json:"country"`
	// This field is from variant [ObjectRestoreResponseFieldValueAddress].
	Latitude float64 `json:"latitude"`
	// This field is from variant [ObjectRestoreResponseFieldValueAddress].
	Longitude float64 `json:"longitude"`
	// This field is from variant [ObjectRestoreResponseFieldValueAddress].
	PostalCode string `json:"postalCode"`
	// This field is from variant [ObjectRestoreResponseFieldValueAddress].
	State string `json:"state"`
	// This field is from variant [ObjectRestoreResponseFieldValueAddress].
	Street string `json:"street"`
	// This field is from variant [ObjectRestoreResponseFieldValueAddress].
	Street2 string `json:"street2"`
	// This field is from variant [ObjectRestoreResponseFieldValueFullName].
	FirstName string `json:"firstName"`
	// This field is from variant [ObjectRestoreResponseFieldValueFullName].
	LastName string `json:"lastName"`
	JSON     struct {
		OfString      respjson.Field
		OfFloat       respjson.Field
		OfBool        respjson.Field
		OfStringArray respjson.Field
		City          respjson.Field
		Country       respjson.Field
		Latitude      respjson.Field
		Longitude     respjson.Field
		PostalCode    respjson.Field
		State         respjson.Field
		Street        respjson.Field
		Street2       respjson.Field
		FirstName     respjson.Field
		LastName      respjson.Field
		raw           string
	} `json:"-"`
}

func (u ObjectRestoreResponseFieldValueUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ObjectRestoreResponseFieldValueUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ObjectRestoreResponseFieldValueUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ObjectRestoreResponseFieldValueUnion) AsStringArray() (v []string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ObjectRestoreResponseFieldValueUnion) AsAddress() (v ObjectRestoreResponseFieldValueAddress) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ObjectRestoreResponseFieldValueUnion) AsFullName() (v ObjectRestoreResponseFieldValueFullName) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ObjectRestoreResponseFieldValueUnion) RawJSON() string { return u.JSON.raw }

func (r *ObjectRestoreResponseFieldValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ObjectRestoreResponseFieldValueAddress struct {
	// City name.
	City string `json:"city" api:"nullable"`
	// 2-letter ISO 3166-1 alpha-2 country code.
	Country string `json:"country" api:"nullable"`
	// Latitude coordinate.
	Latitude float64 `json:"latitude" api:"nullable"`
	// Longitude coordinate.
	Longitude float64 `json:"longitude" api:"nullable"`
	// Postal or ZIP code.
	PostalCode string `json:"postalCode" api:"nullable"`
	// State or province.
	State string `json:"state" api:"nullable"`
	// Street address line 1.
	Street string `json:"street" api:"nullable"`
	// Street address line 2.
	Street2 string `json:"street2" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		City        respjson.Field
		Country     respjson.Field
		Latitude    respjson.Field
		Longitude   respjson.Field
		PostalCode  respjson.Field
		State       respjson.Field
		Street      respjson.Field
		Street2     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ObjectRestoreResponseFieldValueAddress) RawJSON() string { return r.JSON.raw }
func (r *ObjectRestoreResponseFieldValueAddress) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ObjectRestoreResponseFieldValueFullName struct {
	// The contact's first name.
	FirstName string `json:"firstName" api:"nullable"`
	// The contact's last name.
	LastName string `json:"lastName" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FirstName   respjson.Field
		LastName    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ObjectRestoreResponseFieldValueFullName) RawJSON() string { return r.JSON.raw }
func (r *ObjectRestoreResponseFieldValueFullName) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ObjectRestoreResponseRelationship struct {
	// Whether the relationship is `has_one` or `has_many`.
	Cardinality string `json:"cardinality" api:"required"`
	// The type of the related object (e.g. `account`, `contact`).
	ObjectType string `json:"objectType" api:"required"`
	// IDs of the related entities.
	Values []string `json:"values" api:"required"`
	// Directional values for a self-referential relationship. List filters use the
	// relationship key with `:asSource` or `:asDestination`; writes continue to use
	// the bare key.
	Directions ObjectRestoreResponseRelationshipDirections `json:"directions"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Cardinality respjson.Field
		ObjectType  respjson.Field
		Values      respjson.Field
		Directions  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ObjectRestoreResponseRelationship) RawJSON() string { return r.JSON.raw }
func (r *ObjectRestoreResponseRelationship) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Directional values for a self-referential relationship. List filters use the
// relationship key with `:asSource` or `:asDestination`; writes continue to use
// the bare key.
type ObjectRestoreResponseRelationshipDirections struct {
	// The current record is the destination (incoming) endpoint.
	AsDestination ObjectRestoreResponseRelationshipDirectionsAsDestination `json:"asDestination" api:"required"`
	// The current record is the source (outgoing) endpoint.
	AsSource ObjectRestoreResponseRelationshipDirectionsAsSource `json:"asSource" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AsDestination respjson.Field
		AsSource      respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ObjectRestoreResponseRelationshipDirections) RawJSON() string { return r.JSON.raw }
func (r *ObjectRestoreResponseRelationshipDirections) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The current record is the destination (incoming) endpoint.
type ObjectRestoreResponseRelationshipDirectionsAsDestination struct {
	// Whether this endpoint direction is `HAS_ONE` or `HAS_MANY`.
	//
	// Any of "HAS_ONE", "HAS_MANY".
	Cardinality string `json:"cardinality" api:"required"`
	// IDs of the related entities in this direction.
	Values []string `json:"values" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Cardinality respjson.Field
		Values      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ObjectRestoreResponseRelationshipDirectionsAsDestination) RawJSON() string { return r.JSON.raw }
func (r *ObjectRestoreResponseRelationshipDirectionsAsDestination) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The current record is the source (outgoing) endpoint.
type ObjectRestoreResponseRelationshipDirectionsAsSource struct {
	// Whether this endpoint direction is `HAS_ONE` or `HAS_MANY`.
	//
	// Any of "HAS_ONE", "HAS_MANY".
	Cardinality string `json:"cardinality" api:"required"`
	// IDs of the related entities in this direction.
	Values []string `json:"values" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Cardinality respjson.Field
		Values      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ObjectRestoreResponseRelationshipDirectionsAsSource) RawJSON() string { return r.JSON.raw }
func (r *ObjectRestoreResponseRelationshipDirectionsAsSource) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ObjectRetrievePipelineResponse struct {
	// Unique identifier of the pipeline.
	ID string `json:"id" api:"required"`
	// Description of the pipeline, or null.
	Description string `json:"description" api:"required"`
	// Whether stages carry a `forecastPercentage`. Turning it off clears every stage
	// forecast.
	ForecastEnabled bool `json:"forecastEnabled" api:"required"`
	// Display name of the pipeline.
	Name string `json:"name" api:"required"`
	// Key derived from the name when the pipeline was created, unique on the object
	// type. Pipelines provisioned by Lightfield use a `$` prefix (e.g.
	// `$default-pipeline`).
	Slug string `json:"slug" api:"required"`
	// Stages in display order. Each `id` is the option ID records hold in `$stage`.
	Stages []ObjectRetrievePipelineResponseStage `json:"stages" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID              respjson.Field
		Description     respjson.Field
		ForecastEnabled respjson.Field
		Name            respjson.Field
		Slug            respjson.Field
		Stages          respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ObjectRetrievePipelineResponse) RawJSON() string { return r.JSON.raw }
func (r *ObjectRetrievePipelineResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ObjectRetrievePipelineResponseStage struct {
	// Unique identifier of the select option.
	ID string `json:"id" api:"required"`
	// Human-readable display name of the option.
	Label string `json:"label" api:"required"`
	// Lifecycle group of the option, in-progress (`ACTIVE`) or done (`COMPLETED`),
	// `ACTIVE` when none was set. Absent on selects whose options are not grouped.
	//
	// Any of "ACTIVE", "COMPLETED".
	Category string `json:"category"`
	// Color the option is displayed in, `GRAY` when none was chosen. Absent on selects
	// whose options take no color.
	//
	// Any of "RED", "ORANGE", "YELLOW", "LIME", "GREEN", "BLUE", "INDIGO", "PURPLE",
	// "COPPER", "GRAY".
	Color string `json:"color"`
	// Description of the option, or null.
	Description string `json:"description" api:"nullable"`
	// Forecast weighting of the option from 0 to 100, or null when none is set. Absent
	// on selects whose options take no forecast.
	ForecastPercentage float64 `json:"forecastPercentage" api:"nullable"`
	// Icon the option is displayed with, `BACKLOG` when none was chosen. Present only
	// on $stage.
	//
	// Any of "BACKLOG", "QUALIFICATION", "DISCOVERY", "ACTIVE", "NEGOTIATION",
	// "CLOSED_WON", "CLOSED_LOST".
	Icon string `json:"icon"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                 respjson.Field
		Label              respjson.Field
		Category           respjson.Field
		Color              respjson.Field
		Description        respjson.Field
		ForecastPercentage respjson.Field
		Icon               respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ObjectRetrievePipelineResponseStage) RawJSON() string { return r.JSON.raw }
func (r *ObjectRetrievePipelineResponseStage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ObjectRetrieveResponseUnion contains all possible properties and values from
// [ObjectRetrieveResponseEntity], [ObjectRetrieveResponseRestricted].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ObjectRetrieveResponseUnion struct {
	ID string `json:"id"`
	// This field is from variant [ObjectRetrieveResponseEntity].
	CreatedAt string `json:"createdAt"`
	// This field is from variant [ObjectRetrieveResponseEntity].
	Fields map[string]ObjectRetrieveResponseEntityField `json:"fields"`
	// This field is from variant [ObjectRetrieveResponseEntity].
	HTTPLink string `json:"httpLink"`
	// This field is from variant [ObjectRetrieveResponseEntity].
	Relationships map[string]ObjectRetrieveResponseEntityRelationship `json:"relationships"`
	// This field is from variant [ObjectRetrieveResponseEntity].
	UpdatedAt string `json:"updatedAt"`
	// This field is from variant [ObjectRetrieveResponseEntity].
	ExternalID string `json:"externalId"`
	// This field is from variant [ObjectRetrieveResponseRestricted].
	Restricted ObjectRetrieveResponseRestrictedRestricted `json:"restricted"`
	JSON       struct {
		ID            respjson.Field
		CreatedAt     respjson.Field
		Fields        respjson.Field
		HTTPLink      respjson.Field
		Relationships respjson.Field
		UpdatedAt     respjson.Field
		ExternalID    respjson.Field
		Restricted    respjson.Field
		raw           string
	} `json:"-"`
}

func (u ObjectRetrieveResponseUnion) AsEntity() (v ObjectRetrieveResponseEntity) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ObjectRetrieveResponseUnion) AsRestricted() (v ObjectRetrieveResponseRestricted) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ObjectRetrieveResponseUnion) RawJSON() string { return u.JSON.raw }

func (r *ObjectRetrieveResponseUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ObjectRetrieveResponseEntity struct {
	// Unique identifier for the entity.
	ID string `json:"id" api:"required"`
	// ISO 8601 timestamp of when the entity was created.
	CreatedAt string `json:"createdAt" api:"required"`
	// Map of field names to their typed values. System fields are prefixed with `$`
	// (e.g. `$name`, `$email`); custom attributes use their bare slug.
	Fields map[string]ObjectRetrieveResponseEntityField `json:"fields" api:"required"`
	// URL to view the entity in the Lightfield web app, or null.
	HTTPLink string `json:"httpLink" api:"required"`
	// Map of relationship names to their associated entities. System relationships are
	// prefixed with `$` (e.g. `$owner`, `$contact`). A self-referential relationship
	// keeps its merged values and includes optional directional values. List filters
	// select the current record's role with `:asSource` or `:asDestination`, for
	// example `?parent:asSource[contains]=account_id`.
	Relationships map[string]ObjectRetrieveResponseEntityRelationship `json:"relationships" api:"required"`
	// ISO 8601 timestamp of when the entity was last updated, or null.
	UpdatedAt string `json:"updatedAt" api:"required"`
	// External identifier for the entity, or null if unset.
	ExternalID string `json:"externalId" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID            respjson.Field
		CreatedAt     respjson.Field
		Fields        respjson.Field
		HTTPLink      respjson.Field
		Relationships respjson.Field
		UpdatedAt     respjson.Field
		ExternalID    respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ObjectRetrieveResponseEntity) RawJSON() string { return r.JSON.raw }
func (r *ObjectRetrieveResponseEntity) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ObjectRetrieveResponseEntityField struct {
	// The field value, or null if unset.
	Value ObjectRetrieveResponseEntityFieldValueUnion `json:"value" api:"required"`
	// The data type of the field.
	//
	// Any of "ADDRESS", "CHECKBOX", "CURRENCY", "DATE", "DATETIME", "EMAIL",
	// "FULL_NAME", "MARKDOWN", "MULTI_SELECT", "NUMBER", "SINGLE_SELECT",
	// "SOCIAL_HANDLE", "TELEPHONE", "TEXT", "URL", "HTML".
	ValueType string `json:"valueType" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Value       respjson.Field
		ValueType   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ObjectRetrieveResponseEntityField) RawJSON() string { return r.JSON.raw }
func (r *ObjectRetrieveResponseEntityField) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ObjectRetrieveResponseEntityFieldValueUnion contains all possible properties and
// values from [string], [float64], [bool], [[]string],
// [ObjectRetrieveResponseEntityFieldValueAddress],
// [ObjectRetrieveResponseEntityFieldValueFullName].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfStringArray]
type ObjectRetrieveResponseEntityFieldValueUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]string] instead of an object.
	OfStringArray []string `json:",inline"`
	// This field is from variant [ObjectRetrieveResponseEntityFieldValueAddress].
	City string `json:"city"`
	// This field is from variant [ObjectRetrieveResponseEntityFieldValueAddress].
	Country string `json:"country"`
	// This field is from variant [ObjectRetrieveResponseEntityFieldValueAddress].
	Latitude float64 `json:"latitude"`
	// This field is from variant [ObjectRetrieveResponseEntityFieldValueAddress].
	Longitude float64 `json:"longitude"`
	// This field is from variant [ObjectRetrieveResponseEntityFieldValueAddress].
	PostalCode string `json:"postalCode"`
	// This field is from variant [ObjectRetrieveResponseEntityFieldValueAddress].
	State string `json:"state"`
	// This field is from variant [ObjectRetrieveResponseEntityFieldValueAddress].
	Street string `json:"street"`
	// This field is from variant [ObjectRetrieveResponseEntityFieldValueAddress].
	Street2 string `json:"street2"`
	// This field is from variant [ObjectRetrieveResponseEntityFieldValueFullName].
	FirstName string `json:"firstName"`
	// This field is from variant [ObjectRetrieveResponseEntityFieldValueFullName].
	LastName string `json:"lastName"`
	JSON     struct {
		OfString      respjson.Field
		OfFloat       respjson.Field
		OfBool        respjson.Field
		OfStringArray respjson.Field
		City          respjson.Field
		Country       respjson.Field
		Latitude      respjson.Field
		Longitude     respjson.Field
		PostalCode    respjson.Field
		State         respjson.Field
		Street        respjson.Field
		Street2       respjson.Field
		FirstName     respjson.Field
		LastName      respjson.Field
		raw           string
	} `json:"-"`
}

func (u ObjectRetrieveResponseEntityFieldValueUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ObjectRetrieveResponseEntityFieldValueUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ObjectRetrieveResponseEntityFieldValueUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ObjectRetrieveResponseEntityFieldValueUnion) AsStringArray() (v []string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ObjectRetrieveResponseEntityFieldValueUnion) AsAddress() (v ObjectRetrieveResponseEntityFieldValueAddress) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ObjectRetrieveResponseEntityFieldValueUnion) AsFullName() (v ObjectRetrieveResponseEntityFieldValueFullName) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ObjectRetrieveResponseEntityFieldValueUnion) RawJSON() string { return u.JSON.raw }

func (r *ObjectRetrieveResponseEntityFieldValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ObjectRetrieveResponseEntityFieldValueAddress struct {
	// City name.
	City string `json:"city" api:"nullable"`
	// 2-letter ISO 3166-1 alpha-2 country code.
	Country string `json:"country" api:"nullable"`
	// Latitude coordinate.
	Latitude float64 `json:"latitude" api:"nullable"`
	// Longitude coordinate.
	Longitude float64 `json:"longitude" api:"nullable"`
	// Postal or ZIP code.
	PostalCode string `json:"postalCode" api:"nullable"`
	// State or province.
	State string `json:"state" api:"nullable"`
	// Street address line 1.
	Street string `json:"street" api:"nullable"`
	// Street address line 2.
	Street2 string `json:"street2" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		City        respjson.Field
		Country     respjson.Field
		Latitude    respjson.Field
		Longitude   respjson.Field
		PostalCode  respjson.Field
		State       respjson.Field
		Street      respjson.Field
		Street2     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ObjectRetrieveResponseEntityFieldValueAddress) RawJSON() string { return r.JSON.raw }
func (r *ObjectRetrieveResponseEntityFieldValueAddress) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ObjectRetrieveResponseEntityFieldValueFullName struct {
	// The contact's first name.
	FirstName string `json:"firstName" api:"nullable"`
	// The contact's last name.
	LastName string `json:"lastName" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FirstName   respjson.Field
		LastName    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ObjectRetrieveResponseEntityFieldValueFullName) RawJSON() string { return r.JSON.raw }
func (r *ObjectRetrieveResponseEntityFieldValueFullName) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ObjectRetrieveResponseEntityRelationship struct {
	// Whether the relationship is `has_one` or `has_many`.
	Cardinality string `json:"cardinality" api:"required"`
	// The type of the related object (e.g. `account`, `contact`).
	ObjectType string `json:"objectType" api:"required"`
	// IDs of the related entities.
	Values []string `json:"values" api:"required"`
	// Directional values for a self-referential relationship. List filters use the
	// relationship key with `:asSource` or `:asDestination`; writes continue to use
	// the bare key.
	Directions ObjectRetrieveResponseEntityRelationshipDirections `json:"directions"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Cardinality respjson.Field
		ObjectType  respjson.Field
		Values      respjson.Field
		Directions  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ObjectRetrieveResponseEntityRelationship) RawJSON() string { return r.JSON.raw }
func (r *ObjectRetrieveResponseEntityRelationship) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Directional values for a self-referential relationship. List filters use the
// relationship key with `:asSource` or `:asDestination`; writes continue to use
// the bare key.
type ObjectRetrieveResponseEntityRelationshipDirections struct {
	// The current record is the destination (incoming) endpoint.
	AsDestination ObjectRetrieveResponseEntityRelationshipDirectionsAsDestination `json:"asDestination" api:"required"`
	// The current record is the source (outgoing) endpoint.
	AsSource ObjectRetrieveResponseEntityRelationshipDirectionsAsSource `json:"asSource" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AsDestination respjson.Field
		AsSource      respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ObjectRetrieveResponseEntityRelationshipDirections) RawJSON() string { return r.JSON.raw }
func (r *ObjectRetrieveResponseEntityRelationshipDirections) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The current record is the destination (incoming) endpoint.
type ObjectRetrieveResponseEntityRelationshipDirectionsAsDestination struct {
	// Whether this endpoint direction is `HAS_ONE` or `HAS_MANY`.
	//
	// Any of "HAS_ONE", "HAS_MANY".
	Cardinality string `json:"cardinality" api:"required"`
	// IDs of the related entities in this direction.
	Values []string `json:"values" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Cardinality respjson.Field
		Values      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ObjectRetrieveResponseEntityRelationshipDirectionsAsDestination) RawJSON() string {
	return r.JSON.raw
}
func (r *ObjectRetrieveResponseEntityRelationshipDirectionsAsDestination) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The current record is the source (outgoing) endpoint.
type ObjectRetrieveResponseEntityRelationshipDirectionsAsSource struct {
	// Whether this endpoint direction is `HAS_ONE` or `HAS_MANY`.
	//
	// Any of "HAS_ONE", "HAS_MANY".
	Cardinality string `json:"cardinality" api:"required"`
	// IDs of the related entities in this direction.
	Values []string `json:"values" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Cardinality respjson.Field
		Values      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ObjectRetrieveResponseEntityRelationshipDirectionsAsSource) RawJSON() string {
	return r.JSON.raw
}
func (r *ObjectRetrieveResponseEntityRelationshipDirectionsAsSource) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ObjectRetrieveResponseRestricted struct {
	// Unique identifier for the record, the same ID a full response carries.
	ID string `json:"id" api:"required"`
	// Present only on a record the caller may know exists but not read. Check for it
	// before reading `fields` or `relationships`, which a restricted record does not
	// carry.
	Restricted ObjectRetrieveResponseRestrictedRestricted `json:"restricted" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Restricted  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ObjectRetrieveResponseRestricted) RawJSON() string { return r.JSON.raw }
func (r *ObjectRetrieveResponseRestricted) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Present only on a record the caller may know exists but not read. Check for it
// before reading `fields` or `relationships`, which a restricted record does not
// carry.
type ObjectRetrieveResponseRestrictedRestricted struct {
	// A label for the record's type, such as "Restricted opportunity". Never the
	// record's name.
	Name string `json:"name" api:"required"`
	// Why the record is restricted.
	//
	// Any of "NO_ACCESS", "INHERITED".
	Reason string `json:"reason" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name        respjson.Field
		Reason      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ObjectRetrieveResponseRestrictedRestricted) RawJSON() string { return r.JSON.raw }
func (r *ObjectRetrieveResponseRestrictedRestricted) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ObjectRetrieveSettingsResponse struct {
	// Stable public ID of the custom object definition.
	ID string `json:"id" api:"required"`
	// Whether this custom object type is archived.
	Archived bool `json:"archived" api:"required"`
	// ISO 8601 archive timestamp, or null when active.
	ArchivedAt string `json:"archivedAt" api:"required"`
	// ISO 8601 creation timestamp.
	CreatedAt string `json:"createdAt" api:"required"`
	// Description of the custom object type, or null.
	Description string `json:"description" api:"required"`
	// Display icon, or null when no icon is configured.
	Icon string `json:"icon" api:"required"`
	// Plural display label, or null.
	LabelPlural string `json:"labelPlural" api:"required"`
	// Singular display label.
	LabelSingular string `json:"labelSingular" api:"required"`
	// Custom object slug used in API paths.
	Slug string `json:"slug" api:"required"`
	// ISO 8601 last-update timestamp.
	UpdatedAt string `json:"updatedAt" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID            respjson.Field
		Archived      respjson.Field
		ArchivedAt    respjson.Field
		CreatedAt     respjson.Field
		Description   respjson.Field
		Icon          respjson.Field
		LabelPlural   respjson.Field
		LabelSingular respjson.Field
		Slug          respjson.Field
		UpdatedAt     respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ObjectRetrieveSettingsResponse) RawJSON() string { return r.JSON.raw }
func (r *ObjectRetrieveSettingsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ObjectUpdateDefinitionsResponse struct {
	// Map of field keys to their definitions, including both system and custom fields.
	FieldDefinitions map[string]ObjectUpdateDefinitionsResponseFieldDefinition `json:"fieldDefinitions" api:"required"`
	// The object type these definitions belong to (e.g. `account`).
	ObjectType string `json:"objectType" api:"required"`
	// Map of relationship keys to their definitions. Self-referential relationships
	// include optional directional metadata on the bare definition.
	RelationshipDefinitions map[string]ObjectUpdateDefinitionsResponseRelationshipDefinition `json:"relationshipDefinitions" api:"required"`
	// Map of relationship group keys to their definitions. A group collects
	// relationships between the same two object types, such as the roles a contact can
	// hold on an opportunity. Relationships name their group with `groupKey`.
	RelationshipGroupDefinitions map[string]ObjectUpdateDefinitionsResponseRelationshipGroupDefinition `json:"relationshipGroupDefinitions" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FieldDefinitions             respjson.Field
		ObjectType                   respjson.Field
		RelationshipDefinitions      respjson.Field
		RelationshipGroupDefinitions respjson.Field
		ExtraFields                  map[string]respjson.Field
		raw                          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ObjectUpdateDefinitionsResponse) RawJSON() string { return r.JSON.raw }
func (r *ObjectUpdateDefinitionsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ObjectUpdateDefinitionsResponseFieldDefinition struct {
	// Whether the field definition is archived.
	Archived bool `json:"archived" api:"required"`
	// When the field definition was archived, or null if active.
	ArchivedAt time.Time `json:"archivedAt" api:"required" format:"date-time"`
	// Description of the field, or null.
	Description string `json:"description" api:"required"`
	// Human-readable display name of the field.
	Label string `json:"label" api:"required"`
	// Type-specific configuration (e.g. select options, currency code). A key the
	// field type does not support is absent; a supported key with nothing set is null.
	TypeConfiguration ObjectUpdateDefinitionsResponseFieldDefinitionTypeConfiguration `json:"typeConfiguration" api:"required"`
	// Data type of the field.
	//
	// Any of "ADDRESS", "CHECKBOX", "CURRENCY", "DATE", "DATETIME", "EMAIL",
	// "FULL_NAME", "MARKDOWN", "MULTI_SELECT", "NUMBER", "SINGLE_SELECT",
	// "SOCIAL_HANDLE", "TELEPHONE", "TEXT", "URL", "HTML".
	ValueType string `json:"valueType" api:"required"`
	// Unique identifier of the field definition.
	ID string `json:"id"`
	// Whether AI fills this field automatically (`ON`), only suggests values
	// (`SUGGEST`), or leaves it alone (`OFF`). Absent on built-in fields that have no
	// `id`.
	//
	// Any of "ON", "OFF", "SUGGEST".
	AIFillable string `json:"aiFillable"`
	// Instructions AI follows when filling this field, or null when none are set.
	// Absent on built-in fields that have no `id`.
	AIFillablePrompt string `json:"aiFillablePrompt" api:"nullable"`
	// Formula a calculated field evaluates, with fields and relationships referenced
	// by key. Null for fields that are not calculated. Absent on built-in fields that
	// have no `id`.
	Formula string `json:"formula" api:"nullable"`
	// `true` for fields that are not writable via the API (e.g. AI-generated
	// summaries, calculated fields). `false` or absent for writable fields.
	ReadOnly bool `json:"readOnly"`
	// Whether values are entered (`MANUAL`) or calculated from a formula (`COMPUTED`).
	// Absent on built-in fields that have no `id`.
	//
	// Any of "MANUAL", "COMPUTED".
	ValueSource string `json:"valueSource"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Archived          respjson.Field
		ArchivedAt        respjson.Field
		Description       respjson.Field
		Label             respjson.Field
		TypeConfiguration respjson.Field
		ValueType         respjson.Field
		ID                respjson.Field
		AIFillable        respjson.Field
		AIFillablePrompt  respjson.Field
		Formula           respjson.Field
		ReadOnly          respjson.Field
		ValueSource       respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ObjectUpdateDefinitionsResponseFieldDefinition) RawJSON() string { return r.JSON.raw }
func (r *ObjectUpdateDefinitionsResponseFieldDefinition) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Type-specific configuration (e.g. select options, currency code). A key the
// field type does not support is absent; a supported key with nothing set is null.
type ObjectUpdateDefinitionsResponseFieldDefinitionTypeConfiguration struct {
	// ISO 4217 3-letter currency code. Present on `CURRENCY` fields, absent on every
	// other type.
	Currency string `json:"currency"`
	// Number of decimal places values carry, or null when none is set. Present on
	// `NUMBER` and `CURRENCY` fields, absent on every other type.
	DecimalPlaces int64 `json:"decimalPlaces" api:"nullable"`
	// Social platform associated with this handle field, or null when none is set.
	// Present on `SOCIAL_HANDLE` fields, absent on every other type.
	//
	// Any of "TWITTER", "LINKEDIN", "FACEBOOK", "INSTAGRAM".
	HandleService string `json:"handleService" api:"nullable"`
	// Whether this field accepts multiple values. Present on `EMAIL`, `TELEPHONE` and
	// `URL` fields, absent on every other type.
	MultipleValues bool `json:"multipleValues"`
	// Available options. Always present on `SINGLE_SELECT` and `MULTI_SELECT` fields,
	// as an empty list when there are none. Absent on every other type.
	Options []ObjectUpdateDefinitionsResponseFieldDefinitionTypeConfigurationOption `json:"options"`
	// For a dependent select, the field key whose selected option determines which of
	// this field's options apply. Absent for ordinary selects.
	ParentFieldKey string `json:"parentFieldKey"`
	// Whether values for this field must be unique, or null when not set. Present on
	// `EMAIL`, `TELEPHONE`, `URL` and `SOCIAL_HANDLE` fields, absent on every other
	// type.
	Unique bool `json:"unique" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Currency       respjson.Field
		DecimalPlaces  respjson.Field
		HandleService  respjson.Field
		MultipleValues respjson.Field
		Options        respjson.Field
		ParentFieldKey respjson.Field
		Unique         respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ObjectUpdateDefinitionsResponseFieldDefinitionTypeConfiguration) RawJSON() string {
	return r.JSON.raw
}
func (r *ObjectUpdateDefinitionsResponseFieldDefinitionTypeConfiguration) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ObjectUpdateDefinitionsResponseFieldDefinitionTypeConfigurationOption struct {
	// Unique identifier of the select option.
	ID string `json:"id" api:"required"`
	// Human-readable display name of the option.
	Label string `json:"label" api:"required"`
	// Lifecycle group of the option, in-progress (`ACTIVE`) or done (`COMPLETED`),
	// `ACTIVE` when none was set. Absent on selects whose options are not grouped.
	//
	// Any of "ACTIVE", "COMPLETED".
	Category string `json:"category"`
	// Color the option is displayed in, `GRAY` when none was chosen. Absent on selects
	// whose options take no color.
	//
	// Any of "RED", "ORANGE", "YELLOW", "LIME", "GREEN", "BLUE", "INDIGO", "PURPLE",
	// "COPPER", "GRAY".
	Color string `json:"color"`
	// Description of the option, or null.
	Description string `json:"description" api:"nullable"`
	// Forecast weighting of the option from 0 to 100, or null when none is set. Absent
	// on selects whose options take no forecast.
	ForecastPercentage float64 `json:"forecastPercentage" api:"nullable"`
	// Icon the option is displayed with, `BACKLOG` when none was chosen. Present only
	// on $stage.
	//
	// Any of "BACKLOG", "QUALIFICATION", "DISCOVERY", "ACTIVE", "NEGOTIATION",
	// "CLOSED_WON", "CLOSED_LOST".
	Icon string `json:"icon"`
	// For a dependent select, the ID of the option in the parent field named by
	// `typeConfiguration.parentFieldKey` that owns this option. Absent for ordinary
	// selects.
	ParentID string `json:"parentId"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                 respjson.Field
		Label              respjson.Field
		Category           respjson.Field
		Color              respjson.Field
		Description        respjson.Field
		ForecastPercentage respjson.Field
		Icon               respjson.Field
		ParentID           respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ObjectUpdateDefinitionsResponseFieldDefinitionTypeConfigurationOption) RawJSON() string {
	return r.JSON.raw
}
func (r *ObjectUpdateDefinitionsResponseFieldDefinitionTypeConfigurationOption) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ObjectUpdateDefinitionsResponseRelationshipDefinition struct {
	// Whether the relationship definition is archived.
	Archived bool `json:"archived" api:"required"`
	// When the relationship definition was archived, or null if active.
	ArchivedAt time.Time `json:"archivedAt" api:"required" format:"date-time"`
	// Whether this is a `has_one` or `has_many` relationship.
	//
	// Any of "HAS_ONE", "HAS_MANY".
	Cardinality string `json:"cardinality" api:"required"`
	// Description of the relationship, or null.
	Description string `json:"description" api:"required"`
	// Human-readable display name of the relationship, as it reads from this object
	// type.
	Label string `json:"label" api:"required"`
	// The type of the related object (e.g. `account`, `contact`).
	ObjectType string `json:"objectType" api:"required"`
	// Unique identifier of the relationship definition.
	ID string `json:"id"`
	// Directional metadata for a self-referential relationship. List filters use the
	// relationship key with `:asSource` or `:asDestination`; writes continue to use
	// the bare key.
	Directions ObjectUpdateDefinitionsResponseRelationshipDefinitionDirections `json:"directions"`
	// Key of the relationship group this relationship belongs to, as listed under
	// `relationshipGroupDefinitions`. Absent when the relationship is not in a group.
	GroupKey string `json:"groupKey"`
	// How the same relationship reads from the related object type.
	Inverse ObjectUpdateDefinitionsResponseRelationshipDefinitionInverse `json:"inverse"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Archived    respjson.Field
		ArchivedAt  respjson.Field
		Cardinality respjson.Field
		Description respjson.Field
		Label       respjson.Field
		ObjectType  respjson.Field
		ID          respjson.Field
		Directions  respjson.Field
		GroupKey    respjson.Field
		Inverse     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ObjectUpdateDefinitionsResponseRelationshipDefinition) RawJSON() string { return r.JSON.raw }
func (r *ObjectUpdateDefinitionsResponseRelationshipDefinition) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Directional metadata for a self-referential relationship. List filters use the
// relationship key with `:asSource` or `:asDestination`; writes continue to use
// the bare key.
type ObjectUpdateDefinitionsResponseRelationshipDefinitionDirections struct {
	// The current record is the destination (incoming) endpoint.
	AsDestination ObjectUpdateDefinitionsResponseRelationshipDefinitionDirectionsAsDestination `json:"asDestination" api:"required"`
	// The current record is the source (outgoing) endpoint.
	AsSource ObjectUpdateDefinitionsResponseRelationshipDefinitionDirectionsAsSource `json:"asSource" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AsDestination respjson.Field
		AsSource      respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ObjectUpdateDefinitionsResponseRelationshipDefinitionDirections) RawJSON() string {
	return r.JSON.raw
}
func (r *ObjectUpdateDefinitionsResponseRelationshipDefinitionDirections) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The current record is the destination (incoming) endpoint.
type ObjectUpdateDefinitionsResponseRelationshipDefinitionDirectionsAsDestination struct {
	// Whether this endpoint direction is `HAS_ONE` or `HAS_MANY`.
	//
	// Any of "HAS_ONE", "HAS_MANY".
	Cardinality string `json:"cardinality" api:"required"`
	// Human-readable display name for this endpoint direction.
	Label string `json:"label" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Cardinality respjson.Field
		Label       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ObjectUpdateDefinitionsResponseRelationshipDefinitionDirectionsAsDestination) RawJSON() string {
	return r.JSON.raw
}
func (r *ObjectUpdateDefinitionsResponseRelationshipDefinitionDirectionsAsDestination) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The current record is the source (outgoing) endpoint.
type ObjectUpdateDefinitionsResponseRelationshipDefinitionDirectionsAsSource struct {
	// Whether this endpoint direction is `HAS_ONE` or `HAS_MANY`.
	//
	// Any of "HAS_ONE", "HAS_MANY".
	Cardinality string `json:"cardinality" api:"required"`
	// Human-readable display name for this endpoint direction.
	Label string `json:"label" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Cardinality respjson.Field
		Label       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ObjectUpdateDefinitionsResponseRelationshipDefinitionDirectionsAsSource) RawJSON() string {
	return r.JSON.raw
}
func (r *ObjectUpdateDefinitionsResponseRelationshipDefinitionDirectionsAsSource) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// How the same relationship reads from the related object type.
type ObjectUpdateDefinitionsResponseRelationshipDefinitionInverse struct {
	// Whether this endpoint direction is `HAS_ONE` or `HAS_MANY`.
	//
	// Any of "HAS_ONE", "HAS_MANY".
	Cardinality string `json:"cardinality" api:"required"`
	// Human-readable display name for this endpoint direction.
	Label string `json:"label" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Cardinality respjson.Field
		Label       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ObjectUpdateDefinitionsResponseRelationshipDefinitionInverse) RawJSON() string {
	return r.JSON.raw
}
func (r *ObjectUpdateDefinitionsResponseRelationshipDefinitionInverse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ObjectUpdateDefinitionsResponseRelationshipGroupDefinition struct {
	// Unique identifier of the relationship group.
	ID string `json:"id" api:"required"`
	// Whether the relationship group is archived.
	Archived bool `json:"archived" api:"required"`
	// When the relationship group was archived, or null if active.
	ArchivedAt time.Time `json:"archivedAt" api:"required" format:"date-time"`
	// How the same group reads from the related object type.
	Inverse ObjectUpdateDefinitionsResponseRelationshipGroupDefinitionInverse `json:"inverse" api:"required"`
	// Human-readable display name of the group, as it reads from this object type.
	Label string `json:"label" api:"required"`
	// The type of the related object every relationship in the group points to.
	ObjectType string `json:"objectType" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Archived    respjson.Field
		ArchivedAt  respjson.Field
		Inverse     respjson.Field
		Label       respjson.Field
		ObjectType  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ObjectUpdateDefinitionsResponseRelationshipGroupDefinition) RawJSON() string {
	return r.JSON.raw
}
func (r *ObjectUpdateDefinitionsResponseRelationshipGroupDefinition) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// How the same group reads from the related object type.
type ObjectUpdateDefinitionsResponseRelationshipGroupDefinitionInverse struct {
	// Human-readable display name of the group, as it reads from the related object
	// type.
	Label string `json:"label" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Label       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ObjectUpdateDefinitionsResponseRelationshipGroupDefinitionInverse) RawJSON() string {
	return r.JSON.raw
}
func (r *ObjectUpdateDefinitionsResponseRelationshipGroupDefinitionInverse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ObjectUpdatePipelineResponseUnion contains all possible properties and values
// from [ObjectUpdatePipelineResponseObject],
// [ObjectUpdatePipelineResponseAffectedRecordCount].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ObjectUpdatePipelineResponseUnion struct {
	// This field is from variant [ObjectUpdatePipelineResponseObject].
	ID string `json:"id"`
	// This field is from variant [ObjectUpdatePipelineResponseObject].
	Description string `json:"description"`
	// This field is from variant [ObjectUpdatePipelineResponseObject].
	ForecastEnabled bool `json:"forecastEnabled"`
	// This field is from variant [ObjectUpdatePipelineResponseObject].
	Name string `json:"name"`
	// This field is from variant [ObjectUpdatePipelineResponseObject].
	Slug string `json:"slug"`
	// This field is from variant [ObjectUpdatePipelineResponseObject].
	Stages []ObjectUpdatePipelineResponseObjectStage `json:"stages"`
	// This field is from variant [ObjectUpdatePipelineResponseAffectedRecordCount].
	AffectedRecordCount int64 `json:"affectedRecordCount"`
	JSON                struct {
		ID                  respjson.Field
		Description         respjson.Field
		ForecastEnabled     respjson.Field
		Name                respjson.Field
		Slug                respjson.Field
		Stages              respjson.Field
		AffectedRecordCount respjson.Field
		raw                 string
	} `json:"-"`
}

func (u ObjectUpdatePipelineResponseUnion) AsObjectUpdatePipelineResponseObject() (v ObjectUpdatePipelineResponseObject) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ObjectUpdatePipelineResponseUnion) AsObjectUpdatePipelineResponseAffectedRecordCount() (v ObjectUpdatePipelineResponseAffectedRecordCount) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ObjectUpdatePipelineResponseUnion) RawJSON() string { return u.JSON.raw }

func (r *ObjectUpdatePipelineResponseUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ObjectUpdatePipelineResponseObject struct {
	// Unique identifier of the pipeline.
	ID string `json:"id" api:"required"`
	// Description of the pipeline, or null.
	Description string `json:"description" api:"required"`
	// Whether stages carry a `forecastPercentage`. Turning it off clears every stage
	// forecast.
	ForecastEnabled bool `json:"forecastEnabled" api:"required"`
	// Display name of the pipeline.
	Name string `json:"name" api:"required"`
	// Key derived from the name when the pipeline was created, unique on the object
	// type. Pipelines provisioned by Lightfield use a `$` prefix (e.g.
	// `$default-pipeline`).
	Slug string `json:"slug" api:"required"`
	// Stages in display order. Each `id` is the option ID records hold in `$stage`.
	Stages []ObjectUpdatePipelineResponseObjectStage `json:"stages" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID              respjson.Field
		Description     respjson.Field
		ForecastEnabled respjson.Field
		Name            respjson.Field
		Slug            respjson.Field
		Stages          respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ObjectUpdatePipelineResponseObject) RawJSON() string { return r.JSON.raw }
func (r *ObjectUpdatePipelineResponseObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ObjectUpdatePipelineResponseObjectStage struct {
	// Unique identifier of the select option.
	ID string `json:"id" api:"required"`
	// Human-readable display name of the option.
	Label string `json:"label" api:"required"`
	// Lifecycle group of the option, in-progress (`ACTIVE`) or done (`COMPLETED`),
	// `ACTIVE` when none was set. Absent on selects whose options are not grouped.
	//
	// Any of "ACTIVE", "COMPLETED".
	Category string `json:"category"`
	// Color the option is displayed in, `GRAY` when none was chosen. Absent on selects
	// whose options take no color.
	//
	// Any of "RED", "ORANGE", "YELLOW", "LIME", "GREEN", "BLUE", "INDIGO", "PURPLE",
	// "COPPER", "GRAY".
	Color string `json:"color"`
	// Description of the option, or null.
	Description string `json:"description" api:"nullable"`
	// Forecast weighting of the option from 0 to 100, or null when none is set. Absent
	// on selects whose options take no forecast.
	ForecastPercentage float64 `json:"forecastPercentage" api:"nullable"`
	// Icon the option is displayed with, `BACKLOG` when none was chosen. Present only
	// on $stage.
	//
	// Any of "BACKLOG", "QUALIFICATION", "DISCOVERY", "ACTIVE", "NEGOTIATION",
	// "CLOSED_WON", "CLOSED_LOST".
	Icon string `json:"icon"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                 respjson.Field
		Label              respjson.Field
		Category           respjson.Field
		Color              respjson.Field
		Description        respjson.Field
		ForecastPercentage respjson.Field
		Icon               respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ObjectUpdatePipelineResponseObjectStage) RawJSON() string { return r.JSON.raw }
func (r *ObjectUpdatePipelineResponseObjectStage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ObjectUpdatePipelineResponseAffectedRecordCount struct {
	// Number of records whose stage is, or would be, cleared.
	AffectedRecordCount int64 `json:"affectedRecordCount" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AffectedRecordCount respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ObjectUpdatePipelineResponseAffectedRecordCount) RawJSON() string { return r.JSON.raw }
func (r *ObjectUpdatePipelineResponseAffectedRecordCount) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ObjectUpdateResponse struct {
	// Unique identifier for the entity.
	ID string `json:"id" api:"required"`
	// ISO 8601 timestamp of when the entity was created.
	CreatedAt string `json:"createdAt" api:"required"`
	// Map of field names to their typed values. System fields are prefixed with `$`
	// (e.g. `$name`, `$email`); custom attributes use their bare slug.
	Fields map[string]ObjectUpdateResponseField `json:"fields" api:"required"`
	// URL to view the entity in the Lightfield web app, or null.
	HTTPLink string `json:"httpLink" api:"required"`
	// Map of relationship names to their associated entities. System relationships are
	// prefixed with `$` (e.g. `$owner`, `$contact`). A self-referential relationship
	// keeps its merged values and includes optional directional values. List filters
	// select the current record's role with `:asSource` or `:asDestination`, for
	// example `?parent:asSource[contains]=account_id`.
	Relationships map[string]ObjectUpdateResponseRelationship `json:"relationships" api:"required"`
	// ISO 8601 timestamp of when the entity was last updated, or null.
	UpdatedAt string `json:"updatedAt" api:"required"`
	// External identifier for the entity, or null if unset.
	ExternalID string `json:"externalId" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID            respjson.Field
		CreatedAt     respjson.Field
		Fields        respjson.Field
		HTTPLink      respjson.Field
		Relationships respjson.Field
		UpdatedAt     respjson.Field
		ExternalID    respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ObjectUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *ObjectUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ObjectUpdateResponseField struct {
	// The field value, or null if unset.
	Value ObjectUpdateResponseFieldValueUnion `json:"value" api:"required"`
	// The data type of the field.
	//
	// Any of "ADDRESS", "CHECKBOX", "CURRENCY", "DATE", "DATETIME", "EMAIL",
	// "FULL_NAME", "MARKDOWN", "MULTI_SELECT", "NUMBER", "SINGLE_SELECT",
	// "SOCIAL_HANDLE", "TELEPHONE", "TEXT", "URL", "HTML".
	ValueType string `json:"valueType" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Value       respjson.Field
		ValueType   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ObjectUpdateResponseField) RawJSON() string { return r.JSON.raw }
func (r *ObjectUpdateResponseField) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ObjectUpdateResponseFieldValueUnion contains all possible properties and values
// from [string], [float64], [bool], [[]string],
// [ObjectUpdateResponseFieldValueAddress],
// [ObjectUpdateResponseFieldValueFullName].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfStringArray]
type ObjectUpdateResponseFieldValueUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]string] instead of an object.
	OfStringArray []string `json:",inline"`
	// This field is from variant [ObjectUpdateResponseFieldValueAddress].
	City string `json:"city"`
	// This field is from variant [ObjectUpdateResponseFieldValueAddress].
	Country string `json:"country"`
	// This field is from variant [ObjectUpdateResponseFieldValueAddress].
	Latitude float64 `json:"latitude"`
	// This field is from variant [ObjectUpdateResponseFieldValueAddress].
	Longitude float64 `json:"longitude"`
	// This field is from variant [ObjectUpdateResponseFieldValueAddress].
	PostalCode string `json:"postalCode"`
	// This field is from variant [ObjectUpdateResponseFieldValueAddress].
	State string `json:"state"`
	// This field is from variant [ObjectUpdateResponseFieldValueAddress].
	Street string `json:"street"`
	// This field is from variant [ObjectUpdateResponseFieldValueAddress].
	Street2 string `json:"street2"`
	// This field is from variant [ObjectUpdateResponseFieldValueFullName].
	FirstName string `json:"firstName"`
	// This field is from variant [ObjectUpdateResponseFieldValueFullName].
	LastName string `json:"lastName"`
	JSON     struct {
		OfString      respjson.Field
		OfFloat       respjson.Field
		OfBool        respjson.Field
		OfStringArray respjson.Field
		City          respjson.Field
		Country       respjson.Field
		Latitude      respjson.Field
		Longitude     respjson.Field
		PostalCode    respjson.Field
		State         respjson.Field
		Street        respjson.Field
		Street2       respjson.Field
		FirstName     respjson.Field
		LastName      respjson.Field
		raw           string
	} `json:"-"`
}

func (u ObjectUpdateResponseFieldValueUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ObjectUpdateResponseFieldValueUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ObjectUpdateResponseFieldValueUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ObjectUpdateResponseFieldValueUnion) AsStringArray() (v []string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ObjectUpdateResponseFieldValueUnion) AsAddress() (v ObjectUpdateResponseFieldValueAddress) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ObjectUpdateResponseFieldValueUnion) AsFullName() (v ObjectUpdateResponseFieldValueFullName) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ObjectUpdateResponseFieldValueUnion) RawJSON() string { return u.JSON.raw }

func (r *ObjectUpdateResponseFieldValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ObjectUpdateResponseFieldValueAddress struct {
	// City name.
	City string `json:"city" api:"nullable"`
	// 2-letter ISO 3166-1 alpha-2 country code.
	Country string `json:"country" api:"nullable"`
	// Latitude coordinate.
	Latitude float64 `json:"latitude" api:"nullable"`
	// Longitude coordinate.
	Longitude float64 `json:"longitude" api:"nullable"`
	// Postal or ZIP code.
	PostalCode string `json:"postalCode" api:"nullable"`
	// State or province.
	State string `json:"state" api:"nullable"`
	// Street address line 1.
	Street string `json:"street" api:"nullable"`
	// Street address line 2.
	Street2 string `json:"street2" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		City        respjson.Field
		Country     respjson.Field
		Latitude    respjson.Field
		Longitude   respjson.Field
		PostalCode  respjson.Field
		State       respjson.Field
		Street      respjson.Field
		Street2     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ObjectUpdateResponseFieldValueAddress) RawJSON() string { return r.JSON.raw }
func (r *ObjectUpdateResponseFieldValueAddress) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ObjectUpdateResponseFieldValueFullName struct {
	// The contact's first name.
	FirstName string `json:"firstName" api:"nullable"`
	// The contact's last name.
	LastName string `json:"lastName" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FirstName   respjson.Field
		LastName    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ObjectUpdateResponseFieldValueFullName) RawJSON() string { return r.JSON.raw }
func (r *ObjectUpdateResponseFieldValueFullName) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ObjectUpdateResponseRelationship struct {
	// Whether the relationship is `has_one` or `has_many`.
	Cardinality string `json:"cardinality" api:"required"`
	// The type of the related object (e.g. `account`, `contact`).
	ObjectType string `json:"objectType" api:"required"`
	// IDs of the related entities.
	Values []string `json:"values" api:"required"`
	// Directional values for a self-referential relationship. List filters use the
	// relationship key with `:asSource` or `:asDestination`; writes continue to use
	// the bare key.
	Directions ObjectUpdateResponseRelationshipDirections `json:"directions"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Cardinality respjson.Field
		ObjectType  respjson.Field
		Values      respjson.Field
		Directions  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ObjectUpdateResponseRelationship) RawJSON() string { return r.JSON.raw }
func (r *ObjectUpdateResponseRelationship) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Directional values for a self-referential relationship. List filters use the
// relationship key with `:asSource` or `:asDestination`; writes continue to use
// the bare key.
type ObjectUpdateResponseRelationshipDirections struct {
	// The current record is the destination (incoming) endpoint.
	AsDestination ObjectUpdateResponseRelationshipDirectionsAsDestination `json:"asDestination" api:"required"`
	// The current record is the source (outgoing) endpoint.
	AsSource ObjectUpdateResponseRelationshipDirectionsAsSource `json:"asSource" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AsDestination respjson.Field
		AsSource      respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ObjectUpdateResponseRelationshipDirections) RawJSON() string { return r.JSON.raw }
func (r *ObjectUpdateResponseRelationshipDirections) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The current record is the destination (incoming) endpoint.
type ObjectUpdateResponseRelationshipDirectionsAsDestination struct {
	// Whether this endpoint direction is `HAS_ONE` or `HAS_MANY`.
	//
	// Any of "HAS_ONE", "HAS_MANY".
	Cardinality string `json:"cardinality" api:"required"`
	// IDs of the related entities in this direction.
	Values []string `json:"values" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Cardinality respjson.Field
		Values      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ObjectUpdateResponseRelationshipDirectionsAsDestination) RawJSON() string { return r.JSON.raw }
func (r *ObjectUpdateResponseRelationshipDirectionsAsDestination) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The current record is the source (outgoing) endpoint.
type ObjectUpdateResponseRelationshipDirectionsAsSource struct {
	// Whether this endpoint direction is `HAS_ONE` or `HAS_MANY`.
	//
	// Any of "HAS_ONE", "HAS_MANY".
	Cardinality string `json:"cardinality" api:"required"`
	// IDs of the related entities in this direction.
	Values []string `json:"values" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Cardinality respjson.Field
		Values      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ObjectUpdateResponseRelationshipDirectionsAsSource) RawJSON() string { return r.JSON.raw }
func (r *ObjectUpdateResponseRelationshipDirectionsAsSource) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ObjectUpdateSettingsResponse struct {
	// Stable public ID of the custom object definition.
	ID string `json:"id" api:"required"`
	// Whether this custom object type is archived.
	Archived bool `json:"archived" api:"required"`
	// ISO 8601 archive timestamp, or null when active.
	ArchivedAt string `json:"archivedAt" api:"required"`
	// ISO 8601 creation timestamp.
	CreatedAt string `json:"createdAt" api:"required"`
	// Description of the custom object type, or null.
	Description string `json:"description" api:"required"`
	// Display icon, or null when no icon is configured.
	Icon string `json:"icon" api:"required"`
	// Plural display label, or null.
	LabelPlural string `json:"labelPlural" api:"required"`
	// Singular display label.
	LabelSingular string `json:"labelSingular" api:"required"`
	// Custom object slug used in API paths.
	Slug string `json:"slug" api:"required"`
	// ISO 8601 last-update timestamp.
	UpdatedAt string `json:"updatedAt" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID            respjson.Field
		Archived      respjson.Field
		ArchivedAt    respjson.Field
		CreatedAt     respjson.Field
		Description   respjson.Field
		Icon          respjson.Field
		LabelPlural   respjson.Field
		LabelSingular respjson.Field
		Slug          respjson.Field
		UpdatedAt     respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ObjectUpdateSettingsResponse) RawJSON() string { return r.JSON.raw }
func (r *ObjectUpdateSettingsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ObjectNewParams struct {
	// Field names to values for the new record.
	Fields map[string]ObjectNewParamsFieldUnion `json:"fields,omitzero" api:"required"`
	// Relationship names to entity ID(s) to associate.
	Relationships map[string]ObjectNewParamsRelationshipUnion `json:"relationships,omitzero"`
	paramObj
}

func (r ObjectNewParams) MarshalJSON() (data []byte, err error) {
	type shadow ObjectNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ObjectNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ObjectNewParamsFieldUnion struct {
	OfString      param.Opt[string]             `json:",omitzero,inline"`
	OfFloat       param.Opt[float64]            `json:",omitzero,inline"`
	OfBool        param.Opt[bool]               `json:",omitzero,inline"`
	OfStringArray []string                      `json:",omitzero,inline"`
	OfAddress     *ObjectNewParamsFieldAddress  `json:",omitzero,inline"`
	OfFullName    *ObjectNewParamsFieldFullName `json:",omitzero,inline"`
	paramUnion
}

func (u ObjectNewParamsFieldUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString,
		u.OfFloat,
		u.OfBool,
		u.OfStringArray,
		u.OfAddress,
		u.OfFullName)
}
func (u *ObjectNewParamsFieldUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

type ObjectNewParamsFieldAddress struct {
	// City name.
	City param.Opt[string] `json:"city,omitzero"`
	// 2-letter ISO 3166-1 alpha-2 country code.
	Country param.Opt[string] `json:"country,omitzero"`
	// Latitude coordinate.
	Latitude param.Opt[float64] `json:"latitude,omitzero"`
	// Longitude coordinate.
	Longitude param.Opt[float64] `json:"longitude,omitzero"`
	// Postal or ZIP code.
	PostalCode param.Opt[string] `json:"postalCode,omitzero"`
	// State or province.
	State param.Opt[string] `json:"state,omitzero"`
	// Street address line 1.
	Street param.Opt[string] `json:"street,omitzero"`
	// Street address line 2.
	Street2 param.Opt[string] `json:"street2,omitzero"`
	paramObj
}

func (r ObjectNewParamsFieldAddress) MarshalJSON() (data []byte, err error) {
	type shadow ObjectNewParamsFieldAddress
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ObjectNewParamsFieldAddress) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ObjectNewParamsFieldFullName struct {
	// The contact's first name.
	FirstName param.Opt[string] `json:"firstName,omitzero"`
	// The contact's last name.
	LastName param.Opt[string] `json:"lastName,omitzero"`
	paramObj
}

func (r ObjectNewParamsFieldFullName) MarshalJSON() (data []byte, err error) {
	type shadow ObjectNewParamsFieldFullName
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ObjectNewParamsFieldFullName) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ObjectNewParamsRelationshipUnion struct {
	OfString      param.Opt[string] `json:",omitzero,inline"`
	OfStringArray []string          `json:",omitzero,inline"`
	paramUnion
}

func (u ObjectNewParamsRelationshipUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray)
}
func (u *ObjectNewParamsRelationshipUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

type ObjectGetParams struct {
	// The slug of the custom object type.
	EntitySlug string `path:"entitySlug" api:"required" json:"-"`
	paramObj
}

type ObjectUpdateParams struct {
	// The slug of the custom object type.
	EntitySlug string `path:"entitySlug" api:"required" json:"-"`
	// Field names to values. Only provided fields are modified.
	Fields map[string]ObjectUpdateParamsFieldUnion `json:"fields,omitzero"`
	// Relationship names to operations (`add`, `remove`, or `replace`).
	Relationships map[string]ObjectUpdateParamsRelationship `json:"relationships,omitzero"`
	paramObj
}

func (r ObjectUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow ObjectUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ObjectUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ObjectUpdateParamsFieldUnion struct {
	OfString      param.Opt[string]                `json:",omitzero,inline"`
	OfFloat       param.Opt[float64]               `json:",omitzero,inline"`
	OfBool        param.Opt[bool]                  `json:",omitzero,inline"`
	OfStringArray []string                         `json:",omitzero,inline"`
	OfAddress     *ObjectUpdateParamsFieldAddress  `json:",omitzero,inline"`
	OfFullName    *ObjectUpdateParamsFieldFullName `json:",omitzero,inline"`
	paramUnion
}

func (u ObjectUpdateParamsFieldUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString,
		u.OfFloat,
		u.OfBool,
		u.OfStringArray,
		u.OfAddress,
		u.OfFullName)
}
func (u *ObjectUpdateParamsFieldUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

type ObjectUpdateParamsFieldAddress struct {
	// City name.
	City param.Opt[string] `json:"city,omitzero"`
	// 2-letter ISO 3166-1 alpha-2 country code.
	Country param.Opt[string] `json:"country,omitzero"`
	// Latitude coordinate.
	Latitude param.Opt[float64] `json:"latitude,omitzero"`
	// Longitude coordinate.
	Longitude param.Opt[float64] `json:"longitude,omitzero"`
	// Postal or ZIP code.
	PostalCode param.Opt[string] `json:"postalCode,omitzero"`
	// State or province.
	State param.Opt[string] `json:"state,omitzero"`
	// Street address line 1.
	Street param.Opt[string] `json:"street,omitzero"`
	// Street address line 2.
	Street2 param.Opt[string] `json:"street2,omitzero"`
	paramObj
}

func (r ObjectUpdateParamsFieldAddress) MarshalJSON() (data []byte, err error) {
	type shadow ObjectUpdateParamsFieldAddress
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ObjectUpdateParamsFieldAddress) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ObjectUpdateParamsFieldFullName struct {
	// The contact's first name.
	FirstName param.Opt[string] `json:"firstName,omitzero"`
	// The contact's last name.
	LastName param.Opt[string] `json:"lastName,omitzero"`
	paramObj
}

func (r ObjectUpdateParamsFieldFullName) MarshalJSON() (data []byte, err error) {
	type shadow ObjectUpdateParamsFieldFullName
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ObjectUpdateParamsFieldFullName) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An operation to modify a relationship. Provide one of `add`, `remove`, or
// `replace`.
type ObjectUpdateParamsRelationship struct {
	// The entity ID to set as the relationship replacement, or null to clear supported
	// relationships.
	Replace param.Opt[string] `json:"replace,omitzero"`
	// Entity ID(s) to add to the relationship.
	Add ObjectUpdateParamsRelationshipAddUnion `json:"add,omitzero"`
	// Entity ID(s) to remove from the relationship.
	Remove ObjectUpdateParamsRelationshipRemoveUnion `json:"remove,omitzero"`
	paramObj
}

func (r ObjectUpdateParamsRelationship) MarshalJSON() (data []byte, err error) {
	type shadow ObjectUpdateParamsRelationship
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ObjectUpdateParamsRelationship) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ObjectUpdateParamsRelationshipAddUnion struct {
	OfString      param.Opt[string] `json:",omitzero,inline"`
	OfStringArray []string          `json:",omitzero,inline"`
	paramUnion
}

func (u ObjectUpdateParamsRelationshipAddUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray)
}
func (u *ObjectUpdateParamsRelationshipAddUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ObjectUpdateParamsRelationshipRemoveUnion struct {
	OfString      param.Opt[string] `json:",omitzero,inline"`
	OfStringArray []string          `json:",omitzero,inline"`
	paramUnion
}

func (u ObjectUpdateParamsRelationshipRemoveUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray)
}
func (u *ObjectUpdateParamsRelationshipRemoveUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

type ObjectListParams struct {
	// Comma-separated field and relationship keys to return, keyed as in the record
	// (`$name,$stage,region`). Each record then holds only these in `fields` and
	// `relationships`; `id`, timestamps and `httpLink` are always returned. Omit to
	// return every field. Selecting fewer fields makes large pulls much smaller.
	Fields param.Opt[string] `query:"fields,omitzero" json:"-"`
	// A JSON filter tree, for `or` logic or nested groups: one condition
	// `{"field": "$stage", "operator": "equal", "value": "Proposal"}`, or an `and` /
	// `or` group whose items are conditions or one more level of groups, e.g.
	// `{"or": [{"field": "$stage", "value": "Proposal"}, {"field": "$stage", "value": "Negotiation"}]}`.
	// `operator` defaults to `equal`, or `contains` for a field that holds several
	// values; prefix it with `-` to negate. At most 25 conditions. URL-encode the
	// JSON. It is ANDed with any `field[operator]=value` parameters.
	Filter param.Opt[string] `query:"filter,omitzero" json:"-"`
	// Maximum number of records to return. Defaults to 25, maximum 25.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Number of records to skip for pagination. Defaults to 0.
	Offset param.Opt[int64] `query:"offset,omitzero" json:"-"`
	// Result order: `$createdAt` or `$updatedAt`, oldest first, or prefixed with `-`
	// for newest first. Defaults to `-$createdAt`. Use `$updatedAt` to process changes
	// oldest first.
	//
	// Any of "$createdAt", "-$createdAt", "$updatedAt", "-$updatedAt".
	Sort ObjectListParamsSort `query:"sort,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ObjectListParams]'s query parameters as `url.Values`.
func (r ObjectListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Result order: `$createdAt` or `$updatedAt`, oldest first, or prefixed with `-`
// for newest first. Defaults to `-$createdAt`. Use `$updatedAt` to process changes
// oldest first.
type ObjectListParamsSort string

const (
	ObjectListParamsSortCreatedAt      ObjectListParamsSort = "$createdAt"
	ObjectListParamsSortMinusCreatedAt ObjectListParamsSort = "-$createdAt"
	ObjectListParamsSortUpdatedAt      ObjectListParamsSort = "$updatedAt"
	ObjectListParamsSortMinusUpdatedAt ObjectListParamsSort = "-$updatedAt"
)

type ObjectDeleteParams struct {
	// The slug of the custom object type.
	EntitySlug string `path:"entitySlug" api:"required" json:"-"`
	Body       ObjectDeleteParamsBody
	paramObj
}

func (r ObjectDeleteParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.Body)
}
func (r *ObjectDeleteParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ObjectDeleteParamsBody struct {
	paramObj
}

func (r ObjectDeleteParamsBody) MarshalJSON() (data []byte, err error) {
	type shadow ObjectDeleteParamsBody
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ObjectDeleteParamsBody) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ObjectArchiveTypeParams struct {
	Body ObjectArchiveTypeParamsBody
	paramObj
}

func (r ObjectArchiveTypeParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.Body)
}
func (r *ObjectArchiveTypeParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ObjectArchiveTypeParamsBody struct {
	paramObj
}

func (r ObjectArchiveTypeParamsBody) MarshalJSON() (data []byte, err error) {
	type shadow ObjectArchiveTypeParamsBody
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ObjectArchiveTypeParamsBody) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ObjectNewPipelineParams struct {
	// Display name of the pipeline, unique among pipelines on the object type.
	Name string `json:"name" api:"required"`
	// Description of the pipeline. Send null to clear it; an empty or whitespace-only
	// string returns 422.
	Description param.Opt[string] `json:"description,omitzero"`
	// Whether stages carry a forecast. Defaults to false.
	ForecastEnabled param.Opt[bool] `json:"forecastEnabled,omitzero"`
	// Key of the pipeline, fixed once created. A slug another pipeline on the object
	// type already has is refused with 409 `crm.pipeline.slug_taken`, so a create sent
	// again after a lost response is recognized. Derived from `name` when left out.
	Slug param.Opt[string] `json:"slug,omitzero"`
	// Stages in display order. Defaults to none.
	Stages []ObjectNewPipelineParamsStage `json:"stages,omitzero"`
	paramObj
}

func (r ObjectNewPipelineParams) MarshalJSON() (data []byte, err error) {
	type shadow ObjectNewPipelineParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ObjectNewPipelineParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Category, Color, Icon, Label are required.
type ObjectNewPipelineParamsStage struct {
	// Lifecycle group of the stage, in-progress (`ACTIVE`) or done (`COMPLETED`).
	//
	// Any of "ACTIVE", "COMPLETED".
	Category string `json:"category,omitzero" api:"required"`
	// Color the stage is displayed in.
	//
	// Any of "RED", "ORANGE", "YELLOW", "LIME", "GREEN", "BLUE", "INDIGO", "PURPLE",
	// "COPPER", "GRAY".
	Color string `json:"color,omitzero" api:"required"`
	// Icon the stage is displayed with.
	//
	// Any of "BACKLOG", "QUALIFICATION", "DISCOVERY", "ACTIVE", "NEGOTIATION",
	// "CLOSED_WON", "CLOSED_LOST".
	Icon string `json:"icon,omitzero" api:"required"`
	// Display name of the stage.
	Label string `json:"label" api:"required"`
	// Description of the stage. Null clears it.
	Description param.Opt[string] `json:"description,omitzero"`
	// Forecast weighting from 0 to 100. Only allowed when `forecastEnabled`. Null
	// clears it.
	ForecastPercentage param.Opt[float64] `json:"forecastPercentage,omitzero"`
	// Omit the ID to create a new stage with a random UUIDv4. To keep or edit a stage,
	// send its published ID. A caller-chosen ID starts with `opt_` followed by either
	// a UUID in any letter case or groups of lowercase letters and digits separated by
	// single hyphens or underscores. The ID must be unique on the object type,
	// ignoring case, hyphens, and underscores.
	ID param.Opt[string] `json:"id,omitzero"`
	paramObj
}

func (r ObjectNewPipelineParamsStage) MarshalJSON() (data []byte, err error) {
	type shadow ObjectNewPipelineParamsStage
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ObjectNewPipelineParamsStage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ObjectNewPipelineParamsStage](
		"category", "ACTIVE", "COMPLETED",
	)
	apijson.RegisterFieldValidator[ObjectNewPipelineParamsStage](
		"color", "RED", "ORANGE", "YELLOW", "LIME", "GREEN", "BLUE", "INDIGO", "PURPLE", "COPPER", "GRAY",
	)
	apijson.RegisterFieldValidator[ObjectNewPipelineParamsStage](
		"icon", "BACKLOG", "QUALIFICATION", "DISCOVERY", "ACTIVE", "NEGOTIATION", "CLOSED_WON", "CLOSED_LOST",
	)
}

type ObjectNewTypeParams struct {
	// Singular display label.
	LabelSingular string `json:"labelSingular" api:"required"`
	// Unique custom object slug within the workspace. Immutable after creation.
	Slug string `json:"slug" api:"required"`
	// Description of the custom object type, or null.
	Description param.Opt[string] `json:"description,omitzero"`
	// Plural label. Defaults to the pluralized singular label when omitted or null.
	LabelPlural param.Opt[string] `json:"labelPlural,omitzero"`
	// Display icon, or null to use the default.
	//
	// Any of "Add", "Agent", "AlarmClock", "Archive", "Attachment", "Bell", "Book",
	// "Briefcase", "Calendar", "CameraOutline", "ChartArea", "ChartLine", "ChartPie",
	// "Checkmark", "ChevronRightDouble", "CircleAdd", "Clock", "Code", "Command",
	// "Competitors", "Compose", "CreditCard", "DealValue", "DeviceDesktop", "Dice",
	// "Document", "Download", "Draggable", "EmployeeBadge", "EngagementAnalyticsAlt",
	// "Enrich", "Folder", "Funding", "Grid", "Hashtag", "Help", "Key", "Keyboard",
	// "LightBulb", "Link", "Mail", "MicrophoneOutline", "Money", "Moon",
	// "MultiplePagesPortrait", "OpenQuote", "PageLandscape", "Paperplane", "People",
	// "Phone", "Preferences", "Search", "Shapes", "Sparkle", "SquareStack", "Star",
	// "Suggestions", "Sun", "Table", "TableColumn", "Tag", "Teams", "ThreeDCube",
	// "ThreeDStack", "VideoAlt", "View", "Waveform".
	Icon ObjectNewTypeParamsIcon `json:"icon,omitzero"`
	paramObj
}

func (r ObjectNewTypeParams) MarshalJSON() (data []byte, err error) {
	type shadow ObjectNewTypeParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ObjectNewTypeParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Display icon, or null to use the default.
type ObjectNewTypeParamsIcon string

const (
	ObjectNewTypeParamsIconAdd                    ObjectNewTypeParamsIcon = "Add"
	ObjectNewTypeParamsIconAgent                  ObjectNewTypeParamsIcon = "Agent"
	ObjectNewTypeParamsIconAlarmClock             ObjectNewTypeParamsIcon = "AlarmClock"
	ObjectNewTypeParamsIconArchive                ObjectNewTypeParamsIcon = "Archive"
	ObjectNewTypeParamsIconAttachment             ObjectNewTypeParamsIcon = "Attachment"
	ObjectNewTypeParamsIconBell                   ObjectNewTypeParamsIcon = "Bell"
	ObjectNewTypeParamsIconBook                   ObjectNewTypeParamsIcon = "Book"
	ObjectNewTypeParamsIconBriefcase              ObjectNewTypeParamsIcon = "Briefcase"
	ObjectNewTypeParamsIconCalendar               ObjectNewTypeParamsIcon = "Calendar"
	ObjectNewTypeParamsIconCameraOutline          ObjectNewTypeParamsIcon = "CameraOutline"
	ObjectNewTypeParamsIconChartArea              ObjectNewTypeParamsIcon = "ChartArea"
	ObjectNewTypeParamsIconChartLine              ObjectNewTypeParamsIcon = "ChartLine"
	ObjectNewTypeParamsIconChartPie               ObjectNewTypeParamsIcon = "ChartPie"
	ObjectNewTypeParamsIconCheckmark              ObjectNewTypeParamsIcon = "Checkmark"
	ObjectNewTypeParamsIconChevronRightDouble     ObjectNewTypeParamsIcon = "ChevronRightDouble"
	ObjectNewTypeParamsIconCircleAdd              ObjectNewTypeParamsIcon = "CircleAdd"
	ObjectNewTypeParamsIconClock                  ObjectNewTypeParamsIcon = "Clock"
	ObjectNewTypeParamsIconCode                   ObjectNewTypeParamsIcon = "Code"
	ObjectNewTypeParamsIconCommand                ObjectNewTypeParamsIcon = "Command"
	ObjectNewTypeParamsIconCompetitors            ObjectNewTypeParamsIcon = "Competitors"
	ObjectNewTypeParamsIconCompose                ObjectNewTypeParamsIcon = "Compose"
	ObjectNewTypeParamsIconCreditCard             ObjectNewTypeParamsIcon = "CreditCard"
	ObjectNewTypeParamsIconDealValue              ObjectNewTypeParamsIcon = "DealValue"
	ObjectNewTypeParamsIconDeviceDesktop          ObjectNewTypeParamsIcon = "DeviceDesktop"
	ObjectNewTypeParamsIconDice                   ObjectNewTypeParamsIcon = "Dice"
	ObjectNewTypeParamsIconDocument               ObjectNewTypeParamsIcon = "Document"
	ObjectNewTypeParamsIconDownload               ObjectNewTypeParamsIcon = "Download"
	ObjectNewTypeParamsIconDraggable              ObjectNewTypeParamsIcon = "Draggable"
	ObjectNewTypeParamsIconEmployeeBadge          ObjectNewTypeParamsIcon = "EmployeeBadge"
	ObjectNewTypeParamsIconEngagementAnalyticsAlt ObjectNewTypeParamsIcon = "EngagementAnalyticsAlt"
	ObjectNewTypeParamsIconEnrich                 ObjectNewTypeParamsIcon = "Enrich"
	ObjectNewTypeParamsIconFolder                 ObjectNewTypeParamsIcon = "Folder"
	ObjectNewTypeParamsIconFunding                ObjectNewTypeParamsIcon = "Funding"
	ObjectNewTypeParamsIconGrid                   ObjectNewTypeParamsIcon = "Grid"
	ObjectNewTypeParamsIconHashtag                ObjectNewTypeParamsIcon = "Hashtag"
	ObjectNewTypeParamsIconHelp                   ObjectNewTypeParamsIcon = "Help"
	ObjectNewTypeParamsIconKey                    ObjectNewTypeParamsIcon = "Key"
	ObjectNewTypeParamsIconKeyboard               ObjectNewTypeParamsIcon = "Keyboard"
	ObjectNewTypeParamsIconLightBulb              ObjectNewTypeParamsIcon = "LightBulb"
	ObjectNewTypeParamsIconLink                   ObjectNewTypeParamsIcon = "Link"
	ObjectNewTypeParamsIconMail                   ObjectNewTypeParamsIcon = "Mail"
	ObjectNewTypeParamsIconMicrophoneOutline      ObjectNewTypeParamsIcon = "MicrophoneOutline"
	ObjectNewTypeParamsIconMoney                  ObjectNewTypeParamsIcon = "Money"
	ObjectNewTypeParamsIconMoon                   ObjectNewTypeParamsIcon = "Moon"
	ObjectNewTypeParamsIconMultiplePagesPortrait  ObjectNewTypeParamsIcon = "MultiplePagesPortrait"
	ObjectNewTypeParamsIconOpenQuote              ObjectNewTypeParamsIcon = "OpenQuote"
	ObjectNewTypeParamsIconPageLandscape          ObjectNewTypeParamsIcon = "PageLandscape"
	ObjectNewTypeParamsIconPaperplane             ObjectNewTypeParamsIcon = "Paperplane"
	ObjectNewTypeParamsIconPeople                 ObjectNewTypeParamsIcon = "People"
	ObjectNewTypeParamsIconPhone                  ObjectNewTypeParamsIcon = "Phone"
	ObjectNewTypeParamsIconPreferences            ObjectNewTypeParamsIcon = "Preferences"
	ObjectNewTypeParamsIconSearch                 ObjectNewTypeParamsIcon = "Search"
	ObjectNewTypeParamsIconShapes                 ObjectNewTypeParamsIcon = "Shapes"
	ObjectNewTypeParamsIconSparkle                ObjectNewTypeParamsIcon = "Sparkle"
	ObjectNewTypeParamsIconSquareStack            ObjectNewTypeParamsIcon = "SquareStack"
	ObjectNewTypeParamsIconStar                   ObjectNewTypeParamsIcon = "Star"
	ObjectNewTypeParamsIconSuggestions            ObjectNewTypeParamsIcon = "Suggestions"
	ObjectNewTypeParamsIconSun                    ObjectNewTypeParamsIcon = "Sun"
	ObjectNewTypeParamsIconTable                  ObjectNewTypeParamsIcon = "Table"
	ObjectNewTypeParamsIconTableColumn            ObjectNewTypeParamsIcon = "TableColumn"
	ObjectNewTypeParamsIconTag                    ObjectNewTypeParamsIcon = "Tag"
	ObjectNewTypeParamsIconTeams                  ObjectNewTypeParamsIcon = "Teams"
	ObjectNewTypeParamsIconThreeDCube             ObjectNewTypeParamsIcon = "ThreeDCube"
	ObjectNewTypeParamsIconThreeDStack            ObjectNewTypeParamsIcon = "ThreeDStack"
	ObjectNewTypeParamsIconVideoAlt               ObjectNewTypeParamsIcon = "VideoAlt"
	ObjectNewTypeParamsIconView                   ObjectNewTypeParamsIcon = "View"
	ObjectNewTypeParamsIconWaveform               ObjectNewTypeParamsIcon = "Waveform"
)

type ObjectDefinitionsParams struct {
	// Include archived field and relationship definitions, and those of an archived
	// custom object type or whose other end is one. Defaults to false.
	IncludeArchived param.Opt[bool] `query:"includeArchived,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ObjectDefinitionsParams]'s query parameters as
// `url.Values`.
func (r ObjectDefinitionsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ObjectDeletePipelineParams struct {
	// The slug of the custom object type.
	EntitySlug string `path:"entitySlug" api:"required" json:"-"`
	// Check the request and count the records it would clear, without changing
	// anything. Defaults to false.
	DryRun param.Opt[bool] `query:"dryRun,omitzero" json:"-"`
	Body   ObjectDeletePipelineParamsBody
	paramObj
}

func (r ObjectDeletePipelineParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.Body)
}
func (r *ObjectDeletePipelineParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [ObjectDeletePipelineParams]'s query parameters as
// `url.Values`.
func (r ObjectDeletePipelineParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ObjectDeletePipelineParamsBody struct {
	paramObj
}

func (r ObjectDeletePipelineParamsBody) MarshalJSON() (data []byte, err error) {
	type shadow ObjectDeletePipelineParamsBody
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ObjectDeletePipelineParamsBody) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ObjectFieldHistoryParams struct {
	// The slug of the custom object type.
	EntitySlug string `path:"entitySlug" api:"required" json:"-"`
	// Unique identifier of the record.
	ID string `path:"id" api:"required" json:"-"`
	// Cursor from a previous response’s `nextCursor` to fetch the next page.
	After param.Opt[string] `query:"after,omitzero" json:"-"`
	// Maximum number of history entries to return. Defaults to 20, maximum 100.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ObjectFieldHistoryParams]'s query parameters as
// `url.Values`.
func (r ObjectFieldHistoryParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ObjectRestoreParams struct {
	// The slug of the custom object type.
	EntitySlug string `path:"entitySlug" api:"required" json:"-"`
	Body       ObjectRestoreParamsBody
	paramObj
}

func (r ObjectRestoreParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.Body)
}
func (r *ObjectRestoreParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ObjectRestoreParamsBody struct {
	paramObj
}

func (r ObjectRestoreParamsBody) MarshalJSON() (data []byte, err error) {
	type shadow ObjectRestoreParamsBody
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ObjectRestoreParamsBody) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ObjectGetPipelineParams struct {
	// The slug of the custom object type.
	EntitySlug string `path:"entitySlug" api:"required" json:"-"`
	paramObj
}

type ObjectGetSettingsParams struct {
	// Include archived definitions. Defaults to false.
	IncludeArchived param.Opt[bool] `query:"includeArchived,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ObjectGetSettingsParams]'s query parameters as
// `url.Values`.
func (r ObjectGetSettingsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ObjectUpdateDefinitionsParams struct {
	// Field definitions to create or change, keyed by field key. A key that names no
	// field creates one with that key, and needs `label` and `valueType`. `filter`,
	// `fields` and `sort` can't be new keys, since list endpoints use them as query
	// params. A key that names a field changes only the properties sent. `archived`
	// archives or restores it. At most 10 definitions per request.
	FieldDefinitions map[string]ObjectUpdateDefinitionsParamsFieldDefinition `json:"fieldDefinitions,omitzero"`
	// Relationship definitions to create or change, keyed by relationship key.
	// Creating one needs `label`, `cardinality`, `objectType` and
	// `inverse.cardinality`.
	RelationshipDefinitions map[string]ObjectUpdateDefinitionsParamsRelationshipDefinition `json:"relationshipDefinitions,omitzero"`
	// Relationship group definitions to create or change, keyed by group key. Creating
	// one needs `label` and `objectType`. Groups are written before relationships, so
	// a relationship in the same request can name a new group.
	RelationshipGroupDefinitions map[string]ObjectUpdateDefinitionsParamsRelationshipGroupDefinition `json:"relationshipGroupDefinitions,omitzero"`
	paramObj
}

func (r ObjectUpdateDefinitionsParams) MarshalJSON() (data []byte, err error) {
	type shadow ObjectUpdateDefinitionsParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ObjectUpdateDefinitionsParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ObjectUpdateDefinitionsParamsFieldDefinition struct {
	// Instructions AI follows when filling this field, or null when none are set.
	// Absent on built-in fields that have no `id`.
	AIFillablePrompt param.Opt[string] `json:"aiFillablePrompt,omitzero"`
	// Description of the field, or null.
	Description param.Opt[string] `json:"description,omitzero"`
	// Formula a calculated field evaluates, with fields and relationships referenced
	// by key. Null for fields that are not calculated. Absent on built-in fields that
	// have no `id`.
	Formula param.Opt[string] `json:"formula,omitzero"`
	// Whether the field definition is archived.
	Archived param.Opt[bool] `json:"archived,omitzero"`
	// Human-readable display name of the field.
	Label param.Opt[string] `json:"label,omitzero"`
	// Whether AI fills this field automatically (`ON`), only suggests values
	// (`SUGGEST`), or leaves it alone (`OFF`). Absent on built-in fields that have no
	// `id`.
	//
	// Any of "ON", "OFF", "SUGGEST".
	AIFillable        string                                                        `json:"aiFillable,omitzero"`
	TypeConfiguration ObjectUpdateDefinitionsParamsFieldDefinitionTypeConfiguration `json:"typeConfiguration,omitzero"`
	// Data type of the field.
	//
	// Any of "ADDRESS", "CHECKBOX", "CURRENCY", "DATE", "DATETIME", "EMAIL",
	// "FULL_NAME", "MARKDOWN", "MULTI_SELECT", "NUMBER", "SINGLE_SELECT",
	// "SOCIAL_HANDLE", "TELEPHONE", "TEXT", "URL", "HTML".
	ValueType string `json:"valueType,omitzero"`
	paramObj
}

func (r ObjectUpdateDefinitionsParamsFieldDefinition) MarshalJSON() (data []byte, err error) {
	type shadow ObjectUpdateDefinitionsParamsFieldDefinition
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ObjectUpdateDefinitionsParamsFieldDefinition) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ObjectUpdateDefinitionsParamsFieldDefinition](
		"aiFillable", "ON", "OFF", "SUGGEST",
	)
	apijson.RegisterFieldValidator[ObjectUpdateDefinitionsParamsFieldDefinition](
		"valueType", "ADDRESS", "CHECKBOX", "CURRENCY", "DATE", "DATETIME", "EMAIL", "FULL_NAME", "MARKDOWN", "MULTI_SELECT", "NUMBER", "SINGLE_SELECT", "SOCIAL_HANDLE", "TELEPHONE", "TEXT", "URL", "HTML",
	)
}

type ObjectUpdateDefinitionsParamsFieldDefinitionTypeConfiguration struct {
	// Number of decimal places values carry, or null when none is set. Present on
	// `NUMBER` and `CURRENCY` fields, absent on every other type.
	DecimalPlaces param.Opt[int64] `json:"decimalPlaces,omitzero"`
	// ISO 4217 3-letter currency code. Present on `CURRENCY` fields, absent on every
	// other type.
	Currency param.Opt[string] `json:"currency,omitzero"`
	// Whether this field accepts multiple values. Present on `EMAIL`, `TELEPHONE` and
	// `URL` fields, absent on every other type.
	MultipleValues param.Opt[bool] `json:"multipleValues,omitzero"`
	// For a dependent select, the field key whose selected option determines which of
	// this field's options apply. Absent for ordinary selects.
	ParentFieldKey param.Opt[string] `json:"parentFieldKey,omitzero"`
	// Social platform associated with this handle field, or null when none is set.
	// Present on `SOCIAL_HANDLE` fields, absent on every other type.
	//
	// Any of "TWITTER", "LINKEDIN", "FACEBOOK", "INSTAGRAM".
	HandleService string `json:"handleService,omitzero"`
	// The full option list of a `SINGLE_SELECT` or `MULTI_SELECT` field, matched to
	// the options the field has by `id`. An `id` the field does not have creates an
	// option with that `id`, which must be `opt_` followed by lowercase letters and
	// digits separated by single hyphens or underscores. A new option cannot take the
	// label of an existing one unless the same request relabels that one. A
	// `description` sent as null is cleared; a property left out is kept. Options left
	// out are removed. `SINGLE_SELECT` options are displayed in the order sent.
	Options []ObjectUpdateDefinitionsParamsFieldDefinitionTypeConfigurationOption `json:"options,omitzero"`
	paramObj
}

func (r ObjectUpdateDefinitionsParamsFieldDefinitionTypeConfiguration) MarshalJSON() (data []byte, err error) {
	type shadow ObjectUpdateDefinitionsParamsFieldDefinitionTypeConfiguration
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ObjectUpdateDefinitionsParamsFieldDefinitionTypeConfiguration) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ObjectUpdateDefinitionsParamsFieldDefinitionTypeConfiguration](
		"handleService", "TWITTER", "LINKEDIN", "FACEBOOK", "INSTAGRAM",
	)
}

// The properties ID, Label are required.
type ObjectUpdateDefinitionsParamsFieldDefinitionTypeConfigurationOption struct {
	// Unique identifier of the select option.
	ID string `json:"id" api:"required"`
	// Human-readable display name of the option.
	Label string `json:"label" api:"required"`
	// Description of the option, or null.
	Description param.Opt[string] `json:"description,omitzero"`
	// Color the option is displayed in, `GRAY` when none was chosen. Absent on selects
	// whose options take no color.
	//
	// Any of "RED", "ORANGE", "YELLOW", "LIME", "GREEN", "BLUE", "INDIGO", "PURPLE",
	// "COPPER", "GRAY".
	Color string `json:"color,omitzero"`
	paramObj
}

func (r ObjectUpdateDefinitionsParamsFieldDefinitionTypeConfigurationOption) MarshalJSON() (data []byte, err error) {
	type shadow ObjectUpdateDefinitionsParamsFieldDefinitionTypeConfigurationOption
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ObjectUpdateDefinitionsParamsFieldDefinitionTypeConfigurationOption) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ObjectUpdateDefinitionsParamsFieldDefinitionTypeConfigurationOption](
		"color", "RED", "ORANGE", "YELLOW", "LIME", "GREEN", "BLUE", "INDIGO", "PURPLE", "COPPER", "GRAY",
	)
}

type ObjectUpdateDefinitionsParamsRelationshipDefinition struct {
	// Description of the relationship, or null.
	Description param.Opt[string] `json:"description,omitzero"`
	// Key of the relationship group to create this relationship in: one that exists,
	// or one created by the same request. Fixed once the relationship exists.
	GroupKey param.Opt[string] `json:"groupKey,omitzero"`
	// Whether the relationship definition is archived.
	Archived param.Opt[bool] `json:"archived,omitzero"`
	// Human-readable display name of the relationship, as it reads from this object
	// type.
	Label param.Opt[string] `json:"label,omitzero"`
	// The type of the related object (e.g. `account`, `contact`).
	ObjectType param.Opt[string] `json:"objectType,omitzero"`
	// Whether this is a `has_one` or `has_many` relationship.
	//
	// Any of "HAS_ONE", "HAS_MANY".
	Cardinality string                                                     `json:"cardinality,omitzero"`
	Inverse     ObjectUpdateDefinitionsParamsRelationshipDefinitionInverse `json:"inverse,omitzero"`
	paramObj
}

func (r ObjectUpdateDefinitionsParamsRelationshipDefinition) MarshalJSON() (data []byte, err error) {
	type shadow ObjectUpdateDefinitionsParamsRelationshipDefinition
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ObjectUpdateDefinitionsParamsRelationshipDefinition) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ObjectUpdateDefinitionsParamsRelationshipDefinition](
		"cardinality", "HAS_ONE", "HAS_MANY",
	)
}

type ObjectUpdateDefinitionsParamsRelationshipDefinitionInverse struct {
	// Display name of the relationship as it reads from the related object type.
	// Defaults to `label`.
	Label param.Opt[string] `json:"label,omitzero"`
	// Whether the related object type reads this relationship as `HAS_ONE` or
	// `HAS_MANY`. Required to create a relationship.
	//
	// Any of "HAS_ONE", "HAS_MANY".
	Cardinality string `json:"cardinality,omitzero"`
	paramObj
}

func (r ObjectUpdateDefinitionsParamsRelationshipDefinitionInverse) MarshalJSON() (data []byte, err error) {
	type shadow ObjectUpdateDefinitionsParamsRelationshipDefinitionInverse
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ObjectUpdateDefinitionsParamsRelationshipDefinitionInverse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ObjectUpdateDefinitionsParamsRelationshipDefinitionInverse](
		"cardinality", "HAS_ONE", "HAS_MANY",
	)
}

type ObjectUpdateDefinitionsParamsRelationshipGroupDefinition struct {
	// Whether the relationship group is archived.
	Archived param.Opt[bool] `json:"archived,omitzero"`
	// Human-readable display name of the group, as it reads from this object type.
	Label param.Opt[string] `json:"label,omitzero"`
	// The type of the related object every relationship in the group points to.
	ObjectType param.Opt[string]                                               `json:"objectType,omitzero"`
	Inverse    ObjectUpdateDefinitionsParamsRelationshipGroupDefinitionInverse `json:"inverse,omitzero"`
	paramObj
}

func (r ObjectUpdateDefinitionsParamsRelationshipGroupDefinition) MarshalJSON() (data []byte, err error) {
	type shadow ObjectUpdateDefinitionsParamsRelationshipGroupDefinition
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ObjectUpdateDefinitionsParamsRelationshipGroupDefinition) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ObjectUpdateDefinitionsParamsRelationshipGroupDefinitionInverse struct {
	// Display name of the group as it reads from the related object type. Defaults to
	// `label`.
	Label param.Opt[string] `json:"label,omitzero"`
	paramObj
}

func (r ObjectUpdateDefinitionsParamsRelationshipGroupDefinitionInverse) MarshalJSON() (data []byte, err error) {
	type shadow ObjectUpdateDefinitionsParamsRelationshipGroupDefinitionInverse
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ObjectUpdateDefinitionsParamsRelationshipGroupDefinitionInverse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ObjectUpdatePipelineParams struct {
	// The slug of the custom object type.
	EntitySlug string `path:"entitySlug" api:"required" json:"-"`
	// Description of the pipeline. Send null to clear it; an empty or whitespace-only
	// string returns 422.
	Description param.Opt[string] `json:"description,omitzero"`
	// Check the request and count the records it would clear, without changing
	// anything. Defaults to false.
	DryRun param.Opt[bool] `query:"dryRun,omitzero" json:"-"`
	// Confirm deletion of existing stages omitted from the stage list.
	Force param.Opt[bool] `query:"force,omitzero" json:"-"`
	// Whether stages carry a forecast. Turning it off clears every stage forecast.
	ForecastEnabled param.Opt[bool]   `json:"forecastEnabled,omitzero"`
	Name            param.Opt[string] `json:"name,omitzero"`
	// The full stage list, in display order. Stages are matched by `id`. Omitted
	// properties on a retained stage are preserved. Omitting an existing stage
	// requires `force=true` to delete it and clear it from records.
	Stages []ObjectUpdatePipelineParamsStage `json:"stages,omitzero"`
	paramObj
}

func (r ObjectUpdatePipelineParams) MarshalJSON() (data []byte, err error) {
	type shadow ObjectUpdatePipelineParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ObjectUpdatePipelineParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [ObjectUpdatePipelineParams]'s query parameters as
// `url.Values`.
func (r ObjectUpdatePipelineParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ObjectUpdatePipelineParamsStage struct {
	// Description of the stage. Null clears it.
	Description param.Opt[string] `json:"description,omitzero"`
	// Forecast weighting from 0 to 100. Only allowed when `forecastEnabled`. Null
	// clears it.
	ForecastPercentage param.Opt[float64] `json:"forecastPercentage,omitzero"`
	// Omit the ID to create a new stage with a random UUIDv4. To keep or edit a stage,
	// send its published ID. A caller-chosen ID starts with `opt_` followed by either
	// a UUID in any letter case or groups of lowercase letters and digits separated by
	// single hyphens or underscores. The ID must be unique on the object type,
	// ignoring case, hyphens, and underscores.
	ID param.Opt[string] `json:"id,omitzero"`
	// Display name of the stage.
	Label param.Opt[string] `json:"label,omitzero"`
	// Lifecycle group of the stage, in-progress (`ACTIVE`) or done (`COMPLETED`).
	//
	// Any of "ACTIVE", "COMPLETED".
	Category string `json:"category,omitzero"`
	// Color the stage is displayed in.
	//
	// Any of "RED", "ORANGE", "YELLOW", "LIME", "GREEN", "BLUE", "INDIGO", "PURPLE",
	// "COPPER", "GRAY".
	Color string `json:"color,omitzero"`
	// Icon the stage is displayed with.
	//
	// Any of "BACKLOG", "QUALIFICATION", "DISCOVERY", "ACTIVE", "NEGOTIATION",
	// "CLOSED_WON", "CLOSED_LOST".
	Icon string `json:"icon,omitzero"`
	paramObj
}

func (r ObjectUpdatePipelineParamsStage) MarshalJSON() (data []byte, err error) {
	type shadow ObjectUpdatePipelineParamsStage
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ObjectUpdatePipelineParamsStage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ObjectUpdatePipelineParamsStage](
		"category", "ACTIVE", "COMPLETED",
	)
	apijson.RegisterFieldValidator[ObjectUpdatePipelineParamsStage](
		"color", "RED", "ORANGE", "YELLOW", "LIME", "GREEN", "BLUE", "INDIGO", "PURPLE", "COPPER", "GRAY",
	)
	apijson.RegisterFieldValidator[ObjectUpdatePipelineParamsStage](
		"icon", "BACKLOG", "QUALIFICATION", "DISCOVERY", "ACTIVE", "NEGOTIATION", "CLOSED_WON", "CLOSED_LOST",
	)
}

type ObjectUpdateSettingsParams struct {
	// Description of the custom object type, or null.
	Description param.Opt[string] `json:"description,omitzero"`
	// Plural display label, or null.
	LabelPlural param.Opt[string] `json:"labelPlural,omitzero"`
	// Singular display label.
	LabelSingular param.Opt[string] `json:"labelSingular,omitzero"`
	// Display icon, or null to use the default.
	//
	// Any of "Add", "Agent", "AlarmClock", "Archive", "Attachment", "Bell", "Book",
	// "Briefcase", "Calendar", "CameraOutline", "ChartArea", "ChartLine", "ChartPie",
	// "Checkmark", "ChevronRightDouble", "CircleAdd", "Clock", "Code", "Command",
	// "Competitors", "Compose", "CreditCard", "DealValue", "DeviceDesktop", "Dice",
	// "Document", "Download", "Draggable", "EmployeeBadge", "EngagementAnalyticsAlt",
	// "Enrich", "Folder", "Funding", "Grid", "Hashtag", "Help", "Key", "Keyboard",
	// "LightBulb", "Link", "Mail", "MicrophoneOutline", "Money", "Moon",
	// "MultiplePagesPortrait", "OpenQuote", "PageLandscape", "Paperplane", "People",
	// "Phone", "Preferences", "Search", "Shapes", "Sparkle", "SquareStack", "Star",
	// "Suggestions", "Sun", "Table", "TableColumn", "Tag", "Teams", "ThreeDCube",
	// "ThreeDStack", "VideoAlt", "View", "Waveform".
	Icon ObjectUpdateSettingsParamsIcon `json:"icon,omitzero"`
	paramObj
}

func (r ObjectUpdateSettingsParams) MarshalJSON() (data []byte, err error) {
	type shadow ObjectUpdateSettingsParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ObjectUpdateSettingsParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Display icon, or null to use the default.
type ObjectUpdateSettingsParamsIcon string

const (
	ObjectUpdateSettingsParamsIconAdd                    ObjectUpdateSettingsParamsIcon = "Add"
	ObjectUpdateSettingsParamsIconAgent                  ObjectUpdateSettingsParamsIcon = "Agent"
	ObjectUpdateSettingsParamsIconAlarmClock             ObjectUpdateSettingsParamsIcon = "AlarmClock"
	ObjectUpdateSettingsParamsIconArchive                ObjectUpdateSettingsParamsIcon = "Archive"
	ObjectUpdateSettingsParamsIconAttachment             ObjectUpdateSettingsParamsIcon = "Attachment"
	ObjectUpdateSettingsParamsIconBell                   ObjectUpdateSettingsParamsIcon = "Bell"
	ObjectUpdateSettingsParamsIconBook                   ObjectUpdateSettingsParamsIcon = "Book"
	ObjectUpdateSettingsParamsIconBriefcase              ObjectUpdateSettingsParamsIcon = "Briefcase"
	ObjectUpdateSettingsParamsIconCalendar               ObjectUpdateSettingsParamsIcon = "Calendar"
	ObjectUpdateSettingsParamsIconCameraOutline          ObjectUpdateSettingsParamsIcon = "CameraOutline"
	ObjectUpdateSettingsParamsIconChartArea              ObjectUpdateSettingsParamsIcon = "ChartArea"
	ObjectUpdateSettingsParamsIconChartLine              ObjectUpdateSettingsParamsIcon = "ChartLine"
	ObjectUpdateSettingsParamsIconChartPie               ObjectUpdateSettingsParamsIcon = "ChartPie"
	ObjectUpdateSettingsParamsIconCheckmark              ObjectUpdateSettingsParamsIcon = "Checkmark"
	ObjectUpdateSettingsParamsIconChevronRightDouble     ObjectUpdateSettingsParamsIcon = "ChevronRightDouble"
	ObjectUpdateSettingsParamsIconCircleAdd              ObjectUpdateSettingsParamsIcon = "CircleAdd"
	ObjectUpdateSettingsParamsIconClock                  ObjectUpdateSettingsParamsIcon = "Clock"
	ObjectUpdateSettingsParamsIconCode                   ObjectUpdateSettingsParamsIcon = "Code"
	ObjectUpdateSettingsParamsIconCommand                ObjectUpdateSettingsParamsIcon = "Command"
	ObjectUpdateSettingsParamsIconCompetitors            ObjectUpdateSettingsParamsIcon = "Competitors"
	ObjectUpdateSettingsParamsIconCompose                ObjectUpdateSettingsParamsIcon = "Compose"
	ObjectUpdateSettingsParamsIconCreditCard             ObjectUpdateSettingsParamsIcon = "CreditCard"
	ObjectUpdateSettingsParamsIconDealValue              ObjectUpdateSettingsParamsIcon = "DealValue"
	ObjectUpdateSettingsParamsIconDeviceDesktop          ObjectUpdateSettingsParamsIcon = "DeviceDesktop"
	ObjectUpdateSettingsParamsIconDice                   ObjectUpdateSettingsParamsIcon = "Dice"
	ObjectUpdateSettingsParamsIconDocument               ObjectUpdateSettingsParamsIcon = "Document"
	ObjectUpdateSettingsParamsIconDownload               ObjectUpdateSettingsParamsIcon = "Download"
	ObjectUpdateSettingsParamsIconDraggable              ObjectUpdateSettingsParamsIcon = "Draggable"
	ObjectUpdateSettingsParamsIconEmployeeBadge          ObjectUpdateSettingsParamsIcon = "EmployeeBadge"
	ObjectUpdateSettingsParamsIconEngagementAnalyticsAlt ObjectUpdateSettingsParamsIcon = "EngagementAnalyticsAlt"
	ObjectUpdateSettingsParamsIconEnrich                 ObjectUpdateSettingsParamsIcon = "Enrich"
	ObjectUpdateSettingsParamsIconFolder                 ObjectUpdateSettingsParamsIcon = "Folder"
	ObjectUpdateSettingsParamsIconFunding                ObjectUpdateSettingsParamsIcon = "Funding"
	ObjectUpdateSettingsParamsIconGrid                   ObjectUpdateSettingsParamsIcon = "Grid"
	ObjectUpdateSettingsParamsIconHashtag                ObjectUpdateSettingsParamsIcon = "Hashtag"
	ObjectUpdateSettingsParamsIconHelp                   ObjectUpdateSettingsParamsIcon = "Help"
	ObjectUpdateSettingsParamsIconKey                    ObjectUpdateSettingsParamsIcon = "Key"
	ObjectUpdateSettingsParamsIconKeyboard               ObjectUpdateSettingsParamsIcon = "Keyboard"
	ObjectUpdateSettingsParamsIconLightBulb              ObjectUpdateSettingsParamsIcon = "LightBulb"
	ObjectUpdateSettingsParamsIconLink                   ObjectUpdateSettingsParamsIcon = "Link"
	ObjectUpdateSettingsParamsIconMail                   ObjectUpdateSettingsParamsIcon = "Mail"
	ObjectUpdateSettingsParamsIconMicrophoneOutline      ObjectUpdateSettingsParamsIcon = "MicrophoneOutline"
	ObjectUpdateSettingsParamsIconMoney                  ObjectUpdateSettingsParamsIcon = "Money"
	ObjectUpdateSettingsParamsIconMoon                   ObjectUpdateSettingsParamsIcon = "Moon"
	ObjectUpdateSettingsParamsIconMultiplePagesPortrait  ObjectUpdateSettingsParamsIcon = "MultiplePagesPortrait"
	ObjectUpdateSettingsParamsIconOpenQuote              ObjectUpdateSettingsParamsIcon = "OpenQuote"
	ObjectUpdateSettingsParamsIconPageLandscape          ObjectUpdateSettingsParamsIcon = "PageLandscape"
	ObjectUpdateSettingsParamsIconPaperplane             ObjectUpdateSettingsParamsIcon = "Paperplane"
	ObjectUpdateSettingsParamsIconPeople                 ObjectUpdateSettingsParamsIcon = "People"
	ObjectUpdateSettingsParamsIconPhone                  ObjectUpdateSettingsParamsIcon = "Phone"
	ObjectUpdateSettingsParamsIconPreferences            ObjectUpdateSettingsParamsIcon = "Preferences"
	ObjectUpdateSettingsParamsIconSearch                 ObjectUpdateSettingsParamsIcon = "Search"
	ObjectUpdateSettingsParamsIconShapes                 ObjectUpdateSettingsParamsIcon = "Shapes"
	ObjectUpdateSettingsParamsIconSparkle                ObjectUpdateSettingsParamsIcon = "Sparkle"
	ObjectUpdateSettingsParamsIconSquareStack            ObjectUpdateSettingsParamsIcon = "SquareStack"
	ObjectUpdateSettingsParamsIconStar                   ObjectUpdateSettingsParamsIcon = "Star"
	ObjectUpdateSettingsParamsIconSuggestions            ObjectUpdateSettingsParamsIcon = "Suggestions"
	ObjectUpdateSettingsParamsIconSun                    ObjectUpdateSettingsParamsIcon = "Sun"
	ObjectUpdateSettingsParamsIconTable                  ObjectUpdateSettingsParamsIcon = "Table"
	ObjectUpdateSettingsParamsIconTableColumn            ObjectUpdateSettingsParamsIcon = "TableColumn"
	ObjectUpdateSettingsParamsIconTag                    ObjectUpdateSettingsParamsIcon = "Tag"
	ObjectUpdateSettingsParamsIconTeams                  ObjectUpdateSettingsParamsIcon = "Teams"
	ObjectUpdateSettingsParamsIconThreeDCube             ObjectUpdateSettingsParamsIcon = "ThreeDCube"
	ObjectUpdateSettingsParamsIconThreeDStack            ObjectUpdateSettingsParamsIcon = "ThreeDStack"
	ObjectUpdateSettingsParamsIconVideoAlt               ObjectUpdateSettingsParamsIcon = "VideoAlt"
	ObjectUpdateSettingsParamsIconView                   ObjectUpdateSettingsParamsIcon = "View"
	ObjectUpdateSettingsParamsIconWaveform               ObjectUpdateSettingsParamsIcon = "Waveform"
)
