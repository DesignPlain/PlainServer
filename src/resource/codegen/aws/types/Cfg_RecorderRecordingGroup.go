package types

type Cfg_RecorderRecordingGroup struct {
	// Specifies whether AWS Config records configuration changes for every supported type of regional resource (which includes any new type that will become supported in the future). Conflicts with `resource_types`. Defaults to `true`.
	AllSupported bool `json:"allSupported,omitempty" yaml:"allSupported,omitempty"`

	// An object that specifies how AWS Config excludes resource types from being recorded by the configuration recorder.To use this option, you must set the useOnly field of RecordingStrategy to `EXCLUSION_BY_RESOURCE_TYPES` Requires `all_supported = false`. Conflicts with `resource_types`.
	ExclusionByResourceTypes []Cfg_RecorderRecordingGroupExclusionByResourceType `json:"exclusionByResourceTypes,omitempty" yaml:"exclusionByResourceTypes,omitempty"`

	// Specifies whether AWS Config includes all supported types of _global resources_ with the resources that it records. Requires `all_supported = true`. Conflicts with `resource_types`.
	IncludeGlobalResourceTypes bool `json:"includeGlobalResourceTypes,omitempty" yaml:"includeGlobalResourceTypes,omitempty"`

	// Recording Strategy. Detailed below.
	RecordingStrategies []Cfg_RecorderRecordingGroupRecordingStrategy `json:"recordingStrategies,omitempty" yaml:"recordingStrategies,omitempty"`

	// A list that specifies the types of AWS resources for which AWS Config records configuration changes (for example, `AWS::EC2::Instance` or `AWS::CloudTrail::Trail`). See [relevant part of AWS Docs](http://docs.aws.amazon.com/config/latest/APIReference/API_ResourceIdentifier.html#config-Type-ResourceIdentifier-resourceType) for available types. In order to use this attribute, `all_supported` must be set to false.
	ResourceTypes []string `json:"resourceTypes,omitempty" yaml:"resourceTypes,omitempty"`
}
