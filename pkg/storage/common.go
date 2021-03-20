package storage

import (
	"errors"

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

func filterAssignmentByCourseID(courseID string) bson.M {
	return bson.M{"course_id": courseID}
}

func filterSubmissions(userID, assignmentID string) (bson.M, error) {
	if userID != "" && assignmentID != "" {
		return bson.M{"user_id": userID, "assignment_id": assignmentID}, nil
	}
	if userID != "" {
		return bson.M{"user_id": userID}, nil
	}
	if assignmentID != "" {
		return bson.M{"assignment_id": assignmentID}, nil
	}
	return bson.M{}, errors.New("pass user_id and/or assignment_id to filter submissions")
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
