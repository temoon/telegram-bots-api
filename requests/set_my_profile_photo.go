package requests

import (
	"context"
	"encoding/json"
	"io"

	"github.com/temoon/telegram-bots-api"
)

type SetMyProfilePhoto struct {
	Photo interface{}
}

func (r *SetMyProfilePhoto) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new(bool)
	err = b.CallMethod(ctx, "setMyProfilePhoto", r, response)
	return
}

func (r *SetMyProfilePhoto) GetValues() (values map[string]string, err error) {
	values = make(map[string]string)

	var dataPhoto []byte
	if dataPhoto, err = json.Marshal(r.Photo); err != nil {
		return
	}

	values["photo"] = string(dataPhoto)

	return
}

func (r *SetMyProfilePhoto) GetFiles() (files map[string]io.Reader) {
	files = make(map[string]io.Reader)

	switch value := r.Photo.(type) {
	case telegram.InputProfilePhotoStatic:
		if value.Photo.HasFile() {
			files[value.Photo.GetFormFieldName()] = value.Photo.GetFile()
		}
	case telegram.InputProfilePhotoAnimated:
		if value.Animation.HasFile() {
			files[value.Animation.GetFormFieldName()] = value.Animation.GetFile()
		}
	}

	return
}
