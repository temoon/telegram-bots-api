package telegram

type Animation struct {
	Duration     uint64     `json:"duration"`
	FileId       string     `json:"file_id"`
	FileUniqueId string     `json:"file_unique_id"`
	Height       uint64     `json:"height"`
	Width        uint64     `json:"width"`
	FileName     string     `json:"file_name,omitempty"`
	FileSize     uint64     `json:"file_size,omitempty"`
	MimeType     string     `json:"mime_type,omitempty"`
	Thumb        *PhotoSize `json:"thumb,omitempty"`
}

type Audio struct {
	Duration     uint64     `json:"duration"`
	FileId       string     `json:"file_id"`
	FileUniqueId string     `json:"file_unique_id"`
	FileName     string     `json:"file_name,omitempty"`
	FileSize     uint64     `json:"file_size,omitempty"`
	MimeType     string     `json:"mime_type,omitempty"`
	Performer    string     `json:"performer,omitempty"`
	Thumb        *PhotoSize `json:"thumb,omitempty"`
	Title        string     `json:"title,omitempty"`
}

type BotCommand struct {
	Command     string `json:"command"`
	Description string `json:"description"`
}

type CallbackGame struct {
	// Placeholder
}

type CallbackQuery struct {
	ChatInstance    string   `json:"chat_instance"`
	From            User     `json:"from"`
	Id              string   `json:"id"`
	Data            string   `json:"data,omitempty"`
	GameShortName   string   `json:"game_short_name,omitempty"`
	InlineMessageId string   `json:"inline_message_id,omitempty"`
	Message         *Message `json:"message,omitempty"`
}

type Chat struct {
	Id               uint64           `json:"id"`
	Type             string           `json:"type"`
	Bio              string           `json:"bio,omitempty"`
	CanSetStickerSet bool             `json:"can_set_sticker_set,omitempty"`
	Description      string           `json:"description,omitempty"`
	FirstName        string           `json:"first_name,omitempty"`
	InviteLink       string           `json:"invite_link,omitempty"`
	LastName         string           `json:"last_name,omitempty"`
	LinkedChatId     uint64           `json:"linked_chat_id,omitempty"`
	Location         *ChatLocation    `json:"location,omitempty"`
	Permissions      *ChatPermissions `json:"permissions,omitempty"`
	Photo            *ChatPhoto       `json:"photo,omitempty"`
	PinnedMessage    *Message         `json:"pinned_message,omitempty"`
	SlowModeDelay    uint64           `json:"slow_mode_delay,omitempty"`
	StickerSetName   string           `json:"sticker_set_name,omitempty"`
	Title            string           `json:"title,omitempty"`
	Username         string           `json:"username,omitempty"`
}

type ChatInviteLink struct {
	Creator     User   `json:"creator"`
	InviteLink  string `json:"invite_link"`
	IsPrimary   bool   `json:"is_primary"`
	IsRevoked   bool   `json:"is_revoked"`
	ExpireDate  uint64 `json:"expire_date,omitempty"`
	MemberLimit uint64 `json:"member_limit,omitempty"`
}

type ChatLocation struct {
	Address  string   `json:"address"`
	Location Location `json:"location"`
}

type ChatMember struct {
	Status                string `json:"status"`
	User                  User   `json:"user"`
	CanAddWebPagePreviews bool   `json:"can_add_web_page_previews,omitempty"`
	CanBeEdited           bool   `json:"can_be_edited,omitempty"`
	CanChangeInfo         bool   `json:"can_change_info,omitempty"`
	CanDeleteMessages     bool   `json:"can_delete_messages,omitempty"`
	CanEditMessages       bool   `json:"can_edit_messages,omitempty"`
	CanInviteUsers        bool   `json:"can_invite_users,omitempty"`
	CanManageChat         bool   `json:"can_manage_chat,omitempty"`
	CanManageVoiceChats   bool   `json:"can_manage_voice_chats,omitempty"`
	CanPinMessages        bool   `json:"can_pin_messages,omitempty"`
	CanPostMessages       bool   `json:"can_post_messages,omitempty"`
	CanPromoteMembers     bool   `json:"can_promote_members,omitempty"`
	CanRestrictMembers    bool   `json:"can_restrict_members,omitempty"`
	CanSendMediaMessages  bool   `json:"can_send_media_messages,omitempty"`
	CanSendMessages       bool   `json:"can_send_messages,omitempty"`
	CanSendOtherMessages  bool   `json:"can_send_other_messages,omitempty"`
	CanSendPolls          bool   `json:"can_send_polls,omitempty"`
	CustomTitle           string `json:"custom_title,omitempty"`
	IsAnonymous           bool   `json:"is_anonymous,omitempty"`
	IsMember              bool   `json:"is_member,omitempty"`
	UntilDate             uint64 `json:"until_date,omitempty"`
}

type ChatMemberUpdated struct {
	Chat          Chat            `json:"chat"`
	Date          uint64          `json:"date"`
	From          User            `json:"from"`
	NewChatMember ChatMember      `json:"new_chat_member"`
	OldChatMember ChatMember      `json:"old_chat_member"`
	InviteLink    *ChatInviteLink `json:"invite_link,omitempty"`
}

type ChatPermissions struct {
	CanAddWebPagePreviews bool `json:"can_add_web_page_previews,omitempty"`
	CanChangeInfo         bool `json:"can_change_info,omitempty"`
	CanInviteUsers        bool `json:"can_invite_users,omitempty"`
	CanPinMessages        bool `json:"can_pin_messages,omitempty"`
	CanSendMediaMessages  bool `json:"can_send_media_messages,omitempty"`
	CanSendMessages       bool `json:"can_send_messages,omitempty"`
	CanSendOtherMessages  bool `json:"can_send_other_messages,omitempty"`
	CanSendPolls          bool `json:"can_send_polls,omitempty"`
}

