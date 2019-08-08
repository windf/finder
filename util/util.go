package util

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"strconv"
	"time"
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

	dir, err := os.Getwd()
	if err != nil {
		return
	}

	imgPath := "/static/img/"

	ext := filepath.Ext(file.Filename)

	path = imgPath + Md5(file.Filename+fmt.Sprintf("%d", time.Now().Unix())) + ext

	// Destination
	dst, err := os.Create(dir + path)
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
