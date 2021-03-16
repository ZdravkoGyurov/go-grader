package model

type Session struct {
	ID     string `json:"id" bson:"_id,omitempty"`
	UserID string `json:"userId" bson:"userId,omitempty"`
}
