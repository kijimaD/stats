package main

import (
	"context"
	"fmt"
	"github.com/google/go-github/v47/github"
)

type Scrape struct {
	Repos []Repo
}

type Repo struct {
	Name            string
	Description     string
	Language        string
	HTMLURL         string
	DefaultBranch   string
	ForksCount      int
	StargazersCount int
	CommitCount     int
}

func newScrape() Scrape {
	var scrape Scrape
	client := login()
	repos, _, _ := client.Repositories.List(context.Background(), "kijimaD", nil)
	for _, r := range repos {
		scrape.Repos = append(scrape.Repos, newRepo(r))
	}
	return scrape
}

func newRepo(res *github.Repository) Repo {
	// ヌルポ対策
	var desc string
	var lang string
	if res.Description == nil {
		desc = ""
	} else {
		desc = *res.Description
	}
	if res.Language == nil {
		lang = ""
	} else {
		lang = *res.Language
	}

	result := Repo{
		*res.Name,
		desc,
		lang,
		*res.HTMLURL,
		*res.DefaultBranch,
		*res.ForksCount,
		*res.StargazersCount,
		0, // *contributors[0].Total,
	}
	return result
}

func contrib(reponame string) string {
	client := login()
	context := context.Background()
	// TODO: ContributionStatsは、キャッシュが未生成の場合少し待って再リトライしてあげないといけない
	contributors, _, err := client.Repositories.ListContributorsStats(context, "kijimaD", reponame)
	if err != nil {
		// panic(err)
		fmt.Println(err)
	} else {
		fmt.Println(*contributors[0].Total)
	}

	return "a"
}