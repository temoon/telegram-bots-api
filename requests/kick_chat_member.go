package requests

import (
    "errors"
    "strconv"
)

type KickChatMember struct {
    ChatID    interface{}
    UserID    int
    UntilDate int
}

func (r *KickChatMember) IsMultipart() bool {
    return false
}

func (r *KickChatMember) GetValues() (values map[string][]interface{}, err error) {
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

    if r.UntilDate != 0 {
        values["until_date"] = []interface{}{strconv.Itoa(r.UntilDate)}
    }

    return
}
