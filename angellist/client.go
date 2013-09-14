package angellist

// New creates an instance of the AngelList Client
func New(token string) *Client {
	c := &Client{}
	c.Token = token
	c.Users = &UserResource{c}
	c.Startups = &StartupResource{c}
	return c
}

// New creates an instance of the AngelList Client
// using a Guest / Anonymous account.
func NewGuest() *Client {
	return New("")
}

type Client struct {
	Token string

	Users    *UserResource
	Startups *StartupResource
}

