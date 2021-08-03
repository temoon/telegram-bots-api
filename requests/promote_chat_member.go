package requests

import (
	"strconv"
)

type PromoteChatMember struct {
	CanChangeInfo       bool
	CanDeleteMessages   bool
	CanEditMessages     bool
	CanInviteUsers      bool
	CanManageChat       bool
	CanManageVoiceChats bool
	CanPinMessages      bool
	CanPostMessages     bool
	CanPromoteMembers   bool
	CanRestrictMembers  bool
	ChatId              interface{}
	IsAnonymous         bool
	UserId              uint64
}

func (r *PromoteChatMember) IsMultipart() bool {
	return false
}

func (r *PromoteChatMember) GetValues() (values map[string]interface{}, err error) {
	values = make(map[string]interface{})

	if r.CanChangeInfo {
		values["can_change_info"] = "1"
	}

	if r.CanDeleteMessages {
		values["can_delete_messages"] = "1"
	}

	if r.CanEditMessages {
		values["can_edit_messages"] = "1"
	}

	if r.CanInviteUsers {
		values["can_invite_users"] = "1"
	}

	if r.CanManageChat {
		values["can_manage_chat"] = "1"
	}

	if r.CanManageVoiceChats {
		values["can_manage_voice_chats"] = "1"
	}

	if r.CanPinMessages {
		values["can_pin_messages"] = "1"
	}

	if r.CanPostMessages {
		values["can_post_messages"] = "1"
	}

	if r.CanPromoteMembers {
		values["can_promote_members"] = "1"
	}

	if r.CanRestrictMembers {
		values["can_restrict_members"] = "1"
	}

	switch value := r.ChatId.(type) {
	case uint64:
		values["chat_id"] = strconv.FormatUint(value, 10)
	case string:
		values["chat_id"] = value
	}

	if r.IsAnonymous {
		values["is_anonymous"] = "1"
	}

	values["user_id"] = strconv.FormatUint(r.UserId, 10)

	return
}
