package telegram

type Animation struct {
	Duration     int64      `json:"duration"`
	FileId       string     `json:"file_id"`
	FileUniqueId string     `json:"file_unique_id"`
	Height       int64      `json:"height"`
	Width        int64      `json:"width"`
	FileName     *string    `json:"file_name,omitempty"`
	FileSize     *int64     `json:"file_size,omitempty"`
	MimeType     *string    `json:"mime_type,omitempty"`
	Thumbnail    *PhotoSize `json:"thumbnail,omitempty"`
}

type Audio struct {
	Duration     int64      `json:"duration"`
	FileId       string     `json:"file_id"`
	FileUniqueId string     `json:"file_unique_id"`
	FileName     *string    `json:"file_name,omitempty"`
	FileSize     *int64     `json:"file_size,omitempty"`
	MimeType     *string    `json:"mime_type,omitempty"`
	Performer    *string    `json:"performer,omitempty"`
	Thumbnail    *PhotoSize `json:"thumbnail,omitempty"`
	Title        *string    `json:"title,omitempty"`
}

type BotCommand struct {
	Command     string `json:"command"`
	Description string `json:"description"`
}

type BotCommandScopeAllChatAdministrators struct {
	Type string `json:"type"`
}

type BotCommandScopeAllGroupChats struct {
	Type string `json:"type"`
}

type BotCommandScopeAllPrivateChats struct {
	Type string `json:"type"`
}

type BotCommandScopeChat struct {
	ChatId ChatId `json:"chat_id"`
	Type   string `json:"type"`
}

type BotCommandScopeChatAdministrators struct {
	ChatId ChatId `json:"chat_id"`
	Type   string `json:"type"`
}

type BotCommandScopeChatMember struct {
	ChatId ChatId `json:"chat_id"`
	Type   string `json:"type"`
	UserId int64  `json:"user_id"`
}

type BotCommandScopeDefault struct {
	Type string `json:"type"`
}

type BotDescription struct {
	Description string `json:"description"`
}

type BotName struct {
	Name string `json:"name"`
}

type BotShortDescription struct {
	ShortDescription string `json:"short_description"`
}

type CallbackGame struct {
	// No fields
}

type CallbackQuery struct {
	ChatInstance    string      `json:"chat_instance"`
	From            User        `json:"from"`
	Id              string      `json:"id"`
	Data            *string     `json:"data,omitempty"`
	GameShortName   *string     `json:"game_short_name,omitempty"`
	InlineMessageId *string     `json:"inline_message_id,omitempty"`
	Message         interface{} `json:"message,omitempty"`
}

type Chat struct {
	Id                                 int64            `json:"id"`
	Type                               string           `json:"type"`
	AccentColorId                      *int64           `json:"accent_color_id,omitempty"`
	ActiveUsernames                    []string         `json:"active_usernames,omitempty"`
	AvailableReactions                 []interface{}    `json:"available_reactions,omitempty"`
	BackgroundCustomEmojiId            *string          `json:"background_custom_emoji_id,omitempty"`
	Bio                                *string          `json:"bio,omitempty"`
	CanSetStickerSet                   *bool            `json:"can_set_sticker_set,omitempty"`
	CustomEmojiStickerSetName          *string          `json:"custom_emoji_sticker_set_name,omitempty"`
	Description                        *string          `json:"description,omitempty"`
	EmojiStatusCustomEmojiId           *string          `json:"emoji_status_custom_emoji_id,omitempty"`
	EmojiStatusExpirationDate          *int64           `json:"emoji_status_expiration_date,omitempty"`
	FirstName                          *string          `json:"first_name,omitempty"`
	HasAggressiveAntiSpamEnabled       *bool            `json:"has_aggressive_anti_spam_enabled,omitempty"`
	HasHiddenMembers                   *bool            `json:"has_hidden_members,omitempty"`
	HasPrivateForwards                 *bool            `json:"has_private_forwards,omitempty"`
	HasProtectedContent                *bool            `json:"has_protected_content,omitempty"`
	HasRestrictedVoiceAndVideoMessages *bool            `json:"has_restricted_voice_and_video_messages,omitempty"`
	HasVisibleHistory                  *bool            `json:"has_visible_history,omitempty"`
	InviteLink                         *string          `json:"invite_link,omitempty"`
	IsForum                            *bool            `json:"is_forum,omitempty"`
	JoinByRequest                      *bool            `json:"join_by_request,omitempty"`
	JoinToSendMessages                 *bool            `json:"join_to_send_messages,omitempty"`
	LastName                           *string          `json:"last_name,omitempty"`
	LinkedChatId                       *int64           `json:"linked_chat_id,omitempty"`
	Location                           *ChatLocation    `json:"location,omitempty"`
	MessageAutoDeleteTime              *int64           `json:"message_auto_delete_time,omitempty"`
	Permissions                        *ChatPermissions `json:"permissions,omitempty"`
	Photo                              *ChatPhoto       `json:"photo,omitempty"`
	PinnedMessage                      *Message         `json:"pinned_message,omitempty"`
	ProfileAccentColorId               *int64           `json:"profile_accent_color_id,omitempty"`
	ProfileBackgroundCustomEmojiId     *string          `json:"profile_background_custom_emoji_id,omitempty"`
	SlowModeDelay                      *int64           `json:"slow_mode_delay,omitempty"`
	StickerSetName                     *string          `json:"sticker_set_name,omitempty"`
	Title                              *string          `json:"title,omitempty"`
	UnrestrictBoostCount               *int64           `json:"unrestrict_boost_count,omitempty"`
	Username                           *string          `json:"username,omitempty"`
}

type ChatAdministratorRights struct {
	CanChangeInfo       bool  `json:"can_change_info"`
	CanDeleteMessages   bool  `json:"can_delete_messages"`
	CanDeleteStories    bool  `json:"can_delete_stories"`
	CanEditStories      bool  `json:"can_edit_stories"`
	CanInviteUsers      bool  `json:"can_invite_users"`
	CanManageChat       bool  `json:"can_manage_chat"`
	CanManageVideoChats bool  `json:"can_manage_video_chats"`
	CanPostStories      bool  `json:"can_post_stories"`
	CanPromoteMembers   bool  `json:"can_promote_members"`
	CanRestrictMembers  bool  `json:"can_restrict_members"`
	IsAnonymous         bool  `json:"is_anonymous"`
	CanEditMessages     *bool `json:"can_edit_messages,omitempty"`
	CanManageTopics     *bool `json:"can_manage_topics,omitempty"`
	CanPinMessages      *bool `json:"can_pin_messages,omitempty"`
	CanPostMessages     *bool `json:"can_post_messages,omitempty"`
}

type ChatBoost struct {
	AddDate        int64       `json:"add_date"`
	BoostId        string      `json:"boost_id"`
	ExpirationDate int64       `json:"expiration_date"`
	Source         interface{} `json:"source"`
}

type ChatBoostAdded struct {
	BoostCount int64 `json:"boost_count"`
}

type ChatBoostRemoved struct {
	BoostId    string      `json:"boost_id"`
	Chat       Chat        `json:"chat"`
	RemoveDate int64       `json:"remove_date"`
	Source     interface{} `json:"source"`
}

type ChatBoostSourceGiftCode struct {
	Source string `json:"source"`
	User   User   `json:"user"`
}

type ChatBoostSourceGiveaway struct {
	GiveawayMessageId int64  `json:"giveaway_message_id"`
	Source            string `json:"source"`
	IsUnclaimed       *bool  `json:"is_unclaimed,omitempty"`
	User              *User  `json:"user,omitempty"`
}

type ChatBoostSourcePremium struct {
	Source string `json:"source"`
	User   User   `json:"user"`
}

type ChatBoostUpdated struct {
	Boost ChatBoost `json:"boost"`
	Chat  Chat      `json:"chat"`
}

type ChatInviteLink struct {
	CreatesJoinRequest      bool    `json:"creates_join_request"`
	Creator                 User    `json:"creator"`
	InviteLink              string  `json:"invite_link"`
	IsPrimary               bool    `json:"is_primary"`
	IsRevoked               bool    `json:"is_revoked"`
	ExpireDate              *int64  `json:"expire_date,omitempty"`
	MemberLimit             *int64  `json:"member_limit,omitempty"`
	Name                    *string `json:"name,omitempty"`
	PendingJoinRequestCount *int64  `json:"pending_join_request_count,omitempty"`
}

