package requests

import (
	"context"
	"github.com/temoon/telegram-bots-api"
	"io"
	"strconv"
)

type GetUserProfilePhotos struct {
	UserId int64
	Limit  *int64
	Offset *int64
}

func (r *GetUserProfilePhotos) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new(telegram.UserProfilePhotos)
	err = b.CallMethod(ctx, "getUserProfilePhotos", r, response)
	return
}

func (r *GetUserProfilePhotos) GetValues() (values map[string]interface{}, err error) {
	values = make(map[string]interface{})

	values["user_id"] = strconv.FormatInt(r.UserId, 10)

	if r.Limit != nil {
		values["limit"] = strconv.FormatInt(*r.Limit, 10)
	}

	if r.Offset != nil {
		values["offset"] = strconv.FormatInt(*r.Offset, 10)
	}

	return
}

func (r *GetUserProfilePhotos) GetFiles() (files map[string]io.Reader) {
	return
}
