package types

type Integrationconnectors_ConnectionConnectorVersionInfraConfig struct {
	/*
	   (Output)
	   Max QPS supported by the connector version before throttling of requests.
	*/
	RatelimitThreshold string `json:"ratelimitThreshold,omitempty" yaml:"ratelimitThreshold,omitempty"`
}
