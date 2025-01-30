package entity

type ClientInterface interface {
	ReadMessage() (message []byte, err error)
	SendMessage(message string) error
	Close() error
}
