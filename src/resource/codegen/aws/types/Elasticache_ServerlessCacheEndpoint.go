package types

type Elasticache_ServerlessCacheEndpoint struct {
	// The DNS hostname of the cache node.
	Address string `json:"address,omitempty" yaml:"address,omitempty"`

	// The port number that the cache engine is listening on. Set as integer.
	Port int `json:"port,omitempty" yaml:"port,omitempty"`
}
