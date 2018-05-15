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
    "strings"

    "go-telegram-bots-api/requests"
)

type bot struct {
    token  string
    client http.Client
}

type Bot interface {
    GetUpdates(*requests.GetUpdates) ([]Update, error)
    SetWebhook(*requests.SetWebhook) (bool, error)
    DeleteWebhook() (bool, error)
    GetWebhookInfo() (WebhookInfo, error)
    GetMe() (User, error)
    SendMessage(*requests.SendMessage) (Message, error)
    ForwardMessage(*requests.ForwardMessage) (Message, error)
    SendPhoto(*requests.SendPhoto) (Message, error)
    SendAudio(*requests.SendAudio) (Message, error)
    SendDocument(*requests.SendDocument) (Message, error)
    SendVideo(*requests.SendVideo) (Message, error)
    SendVoice(*requests.SendVoice) (Message, error)
    SendVideoNote(*requests.SendVideoNote) (Message, error)
    SendMediaGroup(*requests.SendMediaGroup) (Message, error)
    SendLocation(*requests.SendLocation) (Message, error)
    EditMessageLiveLocation(*requests.EditMessageLiveLocation) (Message, error)
    StopMessageLiveLocation(*requests.StopMessageLiveLocation) (Message, error)
    SendVenue(*requests.SendVenue) (Message, error)
    SendContact(*requests.SendContact) (Message, error)
    SendChatAction(*requests.SendChatAction) (bool, error)
    GetUserProfilePhotos(*requests.GetUserProfilePhotos) (UserProfilePhotos, error)
    GetFile(*requests.GetFile) (File, error)
    KickChatMember(*requests.KickChatMember) (bool, error)
    UnbanChatMember(*requests.UnbanChatMember) (bool, error)
    RestrictChatMember(*requests.RestrictChatMember) (bool, error)
    PromoteChatMember(*requests.PromoteChatMember) (bool, error)
    ExportChatInviteLink(*requests.ExportChatInviteLink) (string, error)
    SetChatPhoto(*requests.SetChatPhoto) (bool, error)
    DeleteChatPhoto(*requests.DeleteChatPhoto) (bool, error)
    SetChatTitle(*requests.SetChatTitle) (bool, error)
    SetChatDescription(*requests.SetChatDescription) (bool, error)
    PinChatMessage(*requests.PinChatMessage) (bool, error)
    UnpinChatMessage(*requests.UnpinChatMessage) (bool, error)
    LeaveChat(*requests.LeaveChat) (bool, error)
    GetChat(*requests.GetChat) (Chat, error)
    GetChatAdministrators(*requests.GetChatAdministrators) ([]ChatMember, error)
    GetChatMembersCount(*requests.GetChatMembersCount) (int, error)
    GetChatMember(*requests.GetChatMember) (ChatMember, error)
    SetChatStickerSet(*requests.SetChatStickerSet) (bool, error)
    DeleteChatStickerSet(*requests.DeleteChatStickerSet) (bool, error)
    AnswerCallbackQuery(*requests.AnswerCallbackQuery) (bool, error)
    EditMessageText(*requests.EditMessageText) (Message, error)
    EditMessageCaption(*requests.EditMessageCaption) (Message, error)
    EditMessageReplyMarkup(*requests.EditMessageReplyMarkup) (Message, error)
    DeleteMessage(*requests.DeleteMessage) (bool, error)
    SendSticker(*requests.SendSticker) (Message, error)
    GetStickerSet(*requests.GetStickerSet) (StickerSet, error)
    UploadStickerFile(*requests.UploadStickerFile) (File, error)
    CreateNewStickerSet(*requests.CreateNewStickerSet) (bool, error)
    AddStickerToSet(*requests.AddStickerToSet) (bool, error)
    SetStickerPositionInSet(*requests.SetStickerPositionInSet) (bool, error)
    DeleteStickerFromSet(*requests.DeleteStickerFromSet) (bool, error)
    SendInvoice(*requests.SendInvoice) (Message, error)
    AnswerShippingQuery(*requests.AnswerShippingQuery) (bool, error)
    AnswerPreCheckoutQuery(*requests.AnswerPreCheckoutQuery) (bool, error)
    SendGame(*requests.SendGame) (Message, error)
    SetGameScore(*requests.SetGameScore) (Message, error)
    GetGameHighScores(*requests.GetGameHighScores) ([]GameHighScore, error)
}

