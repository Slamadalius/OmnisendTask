package main

import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"

	"github.com/julienschmidt/httprouter"
	"go.mongodb.org/mongo-driver/bson"

	"github.com/SlamaDalius/OmnisendTask/config"
	"github.com/SlamaDalius/OmnisendTask/models"
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params){
	// Creating slice of struct to store all reviews
	var allReviews []models.Review

	// Getting context and db connection from config
	ctx := config.CTX
	db  := config.DB

	// Running Find query to fetch all reviews from review collection
	coll, err := db.Collection("reviews").Find(ctx, bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	defer coll.Close(ctx)

	// Find returns a mongo cursor that can be used to iterate over a collections using Next()
	for coll.Next(ctx) {
		// Setting elem to a bson Document address
		elem := &bson.D{}
		
		// Decoding document into val
		err = coll.Decode(elem)
		if err != nil {
			log.Fatal(err)
		}

		m := elem.Map()

		review := models.Review{
			Author: m["author"].(string),
			Rating: m["rating"].(string),
			Date:   m["date"].(string),
			Body:   m["body"].(string),
		}

		// Pushing single review to allReviews slice
		allReviews = append(allReviews, review)
	}

	// Encoding allReviews to Json format
	reviewsJson, err := json.Marshal(allReviews)
	if err != nil {
		fmt.Println(err)
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