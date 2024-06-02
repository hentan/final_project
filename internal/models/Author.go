package models

type Author struct {
	ID            int    `json:"id"`
	NameAuthor    string `json:"name_author"`
	SirnameAuthor string `json:"sirname_author"`
	Biography     string `json:"biography"`
	Birthday      string `json:"birthday"`
}

type AuthorAndBook struct {
	NameAuthor    string `json:"name_author"`
	SirnameAuthor string `json:"sirname_author"`
	Biography     string `json:"biography"`
	Birthday      string `json:"birthday"`
	Title         string `json:"title"`
	AuthorID      int    `json:"author_id"`
	Year          int    `json:"year"`
	ISBN          string `json:"isbn"`
}
