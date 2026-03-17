// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package githubcomlightfldlightfieldgo_test

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"testing"
	"time"

	"github.com/Lightfld/lightfield-go"
	"github.com/Lightfld/lightfield-go/internal"
	"github.com/Lightfld/lightfield-go/option"
)

type closureTransport struct {
	fn func(req *http.Request) (*http.Response, error)
}

func (t *closureTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	return t.fn(req)
}

func TestUserAgentHeader(t *testing.T) {
	var userAgent string
	client := githubcomlightfldlightfieldgo.NewClient(
		option.WithAPIKey("My API Key"),
		option.WithHTTPClient(&http.Client{
			Transport: &closureTransport{
				fn: func(req *http.Request) (*http.Response, error) {
					userAgent = req.Header.Get("User-Agent")
					return &http.Response{
						StatusCode: http.StatusOK,
					}, nil
				},
			},
		}),
	)
	client.Account.New(context.Background(), githubcomlightfldlightfieldgo.AccountNewParams{
		Fields: githubcomlightfldlightfieldgo.AccountNewParamsFields{
			SystemName:      "Acme Corp",
			SystemWebsite:   []string{"https://acme.com"},
			SystemIndustry:  []string{"Technology", "SaaS"},
			SystemHeadcount: githubcomlightfldlightfieldgo.String("51-200"),
			SystemLinkedIn:  githubcomlightfldlightfieldgo.String("https://linkedin.com/company/acme"),
			SystemPrimaryAddress: map[string]string{
				"street":  "123 Market St",
				"city":    "San Francisco",
				"state":   "CA",
				"zip":     "94105",
				"country": "US",
			},
		},
		Relationships: githubcomlightfldlightfieldgo.AccountNewParamsRelationships{
			SystemOwner: githubcomlightfldlightfieldgo.AccountNewParamsRelationshipsSystemOwnerUnion{
				OfString: githubcomlightfldlightfieldgo.String("user_cm1abc123def456"),
			},
			SystemContact: githubcomlightfldlightfieldgo.AccountNewParamsRelationshipsSystemContactUnion{
				OfStringArray: []string{"contact_cm2ghi789jkl012", "contact_cm3mno345pqr678"},
			},
		},
	})
	if userAgent != fmt.Sprintf("Lightfield/Go %s", internal.PackageVersion) {
		t.Errorf("Expected User-Agent to be correct, but got: %#v", userAgent)
	}
}

func TestRetryAfter(t *testing.T) {
	retryCountHeaders := make([]string, 0)
	client := githubcomlightfldlightfieldgo.NewClient(
		option.WithAPIKey("My API Key"),
		option.WithHTTPClient(&http.Client{
			Transport: &closureTransport{
				fn: func(req *http.Request) (*http.Response, error) {
					retryCountHeaders = append(retryCountHeaders, req.Header.Get("X-Stainless-Retry-Count"))
					return &http.Response{
						StatusCode: http.StatusTooManyRequests,
						Header: http.Header{
							http.CanonicalHeaderKey("Retry-After"): []string{"0.1"},
						},
					}, nil
				},
			},
		}),
	)
	_, err := client.Account.New(context.Background(), githubcomlightfldlightfieldgo.AccountNewParams{
		Fields: githubcomlightfldlightfieldgo.AccountNewParamsFields{
			SystemName:      "Acme Corp",
			SystemWebsite:   []string{"https://acme.com"},
			SystemIndustry:  []string{"Technology", "SaaS"},
			SystemHeadcount: githubcomlightfldlightfieldgo.String("51-200"),
			SystemLinkedIn:  githubcomlightfldlightfieldgo.String("https://linkedin.com/company/acme"),
			SystemPrimaryAddress: map[string]string{
				"street":  "123 Market St",
				"city":    "San Francisco",
				"state":   "CA",
				"zip":     "94105",
				"country": "US",
			},
		},
		Relationships: githubcomlightfldlightfieldgo.AccountNewParamsRelationships{
			SystemOwner: githubcomlightfldlightfieldgo.AccountNewParamsRelationshipsSystemOwnerUnion{
				OfString: githubcomlightfldlightfieldgo.String("user_cm1abc123def456"),
			},
			SystemContact: githubcomlightfldlightfieldgo.AccountNewParamsRelationshipsSystemContactUnion{
				OfStringArray: []string{"contact_cm2ghi789jkl012", "contact_cm3mno345pqr678"},
			},
		},
	})
	if err == nil {
		t.Error("Expected there to be a cancel error")
	}

	attempts := len(retryCountHeaders)
	if attempts != 3 {
		t.Errorf("Expected %d attempts, got %d", 3, attempts)
	}

	expectedRetryCountHeaders := []string{"0", "1", "2"}
	if !reflect.DeepEqual(retryCountHeaders, expectedRetryCountHeaders) {
		t.Errorf("Expected %v retry count headers, got %v", expectedRetryCountHeaders, retryCountHeaders)
	}
}

