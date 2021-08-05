package requests

import (
	"context"
	"github.com/temoon/go-telegram-bots-api"
)

type GetWebhookInfo struct {
}

func (r *GetWebhookInfo) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new(telegram.WebhookInfo)
	err = b.CallMethod(ctx, "getWebhookInfo", nil, response)
	return
}

func (r *GetWebhookInfo) IsMultipart() (multipart bool) {
	return
}

func (r *GetWebhookInfo) GetValues() (values map[string]interface{}, err error) {
	return
}
