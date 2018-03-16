package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true }, // bad hack to accept all websockets during development
}

func (h *holder) ws(w http.ResponseWriter, r *http.Request) {
	connection, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("could not start dev-websocket:", err)
		return
	}
	defer connection.Close()

	h.connection = connection

	for {
		// mt, message, err := connection.ReadMessage()
		_, message, err := connection.ReadMessage()
		if err != nil {
			// log.Println("read:", err)
			break
		}
		log.Printf("recv: %s", message)
		// err = connection.WriteMessage(mt, message)
		// if err != nil {
		// 	log.Println("write:", err)
		// 	break
		// }
	}
}

func (h *holder) startFileWatcher(dir string) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	eventChan := make(chan string)
	go debounce(time.Second, eventChan, func(name string) {
		if h.connection != nil {
			fmt.Println("found a change. sending to browser...", name)

			content, _ := ioutil.ReadFile(name)
			data := parse(string(content))
			h.connection.WriteJSON(data)
		}
	})

	done := make(chan bool)
	go func() {
		for {
			select {
			case event := <-watcher.Events:
				eventChan <- event.Name
			case err := <-watcher.Errors:
				log.Println("error:", err)
			}
		}
	}()

	err = watcher.Add(dir)
	if err != nil {
		log.Fatal(err)
	}
	<-done
}

func debounce(interval time.Duration, input chan string, cb func(arg string)) {
	var item string
	timer := time.NewTimer(interval)
	for {
		select {
		case item = <-input:
			timer.Reset(interval)
		case <-timer.C:
			if item != "" {
				cb(item)
			}
		}
	}
}
