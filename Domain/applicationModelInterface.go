package Domain

type HelloWorldApplicationModel interface {
	GetMessage() string
	SetMessage(message string)
}
