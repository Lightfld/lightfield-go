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

func TestNoteNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Note.New(context.TODO(), githubcomlightfldlightfieldgo.NoteNewParams{
		Fields: githubcomlightfldlightfieldgo.NoteNewParamsFields{
			Title:   "$title",
			Content: githubcomlightfldlightfieldgo.String("$content"),
		},
		Relationships: githubcomlightfldlightfieldgo.NoteNewParamsRelationships{
			Account: githubcomlightfldlightfieldgo.NoteNewParamsRelationshipsAccountUnion{
				OfString: githubcomlightfldlightfieldgo.String("string"),
			},
			Opportunity: githubcomlightfldlightfieldgo.NoteNewParamsRelationshipsOpportunityUnion{
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

func TestNoteGet(t *testing.T) {
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
	_, err := client.Note.Get(context.TODO(), "id")
	if err != nil {
		var apierr *githubcomlightfldlightfieldgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestNoteUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Note.Update(
		context.TODO(),
		"id",
		githubcomlightfldlightfieldgo.NoteUpdateParams{
			Fields: githubcomlightfldlightfieldgo.NoteUpdateParamsFields{
				Content: githubcomlightfldlightfieldgo.String("$content"),
				Title:   githubcomlightfldlightfieldgo.String("$title"),
			},
			Relationships: githubcomlightfldlightfieldgo.NoteUpdateParamsRelationships{
				Account: githubcomlightfldlightfieldgo.NoteUpdateParamsRelationshipsAccountUnion{
					OfNoteUpdatesRelationshipsAccountAdd: &githubcomlightfldlightfieldgo.NoteUpdateParamsRelationshipsAccountAdd{
						Add: githubcomlightfldlightfieldgo.NoteUpdateParamsRelationshipsAccountAddAddUnion{
							OfString: githubcomlightfldlightfieldgo.String("string"),
						},
					},
				},
				Opportunity: githubcomlightfldlightfieldgo.NoteUpdateParamsRelationshipsOpportunityUnion{
					OfNoteUpdatesRelationshipsOpportunityAdd: &githubcomlightfldlightfieldgo.NoteUpdateParamsRelationshipsOpportunityAdd{
						Add: githubcomlightfldlightfieldgo.NoteUpdateParamsRelationshipsOpportunityAddAddUnion{
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

func TestNoteListWithOptionalParams(t *testing.T) {
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
	_, err := client.Note.List(context.TODO(), githubcomlightfldlightfieldgo.NoteListParams{
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
