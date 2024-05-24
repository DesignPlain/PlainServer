package types

type Batch_JobDefinitionEksPropertiesPodPropertiesVolumeEmptyDir struct {
	// The maximum size of the volume. By default, there's no maximum size defined.
	SizeLimit string `json:"sizeLimit,omitempty" yaml:"sizeLimit,omitempty"`

	// The medium to store the volume. The default value is an empty string, which uses the storage of the node.
	Medium string `json:"medium,omitempty" yaml:"medium,omitempty"`
}