func TestDeleteRetryCountHeader(t *testing.T) {
	retryCountHeaders := make([]string, 0)
	client := githubcomlightfldlightfieldgo.NewClient(
		option.WithAPIKey("My API Key"),
		option.WithHTTPClient(&http.Client{
			Transport: &closureTransport{
				fn: func(req *http.Request) (*http.Response, error) {
					retryCountHeaders = append(retryCountHeaders, req.Header.Get("X-Stainless-Retry-Count"))
					return &http.Response{
						StatusCode: http.StatusTooManyRequests,
						Header: http.Header{
							http.CanonicalHeaderKey("Retry-After"): []string{"0.1"},
						},
					}, nil
				},
			},
		}),
		option.WithHeaderDel("X-Stainless-Retry-Count"),
	)
	_, err := client.Account.New(context.Background(), githubcomlightfldlightfieldgo.AccountNewParams{
		Fields: githubcomlightfldlightfieldgo.AccountNewParamsFields{
			SystemName:      "Acme Corp",
			SystemWebsite:   []string{"https://acme.com"},
			SystemIndustry:  []string{"Technology", "SaaS"},
			SystemHeadcount: githubcomlightfldlightfieldgo.String("51-200"),
			SystemLinkedIn:  githubcomlightfldlightfieldgo.String("https://linkedin.com/company/acme"),
			SystemPrimaryAddress: map[string]string{
				"street":  "123 Market St",
				"city":    "San Francisco",
				"state":   "CA",
				"zip":     "94105",
				"country": "US",
			},
		},
		Relationships: githubcomlightfldlightfieldgo.AccountNewParamsRelationships{
			SystemOwner: githubcomlightfldlightfieldgo.AccountNewParamsRelationshipsSystemOwnerUnion{
				OfString: githubcomlightfldlightfieldgo.String("user_cm1abc123def456"),
			},
			SystemContact: githubcomlightfldlightfieldgo.AccountNewParamsRelationshipsSystemContactUnion{
				OfStringArray: []string{"contact_cm2ghi789jkl012", "contact_cm3mno345pqr678"},
			},
		},
	})
	if err == nil {
		t.Error("Expected there to be a cancel error")
	}

	expectedRetryCountHeaders := []string{"", "", ""}
	if !reflect.DeepEqual(retryCountHeaders, expectedRetryCountHeaders) {
		t.Errorf("Expected %v retry count headers, got %v", expectedRetryCountHeaders, retryCountHeaders)
	}
}

func TestOverwriteRetryCountHeader(t *testing.T) {
	retryCountHeaders := make([]string, 0)
	client := githubcomlightfldlightfieldgo.NewClient(
		option.WithAPIKey("My API Key"),
		option.WithHTTPClient(&http.Client{
			Transport: &closureTransport{
				fn: func(req *http.Request) (*http.Response, error) {
					retryCountHeaders = append(retryCountHeaders, req.Header.Get("X-Stainless-Retry-Count"))
					return &http.Response{
						StatusCode: http.StatusTooManyRequests,
						Header: http.Header{
							http.CanonicalHeaderKey("Retry-After"): []string{"0.1"},
						},
					}, nil
				},
			},
		}),
		option.WithHeader("X-Stainless-Retry-Count", "42"),
	)
	_, err := client.Account.New(context.Background(), githubcomlightfldlightfieldgo.AccountNewParams{
		Fields: githubcomlightfldlightfieldgo.AccountNewParamsFields{
			SystemName:      "Acme Corp",
			SystemWebsite:   []string{"https://acme.com"},
			SystemIndustry:  []string{"Technology", "SaaS"},
			SystemHeadcount: githubcomlightfldlightfieldgo.String("51-200"),
			SystemLinkedIn:  githubcomlightfldlightfieldgo.String("https://linkedin.com/company/acme"),
			SystemPrimaryAddress: map[string]string{
				"street":  "123 Market St",
				"city":    "San Francisco",
				"state":   "CA",
				"zip":     "94105",
				"country": "US",
			},
		},
		Relationships: githubcomlightfldlightfieldgo.AccountNewParamsRelationships{
			SystemOwner: githubcomlightfldlightfieldgo.AccountNewParamsRelationshipsSystemOwnerUnion{
				OfString: githubcomlightfldlightfieldgo.String("user_cm1abc123def456"),
			},
			SystemContact: githubcomlightfldlightfieldgo.AccountNewParamsRelationshipsSystemContactUnion{
				OfStringArray: []string{"contact_cm2ghi789jkl012", "contact_cm3mno345pqr678"},
			},
		},
	})
	if err == nil {
		t.Error("Expected there to be a cancel error")
	}

	expectedRetryCountHeaders := []string{"42", "42", "42"}
	if !reflect.DeepEqual(retryCountHeaders, expectedRetryCountHeaders) {
		t.Errorf("Expected %v retry count headers, got %v", expectedRetryCountHeaders, retryCountHeaders)
	}
}

