package requests

import (
	"context"
	"encoding/json"
	"github.com/temoon/telegram-bots-api"
	"strconv"
)

type SetWebhook struct {
	AllowedUpdates     []string
	Certificate        interface{}
	DropPendingUpdates *bool
	IpAddress          *string
	MaxConnections     *int64
	SecretToken        *string
	Url                string
}

func (r *SetWebhook) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new(bool)
	err = b.CallMethod(ctx, "setWebhook", r, response)
	return
}

func (r *SetWebhook) IsMultipart() bool {
	return false
}

func (r *SetWebhook) GetValues() (values map[string]interface{}, err error) {
	values = make(map[string]interface{})

	if r.AllowedUpdates != nil {
		var dataAllowedUpdates []byte
		if dataAllowedUpdates, err = json.Marshal(r.AllowedUpdates); err != nil {
			return
		}

		values["allowed_updates"] = string(dataAllowedUpdates)
	}

	if r.Certificate != nil {
		values["certificate"] = r.Certificate
	}

	if r.DropPendingUpdates != nil {
		if *r.DropPendingUpdates {
			values["drop_pending_updates"] = "1"
		} else {
			values["drop_pending_updates"] = "0"
		}
	}

	if r.IpAddress != nil {
		values["ip_address"] = *r.IpAddress
	}

	if r.MaxConnections != nil {
		values["max_connections"] = strconv.FormatInt(*r.MaxConnections, 10)
	}

	if r.SecretToken != nil {
		values["secret_token"] = *r.SecretToken
	}

	values["url"] = r.Url

	return
}