type ChatPhoto struct {
	BigFileId         string `json:"big_file_id"`
	BigFileUniqueId   string `json:"big_file_unique_id"`
	SmallFileId       string `json:"small_file_id"`
	SmallFileUniqueId string `json:"small_file_unique_id"`
}

type ChosenInlineResult struct {
	From            User      `json:"from"`
	Query           string    `json:"query"`
	ResultId        string    `json:"result_id"`
	InlineMessageId string    `json:"inline_message_id,omitempty"`
	Location        *Location `json:"location,omitempty"`
}

type Contact struct {
	FirstName   string `json:"first_name"`
	PhoneNumber string `json:"phone_number"`
	LastName    string `json:"last_name,omitempty"`
	UserId      uint64 `json:"user_id,omitempty"`
	Vcard       string `json:"vcard,omitempty"`
}

type Dice struct {
	Emoji string `json:"emoji"`
	Value uint64 `json:"value"`
}

type Document struct {
	FileId       string     `json:"file_id"`
	FileUniqueId string     `json:"file_unique_id"`
	FileName     string     `json:"file_name,omitempty"`
	FileSize     uint64     `json:"file_size,omitempty"`
	MimeType     string     `json:"mime_type,omitempty"`
	Thumb        *PhotoSize `json:"thumb,omitempty"`
}

type EncryptedCredentials struct {
	Data   string `json:"data"`
	Hash   string `json:"hash"`
	Secret string `json:"secret"`
}

type EncryptedPassportElement struct {
	Hash        string         `json:"hash"`
	Type        string         `json:"type"`
	Data        string         `json:"data,omitempty"`
	Email       string         `json:"email,omitempty"`
	Files       []PassportFile `json:"files,omitempty"`
	FrontSide   *PassportFile  `json:"front_side,omitempty"`
	PhoneNumber string         `json:"phone_number,omitempty"`
	ReverseSide *PassportFile  `json:"reverse_side,omitempty"`
	Selfie      *PassportFile  `json:"selfie,omitempty"`
	Translation []PassportFile `json:"translation,omitempty"`
}

type Error struct {
	Description string              `json:"description"`
	ErrorCode   uint64              `json:"error_code"`
	Ok          bool                `json:"ok"`
	Parameters  *ResponseParameters `json:"parameters,omitempty"`
}

type File struct {
	FileId       string `json:"file_id"`
	FileUniqueId string `json:"file_unique_id"`
	FilePath     string `json:"file_path,omitempty"`
	FileSize     uint64 `json:"file_size,omitempty"`
}

type ForceReply struct {
	ForceReply bool `json:"force_reply"`
	Selective  bool `json:"selective,omitempty"`
}

type Game struct {
	Description  string          `json:"description"`
	Photo        []PhotoSize     `json:"photo"`
	Title        string          `json:"title"`
	Animation    *Animation      `json:"animation,omitempty"`
	Text         string          `json:"text,omitempty"`
	TextEntities []MessageEntity `json:"text_entities,omitempty"`
}

type GameHighScore struct {
	Position uint64 `json:"position"`
	Score    uint64 `json:"score"`
	User     User   `json:"user"`
}

type InlineKeyboardButton struct {
	Text                         string        `json:"text"`
	CallbackData                 string        `json:"callback_data,omitempty"`
	CallbackGame                 *CallbackGame `json:"callback_game,omitempty"`
	LoginUrl                     *LoginUrl     `json:"login_url,omitempty"`
	Pay                          bool          `json:"pay,omitempty"`
	SwitchInlineQuery            string        `json:"switch_inline_query,omitempty"`
	SwitchInlineQueryCurrentChat string        `json:"switch_inline_query_current_chat,omitempty"`
	Url                          string        `json:"url,omitempty"`
}

type InlineKeyboardMarkup struct {
	InlineKeyboard [][]InlineKeyboardButton `json:"inline_keyboard"`
}

type InlineQuery struct {
	From     User      `json:"from"`
	Id       string    `json:"id"`
	Offset   string    `json:"offset"`
	Query    string    `json:"query"`
	ChatType string    `json:"chat_type,omitempty"`
	Location *Location `json:"location,omitempty"`
}

type InlineQueryResult struct {
	// Placeholder
}

