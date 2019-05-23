package Domain

type ValidationInterface interface {
	AddMessage(msg MessageInterface)
	GetMessages() []MessageInterface
	SetMessage(msg string)
	MarshalJSON() (b []byte, e error)
	Empty() bool
}

type MessageInterface interface {
	MarshalJSON() (b []byte, e error)
}
