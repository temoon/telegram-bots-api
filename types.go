package telegram

import (
    "encoding/json"
)

type Request interface {
    IsMultipart() bool
    GetValues() (map[string]interface{}, error)
}

type Response struct {
    OK          bool                `json:"ok"`
    Description string              `json:"description,omitempty"`
    Result      json.RawMessage     `json:"result,omitempty"`
    ErrorCode   uint32              `json:"error_code,omitempty"`
    Parameters  *ResponseParameters `json:"parameters,omitempty"`
}

type Update struct {
    UpdateID           uint32              `json:"update_id"`
    Message            *Message            `json:"message,omitempty"`
    EditedMessage      *Message            `json:"edited_message,omitempty"`
    ChannelPost        *Message            `json:"channel_post,omitempty"`
    EditedChannelPost  *Message            `json:"edited_channel_post,omitempty"`
    InlineQuery        *InlineQuery        `json:"inline_query,omitempty"`
    ChosenInlineResult *ChosenInlineResult `json:"chosen_inline_result,omitempty"`
    CallbackQuery      *CallbackQuery      `json:"callback_query,omitempty"`
    ShippingQuery      *ShippingQuery      `json:"shipping_query,omitempty"`
    PreCheckoutQuery   *PreCheckoutQuery   `json:"pre_checkout_query,omitempty"`
}

type WebhookInfo struct {
    URL                  string   `json:"url"`
    HasCustomCertificate bool     `json:"has_custom_certificate"`
    PendingUpdateCount   uint32   `json:"pending_update_count"`
    LastErrorDate        uint32   `json:"last_error_date,omitempty"`
    LastErrorMessage     string   `json:"last_error_message,omitempty"`
    MaxConnections       uint32   `json:"max_connections,omitempty"`
    AllowedUpdates       []string `json:"allowed_updates,omitempty"`
}

type User struct {
    ID           uint32 `json:"id"`
    IsBot        bool   `json:"is_bot"`
    FirstName    string `json:"first_name"`
    LastName     string `json:"last_name,omitempty"`
    Username     string `json:"username,omitempty"`
    LanguageCode string `json:"language_code,omitempty"`
}

type Chat struct {
    ID                          uint64     `json:"id"`
    Type                        string     `json:"type"`
    Title                       string     `json:"title,omitempty"`
    Username                    string     `json:"username,omitempty"`
    FirstName                   string     `json:"first_name,omitempty"`
    LastName                    string     `json:"last_name,omitempty"`
    AllMembersAreAdministrators bool       `json:"all_members_are_administrators,omitempty"`
    Photo                       *ChatPhoto `json:"photo,omitempty"`
    Description                 string     `json:"description,omitempty"`
    InviteLink                  string     `json:"invite_link,omitempty"`
    PinnedMessage               *Message   `json:"pinned_message,omitempty"`
    StickerSetName              string     `json:"sticker_set_name,omitempty"`
    CanSetStickerSet            bool       `json:"can_set_sticker_set,omitempty"`
}

