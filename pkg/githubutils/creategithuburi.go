package githubutils

import (
	"fmt"
	"net/url"
)

// CreateGithubURI will construct a valid uri with the given information
func CreateGithubURI(account string, repo string, commit string) string {
	return "https://github.com/" + account + "/" + repo + "/commit/" + commit
}

// IsGithubURI checks to see if the given uri is a valid github uri and returns
// true if so
func IsGithubURI(uri string) (bool, error) {
	u, err := url.Parse(uri)
	if err != nil {
		err = fmt.Errorf("Error parsing uri: '%v'", err)
		return false, err
	}
	b := u.Hostname() == "github.com"
	return b, nil
}
