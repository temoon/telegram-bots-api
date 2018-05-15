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

func (r *SendInvoice) GetValues() (values map[string]interface{}, err error) {
    values = make(map[string]interface{})

    values["chat_id"] = strconv.FormatInt(r.ChatID, 10)
    values["title"] = r.Title
    values["description"] = r.Description
    values["payload"] = r.Payload
    values["provider_token"] = r.ProviderToken
    values["start_parameter"] = r.StartParameter
    values["currency"] = r.Currency

    var data []byte
    if data, err = json.Marshal(r.Prices); err != nil {
        return
    }

    values["prices"] = string(data)

    if r.ProviderData != "" {
        values["provider_data"] = r.ProviderData
    }

    if r.PhotoURL != "" {
        values["photo_url"] = r.PhotoURL
    }

    if r.PhotoSize != 0 {
        values["photo_size"] = strconv.Itoa(r.PhotoSize)
    }

    if r.PhotoWidth != 0 {
        values["photo_width"] = strconv.Itoa(r.PhotoWidth)
    }

    if r.PhotoHeight != 0 {
        values["photo_height"] = strconv.Itoa(r.PhotoHeight)
    }

    if r.NeedName {
        values["need_name"] = "1"
    }

    if r.NeedPhoneNumber {
        values["need_phone_number"] = "1"
    }

    if r.NeedEmail {
        values["need_email"] = "1"
    }

    if r.NeedShippingAddress {
        values["need_shipping_address"] = "1"
    }

    if r.SendPhoneNumberToProvider {
        values["send_phone_number_to_provider"] = "1"
    }

    if r.SendEmailToProvider {
        values["send_email_to_provider"] = "1"
    }

    if r.IsFlexible {
        values["is_flexible"] = "1"
    }

    if r.DisableNotification {
        values["disable_notification"] = "1"
    }

    if r.ReplyToMessageID != 0 {
        values["reply_to_message_id"] = strconv.Itoa(r.ReplyToMessageID)
    }

    if r.ReplyMarkup != nil {
        if data, err = json.Marshal(r.ReplyMarkup); err != nil {
            return
        }

        values["reply_markup"] = string(data)
    }

    return
}
