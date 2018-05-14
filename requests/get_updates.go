package requests

import (
    "encoding/json"
    "strconv"
)

type GetUpdates struct {
    Offset         int
    Limit          int
    Timeout        int
    AllowedUpdates []string
}

func (r *GetUpdates) IsMultipart() bool {
    return false
}

func (r *GetUpdates) GetValues() (values map[string][]interface{}, err error) {
    values = make(map[string][]interface{})

    if r.Offset != 0 {
        values["offset"] = []interface{}{strconv.Itoa(r.Offset)}
    }

    if r.Limit != 0 {
        values["limit"] = []interface{}{strconv.Itoa(r.Limit)}
    }

    if r.Timeout != 0 {
        values["timeout"] = []interface{}{strconv.Itoa(r.Timeout)}
    }

    if r.AllowedUpdates != nil {
        var data []byte
        if data, err = json.Marshal(r.AllowedUpdates); err != nil {
            return
        }

        values["allowed_updates"] = []interface{}{string(data)}
    }

    return
}
