package requests

import (
    "encoding/json"
    "errors"
    "strconv"
)

type SendMessage struct {
    ChatID                interface{}
    Text                  string
    ParseMode             string
    DisableWebPagePreview bool
    DisableNotification   bool
    ReplyToMessageID      uint32
    ReplyMarkup           interface{}
}

func (r *SendMessage) IsMultipart() bool {
    return false
}

func (r *SendMessage) GetValues() (values map[string]interface{}, err error) {
    values = make(map[string]interface{})

    switch chatID := r.ChatID.(type) {
    case uint64:
        values["chat_id"] = strconv.FormatUint(chatID, 10)
    case string:
        values["chat_id"] = chatID
    default:
        return nil, errors.New("invalid chat_id")
    }

    values["text"] = r.Text

    if r.ParseMode != "" {
        values["parse_mode"] = r.ParseMode
    }

    if r.DisableWebPagePreview {
        values["disable_web_page_preview"] = "1"
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
