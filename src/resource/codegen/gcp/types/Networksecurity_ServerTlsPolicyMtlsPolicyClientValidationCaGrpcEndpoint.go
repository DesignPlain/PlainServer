package types

type Networksecurity_ServerTlsPolicyMtlsPolicyClientValidationCaGrpcEndpoint struct {
	// The target URI of the gRPC endpoint. Only UDS path is supported, and should start with "unix:".
	TargetUri string `json:"targetUri,omitempty" yaml:"targetUri,omitempty"`
}
