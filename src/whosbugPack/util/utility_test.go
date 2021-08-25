package util

import (
	"fmt"
	"testing"
)

func TestGenerateKIV(t *testing.T) {
	projectId := []byte("new_Project")
	key := []byte("password")
	K, IV := GenerateKIV(
		projectId,
		key,
	)
	if fmt.Sprintf("%x", K) != "93a3ad5bb81b0ce909b7e3b8df1a98853a462a8e5b064f2fe52f9cda640b97e0" || fmt.Sprintf("%x", IV) != "8dcb87e90d20cb6c9169891286741741" {
		t.Errorf("Generate K=%x, \tIV=%x\nexpected K=%s, \tIV=%s", string(K), string(IV), "93a3ad5bb81b0ce909b7e3b8df1a98853a462a8e5b064f2fe52f9cda640b97e0",
			"8dcb87e90d20cb6c9169891286741741")
	}
}

func TestEncryptAndDecrypt(t *testing.T) {
	projectId := "new_Project"
	key := "password"
	var cryptTests = []struct {
		input    string
		expected string
	}{
		{"abc12", "abc12"},
		{"1abasd12", "1abasd12"},
		{"@34sd", "@34sd"},
		{"^^&*)", "^^&*)"},
	}

	for _, testItem := range cryptTests {
		actual := Encrypt(projectId, key, testItem.input)
		actual = Decrypt(projectId, key, actual)
		if actual != testItem.expected {
			t.Errorf("Crypt Wrong!")
		}
	}
}
