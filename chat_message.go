package twitchirc

// ChatMessage is a struct of useful message information
//
//	Tags: Tags - Meta data associated with the message and user
//	Username: string - The chatter's username
//	Message: string - The actual message content
type ChatMessage struct {
	Tags     Tags   `json:"tags"`
	Username string `json:"username"`
	Message  string `json:"message"`
}

// Tags is a struct containing all possible metadata tags for a chat message
//
//	BadgeInfo: []Badge - Metadata related to chat badges (e.g. subscriber months)
//	Badges: []Badge - Chat badges in the format <badge>/<version>
//	Bits: string - Amount of Bits the user cheered
//	Color: string - User's name color in hex RGB format (e.g. #1E90FF)
//	DisplayName: string - The user’s display name
//	Emotes: []Emote - Emotes in the message with their position ranges
//	ID: string - Unique identifier for the message
//	Mod: bool - Whether the user is a moderator
//	ReplyParentMsgID: string - ID of the direct parent message being replied to
//	ReplyParentUserID: string - ID of the sender of the direct parent message
//	ReplyParentUserLogin: string - Login name of the direct parent sender
//	ReplyParentDisplayName: string - Display name of the direct parent sender
//	ReplyParentMsgBody: string - Body text of the parent message
//	ReplyThreadParentMsgID: string - ID of the top-level parent message
//	ReplyThreadParentUserLogin: string - Login name of top-level parent sender
//	RoomID: string - ID of the chat room (channel) where the message was sent
//	SourceBadges: []Badge - Badges from the source chat room (for shared chat)
//	SourceBadgeInfo: []Badge - Metadata related to source chat room badges
//	SourceID: string - ID of the source message (shared chat)
//	SourceOnly: bool - True if message was only sent to the source channel
//	SourceRoomID: string - ID of the chat room the message was sent from
//	Subscriber: bool - Whether the user is a subscriber
//	TmiSentTS: string - UNIX timestamp of when the message was sent
//	Turbo: bool - Whether user has commercial-free (Turbo) enabled
//	UserID: string - ID of the user sending the message
//	UserType: string - Type of user (admin, global_mod, staff, "")
//	Vip: bool - Whether the user is a VIP
//	ClientNonce: string - Unique identifier for verifying message from client
//	FirstMsg: string - Whether this is the user's first message ("1") or not
//	Flags: string - Internal system flags for the message (if any)
type Tags struct {
	BadgeInfo                  []Badge `json:"badge-info,omitempty"`
	Badges                     []Badge `json:"badges,omitempty"`
	Bits                       string  `json:"bits,omitempty"`
	Color                      string  `json:"color,omitempty"`
	DisplayName                string  `json:"display-name,omitempty"`
	Emotes                     []Emote `json:"emotes,omitempty"`
	ID                         string  `json:"id,omitempty"`
	Mod                        bool    `json:"mod,omitempty"`
	ReplyParentMsgID           string  `json:"reply-parent-msg-id,omitempty"`
	ReplyParentUserID          string  `json:"reply-parent-user-id,omitempty"`
	ReplyParentUserLogin       string  `json:"reply-parent-user-login,omitempty"`
	ReplyParentDisplayName     string  `json:"reply-parent-display-name,omitempty"`
	ReplyParentMsgBody         string  `json:"reply-parent-msg-body,omitempty"`
	ReplyThreadParentMsgID     string  `json:"reply-thread-parent-msg-id,omitempty"`
	ReplyThreadParentUserLogin string  `json:"reply-thread-parent-user-login,omitempty"`
	RoomID                     string  `json:"room-id,omitempty"`
	SourceBadges               []Badge `json:"source-badges,omitempty"`
	SourceBadgeInfo            []Badge `json:"source-badge-info,omitempty"`
	SourceID                   string  `json:"source-id,omitempty"`
	SourceOnly                 bool    `json:"source-only,omitempty"`
	SourceRoomID               string  `json:"source-room-id,omitempty"`
	Subscriber                 bool    `json:"subscriber,omitempty"`
	TmiSentTS                  string  `json:"tmi-sent-ts,omitempty"`
	Turbo                      bool    `json:"turbo,omitempty"`
	UserID                     string  `json:"user-id,omitempty"`
	UserType                   string  `json:"user-type,omitempty"`
	Vip                        bool    `json:"vip,omitempty"`
	ClientNonce                string  `json:"client-nonce,omitempty"`
	FirstMsg                   string  `json:"first-msg,omitempty"`
	Flags                      string  `json:"flags,omitempty"`
}

// Emote represents a Twitch emote and its position within a message
//
//	ID: string - The unique emote ID
//	RangeStart: int - Start index of the emote in the message
//	RangeEnd: int - End index of the emote in the message
type Emote struct {
	ID         string
	RangeStart int
	RangeEnd   int
}

// Badge represents a chat badge
//
//	Title: string - Type of badge (e.g. moderator, subscriber)
//	Data: string - Version or value associated with the badge
type Badge struct {
	Title string
	Data  string
}
