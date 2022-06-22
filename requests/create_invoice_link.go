package requests

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/temoon/telegram-bots-api"
)

type CreateInvoiceLink struct {
	Currency                  string
	Description               string
	IsFlexible                *bool
	MaxTipAmount              *int32
	NeedEmail                 *bool
	NeedName                  *bool
	NeedPhoneNumber           *bool
	NeedShippingAddress       *bool
	Payload                   string
	PhotoHeight               *int32
	PhotoSize                 *int32
	PhotoUrl                  *string
	PhotoWidth                *int32
	Prices                    []telegram.LabeledPrice
	ProviderData              *string
	ProviderToken             string
	SendEmailToProvider       *bool
	SendPhoneNumberToProvider *bool
	SuggestedTipAmounts       []int32
	Title                     string
}

func (r *CreateInvoiceLink) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new(string)
	err = b.CallMethod(ctx, "createInvoiceLink", r, response)
	return
}

func (r *CreateInvoiceLink) IsMultipart() (multipart bool) {
	return false
}

func (r *CreateInvoiceLink) GetValues() (values map[string]interface{}, err error) {
	values = make(map[string]interface{})

	values["currency"] = r.Currency

	values["description"] = r.Description

	if r.IsFlexible != nil {
		if *r.IsFlexible {
			values["is_flexible"] = "1"
		} else {
			values["is_flexible"] = "0"
		}
	}

	if r.MaxTipAmount != nil {
		values["max_tip_amount"] = strconv.FormatInt(int64(*r.MaxTipAmount), 10)
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

	values["payload"] = r.Payload

	if r.PhotoHeight != nil {
		values["photo_height"] = strconv.FormatInt(int64(*r.PhotoHeight), 10)
	}

	if r.PhotoSize != nil {
		values["photo_size"] = strconv.FormatInt(int64(*r.PhotoSize), 10)
	}

	if r.PhotoUrl != nil {
		values["photo_url"] = *r.PhotoUrl
	}

	if r.PhotoWidth != nil {
		values["photo_width"] = strconv.FormatInt(int64(*r.PhotoWidth), 10)
	}

	var dataPrices []byte
	if dataPrices, err = json.Marshal(r.Prices); err != nil {
		return
	}

	values["prices"] = string(dataPrices)

	if r.ProviderData != nil {
		values["provider_data"] = *r.ProviderData
	}

	values["provider_token"] = r.ProviderToken

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

	values["title"] = r.Title

	return
}
