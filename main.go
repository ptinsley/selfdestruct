package main

import (
	"github.com/ptinsley/selfdestruct/kv"
	"github.com/ptinsley/selfdestruct/server"
	"github.com/ptinsley/selfdestruct/storage"
)

func main() {
	kv.Init()
	storage.Init()
	server.Init()
}
