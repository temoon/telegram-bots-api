package requests

import (
	"encoding/json"
	"io"
	"strconv"
)

type SetStickerSetThumb struct {
	Name   string
	Thumb  interface{}
	UserId uint64
}

func (r *SetStickerSetThumb) IsMultipart() bool {
	return true
}

func (r *SetStickerSetThumb) GetValues() (values map[string]interface{}, err error) {
	values = make(map[string]interface{})

	values["name"] = r.Name

	switch value := r.Thumb.(type) {
	case io.Reader:
		values["thumb"] = value
	case string:
		values["thumb"] = value
	}

	values["user_id"] = strconv.FormatUint(r.UserId, 10)

	return
}
