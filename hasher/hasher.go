package hasher

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

// EmailHasher takes an email string and returns it's MD5 hash value
func EmailHasher(email string) string {
	hasher := md5.New()                                // instantiate a new hasher from the md5 package
	data := []byte(email)                              // case the string to bytes
	hasher.Write(data)                                 // add the byte data to the write interface of the hasher
	hashedValue := hex.EncodeToString(hasher.Sum(nil)) // call the checksum method to computer the hash
	return hashedValue
}

// CreateURL takes the hashed email value returns the gravatar endpoint
func CreateURL(hashedValue string, size int32) string {
	gURL := fmt.Sprintf("https://www.gravatar.com/avatar/%v?s=%d", hashedValue, size)
	return gURL
}

// Gravatar calls the CreateURL() to generate the endpoint
func Gravatar(email string, size int32) string {
	hashedValue := EmailHasher(email)
	return CreateURL(hashedValue, size)
}
