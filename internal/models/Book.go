package models

type Book struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	AuthorId int    `json:"author_id"`
	Year     int    `json:"year"`
	ISBN     string `json:"isbn"`
}
