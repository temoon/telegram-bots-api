package telegram

type SuccessfulPayment struct {
	Currency                string     `json:"currency"`
	InvoicePayload          string     `json:"invoice_payload"`
	OrderInfo               *OrderInfo `json:"order_info"`
	ProviderPaymentChargeId string     `json:"provider_payment_charge_id"`
	ShippingOptionId        string     `json:"shipping_option_id"`
	TelegramPaymentChargeId string     `json:"telegram_payment_charge_id"`
	TotalAmount             uint64     `json:"total_amount"`
}

type PassportFile struct {
	FileDate     uint64 `json:"file_date"`
	FileId       string `json:"file_id"`
	FileSize     uint64 `json:"file_size"`
	FileUniqueId string `json:"file_unique_id"`
}

type ReplyKeyboardMarkup struct {
	Keyboard        [][]KeyboardButton `json:"keyboard"`
	OneTimeKeyboard bool               `json:"one_time_keyboard"`
	ResizeKeyboard  bool               `json:"resize_keyboard"`
	Selective       bool               `json:"selective"`
}

type PassportElementErrorDataField struct {
	DataHash  string `json:"data_hash"`
	FieldName string `json:"field_name"`
	Message   string `json:"message"`
	Source    string `json:"source"`
	Type      string `json:"type"`
}

type File struct {
	FileId       string `json:"file_id"`
	FilePath     string `json:"file_path"`
	FileSize     uint64 `json:"file_size"`
	FileUniqueId string `json:"file_unique_id"`
}

type MaskPosition struct {
	Point  string  `json:"point"`
	Scale  float64 `json:"scale"`
	XShift float64 `json:"x_shift"`
	YShift float64 `json:"y_shift"`
}

type VideoNote struct {
	Duration     uint64     `json:"duration"`
	FileId       string     `json:"file_id"`
	FileSize     uint64     `json:"file_size"`
	FileUniqueId string     `json:"file_unique_id"`
	Length       uint64     `json:"length"`
	Thumb        *PhotoSize `json:"thumb"`
}

type InlineQueryResultCachedAudio struct {
	AudioFileId         string                `json:"audio_file_id"`
	Caption             string                `json:"caption"`
	CaptionEntities     []MessageEntity       `json:"caption_entities"`
	Id                  string                `json:"id"`
	InputMessageContent *InputMessageContent  `json:"input_message_content"`
	ParseMode           string                `json:"parse_mode"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup"`
	Type                string                `json:"type"`
}

type InputTextMessageContent struct {
	DisableWebPagePreview bool            `json:"disable_web_page_preview"`
	Entities              []MessageEntity `json:"entities"`
	MessageText           string          `json:"message_text"`
	ParseMode             string          `json:"parse_mode"`
}

type InputMedia struct {
	// Placeholder
}

type MessageId struct {
	MessageId uint64 `json:"message_id"`
}

type VoiceChatEnded struct {
	Duration uint64 `json:"duration"`
}

type VoiceChatParticipantsInvited struct {
	Users []User `json:"users"`
}

type InlineQuery struct {
	From     User      `json:"from"`
	Id       string    `json:"id"`
	Location *Location `json:"location"`
	Offset   string    `json:"offset"`
	Query    string    `json:"query"`
}

type InlineQueryResultCachedVideo struct {
	Caption             string                `json:"caption"`
	CaptionEntities     []MessageEntity       `json:"caption_entities"`
	Description         string                `json:"description"`
	Id                  string                `json:"id"`
	InputMessageContent *InputMessageContent  `json:"input_message_content"`
	ParseMode           string                `json:"parse_mode"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup"`
	Title               string                `json:"title"`
	Type                string                `json:"type"`
	VideoFileId         string                `json:"video_file_id"`
}

type InlineQueryResultVenue struct {
	Address             string                `json:"address"`
	FoursquareId        string                `json:"foursquare_id"`
	FoursquareType      string                `json:"foursquare_type"`
	GooglePlaceId       string                `json:"google_place_id"`
	GooglePlaceType     string                `json:"google_place_type"`
	Id                  string                `json:"id"`
	InputMessageContent *InputMessageContent  `json:"input_message_content"`
	Latitude            float64               `json:"latitude"`
	Longitude           float64               `json:"longitude"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup"`
	ThumbHeight         uint64                `json:"thumb_height"`
	ThumbUrl            string                `json:"thumb_url"`
	ThumbWidth          uint64                `json:"thumb_width"`
	Title               string                `json:"title"`
	Type                string                `json:"type"`
}

type InputLocationMessageContent struct {
	Heading              uint64  `json:"heading"`
	HorizontalAccuracy   float64 `json:"horizontal_accuracy"`
	Latitude             float64 `json:"latitude"`
	LivePeriod           uint64  `json:"live_period"`
	Longitude            float64 `json:"longitude"`
	ProximityAlertRadius uint64  `json:"proximity_alert_radius"`
}

