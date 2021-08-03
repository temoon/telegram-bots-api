package requests

import (
	"encoding/json"
	"io"
	"strconv"
)

type SetWebhook struct {
	AllowedUpdates     []string
	Certificate        interface{}
	DropPendingUpdates bool
	IpAddress          string
	MaxConnections     uint64
	Url                string
}

func (r *SetWebhook) IsMultipart() bool {
	return true
}

func (r *SetWebhook) GetValues() (values map[string]interface{}, err error) {
	values = make(map[string]interface{})

	if r.AllowedUpdates != nil {
		var data []byte
		if data, err = json.Marshal(r.AllowedUpdates); err != nil {
			return
		}

		values["allowed_updates"] = string(data)
	}

	if r.Certificate != nil {
		values["certificate"] = r.Certificate
	}

	if r.DropPendingUpdates {
		values["drop_pending_updates"] = "1"
	}

	if r.IpAddress != "" {
		values["ip_address"] = r.IpAddress
	}

	if r.MaxConnections != 0 {
		values["max_connections"] = strconv.FormatUint(r.MaxConnections, 10)
	}

	values["url"] = r.Url

	return
}
