package message

// DataMap is an helper for interacting with map type message data
type DataMap map[string]interface{}

// Get returns a value from the map by key
func (d DataMap) Get(key string) (v interface{}, ok bool) {
	v, ok = d[key]
	return
}

// GetString returns a value from the map by key as a string
func (d DataMap) GetString(key string) (v string, ok bool) {
	if iv, iok := d.Get(key); iok {
		v, ok = iv.(string)
	}
	return
}

// GetBool returns a value from the map by key as a bool
func (d DataMap) GetBool(key string) (v, ok bool) {
	if iv, iok := d.Get(key); iok {
		v, ok = iv.(bool)
	}
	return
}

// GetInt returns a value from the map by key as a int
func (d DataMap) GetInt(key string) (v int, ok bool) {
	fv, ok := d.GetFloat64(key)
	if ok {
		v = int(fv)
	}
	return
}

// GetFloat64 returns a value from the map by key as a float64
func (d DataMap) GetFloat64(key string) (v float64, ok bool) {
	if iv, iok := d.Get(key); iok {
		v, ok = iv.(float64)
	}
	return
}
