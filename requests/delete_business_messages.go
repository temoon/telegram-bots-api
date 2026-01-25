package requests

import (
	"context"
	"encoding/json"
	"io"

	"github.com/temoon/telegram-bots-api"
)

type DeleteBusinessMessages struct {
	BusinessConnectionId string
	MessageIds           []int64
}

func (r *DeleteBusinessMessages) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new(bool)
	err = b.CallMethod(ctx, "deleteBusinessMessages", r, response)
	return
}

func (r *DeleteBusinessMessages) GetValues() (values map[string]string, err error) {
	values = make(map[string]string)

	values["business_connection_id"] = r.BusinessConnectionId

	var dataMessageIds []byte
	if dataMessageIds, err = json.Marshal(r.MessageIds); err != nil {
		return
	}

	values["message_ids"] = string(dataMessageIds)

	return
}

func (r *DeleteBusinessMessages) GetFiles() (files map[string]io.Reader) {
	return
}
