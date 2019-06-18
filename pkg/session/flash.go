package session

import (
	sess "echo_shop/pkg/session/echo-session"

	"github.com/labstack/echo/v4"
)

const (
	// DefaultFlashName -
	DefaultFlashName = "_flash"
)

type flash struct {
	Name    string
	Session sess.Session
}

func NewFlash(c echo.Context, name string) *flash {
	if name == "" {
		name = DefaultFlashName
	}

	s := sess.Default(c)
	return &flash{
		Name:    name,
		Session: s,
	}
}

func (f *flash) Add(val string) {
	f.Session.AddFlash(val, f.Name)
}

func (f *flash) Save() {
	f.Session.Save()
}

func (f *flash) Read() interface{} {
	if f.Session == nil {
		return nil
	}

	flashes := f.Session.Flashes()
	if len(flashes) == 0 {
		return nil
	}

	// v := flashes[0].(string)
	// f.Save()
	// return v
	return flashes
}
