package types

type Batch_ComputeEnvironmentComputeResourcesLaunchTemplate struct {
	// ID of the launch template. You must specify either the launch template ID or launch template name in the request, but not both.
	LaunchTemplateId string `json:"launchTemplateId,omitempty" yaml:"launchTemplateId,omitempty"`

	// Name of the launch template.
	LaunchTemplateName string `json:"launchTemplateName,omitempty" yaml:"launchTemplateName,omitempty"`

	// The version number of the launch template. Default: The default version of the launch template.
	Version string `json:"version,omitempty" yaml:"version,omitempty"`
}
