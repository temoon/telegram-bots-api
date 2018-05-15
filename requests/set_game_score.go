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

func (r *SetGameScore) GetValues() (values map[string][]interface{}, err error) {
    values = make(map[string][]interface{})

    values["user_id"] = []interface{}{strconv.Itoa(r.UserID)}
    values["score"] = []interface{}{strconv.Itoa(r.Score)}

    if r.Force {
        values["force"] = []interface{}{"1"}
    }

    if r.DisableEditMessage {
        values["disable_edit_message"] = []interface{}{"1"}
    }

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
