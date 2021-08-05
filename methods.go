package telegram

import (
	"github.com/temoon/go-telegram-bots-api/requests"
)

func (b *Bot) AddStickerToSet(request *requests.AddStickerToSet) (response bool, err error) {
	err = b.callMethod(b.getMethodURL("addStickerToSet"), request, &response)
	return
}

func (b *Bot) AnswerCallbackQuery(request *requests.AnswerCallbackQuery) (response bool, err error) {
	err = b.callMethod(b.getMethodURL("answerCallbackQuery"), request, &response)
	return
}

func (b *Bot) AnswerInlineQuery(request *requests.AnswerInlineQuery) (response bool, err error) {
	err = b.callMethod(b.getMethodURL("answerInlineQuery"), request, &response)
	return
}

func (b *Bot) AnswerPreCheckoutQuery(request *requests.AnswerPreCheckoutQuery) (response bool, err error) {
	err = b.callMethod(b.getMethodURL("answerPreCheckoutQuery"), request, &response)
	return
}

func (b *Bot) AnswerShippingQuery(request *requests.AnswerShippingQuery) (response bool, err error) {
	err = b.callMethod(b.getMethodURL("answerShippingQuery"), request, &response)
	return
}

func (b *Bot) Close() (response bool, err error) {
	err = b.callMethod(b.getMethodURL("close"), nil, &response)
	return
}

func (b *Bot) CopyMessage(request *requests.CopyMessage) (response *MessageId, err error) {
	response = new(MessageId)
	err = b.callMethod(b.getMethodURL("copyMessage"), request, response)
	return
}

func (b *Bot) CreateChatInviteLink(request *requests.CreateChatInviteLink) (response string, err error) {
	err = b.callMethod(b.getMethodURL("createChatInviteLink"), request, &response)
	return
}

func (b *Bot) CreateNewStickerSet(request *requests.CreateNewStickerSet) (response bool, err error) {
	err = b.callMethod(b.getMethodURL("createNewStickerSet"), request, &response)
	return
}

func (b *Bot) DeleteChatPhoto(request *requests.DeleteChatPhoto) (response bool, err error) {
	err = b.callMethod(b.getMethodURL("deleteChatPhoto"), request, &response)
	return
}

func (b *Bot) DeleteChatStickerSet(request *requests.DeleteChatStickerSet) (response bool, err error) {
	err = b.callMethod(b.getMethodURL("deleteChatStickerSet"), request, &response)
	return
}

func (b *Bot) DeleteMessage(request *requests.DeleteMessage) (response bool, err error) {
	err = b.callMethod(b.getMethodURL("deleteMessage"), request, &response)
	return
}

func (b *Bot) DeleteStickerFromSet(request *requests.DeleteStickerFromSet) (response bool, err error) {
	err = b.callMethod(b.getMethodURL("deleteStickerFromSet"), request, &response)
	return
}

func (b *Bot) DeleteWebhook(request *requests.DeleteWebhook) (response bool, err error) {
	err = b.callMethod(b.getMethodURL("deleteWebhook"), request, &response)
	return
}

func (b *Bot) EditChatInviteLink(request *requests.EditChatInviteLink) (response string, err error) {
	err = b.callMethod(b.getMethodURL("editChatInviteLink"), request, &response)
	return
}

func (b *Bot) EditMessageCaption(request *requests.EditMessageCaption) (response interface{}, err error) {
	err = b.callMethod(b.getMethodURL("editMessageCaption"), request, &response)
	return
}

func (b *Bot) EditMessageLiveLocation(request *requests.EditMessageLiveLocation) (response interface{}, err error) {
	err = b.callMethod(b.getMethodURL("editMessageLiveLocation"), request, &response)
	return
}

func (b *Bot) EditMessageMedia(request *requests.EditMessageMedia) (response interface{}, err error) {
	err = b.callMethod(b.getMethodURL("editMessageMedia"), request, &response)
	return
}

