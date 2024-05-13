package helplerx

import (
	"crypto/md5" //nolint:gosec
	"crypto/sha256"
	"encoding/hex"
)

func EncodeMD5(str string) string {
	m := md5.New() //nolint:gosec

	m.Write([]byte(str))

	return hex.EncodeToString(m.Sum(nil))
}

func EncodeSha256(raw string) string {
	encoder := sha256.New()

	encoder.Write([]byte(raw))

	return hex.EncodeToString(encoder.Sum(nil))
}
