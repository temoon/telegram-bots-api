package requests

import (
    "encoding/json"
    "errors"
    "strconv"
)

type SendVenue struct {
    ChatID              interface{}
    Latitude            float64
    Longitude           float64
    Title               string
    Address             string
    FoursquareID        string
    DisableNotification bool
    ReplyToMessageID    int
    ReplyMarkup         interface{}
}

func (r *SendVenue) IsMultipart() bool {
    return false
}

func (r *SendVenue) GetValues() (values map[string][]interface{}, err error) {
    values = make(map[string][]interface{})

    switch r.ChatID.(type) {
    case int64:
        values["chat_id"] = []interface{}{strconv.FormatInt(r.ChatID.(int64), 10)}
    case string:
        values["chat_id"] = []interface{}{r.ChatID.(string)}
    default:
        return nil, errors.New("invalid chat_id")
    }

    values["latitude"] = []interface{}{strconv.FormatFloat(r.Latitude, 'f', -1, 64)}
    values["longitude"] = []interface{}{strconv.FormatFloat(r.Longitude, 'f', -1, 64)}
    values["title"] = []interface{}{r.Title}
    values["address"] = []interface{}{r.Address}

    if r.FoursquareID != "" {
        values["foursquare_id"] = []interface{}{r.FoursquareID}
    }

    if r.DisableNotification == true {
        values["disable_notification"] = []interface{}{"1"}
    }

    values["reply_to_message_id"] = []interface{}{strconv.Itoa(r.ReplyToMessageID)}

    if r.ReplyMarkup != nil {
        var data []byte
        if data, err = json.Marshal(r.ReplyMarkup); err != nil {
            return
        }

        values["reply_markup"] = []interface{}{string(data)}
    }

    return
}
