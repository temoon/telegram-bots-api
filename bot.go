package telegram

import (
    "bytes"
    "encoding/json"
    "errors"
    "io"
    "io/ioutil"
    "mime/multipart"
    "net/http"
    "net/url"
    "os"
    "strings"

    "go-telegram-bots-api/requests"
)

type bot struct {
    token  string
    client http.Client
}

type Bot interface {
    GetUpdates(*requests.GetUpdates) ([]Update, error)
    GetMe() (User, error)
    SendMessage(*requests.SendMessage) (Message, error)
    ForwardMessage(*requests.ForwardMessage) (Message, error)
    SendPhoto(request *requests.SendPhoto) (Message, error)
}

func NewBot(token string) Bot {
    b := &bot{
        token:  token,
        client: http.Client{},
    }

    return b
}

func (b *bot) GetUpdates(request *requests.GetUpdates) (updates []Update, err error) {
    err = b.callMethod(b.getMethodURL("getUpdates"), request, &updates)

    return
}

func (b *bot) GetMe() (me User, err error) {
    err = b.callMethod(b.getMethodURL("getMe"), nil, &me)

    return
}

func (b *bot) SendMessage(request *requests.SendMessage) (message Message, err error) {
    err = b.callMethod(b.getMethodURL("sendMessage"), request, &message)

    return
}

func (b *bot) ForwardMessage(request *requests.ForwardMessage) (message Message, err error) {
    err = b.callMethod(b.getMethodURL("forwardMessage"), request, &message)

    return
}

func (b *bot) SendPhoto(request *requests.SendPhoto) (message Message, err error) {
    err = b.callMethod(b.getMethodURL("sendPhoto"), request, &message)

    return
}

func (b *bot) getMethodURL(method string) string {
    return "https://api.telegram.org/bot" + b.token + "/" + method
}

func (b *bot) callMethod(methodUrl string, request Request, response interface{}) (err error) {
    var contentType string
    var query *bytes.Buffer

    if request != nil && request.IsMultipart() {
        contentType, query, err = b.getFormMultipart(request)
    } else {
        contentType, query, err = b.getForm(request)
    }

    if err != nil {
        return
    }

    var httpResponse *http.Response
    if httpResponse, err = b.client.Post(methodUrl, contentType, query); err != nil {
        return
    }
    defer httpResponse.Body.Close()

    var data []byte
    if data, err = ioutil.ReadAll(httpResponse.Body); err != nil {
        return
    }

    var telegramResponse Response
    if err = json.Unmarshal(data, &telegramResponse); err != nil {
        return
    }

    if !telegramResponse.Ok {
        return errors.New(telegramResponse.Description)
    }

    return json.Unmarshal(telegramResponse.Result, response)
}

func (b *bot) getForm(request Request) (contentType string, query *bytes.Buffer, err error) {
    contentType = "application/x-www-form-urlencoded"
    query = &bytes.Buffer{}

    if request == nil {
        return
    }

    var rv map[string][]interface{}
    if rv, err = request.GetValues(); err != nil {
        return
    }

    var prefix string

    for key, values := range rv {
        prefix = url.QueryEscape(key) + "="

        for _, value := range values {
            if query.Len() > 0 {
                query.WriteByte('&')
            }

            query.WriteString(prefix)
            query.WriteString(url.QueryEscape(value.(string)))
        }
    }

    return
}

func (b *bot) getFormMultipart(request Request) (contentType string, query *bytes.Buffer, err error) {
    var rv map[string][]interface{}
    if rv, err = request.GetValues(); err != nil {
        return
    }

    query = &bytes.Buffer{}

    var mw = multipart.NewWriter(query)
    var fw io.Writer

    for key, values := range rv {
        for _, value := range values {
            if file, ok := value.(*os.File); ok {
                if fw, err = mw.CreateFormFile(key, file.Name()); err != nil {
                    return
                }
            } else {
                if fw, err = mw.CreateFormField(key); err != nil {
                    return
                }
            }

            reader, ok := value.(io.Reader)
            if !ok {
                reader = strings.NewReader(value.(string))
            }

            if _, err = io.Copy(fw, reader); err != nil {
                return
            }
        }
    }

    mw.Close()

    contentType = mw.FormDataContentType()

    return
}
