package inbound

type Event struct {
	Type    string      `json:"type"`
	Payload interface{} `json:"payload"`
}
