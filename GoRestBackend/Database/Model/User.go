package model

import (
	"context"
	"time"

	MongoConnInit "github.com/NirmalVatsyayan/GoRestBackend/Database/MongoConn"
	"go.mongodb.org/mongo-driver/bson"
)

type User struct {
	ID        string    `bson:"ID" json:"ID"`
	FirstName string    `bson:"FirstName" json:"FirstName"`
	LastName  string    `bson:"LastName" json:"LastName"`
	Email     string    `bson:"Email" json:"Email"`
	Password  string    `bson:"Password" json:"Password"`
	Created   time.Time `bson:"Created" json:"Created"`
	Updated   time.Time `bson:"Updated" json:"Updated"`
}

func (a *User) CollectionName() string {
	return "User"
}

func InsertUser(a *User) (err error) {

	_, err = MongoConnInit.MDB.Collection(a.CollectionName()).InsertOne(context.TODO(), a)
	if err != nil {
		return err
	}
	return nil

}

func CountUser(email string) (int, error) {

	var a User
	var count int64
	var err error

	if count, err = MongoConnInit.MDB.Collection(a.CollectionName()).CountDocuments(context.TODO(), bson.M{"Email": email}); err != nil {
		return 0, err
	}

	return int(count), nil
}

func GetUser(a *User, email string) (err error) {

	if err := MongoConnInit.MDB.Collection(a.CollectionName()).FindOne(context.TODO(), bson.M{"Email": email}).Decode(&a); err != nil {
		return err
	}

	return nil
}

func UpdateUser(a *User, updateData *map[string]interface{}) (err error) {

	updateBson := make(bson.M)
	for key, value := range *updateData {
		updateBson[key] = value
	}

	update := bson.M{
		"$set": updateBson,
	}

	if _, err := MongoConnInit.MDB.Collection(a.CollectionName()).UpdateOne(context.Background(), bson.M{"ID": a.ID}, update); err != nil {
		return err
	}

	return nil
}
