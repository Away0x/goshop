package viewmodel

type ViewModelSerialize map[string]interface{}

// ViewModel response face
type ViewModel interface {
	// Serialize 序列化方法
	Serialize() ViewModelSerialize
}

func (v *ViewModelSerialize) Include(key string, val interface{}) *ViewModelSerialize {
	if vm, ok := val.(ViewModel); ok {
		val = vm.Serialize()
	}

	(*v)[key] = val
	return v
}
