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
			Name:            "$name",
			Facebook:        githubcomlightfldlightfieldgo.String("$facebook"),
			Headcount:       githubcomlightfldlightfieldgo.String("$headcount"),
			Industry:        []string{"string"},
			Instagram:       githubcomlightfldlightfieldgo.String("$instagram"),
			LastFundingType: githubcomlightfldlightfieldgo.String("$lastFundingType"),
			LinkedIn:        githubcomlightfldlightfieldgo.String("$linkedIn"),
			PrimaryAddress: map[string]string{
				"foo": "string",
			},
			Twitter: githubcomlightfldlightfieldgo.String("$twitter"),
			Website: []string{"string"},
		},
		Relationships: githubcomlightfldlightfieldgo.AccountNewParamsRelationships{
			Contacts: githubcomlightfldlightfieldgo.AccountNewParamsRelationshipsContactsUnion{
				OfString: githubcomlightfldlightfieldgo.String("string"),
			},
			Owner: githubcomlightfldlightfieldgo.AccountNewParamsRelationshipsOwnerUnion{
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
				Facebook:        githubcomlightfldlightfieldgo.String("$facebook"),
				Headcount:       githubcomlightfldlightfieldgo.String("$headcount"),
				Industry:        []string{"string"},
				Instagram:       githubcomlightfldlightfieldgo.String("$instagram"),
				LastFundingType: githubcomlightfldlightfieldgo.String("$lastFundingType"),
				LinkedIn:        githubcomlightfldlightfieldgo.String("$linkedIn"),
				Name:            githubcomlightfldlightfieldgo.String("$name"),
				PrimaryAddress: map[string]string{
					"foo": "string",
				},
				Twitter: githubcomlightfldlightfieldgo.String("$twitter"),
				Website: []string{"string"},
			},
			Relationships: githubcomlightfldlightfieldgo.AccountUpdateParamsRelationships{
				Contacts: githubcomlightfldlightfieldgo.AccountUpdateParamsRelationshipsContacts{
					Add: githubcomlightfldlightfieldgo.AccountUpdateParamsRelationshipsContactsAddUnion{
						OfString: githubcomlightfldlightfieldgo.String("string"),
					},
					Remove: githubcomlightfldlightfieldgo.AccountUpdateParamsRelationshipsContactsRemoveUnion{
						OfString: githubcomlightfldlightfieldgo.String("string"),
					},
					Replace: githubcomlightfldlightfieldgo.AccountUpdateParamsRelationshipsContactsReplaceUnion{
						OfString: githubcomlightfldlightfieldgo.String("string"),
					},
				},
				Owner: githubcomlightfldlightfieldgo.AccountUpdateParamsRelationshipsOwner{
					Add: githubcomlightfldlightfieldgo.AccountUpdateParamsRelationshipsOwnerAddUnion{
						OfString: githubcomlightfldlightfieldgo.String("string"),
					},
					Remove: githubcomlightfldlightfieldgo.AccountUpdateParamsRelationshipsOwnerRemoveUnion{
						OfString: githubcomlightfldlightfieldgo.String("string"),
					},
					Replace: githubcomlightfldlightfieldgo.AccountUpdateParamsRelationshipsOwnerReplaceUnion{
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

func TestAccountDefinitions(t *testing.T) {
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
	_, err := client.Account.Definitions(context.TODO())
	if err != nil {
		var apierr *githubcomlightfldlightfieldgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
