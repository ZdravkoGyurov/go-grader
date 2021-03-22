package model

import "time"

type Assignment struct {
	ID            string    `json:"id" bson:"_id,omitempty"`
	Name          string    `json:"name" bson:"name,omitempty"`
	Description   string    `json:"description" bson:"description,omitempty"`
	CreatedOn     time.Time `json:"created_on" bson:"created_on,omitempty"`
	LastUpdatedOn time.Time `json:"last_updated_on" bson:"last_updated_on,omitempty"`
	DueDate       time.Time `json:"due_date" bson:"due_date,omitempty"`
	CourseID      string    `json:"course_id" bson:"course_id,omitempty"`
}
