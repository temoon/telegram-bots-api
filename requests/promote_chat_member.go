package requests

import (
    "errors"
    "strconv"
)

type PromoteChatMember struct {
    ChatID             interface{}
    UserID             uint32
    CanChangeInfo      bool
    CanPostMessages    bool
    CanEditMessages    bool
    CanDeleteMessages  bool
    CanInviteUsers     bool
    CanRestrictMembers bool
    CanPinMessages     bool
    CanPromoteMembers  bool
}

func (r *PromoteChatMember) IsMultipart() bool {
    return false
}

func (r *PromoteChatMember) GetValues() (values map[string]interface{}, err error) {
    values = make(map[string]interface{})

    switch chatID := r.ChatID.(type) {
    case uint64:
        values["chat_id"] = strconv.FormatUint(chatID, 10)
    case string:
        values["chat_id"] = chatID
    default:
        return nil, errors.New("invalid chat_id")
    }

    values["user_id"] = strconv.FormatUint(uint64(r.UserID), 10)

    if r.CanChangeInfo {
        values["can_change_info"] = "1"
    }

    if r.CanPostMessages {
        values["can_post_messages"] = "1"
    }

    if r.CanEditMessages {
        values["can_edit_messages"] = "1"
    }

    if r.CanDeleteMessages {
        values["can_delete_messages"] = "1"
    }

    if r.CanInviteUsers {
        values["can_invite_users"] = "1"
    }

    if r.CanRestrictMembers {
        values["can_restrict_members"] = "1"
    }

    if r.CanPinMessages {
        values["can_pin_messages"] = "1"
    }

    if r.CanPromoteMembers {
        values["can_promote_members"] = "1"
    }

    return
}
