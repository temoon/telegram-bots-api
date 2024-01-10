package requests

import (
"errors"
"io"
"strconv"
"context"
	"github.com/temoon/telegram-bots-api"
)

type SetStickerSetThumbnail struct {
Name string
Thumbnail interface{}
UserId int32
}

func (r *SetStickerSetThumbnail) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new(bool)
	err = b.CallMethod(ctx, "setStickerSetThumbnail", r, response)
	return
}



func (r *SetStickerSetThumbnail) IsMultipart() (multipart bool) {
	return false
	}

func (r *SetStickerSetThumbnail) GetValues() (values map[string]interface{}, err error) {
	values = make(map[string]interface{})

	
			values["name"] = r.Name
			
			switch value := r.Thumbnail.(type) {
			case *int64:
					values["thumbnail"] = strconv.FormatInt(*value, 10)
				case *string:
					values["thumbnail"] = *value
				default:
				err = errors.New("invalid thumbnail field type")
				return
			}
		
			values["user_id"] = strconv.FormatInt(int64(r.UserId), 10)
			

	return
}
