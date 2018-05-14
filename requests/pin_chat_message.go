package requests

import (
    "errors"
    "strconv"
)

type PinChatMessage struct {
    ChatID              interface{}
    MessageID           int
    DisableNotification bool
}

func (r *PinChatMessage) IsMultipart() bool {
    return false
}

func (r *PinChatMessage) GetValues() (values map[string][]interface{}, err error) {
    values = make(map[string][]interface{})

    switch r.ChatID.(type) {
    case int64:
        values["chat_id"] = []interface{}{strconv.FormatInt(r.ChatID.(int64), 10)}
    case string:
        values["chat_id"] = []interface{}{r.ChatID.(string)}
    default:
        return nil, errors.New("invalid chat_id")
    }

    values["message_id"] = []interface{}{strconv.Itoa(r.MessageID)}

    if r.DisableNotification {
        values["disable_notification"] = []interface{}{"1"}
    }

    return
}
