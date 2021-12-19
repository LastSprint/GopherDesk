package Utils

import (
	"crypto/hmac"
	"encoding/base64"
	"hash"
)

func ValidatePayload(payload []byte, secret, digest, url string, hash func() hash.Hash) bool {
	str := string(payload) + url

	var encoded []byte

	base64.StdEncoding.Encode([]byte(str), encoded)

	crp := hmac.New(hash, []byte(secret))

	crp.Write(encoded)

	result := string(crp.Sum(nil))

	return result == digest
}
