package requests

import (
	"context"
	"github.com/temoon/telegram-bots-api"
)

type GetWebhookInfo struct {
}

func (r *GetWebhookInfo) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new(telegram.WebhookInfo)
	err = b.CallMethod(ctx, "getWebhookInfo", nil, response)
	return
}

func (r *GetWebhookInfo) IsMultipart() bool {
	return false
}

func (r *GetWebhookInfo) GetValues() (values map[string]interface{}, err error) {

	return
}
