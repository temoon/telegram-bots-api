package requests

import (
    "encoding/json"
    "errors"
    "os"
    "strconv"
)

type SendAudio struct {
    ChatID              interface{}
    Audio               interface{}
    Caption             string
    ParseMode           string
    Duration            int
    Performer           string
    Title               string
    DisableNotification bool
    ReplyToMessageID    int
    ReplyMarkup         interface{}
}

func (r *SendAudio) IsMultipart() bool {
    _, ok := r.Audio.(*os.File)

    return ok
}

func (r *SendAudio) GetValues() (values map[string]interface{}, err error) {
    values = make(map[string]interface{})

    switch r.ChatID.(type) {
    case int64:
        values["chat_id"] = strconv.FormatInt(r.ChatID.(int64), 10)
    case string:
        values["chat_id"] = r.ChatID.(string)
    default:
        return nil, errors.New("invalid chat_id")
    }

    switch r.Audio.(type) {
    case string:
        values["audio"] = r.Audio.(string)
    case *os.File:
        values["audio"] = r.Audio.(*os.File)
    default:
        return nil, errors.New("invalid audio")
    }

    if r.Caption != "" {
        values["caption"] = r.Caption
    }

    if r.ParseMode != "" {
        values["parse_mode"] = r.ParseMode
    }

    if r.Duration != 0 {
        values["duration"] = strconv.Itoa(r.Duration)
    }

    if r.Performer != "" {
        values["performer"] = r.Performer
    }

    if r.Title != "" {
        values["title"] = r.Title
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
