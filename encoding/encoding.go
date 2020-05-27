package encoding

import (
	"encoding/base64"
)

// DaASCIIABase64 prende una stringa in ascii e la trasforma in base64
func DaASCIIABase64(msg string) string {
	return base64.StdEncoding.EncodeToString([]byte(msg))
}
