package models

// Team Is there a team?
type Team struct {
	Name    string   `json:"name"`
	Logo    string   `json:"logo"`
	Members []string `json:"members"`
}
