/*
2024 Moopinger
*/

package lib

import (
	"os"
	"strings"

	"net/url"
)

type Payload struct {
	HeaderName  []byte
	HeaderValue []byte
	Name        string
	Response    string
}

func ReadPayloadsFile(fileName string) ([]string, error) {
	content, err := os.ReadFile(fileName)
	if err != nil {
		return nil, err
	}

	unifiedContent := strings.ReplaceAll(string(content), "\r\n", "\n")
	lines := strings.Split(unifiedContent, "\n")

	return lines, nil
}

func WritePayloadsToFile(payloads []Payload, hostname string, fileName string) error {
	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	for _, payload := range payloads {
		_, err := file.WriteString("Payload: " + payload.Name + "\nHost: " + hostname + "\n\n")
		if err != nil {
			return err
		}
	}

	return nil
}

func ImportSinglePayloads(payloadString string) (Payload, error) {
	parts := strings.Split(payloadString, "; ")

	if len(parts) != 2 {
		return Payload{}, nil
	}

	return NewPayload(parts[0], parts[1], payloadString), nil
}

func BulkImportPayloads(payloadsFromFile []string, hostname string) ([]Payload, error) {

	payloads := []Payload{}

	for _, line := range payloadsFromFile {
		if line == "" {
			continue
		}

		//replace [HOSTNAME] with the actual hostname
		line = strings.ReplaceAll(line, "[HOSTNAME]", hostname)

		parts := strings.Split(line, "; ")

		if len(parts) != 2 {
			continue
		}

		payloads = append(payloads, NewPayload(parts[0], parts[1], line))
	}

	return payloads, nil
}

func NewPayload(headerName string, headerValue string, name string) Payload {

	// match and replace
	headerName = strings.ReplaceAll(headerName, "\\r", "\r")
	headerName = strings.ReplaceAll(headerName, "\\n", "\n")
	headerName = strings.ReplaceAll(headerName, "\\t", "\t")

	headerValue = strings.ReplaceAll(headerValue, "\\r", "\r")
	headerValue = strings.ReplaceAll(headerValue, "\\n", "\n")
	headerValue = strings.ReplaceAll(headerValue, "\\t", "\t")

	// Replace URL encoded values (%20) in the string, and replace with actual byte (0x20)
	headerName, _ = url.PathUnescape(headerName)
	headerValue, _ = url.PathUnescape(headerValue)

	return Payload{
		HeaderName:  []byte(headerName),
		HeaderValue: []byte(headerValue),
		Name:        name,
		Response:    "",
	}
}

func ChunkPayloads(payloads []Payload, routines int) [][]Payload {
	var divided [][]Payload

	size := len(payloads) / routines
	for i := 0; i < routines; i++ {
		start := i * size
		end := start + size

		// For the last routine, add the leftover payloads
		if i == routines-1 {
			end = len(payloads)
		}

		divided = append(divided, payloads[start:end])
	}

	return divided
}
