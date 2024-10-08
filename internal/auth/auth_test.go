package auth_test

import (
	"errors"
	"fmt"
	"net/http"
	"testing"

	"github.com/bootdotdev/learn-cicd-starter/internal/auth"
)

func TestGetApiKey(t *testing.T) {
	t.Parallel()

	testApiKey := "tw6ryDIjmDNB4jOE50231osrLaOB23CJI"

	headers := http.Header{}
	headers.Set("Authorization", fmt.Sprintf("ApiKey %s", testApiKey))

	gotApiKey, err := auth.GetAPIKey(headers)
	if err != nil {
		t.Fatalf(
			"FATAL: Unexpected error received from auth.GetAPIKey: got %v",
			err,
		)
	}

	if gotApiKey != testApiKey {
		t.Errorf(
			"FAILED: Unexpected API key received from auth.GetAPIKey: want %s, got %s",
			testApiKey,
			gotApiKey,
		)
	} else {
		t.Logf(
			"PASSED: Expected API key received from auth.GetAPIKey: got %s",
			gotApiKey,
		)
	}
}

func TestGetApiKeyNoAuthHeader(t *testing.T) {
	t.Parallel()

	_, err := auth.GetAPIKey(http.Header{})

	if err == nil {
		t.Error("FAILED: Did not receive an error from auth.GetAPIKey with missing Authorization header")
	} else {
		if !errors.Is(err, auth.ErrNoAuthHeaderIncluded) {
			t.Errorf(
				"FAILED: Unexpected error received from auth.GetAPIKey with missing Authorization header: got %v",
				err,
			)
		} else {
			t.Logf(
				"PASSED: Expected error received from auth.GetAPIKey with missing Authorization header: got %v",
				err,
			)
		}
	}
}
