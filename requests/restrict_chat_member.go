package requests

import (
    "errors"
    "strconv"
)

type RestrictChatMember struct {
    ChatID                interface{}
    UserID                uint32
    UntilDate             uint32
    CanSendMessages       bool
    CanSendMediaMessages  bool
    CanSendOtherMessages  bool
    CanAddWebPagePreviews bool
}

func (r *RestrictChatMember) IsMultipart() bool {
    return false
}

func (r *RestrictChatMember) GetValues() (values map[string]interface{}, err error) {
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

    if r.UntilDate != 0 {
        values["until_date"] = strconv.FormatUint(uint64(r.UntilDate), 10)
    }

    if r.CanSendMessages {
        values["can_send_messages"] = "1"
    }

    if r.CanSendMediaMessages {
        values["can_send_media_messages"] = "1"
    }

    if r.CanSendOtherMessages {
        values["can_send_other_messages"] = "1"
    }

    if r.CanAddWebPagePreviews {
        values["can_add_web_page_previews"] = "1"
    }

    return
}
