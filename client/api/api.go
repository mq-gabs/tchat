package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
	"tchat.com/server/modules/messages"
	"tchat.com/server/modules/users"
	"tchat.com/server/router"
	"tchat.com/server/router/handlers"
	"tchat.com/server/utils"
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

func (api *TChatAPI) httpHost() string {
	return "http://" + api.host
}

func (api *TChatAPI) wsHost() string {
	return "ws://" + api.host
}

func (api *TChatAPI) do(req *http.Request) (*http.Response, error) {
	resp, err := api.client.Do(req)
	if err != nil {
		return nil, errors.Join(errRequestFailed, err)
	}

	return resp, nil
}

func (api *TChatAPI) Ping() error {
	req, err := NewGet(api.httpHost() + router.PathPing)
	if err != nil {
		return err
	}

	res, err := api.do(req)
	if err != nil {
		return err
	}

	_, err = ProcessResponseData[any](res)

	return err
}

func (api *TChatAPI) ReadChat(readChatQuery *handlers.ReadChatQuery) ([]messages.Message, error) {
	req, err := NewGet(api.httpHost() + router.PathMessagesGet)
	if err != nil {
		return nil, err
	}

	AddQuery(req, map[string]string{
		"user_1": string(readChatQuery.User1),
		"user_2": string(readChatQuery.User2),
	})

	resp, err := api.do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	return ProcessResponseData[[]messages.Message](resp)
}

func (api *TChatAPI) SendMessage(sendMessageBody *handlers.SendMessageBody) error {
	req, err := NewPost(api.httpHost()+router.PathMessagesSend, sendMessageBody)
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

func (api *TChatAPI) SaveUser(saveUserBody *handlers.SaveUserBody) (*users.User, error) {
	req, err := NewPost(api.httpHost()+router.PathUsersSave, saveUserBody)
	if err != nil {
		return nil, err
	}

	resp, err := api.do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	user, err := ProcessResponseData[*users.User](resp)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (api *TChatAPI) FindUserByID(findUserByIDQuery *handlers.FindUserByIDQuery) (*users.User, error) {
	req, err := NewGet(api.httpHost() + router.PathUsersSave)
	if err != nil {
		return nil, err
	}

	AddQuery(req, map[string]string{
		"user_id": string(findUserByIDQuery.UserID),
	})

	resp, err := api.do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	user, err := ProcessResponseData[users.User](resp)
	if err != nil {
		return nil, err
	}

	return &user, nil

}

func (api *TChatAPI) WebsocketChat(mergedIds utils.MergedIDs) (chan *messages.Message, error) {
	url := api.wsHost() + router.PathWebsocketChatBase + "/" + string(mergedIds)
	fmt.Println(url)
	conn, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		return nil, errors.Join(errCannotConnectToWebsocket, err)
	}

	newMsgs := make(chan *messages.Message)

	go func() {
		defer conn.Close()
		defer close(newMsgs)

		for {
			_, reader, err := conn.NextReader()
			if err != nil {
				fmt.Printf("cannot read websocket message\n")
				return
			}

			var msg messages.Message
			if err := json.NewDecoder(reader).Decode(&msg); err != nil {
				fmt.Printf("cannot decode websocket message")
				return
			}

			newMsgs <- &msg
		}
	}()

	return newMsgs, nil
}
