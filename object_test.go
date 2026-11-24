// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package githubcomlightfldlightfieldgo_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/Lightfld/lightfield-go"
	"github.com/Lightfld/lightfield-go/internal/testutil"
	"github.com/Lightfld/lightfield-go/option"
)

func TestObjectNewWithOptionalParams(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := githubcomlightfldlightfieldgo.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Object.New(
		context.TODO(),
		"entitySlug",
		githubcomlightfldlightfieldgo.ObjectNewParams{
			Fields: map[string]githubcomlightfldlightfieldgo.ObjectNewParamsFieldUnion{
				"foo": {
					OfString: githubcomlightfldlightfieldgo.String("string"),
				},
			},
			Relationships: map[string]githubcomlightfldlightfieldgo.ObjectNewParamsRelationshipUnion{
				"foo": {
					OfString: githubcomlightfldlightfieldgo.String("string"),
				},
			},
		},
	)
	if err != nil {
		var apierr *githubcomlightfldlightfieldgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestObjectGet(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := githubcomlightfldlightfieldgo.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Object.Get(
		context.TODO(),
		"id",
		githubcomlightfldlightfieldgo.ObjectGetParams{
			EntitySlug: "entitySlug",
		},
	)
	if err != nil {
		var apierr *githubcomlightfldlightfieldgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestObjectUpdateWithOptionalParams(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := githubcomlightfldlightfieldgo.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Object.Update(
		context.TODO(),
		"id",
		githubcomlightfldlightfieldgo.ObjectUpdateParams{
			EntitySlug: "entitySlug",
			Fields: map[string]githubcomlightfldlightfieldgo.ObjectUpdateParamsFieldUnion{
				"foo": {
					OfString: githubcomlightfldlightfieldgo.String("string"),
				},
			},
			Relationships: map[string]githubcomlightfldlightfieldgo.ObjectUpdateParamsRelationship{
				"foo": {
					Add: githubcomlightfldlightfieldgo.ObjectUpdateParamsRelationshipAddUnion{
						OfString: githubcomlightfldlightfieldgo.String("string"),
					},
					Remove: githubcomlightfldlightfieldgo.ObjectUpdateParamsRelationshipRemoveUnion{
						OfString: githubcomlightfldlightfieldgo.String("string"),
					},
					Replace: githubcomlightfldlightfieldgo.String("replace"),
				},
			},
		},
	)
	if err != nil {
		var apierr *githubcomlightfldlightfieldgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestObjectListWithOptionalParams(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := githubcomlightfldlightfieldgo.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Object.List(
		context.TODO(),
		"entitySlug",
		githubcomlightfldlightfieldgo.ObjectListParams{
			Fields: githubcomlightfldlightfieldgo.String("fields"),
			Filter: githubcomlightfldlightfieldgo.String("filter"),
			Limit:  githubcomlightfldlightfieldgo.Int(1),
			Offset: githubcomlightfldlightfieldgo.Int(0),
			Sort:   githubcomlightfldlightfieldgo.ObjectListParamsSortCreatedAt,
		},
	)
	if err != nil {
		var apierr *githubcomlightfldlightfieldgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestObjectDeleteWithOptionalParams(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := githubcomlightfldlightfieldgo.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Object.Delete(
		context.TODO(),
		"id",
		githubcomlightfldlightfieldgo.ObjectDeleteParams{
			EntitySlug: "entitySlug",
			Body:       githubcomlightfldlightfieldgo.ObjectDeleteParamsBody{},
		},
	)
	if err != nil {
		var apierr *githubcomlightfldlightfieldgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestObjectArchiveTypeWithOptionalParams(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := githubcomlightfldlightfieldgo.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Object.ArchiveType(
		context.TODO(),
		"entitySlug",
		githubcomlightfldlightfieldgo.ObjectArchiveTypeParams{
			Body: githubcomlightfldlightfieldgo.ObjectArchiveTypeParamsBody{},
		},
	)
	if err != nil {
		var apierr *githubcomlightfldlightfieldgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestObjectNewPipelineWithOptionalParams(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := githubcomlightfldlightfieldgo.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Object.NewPipeline(
		context.TODO(),
		"entitySlug",
		githubcomlightfldlightfieldgo.ObjectNewPipelineParams{
			Name:            "name",
			Description:     githubcomlightfldlightfieldgo.String("description"),
			ForecastEnabled: githubcomlightfldlightfieldgo.Bool(true),
			Slug:            githubcomlightfldlightfieldgo.String("slug"),
			Stages: []githubcomlightfldlightfieldgo.ObjectNewPipelineParamsStage{{
				Category:           "ACTIVE",
				Color:              "RED",
				Icon:               "BACKLOG",
				Label:              "label",
				ID:                 githubcomlightfldlightfieldgo.String("id"),
				Description:        githubcomlightfldlightfieldgo.String("description"),
				ForecastPercentage: githubcomlightfldlightfieldgo.Float(0),
			}},
		},
	)
	if err != nil {
		var apierr *githubcomlightfldlightfieldgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestObjectNewTypeWithOptionalParams(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := githubcomlightfldlightfieldgo.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Object.NewType(context.TODO(), githubcomlightfldlightfieldgo.ObjectNewTypeParams{
		LabelSingular: "labelSingular",
		Slug:          "slug",
		Description:   githubcomlightfldlightfieldgo.String("description"),
		Icon:          githubcomlightfldlightfieldgo.ObjectNewTypeParamsIconAdd,
		LabelPlural:   githubcomlightfldlightfieldgo.String("labelPlural"),
	})
	if err != nil {
		var apierr *githubcomlightfldlightfieldgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestObjectDefinitionsWithOptionalParams(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := githubcomlightfldlightfieldgo.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Object.Definitions(
		context.TODO(),
		"entitySlug",
		githubcomlightfldlightfieldgo.ObjectDefinitionsParams{
			IncludeArchived: githubcomlightfldlightfieldgo.Bool(true),
		},
	)
	if err != nil {
		var apierr *githubcomlightfldlightfieldgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestObjectDeletePipelineWithOptionalParams(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := githubcomlightfldlightfieldgo.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Object.DeletePipeline(
		context.TODO(),
		"id",
		githubcomlightfldlightfieldgo.ObjectDeletePipelineParams{
			EntitySlug: "entitySlug",
			DryRun:     githubcomlightfldlightfieldgo.Bool(true),
			Body:       githubcomlightfldlightfieldgo.ObjectDeletePipelineParamsBody{},
		},
	)
	if err != nil {
		var apierr *githubcomlightfldlightfieldgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestObjectFieldHistoryWithOptionalParams(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := githubcomlightfldlightfieldgo.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Object.FieldHistory(
		context.TODO(),
		"fieldKey",
		githubcomlightfldlightfieldgo.ObjectFieldHistoryParams{
			EntitySlug: "entitySlug",
			ID:         "id",
			After:      githubcomlightfldlightfieldgo.String("after"),
			Limit:      githubcomlightfldlightfieldgo.Int(1),
		},
	)
	if err != nil {
		var apierr *githubcomlightfldlightfieldgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestObjectListDefinitions(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := githubcomlightfldlightfieldgo.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Object.ListDefinitions(context.TODO())
	if err != nil {
		var apierr *githubcomlightfldlightfieldgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestObjectListPipelines(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := githubcomlightfldlightfieldgo.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Object.ListPipelines(context.TODO(), "entitySlug")
	if err != nil {
		var apierr *githubcomlightfldlightfieldgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestObjectRestoreWithOptionalParams(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := githubcomlightfldlightfieldgo.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Object.Restore(
		context.TODO(),
		"id",
		githubcomlightfldlightfieldgo.ObjectRestoreParams{
			EntitySlug: "entitySlug",
			Body:       githubcomlightfldlightfieldgo.ObjectRestoreParamsBody{},
		},
	)
	if err != nil {
		var apierr *githubcomlightfldlightfieldgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestObjectGetPipeline(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := githubcomlightfldlightfieldgo.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Object.GetPipeline(
		context.TODO(),
		"id",
		githubcomlightfldlightfieldgo.ObjectGetPipelineParams{
			EntitySlug: "entitySlug",
		},
	)
	if err != nil {
		var apierr *githubcomlightfldlightfieldgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestObjectGetSettingsWithOptionalParams(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := githubcomlightfldlightfieldgo.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Object.GetSettings(
		context.TODO(),
		"entitySlug",
		githubcomlightfldlightfieldgo.ObjectGetSettingsParams{
			IncludeArchived: githubcomlightfldlightfieldgo.Bool(true),
		},
	)
	if err != nil {
		var apierr *githubcomlightfldlightfieldgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestObjectUpdateDefinitionsWithOptionalParams(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := githubcomlightfldlightfieldgo.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Object.UpdateDefinitions(
		context.TODO(),
		"entitySlug",
		githubcomlightfldlightfieldgo.ObjectUpdateDefinitionsParams{
			FieldDefinitions: map[string]githubcomlightfldlightfieldgo.ObjectUpdateDefinitionsParamsFieldDefinition{
				"foo": {
					AIFillable:       "ON",
					AIFillablePrompt: githubcomlightfldlightfieldgo.String("aiFillablePrompt"),
					Archived:         githubcomlightfldlightfieldgo.Bool(true),
					Description:      githubcomlightfldlightfieldgo.String("description"),
					Formula:          githubcomlightfldlightfieldgo.String("formula"),
					Label:            githubcomlightfldlightfieldgo.String("label"),
					TypeConfiguration: githubcomlightfldlightfieldgo.ObjectUpdateDefinitionsParamsFieldDefinitionTypeConfiguration{
						Currency:       githubcomlightfldlightfieldgo.String("currency"),
						DecimalPlaces:  githubcomlightfldlightfieldgo.Int(-9007199254740991),
						HandleService:  "TWITTER",
						MultipleValues: githubcomlightfldlightfieldgo.Bool(true),
						Options: []githubcomlightfldlightfieldgo.ObjectUpdateDefinitionsParamsFieldDefinitionTypeConfigurationOption{{
							ID:          "id",
							Label:       "label",
							Color:       "RED",
							Description: githubcomlightfldlightfieldgo.String("description"),
						}},
						ParentFieldKey: githubcomlightfldlightfieldgo.String("parentFieldKey"),
					},
					ValueType: "ADDRESS",
				},
			},
			RelationshipDefinitions: map[string]githubcomlightfldlightfieldgo.ObjectUpdateDefinitionsParamsRelationshipDefinition{
				"foo": {
					Archived:    githubcomlightfldlightfieldgo.Bool(true),
					Cardinality: "HAS_ONE",
					Description: githubcomlightfldlightfieldgo.String("description"),
					GroupKey:    githubcomlightfldlightfieldgo.String("groupKey"),
					Inverse: githubcomlightfldlightfieldgo.ObjectUpdateDefinitionsParamsRelationshipDefinitionInverse{
						Cardinality: "HAS_ONE",
						Label:       githubcomlightfldlightfieldgo.String("label"),
					},
					Label:      githubcomlightfldlightfieldgo.String("label"),
					ObjectType: githubcomlightfldlightfieldgo.String("objectType"),
				},
			},
			RelationshipGroupDefinitions: map[string]githubcomlightfldlightfieldgo.ObjectUpdateDefinitionsParamsRelationshipGroupDefinition{
				"foo": {
					Archived: githubcomlightfldlightfieldgo.Bool(true),
					Inverse: githubcomlightfldlightfieldgo.ObjectUpdateDefinitionsParamsRelationshipGroupDefinitionInverse{
						Label: githubcomlightfldlightfieldgo.String("label"),
					},
					Label:      githubcomlightfldlightfieldgo.String("label"),
					ObjectType: githubcomlightfldlightfieldgo.String("objectType"),
				},
			},
		},
	)
	if err != nil {
		var apierr *githubcomlightfldlightfieldgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestObjectUpdatePipelineWithOptionalParams(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := githubcomlightfldlightfieldgo.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Object.UpdatePipeline(
		context.TODO(),
		"id",
		githubcomlightfldlightfieldgo.ObjectUpdatePipelineParams{
			EntitySlug:      "entitySlug",
			DryRun:          githubcomlightfldlightfieldgo.Bool(true),
			Force:           githubcomlightfldlightfieldgo.Bool(true),
			Description:     githubcomlightfldlightfieldgo.String("description"),
			ForecastEnabled: githubcomlightfldlightfieldgo.Bool(true),
			Name:            githubcomlightfldlightfieldgo.String("name"),
			Stages: []githubcomlightfldlightfieldgo.ObjectUpdatePipelineParamsStage{{
				ID:                 githubcomlightfldlightfieldgo.String("id"),
				Category:           "ACTIVE",
				Color:              "RED",
				Description:        githubcomlightfldlightfieldgo.String("description"),
				ForecastPercentage: githubcomlightfldlightfieldgo.Float(0),
				Icon:               "BACKLOG",
				Label:              githubcomlightfldlightfieldgo.String("label"),
			}},
		},
	)
	if err != nil {
		var apierr *githubcomlightfldlightfieldgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestObjectUpdateSettingsWithOptionalParams(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := githubcomlightfldlightfieldgo.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Object.UpdateSettings(
		context.TODO(),
		"entitySlug",
		githubcomlightfldlightfieldgo.ObjectUpdateSettingsParams{
			Description:   githubcomlightfldlightfieldgo.String("description"),
			Icon:          githubcomlightfldlightfieldgo.ObjectUpdateSettingsParamsIconAdd,
			LabelPlural:   githubcomlightfldlightfieldgo.String("labelPlural"),
			LabelSingular: githubcomlightfldlightfieldgo.String("labelSingular"),
		},
	)
	if err != nil {
		var apierr *githubcomlightfldlightfieldgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
