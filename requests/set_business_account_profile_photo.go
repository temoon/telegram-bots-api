package requests

import (
	"context"
	"encoding/json"
	"github.com/temoon/telegram-bots-api"
	"io"
)

type SetBusinessAccountProfilePhoto struct {
	BusinessConnectionId string
	Photo                interface{}
	IsPublic             *bool
}

func (r *SetBusinessAccountProfilePhoto) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new(bool)
	err = b.CallMethod(ctx, "setBusinessAccountProfilePhoto", r, response)
	return
}

func (r *SetBusinessAccountProfilePhoto) GetValues() (values map[string]interface{}, err error) {
	values = make(map[string]interface{})

	values["business_connection_id"] = r.BusinessConnectionId

	var dataPhoto []byte
	if dataPhoto, err = json.Marshal(r.Photo); err != nil {
		return
	}

	values["photo"] = string(dataPhoto)

	if r.IsPublic != nil {
		if *r.IsPublic {
			values["is_public"] = "1"
		} else {
			values["is_public"] = "0"
		}
	}

	return
}

func (r *SetBusinessAccountProfilePhoto) GetFiles() (files map[string]io.Reader) {
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
