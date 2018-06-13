package requests

import (
    "encoding/json"
    "errors"
    "io"
    "strconv"
)

type SendVideo struct {
    ChatID              interface{}
    Video               interface{}
    Duration            uint32
    Width               uint32
    Height              uint32
    Caption             string
    ParseMode           string
    SupportsStreaming   bool
    DisableNotification bool
    ReplyToMessageID    uint32
    ReplyMarkup         interface{}
}

func (r *SendVideo) IsMultipart() bool {
    _, ok := r.Video.(io.Reader)

    return ok
}

func (r *SendVideo) GetValues() (values map[string]interface{}, err error) {
    values = make(map[string]interface{})

    switch chatID := r.ChatID.(type) {
    case uint64:
        values["chat_id"] = strconv.FormatUint(chatID, 10)
    case string:
        values["chat_id"] = chatID
    default:
        return nil, errors.New("invalid chat_id")
    }

    switch video := r.Video.(type) {
    case string:
        values["video"] = video
    case io.Reader:
        values["video"] = video
    default:
        return nil, errors.New("invalid video")
    }

    if r.Duration != 0 {
        values["duration"] = strconv.FormatUint(uint64(r.Duration), 10)
    }

    if r.Width != 0 {
        values["width"] = strconv.FormatUint(uint64(r.Width), 10)
    }

    if r.Height != 0 {
        values["height"] = strconv.FormatUint(uint64(r.Height), 10)
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