func TestRetryAfterMs(t *testing.T) {
	attempts := 0
	client := githubcomlightfldlightfieldgo.NewClient(
		option.WithAPIKey("My API Key"),
		option.WithHTTPClient(&http.Client{
			Transport: &closureTransport{
				fn: func(req *http.Request) (*http.Response, error) {
					attempts++
					return &http.Response{
						StatusCode: http.StatusTooManyRequests,
						Header: http.Header{
							http.CanonicalHeaderKey("Retry-After-Ms"): []string{"100"},
						},
					}, nil
				},
			},
		}),
	)
	_, err := client.Account.New(context.Background(), githubcomlightfldlightfieldgo.AccountNewParams{
		Fields: githubcomlightfldlightfieldgo.AccountNewParamsFields{
			SystemName:      "Acme Corp",
			SystemWebsite:   []string{"https://acme.com"},
			SystemIndustry:  []string{"Technology", "SaaS"},
			SystemHeadcount: githubcomlightfldlightfieldgo.String("51-200"),
			SystemLinkedIn:  githubcomlightfldlightfieldgo.String("https://linkedin.com/company/acme"),
			SystemPrimaryAddress: map[string]string{
				"street":  "123 Market St",
				"city":    "San Francisco",
				"state":   "CA",
				"zip":     "94105",
				"country": "US",
			},
		},
		Relationships: githubcomlightfldlightfieldgo.AccountNewParamsRelationships{
			SystemOwner: githubcomlightfldlightfieldgo.AccountNewParamsRelationshipsSystemOwnerUnion{
				OfString: githubcomlightfldlightfieldgo.String("user_cm1abc123def456"),
			},
			SystemContact: githubcomlightfldlightfieldgo.AccountNewParamsRelationshipsSystemContactUnion{
				OfStringArray: []string{"contact_cm2ghi789jkl012", "contact_cm3mno345pqr678"},
			},
		},
	})
	if err == nil {
		t.Error("Expected there to be a cancel error")
	}
	if want := 3; attempts != want {
		t.Errorf("Expected %d attempts, got %d", want, attempts)
	}
}

func TestContextCancel(t *testing.T) {
	client := githubcomlightfldlightfieldgo.NewClient(
		option.WithAPIKey("My API Key"),
		option.WithHTTPClient(&http.Client{
			Transport: &closureTransport{
				fn: func(req *http.Request) (*http.Response, error) {
					<-req.Context().Done()
					return nil, req.Context().Err()
				},
			},
		}),
	)
	cancelCtx, cancel := context.WithCancel(context.Background())
	cancel()
	_, err := client.Account.New(cancelCtx, githubcomlightfldlightfieldgo.AccountNewParams{
		Fields: githubcomlightfldlightfieldgo.AccountNewParamsFields{
			SystemName:      "Acme Corp",
			SystemWebsite:   []string{"https://acme.com"},
			SystemIndustry:  []string{"Technology", "SaaS"},
			SystemHeadcount: githubcomlightfldlightfieldgo.String("51-200"),
			SystemLinkedIn:  githubcomlightfldlightfieldgo.String("https://linkedin.com/company/acme"),
			SystemPrimaryAddress: map[string]string{
				"street":  "123 Market St",
				"city":    "San Francisco",
				"state":   "CA",
				"zip":     "94105",
				"country": "US",
			},
		},
		Relationships: githubcomlightfldlightfieldgo.AccountNewParamsRelationships{
			SystemOwner: githubcomlightfldlightfieldgo.AccountNewParamsRelationshipsSystemOwnerUnion{
				OfString: githubcomlightfldlightfieldgo.String("user_cm1abc123def456"),
			},
			SystemContact: githubcomlightfldlightfieldgo.AccountNewParamsRelationshipsSystemContactUnion{
				OfStringArray: []string{"contact_cm2ghi789jkl012", "contact_cm3mno345pqr678"},
			},
		},
	})
	if err == nil {
		t.Error("Expected there to be a cancel error")
	}
}

