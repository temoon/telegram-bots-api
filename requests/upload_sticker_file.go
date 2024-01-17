package requests

import (
	"context"
	"github.com/temoon/telegram-bots-api"
	"io"
	"strconv"
)

type UploadStickerFile struct {
	StickerFormat string
	UserId        int64
	Sticker       telegram.InputFile
}

func (r *UploadStickerFile) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new(telegram.File)
	err = b.CallMethod(ctx, "uploadStickerFile", r, response)
	return
}

func (r *UploadStickerFile) GetValues() (values map[string]interface{}, err error) {
	values = make(map[string]interface{})

	values["sticker_format"] = r.StickerFormat

	values["user_id"] = strconv.FormatInt(r.UserId, 10)

	values["sticker"] = r.Sticker.GetValue()

	return
}

func (r *UploadStickerFile) GetFiles() (files map[string]io.Reader) {
	return
}
