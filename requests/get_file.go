package requests

import (
	"context"

	"github.com/temoon/go-telegram-bots-api"
)

type GetFile struct {
	FileId string
}

func (r *GetFile) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new(telegram.File)
	err = b.CallMethod(ctx, "getFile", r, response)
	return
}

func (r *GetFile) IsMultipart() (multipart bool) {
	return false
}

func (r *GetFile) GetValues() (values map[string]interface{}, err error) {
	values = make(map[string]interface{})

	values["file_id"] = r.FileId

	return
}
