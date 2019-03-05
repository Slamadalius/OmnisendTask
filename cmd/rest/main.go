package main

import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/SlamaDalius/OmnisendTask/config"
	"github.com/SlamaDalius/OmnisendTask/models"
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params){
	// Creating slice of struct to store all reviews
	var allReviews []models.Review

	// Getting params from query (?skip=) 
	queryValues := r.URL.Query()
	batchSize := int32(20)
	skip, _     := strconv.ParseInt(queryValues.Get("skip"), 10, 32)
	limit, _    := strconv.ParseInt(queryValues.Get("limit"), 10, 32)
	sortBy      := queryValues.Get("sortBy")

	// Getting context and db connection from config
	ctx := config.CTX
	db  := config.DB

	// Batch which set the number of documents to return in every batch.
	// Limit which limits results returned (values is set from query params)
	// Skip which specifies number of documents to skip before returning (values is set from query params)
	options := options.FindOptions{}
	options.BatchSize = &batchSize
	options.Limit = &limit
	options.Skip  = &skip
	
	// Switch statement to sort results by rating in Ascending and Descending order
	switch sortBy{
	case "ratingAsc":
		options.Sort  = bson.M{"rating": 1}
	case "ratingDes":
		options.Sort  = bson.M{"rating": -1}
	default:	
	}

	// Running Find query to fetch all reviews from review collection with options:
	coll, err := db.Collection("reviews").Find(ctx, bson.D{}, &options)
	if err != nil {
		log.Fatal(err)
	}
	defer coll.Close(ctx)

	// Find returns a mongo cursor that can be used to iterate over a collections using Next()
	for coll.Next(ctx) {
		// Setting elem to a bson Document address
		var review models.Review
		
		// Decoding document into val
		err = coll.Decode(&review)
		if err != nil {
			log.Fatal(err)
		}

		// Pushing single review to allReviews slice
		allReviews = append(allReviews, review)
	}

	// Encoding allReviews to Json format
	reviewsJson, err := json.Marshal(allReviews)
	if err != nil {
		log.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s\n", reviewsJson)
}

func main() {
	r := httprouter.New()
	r.GET("/", Index)

	fmt.Println("Server is listening on port :8080")
	err := http.ListenAndServe("localhost:8080", r)
	if err != nil {
		log.Fatal(err)
	}
}