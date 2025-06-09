package requests

import (
	"context"
	"github.com/temoon/telegram-bots-api"
	"io"
	"strconv"
)

type TransferGift struct {
	BusinessConnectionId string
	NewOwnerChatId       int64
	OwnedGiftId          string
	StarCount            *int64
}

func (r *TransferGift) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new(bool)
	err = b.CallMethod(ctx, "transferGift", r, response)
	return
}

func (r *TransferGift) GetValues() (values map[string]interface{}, err error) {
	values = make(map[string]interface{})

	values["business_connection_id"] = r.BusinessConnectionId

	values["new_owner_chat_id"] = strconv.FormatInt(r.NewOwnerChatId, 10)

	values["owned_gift_id"] = r.OwnedGiftId

	if r.StarCount != nil {
		values["star_count"] = strconv.FormatInt(*r.StarCount, 10)
	}

	return
}

func (r *TransferGift) GetFiles() (files map[string]io.Reader) {
	return
}
