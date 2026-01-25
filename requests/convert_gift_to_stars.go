package requests

import (
	"context"
	"io"

	"github.com/temoon/telegram-bots-api"
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

func (r *ConvertGiftToStars) GetValues() (values map[string]string, err error) {
	values = make(map[string]string)

	values["business_connection_id"] = r.BusinessConnectionId

	values["owned_gift_id"] = r.OwnedGiftId

	return
}

func (r *ConvertGiftToStars) GetFiles() (files map[string]io.Reader) {
	return
}
