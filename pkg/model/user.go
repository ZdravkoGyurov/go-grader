package model

type User struct {
	ID             string   `json:"id" bson:"_id,omitempty"`
	Username       string   `json:"username" bson:"username,omitempty"`
	GithubUsername string   `json:"github_username" bson:"github_username,omitempty"`
	Fullname       string   `json:"fullname" bson:"fullname,omitempty"`
	Password       string   `json:"password" bson:"password,omitempty"`
	Permissions    []string `json:"permissions" bson:"permissions,omitempty"`
	Disabled       bool     `json:"disabled" bson:"disabled,omitempty"`
	CourseIDs      []string `json:"course_ids" bson:"course_ids,omitempty"`
}