type Message struct {
    MessageID             uint32             `json:"message_id"`
    From                  *User              `json:"from,omitempty"`
    Date                  uint32             `json:"date"`
    Chat                  Chat               `json:"chat"`
    ForwardFrom           *User              `json:"forward_from,omitempty"`
    ForwardFromChat       *Chat              `json:"forward_from_chat,omitempty"`
    ForwardFromMessageID  uint32             `json:"forward_from_message_id,omitempty"`
    ForwardSignature      string             `json:"forward_signature,omitempty"`
    ForwardDate           uint32             `json:"forward_date,omitempty"`
    ReplyToMessage        *Message           `json:"reply_to_message,omitempty"`
    EditDate              uint32             `json:"edit_date,omitempty"`
    MediaGroupID          string             `json:"media_group_id,omitempty"`
    AuthorSignature       string             `json:"author_signature,omitempty"`
    Text                  string             `json:"text,omitempty"`
    Entities              *[]MessageEntity   `json:"entities,omitempty"`
    CaptionEntities       *[]MessageEntity   `json:"caption_entities,omitempty"`
    Audio                 *Audio             `json:"audio,omitempty"`
    Document              *Document          `json:"document,omitempty"`
    Game                  *Game              `json:"game,omitempty"`
    Photo                 *[]PhotoSize       `json:"photo,omitempty"`
    Sticker               *Sticker           `json:"sticker,omitempty"`
    Video                 *Video             `json:"video,omitempty"`
    Voice                 *Voice             `json:"voice,omitempty"`
    VideoNote             *VideoNote         `json:"video_note,omitempty"`
    Caption               string             `json:"caption,omitempty"`
    Contact               *Contact           `json:"contact,omitempty"`
    Location              *Location          `json:"location,omitempty"`
    Venue                 *Venue             `json:"venue,omitempty"`
    NewChatMembers        *[]User            `json:"new_chat_members,omitempty"`
    LeftChatMember        *User              `json:"left_chat_member,omitempty"`
    NewChatTitle          string             `json:"new_chat_title,omitempty"`
    NewChatPhoto          *[]PhotoSize       `json:"new_chat_photo,omitempty"`
    DeleteChatPhoto       bool               `json:"delete_chat_photo,omitempty"`
    GroupChatCreated      bool               `json:"group_chat_created,omitempty"`
    SupergroupChatCreated bool               `json:"supergroup_chat_created,omitempty"`
    ChannelChatCreated    bool               `json:"channel_chat_created,omitempty"`
    MigrateToChatID       uint64             `json:"migrate_to_chat_id,omitempty"`
    MigrateFromChatID     uint64             `json:"migrate_from_chat_id,omitempty"`
    PinnedMessage         *Message           `json:"pinned_message,omitempty"`
    Invoice               *Invoice           `json:"invoice,omitempty"`
    SuccessfulPayment     *SuccessfulPayment `json:"successful_payment,omitempty"`
    ConnectedWebsite      string             `json:"connected_website,omitempty"`
}

type MessageEntity struct {
    Type   string `json:"type"`
    Offset uint32 `json:"offset"`
    Length uint32 `json:"length"`
    URL    string `json:"url,omitempty"`
    User   *User  `json:"user,omitempty"`
}

type PhotoSize struct {
    FileID   string `json:"file_id"`
    Width    uint32 `json:"width"`
    Height   uint32 `json:"height"`
    FileSize uint32 `json:"file_size,omitempty"`
}

type Audio struct {
    FileID    string `json:"file_id"`
    Duration  uint32 `json:"duration"`
    Performer string `json:"performer,omitempty"`
    Title     string `json:"title,omitempty"`
    MimeType  string `json:"mime_type,omitempty"`
    FileSize  uint32 `json:"file_size,omitempty"`
}

type Document struct {
    FileID   string     `json:"file_id"`
    Thumb    *PhotoSize `json:"thumb,omitempty"`
    FileName string     `json:"file_name,omitempty"`
    MimeType string     `json:"mime_type,omitempty"`
    FileSize uint32     `json:"file_size,omitempty"`
}

type Video struct {
    FileID   string     `json:"file_id"`
    Width    uint32     `json:"width"`
    Height   uint32     `json:"height"`
    Duration uint32     `json:"duration"`
    Thumb    *PhotoSize `json:"thumb,omitempty"`
    MimeType string     `json:"mime_type,omitempty"`
    FileSize uint32     `json:"file_size,omitempty"`
}

type Voice struct {
    FileID   string `json:"file_id"`
    Duration uint32 `json:"duration"`
    MimeType string `json:"mime_type,omitempty"`
    FileSize uint32 `json:"file_size,omitempty"`
}

type VideoNote struct {
    FileID   string     `json:"file_id"`
    Length   uint32     `json:"length"`
    Duration uint32     `json:"duration"`
    Thumb    *PhotoSize `json:"thumb,omitempty"`
    FileSize uint32     `json:"file_size,omitempty"`
}

