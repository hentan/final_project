package models

type Author struct {
	ID            int    `json:"id"`
	NameAuthor    string `json:"name_author"`
	SurnameAuthor string `json:"surname_author"`
	Biography     string `json:"biography"`
	Birthday      string `json:"birthday"`
}
