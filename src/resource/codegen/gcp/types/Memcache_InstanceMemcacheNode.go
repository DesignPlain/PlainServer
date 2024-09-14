package types

type Memcache_InstanceMemcacheNode struct {
	/*
	   (Output)
	   Hostname or IP address of the Memcached node used by the clients to connect to the Memcached server on this node.
	*/
	Host string `json:"host,omitempty" yaml:"host,omitempty"`

	/*
	   (Output)
	   Identifier of the Memcached node. The node id does not include project or location like the Memcached instance name.
	*/
	NodeId string `json:"nodeId,omitempty" yaml:"nodeId,omitempty"`

	/*
	   (Output)
	   The port number of the Memcached server on this node.
	*/
	Port int `json:"port,omitempty" yaml:"port,omitempty"`

	/*
	   (Output)
	   Current state of the Memcached node.
	*/
	State string `json:"state,omitempty" yaml:"state,omitempty"`

	/*
	   (Output)
	   Location (GCP Zone) for the Memcached node.
	*/
	Zone string `json:"zone,omitempty" yaml:"zone,omitempty"`
}
