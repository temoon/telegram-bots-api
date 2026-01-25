package requests

import (
	"context"
	"io"

	"github.com/temoon/telegram-bots-api"
)

type GetMyDefaultAdministratorRights struct {
	ForChannels *bool
}

func (r *GetMyDefaultAdministratorRights) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new(telegram.ChatAdministratorRights)
	err = b.CallMethod(ctx, "getMyDefaultAdministratorRights", r, response)
	return
}

func (r *GetMyDefaultAdministratorRights) GetValues() (values map[string]string, err error) {
	values = make(map[string]string)

	if r.ForChannels != nil {
		if *r.ForChannels {
			values["for_channels"] = "1"
		} else {
			values["for_channels"] = "0"
		}
	}

	return
}

func (r *GetMyDefaultAdministratorRights) GetFiles() (files map[string]io.Reader) {
	return
}
