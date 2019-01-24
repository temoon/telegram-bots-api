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

    . "github.com/temoon/go-telegram-bots-api/requests"
)

type Bot struct {
    opts *BotOpts
}

type BotOpts struct {
    Token  string
    Client *http.Client
}

type BotAPI interface {
    GetUpdates(*GetUpdates) ([]Update, error)
    SetWebhook(*SetWebhook) (bool, error)
    DeleteWebhook() (bool, error)
    GetWebhookInfo() (*WebhookInfo, error)
    GetMe() (*User, error)
    SendMessage(*SendMessage) (*Message, error)
    ForwardMessage(*ForwardMessage) (*Message, error)
    SendPhoto(*SendPhoto) (*Message, error)
    SendAudio(*SendAudio) (*Message, error)
    SendDocument(*SendDocument) (*Message, error)
    SendVideo(*SendVideo) (*Message, error)
    SendVoice(*SendVoice) (*Message, error)
    SendVideoNote(*SendVideoNote) (*Message, error)
    SendMediaGroup(*SendMediaGroup) (*Message, error)
    SendLocation(*SendLocation) (*Message, error)
    EditMessageLiveLocation(*EditMessageLiveLocation) (*Message, error)
    StopMessageLiveLocation(*StopMessageLiveLocation) (*Message, error)
    SendVenue(*SendVenue) (*Message, error)
    SendContact(*SendContact) (*Message, error)
    SendChatAction(*SendChatAction) (bool, error)
    GetUserProfilePhotos(*GetUserProfilePhotos) (*UserProfilePhotos, error)
    GetFile(*GetFile) (*File, error)
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
    GetChat(*GetChat) (*Chat, error)
    GetChatAdministrators(*GetChatAdministrators) ([]ChatMember, error)
    GetChatMembersCount(*GetChatMembersCount) (int, error)
    GetChatMember(*GetChatMember) (*ChatMember, error)
    SetChatStickerSet(*SetChatStickerSet) (bool, error)
    DeleteChatStickerSet(*DeleteChatStickerSet) (bool, error)
    AnswerCallbackQuery(*AnswerCallbackQuery) (bool, error)
    EditMessageText(*EditMessageText) (*Message, error)
    EditMessageCaption(*EditMessageCaption) (*Message, error)
    EditMessageReplyMarkup(*EditMessageReplyMarkup) (*Message, error)
    DeleteMessage(*DeleteMessage) (bool, error)
    SendSticker(*SendSticker) (*Message, error)
    GetStickerSet(*GetStickerSet) (*StickerSet, error)
    UploadStickerFile(*UploadStickerFile) (*File, error)
    CreateNewStickerSet(*CreateNewStickerSet) (bool, error)
    AddStickerToSet(*AddStickerToSet) (bool, error)
    SetStickerPositionInSet(*SetStickerPositionInSet) (bool, error)
    DeleteStickerFromSet(*DeleteStickerFromSet) (bool, error)
    AnswerInlineQuery(*AnswerInlineQuery) (bool, error)
    SendInvoice(*SendInvoice) (*Message, error)
    AnswerShippingQuery(*AnswerShippingQuery) (bool, error)
    AnswerPreCheckoutQuery(*AnswerPreCheckoutQuery) (bool, error)
    SetPassportDataErrors(*SetPassportDataErrors) (bool, error)
    SendGame(*SendGame) (*Message, error)
    SetGameScore(*SetGameScore) (*Message, error)
    GetGameHighScores(*GetGameHighScores) ([]GameHighScore, error)
}

func NewBot(opts *BotOpts) (bot BotAPI) {
    if opts.Client == nil {
        opts.Client = new(http.Client)
    }

    bot = new(Bot)
    bot.(*Bot).opts = opts

    return
}

func (b *Bot) GetUpdates(request *GetUpdates) (updates []Update, err error) {
    err = b.callMethod(b.getMethodURL("getUpdates"), request, &updates)

    return
}

func (b *Bot) SetWebhook(request *SetWebhook) (success bool, err error) {
    err = b.callMethod(b.getMethodURL("setWebhook"), request, &success)

    return
}

func (b *Bot) DeleteWebhook() (success bool, err error) {
    err = b.callMethod(b.getMethodURL("deleteWebhook"), nil, &success)

    return
}

