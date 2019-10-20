package hasher

import (
	"crypto/md5"
	"fmt"
)

// gravatarHash() takes an email string and returns it's MD5 hash value
func EmailHasher(email string) []byte {
	hashedValue := md5.Sum([]byte(email)) // initialized hashedValue, because [:] does not work on unaddressable value
	return hashedValue[:]
}

func CreateURL(hashedValue []byte, size uint32) string {
	gURL := fmt.Sprintf("https://www.gravatar.com/avatar/%x?s=%d", hashedValue, size)
	return gURL
}

func Gravatar(email string, size uint32) string {
	hashedValue := EmailHasher(email)
	return CreateURL(hashedValue, size)
}
