package issues

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestParseIssuesFromJson(t *testing.T) {
	fixtureResponse, _ := os.Open("../fixtures/issue_response.json")
	fixtureBody, _ := ioutil.ReadAll(fixtureResponse)
	defer fixtureResponse.Close()

	result := issuesFromJson(fixtureBody)

	if result[0].Title != "Add short twig runes to lettersToRunes" {
		t.Errorf("Expected %s, got %s", "Add short twig runes to lettersToRunes", result[0].Title)
	}

	if result[0].Link != "https://github.com/stscoundrel/younger-futhark/issues/165" {
		t.Errorf("Expected %s, got %s", "https://github.com/stscoundrel/younger-futhark/issues/165", result[0].Link)
	}

	if result[1].Title != "Check SWC as alternative to TypeScript default build" {
		t.Errorf("Expected %s, got %s", "Check SWC as alternative to TypeScript default build", result[1].Title)
	}

	if result[1].Link != "https://github.com/stscoundrel/old-norwegian-dictionary/issues/38" {
		t.Errorf("Expected %s, got %s", "https://github.com/stscoundrel/old-norwegian-dictionary/issues/38", result[1].Link)
	}
}
