package types

type Sfn_StateMachineTracingConfiguration struct {
	// When set to `true`, AWS X-Ray tracing is enabled. Make sure the State Machine has the correct IAM policies for logging. See the [AWS Step Functions Developer Guide](https://docs.aws.amazon.com/step-functions/latest/dg/xray-iam.html) for details.
	Enabled bool `json:"enabled,omitempty" yaml:"enabled,omitempty"`
}
