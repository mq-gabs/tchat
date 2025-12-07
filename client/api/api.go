package api

import (
	"errors"
	"net/http"
	"time"

	"tchat.com/server/modules/messages"
	"tchat.com/server/router"
	"tchat.com/server/router/handlers"
)

var (
	errCannotCreateRequest = errors.New("cannot create request")
	errRequestFailed       = errors.New("request failed")
)

type TChatAPI struct {
	client *http.Client
	host   string
}

func NewTChatAPI(host string) *TChatAPI {
	return &TChatAPI{
		client: &http.Client{
			Timeout: 5 * time.Second,
		},
		host: host,
	}
}

func (api *TChatAPI) do(req *http.Request) (*http.Response, error) {
	resp, err := api.client.Do(req)
	if err != nil {
		return nil, errors.Join(errRequestFailed, err)
	}

	return resp, nil
}

func (api *TChatAPI) ReadChat(readChatBody *handlers.ReadChatBody) ([]messages.Message, error) {
	req, err := NewGet(api.host + router.PathMessagesGet)
	if err != nil {
		return nil, err
	}

	AddQuery(req, map[string]string{
		"user_1": string(readChatBody.User1),
		"user_2": string(readChatBody.User2),
	})

	resp, err := api.do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	return ProcessResponseData[[]messages.Message](resp)
}

func (api *TChatAPI) SendMessage(sendMessageBody *handlers.SendMessageBody) error {
	req, err := NewPost(api.host+router.PathMessagesSend, sendMessageBody)
	if err != nil {
		return err
	}

	resp, err := api.do(req)
	if err != nil {
		return err
	}

	_, err = ProcessResponseData[any](resp)

	return err
}
