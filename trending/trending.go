package trending

import (
	"errors"
	"log"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

const (
	TrendingEndpoint = "https://github.com/trending/"
	MainEndpoint     = "https://github.com/"

	Today = "daily"
	Week  = "weekly"
	Month = "monthly"
)

type Trending struct {
	Repos []Repo // Trending repos
	lang  string // If lang is set, Repos are only of the particular language. (Default is all)
	since string // The timespan of the trending repos (Monthly, Weekly, Daily). (Default is daily)
}

type Repo struct {
	URL   string // URL repo, eg : github.com/kalbhor/gotrending
	Owner string // Owner of repo
	Name  string // Name of repo
	Stars int    // Number of stars on repo
}

// Returns a Trending type. Sets default values for lang and since
func NewTrending() *Trending {
	t := &Trending{lang: "", since: Today}
	return t
}

/* To do : check if lang is valid, return error */

// Sets the lang on a Trending type
func (t *Trending) SetLang(lang string) error {
	t.lang = lang
	return nil
}

// Sets the Since on a Trending type
func (t *Trending) SetSince(since string) error {
	if since != Today || since != Week || since != Month {
		t.since = Today
		return errors.New("Invalid value for Trending.Since")
	}
	t.since = since
	return nil
}

// Run scrapes github.com/trending to get trending repos.
// Returns a slice of repo with trending repos.
func (t *Trending) Get() []Repo {
	var repos []Repo
	var repo Repo

	query := TrendingEndpoint + t.lang + "/?since=" + t.since

	doc, err := goquery.NewDocument(query)
	if err != nil {
		log.Printf("%v\n", err)
	}

	doc.Find("ol.repo-list li").Each(func(i int, s *goquery.Selection) {
		repoPath, err := getRepoPath(s) // Eg : kalbhor/gotrending
		if err != nil {
			log.Printf("%v\n", err)
		}

		stars, err := getRepoStars(s)
		if err != nil {
			log.Printf("%v\n", err)
		}

		pathSplit := strings.Split(repoPath, "/") // Split repoPath into the owner and name of repo

		repo.URL = MainEndpoint + repoURL // Eg : github.com/ + kalbhor/gotrending
		repo.Owner = pathSplit[0]
		repo.Name = pathSplit[1]
		repo.Stars, err = strconv.Atoi(stars) // Convert stars to an int
		if err != nil {
			log.Printf("%v\n", err)
		}
		repos = append(repos, repo) // Append trending repo to repos
	})

	t.Repos = repos
	return repos

}

/* To do : check if path is valid, return error */

// Parses html to get the repo's path
func getRepoPath(s *goquery.Selection) (string, error) {
	repoPath := s.Find("h3 a").Text()
	repoPath = strings.Replace(repoPath, "\n", "", -1)
	repoPath = strings.Replace(repoPath, " ", "", -1)

	return repoPath, nil
}

/* To do : check if stars is valid, return error */

// Parses html to get number of stars
func getRepoStars(s *goquery.Selection) (string, error) {
	stars := s.Find("div.f6 a").First().Text()
	stars = strings.Replace(stars, "\n", "", -1)
	stars = strings.Replace(stars, " ", "", -1)
	stars = strings.Replace(stars, ",", "", -1)

	return stars, nil
}
