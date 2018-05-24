package requests

import (
    "encoding/json"
    "errors"
    "strconv"
)

type SendLocation struct {
    ChatID              interface{}
    Latitude            float64
    Longitude           float64
    LivePeriod          uint32
    DisableNotification bool
    ReplyToMessageID    uint32
    ReplyMarkup         interface{}
}

func (r *SendLocation) IsMultipart() bool {
    return false
}

func (r *SendLocation) GetValues() (values map[string]interface{}, err error) {
    values = make(map[string]interface{})

    switch chatID := r.ChatID.(type) {
    case uint64:
        values["chat_id"] = strconv.FormatUint(chatID, 10)
    case string:
        values["chat_id"] = chatID
    default:
        return nil, errors.New("invalid chat_id")
    }

    values["latitude"] = strconv.FormatFloat(r.Latitude, 'f', -1, 64)
    values["longitude"] = strconv.FormatFloat(r.Longitude, 'f', -1, 64)

    if r.LivePeriod != 0 {
        values["live_period"] = strconv.FormatUint(uint64(r.LivePeriod), 10)
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
