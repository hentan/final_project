package models

type Author struct {
	ID            int    `json:"id"`
	NameAuthor    string `json:"name_author"`
	SirnameAuthor string `json:"sirname_author"`
	Biography     string `json:"biography"`
	Birthday      string `json:"birthday"`
}
