package main

import (
    "log"
    "net/http"
    "code.google.com/p/go.net/websocket"
)

func RugServer(ws *websocket.Conn) {
    var message string
    world := createWorld()

    log.Print("Launched worker\n")

    for {
        err := websocket.Message.Receive(ws, &message)
        if err != nil {
            log.Print("Receive error - stopping worker: ", err)
            break
        }

        err = websocket.JSON.Send(ws, world.Update(message))
        if err != nil {
            log.Print("Send error - stopping worker: ", err)
            break
        }
    }
}

func main() {
    http.Handle("/ws", websocket.Handler(RugServer))
    err := http.ListenAndServe(":44561", nil)
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}
