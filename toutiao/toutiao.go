package toutiao

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
)

var url = "https://toutiao.io/search?utf8=✓&q="
var link = "https://toutiao.io/"

type Article struct {
	Title   string
	Link    string
	Summary string
	Meta    string
}

func (a Article) ToString() string {
	return fmt.Sprintf("title:%s\nlink:%s\nsummary:%s\nmeta:%s", a.Title, a.Link, a.Summary, a.Meta)
}

func GetArticles(keyword string) (articles []Article) {
	doc, err := goquery.NewDocument(url + keyword)
	if err != nil {
		log.Fatal(err)
	}

	// Find the review items
	doc.Find(".posts .post").Each(func(i int, s *goquery.Selection) {
		article := Article{}
		// For each item found, get the band and title
		article.Title = s.Find(".content .title a").Text()
		article.Link, _ = s.Find(".content .title a").Attr("href")
		article.Link = link + article.Link
		article.Summary = s.Find(".content .summary a").Text()
		article.Meta = s.Find(".content .meta span").Prev().Text()
		articles = append(articles, article)
	})
	return
}
