package requests

import (
    "encoding/json"
    "strconv"
)

type SendInvoice struct {
    ChatID                    int64
    Title                     string
    Description               string
    Payload                   string
    ProviderToken             string
    StartParameter            string
    Currency                  string
    Prices                    []interface{}
    ProviderData              string
    PhotoURL                  string
    PhotoSize                 int
    PhotoWidth                int
    PhotoHeight               int
    NeedName                  bool
    NeedPhoneNumber           bool
    NeedEmail                 bool
    NeedShippingAddress       bool
    SendPhoneNumberToProvider bool
    SendEmailToProvider       bool
    IsFlexible                bool
    DisableNotification       bool
    ReplyToMessageID          int
    ReplyMarkup               interface{}
}

func (r *SendInvoice) IsMultipart() bool {
    return false
}

func (r *SendInvoice) GetValues() (values map[string][]interface{}, err error) {
    values = make(map[string][]interface{})

    values["chat_id"] = []interface{}{strconv.FormatInt(r.ChatID, 10)}
    values["title"] = []interface{}{r.Title}
    values["description"] = []interface{}{r.Description}
    values["payload"] = []interface{}{r.Payload}
    values["provider_token"] = []interface{}{r.ProviderToken}
    values["start_parameter"] = []interface{}{r.StartParameter}
    values["currency"] = []interface{}{r.Currency}

    var data []byte
    if data, err = json.Marshal(r.Prices); err != nil {
        return
    }

    values["prices"] = []interface{}{string(data)}

    if r.ProviderData != "" {
        values["provider_data"] = []interface{}{r.ProviderData}
    }

    if r.PhotoURL != "" {
        values["photo_url"] = []interface{}{r.PhotoURL}
    }

    if r.PhotoSize != 0 {
        values["photo_size"] = []interface{}{strconv.Itoa(r.PhotoSize)}
    }

    if r.PhotoWidth != 0 {
        values["photo_width"] = []interface{}{strconv.Itoa(r.PhotoWidth)}
    }

    if r.PhotoHeight != 0 {
        values["photo_height"] = []interface{}{strconv.Itoa(r.PhotoHeight)}
    }

    if r.NeedName {
        values["need_name"] = []interface{}{"1"}
    }

    if r.NeedPhoneNumber {
        values["need_phone_number"] = []interface{}{"1"}
    }

    if r.NeedEmail {
        values["need_email"] = []interface{}{"1"}
    }

    if r.NeedShippingAddress {
        values["need_shipping_address"] = []interface{}{"1"}
    }

    if r.SendPhoneNumberToProvider {
        values["send_phone_number_to_provider"] = []interface{}{"1"}
    }

    if r.SendEmailToProvider {
        values["send_email_to_provider"] = []interface{}{"1"}
    }

    if r.IsFlexible {
        values["is_flexible"] = []interface{}{"1"}
    }

    if r.DisableNotification {
        values["disable_notification"] = []interface{}{"1"}
    }

    if r.ReplyToMessageID != 0 {
        values["reply_to_message_id"] = []interface{}{strconv.Itoa(r.ReplyToMessageID)}
    }

    if r.ReplyMarkup != nil {
        if data, err = json.Marshal(r.ReplyMarkup); err != nil {
            return
        }

        values["reply_markup"] = []interface{}{string(data)}
    }

    return
}
