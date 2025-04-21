package main

import (
	"context"
	"github.com/eeQuillibrium/wobble/pkg/utils"
)

func main() {
	globalCtx, cancel := context.WithCancel(context.Background())

	pool := initDB(globalCtx)

	srv := initServer(pool, NewFiberApp(NewFiberTemplates()))

	go srv.Run(globalCtx)

	utils.Shutdown(cancel)
}
