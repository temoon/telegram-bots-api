package requests

import (
"context"
	"github.com/temoon/telegram-bots-api"
)

type GetStickerSet struct {
Name string
}

func (r *GetStickerSet) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new(telegram.StickerSet)
	err = b.CallMethod(ctx, "getStickerSet", r, response)
	return
}



func (r *GetStickerSet) IsMultipart() (multipart bool) {
	return false
	}

func (r *GetStickerSet) GetValues() (values map[string]interface{}, err error) {
	values = make(map[string]interface{})

	
			values["name"] = r.Name
			

	return
}
