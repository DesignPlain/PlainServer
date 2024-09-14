package types

type Lambda_FunctionEventInvokeConfigDestinationConfigOnSuccess struct {
	// Amazon Resource Name (ARN) of the destination resource. See the [Lambda Developer Guide](https://docs.aws.amazon.com/lambda/latest/dg/invocation-async.html#invocation-async-destinations) for acceptable resource types and associated IAM permissions.
	Destination string `json:"destination,omitempty" yaml:"destination,omitempty"`
}
