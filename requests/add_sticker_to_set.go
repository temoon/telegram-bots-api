package requests

import (
	"context"
	"encoding/json"
	"io"
	"strconv"

	"github.com/temoon/telegram-bots-api"
)

type AddStickerToSet struct {
	Name    string
	Sticker telegram.InputSticker
	UserId  int64
}

func (r *AddStickerToSet) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new(bool)
	err = b.CallMethod(ctx, "addStickerToSet", r, response)
	return
}

func (r *AddStickerToSet) GetValues() (values map[string]string, err error) {
	values = make(map[string]string)

	values["name"] = r.Name

	var dataSticker []byte
	if dataSticker, err = json.Marshal(r.Sticker); err != nil {
		return
	}

	values["sticker"] = string(dataSticker)

	values["user_id"] = strconv.FormatInt(r.UserId, 10)

	return
}

func (r *AddStickerToSet) GetFiles() (files map[string]io.Reader) {
	files = make(map[string]io.Reader)

	if r.Sticker.Sticker.HasFile() {
		files[r.Sticker.Sticker.GetFormFieldName()] = r.Sticker.Sticker.GetFile()
	}

	return
}
