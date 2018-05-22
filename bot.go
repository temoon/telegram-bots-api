package telegram

import (
    "bytes"
    "encoding/json"
    "errors"
    "io"
    "io/ioutil"
    "mime/multipart"
    "net/http"
    "net/url"
    "os"

    . "github.com/temoon/go-telegram-bots-api/requests"
)

type bot struct {
    token  string
    client http.Client
}

type Bot interface {
    GetUpdates(*GetUpdates) ([]Update, error)
    SetWebhook(*SetWebhook) (bool, error)
    DeleteWebhook() (bool, error)
    GetWebhookInfo() (WebhookInfo, error)
    GetMe() (User, error)
    SendMessage(*SendMessage) (Message, error)
    ForwardMessage(*ForwardMessage) (Message, error)
    SendPhoto(*SendPhoto) (Message, error)
    SendAudio(*SendAudio) (Message, error)
    SendDocument(*SendDocument) (Message, error)
    SendVideo(*SendVideo) (Message, error)
    SendVoice(*SendVoice) (Message, error)
    SendVideoNote(*SendVideoNote) (Message, error)
    SendMediaGroup(*SendMediaGroup) (Message, error)
    SendLocation(*SendLocation) (Message, error)
    EditMessageLiveLocation(*EditMessageLiveLocation) (Message, error)
    StopMessageLiveLocation(*StopMessageLiveLocation) (Message, error)
    SendVenue(*SendVenue) (Message, error)
    SendContact(*SendContact) (Message, error)
    SendChatAction(*SendChatAction) (bool, error)
    GetUserProfilePhotos(*GetUserProfilePhotos) (UserProfilePhotos, error)
    GetFile(*GetFile) (File, error)
    KickChatMember(*KickChatMember) (bool, error)
    UnbanChatMember(*UnbanChatMember) (bool, error)
    RestrictChatMember(*RestrictChatMember) (bool, error)
    PromoteChatMember(*PromoteChatMember) (bool, error)
    ExportChatInviteLink(*ExportChatInviteLink) (string, error)
    SetChatPhoto(*SetChatPhoto) (bool, error)
    DeleteChatPhoto(*DeleteChatPhoto) (bool, error)
    SetChatTitle(*SetChatTitle) (bool, error)
    SetChatDescription(*SetChatDescription) (bool, error)
    PinChatMessage(*PinChatMessage) (bool, error)
    UnpinChatMessage(*UnpinChatMessage) (bool, error)
    LeaveChat(*LeaveChat) (bool, error)
    GetChat(*GetChat) (Chat, error)
    GetChatAdministrators(*GetChatAdministrators) ([]ChatMember, error)
    GetChatMembersCount(*GetChatMembersCount) (int, error)
    GetChatMember(*GetChatMember) (ChatMember, error)
    SetChatStickerSet(*SetChatStickerSet) (bool, error)
    DeleteChatStickerSet(*DeleteChatStickerSet) (bool, error)
    AnswerCallbackQuery(*AnswerCallbackQuery) (bool, error)
    EditMessageText(*EditMessageText) (Message, error)
    EditMessageCaption(*EditMessageCaption) (Message, error)
    EditMessageReplyMarkup(*EditMessageReplyMarkup) (Message, error)
    DeleteMessage(*DeleteMessage) (bool, error)
    SendSticker(*SendSticker) (Message, error)
    GetStickerSet(*GetStickerSet) (StickerSet, error)
    UploadStickerFile(*UploadStickerFile) (File, error)
    CreateNewStickerSet(*CreateNewStickerSet) (bool, error)
    AddStickerToSet(*AddStickerToSet) (bool, error)
    SetStickerPositionInSet(*SetStickerPositionInSet) (bool, error)
    DeleteStickerFromSet(*DeleteStickerFromSet) (bool, error)
    AnswerInlineQuery(*AnswerInlineQuery) (bool, error)
    SendInvoice(*SendInvoice) (Message, error)
    AnswerShippingQuery(*AnswerShippingQuery) (bool, error)
    AnswerPreCheckoutQuery(*AnswerPreCheckoutQuery) (bool, error)
    SendGame(*SendGame) (Message, error)
    SetGameScore(*SetGameScore) (Message, error)
    GetGameHighScores(*GetGameHighScores) ([]GameHighScore, error)
}

func NewBot(token string) Bot {
    b := &bot{
        token:  token,
        client: http.Client{},
    }

    return b
}