type PassportElementError struct {
	// Placeholder
}

type PreCheckoutQuery struct {
	Currency         string     `json:"currency"`
	From             User       `json:"from"`
	Id               string     `json:"id"`
	InvoicePayload   string     `json:"invoice_payload"`
	OrderInfo        *OrderInfo `json:"order_info"`
	ShippingOptionId string     `json:"shipping_option_id"`
	TotalAmount      uint64     `json:"total_amount"`
}

type Sticker struct {
	Emoji        string        `json:"emoji"`
	FileId       string        `json:"file_id"`
	FileSize     uint64        `json:"file_size"`
	FileUniqueId string        `json:"file_unique_id"`
	Height       uint64        `json:"height"`
	IsAnimated   bool          `json:"is_animated"`
	MaskPosition *MaskPosition `json:"mask_position"`
	SetName      string        `json:"set_name"`
	Thumb        *PhotoSize    `json:"thumb"`
	Width        uint64        `json:"width"`
}

type Venue struct {
	Address         string   `json:"address"`
	FoursquareId    string   `json:"foursquare_id"`
	FoursquareType  string   `json:"foursquare_type"`
	GooglePlaceId   string   `json:"google_place_id"`
	GooglePlaceType string   `json:"google_place_type"`
	Location        Location `json:"location"`
	Title           string   `json:"title"`
}

type ChosenInlineResult struct {
	From            User      `json:"from"`
	InlineMessageId string    `json:"inline_message_id"`
	Location        *Location `json:"location"`
	Query           string    `json:"query"`
	ResultId        string    `json:"result_id"`
}

type InlineKeyboardMarkup struct {
	InlineKeyboard [][]InlineKeyboardButton `json:"inline_keyboard"`
}

type EncryptedPassportElement struct {
	Data        string         `json:"data"`
	Email       string         `json:"email"`
	Files       []PassportFile `json:"files"`
	FrontSide   *PassportFile  `json:"front_side"`
	Hash        string         `json:"hash"`
	PhoneNumber string         `json:"phone_number"`
	ReverseSide *PassportFile  `json:"reverse_side"`
	Selfie      *PassportFile  `json:"selfie"`
	Translation []PassportFile `json:"translation"`
	Type        string         `json:"type"`
}

type CallbackQuery struct {
	ChatInstance    string   `json:"chat_instance"`
	Data            string   `json:"data"`
	From            User     `json:"from"`
	GameShortName   string   `json:"game_short_name"`
	Id              string   `json:"id"`
	InlineMessageId string   `json:"inline_message_id"`
	Message         *Message `json:"message"`
}

type InlineQueryResultVoice struct {
	Caption             string                `json:"caption"`
	CaptionEntities     []MessageEntity       `json:"caption_entities"`
	Id                  string                `json:"id"`
	InputMessageContent *InputMessageContent  `json:"input_message_content"`
	ParseMode           string                `json:"parse_mode"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup"`
	Title               string                `json:"title"`
	Type                string                `json:"type"`
	VoiceDuration       uint64                `json:"voice_duration"`
	VoiceUrl            string                `json:"voice_url"`
}

type ReplyKeyboardRemove struct {
	RemoveKeyboard bool `json:"remove_keyboard"`
	Selective      bool `json:"selective"`
}

type VoiceChatStarted struct {
	// Placeholder
}

type InlineQueryResultGame struct {
	GameShortName string                `json:"game_short_name"`
	Id            string                `json:"id"`
	ReplyMarkup   *InlineKeyboardMarkup `json:"reply_markup"`
	Type          string                `json:"type"`
}

type InputMediaAnimation struct {
	Caption         string          `json:"caption"`
	CaptionEntities []MessageEntity `json:"caption_entities"`
	Duration        uint64          `json:"duration"`
	Height          uint64          `json:"height"`
	Media           string          `json:"media"`
	ParseMode       string          `json:"parse_mode"`
	Thumb           interface{}     `json:"thumb"`
	Type            string          `json:"type"`
	Width           uint64          `json:"width"`
}

type MessageEntity struct {
	Language string `json:"language"`
	Length   uint64 `json:"length"`
	Offset   uint64 `json:"offset"`
	Type     string `json:"type"`
	Url      string `json:"url"`
	User     *User  `json:"user"`
}

type ChatPermissions struct {
	CanAddWebPagePreviews bool `json:"can_add_web_page_previews"`
	CanChangeInfo         bool `json:"can_change_info"`
	CanInviteUsers        bool `json:"can_invite_users"`
	CanPinMessages        bool `json:"can_pin_messages"`
	CanSendMediaMessages  bool `json:"can_send_media_messages"`
	CanSendMessages       bool `json:"can_send_messages"`
	CanSendOtherMessages  bool `json:"can_send_other_messages"`
	CanSendPolls          bool `json:"can_send_polls"`
}

