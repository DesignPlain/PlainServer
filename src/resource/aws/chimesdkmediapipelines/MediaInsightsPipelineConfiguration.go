package chimesdkmediapipelines

import types "DesignSphere_Server/src/resource/aws/types"

type MediaInsightsPipelineConfiguration struct {
	// Collection of processors and sinks to transform media and deliver data.
	Elements []types.Chimesdkmediapipelines_MediaInsightsPipelineConfigurationElement `json:"elements,omitempty" yaml:"elements,omitempty"`

	// Configuration name.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Configuration for real-time alert rules to send EventBridge notifications when certain conditions are met.
	RealTimeAlertConfiguration types.Chimesdkmediapipelines_MediaInsightsPipelineConfigurationRealTimeAlertConfiguration `json:"realTimeAlertConfiguration,omitempty" yaml:"realTimeAlertConfiguration,omitempty"`

	// ARN of IAM Role used by service to invoke processors and sinks specified by configuration elements.
	ResourceAccessRoleArn string `json:"resourceAccessRoleArn,omitempty" yaml:"resourceAccessRoleArn,omitempty"`

	// Key-value map of tags for the resource.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`
}
