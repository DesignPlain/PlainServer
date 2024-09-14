package types

type Evidently_LaunchScheduledSplitsConfigStep struct {
	// The traffic allocation percentages among the feature variations during one step of a launch. This is a set of key-value pairs. The keys are variation names. The values represent the percentage of traffic to allocate to that variation during this step. For more information, refer to the [AWS documentation for ScheduledSplitConfig groupWeights](https://docs.aws.amazon.com/cloudwatchevidently/latest/APIReference/API_ScheduledSplitConfig.html).
	GroupWeights map[string]int `json:"groupWeights,omitempty" yaml:"groupWeights,omitempty"`

	// One or up to six blocks that specify different traffic splits for one or more audience segments. A segment is a portion of your audience that share one or more characteristics. Examples could be Chrome browser users, users in Europe, or Firefox browser users in Europe who also fit other criteria that your application collects, such as age. Detailed below.
	SegmentOverrides []Evidently_LaunchScheduledSplitsConfigStepSegmentOverride `json:"segmentOverrides,omitempty" yaml:"segmentOverrides,omitempty"`

	// Specifies the date and time that this step of the launch starts.
	StartTime string `json:"startTime,omitempty" yaml:"startTime,omitempty"`
}
