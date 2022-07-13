package generate_helper

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
	"github.com/go-basic/uuid"
	"io"
	"strings"
)

//表ID
func GenerateUUID() string {
	return strings.ReplaceAll(uuid.New(), "-", "")
}

//sha1
func Sha1Mix(str string) string {
	sha1 := sha1.New()
	sha1.Write([]byte(str))
	return hex.EncodeToString(sha1.Sum([]byte("")))
}

func Md5Mix(str string) string {
	md5 := md5.New()
	md5.Write([]byte(str))
	return hex.EncodeToString(md5.Sum([]byte("")))
}

//sha256
func Sha256(v string) string {
	h := sha256.New()
	io.WriteString(h, v)
	return hex.EncodeToString(h.Sum(nil))
}
