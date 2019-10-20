package hasher

import (
	"crypto/md5"
	"fmt"
	"testing"
)

// Define testValues struct i.e. a fieldset of testValues
type testValues struct {
	gravatarEndpoint string
	email            string
	size             uint32
}

// Implement the testValues struct with actual testPairs
var testValuesFieldSet = []testValues{
	{"https://www.gravatar.com/avatar/cf38500a2cd3b6a2c8c1d4d8259e83f8?s=%v", "kamil@lelonek.me", 10},
}

// TestAverage function tets the Average function in mathPackage
func TestGravatar(t *testing.T) {
	for _, testValuesExtract := range testValuesFieldSet {
		returnedValue := Average(testValuesExtract.testSlice)
		if returnedValue != testValuesExtract.averageValue {
			t.Error(
				"For the test slice", testValuesExtract.testSlice,
				"The expected average was", testValuesExtract.averageValue,
				"But we instead got", returnedValue,
			)
		}
	}
}

// TestAverage function tets the Average function in mathPackage
func TestGravatar(t *testing.T) {
	for _, testValuesExtract := range testValuesFieldSet {
		returnedValue := Average(testValuesExtract.testSlice)
		if returnedValue != testValuesExtract.averageValue {
			t.Error(
				"For the test slice", testValuesExtract.testSlice,
				"The expected average was", testValuesExtract.averageValue,
				"But we instead got", returnedValue,
			)
		}
	}
}

// TestAverage function tets the Average function in mathPackage
func TestGravatar(t *testing.T) {
	for _, testValuesExtract := range testValuesFieldSet {
		returnedValue := Average(testValuesExtract.testSlice)
		if returnedValue != testValuesExtract.averageValue {
			t.Error(
				"For the test slice", testValuesExtract.testSlice,
				"The expected average was", testValuesExtract.averageValue,
				"But we instead got", returnedValue,
			)
		}
	}
}

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
