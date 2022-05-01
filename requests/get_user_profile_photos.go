package requests

import (
	"context"
	"strconv"

	"github.com/temoon/telegram-bots-api"
)

type GetUserProfilePhotos struct {
	Limit  *int32
	Offset *int32
	UserId int64
}

func (r *GetUserProfilePhotos) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new(telegram.UserProfilePhotos)
	err = b.CallMethod(ctx, "getUserProfilePhotos", r, response)
	return
}

func (r *GetUserProfilePhotos) IsMultipart() (multipart bool) {
	return false
}

func (r *GetUserProfilePhotos) GetValues() (values map[string]interface{}, err error) {
	values = make(map[string]interface{})

	if r.Limit != nil {
		values["limit"] = strconv.FormatInt(int64(*r.Limit), 10)
	}

	if r.Offset != nil {
		values["offset"] = strconv.FormatInt(int64(*r.Offset), 10)
	}

	values["user_id"] = strconv.FormatInt(r.UserId, 10)

	return
}
