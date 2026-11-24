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

func TestOpportunityNew(t *testing.T) {
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
	_, err := client.Opportunity.New(context.TODO(), githubcomlightfldlightfieldgo.OpportunityNewParams{
		Fields: map[string]githubcomlightfldlightfieldgo.OpportunityNewParamsFieldUnion{
			"foo": {
				OfString: githubcomlightfldlightfieldgo.String("string"),
			},
		},
		Relationships: map[string]githubcomlightfldlightfieldgo.OpportunityNewParamsRelationshipUnion{
			"foo": {
				OfString: githubcomlightfldlightfieldgo.String("string"),
			},
		},
	})
	if err != nil {
		var apierr *githubcomlightfldlightfieldgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestOpportunityGet(t *testing.T) {
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
	_, err := client.Opportunity.Get(context.TODO(), "id")
	if err != nil {
		var apierr *githubcomlightfldlightfieldgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestOpportunityUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Opportunity.Update(
		context.TODO(),
		"id",
		githubcomlightfldlightfieldgo.OpportunityUpdateParams{
			Fields: map[string]githubcomlightfldlightfieldgo.OpportunityUpdateParamsFieldUnion{
				"foo": {
					OfString: githubcomlightfldlightfieldgo.String("string"),
				},
			},
			Relationships: map[string]githubcomlightfldlightfieldgo.OpportunityUpdateParamsRelationship{
				"foo": {
					Add: githubcomlightfldlightfieldgo.OpportunityUpdateParamsRelationshipAddUnion{
						OfString: githubcomlightfldlightfieldgo.String("string"),
					},
					Remove: githubcomlightfldlightfieldgo.OpportunityUpdateParamsRelationshipRemoveUnion{
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

func TestOpportunityListWithOptionalParams(t *testing.T) {
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
	_, err := client.Opportunity.List(context.TODO(), githubcomlightfldlightfieldgo.OpportunityListParams{
		Fields: githubcomlightfldlightfieldgo.String("fields"),
		Filter: githubcomlightfldlightfieldgo.String("filter"),
		Limit:  githubcomlightfldlightfieldgo.Int(1),
		Offset: githubcomlightfldlightfieldgo.Int(0),
		Sort:   githubcomlightfldlightfieldgo.OpportunityListParamsSortCreatedAt,
		Status: githubcomlightfldlightfieldgo.OpportunityListParamsStatusActive,
	})
	if err != nil {
		var apierr *githubcomlightfldlightfieldgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestOpportunityDeleteWithOptionalParams(t *testing.T) {
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
	_, err := client.Opportunity.Delete(
		context.TODO(),
		"id",
		githubcomlightfldlightfieldgo.OpportunityDeleteParams{
			Body: githubcomlightfldlightfieldgo.OpportunityDeleteParamsBody{},
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

func TestOpportunityNewPipelineWithOptionalParams(t *testing.T) {
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
	_, err := client.Opportunity.NewPipeline(context.TODO(), githubcomlightfldlightfieldgo.OpportunityNewPipelineParams{
		Name:            "name",
		Description:     githubcomlightfldlightfieldgo.String("description"),
		ForecastEnabled: githubcomlightfldlightfieldgo.Bool(true),
		Slug:            githubcomlightfldlightfieldgo.String("slug"),
		Stages: []githubcomlightfldlightfieldgo.OpportunityNewPipelineParamsStage{{
			Category:           "ACTIVE",
			Color:              "RED",
			Icon:               "BACKLOG",
			Label:              "label",
			ID:                 githubcomlightfldlightfieldgo.String("id"),
			Description:        githubcomlightfldlightfieldgo.String("description"),
			ForecastPercentage: githubcomlightfldlightfieldgo.Float(0),
		}},
	})
	if err != nil {
		var apierr *githubcomlightfldlightfieldgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestOpportunityDefinitionsWithOptionalParams(t *testing.T) {
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
	_, err := client.Opportunity.Definitions(context.TODO(), githubcomlightfldlightfieldgo.OpportunityDefinitionsParams{
		IncludeArchived: githubcomlightfldlightfieldgo.Bool(true),
	})
	if err != nil {
		var apierr *githubcomlightfldlightfieldgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestOpportunityDeletePipelineWithOptionalParams(t *testing.T) {
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
	_, err := client.Opportunity.DeletePipeline(
		context.TODO(),
		"id",
		githubcomlightfldlightfieldgo.OpportunityDeletePipelineParams{
			DryRun: githubcomlightfldlightfieldgo.Bool(true),
			Body:   githubcomlightfldlightfieldgo.OpportunityDeletePipelineParamsBody{},
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

func TestOpportunityFieldHistoryWithOptionalParams(t *testing.T) {
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
	_, err := client.Opportunity.FieldHistory(
		context.TODO(),
		"fieldKey",
		githubcomlightfldlightfieldgo.OpportunityFieldHistoryParams{
			ID:    "id",
			After: githubcomlightfldlightfieldgo.String("after"),
			Limit: githubcomlightfldlightfieldgo.Int(1),
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

func TestOpportunityRestoreWithOptionalParams(t *testing.T) {
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
	_, err := client.Opportunity.Restore(
		context.TODO(),
		"id",
		githubcomlightfldlightfieldgo.OpportunityRestoreParams{
			Body: githubcomlightfldlightfieldgo.OpportunityRestoreParamsBody{},
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

func TestOpportunityUpdateDefinitionsWithOptionalParams(t *testing.T) {
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
	_, err := client.Opportunity.UpdateDefinitions(context.TODO(), githubcomlightfldlightfieldgo.OpportunityUpdateDefinitionsParams{
		FieldDefinitions: map[string]githubcomlightfldlightfieldgo.OpportunityUpdateDefinitionsParamsFieldDefinition{
			"foo": {
				AIFillable:       "ON",
				AIFillablePrompt: githubcomlightfldlightfieldgo.String("aiFillablePrompt"),
				Archived:         githubcomlightfldlightfieldgo.Bool(true),
				Description:      githubcomlightfldlightfieldgo.String("description"),
				Formula:          githubcomlightfldlightfieldgo.String("formula"),
				Label:            githubcomlightfldlightfieldgo.String("label"),
				TypeConfiguration: githubcomlightfldlightfieldgo.OpportunityUpdateDefinitionsParamsFieldDefinitionTypeConfiguration{
					Currency:       githubcomlightfldlightfieldgo.String("currency"),
					DecimalPlaces:  githubcomlightfldlightfieldgo.Int(-9007199254740991),
					HandleService:  "TWITTER",
					MultipleValues: githubcomlightfldlightfieldgo.Bool(true),
					Options: []githubcomlightfldlightfieldgo.OpportunityUpdateDefinitionsParamsFieldDefinitionTypeConfigurationOption{{
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
		RelationshipDefinitions: map[string]githubcomlightfldlightfieldgo.OpportunityUpdateDefinitionsParamsRelationshipDefinition{
			"foo": {
				Archived:    githubcomlightfldlightfieldgo.Bool(true),
				Cardinality: "HAS_ONE",
				Description: githubcomlightfldlightfieldgo.String("description"),
				GroupKey:    githubcomlightfldlightfieldgo.String("groupKey"),
				Inverse: githubcomlightfldlightfieldgo.OpportunityUpdateDefinitionsParamsRelationshipDefinitionInverse{
					Cardinality: "HAS_ONE",
					Label:       githubcomlightfldlightfieldgo.String("label"),
				},
				Label:      githubcomlightfldlightfieldgo.String("label"),
				ObjectType: githubcomlightfldlightfieldgo.String("objectType"),
			},
		},
		RelationshipGroupDefinitions: map[string]githubcomlightfldlightfieldgo.OpportunityUpdateDefinitionsParamsRelationshipGroupDefinition{
			"foo": {
				Archived: githubcomlightfldlightfieldgo.Bool(true),
				Inverse: githubcomlightfldlightfieldgo.OpportunityUpdateDefinitionsParamsRelationshipGroupDefinitionInverse{
					Label: githubcomlightfldlightfieldgo.String("label"),
				},
				Label:      githubcomlightfldlightfieldgo.String("label"),
				ObjectType: githubcomlightfldlightfieldgo.String("objectType"),
			},
		},
	})
	if err != nil {
		var apierr *githubcomlightfldlightfieldgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestOpportunityUpdatePipelineWithOptionalParams(t *testing.T) {
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
	_, err := client.Opportunity.UpdatePipeline(
		context.TODO(),
		"id",
		githubcomlightfldlightfieldgo.OpportunityUpdatePipelineParams{
			DryRun:          githubcomlightfldlightfieldgo.Bool(true),
			Force:           githubcomlightfldlightfieldgo.Bool(true),
			Description:     githubcomlightfldlightfieldgo.String("description"),
			ForecastEnabled: githubcomlightfldlightfieldgo.Bool(true),
			Name:            githubcomlightfldlightfieldgo.String("name"),
			Stages: []githubcomlightfldlightfieldgo.OpportunityUpdatePipelineParamsStage{{
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
