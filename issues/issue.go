package issues

import (
	"encoding/json"
	"sort"
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

func issuesFromJson(jsonContent []byte) []Issue {
	var result IssueResponse

	json.Unmarshal(jsonContent, &result)

	issues := sortIssues(result.Issues)

	return issues
}

func sortIssues(issues []Issue) []Issue {
	sort.Slice(issues, func(i, j int) bool {
		return issues[i].Link < issues[j].Link
	})

	return issues
}
