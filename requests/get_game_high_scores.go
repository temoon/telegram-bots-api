package requests

import (
    "strconv"
)

type GetGameHighScores struct {
    UserID          int
    ChatID          int64
    MessageID       int
    InlineMessageID string
}

func (r *GetGameHighScores) IsMultipart() bool {
    return false
}

func (r *GetGameHighScores) GetValues() (values map[string]interface{}, err error) {
    values = make(map[string]interface{})

    values["user_id"] = strconv.Itoa(r.UserID)

    if r.ChatID != 0 {
        values["chat_id"] = strconv.FormatInt(r.ChatID, 10)
    }

    if r.MessageID != 0 {
        values["message_id"] = strconv.Itoa(r.MessageID)
    }

    if r.InlineMessageID != "" {
        values["inline_message_id"] = r.InlineMessageID
    }

    return
}
