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

func TestListNewWithOptionalParams(t *testing.T) {
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
	_, err := client.List.New(context.TODO(), githubcomlightfldlightfieldgo.ListNewParams{
		Fields: githubcomlightfldlightfieldgo.ListNewParamsFields{
			Name:       "$name",
			ObjectType: "account",
		},
		Relationships: githubcomlightfldlightfieldgo.ListNewParamsRelationshipsUnion{
			OfListNewsRelationshipsAccounts: &githubcomlightfldlightfieldgo.ListNewParamsRelationshipsAccounts{
				Accounts: githubcomlightfldlightfieldgo.ListNewParamsRelationshipsAccountsAccountsUnion{
					OfString: githubcomlightfldlightfieldgo.String("string"),
				},
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

func TestListGet(t *testing.T) {
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
	_, err := client.List.Get(context.TODO(), "id")
	if err != nil {
		var apierr *githubcomlightfldlightfieldgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestListUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.List.Update(
		context.TODO(),
		"id",
		githubcomlightfldlightfieldgo.ListUpdateParams{
			Fields: githubcomlightfldlightfieldgo.ListUpdateParamsFields{
				Name: githubcomlightfldlightfieldgo.String("$name"),
			},
			Relationships: githubcomlightfldlightfieldgo.ListUpdateParamsRelationshipsUnion{
				OfListUpdatesRelationshipsAccounts: &githubcomlightfldlightfieldgo.ListUpdateParamsRelationshipsAccounts{
					Accounts: githubcomlightfldlightfieldgo.ListUpdateParamsRelationshipsAccountsAccounts{
						Add: githubcomlightfldlightfieldgo.ListUpdateParamsRelationshipsAccountsAccountsAddUnion{
							OfString: githubcomlightfldlightfieldgo.String("string"),
						},
						Remove: githubcomlightfldlightfieldgo.ListUpdateParamsRelationshipsAccountsAccountsRemoveUnion{
							OfString: githubcomlightfldlightfieldgo.String("string"),
						},
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

func TestListListWithOptionalParams(t *testing.T) {
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
	_, err := client.List.List(context.TODO(), githubcomlightfldlightfieldgo.ListListParams{
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

func TestListListAccountsWithOptionalParams(t *testing.T) {
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
	_, err := client.List.ListAccounts(
		context.TODO(),
		"listId",
		githubcomlightfldlightfieldgo.ListListAccountsParams{
			Limit:  githubcomlightfldlightfieldgo.Int(1),
			Offset: githubcomlightfldlightfieldgo.Int(0),
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

func TestListListContactsWithOptionalParams(t *testing.T) {
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
	_, err := client.List.ListContacts(
		context.TODO(),
		"listId",
		githubcomlightfldlightfieldgo.ListListContactsParams{
			Limit:  githubcomlightfldlightfieldgo.Int(1),
			Offset: githubcomlightfldlightfieldgo.Int(0),
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

func TestListListOpportunitiesWithOptionalParams(t *testing.T) {
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
	_, err := client.List.ListOpportunities(
		context.TODO(),
		"listId",
		githubcomlightfldlightfieldgo.ListListOpportunitiesParams{
			Limit:  githubcomlightfldlightfieldgo.Int(1),
			Offset: githubcomlightfldlightfieldgo.Int(0),
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
