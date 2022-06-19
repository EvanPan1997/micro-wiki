package util

import "encoding/base64"

func EncodeToBase64(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}
