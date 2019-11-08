package models

import "2019_2_Next_Level/internal/model"

type MailID int64

type MailToSend struct {
	To []string `json:"to"`
	Subject string `json:"subject"`
	ReplyTo []string `json:"reply_to,omitempty"`
	Content string `json:"content"`
}

func (m MailToSend) ToMain() model.Email {
	to := model.Email{}
	to.Header.To = m.To
	to.Header.Subject = m.Subject
	to.Body = m.Content
	if len(m.To) > 0{
		to.To = m.To[0]
	}else{
		to.To = ""
	}
	return to
}


type MailToGet struct {
	Id MailID`json:"id"`
	From Sender `json:"from"`
	Subject string `json:"subject"`
	Content string `json:"content"`
	Replies []MailID`json:"replies,omitempty"`
}
func (m MailToGet) FromMain(from *model.Email) MailToGet {
	m.Id = MailID(0)
	m.From = Sender{Name:"", Email:from.From}
	m.Subject = from.Header.Subject
	m.Content = from.Body
	m.Replies = []MailID{}

}

type Sender struct {
	Name string	`json:"name"`
	Email string `json:"email"`
}