type ChatJoinRequest struct {
	Chat       Chat            `json:"chat"`
	Date       int64           `json:"date"`
	From       User            `json:"from"`
	UserChatId int64           `json:"user_chat_id"`
	Bio        *string         `json:"bio,omitempty"`
	InviteLink *ChatInviteLink `json:"invite_link,omitempty"`
}

type ChatLocation struct {
	Address  string   `json:"address"`
	Location Location `json:"location"`
}

type ChatMemberAdministrator struct {
	CanBeEdited         bool    `json:"can_be_edited"`
	CanChangeInfo       bool    `json:"can_change_info"`
	CanDeleteMessages   bool    `json:"can_delete_messages"`
	CanDeleteStories    bool    `json:"can_delete_stories"`
	CanEditStories      bool    `json:"can_edit_stories"`
	CanInviteUsers      bool    `json:"can_invite_users"`
	CanManageChat       bool    `json:"can_manage_chat"`
	CanManageVideoChats bool    `json:"can_manage_video_chats"`
	CanPostStories      bool    `json:"can_post_stories"`
	CanPromoteMembers   bool    `json:"can_promote_members"`
	CanRestrictMembers  bool    `json:"can_restrict_members"`
	IsAnonymous         bool    `json:"is_anonymous"`
	Status              string  `json:"status"`
	User                User    `json:"user"`
	CanEditMessages     *bool   `json:"can_edit_messages,omitempty"`
	CanManageTopics     *bool   `json:"can_manage_topics,omitempty"`
	CanPinMessages      *bool   `json:"can_pin_messages,omitempty"`
	CanPostMessages     *bool   `json:"can_post_messages,omitempty"`
	CustomTitle         *string `json:"custom_title,omitempty"`
}

type ChatMemberBanned struct {
	Status    string `json:"status"`
	UntilDate int64  `json:"until_date"`
	User      User   `json:"user"`
}

type ChatMemberLeft struct {
	Status string `json:"status"`
	User   User   `json:"user"`
}

type ChatMemberMember struct {
	Status string `json:"status"`
	User   User   `json:"user"`
}

type ChatMemberOwner struct {
	IsAnonymous bool    `json:"is_anonymous"`
	Status      string  `json:"status"`
	User        User    `json:"user"`
	CustomTitle *string `json:"custom_title,omitempty"`
}

type ChatMemberRestricted struct {
	CanAddWebPagePreviews bool   `json:"can_add_web_page_previews"`
	CanChangeInfo         bool   `json:"can_change_info"`
	CanInviteUsers        bool   `json:"can_invite_users"`
	CanManageTopics       bool   `json:"can_manage_topics"`
	CanPinMessages        bool   `json:"can_pin_messages"`
	CanSendAudios         bool   `json:"can_send_audios"`
	CanSendDocuments      bool   `json:"can_send_documents"`
	CanSendMessages       bool   `json:"can_send_messages"`
	CanSendOtherMessages  bool   `json:"can_send_other_messages"`
	CanSendPhotos         bool   `json:"can_send_photos"`
	CanSendPolls          bool   `json:"can_send_polls"`
	CanSendVideoNotes     bool   `json:"can_send_video_notes"`
	CanSendVideos         bool   `json:"can_send_videos"`
	CanSendVoiceNotes     bool   `json:"can_send_voice_notes"`
	IsMember              bool   `json:"is_member"`
	Status                string `json:"status"`
	UntilDate             int64  `json:"until_date"`
	User                  User   `json:"user"`
}

type ChatMemberUpdated struct {
	Chat                    Chat            `json:"chat"`
	Date                    int64           `json:"date"`
	From                    User            `json:"from"`
	NewChatMember           interface{}     `json:"new_chat_member"`
	OldChatMember           interface{}     `json:"old_chat_member"`
	InviteLink              *ChatInviteLink `json:"invite_link,omitempty"`
	ViaChatFolderInviteLink *bool           `json:"via_chat_folder_invite_link,omitempty"`
}

type ChatPermissions struct {
	CanAddWebPagePreviews *bool `json:"can_add_web_page_previews,omitempty"`
	CanChangeInfo         *bool `json:"can_change_info,omitempty"`
	CanInviteUsers        *bool `json:"can_invite_users,omitempty"`
	CanManageTopics       *bool `json:"can_manage_topics,omitempty"`
	CanPinMessages        *bool `json:"can_pin_messages,omitempty"`
	CanSendAudios         *bool `json:"can_send_audios,omitempty"`
	CanSendDocuments      *bool `json:"can_send_documents,omitempty"`
	CanSendMessages       *bool `json:"can_send_messages,omitempty"`
	CanSendOtherMessages  *bool `json:"can_send_other_messages,omitempty"`
	CanSendPhotos         *bool `json:"can_send_photos,omitempty"`
	CanSendPolls          *bool `json:"can_send_polls,omitempty"`
	CanSendVideoNotes     *bool `json:"can_send_video_notes,omitempty"`
	CanSendVideos         *bool `json:"can_send_videos,omitempty"`
	CanSendVoiceNotes     *bool `json:"can_send_voice_notes,omitempty"`
}

type ChatPhoto struct {
	BigFileId         string `json:"big_file_id"`
	BigFileUniqueId   string `json:"big_file_unique_id"`
	SmallFileId       string `json:"small_file_id"`
	SmallFileUniqueId string `json:"small_file_unique_id"`
}

type ChatShared struct {
	ChatId    int64 `json:"chat_id"`
	RequestId int64 `json:"request_id"`
}

type ChosenInlineResult struct {
	From            User      `json:"from"`
	Query           string    `json:"query"`
	ResultId        string    `json:"result_id"`
	InlineMessageId *string   `json:"inline_message_id,omitempty"`
	Location        *Location `json:"location,omitempty"`
}

type Contact struct {
	FirstName   string  `json:"first_name"`
	PhoneNumber string  `json:"phone_number"`
	LastName    *string `json:"last_name,omitempty"`
	UserId      *int64  `json:"user_id,omitempty"`
	Vcard       *string `json:"vcard,omitempty"`
}

type Dice struct {
	Emoji string `json:"emoji"`
	Value int64  `json:"value"`
}

type Document struct {
	FileId       string     `json:"file_id"`
	FileUniqueId string     `json:"file_unique_id"`
	FileName     *string    `json:"file_name,omitempty"`
	FileSize     *int64     `json:"file_size,omitempty"`
	MimeType     *string    `json:"mime_type,omitempty"`
	Thumbnail    *PhotoSize `json:"thumbnail,omitempty"`
}

type EncryptedCredentials struct {
	Data   string `json:"data"`
	Hash   string `json:"hash"`
	Secret string `json:"secret"`
}

type EncryptedPassportElement struct {
	Hash        string         `json:"hash"`
	Type        string         `json:"type"`
	Data        *string        `json:"data,omitempty"`
	Email       *string        `json:"email,omitempty"`
	Files       []PassportFile `json:"files,omitempty"`
	FrontSide   *PassportFile  `json:"front_side,omitempty"`
	PhoneNumber *string        `json:"phone_number,omitempty"`
	ReverseSide *PassportFile  `json:"reverse_side,omitempty"`
	Selfie      *PassportFile  `json:"selfie,omitempty"`
	Translation []PassportFile `json:"translation,omitempty"`
}

