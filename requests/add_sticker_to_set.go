package requests

import (
	"context"
	"encoding/json"
	"github.com/temoon/telegram-bots-api"
	"io"
	"strconv"
)

type AddStickerToSet struct {
	UserId  int64
	Name    string
	Sticker telegram.InputSticker
}

func (r *AddStickerToSet) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new(bool)
	err = b.CallMethod(ctx, "addStickerToSet", r, response)
	return
}

func (r *AddStickerToSet) GetValues() (values map[string]interface{}, err error) {
	values = make(map[string]interface{})

	values["user_id"] = strconv.FormatInt(r.UserId, 10)

	values["name"] = r.Name

	var dataSticker []byte
	if dataSticker, err = json.Marshal(r.Sticker); err != nil {
		return
	}

	values["sticker"] = string(dataSticker)

	return
}

func (r *AddStickerToSet) GetFiles() (files map[string]io.Reader) {
	files = make(map[string]io.Reader)

	if r.Sticker.Sticker.HasFile() {
		files[r.Sticker.Sticker.GetFormFieldName()] = r.Sticker.Sticker.GetFile()
	}

	return
}
