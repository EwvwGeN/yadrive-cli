package util

import (
	"crypto/md5"
	"encoding/hex"
)

func GetMD5Hash(input string) string {
	md5Hash := md5.Sum([]byte(input))
	return hex.EncodeToString(md5Hash[:])
}