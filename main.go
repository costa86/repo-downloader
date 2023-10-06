package main

import (
	"flag"
	"sync"
)

func main() {

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
