package issues

import "strings"

type Repository struct {
	Title  string
	Link   string
	Count  int
	Issues []Issue
}

func getRepositoryName(url string) string {
	segments := strings.Split(url, "/")

	return segments[len(segments)-2] + "/" + segments[len(segments)-1]
}

func getRepositoryUrl(repositoryName string) string {
	// API response contains only "api" version of repo url.
	// Parse proper user url based on repo name.
	return "https://github.com/" + repositoryName
}

func repositoriesFromIssues(issues []Issue) []Repository {
	repositoriesToIssues := map[string][]Issue{}

	for _, issue := range issues {
		repositoriesToIssues[issue.RepositoryLink] = append(repositoriesToIssues[issue.RepositoryLink], issue)
	}

	repositories := []Repository{}

	for repositoryUrl, repositoryIssues := range repositoriesToIssues {
		name := getRepositoryName(repositoryUrl)
		repositories = append(repositories, Repository{
			Title:  name,
			Link:   getRepositoryUrl(name),
			Issues: repositoryIssues,
			Count:  len(repositoryIssues),
		})
	}

	return repositories
}
