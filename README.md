# OmnisendTask v1

## Reviews harvesting task. Scraper and Rest endpoint to list reviews

Purpose:
Our support and customer feedback teams constantly analyzing reviews of our app in Shopify (https://apps.shopify.com/omnisend). Various metrics are calculated using internal tools. To do so, reviews should be kept in-house.

## Steps to get a project running

Project uses new go modules (go v1.11).

First clone the repository from (https://github.com/Slamadalius/OmnisendTask) to your desired folder outside of GOPATH.
```
ssh:
git clone git@github.com:Slamadalius/OmnisendTask.git

https:
git clone https://github.com/Slamadalius/OmnisendTask.git
```

You also need to set up ENV Variables MONGO_CONNECTION_STRING and DATABASE_NAME
```
export MONGO_CONNECTION_STRING=  //project owner is going to send it to you
export DATABASE_NAME=            //project owner is going to send it to you (This database is populated with reviews)
```

After that navigate to project folder and from a terminal run:
```
go build ./cmd/scraper/
go build ./cmd/rest/
```

This should create rest.exe and scraper.exe

Then from a terminal run:
```
./rest.exe

output:
Connected to MongoDb
Server is listening on port :8080
```
After this you can visit localhost:8080 from a postman or a browser. ("Query params" section below has info how to use query params)

If you want to test scraper change DATABASE_NAME to your desired name and run:
```
./scraper.exe

output:
Connected to MongoDb
Visiting ...
```
After finish database should be populated.

To run test
```
go test ./cmd/rest/
```

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
localhost:8080
localhost:8080?skip=15
localhost:8080?limit=10&skip=15
localhost:8080?sortBy=ratingDes&skip=15&limit=10
```

## Used third party packages

* [Colly](https://github.com/gocolly/colly) - Scraping Framework.
* [Mongo-go-driver](https://github.com/mongodb/mongo-go-driver) - Official MongoDB supported driver for Go.
* [HttpRouter](https://github.com/julienschmidt/httprouter) - HTTP request router (mux).
* [Testify/assert](https://github.com/stretchr/testify) - A toolkit with common assertions.

## TODO for the future

Learn more about testing with go. Set up a check if the user inserts wrong params and show a message. Create a scraper withouth using third party package. Insert reviews to a database not one by one but in batches. Learn more about MongoDB. Implement more sort options and filters. Deepen knowledge about types (It's all about the types). And more...

