package github

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestGetIssues(t *testing.T) {
	fixtureResponse, _ := os.Open("../../fixtures/issue_response.json")
	fixtureBody, _ := ioutil.ReadAll(fixtureResponse)
	defer fixtureResponse.Close()

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, request *http.Request) {
		queryParams := request.URL.Query()

		if queryParams["q"][0] != "user:stscoundrel is:issue state:open" {
			t.Errorf("Expected to per_page param paged to be user:stscoundrel is:issue state:open, got: " + queryParams["q"][0])
		}

		if queryParams["page"][0] != "1" {
			t.Errorf("Expected to query param paged to be 1, got: " + queryParams["page"][0])
		}

		if queryParams["per_page"][0] != "100" {
			t.Errorf("Expected to per_page param paged to be 100, got: " + queryParams["per_page"][0])
		}

		w.WriteHeader(http.StatusOK)
		w.Write(fixtureBody)
	}))
	defer server.Close()

	result, _ := GetIssues("stscoundrel", server.URL)
	if string(result) != string(fixtureBody) {
		t.Errorf("Expected %s, got %s", string(result), string(fixtureBody))
	}
}
