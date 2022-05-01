package requests

import (
	"context"
	"encoding/json"
	"io"
	"strconv"

	"github.com/temoon/telegram-bots-api"
)

type CreateNewStickerSet struct {
	ContainsMasks *bool
	Emojis        string
	MaskPosition  *telegram.MaskPosition
	Name          string
	PngSticker    interface{}
	TgsSticker    interface{}
	Title         string
	UserId        int64
	WebmSticker   interface{}
}

func (r *CreateNewStickerSet) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new(bool)
	err = b.CallMethod(ctx, "createNewStickerSet", r, response)
	return
}

func (r *CreateNewStickerSet) IsMultipart() (multipart bool) {
	return true
}

func (r *CreateNewStickerSet) GetValues() (values map[string]interface{}, err error) {
	values = make(map[string]interface{})

	if r.ContainsMasks != nil {
		if *r.ContainsMasks {
			values["contains_masks"] = "1"
		} else {
			values["contains_masks"] = "0"
		}
	}

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

	values["title"] = r.Title

	values["user_id"] = strconv.FormatInt(r.UserId, 10)

	if r.WebmSticker != nil {
		values["webm_sticker"] = r.WebmSticker
	}

	return
}
