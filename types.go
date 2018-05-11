package telegram

import (
    "encoding/json"
)

type Response struct {
    Ok          bool               `json:"ok"`
    Description string             `json:"description"`
    Result      json.RawMessage    `json:"result"`
    ErrorCode   int                `json:"error_code"`
    Parameters  ResponseParameters `json:"parameters"`
}

type ResponseParameters struct {
    MigrateToChatID int64 `json:"migrate_to_chat_id"`
    RetryAfter      int   `json:"retry_after"`
}

type User struct {
    ID           int    `json:"id"`
    IsBot        bool   `json:"is_bot"`
    FirstName    string `json:"first_name"`
    LastName     string `json:"last_name"`
    Username     string `json:"username"`
    LanguageCode string `json:"language_code"`
}

type Chat struct {
    ID                          int64      `json:"id"`
    Type                        string     `json:"type"`
    Title                       string     `json:"title"`
    Username                    string     `json:"username"`
    FirstName                   string     `json:"first_name"`
    LastName                    string     `json:"last_name"`
    AllMembersAreAdministrators bool       `json:"all_members_are_administrators"`
    Photo                       *ChatPhoto `json:"photo"`
    Description                 string     `json:"description"`
    InviteLink                  string     `json:"invite_link"`
    PinnedMessage               *Message   `json:"pinned_message"`
    StickerSetName              string     `json:"sticker_set_name"`
    CanSetStickerSet            bool       `json:"can_set_sticker_set"`
}

type Message struct {
    ID                    int                `json:"message_id"`
    From                  *User              `json:"from"`
    Date                  int                `json:"date"`
    Chat                  Chat               `json:"chat"`
    ForwardFrom           *User              `json:"forward_from"`
    ForwardFromChat       *Chat              `json:"forward_from_chat"`
    ForwardFromMessageID  int                `json:"forward_from_message_id"`
    ForwardSignature      string             `json:"forward_signature"`
    ForwardDate           int                `json:"forward_date"`
    ReplyToMessage        *Message           `json:"reply_to_message"`
    EditDate              int                `json:"edit_date"`
    MediaGroupID          string             `json:"media_group_id"`
    AuthorSignature       string             `json:"author_signature"`
    Text                  string             `json:"text"`
    Entities              *[]MessageEntity   `json:"entities"`
    CaptionEntities       *[]MessageEntity   `json:"caption_entities"`
    Audio                 *Audio             `json:"audio"`
    Document              *Document          `json:"document"`
    Game                  *Game              `json:"game"`
    Photo                 *[]PhotoSize       `json:"photo"`
    Sticker               *Sticker           `json:"sticker"`
    Video                 *Video             `json:"video"`
    Voice                 *Voice             `json:"voice"`
    VideoNote             *VideoNote         `json:"video_note"`
    Caption               string             `json:"caption"`
    Contact               *Contact           `json:"contact"`
    Location              *Location          `json:"location"`
    Venue                 *Venue             `json:"venue"`
    NewChatMembers        *[]User            `json:"new_chat_members"`
    LeftChatMember        *User              `json:"left_chat_member"`
    NewChatTitle          string             `json:"new_chat_title"`
    NewChatPhoto          *[]PhotoSize       `json:"new_chat_photo"`
    DeleteChatPhoto       bool               `json:"delete_chat_photo"`
    GroupChatCreated      bool               `json:"group_chat_created"`
    SupergroupChatCreated bool               `json:"supergroup_chat_created"`
    ChannelChatCreated    bool               `json:"channel_chat_created"`
    MigrateToChatID       int                `json:"migrate_to_chat_id"`
    MigrateFromChatID     int                `json:"migrate_from_chat_id"`
    PinnedMessage         *Message           `json:"pinned_message"`
    Invoice               *Invoice           `json:"invoice"`
    SuccessfulPayment     *SuccessfulPayment `json:"successful_payment"`
    ConnectedWebsite      string             `json:"connected_website"`
}

type MessageEntity struct {
    Type   string `json:"type"`
    Offset int    `json:"offset"`
    Length int    `json:"length"`
    URL    string `json:"url"`
    User   *User  `json:"user"`
}

