package types

type Appmesh_VirtualGatewaySpecLogging struct {
	// Access log configuration for a virtual gateway.
	AccessLog Appmesh_VirtualGatewaySpecLoggingAccessLog `json:"accessLog,omitempty" yaml:"accessLog,omitempty"`
}
