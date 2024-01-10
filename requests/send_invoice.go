package requests

import (
"encoding/json"
"errors"
"strconv"
"context"
	"github.com/temoon/telegram-bots-api"
)

type SendInvoice struct {
ChatId interface{}
Currency string
Description string
DisableNotification *bool
IsFlexible *bool
MaxTipAmount *int32
MessageThreadId *int32
NeedEmail *bool
NeedName *bool
NeedPhoneNumber *bool
NeedShippingAddress *bool
Payload string
PhotoHeight *int32
PhotoSize *int32
PhotoUrl *string
PhotoWidth *int32
Prices []telegram.LabeledPrice
ProtectContent *bool
ProviderData *string
ProviderToken string
ReplyMarkup *telegram.InlineKeyboardMarkup
ReplyParameters *telegram.ReplyParameters
SendEmailToProvider *bool
SendPhoneNumberToProvider *bool
StartParameter *string
SuggestedTipAmounts []int32
Title string
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

	
			switch value := r.ChatId.(type) {
			case int64:
					values["chat_id"] = strconv.FormatInt(value, 10)
				case string:
					values["chat_id"] = value
				default:
				err = errors.New("invalid chat_id field type")
				return
			}
		
			values["currency"] = r.Currency
			
			values["description"] = r.Description
			
			if r.DisableNotification != nil {
			if *r.DisableNotification {
					values["disable_notification"] = "1"
				} else {
					values["disable_notification"] = "0"
				}
			}
			
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
			
			if r.MessageThreadId != nil {
			values["message_thread_id"] = strconv.FormatInt(int64(*r.MessageThreadId), 10)
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
			
			values["provider_token"] = r.ProviderToken
			
			if r.ReplyMarkup != nil {
			var dataReplyMarkup []byte
				if dataReplyMarkup, err = json.Marshal(r.ReplyMarkup); err != nil {
					return
				}

				values["reply_markup"] = string(dataReplyMarkup)
			}
			
			if r.ReplyParameters != nil {
			var dataReplyParameters []byte
				if dataReplyParameters, err = json.Marshal(r.ReplyParameters); err != nil {
					return
				}

				values["reply_parameters"] = string(dataReplyParameters)
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
			
			if r.StartParameter != nil {
			values["start_parameter"] = *r.StartParameter
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
