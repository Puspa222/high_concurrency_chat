package models

type Message struct {
	Sender    string `json:"sender"`
	Content   string `json:"content"`
	Room      string `json:"room"`
	Timestamp int64  `json:"timestamp"`
}
