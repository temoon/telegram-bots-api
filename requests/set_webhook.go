package requests

import (
    "encoding/json"
    "os"
    "strconv"
)

type SetWebhook struct {
    URL            string
    Certificate    *os.File
    MaxConnections int
    AllowedUpdates []string
}

func (r *SetWebhook) IsMultipart() bool {
    return r.Certificate != nil
}

func (r *SetWebhook) GetValues() (values map[string][]interface{}, err error) {
    values = make(map[string][]interface{})

    values["url"] = []interface{}{r.URL}

    if r.Certificate != nil {
        values["certificate"] = []interface{}{r.Certificate}
    }

    if r.MaxConnections != 0 {
        values["max_connections"] = []interface{}{strconv.Itoa(r.MaxConnections)}
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