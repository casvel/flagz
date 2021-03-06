package main

type hub struct {
	// Registered connections.
	connections map[*connection]string

	// Inbound messages from the connections.
	broadcast chan []byte

	// Register requests from the connections.
	register chan *connection

	// Unregister requests from connections.
	unregister chan *connection
}

func newHub() *hub {
	return &hub{
		broadcast:   make(chan []byte),
		register:    make(chan *connection),
		unregister:  make(chan *connection),
		connections: make(map[*connection]string),
	}
}

func (h *hub) run() {
	for {
		select {
		//case c := <-h.register:
			//h.connections[c] = true			
		case c := <-h.unregister:
			if _, ok := h.connections[c]; ok {
				delete(connPlayer, h.connections[c])
				delete(h.connections, c)
				close(c.send)
			}
		case m := <-h.broadcast:
			for c := range h.connections {
				select {
				case c.send <- m:
				default:
					delete(h.connections, c)
					close(c.send)
				}
			}
		}
	}
}
