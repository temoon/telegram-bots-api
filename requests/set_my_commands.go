package requests

import (
	"context"
	"encoding/json"

	"github.com/temoon/telegram-bots-api"
)

type SetMyCommands struct {
	Commands     []telegram.BotCommand
	LanguageCode *string
	Scope        interface{}
}

func (r *SetMyCommands) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new(bool)
	err = b.CallMethod(ctx, "setMyCommands", r, response)
	return
}

func (r *SetMyCommands) IsMultipart() (multipart bool) {
	return false
}

func (r *SetMyCommands) GetValues() (values map[string]interface{}, err error) {
	values = make(map[string]interface{})

	var dataCommands []byte
	if dataCommands, err = json.Marshal(r.Commands); err != nil {
		return
	}

	values["commands"] = string(dataCommands)

	if r.LanguageCode != nil {
		values["language_code"] = *r.LanguageCode
	}

	switch value := r.Scope.(type) {
	case *telegram.BotCommandScopeDefault:
		var dataBotCommandScopeDefault []byte
		if dataBotCommandScopeDefault, err = json.Marshal(value); err != nil {
			return
		}

		values["scope"] = string(dataBotCommandScopeDefault)
	case *telegram.BotCommandScopeAllPrivateChats:
		var dataBotCommandScopeAllPrivateChats []byte
		if dataBotCommandScopeAllPrivateChats, err = json.Marshal(value); err != nil {
			return
		}

		values["scope"] = string(dataBotCommandScopeAllPrivateChats)
	case *telegram.BotCommandScopeAllGroupChats:
		var dataBotCommandScopeAllGroupChats []byte
		if dataBotCommandScopeAllGroupChats, err = json.Marshal(value); err != nil {
			return
		}

		values["scope"] = string(dataBotCommandScopeAllGroupChats)
	case *telegram.BotCommandScopeAllChatAdministrators:
		var dataBotCommandScopeAllChatAdministrators []byte
		if dataBotCommandScopeAllChatAdministrators, err = json.Marshal(value); err != nil {
			return
		}

		values["scope"] = string(dataBotCommandScopeAllChatAdministrators)
	case *telegram.BotCommandScopeChat:
		var dataBotCommandScopeChat []byte
		if dataBotCommandScopeChat, err = json.Marshal(value); err != nil {
			return
		}

		values["scope"] = string(dataBotCommandScopeChat)
	case *telegram.BotCommandScopeChatAdministrators:
		var dataBotCommandScopeChatAdministrators []byte
		if dataBotCommandScopeChatAdministrators, err = json.Marshal(value); err != nil {
			return
		}

		values["scope"] = string(dataBotCommandScopeChatAdministrators)
	case *telegram.BotCommandScopeChatMember:
		var dataBotCommandScopeChatMember []byte
		if dataBotCommandScopeChatMember, err = json.Marshal(value); err != nil {
			return
		}

		values["scope"] = string(dataBotCommandScopeChatMember)
	}

	return
}
