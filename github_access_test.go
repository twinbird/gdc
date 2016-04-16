package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGitHubAccess(t *testing.T) {
	const right_response_json = `
	[
		{
			"tag_name": "0.1",
			"name": "Release 0.1",
			"assets": [
				{
					"name": "windows 0.1.zip",
					"download_count": 1
				},
				{
					"name": "linux 0.1.zip",
					"download_count": 2
				}
			]
		},
		{
			"tag_name": "0.2",
			"name": "Release 0.2",
			"assets": [
				{
					"name": "linux 0.2.zip",
					"download_count": 4
				}
			]
		}
	]`

	ts := httptest.NewServer(http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path != "/repos/test_id/test_repo/releases" {
				t.Fatalf("Request path Expect %s, But Actual %s",
					"/repos/test_id/test_repo/releases",
					r.URL.Path)
			}
			if r.URL.Query().Get("access_token") != "test_access_token" {
				t.Errorf("Expect access_token is 'test_access_token' but actual %s",
					r.URL.Query().Get("access_token"))
			}
			w.Header().Set("Content-Type", "application/json")
			fmt.Fprintf(w, right_response_json)
			return
		},
	))
	defer ts.Close()

	repo, err := getGitHubReleaseInfo(ts.URL, "test_repo", GitHubAccount{id: "test_id", access_token: "test_access_token"})

	if err != nil {
		t.Fatalf("Should not occur error in getGitHubReleaseInfo. Err:%s", err)
	}

	if repo.TotalDownloads != 7 {
		t.Errorf("Total Downloads Summarize error. Expect:7, Actual:%d", repo.TotalDownloads)
	}

	if repo.Name != "test_repo" {
		t.Errorf("repo.Name is wrong. Expect:'test_repo', Actual:'%s'", repo.Name)
	}

	if len(repo.Releases) != 2 {
		t.Fatalf("Expect 2 Releases data. But Actual %d", len(repo.Releases))
	}
}
