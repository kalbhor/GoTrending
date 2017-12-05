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
	/* To do : accept and handle errors */
	for _, repo := range repos {
		isStarred, _, _ := client.Activity.IsStarred(ctx, repo.Owner, repo.Name)
		if !isStarred {
			_, _ = client.Activity.Star(ctx, repo.Owner, repo.Name)
		}
		isFollowing, _, _ := client.Users.IsFollowing(ctx, "", repo.Owner)
		if !isFollowing {
			_, _ = client.Users.Follow(ctx, repo.Owner)
		}
	}

}
