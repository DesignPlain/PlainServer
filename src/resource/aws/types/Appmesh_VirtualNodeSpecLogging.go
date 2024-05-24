package types

type Appmesh_VirtualNodeSpecLogging struct {
	// Access log configuration for a virtual node.
	AccessLog Appmesh_VirtualNodeSpecLoggingAccessLog `json:"accessLog,omitempty" yaml:"accessLog,omitempty"`
}
