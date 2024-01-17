package requests

import (
	"context"
	"encoding/json"
	"github.com/temoon/telegram-bots-api"
	"io"
)

type AnswerShippingQuery struct {
	Ok              bool
	ShippingQueryId string
	ErrorMessage    *string
	ShippingOptions []telegram.ShippingOption
}

func (r *AnswerShippingQuery) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new(bool)
	err = b.CallMethod(ctx, "answerShippingQuery", r, response)
	return
}

func (r *AnswerShippingQuery) GetValues() (values map[string]interface{}, err error) {
	values = make(map[string]interface{})

	if r.Ok {
		values["ok"] = "1"
	} else {
		values["ok"] = "0"
	}

	values["shipping_query_id"] = r.ShippingQueryId

	if r.ErrorMessage != nil {
		values["error_message"] = *r.ErrorMessage
	}

	if r.ShippingOptions != nil {
		var dataShippingOptions []byte
		if dataShippingOptions, err = json.Marshal(r.ShippingOptions); err != nil {
			return
		}

		values["shipping_options"] = string(dataShippingOptions)
	}

	return
}

func (r *AnswerShippingQuery) GetFiles() (files map[string]io.Reader) {
	return
}
