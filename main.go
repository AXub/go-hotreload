package gohotreload

import (
	"context"
	"net/http"

	"nhooyr.io/websocket"
)

func WsConnection(w http.ResponseWriter, r *http.Request) {
	_, err := websocket.Accept(w, r, nil)
	if err != nil {
		http.Error(w, "error", http.StatusInternalServerError)
		return
	}

	_, cancel := context.WithCancel(r.Context())
	defer cancel()

	msgChan := make(chan interface{})

	go func() {
		defer close(msgChan)
		msgChan <- struct{}{}
	}()

	<-msgChan
}
