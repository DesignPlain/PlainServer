package types

type Cfg_RecorderRecordingModeRecordingModeOverride struct {
	// The recording frequency for the resources in the override block. `CONTINUOUS` or `DAILY`.
	RecordingFrequency string `json:"recordingFrequency,omitempty" yaml:"recordingFrequency,omitempty"`

	// A list that specifies the types of AWS resources for which the override applies to.  See [restrictions in the AWS Docs](https://docs.aws.amazon.com/config/latest/APIReference/API_RecordingModeOverride.html)
	ResourceTypes []string `json:"resourceTypes,omitempty" yaml:"resourceTypes,omitempty"`

	// A description you provide of the override.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
}
