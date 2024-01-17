package requests

import (
	"context"
	"encoding/json"
	"github.com/temoon/telegram-bots-api"
	"io"
)

type SetStickerMaskPosition struct {
	Sticker      string
	MaskPosition *telegram.MaskPosition
}

func (r *SetStickerMaskPosition) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new(bool)
	err = b.CallMethod(ctx, "setStickerMaskPosition", r, response)
	return
}

func (r *SetStickerMaskPosition) GetValues() (values map[string]interface{}, err error) {
	values = make(map[string]interface{})

	values["sticker"] = r.Sticker

	if r.MaskPosition != nil {
		var dataMaskPosition []byte
		if dataMaskPosition, err = json.Marshal(r.MaskPosition); err != nil {
			return
		}

		values["mask_position"] = string(dataMaskPosition)
	}

	return
}

func (r *SetStickerMaskPosition) GetFiles() (files map[string]io.Reader) {
	return
}
