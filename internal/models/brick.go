package models

type Brick struct {
	ID   string `json:"id"`
	Type string `json:"type"`
	Payload string `json:"payload"`
	Params string `json:"params"`
	Children []string `json:"children"`
}