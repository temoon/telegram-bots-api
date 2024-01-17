package requests

import (
	"context"
	"encoding/json"
	"github.com/temoon/telegram-bots-api"
	"io"
	"strconv"
)

type CreateNewStickerSet struct {
	StickerFormat   string
	StickerType     *string
	NeedsRepainting *bool
	UserId          int64
	Name            string
	Title           string
	Stickers        []telegram.InputSticker
}

func (r *CreateNewStickerSet) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new(bool)
	err = b.CallMethod(ctx, "createNewStickerSet", r, response)
	return
}

func (r *CreateNewStickerSet) GetValues() (values map[string]interface{}, err error) {
	values = make(map[string]interface{})

	values["sticker_format"] = r.StickerFormat

	if r.StickerType != nil {
		values["sticker_type"] = *r.StickerType
	}

	if r.NeedsRepainting != nil {
		if *r.NeedsRepainting {
			values["needs_repainting"] = "1"
		} else {
			values["needs_repainting"] = "0"
		}
	}

	values["user_id"] = strconv.FormatInt(r.UserId, 10)

	values["name"] = r.Name

	values["title"] = r.Title

	var dataStickers []byte
	if dataStickers, err = json.Marshal(r.Stickers); err != nil {
		return
	}

	values["stickers"] = string(dataStickers)

	return
}

func (r *CreateNewStickerSet) GetFiles() (files map[string]io.Reader) {
	files = make(map[string]io.Reader)

	for _, item := range r.Stickers {
		if item.Sticker.HasFile() {
			files[item.Sticker.GetFormFieldName()] = item.Sticker.GetFile()
		}
	}

	return
}
