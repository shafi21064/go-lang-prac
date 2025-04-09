package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/shafi21064/go_mongo_api/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

const connectionString = "mongodb+srv://mdshadikulislamshafi:learngo@gomongoapidatabase.wusvgms.mongodb.net/?retryWrites=true&w=majority"

// uri := "mongodb+srv://mdshadikulislamshafi:learngo%40gomongo@cluster0.mongodb.net/?retryWrites=true&w=majority"

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

// update single data
func updateSingleData(movieId string) {
	id, _ := primitive.ObjectIDFromHex(movieId)
	filter := bson.M{"_id": id}

	update := bson.M{"$set": bson.M{"watched": true}}

	result, err := collection.UpdateOne(context.Background(), filter, update)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Modified count: ", result.ModifiedCount)
}

// delete single moovie
func deleteSingleMovie(movieId string) {
	id, _ := primitive.ObjectIDFromHex(movieId)
	filter := bson.M{"_id": id}
	deletedCount, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Delete count: ", deletedCount)
}

// delete all
func deleteAllData() int64 {
	deletedResult, err := collection.DeleteMany(context.Background(), bson.D{{}}, nil)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("delete count: ", deletedResult)
	return deletedResult.DeletedCount
}

// get all movies
func getAllMovies() *[]bson.M {
	cursor, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	var movies []bson.M

	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var movie bson.M
		err := cursor.Decode(&movie)

		if err != nil {
			log.Fatal(err)
		}
		movies = append(movies, movie)
	}
	fmt.Println("all movies list")
	return &movies
}

/// Controllers function

// get all movies controller
func GetAllMoviesController(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	aLlMovies := getAllMovies()
	json.NewEncoder(w).Encode(aLlMovies)
}

// create movie controller
func CreateMovieCotroller(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow_Method", "POST")

	var movie model.NetflixModel
	json.NewDecoder(r.Body).Decode(&movie)
	insertSingleData(movie)
	json.NewEncoder(w).Encode(movie)
}

// mark as watched controller
func MarkAsWatchedControler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Method", "PUT")
	var movie *model.NetflixModel
	params := mux.Vars(r)
	updateSingleData(params["id"])
	json.NewEncoder(w).Encode(&movie)
}

// delete single movie controller
func DeleteSingleMovieController(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Method", "DELETE")
	params := mux.Vars(r)
	deleteSingleMovie(params["id"])
	json.NewEncoder(w).Encode("Deleted the data")
}

// delete all movie controller
func DeleteAllMovieController(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Method", "DELETE")
	count := deleteAllData()
	json.NewEncoder(w).Encode(count)
}
