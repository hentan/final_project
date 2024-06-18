package models

type AuthorAndBook struct {
	Author Author `json:"author"`
	Book   BookID `json:"book"`
}
