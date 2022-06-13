package requests

import (
	"context"
	"encoding/json"

	"github.com/temoon/telegram-bots-api"
)

type DeleteMyCommands struct {
	LanguageCode *string
	Scope        interface{}
}

func (r *DeleteMyCommands) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new([]telegram.BotCommand)
	err = b.CallMethod(ctx, "deleteMyCommands", r, response)
	return
}

func (r *DeleteMyCommands) IsMultipart() (multipart bool) {
	return false
}

func (r *DeleteMyCommands) GetValues() (values map[string]interface{}, err error) {
	values = make(map[string]interface{})

	if r.LanguageCode != nil {
		values["language_code"] = *r.LanguageCode
	}

	switch value := r.Scope.(type) {
	case *telegram.BotCommandScopeDefault:
		if value != nil {
			var dataBotCommandScopeDefault []byte
			if dataBotCommandScopeDefault, err = json.Marshal(value); err != nil {
				return
			}

			values["scope"] = string(dataBotCommandScopeDefault)
		}
	case *telegram.BotCommandScopeAllPrivateChats:
		if value != nil {
			var dataBotCommandScopeAllPrivateChats []byte
			if dataBotCommandScopeAllPrivateChats, err = json.Marshal(value); err != nil {
				return
			}

			values["scope"] = string(dataBotCommandScopeAllPrivateChats)
		}
	case *telegram.BotCommandScopeAllGroupChats:
		if value != nil {
			var dataBotCommandScopeAllGroupChats []byte
			if dataBotCommandScopeAllGroupChats, err = json.Marshal(value); err != nil {
				return
			}

			values["scope"] = string(dataBotCommandScopeAllGroupChats)
		}
	case *telegram.BotCommandScopeAllChatAdministrators:
		if value != nil {
			var dataBotCommandScopeAllChatAdministrators []byte
			if dataBotCommandScopeAllChatAdministrators, err = json.Marshal(value); err != nil {
				return
			}

			values["scope"] = string(dataBotCommandScopeAllChatAdministrators)
		}
	case *telegram.BotCommandScopeChat:
		if value != nil {
			var dataBotCommandScopeChat []byte
			if dataBotCommandScopeChat, err = json.Marshal(value); err != nil {
				return
			}

			values["scope"] = string(dataBotCommandScopeChat)
		}
	case *telegram.BotCommandScopeChatAdministrators:
		if value != nil {
			var dataBotCommandScopeChatAdministrators []byte
			if dataBotCommandScopeChatAdministrators, err = json.Marshal(value); err != nil {
				return
			}

			values["scope"] = string(dataBotCommandScopeChatAdministrators)
		}
	case *telegram.BotCommandScopeChatMember:
		if value != nil {
			var dataBotCommandScopeChatMember []byte
			if dataBotCommandScopeChatMember, err = json.Marshal(value); err != nil {
				return
			}

			values["scope"] = string(dataBotCommandScopeChatMember)
		}
	}

	return
}
