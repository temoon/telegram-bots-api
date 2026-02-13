package requests

import (
	"context"
	"io"
	"strconv"

	"github.com/temoon/telegram-bots-api"
)

type GetUserProfileAudios struct {
	UserId int64
	Limit  *int64
	Offset *int64
}

func (r *GetUserProfileAudios) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new(telegram.UserProfileAudios)
	err = b.CallMethod(ctx, "getUserProfileAudios", r, response)
	return
}

func (r *GetUserProfileAudios) GetValues() (values map[string]string, err error) {
	values = make(map[string]string)

	values["user_id"] = strconv.FormatInt(r.UserId, 10)

	if r.Limit != nil {
		values["limit"] = strconv.FormatInt(*r.Limit, 10)
	}

	if r.Offset != nil {
		values["offset"] = strconv.FormatInt(*r.Offset, 10)
	}

	return
}

func (r *GetUserProfileAudios) GetFiles() (files map[string]io.Reader) {
	return
}
