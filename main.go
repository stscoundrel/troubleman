package main

import (
	"github.com/stscoundrel/troubleman/github"
	"github.com/stscoundrel/troubleman/issues"
)

func GetIssues(username string) []issues.Issue {
	rawIssues, _ := github.GetIssues("stscoundrel", github.BASE_ISSUE_SEARCH_URL)
	return issues.FromJson(rawIssues)
}
