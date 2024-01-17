package requests

import (
	"context"
	"encoding/json"
	"github.com/temoon/telegram-bots-api"
	"io"
	"strconv"
)

type GetUpdates struct {
	AllowedUpdates []string
	Limit          *int64
	Offset         *int64
	Timeout        *int64
}

func (r *GetUpdates) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new([]telegram.Update)
	err = b.CallMethod(ctx, "getUpdates", r, response)
	return
}

func (r *GetUpdates) GetValues() (values map[string]interface{}, err error) {
	values = make(map[string]interface{})

	if r.AllowedUpdates != nil {
		var dataAllowedUpdates []byte
		if dataAllowedUpdates, err = json.Marshal(r.AllowedUpdates); err != nil {
			return
		}

		values["allowed_updates"] = string(dataAllowedUpdates)
	}

	if r.Limit != nil {
		values["limit"] = strconv.FormatInt(*r.Limit, 10)
	}

	if r.Offset != nil {
		values["offset"] = strconv.FormatInt(*r.Offset, 10)
	}

	if r.Timeout != nil {
		values["timeout"] = strconv.FormatInt(*r.Timeout, 10)
	}

	return
}

func (r *GetUpdates) GetFiles() (files map[string]io.Reader) {
	return
}
