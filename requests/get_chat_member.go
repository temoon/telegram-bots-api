package requests

import (
    "errors"
    "strconv"
)

type GetChatMember struct {
    ChatID interface{}
    UserID int
}

func (r *GetChatMember) IsMultipart() bool {
    return false
}

func (r *GetChatMember) GetValues() (values map[string][]interface{}, err error) {
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

    return
}
