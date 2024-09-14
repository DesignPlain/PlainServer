package types

type Appmesh_VirtualNodeSpecListenerTlsValidationTrustFile struct {
	// Certificate trust chain for a certificate stored on the file system of the mesh endpoint that the proxy is running on. Must be between 1 and 255 characters in length.
	CertificateChain string `json:"certificateChain,omitempty" yaml:"certificateChain,omitempty"`
}