type Contact struct {
    PhoneNumber string `json:"phone_number"`
    FirstName   string `json:"first_name"`
    LastName    string `json:"last_name,omitempty"`
    UserID      uint32 `json:"user_id,omitempty"`
}

type Location struct {
    Longitude float64 `json:"longitude"`
    Latitude  float64 `json:"latitude"`
}

type Venue struct {
    Location     *Location `json:"location"`
    Title        string    `json:"title"`
    Address      string    `json:"address"`
    FoursquareID string    `json:"foursquare_id,omitempty"`
}

type UserProfilePhotos struct {
    TotalCount uint32        `json:"total_count"`
    Photos     [][]PhotoSize `json:"photos"`
}

type File struct {
    FileID   string `json:"file_id"`
    FileSize uint32 `json:"file_size,omitempty"`
    FilePath string `json:"file_path,omitempty"`
}

type ReplyKeyboardMarkup struct {
    Keyboard        [][]KeyboardButton `json:"keyboard"`
    ResizeKeyboard  bool               `json:"resize_keyboard,omitempty"`
    OneTimeKeyboard bool               `json:"one_time_keyboard,omitempty"`
    Selective       bool               `json:"selective,omitempty"`
}

type KeyboardButton struct {
    Text            string `json:"text"`
    RequestContact  bool   `json:"request_contact,omitempty"`
    RequestLocation bool   `json:"request_location,omitempty"`
}

type ReplyKeyboardRemove struct {
    RemoveKeyboard bool `json:"remove_keyboard"`
    Selective      bool `json:"selective,omitempty"`
}

type InlineKeyboardMarkup struct {
    InlineKeyboard [][]InlineKeyboardButton `json:"inline_keyboard"`
}

type InlineKeyboardButton struct {
    Text                         string        `json:"text"`
    URL                          string        `json:"url,omitempty"`
    CallbackData                 string        `json:"callback_data,omitempty"`
    SwitchInlineQuery            string        `json:"switch_inline_query,omitempty"`
    SwitchInlineQueryCurrentChat string        `json:"switch_inline_query_current_chat,omitempty"`
    CallbackGame                 *CallbackGame `json:"callback_game,omitempty"`
    Pay                          bool          `json:"pay,omitempty"`
}

type CallbackQuery struct {
    ID              string   `json:"id"`
    From            User     `json:"from"`
    Message         *Message `json:"message,omitempty"`
    InlineMessageID string   `json:"inline_message_id,omitempty"`
    ChatInstance    string   `json:"chat_instance"`
    Data            string   `json:"data,omitempty"`
    GameShortName   string   `json:"game_short_name,omitempty"`
}

type ForceReply struct {
    ForceReply bool `json:"force_reply"`
    Selective  bool `json:"selective,omitempty"`
}

type ChatPhoto struct {
    SmallFileID string `json:"small_file_id"`
    BigFileID   string `json:"big_file_id"`
}

type ChatMember struct {
    User                  User   `json:"user"`
    Status                string `json:"status"`
    UntilDate             uint32 `json:"until_date,omitempty"`
    CanBeEdited           bool   `json:"can_be_edited,omitempty"`
    CanChangeInfo         bool   `json:"can_change_info,omitempty"`
    CanPostMessages       bool   `json:"can_post_messages,omitempty"`
    CanEditMessages       bool   `json:"can_edit_messages,omitempty"`
    CanDeleteMessages     bool   `json:"can_delete_messages,omitempty"`
    CanInviteUsers        bool   `json:"can_invite_users,omitempty"`
    CanRestrictMembers    bool   `json:"can_restrict_members,omitempty"`
    CanPinMessages        bool   `json:"can_pin_messages,omitempty"`
    CanPromoteMembers     bool   `json:"can_promote_members,omitempty"`
    CanSendMessages       bool   `json:"can_send_messages,omitempty"`
    CanSendMediaMessages  bool   `json:"can_send_media_messages,omitempty"`
    CanSendOtherMessages  bool   `json:"can_send_other_messages,omitempty"`
    CanAddWebPagePreviews bool   `json:"can_add_web_page_previews,omitempty"`
}

