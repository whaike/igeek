package pkg

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"io"
)

func FileMd5(b []byte) (string, error) {
	f := &bytes.Buffer{}
	_, err := f.Read(b)
	if err != nil {
		return "", err
	}
	hash := md5.New()
	_, err = io.Copy(hash, f)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(hash.Sum(nil)), nil
}
