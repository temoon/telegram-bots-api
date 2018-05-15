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

func (r *SendVideo) GetValues() (values map[string]interface{}, err error) {
    values = make(map[string]interface{})

    switch r.ChatID.(type) {
    case int64:
        values["chat_id"] = strconv.FormatInt(r.ChatID.(int64), 10)
    case string:
        values["chat_id"] = r.ChatID.(string)
    default:
        return nil, errors.New("invalid chat_id")
    }

    switch r.Video.(type) {
    case string:
        values["video"] = r.Video.(string)
    case *os.File:
        values["video"] = r.Video.(*os.File)
    default:
        return nil, errors.New("invalid video")
    }

    if r.Duration != 0 {
        values["duration"] = strconv.Itoa(r.Duration)
    }

    if r.Width != 0 {
        values["width"] = strconv.Itoa(r.Width)
    }

    if r.Height != 0 {
        values["height"] = strconv.Itoa(r.Height)
    }

    if r.Caption != "" {
        values["caption"] = r.Caption
    }

    if r.ParseMode != "" {
        values["parse_mode"] = r.ParseMode
    }

    if r.SupportsStreaming {
        values["supports_streaming"] = "1"
    }

    if r.DisableNotification {
        values["disable_notification"] = "1"
    }

    if r.ReplyToMessageID != 0 {
        values["reply_to_message_id"] = strconv.Itoa(r.ReplyToMessageID)
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
