package requests

import (
	"context"
	"encoding/json"
	"github.com/temoon/go-telegram-bots-api"
	"io"
	"strconv"
)

type AddStickerToSet struct {
	Emojis       string
	MaskPosition *telegram.MaskPosition
	Name         string
	PngSticker   interface{}
	TgsSticker   interface{}
	UserId       int64
}

func (r *AddStickerToSet) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new(bool)
	err = b.CallMethod(ctx, "addStickerToSet", r, response)
	return
}

func (r *AddStickerToSet) IsMultipart() (multipart bool) {
	return true
}

func (r *AddStickerToSet) GetValues() (values map[string]interface{}, err error) {
	values = make(map[string]interface{})

	values["emojis"] = r.Emojis

	if r.MaskPosition != nil {
		var dataMaskPosition []byte
		if dataMaskPosition, err = json.Marshal(r.MaskPosition); err != nil {
			return
		}

		values["mask_position"] = string(dataMaskPosition)
	}

	values["name"] = r.Name

	switch value := r.PngSticker.(type) {
	case io.Reader:
		values["png_sticker"] = value
	case string:
		values["png_sticker"] = value
	}

	if r.TgsSticker != nil {
		values["tgs_sticker"] = r.TgsSticker
	}

	values["user_id"] = strconv.FormatInt(r.UserId, 10)

	return
}
