package requests

import (
    "errors"
    "strconv"
)

type SetChatStickerSet struct {
    ChatID         interface{}
    StickerSetName string
}

func (r *SetChatStickerSet) IsMultipart() bool {
    return false
}

func (r *SetChatStickerSet) GetValues() (values map[string]interface{}, err error) {
    values = make(map[string]interface{})

    switch r.ChatID.(type) {
    case uint64:
        values["chat_id"] = strconv.FormatUint(r.ChatID.(uint64), 10)
    case string:
        values["chat_id"] = r.ChatID.(string)
    default:
        return nil, errors.New("invalid chat_id")
    }

    values["sticker_set_name"] = r.StickerSetName

    return
}
