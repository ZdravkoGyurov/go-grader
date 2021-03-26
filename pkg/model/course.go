package model

type Course struct {
	ID             string `json:"id" bson:"_id,omitempty"`
	Name           string `json:"name" bson:"name,omitempty"`
	Description    string `json:"description" bson:"description,omitempty"`
	GithubRepoName string `json:"github_repo_name" bson:"github_repo_name,omitempty"`
}