type InlineQueryResultCachedMpeg4Gif struct {
	Caption             string                `json:"caption"`
	CaptionEntities     []MessageEntity       `json:"caption_entities"`
	Id                  string                `json:"id"`
	InputMessageContent *InputMessageContent  `json:"input_message_content"`
	Mpeg4FileId         string                `json:"mpeg4_file_id"`
	ParseMode           string                `json:"parse_mode"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup"`
	Title               string                `json:"title"`
	Type                string                `json:"type"`
}

type Dice struct {
	Emoji string `json:"emoji"`
	Value uint64 `json:"value"`
}

type GameHighScore struct {
	Position uint64 `json:"position"`
	Score    uint64 `json:"score"`
	User     User   `json:"user"`
}

type InlineQueryResultLocation struct {
	Heading              uint64                `json:"heading"`
	HorizontalAccuracy   float64               `json:"horizontal_accuracy"`
	Id                   string                `json:"id"`
	InputMessageContent  *InputMessageContent  `json:"input_message_content"`
	Latitude             float64               `json:"latitude"`
	LivePeriod           uint64                `json:"live_period"`
	Longitude            float64               `json:"longitude"`
	ProximityAlertRadius uint64                `json:"proximity_alert_radius"`
	ReplyMarkup          *InlineKeyboardMarkup `json:"reply_markup"`
	ThumbHeight          uint64                `json:"thumb_height"`
	ThumbUrl             string                `json:"thumb_url"`
	ThumbWidth           uint64                `json:"thumb_width"`
	Title                string                `json:"title"`
	Type                 string                `json:"type"`
}

type PassportElementErrorTranslationFiles struct {
	FileHashes []string `json:"file_hashes"`
	Message    string   `json:"message"`
	Source     string   `json:"source"`
	Type       string   `json:"type"`
}

type ProximityAlertTriggered struct {
	Distance uint64 `json:"distance"`
	Traveler User   `json:"traveler"`
	Watcher  User   `json:"watcher"`
}

