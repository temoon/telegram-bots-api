package requests

import (
	"context"
	"io"

	"github.com/temoon/telegram-bots-api"
)

type RemoveMyProfilePhoto struct {
}

func (r *RemoveMyProfilePhoto) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new(bool)
	err = b.CallMethod(ctx, "removeMyProfilePhoto", r, response)
	return
}

func (r *RemoveMyProfilePhoto) GetValues() (values map[string]string, err error) {
	return
}

func (r *RemoveMyProfilePhoto) GetFiles() (files map[string]io.Reader) {
	return
}
