package requests

import (
    "encoding/json"
    "errors"
    "strconv"
)

type EditMessageText struct {
    ChatID                interface{}
    MessageID             uint32
    InlineMessageID       string
    Text                  string
    ParseMode             string
    DisableWebPagePreview bool
    ReplyMarkup           interface{}
}

func (r *EditMessageText) IsMultipart() bool {
    return false
}

func (r *EditMessageText) GetValues() (values map[string]interface{}, err error) {
    values = make(map[string]interface{})

    switch chatID := r.ChatID.(type) {
    case uint64:
        values["chat_id"] = strconv.FormatUint(chatID, 10)
    case string:
        values["chat_id"] = chatID
    default:
        return nil, errors.New("invalid chat_id")
    }

    if r.MessageID != 0 {
        values["message_id"] = strconv.FormatUint(uint64(r.MessageID), 10)
    }

    if r.InlineMessageID != "" {
        values["inline_message_id"] = r.InlineMessageID
    }

    values["text"] = r.Text

    if r.ParseMode != "" {
        values["parse_mode"] = r.ParseMode
    }

    if r.DisableWebPagePreview {
        values["disable_web_page_preview"] = "1"
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
