package main

import (
	"context"
	"log"
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

	repos, err := tr.Get()
	if err != nil {
		log.Printf("%v\n", err)
	}

	/* To do : accept and handle errors */
	for _, repo := range repos {
		if _, ok := err.(*github.RateLimitError); ok {
			log.Println("Hit rate limit")
			break
		}

		isStarred, _, err := client.Activity.IsStarred(ctx, repo.Owner, repo.Name)
		if err != nil {
			log.Printf("%v\n", err)
			break
		}

		if !isStarred {
			_, err = client.Activity.Star(ctx, repo.Owner, repo.Name)
			if err != nil {
				log.Printf("%v\n", err)
				break
			}

		}

		isFollowing, _, err := client.Users.IsFollowing(ctx, "", repo.Owner)
		if err != nil {
			log.Printf("%v\n", err)
			break
		}

		if !isFollowing {
			_, err = client.Users.Follow(ctx, repo.Owner)
			if err != nil {
				log.Printf("%v\n", err)
				break
			}

		}
	}

}
