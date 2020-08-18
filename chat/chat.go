package chat

import (
	"crypto/rand"
	"encoding/json"
	"fmt"

	"github.com/gorilla/websocket"
)

// A Client will connect to the Olivia's server using the Token the Information and the Locale
type Client struct {
	Information *map[string]interface{}
	Locale      string
	Token       string
	Connection  *websocket.Conn
	Channel     chan string
}

// RequestMessage is the structure that uses entry connections to chat with the websocket
type RequestMessage struct {
	Type        int                    `json:"type"` // 0 for handshakes and 1 for messages
	Content     string                 `json:"content"`
	Token       string                 `json:"user_token"`
	Information map[string]interface{} `json:"information"`
	Locale      string                 `json:"locale"`
}

// ResponseMessage is the structure used to reply to the user through the websocket
type ResponseMessage struct {
	Content     string                 `json:"content"`
	Tag         string                 `json:"tag"`
	Information map[string]interface{} `json:"information"`
	Actions     []string               `json:"actions"`
}

// NewClient creates a new Client by generating a random token, and setting english as the
// default langauge.
// You need to give a pointer of the information map of your client.
// The host is also required with a boolean, if the SSL certificate is required.
func NewClient(host string, ssl bool, information *map[string]interface{}) (client Client, err error) {
	// Initialite the scheme and the url
	scheme := "ws"
	if ssl {
		scheme += "s"
	}

	url := fmt.Sprintf("%s://%s/websocket", scheme, host)
	fmt.Println(url)

	// Create the websocket connection
	connection, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		return
	}

	client = Client{
		Information: information,
		Locale:      "en",
		Token:       generateToken(),
		Connection:  connection,
		Channel:     make(chan string),
	}

	// Execute the handshake
	err = client.handshake()

	return
}

// Close closes the websocket connection
func (client *Client) Close() {
	client.Connection.Close()
}

// SendMessage sends the given content to the websocket of Olivia and returns her response with
// an error.
func (client *Client) SendMessage(content string) (ResponseMessage, error) {
	// Marshal the RequestMessage with type 1
	message := RequestMessage{
		Type:        1,
		Content:     content,
		Token:       client.Token,
		Information: *client.Information,
		Locale:      client.Locale,
	}

	namefn := func(s string) []byte {
		b, _ := json.Marshal(s)
		return b
	}

	if btest, err2 := json.Marshal(message.Type); err2 == nil {
		fmt.Println("Type :", namefn("type"), btest)
	}
	if btest, err2 := json.Marshal(message.Content); err2 == nil {
		fmt.Println("Content :", namefn("content"), btest)
	}
	if btest, err2 := json.Marshal(message.Token); err2 == nil {
		fmt.Println("Token :", namefn("user_token"), btest)
	}
	if btest, err2 := json.Marshal(message.Information); err2 == nil {
		fmt.Println("Information :", namefn("information"), btest)
	}
	if btest, err2 := json.Marshal(message.Locale); err2 == nil {
		fmt.Println("Locale :", namefn("locale"), btest)
	}
	bytes, err := json.Marshal(message)
	if err != nil {
		return ResponseMessage{}, err
	}
	fmt.Println("bytes :", bytes)

	// Write the message to the websocket
	if err = client.Connection.WriteMessage(websocket.TextMessage, bytes); err != nil {
		return ResponseMessage{}, err
	}

	_, bytes, err = client.Connection.ReadMessage()
	if err != nil {
		return ResponseMessage{}, err
	}

	// Unmarshal the json content of the message
	var response ResponseMessage
	var res2 map[string]interface{}
	json.Unmarshal(bytes, &res2)
	fmt.Println("res :", res2)
	if err = json.Unmarshal(bytes, &response); err != nil {
		return ResponseMessage{}, err
	}

	return response, nil
}

// handshake performs the handshake request to the websocket
func (client *Client) handshake() error {
	// Marshal the RequestMessage with type 0
	bytes, err := json.Marshal(RequestMessage{
		Type:        0,
		Content:     "",
		Token:       client.Token,
		Information: *client.Information,
	})
	if err != nil {
		return err
	}

	// Write the message to the websocket
	if err = client.Connection.WriteMessage(websocket.TextMessage, bytes); err != nil {
		return err
	}

	return nil
}

// generateToken returns a random token of 50 characters
func generateToken() string {
	b := make([]byte, 50)
	rand.Read(b)

	return fmt.Sprintf("%x", b)
}
