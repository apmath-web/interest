package Validation

import (
	"encoding/json"
	"github.com/apmath-web/interests/Domain"
)

type Validation struct {
	message  string
	messages []Domain.MessageInterface
}

func (v *Validation) AddMessage(message Domain.MessageInterface) {
	v.messages = append(v.messages, message)
}

func GenMessage(field, text string) *Message {
	m := new(Message)
	m.field = field
	m.text = text
	return m
}

type Message struct {
	text, field string
}

func (m *Message) MarshalJSON() (b []byte, e error) {
	return json.Marshal(map[string]string{
		m.field: m.text,
	})
}

func (v *Validation) SetMessage(msg string) {
	v.message = msg
}

func (v *Validation) Empty() bool {
	return (len(v.messages) == 0) && (v.message == "")
}

func GenValidation() Domain.ValidationInterface {
	v := new(Validation)
	return v
}

func (v *Validation) GetMessages() []Domain.MessageInterface {
	return v.messages
}

func (v *Validation) MarshalJSON() (b []byte, e error) {
	return json.Marshal(map[string]interface{}{
		"message":     v.message,
		"description": v.messages,
	})
}

func (m *Message) GetStringMessage() string {
	return m.field + ": " + m.text
}
