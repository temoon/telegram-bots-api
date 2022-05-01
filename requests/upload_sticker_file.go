package requests

import (
	"context"
	"strconv"

	"github.com/temoon/telegram-bots-api"
)

type UploadStickerFile struct {
	PngSticker interface{}
	UserId     int64
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

	values["png_sticker"] = r.PngSticker

	values["user_id"] = strconv.FormatInt(r.UserId, 10)

	return
}
