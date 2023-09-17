package requests

import (
	"context"
	"encoding/json"
	"github.com/temoon/telegram-bots-api"
	"strconv"
)

type SetPassportDataErrors struct {
	Errors []interface{}
	UserId int64
}

func (r *SetPassportDataErrors) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new(bool)
	err = b.CallMethod(ctx, "setPassportDataErrors", r, response)
	return
}

func (r *SetPassportDataErrors) IsMultipart() (multipart bool) {
	return false
}

func (r *SetPassportDataErrors) GetValues() (values map[string]interface{}, err error) {
	values = make(map[string]interface{})

	var dataErrors []byte
	if dataErrors, err = json.Marshal(r.Errors); err != nil {
		return
	}

	values["errors"] = string(dataErrors)

	values["user_id"] = strconv.FormatInt(r.UserId, 10)

	return
}
