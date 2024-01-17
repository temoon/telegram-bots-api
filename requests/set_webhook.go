package requests

import (
	"context"
	"encoding/json"
	"github.com/temoon/telegram-bots-api"
	"io"
	"strconv"
)

type SetWebhook struct {
	Url                string
	AllowedUpdates     []string
	Certificate        *telegram.InputFile
	DropPendingUpdates *bool
	IpAddress          *string
	MaxConnections     *int64
	SecretToken        *string
}

func (r *SetWebhook) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new(bool)
	err = b.CallMethod(ctx, "setWebhook", r, response)
	return
}

func (r *SetWebhook) GetValues() (values map[string]interface{}, err error) {
	values = make(map[string]interface{})

	values["url"] = r.Url

	if r.AllowedUpdates != nil {
		var dataAllowedUpdates []byte
		if dataAllowedUpdates, err = json.Marshal(r.AllowedUpdates); err != nil {
			return
		}

		values["allowed_updates"] = string(dataAllowedUpdates)
	}

	if r.Certificate != nil {
		values["certificate"] = r.Certificate.GetValue()
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

	return
}

func (r *SetWebhook) GetFiles() (files map[string]io.Reader) {
	return
}
