package telegram

//go:generate go run generate/parse-api-doc.go
//go:generate go run generate/generate.go
//go:generate gofmt -w types.go
//go:generate gofmt -w requests

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"time"
)

//goland:noinspection GoUnusedConst
const EnvProduction = "prod"
const EnvTest = "test"

type Bot struct {
	opts *BotOpts
}

type BotOpts struct {
	Token   string
	Client  *http.Client
	Timeout time.Duration
	Env     string
}

type Request interface {
	Call(context.Context, *Bot) (interface{}, error)
	IsMultipart() bool
	GetValues() (map[string]interface{}, error)
}

type RequestWithResponse interface {
	Request
	CallWithResponse(context.Context, *Bot, interface{}) error
}

type Response struct {
	OK          bool                `json:"ok"`
	Description string              `json:"description,omitempty"`
	Result      json.RawMessage     `json:"result,omitempty"`
	ErrorCode   uint32              `json:"error_code,omitempty"`
	Parameters  *ResponseParameters `json:"parameters,omitempty"`
}

//goland:noinspection GoUnusedExportedFunction
func NewBot(opts *BotOpts) *Bot {
	if opts.Client == nil {
		opts.Client = &http.Client{
			Timeout: opts.Timeout,
		}
	}

	return &Bot{
		opts: opts,
	}
}

func (b *Bot) CallMethod(ctx context.Context, method string, request Request, response interface{}) (err error) {
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

	methodUrlBase := "https://api.telegram.org/bot" + b.opts.Token + "/"
	if b.opts.Env == EnvTest {
		methodUrlBase += "test/"
	}

	var httpRequest *http.Request
	if httpRequest, err = http.NewRequestWithContext(ctx, "POST", methodUrlBase+method, query); err != nil {
		return
	}
	httpRequest.Header.Set("Content-Type", contentType)

	var httpResponse *http.Response
	if httpResponse, err = b.opts.Client.Do(httpRequest); err != nil {
		return
	}
	//goland:noinspection GoUnhandledErrorResult
	defer httpResponse.Body.Close()

	var data []byte
	if data, err = io.ReadAll(httpResponse.Body); err != nil {
		return
	}

	var telegramResponse Response
	if err = json.Unmarshal(data, &telegramResponse); err != nil {
		return
	}

	if !telegramResponse.OK {
		return errors.New(telegramResponse.Description)
	}

	return json.Unmarshal(telegramResponse.Result, response)
}

func (b *Bot) getForm(request Request) (contentType string, query *bytes.Buffer, err error) {
	contentType = "application/x-www-form-urlencoded"
	query = &bytes.Buffer{}

	if request == nil {
		return
	}

	var values map[string]interface{}

	if values, err = request.GetValues(); err != nil {
		return
	}

	for key, value := range values {
		if query.Len() > 0 {
			query.WriteByte('&')
		}

		query.WriteString(url.QueryEscape(key) + "=")
		query.WriteString(url.QueryEscape(value.(string)))
	}

	return
}

func (b *Bot) getFormMultipart(request Request) (contentType string, query *bytes.Buffer, err error) {
	var values map[string]interface{}

	if values, err = request.GetValues(); err != nil {
		return
	}

	query = new(bytes.Buffer)

	var mw = multipart.NewWriter(query)
	var fw io.Writer
	var data io.Reader
	var ok bool

	for key, value := range values {
		if data, ok = value.(io.Reader); ok {
			if fw, err = mw.CreateFormFile(key, key); err != nil {
				return
			}

			if _, err = io.Copy(fw, data); err != nil {
				return
			}
		} else {
			if fw, err = mw.CreateFormField(key); err != nil {
				return
			}

			if _, err = fw.Write([]byte(value.(string))); err != nil {
				return
			}
		}
	}

	if err = mw.Close(); err != nil {
		return
	}

	contentType = mw.FormDataContentType()

	return
}
