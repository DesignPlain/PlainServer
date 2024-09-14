package types

type Networksecurity_ClientTlsPolicyServerValidationCa struct {
	/*
	   The certificate provider instance specification that will be passed to the data plane, which will be used to load necessary credential information.
	   Structure is documented below.
	*/
	CertificateProviderInstance Networksecurity_ClientTlsPolicyServerValidationCaCertificateProviderInstance `json:"certificateProviderInstance,omitempty" yaml:"certificateProviderInstance,omitempty"`

	/*
	   gRPC specific configuration to access the gRPC server to obtain the cert and private key.
	   Structure is documented below.
	*/
	GrpcEndpoint Networksecurity_ClientTlsPolicyServerValidationCaGrpcEndpoint `json:"grpcEndpoint,omitempty" yaml:"grpcEndpoint,omitempty"`
}
