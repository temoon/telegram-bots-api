package requests

import (
	"context"
	"io"
	"strconv"

	"github.com/temoon/telegram-bots-api"
)

type UploadStickerFile struct {
	Sticker       telegram.InputFile
	StickerFormat string
	UserId        int64
}

func (r *UploadStickerFile) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new(telegram.File)
	err = b.CallMethod(ctx, "uploadStickerFile", r, response)
	return
}

func (r *UploadStickerFile) GetValues() (values map[string]string, err error) {
	values = make(map[string]string)

	values["sticker"] = r.Sticker.String()

	values["sticker_format"] = r.StickerFormat

	values["user_id"] = strconv.FormatInt(r.UserId, 10)

	return
}

func (r *UploadStickerFile) GetFiles() (files map[string]io.Reader) {
	files = make(map[string]io.Reader)

	if r.Sticker.HasFile() {
		files[r.Sticker.GetFormFieldName()] = r.Sticker.GetFile()
	}

	return
}
