package requests

import (
	"context"
	"io"

	"github.com/temoon/telegram-bots-api"
)

type GetFile struct {
	FileId string
}

func (r *GetFile) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new(telegram.File)
	err = b.CallMethod(ctx, "getFile", r, response)
	return
}

func (r *GetFile) GetValues() (values map[string]string, err error) {
	values = make(map[string]string)

	values["file_id"] = r.FileId

	return
}

func (r *GetFile) GetFiles() (files map[string]io.Reader) {
	return
}
