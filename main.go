package main

import (
	"fmt"
	//"strings"

	"github.com/gocolly/colly"
)

func main() {
	c := colly.NewCollector()

	c.OnHTML("div.review-listing", func(e *colly.HTMLElement){
		fmt.Printf("Comment author: %s\n ", e.DOM.Find("h3").Text())
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	c.Visit("https://apps.shopify.com/omnisend/reviews")
}