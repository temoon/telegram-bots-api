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
    ReplyToMessageID    uint32
    ReplyMarkup         interface{}
}

func (r *SendContact) IsMultipart() bool {
    return false
}

func (r *SendContact) GetValues() (values map[string]interface{}, err error) {
    values = make(map[string]interface{})

    switch chatID := r.ChatID.(type) {
    case uint64:
        values["chat_id"] = strconv.FormatUint(chatID, 10)
    case string:
        values["chat_id"] = chatID
    default:
        return nil, errors.New("invalid chat_id")
    }

    values["phone_number"] = r.PhoneNumber
    values["first_name"] = r.FirstName

    if r.LastName != "" {
        values["last_name"] = r.LastName
    }

    if r.DisableNotification {
        values["disable_notification"] = "1"
    }

    values["reply_to_message_id"] = strconv.FormatUint(uint64(r.ReplyToMessageID), 10)

    if r.ReplyMarkup != nil {
        var data []byte
        if data, err = json.Marshal(r.ReplyMarkup); err != nil {
            return
        }

        values["reply_markup"] = string(data)
    }

    return
}