func NewBot(token string) Bot {
    b := &bot{
        token:  token,
        client: http.Client{},
    }

    return b
}

func (b *bot) GetUpdates(request *requests.GetUpdates) (updates []Update, err error) {
    err = b.callMethod(b.getMethodURL("getUpdates"), request, &updates)

    return
}

func (b *bot) SetWebhook(request *requests.SetWebhook) (success bool, err error) {
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

func (b *bot) SendMessage(request *requests.SendMessage) (message Message, err error) {
    err = b.callMethod(b.getMethodURL("sendMessage"), request, &message)

    return
}

func (b *bot) ForwardMessage(request *requests.ForwardMessage) (message Message, err error) {
    err = b.callMethod(b.getMethodURL("forwardMessage"), request, &message)

    return
}

func (b *bot) SendPhoto(request *requests.SendPhoto) (message Message, err error) {
    err = b.callMethod(b.getMethodURL("sendPhoto"), request, &message)

    return
}

func (b *bot) SendAudio(request *requests.SendAudio) (message Message, err error) {
    err = b.callMethod(b.getMethodURL("sendAudio"), request, &message)

    return
}

func (b *bot) SendDocument(request *requests.SendDocument) (message Message, err error) {
    err = b.callMethod(b.getMethodURL("sendDocument"), request, &message)

    return
}

func (b *bot) SendVideo(request *requests.SendVideo) (message Message, err error) {
    err = b.callMethod(b.getMethodURL("sendVideo"), request, &message)

    return
}

func (b *bot) SendVoice(request *requests.SendVoice) (message Message, err error) {
    err = b.callMethod(b.getMethodURL("sendVoice"), request, &message)

    return
}

func (b *bot) SendVideoNote(request *requests.SendVideoNote) (message Message, err error) {
    err = b.callMethod(b.getMethodURL("sendVideoNote"), request, &message)

    return
}

func (b *bot) SendMediaGroup(request *requests.SendMediaGroup) (message Message, err error) {
    err = b.callMethod(b.getMethodURL("sendMediaGroup"), request, &message)

    return
}

func (b *bot) SendLocation(request *requests.SendLocation) (message Message, err error) {
    err = b.callMethod(b.getMethodURL("sendLocation"), request, &message)

    return
}

func (b *bot) EditMessageLiveLocation(request *requests.EditMessageLiveLocation) (message Message, err error) {
    err = b.callMethod(b.getMethodURL("editMessageLiveLocation"), request, &message)

    return
}

func (b *bot) StopMessageLiveLocation(request *requests.StopMessageLiveLocation) (message Message, err error) {
    err = b.callMethod(b.getMethodURL("stopMessageLiveLocation"), request, &message)

    return
}

func (b *bot) SendVenue(request *requests.SendVenue) (message Message, err error) {
    err = b.callMethod(b.getMethodURL("sendVenue"), request, &message)

    return
}

func (b *bot) SendContact(request *requests.SendContact) (message Message, err error) {
    err = b.callMethod(b.getMethodURL("sendContact"), request, &message)

    return
}

func (b *bot) SendChatAction(request *requests.SendChatAction) (success bool, err error) {
    err = b.callMethod(b.getMethodURL("sendChatAction"), request, &success)

    return
}

func (b *bot) GetUserProfilePhotos(request *requests.GetUserProfilePhotos) (photos UserProfilePhotos, err error) {
    err = b.callMethod(b.getMethodURL("getUserProfilePhotos"), request, &photos)

    return
}

func (b *bot) GetFile(request *requests.GetFile) (file File, err error) {
    err = b.callMethod(b.getMethodURL("getFile"), request, &file)

    return
}

func (b *bot) KickChatMember(request *requests.KickChatMember) (success bool, err error) {
    err = b.callMethod(b.getMethodURL("kickChatMember"), request, &success)

    return
}

func (b *bot) UnbanChatMember(request *requests.UnbanChatMember) (success bool, err error) {
    err = b.callMethod(b.getMethodURL("unbanChatMember"), request, &success)

    return
}

func (b *bot) RestrictChatMember(request *requests.RestrictChatMember) (success bool, err error) {
    err = b.callMethod(b.getMethodURL("restrictChatMember"), request, &success)

    return
}

func (b *bot) PromoteChatMember(request *requests.PromoteChatMember) (success bool, err error) {
    err = b.callMethod(b.getMethodURL("promoteChatMember"), request, &success)

    return
}

func (b *bot) ExportChatInviteLink(request *requests.ExportChatInviteLink) (link string, err error) {
    err = b.callMethod(b.getMethodURL("exportChatInviteLink"), request, &link)

    return
}

func (b *bot) SetChatPhoto(request *requests.SetChatPhoto) (success bool, err error) {
    err = b.callMethod(b.getMethodURL("setChatPhoto"), request, &success)

    return
}

func (b *bot) DeleteChatPhoto(request *requests.DeleteChatPhoto) (success bool, err error) {
    err = b.callMethod(b.getMethodURL("deleteChatPhoto"), request, &success)

    return
}

func (b *bot) SetChatTitle(request *requests.SetChatTitle) (success bool, err error) {
    err = b.callMethod(b.getMethodURL("setChatTitle"), request, &success)

    return
}

func (b *bot) SetChatDescription(request *requests.SetChatDescription) (success bool, err error) {
    err = b.callMethod(b.getMethodURL("setChatDescription"), request, &success)

    return
}

func (b *bot) PinChatMessage(request *requests.PinChatMessage) (success bool, err error) {
    err = b.callMethod(b.getMethodURL("pinChatMessage"), request, &success)

    return
}

func (b *bot) UnpinChatMessage(request *requests.UnpinChatMessage) (success bool, err error) {
    err = b.callMethod(b.getMethodURL("unpinChatMessage"), request, &success)

    return
}

func (b *bot) LeaveChat(request *requests.LeaveChat) (success bool, err error) {
    err = b.callMethod(b.getMethodURL("leaveChat"), request, &success)

    return
}

func (b *bot) GetChat(request *requests.GetChat) (chat Chat, err error) {
    err = b.callMethod(b.getMethodURL("getChat"), request, &chat)

    return
}

func (b *bot) GetChatAdministrators(request *requests.GetChatAdministrators) (members []ChatMember, err error) {
    err = b.callMethod(b.getMethodURL("getChatAdministrators"), request, &members)

    return
}

func (b *bot) GetChatMembersCount(request *requests.GetChatMembersCount) (count int, err error) {
    err = b.callMethod(b.getMethodURL("getChatMembersCount"), request, &count)

    return
}

func (b *bot) GetChatMember(request *requests.GetChatMember) (member ChatMember, err error) {
    err = b.callMethod(b.getMethodURL("setChatDescription"), request, &member)

    return
}

func (b *bot) SetChatStickerSet(request *requests.SetChatStickerSet) (success bool, err error) {
    err = b.callMethod(b.getMethodURL("setChatStickerSet"), request, &success)

    return
}

func (b *bot) DeleteChatStickerSet(request *requests.DeleteChatStickerSet) (success bool, err error) {
    err = b.callMethod(b.getMethodURL("deleteChatStickerSet"), request, &success)

    return
}

func (b *bot) AnswerCallbackQuery(request *requests.AnswerCallbackQuery) (success bool, err error) {
    err = b.callMethod(b.getMethodURL("answerCallbackQuery"), request, &success)

    return
}

func (b *bot) EditMessageText(request *requests.EditMessageText) (message Message, err error) {
    err = b.callMethod(b.getMethodURL("editMessageText"), request, &message)

    return
}

func (b *bot) EditMessageCaption(request *requests.EditMessageCaption) (message Message, err error) {
    err = b.callMethod(b.getMethodURL("editMessageCaption"), request, &message)

    return
}

func (b *bot) EditMessageReplyMarkup(request *requests.EditMessageReplyMarkup) (message Message, err error) {
    err = b.callMethod(b.getMethodURL("editMessageReplyMarkup"), request, &message)

    return
}

func (b *bot) DeleteMessage(request *requests.DeleteMessage) (success bool, err error) {
    err = b.callMethod(b.getMethodURL("deleteMessage"), request, &success)

    return
}

func (b *bot) SendSticker(request *requests.SendSticker) (message Message, err error) {
    err = b.callMethod(b.getMethodURL("sendSticker"), request, &message)

    return
}

func (b *bot) GetStickerSet(request *requests.GetStickerSet) (set StickerSet, err error) {
    err = b.callMethod(b.getMethodURL("getStickerSet"), request, &set)

    return
}

func (b *bot) UploadStickerFile(request *requests.UploadStickerFile) (file File, err error) {
    err = b.callMethod(b.getMethodURL("uploadStickerFile"), request, &file)

    return
}

func (b *bot) CreateNewStickerSet(request *requests.CreateNewStickerSet) (success bool, err error) {
    err = b.callMethod(b.getMethodURL("createNewStickerSet"), request, &success)

    return
}

func (b *bot) AddStickerToSet(request *requests.AddStickerToSet) (success bool, err error) {
    err = b.callMethod(b.getMethodURL("addStickerToSet"), request, &success)

    return
}

func (b *bot) SetStickerPositionInSet(request *requests.SetStickerPositionInSet) (success bool, err error) {
    err = b.callMethod(b.getMethodURL("setStickerPositionInSet"), request, &success)

    return
}

func (b *bot) DeleteStickerFromSet(request *requests.DeleteStickerFromSet) (success bool, err error) {
    err = b.callMethod(b.getMethodURL("deleteStickerFromSet"), request, &success)

    return
}

func (b *bot) SendInvoice(request *requests.SendInvoice) (message Message, err error) {
    err = b.callMethod(b.getMethodURL("sendInvoice"), request, &message)

    return
}

func (b *bot) AnswerShippingQuery(request *requests.AnswerShippingQuery) (success bool, err error) {
    err = b.callMethod(b.getMethodURL("answerShippingQuery"), request, &success)

    return
}

func (b *bot) AnswerPreCheckoutQuery(request *requests.AnswerPreCheckoutQuery) (success bool, err error) {
    err = b.callMethod(b.getMethodURL("answerPreCheckoutQuery"), request, &success)

    return
}

func (b *bot) SendGame(request *requests.SendGame) (message Message, err error) {
    err = b.callMethod(b.getMethodURL("sendGame"), request, &message)

    return
}

func (b *bot) SetGameScore(request *requests.SetGameScore) (message Message, err error) {
    err = b.callMethod(b.getMethodURL("setGameScore"), request, &message)

    return
}

func (b *bot) GetGameHighScores(request *requests.GetGameHighScores) (highScores []GameHighScore, err error) {
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

    var rv map[string][]interface{}
    if rv, err = request.GetValues(); err != nil {
        return
    }

    var prefix string

    for key, values := range rv {
        prefix = url.QueryEscape(key) + "="

        for _, value := range values {
            if query.Len() > 0 {
                query.WriteByte('&')
            }

            query.WriteString(prefix)
            query.WriteString(url.QueryEscape(value.(string)))
        }
    }

    return
}

func (b *bot) getFormMultipart(request Request) (contentType string, query *bytes.Buffer, err error) {
    var rv map[string][]interface{}
    if rv, err = request.GetValues(); err != nil {
        return
    }

    query = &bytes.Buffer{}

    var mw = multipart.NewWriter(query)
    var fw io.Writer

    for key, values := range rv {
        for _, value := range values {
            if file, ok := value.(*os.File); ok {
                if fw, err = mw.CreateFormFile(key, file.Name()); err != nil {
                    return
                }
            } else {
                if fw, err = mw.CreateFormField(key); err != nil {
                    return
                }
            }

            reader, ok := value.(io.Reader)
            if !ok {
                reader = strings.NewReader(value.(string))
            }

            if _, err = io.Copy(fw, reader); err != nil {
                return
            }
        }
    }

    mw.Close()

    contentType = mw.FormDataContentType()

    return
}
