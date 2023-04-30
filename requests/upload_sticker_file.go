package requests

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/temoon/telegram-bots-api"
)

type UploadStickerFile struct {
	Sticker       telegram.InputSticker
	StickerFormat string
	UserId        int64
}

func (r *UploadStickerFile) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new(telegram.File)
	err = b.CallMethod(ctx, "uploadStickerFile", r, response)
	return
}

func (r *UploadStickerFile) IsMultipart() (multipart bool) {
	return true
}

func (r *UploadStickerFile) GetValues() (values map[string]interface{}, err error) {
	values = make(map[string]interface{})

	var dataSticker []byte
	if dataSticker, err = json.Marshal(r.Sticker); err != nil {
		return
	}

	values["sticker"] = string(dataSticker)

	values["sticker_format"] = r.StickerFormat

	values["user_id"] = strconv.FormatInt(r.UserId, 10)

	return
}
