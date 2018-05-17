package session

import (
	"log"
	"math/rand"
	"strconv"
	"testing"
)

func TestSum(t *testing.T) {
	CIPHER_KEY := []byte("0123456789012345")
	msg := strconv.Itoa(rand.Intn(100))

	if encrypted, err := encrypt(CIPHER_KEY, msg); err != nil {
		log.Println(err)
	} else {
		log.Printf("CIPHER KEY: %s\n", string(CIPHER_KEY))
		log.Printf("ENCRYPTED: %s\n", encrypted)

		if decrypted, err := decrypt(CIPHER_KEY, encrypted); err != nil {
			log.Println(err)
		} else {
			log.Printf("DECRYPTED: %s\n", decrypted)

			if decrypted != msg {
				t.Errorf("Encryption was incorrect")
			}
		}
	}

}
