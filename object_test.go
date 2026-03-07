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

func TestObjectNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Object.New(
		context.TODO(),
		lightfield.ObjectNewParamsEntityTypeAccounts,
		lightfield.ObjectNewParams{
			Fields: map[string]lightfield.ObjectNewParamsFieldUnion{
				"foo": {
					OfString: lightfield.String("string"),
				},
			},
			Relationships: map[string]lightfield.ObjectNewParamsRelationshipUnion{
				"foo": {
					OfString: lightfield.String("string"),
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

func TestObjectGet(t *testing.T) {
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
	_, err := client.Object.Get(
		context.TODO(),
		"id",
		lightfield.ObjectGetParams{
			EntityType: lightfield.ObjectGetParamsEntityTypeAccounts,
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

func TestObjectUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Object.Update(
		context.TODO(),
		"id",
		lightfield.ObjectUpdateParams{
			EntityType: lightfield.ObjectUpdateParamsEntityTypeAccounts,
			Fields: map[string]lightfield.ObjectUpdateParamsFieldUnion{
				"foo": {
					OfString: lightfield.String("string"),
				},
			},
			Relationships: map[string]lightfield.ObjectUpdateParamsRelationship{
				"foo": {
					Add: lightfield.ObjectUpdateParamsRelationshipAddUnion{
						OfString: lightfield.String("string"),
					},
					Remove: lightfield.ObjectUpdateParamsRelationshipRemoveUnion{
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

func TestObjectListWithOptionalParams(t *testing.T) {
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
	_, err := client.Object.List(
		context.TODO(),
		lightfield.ObjectListParamsEntityTypeAccounts,
		lightfield.ObjectListParams{
			Limit:  lightfield.Int(1),
			Offset: lightfield.Int(0),
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
