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

func TestSequenceNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Sequence.New(context.TODO(), githubcomlightfldlightfieldgo.SequenceNewParams{
		Definition: githubcomlightfldlightfieldgo.SequenceNewParamsDefinition{
			Description: "x",
			Steps: []githubcomlightfldlightfieldgo.SequenceNewParamsDefinitionStepUnion{{
				OfSequenceNewsDefinitionStepObject: &githubcomlightfldlightfieldgo.SequenceNewParamsDefinitionStepObject{
					DeliveryMode: "AUTO",
					Name:         "x",
					Prompt:       "x",
					Type:         "EMAIL",
				},
			}},
			Messaging: githubcomlightfldlightfieldgo.SequenceNewParamsDefinitionMessaging{
				CallToAction: githubcomlightfldlightfieldgo.String("x"),
				Persona:      githubcomlightfldlightfieldgo.String("x"),
				Trigger:      githubcomlightfldlightfieldgo.String("x"),
				ValueProp:    githubcomlightfldlightfieldgo.String("x"),
			},
			Pacing: githubcomlightfldlightfieldgo.SequenceNewParamsDefinitionPacing{
				Email: githubcomlightfldlightfieldgo.SequenceNewParamsDefinitionPacingEmail{
					DailyLimit: githubcomlightfldlightfieldgo.Int(1),
				},
				LinkedinInvite: githubcomlightfldlightfieldgo.SequenceNewParamsDefinitionPacingLinkedinInvite{
					DailyLimit: githubcomlightfldlightfieldgo.Int(1),
				},
				LinkedinMessage: githubcomlightfldlightfieldgo.SequenceNewParamsDefinitionPacingLinkedinMessage{
					DailyLimit: githubcomlightfldlightfieldgo.Int(1),
				},
			},
			Priority: "HIGH",
			SendWindow: githubcomlightfldlightfieldgo.SequenceNewParamsDefinitionSendWindow{
				DaysOfWeek: []int64{0},
				EndHour:    githubcomlightfldlightfieldgo.Int(1),
				StartHour:  githubcomlightfldlightfieldgo.Int(0),
				Timezone:   githubcomlightfldlightfieldgo.String("timezone"),
			},
		},
		Name:      "x",
		Reviewers: []string{"x"},
		Senders:   []string{"x"},
	})
	if err != nil {
		var apierr *githubcomlightfldlightfieldgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSequenceGet(t *testing.T) {
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
	_, err := client.Sequence.Get(context.TODO(), "id")
	if err != nil {
		var apierr *githubcomlightfldlightfieldgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSequenceListWithOptionalParams(t *testing.T) {
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
	_, err := client.Sequence.List(context.TODO(), githubcomlightfldlightfieldgo.SequenceListParams{
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

func TestSequenceActivateWithOptionalParams(t *testing.T) {
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
	_, err := client.Sequence.Activate(
		context.TODO(),
		"id",
		githubcomlightfldlightfieldgo.SequenceActivateParams{
			Body: githubcomlightfldlightfieldgo.SequenceActivateParamsBody{},
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

func TestSequenceDeactivateWithOptionalParams(t *testing.T) {
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
	_, err := client.Sequence.Deactivate(
		context.TODO(),
		"id",
		githubcomlightfldlightfieldgo.SequenceDeactivateParams{
			Body: githubcomlightfldlightfieldgo.SequenceDeactivateParamsBody{},
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

func TestSequenceDiscardDraft(t *testing.T) {
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
	_, err := client.Sequence.DiscardDraft(
		context.TODO(),
		"id",
		githubcomlightfldlightfieldgo.SequenceDiscardDraftParams{
			ExpectedVersion: "expectedVersion",
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

func TestSequenceListEnrollmentsWithOptionalParams(t *testing.T) {
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
	_, err := client.Sequence.ListEnrollments(
		context.TODO(),
		"id",
		githubcomlightfldlightfieldgo.SequenceListEnrollmentsParams{
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

func TestSequenceListMessagesWithOptionalParams(t *testing.T) {
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
	_, err := client.Sequence.ListMessages(
		context.TODO(),
		"id",
		githubcomlightfldlightfieldgo.SequenceListMessagesParams{
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

func TestSequencePublishDraft(t *testing.T) {
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
	_, err := client.Sequence.PublishDraft(
		context.TODO(),
		"id",
		githubcomlightfldlightfieldgo.SequencePublishDraftParams{
			ExpectedVersion: "expectedVersion",
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

func TestSequenceGetDraft(t *testing.T) {
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
	_, err := client.Sequence.GetDraft(context.TODO(), "id")
	if err != nil {
		var apierr *githubcomlightfldlightfieldgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSequenceGetEnrollment(t *testing.T) {
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
	_, err := client.Sequence.GetEnrollment(
		context.TODO(),
		"enrollmentId",
		githubcomlightfldlightfieldgo.SequenceGetEnrollmentParams{
			ID: "id",
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

func TestSequenceGetMessage(t *testing.T) {
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
	_, err := client.Sequence.GetMessage(
		context.TODO(),
		"messageId",
		githubcomlightfldlightfieldgo.SequenceGetMessageParams{
			ID: "id",
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

func TestSequenceGetPerformance(t *testing.T) {
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
	_, err := client.Sequence.GetPerformance(context.TODO(), "id")
	if err != nil {
		var apierr *githubcomlightfldlightfieldgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSequenceUpdateDraftWithOptionalParams(t *testing.T) {
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
	_, err := client.Sequence.UpdateDraft(
		context.TODO(),
		"id",
		githubcomlightfldlightfieldgo.SequenceUpdateDraftParams{
			Definition: githubcomlightfldlightfieldgo.SequenceUpdateDraftParamsDefinition{
				Description: githubcomlightfldlightfieldgo.String("x"),
				Messaging: githubcomlightfldlightfieldgo.SequenceUpdateDraftParamsDefinitionMessaging{
					CallToAction: githubcomlightfldlightfieldgo.String("x"),
					Persona:      githubcomlightfldlightfieldgo.String("x"),
					Trigger:      githubcomlightfldlightfieldgo.String("x"),
					ValueProp:    githubcomlightfldlightfieldgo.String("x"),
				},
				Pacing: githubcomlightfldlightfieldgo.SequenceUpdateDraftParamsDefinitionPacing{
					Email: githubcomlightfldlightfieldgo.SequenceUpdateDraftParamsDefinitionPacingEmail{
						DailyLimit: githubcomlightfldlightfieldgo.Int(1),
					},
					LinkedinInvite: githubcomlightfldlightfieldgo.SequenceUpdateDraftParamsDefinitionPacingLinkedinInvite{
						DailyLimit: githubcomlightfldlightfieldgo.Int(1),
					},
					LinkedinMessage: githubcomlightfldlightfieldgo.SequenceUpdateDraftParamsDefinitionPacingLinkedinMessage{
						DailyLimit: githubcomlightfldlightfieldgo.Int(1),
					},
				},
				Priority: "HIGH",
				SendWindow: githubcomlightfldlightfieldgo.SequenceUpdateDraftParamsDefinitionSendWindow{
					DaysOfWeek: []int64{0},
					EndHour:    githubcomlightfldlightfieldgo.Int(1),
					StartHour:  githubcomlightfldlightfieldgo.Int(0),
					Timezone:   githubcomlightfldlightfieldgo.String("timezone"),
				},
				Steps: []githubcomlightfldlightfieldgo.SequenceUpdateDraftParamsDefinitionStepUnion{{
					OfSequenceUpdateDraftsDefinitionStepObject: &githubcomlightfldlightfieldgo.SequenceUpdateDraftParamsDefinitionStepObject{
						DeliveryMode: "AUTO",
						Name:         "x",
						Prompt:       "x",
						Type:         "EMAIL",
						ID:           githubcomlightfldlightfieldgo.String("x"),
					},
				}},
			},
			ExpectedVersion: githubcomlightfldlightfieldgo.String("expectedVersion"),
			Name:            githubcomlightfldlightfieldgo.String("x"),
			Reviewers: githubcomlightfldlightfieldgo.SequenceUpdateDraftParamsReviewers{
				Add:    []string{"x"},
				Remove: []string{"x"},
			},
			Senders: githubcomlightfldlightfieldgo.SequenceUpdateDraftParamsSenders{
				Add:    []string{"x"},
				Remove: []string{"x"},
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

func TestSequenceUpdateEnrollmentsWithOptionalParams(t *testing.T) {
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
	_, err := client.Sequence.UpdateEnrollments(
		context.TODO(),
		"id",
		githubcomlightfldlightfieldgo.SequenceUpdateEnrollmentsParams{
			Add:    []string{"x"},
			Remove: []string{"x"},
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