type ResponseParameters struct {
    MigrateToChatID uint64 `json:"migrate_to_chat_id,omitempty"`
    RetryAfter      uint32 `json:"retry_after,omitempty"`
}

type InputMedia interface{}

type InputMediaPhoto struct {
    Type      string `json:"type"`
    Media     string `json:"media"`
    Caption   string `json:"caption,omitempty"`
    ParseMode string `json:"parse_mode,omitempty"`
}

type InputMediaVideo struct {
    Type              string `json:"type"`
    Media             string `json:"media"`
    Caption           string `json:"caption,omitempty"`
    ParseMode         string `json:"parse_mode,omitempty"`
    Width             uint32 `json:"width,omitempty"`
    Height            uint32 `json:"height,omitempty"`
    Duration          uint32 `json:"duration,omitempty"`
    SupportsStreaming bool   `json:"supports_streaming,omitempty"`
}

type Sticker struct {
    FileID       string        `json:"file_id"`
    Width        uint32        `json:"width"`
    Height       uint32        `json:"height"`
    Thumb        *PhotoSize    `json:"thumb,omitempty"`
    Emoji        string        `json:"emoji,omitempty"`
    SetName      string        `json:"set_name,omitempty"`
    MaskPosition *MaskPosition `json:"mask_position,omitempty"`
    FileSize     uint32        `json:"file_size,omitempty"`
}

type StickerSet struct {
    Name          string    `json:"name"`
    Title         string    `json:"title"`
    ContainsMasks bool      `json:"contains_masks"`
    Stickers      []Sticker `json:"stickers"`
}

type MaskPosition struct {
    Point  string  `json:"point"`
    XShift float64 `json:"x_shift"`
    YShift float64 `json:"y_shift"`
    Scale  float64 `json:"scale"`
}

type InlineQuery struct {
    ID       string    `json:"id"`
    From     User      `json:"from"`
    Location *Location `json:"location,omitempty"`
    Query    string    `json:"query"`
    Offset   string    `json:"offset"`
}

type InlineQueryResult interface{}

type InlineQueryResultArticle struct {
    Type                string                `json:"type"`
    ID                  string                `json:"id"`
    Title               string                `json:"title"`
    InputMessageContent InputMessageContent   `json:"input_message_content"`
    ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
    URL                 string                `json:"url,omitempty"`
    HideURL             bool                  `json:"hide_url,omitempty"`
    Description         string                `json:"description,omitempty"`
    ThumbURL            string                `json:"thumb_url,omitempty"`
    ThumbWidth          uint32                `json:"thumb_width,omitempty"`
    ThumbHeight         uint32                `json:"thumb_height,omitempty"`
}

type InlineQueryResultPhoto struct {
    Type                string                `json:"type"`
    ID                  string                `json:"id"`
    PhotoURL            string                `json:"photo_url"`
    ThumbURL            string                `json:"thumb_url"`
    PhotoWidth          uint32                `json:"photo_width,omitempty"`
    PhotoHeight         uint32                `json:"photo_height,omitempty"`
    Title               string                `json:"title,omitempty"`
    Description         string                `json:"description,omitempty"`
    Caption             string                `json:"caption,omitempty"`
    ParseMode           string                `json:"parse_mode,omitempty"`
    ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
    InputMessageContent *InputMessageContent  `json:"input_message_content,omitempty"`
}

type InlineQueryResultGif struct {
    Type                string                `json:"type"`
    ID                  string                `json:"id"`
    GifURL              string                `json:"gif_url"`
    GifWidth            uint32                `json:"gif_width,omitempty"`
    GifHeight           uint32                `json:"gif_height,omitempty"`
    GifDuration         uint32                `json:"gif_duration,omitempty"`
    ThumbURL            string                `json:"thumb_url"`
    Title               string                `json:"title,omitempty"`
    Caption             string                `json:"caption,omitempty"`
    ParseMode           string                `json:"parse_mode,omitempty"`
    ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
    InputMessageContent *InputMessageContent  `json:"input_message_content,omitempty"`
}

