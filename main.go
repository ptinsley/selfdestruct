package main

import (
	"fmt"

	"github.com/ptinsley/selfdestruct/kv"
	"github.com/ptinsley/selfdestruct/server"
	"github.com/ptinsley/selfdestruct/storage"
	"github.com/ptinsley/selfdestruct/utils"
)

func main() {
	kv.Init()
	storage.Init()
	florp := utils.Env("FLORP")
	if florp == "" {
		fmt.Println("blank")
	} else {
		fmt.Println(florp)
	}
	server.Init()
}
