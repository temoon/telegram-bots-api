package requests

import (
    "errors"
    "strconv"
)

type GetChatMembersCount struct {
    ChatID interface{}
}

func (r *GetChatMembersCount) IsMultipart() bool {
    return false
}

func (r *GetChatMembersCount) GetValues() (values map[string]interface{}, err error) {
    values = make(map[string]interface{})

    switch r.ChatID.(type) {
    case uint64:
        values["chat_id"] = strconv.FormatUint(r.ChatID.(uint64), 10)
    case string:
        values["chat_id"] = r.ChatID.(string)
    default:
        return nil, errors.New("invalid chat_id")
    }

    return
}
