package telegram

import (
	"encoding/json"
	"strconv"
)

type ChatId struct {
	id   int64
	name string
}

//goland:noinspection GoUnusedExportedFunction
func NewChatId(id int64, name string) (chatId ChatId) {
	if id != 0 {
		chatId.id = id
	} else if len(name) > 0 {
		chatId.name = name
	}

	return
}

//goland:noinspection GoMixedReceiverTypes
func (c *ChatId) GetId() int64 {
	return c.id
}

//goland:noinspection GoMixedReceiverTypes
func (c *ChatId) GetName() string {
	return c.name
}

//goland:noinspection GoMixedReceiverTypes
func (c ChatId) MarshalJSON() (data []byte, err error) {
	return json.Marshal(c.String())
}

//goland:noinspection GoMixedReceiverTypes
func (c *ChatId) UnmarshalJSON(data []byte) (err error) {
	if len(data) >= 2 && data[0] == '"' && data[len(data)-1] == '"' { // len(`""`) == 2
		data = data[1 : len(data)-1]
	}

	if len(data) > 0 {
		if data[0] == '@' {
			c.name = string(data)
		} else {
			c.id, err = strconv.ParseInt(string(data), 10, 64)
		}
	}

	return
}

//goland:noinspection GoMixedReceiverTypes
func (c *ChatId) String() string {
	if c.id != 0 {
		return strconv.FormatInt(c.id, 10)
	} else if len(c.name) > 0 {
		return c.name
	}

	return ""
}