type PhotoSize struct {
    FileID   string `json:"file_id"`
    Width    int    `json:"width"`
    Height   int    `json:"height"`
    FileSize int    `json:"file_size"`
}

type Audio struct {
    FileID    string `json:"file_id"`
    Duration  int    `json:"duration"`
    Performer string `json:"performer"`
    Title     string `json:"title"`
    MimeType  string `json:"mime_type"`
    FileSize  int    `json:"file_size"`
}

type Document struct {
    FileID   string     `json:"file_id"`
    Thumb    *PhotoSize `json:"thumb"`
    FileName string     `json:"file_name"`
    MimeType string     `json:"mime_type"`
    FileSize int        `json:"file_size"`
}

type Video struct {
    FileID   string     `json:"file_id"`
    Width    int        `json:"width"`
    Height   int        `json:"height"`
    Duration int        `json:"duration"`
    Thumb    *PhotoSize `json:"thumb"`
    MimeType string     `json:"mime_type"`
    FileSize int        `json:"file_size"`
}

type Voice struct {
    FileID   string `json:"file_id"`
    Duration int    `json:"duration"`
    MimeType string `json:"mime_type"`
    FileSize int    `json:"file_size"`
}

type VideoNote struct {
    FileID   string     `json:"file_id"`
    Length   int        `json:"length"`
    Duration int        `json:"duration"`
    Thumb    *PhotoSize `json:"thumb"`
    FileSize int        `json:"file_size"`
}

type Contact struct {
    PhoneNumber string `json:"phone_number"`
    FirstName   string `json:"first_name"`
    LastName    string `json:"last_name"`
    UserID      int    `json:"user_id"`
}

type Location struct {
    Longitude float64 `json:"longitude"`
    Latitude  float64 `json:"latitude"`
}

type Venue struct {
    Location     *Location `json:"location"`
    Title        string    `json:"title"`
    Address      string    `json:"address"`
    FoursquareID string    `json:"foursquare_id"`
}

type UserProfilePhotos struct {
    TotalCount int            `json:"total_count"`
    Photos     *[][]PhotoSize `json:"photos"`
}

type File struct {
    FileID   string `json:"file_id"`
    FileSize int    `json:"file_size"`
    FilePath string `json:"file_path"`
}

type ReplyKeyboardMarkup struct {
    Keyboard        [][]KeyboardButton `json:"keyboard"`
    ResizeKeyboard  bool               `json:"resize_keyboard"`
    OneTimeKeyboard bool               `json:"one_time_keyboard"`
    Selective       bool               `json:"selective"`
}

type KeyboardButton struct {
    Text            string `json:"text"`
    RequestContact  bool   `json:"request_contact"`
    RequestLocation bool   `json:"request_location"`
}

type ReplyKeyboardRemove struct {
    RemoveKeyboard bool `json:"remove_keyboard"`
    Selective      bool `json:"selective"`
}

type InlineKeyboardMarkup struct {
    InlineKeyboard [][]InlineKeyboardButton `json:"inline_keyboard"`
}

type InlineKeyboardButton struct {
    Text                         string        `json:"text"`
    URL                          string        `json:"url"`
    CallbackData                 string        `json:"callback_data"`
    SwitchInlineQuery            string        `json:"switch_inline_query"`
    SwitchInlineQueryCurrentChat string        `json:"switch_inline_query_current_chat"`
    CallbackGame                 *CallbackGame `json:"callback_game"`
    Pay                          bool          `json:"pay"`
}

type CallbackQuery struct {
    ID              string   `json:"id"`
    From            User     `json:"from"`
    Message         *Message `json:"message"`
    InlineMessageID string   `json:"inline_message_id"`
    ChatInstance    string   `json:"chat_instance"`
    Data            string   `json:"data"`
    GameShortName   string   `json:"game_short_name"`
}

type ForceReply struct {
    ForceReply bool `json:"force_reply"`
    Selective  bool `json:"selective"`
}

type ChatPhoto struct {
    SmallFileID string `json:"small_file_id"`
    BigFileID   string `json:"big_file_id"`
}

