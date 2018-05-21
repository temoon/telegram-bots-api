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

func (r *SetChatPhoto) GetValues() (values map[string]interface{}, err error) {
    values = make(map[string]interface{})

    switch r.ChatID.(type) {
    case uint64:
        values["chat_id"] = strconv.FormatUint(r.ChatID.(uint64), 10)
    case string:
        values["chat_id"] = r.ChatID.(string)
    default:
        return nil, errors.New("invalid chat_id")
    }

    values["photo"] = r.Photo

    return
}
