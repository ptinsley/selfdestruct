package storage

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"

	"cloud.google.com/go/storage"
	"github.com/ptinsley/selfdestruct/utils"
)

var bucket *storage.BucketHandle

// Init - setup storage client
func Init() {
	ctx := context.Background()

	client, err := storage.NewClient(ctx)
	if err != nil {
		fmt.Println(fmt.Sprintf("Failed to create Cloud Storage Client (%s)", err))
		os.Exit(-1)
	}

	bucket = client.Bucket(utils.Env("SECRET_BUCKET"))
}

// SetString - value passed in as a string vs []byte
func SetString(objectName string, objectValue string) error {
	return Set(objectName, []byte(objectValue))
}

// Set - Set an object value in cloud storage
func Set(objectName string, objectValue []byte) error {
	ctx := context.Background()
	object := bucket.Object(objectName)

	writer := object.NewWriter(ctx)

	if _, err := writer.Write(objectValue); err != nil {
		fmt.Println(fmt.Sprintf("Couldn't set '%s' [Write] (%s)", objectName, err))
		return err
	}

	if err := writer.Close(); err != nil {
		fmt.Println(fmt.Sprintf("Couldn't set '%s' [Close] (%s)", objectName, err))
		return err
	}

	return nil
}

// GetString - value returned as a string vs []byte
func GetString(objectName string) string {
	return string(Get(objectName))
}

// Get - Retrieve an object value from cloud storage
func Get(objectName string) []byte {
	ctx := context.Background()
	object := bucket.Object(objectName)

	reader, err := object.NewReader(ctx)
	if err != nil {
		fmt.Println(fmt.Sprintf("Couldn't retrieve '%s' [NewReader] (%s)", objectName, err))
		return []byte{}
	}
	defer reader.Close()

	objectValue, err := ioutil.ReadAll(reader)
	if err != nil {
		fmt.Println(fmt.Sprintf("Couldn't retrieve '%s' [ReadAll] (%s)", objectName, err))
		return []byte{}
	}

	return objectValue
}

// Delete - Delete an object from cloud storage
func Delete(objectName string) error {
	ctx := context.Background()
	object := bucket.Object(objectName)

	if err := object.Delete(ctx); err != nil {
		fmt.Println(fmt.Sprintf("Couldn't delete '%s' (%s)", objectName, err))
		return err
	}

	return nil
}

// TakeString - Retrieve and Delete object from cloud storage
// (returned as a string vs []byte)
func TakeString(objectName string) string {
	return string(Take(objectName))
}

// Take - Retrieve and Delete an object from cloud storage
func Take(objectName string) []byte {
	value := Get(objectName)

	if len(value) > 0 {
		Delete(objectName)
	}

	return value
}
