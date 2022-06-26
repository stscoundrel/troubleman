package github

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const BASE_ISSUE_SEARCH_URL = "https://api.github.com/search/issues"

func get(url string) ([]byte, error) {
	resp, responseErr := http.Get(url)
	if responseErr != nil {
		fmt.Println(responseErr)
		return []byte{}, responseErr
	}
	defer resp.Body.Close()
	body, parseErr := ioutil.ReadAll(resp.Body)

	if parseErr != nil {
		fmt.Println(parseErr)
		return []byte{}, parseErr
	}

	return body, nil
}

func GetIssues(username string, baseApiUrl string) ([]byte, error) {
	api_url := baseApiUrl + "?q=user:" + username + "+is:issue+state:open&per_page=100&page=1"
	return get(api_url)
}
