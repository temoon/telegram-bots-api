package requests

import (
	"context"
	"github.com/temoon/telegram-bots-api"
	"io"
	"strconv"
)

type PromoteChatMember struct {
	UserId              int64
	CanManageChat       *bool
	CanEditMessages     *bool
	CanInviteUsers      *bool
	CanPostMessages     *bool
	CanManageTopics     *bool
	CanDeleteMessages   *bool
	CanManageVideoChats *bool
	CanPromoteMembers   *bool
	CanChangeInfo       *bool
	ChatId              telegram.ChatId
	CanRestrictMembers  *bool
	CanDeleteStories    *bool
	IsAnonymous         *bool
	CanPinMessages      *bool
	CanPostStories      *bool
	CanEditStories      *bool
}

func (r *PromoteChatMember) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new(bool)
	err = b.CallMethod(ctx, "promoteChatMember", r, response)
	return
}

func (r *PromoteChatMember) GetValues() (values map[string]interface{}, err error) {
	values = make(map[string]interface{})

	values["user_id"] = strconv.FormatInt(r.UserId, 10)

	if r.CanManageChat != nil {
		if *r.CanManageChat {
			values["can_manage_chat"] = "1"
		} else {
			values["can_manage_chat"] = "0"
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

	if r.CanPostMessages != nil {
		if *r.CanPostMessages {
			values["can_post_messages"] = "1"
		} else {
			values["can_post_messages"] = "0"
		}
	}

	if r.CanManageTopics != nil {
		if *r.CanManageTopics {
			values["can_manage_topics"] = "1"
		} else {
			values["can_manage_topics"] = "0"
		}
	}

	if r.CanDeleteMessages != nil {
		if *r.CanDeleteMessages {
			values["can_delete_messages"] = "1"
		} else {
			values["can_delete_messages"] = "0"
		}
	}

	if r.CanManageVideoChats != nil {
		if *r.CanManageVideoChats {
			values["can_manage_video_chats"] = "1"
		} else {
			values["can_manage_video_chats"] = "0"
		}
	}

	if r.CanPromoteMembers != nil {
		if *r.CanPromoteMembers {
			values["can_promote_members"] = "1"
		} else {
			values["can_promote_members"] = "0"
		}
	}

	if r.CanChangeInfo != nil {
		if *r.CanChangeInfo {
			values["can_change_info"] = "1"
		} else {
			values["can_change_info"] = "0"
		}
	}

	values["chat_id"] = r.ChatId.String()

	if r.CanRestrictMembers != nil {
		if *r.CanRestrictMembers {
			values["can_restrict_members"] = "1"
		} else {
			values["can_restrict_members"] = "0"
		}
	}

	if r.CanDeleteStories != nil {
		if *r.CanDeleteStories {
			values["can_delete_stories"] = "1"
		} else {
			values["can_delete_stories"] = "0"
		}
	}

	if r.IsAnonymous != nil {
		if *r.IsAnonymous {
			values["is_anonymous"] = "1"
		} else {
			values["is_anonymous"] = "0"
		}
	}

	if r.CanPinMessages != nil {
		if *r.CanPinMessages {
			values["can_pin_messages"] = "1"
		} else {
			values["can_pin_messages"] = "0"
		}
	}

	if r.CanPostStories != nil {
		if *r.CanPostStories {
			values["can_post_stories"] = "1"
		} else {
			values["can_post_stories"] = "0"
		}
	}

	if r.CanEditStories != nil {
		if *r.CanEditStories {
			values["can_edit_stories"] = "1"
		} else {
			values["can_edit_stories"] = "0"
		}
	}

	return
}

func (r *PromoteChatMember) GetFiles() (files map[string]io.Reader) {
	return
}
