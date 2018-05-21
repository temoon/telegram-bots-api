package requests

import (
    "strconv"
)

type GetGameHighScores struct {
    UserID          uint32
    ChatID          int64
    MessageID       uint32
    InlineMessageID string
}

func (r *GetGameHighScores) IsMultipart() bool {
    return false
}

func (r *GetGameHighScores) GetValues() (values map[string]interface{}, err error) {
    values = make(map[string]interface{})

    values["user_id"] = strconv.FormatUint(uint64(r.UserID), 10)

    if r.ChatID != 0 {
        values["chat_id"] = strconv.FormatInt(r.ChatID, 10)
    }

    if r.MessageID != 0 {
        values["message_id"] = strconv.FormatUint(uint64(r.MessageID), 10)
    }

    if r.InlineMessageID != "" {
        values["inline_message_id"] = r.InlineMessageID
    }

    return
}
