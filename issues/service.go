package issues

import "github.com/stscoundrel/troubleman/internal/github"

func GetIssues(username string) ([]Issue, error) {
	rawIssues, error := github.GetIssues(username, github.BASE_ISSUE_SEARCH_URL)

	if error != nil {
		return []Issue{}, error
	}

	return issuesFromJson(rawIssues), nil
}

func GetRepositories(username string) ([]Repository, error) {
	issues, error := GetIssues(username)

	if error != nil {
		return []Repository{}, error
	}

	return repositoriesFromIssues(issues), nil
}
