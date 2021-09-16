package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

//scrape golangcode.com for blog titles

func main() {

	blogTitles, err := GetLatestBlogTitles("https://golangcode.com")
	if err != nil {
		log.Println(err)
	}

	fmt.Println("Blog Titles:")
	fmt.Println(blogTitles)

}

func GetLatestBlogTitles(url string) (string, error) {
	res, err := http.Get(url)
	if err != nil {
		log.Println(err)
		return "", err
	}
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Println(err)
		return "", err
	}
	titles := ""

	//find posts by the html id
	doc.Find(".post-title").Each(func(i int, s *goquery.Selection) {
		titles += s.Text() + "\n"
	})
	return titles, nil
}
