package utils

import "errors"

type UserID string

func (u UserID) Validate() error {
	if u == "" {
		return errors.New("user id is empty")
	}

	return nil
}

type UserName string

func (u UserName) Validate() error {
	if u == "" {
		return errors.New("user name is empty")
	}

	return nil
}

type MessageID string

func (m MessageID) Validate() error {
	if m == "" {
		return errors.New("message id is empty")
	}

	return nil
}

type MessageBody string

func (m MessageBody) Validate() error {
	if m == "" {
		return errors.New("message body is empty")
	}

	return nil
}

type ChatID string

func (m ChatID) Validate() error {
	if m == "" {
		return errors.New("chatID is empty")
	}

	return nil
}
