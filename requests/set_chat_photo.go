package requests

import (
    "errors"
    "os"
    "strconv"
)

type SetChatPhoto struct {
    ChatID interface{}
    Photo  *os.File
}

func (r *SetChatPhoto) IsMultipart() bool {
    return true
}

func (r *SetChatPhoto) GetValues() (values map[string][]interface{}, err error) {
    values = make(map[string][]interface{})

    switch r.ChatID.(type) {
    case int64:
        values["chat_id"] = []interface{}{strconv.FormatInt(r.ChatID.(int64), 10)}
    case string:
        values["chat_id"] = []interface{}{r.ChatID.(string)}
    default:
        return nil, errors.New("invalid chat_id")
    }

    values["photo"] = []interface{}{r.Photo}

    return
}
