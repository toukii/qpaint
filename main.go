package main

import (
	common "./common"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"

	ezlog "git.ezbuy.me/ezbuy/base/misc/log"
	"github.com/toukii/bezier"
	"github.com/zserge/lorca"
)

func main() {
	ui, err := lorca.New("", "", 1024, 650)
	if err != nil {
		log.Fatal(err)
	}

	_ = common.FS

	ui.Bind("bezierPath", bezierPath)

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

type Point []struct {
	X int64 `json:x`
	Y int64 `json:y`
}

// type Points []*Point
type Points []*bezier.Point

func bezierPath(i interface{}) string {
	bs, err := json.Marshal(i)
	ezlog.Infof("err:%+v", err)
	ezlog.Infof("%s", bs)
	var v Points
	err = json.Unmarshal(bs, &v)
	ezlog.Infof("err:%+v", err)
	ezlog.JSON(v)
	pathBs := bezier.Trhs(2, v...)
	ezlog.Infof("path: %s", pathBs)
	return string(pathBs)
}
