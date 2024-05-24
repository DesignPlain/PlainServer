package types

type Cfg_RecorderRecordingGroupExclusionByResourceType struct {
	// A list that specifies the types of AWS resources for which AWS Config excludes records configuration changes. See [relevant part of AWS Docs](http://docs.aws.amazon.com/config/latest/APIReference/API_ResourceIdentifier.html#config-Type-ResourceIdentifier-resourceType) for available
	ResourceTypes []string `json:"resourceTypes,omitempty" yaml:"resourceTypes,omitempty"`
}
