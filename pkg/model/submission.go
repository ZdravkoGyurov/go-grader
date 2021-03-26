package model

type Submission struct {
	ID           string `json:"id" bson:"_id,omitempty"`
	Status       string `json:"status" bson:"status,omitempty"`
	Result       string `json:"result" bson:"result,omitempty"`
	UserID       string `json:"user_id" bson:"user_id,omitempty"`
	AssignmentID string `json:"assignment_id" bson:"assignment_id,omitempty"`
}
