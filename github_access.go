package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Repo struct {
	Releases       []Release
	TotalDownloads int
	Name           string
}

type Release struct {
	TagName     string  `json:"tag_name"`
	ReleaseName string  `json:"name"`
	Assets      []Asset `json:"assets"`
	Dl_total    int
}

type Asset struct {
	Name          string `json:"name"`
	DownloadCount int    `json:"download_count"`
}

type GitHubAccount struct {
	id           string
	access_token string
}

const GITHUB_DOMAIN = "https://api.github.com"

func getGitHubReleaseInfo(domain string, repo_name string, account GitHubAccount) (Repo, error) {
	var repo Repo

	// Access github Release API
	res, err := http.Get(domain + "/repos/" + account.id + "/" + repo_name + "/releases?access_token=" + account.access_token)
	if err != nil {
		return repo, err
	}
	defer res.Body.Close()

	// Decode
	dec := json.NewDecoder(res.Body)
	var releases []Release
	decode_err := dec.Decode(&releases)
	if decode_err != nil {
		return repo, decode_err
	}

	// Summary download count every release
	var total_downloads int
	for i := 0; i < len(releases); i++ {
		for k := 0; k < len(releases[i].Assets); k++ {
			releases[i].Dl_total += releases[i].Assets[k].DownloadCount
		}
		total_downloads += releases[i].Dl_total
	}

	repo.Releases = releases
	repo.Name = repo_name
	repo.TotalDownloads = total_downloads

	return repo, nil
}

func printReleaseInfo(repo Repo) {
	fmt.Println("[Name]:", repo.Name)
	fmt.Println("[Total]:", repo.TotalDownloads)
	for i := 0; i < len(repo.Releases); i++ {
		fmt.Println("**************************************")
		fmt.Println("[Tag Name]:", repo.Releases[i].TagName)
		fmt.Println("[Release Name]:", repo.Releases[i].ReleaseName)
		fmt.Println("[Total Downloads]:", repo.Releases[i].Dl_total)
		fmt.Println("--------------------------------------")
		for k := 0; k < len(repo.Releases[i].Assets); k++ {
			fmt.Println("[Name]:", repo.Releases[i].Assets[k].Name)
			fmt.Println("[Download Count]:", repo.Releases[i].Assets[k].DownloadCount)
		}
	}
}
