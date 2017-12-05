package main

import (
	"context"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

const (
	Token = os.Env("GITHUB_ACCESS_TOKEN")
)

func main() {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: Token},
	)
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)

}
