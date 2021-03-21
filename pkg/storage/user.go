package storage

import (
	"context"

	"github.com/ZdravkoGyurov/go-grader/pkg/errors"
	"github.com/ZdravkoGyurov/go-grader/pkg/model"
)

const userCollection = "users"

func (s *Storage) CreateUser(ctx context.Context, user *model.User) error {
	collection := s.mongoClient.Database(s.config.DatabaseName).Collection(userCollection)
	_, err := collection.InsertOne(ctx, user)
	if err != nil {
		return errors.Wrap(storageError(err), "failed to insert user")
	}

	return nil
}

func (s *Storage) ReadUserByID(ctx context.Context, userID string) (*model.User, error) {
	collection := s.mongoClient.Database(s.config.DatabaseName).Collection(userCollection)
	var user model.User
	if err := collection.FindOne(ctx, filterByID(userID)).Decode(&user); err != nil {
		return nil, errors.Wrapf(storageError(err), "failed to find user with id %s", userID)
	}

	return &user, nil
}

func (s *Storage) ReadUserByUsername(ctx context.Context, username string) (*model.User, error) {
	collection := s.mongoClient.Database(s.config.DatabaseName).Collection(userCollection)
	var user model.User

	result := collection.FindOne(ctx, filterByUsername(username))
	if err := result.Decode(&user); err != nil {
		return nil, errors.Wrapf(storageError(err), "failed to find user with username %s", username)
	}

	return &user, nil
}

func (s *Storage) ReadAllUsers(ctx context.Context, courseID string) ([]*model.User, error) {
	collection := s.mongoClient.Database(s.config.DatabaseName).Collection(userCollection)

	cursor, err := collection.Find(ctx, filterUsersByCourseID(courseID))
	if err != nil {
		return nil, errors.Wrapf(storageError(err), "failed to find all users with course_id %s", courseID)
	}

	users := make([]*model.User, 0)
	if err = cursor.All(ctx, &users); err != nil {
		return nil, errors.Wrapf(storageError(err), "failed to decode all users with course_id %s", courseID)
	}

	return users, nil
}

func (s *Storage) UpdateUser(ctx context.Context, userID string, user *model.User) (*model.User, error) {
	collection := s.mongoClient.Database(s.config.DatabaseName).Collection(userCollection)
	var updatedUser model.User
	result := collection.FindOneAndUpdate(ctx, filterByID(userID), update(user), updateOpts())
	if err := result.Decode(&updatedUser); err != nil {
		return nil, errors.Wrapf(storageError(err), "failed to find and update user with id %s", userID)
	}

	return &updatedUser, nil
}

func (s *Storage) DeleteUser(ctx context.Context, userID string) error {
	collection := s.mongoClient.Database(s.config.DatabaseName).Collection(userCollection)
	if _, err := collection.DeleteOne(ctx, filterByID(userID)); err != nil {
		return errors.Wrapf(storageError(err), "failed to delete user with id %s", userID)
	}
	return nil
}
