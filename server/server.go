package server

import (
	"fmt"

	"github.com/ptinsley/selfdestruct/utils"
)

// Init - start up the server
func Init() {
	r := NewRouter()
	r.Run(fmt.Sprintf("%s:%s", utils.EnvOr("LISTEN", "0.0.0.0"), utils.EnvOr("PORT", "8080")))
}
