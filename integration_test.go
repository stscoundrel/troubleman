package tests

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stscoundrel/troubleman/issues"
)

func TestIssuesFromGithub(t *testing.T) {
	result, _ := issues.GetIssues("stscoundrel")

	// Should always have at least test issue.
	assert.True(t, len(result) > 0)

	found_expected_issue := false

	for _, issue := range result {
		if issue.Title == "Sample issue for integration tests" {
			found_expected_issue = true

			assert.Equal(t, issue.Content, "Exists only for integration test to find this")
			assert.Equal(t, issue.Link, "https://github.com/stscoundrel/troubleman/issues/6")
		}
	}

	assert.True(t, found_expected_issue)
}

func TestRepositoriesFromGithub(t *testing.T) {
	result, _ := issues.GetRepositories("stscoundrel")

	// Should always have at least repo with test issue.
	assert.True(t, len(result) > 0)

	found_expected_repo := false

	for _, repository := range result {
		if repository.Title == "stscoundrel/troubleman" {
			found_expected_repo = true

			assert.True(t, repository.Count > 0)
			assert.Equal(t, repository.Link, "https://github.com/stscoundrel/troubleman")
		}
	}

	assert.True(t, found_expected_repo)
}