type InlineQueryResultArticle struct {
	Id                  string                `json:"id"`
	InputMessageContent InputMessageContent   `json:"input_message_content"`
	Title               string                `json:"title"`
	Type                string                `json:"type"`
	Description         string                `json:"description,omitempty"`
	HideUrl             bool                  `json:"hide_url,omitempty"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	ThumbHeight         uint64                `json:"thumb_height,omitempty"`
	ThumbUrl            string                `json:"thumb_url,omitempty"`
	ThumbWidth          uint64                `json:"thumb_width,omitempty"`
	Url                 string                `json:"url,omitempty"`
}

type InlineQueryResultAudio struct {
	AudioUrl            string                `json:"audio_url"`
	Id                  string                `json:"id"`
	Title               string                `json:"title"`
	Type                string                `json:"type"`
	AudioDuration       uint64                `json:"audio_duration,omitempty"`
	Caption             string                `json:"caption,omitempty"`
	CaptionEntities     []MessageEntity       `json:"caption_entities,omitempty"`
	InputMessageContent *InputMessageContent  `json:"input_message_content,omitempty"`
	ParseMode           string                `json:"parse_mode,omitempty"`
	Performer           string                `json:"performer,omitempty"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

type InlineQueryResultCachedAudio struct {
	AudioFileId         string                `json:"audio_file_id"`
	Id                  string                `json:"id"`
	Type                string                `json:"type"`
	Caption             string                `json:"caption,omitempty"`
	CaptionEntities     []MessageEntity       `json:"caption_entities,omitempty"`
	InputMessageContent *InputMessageContent  `json:"input_message_content,omitempty"`
	ParseMode           string                `json:"parse_mode,omitempty"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

type InlineQueryResultCachedDocument struct {
	DocumentFileId      string                `json:"document_file_id"`
	Id                  string                `json:"id"`
	Title               string                `json:"title"`
	Type                string                `json:"type"`
	Caption             string                `json:"caption,omitempty"`
	CaptionEntities     []MessageEntity       `json:"caption_entities,omitempty"`
	Description         string                `json:"description,omitempty"`
	InputMessageContent *InputMessageContent  `json:"input_message_content,omitempty"`
	ParseMode           string                `json:"parse_mode,omitempty"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

type InlineQueryResultCachedGif struct {
	GifFileId           string                `json:"gif_file_id"`
	Id                  string                `json:"id"`
	Type                string                `json:"type"`
	Caption             string                `json:"caption,omitempty"`
	CaptionEntities     []MessageEntity       `json:"caption_entities,omitempty"`
	InputMessageContent *InputMessageContent  `json:"input_message_content,omitempty"`
	ParseMode           string                `json:"parse_mode,omitempty"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	Title               string                `json:"title,omitempty"`
}

type InlineQueryResultCachedMpeg4Gif struct {
	Id                  string                `json:"id"`
	Mpeg4FileId         string                `json:"mpeg4_file_id"`
	Type                string                `json:"type"`
	Caption             string                `json:"caption,omitempty"`
	CaptionEntities     []MessageEntity       `json:"caption_entities,omitempty"`
	InputMessageContent *InputMessageContent  `json:"input_message_content,omitempty"`
	ParseMode           string                `json:"parse_mode,omitempty"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	Title               string                `json:"title,omitempty"`
}

type InlineQueryResultCachedPhoto struct {
	Id                  string                `json:"id"`
	PhotoFileId         string                `json:"photo_file_id"`
	Type                string                `json:"type"`
	Caption             string                `json:"caption,omitempty"`
	CaptionEntities     []MessageEntity       `json:"caption_entities,omitempty"`
	Description         string                `json:"description,omitempty"`
	InputMessageContent *InputMessageContent  `json:"input_message_content,omitempty"`
	ParseMode           string                `json:"parse_mode,omitempty"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	Title               string                `json:"title,omitempty"`
}

type InlineQueryResultCachedSticker struct {
	Id                  string                `json:"id"`
	StickerFileId       string                `json:"sticker_file_id"`
	Type                string                `json:"type"`
	InputMessageContent *InputMessageContent  `json:"input_message_content,omitempty"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

type InlineQueryResultCachedVideo struct {
	Id                  string                `json:"id"`
	Title               string                `json:"title"`
	Type                string                `json:"type"`
	VideoFileId         string                `json:"video_file_id"`
	Caption             string                `json:"caption,omitempty"`
	CaptionEntities     []MessageEntity       `json:"caption_entities,omitempty"`
	Description         string                `json:"description,omitempty"`
	InputMessageContent *InputMessageContent  `json:"input_message_content,omitempty"`
	ParseMode           string                `json:"parse_mode,omitempty"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

type InlineQueryResultCachedVoice struct {
	Id                  string                `json:"id"`
	Title               string                `json:"title"`
	Type                string                `json:"type"`
	VoiceFileId         string                `json:"voice_file_id"`
	Caption             string                `json:"caption,omitempty"`
	CaptionEntities     []MessageEntity       `json:"caption_entities,omitempty"`
	InputMessageContent *InputMessageContent  `json:"input_message_content,omitempty"`
	ParseMode           string                `json:"parse_mode,omitempty"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

type InlineQueryResultContact struct {
	FirstName           string                `json:"first_name"`
	Id                  string                `json:"id"`
	PhoneNumber         string                `json:"phone_number"`
	Type                string                `json:"type"`
	InputMessageContent *InputMessageContent  `json:"input_message_content,omitempty"`
	LastName            string                `json:"last_name,omitempty"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	ThumbHeight         uint64                `json:"thumb_height,omitempty"`
	ThumbUrl            string                `json:"thumb_url,omitempty"`
	ThumbWidth          uint64                `json:"thumb_width,omitempty"`
	Vcard               string                `json:"vcard,omitempty"`
}

type InlineQueryResultDocument struct {
	DocumentUrl         string                `json:"document_url"`
	Id                  string                `json:"id"`
	MimeType            string                `json:"mime_type"`
	Title               string                `json:"title"`
	Type                string                `json:"type"`
	Caption             string                `json:"caption,omitempty"`
	CaptionEntities     []MessageEntity       `json:"caption_entities,omitempty"`
	Description         string                `json:"description,omitempty"`
	InputMessageContent *InputMessageContent  `json:"input_message_content,omitempty"`
	ParseMode           string                `json:"parse_mode,omitempty"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	ThumbHeight         uint64                `json:"thumb_height,omitempty"`
	ThumbUrl            string                `json:"thumb_url,omitempty"`
	ThumbWidth          uint64                `json:"thumb_width,omitempty"`
}

type InlineQueryResultGame struct {
	GameShortName string                `json:"game_short_name"`
	Id            string                `json:"id"`
	Type          string                `json:"type"`
	ReplyMarkup   *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

type InlineQueryResultGif struct {
	GifUrl              string                `json:"gif_url"`
	Id                  string                `json:"id"`
	ThumbUrl            string                `json:"thumb_url"`
	Type                string                `json:"type"`
	Caption             string                `json:"caption,omitempty"`
	CaptionEntities     []MessageEntity       `json:"caption_entities,omitempty"`
	GifDuration         uint64                `json:"gif_duration,omitempty"`
	GifHeight           uint64                `json:"gif_height,omitempty"`
	GifWidth            uint64                `json:"gif_width,omitempty"`
	InputMessageContent *InputMessageContent  `json:"input_message_content,omitempty"`
	ParseMode           string                `json:"parse_mode,omitempty"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	ThumbMimeType       string                `json:"thumb_mime_type,omitempty"`
	Title               string                `json:"title,omitempty"`
}

type InlineQueryResultLocation struct {
	Id                   string                `json:"id"`
	Latitude             float64               `json:"latitude"`
	Longitude            float64               `json:"longitude"`
	Title                string                `json:"title"`
	Type                 string                `json:"type"`
	Heading              uint64                `json:"heading,omitempty"`
	HorizontalAccuracy   float64               `json:"horizontal_accuracy,omitempty"`
	InputMessageContent  *InputMessageContent  `json:"input_message_content,omitempty"`
	LivePeriod           uint64                `json:"live_period,omitempty"`
	ProximityAlertRadius uint64                `json:"proximity_alert_radius,omitempty"`
	ReplyMarkup          *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	ThumbHeight          uint64                `json:"thumb_height,omitempty"`
	ThumbUrl             string                `json:"thumb_url,omitempty"`
	ThumbWidth           uint64                `json:"thumb_width,omitempty"`
}

type InlineQueryResultMpeg4Gif struct {
	Id                  string                `json:"id"`
	Mpeg4Url            string                `json:"mpeg4_url"`
	ThumbUrl            string                `json:"thumb_url"`
	Type                string                `json:"type"`
	Caption             string                `json:"caption,omitempty"`
	CaptionEntities     []MessageEntity       `json:"caption_entities,omitempty"`
	InputMessageContent *InputMessageContent  `json:"input_message_content,omitempty"`
	Mpeg4Duration       uint64                `json:"mpeg4_duration,omitempty"`
	Mpeg4Height         uint64                `json:"mpeg4_height,omitempty"`
	Mpeg4Width          uint64                `json:"mpeg4_width,omitempty"`
	ParseMode           string                `json:"parse_mode,omitempty"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	ThumbMimeType       string                `json:"thumb_mime_type,omitempty"`
	Title               string                `json:"title,omitempty"`
}

type InlineQueryResultPhoto struct {
	Id                  string                `json:"id"`
	PhotoUrl            string                `json:"photo_url"`
	ThumbUrl            string                `json:"thumb_url"`
	Type                string                `json:"type"`
	Caption             string                `json:"caption,omitempty"`
	CaptionEntities     []MessageEntity       `json:"caption_entities,omitempty"`
	Description         string                `json:"description,omitempty"`
	InputMessageContent *InputMessageContent  `json:"input_message_content,omitempty"`
	ParseMode           string                `json:"parse_mode,omitempty"`
	PhotoHeight         uint64                `json:"photo_height,omitempty"`
	PhotoWidth          uint64                `json:"photo_width,omitempty"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	Title               string                `json:"title,omitempty"`
}

type InlineQueryResultVenue struct {
	Address             string                `json:"address"`
	Id                  string                `json:"id"`
	Latitude            float64               `json:"latitude"`
	Longitude           float64               `json:"longitude"`
	Title               string                `json:"title"`
	Type                string                `json:"type"`
	FoursquareId        string                `json:"foursquare_id,omitempty"`
	FoursquareType      string                `json:"foursquare_type,omitempty"`
	GooglePlaceId       string                `json:"google_place_id,omitempty"`
	GooglePlaceType     string                `json:"google_place_type,omitempty"`
	InputMessageContent *InputMessageContent  `json:"input_message_content,omitempty"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	ThumbHeight         uint64                `json:"thumb_height,omitempty"`
	ThumbUrl            string                `json:"thumb_url,omitempty"`
	ThumbWidth          uint64                `json:"thumb_width,omitempty"`
}

type InlineQueryResultVideo struct {
	Id                  string                `json:"id"`
	MimeType            string                `json:"mime_type"`
	ThumbUrl            string                `json:"thumb_url"`
	Title               string                `json:"title"`
	Type                string                `json:"type"`
	VideoUrl            string                `json:"video_url"`
	Caption             string                `json:"caption,omitempty"`
	CaptionEntities     []MessageEntity       `json:"caption_entities,omitempty"`
	Description         string                `json:"description,omitempty"`
	InputMessageContent *InputMessageContent  `json:"input_message_content,omitempty"`
	ParseMode           string                `json:"parse_mode,omitempty"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	VideoDuration       uint64                `json:"video_duration,omitempty"`
	VideoHeight         uint64                `json:"video_height,omitempty"`
	VideoWidth          uint64                `json:"video_width,omitempty"`
}

type InlineQueryResultVoice struct {
	Id                  string                `json:"id"`
	Title               string                `json:"title"`
	Type                string                `json:"type"`
	VoiceUrl            string                `json:"voice_url"`
	Caption             string                `json:"caption,omitempty"`
	CaptionEntities     []MessageEntity       `json:"caption_entities,omitempty"`
	InputMessageContent *InputMessageContent  `json:"input_message_content,omitempty"`
	ParseMode           string                `json:"parse_mode,omitempty"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	VoiceDuration       uint64                `json:"voice_duration,omitempty"`
}

type InputContactMessageContent struct {
	FirstName   string `json:"first_name"`
	PhoneNumber string `json:"phone_number"`
	LastName    string `json:"last_name,omitempty"`
	Vcard       string `json:"vcard,omitempty"`
}

type InputInvoiceMessageContent struct {
	Currency                  string         `json:"currency"`
	Description               string         `json:"description"`
	Payload                   string         `json:"payload"`
	Prices                    []LabeledPrice `json:"prices"`
	ProviderToken             string         `json:"provider_token"`
	Title                     string         `json:"title"`
	IsFlexible                bool           `json:"is_flexible,omitempty"`
	MaxTipAmount              uint64         `json:"max_tip_amount,omitempty"`
	NeedEmail                 bool           `json:"need_email,omitempty"`
	NeedName                  bool           `json:"need_name,omitempty"`
	NeedPhoneNumber           bool           `json:"need_phone_number,omitempty"`
	NeedShippingAddress       bool           `json:"need_shipping_address,omitempty"`
	PhotoHeight               uint64         `json:"photo_height,omitempty"`
	PhotoSize                 uint64         `json:"photo_size,omitempty"`
	PhotoUrl                  string         `json:"photo_url,omitempty"`
	PhotoWidth                uint64         `json:"photo_width,omitempty"`
	ProviderData              string         `json:"provider_data,omitempty"`
	SendEmailToProvider       bool           `json:"send_email_to_provider,omitempty"`
	SendPhoneNumberToProvider bool           `json:"send_phone_number_to_provider,omitempty"`
	SuggestedTipAmounts       []uint64       `json:"suggested_tip_amounts,omitempty"`
}

type InputLocationMessageContent struct {
	Latitude             float64 `json:"latitude"`
	Longitude            float64 `json:"longitude"`
	Heading              uint64  `json:"heading,omitempty"`
	HorizontalAccuracy   float64 `json:"horizontal_accuracy,omitempty"`
	LivePeriod           uint64  `json:"live_period,omitempty"`
	ProximityAlertRadius uint64  `json:"proximity_alert_radius,omitempty"`
}

type InputMedia struct {
	// Placeholder
}

type InputMediaAnimation struct {
	Media           string          `json:"media"`
	Type            string          `json:"type"`
	Caption         string          `json:"caption,omitempty"`
	CaptionEntities []MessageEntity `json:"caption_entities,omitempty"`
	Duration        uint64          `json:"duration,omitempty"`
	Height          uint64          `json:"height,omitempty"`
	ParseMode       string          `json:"parse_mode,omitempty"`
	Thumb           interface{}     `json:"thumb,omitempty"`
	Width           uint64          `json:"width,omitempty"`
}

type InputMediaAudio struct {
	Media           string          `json:"media"`
	Type            string          `json:"type"`
	Caption         string          `json:"caption,omitempty"`
	CaptionEntities []MessageEntity `json:"caption_entities,omitempty"`
	Duration        uint64          `json:"duration,omitempty"`
	ParseMode       string          `json:"parse_mode,omitempty"`
	Performer       string          `json:"performer,omitempty"`
	Thumb           interface{}     `json:"thumb,omitempty"`
	Title           string          `json:"title,omitempty"`
}

type InputMediaDocument struct {
	Media                       string          `json:"media"`
	Type                        string          `json:"type"`
	Caption                     string          `json:"caption,omitempty"`
	CaptionEntities             []MessageEntity `json:"caption_entities,omitempty"`
	DisableContentTypeDetection bool            `json:"disable_content_type_detection,omitempty"`
	ParseMode                   string          `json:"parse_mode,omitempty"`
	Thumb                       interface{}     `json:"thumb,omitempty"`
}

type InputMediaPhoto struct {
	Media           string          `json:"media"`
	Type            string          `json:"type"`
	Caption         string          `json:"caption,omitempty"`
	CaptionEntities []MessageEntity `json:"caption_entities,omitempty"`
	ParseMode       string          `json:"parse_mode,omitempty"`
}

type InputMediaVideo struct {
	Media             string          `json:"media"`
	Type              string          `json:"type"`
	Caption           string          `json:"caption,omitempty"`
	CaptionEntities   []MessageEntity `json:"caption_entities,omitempty"`
	Duration          uint64          `json:"duration,omitempty"`
	Height            uint64          `json:"height,omitempty"`
	ParseMode         string          `json:"parse_mode,omitempty"`
	SupportsStreaming bool            `json:"supports_streaming,omitempty"`
	Thumb             interface{}     `json:"thumb,omitempty"`
	Width             uint64          `json:"width,omitempty"`
}

type InputMessageContent struct {
	// Placeholder
}

type InputTextMessageContent struct {
	MessageText           string          `json:"message_text"`
	DisableWebPagePreview bool            `json:"disable_web_page_preview,omitempty"`
	Entities              []MessageEntity `json:"entities,omitempty"`
	ParseMode             string          `json:"parse_mode,omitempty"`
}

type InputVenueMessageContent struct {
	Address         string  `json:"address"`
	Latitude        float64 `json:"latitude"`
	Longitude       float64 `json:"longitude"`
	Title           string  `json:"title"`
	FoursquareId    string  `json:"foursquare_id,omitempty"`
	FoursquareType  string  `json:"foursquare_type,omitempty"`
	GooglePlaceId   string  `json:"google_place_id,omitempty"`
	GooglePlaceType string  `json:"google_place_type,omitempty"`
}

type Invoice struct {
	Currency       string `json:"currency"`
	Description    string `json:"description"`
	StartParameter string `json:"start_parameter"`
	Title          string `json:"title"`
	TotalAmount    uint64 `json:"total_amount"`
}

type KeyboardButton struct {
	Text            string                  `json:"text"`
	RequestContact  bool                    `json:"request_contact,omitempty"`
	RequestLocation bool                    `json:"request_location,omitempty"`
	RequestPoll     *KeyboardButtonPollType `json:"request_poll,omitempty"`
}

type KeyboardButtonPollType struct {
	Type string `json:"type,omitempty"`
}

type LabeledPrice struct {
	Amount uint64 `json:"amount"`
	Label  string `json:"label"`
}

type Location struct {
	Latitude             float64 `json:"latitude"`
	Longitude            float64 `json:"longitude"`
	Heading              uint64  `json:"heading,omitempty"`
	HorizontalAccuracy   float64 `json:"horizontal_accuracy,omitempty"`
	LivePeriod           uint64  `json:"live_period,omitempty"`
	ProximityAlertRadius uint64  `json:"proximity_alert_radius,omitempty"`
}

type LoginUrl struct {
	Url                string `json:"url"`
	BotUsername        string `json:"bot_username,omitempty"`
	ForwardText        string `json:"forward_text,omitempty"`
	RequestWriteAccess bool   `json:"request_write_access,omitempty"`
}

type MaskPosition struct {
	Point  string  `json:"point"`
	Scale  float64 `json:"scale"`
	XShift float64 `json:"x_shift"`
	YShift float64 `json:"y_shift"`
}

type Message struct {
	Chat                          Chat                           `json:"chat"`
	Date                          uint64                         `json:"date"`
	MessageId                     uint64                         `json:"message_id"`
	Animation                     *Animation                     `json:"animation,omitempty"`
	Audio                         *Audio                         `json:"audio,omitempty"`
	AuthorSignature               string                         `json:"author_signature,omitempty"`
	Caption                       string                         `json:"caption,omitempty"`
	CaptionEntities               []MessageEntity                `json:"caption_entities,omitempty"`
	ChannelChatCreated            bool                           `json:"channel_chat_created,omitempty"`
	ConnectedWebsite              string                         `json:"connected_website,omitempty"`
	Contact                       *Contact                       `json:"contact,omitempty"`
	DeleteChatPhoto               bool                           `json:"delete_chat_photo,omitempty"`
	Dice                          *Dice                          `json:"dice,omitempty"`
	Document                      *Document                      `json:"document,omitempty"`
	EditDate                      uint64                         `json:"edit_date,omitempty"`
	Entities                      []MessageEntity                `json:"entities,omitempty"`
	ForwardDate                   uint64                         `json:"forward_date,omitempty"`
	ForwardFrom                   *User                          `json:"forward_from,omitempty"`
	ForwardFromChat               *Chat                          `json:"forward_from_chat,omitempty"`
	ForwardFromMessageId          uint64                         `json:"forward_from_message_id,omitempty"`
	ForwardSenderName             string                         `json:"forward_sender_name,omitempty"`
	ForwardSignature              string                         `json:"forward_signature,omitempty"`
	From                          *User                          `json:"from,omitempty"`
	Game                          *Game                          `json:"game,omitempty"`
	GroupChatCreated              bool                           `json:"group_chat_created,omitempty"`
	Invoice                       *Invoice                       `json:"invoice,omitempty"`
	LeftChatMember                *User                          `json:"left_chat_member,omitempty"`
	Location                      *Location                      `json:"location,omitempty"`
	MediaGroupId                  string                         `json:"media_group_id,omitempty"`
	MessageAutoDeleteTimerChanged *MessageAutoDeleteTimerChanged `json:"message_auto_delete_timer_changed,omitempty"`
	MigrateFromChatId             uint64                         `json:"migrate_from_chat_id,omitempty"`
	MigrateToChatId               uint64                         `json:"migrate_to_chat_id,omitempty"`
	NewChatMembers                []User                         `json:"new_chat_members,omitempty"`
	NewChatPhoto                  []PhotoSize                    `json:"new_chat_photo,omitempty"`
	NewChatTitle                  string                         `json:"new_chat_title,omitempty"`
	PassportData                  *PassportData                  `json:"passport_data,omitempty"`
	Photo                         []PhotoSize                    `json:"photo,omitempty"`
	PinnedMessage                 *Message                       `json:"pinned_message,omitempty"`
	Poll                          *Poll                          `json:"poll,omitempty"`
	ProximityAlertTriggered       *ProximityAlertTriggered       `json:"proximity_alert_triggered,omitempty"`
	ReplyMarkup                   *InlineKeyboardMarkup          `json:"reply_markup,omitempty"`
	ReplyToMessage                *Message                       `json:"reply_to_message,omitempty"`
	SenderChat                    *Chat                          `json:"sender_chat,omitempty"`
	Sticker                       *Sticker                       `json:"sticker,omitempty"`
	SuccessfulPayment             *SuccessfulPayment             `json:"successful_payment,omitempty"`
	SupergroupChatCreated         bool                           `json:"supergroup_chat_created,omitempty"`
	Text                          string                         `json:"text,omitempty"`
	Venue                         *Venue                         `json:"venue,omitempty"`
	ViaBot                        *User                          `json:"via_bot,omitempty"`
	Video                         *Video                         `json:"video,omitempty"`
	VideoNote                     *VideoNote                     `json:"video_note,omitempty"`
	Voice                         *Voice                         `json:"voice,omitempty"`
	VoiceChatEnded                *VoiceChatEnded                `json:"voice_chat_ended,omitempty"`
	VoiceChatParticipantsInvited  *VoiceChatParticipantsInvited  `json:"voice_chat_participants_invited,omitempty"`
	VoiceChatScheduled            *VoiceChatScheduled            `json:"voice_chat_scheduled,omitempty"`
	VoiceChatStarted              *VoiceChatStarted              `json:"voice_chat_started,omitempty"`
}

type MessageAutoDeleteTimerChanged struct {
	MessageAutoDeleteTime uint64 `json:"message_auto_delete_time"`
}

type MessageEntity struct {
	Length   uint64 `json:"length"`
	Offset   uint64 `json:"offset"`
	Type     string `json:"type"`
	Language string `json:"language,omitempty"`
	Url      string `json:"url,omitempty"`
	User     *User  `json:"user,omitempty"`
}

type MessageId struct {
	MessageId uint64 `json:"message_id"`
}

type OrderInfo struct {
	Email           string           `json:"email,omitempty"`
	Name            string           `json:"name,omitempty"`
	PhoneNumber     string           `json:"phone_number,omitempty"`
	ShippingAddress *ShippingAddress `json:"shipping_address,omitempty"`
}

type PassportData struct {
	Credentials EncryptedCredentials       `json:"credentials"`
	Data        []EncryptedPassportElement `json:"data"`
}

type PassportElementError struct {
	// Placeholder
}

type PassportElementErrorDataField struct {
	DataHash  string `json:"data_hash"`
	FieldName string `json:"field_name"`
	Message   string `json:"message"`
	Source    string `json:"source"`
	Type      string `json:"type"`
}

type PassportElementErrorFile struct {
	FileHash string `json:"file_hash"`
	Message  string `json:"message"`
	Source   string `json:"source"`
	Type     string `json:"type"`
}

type PassportElementErrorFiles struct {
	FileHashes []string `json:"file_hashes"`
	Message    string   `json:"message"`
	Source     string   `json:"source"`
	Type       string   `json:"type"`
}

type PassportElementErrorFrontSide struct {
	FileHash string `json:"file_hash"`
	Message  string `json:"message"`
	Source   string `json:"source"`
	Type     string `json:"type"`
}

type PassportElementErrorReverseSide struct {
	FileHash string `json:"file_hash"`
	Message  string `json:"message"`
	Source   string `json:"source"`
	Type     string `json:"type"`
}

type PassportElementErrorSelfie struct {
	FileHash string `json:"file_hash"`
	Message  string `json:"message"`
	Source   string `json:"source"`
	Type     string `json:"type"`
}

type PassportElementErrorTranslationFile struct {
	FileHash string `json:"file_hash"`
	Message  string `json:"message"`
	Source   string `json:"source"`
	Type     string `json:"type"`
}

type PassportElementErrorTranslationFiles struct {
	FileHashes []string `json:"file_hashes"`
	Message    string   `json:"message"`
	Source     string   `json:"source"`
	Type       string   `json:"type"`
}

type PassportElementErrorUnspecified struct {
	ElementHash string `json:"element_hash"`
	Message     string `json:"message"`
	Source      string `json:"source"`
	Type        string `json:"type"`
}

type PassportFile struct {
	FileDate     uint64 `json:"file_date"`
	FileId       string `json:"file_id"`
	FileSize     uint64 `json:"file_size"`
	FileUniqueId string `json:"file_unique_id"`
}

type PhotoSize struct {
	FileId       string `json:"file_id"`
	FileUniqueId string `json:"file_unique_id"`
	Height       uint64 `json:"height"`
	Width        uint64 `json:"width"`
	FileSize     uint64 `json:"file_size,omitempty"`
}

type Poll struct {
	AllowsMultipleAnswers bool            `json:"allows_multiple_answers"`
	Id                    string          `json:"id"`
	IsAnonymous           bool            `json:"is_anonymous"`
	IsClosed              bool            `json:"is_closed"`
	Options               []PollOption    `json:"options"`
	Question              string          `json:"question"`
	TotalVoterCount       uint64          `json:"total_voter_count"`
	Type                  string          `json:"type"`
	CloseDate             uint64          `json:"close_date,omitempty"`
	CorrectOptionId       uint64          `json:"correct_option_id,omitempty"`
	Explanation           string          `json:"explanation,omitempty"`
	ExplanationEntities   []MessageEntity `json:"explanation_entities,omitempty"`
	OpenPeriod            uint64          `json:"open_period,omitempty"`
}

type PollAnswer struct {
	OptionIds []uint64 `json:"option_ids"`
	PollId    string   `json:"poll_id"`
	User      User     `json:"user"`
}

type PollOption struct {
	Text       string `json:"text"`
	VoterCount uint64 `json:"voter_count"`
}

type PreCheckoutQuery struct {
	Currency         string     `json:"currency"`
	From             User       `json:"from"`
	Id               string     `json:"id"`
	InvoicePayload   string     `json:"invoice_payload"`
	TotalAmount      uint64     `json:"total_amount"`
	OrderInfo        *OrderInfo `json:"order_info,omitempty"`
	ShippingOptionId string     `json:"shipping_option_id,omitempty"`
}

type ProximityAlertTriggered struct {
	Distance uint64 `json:"distance"`
	Traveler User   `json:"traveler"`
	Watcher  User   `json:"watcher"`
}

type ReplyKeyboardMarkup struct {
	Keyboard        [][]KeyboardButton `json:"keyboard"`
	OneTimeKeyboard bool               `json:"one_time_keyboard,omitempty"`
	ResizeKeyboard  bool               `json:"resize_keyboard,omitempty"`
	Selective       bool               `json:"selective,omitempty"`
}

type ReplyKeyboardRemove struct {
	RemoveKeyboard bool `json:"remove_keyboard"`
	Selective      bool `json:"selective,omitempty"`
}

type ResponseParameters struct {
	MigrateToChatId uint64 `json:"migrate_to_chat_id,omitempty"`
	RetryAfter      uint64 `json:"retry_after,omitempty"`
}

type ShippingAddress struct {
	City        string `json:"city"`
	CountryCode string `json:"country_code"`
	PostCode    string `json:"post_code"`
	State       string `json:"state"`
	StreetLine1 string `json:"street_line1"`
	StreetLine2 string `json:"street_line2"`
}

type ShippingOption struct {
	Id     string         `json:"id"`
	Prices []LabeledPrice `json:"prices"`
	Title  string         `json:"title"`
}

type ShippingQuery struct {
	From            User            `json:"from"`
	Id              string          `json:"id"`
	InvoicePayload  string          `json:"invoice_payload"`
	ShippingAddress ShippingAddress `json:"shipping_address"`
}

type Sticker struct {
	FileId       string        `json:"file_id"`
	FileUniqueId string        `json:"file_unique_id"`
	Height       uint64        `json:"height"`
	IsAnimated   bool          `json:"is_animated"`
	Width        uint64        `json:"width"`
	Emoji        string        `json:"emoji,omitempty"`
	FileSize     uint64        `json:"file_size,omitempty"`
	MaskPosition *MaskPosition `json:"mask_position,omitempty"`
	SetName      string        `json:"set_name,omitempty"`
	Thumb        *PhotoSize    `json:"thumb,omitempty"`
}

type StickerSet struct {
	ContainsMasks bool       `json:"contains_masks"`
	IsAnimated    bool       `json:"is_animated"`
	Name          string     `json:"name"`
	Stickers      []Sticker  `json:"stickers"`
	Title         string     `json:"title"`
	Thumb         *PhotoSize `json:"thumb,omitempty"`
}

type SuccessfulPayment struct {
	Currency                string     `json:"currency"`
	InvoicePayload          string     `json:"invoice_payload"`
	ProviderPaymentChargeId string     `json:"provider_payment_charge_id"`
	TelegramPaymentChargeId string     `json:"telegram_payment_charge_id"`
	TotalAmount             uint64     `json:"total_amount"`
	OrderInfo               *OrderInfo `json:"order_info,omitempty"`
	ShippingOptionId        string     `json:"shipping_option_id,omitempty"`
}

type Update struct {
	UpdateId           uint64              `json:"update_id"`
	CallbackQuery      *CallbackQuery      `json:"callback_query,omitempty"`
	ChannelPost        *Message            `json:"channel_post,omitempty"`
	ChatMember         *ChatMemberUpdated  `json:"chat_member,omitempty"`
	ChosenInlineResult *ChosenInlineResult `json:"chosen_inline_result,omitempty"`
	EditedChannelPost  *Message            `json:"edited_channel_post,omitempty"`
	EditedMessage      *Message            `json:"edited_message,omitempty"`
	InlineQuery        *InlineQuery        `json:"inline_query,omitempty"`
	Message            *Message            `json:"message,omitempty"`
	MyChatMember       *ChatMemberUpdated  `json:"my_chat_member,omitempty"`
	Poll               *Poll               `json:"poll,omitempty"`
	PollAnswer         *PollAnswer         `json:"poll_answer,omitempty"`
	PreCheckoutQuery   *PreCheckoutQuery   `json:"pre_checkout_query,omitempty"`
	ShippingQuery      *ShippingQuery      `json:"shipping_query,omitempty"`
}

type User struct {
	FirstName               string `json:"first_name"`
	Id                      uint64 `json:"id"`
	IsBot                   bool   `json:"is_bot"`
	CanJoinGroups           bool   `json:"can_join_groups,omitempty"`
	CanReadAllGroupMessages bool   `json:"can_read_all_group_messages,omitempty"`
	LanguageCode            string `json:"language_code,omitempty"`
	LastName                string `json:"last_name,omitempty"`
	SupportsInlineQueries   bool   `json:"supports_inline_queries,omitempty"`
	Username                string `json:"username,omitempty"`
}

type UserProfilePhotos struct {
	Photos     [][]PhotoSize `json:"photos"`
	TotalCount uint64        `json:"total_count"`
}

type Venue struct {
	Address         string   `json:"address"`
	Location        Location `json:"location"`
	Title           string   `json:"title"`
	FoursquareId    string   `json:"foursquare_id,omitempty"`
	FoursquareType  string   `json:"foursquare_type,omitempty"`
	GooglePlaceId   string   `json:"google_place_id,omitempty"`
	GooglePlaceType string   `json:"google_place_type,omitempty"`
}

type Video struct {
	Duration     uint64     `json:"duration"`
	FileId       string     `json:"file_id"`
	FileUniqueId string     `json:"file_unique_id"`
	Height       uint64     `json:"height"`
	Width        uint64     `json:"width"`
	FileName     string     `json:"file_name,omitempty"`
	FileSize     uint64     `json:"file_size,omitempty"`
	MimeType     string     `json:"mime_type,omitempty"`
	Thumb        *PhotoSize `json:"thumb,omitempty"`
}

type VideoNote struct {
	Duration     uint64     `json:"duration"`
	FileId       string     `json:"file_id"`
	FileUniqueId string     `json:"file_unique_id"`
	Length       uint64     `json:"length"`
	FileSize     uint64     `json:"file_size,omitempty"`
	Thumb        *PhotoSize `json:"thumb,omitempty"`
}

type Voice struct {
	Duration     uint64 `json:"duration"`
	FileId       string `json:"file_id"`
	FileUniqueId string `json:"file_unique_id"`
	FileSize     uint64 `json:"file_size,omitempty"`
	MimeType     string `json:"mime_type,omitempty"`
}

type VoiceChatEnded struct {
	Duration uint64 `json:"duration"`
}

type VoiceChatParticipantsInvited struct {
	Users []User `json:"users"`
}

type VoiceChatScheduled struct {
	StartDate uint64 `json:"start_date"`
}

type VoiceChatStarted struct {
	// Placeholder
}

type WebhookInfo struct {
	HasCustomCertificate bool     `json:"has_custom_certificate"`
	PendingUpdateCount   uint64   `json:"pending_update_count"`
	Url                  string   `json:"url"`
	AllowedUpdates       []string `json:"allowed_updates,omitempty"`
	IpAddress            string   `json:"ip_address,omitempty"`
	LastErrorDate        uint64   `json:"last_error_date,omitempty"`
	LastErrorMessage     string   `json:"last_error_message,omitempty"`
	MaxConnections       uint64   `json:"max_connections,omitempty"`
}
