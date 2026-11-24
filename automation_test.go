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

func TestAutomationNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Automation.New(context.TODO(), githubcomlightfldlightfieldgo.AutomationNewParams{
		Name: "x",
		Steps: []githubcomlightfldlightfieldgo.AutomationNewParamsStepUnion{{
			OfAgentStep: &githubcomlightfldlightfieldgo.AutomationNewParamsStepAgentStep{
				Name:   "x",
				Prompt: "x",
				Type:   "AGENT",
				Permissions: []githubcomlightfldlightfieldgo.AutomationNewParamsStepAgentStepPermissionUnion{{
					OfToolPermission: &githubcomlightfldlightfieldgo.AutomationNewParamsStepAgentStepPermissionToolPermission{
						Name:   "x",
						Type:   "TOOL",
						Reason: githubcomlightfldlightfieldgo.String("reason"),
					},
				}},
			},
		}},
		Triggers: []githubcomlightfldlightfieldgo.AutomationNewParamsTriggerUnion{{
			OfManualTrigger: &githubcomlightfldlightfieldgo.AutomationNewParamsTriggerManualTrigger{
				Type: "MANUAL",
			},
		}},
		Description: githubcomlightfldlightfieldgo.String("description"),
		Model:       githubcomlightfldlightfieldgo.AutomationNewParamsModelAuto,
		RunAs: githubcomlightfldlightfieldgo.AutomationNewParamsRunAsUnion{
			OfWorkspaceRunAs: &githubcomlightfldlightfieldgo.AutomationNewParamsRunAsWorkspaceRunAs{
				Type: "WORKSPACE",
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

func TestAutomationGet(t *testing.T) {
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
	_, err := client.Automation.Get(context.TODO(), "id")
	if err != nil {
		var apierr *githubcomlightfldlightfieldgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAutomationUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Automation.Update(
		context.TODO(),
		"id",
		githubcomlightfldlightfieldgo.AutomationUpdateParams{
			Description: githubcomlightfldlightfieldgo.String("description"),
			Model:       githubcomlightfldlightfieldgo.AutomationUpdateParamsModelAuto,
			Name:        githubcomlightfldlightfieldgo.String("x"),
			RunAs: githubcomlightfldlightfieldgo.AutomationUpdateParamsRunAsUnion{
				OfWorkspaceRunAs: &githubcomlightfldlightfieldgo.AutomationUpdateParamsRunAsWorkspaceRunAs{
					Type: "WORKSPACE",
				},
			},
			Steps: []githubcomlightfldlightfieldgo.AutomationUpdateParamsStepUnion{{
				OfAgentStep: &githubcomlightfldlightfieldgo.AutomationUpdateParamsStepAgentStep{
					Name:   "x",
					Prompt: "x",
					Type:   "AGENT",
					Permissions: []githubcomlightfldlightfieldgo.AutomationUpdateParamsStepAgentStepPermissionUnion{{
						OfToolPermission: &githubcomlightfldlightfieldgo.AutomationUpdateParamsStepAgentStepPermissionToolPermission{
							Name:   "x",
							Type:   "TOOL",
							Reason: githubcomlightfldlightfieldgo.String("reason"),
						},
					}},
				},
			}},
			Triggers: []githubcomlightfldlightfieldgo.AutomationUpdateParamsTriggerUnion{{
				OfManualTrigger: &githubcomlightfldlightfieldgo.AutomationUpdateParamsTriggerManualTrigger{
					Type: "MANUAL",
				},
			}},
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

func TestAutomationListWithOptionalParams(t *testing.T) {
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
	_, err := client.Automation.List(context.TODO(), githubcomlightfldlightfieldgo.AutomationListParams{
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

func TestAutomationDeleteWithOptionalParams(t *testing.T) {
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
	_, err := client.Automation.Delete(
		context.TODO(),
		"id",
		githubcomlightfldlightfieldgo.AutomationDeleteParams{
			Body: githubcomlightfldlightfieldgo.AutomationDeleteParamsBody{},
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

func TestAutomationPauseWithOptionalParams(t *testing.T) {
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
	_, err := client.Automation.Pause(
		context.TODO(),
		"id",
		githubcomlightfldlightfieldgo.AutomationPauseParams{
			Body: githubcomlightfldlightfieldgo.AutomationPauseParamsBody{},
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

func TestAutomationResumeWithOptionalParams(t *testing.T) {
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
	_, err := client.Automation.Resume(
		context.TODO(),
		"id",
		githubcomlightfldlightfieldgo.AutomationResumeParams{
			Body: githubcomlightfldlightfieldgo.AutomationResumeParamsBody{},
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
