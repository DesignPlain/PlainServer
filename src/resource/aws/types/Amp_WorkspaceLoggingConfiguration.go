package types

type Amp_WorkspaceLoggingConfiguration struct {
	// The ARN of the CloudWatch log group to which the vended log data will be published. This log group must exist.
	LogGroupArn string `json:"logGroupArn,omitempty" yaml:"logGroupArn,omitempty"`
}
