package requests

import (
    "encoding/json"
    "errors"
    "os"
    "strconv"
)

type SendPhoto struct {
    ChatID              interface{}
    Photo               interface{}
    Caption             string
    ParseMode           string
    DisableNotification bool
    ReplyToMessageID    int
    ReplyMarkup         interface{}
}

func (r *SendPhoto) IsMultipart() bool {
    _, ok := r.Photo.(*os.File)

    return ok
}

func (r *SendPhoto) GetValues() (values map[string][]interface{}, err error) {
    values = make(map[string][]interface{})

    switch r.ChatID.(type) {
    case int64:
        values["chat_id"] = []interface{}{strconv.FormatInt(r.ChatID.(int64), 10)}
    case string:
        values["chat_id"] = []interface{}{r.ChatID.(string)}
    default:
        return nil, errors.New("invalid chat_id")
    }

    switch r.Photo.(type) {
    case string:
        values["photo"] = []interface{}{r.Photo.(string)}
    case *os.File:
        values["photo"] = []interface{}{r.Photo.(*os.File)}
    default:
        return nil, errors.New("invalid photo")
    }

    if r.Caption != "" {
        values["caption"] = []interface{}{r.Caption}
    }

    if r.ParseMode != "" {
        values["parse_mode"] = []interface{}{r.ParseMode}
    }

    if r.DisableNotification == true {
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
