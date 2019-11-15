package handlers

import (
	"2019_2_Next_Level/internal/serverapi/log"
	hr "2019_2_Next_Level/internal/serverapi/server/Error/httpError"
	mailbox "2019_2_Next_Level/internal/serverapi/server/MailBox"
	"2019_2_Next_Level/internal/serverapi/server/MailBox/models"
	"2019_2_Next_Level/pkg/HttpTools"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type MailHandler struct {
	usecase mailbox.MailBoxUseCase
	resp    *HttpTools.Response
}

// NewMailHandler : sets handlers for specified routes (prefix = "/mail")
func NewMailHandler(router *mux.Router, usecase mailbox.MailBoxUseCase) {
	handler := MailHandler{usecase: usecase}
	handler.resp = (&HttpTools.Response{}).SetError(hr.DefaultResponse)

	router.HandleFunc("/send", handler.SendMail).Methods("POST")
	router.HandleFunc("/getByPage", handler.GetMailList).Methods("GET")
	router.HandleFunc("/get", handler.GetEmail).Methods("GET")
	router.HandleFunc("/getUnreadCount", handler.GetUnreadCount).Methods("GET")
	router.HandleFunc("/read", handler.MarkMailRead).Methods("POST")
	router.HandleFunc("/unread", handler.MarkMailUnRead).Methods("POST")
	router.HandleFunc("/remove", handler.DeleteEmail).Methods("POST")

}

func (h *MailHandler) SendMail(w http.ResponseWriter, r *http.Request) {
	resp := h.resp.SetWriter(w).Copy()
	defer resp.Send()
	login := h.getLogin(r)
	if login=="" {
		resp.SetError(hr.GetError(hr.BadSession))
		return
	}
	mail := models.MailToSend{}
	req := struct{
		Message models.MailToSend `json:"message"`
	}{mail}
	err := HttpTools.StructFromBody(*r, &req)
	mail = req.Message
	if err !=nil {
		log.Log().E(err)
		resp.SetError(hr.GetError(hr.BadParam))
		return
	}
	email := mail.ToMain()
	email.SetFrom(login)
	err = h.usecase.SendMail(&email)
	if err != nil {
		log.Log().E("Cannot send email")
		resp.SetError(hr.GetError(hr.UnknownError))
	}
}

func (h *MailHandler) GetMailList(w http.ResponseWriter, r *http.Request) {
	resp := h.resp.SetWriter(w).Copy()
	defer resp.Send()
	login := h.getLogin(r)
	if login=="" {
		resp.SetError(hr.GetError(hr.BadSession))
		return
	}
	pageTemp := r.FormValue("page")
	page, err := strconv.ParseInt(pageTemp, 10, 8)
	if err != nil {
		resp.SetError(hr.GetError(hr.BadParam))
		return
	}
	perPage, err := strconv.ParseInt(r.FormValue("perPage"), 10, 8)
	if err != nil {
		resp.SetError(hr.GetError(hr.BadParam))
		return
	}
	folder := r.FormValue("folder")

	//count, page, list, err := h.usecase.GetMailListPlain(login, int(pg))
	startLetter := perPage*(page-1)+1
	list, err := h.usecase.GetMailList(login, folder, "", int(startLetter), int(perPage))
	if err != nil {
		resp.SetError(hr.BadParam)
		return
	}
	resp.SetAnswer(struct{
		Status string `json:"status"`
		PagesCount int `json:"pagesCount"`
		Page int `json:"page"`
		Messages []models.MailToGet `json:"messages"`
	}{
		Status:"ok",
		PagesCount:10,
		Page: int(page),
		Messages: func()[]models.MailToGet{
			localList := make([]models.MailToGet, 0, len(list))
			for _, i := range list {
				localList = append(localList, models.MailToGet{}.FromMain(&i))
			}
			return localList
		}(),
	})
}

func (h *MailHandler) GetUnreadCount(w http.ResponseWriter, r *http.Request) {
	resp := h.resp.SetWriter(w).Copy()
	defer resp.Send()
	login := h.getLogin(r)
	if login=="" {
		resp.SetError(hr.GetError(hr.BadSession))
		return
	}
	count, err := h.usecase.GetUnreadCount(login)
	if err != nil {
		resp.SetError(hr.GetError(hr.UnknownError))
		log.Log().E(err)
		return
	}
	resp.SetAnswer(struct{
		Status string
		Count int
	}{Status:"ok", Count:count})
}

func (h *MailHandler) GetEmail(w http.ResponseWriter, r *http.Request) {
	resp := h.resp.SetWriter(w).Copy()
	defer resp.Send()
	login := h.getLogin(r)
	if login=="" {
		resp.SetError(hr.GetError(hr.BadSession))
		return
	}
	idTemp := r.FormValue("message")
	id, err := strconv.ParseInt(idTemp, 10, 8)
	if err != nil {
		resp.SetError(hr.GetError(hr.BadParam))
		return
	}
	//id := struct{
	//	Id models.MailID
	//}{}
	//err := HttpTools.StructFromBody(*r, &id)
	//if err != nil {
	//	resp.SetError(hr.GetError(hr.BadSession))
	//	return
	//}
	mail, err := h.usecase.GetMail(login, models.MailID(id))
	if err != nil {
		resp.SetError(hr.GetError(hr.BadParam))
		return
	}

	answer := struct {
		Status string `json:"status"`
		Message models.MailToGet `json:"message"`
	}{
		Status: "ok",
		Message:models.MailToGet{
			Id: models.MailID(id),
			From: models.Sender{
				Name:  "",
				Email: mail.From,
			},
			Subject:mail.Header.Subject,
			Content:mail.Body,
			Replies:[]models.MailID{},
		},
	}

	resp.SetAnswer(answer)
}

func (h *MailHandler) MarkMailRead(w http.ResponseWriter, r *http.Request) {
	h.markMail(w, r, models.MarkMessageRead)
}

func (h *MailHandler) MarkMailUnRead(w http.ResponseWriter, r *http.Request) {
	h.markMail(w, r, models.MarkMessageUnread)
}

func (h *MailHandler) DeleteEmail(w http.ResponseWriter, r *http.Request) {
	h.markMail(w, r, models.MarkMessageDeleted)
}

func (h *MailHandler) markMail(w http.ResponseWriter, r *http.Request, mark int) {
	resp := h.resp.SetWriter(w).Copy()
	defer resp.Send()
	login := h.getLogin(r)
	if login=="" {
		resp.SetError(hr.GetError(hr.BadSession))
		return
	}

	req := struct {
		Messages []models.MailID
	}{}
	err := HttpTools.StructFromBody(*r, &req)
	if err != nil {
		resp.SetError(hr.GetError(hr.BadSession))
		return
	}
	err = h.usecase.MarkMail(login, req.Messages, mark)
	if err != nil {
		resp.SetError(hr.GetError(hr.UnknownError))
		return
	}
}

func (h *MailHandler) getLogin(r *http.Request) string {
	return r.Header.Get("X-Login")
}