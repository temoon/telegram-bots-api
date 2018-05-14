package requests

import (
    "encoding/json"
    "errors"
    "os"
    "strconv"
)

type SendVideo struct {
    ChatID              interface{}
    Video               interface{}
    Duration            int
    Width               int
    Height              int
    Caption             string
    ParseMode           string
    SupportsStreaming   bool
    DisableNotification bool
    ReplyToMessageID    int
    ReplyMarkup         interface{}
}

func (r *SendVideo) IsMultipart() bool {
    _, ok := r.Video.(*os.File)

    return ok
}

func (r *SendVideo) GetValues() (values map[string][]interface{}, err error) {
    values = make(map[string][]interface{})

    switch r.ChatID.(type) {
    case int64:
        values["chat_id"] = []interface{}{strconv.FormatInt(r.ChatID.(int64), 10)}
    case string:
        values["chat_id"] = []interface{}{r.ChatID.(string)}
    default:
        return nil, errors.New("invalid chat_id")
    }

    switch r.Video.(type) {
    case string:
        values["video"] = []interface{}{r.Video.(string)}
    case *os.File:
        values["video"] = []interface{}{r.Video.(*os.File)}
    default:
        return nil, errors.New("invalid video")
    }

    if r.Duration != 0 {
        values["duration"] = []interface{}{strconv.Itoa(r.Duration)}
    }

    if r.Width != 0 {
        values["width"] = []interface{}{strconv.Itoa(r.Width)}
    }

    if r.Height != 0 {
        values["height"] = []interface{}{strconv.Itoa(r.Height)}
    }

    if r.Caption != "" {
        values["caption"] = []interface{}{r.Caption}
    }

    if r.ParseMode != "" {
        values["parse_mode"] = []interface{}{r.ParseMode}
    }

    if r.SupportsStreaming == true {
        values["supports_streaming"] = []interface{}{"1"}
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
