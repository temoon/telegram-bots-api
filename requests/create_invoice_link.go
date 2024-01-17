package requests

import (
	"context"
	"encoding/json"
	"github.com/temoon/telegram-bots-api"
	"io"
	"strconv"
)

type CreateInvoiceLink struct {
	Currency                  string
	Description               string
	Payload                   string
	Prices                    []telegram.LabeledPrice
	ProviderToken             string
	Title                     string
	IsFlexible                *bool
	MaxTipAmount              *int64
	NeedEmail                 *bool
	NeedName                  *bool
	NeedPhoneNumber           *bool
	NeedShippingAddress       *bool
	PhotoHeight               *int64
	PhotoSize                 *int64
	PhotoUrl                  *string
	PhotoWidth                *int64
	ProviderData              *string
	SendEmailToProvider       *bool
	SendPhoneNumberToProvider *bool
	SuggestedTipAmounts       []int64
}

func (r *CreateInvoiceLink) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new(string)
	err = b.CallMethod(ctx, "createInvoiceLink", r, response)
	return
}

func (r *CreateInvoiceLink) GetValues() (values map[string]interface{}, err error) {
	values = make(map[string]interface{})

	values["currency"] = r.Currency

	values["description"] = r.Description

	values["payload"] = r.Payload

	var dataPrices []byte
	if dataPrices, err = json.Marshal(r.Prices); err != nil {
		return
	}

	values["prices"] = string(dataPrices)

	values["provider_token"] = r.ProviderToken

	values["title"] = r.Title

	if r.IsFlexible != nil {
		if *r.IsFlexible {
			values["is_flexible"] = "1"
		} else {
			values["is_flexible"] = "0"
		}
	}

	if r.MaxTipAmount != nil {
		values["max_tip_amount"] = strconv.FormatInt(*r.MaxTipAmount, 10)
	}

	if r.NeedEmail != nil {
		if *r.NeedEmail {
			values["need_email"] = "1"
		} else {
			values["need_email"] = "0"
		}
	}

	if r.NeedName != nil {
		if *r.NeedName {
			values["need_name"] = "1"
		} else {
			values["need_name"] = "0"
		}
	}

	if r.NeedPhoneNumber != nil {
		if *r.NeedPhoneNumber {
			values["need_phone_number"] = "1"
		} else {
			values["need_phone_number"] = "0"
		}
	}

	if r.NeedShippingAddress != nil {
		if *r.NeedShippingAddress {
			values["need_shipping_address"] = "1"
		} else {
			values["need_shipping_address"] = "0"
		}
	}

	if r.PhotoHeight != nil {
		values["photo_height"] = strconv.FormatInt(*r.PhotoHeight, 10)
	}

	if r.PhotoSize != nil {
		values["photo_size"] = strconv.FormatInt(*r.PhotoSize, 10)
	}

	if r.PhotoUrl != nil {
		values["photo_url"] = *r.PhotoUrl
	}

	if r.PhotoWidth != nil {
		values["photo_width"] = strconv.FormatInt(*r.PhotoWidth, 10)
	}

	if r.ProviderData != nil {
		values["provider_data"] = *r.ProviderData
	}

	if r.SendEmailToProvider != nil {
		if *r.SendEmailToProvider {
			values["send_email_to_provider"] = "1"
		} else {
			values["send_email_to_provider"] = "0"
		}
	}

	if r.SendPhoneNumberToProvider != nil {
		if *r.SendPhoneNumberToProvider {
			values["send_phone_number_to_provider"] = "1"
		} else {
			values["send_phone_number_to_provider"] = "0"
		}
	}

	if r.SuggestedTipAmounts != nil {
		var dataSuggestedTipAmounts []byte
		if dataSuggestedTipAmounts, err = json.Marshal(r.SuggestedTipAmounts); err != nil {
			return
		}

		values["suggested_tip_amounts"] = string(dataSuggestedTipAmounts)
	}

	return
}

func (r *CreateInvoiceLink) GetFiles() (files map[string]io.Reader) {
	return
}
