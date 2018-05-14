package requests

import (
    "errors"
    "strconv"
)

const ChatActionTyping = "typing"
const ChatActionUploadPhoto = "upload_photo"
const ChatActionRecordVideo = "record_video"
const ChatActionUploadVideo = "upload_video"
const ChatActionRecordAudio = "record_audio"
const ChatActionUploadAudio = "upload_audio"
const ChatActionUploadDocument = "upload_document"
const ChatActionFindLocation = "find_location"
const ChatActionRecordVideoNote = "record_video_note"
const ChatActionUploadVideoNote = "upload_video_note"

type SendChatAction struct {
    ChatID interface{}
    Action string
}

func (r *SendChatAction) IsMultipart() bool {
    return false
}

func (r *SendChatAction) GetValues() (values map[string][]interface{}, err error) {
    values = make(map[string][]interface{})

    switch r.ChatID.(type) {
    case int64:
        values["chat_id"] = []interface{}{strconv.FormatInt(r.ChatID.(int64), 10)}
    case string:
        values["chat_id"] = []interface{}{r.ChatID.(string)}
    default:
        return nil, errors.New("invalid chat_id")
    }

    values["action"] = []interface{}{r.Action}

    return
}
