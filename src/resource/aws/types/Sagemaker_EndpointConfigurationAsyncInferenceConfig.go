package types

type Sagemaker_EndpointConfigurationAsyncInferenceConfig struct {
	// Configures the behavior of the client used by Amazon SageMaker to interact with the model container during asynchronous inference.
	ClientConfig Sagemaker_EndpointConfigurationAsyncInferenceConfigClientConfig `json:"clientConfig,omitempty" yaml:"clientConfig,omitempty"`

	// Specifies the configuration for asynchronous inference invocation outputs.
	OutputConfig Sagemaker_EndpointConfigurationAsyncInferenceConfigOutputConfig `json:"outputConfig,omitempty" yaml:"outputConfig,omitempty"`
}
