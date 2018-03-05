package main

import (
	"fmt"
	"log"
	"time"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	doc, err := goquery.NewDocument("https://hckrnews.com")
	if err != nil {
		log.Fatal(err)
	}

	doc.Find("a.link").EachWithBreak(func(i int, s *goquery.Selection) bool {
		if i > 10 {
			return false
		}
		link, exists := s.Attr("href")
		if exists {
			readArticle(link)
		}
		time.Sleep(3 * time.Second)
		return true
	})
}

func readArticle(link string) {
	doc, err := goquery.NewDocument(link)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s --\n", doc.Find("h1").Text())
	doc.Find("p").Each(func(i int, s *goquery.Selection) {
		fmt.Printf("%s\n", s.Text())
	})
}
