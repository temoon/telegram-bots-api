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

func (r *AnswerShippingQuery) GetValues() (values map[string][]interface{}, err error) {
    values = make(map[string][]interface{})

    values["shipping_query_id"] = []interface{}{r.ShippingQueryID}

    if r.OK {
        values["ok"] = []interface{}{"1"}
    } else {
        values["ok"] = []interface{}{"0"}
    }

    if r.ShippingOptions != nil {
        var data []byte
        if data, err = json.Marshal(r.ShippingOptions); err != nil {
            return
        }

        values["shipping_options"] = []interface{}{string(data)}
    }

    if r.ErrorMessage != "" {
        values["error_message"] = []interface{}{r.ErrorMessage}
    }

    return
}
