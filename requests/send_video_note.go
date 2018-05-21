package requests

import (
    "encoding/json"
    "errors"
    "os"
    "strconv"
)

type SendVideoNote struct {
    ChatID              interface{}
    VideoNote           interface{}
    Duration            uint32
    Length              uint32
    DisableNotification bool
    ReplyToMessageID    uint32
    ReplyMarkup         interface{}
}

func (r *SendVideoNote) IsMultipart() bool {
    _, ok := r.VideoNote.(*os.File)

    return ok
}

func (r *SendVideoNote) GetValues() (values map[string]interface{}, err error) {
    values = make(map[string]interface{})

    switch r.ChatID.(type) {
    case uint64:
        values["chat_id"] = strconv.FormatUint(r.ChatID.(uint64), 10)
    case string:
        values["chat_id"] = r.ChatID.(string)
    default:
        return nil, errors.New("invalid chat_id")
    }

    switch r.VideoNote.(type) {
    case string:
        values["video_note"] = r.VideoNote.(string)
    case *os.File:
        values["video_note"] = r.VideoNote.(*os.File)
    default:
        return nil, errors.New("invalid video_note")
    }

    if r.Duration != 0 {
        values["duration"] = strconv.FormatUint(uint64(r.Duration), 10)
    }

    if r.Length != 0 {
        values["length"] = strconv.FormatUint(uint64(r.Length), 10)
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
