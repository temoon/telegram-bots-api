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
    ReplyToMessageID      int
    ReplyMarkup           interface{}
}

func (r *SendMessage) IsMultipart() bool {
    return false
}

func (r *SendMessage) GetValues() (values map[string]interface{}, err error) {
    values = make(map[string]interface{})

    switch r.ChatID.(type) {
    case int64:
        values["chat_id"] = strconv.FormatInt(r.ChatID.(int64), 10)
    case string:
        values["chat_id"] = r.ChatID.(string)
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
