package requests

import (
	"context"
	"io"
	"strconv"

	"github.com/temoon/telegram-bots-api"
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

func (r *TransferBusinessAccountStars) GetValues() (values map[string]string, err error) {
	values = make(map[string]string)

	values["business_connection_id"] = r.BusinessConnectionId

	values["star_count"] = strconv.FormatInt(r.StarCount, 10)

	return
}

func (r *TransferBusinessAccountStars) GetFiles() (files map[string]io.Reader) {
	return
}
