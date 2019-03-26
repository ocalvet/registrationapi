package models

// Registration containes information regarding an Registration
type Registration struct {
	ID          string `json:"id"`
	Participant User   `json:"user"`
	Idea        Idea   `json:"idea"`
	Team        Team   `json:"team"`
}
