package requests

import (
	"context"
	"encoding/json"
	"io"

	"github.com/temoon/telegram-bots-api"
)

type SetBusinessAccountGiftSettings struct {
	AcceptedGiftTypes    telegram.AcceptedGiftTypes
	BusinessConnectionId string
	ShowGiftButton       bool
}

func (r *SetBusinessAccountGiftSettings) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new(bool)
	err = b.CallMethod(ctx, "setBusinessAccountGiftSettings", r, response)
	return
}

func (r *SetBusinessAccountGiftSettings) GetValues() (values map[string]interface{}, err error) {
	values = make(map[string]interface{})

	var dataAcceptedGiftTypes []byte
	if dataAcceptedGiftTypes, err = json.Marshal(r.AcceptedGiftTypes); err != nil {
		return
	}

	values["accepted_gift_types"] = string(dataAcceptedGiftTypes)

	values["business_connection_id"] = r.BusinessConnectionId

	if r.ShowGiftButton {
		values["show_gift_button"] = "1"
	} else {
		values["show_gift_button"] = "0"
	}

	return
}

func (r *SetBusinessAccountGiftSettings) GetFiles() (files map[string]io.Reader) {
	return
}
