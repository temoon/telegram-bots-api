package requests

import (
	"context"
	"encoding/json"
	"io"
	"strconv"

	"github.com/temoon/telegram-bots-api"
)

type SavePreparedKeyboardButton struct {
	Button telegram.KeyboardButton
	UserId int64
}

func (r *SavePreparedKeyboardButton) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new(telegram.PreparedKeyboardButton)
	err = b.CallMethod(ctx, "savePreparedKeyboardButton", r, response)
	return
}

func (r *SavePreparedKeyboardButton) GetValues() (values map[string]string, err error) {
	values = make(map[string]string)

	var dataButton []byte
	if dataButton, err = json.Marshal(r.Button); err != nil {
		return
	}

	values["button"] = string(dataButton)

	values["user_id"] = strconv.FormatInt(r.UserId, 10)

	return
}

func (r *SavePreparedKeyboardButton) GetFiles() (files map[string]io.Reader) {
	return
}
