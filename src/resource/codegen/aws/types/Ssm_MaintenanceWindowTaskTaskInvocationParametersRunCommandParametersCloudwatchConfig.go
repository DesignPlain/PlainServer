package types

type Ssm_MaintenanceWindowTaskTaskInvocationParametersRunCommandParametersCloudwatchConfig struct {
	// The name of the CloudWatch log group where you want to send command output. If you don't specify a group name, Systems Manager automatically creates a log group for you. The log group uses the following naming format: aws/ssm/SystemsManagerDocumentName.
	CloudwatchLogGroupName string `json:"cloudwatchLogGroupName,omitempty" yaml:"cloudwatchLogGroupName,omitempty"`

	// Enables Systems Manager to send command output to CloudWatch Logs.
	CloudwatchOutputEnabled bool `json:"cloudwatchOutputEnabled,omitempty" yaml:"cloudwatchOutputEnabled,omitempty"`
}
