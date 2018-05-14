package requests

import (
    "errors"
    "strconv"
)

type PromoteChatMember struct {
    ChatID             interface{}
    UserID             int
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

func (r *PromoteChatMember) GetValues() (values map[string][]interface{}, err error) {
    values = make(map[string][]interface{})

    switch r.ChatID.(type) {
    case int64:
        values["chat_id"] = []interface{}{strconv.FormatInt(r.ChatID.(int64), 10)}
    case string:
        values["chat_id"] = []interface{}{r.ChatID.(string)}
    default:
        return nil, errors.New("invalid chat_id")
    }

    values["user_id"] = []interface{}{strconv.Itoa(r.UserID)}

    if r.CanChangeInfo {
        values["can_change_info"] = []interface{}{"1"}
    }

    if r.CanPostMessages {
        values["can_post_messages"] = []interface{}{"1"}
    }

    if r.CanEditMessages {
        values["can_edit_messages"] = []interface{}{"1"}
    }

    if r.CanDeleteMessages {
        values["can_delete_messages"] = []interface{}{"1"}
    }

    if r.CanInviteUsers {
        values["can_invite_users"] = []interface{}{"1"}
    }

    if r.CanRestrictMembers {
        values["can_restrict_members"] = []interface{}{"1"}
    }

    if r.CanPinMessages {
        values["can_pin_messages"] = []interface{}{"1"}
    }

    if r.CanPromoteMembers {
        values["can_promote_members"] = []interface{}{"1"}
    }

    return
}