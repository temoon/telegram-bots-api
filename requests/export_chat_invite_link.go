package requests

import (
	"context"
	"github.com/temoon/telegram-bots-api"
	"io"
)

type ExportChatInviteLink struct {
	ChatId telegram.ChatId
}

func (r *ExportChatInviteLink) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new(string)
	err = b.CallMethod(ctx, "exportChatInviteLink", r, response)
	return
}

func (r *ExportChatInviteLink) GetValues() (values map[string]interface{}, err error) {
	values = make(map[string]interface{})

	values["chat_id"] = r.ChatId.String()

	return
}

func (r *ExportChatInviteLink) GetFiles() (files map[string]io.Reader) {
	return
}
