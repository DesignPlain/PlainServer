package types

type Networksecurity_ServerTlsPolicyMtlsPolicyClientValidationCa struct {
	/*
	   gRPC specific configuration to access the gRPC server to obtain the cert and private key.
	   Structure is documented below.
	*/
	GrpcEndpoint Networksecurity_ServerTlsPolicyMtlsPolicyClientValidationCaGrpcEndpoint `json:"grpcEndpoint,omitempty" yaml:"grpcEndpoint,omitempty"`

	/*
	   Optional if policy is to be used with Traffic Director. For external HTTPS load balancer must be empty.
	   Defines a mechanism to provision server identity (public and private keys). Cannot be combined with allowOpen as a permissive mode that allows both plain text and TLS is not supported.
	   Structure is documented below.
	*/
	CertificateProviderInstance Networksecurity_ServerTlsPolicyMtlsPolicyClientValidationCaCertificateProviderInstance `json:"certificateProviderInstance,omitempty" yaml:"certificateProviderInstance,omitempty"`
}
