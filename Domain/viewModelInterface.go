package Domain

type HelloWorldViewModel interface {
	GetMessage() string
	SetMessage(message string)
	GetStatus() string
	SetStatus(status string)
}
