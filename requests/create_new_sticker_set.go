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

func (r *CreateNewStickerSet) GetValues() (values map[string][]interface{}, err error) {
    values = make(map[string][]interface{})

    values["user_id"] = []interface{}{strconv.Itoa(r.UserID)}
    values["name"] = []interface{}{r.Name}
    values["title"] = []interface{}{r.Title}
    values["png_sticker"] = []interface{}{r.PNGSticker}
    values["emojis"] = []interface{}{r.Emojis}

    if r.ContainsMasks {
        values["contains_masks"] = []interface{}{"1"}
    }

    if r.MaskPosition != nil {
        var data []byte
        if data, err = json.Marshal(r.MaskPosition); err != nil {
            return
        }

        values["mask_position"] = []interface{}{string(data)}
    }

    return
}
