package controller

import (
	"context"
	"fmt"
	"log"

	"github.com/shafi21064/go_mongo_api/model"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

const connectionString = "mongodb+srv://atlas-sample-dataset-load-67f54a292966397e642892f1:learngo@gomongoapidatabase.wusvgms.mongodb.net/?retryWrites=true&w=majority&appName=GoMongoApiDatabase"

const dbName = "netflix"

const collectionName = "watchlist"

// Most important part
var collection *mongo.Collection

//connect with mongodb

func init() {
	//clint option
	clintOption := options.Client().ApplyURI(connectionString)

	//connct to mongodb
	client, err := mongo.Connect(clintOption)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Mongodb connection success")

	collection = client.Database(dbName).Collection(collectionName)

	fmt.Println("collection ref is ready")

}

// mongodb helpers

// insert single data

func insertSingleData(netflixModel model.NetflixModel) {
	inserted, err := collection.InsertOne(context.Background(), netflixModel)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted data id is: ", inserted.InsertedID)
}
