package healthchecker

import (
	"log"
	"math/rand"
	"net/url"

	"github.com/dustin/go-coap"
)

// Healthy returns true if the CoAP request
// succeeds, false otherwise
func Healthy(uriStr string) bool {
	uri, err := url.Parse(uriStr)
	fatalIfError("Error parsing uri", err)

	messageID := uint16(rand.Int() / 2)

	request := coap.Message{
		Type:      coap.Confirmable,
		Code:      coap.GET,
		MessageID: messageID,
	}
	request.SetPathString("healthcheck")

	client, err := coap.Dial("udp", uri.Host)
	fatalIfError("Error on coap.Dial", err)

	response, err := client.Send(request)
	fatalIfError("Error on client.Send", err)

	return response.Code == coap.Content
}

func fatalIfError(msg string, err error) {
	if err == nil {
		return
	}

	log.Fatalln(msg, err.Error())
}
