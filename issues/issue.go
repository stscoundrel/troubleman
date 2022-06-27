package issues

import (
	"encoding/json"

	"github.com/stscoundrel/troubleman/internal/github"
)

type Issue struct {
	Title          string `json:"title"`
	Content        string `json:"body"`
	Link           string `json:"htmL_url"`
	RepositoryLink string `json:"repository_url"`
}

type IssueResponse struct {
	TotalCount int     `json:"total_count"`
	InComplete bool    `json:"incomplete_results"`
	Issues     []Issue `json:"items"`
}

func IssuesFromJson(jsonContent []byte) []Issue {
	var result IssueResponse

	json.Unmarshal(jsonContent, &result)

	return result.Issues
}

func GetIssues(username string) []Issue {
	rawIssues, _ := github.GetIssues("stscoundrel", github.BASE_ISSUE_SEARCH_URL)
	return IssuesFromJson(rawIssues)
}
