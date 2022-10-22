package models

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

type MongoDBProvider struct {
	ConnectionURI string
	Collection    *mongo.Collection
	Context       context.Context
}

func (p MongoDBProvider) GetAll() (dbHosts Hosts) {
	cursor, err := p.Collection.Find(p.Context, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	for cursor.Next(p.Context) {
		var oneHost Host
		var readHost bson.M
		if err = cursor.Decode(&readHost); err != nil {
			log.Fatal(err)
		}
		oneHost.Name = readHost["name"].(string)
		oneHost.Ip = readHost["ip"].(string)
		oneHost.Mac = readHost["mac_address"].(string)
		oneHost.Hw = readHost["vendor"].(string)
		oneHost.Date = readHost["last_seen"].(string)
		oneHost.Id = readHost["_id"].(primitive.ObjectID).Hex()
		oneHost.Known = uint16(readHost["known"].(int32))
		oneHost.Now = uint16(readHost["now"].(int32))
		dbHosts = append(dbHosts, oneHost)
	}
	return dbHosts
}

func (p MongoDBProvider) Set(h Host) {
	objID, _ := primitive.ObjectIDFromHex(h.Id)
	filter := bson.M{"_id": bson.M{"$eq": objID}}
	update := bson.M{
		"$set": bson.M{
			"name":        h.Name,
			"ip":          h.Ip,
			"mac_address": h.Mac,
			"vendor":      h.Hw,
			"last_seen":   h.Date,
			"known":       h.Known,
			"now":         h.Now,
		},
	}
	_, _ = p.Collection.UpdateOne(
		context.Background(),
		filter,
		update,
	)
}

func (p MongoDBProvider) SetLastSeen() {
	//sqlStatement := `UPDATE "now" set NOW = '0';`
	//p.execute(sqlStatement)
}

func (p MongoDBProvider) Add(h Host) {
	_, _ = p.Collection.InsertOne(p.Context, bson.M{
		"name":        h.Name,
		"ip":          h.Ip,
		"mac_address": h.Mac,
		"vendor":      h.Hw,
		"last_seen":   h.Date,
		"known":       h.Known,
		"now":         h.Now,
	})
}

func getValueFromMap(key string, dictionary map[string]interface{}) interface{} {
	if val, ok := dictionary[key]; ok {
		return val
	}
	return nil
}

func (p MongoDBProvider) Initialize(connectionString map[string]interface{}) interface{} {
	connectionURI := getValueFromMap("connectionURI", connectionString)
	dbName := getValueFromMap("database", connectionString)
	collectionName := getValueFromMap("collection", connectionString)
	if connectionURI == nil || dbName == nil || collectionName == nil {
		log.Fatalln("Invalid connection string")
	}
	client, err := mongo.NewClient(options.Client().ApplyURI(connectionURI.(string)))
	if err != nil {
		log.Fatal(err)
	}
	p.Context = context.TODO()
	err = client.Connect(p.Context)
	if err != nil {
		log.Fatal(err)
	}

	db := client.Database(dbName.(string))
	p.Collection = db.Collection(collectionName.(string))

	return p
}