func (b *Bot) GetWebhookInfo() (info *WebhookInfo, err error) {
    info = new(WebhookInfo)
    err = b.callMethod(b.getMethodURL("getWebhookInfo"), nil, info)

    return
}

func (b *Bot) GetMe() (me *User, err error) {
    me = new(User)
    err = b.callMethod(b.getMethodURL("getMe"), nil, me)

    return
}

func (b *Bot) SendMessage(request *SendMessage) (message *Message, err error) {
    message = new(Message)
    err = b.callMethod(b.getMethodURL("sendMessage"), request, message)

    return
}

func (b *Bot) ForwardMessage(request *ForwardMessage) (message *Message, err error) {
    message = new(Message)
    err = b.callMethod(b.getMethodURL("forwardMessage"), request, message)

    return
}

func (b *Bot) SendPhoto(request *SendPhoto) (message *Message, err error) {
    message = new(Message)
    err = b.callMethod(b.getMethodURL("sendPhoto"), request, message)

    return
}

func (b *Bot) SendAudio(request *SendAudio) (message *Message, err error) {
    message = new(Message)
    err = b.callMethod(b.getMethodURL("sendAudio"), request, message)

    return
}

func (b *Bot) SendDocument(request *SendDocument) (message *Message, err error) {
    message = new(Message)
    err = b.callMethod(b.getMethodURL("sendDocument"), request, message)

    return
}

func (b *Bot) SendVideo(request *SendVideo) (message *Message, err error) {
    message = new(Message)
    err = b.callMethod(b.getMethodURL("sendVideo"), request, message)

    return
}

func (b *Bot) SendVoice(request *SendVoice) (message *Message, err error) {
    message = new(Message)
    err = b.callMethod(b.getMethodURL("sendVoice"), request, message)

    return
}

func (b *Bot) SendVideoNote(request *SendVideoNote) (message *Message, err error) {
    message = new(Message)
    err = b.callMethod(b.getMethodURL("sendVideoNote"), request, message)

    return
}

func (b *Bot) SendMediaGroup(request *SendMediaGroup) (message *Message, err error) {
    message = new(Message)
    err = b.callMethod(b.getMethodURL("sendMediaGroup"), request, message)

    return
}

func (b *Bot) SendLocation(request *SendLocation) (message *Message, err error) {
    message = new(Message)
    err = b.callMethod(b.getMethodURL("sendLocation"), request, message)

    return
}

func (b *Bot) EditMessageLiveLocation(request *EditMessageLiveLocation) (message *Message, err error) {
    message = new(Message)
    err = b.callMethod(b.getMethodURL("editMessageLiveLocation"), request, message)

    return
}

func (b *Bot) StopMessageLiveLocation(request *StopMessageLiveLocation) (message *Message, err error) {
    message = new(Message)
    err = b.callMethod(b.getMethodURL("stopMessageLiveLocation"), request, message)

    return
}

func (b *Bot) SendVenue(request *SendVenue) (message *Message, err error) {
    message = new(Message)
    err = b.callMethod(b.getMethodURL("sendVenue"), request, message)

    return
}

func (b *Bot) SendContact(request *SendContact) (message *Message, err error) {
    message = new(Message)
    err = b.callMethod(b.getMethodURL("sendContact"), request, message)

    return
}

func (b *Bot) SendChatAction(request *SendChatAction) (success bool, err error) {
    err = b.callMethod(b.getMethodURL("sendChatAction"), request, &success)

    return
}

func (b *Bot) GetUserProfilePhotos(request *GetUserProfilePhotos) (photos *UserProfilePhotos, err error) {
    photos = new(UserProfilePhotos)
    err = b.callMethod(b.getMethodURL("getUserProfilePhotos"), request, photos)

    return
}

func (b *Bot) GetFile(request *GetFile) (file *File, err error) {
    file = new(File)
    err = b.callMethod(b.getMethodURL("getFile"), request, file)

    return
}

func (b *Bot) KickChatMember(request *KickChatMember) (success bool, err error) {
    err = b.callMethod(b.getMethodURL("kickChatMember"), request, &success)

    return
}

func (b *Bot) UnbanChatMember(request *UnbanChatMember) (success bool, err error) {
    err = b.callMethod(b.getMethodURL("unbanChatMember"), request, &success)

    return
}

