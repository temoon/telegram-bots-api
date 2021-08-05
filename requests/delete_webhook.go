package requests

import (
	"context"
	"github.com/temoon/go-telegram-bots-api"
)

type DeleteWebhook struct {
	DropPendingUpdates bool
}

func (r *DeleteWebhook) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new(bool)
	err = b.CallMethod(ctx, "deleteWebhook", r, response)
	return
}

func (r *DeleteWebhook) IsMultipart() (multipart bool) {
	return false
}

func (r *DeleteWebhook) GetValues() (values map[string]interface{}, err error) {
	values = make(map[string]interface{})

	if r.DropPendingUpdates {
		values["drop_pending_updates"] = "1"
	}

	return
}
