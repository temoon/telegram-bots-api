package requests

import (
    "encoding/json"
    "errors"
    "io"
    "strconv"
)

type SendVoice struct {
    ChatID              interface{}
    Voice               interface{}
    Caption             string
    ParseMode           string
    Duration            uint32
    DisableNotification bool
    ReplyToMessageID    uint32
    ReplyMarkup         interface{}
}

func (r *SendVoice) IsMultipart() bool {
    _, ok := r.Voice.(io.Reader)

    return ok
}

func (r *SendVoice) GetValues() (values map[string]interface{}, err error) {
    values = make(map[string]interface{})

    switch chatID := r.ChatID.(type) {
    case uint64:
        values["chat_id"] = strconv.FormatUint(chatID, 10)
    case string:
        values["chat_id"] = chatID
    default:
        return nil, errors.New("invalid chat_id")
    }

    switch voice := r.Voice.(type) {
    case string:
        values["voice"] = voice
    case io.Reader:
        values["voice"] = voice
    default:
        return nil, errors.New("invalid voice")
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