func (b *Bot) RestrictChatMember(request *RestrictChatMember) (success bool, err error) {
    err = b.callMethod(b.getMethodURL("restrictChatMember"), request, &success)

    return
}

func (b *Bot) PromoteChatMember(request *PromoteChatMember) (success bool, err error) {
    err = b.callMethod(b.getMethodURL("promoteChatMember"), request, &success)

    return
}

func (b *Bot) ExportChatInviteLink(request *ExportChatInviteLink) (link string, err error) {
    err = b.callMethod(b.getMethodURL("exportChatInviteLink"), request, &link)

    return
}

func (b *Bot) SetChatPhoto(request *SetChatPhoto) (success bool, err error) {
    err = b.callMethod(b.getMethodURL("setChatPhoto"), request, &success)

    return
}

func (b *Bot) DeleteChatPhoto(request *DeleteChatPhoto) (success bool, err error) {
    err = b.callMethod(b.getMethodURL("deleteChatPhoto"), request, &success)

    return
}

func (b *Bot) SetChatTitle(request *SetChatTitle) (success bool, err error) {
    err = b.callMethod(b.getMethodURL("setChatTitle"), request, &success)

    return
}

func (b *Bot) SetChatDescription(request *SetChatDescription) (success bool, err error) {
    err = b.callMethod(b.getMethodURL("setChatDescription"), request, &success)

    return
}

func (b *Bot) PinChatMessage(request *PinChatMessage) (success bool, err error) {
    err = b.callMethod(b.getMethodURL("pinChatMessage"), request, &success)

    return
}

func (b *Bot) UnpinChatMessage(request *UnpinChatMessage) (success bool, err error) {
    err = b.callMethod(b.getMethodURL("unpinChatMessage"), request, &success)

    return
}

func (b *Bot) LeaveChat(request *LeaveChat) (success bool, err error) {
    err = b.callMethod(b.getMethodURL("leaveChat"), request, &success)

    return
}

func (b *Bot) GetChat(request *GetChat) (chat *Chat, err error) {
    chat = new(Chat)
    err = b.callMethod(b.getMethodURL("getChat"), request, chat)

    return
}

func (b *Bot) GetChatAdministrators(request *GetChatAdministrators) (members []ChatMember, err error) {
    err = b.callMethod(b.getMethodURL("getChatAdministrators"), request, &members)

    return
}

func (b *Bot) GetChatMembersCount(request *GetChatMembersCount) (count int, err error) {
    err = b.callMethod(b.getMethodURL("getChatMembersCount"), request, &count)

    return
}

func (b *Bot) GetChatMember(request *GetChatMember) (member *ChatMember, err error) {
    member = new(ChatMember)
    err = b.callMethod(b.getMethodURL("setChatDescription"), request, member)

    return
}

func (b *Bot) SetChatStickerSet(request *SetChatStickerSet) (success bool, err error) {
    err = b.callMethod(b.getMethodURL("setChatStickerSet"), request, &success)

    return
}

func (b *Bot) DeleteChatStickerSet(request *DeleteChatStickerSet) (success bool, err error) {
    err = b.callMethod(b.getMethodURL("deleteChatStickerSet"), request, &success)

    return
}

func (b *Bot) AnswerCallbackQuery(request *AnswerCallbackQuery) (success bool, err error) {
    err = b.callMethod(b.getMethodURL("answerCallbackQuery"), request, &success)

    return
}

func (b *Bot) EditMessageText(request *EditMessageText) (message *Message, err error) {
    message = new(Message)
    err = b.callMethod(b.getMethodURL("editMessageText"), request, message)

    return
}

func (b *Bot) EditMessageCaption(request *EditMessageCaption) (message *Message, err error) {
    message = new(Message)
    err = b.callMethod(b.getMethodURL("editMessageCaption"), request, message)

    return
}

func (b *Bot) EditMessageReplyMarkup(request *EditMessageReplyMarkup) (message *Message, err error) {
    message = new(Message)
    err = b.callMethod(b.getMethodURL("editMessageReplyMarkup"), request, message)

    return
}

