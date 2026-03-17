// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package lightfield_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/stainless-sdks/lightfield-go"
	"github.com/stainless-sdks/lightfield-go/internal/testutil"
	"github.com/stainless-sdks/lightfield-go/option"
)

func TestAccountNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Account.New(context.TODO(), lightfield.AccountNewParams{
		Fields: lightfield.AccountNewParamsFields{
			SystemName:            "system_name",
			SystemFacebook:        lightfield.String("system_facebook"),
			SystemHeadcount:       lightfield.String("system_headcount"),
			SystemIndustry:        []string{"string"},
			SystemInstagram:       lightfield.String("system_instagram"),
			SystemLastFundingType: lightfield.String("system_lastFundingType"),
			SystemLinkedIn:        lightfield.String("system_linkedIn"),
			SystemPrimaryAddress: map[string]string{
				"foo": "string",
			},
			SystemTwitter: lightfield.String("system_twitter"),
			SystemWebsite: []string{"string"},
		},
		Relationships: lightfield.AccountNewParamsRelationships{
			SystemContact: lightfield.AccountNewParamsRelationshipsSystemContactUnion{
				OfString: lightfield.String("string"),
			},
			SystemOwner: lightfield.AccountNewParamsRelationshipsSystemOwnerUnion{
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

func TestAccountGet(t *testing.T) {
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
	_, err := client.Account.Get(context.TODO(), "id")
	if err != nil {
		var apierr *lightfield.Error
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
	client := lightfield.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Account.Update(
		context.TODO(),
		"id",
		lightfield.AccountUpdateParams{
			Fields: lightfield.AccountUpdateParamsFields{
				SystemFacebook:        lightfield.String("system_facebook"),
				SystemHeadcount:       lightfield.String("system_headcount"),
				SystemIndustry:        []string{"string"},
				SystemInstagram:       lightfield.String("system_instagram"),
				SystemLastFundingType: lightfield.String("system_lastFundingType"),
				SystemLinkedIn:        lightfield.String("system_linkedIn"),
				SystemName:            lightfield.String("system_name"),
				SystemPrimaryAddress: map[string]string{
					"foo": "string",
				},
				SystemTwitter: lightfield.String("system_twitter"),
				SystemWebsite: []string{"string"},
			},
			Relationships: lightfield.AccountUpdateParamsRelationships{
				SystemContact: lightfield.AccountUpdateParamsRelationshipsSystemContact{
					Add: lightfield.AccountUpdateParamsRelationshipsSystemContactAddUnion{
						OfString: lightfield.String("string"),
					},
					Remove: lightfield.AccountUpdateParamsRelationshipsSystemContactRemoveUnion{
						OfString: lightfield.String("string"),
					},
					Replace: lightfield.AccountUpdateParamsRelationshipsSystemContactReplaceUnion{
						OfString: lightfield.String("string"),
					},
				},
				SystemOwner: lightfield.AccountUpdateParamsRelationshipsSystemOwner{
					Add: lightfield.AccountUpdateParamsRelationshipsSystemOwnerAddUnion{
						OfString: lightfield.String("string"),
					},
					Remove: lightfield.AccountUpdateParamsRelationshipsSystemOwnerRemoveUnion{
						OfString: lightfield.String("string"),
					},
					Replace: lightfield.AccountUpdateParamsRelationshipsSystemOwnerReplaceUnion{
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

func TestAccountListWithOptionalParams(t *testing.T) {
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
	_, err := client.Account.List(context.TODO(), lightfield.AccountListParams{
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