type ChatMember struct {
    User                  User   `json:"user"`
    Status                string `json:"status"`
    UntilDate             int    `json:"until_date"`
    CanBeEdited           bool   `json:"can_be_edited"`
    CanChangeInfo         bool   `json:"can_change_info"`
    CanPostMessages       bool   `json:"can_post_messages"`
    CanEditMessages       bool   `json:"can_edit_messages"`
    CanDeleteMessages     bool   `json:"can_delete_messages"`
    CanInviteUsers        bool   `json:"can_invite_users"`
    CanRestrictMembers    bool   `json:"can_restrict_members"`
    CanPinMessages        bool   `json:"can_pin_messages"`
    CanPromoteMembers     bool   `json:"can_promote_members"`
    CanSendMessages       bool   `json:"can_send_messages"`
    CanSendMediaMessages  bool   `json:"can_send_media_messages"`
    CanSendOtherMessages  bool   `json:"can_send_other_messages"`
    CanAddWebPagePreviews bool   `json:"can_add_web_page_previews"`
}

type InputMedia interface{}

type InputMediaPhoto struct {
    InputMedia

    Type      string `json:"type"`
    Media     string `json:"media"`
    Caption   string `json:"caption"`
    ParseMode string `json:"parse_mode"`
}

type InputMediaVideo struct {
    Type              string `json:"type"`
    Media             string `json:"media"`
    Caption           string `json:"caption"`
    ParseMode         string `json:"parse_mode"`
    Width             int    `json:"width"`
    Height            int    `json:"height"`
    Duration          int    `json:"duration"`
    SupportsStreaming bool   `json:"supports_streaming"`
}

type InputFile interface{}

type Sticker struct {
    FileID       string        `json:"file_id"`
    Width        int           `json:"width"`
    Height       int           `json:"height"`
    Thumb        *PhotoSize    `json:"thumb"`
    Emoji        string        `json:"emoji"`
    SetName      string        `json:"set_name"`
    MaskPosition *MaskPosition `json:"mask_position"`
    FileSize     int           `json:"file_size"`
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
    Location *Location `json:"location"`
    Query    string    `json:"query"`
    Offset   string    `json:"offset"`
}

type InlineQueryResult interface{}

// ...

type LabeledPrice struct {
    Label  string `json:"label"`
    Amount int    `json:"amount"`
}

type Invoice struct {
    Title          string `json:"title"`
    Description    string `json:"description"`
    StartParameter string `json:"start_parameter"`
    Currency       string `json:"currency"`
    TotalAmount    int    `json:"total_amount"`
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
    Name            string           `json:"name"`
    PhoneNumber     string           `json:"phone_number"`
    Email           string           `json:"email"`
    ShippingAddress *ShippingAddress `json:"shipping_address"`
}

type ShippingOption struct {
    ID     string         `json:"id"`
    Title  string         `json:"title"`
    Prices []LabeledPrice `json:"prices"`
}

type SuccessfulPayment struct {
    Currency                string     `json:"currency"`
    TotalAmount             int        `json:"total_amount"`
    InvoicePayload          string     `json:"invoice_payload"`
    ShippingOptionID        string     `json:"shipping_option_id"`
    OrderInfo               *OrderInfo `json:"order_info"`
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
    TotalAmount      int        `json:"total_amount"`
    InvoicePayload   string     `json:"invoice_payload"`
    ShippingOptionID string     `json:"shipping_option_id"`
    OrderInfo        *OrderInfo `json:"order_info"`
}

type Game struct {
    Title        string           `json:"title"`
    Description  string           `json:"description"`
    Photo        *[]PhotoSize     `json:"photo"`
    Text         string           `json:"text"`
    TextEntities *[]MessageEntity `json:"text_entities"`
    Animation    *Animation       `json:"animation"`
}

type Animation struct {
    FileID   string     `json:"file_id"`
    Thumb    *PhotoSize `json:"thumb"`
    FileName string     `json:"file_name"`
    MimeType string     `json:"mime_type"`
    FileSize int        `json:"file_size"`
}

type CallbackGame struct {
    // Placeholder
}

type GameHighScore struct {
    Position int  `json:"position"`
    User     User `json:"user"`
    Score    int  `json:"score"`
}
