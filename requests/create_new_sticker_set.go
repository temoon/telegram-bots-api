package requests

import (
    "encoding/json"
    "strconv"
)

type CreateNewStickerSet struct {
    UserID        int
    Name          string
    Title         string
    PNGSticker    interface{}
    Emojis        string
    ContainsMasks bool
    MaskPosition  interface{}
}

func (r *CreateNewStickerSet) IsMultipart() bool {
    return true
}

func (r *CreateNewStickerSet) GetValues() (values map[string]interface{}, err error) {
    values = make(map[string]interface{})

    values["user_id"] = strconv.Itoa(r.UserID)
    values["name"] = r.Name
    values["title"] = r.Title
    values["png_sticker"] = r.PNGSticker
    values["emojis"] = r.Emojis

    if r.ContainsMasks {
        values["contains_masks"] = "1"
    }

    if r.MaskPosition != nil {
        var data []byte
        if data, err = json.Marshal(r.MaskPosition); err != nil {
            return
        }

        values["mask_position"] = string(data)
    }

    return
}
