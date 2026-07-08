package models

//to get, delete and update data
type Posts struct {
	ID       string   `json:"id"`
	Title    string   `json:"title"`
	Content  string   `json:"content"`
	Category string   `json:"category"`
	Tags     []string `json:"tags"`
}

type PostFilter struct {
	Query    string
	Category string
	Tag      string
}
