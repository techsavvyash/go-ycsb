package stylusdb

import (
	"bufio"
	"fmt"
	"net"

	"github.com/google/uuid"
)

func writeToSocket() {
	// Connect to the server
	conn, err := net.Dial("tcp", "localhost:6767") //stylus db proxy runs on 6767
	if err != nil {
		fmt.Println("Error connecting:", err)
		return
	}
	defer conn.Close()

	// Generate payloads
	events := GenerateEvents(100000)

	// Convert payloads to JSON and send to the server
	for _, event := range events {
		newUUID := uuid.New()
		query := ConvertToQuery(newUUID.String(), event)
		_, err := conn.Write([]byte(query))
		if err != nil {
			fmt.Println("Error writing to socket:", err)
			return
		}
	}
	for {
		// Read response from the server
		// (Assuming the server sends a response back)
		response, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			fmt.Println("Error reading response:", err)
			return
		}

		fmt.Println("Response from server:", response)
	}
}
