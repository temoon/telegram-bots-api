package requests

import (
	"context"
	"github.com/temoon/telegram-bots-api"
	"io"
	"strconv"
)

type TransferBusinessAccountStars struct {
	BusinessConnectionId string
	StarCount            int64
}

func (r *TransferBusinessAccountStars) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new(bool)
	err = b.CallMethod(ctx, "transferBusinessAccountStars", r, response)
	return
}

func (r *TransferBusinessAccountStars) GetValues() (values map[string]interface{}, err error) {
	values = make(map[string]interface{})

	values["business_connection_id"] = r.BusinessConnectionId

	values["star_count"] = strconv.FormatInt(r.StarCount, 10)

	return
}

func (r *TransferBusinessAccountStars) GetFiles() (files map[string]io.Reader) {
	return
}
