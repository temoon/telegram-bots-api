package requests

import (
	"context"
	"encoding/json"
	"github.com/temoon/telegram-bots-api"
	"io"
	"strconv"
)

type CreateInvoiceLink struct {
	Description               string
	Currency                  string
	PhotoHeight               *int64
	NeedPhoneNumber           *bool
	SendPhoneNumberToProvider *bool
	SendEmailToProvider       *bool
	Title                     string
	Payload                   string
	ProviderData              *string
	NeedName                  *bool
	ProviderToken             string
	PhotoSize                 *int64
	PhotoWidth                *int64
	IsFlexible                *bool
	Prices                    []telegram.LabeledPrice
	MaxTipAmount              *int64
	SuggestedTipAmounts       []int64
	PhotoUrl                  *string
	NeedEmail                 *bool
	NeedShippingAddress       *bool
}

func (r *CreateInvoiceLink) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new(string)
	err = b.CallMethod(ctx, "createInvoiceLink", r, response)
	return
}

func (r *CreateInvoiceLink) GetValues() (values map[string]interface{}, err error) {
	values = make(map[string]interface{})

	values["description"] = r.Description

	values["currency"] = r.Currency

	if r.PhotoHeight != nil {
		values["photo_height"] = strconv.FormatInt(*r.PhotoHeight, 10)
	}

	if r.NeedPhoneNumber != nil {
		if *r.NeedPhoneNumber {
			values["need_phone_number"] = "1"
		} else {
			values["need_phone_number"] = "0"
		}
	}

	if r.SendPhoneNumberToProvider != nil {
		if *r.SendPhoneNumberToProvider {
			values["send_phone_number_to_provider"] = "1"
		} else {
			values["send_phone_number_to_provider"] = "0"
		}
	}

	if r.SendEmailToProvider != nil {
		if *r.SendEmailToProvider {
			values["send_email_to_provider"] = "1"
		} else {
			values["send_email_to_provider"] = "0"
		}
	}

	values["title"] = r.Title

	values["payload"] = r.Payload

	if r.ProviderData != nil {
		values["provider_data"] = *r.ProviderData
	}

	if r.NeedName != nil {
		if *r.NeedName {
			values["need_name"] = "1"
		} else {
			values["need_name"] = "0"
		}
	}

	values["provider_token"] = r.ProviderToken

	if r.PhotoSize != nil {
		values["photo_size"] = strconv.FormatInt(*r.PhotoSize, 10)
	}

	if r.PhotoWidth != nil {
		values["photo_width"] = strconv.FormatInt(*r.PhotoWidth, 10)
	}

	if r.IsFlexible != nil {
		if *r.IsFlexible {
			values["is_flexible"] = "1"
		} else {
			values["is_flexible"] = "0"
		}
	}

	var dataPrices []byte
	if dataPrices, err = json.Marshal(r.Prices); err != nil {
		return
	}

	values["prices"] = string(dataPrices)

	if r.MaxTipAmount != nil {
		values["max_tip_amount"] = strconv.FormatInt(*r.MaxTipAmount, 10)
	}

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

	if r.NeedEmail != nil {
		if *r.NeedEmail {
			values["need_email"] = "1"
		} else {
			values["need_email"] = "0"
		}
	}

	if r.NeedShippingAddress != nil {
		if *r.NeedShippingAddress {
			values["need_shipping_address"] = "1"
		} else {
			values["need_shipping_address"] = "0"
		}
	}

	return
}

func (r *CreateInvoiceLink) GetFiles() (files map[string]io.Reader) {
	return
}
