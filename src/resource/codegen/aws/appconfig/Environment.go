package appconfig

import types "libds/aws/types"

type Environment struct {
	// AppConfig application ID. Must be between 4 and 7 characters in length.
	ApplicationId string `json:"applicationId,omitempty" yaml:"applicationId,omitempty"`

	// Description of the environment. Can be at most 1024 characters.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// Set of Amazon CloudWatch alarms to monitor during the deployment process. Maximum of 5. See Monitor below for more details.
	Monitors []types.Appconfig_EnvironmentMonitor `json:"monitors,omitempty" yaml:"monitors,omitempty"`

	// Name for the environment. Must be between 1 and 64 characters in length.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`
}
