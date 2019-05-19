package viewModels

type HelloWorld struct {
	Message string `json:"message"`
	Status  string `json:"status"`
}

func (hw *HelloWorld) GetMessage() string {
	return hw.Message
}

func (hw *HelloWorld) GetStatus() string {
	return hw.Status
}

func (hw *HelloWorld) SetMessage(message string) {
	hw.Message = message
}

func (hw *HelloWorld) SetStatus(status string) {
	hw.Status = status
}

func GenHelloWorldViewModel(message string, status string) *HelloWorld {
	hw := new(HelloWorld)
	hw.SetStatus(status)
	hw.SetMessage(message)
	return hw
}
