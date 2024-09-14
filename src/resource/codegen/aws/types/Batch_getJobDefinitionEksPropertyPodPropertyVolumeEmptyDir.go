package types

type Batch_getJobDefinitionEksPropertyPodPropertyVolumeEmptyDir struct {
	// The medium to store the volume.
	Medium string `json:"medium,omitempty" yaml:"medium,omitempty"`

	// The maximum size of the volume. By default, there's no maximum size defined.
	SizeLimit string `json:"sizeLimit,omitempty" yaml:"sizeLimit,omitempty"`
}
