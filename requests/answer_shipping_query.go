package requests

import (
	"context"
	"encoding/json"

	"github.com/temoon/telegram-bots-api"
)

type AnswerShippingQuery struct {
	ErrorMessage    *string
	Ok              bool
	ShippingOptions []telegram.ShippingOption
	ShippingQueryId string
}

func (r *AnswerShippingQuery) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new(bool)
	err = b.CallMethod(ctx, "answerShippingQuery", r, response)
	return
}

func (r *AnswerShippingQuery) IsMultipart() (multipart bool) {
	return false
}

func (r *AnswerShippingQuery) GetValues() (values map[string]interface{}, err error) {
	values = make(map[string]interface{})

	if r.ErrorMessage != nil {
		values["error_message"] = *r.ErrorMessage
	}

	if r.Ok {
		values["ok"] = "1"
	} else {
		values["ok"] = "0"
	}

	if r.ShippingOptions != nil {
		var dataShippingOptions []byte
		if dataShippingOptions, err = json.Marshal(r.ShippingOptions); err != nil {
			return
		}

		values["shipping_options"] = string(dataShippingOptions)
	}

	values["shipping_query_id"] = r.ShippingQueryId

	return
}
