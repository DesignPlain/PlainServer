package apprunner

type DefaultAutoScalingConfigurationVersion struct {
	// The ARN of the App Runner auto scaling configuration that you want to set as the default.
	AutoScalingConfigurationArn string `json:"autoScalingConfigurationArn,omitempty" yaml:"autoScalingConfigurationArn,omitempty"`
}
