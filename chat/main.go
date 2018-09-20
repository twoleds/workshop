package main

import (
	"github.com/gorilla/websocket"
	"io"
	"log"
	"net/http"
	"os"
)

type Message struct {
	Name    string `json:"name"`
	Content string `json:"content"`
}

type Chat struct {
	upgrader       websocket.Upgrader
	clients        map[chan Message]struct{}
	enterChannel   chan chan Message
	leaveChannel   chan chan Message
	messageChannel chan Message
}

func (c *Chat) Run() {
	for {
		select {
		case ch := <-c.enterChannel:
			c.clients[ch] = struct{}{}
		case ch := <-c.leaveChannel:
			delete(c.clients, ch)
			close(ch)
		case msg := <-c.messageChannel:
			for ch := range c.clients {
				ch <- msg
			}
		}
	}
}

func (c *Chat) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	conn, err := c.upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()

	channel := make(chan Message, 128)

	c.enterChannel <- channel
	defer func() {
		c.leaveChannel <- channel
	}()

	c.messageChannel <- Message{
		Name:    "SYSTEM",
		Content: "New user entered from " + r.RemoteAddr,
	}

	defer func() {
		c.messageChannel <- Message{
			Name:    "SYSTEM",
			Content: "User " + r.RemoteAddr + " left",
		}
	}()

	go func() {
		for msg := range channel {
			err := conn.WriteJSON(msg)
			if err != nil {
				log.Println(err)
			}
		}
	}()

	msg := Message{}

	for {

		err := conn.ReadJSON(&msg)
		if err != nil {
			log.Println(err)
			return
		}

		c.messageChannel <- msg

	}

}

// $ go get github.com/gorilla/websocket
func main() {

	chat := &Chat{
		upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		},
		clients:        make(map[chan Message]struct{}),
		enterChannel:   make(chan chan Message, 128),
		leaveChannel:   make(chan chan Message, 128),
		messageChannel: make(chan Message, 128),
	}
	go chat.Run()

	http.Handle("/socket", chat)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		f, err := os.Open("./index.html")
		if err != nil {
			log.Println(err)
			return
		}
		defer f.Close()

		io.Copy(w, f)

	})

	http.ListenAndServe(":8080", nil)

}
