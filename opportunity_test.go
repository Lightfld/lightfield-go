// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package lightfield_test

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
	client := lightfield.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Opportunity.New(context.TODO(), lightfield.OpportunityNewParams{
		Fields: lightfield.OpportunityNewParamsFields{
			SystemName:  "system_name",
			SystemStage: "system_stage",
		},
		Relationships: lightfield.OpportunityNewParamsRelationships{
			SystemAccount: lightfield.OpportunityNewParamsRelationshipsSystemAccountUnion{
				OfString: lightfield.String("string"),
			},
			SystemChampion: lightfield.OpportunityNewParamsRelationshipsSystemChampionUnion{
				OfString: lightfield.String("string"),
			},
			SystemCreatedBy: lightfield.OpportunityNewParamsRelationshipsSystemCreatedByUnion{
				OfString: lightfield.String("string"),
			},
			SystemEvaluator: lightfield.OpportunityNewParamsRelationshipsSystemEvaluatorUnion{
				OfString: lightfield.String("string"),
			},
			SystemOwner: lightfield.OpportunityNewParamsRelationshipsSystemOwnerUnion{
				OfString: lightfield.String("string"),
			},
		},
	})
	if err != nil {
		var apierr *lightfield.Error
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
	client := lightfield.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Opportunity.Get(context.TODO(), "id")
	if err != nil {
		var apierr *lightfield.Error
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
	client := lightfield.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Opportunity.Update(
		context.TODO(),
		"id",
		lightfield.OpportunityUpdateParams{
			Fields: lightfield.OpportunityUpdateParamsFields{
				SystemName:  lightfield.String("system_name"),
				SystemStage: lightfield.String("system_stage"),
			},
			Relationships: lightfield.OpportunityUpdateParamsRelationships{
				SystemAccount: lightfield.OpportunityUpdateParamsRelationshipsSystemAccount{
					Add: lightfield.OpportunityUpdateParamsRelationshipsSystemAccountAddUnion{
						OfString: lightfield.String("string"),
					},
					Remove: lightfield.OpportunityUpdateParamsRelationshipsSystemAccountRemoveUnion{
						OfString: lightfield.String("string"),
					},
					Replace: lightfield.OpportunityUpdateParamsRelationshipsSystemAccountReplaceUnion{
						OfString: lightfield.String("string"),
					},
				},
				SystemChampion: lightfield.OpportunityUpdateParamsRelationshipsSystemChampion{
					Add: lightfield.OpportunityUpdateParamsRelationshipsSystemChampionAddUnion{
						OfString: lightfield.String("string"),
					},
					Remove: lightfield.OpportunityUpdateParamsRelationshipsSystemChampionRemoveUnion{
						OfString: lightfield.String("string"),
					},
					Replace: lightfield.OpportunityUpdateParamsRelationshipsSystemChampionReplaceUnion{
						OfString: lightfield.String("string"),
					},
				},
				SystemCreatedBy: lightfield.OpportunityUpdateParamsRelationshipsSystemCreatedBy{
					Add: lightfield.OpportunityUpdateParamsRelationshipsSystemCreatedByAddUnion{
						OfString: lightfield.String("string"),
					},
					Remove: lightfield.OpportunityUpdateParamsRelationshipsSystemCreatedByRemoveUnion{
						OfString: lightfield.String("string"),
					},
					Replace: lightfield.OpportunityUpdateParamsRelationshipsSystemCreatedByReplaceUnion{
						OfString: lightfield.String("string"),
					},
				},
				SystemEvaluator: lightfield.OpportunityUpdateParamsRelationshipsSystemEvaluator{
					Add: lightfield.OpportunityUpdateParamsRelationshipsSystemEvaluatorAddUnion{
						OfString: lightfield.String("string"),
					},
					Remove: lightfield.OpportunityUpdateParamsRelationshipsSystemEvaluatorRemoveUnion{
						OfString: lightfield.String("string"),
					},
					Replace: lightfield.OpportunityUpdateParamsRelationshipsSystemEvaluatorReplaceUnion{
						OfString: lightfield.String("string"),
					},
				},
				SystemOwner: lightfield.OpportunityUpdateParamsRelationshipsSystemOwner{
					Add: lightfield.OpportunityUpdateParamsRelationshipsSystemOwnerAddUnion{
						OfString: lightfield.String("string"),
					},
					Remove: lightfield.OpportunityUpdateParamsRelationshipsSystemOwnerRemoveUnion{
						OfString: lightfield.String("string"),
					},
					Replace: lightfield.OpportunityUpdateParamsRelationshipsSystemOwnerReplaceUnion{
						OfString: lightfield.String("string"),
					},
				},
			},
		},
	)
	if err != nil {
		var apierr *lightfield.Error
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
	client := lightfield.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Opportunity.List(context.TODO(), lightfield.OpportunityListParams{
		Limit:  lightfield.Int(1),
		Offset: lightfield.Int(0),
	})
	if err != nil {
		var apierr *lightfield.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
