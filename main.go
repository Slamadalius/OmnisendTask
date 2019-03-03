package main

import (
	"fmt"
	"strings"

	"github.com/gocolly/colly"
)

// Declaring struct type for Review structure
type Review struct {
	Author string
	Rating string
	Date   string
	Body   string
}

func main() {
	// Creating colly collector, here you can specify other options
	// like AllowedDomains, AllowRevisits or Async
	c := colly.NewCollector()

	// Creating a callback function on every div with class review-listing
	c.OnHTML("div.review-listing", func(e *colly.HTMLElement){
		// Declaring review with values from scraped reviews
		review := Review{
			Author: e.ChildText("h3.review-listing-header__text"),
			Rating: e.ChildAttr("div.ui-star-rating", "data-rating"),
			//e.ChildText trims the values. In this case I am using strings package to trim space and new lines
			Date:   strings.Trim(e.DOM.Find("div.review-metadata__item .review-metadata__item-value").Last().Text(), " \n"), 
			Body:   e.ChildText("p"),
		}

		fmt.Printf("On %s Comment author: %s\nGave a rating of %s stars\n%s\n\n", review.Date, review.Author, review.Rating, review.Body)
	})

	// Callback function to go through pagination links
	c.OnHTML("a.search-pagination__link", func(e *colly.HTMLElement){
		link := e.Attr("href")

		// visit the link specified
		c.Visit(e.Request.AbsoluteURL(link))
	})

	//Logging which page is being visited
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	c.Visit("https://apps.shopify.com/omnisend/reviews")
}