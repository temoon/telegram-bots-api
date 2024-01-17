package requests

import (
	"context"
	"encoding/json"
	"github.com/temoon/telegram-bots-api"
	"io"
	"strconv"
)

type SendInvoice struct {
	SuggestedTipAmounts       []int64
	PhotoUrl                  *string
	NeedShippingAddress       *bool
	ReplyParameters           *telegram.ReplyParameters
	Prices                    []telegram.LabeledPrice
	PhotoSize                 *int64
	NeedPhoneNumber           *bool
	ProtectContent            *bool
	ProviderData              *string
	DisableNotification       *bool
	ProviderToken             string
	StartParameter            *string
	PhotoHeight               *int64
	SendPhoneNumberToProvider *bool
	MessageThreadId           *int64
	Description               string
	PhotoWidth                *int64
	NeedName                  *bool
	NeedEmail                 *bool
	IsFlexible                *bool
	ChatId                    telegram.ChatId
	Currency                  string
	SendEmailToProvider       *bool
	ReplyMarkup               *telegram.InlineKeyboardMarkup
	Title                     string
	Payload                   string
	MaxTipAmount              *int64
}

func (r *SendInvoice) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new(telegram.Message)
	err = b.CallMethod(ctx, "sendInvoice", r, response)
	return
}

func (r *SendInvoice) GetValues() (values map[string]interface{}, err error) {
	values = make(map[string]interface{})

	if r.SuggestedTipAmounts != nil {
		var dataSuggestedTipAmounts []byte
		if dataSuggestedTipAmounts, err = json.Marshal(r.SuggestedTipAmounts); err != nil {
			return
		}

		values["suggested_tip_amounts"] = string(dataSuggestedTipAmounts)
	}

	if r.PhotoUrl != nil {
		values["photo_url"] = *r.PhotoUrl
	}

	if r.NeedShippingAddress != nil {
		if *r.NeedShippingAddress {
			values["need_shipping_address"] = "1"
		} else {
			values["need_shipping_address"] = "0"
		}
	}

	if r.ReplyParameters != nil {
		var dataReplyParameters []byte
		if dataReplyParameters, err = json.Marshal(r.ReplyParameters); err != nil {
			return
		}

		values["reply_parameters"] = string(dataReplyParameters)
	}

	var dataPrices []byte
	if dataPrices, err = json.Marshal(r.Prices); err != nil {
		return
	}

	values["prices"] = string(dataPrices)

	if r.PhotoSize != nil {
		values["photo_size"] = strconv.FormatInt(*r.PhotoSize, 10)
	}

	if r.NeedPhoneNumber != nil {
		if *r.NeedPhoneNumber {
			values["need_phone_number"] = "1"
		} else {
			values["need_phone_number"] = "0"
		}
	}

	if r.ProtectContent != nil {
		if *r.ProtectContent {
			values["protect_content"] = "1"
		} else {
			values["protect_content"] = "0"
		}
	}

	if r.ProviderData != nil {
		values["provider_data"] = *r.ProviderData
	}

	if r.DisableNotification != nil {
		if *r.DisableNotification {
			values["disable_notification"] = "1"
		} else {
			values["disable_notification"] = "0"
		}
	}

	values["provider_token"] = r.ProviderToken

	if r.StartParameter != nil {
		values["start_parameter"] = *r.StartParameter
	}

	if r.PhotoHeight != nil {
		values["photo_height"] = strconv.FormatInt(*r.PhotoHeight, 10)
	}

	if r.SendPhoneNumberToProvider != nil {
		if *r.SendPhoneNumberToProvider {
			values["send_phone_number_to_provider"] = "1"
		} else {
			values["send_phone_number_to_provider"] = "0"
		}
	}

	if r.MessageThreadId != nil {
		values["message_thread_id"] = strconv.FormatInt(*r.MessageThreadId, 10)
	}

	values["description"] = r.Description

	if r.PhotoWidth != nil {
		values["photo_width"] = strconv.FormatInt(*r.PhotoWidth, 10)
	}

	if r.NeedName != nil {
		if *r.NeedName {
			values["need_name"] = "1"
		} else {
			values["need_name"] = "0"
		}
	}

	if r.NeedEmail != nil {
		if *r.NeedEmail {
			values["need_email"] = "1"
		} else {
			values["need_email"] = "0"
		}
	}

	if r.IsFlexible != nil {
		if *r.IsFlexible {
			values["is_flexible"] = "1"
		} else {
			values["is_flexible"] = "0"
		}
	}

	values["chat_id"] = r.ChatId.String()

	values["currency"] = r.Currency

	if r.SendEmailToProvider != nil {
		if *r.SendEmailToProvider {
			values["send_email_to_provider"] = "1"
		} else {
			values["send_email_to_provider"] = "0"
		}
	}

	if r.ReplyMarkup != nil {
		var dataReplyMarkup []byte
		if dataReplyMarkup, err = json.Marshal(r.ReplyMarkup); err != nil {
			return
		}

		values["reply_markup"] = string(dataReplyMarkup)
	}

	values["title"] = r.Title

	values["payload"] = r.Payload

	if r.MaxTipAmount != nil {
		values["max_tip_amount"] = strconv.FormatInt(*r.MaxTipAmount, 10)
	}

	return
}

func (r *SendInvoice) GetFiles() (files map[string]io.Reader) {
	return
}
