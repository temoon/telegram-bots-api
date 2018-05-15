package requests

import (
    "encoding/json"
)

type AnswerShippingQuery struct {
    ShippingQueryID string
    OK              bool
    ShippingOptions []interface{}
    ErrorMessage    string
}

func (r *AnswerShippingQuery) IsMultipart() bool {
    return false
}

func (r *AnswerShippingQuery) GetValues() (values map[string]interface{}, err error) {
    values = make(map[string]interface{})

    values["shipping_query_id"] = r.ShippingQueryID

    if r.OK {
        values["ok"] = "1"
    } else {
        values["ok"] = "0"
    }

    if r.ShippingOptions != nil {
        var data []byte
        if data, err = json.Marshal(r.ShippingOptions); err != nil {
            return
        }

        values["shipping_options"] = string(data)
    }

    if r.ErrorMessage != "" {
        values["error_message"] = r.ErrorMessage
    }

    return
}
