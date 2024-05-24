package types

type Lambda_EventSourceMappingSelfManagedEventSource struct {
	// A map of endpoints for the self managed source.  For Kafka self-managed sources, the key should be `KAFKA_BOOTSTRAP_SERVERS` and the value should be a string with a comma separated list of broker endpoints.
	Endpoints map[string]string `json:"endpoints,omitempty" yaml:"endpoints,omitempty"`
}
