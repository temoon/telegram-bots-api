package requests

import (
    "encoding/json"
    "errors"
    "io"
    "strconv"
)

type SendAnimation struct {
    ChatID              interface{}
    Animation           interface{}
    Duration            uint32
    Width               uint32
    Height              uint32
    Thumb               interface{}
    Caption             string
    ParseMode           string
    DisableNotification bool
    ReplyToMessageID    uint32
    ReplyMarkup         interface{}
}

func (r *SendAnimation) IsMultipart() bool {
    _, okAnimation := r.Animation.(io.Reader)
    _, okThumb := r.Thumb.(io.Reader)

    return okAnimation || okThumb
}

func (r *SendAnimation) GetValues() (values map[string]interface{}, err error) {
    values = make(map[string]interface{})

    switch chatID := r.ChatID.(type) {
    case uint64:
        values["chat_id"] = strconv.FormatUint(chatID, 10)
    case string:
        values["chat_id"] = chatID
    default:
        return nil, errors.New("invalid chat_id")
    }

    switch animation := r.Animation.(type) {
    case string:
        values["animation"] = animation
    case io.Reader:
        values["animation"] = animation
    default:
        return nil, errors.New("invalid animation")
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

    switch thumb := r.Thumb.(type) {
    case string:
        values["thumb"] = thumb
    case io.Reader:
        values["thumb"] = thumb
    default:
        return nil, errors.New("invalid thumb")
    }

    if r.Caption != "" {
        values["caption"] = r.Caption
    }

    if r.ParseMode != "" {
        values["parse_mode"] = r.ParseMode
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
