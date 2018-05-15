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

func (r *AddStickerToSet) GetValues() (values map[string][]interface{}, err error) {
    values = make(map[string][]interface{})

    values["user_id"] = []interface{}{strconv.Itoa(r.UserID)}
    values["name"] = []interface{}{r.Name}
    values["png_sticker"] = []interface{}{r.PNGSticker}
    values["emojis"] = []interface{}{r.Emojis}

    if r.MaskPosition != nil {
        var data []byte
        if data, err = json.Marshal(r.MaskPosition); err != nil {
            return
        }

        values["mask_position"] = []interface{}{string(data)}
    }

    return
}
