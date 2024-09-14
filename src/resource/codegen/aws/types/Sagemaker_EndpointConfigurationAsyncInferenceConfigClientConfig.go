package types

type Sagemaker_EndpointConfigurationAsyncInferenceConfigClientConfig struct {
	// The maximum number of concurrent requests sent by the SageMaker client to the model container. If no value is provided, Amazon SageMaker will choose an optimal value for you.
	MaxConcurrentInvocationsPerInstance int `json:"maxConcurrentInvocationsPerInstance,omitempty" yaml:"maxConcurrentInvocationsPerInstance,omitempty"`
}
