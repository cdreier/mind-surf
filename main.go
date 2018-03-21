package main

import (
	"flag"
	"html/template"
	"log"
	"net/http"

	"github.com/gobuffalo/packr"
	"github.com/gorilla/websocket"
)

type holder struct {
	dev        bool
	connection *websocket.Conn
}

func main() {

	dev := flag.Bool("dev", false, "set dev to load from webpack instead of dist bundle")
	port := flag.String("port", "3000", "the port to start the server")
	flag.Parse()
	arg := flag.Args()

	if len(arg) != 1 {
		log.Fatalln("please give file")
	}

	file := arg[0]
	h := holder{
		dev: *dev,
	}

	go h.startFileWatcher(file)

	http.HandleFunc("/", h.index)
	http.HandleFunc("/ws", h.ws)

	distBox := packr.NewBox("./frontend/dist")
	http.Handle("/dist/", http.StripPrefix("/dist", http.FileServer(distBox)))

	log.Println("starting webserver on port", *port)

	http.ListenAndServe(":"+*port, nil)
}

func (h *holder) index(w http.ResponseWriter, r *http.Request) {
	box := packr.NewBox("./www")
	t, _ := template.New("index").Parse(box.String("index.html"))

	data := map[string]string{
		"devServer": "",
	}
	if h.dev {
		data["devServer"] = "http://localhost:8080"
	}

	t.Execute(w, data)
}
