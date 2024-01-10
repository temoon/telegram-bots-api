package requests

import (
"strconv"
"context"
	"github.com/temoon/telegram-bots-api"
)

type SetStickerPositionInSet struct {
Position int32
Sticker string
}

func (r *SetStickerPositionInSet) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new(bool)
	err = b.CallMethod(ctx, "setStickerPositionInSet", r, response)
	return
}



func (r *SetStickerPositionInSet) IsMultipart() (multipart bool) {
	return false
	}

func (r *SetStickerPositionInSet) GetValues() (values map[string]interface{}, err error) {
	values = make(map[string]interface{})

	
			values["position"] = strconv.FormatInt(int64(r.Position), 10)
			
			values["sticker"] = r.Sticker
			

	return
}
