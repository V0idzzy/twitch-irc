package twitchirc

import (
	"strconv"
	"strings"
)

// parser parses an IRC string into a ChatMessage struct
//
// rawIRC: is the raw string recieved from IRC
//
// Returns a ChatMessage struct
func parser(rawIRC string) ChatMessage {
	msg := ChatMessage{}
	parts := strings.SplitN(rawIRC, " :", 2)
	if len(parts) != 2 {
		return msg
	}

	tagsString := parts[0]
	messageString := parts[1]

	//parse tags
	if strings.HasPrefix(tagsString, "@") {
		msg.Tags = tagsParser(tagsString)
	}

	//parse message
	messageParts := strings.SplitN(messageString, ":", 2)
	if len(parts) != 2 {
		return msg
	}
	msg.Username = strings.SplitN(messageParts[0], "!", 2)[0]
	msg.Message = cleanString(messageParts[1])

	return msg
}

// tagsParser parses an IRC tag string into a Tag struct.
//
// tagsString: is the raw string recieved from IRC
//
// Returns a Tag struct
func tagsParser(tagsString string) Tags {
	tags := Tags{}
	tagsString = strings.TrimPrefix(tagsString, "@")
	tagsString = strings.TrimSpace(tagsString)
	parts := strings.Split(tagsString, ";")

	tagMap := make(map[string]string)
	for _, part := range parts {
		kv := strings.Split(part, "=")
		key := kv[0]
		value := ""
		if len(kv) > 1 {
			value = kv[1]
		}
		tagMap[key] = value
	}

	tags.BadgeInfo = badgeParser(tagMap["badge-info"])
	tags.Badges = badgeParser(tagMap["badges"])
	tags.Bits = tagMap["bits"]
	tags.Color = tagMap["color"]
	tags.DisplayName = tagMap["display-name"]
	tags.Emotes = emoteParser(tagMap["emotes"])
	tags.ID = tagMap["id"]
	tags.Mod = tagMap["mod"] == "1"
	tags.ReplyParentMsgID = tagMap["reply-parent-msg-id"]
	tags.ReplyParentUserID = tagMap["reply-parent-user-id"]
	tags.ReplyParentUserLogin = tagMap["reply-parent-user-login"]
	tags.ReplyParentDisplayName = tagMap["reply-parent-display-name"]
	tags.ReplyParentMsgBody = tagMap["reply-parent-msg-body"]
	tags.ReplyThreadParentMsgID = tagMap["reply-thread-parent-msg-id"]
	tags.ReplyThreadParentUserLogin = tagMap["reply-thread-parent-user-login"]
	tags.RoomID = tagMap["room-id"]
	tags.SourceBadges = badgeParser(tagMap["source-badges"])
	tags.SourceBadgeInfo = badgeParser(tagMap["source-badge-info"])
	tags.SourceID = tagMap["source-id"]
	tags.SourceOnly = tagMap["source-only"] == "1"
	tags.SourceRoomID = tagMap["source-room-id"]
	tags.Subscriber = tagMap["subscriber"] == "1"
	tags.TmiSentTS = tagMap["tmi-sent-ts"]
	tags.Turbo = tagMap["turbo"] == "1"
	tags.UserID = tagMap["user-id"]
	tags.UserType = tagMap["user-type"]
	tags.Vip = tagMap["vip"] == "1"
	tags.ClientNonce = tagMap["client-nonce"]
	tags.FirstMsg = tagMap["first-msg"]
	tags.Flags = tagMap["flags"]

	return tags
}

// emoteParser parses an IRC emote string into a slice of Emote structs.
//
// emoteString: is the raw string received from IRC
//
// Returns a slice of Emote structs representing each emote and its position
func emoteParser(emoteString string) []Emote {
	var emotes []Emote

	if emoteString == "" {
		return emotes
	}
	for _, emoteData := range strings.Split(emoteString, "/") {
		rangeParts := strings.Split(strings.Split(emoteData, ":")[1], "-")
		range1, err := strconv.Atoi(rangeParts[0])
		if err != nil {
			return emotes
		}
		range2, err := strconv.Atoi(rangeParts[1])
		if err != nil {
			return emotes
		}
		emote := Emote{
			strings.Split(emoteData, ":")[0],
			range1,
			range2,
		}
		emotes = append(emotes, emote)
	}

	return emotes
}

// badgeParser parses a RawrIRC badge string into a slice of Badge structs
//
// badgeString: is the raw badge string recieved from IRC
//
// returns a slice of Badge structs
func badgeParser(badgeString string) []Badge {
	var badges []Badge
	if badgeString == "" {
		return badges
	}
	for _, badgeData := range strings.Split(badgeString, ",") {
		badge := Badge{
			strings.Split(badgeData, "/")[0],
			strings.Split(badgeData, "/")[1],
		}
		badges = append(badges, badge)
	}

	return badges
}

// cleanString cleans up the RawIRC string for easy string comparisons
//
// str: is the raw IRC string
//
// Returns a cleaned string for comparisons
func cleanString(str string) string {
	clean := strings.ToLower(str)
	clean = strings.TrimRight(clean, "\r\n")
	clean = strings.TrimSpace(clean)
	return clean
}
