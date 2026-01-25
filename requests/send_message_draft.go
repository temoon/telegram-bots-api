package requests

import (
	"context"
	"encoding/json"
	"io"
	"strconv"

	"github.com/temoon/telegram-bots-api"
)

type SendMessageDraft struct {
	ChatId          int64
	DraftId         int64
	Text            string
	Entities        []telegram.MessageEntity
	MessageThreadId *int64
	ParseMode       *string
}

func (r *SendMessageDraft) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new(bool)
	err = b.CallMethod(ctx, "sendMessageDraft", r, response)
	return
}

func (r *SendMessageDraft) GetValues() (values map[string]string, err error) {
	values = make(map[string]string)

	values["chat_id"] = strconv.FormatInt(r.ChatId, 10)

	values["draft_id"] = strconv.FormatInt(r.DraftId, 10)

	values["text"] = r.Text

	if r.Entities != nil {
		var dataEntities []byte
		if dataEntities, err = json.Marshal(r.Entities); err != nil {
			return
		}

		values["entities"] = string(dataEntities)
	}

	if r.MessageThreadId != nil {
		values["message_thread_id"] = strconv.FormatInt(*r.MessageThreadId, 10)
	}

	if r.ParseMode != nil {
		values["parse_mode"] = *r.ParseMode
	}

	return
}

func (r *SendMessageDraft) GetFiles() (files map[string]io.Reader) {
	return
}
