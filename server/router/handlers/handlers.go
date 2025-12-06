package handlers

import (
	"errors"
	"fmt"
	"net/http"

	"tchat.com/server/modules/messages"
	"tchat.com/server/modules/users"
	"tchat.com/server/store"
	"tchat.com/server/utils"
)

type Handlers struct {
	store store.Store
}

func NewHandler(store store.Store) *Handlers {
	return &Handlers{store}
}

type SendMessageBody struct {
	Body       utils.MessageBody `json:"body"`
	SenderID   utils.UserID      `json:"sent_by"`
	ReceiverID utils.UserID      `json:"sent_to"`
}

func (s SendMessageBody) Validate() error {
	err1 := s.Body.Validate()
	err2 := s.SenderID.Validate()
	err3 := s.ReceiverID.Validate()

	return errors.Join(err1, err2, err3)
}

func (h *Handlers) SendMessage(w http.ResponseWriter, r *http.Request) {
	b := SendMessageBody{}

	if err := ReadBody(r, &b); err != nil {
		WriteInternalServerError(w, err)
		return
	}

	if err := b.Validate(); err != nil {
		WriteBadRequest(w, err)
		return
	}

	sender, err := h.store.FindUserByID(b.SenderID)
	if err != nil {
		WriteNotFound(w, "Sender not found", err)
		return
	}

	receiver, err := h.store.FindUserByID(b.ReceiverID)
	if err != nil {
		WriteNotFound(w, "Receiver not found", err)
		return
	}

	m := messages.New(b.Body, sender, receiver)

	err = h.store.SendMessage(m)
	if err != nil {
		WriteInternalServerError(w, err)
		return
	}

	WriteOKEmpty(w, "Message sent")
}

type ReadChatBody struct {
	User1 utils.UserID `json:"user_1"`
	User2 utils.UserID `json:"user_2"`
}

func (r ReadChatBody) Validate() error {
	err1 := r.User1.Validate()
	err2 := r.User2.Validate()

	return errors.Join(err1, err2)
}

func (h *Handlers) ReadChat(w http.ResponseWriter, r *http.Request) {
	b := ReadChatBody{}

	if err := ReadBody(r, &b); err != nil {
		WriteInternalServerError(w, err)
		return
	}

	if err := b.Validate(); err != nil {
		WriteBadRequest(w, err)
		return
	}

	u1, err := h.store.FindUserByID(b.User1)
	if err != nil {
		WriteNotFound(w, fmt.Sprintf("User not found with id: %v", b.User1), err)
		return
	}

	u2, err := h.store.FindUserByID(b.User2)
	if err != nil {
		WriteNotFound(w, fmt.Sprintf("User not found with id: %v", b.User2), err)
		return
	}

	ms, err := h.store.ReadChat(u1, u2)
	if err != nil {
		WriteInternalServerError(w, err)
		return
	}

	WriteOKWithBody(w, ms)
}

type SaveUserBody struct {
	Name utils.UserName `json:"name"`
}

func (s SaveUserBody) Validate() error {
	err := s.Name.Validate()

	return err
}

func (h *Handlers) SaveUser(w http.ResponseWriter, r *http.Request) {
	b := SaveUserBody{}

	if err := ReadBody(r, &b); err != nil {
		WriteInternalServerError(w, err)
		return
	}

	if err := b.Validate(); err != nil {
		WriteBadRequest(w, err)
		return
	}

	u := users.New(b.Name)

	err := h.store.SaveUser(u)
	if err != nil {
		WriteInternalServerError(w, err)
		return
	}

	WriteOKWithBody(w, u)
}

type FindUserByIDBody struct {
	UserID utils.UserID `json:"user_id"`
}

func (f FindUserByIDBody) Validate() error {
	err := f.UserID.Validate()

	return err
}

func (h *Handlers) FindUserByID(w http.ResponseWriter, r *http.Request) {
	b := FindUserByIDBody{}

	if err := ReadBody(r, &b); err != nil {
		WriteInternalServerError(w, err)
		return
	}

	if err := b.Validate(); err != nil {
		WriteBadRequest(w, err)
		return
	}

	u, err := h.store.FindUserByID(b.UserID)
	if err != nil {
		WriteInternalServerError(w, err)
		return
	}

	WriteOKWithBody(w, u)
}
