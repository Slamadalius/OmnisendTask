# OmnisendTask v1

## Reviews harvesting task. Scraper and Rest endpoint to list reviews

Purpose:
Our support and customer feedback teams constantly analyzing reviews of our app in Shopify (https://apps.shopify.com/omnisend). Various metrics are calculated using internal tools. To do so, reviews should be kept in-house.

## Steps to get a project running

## Used third party packages

* [Colly](https://github.com/gocolly/colly) - Scraping Framework.
* [Mongo-go-driver](https://github.com/mongodb/mongo-go-driver) - Official MongoDB supported driver for Go.
* [HttpRouter](https://github.com/julienschmidt/httprouter) - HTTP request router (mux).
* [Testify/assert](https://github.com/stretchr/testify) - A toolkit with common assertions.

## Query params

To set limit, skip or sortBy you need to pass those values through query params.

Max number of results returned set to 30.

```
limit  int                        (limits how many results returned)
skip   int                        (skips as many results as provided in skip param)
sortBy string ratingDes/ratingAsc (sorts results by rating in descending or ascending order)
```
Without any params results returned by newest date. 
Params can be used individually.

Examples
```
localhost:8080?skip=15
localhost:8080?limit=10&skip=15
localhost:8080?sortBy=ratingDes&skip=15&limit=10
```

## TODO for the future

Learn more about testing with go. Set up a check if the user inserts wrong params and show a message. Create a scraper withouth using third party package. Learn more about MongoDB. Implement more sort options and filters. Deepen knowledge about types (It's all about the types). And more...

