package types

type Memcache_InstanceMemcacheParameters struct {
	/*
	   (Output)
	   This is a unique ID associated with this set of parameters.
	*/
	Id string `json:"id,omitempty" yaml:"id,omitempty"`

	// User-defined set of parameters to use in the memcache process.
	Params map[string]string `json:"params,omitempty" yaml:"params,omitempty"`
}
