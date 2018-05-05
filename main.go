package main

import (
	"log"
	"net/http"

	"github.com/googollee/go-socket.io"
)

func main() {
	server, err := socketio.NewServer(nil)

	if err != nil {
		log.Fatal(err)
	}
	server.On("connection", func(so socketio.Socket) {
		log.Println("on connection, max: ", server.GetMaxConnection())
		so.Join("chat")
		so.On("chat message", func(msg string) {
			log.Println("Receive: " + so.Id() + " -> " + msg)
			// so.BroadcastTo("chat", "chat message", msg)
			err := so.Emit("chat message", msg)

			if err != nil {
				log.Println("Error Emitting:", err)
			}
		})
		// Socket.io acknowledgement example
		// The return type may vary depending on whether you will return
		// For this example it is "string" type
		so.On("chat message with ack", func(msg string) string {
			return msg
		})
		so.On("disconnection", func() {
			log.Println("on disconnect")
		})
	})
	server.On("error", func(so socketio.Socket, err error) {
		log.Println("error:", err)
	})

	http.Handle("/socket.io/", server)
	http.Handle("/", http.FileServer(http.Dir("./asset")))
	log.Println("Serving at localhost:5000...")
	log.Fatal(http.ListenAndServe(":5000", nil))
}
