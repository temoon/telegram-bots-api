package requests

import (
	"context"
	"encoding/json"
	"github.com/temoon/go-telegram-bots-api"
)

type GetMyCommands struct {
	LanguageCode string
	Scope        interface{}
}

func (r *GetMyCommands) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new([]telegram.BotCommand)
	err = b.CallMethod(ctx, "getMyCommands", r, response)
	return
}

func (r *GetMyCommands) IsMultipart() (multipart bool) {
	return false
}

func (r *GetMyCommands) GetValues() (values map[string]interface{}, err error) {
	values = make(map[string]interface{})

	if r.LanguageCode != "" {
		values["language_code"] = r.LanguageCode
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
