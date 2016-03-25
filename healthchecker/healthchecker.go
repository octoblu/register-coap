package healthchecker

import (
	"fmt"
	"log"
	"math/rand"

	"github.com/dustin/go-coap"
)

// Healthy returns true if the CoAP request
// succeeds, false otherwise
func Healthy(uriStr string) bool {
	messageID := uint16(rand.Int() / 2)

	request := coap.Message{
		Type:      coap.Confirmable,
		Code:      coap.GET,
		MessageID: messageID,
	}
	request.SetPathString("healthcheck")

	client, err := coap.Dial("udp", uriStr)
	if err != nil {
		log.Fatalln("Error on coap.Dial", err.Error())
	}

	response, err := client.Send(request)
	if err != nil {
		log.Fatalln("Error on client.Send", err.Error())
	}

	fmt.Println("response", response.Code)
	return response.Code == coap.Content
}
