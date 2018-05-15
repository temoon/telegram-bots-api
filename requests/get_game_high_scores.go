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

func (r *GetGameHighScores) GetValues() (values map[string][]interface{}, err error) {
    values = make(map[string][]interface{})

    values["user_id"] = []interface{}{strconv.Itoa(r.UserID)}

    if r.ChatID != 0 {
        values["chat_id"] = []interface{}{strconv.FormatInt(r.ChatID, 10)}
    }

    if r.MessageID != 0 {
        values["message_id"] = []interface{}{strconv.Itoa(r.MessageID)}
    }

    if r.InlineMessageID != "" {
        values["inline_message_id"] = []interface{}{r.InlineMessageID}
    }

    return
}
