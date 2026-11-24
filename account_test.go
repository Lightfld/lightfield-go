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

func TestAccountNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Account.New(context.TODO(), githubcomlightfldlightfieldgo.AccountNewParams{
		Fields: map[string]githubcomlightfldlightfieldgo.AccountNewParamsFieldUnion{
			"foo": {
				OfString: githubcomlightfldlightfieldgo.String("string"),
			},
		},
		Relationships: map[string]githubcomlightfldlightfieldgo.AccountNewParamsRelationshipUnion{
			"foo": {
				OfString: githubcomlightfldlightfieldgo.String("string"),
			},
		},
		SkipEnrichment: githubcomlightfldlightfieldgo.Bool(true),
	})
	if err != nil {
		var apierr *githubcomlightfldlightfieldgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAccountGet(t *testing.T) {
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
	_, err := client.Account.Get(context.TODO(), "id")
	if err != nil {
		var apierr *githubcomlightfldlightfieldgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAccountUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Account.Update(
		context.TODO(),
		"id",
		githubcomlightfldlightfieldgo.AccountUpdateParams{
			Fields: map[string]githubcomlightfldlightfieldgo.AccountUpdateParamsFieldUnion{
				"foo": {
					OfString: githubcomlightfldlightfieldgo.String("string"),
				},
			},
			Relationships: map[string]githubcomlightfldlightfieldgo.AccountUpdateParamsRelationship{
				"foo": {
					Add: githubcomlightfldlightfieldgo.AccountUpdateParamsRelationshipAddUnion{
						OfString: githubcomlightfldlightfieldgo.String("string"),
					},
					Remove: githubcomlightfldlightfieldgo.AccountUpdateParamsRelationshipRemoveUnion{
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

func TestAccountListWithOptionalParams(t *testing.T) {
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
	_, err := client.Account.List(context.TODO(), githubcomlightfldlightfieldgo.AccountListParams{
		Fields: githubcomlightfldlightfieldgo.String("fields"),
		Filter: githubcomlightfldlightfieldgo.String("filter"),
		Limit:  githubcomlightfldlightfieldgo.Int(1),
		Offset: githubcomlightfldlightfieldgo.Int(0),
		Sort:   githubcomlightfldlightfieldgo.AccountListParamsSortCreatedAt,
		Status: githubcomlightfldlightfieldgo.AccountListParamsStatusActive,
	})
	if err != nil {
		var apierr *githubcomlightfldlightfieldgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAccountDeleteWithOptionalParams(t *testing.T) {
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
	_, err := client.Account.Delete(
		context.TODO(),
		"id",
		githubcomlightfldlightfieldgo.AccountDeleteParams{
			Body: githubcomlightfldlightfieldgo.AccountDeleteParamsBody{},
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

func TestAccountNewPipelineWithOptionalParams(t *testing.T) {
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
	_, err := client.Account.NewPipeline(context.TODO(), githubcomlightfldlightfieldgo.AccountNewPipelineParams{
		Name:            "name",
		Description:     githubcomlightfldlightfieldgo.String("description"),
		ForecastEnabled: githubcomlightfldlightfieldgo.Bool(true),
		Slug:            githubcomlightfldlightfieldgo.String("slug"),
		Stages: []githubcomlightfldlightfieldgo.AccountNewPipelineParamsStage{{
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

func TestAccountDefinitionsWithOptionalParams(t *testing.T) {
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
	_, err := client.Account.Definitions(context.TODO(), githubcomlightfldlightfieldgo.AccountDefinitionsParams{
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

func TestAccountDeletePipelineWithOptionalParams(t *testing.T) {
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
	_, err := client.Account.DeletePipeline(
		context.TODO(),
		"id",
		githubcomlightfldlightfieldgo.AccountDeletePipelineParams{
			DryRun: githubcomlightfldlightfieldgo.Bool(true),
			Body:   githubcomlightfldlightfieldgo.AccountDeletePipelineParamsBody{},
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

func TestAccountFieldHistoryWithOptionalParams(t *testing.T) {
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
	_, err := client.Account.FieldHistory(
		context.TODO(),
		"fieldKey",
		githubcomlightfldlightfieldgo.AccountFieldHistoryParams{
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

func TestAccountListPipelines(t *testing.T) {
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
	_, err := client.Account.ListPipelines(context.TODO())
	if err != nil {
		var apierr *githubcomlightfldlightfieldgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAccountRestoreWithOptionalParams(t *testing.T) {
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
	_, err := client.Account.Restore(
		context.TODO(),
		"id",
		githubcomlightfldlightfieldgo.AccountRestoreParams{
			Body: githubcomlightfldlightfieldgo.AccountRestoreParamsBody{},
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

func TestAccountGetPipeline(t *testing.T) {
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
	_, err := client.Account.GetPipeline(context.TODO(), "id")
	if err != nil {
		var apierr *githubcomlightfldlightfieldgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAccountUpdateDefinitionsWithOptionalParams(t *testing.T) {
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
	_, err := client.Account.UpdateDefinitions(context.TODO(), githubcomlightfldlightfieldgo.AccountUpdateDefinitionsParams{
		FieldDefinitions: map[string]githubcomlightfldlightfieldgo.AccountUpdateDefinitionsParamsFieldDefinition{
			"foo": {
				AIFillable:       "ON",
				AIFillablePrompt: githubcomlightfldlightfieldgo.String("aiFillablePrompt"),
				Archived:         githubcomlightfldlightfieldgo.Bool(true),
				Description:      githubcomlightfldlightfieldgo.String("description"),
				Formula:          githubcomlightfldlightfieldgo.String("formula"),
				Label:            githubcomlightfldlightfieldgo.String("label"),
				TypeConfiguration: githubcomlightfldlightfieldgo.AccountUpdateDefinitionsParamsFieldDefinitionTypeConfiguration{
					Currency:       githubcomlightfldlightfieldgo.String("currency"),
					DecimalPlaces:  githubcomlightfldlightfieldgo.Int(-9007199254740991),
					HandleService:  "TWITTER",
					MultipleValues: githubcomlightfldlightfieldgo.Bool(true),
					Options: []githubcomlightfldlightfieldgo.AccountUpdateDefinitionsParamsFieldDefinitionTypeConfigurationOption{{
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
		RelationshipDefinitions: map[string]githubcomlightfldlightfieldgo.AccountUpdateDefinitionsParamsRelationshipDefinition{
			"foo": {
				Archived:    githubcomlightfldlightfieldgo.Bool(true),
				Cardinality: "HAS_ONE",
				Description: githubcomlightfldlightfieldgo.String("description"),
				GroupKey:    githubcomlightfldlightfieldgo.String("groupKey"),
				Inverse: githubcomlightfldlightfieldgo.AccountUpdateDefinitionsParamsRelationshipDefinitionInverse{
					Cardinality: "HAS_ONE",
					Label:       githubcomlightfldlightfieldgo.String("label"),
				},
				Label:      githubcomlightfldlightfieldgo.String("label"),
				ObjectType: githubcomlightfldlightfieldgo.String("objectType"),
			},
		},
		RelationshipGroupDefinitions: map[string]githubcomlightfldlightfieldgo.AccountUpdateDefinitionsParamsRelationshipGroupDefinition{
			"foo": {
				Archived: githubcomlightfldlightfieldgo.Bool(true),
				Inverse: githubcomlightfldlightfieldgo.AccountUpdateDefinitionsParamsRelationshipGroupDefinitionInverse{
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

func TestAccountUpdatePipelineWithOptionalParams(t *testing.T) {
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
	_, err := client.Account.UpdatePipeline(
		context.TODO(),
		"id",
		githubcomlightfldlightfieldgo.AccountUpdatePipelineParams{
			DryRun:          githubcomlightfldlightfieldgo.Bool(true),
			Force:           githubcomlightfldlightfieldgo.Bool(true),
			Description:     githubcomlightfldlightfieldgo.String("description"),
			ForecastEnabled: githubcomlightfldlightfieldgo.Bool(true),
			Name:            githubcomlightfldlightfieldgo.String("name"),
			Stages: []githubcomlightfldlightfieldgo.AccountUpdatePipelineParamsStage{{
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
