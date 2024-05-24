package types

type Servicecatalog_ProvisionedProductOutput struct {
	// The description of the output.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// Parameter key.
	Key string `json:"key,omitempty" yaml:"key,omitempty"`

	// Parameter value.
	Value string `json:"value,omitempty" yaml:"value,omitempty"`
}
