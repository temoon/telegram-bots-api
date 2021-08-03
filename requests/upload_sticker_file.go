package requests

import (
	"io"
	"strconv"
)

type UploadStickerFile struct {
	PngSticker interface{}
	UserId     uint64
}

func (r *UploadStickerFile) IsMultipart() bool {
	return true
}

func (r *UploadStickerFile) GetValues() (values map[string]interface{}, err error) {
	values = make(map[string]interface{})

	values["png_sticker"] = r.PngSticker

	values["user_id"] = strconv.FormatUint(r.UserId, 10)

	return
}
