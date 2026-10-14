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

func TestEmailGet(t *testing.T) {
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
	_, err := client.Email.Get(context.TODO(), "id")
	if err != nil {
		var apierr *githubcomlightfldlightfieldgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestEmailListWithOptionalParams(t *testing.T) {
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
	_, err := client.Email.List(context.TODO(), githubcomlightfldlightfieldgo.EmailListParams{
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

func TestEmailDeleteWithOptionalParams(t *testing.T) {
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
	_, err := client.Email.Delete(
		context.TODO(),
		"id",
		githubcomlightfldlightfieldgo.EmailDeleteParams{
			Body: githubcomlightfldlightfieldgo.EmailDeleteParamsBody{},
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

func TestEmailDraftWithOptionalParams(t *testing.T) {
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
	_, err := client.Email.Draft(context.TODO(), githubcomlightfldlightfieldgo.EmailDraftParams{
		From:        "sales@acme.com",
		Attachments: []string{"file_01abc2def3ghi4jkl5mno6pqr"},
		Bcc:         []string{"lead@example.com"},
		Cc:          []string{"lead@example.com"},
		MessageBody: githubcomlightfldlightfieldgo.EmailDraftParamsMessageBody{
			Content:     "<p>Hi there,</p><p>Following up on our chat earlier this week.</p>",
			ContentType: "HTML",
		},
		Subject: githubcomlightfldlightfieldgo.String("subject"),
		To:      []string{"lead@example.com"},
	})
	if err != nil {
		var apierr *githubcomlightfldlightfieldgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestEmailLogWithOptionalParams(t *testing.T) {
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
	_, err := client.Email.Log(context.TODO(), githubcomlightfldlightfieldgo.EmailLogParams{
		Fields: githubcomlightfldlightfieldgo.EmailLogParamsFields{
			From:   "$from",
			SentAt: "$sentAt",
			Bcc:    []string{"string"},
			Body: githubcomlightfldlightfieldgo.EmailLogParamsFieldsBody{
				HTML: githubcomlightfldlightfieldgo.String("html"),
				Text: githubcomlightfldlightfieldgo.String("text"),
			},
			Cc:             []string{"string"},
			PrivacySetting: "FULL",
			Subject:        githubcomlightfldlightfieldgo.String("$subject"),
			ThreadID:       githubcomlightfldlightfieldgo.String("x"),
			To:             []string{"string"},
		},
		AutoCreateRecords: githubcomlightfldlightfieldgo.Bool(true),
		Headers: map[string]string{
			"foo": "string",
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

func TestEmailRestoreWithOptionalParams(t *testing.T) {
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
	_, err := client.Email.Restore(
		context.TODO(),
		"id",
		githubcomlightfldlightfieldgo.EmailRestoreParams{
			Body: githubcomlightfldlightfieldgo.EmailRestoreParamsBody{},
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

func TestEmailSendWithOptionalParams(t *testing.T) {
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
	_, err := client.Email.Send(context.TODO(), githubcomlightfldlightfieldgo.EmailSendParams{
		From: "sales@acme.com",
		MessageBody: githubcomlightfldlightfieldgo.EmailSendParamsMessageBody{
			Content:     "<p>Hi there,</p><p>Following up on our chat earlier this week.</p>",
			ContentType: "HTML",
		},
		Subject:     "Following up on our chat",
		To:          []string{"lead@example.com"},
		Attachments: []string{"file_01abc2def3ghi4jkl5mno6pqr"},
		Bcc:         []string{"lead@example.com"},
		Cc:          []string{"lead@example.com"},
	})
	if err != nil {
		var apierr *githubcomlightfldlightfieldgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
