package requests

import (
	"encoding/json"
	"strconv"
)

type SendInvoice struct {
	AllowSendingWithoutReply  bool
	ChatId                    uint64
	Currency                  string
	Description               string
	DisableNotification       bool
	IsFlexible                bool
	NeedEmail                 bool
	NeedName                  bool
	NeedPhoneNumber           bool
	NeedShippingAddress       bool
	Payload                   string
	PhotoHeight               uint64
	PhotoSize                 uint64
	PhotoUrl                  string
	PhotoWidth                uint64
	Prices                    []interface{}
	ProviderData              string
	ProviderToken             string
	ReplyMarkup               interface{}
	ReplyToMessageId          uint64
	SendEmailToProvider       bool
	SendPhoneNumberToProvider bool
	StartParameter            string
	Title                     string
}

func (r *SendInvoice) IsMultipart() bool {
	return false
}

func (r *SendInvoice) GetValues() (values map[string]interface{}, err error) {
	values = make(map[string]interface{})

	if r.AllowSendingWithoutReply {
		values["allow_sending_without_reply"] = "1"
	}

	values["chat_id"] = strconv.FormatUint(r.ChatId, 10)

	values["currency"] = r.Currency

	values["description"] = r.Description

	if r.DisableNotification {
		values["disable_notification"] = "1"
	}

	if r.IsFlexible {
		values["is_flexible"] = "1"
	}

	if r.NeedEmail {
		values["need_email"] = "1"
	}

	if r.NeedName {
		values["need_name"] = "1"
	}

	if r.NeedPhoneNumber {
		values["need_phone_number"] = "1"
	}

	if r.NeedShippingAddress {
		values["need_shipping_address"] = "1"
	}

	values["payload"] = r.Payload

	if r.PhotoHeight != 0 {
		values["photo_height"] = strconv.FormatUint(r.PhotoHeight, 10)
	}

	if r.PhotoSize != 0 {
		values["photo_size"] = strconv.FormatUint(r.PhotoSize, 10)
	}

	if r.PhotoUrl != "" {
		values["photo_url"] = r.PhotoUrl
	}

	if r.PhotoWidth != 0 {
		values["photo_width"] = strconv.FormatUint(r.PhotoWidth, 10)
	}

	if r.Prices != nil {
		var data []byte
		if data, err = json.Marshal(r.Prices); err != nil {
			return
		}

		values["prices"] = string(data)
	}

	if r.ProviderData != "" {
		values["provider_data"] = r.ProviderData
	}

	values["provider_token"] = r.ProviderToken

	if r.ReplyMarkup != nil {
		var data []byte
		if data, err = json.Marshal(r.ReplyMarkup); err != nil {
			return
		}

		values["reply_markup"] = string(data)
	}

	if r.ReplyToMessageId != 0 {
		values["reply_to_message_id"] = strconv.FormatUint(r.ReplyToMessageId, 10)
	}

	if r.SendEmailToProvider {
		values["send_email_to_provider"] = "1"
	}

	if r.SendPhoneNumberToProvider {
		values["send_phone_number_to_provider"] = "1"
	}

	values["start_parameter"] = r.StartParameter

	values["title"] = r.Title

	return
}