type InlineQueryResultMpeg4Gif struct {
    Type                string                `json:"type"`
    ID                  string                `json:"id"`
    Mpeg4URL            string                `json:"mpeg4_url"`
    Mpeg4Width          uint32                `json:"mpeg4_width,omitempty"`
    Mpeg4Height         uint32                `json:"mpeg4_height,omitempty"`
    Mpeg4Duration       uint32                `json:"mpeg4_duration,omitempty"`
    ThumbURL            string                `json:"thumb_url"`
    Title               string                `json:"title,omitempty"`
    Caption             string                `json:"caption,omitempty"`
    ParseMode           string                `json:"parse_mode,omitempty"`
    ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
    InputMessageContent *InputMessageContent  `json:"input_message_content,omitempty"`
}

type InlineQueryResultVideo struct {
    Type                string                `json:"type"`
    ID                  string                `json:"id"`
    VideoURL            string                `json:"video_url"`
    MimeType            string                `json:"mime_type"`
    ThumbURL            string                `json:"thumb_url"`
    Title               string                `json:"title"`
    Caption             string                `json:"caption,omitempty"`
    ParseMode           string                `json:"parse_mode,omitempty"`
    VideoWidth          uint32                `json:"video_width,omitempty"`
    VideoHeight         uint32                `json:"video_height,omitempty"`
    VideoDuration       uint32                `json:"video_duration,omitempty"`
    Description         string                `json:"description,omitempty"`
    ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
    InputMessageContent *InputMessageContent  `json:"input_message_content,omitempty"`
}

type InlineQueryResultAudio struct {
    Type                string                `json:"type"`
    ID                  string                `json:"id"`
    AudioURL            string                `json:"audio_url"`
    Title               string                `json:"title"`
    Caption             string                `json:"caption,omitempty"`
    ParseMode           string                `json:"parse_mode,omitempty"`
    Performer           string                `json:"performer,omitempty"`
    AudioDuration       uint32                `json:"audio_duration,omitempty"`
    ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
    InputMessageContent *InputMessageContent  `json:"input_message_content,omitempty"`
}

type InlineQueryResultVoice struct {
    Type                string                `json:"type"`
    ID                  string                `json:"id"`
    VoiceURL            string                `json:"voice_url"`
    Title               string                `json:"title"`
    Caption             string                `json:"caption,omitempty"`
    ParseMode           string                `json:"parse_mode,omitempty"`
    VoiceDuration       uint32                `json:"voice_duration,omitempty"`
    ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
    InputMessageContent *InputMessageContent  `json:"input_message_content,omitempty"`
}

type InlineQueryResultDocument struct {
    Type                string                `json:"type"`
    ID                  string                `json:"id"`
    Title               string                `json:"title"`
    Caption             string                `json:"caption,omitempty"`
    ParseMode           string                `json:"parse_mode,omitempty"`
    DocumentURL         string                `json:"document_url"`
    MimeType            string                `json:"mime_type"`
    Description         string                `json:"description,omitempty"`
    ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
    InputMessageContent *InputMessageContent  `json:"input_message_content,omitempty"`
    ThumbURL            string                `json:"thumb_url,omitempty"`
    ThumbWidth          uint32                `json:"thumb_width,omitempty"`
    ThumbHeight         uint32                `json:"thumb_height,omitempty"`
}

type InlineQueryResultLocation struct {
    Type                string                `json:"type"`
    ID                  string                `json:"id"`
    Latitude            float64               `json:"latitude"`
    Longitude           float64               `json:"longitude"`
    Title               string                `json:"title"`
    LivePeriod          uint32                `json:"live_period,omitempty"`
    ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
    InputMessageContent *InputMessageContent  `json:"input_message_content,omitempty"`
    ThumbURL            string                `json:"thumb_url,omitempty"`
    ThumbWidth          uint32                `json:"thumb_width,omitempty"`
    ThumbHeight         uint32                `json:"thumb_height,omitempty"`
}

