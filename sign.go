package mturk

import (
	"fmt"
	"crypto/hmac"
	"encoding/base64"
	"launchpad.net/goamz/aws"
)

var b64 = base64.StdEncoding

// ----------------------------------------------------------------------------
// Mechanical Turk signing (http://goo.gl/wrzfn)
func sign(auth aws.Auth, service, method, timestamp string, params map[string] string) {
	payload := service + method + timestamp
	hash := hmac.NewSHA1([]byte(auth.SecretKey))
	hash.Write([]byte(payload))
	signature := make([]byte, b64.EncodedLen(hash.Size()))
	b64.Encode(signature, hash.Sum())

	params["Signature"] = string(signature)
}
