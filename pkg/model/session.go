package model

type Session struct {
	ID     string `json:"id" bson:"_id,omitempty"`
	UserID string `json:"user_id" bson:"user_id,omitempty"`
}
