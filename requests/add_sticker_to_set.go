package requests

import (
	"context"
	"encoding/json"
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

func (r *AddStickerToSet) IsMultipart() (multipart bool) {
	return true
}

func (r *AddStickerToSet) GetValues() (values map[string]interface{}, err error) {
	values = make(map[string]interface{})

	values["name"] = r.Name

	var dataSticker []byte
	if dataSticker, err = json.Marshal(r.Sticker); err != nil {
		return
	}

	values["sticker"] = string(dataSticker)

	values["user_id"] = strconv.FormatInt(r.UserId, 10)

	return
}
