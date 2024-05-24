package main

import (
	"DesignSphere_Server/src/server"
	"DesignSphere_Server/src/utils"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	SetupSignalHandler()
	server.StartServer()
}

// Setup up signal handler to handler process interrupts
//
// TODO: Gracefully handle the resource state and othe data
func SetupSignalHandler() {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		sig := <-sigs
		utils.Log(utils.INFO, "Received signal: "+sig.String())
		signal.Reset(syscall.SIGINT)
	}()
}
