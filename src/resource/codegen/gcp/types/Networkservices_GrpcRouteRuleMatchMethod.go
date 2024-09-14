package types

type Networkservices_GrpcRouteRuleMatchMethod struct {
	// Specifies that matches are case sensitive. The default value is true.
	CaseSensitive bool `json:"caseSensitive,omitempty" yaml:"caseSensitive,omitempty"`

	// Required. Name of the method to match against.
	GrpcMethod string `json:"grpcMethod,omitempty" yaml:"grpcMethod,omitempty"`

	// Required. Name of the service to match against.
	GrpcService string `json:"grpcService,omitempty" yaml:"grpcService,omitempty"`
}
