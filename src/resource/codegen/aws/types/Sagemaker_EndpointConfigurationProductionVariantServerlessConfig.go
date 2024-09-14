package types

type Sagemaker_EndpointConfigurationProductionVariantServerlessConfig struct {
	// The maximum number of concurrent invocations your serverless endpoint can process. Valid values are between `1` and `200`.
	MaxConcurrency int `json:"maxConcurrency,omitempty" yaml:"maxConcurrency,omitempty"`

	// The memory size of your serverless endpoint. Valid values are in 1 GB increments: `1024` MB, `2048` MB, `3072` MB, `4096` MB, `5120` MB, or `6144` MB.
	MemorySizeInMb int `json:"memorySizeInMb,omitempty" yaml:"memorySizeInMb,omitempty"`

	// The amount of provisioned concurrency to allocate for the serverless endpoint. Should be less than or equal to `max_concurrency`. Valid values are between `1` and `200`.
	ProvisionedConcurrency int `json:"provisionedConcurrency,omitempty" yaml:"provisionedConcurrency,omitempty"`
}
