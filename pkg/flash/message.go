// Package flash : message 提示信息
package flash

import (
	"github.com/Away0x/flash"
	"github.com/labstack/echo/v4"
)

const (
	flashMessageKeyName = "_flash_message"

	// FlashMessageSuccessType -
	FlashMessageSuccessType = "success"
	// FlashMessageInfoType -
	FlashMessageInfoType = "info"
	// FlashMessageWarningType -
	FlashMessageWarningType = "warning"
	// FlashMessageDangerType -
	FlashMessageDangerType = "danger"
)

// MessageFlash -
type MessageFlash struct {
	EchoContext echo.Context
	Data        flash.DataValue
}

// NewMessageFlash -
func NewMessageFlash(c echo.Context) *MessageFlash {
	return &MessageFlash{
		EchoContext: c,
		Data:        make(flash.DataValue),
	}
}

func (m *MessageFlash) add(typeName string, msg string) {
	if m.Data[typeName] == nil {
		m.Data[typeName] = make([]string, 0)
	}
	m.Data[typeName] = append(m.Data[typeName], msg)
}

func (m *MessageFlash) Save() {
	flash.NewFlash(flashMessageKeyName, m.EchoContext.Request()).
		Set(m.Data).
		Save(m.EchoContext.Response())
}

func (m *MessageFlash) Read() flash.DataValue {
	return flash.NewFlash(flashMessageKeyName, m.EchoContext.Request()).
		Read(m.EchoContext.Response())
}

func (m *MessageFlash) AddSuccess(msg string) *MessageFlash {
	m.add(FlashMessageSuccessType, msg)
	return m
}

func (m *MessageFlash) AddInfo(msg string) *MessageFlash {
	m.add(FlashMessageInfoType, msg)
	return m
}

func (m *MessageFlash) AddWarning(msg string) *MessageFlash {
	m.add(FlashMessageWarningType, msg)
	return m
}

func (m *MessageFlash) AddDanger(msg string) *MessageFlash {
	m.add(FlashMessageDangerType, msg)
	return m
}

func (m *MessageFlash) Add(key, msg string) *MessageFlash {
	m.add(key, msg)
	return m
}

// NewSuccessMessage -
func NewSuccessMessage(c echo.Context, msg string) {
	NewMessageFlash(c).AddSuccess(msg).Save()
}

// NewInfoMessage -
func NewInfoMessage(c echo.Context, msg string) {
	NewMessageFlash(c).AddInfo(msg).Save()
}

// NewWarningMessage -
func NewWarningMessage(c echo.Context, msg string) {
	NewMessageFlash(c).AddWarning(msg).Save()
}

// NewDangerMessage -
func NewDangerMessage(c echo.Context, msg string) {
	NewMessageFlash(c).AddDanger(msg).Save()
}

// NewMessage -
func NewMessage(c echo.Context, key, msg string) {
	NewMessageFlash(c).Add(key, msg).Save()
}
