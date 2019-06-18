package session

// const (
// 	// MessageFlashName -
// 	MessageFlashName = "message"
// )

// // MessageFlashData -
// type MessageFlashData struct {
// 	Data        map[string][]string
// 	EchoContext echo.Context
// }

// // NewMessage -
// func NewMessage(c echo.Context, data ...map[string][]string) *MessageFlashData {
// 	m := &MessageFlashData{
// 		EchoContext: c,
// 	}
// 	if len(data) != 0 {
// 		m.Data = data[0]
// 	}

// 	return m
// }

// // Save -
// func (m *MessageFlashData) Save() {
// 	val, err := json.Marshal(m.Data)
// 	if err != nil {
// 		return
// 	}

// 	f := NewFlashData(m.EchoContext, MessageFlashName, string(val))
// 	f.Save()
// }

// // Read -
// func (m *MessageFlashData) Read() map[string][]string {
// 	f := NewFlashData(m.EchoContext, MessageFlashName)
// 	val := f.Read()
// 	if val == "" {
// 		return nil
// 	}

// 	var data map[string][]string
// 	if err := json.Unmarshal([]byte(val), &data); err != nil {
// 		return nil
// 	}

// 	return data
// }
