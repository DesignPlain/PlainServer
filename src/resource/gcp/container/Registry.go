package container

type Registry struct {
	// The location of the registry. One of `ASIA`, `EU`, `US` or not specified. See [the official documentation](https://cloud.google.com/container-registry/docs/pushing-and-pulling#pushing_an_image_to_a_registry) for more information on registry locations.
	Location string `json:"location,omitempty" yaml:"location,omitempty"`

	// The ID of the project in which the resource belongs. If it is not provided, the provider project is used.
	Project string `json:"project,omitempty" yaml:"project,omitempty"`
}
