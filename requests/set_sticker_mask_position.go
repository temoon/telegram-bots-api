package requests

import (
"encoding/json"
"context"
	"github.com/temoon/telegram-bots-api"
)

type SetStickerMaskPosition struct {
MaskPosition *telegram.MaskPosition
Sticker string
}

func (r *SetStickerMaskPosition) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new(bool)
	err = b.CallMethod(ctx, "setStickerMaskPosition", r, response)
	return
}



func (r *SetStickerMaskPosition) IsMultipart() (multipart bool) {
	return false
	}

func (r *SetStickerMaskPosition) GetValues() (values map[string]interface{}, err error) {
	values = make(map[string]interface{})

	
			if r.MaskPosition != nil {
			var dataMaskPosition []byte
				if dataMaskPosition, err = json.Marshal(r.MaskPosition); err != nil {
					return
				}

				values["mask_position"] = string(dataMaskPosition)
			}
			
			values["sticker"] = r.Sticker
			

	return
}
