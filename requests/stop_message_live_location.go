package requests

import (
    "encoding/json"
    "errors"
    "strconv"
)

type StopMessageLiveLocation struct {
    ChatID          interface{}
    MessageID       uint32
    InlineMessageID string
    ReplyMarkup     interface{}
}

func (r *StopMessageLiveLocation) IsMultipart() bool {
    return false
}

func (r *StopMessageLiveLocation) GetValues() (values map[string]interface{}, err error) {
    values = make(map[string]interface{})

    if r.ChatID != nil {
        switch r.ChatID.(type) {
        case uint64:
            values["chat_id"] = strconv.FormatUint(r.ChatID.(uint64), 10)
        case string:
            values["chat_id"] = r.ChatID.(string)
        default:
            return nil, errors.New("invalid chat_id")
        }
    }

    if r.MessageID != 0 {
        values["message_id"] = strconv.FormatUint(uint64(r.MessageID), 10)
    }

    if r.InlineMessageID != "" {
        values["inline_message_id"] = r.InlineMessageID
    }

    if r.ReplyMarkup != nil {
        var data []byte
        if data, err = json.Marshal(r.ReplyMarkup); err != nil {
            return
        }

        values["reply_markup"] = string(data)
    }

    return
}
