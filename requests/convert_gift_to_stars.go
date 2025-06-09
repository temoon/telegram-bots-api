package requests

import (
	"context"
	"github.com/temoon/telegram-bots-api"
	"io"
)

type ConvertGiftToStars struct {
	BusinessConnectionId string
	OwnedGiftId          string
}

func (r *ConvertGiftToStars) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new(bool)
	err = b.CallMethod(ctx, "convertGiftToStars", r, response)
	return
}

func (r *ConvertGiftToStars) GetValues() (values map[string]interface{}, err error) {
	values = make(map[string]interface{})

	values["business_connection_id"] = r.BusinessConnectionId

	values["owned_gift_id"] = r.OwnedGiftId

	return
}

func (r *ConvertGiftToStars) GetFiles() (files map[string]io.Reader) {
	return
}