type ExternalReplyInfo struct {
	Origin             interface{}         `json:"origin"`
	Animation          *Animation          `json:"animation,omitempty"`
	Audio              *Audio              `json:"audio,omitempty"`
	Chat               *Chat               `json:"chat,omitempty"`
	Contact            *Contact            `json:"contact,omitempty"`
	Dice               *Dice               `json:"dice,omitempty"`
	Document           *Document           `json:"document,omitempty"`
	Game               *Game               `json:"game,omitempty"`
	Giveaway           *Giveaway           `json:"giveaway,omitempty"`
	GiveawayWinners    *GiveawayWinners    `json:"giveaway_winners,omitempty"`
	HasMediaSpoiler    *bool               `json:"has_media_spoiler,omitempty"`
	Invoice            *Invoice            `json:"invoice,omitempty"`
	LinkPreviewOptions *LinkPreviewOptions `json:"link_preview_options,omitempty"`
	Location           *Location           `json:"location,omitempty"`
	MessageId          *int64              `json:"message_id,omitempty"`
	Photo              []PhotoSize         `json:"photo,omitempty"`
	Poll               *Poll               `json:"poll,omitempty"`
	Sticker            *Sticker            `json:"sticker,omitempty"`
	Story              *Story              `json:"story,omitempty"`
	Venue              *Venue              `json:"venue,omitempty"`
	Video              *Video              `json:"video,omitempty"`
	VideoNote          *VideoNote          `json:"video_note,omitempty"`
	Voice              *Voice              `json:"voice,omitempty"`
}

type File struct {
	FileId       string  `json:"file_id"`
	FileUniqueId string  `json:"file_unique_id"`
	FilePath     *string `json:"file_path,omitempty"`
	FileSize     *int64  `json:"file_size,omitempty"`
}

type ForceReply struct {
	ForceReply            bool    `json:"force_reply"`
	InputFieldPlaceholder *string `json:"input_field_placeholder,omitempty"`
	Selective             *bool   `json:"selective,omitempty"`
}

type ForumTopic struct {
	IconColor         int64   `json:"icon_color"`
	MessageThreadId   int64   `json:"message_thread_id"`
	Name              string  `json:"name"`
	IconCustomEmojiId *string `json:"icon_custom_emoji_id,omitempty"`
}

type ForumTopicClosed struct {
	// No fields
}

type ForumTopicCreated struct {
	IconColor         int64   `json:"icon_color"`
	Name              string  `json:"name"`
	IconCustomEmojiId *string `json:"icon_custom_emoji_id,omitempty"`
}

type ForumTopicEdited struct {
	IconCustomEmojiId *string `json:"icon_custom_emoji_id,omitempty"`
	Name              *string `json:"name,omitempty"`
}

type ForumTopicReopened struct {
	// No fields
}

type Game struct {
	Description  string          `json:"description"`
	Photo        []PhotoSize     `json:"photo"`
	Title        string          `json:"title"`
	Animation    *Animation      `json:"animation,omitempty"`
	Text         *string         `json:"text,omitempty"`
	TextEntities []MessageEntity `json:"text_entities,omitempty"`
}

type GameHighScore struct {
	Position int64 `json:"position"`
	Score    int64 `json:"score"`
	User     User  `json:"user"`
}

type GeneralForumTopicHidden struct {
	// No fields
}

type GeneralForumTopicUnhidden struct {
	// No fields
}

type Giveaway struct {
	Chats                         []Chat   `json:"chats"`
	WinnerCount                   int64    `json:"winner_count"`
	WinnersSelectionDate          int64    `json:"winners_selection_date"`
	CountryCodes                  []string `json:"country_codes,omitempty"`
	HasPublicWinners              *bool    `json:"has_public_winners,omitempty"`
	OnlyNewMembers                *bool    `json:"only_new_members,omitempty"`
	PremiumSubscriptionMonthCount *int64   `json:"premium_subscription_month_count,omitempty"`
	PrizeDescription              *string  `json:"prize_description,omitempty"`
}

type GiveawayCompleted struct {
	WinnerCount         int64    `json:"winner_count"`
	GiveawayMessage     *Message `json:"giveaway_message,omitempty"`
	UnclaimedPrizeCount *int64   `json:"unclaimed_prize_count,omitempty"`
}

type GiveawayCreated struct {
	// No fields
}

type GiveawayWinners struct {
	Chat                          Chat    `json:"chat"`
	GiveawayMessageId             int64   `json:"giveaway_message_id"`
	WinnerCount                   int64   `json:"winner_count"`
	Winners                       []User  `json:"winners"`
	WinnersSelectionDate          int64   `json:"winners_selection_date"`
	AdditionalChatCount           *int64  `json:"additional_chat_count,omitempty"`
	OnlyNewMembers                *bool   `json:"only_new_members,omitempty"`
	PremiumSubscriptionMonthCount *int64  `json:"premium_subscription_month_count,omitempty"`
	PrizeDescription              *string `json:"prize_description,omitempty"`
	UnclaimedPrizeCount           *int64  `json:"unclaimed_prize_count,omitempty"`
	WasRefunded                   *bool   `json:"was_refunded,omitempty"`
}

type InaccessibleMessage struct {
	Chat      Chat  `json:"chat"`
	Date      int64 `json:"date"`
	MessageId int64 `json:"message_id"`
}

type InlineKeyboardButton struct {
	Text                         string                       `json:"text"`
	CallbackData                 *string                      `json:"callback_data,omitempty"`
	CallbackGame                 *CallbackGame                `json:"callback_game,omitempty"`
	LoginUrl                     *LoginUrl                    `json:"login_url,omitempty"`
	Pay                          *bool                        `json:"pay,omitempty"`
	SwitchInlineQuery            *string                      `json:"switch_inline_query,omitempty"`
	SwitchInlineQueryChosenChat  *SwitchInlineQueryChosenChat `json:"switch_inline_query_chosen_chat,omitempty"`
	SwitchInlineQueryCurrentChat *string                      `json:"switch_inline_query_current_chat,omitempty"`
	Url                          *string                      `json:"url,omitempty"`
	WebApp                       *WebAppInfo                  `json:"web_app,omitempty"`
}

type InlineKeyboardMarkup struct {
	InlineKeyboard [][]InlineKeyboardButton `json:"inline_keyboard"`
}

type InlineQuery struct {
	From     User      `json:"from"`
	Id       string    `json:"id"`
	Offset   string    `json:"offset"`
	Query    string    `json:"query"`
	ChatType *string   `json:"chat_type,omitempty"`
	Location *Location `json:"location,omitempty"`
}

