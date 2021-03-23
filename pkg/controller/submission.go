package controller

import (
	"context"
	"time"

	"github.com/ZdravkoGyurov/go-grader/pkg/dexec"
	"github.com/ZdravkoGyurov/go-grader/pkg/log"
	"github.com/ZdravkoGyurov/go-grader/pkg/model"
	"github.com/ZdravkoGyurov/go-grader/pkg/random"
	"github.com/google/uuid"
)

func (c *Controller) CreateSubmission(ctx context.Context, submission *model.Submission) error {
	submission.ID = uuid.NewString()
	submission.Result = ""
	submission.Status = "PENDING"

	if err := c.storage.CreateSubmission(ctx, submission); err != nil {
		return err
	}

	jobName := "run tests in docker"
	jobFunc := func() {
		testsConfig := dexec.TestsRunConfig{
			ImageName:       random.LowercaseString(10),
			ContainerName:   random.LowercaseString(10),
			Assignment:      "assignment1",             // get assignmentID from body and get assignmentName from db
			SolutionGitUser: "ZdravkoGyurov",           // get userID from session and get gitUsername from db
			SolutionGitRepo: "grader-docker-solutions", // get assignmentID from body and get gitCourseName from db
			TestsGitUser:    c.Config.TestsGitUser,
			TestsGitRepo:    c.Config.TestsGitRepo,
		}
		output, err := dexec.RunTests(testsConfig)
		if err != nil {
			log.Error().Println(err) // log status in db
			log.Error().Println(output)
			return
		}

		submission.Result = output
		submission.Status = "DONE"
		updateCtx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()
		_, err = c.UpdateSubmission(updateCtx, submission.ID, submission)
		if err != nil {
			log.Error().Println(err)
		}
	}
	_, err := c.executor.QueueJob(jobName, jobFunc)
	if err != nil {
		return err
	}

	return nil
}

func (c *Controller) GetAllSubmissions(ctx context.Context, userID, assignmentID string) ([]*model.Submission, error) {
	submissions, err := c.storage.ReadAllSubmissions(ctx, userID, assignmentID)
	if err != nil {
		return nil, err
	}

	return submissions, nil
}

func (c *Controller) GetSubmission(ctx context.Context, submissionID string) (*model.Submission, error) {
	submission, err := c.storage.ReadSubmission(ctx, submissionID)
	if err != nil {
		return nil, err
	}

	return submission, nil
}

func (c *Controller) UpdateSubmission(ctx context.Context, submissionID string, submission *model.Submission) (*model.Submission, error) {
	updatedSubmission, err := c.storage.UpdateSubmission(ctx, submissionID, submission)
	if err != nil {
		return nil, err
	}

	return updatedSubmission, nil
}

func (c *Controller) DeleteSubmission(ctx context.Context, submissionID string) error {
	if err := c.storage.DeleteSubmission(ctx, submissionID); err != nil {
		return err
	}

	return nil
}
