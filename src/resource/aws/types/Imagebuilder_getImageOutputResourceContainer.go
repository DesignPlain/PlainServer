package types

type Imagebuilder_getImageOutputResourceContainer struct {
	// Set of URIs for created containers.
	ImageUris []string `json:"imageUris,omitempty" yaml:"imageUris,omitempty"`

	// Region of the container image.
	Region string `json:"region,omitempty" yaml:"region,omitempty"`
}
