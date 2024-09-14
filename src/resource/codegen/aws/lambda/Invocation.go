package lambda

type Invocation struct {
	// Name of the lambda function.
	FunctionName string `json:"functionName,omitempty" yaml:"functionName,omitempty"`

	/*
	   JSON payload to the lambda function.

	   The following arguments are optional:
	*/
	Input string `json:"input,omitempty" yaml:"input,omitempty"`

	// Lifecycle scope of the resource to manage. Valid values are `CREATE_ONLY` and `CRUD`. Defaults to `CREATE_ONLY`. `CREATE_ONLY` will invoke the function only on creation or replacement. `CRUD` will invoke the function on each lifecycle event, and augment the input JSON payload with additional lifecycle information.
	LifecycleScope string `json:"lifecycleScope,omitempty" yaml:"lifecycleScope,omitempty"`

	// Qualifier (i.e., version) of the lambda function. Defaults to `$LATEST`.
	Qualifier string `json:"qualifier,omitempty" yaml:"qualifier,omitempty"`

	//
	TerraformKey string `json:"terraformKey,omitempty" yaml:"terraformKey,omitempty"`

	// Map of arbitrary keys and values that, when changed, will trigger a re-invocation.
	Triggers map[string]string `json:"triggers,omitempty" yaml:"triggers,omitempty"`
}
