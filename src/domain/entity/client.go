package entity

type Client interface {
	ReadMessage() (message []byte, err error)
	SendMessage(message string) error
	Close() error
}
