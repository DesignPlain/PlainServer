package lambda

type ProvisionedConcurrencyConfig struct {
	// Name or Amazon Resource Name (ARN) of the Lambda Function.
	FunctionName string `json:"functionName,omitempty" yaml:"functionName,omitempty"`

	// Amount of capacity to allocate. Must be greater than or equal to `1`.
	ProvisionedConcurrentExecutions int `json:"provisionedConcurrentExecutions,omitempty" yaml:"provisionedConcurrentExecutions,omitempty"`

	/*
	   Lambda Function version or Lambda Alias name.

	   The following arguments are optional:
	*/
	Qualifier string `json:"qualifier,omitempty" yaml:"qualifier,omitempty"`

	// Whether to retain the provisoned concurrency configuration upon destruction. Defaults to `false`. If set to `true`, the resource in simply removed from state instead.
	SkipDestroy bool `json:"skipDestroy,omitempty" yaml:"skipDestroy,omitempty"`
}
