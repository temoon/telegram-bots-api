package requests

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/temoon/telegram-bots-api"
)

type CreateNewStickerSet struct {
	Name            string
	NeedsRepainting *bool
	StickerFormat   string
	StickerType     *string
	Stickers        []telegram.InputSticker
	Title           string
	UserId          int64
}

func (r *CreateNewStickerSet) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new(bool)
	err = b.CallMethod(ctx, "createNewStickerSet", r, response)
	return
}

func (r *CreateNewStickerSet) IsMultipart() (multipart bool) {
	return true
}

func (r *CreateNewStickerSet) GetValues() (values map[string]interface{}, err error) {
	values = make(map[string]interface{})

	values["name"] = r.Name

	if r.NeedsRepainting != nil {
		if *r.NeedsRepainting {
			values["needs_repainting"] = "1"
		} else {
			values["needs_repainting"] = "0"
		}
	}

	values["sticker_format"] = r.StickerFormat

	if r.StickerType != nil {
		values["sticker_type"] = *r.StickerType
	}

	var dataStickers []byte
	if dataStickers, err = json.Marshal(r.Stickers); err != nil {
		return
	}

	values["stickers"] = string(dataStickers)

	values["title"] = r.Title

	values["user_id"] = strconv.FormatInt(r.UserId, 10)

	return
}
