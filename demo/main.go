package main

import (
	"fmt"
	"github.com/mh-cbon/mdl-go-components/demo/handlers"
	"github.com/mh-cbon/guess-path"
	"log"
	"net/http"
	"os/exec"
	"runtime"
	"time"
)

func main() {

	handlers.SetRoutes()

	fs := http.FileServer(http.Dir(GuessPath()))
	http.Handle("/static/", http.StripPrefix("/", fs))

	go httpServe(":8080")
	go openbrowser("http://localhost:8080/")
	<-make(chan bool)
}

var Demopath = ""

func GuessPath() string {
	p := guesspath.Path(Demopath, "")
	if p == "" {
		panic("path not matched")
	}
	return p
}

func httpServe(addr string) {
	// configure the http server
	srv := &http.Server{
		Addr:         addr,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	fmt.Printf("Server running on %v\n", srv.Addr)

	// Start the HTTP server
	log.Fatal(srv.ListenAndServe())
}

func openbrowser(url string) {
	var err error

	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}
	if err != nil {
		log.Fatal(err)
	}
}
