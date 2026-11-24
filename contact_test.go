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

func TestContactNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Contact.New(context.TODO(), githubcomlightfldlightfieldgo.ContactNewParams{
		Fields: map[string]githubcomlightfldlightfieldgo.ContactNewParamsFieldUnion{
			"foo": {
				OfString: githubcomlightfldlightfieldgo.String("string"),
			},
		},
		Relationships: map[string]githubcomlightfldlightfieldgo.ContactNewParamsRelationshipUnion{
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

func TestContactGet(t *testing.T) {
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
	_, err := client.Contact.Get(context.TODO(), "id")
	if err != nil {
		var apierr *githubcomlightfldlightfieldgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestContactUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Contact.Update(
		context.TODO(),
		"id",
		githubcomlightfldlightfieldgo.ContactUpdateParams{
			Fields: map[string]githubcomlightfldlightfieldgo.ContactUpdateParamsFieldUnion{
				"foo": {
					OfString: githubcomlightfldlightfieldgo.String("string"),
				},
			},
			Relationships: map[string]githubcomlightfldlightfieldgo.ContactUpdateParamsRelationship{
				"foo": {
					Add: githubcomlightfldlightfieldgo.ContactUpdateParamsRelationshipAddUnion{
						OfString: githubcomlightfldlightfieldgo.String("string"),
					},
					Remove: githubcomlightfldlightfieldgo.ContactUpdateParamsRelationshipRemoveUnion{
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

func TestContactListWithOptionalParams(t *testing.T) {
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
	_, err := client.Contact.List(context.TODO(), githubcomlightfldlightfieldgo.ContactListParams{
		Fields: githubcomlightfldlightfieldgo.String("fields"),
		Filter: githubcomlightfldlightfieldgo.String("filter"),
		Limit:  githubcomlightfldlightfieldgo.Int(1),
		Offset: githubcomlightfldlightfieldgo.Int(0),
		Sort:   githubcomlightfldlightfieldgo.ContactListParamsSortCreatedAt,
		Status: githubcomlightfldlightfieldgo.ContactListParamsStatusActive,
	})
	if err != nil {
		var apierr *githubcomlightfldlightfieldgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestContactDeleteWithOptionalParams(t *testing.T) {
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
	_, err := client.Contact.Delete(
		context.TODO(),
		"id",
		githubcomlightfldlightfieldgo.ContactDeleteParams{
			Body: githubcomlightfldlightfieldgo.ContactDeleteParamsBody{},
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

func TestContactNewPipelineWithOptionalParams(t *testing.T) {
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
	_, err := client.Contact.NewPipeline(context.TODO(), githubcomlightfldlightfieldgo.ContactNewPipelineParams{
		Name:            "name",
		Description:     githubcomlightfldlightfieldgo.String("description"),
		ForecastEnabled: githubcomlightfldlightfieldgo.Bool(true),
		Slug:            githubcomlightfldlightfieldgo.String("slug"),
		Stages: []githubcomlightfldlightfieldgo.ContactNewPipelineParamsStage{{
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

func TestContactDefinitionsWithOptionalParams(t *testing.T) {
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
	_, err := client.Contact.Definitions(context.TODO(), githubcomlightfldlightfieldgo.ContactDefinitionsParams{
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

func TestContactDeletePipelineWithOptionalParams(t *testing.T) {
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
	_, err := client.Contact.DeletePipeline(
		context.TODO(),
		"id",
		githubcomlightfldlightfieldgo.ContactDeletePipelineParams{
			DryRun: githubcomlightfldlightfieldgo.Bool(true),
			Body:   githubcomlightfldlightfieldgo.ContactDeletePipelineParamsBody{},
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

func TestContactFieldHistoryWithOptionalParams(t *testing.T) {
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
	_, err := client.Contact.FieldHistory(
		context.TODO(),
		"fieldKey",
		githubcomlightfldlightfieldgo.ContactFieldHistoryParams{
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

func TestContactListPipelines(t *testing.T) {
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
	_, err := client.Contact.ListPipelines(context.TODO())
	if err != nil {
		var apierr *githubcomlightfldlightfieldgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestContactRestoreWithOptionalParams(t *testing.T) {
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
	_, err := client.Contact.Restore(
		context.TODO(),
		"id",
		githubcomlightfldlightfieldgo.ContactRestoreParams{
			Body: githubcomlightfldlightfieldgo.ContactRestoreParamsBody{},
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

func TestContactGetPipeline(t *testing.T) {
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
	_, err := client.Contact.GetPipeline(context.TODO(), "id")
	if err != nil {
		var apierr *githubcomlightfldlightfieldgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestContactUpdateDefinitionsWithOptionalParams(t *testing.T) {
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
	_, err := client.Contact.UpdateDefinitions(context.TODO(), githubcomlightfldlightfieldgo.ContactUpdateDefinitionsParams{
		FieldDefinitions: map[string]githubcomlightfldlightfieldgo.ContactUpdateDefinitionsParamsFieldDefinition{
			"foo": {
				AIFillable:       "ON",
				AIFillablePrompt: githubcomlightfldlightfieldgo.String("aiFillablePrompt"),
				Archived:         githubcomlightfldlightfieldgo.Bool(true),
				Description:      githubcomlightfldlightfieldgo.String("description"),
				Formula:          githubcomlightfldlightfieldgo.String("formula"),
				Label:            githubcomlightfldlightfieldgo.String("label"),
				TypeConfiguration: githubcomlightfldlightfieldgo.ContactUpdateDefinitionsParamsFieldDefinitionTypeConfiguration{
					Currency:       githubcomlightfldlightfieldgo.String("currency"),
					DecimalPlaces:  githubcomlightfldlightfieldgo.Int(-9007199254740991),
					HandleService:  "TWITTER",
					MultipleValues: githubcomlightfldlightfieldgo.Bool(true),
					Options: []githubcomlightfldlightfieldgo.ContactUpdateDefinitionsParamsFieldDefinitionTypeConfigurationOption{{
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
		RelationshipDefinitions: map[string]githubcomlightfldlightfieldgo.ContactUpdateDefinitionsParamsRelationshipDefinition{
			"foo": {
				Archived:    githubcomlightfldlightfieldgo.Bool(true),
				Cardinality: "HAS_ONE",
				Description: githubcomlightfldlightfieldgo.String("description"),
				GroupKey:    githubcomlightfldlightfieldgo.String("groupKey"),
				Inverse: githubcomlightfldlightfieldgo.ContactUpdateDefinitionsParamsRelationshipDefinitionInverse{
					Cardinality: "HAS_ONE",
					Label:       githubcomlightfldlightfieldgo.String("label"),
				},
				Label:      githubcomlightfldlightfieldgo.String("label"),
				ObjectType: githubcomlightfldlightfieldgo.String("objectType"),
			},
		},
		RelationshipGroupDefinitions: map[string]githubcomlightfldlightfieldgo.ContactUpdateDefinitionsParamsRelationshipGroupDefinition{
			"foo": {
				Archived: githubcomlightfldlightfieldgo.Bool(true),
				Inverse: githubcomlightfldlightfieldgo.ContactUpdateDefinitionsParamsRelationshipGroupDefinitionInverse{
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

func TestContactUpdatePipelineWithOptionalParams(t *testing.T) {
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
	_, err := client.Contact.UpdatePipeline(
		context.TODO(),
		"id",
		githubcomlightfldlightfieldgo.ContactUpdatePipelineParams{
			DryRun:          githubcomlightfldlightfieldgo.Bool(true),
			Force:           githubcomlightfldlightfieldgo.Bool(true),
			Description:     githubcomlightfldlightfieldgo.String("description"),
			ForecastEnabled: githubcomlightfldlightfieldgo.Bool(true),
			Name:            githubcomlightfldlightfieldgo.String("name"),
			Stages: []githubcomlightfldlightfieldgo.ContactUpdatePipelineParamsStage{{
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
