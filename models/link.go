package model

type Link struct {
	ID		int		`json:"id, omitempty"`
	Title	string	`json:"title"`
	Url		string	`json:"url"`
}
