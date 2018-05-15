package requests

import (
    "encoding/json"
    "strconv"
)

type AddStickerToSet struct {
    UserID       int
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

    values["user_id"] = strconv.Itoa(r.UserID)
    values["name"] = r.Name
    values["png_sticker"] = r.PNGSticker
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
