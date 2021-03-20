package storage

import (
	"context"
	"fmt"

	"github.com/ZdravkoGyurov/go-grader/pkg/model"
)

const userCollection = "users"

func (s *Storage) CreateUser(ctx context.Context, user *model.User) error {
	collection := s.mongoClient.Database(s.config.DatabaseName).Collection(userCollection)
	_, err := collection.InsertOne(ctx, user)
	if err != nil {
		return fmt.Errorf("failed to insert user: %w", err)
	}

	return nil
}

func (s *Storage) ReadUserByID(ctx context.Context, userID string) (*model.User, error) {
	collection := s.mongoClient.Database(s.config.DatabaseName).Collection(userCollection)
	var user model.User
	if err := collection.FindOne(ctx, filterByID(userID)).Decode(&user); err != nil {
		return nil, fmt.Errorf("failed to find user with id %s: %w", userID, err)
	}

	return &user, nil
}

func (s *Storage) ReadUserByUsername(ctx context.Context, username string) (*model.User, error) {
	collection := s.mongoClient.Database(s.config.DatabaseName).Collection(userCollection)
	var user model.User

	result := collection.FindOne(ctx, filterByUsername(username))
	if err := result.Decode(&user); err != nil {
		return nil, fmt.Errorf("failed to find user with username %s: %w", username, err)
	}

	return &user, nil
}

func (s *Storage) ReadAllUsers(ctx context.Context, courseID string) ([]*model.User, error) {
	collection := s.mongoClient.Database(s.config.DatabaseName).Collection(userCollection)

	cursor, err := collection.Find(ctx, filterUsersByCourseID(courseID))
	if err != nil {
		return nil, fmt.Errorf("failed to find all users with course_id %s: %w", courseID, err)
	}

	var users []*model.User
	if err = cursor.All(ctx, &users); err != nil {
		return nil, fmt.Errorf("failed to decode all users with course_id %s: %w", courseID, err)
	}

	return users, nil
}

func (s *Storage) UpdateUser(ctx context.Context, userID string, user *model.User) (*model.User, error) {
	collection := s.mongoClient.Database(s.config.DatabaseName).Collection(userCollection)
	var updatedUser model.User
	result := collection.FindOneAndUpdate(ctx, filterByID(userID), update(user), updateOpts())
	if err := result.Decode(&updatedUser); err != nil {
		return nil, fmt.Errorf("failed to find and update user with id %s: %w", userID, err)
	}

	return &updatedUser, nil
}

func (s *Storage) DeleteUser(ctx context.Context, userID string) error {
	collection := s.mongoClient.Database(s.config.DatabaseName).Collection(userCollection)
	if _, err := collection.DeleteOne(ctx, filterByID(userID)); err != nil {
		return fmt.Errorf("failed to delete user with id %s: %w", userID, err)
	}
	return nil
}
