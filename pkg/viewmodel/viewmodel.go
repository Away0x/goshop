package viewmodel

type Serialize map[string]interface{}

// ViewModel response face
type ViewModel interface {
	// Serialize 序列化方法
	Serialize() Serialize
}
