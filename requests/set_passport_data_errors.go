package requests

import (
	"encoding/json"
	"strconv"
)

type SetPassportDataErrors struct {
	Errors []interface{}
	UserId uint64
}

func (r *SetPassportDataErrors) IsMultipart() bool {
	return false
}

func (r *SetPassportDataErrors) GetValues() (values map[string]interface{}, err error) {
	values = make(map[string]interface{})

	if r.Errors != nil {
		var data []byte
		if data, err = json.Marshal(r.Errors); err != nil {
			return
		}

		values["errors"] = string(data)
	}

	values["user_id"] = strconv.FormatUint(r.UserId, 10)

	return
}
