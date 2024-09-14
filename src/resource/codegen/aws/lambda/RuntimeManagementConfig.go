package lambda

type RuntimeManagementConfig struct {
	// Runtime update mode. Valid values are `Auto`, `FunctionUpdate`, and `Manual`. When a function is created, the default mode is `Auto`.
	UpdateRuntimeOn string `json:"updateRuntimeOn,omitempty" yaml:"updateRuntimeOn,omitempty"`

	/*
	   Name or ARN of the Lambda function.

	   The following arguments are optional:
	*/
	FunctionName string `json:"functionName,omitempty" yaml:"functionName,omitempty"`

	// Version of the function. This can be `$LATEST` or a published version number. If omitted, this resource will manage the runtime configuration for `$LATEST`.
	Qualifier string `json:"qualifier,omitempty" yaml:"qualifier,omitempty"`

	// ARN of the runtime version. Only required when `update_runtime_on` is `Manual`.
	RuntimeVersionArn string `json:"runtimeVersionArn,omitempty" yaml:"runtimeVersionArn,omitempty"`
}
