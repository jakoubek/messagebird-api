package messagebird_api

type Sms struct {
	Originator string `json:"originator"`
	MessageBody string `json:"body"`
	Recipients []string `json:"recipients"`
}

func (s *Sms) AddRecipient(recipient string) {
	s.Recipients = append(s.Recipients, recipient)
}

func (s *Sms) SetMessageBody(messageBody string) {
	s.MessageBody = messageBody
}

func (s *Sms) SetOriginator(originator string) {
	s.Originator = originator
}

func NewSms() *Sms {
	return &Sms{
		Originator:  "",
		MessageBody: "",
		Recipients:  nil,
	}
}

