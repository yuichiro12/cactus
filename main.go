package main

import (
	"fmt"
	"log"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	doc, err := goquery.NewDocument("https://www.nytimes.com/2018/03/04/technology/silicon-valley-midwest.html?rref=collection%2Fsectioncollection%2Ftechnology&action=click&contentCollection=technology&region=stream&module=stream_unit&version=latest&contentPlacement=4&pgtype=sectionfront")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s --\n", doc.Find("#headline").Text())
	doc.Find(".story-body-text").Each(func(i int, s *goquery.Selection) {
		fmt.Printf("%s\n", s.Text())
	})
}
