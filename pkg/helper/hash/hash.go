package hash

import (
	"crypto/md5"
	"crypto/sha1"
	"fmt"

	hashers "github.com/meehow/go-django-hashers"
)

// MD5 MD5哈希值
func MD5(b []byte) string {
	h := md5.New()
	_, _ = h.Write(b)
	return fmt.Sprintf("%x", h.Sum(nil))
}

// MD5String MD5哈希值
func MD5String(s string) string {
	return MD5([]byte(s))
}

// SHA1 SHA1哈希值
func SHA1(b []byte) string {
	h := sha1.New()
	_, _ = h.Write(b)
	return fmt.Sprintf("%x", h.Sum(nil))
}

// SHA1String SHA1哈希值
func SHA1String(s string) string {
	return SHA1([]byte(s))
}

func CheckPassword(cleartext, encoded string) (bool, error) {
	ok, err := hashers.CheckPassword(cleartext, encoded)
	return ok, err
}

// django like password
func MakePassword(cleartext string) (string, error) {
	encoded, err := hashers.MakePassword(cleartext)
	return encoded, err
}
