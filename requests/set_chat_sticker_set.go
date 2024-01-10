package requests

import (
"errors"
"strconv"
"context"
	"github.com/temoon/telegram-bots-api"
)

type SetChatStickerSet struct {
ChatId interface{}
StickerSetName string
}

func (r *SetChatStickerSet) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new(bool)
	err = b.CallMethod(ctx, "setChatStickerSet", r, response)
	return
}



func (r *SetChatStickerSet) IsMultipart() (multipart bool) {
	return false
	}

func (r *SetChatStickerSet) GetValues() (values map[string]interface{}, err error) {
	values = make(map[string]interface{})

	
			switch value := r.ChatId.(type) {
			case int64:
					values["chat_id"] = strconv.FormatInt(value, 10)
				case string:
					values["chat_id"] = value
				default:
				err = errors.New("invalid chat_id field type")
				return
			}
		
			values["sticker_set_name"] = r.StickerSetName
			

	return
}
