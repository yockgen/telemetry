package main

// Client is the client to interact with machine service.
type Client struct {
	token string
	url   string
}

type Node struct {
	ID    string
	IP    string
	State string
}

type Filter struct {
	CPU  string
	RAM  string
	Disk string
}

// NewClient returns client to interact with machine service
// with token for auth.
func NewClient(token, url string) *Client {
	return &Client{token, url}
}

func (c *Client) AcquireNode(filter Filter) (Node, error) {
	// acquire node from machine service based on filter
	return Node{}, nil
}

func (c *Client) GetNodeByID(nodeID string) (Node, error) {
	// lookup node by id from machine service
	return Node{}, nil
}

func (c *Client) Action(nodeID, action string) error {
	// do action on node by id from machine service
	// poweron / poweroff
	return nil
}

func (c *Client) Execute(nodeID, cmd string) error {
	// execute cmd on node by id from machine service
	return nil
}

func (c *Client) ReleaseNode(nodeID string) error {
	// release node by id to machine service
	return nil
}
