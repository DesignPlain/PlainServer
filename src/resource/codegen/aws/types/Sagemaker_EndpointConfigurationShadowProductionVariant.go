package types

type Sagemaker_EndpointConfigurationShadowProductionVariant struct {
	// The type of instance to start.
	InstanceType string `json:"instanceType,omitempty" yaml:"instanceType,omitempty"`

	// Sets how the endpoint routes incoming traffic. See routing_config below.
	RoutingConfigs []Sagemaker_EndpointConfigurationShadowProductionVariantRoutingConfig `json:"routingConfigs,omitempty" yaml:"routingConfigs,omitempty"`

	// You can use this parameter to turn on native Amazon Web Services Systems Manager (SSM) access for a production variant behind an endpoint. By default, SSM access is disabled for all production variants behind an endpoints.
	EnableSsmAccess bool `json:"enableSsmAccess,omitempty" yaml:"enableSsmAccess,omitempty"`

	// The size, in GB, of the ML storage volume attached to individual inference instance associated with the production variant. Valid values between `1` and `512`.
	VolumeSizeInGb int `json:"volumeSizeInGb,omitempty" yaml:"volumeSizeInGb,omitempty"`

	// Specifies configuration for how an endpoint performs asynchronous inference.
	ServerlessConfig Sagemaker_EndpointConfigurationShadowProductionVariantServerlessConfig `json:"serverlessConfig,omitempty" yaml:"serverlessConfig,omitempty"`

	// Specifies configuration for a core dump from the model container when the process crashes. Fields are documented below.
	CoreDumpConfig Sagemaker_EndpointConfigurationShadowProductionVariantCoreDumpConfig `json:"coreDumpConfig,omitempty" yaml:"coreDumpConfig,omitempty"`

	// Specifies an option from a collection of preconfigured Amazon Machine Image (AMI) images. Each image is configured by Amazon Web Services with a set of software and driver versions. Amazon Web Services optimizes these configurations for different machine learning workloads.
	InferenceAmiVersion string `json:"inferenceAmiVersion,omitempty" yaml:"inferenceAmiVersion,omitempty"`

	// Initial number of instances used for auto-scaling.
	InitialInstanceCount int `json:"initialInstanceCount,omitempty" yaml:"initialInstanceCount,omitempty"`

	// Determines initial traffic distribution among all of the models that you specify in the endpoint configuration. If unspecified, it defaults to `1.0`.
	InitialVariantWeight float64 `json:"initialVariantWeight,omitempty" yaml:"initialVariantWeight,omitempty"`

	// The name of the model to use.
	ModelName string `json:"modelName,omitempty" yaml:"modelName,omitempty"`

	// The timeout value, in seconds, for your inference container to pass health check by SageMaker Hosting. For more information about health check, see [How Your Container Should Respond to Health Check (Ping) Requests](https://docs.aws.amazon.com/sagemaker/latest/dg/your-algorithms-inference-code.html#your-algorithms-inference-algo-ping-requests). Valid values between `60` and `3600`.
	ContainerStartupHealthCheckTimeoutInSeconds int `json:"containerStartupHealthCheckTimeoutInSeconds,omitempty" yaml:"containerStartupHealthCheckTimeoutInSeconds,omitempty"`

	// The timeout value, in seconds, to download and extract the model that you want to host from Amazon S3 to the individual inference instance associated with this production variant. Valid values between `60` and `3600`.
	ModelDataDownloadTimeoutInSeconds int `json:"modelDataDownloadTimeoutInSeconds,omitempty" yaml:"modelDataDownloadTimeoutInSeconds,omitempty"`

	// The name of the variant. If omitted, this provider will assign a random, unique name.
	VariantName string `json:"variantName,omitempty" yaml:"variantName,omitempty"`

	// The size of the Elastic Inference (EI) instance to use for the production variant.
	AcceleratorType string `json:"acceleratorType,omitempty" yaml:"acceleratorType,omitempty"`
}
