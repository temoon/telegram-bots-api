package telegram

type SendMessageRequest struct {
    ChatID                interface{} `json:"chat_id"`
    Text                  string      `json:"text"`
    ParseMode             string      `json:"parse_mode,omitempty"`
    DisableWebPagePreview bool        `json:"disable_web_page_preview,omitempty"`
    DisableNotification   bool        `json:"disable_notification,omitempty"`
    ReplyToMessageID      int         `json:"reply_to_message_id,omitempty"`
    ReplyMarkup           interface{} `json:"reply_markup,omitempty"`
}

type ForwardMessageRequest struct {
    ChatID              interface{} `json:"chat_id"`
    FromChatID          interface{} `json:"from_chat_id"`
    DisableNotification bool        `json:"disable_notification,omitempty"`
    MessageID           int         `json:"message_id"`
}
