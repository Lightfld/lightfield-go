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

func TestMergeGetMerge(t *testing.T) {
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
	_, err := client.Merge.GetMerge(context.TODO(), "id")
	if err != nil {
		var apierr *githubcomlightfldlightfieldgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestMergeMergeAccountsWithOptionalParams(t *testing.T) {
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
	_, err := client.Merge.MergeAccounts(context.TODO(), githubcomlightfldlightfieldgo.MergeMergeAccountsParams{
		DuplicateID: "duplicateId",
		PrimaryID:   "primaryId",
		FieldResolutions: map[string]githubcomlightfldlightfieldgo.MergeMergeAccountsParamsFieldResolutionUnion{
			"foo": {
				OfMergeMergeAccountssFieldResolutionString: githubcomlightfldlightfieldgo.String("primary"),
			},
		},
		Options: githubcomlightfldlightfieldgo.MergeMergeAccountsParamsOptions{
			MultiSelectUnion: githubcomlightfldlightfieldgo.Bool(true),
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

func TestMergeMergeContactsWithOptionalParams(t *testing.T) {
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
	_, err := client.Merge.MergeContacts(context.TODO(), githubcomlightfldlightfieldgo.MergeMergeContactsParams{
		DuplicateID: "duplicateId",
		PrimaryID:   "primaryId",
		FieldResolutions: map[string]githubcomlightfldlightfieldgo.MergeMergeContactsParamsFieldResolutionUnion{
			"foo": {
				OfMergeMergeContactssFieldResolutionString: githubcomlightfldlightfieldgo.String("primary"),
			},
		},
		Options: githubcomlightfldlightfieldgo.MergeMergeContactsParamsOptions{
			MultiSelectUnion: githubcomlightfldlightfieldgo.Bool(true),
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

func TestMergeMergeObjectValuesWithOptionalParams(t *testing.T) {
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
	_, err := client.Merge.MergeObjectValues(
		context.TODO(),
		"entitySlug",
		githubcomlightfldlightfieldgo.MergeMergeObjectValuesParams{
			DuplicateID: "duplicateId",
			PrimaryID:   "primaryId",
			FieldResolutions: map[string]githubcomlightfldlightfieldgo.MergeMergeObjectValuesParamsFieldResolutionUnion{
				"foo": {
					OfMergeMergeObjectValuessFieldResolutionString: githubcomlightfldlightfieldgo.String("primary"),
				},
			},
			Options: githubcomlightfldlightfieldgo.MergeMergeObjectValuesParamsOptions{
				MultiSelectUnion: githubcomlightfldlightfieldgo.Bool(true),
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

func TestMergeMergeOpportunitiesWithOptionalParams(t *testing.T) {
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
	_, err := client.Merge.MergeOpportunities(context.TODO(), githubcomlightfldlightfieldgo.MergeMergeOpportunitiesParams{
		DuplicateID: "duplicateId",
		PrimaryID:   "primaryId",
		FieldResolutions: map[string]githubcomlightfldlightfieldgo.MergeMergeOpportunitiesParamsFieldResolutionUnion{
			"foo": {
				OfMergeMergeOpportunitiessFieldResolutionString: githubcomlightfldlightfieldgo.String("primary"),
			},
		},
		Options: githubcomlightfldlightfieldgo.MergeMergeOpportunitiesParamsOptions{
			MultiSelectUnion: githubcomlightfldlightfieldgo.Bool(true),
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
