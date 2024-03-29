package requests

import (
	"context"
	"encoding/json"
	"github.com/temoon/telegram-bots-api"
	"io"
)

type SetStickerKeywords struct {
	Sticker  string
	Keywords []string
}

func (r *SetStickerKeywords) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new(bool)
	err = b.CallMethod(ctx, "setStickerKeywords", r, response)
	return
}

func (r *SetStickerKeywords) GetValues() (values map[string]interface{}, err error) {
	values = make(map[string]interface{})

	values["sticker"] = r.Sticker

	if r.Keywords != nil {
		var dataKeywords []byte
		if dataKeywords, err = json.Marshal(r.Keywords); err != nil {
			return
		}

		values["keywords"] = string(dataKeywords)
	}

	return
}

func (r *SetStickerKeywords) GetFiles() (files map[string]io.Reader) {
	return
}
