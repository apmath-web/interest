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

func Unique(intSlice []int) []int {
	keys := make(map[int]bool)
	list := []int{}
	for _, entry := range intSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
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

func (v *Validation) unique(intSlice []int) []int {
	keys := make(map[int]bool)
	list := []int{}
	for _, entry := range intSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}
