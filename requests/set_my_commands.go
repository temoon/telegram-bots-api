package requests

import (
	"encoding/json"
)

type SetMyCommands struct {
	Commands []interface{}
}

func (r *SetMyCommands) IsMultipart() bool {
	return false
}

func (r *SetMyCommands) GetValues() (values map[string]interface{}, err error) {
	values = make(map[string]interface{})

	if r.Commands != nil {
		var data []byte
		if data, err = json.Marshal(r.Commands); err != nil {
			return
		}

		values["commands"] = string(data)
	}

	return
}
