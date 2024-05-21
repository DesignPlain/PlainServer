package types



type Networksecurity_ServerTlsPolicyServerCertificate struct {
	/*
	   gRPC specific configuration to access the gRPC server to obtain the cert and private key.
	   Structure is documented below.
	*/
	GrpcEndpoint Networksecurity_ServerTlsPolicyServerCertificateGrpcEndpoint `json:"grpcEndpoint,omitempty" yaml:"grpcEndpoint,omitempty"`

	/*
	   Optional if policy is to be used with Traffic Director. For external HTTPS load balancer must be empty.
	   Defines a mechanism to provision server identity (public and private keys). Cannot be combined with allowOpen as a permissive mode that allows both plain text and TLS is not supported.
	   Structure is documented below.
	*/
	CertificateProviderInstance Networksecurity_ServerTlsPolicyServerCertificateCertificateProviderInstance `json:"certificateProviderInstance,omitempty" yaml:"certificateProviderInstance,omitempty"`
}
