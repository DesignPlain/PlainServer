package types

type Msk_ClusterConfigurationInfo struct {
	// Amazon Resource Name (ARN) of the MSK Configuration to use in the cluster.
	Arn string `json:"arn,omitempty" yaml:"arn,omitempty"`

	// Revision of the MSK Configuration to use in the cluster.
	Revision int `json:"revision,omitempty" yaml:"revision,omitempty"`
}
