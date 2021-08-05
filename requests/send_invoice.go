package requests

import (
	"context"
	"encoding/json"
	"github.com/temoon/go-telegram-bots-api"
	"strconv"
)

type SendInvoice struct {
	AllowSendingWithoutReply  bool
	ChatId                    int64
	Currency                  string
	Description               string
	DisableNotification       bool
	IsFlexible                bool
	MaxTipAmount              int32
	NeedEmail                 bool
	NeedName                  bool
	NeedPhoneNumber           bool
	NeedShippingAddress       bool
	Payload                   string
	PhotoHeight               int32
	PhotoSize                 int32
	PhotoUrl                  string
	PhotoWidth                int32
	Prices                    []telegram.LabeledPrice
	ProviderData              string
	ProviderToken             string
	ReplyMarkup               *telegram.InlineKeyboardMarkup
	ReplyToMessageId          int32
	SendEmailToProvider       bool
	SendPhoneNumberToProvider bool
	StartParameter            string
	SuggestedTipAmounts       []int32
	Title                     string
}

func (r *SendInvoice) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new(telegram.Message)
	err = b.CallMethod(ctx, "sendInvoice", r, response)
	return
}

func (r *SendInvoice) IsMultipart() (multipart bool) {
	return false
}

func (r *SendInvoice) GetValues() (values map[string]interface{}, err error) {
	values = make(map[string]interface{})

	if r.AllowSendingWithoutReply {
		values["allow_sending_without_reply"] = "1"
	}

	values["chat_id"] = strconv.FormatInt(r.ChatId, 10)

	values["currency"] = r.Currency

	values["description"] = r.Description

	if r.DisableNotification {
		values["disable_notification"] = "1"
	}

	if r.IsFlexible {
		values["is_flexible"] = "1"
	}

	if r.MaxTipAmount != 0 {
		values["max_tip_amount"] = strconv.FormatInt(int64(r.MaxTipAmount), 10)
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
		values["photo_height"] = strconv.FormatInt(int64(r.PhotoHeight), 10)
	}

	if r.PhotoSize != 0 {
		values["photo_size"] = strconv.FormatInt(int64(r.PhotoSize), 10)
	}

	if r.PhotoUrl != "" {
		values["photo_url"] = r.PhotoUrl
	}

	if r.PhotoWidth != 0 {
		values["photo_width"] = strconv.FormatInt(int64(r.PhotoWidth), 10)
	}

	var dataPrices []byte
	if dataPrices, err = json.Marshal(r.Prices); err != nil {
		return
	}

	values["prices"] = string(dataPrices)

	if r.ProviderData != "" {
		values["provider_data"] = r.ProviderData
	}

	values["provider_token"] = r.ProviderToken

	if r.ReplyMarkup != nil {
		var dataReplyMarkup []byte
		if dataReplyMarkup, err = json.Marshal(r.ReplyMarkup); err != nil {
			return
		}

		values["reply_markup"] = string(dataReplyMarkup)
	}

	if r.ReplyToMessageId != 0 {
		values["reply_to_message_id"] = strconv.FormatInt(int64(r.ReplyToMessageId), 10)
	}

	if r.SendEmailToProvider {
		values["send_email_to_provider"] = "1"
	}

	if r.SendPhoneNumberToProvider {
		values["send_phone_number_to_provider"] = "1"
	}

	if r.StartParameter != "" {
		values["start_parameter"] = r.StartParameter
	}

	if r.SuggestedTipAmounts != nil {
		var dataSuggestedTipAmounts []byte
		if dataSuggestedTipAmounts, err = json.Marshal(r.SuggestedTipAmounts); err != nil {
			return
		}

		values["suggested_tip_amounts"] = string(dataSuggestedTipAmounts)
	}

	values["title"] = r.Title

	return
}
