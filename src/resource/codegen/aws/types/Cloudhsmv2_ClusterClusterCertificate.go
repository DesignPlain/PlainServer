package types

type Cloudhsmv2_ClusterClusterCertificate struct {
	// The HSM hardware certificate issued (signed) by AWS CloudHSM.
	AwsHardwareCertificate string `json:"awsHardwareCertificate,omitempty" yaml:"awsHardwareCertificate,omitempty"`

	// The cluster certificate issued (signed) by the issuing certificate authority (CA) of the cluster's owner.
	ClusterCertificate string `json:"clusterCertificate,omitempty" yaml:"clusterCertificate,omitempty"`

	// The certificate signing request (CSR). Available only in `UNINITIALIZED` state after an HSM instance is added to the cluster.
	ClusterCsr string `json:"clusterCsr,omitempty" yaml:"clusterCsr,omitempty"`

	// The HSM certificate issued (signed) by the HSM hardware.
	HsmCertificate string `json:"hsmCertificate,omitempty" yaml:"hsmCertificate,omitempty"`

	// The HSM hardware certificate issued (signed) by the hardware manufacturer.
	ManufacturerHardwareCertificate string `json:"manufacturerHardwareCertificate,omitempty" yaml:"manufacturerHardwareCertificate,omitempty"`
}
