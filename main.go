package main

import (
	"encoding/json"
	"fmt"
	"os"
	"sync"

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

func getParameters(args []string) UserParameters {
	var userParameters UserParameters

	switch len(args) {
	case 4:
		userParameters.RepoFile = args[1]
		userParameters.Username = args[2]
		userParameters.Password = args[3]

	case 3:
		fmt.Println("Enter password:")
		pwd, err := term.ReadPassword(int(syscall.Stdin))
		if err != nil {
			panic(err)
		}
		userParameters.RepoFile = args[1]
		userParameters.Username = args[2]
		userParameters.Password = string(pwd)
	case 2:
		userParameters.RepoFile = args[1]
		userParameters.Username = ""
		userParameters.Password = ""

	default:
		print("Invalid number of arguments. Options:")
		print("1 argument (no credentials required): <filename>")
		print("2 arguments: <filename> <username>")
		print("3 arguments: <filename> <username> <password>")
		print("filename structure is: [{\"folder\":\"folderName\":\"url\":\"repoUrl\"}]")
		os.Exit(1)
	}

	return userParameters

}

func getRepos(repoFile string) []Repo {
	var repos []Repo
	content, err := os.Open(repoFile)
	if err != nil {
		panic(err)
	}
	json.NewDecoder(content).Decode(&repos)
	return repos
}

func downloadRepo(repo Repo, username string, password string) {
	authRequired := (username != "") || (password != "")

	if authRequired {
		auth := &http.BasicAuth{
			Username: username,
			Password: password,
		}
		_, err := git.PlainClone(repo.Folder, false, &git.CloneOptions{URL: repo.Url, Progress: os.Stdout, Auth: auth})
		if err != nil {
			print(err)
			os.Remove(repo.Folder)
			os.Exit(1)
		}
		return
	}
	_, err := git.PlainClone(repo.Folder, false, &git.CloneOptions{URL: repo.Url, Progress: os.Stdout})
	if err != nil {
		print(err)
		os.Remove(repo.Folder)
		os.Exit(1)
	}

}

func main() {

	userParameters := getParameters(os.Args)

	projects := getRepos(userParameters.RepoFile)
	var wg sync.WaitGroup
	wg.Add(len(projects))

	for i, v := range projects {
		v := v
		i := i + 1
		go func() {
			downloadRepo(v, userParameters.Username, userParameters.Password)
			print(i, "/", len(projects), ":", v.Url, "downloaded to", v.Folder)
			defer wg.Done()
		}()
	}
	wg.Wait()

}
