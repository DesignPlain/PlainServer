package types

type Container_ClusterGatewayApiConfig struct {
	// Which Gateway Api channel should be used. `CHANNEL_DISABLED`, `CHANNEL_EXPERIMENTAL` or `CHANNEL_STANDARD`.
	Channel string `json:"channel,omitempty" yaml:"channel,omitempty"`
}
