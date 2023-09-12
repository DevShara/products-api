package controller

import (
	"context"
	"fmt"
	"log"
	"os"
	"products-api/model"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func goDotEnvVariable(key string) string {

	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}

var connectionString = goDotEnvVariable("DB_URI")

const dbName = "products-api"
const colName = "products"

var collection *mongo.Collection

//connect with MongoDB

func Init() {
	//client option
	clientOption := options.Client().ApplyURI(connectionString)

	//connect to mongodb
	client, err := mongo.Connect(context.TODO(), clientOption)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("MongoDB Connection Success")

	collection = client.Database(dbName).Collection(colName)

	fmt.Println("Collection instance is ready")

}

func InsertProduct(product model.Product) {
	inserted, err := collection.InsertOne(context.Background(), product)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted 1 product", inserted.InsertedID)
}

// func UpdateProduct(product string) {
// 	id, _ := primitive.ObjectIDFromHex(product)

// 	fmt.Println(id)
// }

func GetAllProducts() []bson.M {
	cur, err := collection.Find(context.Background(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}

	var products []bson.M

	for cur.Next(context.Background()) {
		var product bson.M
		err := cur.Decode(&product)
		if err != nil {
			log.Fatal(err)
		}

		products = append(products, product)
	}

	defer cur.Close(context.Background())

	return products
}
