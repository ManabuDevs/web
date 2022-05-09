/*package main

import (
	"os"
	"runtime"

	"github.com/alexflint/gallium"
)

func main() {
	runtime.LockOSThread()         // must be the first statement in main - see below
	gallium.Loop(os.Args, onReady) // must be called from main function
}

func onReady(app *gallium.App) {
	app.OpenWindow("http:/localhost:9000/", gallium.FramedWindow)
}*/

//go:generate fileb0x b0x.yml
package main

import (
	"github.com/ImVexed/muon"

	//"cra-go/webfiles"
	"net/http"
)

func main() {
	fileHandler := http.FileServer(http.Dir("/"))

	cfg := &muon.Config{
		Title:      "Hello, World!",
		Height:     500,
		Width:      500,
		Titled:     true,
		Resizeable: true,
	}

	m := muon.New(cfg, fileHandler)

	//m.Bind("add", add)

	if err := m.Start(); err != nil {
		panic(err)
	}
}
