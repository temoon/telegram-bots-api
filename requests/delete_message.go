package requests

import (
    "errors"
    "strconv"
)

type DeleteMessage struct {
    ChatID    interface{}
    MessageID uint32
}

func (r *DeleteMessage) IsMultipart() bool {
    return false
}

func (r *DeleteMessage) GetValues() (values map[string]interface{}, err error) {
    values = make(map[string]interface{})

    switch chatID := r.ChatID.(type) {
    case uint64:
        values["chat_id"] = strconv.FormatUint(chatID, 10)
    case string:
        values["chat_id"] = chatID
    default:
        return nil, errors.New("invalid chat_id")
    }

    if r.MessageID != 0 {
        values["message_id"] = strconv.FormatUint(uint64(r.MessageID), 10)
    }

    return
}
