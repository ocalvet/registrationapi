package models

// Idea containes information regarding an Idea
type Idea struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Path        string `json:"path"`
}
