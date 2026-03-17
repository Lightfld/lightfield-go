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
		Fields: githubcomlightfldlightfieldgo.AccountNewParamsFields{
			SystemName:            "system_name",
			SystemFacebook:        githubcomlightfldlightfieldgo.String("system_facebook"),
			SystemHeadcount:       githubcomlightfldlightfieldgo.String("system_headcount"),
			SystemIndustry:        []string{"string"},
			SystemInstagram:       githubcomlightfldlightfieldgo.String("system_instagram"),
			SystemLastFundingType: githubcomlightfldlightfieldgo.String("system_lastFundingType"),
			SystemLinkedIn:        githubcomlightfldlightfieldgo.String("system_linkedIn"),
			SystemPrimaryAddress: map[string]string{
				"foo": "string",
			},
			SystemTwitter: githubcomlightfldlightfieldgo.String("system_twitter"),
			SystemWebsite: []string{"string"},
		},
		Relationships: githubcomlightfldlightfieldgo.AccountNewParamsRelationships{
			SystemContact: githubcomlightfldlightfieldgo.AccountNewParamsRelationshipsSystemContactUnion{
				OfString: githubcomlightfldlightfieldgo.String("string"),
			},
			SystemOwner: githubcomlightfldlightfieldgo.AccountNewParamsRelationshipsSystemOwnerUnion{
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
			Fields: githubcomlightfldlightfieldgo.AccountUpdateParamsFields{
				SystemFacebook:        githubcomlightfldlightfieldgo.String("system_facebook"),
				SystemHeadcount:       githubcomlightfldlightfieldgo.String("system_headcount"),
				SystemIndustry:        []string{"string"},
				SystemInstagram:       githubcomlightfldlightfieldgo.String("system_instagram"),
				SystemLastFundingType: githubcomlightfldlightfieldgo.String("system_lastFundingType"),
				SystemLinkedIn:        githubcomlightfldlightfieldgo.String("system_linkedIn"),
				SystemName:            githubcomlightfldlightfieldgo.String("system_name"),
				SystemPrimaryAddress: map[string]string{
					"foo": "string",
				},
				SystemTwitter: githubcomlightfldlightfieldgo.String("system_twitter"),
				SystemWebsite: []string{"string"},
			},
			Relationships: githubcomlightfldlightfieldgo.AccountUpdateParamsRelationships{
				SystemContact: githubcomlightfldlightfieldgo.AccountUpdateParamsRelationshipsSystemContact{
					Add: githubcomlightfldlightfieldgo.AccountUpdateParamsRelationshipsSystemContactAddUnion{
						OfString: githubcomlightfldlightfieldgo.String("string"),
					},
					Remove: githubcomlightfldlightfieldgo.AccountUpdateParamsRelationshipsSystemContactRemoveUnion{
						OfString: githubcomlightfldlightfieldgo.String("string"),
					},
					Replace: githubcomlightfldlightfieldgo.AccountUpdateParamsRelationshipsSystemContactReplaceUnion{
						OfString: githubcomlightfldlightfieldgo.String("string"),
					},
				},
				SystemOwner: githubcomlightfldlightfieldgo.AccountUpdateParamsRelationshipsSystemOwner{
					Add: githubcomlightfldlightfieldgo.AccountUpdateParamsRelationshipsSystemOwnerAddUnion{
						OfString: githubcomlightfldlightfieldgo.String("string"),
					},
					Remove: githubcomlightfldlightfieldgo.AccountUpdateParamsRelationshipsSystemOwnerRemoveUnion{
						OfString: githubcomlightfldlightfieldgo.String("string"),
					},
					Replace: githubcomlightfldlightfieldgo.AccountUpdateParamsRelationshipsSystemOwnerReplaceUnion{
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
