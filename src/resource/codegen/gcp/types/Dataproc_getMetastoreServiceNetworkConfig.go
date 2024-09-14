package types

type Dataproc_getMetastoreServiceNetworkConfig struct {
	// The consumer-side network configuration for the Dataproc Metastore instance.
	Consumers []Dataproc_getMetastoreServiceNetworkConfigConsumer `json:"consumers,omitempty" yaml:"consumers,omitempty"`

	// Enables custom routes to be imported and exported for the Dataproc Metastore service's peered VPC network.
	CustomRoutesEnabled bool `json:"customRoutesEnabled,omitempty" yaml:"customRoutesEnabled,omitempty"`
}
