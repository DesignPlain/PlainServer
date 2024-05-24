package types

type Sagemaker_EndpointConfigurationShadowProductionVariantRoutingConfig struct {
	// Sets how the endpoint routes incoming traffic. Valid values are `LEAST_OUTSTANDING_REQUESTS` and `RANDOM`. `LEAST_OUTSTANDING_REQUESTS` routes requests to the specific instances that have more capacity to process them. `RANDOM` routes each request to a randomly chosen instance.
	RoutingStrategy string `json:"routingStrategy,omitempty" yaml:"routingStrategy,omitempty"`
}
