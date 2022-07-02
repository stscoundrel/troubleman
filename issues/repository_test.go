package issues

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRepositoriesFromIssues(t *testing.T) {
	issues := []Issue{
		{
			Title:          "Test issue 2",
			Link:           "https://github.com/stscoundrel/cde-test-repo/issues/2",
			Content:        "",
			RepositoryLink: "https://api.github.com/stscoundrel/cde-test-repo",
		},
		{
			Title:          "Test issue 1",
			Link:           "https://github.com/stscoundrel/abc-test-repo/issues/1",
			Content:        "",
			RepositoryLink: "https://api.github.com/stscoundrel/abc-test-repo",
		},
		{
			Title:          "Test issue 3",
			Link:           "https://github.com/stscoundrel/abc-test-repo/issues/3",
			Content:        "",
			RepositoryLink: "https://api.github.com/stscoundrel/abc-test-repo",
		},
		{
			Title:          "Test issue 4",
			Link:           "https://github.com/stscoundrel/cde-test-repo/issues/4",
			Content:        "",
			RepositoryLink: "https://api.github.com/stscoundrel/cde-test-repo",
		},
		{
			Title:          "Test issue 5",
			Link:           "https://github.com/stscoundrel/abc-test-repo/issues/5",
			Content:        "",
			RepositoryLink: "https://api.github.com/stscoundrel/abc-test-repo",
		},
	}

	result := repositoriesFromIssues(issues)

	assert.Equal(t, len(result), 2)

	assert.Equal(t, result[0].Title, "stscoundrel/abc-test-repo")
	assert.Equal(t, result[0].Count, 3)
	assert.Equal(t, result[0].Link, "https://github.com/stscoundrel/abc-test-repo")
	assert.Equal(t, result[0].Issues, []Issue{
		issues[1],
		issues[2],
		issues[4],
	})

	assert.Equal(t, result[1].Title, "stscoundrel/cde-test-repo")
	assert.Equal(t, result[1].Count, 2)
	assert.Equal(t, result[1].Link, "https://github.com/stscoundrel/cde-test-repo")
	assert.Equal(t, result[1].Issues, []Issue{
		issues[0],
		issues[3],
	})
}
