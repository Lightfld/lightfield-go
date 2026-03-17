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
		Fields: githubcomlightfldlightfieldgo.ContactNewParamsFields{
			SystemEmail: []string{"string"},
			SystemName: githubcomlightfldlightfieldgo.ContactNewParamsFieldsSystemName{
				FirstName: githubcomlightfldlightfieldgo.String("firstName"),
				LastName:  githubcomlightfldlightfieldgo.String("lastName"),
			},
			SystemProfilePhotoURL: githubcomlightfldlightfieldgo.String("system_profilePhotoUrl"),
		},
		Relationships: githubcomlightfldlightfieldgo.ContactNewParamsRelationships{
			SystemAccount: githubcomlightfldlightfieldgo.ContactNewParamsRelationshipsSystemAccountUnion{
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
			Fields: githubcomlightfldlightfieldgo.ContactUpdateParamsFields{
				SystemEmail: []string{"string"},
				SystemName: githubcomlightfldlightfieldgo.ContactUpdateParamsFieldsSystemName{
					FirstName: githubcomlightfldlightfieldgo.String("firstName"),
					LastName:  githubcomlightfldlightfieldgo.String("lastName"),
				},
				SystemProfilePhotoURL: githubcomlightfldlightfieldgo.String("system_profilePhotoUrl"),
			},
			Relationships: githubcomlightfldlightfieldgo.ContactUpdateParamsRelationships{
				SystemAccount: githubcomlightfldlightfieldgo.ContactUpdateParamsRelationshipsSystemAccount{
					Add: githubcomlightfldlightfieldgo.ContactUpdateParamsRelationshipsSystemAccountAddUnion{
						OfString: githubcomlightfldlightfieldgo.String("string"),
					},
					Remove: githubcomlightfldlightfieldgo.ContactUpdateParamsRelationshipsSystemAccountRemoveUnion{
						OfString: githubcomlightfldlightfieldgo.String("string"),
					},
					Replace: githubcomlightfldlightfieldgo.ContactUpdateParamsRelationshipsSystemAccountReplaceUnion{
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
