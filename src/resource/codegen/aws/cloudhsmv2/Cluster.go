package cloudhsmv2

type Cluster struct {
	// The type of HSM module in the cluster. Currently, only `hsm1.medium` is supported.
	HsmType string `json:"hsmType,omitempty" yaml:"hsmType,omitempty"`

	// ID of Cloud HSM v2 cluster backup to be restored.
	SourceBackupIdentifier string `json:"sourceBackupIdentifier,omitempty" yaml:"sourceBackupIdentifier,omitempty"`

	// The IDs of subnets in which cluster will operate.
	SubnetIds []string `json:"subnetIds,omitempty" yaml:"subnetIds,omitempty"`

	// A map of tags to assign to the resource. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`
}
