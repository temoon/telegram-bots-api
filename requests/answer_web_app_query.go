package requests

import (
	"context"
	"encoding/json"
	"github.com/temoon/telegram-bots-api"
)

type AnswerWebAppQuery struct {
	Result        interface{}
	WebAppQueryId string
}

func (r *AnswerWebAppQuery) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new(telegram.SentWebAppMessage)
	err = b.CallMethod(ctx, "answerWebAppQuery", r, response)
	return
}

func (r *AnswerWebAppQuery) IsMultipart() (multipart bool) {
	return false
}

func (r *AnswerWebAppQuery) GetValues() (values map[string]interface{}, err error) {
	values = make(map[string]interface{})

	switch value := r.Result.(type) {
	case telegram.InlineQueryResultCachedAudio:
		var dataInlineQueryResultCachedAudio []byte
		if dataInlineQueryResultCachedAudio, err = json.Marshal(value); err != nil {
			return
		}

		values["result"] = string(dataInlineQueryResultCachedAudio)
	case telegram.InlineQueryResultCachedDocument:
		var dataInlineQueryResultCachedDocument []byte
		if dataInlineQueryResultCachedDocument, err = json.Marshal(value); err != nil {
			return
		}

		values["result"] = string(dataInlineQueryResultCachedDocument)
	case telegram.InlineQueryResultCachedGif:
		var dataInlineQueryResultCachedGif []byte
		if dataInlineQueryResultCachedGif, err = json.Marshal(value); err != nil {
			return
		}

		values["result"] = string(dataInlineQueryResultCachedGif)
	case telegram.InlineQueryResultCachedMpeg4Gif:
		var dataInlineQueryResultCachedMpeg4Gif []byte
		if dataInlineQueryResultCachedMpeg4Gif, err = json.Marshal(value); err != nil {
			return
		}

		values["result"] = string(dataInlineQueryResultCachedMpeg4Gif)
	case telegram.InlineQueryResultCachedPhoto:
		var dataInlineQueryResultCachedPhoto []byte
		if dataInlineQueryResultCachedPhoto, err = json.Marshal(value); err != nil {
			return
		}

		values["result"] = string(dataInlineQueryResultCachedPhoto)
	case telegram.InlineQueryResultCachedSticker:
		var dataInlineQueryResultCachedSticker []byte
		if dataInlineQueryResultCachedSticker, err = json.Marshal(value); err != nil {
			return
		}

		values["result"] = string(dataInlineQueryResultCachedSticker)
	case telegram.InlineQueryResultCachedVideo:
		var dataInlineQueryResultCachedVideo []byte
		if dataInlineQueryResultCachedVideo, err = json.Marshal(value); err != nil {
			return
		}

		values["result"] = string(dataInlineQueryResultCachedVideo)
	case telegram.InlineQueryResultCachedVoice:
		var dataInlineQueryResultCachedVoice []byte
		if dataInlineQueryResultCachedVoice, err = json.Marshal(value); err != nil {
			return
		}

		values["result"] = string(dataInlineQueryResultCachedVoice)
	case telegram.InlineQueryResultArticle:
		var dataInlineQueryResultArticle []byte
		if dataInlineQueryResultArticle, err = json.Marshal(value); err != nil {
			return
		}

		values["result"] = string(dataInlineQueryResultArticle)
	case telegram.InlineQueryResultAudio:
		var dataInlineQueryResultAudio []byte
		if dataInlineQueryResultAudio, err = json.Marshal(value); err != nil {
			return
		}

		values["result"] = string(dataInlineQueryResultAudio)
	case telegram.InlineQueryResultContact:
		var dataInlineQueryResultContact []byte
		if dataInlineQueryResultContact, err = json.Marshal(value); err != nil {
			return
		}

		values["result"] = string(dataInlineQueryResultContact)
	case telegram.InlineQueryResultGame:
		var dataInlineQueryResultGame []byte
		if dataInlineQueryResultGame, err = json.Marshal(value); err != nil {
			return
		}

		values["result"] = string(dataInlineQueryResultGame)
	case telegram.InlineQueryResultDocument:
		var dataInlineQueryResultDocument []byte
		if dataInlineQueryResultDocument, err = json.Marshal(value); err != nil {
			return
		}

		values["result"] = string(dataInlineQueryResultDocument)
	case telegram.InlineQueryResultGif:
		var dataInlineQueryResultGif []byte
		if dataInlineQueryResultGif, err = json.Marshal(value); err != nil {
			return
		}

		values["result"] = string(dataInlineQueryResultGif)
	case telegram.InlineQueryResultLocation:
		var dataInlineQueryResultLocation []byte
		if dataInlineQueryResultLocation, err = json.Marshal(value); err != nil {
			return
		}

		values["result"] = string(dataInlineQueryResultLocation)
	case telegram.InlineQueryResultMpeg4Gif:
		var dataInlineQueryResultMpeg4Gif []byte
		if dataInlineQueryResultMpeg4Gif, err = json.Marshal(value); err != nil {
			return
		}

		values["result"] = string(dataInlineQueryResultMpeg4Gif)
	case telegram.InlineQueryResultPhoto:
		var dataInlineQueryResultPhoto []byte
		if dataInlineQueryResultPhoto, err = json.Marshal(value); err != nil {
			return
		}

		values["result"] = string(dataInlineQueryResultPhoto)
	case telegram.InlineQueryResultVenue:
		var dataInlineQueryResultVenue []byte
		if dataInlineQueryResultVenue, err = json.Marshal(value); err != nil {
			return
		}

		values["result"] = string(dataInlineQueryResultVenue)
	case telegram.InlineQueryResultVideo:
		var dataInlineQueryResultVideo []byte
		if dataInlineQueryResultVideo, err = json.Marshal(value); err != nil {
			return
		}

		values["result"] = string(dataInlineQueryResultVideo)
	case telegram.InlineQueryResultVoice:
		var dataInlineQueryResultVoice []byte
		if dataInlineQueryResultVoice, err = json.Marshal(value); err != nil {
			return
		}

		values["result"] = string(dataInlineQueryResultVoice)
	}

	values["web_app_query_id"] = r.WebAppQueryId

	return
}
