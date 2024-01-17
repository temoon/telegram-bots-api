package requests

import (
	"context"
	"github.com/temoon/telegram-bots-api"
	"io"
)

type GetWebhookInfo struct {
}

func (r *GetWebhookInfo) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new(telegram.WebhookInfo)
	err = b.CallMethod(ctx, "getWebhookInfo", nil, response)
	return
}

func (r *GetWebhookInfo) GetValues() (values map[string]interface{}, err error) {

	return
}

func (r *GetWebhookInfo) GetFiles() (files map[string]io.Reader) {
	return
}
