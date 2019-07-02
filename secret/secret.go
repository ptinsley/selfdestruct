package secret

import (
	"encoding/hex"
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/ptinsley/selfdestruct/kv"
	"github.com/ptinsley/selfdestruct/storage"
	"github.com/ptinsley/selfdestruct/utils"
)

var masterKey [32]byte

// Init - store master key
func Init() {
	masterKeyString := utils.Env("MASTERKEY")
	blah := NewEncryptionKey()
	fmt.Println(encode(blah[:]))

	if masterKeyString == "" {
		fmt.Println(fmt.Sprintf("Master key not provided, cannot continue"))
		os.Exit(-1)
	}

	copy(masterKey[:], decode(masterKeyString))
}

// Create - Create a lookup in kv and store the resulting encrypted data
func Create(value string) (secretID string, err error) {
	secretID = utils.UUID()

	encryptionKey := NewEncryptionKey()

	//Encrypt the data
	cypherText, err := Encrypt([]byte(value), encryptionKey)
	if err != nil {
		fmt.Println(fmt.Sprintf("[secret.Create] Unable to encrypt (%s)", err))
		return "", err
	}

	//store the encryption key in the kv store uuid -> encryptionkey
	//err = kv.Set(key, encode(encryptionKey[:]))
	encryptedKey, err := encryptKey(encryptionKey)
	if err != nil {
		fmt.Println(fmt.Sprintf("[secret.Create] Failed to encrypt data encryption key (%s)", err))
		return "", err
	}

	err = kv.Set(secretID, encryptedKey)
	if err != nil {
		fmt.Println(fmt.Sprintf("secret.Create] Unable to store key in kv (%s)", err))
		return "", err
	}

	//store the encrypted data in cloud storage uuid -> cyphertext
	err = storage.Set(secretID, cypherText)
	if err != nil {
		fmt.Println(fmt.Sprintf("secret.Create] Unable to store encrypted data in cloud storage (%s)", err))
		return "", err
	}

	return secretID, nil
}

// Take - Retrieve and delete encryption key and encrypted secret
func Take(key string) (string, error) {
	encryptionString := kv.Take(key)
	if encryptionString == "" {
		fmt.Println(fmt.Sprintf("[secret.Take] couldn't retrieve encryption key for %s", key))
		return "", errors.New("kv retrieval failed")
	}

	encryptionKey, err := decryptKey(encryptionString)
	if err != nil {
		fmt.Println(fmt.Sprintf("[secret.Take] unable to decrypt object encryption key (%s)", err))
		return "", err
	}

	// var encryptionKey [32]byte
	// copy(encryptionKey[:], decode(encryptionString))

	cypherText := storage.Take(key)
	if len(cypherText) == 0 {
		fmt.Println(fmt.Sprintf("[secret.Take] couldn't retrieve cyphertext for %s", key))
		return "", errors.New("storage retrieval failed")
	}

	plaintext, err := Decrypt(cypherText, encryptionKey)
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
