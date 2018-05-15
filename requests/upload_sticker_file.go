package requests

import (
    "os"
    "strconv"
)

type UploadStickerFile struct {
    UserID     int
    PNGSticker *os.File
}

func (r *UploadStickerFile) IsMultipart() bool {
    return true
}

func (r *UploadStickerFile) GetValues() (values map[string][]interface{}, err error) {
    values = make(map[string][]interface{})

    values["user_id"] = []interface{}{strconv.Itoa(r.UserID)}
    values["png_sticker"] = []interface{}{r.PNGSticker}

    return
}
