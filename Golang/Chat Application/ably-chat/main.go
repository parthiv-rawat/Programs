package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/ably/ably-go/ably"
)

func main() {
	fmt.Println("Type your clientID")
	reader := bufio.NewReader(os.Stdin)
	clientID, _ := reader.ReadString('\n')
	clientID = strings.Replace(clientID, "\n", "", -1)

	// Connect to Ably using the API key and ClientID specified above
	client, err := ably.NewRealtime(
		// If you have an Ably account, you can find
		// your API key at https://www.ably.io/accounts/any/apps/any/app_keys
		ably.WithKey("4bDMtQ.4jXJ8w:3_UmbUYSdvTJUUFIk4xP73S-wbWojcaYMI4g8T4bC9s"),
		// ably.WithEchoMessages(false), // Uncomment to stop messages you send from being sent back
		ably.WithClientID(clientID))
	if err != nil {
		panic(err)
	}
}
