package models

type Book struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	AuthorID int    `json:"author_id,omitempty"`
	Author   string `json:"author,omitempty"`
	Year     int    `json:"year"`
	ISBN     string `json:"isbn"`
}
