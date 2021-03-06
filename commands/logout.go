package commands

import (
	"github.com/emersion/go-imap"
)

// A LOGOUT command.
// See RFC 3501 section 6.1.3
type Logout struct{}

func (c *Logout) Command() *imap.Command {
	return &imap.Command{
		Name: imap.Logout,
	}
}

func (c *Logout) Parse(fields []interface{}) error {
	return nil
}
