package types

type Eks_ClusterCertificateAuthority struct {
	// Base64 encoded certificate data required to communicate with your cluster. Add this to the `certificate-authority-data` section of the `kubeconfig` file for your cluster.
	Data string `json:"data,omitempty" yaml:"data,omitempty"`
}
