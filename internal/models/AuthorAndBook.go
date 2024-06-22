package models

type AuthorAndBook struct {
	Author Author `json:"author"`
	Book   Book   `json:"book"`
}
