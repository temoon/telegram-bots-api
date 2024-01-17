package requests

import (
	"context"
	"encoding/json"
	"github.com/temoon/telegram-bots-api"
	"io"
)

type SetMyDefaultAdministratorRights struct {
	ForChannels *bool
	Rights      *telegram.ChatAdministratorRights
}

func (r *SetMyDefaultAdministratorRights) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new(bool)
	err = b.CallMethod(ctx, "setMyDefaultAdministratorRights", r, response)
	return
}

func (r *SetMyDefaultAdministratorRights) GetValues() (values map[string]interface{}, err error) {
	values = make(map[string]interface{})

	if r.ForChannels != nil {
		if *r.ForChannels {
			values["for_channels"] = "1"
		} else {
			values["for_channels"] = "0"
		}
	}

	if r.Rights != nil {
		var dataRights []byte
		if dataRights, err = json.Marshal(r.Rights); err != nil {
			return
		}

		values["rights"] = string(dataRights)
	}

	return
}

func (r *SetMyDefaultAdministratorRights) GetFiles() (files map[string]io.Reader) {
	return
}
