package requests

import (
    "encoding/json"
    "strconv"
)

type GetUpdates struct {
    Offset         uint32
    Limit          uint32
    Timeout        uint32
    AllowedUpdates []string
}

func (r *GetUpdates) IsMultipart() bool {
    return false
}

func (r *GetUpdates) GetValues() (values map[string]interface{}, err error) {
    values = make(map[string]interface{})

    if r.Offset != 0 {
        values["offset"] = strconv.FormatUint(uint64(r.Offset), 10)
    }

    if r.Limit != 0 {
        values["limit"] = strconv.FormatUint(uint64(r.Limit), 10)
    }

    if r.Timeout != 0 {
        values["timeout"] = strconv.FormatUint(uint64(r.Timeout), 10)
    }

    if r.AllowedUpdates != nil {
        var data []byte
        if data, err = json.Marshal(r.AllowedUpdates); err != nil {
            return
        }

        values["allowed_updates"] = string(data)
    }

    return
}