type PassportElementErrorFile struct {
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

type InlineQueryResultCachedSticker struct {
	Id                  string                `json:"id"`
	InputMessageContent *InputMessageContent  `json:"input_message_content"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup"`
	StickerFileId       string                `json:"sticker_file_id"`
	Type                string                `json:"type"`
}

type InlineQueryResultVideo struct {
	Caption             string                `json:"caption"`
	CaptionEntities     []MessageEntity       `json:"caption_entities"`
	Description         string                `json:"description"`
	Id                  string                `json:"id"`
	InputMessageContent *InputMessageContent  `json:"input_message_content"`
	MimeType            string                `json:"mime_type"`
	ParseMode           string                `json:"parse_mode"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup"`
	ThumbUrl            string                `json:"thumb_url"`
	Title               string                `json:"title"`
	Type                string                `json:"type"`
	VideoDuration       uint64                `json:"video_duration"`
	VideoHeight         uint64                `json:"video_height"`
	VideoUrl            string                `json:"video_url"`
	VideoWidth          uint64                `json:"video_width"`
}

type InputContactMessageContent struct {
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	PhoneNumber string `json:"phone_number"`
	Vcard       string `json:"vcard"`
}

type Message struct {
	Animation                     *Animation                     `json:"animation"`
	Audio                         *Audio                         `json:"audio"`
	AuthorSignature               string                         `json:"author_signature"`
	Caption                       string                         `json:"caption"`
	CaptionEntities               []MessageEntity                `json:"caption_entities"`
	ChannelChatCreated            bool                           `json:"channel_chat_created"`
	Chat                          Chat                           `json:"chat"`
	ConnectedWebsite              string                         `json:"connected_website"`
	Contact                       *Contact                       `json:"contact"`
	Date                          uint64                         `json:"date"`
	DeleteChatPhoto               bool                           `json:"delete_chat_photo"`
	Dice                          *Dice                          `json:"dice"`
	Document                      *Document                      `json:"document"`
	EditDate                      uint64                         `json:"edit_date"`
	Entities                      []MessageEntity                `json:"entities"`
	ForwardDate                   uint64                         `json:"forward_date"`
	ForwardFrom                   *User                          `json:"forward_from"`
	ForwardFromChat               *Chat                          `json:"forward_from_chat"`
	ForwardFromMessageId          uint64                         `json:"forward_from_message_id"`
	ForwardSenderName             string                         `json:"forward_sender_name"`
	ForwardSignature              string                         `json:"forward_signature"`
	From                          *User                          `json:"from"`
	Game                          *Game                          `json:"game"`
	GroupChatCreated              bool                           `json:"group_chat_created"`
	Invoice                       *Invoice                       `json:"invoice"`
	LeftChatMember                *User                          `json:"left_chat_member"`
	Location                      *Location                      `json:"location"`
	MediaGroupId                  string                         `json:"media_group_id"`
	MessageAutoDeleteTimerChanged *MessageAutoDeleteTimerChanged `json:"message_auto_delete_timer_changed"`
	MessageId                     uint64                         `json:"message_id"`
	MigrateFromChatId             uint64                         `json:"migrate_from_chat_id"`
	MigrateToChatId               uint64                         `json:"migrate_to_chat_id"`
	NewChatMembers                []User                         `json:"new_chat_members"`
	NewChatPhoto                  []PhotoSize                    `json:"new_chat_photo"`
	NewChatTitle                  string                         `json:"new_chat_title"`
	PassportData                  *PassportData                  `json:"passport_data"`
	Photo                         []PhotoSize                    `json:"photo"`
	PinnedMessage                 *Message                       `json:"pinned_message"`
	Poll                          *Poll                          `json:"poll"`
	ProximityAlertTriggered       *ProximityAlertTriggered       `json:"proximity_alert_triggered"`
	ReplyMarkup                   *InlineKeyboardMarkup          `json:"reply_markup"`
	ReplyToMessage                *Message                       `json:"reply_to_message"`
	SenderChat                    *Chat                          `json:"sender_chat"`
	Sticker                       *Sticker                       `json:"sticker"`
	SuccessfulPayment             *SuccessfulPayment             `json:"successful_payment"`
	SupergroupChatCreated         bool                           `json:"supergroup_chat_created"`
	Text                          string                         `json:"text"`
	Venue                         *Venue                         `json:"venue"`
	ViaBot                        *User                          `json:"via_bot"`
	Video                         *Video                         `json:"video"`
	VideoNote                     *VideoNote                     `json:"video_note"`
	Voice                         *Voice                         `json:"voice"`
	VoiceChatEnded                *VoiceChatEnded                `json:"voice_chat_ended"`
	VoiceChatParticipantsInvited  *VoiceChatParticipantsInvited  `json:"voice_chat_participants_invited"`
	VoiceChatStarted              *VoiceChatStarted              `json:"voice_chat_started"`
}

type ResponseParameters struct {
	MigrateToChatId uint64 `json:"migrate_to_chat_id"`
	RetryAfter      uint64 `json:"retry_after"`
}

type InlineQueryResultPhoto struct {
	Caption             string                `json:"caption"`
	CaptionEntities     []MessageEntity       `json:"caption_entities"`
	Description         string                `json:"description"`
	Id                  string                `json:"id"`
	InputMessageContent *InputMessageContent  `json:"input_message_content"`
	ParseMode           string                `json:"parse_mode"`
	PhotoHeight         uint64                `json:"photo_height"`
	PhotoUrl            string                `json:"photo_url"`
	PhotoWidth          uint64                `json:"photo_width"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup"`
	ThumbUrl            string                `json:"thumb_url"`
	Title               string                `json:"title"`
	Type                string                `json:"type"`
}

type PollOption struct {
	Text       string `json:"text"`
	VoterCount uint64 `json:"voter_count"`
}

type ShippingOption struct {
	Id     string         `json:"id"`
	Prices []LabeledPrice `json:"prices"`
	Title  string         `json:"title"`
}

type Error struct {
	Description string              `json:"description"`
	ErrorCode   uint64              `json:"error_code"`
	Ok          bool                `json:"ok"`
	Parameters  *ResponseParameters `json:"parameters"`
}

type UserProfilePhotos struct {
	Photos     [][]PhotoSize `json:"photos"`
	TotalCount uint64        `json:"total_count"`
}

type Video struct {
	Duration     uint64     `json:"duration"`
	FileId       string     `json:"file_id"`
	FileName     string     `json:"file_name"`
	FileSize     uint64     `json:"file_size"`
	FileUniqueId string     `json:"file_unique_id"`
	Height       uint64     `json:"height"`
	MimeType     string     `json:"mime_type"`
	Thumb        *PhotoSize `json:"thumb"`
	Width        uint64     `json:"width"`
}

type OrderInfo struct {
	Email           string           `json:"email"`
	Name            string           `json:"name"`
	PhoneNumber     string           `json:"phone_number"`
	ShippingAddress *ShippingAddress `json:"shipping_address"`
}

type PassportData struct {
	Credentials EncryptedCredentials       `json:"credentials"`
	Data        []EncryptedPassportElement `json:"data"`
}

type PassportElementErrorSelfie struct {
	FileHash string `json:"file_hash"`
	Message  string `json:"message"`
	Source   string `json:"source"`
	Type     string `json:"type"`
}

type StickerSet struct {
	ContainsMasks bool       `json:"contains_masks"`
	IsAnimated    bool       `json:"is_animated"`
	Name          string     `json:"name"`
	Stickers      []Sticker  `json:"stickers"`
	Thumb         *PhotoSize `json:"thumb"`
	Title         string     `json:"title"`
}

type InputMediaDocument struct {
	Caption                     string          `json:"caption"`
	CaptionEntities             []MessageEntity `json:"caption_entities"`
	DisableContentTypeDetection bool            `json:"disable_content_type_detection"`
	Media                       string          `json:"media"`
	ParseMode                   string          `json:"parse_mode"`
	Thumb                       interface{}     `json:"thumb"`
	Type                        string          `json:"type"`
}

type InlineQueryResultContact struct {
	FirstName           string                `json:"first_name"`
	Id                  string                `json:"id"`
	InputMessageContent *InputMessageContent  `json:"input_message_content"`
	LastName            string                `json:"last_name"`
	PhoneNumber         string                `json:"phone_number"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup"`
	ThumbHeight         uint64                `json:"thumb_height"`
	ThumbUrl            string                `json:"thumb_url"`
	ThumbWidth          uint64                `json:"thumb_width"`
	Type                string                `json:"type"`
	Vcard               string                `json:"vcard"`
}

type Invoice struct {
	Currency       string `json:"currency"`
	Description    string `json:"description"`
	StartParameter string `json:"start_parameter"`
	Title          string `json:"title"`
	TotalAmount    uint64 `json:"total_amount"`
}

type ChatMemberUpdated struct {
	Chat          Chat            `json:"chat"`
	Date          uint64          `json:"date"`
	From          User            `json:"from"`
	InviteLink    *ChatInviteLink `json:"invite_link"`
	NewChatMember ChatMember      `json:"new_chat_member"`
	OldChatMember ChatMember      `json:"old_chat_member"`
}

type InlineQueryResultArticle struct {
	Description         string                `json:"description"`
	HideUrl             bool                  `json:"hide_url"`
	Id                  string                `json:"id"`
	InputMessageContent InputMessageContent   `json:"input_message_content"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup"`
	ThumbHeight         uint64                `json:"thumb_height"`
	ThumbUrl            string                `json:"thumb_url"`
	ThumbWidth          uint64                `json:"thumb_width"`
	Title               string                `json:"title"`
	Type                string                `json:"type"`
	Url                 string                `json:"url"`
}

type Voice struct {
	Duration     uint64 `json:"duration"`
	FileId       string `json:"file_id"`
	FileSize     uint64 `json:"file_size"`
	FileUniqueId string `json:"file_unique_id"`
	MimeType     string `json:"mime_type"`
}

type Game struct {
	Animation    *Animation      `json:"animation"`
	Description  string          `json:"description"`
	Photo        []PhotoSize     `json:"photo"`
	Text         string          `json:"text"`
	TextEntities []MessageEntity `json:"text_entities"`
	Title        string          `json:"title"`
}

type PollAnswer struct {
	OptionIds []uint64 `json:"option_ids"`
	PollId    string   `json:"poll_id"`
	User      User     `json:"user"`
}

type LoginUrl struct {
	BotUsername        string `json:"bot_username"`
	ForwardText        string `json:"forward_text"`
	RequestWriteAccess bool   `json:"request_write_access"`
	Url                string `json:"url"`
}

type InputMessageContent struct {
	// Placeholder
}

type InlineQueryResultCachedGif struct {
	Caption             string                `json:"caption"`
	CaptionEntities     []MessageEntity       `json:"caption_entities"`
	GifFileId           string                `json:"gif_file_id"`
	Id                  string                `json:"id"`
	InputMessageContent *InputMessageContent  `json:"input_message_content"`
	ParseMode           string                `json:"parse_mode"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup"`
	Title               string                `json:"title"`
	Type                string                `json:"type"`
}

type BotCommand struct {
	Command     string `json:"command"`
	Description string `json:"description"`
}

type Contact struct {
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	PhoneNumber string `json:"phone_number"`
	UserId      uint64 `json:"user_id"`
	Vcard       string `json:"vcard"`
}

type PhotoSize struct {
	FileId       string `json:"file_id"`
	FileSize     uint64 `json:"file_size"`
	FileUniqueId string `json:"file_unique_id"`
	Height       uint64 `json:"height"`
	Width        uint64 `json:"width"`
}

type LabeledPrice struct {
	Amount uint64 `json:"amount"`
	Label  string `json:"label"`
}

type PassportElementErrorReverseSide struct {
	FileHash string `json:"file_hash"`
	Message  string `json:"message"`
	Source   string `json:"source"`
	Type     string `json:"type"`
}

type InlineQueryResultMpeg4Gif struct {
	Caption             string                `json:"caption"`
	CaptionEntities     []MessageEntity       `json:"caption_entities"`
	Id                  string                `json:"id"`
	InputMessageContent *InputMessageContent  `json:"input_message_content"`
	Mpeg4Duration       uint64                `json:"mpeg4_duration"`
	Mpeg4Height         uint64                `json:"mpeg4_height"`
	Mpeg4Url            string                `json:"mpeg4_url"`
	Mpeg4Width          uint64                `json:"mpeg4_width"`
	ParseMode           string                `json:"parse_mode"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup"`
	ThumbMimeType       string                `json:"thumb_mime_type"`
	ThumbUrl            string                `json:"thumb_url"`
	Title               string                `json:"title"`
	Type                string                `json:"type"`
}

type KeyboardButton struct {
	RequestContact  bool                    `json:"request_contact"`
	RequestLocation bool                    `json:"request_location"`
	RequestPoll     *KeyboardButtonPollType `json:"request_poll"`
	Text            string                  `json:"text"`
}

type Audio struct {
	Duration     uint64     `json:"duration"`
	FileId       string     `json:"file_id"`
	FileName     string     `json:"file_name"`
	FileSize     uint64     `json:"file_size"`
	FileUniqueId string     `json:"file_unique_id"`
	MimeType     string     `json:"mime_type"`
	Performer    string     `json:"performer"`
	Thumb        *PhotoSize `json:"thumb"`
	Title        string     `json:"title"`
}

type Poll struct {
	AllowsMultipleAnswers bool            `json:"allows_multiple_answers"`
	CloseDate             uint64          `json:"close_date"`
	CorrectOptionId       uint64          `json:"correct_option_id"`
	Explanation           string          `json:"explanation"`
	ExplanationEntities   []MessageEntity `json:"explanation_entities"`
	Id                    string          `json:"id"`
	IsAnonymous           bool            `json:"is_anonymous"`
	IsClosed              bool            `json:"is_closed"`
	OpenPeriod            uint64          `json:"open_period"`
	Options               []PollOption    `json:"options"`
	Question              string          `json:"question"`
	TotalVoterCount       uint64          `json:"total_voter_count"`
	Type                  string          `json:"type"`
}

type InputMediaVideo struct {
	Caption           string          `json:"caption"`
	CaptionEntities   []MessageEntity `json:"caption_entities"`
	Duration          uint64          `json:"duration"`
	Height            uint64          `json:"height"`
	Media             string          `json:"media"`
	ParseMode         string          `json:"parse_mode"`
	SupportsStreaming bool            `json:"supports_streaming"`
	Thumb             interface{}     `json:"thumb"`
	Type              string          `json:"type"`
	Width             uint64          `json:"width"`
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

type PassportElementErrorUnspecified struct {
	ElementHash string `json:"element_hash"`
	Message     string `json:"message"`
	Source      string `json:"source"`
	Type        string `json:"type"`
}

type CallbackGame struct {
	// Placeholder
}

type InlineQueryResultCachedVoice struct {
	Caption             string                `json:"caption"`
	CaptionEntities     []MessageEntity       `json:"caption_entities"`
	Id                  string                `json:"id"`
	InputMessageContent *InputMessageContent  `json:"input_message_content"`
	ParseMode           string                `json:"parse_mode"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup"`
	Title               string                `json:"title"`
	Type                string                `json:"type"`
	VoiceFileId         string                `json:"voice_file_id"`
}

type ShippingQuery struct {
	From            User            `json:"from"`
	Id              string          `json:"id"`
	InvoicePayload  string          `json:"invoice_payload"`
	ShippingAddress ShippingAddress `json:"shipping_address"`
}

type WebhookInfo struct {
	AllowedUpdates       []string `json:"allowed_updates"`
	HasCustomCertificate bool     `json:"has_custom_certificate"`
	IpAddress            string   `json:"ip_address"`
	LastErrorDate        uint64   `json:"last_error_date"`
	LastErrorMessage     string   `json:"last_error_message"`
	MaxConnections       uint64   `json:"max_connections"`
	PendingUpdateCount   uint64   `json:"pending_update_count"`
	Url                  string   `json:"url"`
}

type InlineQueryResult struct {
	// Placeholder
}

type InlineQueryResultDocument struct {
	Caption             string                `json:"caption"`
	CaptionEntities     []MessageEntity       `json:"caption_entities"`
	Description         string                `json:"description"`
	DocumentUrl         string                `json:"document_url"`
	Id                  string                `json:"id"`
	InputMessageContent *InputMessageContent  `json:"input_message_content"`
	MimeType            string                `json:"mime_type"`
	ParseMode           string                `json:"parse_mode"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup"`
	ThumbHeight         uint64                `json:"thumb_height"`
	ThumbUrl            string                `json:"thumb_url"`
	ThumbWidth          uint64                `json:"thumb_width"`
	Title               string                `json:"title"`
	Type                string                `json:"type"`
}

type Update struct {
	CallbackQuery      *CallbackQuery      `json:"callback_query"`
	ChannelPost        *Message            `json:"channel_post"`
	ChatMember         *ChatMemberUpdated  `json:"chat_member"`
	ChosenInlineResult *ChosenInlineResult `json:"chosen_inline_result"`
	EditedChannelPost  *Message            `json:"edited_channel_post"`
	EditedMessage      *Message            `json:"edited_message"`
	InlineQuery        *InlineQuery        `json:"inline_query"`
	Message            *Message            `json:"message"`
	MyChatMember       *ChatMemberUpdated  `json:"my_chat_member"`
	Poll               *Poll               `json:"poll"`
	PollAnswer         *PollAnswer         `json:"poll_answer"`
	PreCheckoutQuery   *PreCheckoutQuery   `json:"pre_checkout_query"`
	ShippingQuery      *ShippingQuery      `json:"shipping_query"`
	UpdateId           uint64              `json:"update_id"`
}

type Animation struct {
	Duration     uint64     `json:"duration"`
	FileId       string     `json:"file_id"`
	FileName     string     `json:"file_name"`
	FileSize     uint64     `json:"file_size"`
	FileUniqueId string     `json:"file_unique_id"`
	Height       uint64     `json:"height"`
	MimeType     string     `json:"mime_type"`
	Thumb        *PhotoSize `json:"thumb"`
	Width        uint64     `json:"width"`
}

type InlineQueryResultGif struct {
	Caption             string                `json:"caption"`
	CaptionEntities     []MessageEntity       `json:"caption_entities"`
	GifDuration         uint64                `json:"gif_duration"`
	GifHeight           uint64                `json:"gif_height"`
	GifUrl              string                `json:"gif_url"`
	GifWidth            uint64                `json:"gif_width"`
	Id                  string                `json:"id"`
	InputMessageContent *InputMessageContent  `json:"input_message_content"`
	ParseMode           string                `json:"parse_mode"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup"`
	ThumbMimeType       string                `json:"thumb_mime_type"`
	ThumbUrl            string                `json:"thumb_url"`
	Title               string                `json:"title"`
	Type                string                `json:"type"`
}

type InlineQueryResultAudio struct {
	AudioDuration       uint64                `json:"audio_duration"`
	AudioUrl            string                `json:"audio_url"`
	Caption             string                `json:"caption"`
	CaptionEntities     []MessageEntity       `json:"caption_entities"`
	Id                  string                `json:"id"`
	InputMessageContent *InputMessageContent  `json:"input_message_content"`
	ParseMode           string                `json:"parse_mode"`
	Performer           string                `json:"performer"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup"`
	Title               string                `json:"title"`
	Type                string                `json:"type"`
}

type Location struct {
	Heading              uint64  `json:"heading"`
	HorizontalAccuracy   float64 `json:"horizontal_accuracy"`
	Latitude             float64 `json:"latitude"`
	LivePeriod           uint64  `json:"live_period"`
	Longitude            float64 `json:"longitude"`
	ProximityAlertRadius uint64  `json:"proximity_alert_radius"`
}

type ChatInviteLink struct {
	Creator     User   `json:"creator"`
	ExpireDate  uint64 `json:"expire_date"`
	InviteLink  string `json:"invite_link"`
	IsPrimary   bool   `json:"is_primary"`
	IsRevoked   bool   `json:"is_revoked"`
	MemberLimit uint64 `json:"member_limit"`
}

type ChatPhoto struct {
	BigFileId         string `json:"big_file_id"`
	BigFileUniqueId   string `json:"big_file_unique_id"`
	SmallFileId       string `json:"small_file_id"`
	SmallFileUniqueId string `json:"small_file_unique_id"`
}

type InlineQueryResultCachedPhoto struct {
	Caption             string                `json:"caption"`
	CaptionEntities     []MessageEntity       `json:"caption_entities"`
	Description         string                `json:"description"`
	Id                  string                `json:"id"`
	InputMessageContent *InputMessageContent  `json:"input_message_content"`
	ParseMode           string                `json:"parse_mode"`
	PhotoFileId         string                `json:"photo_file_id"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup"`
	Title               string                `json:"title"`
	Type                string                `json:"type"`
}

type InputMediaAudio struct {
	Caption         string          `json:"caption"`
	CaptionEntities []MessageEntity `json:"caption_entities"`
	Duration        uint64          `json:"duration"`
	Media           string          `json:"media"`
	ParseMode       string          `json:"parse_mode"`
	Performer       string          `json:"performer"`
	Thumb           interface{}     `json:"thumb"`
	Title           string          `json:"title"`
	Type            string          `json:"type"`
}

type ForceReply struct {
	ForceReply bool `json:"force_reply"`
	Selective  bool `json:"selective"`
}

type InlineKeyboardButton struct {
	CallbackData                 string        `json:"callback_data"`
	CallbackGame                 *CallbackGame `json:"callback_game"`
	LoginUrl                     *LoginUrl     `json:"login_url"`
	Pay                          bool          `json:"pay"`
	SwitchInlineQuery            string        `json:"switch_inline_query"`
	SwitchInlineQueryCurrentChat string        `json:"switch_inline_query_current_chat"`
	Text                         string        `json:"text"`
	Url                          string        `json:"url"`
}

type User struct {
	CanJoinGroups           bool   `json:"can_join_groups"`
	CanReadAllGroupMessages bool   `json:"can_read_all_group_messages"`
	FirstName               string `json:"first_name"`
	Id                      uint64 `json:"id"`
	IsBot                   bool   `json:"is_bot"`
	LanguageCode            string `json:"language_code"`
	LastName                string `json:"last_name"`
	SupportsInlineQueries   bool   `json:"supports_inline_queries"`
	Username                string `json:"username"`
}

type KeyboardButtonPollType struct {
	Type string `json:"type"`
}

type MessageAutoDeleteTimerChanged struct {
	MessageAutoDeleteTime uint64 `json:"message_auto_delete_time"`
}

type ShippingAddress struct {
	City        string `json:"city"`
	CountryCode string `json:"country_code"`
	PostCode    string `json:"post_code"`
	State       string `json:"state"`
	StreetLine1 string `json:"street_line1"`
	StreetLine2 string `json:"street_line2"`
}

type Document struct {
	FileId       string     `json:"file_id"`
	FileName     string     `json:"file_name"`
	FileSize     uint64     `json:"file_size"`
	FileUniqueId string     `json:"file_unique_id"`
	MimeType     string     `json:"mime_type"`
	Thumb        *PhotoSize `json:"thumb"`
}

type ChatMember struct {
	CanAddWebPagePreviews bool   `json:"can_add_web_page_previews"`
	CanBeEdited           bool   `json:"can_be_edited"`
	CanChangeInfo         bool   `json:"can_change_info"`
	CanDeleteMessages     bool   `json:"can_delete_messages"`
	CanEditMessages       bool   `json:"can_edit_messages"`
	CanInviteUsers        bool   `json:"can_invite_users"`
	CanManageChat         bool   `json:"can_manage_chat"`
	CanManageVoiceChats   bool   `json:"can_manage_voice_chats"`
	CanPinMessages        bool   `json:"can_pin_messages"`
	CanPostMessages       bool   `json:"can_post_messages"`
	CanPromoteMembers     bool   `json:"can_promote_members"`
	CanRestrictMembers    bool   `json:"can_restrict_members"`
	CanSendMediaMessages  bool   `json:"can_send_media_messages"`
	CanSendMessages       bool   `json:"can_send_messages"`
	CanSendOtherMessages  bool   `json:"can_send_other_messages"`
	CanSendPolls          bool   `json:"can_send_polls"`
	CustomTitle           string `json:"custom_title"`
	IsAnonymous           bool   `json:"is_anonymous"`
	IsMember              bool   `json:"is_member"`
	Status                string `json:"status"`
	UntilDate             uint64 `json:"until_date"`
	User                  User   `json:"user"`
}

type EncryptedCredentials struct {
	Data   string `json:"data"`
	Hash   string `json:"hash"`
	Secret string `json:"secret"`
}

type InlineQueryResultCachedDocument struct {
	Caption             string                `json:"caption"`
	CaptionEntities     []MessageEntity       `json:"caption_entities"`
	Description         string                `json:"description"`
	DocumentFileId      string                `json:"document_file_id"`
	Id                  string                `json:"id"`
	InputMessageContent *InputMessageContent  `json:"input_message_content"`
	ParseMode           string                `json:"parse_mode"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup"`
	Title               string                `json:"title"`
	Type                string                `json:"type"`
}

type Chat struct {
	Bio              string           `json:"bio"`
	CanSetStickerSet bool             `json:"can_set_sticker_set"`
	Description      string           `json:"description"`
	FirstName        string           `json:"first_name"`
	Id               uint64           `json:"id"`
	InviteLink       string           `json:"invite_link"`
	LastName         string           `json:"last_name"`
	LinkedChatId     uint64           `json:"linked_chat_id"`
	Location         *ChatLocation    `json:"location"`
	Permissions      *ChatPermissions `json:"permissions"`
	Photo            *ChatPhoto       `json:"photo"`
	PinnedMessage    *Message         `json:"pinned_message"`
	SlowModeDelay    uint64           `json:"slow_mode_delay"`
	StickerSetName   string           `json:"sticker_set_name"`
	Title            string           `json:"title"`
	Type             string           `json:"type"`
	Username         string           `json:"username"`
}

type ChatLocation struct {
	Address  string   `json:"address"`
	Location Location `json:"location"`
}

type InputVenueMessageContent struct {
	Address         string  `json:"address"`
	FoursquareId    string  `json:"foursquare_id"`
	FoursquareType  string  `json:"foursquare_type"`
	GooglePlaceId   string  `json:"google_place_id"`
	GooglePlaceType string  `json:"google_place_type"`
	Latitude        float64 `json:"latitude"`
	Longitude       float64 `json:"longitude"`
	Title           string  `json:"title"`
}

type InputMediaPhoto struct {
	Caption         string          `json:"caption"`
	CaptionEntities []MessageEntity `json:"caption_entities"`
	Media           string          `json:"media"`
	ParseMode       string          `json:"parse_mode"`
	Type            string          `json:"type"`
}
