package integration

import (
	"crypto/rand"
	"encoding/hex"
)

func randstr(l int) string {
    b := make([]byte, l)
    rand.Read(b)
    return hex.EncodeToString(b)[:l]
}
