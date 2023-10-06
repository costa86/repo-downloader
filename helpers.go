package main

import (
	"encoding/json"
	"fmt"
	"os"
	"syscall"

	"github.com/go-git/go-git/v5/plumbing/transport/http"
	"gopkg.in/src-d/go-git.v4"

	"golang.org/x/term"
)

var print = fmt.Println

type Repo struct {
	Folder string `json:"folder"`
	Url    string `json:"url"`
}

type UserParameters struct {
	Username, Password, RepoFile string
}

/*
Get username, password

Returning ("", "") means that no authentication will be used to clone the repositories
*/
func getCredentials(username, password string) (string, string) {
	empty := ""

	userIsProvided := (username != empty)
	passwordIsProvided := (password != empty)

	if !userIsProvided && !passwordIsProvided {
		return empty, empty
	}

	if userIsProvided && passwordIsProvided {
		return username, password
	}

	if passwordIsProvided && !userIsProvided {
		print("If a password is provided, a username must be provided")
		os.Exit(1)
	}

	fmt.Println("Enter password:")
	pwd, err := term.ReadPassword(int(syscall.Stdin))
	if err != nil {
		print(err)
		os.Exit(1)
	}

	return username, string(pwd)
}

// Get list of repositories do be cloned
func getRepos(repoFile string) []Repo {
	var repos []Repo
	content, err := os.Open(repoFile)
	if err != nil {
		print(err)
		os.Exit(1)
	}
	json.NewDecoder(content).Decode(&repos)
	return repos
}

func cleanUp(err error, folder string) {
	if err != nil {
		print(err)
		os.Remove(folder)
		os.Exit(1)
	}
}

/*
Download remote repository:

  - username and password are set: clones repo with credentials
  - username and password are "": clones repo without credentials
*/
func downloadRepo(repo Repo, username string, password string) {
	empty := ""
	authRequired := (username != empty) && (password != empty)

	if authRequired {
		auth := &http.BasicAuth{
			Username: username,
			Password: password,
		}
		_, err := git.PlainClone(repo.Folder, false, &git.CloneOptions{URL: repo.Url, Progress: os.Stdout, Auth: auth})
		cleanUp(err, repo.Folder)
		return
	}
	_, err := git.PlainClone(repo.Folder, false, &git.CloneOptions{URL: repo.Url, Progress: os.Stdout})
	cleanUp(err, repo.Folder)

}
