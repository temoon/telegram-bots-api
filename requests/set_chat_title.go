package requests

import (
    "errors"
    "strconv"
)

type SetChatTitle struct {
    ChatID interface{}
    Title  string
}

func (r *SetChatTitle) IsMultipart() bool {
    return false
}

func (r *SetChatTitle) GetValues() (values map[string]interface{}, err error) {
    values = make(map[string]interface{})

    switch chatID := r.ChatID.(type) {
    case uint64:
        values["chat_id"] = strconv.FormatUint(chatID, 10)
    case string:
        values["chat_id"] = chatID
    default:
        return nil, errors.New("invalid chat_id")
    }

    values["title"] = r.Title

    return
}
