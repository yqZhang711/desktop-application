package main

import (
	"embed"
	"fmt"
	"github.com/mpc-desktop-application/server"
	"github.com/mpc-desktop-application/view"
)

var (
	//go:embed static
	Static embed.FS

	//go:embed templates
	Templates embed.FS
)

func main() {

	// Obtain an available port
	port := server.GetAvailablePort()
	ginServer := server.NewGinServer(port, Static, Templates)
	fmt.Println(port)

	// Run gin
	go ginServer.Run()

	// Run WebView
	view.WebView(port)
}