func (b *Bot) EditMessageReplyMarkup(request *requests.EditMessageReplyMarkup) (response interface{}, err error) {
	err = b.callMethod(b.getMethodURL("editMessageReplyMarkup"), request, &response)
	return
}

func (b *Bot) EditMessageText(request *requests.EditMessageText) (response interface{}, err error) {
	err = b.callMethod(b.getMethodURL("editMessageText"), request, &response)
	return
}

func (b *Bot) ExportChatInviteLink(request *requests.ExportChatInviteLink) (response string, err error) {
	err = b.callMethod(b.getMethodURL("exportChatInviteLink"), request, &response)
	return
}

func (b *Bot) ForwardMessage(request *requests.ForwardMessage) (response *Message, err error) {
	response = new(Message)
	err = b.callMethod(b.getMethodURL("forwardMessage"), request, response)
	return
}

func (b *Bot) GetChat(request *requests.GetChat) (response *Chat, err error) {
	response = new(Chat)
	err = b.callMethod(b.getMethodURL("getChat"), request, response)
	return
}

func (b *Bot) GetChatAdministrators(request *requests.GetChatAdministrators) (response []ChatMember, err error) {
	err = b.callMethod(b.getMethodURL("getChatAdministrators"), request, &response)
	return
}

func (b *Bot) GetChatMember(request *requests.GetChatMember) (response *ChatMember, err error) {
	response = new(ChatMember)
	err = b.callMethod(b.getMethodURL("getChatMember"), request, response)
	return
}

func (b *Bot) GetChatMembersCount(request *requests.GetChatMembersCount) (response uint64, err error) {
	err = b.callMethod(b.getMethodURL("getChatMembersCount"), request, &response)
	return
}

func (b *Bot) GetFile(request *requests.GetFile) (response *File, err error) {
	response = new(File)
	err = b.callMethod(b.getMethodURL("getFile"), request, response)
	return
}

func (b *Bot) GetGameHighScores(request *requests.GetGameHighScores) (response []GameHighScore, err error) {
	err = b.callMethod(b.getMethodURL("getGameHighScores"), request, &response)
	return
}

func (b *Bot) GetMe() (response *User, err error) {
	response = new(User)
	err = b.callMethod(b.getMethodURL("getMe"), nil, response)
	return
}

func (b *Bot) GetMyCommands() (response []BotCommand, err error) {
	err = b.callMethod(b.getMethodURL("getMyCommands"), nil, &response)
	return
}

func (b *Bot) GetStickerSet(request *requests.GetStickerSet) (response *StickerSet, err error) {
	response = new(StickerSet)
	err = b.callMethod(b.getMethodURL("getStickerSet"), request, response)
	return
}

func (b *Bot) GetUpdates(request *requests.GetUpdates) (response []Update, err error) {
	err = b.callMethod(b.getMethodURL("getUpdates"), request, &response)
	return
}

func (b *Bot) GetUserProfilePhotos(request *requests.GetUserProfilePhotos) (response *UserProfilePhotos, err error) {
	response = new(UserProfilePhotos)
	err = b.callMethod(b.getMethodURL("getUserProfilePhotos"), request, response)
	return
}

func (b *Bot) GetWebhookInfo() (response *WebhookInfo, err error) {
	response = new(WebhookInfo)
	err = b.callMethod(b.getMethodURL("getWebhookInfo"), nil, response)
	return
}

func (b *Bot) KickChatMember(request *requests.KickChatMember) (response bool, err error) {
	err = b.callMethod(b.getMethodURL("kickChatMember"), request, &response)
	return
}

func (b *Bot) LeaveChat(request *requests.LeaveChat) (response bool, err error) {
	err = b.callMethod(b.getMethodURL("leaveChat"), request, &response)
	return
}

func (b *Bot) LogOut() (response bool, err error) {
	err = b.callMethod(b.getMethodURL("logOut"), nil, &response)
	return
}

func (b *Bot) PinChatMessage(request *requests.PinChatMessage) (response bool, err error) {
	err = b.callMethod(b.getMethodURL("pinChatMessage"), request, &response)
	return
}

