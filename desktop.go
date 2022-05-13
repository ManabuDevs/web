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

func LoginBounce(user string, pass string, url string) string {
	client := &http.Client{}
	u := notificationmaildomain.User{User: user, Password: pass}
	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(u)
	req, err := http.NewRequest("POST", url+"login", b)
	if err != nil {
		//os.Exit(1)
		fmt.Println("el error al generar req: ", err)

	}

	req.Header = http.Header{
		"Content-Type": []string{"application/json"},
	}

	resp, err := client.Do(req)
	if err != nil {
		//os.Exit(1)
		fmt.Println("el error al generar req: ", err)
	}
	defer resp.Body.Close()

	var data notificationmaildomain.EmailSendApiLoginResponse
	json.NewDecoder(resp.Body).Decode(&data)
	fmt.Println(data.Token)

	return data.Token
}