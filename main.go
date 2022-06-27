package main

import (
	"fmt"

	"github.com/stscoundrel/troubleman/issues"
)

func main() {
	issueList, _ := issues.GetIssues("stscoundrel")

	for _, issue := range issueList {
		fmt.Println(issue.Title)
		fmt.Println(issue.Content)
		fmt.Println(issue.Link)

		fmt.Println("###########################")
	}

	repositories, _ := issues.GetRepositories("stscoundrel")

	for _, repository := range repositories {
		fmt.Println(repository.Title)
		fmt.Println(repository.Link)
		fmt.Println(repository.Count)
		fmt.Println(repository.Issues)

		fmt.Println("###########################")
	}
}
