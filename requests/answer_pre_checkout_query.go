package requests

type AnswerPreCheckoutQuery struct {
    PreCheckoutQueryID string
    OK                 bool
    ErrorMessage       string
}

func (r *AnswerPreCheckoutQuery) IsMultipart() bool {
    return false
}

func (r *AnswerPreCheckoutQuery) GetValues() (values map[string][]interface{}, err error) {
    values = make(map[string][]interface{})

    values["shipping_query_id"] = []interface{}{r.PreCheckoutQueryID}

    if r.OK {
        values["ok"] = []interface{}{"1"}
    } else {
        values["ok"] = []interface{}{"0"}
    }

    if r.ErrorMessage != "" {
        values["error_message"] = []interface{}{r.ErrorMessage}
    }

    return
}
