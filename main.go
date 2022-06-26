package main

import (
	"fmt"

	"github.com/stscoundrel/troubleman/github"
	"github.com/stscoundrel/troubleman/issues"
)

func GetIssues(username string) []issues.Issue {
	rawIssues, _ := github.GetIssues("stscoundrel", github.BASE_ISSUE_SEARCH_URL)
	return issues.FromJson(rawIssues)
}

func main() {
	issueList := GetIssues("stscoundrel")

	for _, issue := range issueList {
		fmt.Println(issue.Title)
		fmt.Println(issue.Content)
		fmt.Println(issue.Link)

		fmt.Println("###########################")
	}
}
