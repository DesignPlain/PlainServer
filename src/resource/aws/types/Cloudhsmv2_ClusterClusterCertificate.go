package types

type Cloudhsmv2_ClusterClusterCertificate struct {
	//
	ClusterCertificate string `json:"clusterCertificate,omitempty" yaml:"clusterCertificate,omitempty"`

	//
	ClusterCsr string `json:"clusterCsr,omitempty" yaml:"clusterCsr,omitempty"`

	//
	HsmCertificate string `json:"hsmCertificate,omitempty" yaml:"hsmCertificate,omitempty"`

	//
	ManufacturerHardwareCertificate string `json:"manufacturerHardwareCertificate,omitempty" yaml:"manufacturerHardwareCertificate,omitempty"`

	//
	AwsHardwareCertificate string `json:"awsHardwareCertificate,omitempty" yaml:"awsHardwareCertificate,omitempty"`
}
