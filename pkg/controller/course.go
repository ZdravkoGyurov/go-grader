package controller

import (
	"context"

	"github.com/ZdravkoGyurov/go-grader/pkg/model"
	"github.com/google/uuid"
)

func (c *Controller) CreateCourse(ctx context.Context, course *model.Course) error {
	course.ID = uuid.NewString()

	if err := c.storage.CreateCourse(ctx, course); err != nil {
		return err
	}

	return nil
}

func (c *Controller) GetCourse(ctx context.Context, courseID string) (*model.Course, error) {
	course, err := c.storage.ReadCourse(ctx, courseID)
	if err != nil {
		return nil, err
	}

	return course, nil
}

func (c *Controller) GetAllCourses(ctx context.Context) ([]*model.Course, error) {
	courses, err := c.storage.ReadAllCourses(ctx)
	if err != nil {
		return nil, err
	}

	return courses, nil
}

func (c *Controller) UpdateCourse(ctx context.Context, courseID string, course *model.Course) (*model.Course, error) {
	updatedCourse, err := c.storage.UpdateCourse(ctx, courseID, course)
	if err != nil {
		return nil, err
	}

	return updatedCourse, nil
}

func (c *Controller) DeleteCourse(ctx context.Context, courseID string) error {
	if err := c.storage.DeleteCourse(ctx, courseID); err != nil {
		return err
	}

	return nil
}
