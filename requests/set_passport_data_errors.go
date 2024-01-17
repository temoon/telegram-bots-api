package requests

import (
	"context"
	"encoding/json"
	"github.com/temoon/telegram-bots-api"
	"io"
	"strconv"
)

type SetPassportDataErrors struct {
	UserId int64
	Errors []interface{}
}

func (r *SetPassportDataErrors) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new(bool)
	err = b.CallMethod(ctx, "setPassportDataErrors", r, response)
	return
}

func (r *SetPassportDataErrors) GetValues() (values map[string]interface{}, err error) {
	values = make(map[string]interface{})

	values["user_id"] = strconv.FormatInt(r.UserId, 10)

	var dataErrors []byte
	if dataErrors, err = json.Marshal(r.Errors); err != nil {
		return
	}

	values["errors"] = string(dataErrors)

	return
}

func (r *SetPassportDataErrors) GetFiles() (files map[string]io.Reader) {
	return
}
