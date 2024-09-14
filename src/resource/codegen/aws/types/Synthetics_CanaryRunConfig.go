package types

type Synthetics_CanaryRunConfig struct {
	// Whether this canary is to use active AWS X-Ray tracing when it runs. You can enable active tracing only for canaries that use version syn-nodejs-2.0 or later for their canary runtime.
	ActiveTracing bool `json:"activeTracing,omitempty" yaml:"activeTracing,omitempty"`

	// Map of environment variables that are accessible from the canary during execution. Please see [AWS Docs](https://docs.aws.amazon.com/lambda/latest/dg/configuration-envvars.html#configuration-envvars-runtime) for variables reserved for Lambda.
	EnvironmentVariables map[string]string `json:"environmentVariables,omitempty" yaml:"environmentVariables,omitempty"`

	// Maximum amount of memory available to the canary while it is running, in MB. The value you specify must be a multiple of 64.
	MemoryInMb int `json:"memoryInMb,omitempty" yaml:"memoryInMb,omitempty"`

	// Number of seconds the canary is allowed to run before it must stop. If you omit this field, the frequency of the canary is used, up to a maximum of 840 (14 minutes).
	TimeoutInSeconds int `json:"timeoutInSeconds,omitempty" yaml:"timeoutInSeconds,omitempty"`
}
