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
		Body: githubcomlightfldlightfieldgo.EmailSendParamsBody{
			Content:     "x",
			ContentType: "HTML",
		},
		From:        `"S?oC"g*W"5"@m-0-9.V9.w4-o2-l7--.TJdq6.1k.H8n-SjA-.1.U3k7.-F86WnfNI.-R6O-N68g-4-.AmqyytAVIw`,
		Subject:     "x",
		To:          []string{`"S?oC"g*W"5"@m-0-9.V9.w4-o2-l7--.TJdq6.1k.H8n-SjA-.1.U3k7.-F86WnfNI.-R6O-N68g-4-.AmqyytAVIw`},
		Attachments: []string{"string"},
		Bcc:         []string{`"S?oC"g*W"5"@m-0-9.V9.w4-o2-l7--.TJdq6.1k.H8n-SjA-.1.U3k7.-F86WnfNI.-R6O-N68g-4-.AmqyytAVIw`},
		Cc:          []string{`"S?oC"g*W"5"@m-0-9.V9.w4-o2-l7--.TJdq6.1k.H8n-SjA-.1.U3k7.-F86WnfNI.-R6O-N68g-4-.AmqyytAVIw`},
	})
	if err != nil {
		var apierr *githubcomlightfldlightfieldgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
