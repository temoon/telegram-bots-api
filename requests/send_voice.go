package requests

import (
    "encoding/json"
    "errors"
    "os"
    "strconv"
)

type SendVoice struct {
    ChatID              interface{}
    Voice               interface{}
    Caption             string
    ParseMode           string
    Duration            int
    DisableNotification bool
    ReplyToMessageID    int
    ReplyMarkup         interface{}
}

func (r *SendVoice) IsMultipart() bool {
    _, ok := r.Voice.(*os.File)

    return ok
}

func (r *SendVoice) GetValues() (values map[string][]interface{}, err error) {
    values = make(map[string][]interface{})

    switch r.ChatID.(type) {
    case int64:
        values["chat_id"] = []interface{}{strconv.FormatInt(r.ChatID.(int64), 10)}
    case string:
        values["chat_id"] = []interface{}{r.ChatID.(string)}
    default:
        return nil, errors.New("invalid chat_id")
    }

    switch r.Voice.(type) {
    case string:
        values["voice"] = []interface{}{r.Voice.(string)}
    case *os.File:
        values["voice"] = []interface{}{r.Voice.(*os.File)}
    default:
        return nil, errors.New("invalid voice")
    }

    if r.Caption != "" {
        values["caption"] = []interface{}{r.Caption}
    }

    if r.ParseMode != "" {
        values["parse_mode"] = []interface{}{r.ParseMode}
    }

    if r.Duration != 0 {
        values["duration"] = []interface{}{strconv.Itoa(r.Duration)}
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
