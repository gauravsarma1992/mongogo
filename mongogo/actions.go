package mongogo

import (
	"context"
	"time"

	"github.com/gauravsarma1992/gostructs"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func (db *DB) Create(model interface{}) (uuid string, err error) {
	var (
		decodedRes *gostructs.DecodedResult
	)
	if decodedRes, err = db.decoder.Decode(model); err != nil {
		return
	}
	decodedRes.Attributes["uuid"] = generateUUID()
	decodedRes.Attributes["created_at"] = time.Now().UTC()
	decodedRes.Attributes["updated_at"] = time.Now().UTC()

	collection := db.Database.Collection(decodedRes.Name)
	if _, err = collection.InsertOne(context.Background(), decodedRes.Attributes); err != nil {
		return
	}
	uuid = decodedRes.Attributes["uuid"].(string)
	return
}

func (db *DB) FindOne(model interface{}) (resModel map[string]interface{}, err error) {
	var (
		decodedRes *gostructs.DecodedResult
	)
	if decodedRes, err = db.decoder.Decode(model); err != nil {
		return
	}
	resModel = make(map[string]interface{})
	collection := db.Database.Collection(decodedRes.Name)

	if err = collection.FindOne(context.Background(), decodedRes.Attributes).Decode(resModel); err != nil {
		return
	}
	return
}

func (db *DB) FindMany(model interface{}) (resModels []map[string]interface{}, err error) {
	var (
		decodedRes *gostructs.DecodedResult
		cur        *mongo.Cursor
	)
	if decodedRes, err = db.decoder.Decode(model); err != nil {
		return
	}
	collection := db.Database.Collection(decodedRes.Name)

	if cur, err = collection.Find(context.Background(), decodedRes.Attributes); err != nil {
		return
	}
	for cur.Next(context.Background()) {
		result := make(map[string]interface{})
		if err = cur.Decode(&result); err != nil {
			return
		}
		resModels = append(resModels, result)
	}
	return
}

func (db *DB) UpdateOne(updatedAttrs map[string]interface{}, model interface{}) (err error) {
	var (
		decodedRes *gostructs.DecodedResult
	)
	if decodedRes, err = db.decoder.Decode(model); err != nil {
		return
	}
	collection := db.Database.Collection(decodedRes.Name)

	if _, err = collection.UpdateOne(context.Background(), decodedRes.Attributes, bson.M{"$set": updatedAttrs}); err != nil {
		return
	}
	return
}
