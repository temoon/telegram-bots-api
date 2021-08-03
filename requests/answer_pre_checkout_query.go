package requests

type AnswerPreCheckoutQuery struct {
	ErrorMessage       string
	Ok                 bool
	PreCheckoutQueryId string
}

func (r *AnswerPreCheckoutQuery) IsMultipart() bool {
	return false
}

func (r *AnswerPreCheckoutQuery) GetValues() (values map[string]interface{}, err error) {
	values = make(map[string]interface{})

	if r.ErrorMessage != "" {
		values["error_message"] = r.ErrorMessage
	}

	if r.Ok {
		values["ok"] = "1"
	}

	values["pre_checkout_query_id"] = r.PreCheckoutQueryId

	return
}
