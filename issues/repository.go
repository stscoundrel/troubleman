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

func RepositoriesFromIssues(issues []Issue) []Repository {
	repositoriesToIssues := map[string][]Issue{}

	for _, issue := range issues {
		repositoriesToIssues[issue.RepositoryLink] = append(repositoriesToIssues[issue.RepositoryLink], issue)
	}

	repositories := []Repository{}

	for repositoryUrl, repositoryIssues := range repositoriesToIssues {
		repositories = append(repositories, Repository{
			Title:  getRepositoryName(repositoryUrl),
			Link:   repositoryUrl,
			Issues: repositoryIssues,
			Count:  len(repositoryIssues),
		})
	}

	return repositories
}
