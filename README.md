# Troubleman

List issues in Github by username. For the times when you can't recall all the troubles you're in.

## Install

`go get -d github.com/stscoundrel/troubleman`

## Usage

Troubleman exposes two functions: one for listing issues by user, other for listing same data grouped by repositories.

```go
package main

import (
    "fmt"

    "github.com/stscoundrel/troubleman/issues"
)

func main() {
    // Get list of all issues
    issuesResult, issuesError := issues.GetIssues("stscoundrel")
    
    fmt.Println(issuesResult[0])
    // {
    // 	Title:          "Test issue 1",
    // 	Link:           "https://github.com/stscoundrel/test-repo/issues/1",
    // 	Content:        "Lorem ipsum dolor sit issue",
    // 	RepositoryLink: "https://github.com/stscoundrel/test-repo",
    // },

    // Get list of issues of issues grouped by repositories.
    repositoriesResult, repositoriesError := issues.GetRepositories("stscoundrel")
    
    fmt.Println(repositoriesResult[0])
    // {
    // 	Title:          "stscoundrel/test-repo",
    // 	Link:           "https://github.com/stscoundrel/test-repo",
    // 	Count:          666,
    // 	Issues:         [] // list of issues, per upper structure
    // },
}
```

## What's in the name?

"Troublemaker" was the first thing to come to mind when thinking of "something that looks for issues". Altered it a bit to refer to [a song](https://www.youtube.com/watch?v=3OC2aPCuzjo)
