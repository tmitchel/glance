package glance

import (
	"net/http"
	"time"

	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
)

const (
	// Time allowed to write a message to the peer.
	writeWait = 10 * time.Second

	// Time allowed to read the next pong message from the peer.
	pongWait = 60 * time.Second

	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10

	// Maximum message size allowed from peer.
	maxMessageSize = 512
)

type chathub struct {
	clients    map[*client]bool
	broadcast  chan string
	register   chan *client
	unregister chan *client
}

// NewChathub creates a chathub to handle client Websocket
// connections and broadcasting messages.
func NewChathub() *chathub {
	return &chathub{
		clients:    make(map[*client]bool),
		broadcast:  make(chan string),
		register:   make(chan *client),
		unregister: make(chan *client),
	}
}

func (c *chathub) HandleConnection() http.HandlerFunc {
	var upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin:     func(r *http.Request) bool { return true },
	}

	return func(w http.ResponseWriter, r *http.Request) {
		logrus.Infof("Connecting")
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			logrus.Fatalf("unable to upgrade connection %v", err)
		}

		cl := &client{
			conn: conn,
			send: make(chan string),
			hub:  c,
		}

		c.register <- cl

		go cl.writePump()
	}
}

func (h *chathub) Run() {
	for {
		select {
		case client := <-h.register:
			h.clients[client] = true
		case client := <-h.unregister:
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				close(client.send)
			}
		case message := <-h.broadcast:
			for client := range h.clients {
				select {
				case client.send <- message:
				default:
					close(client.send)
					delete(h.clients, client)
				}
			}
		}
	}
}

type client struct {
	conn *websocket.Conn
	send chan string
	hub  *chathub
}

// writePump listens for the chathub to broadcast a message then
// sends it out on the Websocket connection.
func (c *client) writePump() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.conn.Close()
	}()

	for {
		select {
		case message, ok := <-c.send:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				c.conn.WriteMessage(websocket.CloseMessage, []byte{})
			}

			if err := c.conn.WriteMessage(websocket.TextMessage, []byte(message)); err != nil {
				logrus.Errorf("Error writing to websocket %v", err)
			}
		case <-ticker.C:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				logrus.Errorf("Error writing ping message. %v", err)
				return
			}
		}
	}
}
