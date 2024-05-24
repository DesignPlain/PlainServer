package types

type Appmesh_VirtualNodeSpecLoggingAccessLog struct {
	// File object to send virtual node access logs to.
	File Appmesh_VirtualNodeSpecLoggingAccessLogFile `json:"file,omitempty" yaml:"file,omitempty"`
}
