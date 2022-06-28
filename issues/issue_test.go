package issues

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseIssuesFromJson(t *testing.T) {
	fixtureResponse, _ := os.Open("../fixtures/issue_response.json")
	fixtureBody, _ := ioutil.ReadAll(fixtureResponse)
	defer fixtureResponse.Close()

	result := issuesFromJson(fixtureBody)

	assert.Equal(t, result[0].Title, "Ignores: allow custom ignores")
	assert.Equal(t, result[0].Link, "https://github.com/stscoundrel/alliser/issues/9")

	assert.Equal(t, result[1].Title, "Add sessionStorage variant")
	assert.Equal(t, result[1].Link, "https://github.com/stscoundrel/emma/issues/10")

	assert.Equal(t, result[5].Title, "Add integration tests to keyboard pages")
	assert.Equal(t, result[5].Link, "https://github.com/stscoundrel/runes/issues/37")
}
