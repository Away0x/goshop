package viewmodels

// ViewModel response face
type ViewModel interface {
	// Serialize 序列化方法
	Serialize() map[string]interface{}
}
