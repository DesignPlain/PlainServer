package types

type Evidently_LaunchScheduledSplitsConfig struct {
	// One or up to six blocks that define the traffic allocation percentages among the feature variations during each step of the launch. This also defines the start time of each step. Detailed below.
	Steps []Evidently_LaunchScheduledSplitsConfigStep `json:"steps,omitempty" yaml:"steps,omitempty"`
}
