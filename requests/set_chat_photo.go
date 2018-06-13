package requests

import (
    "errors"
    "io"
    "strconv"
)

type SetChatPhoto struct {
    ChatID interface{}
    Photo  io.Reader
}

func (r *SetChatPhoto) IsMultipart() bool {
    return true
}

func (r *SetChatPhoto) GetValues() (values map[string]interface{}, err error) {
    values = make(map[string]interface{})

    switch chatID := r.ChatID.(type) {
    case uint64:
        values["chat_id"] = strconv.FormatUint(chatID, 10)
    case string:
        values["chat_id"] = chatID
    default:
        return nil, errors.New("invalid chat_id")
    }

    values["photo"] = r.Photo

    return
}
