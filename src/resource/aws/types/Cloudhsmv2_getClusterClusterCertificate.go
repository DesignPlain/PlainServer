package types

type Cloudhsmv2_getClusterClusterCertificate struct {
	//
	AwsHardwareCertificate string `json:"awsHardwareCertificate,omitempty" yaml:"awsHardwareCertificate,omitempty"`

	//
	ClusterCertificate string `json:"clusterCertificate,omitempty" yaml:"clusterCertificate,omitempty"`

	//
	ClusterCsr string `json:"clusterCsr,omitempty" yaml:"clusterCsr,omitempty"`

	//
	HsmCertificate string `json:"hsmCertificate,omitempty" yaml:"hsmCertificate,omitempty"`

	//
	ManufacturerHardwareCertificate string `json:"manufacturerHardwareCertificate,omitempty" yaml:"manufacturerHardwareCertificate,omitempty"`
}
