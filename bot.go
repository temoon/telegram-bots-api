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
    SendPhoto(request *requests.SendPhoto) (Message, error)
    SendAudio(request *requests.SendAudio) (Message, error)
    SendDocument(request *requests.SendDocument) (Message, error)
    SendVideo(request *requests.SendVideo) (Message, error)
    SendVoice(request *requests.SendVoice) (Message, error)
    SendVideoNote(request *requests.SendVideoNote) (Message, error)
    SendMediaGroup(request *requests.SendMediaGroup) (Message, error)
    SendLocation(request *requests.SendLocation) (Message, error)
    EditMessageLiveLocation(request *requests.EditMessageLiveLocation) (Message, error)
    StopMessageLiveLocation(request *requests.StopMessageLiveLocation) (Message, error)
    SendVenue(request *requests.SendVenue) (Message, error)
    SendContact(request *requests.SendContact) (Message, error)
    SendChatAction(request *requests.SendChatAction) (bool, error)
    GetUserProfilePhotos(request *requests.GetUserProfilePhotos) (UserProfilePhotos, error)
    GetFile(request *requests.GetFile) (File, error)
    KickChatMember(request *requests.KickChatMember) (bool, error)
    UnbanChatMember(request *requests.UnbanChatMember) (bool, error)
    RestrictChatMember(request *requests.RestrictChatMember) (bool, error)
    PromoteChatMember(request *requests.PromoteChatMember) (bool, error)
    ExportChatInviteLink(request *requests.ExportChatInviteLink) (string, error)
    SetChatPhoto(request *requests.SetChatPhoto) (bool, error)
    DeleteChatPhoto(request *requests.DeleteChatPhoto) (bool, error)
    SetChatTitle(request *requests.SetChatTitle) (bool, error)
    SetChatDescription(request *requests.SetChatDescription) (bool, error)
    PinChatMessage(request *requests.PinChatMessage) (bool, error)
    UnpinChatMessage(request *requests.UnpinChatMessage) (bool, error)
    LeaveChat(request *requests.LeaveChat) (bool, error)
    GetChat(request *requests.GetChat) (Chat, error)
    GetChatAdministrators(request *requests.GetChatAdministrators) ([]ChatMember, error)
    GetChatMembersCount(request *requests.GetChatMembersCount) (int, error)
    GetChatMember(request *requests.GetChatMember) (ChatMember, error)
    SetChatStickerSet(request *requests.SetChatStickerSet) (bool, error)
    DeleteChatStickerSet(request *requests.DeleteChatStickerSet) (bool, error)
    AnswerCallbackQuery(request *requests.AnswerCallbackQuery) (bool, error)
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

    if !telegramResponse.Ok {
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
