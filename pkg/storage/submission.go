package storage

import (
	"context"

	"github.com/ZdravkoGyurov/go-grader/pkg/errors"
	"github.com/ZdravkoGyurov/go-grader/pkg/model"
)

const submissionCollection = "submissions"

func (s *Storage) CreateSubmission(ctx context.Context, submission *model.Submission) error {
	collection := s.mongoClient.Database(s.config.DatabaseName).Collection(submissionCollection)
	_, err := collection.InsertOne(ctx, submission)
	if err != nil {
		return errors.Wrap(storageError(err), "failed to insert submission")
	}

	return nil
}

func (s *Storage) ReadSubmission(ctx context.Context, submissionID string) (*model.Submission, error) {
	collection := s.mongoClient.Database(s.config.DatabaseName).Collection(submissionCollection)
	var submission model.Submission
	if err := collection.FindOne(ctx, filterByID(submissionID)).Decode(&submission); err != nil {
		return nil, errors.Wrapf(storageError(err), "failed to find submission with id %s", submissionID)
	}

	return &submission, nil
}

// ReadAllSubmissions reads all submissions, filtered by userID and/or assignmentID.
// UserID and assignmentID cannot be both empty.
func (s *Storage) ReadAllSubmissions(ctx context.Context, userID, assignmentID string) ([]*model.Submission, error) {
	collection := s.mongoClient.Database(s.config.DatabaseName).Collection(submissionCollection)

	filter, err := filterSubmissions(userID, assignmentID)
	if err != nil {
		err = errors.Wrapf(err, "failed to create submissions filter with user_id %s and assignment_id %s", userID, assignmentID)
		return nil, err
	}

	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		err = errors.Wrapf(storageError(err), "failed to find all submissions with user_id %s and assignment_id %s", userID, assignmentID)
		return nil, err
	}

	var submissions []*model.Submission
	if err = cursor.All(ctx, &submissions); err != nil {
		return nil, errors.Wrap(storageError(err), "failed to decode all submissions")
	}

	return submissions, nil
}

func (s *Storage) UpdateSubmission(ctx context.Context, submissionID string, submission *model.Submission) (*model.Submission, error) {
	collection := s.mongoClient.Database(s.config.DatabaseName).Collection(submissionCollection)
	var updatedSubmission model.Submission
	result := collection.FindOneAndUpdate(ctx, filterByID(submissionID), update(submission), updateOpts())
	if err := result.Decode(&updatedSubmission); err != nil {
		err = errors.Wrapf(storageError(err), "failed to find and update submission with id %s", submissionID)
		return nil, err
	}

	return &updatedSubmission, nil
}

func (s *Storage) DeleteSubmission(ctx context.Context, submissionID string) error {
	collection := s.mongoClient.Database(s.config.DatabaseName).Collection(submissionCollection)
	if _, err := collection.DeleteOne(ctx, filterByID(submissionID)); err != nil {
		return errors.Wrapf(storageError(err), "failed to delete submission with id %s", submissionID)
	}
	return nil
}
