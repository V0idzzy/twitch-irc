package twitchirc

// twitchkitIRC is meant to make it simple to conntect to twitch's IRC and read chat messages + tags

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"net"
	"strings"
)

// Start begins the IRC connection to Twitch chat and starts reading messages.
//
// Parameters:
//   - msgChan: A channel that receives ChatMessage objects as they are received.
//   - oauthToken: Your OAuth token with the "chat:read" permission.
//   - username: The bot's username.
//   - channel: The Twitch channel to connect to (without the #).
//
// It returns an error if the connection fails.
func Start(msgChan chan<- ChatMessage, oauthToken, username, channel string) error {
	conn, err := net.Dial("tcp", "irc.chat.twitch.tv:6667")
	if err != nil {
		return err
	}
	defer conn.Close()

	// Send authentication and join channel
	fmt.Fprintf(conn, "PASS oauth:%s\r\n", oauthToken)
	fmt.Fprintf(conn, "NICK %s\r\n", username)
	fmt.Fprintf(conn, "JOIN #%s\r\n", channel)
	fmt.Fprintf(conn, "CAP REQ :twitch.tv/tags\r\n")

	reader := bufio.NewReader(conn)

	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			log.Println("IRC read error: ", err)
			break
		}

		rawMessage := string(line)

		// handle a chatters message
		if strings.Contains(rawMessage, "PRIVMSG") {
			// Parse the irc message and send to the passed channel
			msgChan <- parser(rawMessage)
			continue
		}

		// NOTICE login failure
		if strings.Contains(rawMessage, "NOTICE") {
			ircMessage := strings.Split(rawMessage, ":")
			if len(ircMessage) >= 3 && cleanString(ircMessage[2]) == "login authentication failed" {
				return errors.New("Login authentication error")
			}
			continue
		}

		// Sucessful Login
		if strings.Contains(rawMessage, ":tmi.twitch.tv 001") {
			msgChan <- ChatMessage{
				Tags{},
				"twitch",
				"IRC Connected...",
			}
			continue
		}

		// Respond to PING to stay connected
		if strings.HasPrefix(rawMessage, "PING") {
			fmt.Fprintf(conn, "PONG :tmi.twitch.tv\r\n")
			continue
		}
	}
	return nil
}
