package kv

import (
	"context"
	"fmt"

	"cloud.google.com/go/datastore"
	"github.com/ptinsley/selfdestruct/utils"
)

var ds *datastore.Client

// Entity - K/V entity
type Entity struct {
	Value string
}

// Init - setup Datastore Client
func Init() {
	ctx := context.Background()
	var err error
	dsClient, err := datastore.NewClient(ctx, utils.Env("DATASTORE_PROJECT"))

	if err != nil {
		fmt.Println(err)
	} else {
		ds = dsClient
	}
}

// Set - set a k/v pair in datastore
func Set(key string, value string) error {
	ctx := context.Background()

	dsKey := datastore.NameKey("Entity", key, nil)
	entity := new(Entity)
	entity.Value = value

	if _, err := ds.Put(ctx, dsKey, entity); err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

// Get - Retrieve a value from a key in datastore
func Get(key string) string {
	ctx := context.Background()

	dsKey := datastore.NameKey("Entity", key, nil)
	entity := new(Entity)

	if err := ds.Get(ctx, dsKey, entity); err != nil {
		fmt.Println(fmt.Sprintf("Couldn't retrieve '%s' (%s)", key, err))
	}

	return entity.Value
}

// Delete - Delete a key from the k/v store
func Delete(key string) {
	ctx := context.Background()

	dsKey := datastore.NameKey("Entity", key, nil)
	if err := ds.Delete(ctx, dsKey); err != nil {
		fmt.Println(fmt.Sprintf("Couldn't delete '%s' (%s)", key, err))
	}
}

// Take - Retrieve and Delete a key from the k/v store
func Take(key string) string {
	value := Get(key)

	if value != "" {
		Delete(key)
	}

	return value
}
