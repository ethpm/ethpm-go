package githubutils

import (
	"fmt"
	"net/url"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// GetETHPMManifest takes a complete github commit uri, and clones it into the
// given parent directory. If parentdir is an empty string, it will clone into
// the current working directory, then go into this directory and retrieve
// the ethpm manifest by looking for an 'ethpm.json' file.
func GetETHPMManifest(commithttpsurl string, parentdir string) (jsonstring string, err error) {
	uri, err := url.Parse(commithttpsurl)
	uripart := strings.Split(uri.Path, "/")
	project := uripart[len(uripart)-3]
	commit := uripart[len(uripart)-1]
	newuri := strings.Join(uripart[:len(uripart)-2], "/")
	uri.Path = newuri

	if parentdir == "" {
		parentdir, err = os.Getwd()
		if err != nil {
			err = fmt.Errorf("Could not access working directory: '%v'", err)
			return
		}
	}
	repodir := filepath.FromSlash(parentdir + "/" + project)
	gitlocation, err := exec.LookPath("git")
	if err != nil {
		err = fmt.Errorf("Error getting git bin location: '%v'", err)
		return
	}
	gitCmd := exec.Command(gitlocation, "clone", uri.String(), repodir)
	_, err = gitCmd.Output()
	if err != nil {
		err = fmt.Errorf("Error calling git command: '%v'", err)
	}

	gitCmd = exec.Command(gitlocation, "-C", repodir, "checkout", commit)
	_, err = gitCmd.Output()
	if err != nil {
		err = fmt.Errorf("Error calling git command: '%v'", err)
	}
	ethpmfile := filepath.Join(repodir, "ethpm.json")
	file, err := os.Open(ethpmfile)
	if err != nil {
		err = fmt.Errorf("Could not open ethpm json file: '%v'", err)
		return
	}
	info, _ := file.Stat()
	manifestbytes := make([]byte, info.Size())
	_, err = file.Read(manifestbytes)
	if err != nil {
		err = fmt.Errorf("Could not read ethpmfile: '%v'", err)
		return
	}
	jsonstring = string(manifestbytes)
	return
}