type InlineQueryResultArticle struct {
	Id                  string                `json:"id"`
	InputMessageContent interface{}           `json:"input_message_content"`
	Title               string                `json:"title"`
	Type                string                `json:"type"`
	Description         *string               `json:"description,omitempty"`
	HideUrl             *bool                 `json:"hide_url,omitempty"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	ThumbnailHeight     *int64                `json:"thumbnail_height,omitempty"`
	ThumbnailUrl        *string               `json:"thumbnail_url,omitempty"`
	ThumbnailWidth      *int64                `json:"thumbnail_width,omitempty"`
	Url                 *string               `json:"url,omitempty"`
}

type InlineQueryResultAudio struct {
	AudioUrl            string                `json:"audio_url"`
	Id                  string                `json:"id"`
	Title               string                `json:"title"`
	Type                string                `json:"type"`
	AudioDuration       *int64                `json:"audio_duration,omitempty"`
	Caption             *string               `json:"caption,omitempty"`
	CaptionEntities     []MessageEntity       `json:"caption_entities,omitempty"`
	InputMessageContent interface{}           `json:"input_message_content,omitempty"`
	ParseMode           *string               `json:"parse_mode,omitempty"`
	Performer           *string               `json:"performer,omitempty"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

type InlineQueryResultCachedAudio struct {
	AudioFileId         string                `json:"audio_file_id"`
	Id                  string                `json:"id"`
	Type                string                `json:"type"`
	Caption             *string               `json:"caption,omitempty"`
	CaptionEntities     []MessageEntity       `json:"caption_entities,omitempty"`
	InputMessageContent interface{}           `json:"input_message_content,omitempty"`
	ParseMode           *string               `json:"parse_mode,omitempty"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

type InlineQueryResultCachedDocument struct {
	DocumentFileId      string                `json:"document_file_id"`
	Id                  string                `json:"id"`
	Title               string                `json:"title"`
	Type                string                `json:"type"`
	Caption             *string               `json:"caption,omitempty"`
	CaptionEntities     []MessageEntity       `json:"caption_entities,omitempty"`
	Description         *string               `json:"description,omitempty"`
	InputMessageContent interface{}           `json:"input_message_content,omitempty"`
	ParseMode           *string               `json:"parse_mode,omitempty"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

type InlineQueryResultCachedGif struct {
	GifFileId           string                `json:"gif_file_id"`
	Id                  string                `json:"id"`
	Type                string                `json:"type"`
	Caption             *string               `json:"caption,omitempty"`
	CaptionEntities     []MessageEntity       `json:"caption_entities,omitempty"`
	InputMessageContent interface{}           `json:"input_message_content,omitempty"`
	ParseMode           *string               `json:"parse_mode,omitempty"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	Title               *string               `json:"title,omitempty"`
}

type InlineQueryResultCachedMpeg4Gif struct {
	Id                  string                `json:"id"`
	Mpeg4FileId         string                `json:"mpeg4_file_id"`
	Type                string                `json:"type"`
	Caption             *string               `json:"caption,omitempty"`
	CaptionEntities     []MessageEntity       `json:"caption_entities,omitempty"`
	InputMessageContent interface{}           `json:"input_message_content,omitempty"`
	ParseMode           *string               `json:"parse_mode,omitempty"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	Title               *string               `json:"title,omitempty"`
}

type InlineQueryResultCachedPhoto struct {
	Id                  string                `json:"id"`
	PhotoFileId         string                `json:"photo_file_id"`
	Type                string                `json:"type"`
	Caption             *string               `json:"caption,omitempty"`
	CaptionEntities     []MessageEntity       `json:"caption_entities,omitempty"`
	Description         *string               `json:"description,omitempty"`
	InputMessageContent interface{}           `json:"input_message_content,omitempty"`
	ParseMode           *string               `json:"parse_mode,omitempty"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	Title               *string               `json:"title,omitempty"`
}

type InlineQueryResultCachedSticker struct {
	Id                  string                `json:"id"`
	StickerFileId       string                `json:"sticker_file_id"`
	Type                string                `json:"type"`
	InputMessageContent interface{}           `json:"input_message_content,omitempty"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

type InlineQueryResultCachedVideo struct {
	Id                  string                `json:"id"`
	Title               string                `json:"title"`
	Type                string                `json:"type"`
	VideoFileId         string                `json:"video_file_id"`
	Caption             *string               `json:"caption,omitempty"`
	CaptionEntities     []MessageEntity       `json:"caption_entities,omitempty"`
	Description         *string               `json:"description,omitempty"`
	InputMessageContent interface{}           `json:"input_message_content,omitempty"`
	ParseMode           *string               `json:"parse_mode,omitempty"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

type InlineQueryResultCachedVoice struct {
	Id                  string                `json:"id"`
	Title               string                `json:"title"`
	Type                string                `json:"type"`
	VoiceFileId         string                `json:"voice_file_id"`
	Caption             *string               `json:"caption,omitempty"`
	CaptionEntities     []MessageEntity       `json:"caption_entities,omitempty"`
	InputMessageContent interface{}           `json:"input_message_content,omitempty"`
	ParseMode           *string               `json:"parse_mode,omitempty"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

type InlineQueryResultContact struct {
	FirstName           string                `json:"first_name"`
	Id                  string                `json:"id"`
	PhoneNumber         string                `json:"phone_number"`
	Type                string                `json:"type"`
	InputMessageContent interface{}           `json:"input_message_content,omitempty"`
	LastName            *string               `json:"last_name,omitempty"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	ThumbnailHeight     *int64                `json:"thumbnail_height,omitempty"`
	ThumbnailUrl        *string               `json:"thumbnail_url,omitempty"`
	ThumbnailWidth      *int64                `json:"thumbnail_width,omitempty"`
	Vcard               *string               `json:"vcard,omitempty"`
}

type InlineQueryResultDocument struct {
	DocumentUrl         string                `json:"document_url"`
	Id                  string                `json:"id"`
	MimeType            string                `json:"mime_type"`
	Title               string                `json:"title"`
	Type                string                `json:"type"`
	Caption             *string               `json:"caption,omitempty"`
	CaptionEntities     []MessageEntity       `json:"caption_entities,omitempty"`
	Description         *string               `json:"description,omitempty"`
	InputMessageContent interface{}           `json:"input_message_content,omitempty"`
	ParseMode           *string               `json:"parse_mode,omitempty"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	ThumbnailHeight     *int64                `json:"thumbnail_height,omitempty"`
	ThumbnailUrl        *string               `json:"thumbnail_url,omitempty"`
	ThumbnailWidth      *int64                `json:"thumbnail_width,omitempty"`
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
	ThumbnailUrl        string                `json:"thumbnail_url"`
	Type                string                `json:"type"`
	Caption             *string               `json:"caption,omitempty"`
	CaptionEntities     []MessageEntity       `json:"caption_entities,omitempty"`
	GifDuration         *int64                `json:"gif_duration,omitempty"`
	GifHeight           *int64                `json:"gif_height,omitempty"`
	GifWidth            *int64                `json:"gif_width,omitempty"`
	InputMessageContent interface{}           `json:"input_message_content,omitempty"`
	ParseMode           *string               `json:"parse_mode,omitempty"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	ThumbnailMimeType   *string               `json:"thumbnail_mime_type,omitempty"`
	Title               *string               `json:"title,omitempty"`
}

type InlineQueryResultLocation struct {
	Id                   string                `json:"id"`
	Latitude             float64               `json:"latitude"`
	Longitude            float64               `json:"longitude"`
	Title                string                `json:"title"`
	Type                 string                `json:"type"`
	Heading              *int64                `json:"heading,omitempty"`
	HorizontalAccuracy   *float64              `json:"horizontal_accuracy,omitempty"`
	InputMessageContent  interface{}           `json:"input_message_content,omitempty"`
	LivePeriod           *int64                `json:"live_period,omitempty"`
	ProximityAlertRadius *int64                `json:"proximity_alert_radius,omitempty"`
	ReplyMarkup          *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	ThumbnailHeight      *int64                `json:"thumbnail_height,omitempty"`
	ThumbnailUrl         *string               `json:"thumbnail_url,omitempty"`
	ThumbnailWidth       *int64                `json:"thumbnail_width,omitempty"`
}

type InlineQueryResultMpeg4Gif struct {
	Id                  string                `json:"id"`
	Mpeg4Url            string                `json:"mpeg4_url"`
	ThumbnailUrl        string                `json:"thumbnail_url"`
	Type                string                `json:"type"`
	Caption             *string               `json:"caption,omitempty"`
	CaptionEntities     []MessageEntity       `json:"caption_entities,omitempty"`
	InputMessageContent interface{}           `json:"input_message_content,omitempty"`
	Mpeg4Duration       *int64                `json:"mpeg4_duration,omitempty"`
	Mpeg4Height         *int64                `json:"mpeg4_height,omitempty"`
	Mpeg4Width          *int64                `json:"mpeg4_width,omitempty"`
	ParseMode           *string               `json:"parse_mode,omitempty"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	ThumbnailMimeType   *string               `json:"thumbnail_mime_type,omitempty"`
	Title               *string               `json:"title,omitempty"`
}

type InlineQueryResultPhoto struct {
	Id                  string                `json:"id"`
	PhotoUrl            string                `json:"photo_url"`
	ThumbnailUrl        string                `json:"thumbnail_url"`
	Type                string                `json:"type"`
	Caption             *string               `json:"caption,omitempty"`
	CaptionEntities     []MessageEntity       `json:"caption_entities,omitempty"`
	Description         *string               `json:"description,omitempty"`
	InputMessageContent interface{}           `json:"input_message_content,omitempty"`
	ParseMode           *string               `json:"parse_mode,omitempty"`
	PhotoHeight         *int64                `json:"photo_height,omitempty"`
	PhotoWidth          *int64                `json:"photo_width,omitempty"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	Title               *string               `json:"title,omitempty"`
}

type InlineQueryResultVenue struct {
	Address             string                `json:"address"`
	Id                  string                `json:"id"`
	Latitude            float64               `json:"latitude"`
	Longitude           float64               `json:"longitude"`
	Title               string                `json:"title"`
	Type                string                `json:"type"`
	FoursquareId        *string               `json:"foursquare_id,omitempty"`
	FoursquareType      *string               `json:"foursquare_type,omitempty"`
	GooglePlaceId       *string               `json:"google_place_id,omitempty"`
	GooglePlaceType     *string               `json:"google_place_type,omitempty"`
	InputMessageContent interface{}           `json:"input_message_content,omitempty"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	ThumbnailHeight     *int64                `json:"thumbnail_height,omitempty"`
	ThumbnailUrl        *string               `json:"thumbnail_url,omitempty"`
	ThumbnailWidth      *int64                `json:"thumbnail_width,omitempty"`
}

type InlineQueryResultVideo struct {
	Id                  string                `json:"id"`
	MimeType            string                `json:"mime_type"`
	ThumbnailUrl        string                `json:"thumbnail_url"`
	Title               string                `json:"title"`
	Type                string                `json:"type"`
	VideoUrl            string                `json:"video_url"`
	Caption             *string               `json:"caption,omitempty"`
	CaptionEntities     []MessageEntity       `json:"caption_entities,omitempty"`
	Description         *string               `json:"description,omitempty"`
	InputMessageContent interface{}           `json:"input_message_content,omitempty"`
	ParseMode           *string               `json:"parse_mode,omitempty"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	VideoDuration       *int64                `json:"video_duration,omitempty"`
	VideoHeight         *int64                `json:"video_height,omitempty"`
	VideoWidth          *int64                `json:"video_width,omitempty"`
}

type InlineQueryResultVoice struct {
	Id                  string                `json:"id"`
	Title               string                `json:"title"`
	Type                string                `json:"type"`
	VoiceUrl            string                `json:"voice_url"`
	Caption             *string               `json:"caption,omitempty"`
	CaptionEntities     []MessageEntity       `json:"caption_entities,omitempty"`
	InputMessageContent interface{}           `json:"input_message_content,omitempty"`
	ParseMode           *string               `json:"parse_mode,omitempty"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	VoiceDuration       *int64                `json:"voice_duration,omitempty"`
}

type InlineQueryResultsButton struct {
	Text           string      `json:"text"`
	StartParameter *string     `json:"start_parameter,omitempty"`
	WebApp         *WebAppInfo `json:"web_app,omitempty"`
}

type InputContactMessageContent struct {
	FirstName   string  `json:"first_name"`
	PhoneNumber string  `json:"phone_number"`
	LastName    *string `json:"last_name,omitempty"`
	Vcard       *string `json:"vcard,omitempty"`
}

type InputInvoiceMessageContent struct {
	Currency                  string         `json:"currency"`
	Description               string         `json:"description"`
	Payload                   string         `json:"payload"`
	Prices                    []LabeledPrice `json:"prices"`
	ProviderToken             string         `json:"provider_token"`
	Title                     string         `json:"title"`
	IsFlexible                *bool          `json:"is_flexible,omitempty"`
	MaxTipAmount              *int64         `json:"max_tip_amount,omitempty"`
	NeedEmail                 *bool          `json:"need_email,omitempty"`
	NeedName                  *bool          `json:"need_name,omitempty"`
	NeedPhoneNumber           *bool          `json:"need_phone_number,omitempty"`
	NeedShippingAddress       *bool          `json:"need_shipping_address,omitempty"`
	PhotoHeight               *int64         `json:"photo_height,omitempty"`
	PhotoSize                 *int64         `json:"photo_size,omitempty"`
	PhotoUrl                  *string        `json:"photo_url,omitempty"`
	PhotoWidth                *int64         `json:"photo_width,omitempty"`
	ProviderData              *string        `json:"provider_data,omitempty"`
	SendEmailToProvider       *bool          `json:"send_email_to_provider,omitempty"`
	SendPhoneNumberToProvider *bool          `json:"send_phone_number_to_provider,omitempty"`
	SuggestedTipAmounts       []int64        `json:"suggested_tip_amounts,omitempty"`
}

type InputLocationMessageContent struct {
	Latitude             float64  `json:"latitude"`
	Longitude            float64  `json:"longitude"`
	Heading              *int64   `json:"heading,omitempty"`
	HorizontalAccuracy   *float64 `json:"horizontal_accuracy,omitempty"`
	LivePeriod           *int64   `json:"live_period,omitempty"`
	ProximityAlertRadius *int64   `json:"proximity_alert_radius,omitempty"`
}

type InputMediaAnimation struct {
	Media           InputFile       `json:"media"`
	Type            string          `json:"type"`
	Caption         *string         `json:"caption,omitempty"`
	CaptionEntities []MessageEntity `json:"caption_entities,omitempty"`
	Duration        *int64          `json:"duration,omitempty"`
	HasSpoiler      *bool           `json:"has_spoiler,omitempty"`
	Height          *int64          `json:"height,omitempty"`
	ParseMode       *string         `json:"parse_mode,omitempty"`
	Thumbnail       *InputFile      `json:"thumbnail,omitempty"`
	Width           *int64          `json:"width,omitempty"`
}

type InputMediaAudio struct {
	Media           InputFile       `json:"media"`
	Type            string          `json:"type"`
	Caption         *string         `json:"caption,omitempty"`
	CaptionEntities []MessageEntity `json:"caption_entities,omitempty"`
	Duration        *int64          `json:"duration,omitempty"`
	ParseMode       *string         `json:"parse_mode,omitempty"`
	Performer       *string         `json:"performer,omitempty"`
	Thumbnail       *InputFile      `json:"thumbnail,omitempty"`
	Title           *string         `json:"title,omitempty"`
}

type InputMediaDocument struct {
	Media                       InputFile       `json:"media"`
	Type                        string          `json:"type"`
	Caption                     *string         `json:"caption,omitempty"`
	CaptionEntities             []MessageEntity `json:"caption_entities,omitempty"`
	DisableContentTypeDetection *bool           `json:"disable_content_type_detection,omitempty"`
	ParseMode                   *string         `json:"parse_mode,omitempty"`
	Thumbnail                   *InputFile      `json:"thumbnail,omitempty"`
}

type InputMediaPhoto struct {
	Media           InputFile       `json:"media"`
	Type            string          `json:"type"`
	Caption         *string         `json:"caption,omitempty"`
	CaptionEntities []MessageEntity `json:"caption_entities,omitempty"`
	HasSpoiler      *bool           `json:"has_spoiler,omitempty"`
	ParseMode       *string         `json:"parse_mode,omitempty"`
}

type InputMediaVideo struct {
	Media             InputFile       `json:"media"`
	Type              string          `json:"type"`
	Caption           *string         `json:"caption,omitempty"`
	CaptionEntities   []MessageEntity `json:"caption_entities,omitempty"`
	Duration          *int64          `json:"duration,omitempty"`
	HasSpoiler        *bool           `json:"has_spoiler,omitempty"`
	Height            *int64          `json:"height,omitempty"`
	ParseMode         *string         `json:"parse_mode,omitempty"`
	SupportsStreaming *bool           `json:"supports_streaming,omitempty"`
	Thumbnail         *InputFile      `json:"thumbnail,omitempty"`
	Width             *int64          `json:"width,omitempty"`
}

type InputSticker struct {
	EmojiList    []string      `json:"emoji_list"`
	Sticker      InputFile     `json:"sticker"`
	Keywords     []string      `json:"keywords,omitempty"`
	MaskPosition *MaskPosition `json:"mask_position,omitempty"`
}

type InputTextMessageContent struct {
	MessageText        string              `json:"message_text"`
	Entities           []MessageEntity     `json:"entities,omitempty"`
	LinkPreviewOptions *LinkPreviewOptions `json:"link_preview_options,omitempty"`
	ParseMode          *string             `json:"parse_mode,omitempty"`
}

type InputVenueMessageContent struct {
	Address         string  `json:"address"`
	Latitude        float64 `json:"latitude"`
	Longitude       float64 `json:"longitude"`
	Title           string  `json:"title"`
	FoursquareId    *string `json:"foursquare_id,omitempty"`
	FoursquareType  *string `json:"foursquare_type,omitempty"`
	GooglePlaceId   *string `json:"google_place_id,omitempty"`
	GooglePlaceType *string `json:"google_place_type,omitempty"`
}

type Invoice struct {
	Currency       string `json:"currency"`
	Description    string `json:"description"`
	StartParameter string `json:"start_parameter"`
	Title          string `json:"title"`
	TotalAmount    int64  `json:"total_amount"`
}

type KeyboardButton struct {
	Text            string                      `json:"text"`
	RequestChat     *KeyboardButtonRequestChat  `json:"request_chat,omitempty"`
	RequestContact  *bool                       `json:"request_contact,omitempty"`
	RequestLocation *bool                       `json:"request_location,omitempty"`
	RequestPoll     *KeyboardButtonPollType     `json:"request_poll,omitempty"`
	RequestUsers    *KeyboardButtonRequestUsers `json:"request_users,omitempty"`
	WebApp          *WebAppInfo                 `json:"web_app,omitempty"`
}

type KeyboardButtonPollType struct {
	Type *string `json:"type,omitempty"`
}

type KeyboardButtonRequestChat struct {
	ChatIsChannel           bool                     `json:"chat_is_channel"`
	RequestId               int64                    `json:"request_id"`
	BotAdministratorRights  *ChatAdministratorRights `json:"bot_administrator_rights,omitempty"`
	BotIsMember             *bool                    `json:"bot_is_member,omitempty"`
	ChatHasUsername         *bool                    `json:"chat_has_username,omitempty"`
	ChatIsCreated           *bool                    `json:"chat_is_created,omitempty"`
	ChatIsForum             *bool                    `json:"chat_is_forum,omitempty"`
	UserAdministratorRights *ChatAdministratorRights `json:"user_administrator_rights,omitempty"`
}

type KeyboardButtonRequestUsers struct {
	RequestId     int64  `json:"request_id"`
	MaxQuantity   *int64 `json:"max_quantity,omitempty"`
	UserIsBot     *bool  `json:"user_is_bot,omitempty"`
	UserIsPremium *bool  `json:"user_is_premium,omitempty"`
}

type LabeledPrice struct {
	Amount int64  `json:"amount"`
	Label  string `json:"label"`
}

type LinkPreviewOptions struct {
	IsDisabled       *bool   `json:"is_disabled,omitempty"`
	PreferLargeMedia *bool   `json:"prefer_large_media,omitempty"`
	PreferSmallMedia *bool   `json:"prefer_small_media,omitempty"`
	ShowAboveText    *bool   `json:"show_above_text,omitempty"`
	Url              *string `json:"url,omitempty"`
}

type Location struct {
	Latitude             float64  `json:"latitude"`
	Longitude            float64  `json:"longitude"`
	Heading              *int64   `json:"heading,omitempty"`
	HorizontalAccuracy   *float64 `json:"horizontal_accuracy,omitempty"`
	LivePeriod           *int64   `json:"live_period,omitempty"`
	ProximityAlertRadius *int64   `json:"proximity_alert_radius,omitempty"`
}

type LoginUrl struct {
	Url                string  `json:"url"`
	BotUsername        *string `json:"bot_username,omitempty"`
	ForwardText        *string `json:"forward_text,omitempty"`
	RequestWriteAccess *bool   `json:"request_write_access,omitempty"`
}

type MaskPosition struct {
	Point  string  `json:"point"`
	Scale  float64 `json:"scale"`
	XShift float64 `json:"x_shift"`
	YShift float64 `json:"y_shift"`
}

type MenuButtonCommands struct {
	Type string `json:"type"`
}

type MenuButtonDefault struct {
	Type string `json:"type"`
}

type MenuButtonWebApp struct {
	Text   string     `json:"text"`
	Type   string     `json:"type"`
	WebApp WebAppInfo `json:"web_app"`
}

type Message struct {
	Chat                          Chat                           `json:"chat"`
	Date                          int64                          `json:"date"`
	MessageId                     int64                          `json:"message_id"`
	Animation                     *Animation                     `json:"animation,omitempty"`
	Audio                         *Audio                         `json:"audio,omitempty"`
	AuthorSignature               *string                        `json:"author_signature,omitempty"`
	BoostAdded                    *ChatBoostAdded                `json:"boost_added,omitempty"`
	Caption                       *string                        `json:"caption,omitempty"`
	CaptionEntities               []MessageEntity                `json:"caption_entities,omitempty"`
	ChannelChatCreated            *bool                          `json:"channel_chat_created,omitempty"`
	ChatShared                    *ChatShared                    `json:"chat_shared,omitempty"`
	ConnectedWebsite              *string                        `json:"connected_website,omitempty"`
	Contact                       *Contact                       `json:"contact,omitempty"`
	DeleteChatPhoto               *bool                          `json:"delete_chat_photo,omitempty"`
	Dice                          *Dice                          `json:"dice,omitempty"`
	Document                      *Document                      `json:"document,omitempty"`
	EditDate                      *int64                         `json:"edit_date,omitempty"`
	Entities                      []MessageEntity                `json:"entities,omitempty"`
	ExternalReply                 *ExternalReplyInfo             `json:"external_reply,omitempty"`
	ForumTopicClosed              *ForumTopicClosed              `json:"forum_topic_closed,omitempty"`
	ForumTopicCreated             *ForumTopicCreated             `json:"forum_topic_created,omitempty"`
	ForumTopicEdited              *ForumTopicEdited              `json:"forum_topic_edited,omitempty"`
	ForumTopicReopened            *ForumTopicReopened            `json:"forum_topic_reopened,omitempty"`
	ForwardOrigin                 interface{}                    `json:"forward_origin,omitempty"`
	From                          *User                          `json:"from,omitempty"`
	Game                          *Game                          `json:"game,omitempty"`
	GeneralForumTopicHidden       *GeneralForumTopicHidden       `json:"general_forum_topic_hidden,omitempty"`
	GeneralForumTopicUnhidden     *GeneralForumTopicUnhidden     `json:"general_forum_topic_unhidden,omitempty"`
	Giveaway                      *Giveaway                      `json:"giveaway,omitempty"`
	GiveawayCompleted             *GiveawayCompleted             `json:"giveaway_completed,omitempty"`
	GiveawayCreated               *GiveawayCreated               `json:"giveaway_created,omitempty"`
	GiveawayWinners               *GiveawayWinners               `json:"giveaway_winners,omitempty"`
	GroupChatCreated              *bool                          `json:"group_chat_created,omitempty"`
	HasMediaSpoiler               *bool                          `json:"has_media_spoiler,omitempty"`
	HasProtectedContent           *bool                          `json:"has_protected_content,omitempty"`
	Invoice                       *Invoice                       `json:"invoice,omitempty"`
	IsAutomaticForward            *bool                          `json:"is_automatic_forward,omitempty"`
	IsTopicMessage                *bool                          `json:"is_topic_message,omitempty"`
	LeftChatMember                *User                          `json:"left_chat_member,omitempty"`
	LinkPreviewOptions            *LinkPreviewOptions            `json:"link_preview_options,omitempty"`
	Location                      *Location                      `json:"location,omitempty"`
	MediaGroupId                  *string                        `json:"media_group_id,omitempty"`
	MessageAutoDeleteTimerChanged *MessageAutoDeleteTimerChanged `json:"message_auto_delete_timer_changed,omitempty"`
	MessageThreadId               *int64                         `json:"message_thread_id,omitempty"`
	MigrateFromChatId             *int64                         `json:"migrate_from_chat_id,omitempty"`
	MigrateToChatId               *int64                         `json:"migrate_to_chat_id,omitempty"`
	NewChatMembers                []User                         `json:"new_chat_members,omitempty"`
	NewChatPhoto                  []PhotoSize                    `json:"new_chat_photo,omitempty"`
	NewChatTitle                  *string                        `json:"new_chat_title,omitempty"`
	PassportData                  *PassportData                  `json:"passport_data,omitempty"`
	Photo                         []PhotoSize                    `json:"photo,omitempty"`
	PinnedMessage                 interface{}                    `json:"pinned_message,omitempty"`
	Poll                          *Poll                          `json:"poll,omitempty"`
	ProximityAlertTriggered       *ProximityAlertTriggered       `json:"proximity_alert_triggered,omitempty"`
	Quote                         *TextQuote                     `json:"quote,omitempty"`
	ReplyMarkup                   *InlineKeyboardMarkup          `json:"reply_markup,omitempty"`
	ReplyToMessage                *Message                       `json:"reply_to_message,omitempty"`
	ReplyToStory                  *Story                         `json:"reply_to_story,omitempty"`
	SenderBoostCount              *int64                         `json:"sender_boost_count,omitempty"`
	SenderChat                    *Chat                          `json:"sender_chat,omitempty"`
	Sticker                       *Sticker                       `json:"sticker,omitempty"`
	Story                         *Story                         `json:"story,omitempty"`
	SuccessfulPayment             *SuccessfulPayment             `json:"successful_payment,omitempty"`
	SupergroupChatCreated         *bool                          `json:"supergroup_chat_created,omitempty"`
	Text                          *string                        `json:"text,omitempty"`
	UsersShared                   *UsersShared                   `json:"users_shared,omitempty"`
	Venue                         *Venue                         `json:"venue,omitempty"`
	ViaBot                        *User                          `json:"via_bot,omitempty"`
	Video                         *Video                         `json:"video,omitempty"`
	VideoChatEnded                *VideoChatEnded                `json:"video_chat_ended,omitempty"`
	VideoChatParticipantsInvited  *VideoChatParticipantsInvited  `json:"video_chat_participants_invited,omitempty"`
	VideoChatScheduled            *VideoChatScheduled            `json:"video_chat_scheduled,omitempty"`
	VideoChatStarted              *VideoChatStarted              `json:"video_chat_started,omitempty"`
	VideoNote                     *VideoNote                     `json:"video_note,omitempty"`
	Voice                         *Voice                         `json:"voice,omitempty"`
	WebAppData                    *WebAppData                    `json:"web_app_data,omitempty"`
	WriteAccessAllowed            *WriteAccessAllowed            `json:"write_access_allowed,omitempty"`
}

type MessageAutoDeleteTimerChanged struct {
	MessageAutoDeleteTime int64 `json:"message_auto_delete_time"`
}

type MessageEntity struct {
	Length        int64   `json:"length"`
	Offset        int64   `json:"offset"`
	Type          string  `json:"type"`
	CustomEmojiId *string `json:"custom_emoji_id,omitempty"`
	Language      *string `json:"language,omitempty"`
	Url           *string `json:"url,omitempty"`
	User          *User   `json:"user,omitempty"`
}

type MessageId struct {
	MessageId int64 `json:"message_id"`
}

type MessageOriginChannel struct {
	Chat            Chat    `json:"chat"`
	Date            int64   `json:"date"`
	MessageId       int64   `json:"message_id"`
	Type            string  `json:"type"`
	AuthorSignature *string `json:"author_signature,omitempty"`
}

type MessageOriginChat struct {
	Date            int64   `json:"date"`
	SenderChat      Chat    `json:"sender_chat"`
	Type            string  `json:"type"`
	AuthorSignature *string `json:"author_signature,omitempty"`
}

type MessageOriginHiddenUser struct {
	Date           int64  `json:"date"`
	SenderUserName string `json:"sender_user_name"`
	Type           string `json:"type"`
}

type MessageOriginUser struct {
	Date       int64  `json:"date"`
	SenderUser User   `json:"sender_user"`
	Type       string `json:"type"`
}

type MessageReactionCountUpdated struct {
	Chat      Chat            `json:"chat"`
	Date      int64           `json:"date"`
	MessageId int64           `json:"message_id"`
	Reactions []ReactionCount `json:"reactions"`
}

type MessageReactionUpdated struct {
	Chat        Chat          `json:"chat"`
	Date        int64         `json:"date"`
	MessageId   int64         `json:"message_id"`
	NewReaction []interface{} `json:"new_reaction"`
	OldReaction []interface{} `json:"old_reaction"`
	ActorChat   *Chat         `json:"actor_chat,omitempty"`
	User        *User         `json:"user,omitempty"`
}

type OrderInfo struct {
	Email           *string          `json:"email,omitempty"`
	Name            *string          `json:"name,omitempty"`
	PhoneNumber     *string          `json:"phone_number,omitempty"`
	ShippingAddress *ShippingAddress `json:"shipping_address,omitempty"`
}

type PassportData struct {
	Credentials EncryptedCredentials       `json:"credentials"`
	Data        []EncryptedPassportElement `json:"data"`
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
	FileDate     int64  `json:"file_date"`
	FileId       string `json:"file_id"`
	FileSize     int64  `json:"file_size"`
	FileUniqueId string `json:"file_unique_id"`
}

type PhotoSize struct {
	FileId       string `json:"file_id"`
	FileUniqueId string `json:"file_unique_id"`
	Height       int64  `json:"height"`
	Width        int64  `json:"width"`
	FileSize     *int64 `json:"file_size,omitempty"`
}

type Poll struct {
	AllowsMultipleAnswers bool            `json:"allows_multiple_answers"`
	Id                    string          `json:"id"`
	IsAnonymous           bool            `json:"is_anonymous"`
	IsClosed              bool            `json:"is_closed"`
	Options               []PollOption    `json:"options"`
	Question              string          `json:"question"`
	TotalVoterCount       int64           `json:"total_voter_count"`
	Type                  string          `json:"type"`
	CloseDate             *int64          `json:"close_date,omitempty"`
	CorrectOptionId       *int64          `json:"correct_option_id,omitempty"`
	Explanation           *string         `json:"explanation,omitempty"`
	ExplanationEntities   []MessageEntity `json:"explanation_entities,omitempty"`
	OpenPeriod            *int64          `json:"open_period,omitempty"`
}

type PollAnswer struct {
	OptionIds []int64 `json:"option_ids"`
	PollId    string  `json:"poll_id"`
	User      *User   `json:"user,omitempty"`
	VoterChat *Chat   `json:"voter_chat,omitempty"`
}

type PollOption struct {
	Text       string `json:"text"`
	VoterCount int64  `json:"voter_count"`
}

type PreCheckoutQuery struct {
	Currency         string     `json:"currency"`
	From             User       `json:"from"`
	Id               string     `json:"id"`
	InvoicePayload   string     `json:"invoice_payload"`
	TotalAmount      int64      `json:"total_amount"`
	OrderInfo        *OrderInfo `json:"order_info,omitempty"`
	ShippingOptionId *string    `json:"shipping_option_id,omitempty"`
}

type ProximityAlertTriggered struct {
	Distance int64 `json:"distance"`
	Traveler User  `json:"traveler"`
	Watcher  User  `json:"watcher"`
}

type ReactionCount struct {
	TotalCount int64       `json:"total_count"`
	Type       interface{} `json:"type"`
}

type ReactionTypeCustomEmoji struct {
	CustomEmojiId string `json:"custom_emoji_id"`
	Type          string `json:"type"`
}

type ReactionTypeEmoji struct {
	Emoji string `json:"emoji"`
	Type  string `json:"type"`
}

type ReplyKeyboardMarkup struct {
	Keyboard              [][]KeyboardButton `json:"keyboard"`
	InputFieldPlaceholder *string            `json:"input_field_placeholder,omitempty"`
	IsPersistent          *bool              `json:"is_persistent,omitempty"`
	OneTimeKeyboard       *bool              `json:"one_time_keyboard,omitempty"`
	ResizeKeyboard        *bool              `json:"resize_keyboard,omitempty"`
	Selective             *bool              `json:"selective,omitempty"`
}

type ReplyKeyboardRemove struct {
	RemoveKeyboard bool  `json:"remove_keyboard"`
	Selective      *bool `json:"selective,omitempty"`
}

type ReplyParameters struct {
	MessageId                int64           `json:"message_id"`
	AllowSendingWithoutReply *bool           `json:"allow_sending_without_reply,omitempty"`
	ChatId                   *ChatId         `json:"chat_id,omitempty"`
	Quote                    *string         `json:"quote,omitempty"`
	QuoteEntities            []MessageEntity `json:"quote_entities,omitempty"`
	QuoteParseMode           *string         `json:"quote_parse_mode,omitempty"`
	QuotePosition            *int64          `json:"quote_position,omitempty"`
}

type ResponseParameters struct {
	MigrateToChatId *int64 `json:"migrate_to_chat_id,omitempty"`
	RetryAfter      *int64 `json:"retry_after,omitempty"`
}

type SentWebAppMessage struct {
	InlineMessageId *string `json:"inline_message_id,omitempty"`
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
	FileId           string        `json:"file_id"`
	FileUniqueId     string        `json:"file_unique_id"`
	Height           int64         `json:"height"`
	IsAnimated       bool          `json:"is_animated"`
	IsVideo          bool          `json:"is_video"`
	Type             string        `json:"type"`
	Width            int64         `json:"width"`
	CustomEmojiId    *string       `json:"custom_emoji_id,omitempty"`
	Emoji            *string       `json:"emoji,omitempty"`
	FileSize         *int64        `json:"file_size,omitempty"`
	MaskPosition     *MaskPosition `json:"mask_position,omitempty"`
	NeedsRepainting  *bool         `json:"needs_repainting,omitempty"`
	PremiumAnimation *File         `json:"premium_animation,omitempty"`
	SetName          *string       `json:"set_name,omitempty"`
	Thumbnail        *PhotoSize    `json:"thumbnail,omitempty"`
}

type StickerSet struct {
	IsAnimated  bool       `json:"is_animated"`
	IsVideo     bool       `json:"is_video"`
	Name        string     `json:"name"`
	StickerType string     `json:"sticker_type"`
	Stickers    []Sticker  `json:"stickers"`
	Title       string     `json:"title"`
	Thumbnail   *PhotoSize `json:"thumbnail,omitempty"`
}

type Story struct {
	Chat Chat  `json:"chat"`
	Id   int64 `json:"id"`
}

type SuccessfulPayment struct {
	Currency                string     `json:"currency"`
	InvoicePayload          string     `json:"invoice_payload"`
	ProviderPaymentChargeId string     `json:"provider_payment_charge_id"`
	TelegramPaymentChargeId string     `json:"telegram_payment_charge_id"`
	TotalAmount             int64      `json:"total_amount"`
	OrderInfo               *OrderInfo `json:"order_info,omitempty"`
	ShippingOptionId        *string    `json:"shipping_option_id,omitempty"`
}

type SwitchInlineQueryChosenChat struct {
	AllowBotChats     *bool   `json:"allow_bot_chats,omitempty"`
	AllowChannelChats *bool   `json:"allow_channel_chats,omitempty"`
	AllowGroupChats   *bool   `json:"allow_group_chats,omitempty"`
	AllowUserChats    *bool   `json:"allow_user_chats,omitempty"`
	Query             *string `json:"query,omitempty"`
}

type TextQuote struct {
	Position int64           `json:"position"`
	Text     string          `json:"text"`
	Entities []MessageEntity `json:"entities,omitempty"`
	IsManual *bool           `json:"is_manual,omitempty"`
}

type Update struct {
	UpdateId             int64                        `json:"update_id"`
	CallbackQuery        *CallbackQuery               `json:"callback_query,omitempty"`
	ChannelPost          *Message                     `json:"channel_post,omitempty"`
	ChatBoost            *ChatBoostUpdated            `json:"chat_boost,omitempty"`
	ChatJoinRequest      *ChatJoinRequest             `json:"chat_join_request,omitempty"`
	ChatMember           *ChatMemberUpdated           `json:"chat_member,omitempty"`
	ChosenInlineResult   *ChosenInlineResult          `json:"chosen_inline_result,omitempty"`
	EditedChannelPost    *Message                     `json:"edited_channel_post,omitempty"`
	EditedMessage        *Message                     `json:"edited_message,omitempty"`
	InlineQuery          *InlineQuery                 `json:"inline_query,omitempty"`
	Message              *Message                     `json:"message,omitempty"`
	MessageReaction      *MessageReactionUpdated      `json:"message_reaction,omitempty"`
	MessageReactionCount *MessageReactionCountUpdated `json:"message_reaction_count,omitempty"`
	MyChatMember         *ChatMemberUpdated           `json:"my_chat_member,omitempty"`
	Poll                 *Poll                        `json:"poll,omitempty"`
	PollAnswer           *PollAnswer                  `json:"poll_answer,omitempty"`
	PreCheckoutQuery     *PreCheckoutQuery            `json:"pre_checkout_query,omitempty"`
	RemovedChatBoost     *ChatBoostRemoved            `json:"removed_chat_boost,omitempty"`
	ShippingQuery        *ShippingQuery               `json:"shipping_query,omitempty"`
}

type User struct {
	FirstName               string  `json:"first_name"`
	Id                      int64   `json:"id"`
	IsBot                   bool    `json:"is_bot"`
	AddedToAttachmentMenu   *bool   `json:"added_to_attachment_menu,omitempty"`
	CanJoinGroups           *bool   `json:"can_join_groups,omitempty"`
	CanReadAllGroupMessages *bool   `json:"can_read_all_group_messages,omitempty"`
	IsPremium               *bool   `json:"is_premium,omitempty"`
	LanguageCode            *string `json:"language_code,omitempty"`
	LastName                *string `json:"last_name,omitempty"`
	SupportsInlineQueries   *bool   `json:"supports_inline_queries,omitempty"`
	Username                *string `json:"username,omitempty"`
}

type UserChatBoosts struct {
	Boosts []ChatBoost `json:"boosts"`
}

type UserProfilePhotos struct {
	Photos     [][]PhotoSize `json:"photos"`
	TotalCount int64         `json:"total_count"`
}

type UsersShared struct {
	RequestId int64   `json:"request_id"`
	UserIds   []int64 `json:"user_ids"`
}

type Venue struct {
	Address         string   `json:"address"`
	Location        Location `json:"location"`
	Title           string   `json:"title"`
	FoursquareId    *string  `json:"foursquare_id,omitempty"`
	FoursquareType  *string  `json:"foursquare_type,omitempty"`
	GooglePlaceId   *string  `json:"google_place_id,omitempty"`
	GooglePlaceType *string  `json:"google_place_type,omitempty"`
}

type Video struct {
	Duration     int64      `json:"duration"`
	FileId       string     `json:"file_id"`
	FileUniqueId string     `json:"file_unique_id"`
	Height       int64      `json:"height"`
	Width        int64      `json:"width"`
	FileName     *string    `json:"file_name,omitempty"`
	FileSize     *int64     `json:"file_size,omitempty"`
	MimeType     *string    `json:"mime_type,omitempty"`
	Thumbnail    *PhotoSize `json:"thumbnail,omitempty"`
}

type VideoChatEnded struct {
	Duration int64 `json:"duration"`
}

type VideoChatParticipantsInvited struct {
	Users []User `json:"users"`
}

type VideoChatScheduled struct {
	StartDate int64 `json:"start_date"`
}

type VideoChatStarted struct {
	// No fields
}

type VideoNote struct {
	Duration     int64      `json:"duration"`
	FileId       string     `json:"file_id"`
	FileUniqueId string     `json:"file_unique_id"`
	Length       int64      `json:"length"`
	FileSize     *int64     `json:"file_size,omitempty"`
	Thumbnail    *PhotoSize `json:"thumbnail,omitempty"`
}

type Voice struct {
	Duration     int64   `json:"duration"`
	FileId       string  `json:"file_id"`
	FileUniqueId string  `json:"file_unique_id"`
	FileSize     *int64  `json:"file_size,omitempty"`
	MimeType     *string `json:"mime_type,omitempty"`
}

type WebAppData struct {
	ButtonText string `json:"button_text"`
	Data       string `json:"data"`
}

type WebAppInfo struct {
	Url string `json:"url"`
}

type WebhookInfo struct {
	HasCustomCertificate         bool     `json:"has_custom_certificate"`
	PendingUpdateCount           int64    `json:"pending_update_count"`
	Url                          string   `json:"url"`
	AllowedUpdates               []string `json:"allowed_updates,omitempty"`
	IpAddress                    *string  `json:"ip_address,omitempty"`
	LastErrorDate                *int64   `json:"last_error_date,omitempty"`
	LastErrorMessage             *string  `json:"last_error_message,omitempty"`
	LastSynchronizationErrorDate *int64   `json:"last_synchronization_error_date,omitempty"`
	MaxConnections               *int64   `json:"max_connections,omitempty"`
}

type WriteAccessAllowed struct {
	FromAttachmentMenu *bool   `json:"from_attachment_menu,omitempty"`
	FromRequest        *bool   `json:"from_request,omitempty"`
	WebAppName         *string `json:"web_app_name,omitempty"`
}
