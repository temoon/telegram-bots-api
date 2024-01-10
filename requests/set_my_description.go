package requests

import (
"context"
	"github.com/temoon/telegram-bots-api"
)

type SetMyDescription struct {
Description *string
LanguageCode *string
}

func (r *SetMyDescription) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new(bool)
	err = b.CallMethod(ctx, "setMyDescription", r, response)
	return
}



func (r *SetMyDescription) IsMultipart() (multipart bool) {
	return false
	}

func (r *SetMyDescription) GetValues() (values map[string]interface{}, err error) {
	values = make(map[string]interface{})

	
			if r.Description != nil {
			values["description"] = *r.Description
			}
			
			if r.LanguageCode != nil {
			values["language_code"] = *r.LanguageCode
			}
			

	return
}
