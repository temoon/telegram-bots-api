package requests

import (
	"encoding/json"
	"io"
	"strconv"
)

type CreateNewStickerSet struct {
	ContainsMasks bool
	Emojis        string
	MaskPosition  interface{}
	Name          string
	PngSticker    interface{}
	TgsSticker    interface{}
	Title         string
	UserId        uint64
}

func (r *CreateNewStickerSet) IsMultipart() bool {
	return true
}

func (r *CreateNewStickerSet) GetValues() (values map[string]interface{}, err error) {
	values = make(map[string]interface{})

	if r.ContainsMasks {
		values["contains_masks"] = "1"
	}

	values["emojis"] = r.Emojis

	if r.MaskPosition != nil {
		var data []byte
		if data, err = json.Marshal(r.MaskPosition); err != nil {
			return
		}

		values["mask_position"] = string(data)
	}

	values["name"] = r.Name

	switch value := r.PngSticker.(type) {
	case io.Reader:
		values["png_sticker"] = value
	case string:
		values["png_sticker"] = value
	}

	if r.TgsSticker != nil {
		values["tgs_sticker"] = r.TgsSticker
	}

	values["title"] = r.Title

	values["user_id"] = strconv.FormatUint(r.UserId, 10)

	return
}
