package types

type Lambda_FunctionLoggingConfig struct {
	// for JSON structured logs, choose the detail level of the logs your application sends to CloudWatch when using supported logging libraries.
	ApplicationLogLevel string `json:"applicationLogLevel,omitempty" yaml:"applicationLogLevel,omitempty"`

	// select between `Text` and structured `JSON` format for your function's logs.
	LogFormat string `json:"logFormat,omitempty" yaml:"logFormat,omitempty"`

	// the CloudWatch log group your function sends logs to.
	LogGroup string `json:"logGroup,omitempty" yaml:"logGroup,omitempty"`

	// for JSON structured logs, choose the detail level of the Lambda platform event logs sent to CloudWatch, such as `ERROR`, `DEBUG`, or `INFO`.
	SystemLogLevel string `json:"systemLogLevel,omitempty" yaml:"systemLogLevel,omitempty"`
}
