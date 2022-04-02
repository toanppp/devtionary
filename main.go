package main

import (
	"context"
	"flag"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
)

func main() {
	log.SetFlags(0)

	keywordPtr := flag.String("k", "", "keyword")
	valuePtr := flag.String("v", "", "value")
	needDeletePtr := flag.Bool("d", false, "delete")

	flag.Parse()

	keyword := *keywordPtr
	value := *valuePtr
	needDelete := *needDeletePtr

	if keyword == "" {
		log.Fatal("Missing keyword")
	}

	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		log.Fatal("Missing MONGODB_URI environment variable")
	}

	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()

	if err := client.Connect(ctx); err != nil {
		log.Fatal(err)
	}

	closeDB := func() {
		if err := client.Disconnect(ctx); err != nil {
			log.Fatal(err)
		}
	}
	defer closeDB()

	collection := client.Database("devtionary").Collection("devtionary")

	filter := bson.D{{"keyword", keyword}}

	switch {
	case needDelete:
		result, err := collection.DeleteOne(ctx, filter)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Number of deleted documents: %v\n", result.DeletedCount)

	case value != "":
		update := bson.D{{"$set", bson.D{{"value", value}}}}
		opts := options.Update().SetUpsert(true)

		result, err := collection.UpdateOne(ctx, filter, update, opts)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Number of updated documents: %v\n", result.ModifiedCount)
		fmt.Printf("Number of upserted documents: %v\n", result.UpsertedCount)

	default:
		cur := collection.FindOne(ctx, filter)

		var document bson.M
		if err = cur.Decode(&document); err != nil {
			log.Fatal(err)
		}

		fmt.Println(document["value"])
	}
}
