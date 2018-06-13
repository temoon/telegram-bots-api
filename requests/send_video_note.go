package requests

import (
    "encoding/json"
    "errors"
    "io"
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
    _, ok := r.VideoNote.(io.Reader)

    return ok
}

func (r *SendVideoNote) GetValues() (values map[string]interface{}, err error) {
    values = make(map[string]interface{})

    switch chatID := r.ChatID.(type) {
    case uint64:
        values["chat_id"] = strconv.FormatUint(chatID, 10)
    case string:
        values["chat_id"] = chatID
    default:
        return nil, errors.New("invalid chat_id")
    }

    switch videoNote := r.VideoNote.(type) {
    case string:
        values["video_note"] = videoNote
    case io.Reader:
        values["video_note"] = videoNote
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
