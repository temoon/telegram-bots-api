package requests

import (
    "errors"
    "strconv"
)

type KickChatMember struct {
    ChatID    interface{}
    UserID    uint32
    UntilDate uint32
}

func (r *KickChatMember) IsMultipart() bool {
    return false
}

func (r *KickChatMember) GetValues() (values map[string]interface{}, err error) {
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

    return
}
