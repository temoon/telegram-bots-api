package requests

import (
	"encoding/json"
)

type AnswerShippingQuery struct {
	ErrorMessage    string
	Ok              bool
	ShippingOptions []interface{}
	ShippingQueryId string
}

func (r *AnswerShippingQuery) IsMultipart() bool {
	return false
}

func (r *AnswerShippingQuery) GetValues() (values map[string]interface{}, err error) {
	values = make(map[string]interface{})

	if r.ErrorMessage != "" {
		values["error_message"] = r.ErrorMessage
	}

	if r.Ok {
		values["ok"] = "1"
	}

	if r.ShippingOptions != nil {
		var data []byte
		if data, err = json.Marshal(r.ShippingOptions); err != nil {
			return
		}

		values["shipping_options"] = string(data)
	}

	values["shipping_query_id"] = r.ShippingQueryId

	return
}
