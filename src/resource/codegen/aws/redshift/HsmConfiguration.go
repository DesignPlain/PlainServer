package redshift

type HsmConfiguration struct {
	// The IP address that the Amazon Redshift cluster must use to access the HSM.
	HsmIpAddress string `json:"hsmIpAddress,omitempty" yaml:"hsmIpAddress,omitempty"`

	// The name of the partition in the HSM where the Amazon Redshift clusters will store their database encryption keys.
	HsmPartitionName string `json:"hsmPartitionName,omitempty" yaml:"hsmPartitionName,omitempty"`

	// The password required to access the HSM partition.
	HsmPartitionPassword string `json:"hsmPartitionPassword,omitempty" yaml:"hsmPartitionPassword,omitempty"`

	// The HSMs public certificate file. When using Cloud HSM, the file name is server.pem.
	HsmServerPublicCertificate string `json:"hsmServerPublicCertificate,omitempty" yaml:"hsmServerPublicCertificate,omitempty"`

	// A map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// A text description of the HSM configuration to be created.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// The identifier to be assigned to the new Amazon Redshift HSM configuration.
	HsmConfigurationIdentifier string `json:"hsmConfigurationIdentifier,omitempty" yaml:"hsmConfigurationIdentifier,omitempty"`
}
