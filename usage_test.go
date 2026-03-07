// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package lightfield_test

import (
	"context"
	"os"
	"testing"

	"github.com/stainless-sdks/lightfield-go"
	"github.com/stainless-sdks/lightfield-go/internal/testutil"
	"github.com/stainless-sdks/lightfield-go/option"
)

func TestUsage(t *testing.T) {
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
	object, err := client.Object.New(
		context.TODO(),
		lightfield.ObjectNewParamsEntityTypeOpportunities,
		lightfield.ObjectNewParams{
			Fields: map[string]lightfield.ObjectNewParamsFieldUnion{
				"name": {
					OfString: lightfield.String("Big Deal"),
				},
			},
		},
	)
	if err != nil {
		t.Fatalf("err should be nil: %s", err.Error())
	}
	t.Logf("%+v\n", object)
}