func (b *bot) GetUpdates(request *GetUpdates) (updates []Update, err error) {
    err = b.callMethod(b.getMethodURL("getUpdates"), request, &updates)

    return
}

func (b *bot) SetWebhook(request *SetWebhook) (success bool, err error) {
    err = b.callMethod(b.getMethodURL("setWebhook"), request, &success)

    return
}

func (b *bot) DeleteWebhook() (success bool, err error) {
    err = b.callMethod(b.getMethodURL("deleteWebhook"), nil, &success)

    return
}

func (b *bot) GetWebhookInfo() (info WebhookInfo, err error) {
    err = b.callMethod(b.getMethodURL("getWebhookInfo"), nil, &info)

    return
}

func (b *bot) GetMe() (me User, err error) {
    err = b.callMethod(b.getMethodURL("getMe"), nil, &me)

    return
}

func (b *bot) SendMessage(request *SendMessage) (message Message, err error) {
    err = b.callMethod(b.getMethodURL("sendMessage"), request, &message)

    return
}

func (b *bot) ForwardMessage(request *ForwardMessage) (message Message, err error) {
    err = b.callMethod(b.getMethodURL("forwardMessage"), request, &message)

    return
}

func (b *bot) SendPhoto(request *SendPhoto) (message Message, err error) {
    err = b.callMethod(b.getMethodURL("sendPhoto"), request, &message)

    return
}

func (b *bot) SendAudio(request *SendAudio) (message Message, err error) {
    err = b.callMethod(b.getMethodURL("sendAudio"), request, &message)

    return
}

func (b *bot) SendDocument(request *SendDocument) (message Message, err error) {
    err = b.callMethod(b.getMethodURL("sendDocument"), request, &message)

    return
}

func (b *bot) SendVideo(request *SendVideo) (message Message, err error) {
    err = b.callMethod(b.getMethodURL("sendVideo"), request, &message)

    return
}

func (b *bot) SendVoice(request *SendVoice) (message Message, err error) {
    err = b.callMethod(b.getMethodURL("sendVoice"), request, &message)

    return
}

func (b *bot) SendVideoNote(request *SendVideoNote) (message Message, err error) {
    err = b.callMethod(b.getMethodURL("sendVideoNote"), request, &message)

    return
}

func (b *bot) SendMediaGroup(request *SendMediaGroup) (message Message, err error) {
    err = b.callMethod(b.getMethodURL("sendMediaGroup"), request, &message)

    return
}

func (b *bot) SendLocation(request *SendLocation) (message Message, err error) {
    err = b.callMethod(b.getMethodURL("sendLocation"), request, &message)

    return
}

func (b *bot) EditMessageLiveLocation(request *EditMessageLiveLocation) (message Message, err error) {
    err = b.callMethod(b.getMethodURL("editMessageLiveLocation"), request, &message)

    return
}

func (b *bot) StopMessageLiveLocation(request *StopMessageLiveLocation) (message Message, err error) {
    err = b.callMethod(b.getMethodURL("stopMessageLiveLocation"), request, &message)

    return
}

func (b *bot) SendVenue(request *SendVenue) (message Message, err error) {
    err = b.callMethod(b.getMethodURL("sendVenue"), request, &message)

    return
}

func (b *bot) SendContact(request *SendContact) (message Message, err error) {
    err = b.callMethod(b.getMethodURL("sendContact"), request, &message)

    return
}

func (b *bot) SendChatAction(request *SendChatAction) (success bool, err error) {
    err = b.callMethod(b.getMethodURL("sendChatAction"), request, &success)

    return
}

func (b *bot) GetUserProfilePhotos(request *GetUserProfilePhotos) (photos UserProfilePhotos, err error) {
    err = b.callMethod(b.getMethodURL("getUserProfilePhotos"), request, &photos)

    return
}

func (b *bot) GetFile(request *GetFile) (file File, err error) {
    err = b.callMethod(b.getMethodURL("getFile"), request, &file)

    return
}

func (b *bot) KickChatMember(request *KickChatMember) (success bool, err error) {
    err = b.callMethod(b.getMethodURL("kickChatMember"), request, &success)

    return
}

func (b *bot) UnbanChatMember(request *UnbanChatMember) (success bool, err error) {
    err = b.callMethod(b.getMethodURL("unbanChatMember"), request, &success)

    return
}

func (b *bot) RestrictChatMember(request *RestrictChatMember) (success bool, err error) {
    err = b.callMethod(b.getMethodURL("restrictChatMember"), request, &success)

    return
}

