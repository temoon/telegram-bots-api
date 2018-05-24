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
    Duration            uint32
    Performer           string
    Title               string
    DisableNotification bool
    ReplyToMessageID    uint32
    ReplyMarkup         interface{}
}

func (r *SendAudio) IsMultipart() bool {
    _, ok := r.Audio.(*os.File)

    return ok
}

func (r *SendAudio) GetValues() (values map[string]interface{}, err error) {
    values = make(map[string]interface{})

    switch chatID := r.ChatID.(type) {
    case uint64:
        values["chat_id"] = strconv.FormatUint(chatID, 10)
    case string:
        values["chat_id"] = chatID
    default:
        return nil, errors.New("invalid chat_id")
    }

    switch audio := r.Audio.(type) {
    case string:
        values["audio"] = audio
    case *os.File:
        values["audio"] = audio
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
        values["duration"] = strconv.FormatUint(uint64(r.Duration), 10)
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
