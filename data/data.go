package data

import (
	"context"
	"example/ginference-server/config/devconfig"
	"example/ginference-server/models/model"
	"example/ginference-server/models/user"
	"fmt"
	"time"

	"github.com/google/uuid"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

var RegisteredUsers = user.Users{
	{UserID: uuid.New(), UserName: "Tom", CreatedAt: time.Now(), ModifiedAt: time.Now()},
	{UserID: uuid.New(), UserName: "Joe", CreatedAt: time.Now(), ModifiedAt: time.Now()},
	{UserID: uuid.New(), UserName: "Harry", CreatedAt: time.Now(), ModifiedAt: time.Now()},
}

var user1, _ = RegisteredUsers.FindByName("Tom")
var user2, _ = RegisteredUsers.FindByName("Joe")

//var user3, _ = RegisteredUsers.FindByName("Harry")

var SubscribedModels = model.AIModels{
	{ModelID: uuid.New(), ModelName: "pickachu_1", CreatedBy: user1.UserID.String(), CreatedAt: time.Now()},
	{ModelID: uuid.New(), ModelName: "bulbasaur_1", CreatedBy: user2.UserID.String(), CreatedAt: time.Now()},
	{ModelID: uuid.New(), ModelName: "bulbasaur_2", CreatedBy: user2.UserID.String(), CreatedAt: time.Now()},
}

func MongoDBInit() *mongo.Client {
	var dbConnectionString string
	var clientOpts *options.ClientOptions
	if devconfig.WithTLS {
		dbConnectionString = devconfig.DBConnectionStringWithTLS
		// credential := options.Credential{
		// 	AuthMechanism: "MONGODB-X509",
		// }
		clientOpts = options.Client().ApplyURI(dbConnectionString)
		//.SetAuth(credential)
	} else {
		dbConnectionString = devconfig.DBConnectionString
		clientOpts = options.Client().ApplyURI(dbConnectionString)
	}
	dbClient, dbConnectionErr := mongo.Connect(clientOpts)
	if dbConnectionErr != nil {
		panic(dbConnectionErr)
	}
	return dbClient
}

func Find[T any](docs []T, dbName string, collName string, filter bson.D, findOptions *options.FindOptionsBuilder) ([]T, error) {
	dbClient := MongoDBInit()
	defer MongoDBDisconnect(dbClient)
	coll := dbClient.Database(dbName).Collection(collName)
	cur, err := coll.Find(context.TODO(), filter, findOptions)
	if err != nil {
		return docs, fmt.Errorf("error finding the document - %v", err)
	}
	for cur.Next(context.TODO()) {
		var doc T
		if err := cur.Decode(&doc); err != nil {
			return docs, fmt.Errorf("error decoding the document - %v", err)
		}
		docs = append(docs, doc)
	}
	if err := cur.Err(); err != nil {
		return docs, fmt.Errorf("error iterating through the documents - %v", err)
	}
	cur.Close(context.TODO())
	return docs, nil
}

func Create[T any](doc T, dbName string, collName string) error {
	dbClient := MongoDBInit()
	defer MongoDBDisconnect(dbClient)
	coll := dbClient.Database(dbName).Collection(collName)
	_, err := coll.InsertOne(context.TODO(), doc)
	return err
}

func EditOne[T any](doc T, dbName string, collName string, filter bson.D, updateOptions *options.UpdateOneOptionsBuilder) error {
	dbClient := MongoDBInit()
	defer MongoDBDisconnect(dbClient)
	coll := dbClient.Database(dbName).Collection(collName)
	bytArr, err := bson.Marshal(doc)
	if err != nil {
		return err
	}
	var updateData bson.D
	if err := bson.Unmarshal(bytArr, &updateData); err != nil {
		return err
	}
	updateData = bson.D{{Key: "$set", Value: updateData}}
	res, updateErr := coll.UpdateOne(context.TODO(), filter, updateData, updateOptions)
	if updateErr != nil {
		return updateErr
	}
	if res.MatchedCount == 0 {
		return fmt.Errorf("no such user")
	}
	if res.MatchedCount > 0 && res.ModifiedCount == 0 {
		return fmt.Errorf("internal server error")
	}
	return nil
}

func MongoDBDisconnect(dbClient *mongo.Client) {
	if err := dbClient.Disconnect(context.TODO()); err != nil {
		panic(err)
	}
}
