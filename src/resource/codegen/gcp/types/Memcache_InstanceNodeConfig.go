package types

type Memcache_InstanceNodeConfig struct {
	// Number of CPUs per node.
	CpuCount int `json:"cpuCount,omitempty" yaml:"cpuCount,omitempty"`

	/*
	   Memory size in Mebibytes for each memcache node.

	   - - -
	*/
	MemorySizeMb int `json:"memorySizeMb,omitempty" yaml:"memorySizeMb,omitempty"`
}
