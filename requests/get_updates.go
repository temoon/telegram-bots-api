package requests

import (
	"context"
	"encoding/json"
	"github.com/temoon/go-telegram-bots-api"
	"strconv"
)

type GetUpdates struct {
	AllowedUpdates []string
	Limit          int32
	Offset         int32
	Timeout        int32
}

func (r *GetUpdates) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new([]telegram.Update)
	err = b.CallMethod(ctx, "getUpdates", r, response)
	return
}

func (r *GetUpdates) IsMultipart() (multipart bool) {
	return false
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

	if r.Limit != 0 {
		values["limit"] = strconv.FormatInt(int64(r.Limit), 10)
	}

	if r.Offset != 0 {
		values["offset"] = strconv.FormatInt(int64(r.Offset), 10)
	}

	if r.Timeout != 0 {
		values["timeout"] = strconv.FormatInt(int64(r.Timeout), 10)
	}

	return
}
