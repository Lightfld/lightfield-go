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

func TestContactNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Contact.New(context.TODO(), lightfield.ContactNewParams{
		Fields: lightfield.ContactNewParamsFields{
			SystemEmail: []string{"string"},
			SystemName: lightfield.ContactNewParamsFieldsSystemName{
				FirstName: lightfield.String("firstName"),
				LastName:  lightfield.String("lastName"),
			},
			SystemProfilePhotoURL: lightfield.String("system_profilePhotoUrl"),
		},
		Relationships: lightfield.ContactNewParamsRelationships{
			SystemAccount: lightfield.ContactNewParamsRelationshipsSystemAccountUnion{
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

func TestContactGet(t *testing.T) {
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
	_, err := client.Contact.Get(context.TODO(), "id")
	if err != nil {
		var apierr *lightfield.Error
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
	client := lightfield.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Contact.Update(
		context.TODO(),
		"id",
		lightfield.ContactUpdateParams{
			Fields: lightfield.ContactUpdateParamsFields{
				SystemEmail: []string{"string"},
				SystemName: lightfield.ContactUpdateParamsFieldsSystemName{
					FirstName: lightfield.String("firstName"),
					LastName:  lightfield.String("lastName"),
				},
				SystemProfilePhotoURL: lightfield.String("system_profilePhotoUrl"),
			},
			Relationships: lightfield.ContactUpdateParamsRelationships{
				SystemAccount: lightfield.ContactUpdateParamsRelationshipsSystemAccount{
					Add: lightfield.ContactUpdateParamsRelationshipsSystemAccountAddUnion{
						OfString: lightfield.String("string"),
					},
					Remove: lightfield.ContactUpdateParamsRelationshipsSystemAccountRemoveUnion{
						OfString: lightfield.String("string"),
					},
					Replace: lightfield.ContactUpdateParamsRelationshipsSystemAccountReplaceUnion{
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

func TestContactListWithOptionalParams(t *testing.T) {
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
	_, err := client.Contact.List(context.TODO(), lightfield.ContactListParams{
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
