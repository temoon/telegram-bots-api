package requests

import (
    "strconv"
)

type SetGameScore struct {
    UserID             int
    Score              int
    Force              bool
    DisableEditMessage bool
    ChatID             int64
    MessageID          int
    InlineMessageID    string
}

func (r *SetGameScore) IsMultipart() bool {
    return false
}

func (r *SetGameScore) GetValues() (values map[string]interface{}, err error) {
    values = make(map[string]interface{})

    values["user_id"] = strconv.Itoa(r.UserID)
    values["score"] = strconv.Itoa(r.Score)

    if r.Force {
        values["force"] = "1"
    }

    if r.DisableEditMessage {
        values["disable_edit_message"] = "1"
    }

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
