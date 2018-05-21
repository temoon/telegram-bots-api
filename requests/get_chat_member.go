package requests

import (
    "errors"
    "strconv"
)

type GetChatMember struct {
    ChatID interface{}
    UserID uint32
}

func (r *GetChatMember) IsMultipart() bool {
    return false
}

func (r *GetChatMember) GetValues() (values map[string]interface{}, err error) {
    values = make(map[string]interface{})

    switch r.ChatID.(type) {
    case uint64:
        values["chat_id"] = strconv.FormatUint(r.ChatID.(uint64), 10)
    case string:
        values["chat_id"] = r.ChatID.(string)
    default:
        return nil, errors.New("invalid chat_id")
    }

    values["user_id"] = strconv.FormatUint(uint64(r.UserID), 10)

    return
}
