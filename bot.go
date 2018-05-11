package telegram

import (
    "net/http"
    "errors"
    "io/ioutil"
    "encoding/json"
    "bytes"
)

const MethodGetMe = 0
const MethodSendMessage = 1
const MethodForwardMessage = 2
const MethodSendPhoto = 3
const MethodSendAudio = 4
const MethodSendDocument = 5
const MethodSendVideo = 6
const MethodSendVoice = 7
const MethodSendVideoNote = 8
const MethodSendMediaGroup = 9
const MethodSendLocation = 10
const MethodEditMessageLiveLocation = 11
const MethodStopMessageLiveLocation = 12
const MethodSendVenue = 13
const MethodSendContact = 14
const MethodSendChatAction = 15

type bot struct {
    methods [16]string
    client  http.Client
}

type Bot interface {
    GetMe() (User, error)
    SendMessage(*SendMessageRequest) (Message, error)
    ForwardMessage(*ForwardMessageRequest) (Message, error)
}

func NewBot(token string) Bot {
    b := &bot{
        client: http.Client{},
    }

    b.methods[MethodGetMe] = ApiEndpoint + token + "/getMe"
    b.methods[MethodSendMessage] = ApiEndpoint + token + "/sendMessage"
    b.methods[MethodForwardMessage] = ApiEndpoint + token + "/forwardMessage"
    b.methods[MethodSendPhoto] = ApiEndpoint + token + "/sendPhoto"
    b.methods[MethodSendAudio] = ApiEndpoint + token + "/sendAudio"
    b.methods[MethodSendDocument] = ApiEndpoint + token + "/sendDocument"
    b.methods[MethodSendVideo] = ApiEndpoint + token + "/sendVideo"
    b.methods[MethodSendVoice] = ApiEndpoint + token + "/sendVoice"
    b.methods[MethodSendVideoNote] = ApiEndpoint + token + "/sendVideoNote"
    b.methods[MethodSendMediaGroup] = ApiEndpoint + token + "/sendMediaGroup"
    b.methods[MethodSendLocation] = ApiEndpoint + token + "/sendLocation"
    b.methods[MethodEditMessageLiveLocation] = ApiEndpoint + token + "/editMessageLiveLocation"
    b.methods[MethodStopMessageLiveLocation] = ApiEndpoint + token + "/stopMessageLiveLocation"
    b.methods[MethodSendVenue] = ApiEndpoint + token + "/sendVenue"
    b.methods[MethodSendContact] = ApiEndpoint + token + "/sendContact"
    b.methods[MethodSendChatAction] = ApiEndpoint + token + "/sendChatAction"

    return b
}

func (b *bot) GetMe() (me User, err error) {
    resp, err := b.callMethod(MethodGetMe, nil)
    if err != nil {
        return
    }

    if err = json.Unmarshal(resp.Result, &me); err != nil {
        return
    }

    return
}

func (b *bot) SendMessage(request *SendMessageRequest) (message Message, err error) {
    resp, err := b.callMethod(MethodSendMessage, request)
    if err != nil {
        return
    }

    if err = json.Unmarshal(resp.Result, &message); err != nil {
        return
    }

    return
}

func (b *bot) ForwardMessage(request *ForwardMessageRequest) (message Message, err error) {
    resp, err := b.callMethod(MethodForwardMessage, request)
    if err != nil {
        return
    }

    if err = json.Unmarshal(resp.Result, &message); err != nil {
        return
    }

    return
}

func (b *bot) callMethod(method int, request interface{}) (response Response, err error) {
    payload, err := json.Marshal(request)
    if err != nil {
        return
    }

    resp, err := b.client.Post(b.methods[method], "application/json", bytes.NewReader(payload))
    if err != nil {
        return
    }

    defer resp.Body.Close()

    data, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return
    }

    if err = json.Unmarshal(data, &response); err != nil {
        return
    }

    if !response.Ok {
        return response, errors.New(response.Description)
    }

    return
}
