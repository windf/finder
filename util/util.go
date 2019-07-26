package util

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"mime/multipart"
	"os"
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

func UploadFile(file *multipart.FileHeader) (path string, err error) {
	src, err := file.Open()
	if err != nil {
		return
	}
	defer src.Close()

	imgPath := "/static/img"

	path = imgPath + file.Filename

	// Destination
	dst, err := os.Create(path)
	if err != nil {
		return
	}
	defer dst.Close()

	// Copy
	if _, err = io.Copy(dst, src); err != nil {
		return
	}
	return
}
