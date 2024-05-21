package runtimeconfig

type Config struct {
	/*
	   The description to associate with the runtime
	   config.
	*/
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	/*
	   The name of the runtime config.

	   - - -
	*/
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	/*
	   The ID of the project in which the resource belongs. If it
	   is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`
}
