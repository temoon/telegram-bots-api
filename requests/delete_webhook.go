package requests

import (
	"context"
	"github.com/temoon/telegram-bots-api"
	"io"
)

type DeleteWebhook struct {
	DropPendingUpdates *bool
}

func (r *DeleteWebhook) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new(bool)
	err = b.CallMethod(ctx, "deleteWebhook", r, response)
	return
}

func (r *DeleteWebhook) GetValues() (values map[string]interface{}, err error) {
	values = make(map[string]interface{})

	if r.DropPendingUpdates != nil {
		if *r.DropPendingUpdates {
			values["drop_pending_updates"] = "1"
		} else {
			values["drop_pending_updates"] = "0"
		}
	}

	return
}

func (r *DeleteWebhook) GetFiles() (files map[string]io.Reader) {
	return
}
