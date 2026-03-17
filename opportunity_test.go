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
			SystemName:  "system_name",
			SystemStage: "system_stage",
		},
		Relationships: githubcomlightfldlightfieldgo.OpportunityNewParamsRelationships{
			SystemAccount: githubcomlightfldlightfieldgo.OpportunityNewParamsRelationshipsSystemAccountUnion{
				OfString: githubcomlightfldlightfieldgo.String("string"),
			},
			SystemChampion: githubcomlightfldlightfieldgo.OpportunityNewParamsRelationshipsSystemChampionUnion{
				OfString: githubcomlightfldlightfieldgo.String("string"),
			},
			SystemCreatedBy: githubcomlightfldlightfieldgo.OpportunityNewParamsRelationshipsSystemCreatedByUnion{
				OfString: githubcomlightfldlightfieldgo.String("string"),
			},
			SystemEvaluator: githubcomlightfldlightfieldgo.OpportunityNewParamsRelationshipsSystemEvaluatorUnion{
				OfString: githubcomlightfldlightfieldgo.String("string"),
			},
			SystemOwner: githubcomlightfldlightfieldgo.OpportunityNewParamsRelationshipsSystemOwnerUnion{
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
				SystemName:  githubcomlightfldlightfieldgo.String("system_name"),
				SystemStage: githubcomlightfldlightfieldgo.String("system_stage"),
			},
			Relationships: githubcomlightfldlightfieldgo.OpportunityUpdateParamsRelationships{
				SystemAccount: githubcomlightfldlightfieldgo.OpportunityUpdateParamsRelationshipsSystemAccount{
					Add: githubcomlightfldlightfieldgo.OpportunityUpdateParamsRelationshipsSystemAccountAddUnion{
						OfString: githubcomlightfldlightfieldgo.String("string"),
					},
					Remove: githubcomlightfldlightfieldgo.OpportunityUpdateParamsRelationshipsSystemAccountRemoveUnion{
						OfString: githubcomlightfldlightfieldgo.String("string"),
					},
					Replace: githubcomlightfldlightfieldgo.OpportunityUpdateParamsRelationshipsSystemAccountReplaceUnion{
						OfString: githubcomlightfldlightfieldgo.String("string"),
					},
				},
				SystemChampion: githubcomlightfldlightfieldgo.OpportunityUpdateParamsRelationshipsSystemChampion{
					Add: githubcomlightfldlightfieldgo.OpportunityUpdateParamsRelationshipsSystemChampionAddUnion{
						OfString: githubcomlightfldlightfieldgo.String("string"),
					},
					Remove: githubcomlightfldlightfieldgo.OpportunityUpdateParamsRelationshipsSystemChampionRemoveUnion{
						OfString: githubcomlightfldlightfieldgo.String("string"),
					},
					Replace: githubcomlightfldlightfieldgo.OpportunityUpdateParamsRelationshipsSystemChampionReplaceUnion{
						OfString: githubcomlightfldlightfieldgo.String("string"),
					},
				},
				SystemCreatedBy: githubcomlightfldlightfieldgo.OpportunityUpdateParamsRelationshipsSystemCreatedBy{
					Add: githubcomlightfldlightfieldgo.OpportunityUpdateParamsRelationshipsSystemCreatedByAddUnion{
						OfString: githubcomlightfldlightfieldgo.String("string"),
					},
					Remove: githubcomlightfldlightfieldgo.OpportunityUpdateParamsRelationshipsSystemCreatedByRemoveUnion{
						OfString: githubcomlightfldlightfieldgo.String("string"),
					},
					Replace: githubcomlightfldlightfieldgo.OpportunityUpdateParamsRelationshipsSystemCreatedByReplaceUnion{
						OfString: githubcomlightfldlightfieldgo.String("string"),
					},
				},
				SystemEvaluator: githubcomlightfldlightfieldgo.OpportunityUpdateParamsRelationshipsSystemEvaluator{
					Add: githubcomlightfldlightfieldgo.OpportunityUpdateParamsRelationshipsSystemEvaluatorAddUnion{
						OfString: githubcomlightfldlightfieldgo.String("string"),
					},
					Remove: githubcomlightfldlightfieldgo.OpportunityUpdateParamsRelationshipsSystemEvaluatorRemoveUnion{
						OfString: githubcomlightfldlightfieldgo.String("string"),
					},
					Replace: githubcomlightfldlightfieldgo.OpportunityUpdateParamsRelationshipsSystemEvaluatorReplaceUnion{
						OfString: githubcomlightfldlightfieldgo.String("string"),
					},
				},
				SystemOwner: githubcomlightfldlightfieldgo.OpportunityUpdateParamsRelationshipsSystemOwner{
					Add: githubcomlightfldlightfieldgo.OpportunityUpdateParamsRelationshipsSystemOwnerAddUnion{
						OfString: githubcomlightfldlightfieldgo.String("string"),
					},
					Remove: githubcomlightfldlightfieldgo.OpportunityUpdateParamsRelationshipsSystemOwnerRemoveUnion{
						OfString: githubcomlightfldlightfieldgo.String("string"),
					},
					Replace: githubcomlightfldlightfieldgo.OpportunityUpdateParamsRelationshipsSystemOwnerReplaceUnion{
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
