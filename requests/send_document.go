package requests

import (
    "encoding/json"
    "errors"
    "os"
    "strconv"
)

type SendDocument struct {
    ChatID              interface{}
    Document            interface{}
    Caption             string
    ParseMode           string
    DisableNotification bool
    ReplyToMessageID    uint32
    ReplyMarkup         interface{}
}

func (r *SendDocument) IsMultipart() bool {
    _, ok := r.Document.(*os.File)

    return ok
}

func (r *SendDocument) GetValues() (values map[string]interface{}, err error) {
    values = make(map[string]interface{})

    switch chatID := r.ChatID.(type) {
    case uint64:
        values["chat_id"] = strconv.FormatUint(chatID, 10)
    case string:
        values["chat_id"] = chatID
    default:
        return nil, errors.New("invalid chat_id")
    }

    switch document := r.Document.(type) {
    case string:
        values["document"] = document
    case *os.File:
        values["document"] = document
    default:
        return nil, errors.New("invalid document")
    }

    if r.Caption != "" {
        values["caption"] = r.Caption
    }

    if r.ParseMode != "" {
        values["parse_mode"] = r.ParseMode
    }

    if r.DisableNotification {
        values["disable_notification"] = "1"
    }

    if r.ReplyToMessageID != 0 {
        values["reply_to_message_id"] = strconv.FormatUint(uint64(r.ReplyToMessageID), 10)
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
