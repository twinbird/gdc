package main

import (
	"errors"
	"fmt"
	"os"
)

func loadUserAccount() (GitHubAccount, error) {
	gdc_id := os.Getenv("GDC_ID")
	gdc_token := os.Getenv("GDC_ACCESS_TOKEN")

	if gdc_id == "" || gdc_token == "" {
		return GitHubAccount{}, errors.New("Configuration environment variable not found.")
	}

	return GitHubAccount{id: gdc_id, access_token: gdc_token}, nil
}

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Usage: gdc [repository name]")
		os.Exit(1)
	}

	acc, acc_err := loadUserAccount()
	if acc_err != nil {
		fmt.Println("Configuration environment variable not found.")
		os.Exit(1)
	}

	repo, err := getGitHubReleaseInfo(GITHUB_DOMAIN, os.Args[1], acc)
	if err != nil {
		fmt.Println("GitHub API Access failed.")
		os.Exit(1)
	}

	printReleaseInfo(repo)

	os.Exit(0)
}
