package main

import (
	"encoding/json"
	"fmt"
	"os"
	"sync"

	"github.com/go-git/go-git/v5/plumbing/transport/http"
	"gopkg.in/src-d/go-git.v4"
)

var print = fmt.Println

type Repo struct {
	Folder string `json:"folder"`
	Url    string `json:"url"`
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
	auth := &http.BasicAuth{
		Username: username,
		Password: password,
	}
	_, err := git.PlainClone(repo.Folder, false, &git.CloneOptions{URL: repo.Url, Progress: os.Stdout, Auth: auth})
	if err != nil {
		print(err)
		os.Exit(1)
	}
}

// func downloadRepo(repo Repo) {
// 	_, err := git.PlainClone(repo.Folder, false, &git.CloneOptions{URL: repo.Url, Progress: os.Stdout})
// 	if err != nil {
// 		print(err)
// 		os.Exit(1)
// 	}
// }

func main() {

	args := os.Args
	if len(args) != 4 {
		print("JSON file is missing. File structure is: [{\"folder\":\"folderName\":\"url\":\"repoUrl\"}]")
		os.Exit(1)
	}
	jsonFile := args[1]
	username := args[2]
	password := args[3]

	projects := getRepos(jsonFile)
	var wg sync.WaitGroup
	wg.Add(len(projects))

	for i, v := range projects {
		v := v
		i := i + 1
		go func() {
			downloadRepo(v, username, password)
			print(i, "/", len(projects), ":", v.Url, "downloaded to", v.Folder)
			defer wg.Done()
		}()
	}
	wg.Wait()

}
