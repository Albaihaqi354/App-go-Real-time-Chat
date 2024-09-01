package ws

import "log"

type Room struct {
	ID      string             `json:"id"`
	Name    string             `json:"name"`
	Clients map[string]*Client `json:"clients"`
}

type Hub struct {
	Rooms      map[string]*Room
	Register   chan *Client
	Unregister chan *Client
	Broadcast  chan *Message
}

func NewHub() *Hub {
	return &Hub{
		Rooms:      make(map[string]*Room),
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Broadcast:  make(chan *Message),
	}
}

func (h *Hub) Run() {
	for {
		select {
		case cl := <-h.Register:
			if _, ok := h.Rooms[cl.RoomID]; ok {
				r := h.Rooms[cl.RoomID]
				if _, ok := r.Clients[cl.ID]; !ok {
					r.Clients[cl.ID] = cl
					log.Printf("Client %s joined room %s", cl.ID, cl.RoomID)
				}
			} else {
				log.Printf("Room %s not found", cl.RoomID)
			}
		case cl := <-h.Unregister:
			if _, ok := h.Rooms[cl.RoomID]; ok {
				if _, ok := h.Rooms[cl.RoomID].Clients[cl.ID]; ok {
					delete(h.Rooms[cl.RoomID].Clients, cl.ID)
					close(cl.Message)

					if len(h.Rooms[cl.RoomID].Clients) > 0 {
						h.Broadcast <- &Message{
							Content:  "User left the chat",
							RoomID:   cl.RoomID,
							Username: cl.Username,
						}
					}
				}
			}
		case m := <-h.Broadcast:
			if room, ok := h.Rooms[m.RoomID]; ok {
				for _, cl := range room.Clients {
					log.Printf("Broadcating message to client %s in room %s", cl.ID, m.RoomID)
					cl.Message <- m
				}
			} else {
				log.Printf("Room %s not found for broadcasting", m.RoomID)
			}
		}
	}
}
