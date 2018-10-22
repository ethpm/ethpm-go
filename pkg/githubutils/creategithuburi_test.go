package githubutils

import (
	"testing"
)

func TestIsGithubURI(t *testing.T) {
	got, _ := IsGithubURI("https://modularbanking.com/check/us/out")
	want := false
	if got != want {
		t.Fatalf("Got '%v', expected '%v'", got, want)
	}
}
