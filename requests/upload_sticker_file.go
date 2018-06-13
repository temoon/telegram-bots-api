package requests

import (
    "io"
    "strconv"
)

type UploadStickerFile struct {
    UserID     uint32
    PNGSticker io.Reader
}

func (r *UploadStickerFile) IsMultipart() bool {
    return true
}

func (r *UploadStickerFile) GetValues() (values map[string]interface{}, err error) {
    values = make(map[string]interface{})

    values["user_id"] = strconv.FormatUint(uint64(r.UserID), 10)
    values["png_sticker"] = r.PNGSticker

    return
}
