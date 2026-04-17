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

func TestTaskNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Task.New(context.TODO(), githubcomlightfldlightfieldgo.TaskNewParams{
		Fields: githubcomlightfldlightfieldgo.TaskNewParamsFields{
			Status:      "$status",
			Title:       "$title",
			Description: githubcomlightfldlightfieldgo.String("$description"),
			DueAt:       githubcomlightfldlightfieldgo.String("$dueAt"),
		},
		Relationships: map[string]githubcomlightfldlightfieldgo.TaskNewParamsRelationshipUnion{
			"foo": {
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

func TestTaskGet(t *testing.T) {
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
	_, err := client.Task.Get(context.TODO(), "id")
	if err != nil {
		var apierr *githubcomlightfldlightfieldgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestTaskUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Task.Update(
		context.TODO(),
		"id",
		githubcomlightfldlightfieldgo.TaskUpdateParams{
			Fields: githubcomlightfldlightfieldgo.TaskUpdateParamsFields{
				Description: githubcomlightfldlightfieldgo.String("$description"),
				DueAt:       githubcomlightfldlightfieldgo.String("$dueAt"),
				Status:      githubcomlightfldlightfieldgo.String("$status"),
				Title:       githubcomlightfldlightfieldgo.String("$title"),
			},
			Relationships: map[string]githubcomlightfldlightfieldgo.TaskUpdateParamsRelationship{
				"foo": {
					Add: githubcomlightfldlightfieldgo.TaskUpdateParamsRelationshipAddUnion{
						OfString: githubcomlightfldlightfieldgo.String("string"),
					},
					Remove: githubcomlightfldlightfieldgo.TaskUpdateParamsRelationshipRemoveUnion{
						OfString: githubcomlightfldlightfieldgo.String("string"),
					},
					Replace: githubcomlightfldlightfieldgo.TaskUpdateParamsRelationshipReplaceUnion{
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

func TestTaskListWithOptionalParams(t *testing.T) {
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
	_, err := client.Task.List(context.TODO(), githubcomlightfldlightfieldgo.TaskListParams{
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

func TestTaskDefinitions(t *testing.T) {
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
	_, err := client.Task.Definitions(context.TODO())
	if err != nil {
		var apierr *githubcomlightfldlightfieldgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
