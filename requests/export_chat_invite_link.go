package requests

import (
	"context"
	"io"

	"github.com/temoon/telegram-bots-api"
)

type ExportChatInviteLink struct {
	ChatId telegram.ChatId
}

func (r *ExportChatInviteLink) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new(string)
	err = b.CallMethod(ctx, "exportChatInviteLink", r, response)
	return
}

func (r *ExportChatInviteLink) GetValues() (values map[string]string, err error) {
	values = make(map[string]string)

	values["chat_id"] = r.ChatId.String()

	return
}

func (r *ExportChatInviteLink) GetFiles() (files map[string]io.Reader) {
	return
}
