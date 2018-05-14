package requests

import (
    "encoding/json"
    "errors"
    "os"
    "strconv"
)

type SendVideoNote struct {
    ChatID              interface{}
    VideoNote           interface{}
    Duration            int
    Length              int
    DisableNotification bool
    ReplyToMessageID    int
    ReplyMarkup         interface{}
}

func (r *SendVideoNote) IsMultipart() bool {
    _, ok := r.VideoNote.(*os.File)

    return ok
}

func (r *SendVideoNote) GetValues() (values map[string][]interface{}, err error) {
    values = make(map[string][]interface{})

    switch r.ChatID.(type) {
    case int64:
        values["chat_id"] = []interface{}{strconv.FormatInt(r.ChatID.(int64), 10)}
    case string:
        values["chat_id"] = []interface{}{r.ChatID.(string)}
    default:
        return nil, errors.New("invalid chat_id")
    }

    switch r.VideoNote.(type) {
    case string:
        values["video_note"] = []interface{}{r.VideoNote.(string)}
    case *os.File:
        values["video_note"] = []interface{}{r.VideoNote.(*os.File)}
    default:
        return nil, errors.New("invalid video_note")
    }

    if r.Duration != 0 {
        values["duration"] = []interface{}{strconv.Itoa(r.Duration)}
    }

    if r.Length != 0 {
        values["length"] = []interface{}{strconv.Itoa(r.Length)}
    }

    if r.DisableNotification {
        values["disable_notification"] = []interface{}{"1"}
    }

    if r.ReplyToMessageID != 0 {
        values["reply_to_message_id"] = []interface{}{strconv.Itoa(r.ReplyToMessageID)}
    }

    if r.ReplyMarkup != nil {
        var data []byte
        if data, err = json.Marshal(r.ReplyMarkup); err != nil {
            return
        }

        values["reply_markup"] = []interface{}{string(data)}
    }

    return
}