func (b *Bot) DeleteMessage(request *DeleteMessage) (success bool, err error) {
    err = b.callMethod(b.getMethodURL("deleteMessage"), request, &success)

    return
}

func (b *Bot) SendSticker(request *SendSticker) (message *Message, err error) {
    message = new(Message)
    err = b.callMethod(b.getMethodURL("sendSticker"), request, message)

    return
}

func (b *Bot) GetStickerSet(request *GetStickerSet) (set *StickerSet, err error) {
    set = new(StickerSet)
    err = b.callMethod(b.getMethodURL("getStickerSet"), request, set)

    return
}

func (b *Bot) UploadStickerFile(request *UploadStickerFile) (file *File, err error) {
    file = new(File)
    err = b.callMethod(b.getMethodURL("uploadStickerFile"), request, file)

    return
}

func (b *Bot) CreateNewStickerSet(request *CreateNewStickerSet) (success bool, err error) {
    err = b.callMethod(b.getMethodURL("createNewStickerSet"), request, &success)

    return
}

func (b *Bot) AddStickerToSet(request *AddStickerToSet) (success bool, err error) {
    err = b.callMethod(b.getMethodURL("addStickerToSet"), request, &success)

    return
}

func (b *Bot) SetStickerPositionInSet(request *SetStickerPositionInSet) (success bool, err error) {
    err = b.callMethod(b.getMethodURL("setStickerPositionInSet"), request, &success)

    return
}

func (b *Bot) DeleteStickerFromSet(request *DeleteStickerFromSet) (success bool, err error) {
    err = b.callMethod(b.getMethodURL("deleteStickerFromSet"), request, &success)

    return
}

func (b *Bot) AnswerInlineQuery(request *AnswerInlineQuery) (success bool, err error) {
    err = b.callMethod(b.getMethodURL("answerInlineQuery"), request, &success)

    return
}

func (b *Bot) SendInvoice(request *SendInvoice) (message *Message, err error) {
    message = new(Message)
    err = b.callMethod(b.getMethodURL("sendInvoice"), request, message)

    return
}

func (b *Bot) AnswerShippingQuery(request *AnswerShippingQuery) (success bool, err error) {
    err = b.callMethod(b.getMethodURL("answerShippingQuery"), request, &success)

    return
}

func (b *Bot) AnswerPreCheckoutQuery(request *AnswerPreCheckoutQuery) (success bool, err error) {
    err = b.callMethod(b.getMethodURL("answerPreCheckoutQuery"), request, &success)

    return
}

func (b *Bot) SetPassportDataErrors(request *SetPassportDataErrors) (success bool, err error) {
    err = b.callMethod(b.getMethodURL("setPassportDataErrors"), request, &success)

    return
}

func (b *Bot) SendGame(request *SendGame) (message *Message, err error) {
    message = new(Message)
    err = b.callMethod(b.getMethodURL("sendGame"), request, message)

    return
}

func (b *Bot) SetGameScore(request *SetGameScore) (message *Message, err error) {
    message = new(Message)
    err = b.callMethod(b.getMethodURL("setGameScore"), request, message)

    return
}

func (b *Bot) GetGameHighScores(request *GetGameHighScores) (highScores []GameHighScore, err error) {
    err = b.callMethod(b.getMethodURL("getGameHighScores"), request, &highScores)

    return
}

func (b *Bot) getMethodURL(method string) string {
    return "https://api.telegram.org/bot" + b.opts.Token + "/" + method
}

func (b *Bot) callMethod(methodUrl string, request Request, response interface{}) (err error) {
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

    if httpResponse, err = b.opts.Client.Post(methodUrl, contentType, query); err != nil {
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

func (b *Bot) getForm(request Request) (contentType string, query *bytes.Buffer, err error) {
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

func (b *Bot) getFormMultipart(request Request) (contentType string, query *bytes.Buffer, err error) {
    var values map[string]interface{}

    if values, err = request.GetValues(); err != nil {
        return
    }

    query = new(bytes.Buffer)

    var mw = multipart.NewWriter(query)
    var fw io.Writer
    var data io.Reader
    var ok bool

    for key, value := range values {
        if data, ok = value.(io.Reader); ok {
            if fw, err = mw.CreateFormFile(key, key); err != nil {
                return
            }

            if _, err = io.Copy(fw, data); err != nil {
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
