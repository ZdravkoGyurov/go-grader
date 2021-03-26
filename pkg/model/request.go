package model

type Request struct {
	ID          string   `json:"id" bson:"_id,omitempty"`
	Name        string   `json:"name" bson:"name,omitempty"`
	Description string   `json:"description" bson:"description,omitempty"`
	Status      string   `json:"status" bson:"status,omitempty"`
	Type        string   `json:"type" bson:"type,omitempty"`
	Permissions []string `json:"permissions" bson:"permissions,omitempty"`
	CourseID    string   `json:"course_id" bson:"course_id,omitempty"`
	UserID      string   `json:"user_id" bson:"user_id,omitempty"`
}
