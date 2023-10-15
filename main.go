package main

import (
	"flag"
	"runtime"
	"sync"
)

func main() {
	maxWorkers := runtime.NumCPU()

	var (
		file,
		username,
		password string
	)

	flag.StringVar(&file, "file", "repos.json", "JSON file with repositories")
	flag.StringVar(&username, "username", "", "username (case authentication is required)")
	flag.StringVar(&password, "password", "", "password (case authentication is required)")
	flag.Parse()

	username, password = getCredentials(username, password)
	projects := getRepos(file)
	var wg sync.WaitGroup
	wg.Add(len(projects))
	projectChan := make(chan Repo, len(projects))
	defer close(projectChan)

	for _, v := range projects {
		projectChan <- v
	}

	for i := 0; i < maxWorkers; i++ {
		go func() {
			for v := range projectChan {
				downloadRepo(v, username, password)
				print(v.Url, "downloaded to", v.Folder)
				wg.Done()
			}
		}()
	}

	wg.Wait()

}
