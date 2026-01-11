# twitch-irc
"twitch-irc" is a lightweight Go Library that establishes a connection with Twitch IRC over TCP and sends structured data to a Go channel. It simplifies reading real-time messages from Twitch Chat.

For detailed field info, see the [pkg.go.dev documentation](https://pkg.go.dev/github.com/v0idzzy/twitch-irc).
# Features
- Connects to Twitch IRC over TCP
- Parses chat messages into structured Go types
- Supports all metadata including emotes, badges, etc.
# Installation 
```bash
go get github.com/v0idzzy/twitch-irc
```
# Example
```go
package main

import (
    "fmt"
    "log"

    "github.com/v0idzzy/twitch-irc"
)

func main() {
    oauth := "YOUR_TWITCH_OAUTH"       // must have chat:read perm (do not include OAuth2:)
    username := "YOUR_TWITCH_USERNAME" // name of the account linked to your OAuth
    channel := "CHANNEL_NAME"           // name of the channel you want to read messages from

    chatMessages := make(chan twitchirc.ChatMessage) // ChatMessage channel

    if err := twitchirc.Start(chatMessages, oauth, username, channel); err != nil {
        log.Fatal("IRC Error: ", err)
    }

    for message := range chatMessages {
        fmt.Printf("%s: %s \n", message.Username, message.Message)
        fmt.Println("Tags: ", message.Tags) // prints Tags struct
    }
}
```
