package util

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"strconv"
)

func IsNumeric(s string) (int64, bool) {
	result, err := strconv.ParseInt(s, 10, 64)
	return result, err == nil
}

func Md5(str string) string {
	h := md5.New()
	io.WriteString(h, str)

	return hex.EncodeToString(h.Sum(nil))
}
