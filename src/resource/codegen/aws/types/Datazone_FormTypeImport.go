package types

type Datazone_FormTypeImport struct {
	// Name of the form type. Must be the name of the structure in smithy document.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Revision of the Form Type.
	Revision string `json:"revision,omitempty" yaml:"revision,omitempty"`
}
