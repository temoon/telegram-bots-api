package requests

import (
"encoding/json"
"strconv"
"context"
	"github.com/temoon/telegram-bots-api"
)

type SetPassportDataErrors struct {
Errors []interface{}
UserId int32
}

func (r *SetPassportDataErrors) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new(bool)
	err = b.CallMethod(ctx, "setPassportDataErrors", r, response)
	return
}



func (r *SetPassportDataErrors) IsMultipart() (multipart bool) {
	return false
	}

func (r *SetPassportDataErrors) GetValues() (values map[string]interface{}, err error) {
	values = make(map[string]interface{})

	
			var dataErrors []byte
				if dataErrors, err = json.Marshal(r.Errors); err != nil {
					return
				}

				values["errors"] = string(dataErrors)
			
			values["user_id"] = strconv.FormatInt(int64(r.UserId), 10)
			

	return
}
