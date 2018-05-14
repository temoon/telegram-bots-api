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
    ReplyToMessageID    int
    ReplyMarkup         interface{}
}

func (r *SendDocument) IsMultipart() bool {
    _, ok := r.Document.(*os.File)

    return ok
}

func (r *SendDocument) GetValues() (values map[string][]interface{}, err error) {
    values = make(map[string][]interface{})

    switch r.ChatID.(type) {
    case int64:
        values["chat_id"] = []interface{}{strconv.FormatInt(r.ChatID.(int64), 10)}
    case string:
        values["chat_id"] = []interface{}{r.ChatID.(string)}
    default:
        return nil, errors.New("invalid chat_id")
    }

    switch r.Document.(type) {
    case string:
        values["document"] = []interface{}{r.Document.(string)}
    case *os.File:
        values["document"] = []interface{}{r.Document.(*os.File)}
    default:
        return nil, errors.New("invalid document")
    }

    if r.Caption != "" {
        values["caption"] = []interface{}{r.Caption}
    }

    if r.ParseMode != "" {
        values["parse_mode"] = []interface{}{r.ParseMode}
    }

    if r.DisableNotification {
        values["disable_notification"] = []interface{}{"1"}
    }

    if r.ReplyToMessageID != 0 {
        values["reply_to_message_id"] = []interface{}{strconv.Itoa(r.ReplyToMessageID)}
    }

    if r.ReplyMarkup != nil {
        var data []byte
        if data, err = json.Marshal(r.ReplyMarkup); err != nil {
            return
        }

        values["reply_markup"] = []interface{}{string(data)}
    }

    return
}
