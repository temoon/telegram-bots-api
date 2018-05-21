package requests

import (
    "strconv"
)

type SetGameScore struct {
    UserID             uint32
    Score              uint32
    Force              bool
    DisableEditMessage bool
    ChatID             int64
    MessageID          uint32
    InlineMessageID    string
}

func (r *SetGameScore) IsMultipart() bool {
    return false
}

func (r *SetGameScore) GetValues() (values map[string]interface{}, err error) {
    values = make(map[string]interface{})

    values["user_id"] = strconv.FormatUint(uint64(r.UserID), 10)
    values["score"] = strconv.FormatUint(uint64(r.Score), 10)

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
        values["message_id"] = strconv.FormatUint(uint64(r.MessageID), 10)
    }

    if r.InlineMessageID != "" {
        values["inline_message_id"] = r.InlineMessageID
    }

    return
}
