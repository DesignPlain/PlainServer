package types

type Appmesh_VirtualNodeSpecBackendDefaultsClientPolicyTlsCertificateFile struct {
	// Certificate trust chain for a certificate stored on the file system of the mesh endpoint that the proxy is running on. Must be between 1 and 255 characters in length.
	CertificateChain string `json:"certificateChain,omitempty" yaml:"certificateChain,omitempty"`

	// Private key for a certificate stored on the file system of the virtual node that the proxy is running on. Must be between 1 and 255 characters in length.
	PrivateKey string `json:"privateKey,omitempty" yaml:"privateKey,omitempty"`
}