func (b *Bot) PromoteChatMember(request *requests.PromoteChatMember) (response bool, err error) {
	err = b.callMethod(b.getMethodURL("promoteChatMember"), request, &response)
	return
}

func (b *Bot) RestrictChatMember(request *requests.RestrictChatMember) (response bool, err error) {
	err = b.callMethod(b.getMethodURL("restrictChatMember"), request, &response)
	return
}

func (b *Bot) RevokeChatInviteLink(request *requests.RevokeChatInviteLink) (response string, err error) {
	err = b.callMethod(b.getMethodURL("revokeChatInviteLink"), request, &response)
	return
}

func (b *Bot) SendAnimation(request *requests.SendAnimation) (response *Message, err error) {
	response = new(Message)
	err = b.callMethod(b.getMethodURL("sendAnimation"), request, response)
	return
}

func (b *Bot) SendAudio(request *requests.SendAudio) (response *Message, err error) {
	response = new(Message)
	err = b.callMethod(b.getMethodURL("sendAudio"), request, response)
	return
}

func (b *Bot) SendChatAction(request *requests.SendChatAction) (response bool, err error) {
	err = b.callMethod(b.getMethodURL("sendChatAction"), request, &response)
	return
}

func (b *Bot) SendContact(request *requests.SendContact) (response *Message, err error) {
	response = new(Message)
	err = b.callMethod(b.getMethodURL("sendContact"), request, response)
	return
}

func (b *Bot) SendDice(request *requests.SendDice) (response *Message, err error) {
	response = new(Message)
	err = b.callMethod(b.getMethodURL("sendDice"), request, response)
	return
}

func (b *Bot) SendDocument(request *requests.SendDocument) (response *Message, err error) {
	response = new(Message)
	err = b.callMethod(b.getMethodURL("sendDocument"), request, response)
	return
}

func (b *Bot) SendGame(request *requests.SendGame) (response *Message, err error) {
	response = new(Message)
	err = b.callMethod(b.getMethodURL("sendGame"), request, response)
	return
}

func (b *Bot) SendInvoice(request *requests.SendInvoice) (response *Message, err error) {
	response = new(Message)
	err = b.callMethod(b.getMethodURL("sendInvoice"), request, response)
	return
}

func (b *Bot) SendLocation(request *requests.SendLocation) (response *Message, err error) {
	response = new(Message)
	err = b.callMethod(b.getMethodURL("sendLocation"), request, response)
	return
}

func (b *Bot) SendMediaGroup(request *requests.SendMediaGroup) (response []Message, err error) {
	err = b.callMethod(b.getMethodURL("sendMediaGroup"), request, &response)
	return
}

func (b *Bot) SendMessage(request *requests.SendMessage) (response *Message, err error) {
	response = new(Message)
	err = b.callMethod(b.getMethodURL("sendMessage"), request, response)
	return
}

func (b *Bot) SendPhoto(request *requests.SendPhoto) (response *Message, err error) {
	response = new(Message)
	err = b.callMethod(b.getMethodURL("sendPhoto"), request, response)
	return
}

func (b *Bot) SendPoll(request *requests.SendPoll) (response *Message, err error) {
	response = new(Message)
	err = b.callMethod(b.getMethodURL("sendPoll"), request, response)
	return
}

func (b *Bot) SendSticker(request *requests.SendSticker) (response *Message, err error) {
	response = new(Message)
	err = b.callMethod(b.getMethodURL("sendSticker"), request, response)
	return
}

func (b *Bot) SendVenue(request *requests.SendVenue) (response *Message, err error) {
	response = new(Message)
	err = b.callMethod(b.getMethodURL("sendVenue"), request, response)
	return
}

func (b *Bot) SendVideo(request *requests.SendVideo) (response *Message, err error) {
	response = new(Message)
	err = b.callMethod(b.getMethodURL("sendVideo"), request, response)
	return
}

func (b *Bot) SendVideoNote(request *requests.SendVideoNote) (response *Message, err error) {
	response = new(Message)
	err = b.callMethod(b.getMethodURL("sendVideoNote"), request, response)
	return
}

