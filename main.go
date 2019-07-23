package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"

	"github.com/toukii/qpaint/common"
	"github.com/zserge/lorca"
)

func main() {
	ui, err := lorca.New("", "", 1024, 650)
	if err != nil {
		log.Fatal(err)
	}

	_ = common.FS

	ui.Bind("bezierPath", common.BezierPath)
	ui.Bind("shapes", common.BuildShapes)

	defer ui.Close()

	ln, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		log.Fatal(err)
	}
	defer ln.Close()
	go http.Serve(ln, http.FileServer(http.Dir("./www")))
	// go http.Serve(ln, http.FileServer(common.FS))
	ui.Load(fmt.Sprintf("http://%s", ln.Addr()))

	sigc := make(chan os.Signal)
	signal.Notify(sigc, os.Interrupt)
	select {
	case <-sigc:
	case <-ui.Done():
	}
}
