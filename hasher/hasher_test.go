package hasher

import (
	"reflect"
	"testing"
)

// Define testValues struct i.e. a fieldset of testValues
type testValues struct {
	gravatarEndpoint string
	email            string
	size             int32
	hashedEmail      string
}

// Implement the testValues struct with actual testPairs
var testValuesFieldSet = []testValues{
	{"https://www.gravatar.com/avatar/606e3533635c8146e6f5e11dcfcb95ac?s=10", "awokogbon@lana.ng", 10, "606e3533635c8146e6f5e11dcfcb95ac"},
}

// TestEmailHasher function tests the EmailHasher() function in hasher package
func TestEmailHasher(t *testing.T) {
	for _, testValuesExtract := range testValuesFieldSet {
		returnedValue := EmailHasher(testValuesExtract.email)
		if !reflect.DeepEqual(returnedValue, testValuesExtract.hashedEmail) {
			t.Error(
				"For the email", testValuesExtract.email,
				"The expected hash was", testValuesExtract.hashedEmail,
				"But we instead got", returnedValue,
			)
		}
	}
}

// TestCreateURL function tests the CreateURL() in hasher package
func TestCreateURL(t *testing.T) {
	for _, testValuesExtract := range testValuesFieldSet {
		returnedValue := CreateURL(testValuesExtract.hashedEmail, testValuesExtract.size)
		if !reflect.DeepEqual(returnedValue, testValuesExtract.gravatarEndpoint) {
			t.Error(
				"For the hashedEmail", testValuesExtract.hashedEmail,
				"and the size", testValuesExtract.size,
				"The expected gravatar endpoint was", testValuesExtract.gravatarEndpoint,
				"But we instead got", returnedValue,
			)
		}
	}
}

// TestGravatar function tests the Gravatar() in hasher package
func TestGravatar(t *testing.T) {
	for _, testValuesExtract := range testValuesFieldSet {
		returnedValue := Gravatar(testValuesExtract.email, testValuesExtract.size)
		if !reflect.DeepEqual(returnedValue, testValuesExtract.gravatarEndpoint) {
			t.Error(
				"For the hashedEmail", testValuesExtract.hashedEmail,
				"and the size", testValuesExtract.size,
				"The expected gravatar endpoint was", testValuesExtract.gravatarEndpoint,
				"But we instead got", returnedValue,
			)
		}
	}
}
