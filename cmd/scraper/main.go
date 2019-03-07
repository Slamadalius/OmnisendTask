package main

import (
	"context"
	"log"
    "fmt"
	"strings"
	"strconv"

	"github.com/Slamadalius/OmnisendTask/config"
	"github.com/Slamadalius/OmnisendTask/models"

    "github.com/gocolly/colly"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson"
)


func writeToDB(ctx context.Context, db *mongo.Database, review models.Review) {
	collection := db.Collection("reviews")

	// Inserting reviews
	_, err := collection.InsertOne(ctx, bson.D{
		{Key:"author", Value:review.Author},
		{Key:"rating", Value:review.Rating},
		{Key:"date",   Value:review.Date},
		{Key:"body",   Value:review.Body},
	})
	if err != nil { 
		log.Fatal(err) 
	}
}

func main() {
	// getting context and db connection from config
	ctx := config.CTX
	db  := config.DB

    // Creating colly collector, here you can specify other options
	// like AllowedDomains, AllowRevisits or Async
	// Async turns on asynchronous network communication.
	// if async is set to true it launch func fetch to a go routine
	c := colly.NewCollector(
		colly.MaxDepth(2),
		colly.Async(true),
	)

	// Parallelism is the number of the maximum allowed concurrent requests of the matching domains
	c.Limit(&colly.LimitRule{Parallelism: 2})

    // Creating a callback function on every div with class review-listing
    c.OnHTML("div.review-listing", func(e *colly.HTMLElement){
		author       := e.ChildText("h3.review-listing-header__text")
		ratingString := e.ChildAttr("div.ui-star-rating", "data-rating")
		//e.ChildText trims the values. In this case I am using strings package to trim space and new lines
		date         := strings.Trim(e.DOM.Find("div.review-metadata__item .review-metadata__item-value").Last().Text(), " \n")
		body         := e.ChildText("p")

		ratingInt, err := strconv.Atoi(ratingString)
		if err != nil {
			log.Println(err)
		}
		rating := uint8(ratingInt)

        // Declaring review with values from scraped reviews
        review := models.Review{
            Author: author,
            Rating: rating,
            Date:   date,
            Body:   body,
		}
		
		// Writing scraped review to a database
		writeToDB(ctx, db, review)

		//fmt.Printf("On %s Comment author: %s\nGave a rating of %d stars\n%s\n\n", review.Date, review.Author, review.Rating, review.Body)
	})
	
	
    // Callback function to go through pagination links
    c.OnHTML("a.search-pagination__link", func(e *colly.HTMLElement){
        link := e.Attr("href")

        // visit the link
        c.Visit(e.Request.AbsoluteURL(link))
    })

    //Logging which page is being visited
    c.OnRequest(func(r *colly.Request) {
        fmt.Println("Visiting", r.URL.String())
	})

	c.Visit("https://apps.shopify.com/omnisend/reviews")
	
	// Returns when the jobs are finished
	c.Wait()

	fmt.Println("Finished")
}