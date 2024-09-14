package types

type Appmesh_VirtualGatewaySpecLoggingAccessLogFile struct {
	// File path to write access logs to. You can use `/dev/stdout` to send access logs to standard out. Must be between 1 and 255 characters in length.
	Path string `json:"path,omitempty" yaml:"path,omitempty"`

	// The specified format for the logs.
	Format Appmesh_VirtualGatewaySpecLoggingAccessLogFileFormat `json:"format,omitempty" yaml:"format,omitempty"`
}