func (b *Bot) SendVoice(request *requests.SendVoice) (response *Message, err error) {
	response = new(Message)
	err = b.callMethod(b.getMethodURL("sendVoice"), request, response)
	return
}

func (b *Bot) SetChatAdministratorCustomTitle(request *requests.SetChatAdministratorCustomTitle) (response bool, err error) {
	err = b.callMethod(b.getMethodURL("setChatAdministratorCustomTitle"), request, &response)
	return
}

func (b *Bot) SetChatDescription(request *requests.SetChatDescription) (response bool, err error) {
	err = b.callMethod(b.getMethodURL("setChatDescription"), request, &response)
	return
}

func (b *Bot) SetChatPermissions(request *requests.SetChatPermissions) (response bool, err error) {
	err = b.callMethod(b.getMethodURL("setChatPermissions"), request, &response)
	return
}

func (b *Bot) SetChatPhoto(request *requests.SetChatPhoto) (response bool, err error) {
	err = b.callMethod(b.getMethodURL("setChatPhoto"), request, &response)
	return
}

func (b *Bot) SetChatStickerSet(request *requests.SetChatStickerSet) (response bool, err error) {
	err = b.callMethod(b.getMethodURL("setChatStickerSet"), request, &response)
	return
}

func (b *Bot) SetChatTitle(request *requests.SetChatTitle) (response bool, err error) {
	err = b.callMethod(b.getMethodURL("setChatTitle"), request, &response)
	return
}

func (b *Bot) SetGameScore(request *requests.SetGameScore) (response interface{}, err error) {
	err = b.callMethod(b.getMethodURL("setGameScore"), request, &response)
	return
}

func (b *Bot) SetMyCommands(request *requests.SetMyCommands) (response bool, err error) {
	err = b.callMethod(b.getMethodURL("setMyCommands"), request, &response)
	return
}

func (b *Bot) SetPassportDataErrors(request *requests.SetPassportDataErrors) (response bool, err error) {
	err = b.callMethod(b.getMethodURL("setPassportDataErrors"), request, &response)
	return
}

func (b *Bot) SetStickerPositionInSet(request *requests.SetStickerPositionInSet) (response bool, err error) {
	err = b.callMethod(b.getMethodURL("setStickerPositionInSet"), request, &response)
	return
}

func (b *Bot) SetStickerSetThumb(request *requests.SetStickerSetThumb) (response bool, err error) {
	err = b.callMethod(b.getMethodURL("setStickerSetThumb"), request, &response)
	return
}

func (b *Bot) SetWebhook(request *requests.SetWebhook) (response bool, err error) {
	err = b.callMethod(b.getMethodURL("setWebhook"), request, &response)
	return
}

func (b *Bot) StopMessageLiveLocation(request *requests.StopMessageLiveLocation) (response interface{}, err error) {
	err = b.callMethod(b.getMethodURL("stopMessageLiveLocation"), request, &response)
	return
}

func (b *Bot) StopPoll(request *requests.StopPoll) (response *Poll, err error) {
	response = new(Poll)
	err = b.callMethod(b.getMethodURL("stopPoll"), request, response)
	return
}

func (b *Bot) UnbanChatMember(request *requests.UnbanChatMember) (response bool, err error) {
	err = b.callMethod(b.getMethodURL("unbanChatMember"), request, &response)
	return
}

func (b *Bot) UnpinAllChatMessages(request *requests.UnpinAllChatMessages) (response bool, err error) {
	err = b.callMethod(b.getMethodURL("unpinAllChatMessages"), request, &response)
	return
}

func (b *Bot) UnpinChatMessage(request *requests.UnpinChatMessage) (response bool, err error) {
	err = b.callMethod(b.getMethodURL("unpinChatMessage"), request, &response)
	return
}

func (b *Bot) UploadStickerFile(request *requests.UploadStickerFile) (response *File, err error) {
	response = new(File)
	err = b.callMethod(b.getMethodURL("uploadStickerFile"), request, response)
	return
}
