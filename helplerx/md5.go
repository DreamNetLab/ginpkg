package helplerx

import (
	"crypto/md5" //nolint:gosec
	"encoding/hex"
)

func EncodeMD5(str string) string {
	m := md5.New() //nolint:gosec

	m.Write([]byte(str))

	return hex.EncodeToString(m.Sum(nil))
}