func (b *bot) PromoteChatMember(request *PromoteChatMember) (success bool, err error) {
    err = b.callMethod(b.getMethodURL("promoteChatMember"), request, &success)

    return
}

func (b *bot) ExportChatInviteLink(request *ExportChatInviteLink) (link string, err error) {
    err = b.callMethod(b.getMethodURL("exportChatInviteLink"), request, &link)

    return
}

func (b *bot) SetChatPhoto(request *SetChatPhoto) (success bool, err error) {
    err = b.callMethod(b.getMethodURL("setChatPhoto"), request, &success)

    return
}

func (b *bot) DeleteChatPhoto(request *DeleteChatPhoto) (success bool, err error) {
    err = b.callMethod(b.getMethodURL("deleteChatPhoto"), request, &success)

    return
}

func (b *bot) SetChatTitle(request *SetChatTitle) (success bool, err error) {
    err = b.callMethod(b.getMethodURL("setChatTitle"), request, &success)

    return
}

func (b *bot) SetChatDescription(request *SetChatDescription) (success bool, err error) {
    err = b.callMethod(b.getMethodURL("setChatDescription"), request, &success)

    return
}

func (b *bot) PinChatMessage(request *PinChatMessage) (success bool, err error) {
    err = b.callMethod(b.getMethodURL("pinChatMessage"), request, &success)

    return
}

func (b *bot) UnpinChatMessage(request *UnpinChatMessage) (success bool, err error) {
    err = b.callMethod(b.getMethodURL("unpinChatMessage"), request, &success)

    return
}

func (b *bot) LeaveChat(request *LeaveChat) (success bool, err error) {
    err = b.callMethod(b.getMethodURL("leaveChat"), request, &success)

    return
}

func (b *bot) GetChat(request *GetChat) (chat Chat, err error) {
    err = b.callMethod(b.getMethodURL("getChat"), request, &chat)

    return
}

func (b *bot) GetChatAdministrators(request *GetChatAdministrators) (members []ChatMember, err error) {
    err = b.callMethod(b.getMethodURL("getChatAdministrators"), request, &members)

    return
}

func (b *bot) GetChatMembersCount(request *GetChatMembersCount) (count int, err error) {
    err = b.callMethod(b.getMethodURL("getChatMembersCount"), request, &count)

    return
}

func (b *bot) GetChatMember(request *GetChatMember) (member ChatMember, err error) {
    err = b.callMethod(b.getMethodURL("setChatDescription"), request, &member)

    return
}

func (b *bot) SetChatStickerSet(request *SetChatStickerSet) (success bool, err error) {
    err = b.callMethod(b.getMethodURL("setChatStickerSet"), request, &success)

    return
}

func (b *bot) DeleteChatStickerSet(request *DeleteChatStickerSet) (success bool, err error) {
    err = b.callMethod(b.getMethodURL("deleteChatStickerSet"), request, &success)

    return
}

func (b *bot) AnswerCallbackQuery(request *AnswerCallbackQuery) (success bool, err error) {
    err = b.callMethod(b.getMethodURL("answerCallbackQuery"), request, &success)

    return
}

func (b *bot) EditMessageText(request *EditMessageText) (message Message, err error) {
    err = b.callMethod(b.getMethodURL("editMessageText"), request, &message)

    return
}

func (b *bot) EditMessageCaption(request *EditMessageCaption) (message Message, err error) {
    err = b.callMethod(b.getMethodURL("editMessageCaption"), request, &message)

    return
}

func (b *bot) EditMessageReplyMarkup(request *EditMessageReplyMarkup) (message Message, err error) {
    err = b.callMethod(b.getMethodURL("editMessageReplyMarkup"), request, &message)

    return
}

func (b *bot) DeleteMessage(request *DeleteMessage) (success bool, err error) {
    err = b.callMethod(b.getMethodURL("deleteMessage"), request, &success)

    return
}

func (b *bot) SendSticker(request *SendSticker) (message Message, err error) {
    err = b.callMethod(b.getMethodURL("sendSticker"), request, &message)

    return
}

func (b *bot) GetStickerSet(request *GetStickerSet) (set StickerSet, err error) {
    err = b.callMethod(b.getMethodURL("getStickerSet"), request, &set)

    return
}

func (b *bot) UploadStickerFile(request *UploadStickerFile) (file File, err error) {
    err = b.callMethod(b.getMethodURL("uploadStickerFile"), request, &file)

    return
}

