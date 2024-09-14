package types

type Batch_JobDefinitionEksPropertiesPodPropertiesVolumeEmptyDir struct {
	// Medium to store the volume. The default value is an empty string, which uses the storage of the node.
	Medium string `json:"medium,omitempty" yaml:"medium,omitempty"`

	// Maximum size of the volume. By default, there's no maximum size defined.
	SizeLimit string `json:"sizeLimit,omitempty" yaml:"sizeLimit,omitempty"`
}
