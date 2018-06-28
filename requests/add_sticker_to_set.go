package requests

import (
    "encoding/json"
    "errors"
    "io"
    "strconv"
)

type AddStickerToSet struct {
    UserID       uint32
    Name         string
    PNGSticker   interface{}
    Emojis       string
    MaskPosition interface{}
}

func (r *AddStickerToSet) IsMultipart() bool {
    return true
}

func (r *AddStickerToSet) GetValues() (values map[string]interface{}, err error) {
    values = make(map[string]interface{})

    values["user_id"] = strconv.FormatUint(uint64(r.UserID), 10)
    values["name"] = r.Name

    switch inputFile := r.PNGSticker.(type) {
    case string:
        values["png_sticker"] = inputFile
    case io.Reader:
        values["png_sticker"] = inputFile
    default:
        return nil, errors.New("invalid photo")
    }

    values["emojis"] = r.Emojis

    if r.MaskPosition != nil {
        var data []byte

        if data, err = json.Marshal(r.MaskPosition); err != nil {
            return
        }

        values["mask_position"] = string(data)
    }

    return
}
