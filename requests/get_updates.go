package requests

import (
	"encoding/json"
	"strconv"
)

type GetUpdates struct {
	AllowedUpdates []string
	Limit          uint64
	Offset         uint64
	Timeout        uint64
}

func (r *GetUpdates) IsMultipart() bool {
	return false
}

func (r *GetUpdates) GetValues() (values map[string]interface{}, err error) {
	values = make(map[string]interface{})

	if r.AllowedUpdates != nil {
		var data []byte
		if data, err = json.Marshal(r.AllowedUpdates); err != nil {
			return
		}

		values["allowed_updates"] = string(data)
	}

	if r.Limit != 0 {
		values["limit"] = strconv.FormatUint(r.Limit, 10)
	}

	if r.Offset != 0 {
		values["offset"] = strconv.FormatUint(r.Offset, 10)
	}

	if r.Timeout != 0 {
		values["timeout"] = strconv.FormatUint(r.Timeout, 10)
	}

	return
}
