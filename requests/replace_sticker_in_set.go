package requests

import (
	"context"
	"encoding/json"
	"github.com/temoon/telegram-bots-api"
	"io"
	"strconv"
)

type ReplaceStickerInSet struct {
	Name       string
	OldSticker string
	Sticker    telegram.InputSticker
	UserId     int64
}

func (r *ReplaceStickerInSet) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new(bool)
	err = b.CallMethod(ctx, "replaceStickerInSet", r, response)
	return
}

func (r *ReplaceStickerInSet) GetValues() (values map[string]interface{}, err error) {
	values = make(map[string]interface{})

	values["name"] = r.Name

	values["old_sticker"] = r.OldSticker

	var dataSticker []byte
	if dataSticker, err = json.Marshal(r.Sticker); err != nil {
		return
	}

	values["sticker"] = string(dataSticker)

	values["user_id"] = strconv.FormatInt(r.UserId, 10)

	return
}

func (r *ReplaceStickerInSet) GetFiles() (files map[string]io.Reader) {
	files = make(map[string]io.Reader)

	if r.Sticker.Sticker.HasFile() {
		files[r.Sticker.Sticker.GetFormFieldName()] = r.Sticker.Sticker.GetFile()
	}

	return
}
