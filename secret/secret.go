package secret

import (
	"encoding/hex"
	"errors"
	"fmt"
	"log"

	"github.com/ptinsley/selfdestruct/kv"
	"github.com/ptinsley/selfdestruct/storage"
)

// Create - Create a lookup in kv and store the resulting encrypted data
func Create(key string, value string) error {
	encryptionKey := NewEncryptionKey()

	//Encrypt the data
	cypherText, err := Encrypt([]byte(value), encryptionKey)
	if err != nil {
		fmt.Println(fmt.Sprintf("[secret.Create] Unable to encrypt (%s)", err))
		return err
	}

	//store the encryption key in the kv store uuid -> encryptionkey
	err = kv.Set(key, encode(encryptionKey[:]))
	if err != nil {
		fmt.Println(fmt.Sprintf("secret.Create] Unable to store key in kv (%s)", err))
		return err
	}

	//store the encrypted data in cloud storage uuid -> cyphertext
	err = storage.Set(key, cypherText)
	if err != nil {
		fmt.Println(fmt.Sprintf("secret.Create] Unable to store encrypted data in cloud storage (%s)", err))
		return err
	}

	return nil
}

// Take - Retrieve and delete encryption key and encrypted secret
func Take(key string) (string, error) {
	//FIXME:
	encryptionString := kv.Get(key)
	if encryptionString == "" {
		fmt.Println(fmt.Sprintf("[secret.Take] couldn't retrieve encryption key for %s", key))
		return "", errors.New("kv retrieval failed")
	}

	var encryptionKey [32]byte
	copy(encryptionKey[:], decode(encryptionString))

	//FIXME:
	cypherText := storage.Get(key)
	if len(cypherText) == 0 {
		fmt.Println(fmt.Sprintf("[secret.Take] couldn't retrieve cyphertext for %s", key))
		return "", errors.New("storage retrieval failed")
	}

	plaintext, err := Decrypt(cypherText, &encryptionKey)
	if err != nil {
		fmt.Println(fmt.Sprintf("[secret.Take] couldn't decrypt cyphertext for %s (%s)", key, err))
		return "", err
	}

	return string(plaintext), nil
}

func encode(bytes []byte) string {

	dst := make([]byte, hex.EncodedLen(len(bytes)))
	hex.Encode(dst, bytes)

	return string(dst)
}

func decode(encodedString string) []byte {
	src := []byte(encodedString)

	dst := make([]byte, hex.DecodedLen(len(src)))
	_, err := hex.Decode(dst, src)
	if err != nil {
		log.Fatal(err)
	}

	return dst
}
