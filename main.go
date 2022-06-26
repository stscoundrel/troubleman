package main

import (
	"fmt"

	"github.com/stscoundrel/troubleman/issues"
)

func main() {
	issueList := issues.GetIssues("stscoundrel")

	for _, issue := range issueList {
		fmt.Println(issue.Title)
		fmt.Println(issue.Content)
		fmt.Println(issue.Link)

		fmt.Println("###########################")
	}
}
