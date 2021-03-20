package storage

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func filterByID(id string) bson.M {
	return bson.M{"_id": id}
}

func filterByUsername(username string) bson.M {
	return bson.M{"username": username}
}

func filterUsersByCourseID(courseID string) bson.M {
	return bson.M{"course_ids": courseID}
}

func updateOpts() *options.FindOneAndUpdateOptions {
	returnDocumentOption := options.After
	options := &options.FindOneAndUpdateOptions{
		ReturnDocument: &returnDocumentOption,
	}
	return options
}

func update(entity interface{}) bson.M {
	return bson.M{"$set": entity}
}
