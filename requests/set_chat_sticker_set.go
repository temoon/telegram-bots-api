package requests

import (
	"strconv"
)

type SetChatStickerSet struct {
	ChatId         interface{}
	StickerSetName string
}

func (r *SetChatStickerSet) IsMultipart() bool {
	return false
}

func (r *SetChatStickerSet) GetValues() (values map[string]interface{}, err error) {
	values = make(map[string]interface{})

	switch value := r.ChatId.(type) {
	case uint64:
		values["chat_id"] = strconv.FormatUint(value, 10)
	case string:
		values["chat_id"] = value
	}

	values["sticker_set_name"] = r.StickerSetName

	return
}
