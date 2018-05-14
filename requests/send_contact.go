package requests

import (
    "encoding/json"
    "errors"
    "strconv"
)

type SendContact struct {
    ChatID              interface{}
    PhoneNumber         string
    FirstName           string
    LastName            string
    DisableNotification bool
    ReplyToMessageID    int
    ReplyMarkup         interface{}
}

func (r *SendContact) IsMultipart() bool {
    return false
}

func (r *SendContact) GetValues() (values map[string][]interface{}, err error) {
    values = make(map[string][]interface{})

    switch r.ChatID.(type) {
    case int64:
        values["chat_id"] = []interface{}{strconv.FormatInt(r.ChatID.(int64), 10)}
    case string:
        values["chat_id"] = []interface{}{r.ChatID.(string)}
    default:
        return nil, errors.New("invalid chat_id")
    }

    values["phone_number"] = []interface{}{r.PhoneNumber}
    values["first_name"] = []interface{}{r.FirstName}

    if r.LastName != "" {
        values["last_name"] = []interface{}{r.LastName}
    }

    if r.DisableNotification {
        values["disable_notification"] = []interface{}{"1"}
    }

    values["reply_to_message_id"] = []interface{}{strconv.Itoa(r.ReplyToMessageID)}

    if r.ReplyMarkup != nil {
        var data []byte
        if data, err = json.Marshal(r.ReplyMarkup); err != nil {
            return
        }

        values["reply_markup"] = []interface{}{string(data)}
    }

    return
}
