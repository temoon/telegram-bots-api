package requests

import (
	"context"
	"github.com/temoon/telegram-bots-api"
	"io"
)

type GetFile struct {
	FileId string
}

func (r *GetFile) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new(telegram.File)
	err = b.CallMethod(ctx, "getFile", r, response)
	return
}

func (r *GetFile) GetValues() (values map[string]interface{}, err error) {
	values = make(map[string]interface{})

	values["file_id"] = r.FileId

	return
}

func (r *GetFile) GetFiles() (files map[string]io.Reader) {
	return
}
