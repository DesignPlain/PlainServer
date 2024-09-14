package kinesis

type ResourcePolicy struct {
	// The policy document.
	Policy string `json:"policy,omitempty" yaml:"policy,omitempty"`

	// The Amazon Resource Name (ARN) of the data stream or consumer.
	ResourceArn string `json:"resourceArn,omitempty" yaml:"resourceArn,omitempty"`
}
