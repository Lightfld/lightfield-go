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

func TestOpportunityNewWithOptionalParams(t *testing.T) {
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
		Fields: githubcomlightfldlightfieldgo.OpportunityNewParamsFields{
			Name:  "$name",
			Stage: "$stage",
		},
		Relationships: githubcomlightfldlightfieldgo.OpportunityNewParamsRelationships{
			Account: githubcomlightfldlightfieldgo.OpportunityNewParamsRelationshipsAccountUnion{
				OfString: githubcomlightfldlightfieldgo.String("string"),
			},
			Champion: githubcomlightfldlightfieldgo.OpportunityNewParamsRelationshipsChampionUnion{
				OfString: githubcomlightfldlightfieldgo.String("string"),
			},
			CreatedBy: githubcomlightfldlightfieldgo.OpportunityNewParamsRelationshipsCreatedByUnion{
				OfString: githubcomlightfldlightfieldgo.String("string"),
			},
			Evaluator: githubcomlightfldlightfieldgo.OpportunityNewParamsRelationshipsEvaluatorUnion{
				OfString: githubcomlightfldlightfieldgo.String("string"),
			},
			Owner: githubcomlightfldlightfieldgo.OpportunityNewParamsRelationshipsOwnerUnion{
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
			Fields: githubcomlightfldlightfieldgo.OpportunityUpdateParamsFields{
				Name:  githubcomlightfldlightfieldgo.String("$name"),
				Stage: githubcomlightfldlightfieldgo.String("$stage"),
			},
			Relationships: githubcomlightfldlightfieldgo.OpportunityUpdateParamsRelationships{
				Champion: githubcomlightfldlightfieldgo.OpportunityUpdateParamsRelationshipsChampion{
					Add: githubcomlightfldlightfieldgo.OpportunityUpdateParamsRelationshipsChampionAddUnion{
						OfString: githubcomlightfldlightfieldgo.String("string"),
					},
					Remove: githubcomlightfldlightfieldgo.OpportunityUpdateParamsRelationshipsChampionRemoveUnion{
						OfString: githubcomlightfldlightfieldgo.String("string"),
					},
					Replace: githubcomlightfldlightfieldgo.OpportunityUpdateParamsRelationshipsChampionReplaceUnion{
						OfString: githubcomlightfldlightfieldgo.String("string"),
					},
				},
				Evaluator: githubcomlightfldlightfieldgo.OpportunityUpdateParamsRelationshipsEvaluator{
					Add: githubcomlightfldlightfieldgo.OpportunityUpdateParamsRelationshipsEvaluatorAddUnion{
						OfString: githubcomlightfldlightfieldgo.String("string"),
					},
					Remove: githubcomlightfldlightfieldgo.OpportunityUpdateParamsRelationshipsEvaluatorRemoveUnion{
						OfString: githubcomlightfldlightfieldgo.String("string"),
					},
					Replace: githubcomlightfldlightfieldgo.OpportunityUpdateParamsRelationshipsEvaluatorReplaceUnion{
						OfString: githubcomlightfldlightfieldgo.String("string"),
					},
				},
				Owner: githubcomlightfldlightfieldgo.OpportunityUpdateParamsRelationshipsOwner{
					Add: githubcomlightfldlightfieldgo.OpportunityUpdateParamsRelationshipsOwnerAddUnion{
						OfString: githubcomlightfldlightfieldgo.String("string"),
					},
					Remove: githubcomlightfldlightfieldgo.OpportunityUpdateParamsRelationshipsOwnerRemoveUnion{
						OfString: githubcomlightfldlightfieldgo.String("string"),
					},
					Replace: githubcomlightfldlightfieldgo.OpportunityUpdateParamsRelationshipsOwnerReplaceUnion{
						OfString: githubcomlightfldlightfieldgo.String("string"),
					},
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
		Limit:  githubcomlightfldlightfieldgo.Int(1),
		Offset: githubcomlightfldlightfieldgo.Int(0),
	})
	if err != nil {
		var apierr *githubcomlightfldlightfieldgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestOpportunityDefinitions(t *testing.T) {
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
	_, err := client.Opportunity.Definitions(context.TODO())
	if err != nil {
		var apierr *githubcomlightfldlightfieldgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
