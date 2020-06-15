package entity

type Post struct {
	ID          int64  `json:"ID"`
	Title       string `json:"Title"`
	Description string `json:"Description"`
}
