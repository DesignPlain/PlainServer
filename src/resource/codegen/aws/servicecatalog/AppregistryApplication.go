package servicecatalog

type AppregistryApplication struct {
	/*
	   Name of the application. The name must be unique within an AWS region.

	   The following arguments are optional:
	*/
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Description of the application.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
}
