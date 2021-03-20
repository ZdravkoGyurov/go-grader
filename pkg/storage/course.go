package storage

import (
	"context"
	"fmt"

	"github.com/ZdravkoGyurov/go-grader/pkg/model"
	"go.mongodb.org/mongo-driver/bson"
)

const courseCollection = "courses"

func (s *Storage) CreateCourse(ctx context.Context, course *model.Course) error {
	collection := s.mongoClient.Database(s.config.DatabaseName).Collection(courseCollection)
	_, err := collection.InsertOne(ctx, course)
	if err != nil {
		return fmt.Errorf("failed to insert collection: %w", err)
	}

	return nil
}

func (s *Storage) ReadCourse(ctx context.Context, courseID string) (*model.Course, error) {
	collection := s.mongoClient.Database(s.config.DatabaseName).Collection(courseCollection)
	var course model.Course
	if err := collection.FindOne(ctx, filterByID(courseID)).Decode(&course); err != nil {
		return nil, fmt.Errorf("failed to find course with id %s: %w", courseID, err)
	}

	return &course, nil
}

func (s *Storage) ReadAllCourses(ctx context.Context) ([]*model.Course, error) {
	collection := s.mongoClient.Database(s.config.DatabaseName).Collection(courseCollection)

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, fmt.Errorf("failed to find all courses: %w", err)
	}

	var courses []*model.Course
	if err = cursor.All(ctx, &courses); err != nil {
		return nil, fmt.Errorf("failed to decode all courses: %w", err)
	}

	return courses, nil
}

func (s *Storage) UpdateCourse(ctx context.Context, courseID string, course *model.Course) (*model.Course, error) {
	collection := s.mongoClient.Database(s.config.DatabaseName).Collection(courseCollection)
	var updatedCourse model.Course
	result := collection.FindOneAndUpdate(ctx, filterByID(courseID), update(course), updateOpts())
	if err := result.Decode(&updatedCourse); err != nil {
		return nil, fmt.Errorf("failed to find and update course with id %s: %w", courseID, err)
	}

	return &updatedCourse, nil
}

func (s *Storage) DeleteCourse(ctx context.Context, courseID string) error {
	collection := s.mongoClient.Database(s.config.DatabaseName).Collection(courseCollection)
	if _, err := collection.DeleteOne(ctx, filterByID(courseID)); err != nil {
		return fmt.Errorf("failed to delete course with id %s: %w", courseID, err)
	}
	return nil
}
