package models

const (
	AssignmentsCollectionName = "assignments"
)

// Assignment represents the DB document
type Assignment struct {
	ID          string `json:"id" bson:"_id"`
	Name        string `json:"name" bson:"name"`
	Description string `json:"description" bson:"description"`
}
