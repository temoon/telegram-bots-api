package requests

import (
	"context"
	"strconv"

	"github.com/temoon/go-telegram-bots-api"
)

type PromoteChatMember struct {
	CanChangeInfo       *bool
	CanDeleteMessages   *bool
	CanEditMessages     *bool
	CanInviteUsers      *bool
	CanManageChat       *bool
	CanManageVideoChats *bool
	CanPinMessages      *bool
	CanPostMessages     *bool
	CanPromoteMembers   *bool
	CanRestrictMembers  *bool
	ChatId              interface{}
	IsAnonymous         *bool
	UserId              int64
}

func (r *PromoteChatMember) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new(bool)
	err = b.CallMethod(ctx, "promoteChatMember", r, response)
	return
}

func (r *PromoteChatMember) IsMultipart() (multipart bool) {
	return false
}

func (r *PromoteChatMember) GetValues() (values map[string]interface{}, err error) {
	values = make(map[string]interface{})

	if r.CanChangeInfo != nil {
		if *r.CanChangeInfo {
			values["can_change_info"] = "1"
		} else {
			values["can_change_info"] = "0"
		}
	}

	if r.CanDeleteMessages != nil {
		if *r.CanDeleteMessages {
			values["can_delete_messages"] = "1"
		} else {
			values["can_delete_messages"] = "0"
		}
	}

	if r.CanEditMessages != nil {
		if *r.CanEditMessages {
			values["can_edit_messages"] = "1"
		} else {
			values["can_edit_messages"] = "0"
		}
	}

	if r.CanInviteUsers != nil {
		if *r.CanInviteUsers {
			values["can_invite_users"] = "1"
		} else {
			values["can_invite_users"] = "0"
		}
	}

	if r.CanManageChat != nil {
		if *r.CanManageChat {
			values["can_manage_chat"] = "1"
		} else {
			values["can_manage_chat"] = "0"
		}
	}

	if r.CanManageVideoChats != nil {
		if *r.CanManageVideoChats {
			values["can_manage_video_chats"] = "1"
		} else {
			values["can_manage_video_chats"] = "0"
		}
	}

	if r.CanPinMessages != nil {
		if *r.CanPinMessages {
			values["can_pin_messages"] = "1"
		} else {
			values["can_pin_messages"] = "0"
		}
	}

	if r.CanPostMessages != nil {
		if *r.CanPostMessages {
			values["can_post_messages"] = "1"
		} else {
			values["can_post_messages"] = "0"
		}
	}

	if r.CanPromoteMembers != nil {
		if *r.CanPromoteMembers {
			values["can_promote_members"] = "1"
		} else {
			values["can_promote_members"] = "0"
		}
	}

	if r.CanRestrictMembers != nil {
		if *r.CanRestrictMembers {
			values["can_restrict_members"] = "1"
		} else {
			values["can_restrict_members"] = "0"
		}
	}

	switch value := r.ChatId.(type) {
	case int64:
		values["chat_id"] = strconv.FormatInt(value, 10)
	case string:
		values["chat_id"] = value
	}

	if r.IsAnonymous != nil {
		if *r.IsAnonymous {
			values["is_anonymous"] = "1"
		} else {
			values["is_anonymous"] = "0"
		}
	}

	values["user_id"] = strconv.FormatInt(r.UserId, 10)

	return
}
