package main

import (
	"fmt"
	"log"
	"os/exec"
	"time"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	for {
		fmt.Println("Getting articles from hckrnews.com...")
		doc, err := goquery.NewDocument("https://hckrnews.com")
		if err != nil {
			log.Fatal(err)
		}

		doc.Find("a.link").EachWithBreak(func(i int, s *goquery.Selection) bool {
			if i > 5 {
				return false
			}
			link, exists := s.Attr("href")
			title := s.Text()
			if exists {
				readArticle(title, link)
			}
			time.Sleep(3 * time.Second)
			return true
		})
	}
}

func readArticle(title string, link string) {
	doc, err := goquery.NewDocument(link)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s -- %s\n", title, link)
	err = exec.Command("say", title).Run()
	if err != nil {
		log.Fatal(err)
	}
	doc.Find("p").Each(func(i int, s *goquery.Selection) {
		err = exec.Command("say", s.Text()).Run()
		if err != nil {
			log.Fatal(err)
		}
	})
}
