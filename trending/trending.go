package trending

import (
	"fmt"
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
	lang  string
	Repos []Repo
	since string
}

type Repo struct {
	URL   string
	Owner string
	Name  string
	Stars int
}

func NewTrending() *Trending {
	t := &Trending{lang: "", since: Today}
	return t
}

func (t *Trending) SetLang(lang string) {
	t.lang = lang
}
func (t *Trending) SetSince(since string) {
	t.since = since
}

func (t *Trending) Get() []Repo {
	var repos []Repo
	var repo Repo

	query := TrendingEndpoint + t.lang + "/?since=" + t.since

	doc, err := goquery.NewDocument(query)
	if err != nil {
		log.Printf("%v\n", err)
	}

	doc.Find("ol.repo-list li").Each(func(i int, s *goquery.Selection) {
		repoURL, err := getRepoPath(s)
		if err != nil {
			log.Printf("%v\n", err)
		}

		stars, err := getRepoStars(s)
		if err != nil {
			log.Printf("%v\n", err)
		}

		pathSplit := strings.Split(repoURL, "/")

		repo.URL = MainEndpoint + repoURL
		repo.Owner = pathSplit[0]
		repo.Name = pathSplit[1]
		repo.Stars, err = strconv.Atoi(stars)
		if err != nil {
			log.Printf("%v\n", err)
		}
		repos = append(repos, repo)
	})

	t.Repos = repos
	return repos

}

func getRepoPath(s *goquery.Selection) (string, error) {
	repoPath := s.Find("h3 a").Text()
	repoPath = strings.Replace(repoPath, "\n", "", -1)
	repoPath = strings.Replace(repoPath, " ", "", -1)

	return repoPath, nil
}

func getRepoStars(s *goquery.Selection) (string, error) {
	stars := s.Find("div.f6 a").First().Text()
	stars = strings.Replace(stars, "\n", "", -1)
	stars = strings.Replace(stars, " ", "", -1)
	stars = strings.Replace(stars, ",", "", -1)

	return stars, nil
}

func (r *Repo) Print() {
	fmt.Println("URL : ", r.URL)
	fmt.Println("Owner : ", r.Owner)
	fmt.Println("Repo Name : ", r.Name)
	fmt.Println("Stars : ", r.Stars)
}
