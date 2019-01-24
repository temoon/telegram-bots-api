package requests

import (
    "encoding/json"
    "strconv"
)

type SetPassportDataErrors struct {
    UserID uint32
    Errors []interface{}
}

func (r *SetPassportDataErrors) IsMultipart() bool {
    return false
}

func (r *SetPassportDataErrors) GetValues() (values map[string]interface{}, err error) {
    values = make(map[string]interface{})

    values["user_id"] = strconv.FormatUint(uint64(r.UserID), 10)

    var data []byte
    if data, err = json.Marshal(r.Errors); err != nil {
        return
    }

    values["errors"] = string(data)

    return
}
