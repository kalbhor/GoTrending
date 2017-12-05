package main

import (
	"context"
	"os"

	"github.com/google/go-github/github"
	"github.com/kalbhor/GoTrending/trending"
	"golang.org/x/oauth2"
)

var (
	Token = os.Getenv("GITHUB_ACCESS_TOKEN")
)

func main() {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: Token},
	)
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)
	tr := trending.NewTrending()
	tr.SetLang("go")

	repos := tr.Get()

	for _, repo := range repos {
		client.Activity.Unstar(ctx, repo.Owner, repo.Name)
	}

}
