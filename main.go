package main

import (
	"github.com/ptinsley/selfdestruct/kv"
	"github.com/ptinsley/selfdestruct/secret"
	"github.com/ptinsley/selfdestruct/server"
	"github.com/ptinsley/selfdestruct/storage"
)

func main() {
	secret.Init()
	kv.Init()
	storage.Init()
	server.Init()
}
