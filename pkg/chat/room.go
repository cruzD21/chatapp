package chat

type Room struct {
	ID         string
	members    map[*Client]bool
	register   chan *Client
	unregister chan *Client
}
