package main

type UserClient struct {
	client *Client
	user_id int
}

type Hub struct {
	user_clients map[int]*Client
	clients map[*Client]bool

	login chan *UserClient
	broadcast chan []byte
	send chan []byte
	register chan *Client
	unregister chan *Client
}


func newHub() *Hub {
	return &Hub{
		login:      make(chan *UserClient),
		broadcast:  make(chan []byte),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		clients:    make(map[*Client]bool),
		user_clients: make(map[int]*Client),
	}
}

func (h *Hub) run() {	
	for {
		select {
		case client := <-h.register:
			h.clients[client] = true
		case client := <-h.unregister:
			if _, ok := h.clients[client]; ok {
				// If the user is logged in, log out
				for user_id, client_ptr := range h.user_clients {
					if client_ptr == client {
						delete(h.user_clients, user_id)
					}
				}
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
		case uc := <-h.login:
			h.user_clients[uc.user_id] = uc.client
		}
	}
}
