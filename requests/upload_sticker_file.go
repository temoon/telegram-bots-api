package requests

import (
	"context"
	"github.com/temoon/telegram-bots-api"
	"strconv"
)

type UploadStickerFile struct {
	Sticker       interface{}
	StickerFormat string
	UserId        int64
}

func (r *UploadStickerFile) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new(telegram.File)
	err = b.CallMethod(ctx, "uploadStickerFile", r, response)
	return
}

func (r *UploadStickerFile) IsMultipart() bool {
	return true
}

func (r *UploadStickerFile) GetValues() (values map[string]interface{}, err error) {
	values = make(map[string]interface{})

	values["sticker"] = r.Sticker

	values["sticker_format"] = r.StickerFormat

	values["user_id"] = strconv.FormatInt(r.UserId, 10)

	return
}