type InlineQueryResultVenue struct {
    Type                string                `json:"type"`
    ID                  string                `json:"id"`
    Latitude            float64               `json:"latitude"`
    Longitude           float64               `json:"longitude"`
    Title               string                `json:"title"`
    Address             string                `json:"address"`
    FoursquareID        string                `json:"foursquare_id,omitempty"`
    ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
    InputMessageContent *InputMessageContent  `json:"input_message_content,omitempty"`
    ThumbURL            string                `json:"thumb_url,omitempty"`
    ThumbWidth          uint32                `json:"thumb_width,omitempty"`
    ThumbHeight         uint32                `json:"thumb_height,omitempty"`
}

type InlineQueryResultContact struct {
    Type                string                `json:"type"`
    ID                  string                `json:"id"`
    PhoneNumber         string                `json:"phone_number"`
    FirstName           string                `json:"first_name"`
    LastName            string                `json:"last_name,omitempty"`
    ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
    InputMessageContent *InputMessageContent  `json:"input_message_content,omitempty"`
    ThumbURL            string                `json:"thumb_url,omitempty"`
    ThumbWidth          uint32                `json:"thumb_width,omitempty"`
    ThumbHeight         uint32                `json:"thumb_height,omitempty"`
}

type InlineQueryResultGame struct {
    Type          string                `json:"type"`
    ID            string                `json:"id"`
    GameShortName string                `json:"game_short_name"`
    ReplyMarkup   *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

type InlineQueryResultCachedPhoto struct {
    Type                string                `json:"type"`
    ID                  string                `json:"id"`
    PhotoFileID         string                `json:"photo_file_id"`
    Title               string                `json:"title,omitempty"`
    Description         string                `json:"description,omitempty"`
    Caption             string                `json:"caption,omitempty"`
    ParseMode           string                `json:"parse_mode,omitempty"`
    ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
    InputMessageContent *InputMessageContent  `json:"input_message_content,omitempty"`
}

type InlineQueryResultCachedGif struct {
    Type                string                `json:"type"`
    ID                  string                `json:"id"`
    GifFileID           string                `json:"gif_file_id"`
    Title               string                `json:"title,omitempty"`
    Caption             string                `json:"caption,omitempty"`
    ParseMode           string                `json:"parse_mode,omitempty"`
    ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
    InputMessageContent *InputMessageContent  `json:"input_message_content,omitempty"`
}

type InlineQueryResultCachedMpeg4Gif struct {
    Type                string                `json:"type"`
    ID                  string                `json:"id"`
    Mpeg4FileID         string                `json:"mpeg4_file_id"`
    Title               string                `json:"title,omitempty"`
    Caption             string                `json:"caption,omitempty"`
    ParseMode           string                `json:"parse_mode,omitempty"`
    ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
    InputMessageContent *InputMessageContent  `json:"input_message_content,omitempty"`
}

type InlineQueryResultCachedSticker struct {
    Type                string                `json:"type"`
    ID                  string                `json:"id"`
    StickerFileID       string                `json:"sticker_file_id"`
    ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
    InputMessageContent *InputMessageContent  `json:"input_message_content,omitempty"`
}

type InlineQueryResultCachedDocument struct {
    Type                string                `json:"type"`
    ID                  string                `json:"id"`
    Title               string                `json:"title"`
    DocumentFileID      string                `json:"document_file_id"`
    Description         string                `json:"description,omitempty"`
    Caption             string                `json:"caption,omitempty"`
    ParseMode           string                `json:"parse_mode,omitempty"`
    ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
    InputMessageContent *InputMessageContent  `json:"input_message_content,omitempty"`
}

type InlineQueryResultCachedVideo struct {
    Type                string                `json:"type"`
    ID                  string                `json:"id"`
    VideoFileID         string                `json:"video_file_id"`
    Title               string                `json:"title"`
    Description         string                `json:"description,omitempty"`
    Caption             string                `json:"caption,omitempty"`
    ParseMode           string                `json:"parse_mode,omitempty"`
    ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
    InputMessageContent *InputMessageContent  `json:"input_message_content,omitempty"`
}

type InlineQueryResultCachedVoice struct {
    Type                string                `json:"type"`
    ID                  string                `json:"id"`
    VoiceFileID         string                `json:"voice_file_id"`
    Title               string                `json:"title"`
    Caption             string                `json:"caption,omitempty"`
    ParseMode           string                `json:"parse_mode,omitempty"`
    ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
    InputMessageContent *InputMessageContent  `json:"input_message_content,omitempty"`
}

type InlineQueryResultCachedAudio struct {
    Type                string                `json:"type"`
    ID                  string                `json:"id"`
    AudioFileID         string                `json:"audio_file_id"`
    Caption             string                `json:"caption,omitempty"`
    ParseMode           string                `json:"parse_mode,omitempty"`
    ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
    InputMessageContent *InputMessageContent  `json:"input_message_content,omitempty"`
}

type InputMessageContent interface{}

type InputTextMessageContent struct {
    MessageText           string `json:"message_text"`
    ParseMode             string `json:"parse_mode,omitempty"`
    DisableWebPagePreview bool   `json:"disable_web_page_preview,omitempty"`
}

type InputLocationMessageContent struct {
    Latitude   float64 `json:"latitude"`
    Longitude  float64 `json:"longitude"`
    LivePeriod uint32  `json:"live_period,omitempty"`
}

type InputVenueMessageContent struct {
    Latitude     float64 `json:"latitude"`
    Longitude    float64 `json:"longitude"`
    Title        string  `json:"title"`
    Address      string  `json:"address"`
    FoursquareID string  `json:"foursquare_id,omitempty"`
}

type InputContactMessageContent struct {
    PhoneNumber string `json:"phone_number"`
    FirstName   string `json:"first_name"`
    LastName    string `json:"last_name,omitempty"`
}

type ChosenInlineResult struct {
    ResultID        string    `json:"result_id"`
    From            User      `json:"from"`
    Location        *Location `json:"location,omitempty"`
    InlineMessageID string    `json:"inline_message_id,omitempty"`
    Query           string    `json:"query"`
}

type LabeledPrice struct {
    Label  string `json:"label"`
    Amount uint32 `json:"amount"`
}

type Invoice struct {
    Title          string `json:"title"`
    Description    string `json:"description"`
    StartParameter string `json:"start_parameter"`
    Currency       string `json:"currency"`
    TotalAmount    uint32 `json:"total_amount"`
}

type ShippingAddress struct {
    CountryCode string `json:"country_code"`
    State       string `json:"state"`
    City        string `json:"city"`
    StreetLine1 string `json:"street_line1"`
    StreetLine2 string `json:"street_line2"`
    PostCode    string `json:"post_code"`
}

type OrderInfo struct {
    Name            string           `json:"name,omitempty"`
    PhoneNumber     string           `json:"phone_number,omitempty"`
    Email           string           `json:"email,omitempty"`
    ShippingAddress *ShippingAddress `json:"shipping_address,omitempty"`
}

type ShippingOption struct {
    ID     string         `json:"id"`
    Title  string         `json:"title"`
    Prices []LabeledPrice `json:"prices"`
}

type SuccessfulPayment struct {
    Currency                string     `json:"currency"`
    TotalAmount             uint32     `json:"total_amount"`
    InvoicePayload          string     `json:"invoice_payload"`
    ShippingOptionID        string     `json:"shipping_option_id,omitempty"`
    OrderInfo               *OrderInfo `json:"order_info,omitempty"`
    TelegramPaymentChargeID string     `json:"telegram_payment_charge_id"`
    ProviderPaymentChargeID string     `json:"provider_payment_charge_id"`
}

type ShippingQuery struct {
    ID              string          `json:"id"`
    From            User            `json:"from"`
    InvoicePayload  string          `json:"invoice_payload"`
    ShippingAddress ShippingAddress `json:"shipping_address"`
}

type PreCheckoutQuery struct {
    ID               string     `json:"id"`
    From             User       `json:"from"`
    Currency         string     `json:"currency"`
    TotalAmount      uint32     `json:"total_amount"`
    InvoicePayload   string     `json:"invoice_payload"`
    ShippingOptionID string     `json:"shipping_option_id,omitempty"`
    OrderInfo        *OrderInfo `json:"order_info,omitempty"`
}

type PassportData struct {
    Data        []EncryptedPassportElement `json:"data"`
    Credentials EncryptedCredentials       `json:"credentials"`
}

type PassportFile struct {
    FileID   string `json:"file_id"`
    FileSize uint32 `json:"file_size"`
    FileDate uint32 `json:"file_date"`
}

type EncryptedPassportElement struct {
    Type        string         `json:"type"`
    Data        string         `json:"data,omitempty"`
    PhoneNumber string         `json:"phone_number,omitempty"`
    Email       string         `json:"email,omitempty"`
    Files       []PassportFile `json:"files,omitempty"`
    FrontSide   PassportFile   `json:"front_side,omitempty"`
    ReverseSide PassportFile   `json:"reverse_side,omitempty"`
    Selfie      PassportFile   `json:"selfie,omitempty"`
    Translation []PassportFile `json:"translation,omitempty"`
    Hash        string         `json:"hash"`
}

type EncryptedCredentials struct {
    Data   string `json:"data"`
    Hash   string `json:"hash"`
    Secret string `json:"secret"`
}

type PassportElementError interface{}

type PassportElementErrorDataField struct {
    Source    string `json:"source"`
    Type      string `json:"type"`
    FieldName string `json:"field_name"`
    DataHash  string `json:"data_hash"`
    Message   string `json:"message"`
}

type PassportElementErrorFrontSide struct {
    Source   string `json:"source"`
    Type     string `json:"type"`
    FileHash string `json:"file_hash"`
    Message  string `json:"message"`
}

type PassportElementErrorReverseSide struct {
    Source   string `json:"source"`
    Type     string `json:"type"`
    FileHash string `json:"file_hash"`
    Message  string `json:"message"`
}

type PassportElementErrorSelfie struct {
    Source   string `json:"source"`
    Type     string `json:"type"`
    FileHash string `json:"file_hash"`
    Message  string `json:"message"`
}

type PassportElementErrorFile struct {
    Source   string `json:"source"`
    Type     string `json:"type"`
    FileHash string `json:"file_hash"`
    Message  string `json:"message"`
}

type PassportElementErrorFiles struct {
    Source     string   `json:"source"`
    Type       string   `json:"type"`
    FileHashes []string `json:"file_hashes"`
    Message    string   `json:"message"`
}

type PassportElementErrorTranslationFile struct {
    Source   string `json:"source"`
    Type     string `json:"type"`
    FileHash string `json:"file_hash"`
    Message  string `json:"message"`
}

type PassportElementErrorTranslationFiles struct {
    Source     string   `json:"source"`
    Type       string   `json:"type"`
    FileHashes []string `json:"file_hashes"`
    Message    string   `json:"message"`
}

type PassportElementErrorUnspecified struct {
    Source      string `json:"source"`
    Type        string `json:"type"`
    ElementHash string `json:"element_hash"`
    Message     string `json:"message"`
}

type Game struct {
    Title        string           `json:"title"`
    Description  string           `json:"description"`
    Photo        *[]PhotoSize     `json:"photo"`
    Text         string           `json:"text,omitempty"`
    TextEntities *[]MessageEntity `json:"text_entities,omitempty"`
    Animation    *Animation       `json:"animation,omitempty"`
}

type Animation struct {
    FileID   string     `json:"file_id"`
    Thumb    *PhotoSize `json:"thumb,omitempty"`
    FileName string     `json:"file_name,omitempty"`
    MimeType string     `json:"mime_type,omitempty"`
    FileSize uint32     `json:"file_size,omitempty"`
}

type CallbackGame struct {
    // Placeholder
}

type GameHighScore struct {
    Position uint32 `json:"position"`
    User     User   `json:"user"`
    Score    uint32 `json:"score"`
}