func TestContextCancelDelay(t *testing.T) {
	client := githubcomlightfldlightfieldgo.NewClient(
		option.WithAPIKey("My API Key"),
		option.WithHTTPClient(&http.Client{
			Transport: &closureTransport{
				fn: func(req *http.Request) (*http.Response, error) {
					<-req.Context().Done()
					return nil, req.Context().Err()
				},
			},
		}),
	)
	cancelCtx, cancel := context.WithTimeout(context.Background(), 2*time.Millisecond)
	defer cancel()
	_, err := client.Account.New(cancelCtx, githubcomlightfldlightfieldgo.AccountNewParams{
		Fields: githubcomlightfldlightfieldgo.AccountNewParamsFields{
			SystemName:      "Acme Corp",
			SystemWebsite:   []string{"https://acme.com"},
			SystemIndustry:  []string{"Technology", "SaaS"},
			SystemHeadcount: githubcomlightfldlightfieldgo.String("51-200"),
			SystemLinkedIn:  githubcomlightfldlightfieldgo.String("https://linkedin.com/company/acme"),
			SystemPrimaryAddress: map[string]string{
				"street":  "123 Market St",
				"city":    "San Francisco",
				"state":   "CA",
				"zip":     "94105",
				"country": "US",
			},
		},
		Relationships: githubcomlightfldlightfieldgo.AccountNewParamsRelationships{
			SystemOwner: githubcomlightfldlightfieldgo.AccountNewParamsRelationshipsSystemOwnerUnion{
				OfString: githubcomlightfldlightfieldgo.String("user_cm1abc123def456"),
			},
			SystemContact: githubcomlightfldlightfieldgo.AccountNewParamsRelationshipsSystemContactUnion{
				OfStringArray: []string{"contact_cm2ghi789jkl012", "contact_cm3mno345pqr678"},
			},
		},
	})
	if err == nil {
		t.Error("expected there to be a cancel error")
	}
}

func TestContextDeadline(t *testing.T) {
	testTimeout := time.After(3 * time.Second)
	testDone := make(chan struct{})

	deadline := time.Now().Add(100 * time.Millisecond)
	deadlineCtx, cancel := context.WithDeadline(context.Background(), deadline)
	defer cancel()

	go func() {
		client := githubcomlightfldlightfieldgo.NewClient(
			option.WithAPIKey("My API Key"),
			option.WithHTTPClient(&http.Client{
				Transport: &closureTransport{
					fn: func(req *http.Request) (*http.Response, error) {
						<-req.Context().Done()
						return nil, req.Context().Err()
					},
				},
			}),
		)
		_, err := client.Account.New(deadlineCtx, githubcomlightfldlightfieldgo.AccountNewParams{
			Fields: githubcomlightfldlightfieldgo.AccountNewParamsFields{
				SystemName:      "Acme Corp",
				SystemWebsite:   []string{"https://acme.com"},
				SystemIndustry:  []string{"Technology", "SaaS"},
				SystemHeadcount: githubcomlightfldlightfieldgo.String("51-200"),
				SystemLinkedIn:  githubcomlightfldlightfieldgo.String("https://linkedin.com/company/acme"),
				SystemPrimaryAddress: map[string]string{
					"street":  "123 Market St",
					"city":    "San Francisco",
					"state":   "CA",
					"zip":     "94105",
					"country": "US",
				},
			},
			Relationships: githubcomlightfldlightfieldgo.AccountNewParamsRelationships{
				SystemOwner: githubcomlightfldlightfieldgo.AccountNewParamsRelationshipsSystemOwnerUnion{
					OfString: githubcomlightfldlightfieldgo.String("user_cm1abc123def456"),
				},
				SystemContact: githubcomlightfldlightfieldgo.AccountNewParamsRelationshipsSystemContactUnion{
					OfStringArray: []string{"contact_cm2ghi789jkl012", "contact_cm3mno345pqr678"},
				},
			},
		})
		if err == nil {
			t.Error("expected there to be a deadline error")
		}
		close(testDone)
	}()

	select {
	case <-testTimeout:
		t.Fatal("client didn't finish in time")
	case <-testDone:
		if diff := time.Since(deadline); diff < -30*time.Millisecond || 30*time.Millisecond < diff {
			t.Fatalf("client did not return within 30ms of context deadline, got %s", diff)
		}
	}
}
