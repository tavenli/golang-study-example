package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/**
CRUD Operations
https://www.mongodb.com/docs/drivers/go/current/fundamentals/crud/

*/

func main() {
	fmt.Println("--------------------------")

	MongoDbDemo1()
	MongoDbDemo2()
	MongoDbDemo3()

}

func MongoDbDemo1() {
	//uri := "mongodb://<hostname>:<port>"
	uri := "mongodb://127.0.0.1:27017"
	// Use the SetServerAPIOptions() method to set the Stable API version to 1
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)
	// Create a new client and connect to the server
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()
	// Send a ping to confirm a successful connection
	var result bson.M
	if err := client.Database("admin").RunCommand(context.TODO(), bson.D{{"ping", 1}}).Decode(&result); err != nil {
		panic(err)
	}
	fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")

}

func MongoDbDemo2() {
	uri := "mongodb://127.0.0.1:27017"
	credential := options.Credential{
		AuthSource: "<authenticationDb>",
		Username:   "<username>",
		Password:   "<password>",
	}
	clientOpts := options.Client().ApplyURI(uri).
		SetAuth(credential)

	client, err := mongo.Connect(context.TODO(), clientOpts)
	fmt.Println(client, err)

}

func MongoDbDemo3() {

	//uri := "mongodb://<hostname>:<port>?tls=true"
	//opts := options.Client().ApplyURI(uri).SetTLSConfig(&tls.Config{})

	uri := "mongodb://127.0.0.1:27017"
	opts := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(context.TODO(), opts)
	fmt.Println(client, err)

	coll := client.Database("DbDemo1").Collection("students")
	address1 := Address{"1 Lakewood Way", "Elwood City", "PA"}
	student1 := Student{FirstName: "Arthur", Address: address1, Age: 30}
	res, err := coll.InsertOne(context.TODO(), student1)
	fmt.Println(res.InsertedID, err)

	filter := bson.D{{"age", 30}}

	var result bson.D
	err = coll.FindOne(context.TODO(), filter).Decode(&result)
	fmt.Println(result, err)

	filter = bson.D{{"age", bson.D{{"$gt", 20}}}}
	cursor, err := coll.Find(context.TODO(), filter)
	if err != nil {
		panic(err)
	}
	var results []Student
	if err = cursor.All(context.TODO(), &results); err != nil {
		panic(err)
	}
	for _, _result := range results {
		_res, _ := bson.MarshalExtJSON(_result, false, false)
		fmt.Println(string(_res))
	}

	opts1_2 := options.Find().SetSort(bson.D{{"first_name", 1}}).SetSkip(2).SetLimit(2)
	cursor1_2, err1_2 := coll.Find(context.TODO(), filter, opts1_2)
	if err1_2 != nil {
		panic(err1_2)
	}
	var results1_2 []bson.D
	if err1_2 = cursor1_2.All(context.TODO(), &results1_2); err1_2 != nil {
		panic(err1_2)
	}

	//SetHint 指定一个 index 索引来进行统计
	opts2 := options.Count().SetHint("_id_")
	count, err := coll.CountDocuments(context.TODO(), bson.D{}, opts2)
	if err != nil {
		panic(err)
	}

	fmt.Println(count)

	//filter = bson.D{{"_id", 100}}

	id, err := primitive.ObjectIDFromHex("6630b00590b2c4cadfb5d685")
	if err != nil {
		panic(err)
	}
	filter = bson.D{{"_id", id}}
	update := bson.D{{"$set", bson.D{{"first_name", "Mary Wollstonecraft Shelley"},
		{"city", "Marketing Director"}}}, {"$inc", bson.D{{"age", 2000}}}}
	result2, err := coll.UpdateOne(context.TODO(), filter, update)
	fmt.Printf("Documents matched: %v\n", result2.MatchedCount)
	fmt.Printf("Documents updated: %v\n", result2.ModifiedCount)

	filter = bson.D{{"age", bson.D{{"$gt", 300}}}}
	opts3 := options.Delete().SetHint(bson.D{{"_id", 1}})
	result3, err := coll.DeleteMany(context.TODO(), filter, opts3)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Number of documents deleted: %d\n", result3.DeletedCount)

}

type Address struct {
	Street string
	City   string
	State  string
}

type Student struct {
	FirstName string  `bson:"first_name,omitempty"`
	LastName  string  `bson:"last_name,omitempty"`
	Address   Address `bson:"inline"`
	Age       int
}