func (b *bot) CreateNewStickerSet(request *CreateNewStickerSet) (success bool, err error) {
    err = b.callMethod(b.getMethodURL("createNewStickerSet"), request, &success)

    return
}

func (b *bot) AddStickerToSet(request *AddStickerToSet) (success bool, err error) {
    err = b.callMethod(b.getMethodURL("addStickerToSet"), request, &success)

    return
}

func (b *bot) SetStickerPositionInSet(request *SetStickerPositionInSet) (success bool, err error) {
    err = b.callMethod(b.getMethodURL("setStickerPositionInSet"), request, &success)

    return
}

func (b *bot) DeleteStickerFromSet(request *DeleteStickerFromSet) (success bool, err error) {
    err = b.callMethod(b.getMethodURL("deleteStickerFromSet"), request, &success)

    return
}

func (b *bot) AnswerInlineQuery(request *AnswerInlineQuery) (success bool, err error) {
    err = b.callMethod(b.getMethodURL("answerInlineQuery"), request, &success)

    return
}

func (b *bot) SendInvoice(request *SendInvoice) (message Message, err error) {
    err = b.callMethod(b.getMethodURL("sendInvoice"), request, &message)

    return
}

func (b *bot) AnswerShippingQuery(request *AnswerShippingQuery) (success bool, err error) {
    err = b.callMethod(b.getMethodURL("answerShippingQuery"), request, &success)

    return
}

func (b *bot) AnswerPreCheckoutQuery(request *AnswerPreCheckoutQuery) (success bool, err error) {
    err = b.callMethod(b.getMethodURL("answerPreCheckoutQuery"), request, &success)

    return
}

func (b *bot) SendGame(request *SendGame) (message Message, err error) {
    err = b.callMethod(b.getMethodURL("sendGame"), request, &message)

    return
}

func (b *bot) SetGameScore(request *SetGameScore) (message Message, err error) {
    err = b.callMethod(b.getMethodURL("setGameScore"), request, &message)

    return
}

func (b *bot) GetGameHighScores(request *GetGameHighScores) (highScores []GameHighScore, err error) {
    err = b.callMethod(b.getMethodURL("getGameHighScores"), request, &highScores)

    return
}

func (b *bot) getMethodURL(method string) string {
    return "https://api.telegram.org/bot" + b.token + "/" + method
}

func (b *bot) callMethod(methodUrl string, request Request, response interface{}) (err error) {
    var contentType string
    var query *bytes.Buffer

    if request != nil && request.IsMultipart() {
        contentType, query, err = b.getFormMultipart(request)
    } else {
        contentType, query, err = b.getForm(request)
    }

    if err != nil {
        return
    }

    var httpResponse *http.Response
    if httpResponse, err = b.client.Post(methodUrl, contentType, query); err != nil {
        return
    }
    defer httpResponse.Body.Close()

    var data []byte
    if data, err = ioutil.ReadAll(httpResponse.Body); err != nil {
        return
    }

    var telegramResponse Response
    if err = json.Unmarshal(data, &telegramResponse); err != nil {
        return
    }

    if !telegramResponse.OK {
        return errors.New(telegramResponse.Description)
    }

    return json.Unmarshal(telegramResponse.Result, response)
}

func (b *bot) getForm(request Request) (contentType string, query *bytes.Buffer, err error) {
    contentType = "application/x-www-form-urlencoded"
    query = &bytes.Buffer{}

    if request == nil {
        return
    }

    var values map[string]interface{}
    if values, err = request.GetValues(); err != nil {
        return
    }

    for key, value := range values {
        if query.Len() > 0 {
            query.WriteByte('&')
        }

        query.WriteString(url.QueryEscape(key) + "=")
        query.WriteString(url.QueryEscape(value.(string)))
    }

    return
}

func (b *bot) getFormMultipart(request Request) (contentType string, query *bytes.Buffer, err error) {
    var values map[string]interface{}
    if values, err = request.GetValues(); err != nil {
        return
    }

    query = &bytes.Buffer{}

    var mw = multipart.NewWriter(query)
    var fw io.Writer

    for key, value := range values {
        if file, ok := value.(*os.File); ok {
            if fw, err = mw.CreateFormFile(key, file.Name()); err != nil {
                return
            }

            if _, err = io.Copy(fw, value.(io.Reader)); err != nil {
                return
            }
        } else {
            if fw, err = mw.CreateFormField(key); err != nil {
                return
            }

            fw.Write([]byte(value.(string)))
        }
    }

    mw.Close()

    contentType = mw.FormDataContentType()

    return
}
