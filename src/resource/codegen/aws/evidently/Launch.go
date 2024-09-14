package evidently

import types "libds/aws/types"

type Launch struct {
	// Tags to apply to the launch. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// Specifies the description of the launch.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// One or up to five blocks that contain the feature and variations that are to be used for the launch. Detailed below.
	Groups []types.Evidently_LaunchGroup `json:"groups,omitempty" yaml:"groups,omitempty"`

	// One or up to three blocks that define the metrics that will be used to monitor the launch performance. Detailed below.
	MetricMonitors []types.Evidently_LaunchMetricMonitor `json:"metricMonitors,omitempty" yaml:"metricMonitors,omitempty"`

	// The name for the new launch. Minimum length of `1`. Maximum length of `127`.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// The name or ARN of the project that is to contain the new launch.
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	// When Evidently assigns a particular user session to a launch, it must use a randomization ID to determine which variation the user session is served. This randomization ID is a combination of the entity ID and randomizationSalt. If you omit randomizationSalt, Evidently uses the launch name as the randomizationSalt.
	RandomizationSalt string `json:"randomizationSalt,omitempty" yaml:"randomizationSalt,omitempty"`

	// A block that defines the traffic allocation percentages among the feature variations during each step of the launch. Detailed below.
	ScheduledSplitsConfig types.Evidently_LaunchScheduledSplitsConfig `json:"scheduledSplitsConfig,omitempty" yaml:"scheduledSplitsConfig,omitempty"`
}
